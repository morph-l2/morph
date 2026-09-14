import argparse
import logging
import os
import subprocess
import json
import socket
import hashlib
import fcntl
import time
import re
import http.client
import devnet.log_setup
from devnet.setup_nodes import setup_devnet_nodes

pjoin = os.path.join

parser = argparse.ArgumentParser(description='devnet launcher')
parser.add_argument('--polyrepo-dir', help='Directory of the polyrepo', default=os.getcwd())
parser.add_argument('--only-l1', help='Only bootstrap l1 geth', action="store_true")
parser.add_argument('--legacy-l1-deployment-file', default=os.environ.get('LEGACY_L1_DEPLOYMENT_FILE'),
                    help='Existing L1 deployment records containing the approved Proxy__L1Staking address')
parser.add_argument('--service-action', choices=('start', 'stop', 'rebuild'),
                    help='Manage only the tx-submitter service of a completed devnet')
parser.add_argument('--execution-client', choices=('geth', 'reth'), default='geth',
                    help='L2 execution client implementation to run')
parser.add_argument('--sequencer-private-key',
                    default=os.environ.get(
                        'SEQUENCER_PRIVATE_KEY',
                        '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80',
                    ),
                    help='Private key used by the single devnet sequencer')
parser.add_argument('--sequencer-address',
                    default=os.environ.get(
                        'HA_SEQUENCER_ADDR',
                        '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
                    ),
                    help='L1Sequencer address expected to match --sequencer-private-key')
parser.add_argument('--sequencer-upgrade-offset-seconds', type=int,
                    default=os.environ.get('SEQUENCER_UPGRADE_OFFSET_SECONDS', '0'),
                    help='Seconds from now before single-sequencer mode activates')
parser.add_argument('--deployer-private-key',
                    default=os.environ.get(
                        'DEPLOYER_PRIVATE_KEY',
                        '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80',
                    ),
                    help='Private key for the L1 contract deployer/owner')
parser.add_argument('--batch-submitter-private-key',
                    default=os.environ.get(
                        'BATCH_SUBMITTER_PRIVATE_KEY',
                        '0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d',
                    ),
                    help='Private key for the devnet batch submitter; kept separate from the block sequencer')
parser.add_argument('--batch-block-interval', type=int,
                    default=os.environ.get('TX_SUBMITTER_BATCH_BLOCK_INTERVAL', '200'),
                    help='Explicit batch block-count threshold; service actions reuse the saved value')
parser.add_argument('--batch-timeout', type=int,
                    default=os.environ.get('TX_SUBMITTER_BATCH_TIMEOUT', '600'),
                    help='Explicit batch timestamp-difference threshold in seconds; service actions reuse the saved value')
parser.add_argument('--gas-oracle-private-key', default=os.environ.get('L2_GAS_ORACLE_PRIVATE_KEY', ''),
                    help='L2 gas oracle owner key; defaults to the deployer key when the address matches')
parser.add_argument('--cluster', action="store_true",
                    default=os.environ.get('DEVNET_CLUSTER', '').lower() in ('1', 'true', 'yes'),
                    help='Start an HA sequencer cluster instead of making node-0 the sequencer')
parser.add_argument('--debugccc', help='Whether set the debug log level for ccc', action="store_true")

log = logging.getLogger()


def compose_file_args(execution_client, cluster=False):
    """Return docker-compose -f flags for the chosen L2 execution client."""
    args = ['-f', 'docker-compose-devnet.yml']
    # The cluster topology comes before the execution-client override so that
    # the reth files get the last word on the ha-el-* image, entrypoint and
    # command. Later -f files win, so reversing these two leaves the cluster
    # nodes on geth even when reth was requested.
    if cluster:
        args.extend(['-f', 'docker-compose-cluster.yml'])
    if execution_client == 'reth':
        args.extend(['-f', 'docker-compose-reth.yml'])
        # The ha-el-* reth overrides sit in their own file: compose starts any
        # service a later -f file introduces, so parking them in
        # docker-compose-reth.yml started them in the non-cluster devnet too,
        # without the mounts that only the cluster file supplies.
        if cluster:
            args.extend(['-f', 'docker-compose-cluster-reth.yml'])
    return args


class Bunch:
    """Lightweight attribute container constructed from keyword arguments."""

    def __init__(self, **kwds):
        """Store all keyword arguments as attributes on the instance."""
        self.__dict__.update(kwds)


def main():
    """Entry point: parse CLI arguments and bring up the L1-only or full devnet."""
    args = parser.parse_args()

    polyrepo_dir = os.path.abspath(args.polyrepo_dir)
    L2_dir = pjoin(polyrepo_dir, 'ops', 'l2-genesis')
    devnet_dir = pjoin(polyrepo_dir, 'ops', 'l2-genesis', '.devnet')
    ops_dir = pjoin(polyrepo_dir, 'ops', 'docker')
    contracts_dir = pjoin(polyrepo_dir, 'contracts')

    paths = Bunch(
        polyrepo_dir=polyrepo_dir,
        devnet_dir=devnet_dir,
        L2_dir=L2_dir,
        deploy_config_dir=pjoin(L2_dir, 'deploy-config'),
        ops_dir=ops_dir,
        env_file=pjoin(devnet_dir, 'runtime.env'),
        genesis_l2_path=pjoin(devnet_dir, 'genesis-l2.json'),
        rollup_config_path=pjoin(devnet_dir, 'rollup.json'),
        deployment_dir=pjoin(devnet_dir, 'devnetL1.json'),
        contracts_dir=pjoin(polyrepo_dir, 'contracts'),
        contracts_config=pjoin(contracts_dir, 'src', 'deploy-config', 'l1.ts'),
    )

    if args.service_action:
        if args.only_l1:
            parser.error('--service-action cannot be combined with --only-l1')
        return devnet_service_action(paths, args)

    os.makedirs(devnet_dir, exist_ok=True)
    if args.only_l1:
        devnet_l1(paths)
        return True

    log.info('Devnet with upcoming smart contract deployments')
    devnet_deploy(paths, args)


def devnet_l1(paths):
    """Start L1 without replacing missing genesis files for an existing chain."""
    log.info('Starting L1.')
    
    layer1_dir = pjoin(paths.ops_dir, 'layer1')
    genesis_dir = pjoin(layer1_dir, 'genesis')
    jwt_dir = pjoin(layer1_dir, 'jwt')
    
    # Check if genesis files exist, if not generate them
    genesis_json = pjoin(genesis_dir, 'genesis.json')
    genesis_ssz = pjoin(genesis_dir, 'genesis.ssz')
    jwt_secret = pjoin(jwt_dir, 'jwtsecret')
    validator_definitions = pjoin(layer1_dir, 'keystores', 'layer1', 'keys', 'validator_definitions.yml')

    required_genesis = (genesis_json, genesis_ssz, pjoin(genesis_dir, 'config.yaml'),
                        pjoin(genesis_dir, 'deposit_contract_block.txt'))
    required_files = (*required_genesis, jwt_secret, validator_definitions)
    def complete():
        return all(os.path.isfile(path) and os.path.getsize(path) > 0 for path in required_files)
    if not complete():
        if any(os.path.exists(path) for path in required_genesis):
            raise RuntimeError('L1 genesis/JWT files are incomplete; preserve the existing files and restore the missing files before retrying')
        if not os.path.isfile(validator_definitions) or os.path.getsize(validator_definitions) == 0:
            raise RuntimeError('L1 validator definitions are missing or empty; restore the checked-in validator inputs')
        volumes = run_command_capture_output([
            'docker', 'volume', 'ls', '-q', '--filter', 'label=com.docker.compose.project=docker',
        ]).stdout.splitlines()
        if any('layer1-' in name for name in volumes):
            raise RuntimeError('L1 data volumes already exist without genesis files; restore the original genesis files before retrying')
        log.info('Genesis files not found, generating...')
        generate_script = pjoin(layer1_dir, 'scripts', 'generate-genesis.sh')
        if os.path.exists(generate_script):
            run_command(['bash', generate_script], check=True, cwd=layer1_dir,
                        env={'COMPOSE_PROJECT_NAME': 'docker'})
        else:
            log.error(f'Genesis generation script not found at {generate_script}')
            raise FileNotFoundError(f'Genesis generation script not found')
    
    if not complete():
        raise RuntimeError('L1 generation did not produce all required genesis/JWT files')

    # Start layer1 services
    log.info('Starting layer1 services (layer1-el, layer1-cl, layer1-vc)...')
    compose_env = {
        'NODE_DATA_DIR': '/data', 'JWT_SECRET_PATH': '/jwt-secret.txt',
        'RUST_LOG': 'info', 'BATCH_BLOCK_INTERVAL': '0',
        'BATCH_TIMEOUT': '0', 'L1_SEQUENCER_CONTRACT': '', 'MORPH_SUBMITTER': '',
        'L1_ETH_RPC': 'http://layer1-el:8545', 'L1_BEACON_CHAIN_RPC': 'http://layer1-cl:4000',
        'ACTIVE_SEQUENCER_PRIVATE_KEY': '', 'BATCH_SUBMITTER_PRIVATE_KEY': '',
    }
    run_command(['docker', 'compose', '--project-name', 'docker', '--env-file', os.devnull, '-f', 'docker-compose-devnet.yml',
                 'up', '-d', 'layer1-el', 'layer1-cl', 'layer1-vc'],
                cwd=paths.ops_dir, env=compose_env)
    
    # Wait for EL node to be ready
    log.info('Waiting for layer1-el to be ready...')
    wait_up(9545, retries=60, wait_secs=2)
    wait_for_rpc_server('127.0.0.1:9545')
    
    # Wait for first block to be mined
    log.info('Waiting for first block to be mined...')
    max_retries = 60
    retry_count = 0
    while retry_count < max_retries:
        block_number = eth_blockNumber('127.0.0.1:9545')
        if block_number is not None and block_number >= 1:
            log.info(f'First block mined! Current block number: {block_number}')
            break
        retry_count += 1
        log.info(f'Waiting for first block (current: {block_number if block_number is not None else "N/A"})...')
        time.sleep(3)
    else:
        raise RuntimeError('Timeout waiting for the first L1 block')


ADDRESS_RE = re.compile(r"0x[0-9a-fA-F]{40}\Z")
DEPLOYMENT_PHASES = ('prepared', 'deployed', 'genesis', 'initialized', 'registered', 'complete')
REQUIRED_DEPLOYMENTS = (
    'Proxy__L1Staking',
    'Proxy__L1CrossDomainMessenger', 'Proxy__L1MessageQueueWithGasPriceOracle',
    'Proxy__Rollup', 'Proxy__Submitter', 'Proxy__L1Sequencer',
)


def public_config(config):
    """Exclude legacy private-key fields from persisted configuration."""
    return {key: value for key, value in config.items()
            if 'private' not in key.lower() and not key.lower().endswith('pks')}


def config_digest(config):
    return hashlib.sha256(json.dumps(config, sort_keys=True, separators=(',', ':')).encode()).hexdigest()


def file_digest(path):
    with open(path, 'rb') as source:
        return hashlib.sha256(source.read()).hexdigest()


def require_address(value, label):
    if not isinstance(value, str) or not ADDRESS_RE.fullmatch(value) or int(value[2:], 16) == 0:
        raise RuntimeError(f'{label} must be a nonzero Ethereum address')
    return value.lower()


def legacy_l1_staking_record(path):
    if not path or not os.path.isfile(path):
        raise RuntimeError('A confirmed existing Proxy__L1Staking record is required by centralization-cleanup-spec section 9.2; provide --legacy-l1-deployment-file or restore the original deployment records. New L1 deployment does not create this contract')
    records = read_json(path)
    if not isinstance(records, list):
        raise RuntimeError('Legacy L1 deployment records must be an array')
    selected = [row for row in records if isinstance(row, dict) and row.get('name') == 'Proxy__L1Staking']
    if len(selected) != 1:
        raise RuntimeError('Exactly one confirmed Proxy__L1Staking record is required; preserve the existing deployment records')
    record = selected[0]
    address = require_address(record.get('address'), 'Proxy__L1Staking')
    number = record.get('number')
    if ('pending' in record and record['pending'] is not False) or isinstance(number, bool) or not isinstance(number, int) or not 0 <= number <= 2**53 - 1:
        raise RuntimeError('Proxy__L1Staking requires a deployment block number between 0 and 9007199254740991, with pending absent or exactly false')
    if int(address, 16) == 0xdead or any(isinstance(row, dict) and row.get('name') == 'Proxy__Submitter'
                                       and str(row.get('address', '')).lower() == address for row in records):
        raise RuntimeError('Proxy__L1Staking cannot use a placeholder or the Submitter address; restore the approved legacy contract record')
    return {'address': address, 'number': number}


def requested_legacy_l1_staking(paths, args, source_config):
    source = args.legacy_l1_deployment_file or paths.deployment_dir
    record = legacy_l1_staking_record(source)
    if os.path.isfile(paths.deployment_dir) and record != legacy_l1_staking_record(paths.deployment_dir):
        raise RuntimeError('Legacy L1Staking source differs from the existing deployment records; preserve both files and use the original source')
    configured = source_config.get('l1StakingProxy')
    if configured is not None and require_address(configured, 'l1StakingProxy') != record['address']:
        raise RuntimeError('l1StakingProxy must match the confirmed Proxy__L1Staking deployment record')
    return record


def validate_batch_parameters(block_interval, timeout):
    parameters = {'batchBlockInterval': block_interval, 'batchTimeout': timeout}
    for name, value in parameters.items():
        if isinstance(value, bool) or not isinstance(value, int) or not 0 <= value <= 2**64 - 1:
            raise RuntimeError(f'{name} must be an integer between 0 and 18446744073709551615')
    if block_interval == 0 and timeout == 0:
        raise RuntimeError('at least one explicit batch sealing trigger must be greater than zero')
    return parameters


def saved_batch_parameters(request):
    parameters = request.get('batch_parameters') if isinstance(request, dict) else None
    if not isinstance(parameters, dict) or set(parameters) != {'batchBlockInterval', 'batchTimeout'}:
        raise RuntimeError('Existing devnet has no recorded explicit batch parameters; preserve its deployment files and inspect the original runtime settings and source configuration. Automatic identity migration is not supported')
    return validate_batch_parameters(parameters['batchBlockInterval'], parameters['batchTimeout'])


def validate_parameters(paths, args, deploy_config):
    """Validate chain IDs and signer identities before creating files or sending transactions."""
    if deploy_config.get('l1ChainID') != 900 or deploy_config.get('l2ChainID') != 53077:
        raise RuntimeError('devnet requires l1ChainID=900 and l2ChainID=53077')
    if args.sequencer_upgrade_offset_seconds != 0:
        raise RuntimeError('devnet requires sequencer-upgrade-offset-seconds=0 to switch before the validator set changes')
    validate_batch_parameters(args.batch_block_interval, args.batch_timeout)
    roles = {}
    for role, key in (
        ('sequencer', args.sequencer_private_key),
        ('deployer', args.deployer_private_key),
        ('batch_submitter', args.batch_submitter_private_key),
        ('gas_oracle', args.gas_oracle_private_key or args.deployer_private_key),
    ):
        if not isinstance(key, str) or not re.fullmatch(r'(0x)?[0-9a-fA-F]{64}', key):
            raise RuntimeError(f'{role} private key must contain 32 hexadecimal bytes')
        derived = run_command_capture_output(
            ['cast', 'wallet', 'address', '--private-key', key], cwd=paths.contracts_dir,
        )
        roles[role] = require_address(derived.stdout.strip(), role)
    if roles['sequencer'] != require_address(args.sequencer_address, 'sequencer-address'):
        raise RuntimeError('sequencer private key does not match sequencer-address')
    if roles.pop('gas_oracle') != require_address(deploy_config.get('gasPriceOracleOwner'), 'gasPriceOracleOwner'):
        raise RuntimeError('gas oracle private key does not match gasPriceOracleOwner; set L2_GAS_ORACLE_PRIVATE_KEY')
    return roles


def deployment_addresses(path):
    rows = read_json(path)
    if not isinstance(rows, list):
        raise RuntimeError('deployment output must be an array')
    addresses = {}
    for row in rows:
        if not isinstance(row, dict) or not isinstance(row.get('name'), str) or not row['name']:
            raise RuntimeError('deployment records must contain a nonempty name')
        name = row['name']
        address = require_address(row.get('address'), f'deployment {name}')
        if name in addresses:
            raise RuntimeError(f'deployment output contains duplicate records for {name}')
        if row.get('pending'):
            raise RuntimeError(f'deployment {name} has not been confirmed')
        addresses[name] = address
    for name in REQUIRED_DEPLOYMENTS:
        if name not in addresses:
            raise RuntimeError(f'{name} is missing from deployment output')
    return addresses


def l1_identity():
    chain_id = rpc_result('127.0.0.1:9545', 'eth_chainId', [])
    if int(chain_id, 16) != 900:
        raise RuntimeError('L1 RPC does not report chain ID 900')
    block = rpc_result('127.0.0.1:9545', 'eth_getBlockByNumber', ['0x0', False])
    if not block or not re.fullmatch(r'0x[0-9a-fA-F]{64}', block.get('hash', '')):
        raise RuntimeError('L1 RPC did not return its genesis block hash')
    return block['hash'].lower()


def verify_l1_contracts(paths, roles, addresses):
    """Verify deployed code and the finalized sequencer and submitter configuration."""
    for name in REQUIRED_DEPLOYMENTS:
        code = rpc_result('127.0.0.1:9545', 'eth_getCode', [addresses[name], 'latest'])
        if not isinstance(code, str) or code in ('0x', '0x0'):
            raise RuntimeError(f'{name} has no deployed code on the connected L1')
    latest_l1_block = eth_blockNumber('127.0.0.1:9545')
    if latest_l1_block is None:
        raise RuntimeError('Cannot read L1 block number before checking finalized contract state')
    wait_for_l1_finalized(latest_l1_block)
    for signature, arguments, expected, contract in (
        ('getSequencerAt(uint64)(address)', ['0'], roles['sequencer'], 'Proxy__L1Sequencer'),
        ('owner()(address)', [], roles['deployer'], 'Proxy__L1Sequencer'),
        ('isActive(address)(bool)', [roles['batch_submitter']], 'true', 'Proxy__Submitter'),
    ):
        result = run_command_capture_output([
            'cast', 'call', addresses[contract], signature, *arguments,
            '--block', 'finalized', '--rpc-url', 'http://127.0.0.1:9545',
        ], cwd=paths.contracts_dir).stdout.strip().lower()
        if result != expected:
            raise RuntimeError(f'{contract}.{signature} does not match the requested devnet configuration')


def generated_artifacts(paths):
    return [paths.deployment_dir, paths.genesis_l2_path, paths.rollup_config_path,
            pjoin(paths.devnet_dir, 'genesis-batch-header.json'),
            pjoin(paths.devnet_dir, 'deployment-config.json'),
            pjoin(paths.devnet_dir, 'genesis.done'),
            pjoin(paths.devnet_dir, 'genesis-input.json'),
            pjoin(paths.devnet_dir, 'deploy-config.json')]


def runtime_environment(paths, args, deploy_config, addresses, upgrade_time):
    batch = validate_batch_parameters(args.batch_block_interval, args.batch_timeout)
    return {
        'L1_CROSS_DOMAIN_MESSENGER': addresses['Proxy__L1CrossDomainMessenger'],
        'MORPH_PORTAL': addresses['Proxy__L1MessageQueueWithGasPriceOracle'],
        'MORPH_ROLLUP': addresses['Proxy__Rollup'],
        'MORPH_SUBMITTER': addresses['Proxy__Submitter'],
        'BATCH_BLOCK_INTERVAL': str(batch['batchBlockInterval']),
        'BATCH_TIMEOUT': str(batch['batchTimeout']),
        'RUST_LOG': 'debug' if args.debugccc else 'info',
        'L1_SEQUENCER_CONTRACT': addresses['Proxy__L1Sequencer'],
        'HA_SEQUENCER_ADDR': args.sequencer_address,
        'SEQUENCER_UPGRADE_TIME': str(upgrade_time),
        'NODE_DATA_DIR': '/data',
        'JWT_SECRET_PATH': '/jwt-secret.txt',
        'TX_SUBMITTER_BATCH_V2_UPGRADE_TIME': '0',
        'MORPH_NODE_SYNC_START_HEIGHT': '1',
        'L1_ETH_RPC': 'http://layer1-el:8545',
        'L1_BEACON_CHAIN_RPC': 'http://layer1-cl:4000',
    }


def compose_runtime(paths, args, deploy_config, addresses, upgrade_time, signing_env):
    runtime = runtime_environment(paths, args, deploy_config, addresses, upgrade_time)
    for key, value in runtime.items():
        if '\n' in value or '\r' in value:
            raise RuntimeError(f'invalid newline in runtime setting {key}')
    temporary = f'{paths.env_file}.tmp'
    with open(temporary, 'w') as target:
        for key, value in runtime.items():
            target.write(f'{key}={value}\n')
    os.replace(temporary, paths.env_file)
    env = {
        **runtime,
        'SEQUENCER_PRIVATE_KEY': '', 'ACTIVE_SEQUENCER_PRIVATE_KEY': '',
        'BATCH_SUBMITTER_PRIVATE_KEY': '', 'L2_GAS_ORACLE_PRIVATE_KEY': '', **signing_env,
    }
    command = ['docker', 'compose', '--project-name', 'docker', '--env-file', paths.env_file,
               *compose_file_args(args.execution_client, args.cluster)]
    return command, env


def start_l2(paths, args, deploy_config, addresses, upgrade_time):
    command, env = compose_runtime(paths, args, deploy_config, addresses, upgrade_time, {
        'SEQUENCER_PRIVATE_KEY': args.sequencer_private_key,
        'ACTIVE_SEQUENCER_PRIVATE_KEY': '' if args.cluster else args.sequencer_private_key,
        'BATCH_SUBMITTER_PRIVATE_KEY': args.batch_submitter_private_key,
        'L2_GAS_ORACLE_PRIVATE_KEY': args.gas_oracle_private_key or args.deployer_private_key,
    })
    run_command([*command, 'config', '--quiet'], cwd=paths.ops_dir, env=env)
    run_command([*command, 'up', '-d'], cwd=paths.ops_dir, env=env)
    wait_up(8545, retries=120, wait_secs=2)
    wait_for_rpc_server('127.0.0.1:8545', expected_chain_id=53077)
    for _ in range(120):
        height = eth_blockNumber('127.0.0.1:8545')
        if height is not None and height >= 1:
            return
        time.sleep(2)
    raise RuntimeError('L2 did not produce its first block; preserve the existing deployment and inspect service logs')


def devnet_service_action(paths, args):
    """Stop the submitter, or start/rebuild it after verifying a completed deployment."""
    state_path = pjoin(paths.devnet_dir, 'deployment-state.json')
    if args.service_action == 'stop':
        os.makedirs(paths.devnet_dir, exist_ok=True)
    elif not os.path.isfile(state_path):
        raise RuntimeError('A completed devnet deployment is required before managing tx-submitter')
    with open(pjoin(paths.devnet_dir, '.deployment.lock'), 'a') as lock:
        try:
            fcntl.flock(lock, fcntl.LOCK_EX | fcntl.LOCK_NB)
        except BlockingIOError:
            raise RuntimeError('Another devnet operation is running for this output directory') from None
        if args.service_action == 'stop':
            # The same project/service identifies the submitter in every topology.
            # Stopping must remain available after an interrupted deployment.
            run_command([
                'docker', 'compose', '--project-name', 'docker', '--env-file', os.devnull,
                '-f', 'docker-compose-devnet.yml', 'stop', 'tx-submitter-0',
            ], cwd=paths.ops_dir, env={'NODE_DATA_DIR': '/data', 'JWT_SECRET_PATH': '/jwt-secret.txt'})
            return
        state = read_json(state_path)
        done_path = pjoin(paths.devnet_dir, 'done')
        if state.get('version') != 1 or state.get('phase') != 'complete' or not os.path.isfile(done_path):
            raise RuntimeError('Finish the existing devnet deployment before managing tx-submitter')
        if read_json(done_path).get('phase') != 'complete':
            raise RuntimeError('The devnet completion marker is invalid')
        source_config = public_config(read_json(pjoin(paths.deploy_config_dir, 'devnet-deploy-config.json')))
        request = state['request']
        batch = saved_batch_parameters(request)
        if request.get('legacy_l1_staking') != legacy_l1_staking_record(paths.deployment_dir):
            raise RuntimeError('The saved legacy L1Staking identity does not match the deployment records; preserve all files and restore the original deployment state')
        if request['source_config_sha256'] != config_digest(source_config) or request['l1_config_sha256'] != file_digest(paths.contracts_config):
            raise RuntimeError('Devnet source configuration changed; restore the original configuration before managing services')
        for path in generated_artifacts(paths):
            if not os.path.isfile(path) or state.get('artifact_hashes', {}).get(os.path.basename(path)) != file_digest(path):
                raise RuntimeError(f'Existing deployment artifact changed or is missing: {path}')
        roles = {name: require_address(value, name) for name, value in request['roles'].items()}
        addresses = deployment_addresses(paths.deployment_dir)
        config = read_json(pjoin(paths.devnet_dir, 'genesis-input.json'))
        service_args = Bunch(**vars(args))
        service_args.execution_client = request['execution_client']
        service_args.cluster = request['cluster']
        service_args.sequencer_address = roles['sequencer']
        service_args.batch_block_interval = batch['batchBlockInterval']
        service_args.batch_timeout = batch['batchTimeout']
        key = args.batch_submitter_private_key
        if not isinstance(key, str) or not re.fullmatch(r'(0x)?[0-9a-fA-F]{64}', key):
            raise RuntimeError('batch submitter private key must contain 32 hexadecimal bytes')
        actual = run_command_capture_output([
            'cast', 'wallet', 'address', '--private-key', key,
        ], cwd=paths.contracts_dir).stdout.strip()
        if require_address(actual, 'batch submitter') != roles['batch_submitter']:
            raise RuntimeError('batch submitter private key does not match the completed deployment')
        if l1_identity() != state['l1_genesis_hash']:
            raise RuntimeError('Connected L1 does not match the completed devnet deployment')
        verify_l1_contracts(paths, roles, addresses)
        command, env = compose_runtime(paths, service_args, config, addresses,
                                       state['sequencer_upgrade_time'], {'BATCH_SUBMITTER_PRIVATE_KEY': key})
        run_command([*command, 'config', '--quiet'], cwd=paths.ops_dir, env=env)
        operation = ['up', '-d', '--no-deps']
        if args.service_action == 'rebuild':
            operation.append('--build')
        operation.append('tx-submitter-0')
        run_command([*command, *operation], cwd=paths.ops_dir, env=env)


def devnet_deploy(paths, args):
    with open(pjoin(paths.devnet_dir, '.deployment.lock'), 'a') as lock:
        try:
            fcntl.flock(lock, fcntl.LOCK_EX | fcntl.LOCK_NB)
        except BlockingIOError:
            raise RuntimeError('Another devnet deployment is running for this output directory') from None
        return _devnet_deploy(paths, args)


def _devnet_deploy(paths, args):
    """Resume deployment stages while preserving recorded addresses and genesis files."""
    source_config = read_json(pjoin(paths.deploy_config_dir, 'devnet-deploy-config.json'))
    legacy_staking = requested_legacy_l1_staking(paths, args, source_config)
    roles = validate_parameters(paths, args, source_config)
    request = {
        'roles': roles,
        'legacy_l1_staking': legacy_staking,
        'execution_client': args.execution_client,
        'cluster': args.cluster,
        'batch_parameters': validate_batch_parameters(args.batch_block_interval, args.batch_timeout),
        'source_config_sha256': config_digest(public_config(source_config)),
        'l1_config_sha256': file_digest(paths.contracts_config),
    }
    state_path = pjoin(paths.devnet_dir, 'deployment-state.json')
    state = read_json(state_path) if os.path.isfile(state_path) else None
    if state is None:
        if set(os.listdir(paths.devnet_dir)) - {'.deployment.lock'} or os.path.exists(pjoin(paths.ops_dir, '.devnet')):
            raise RuntimeError('Existing devnet files have no deployment-state.json; preserve them and establish their deployment state before retrying')
        state = {'version': 1, 'request': request, 'phase': 'preparing'}
        write_json(state_path, state)
    else:
        saved_batch_parameters(state.get('request'))
        if state.get('version') != 1 or state.get('request') != request:
            raise RuntimeError('Existing devnet was created with different configuration; keep the existing chain and use its original parameters')
    if state.get('phase') not in ('preparing', *DEPLOYMENT_PHASES):
        raise RuntimeError('Unknown devnet deployment phase; preserve deployment-state.json for inspection')

    if not os.path.isfile(paths.deployment_dir):
        write_json(paths.deployment_dir, [{'name': 'Proxy__L1Staking', **legacy_staking}])
    phase = state['phase']
    if phase == 'preparing':
        setup_devnet_nodes(paths.polyrepo_dir, cluster=args.cluster)
        deploy_config = public_config(source_config)
        deploy_config['l1StartingBlockTag'] = 'earliest'
        deploy_config['l1StakingProxy'] = legacy_staking['address']
        write_json(pjoin(paths.devnet_dir, 'genesis-input.json'), deploy_config)
        state['deploy_config_sha256'] = config_digest(deploy_config)
        state['phase'] = phase = 'prepared'
        write_json(state_path, state)
    else:
        setup_devnet_nodes(paths.polyrepo_dir, cluster=args.cluster)
    deploy_config = read_json(pjoin(paths.devnet_dir, 'genesis-input.json'))
    if state.get('deploy_config_sha256') != config_digest(deploy_config):
        raise RuntimeError('Generated genesis-input.json changed; restore the original configuration before retrying')
    if not test_port(9545):
        devnet_l1(paths)
    identity = l1_identity()
    l1_head = rpc_result('127.0.0.1:9545', 'eth_blockNumber', [])
    if not isinstance(l1_head, str) or not re.fullmatch(r'0x(?:0|[1-9a-fA-F][0-9a-fA-F]*)', l1_head):
        raise RuntimeError('L1 RPC did not return a canonical hexadecimal block number; no deployment transaction was sent')
    if legacy_staking['number'] > int(l1_head, 16):
        raise RuntimeError('Proxy__L1Staking deployment block is newer than the connected L1 head; restore the original L1 or deployment records before retrying')
    legacy_code = rpc_result('127.0.0.1:9545', 'eth_getCode', [legacy_staking['address'], l1_head])
    if not isinstance(legacy_code, str) or not re.fullmatch(r'0x(?:[0-9a-fA-F]{2})+', legacy_code) or int(legacy_code, 16) == 0:
        raise RuntimeError('The approved Proxy__L1Staking address has no valid nonzero bytecode on the connected L1; restore the original L1 before deploying contracts')
    if 'l1_genesis_hash' in state and state['l1_genesis_hash'] != identity:
        raise RuntimeError('Connected L1 has a different genesis block; restore the original L1 before retrying')
    state['l1_genesis_hash'] = identity
    write_json(state_path, state)

    override_path = pjoin(paths.devnet_dir, 'devnet-contract-config.json')
    overrides = {'contractAdmin': roles['deployer'], 'submitterOwner': roles['deployer'],
                 'firstSequencerAddress': roles['sequencer']}
    env = {
        'L1_RPC_URL': 'http://localhost:9545',
        'DOTENV_CONFIG_PATH': os.devnull,
        'DEPLOYER_PRIVATE_KEY': args.deployer_private_key,
        'DEPLOY_CONFIG_OVERRIDE': override_path,
        'firstSequencerAddress': roles['sequencer'],
        'batchSubmitterPks': json.dumps([args.batch_submitter_private_key]),
    }
    def completed(next_phase):
        state['phase'] = next_phase
        if 'artifact_hashes' in state:
            state['artifact_hashes'] = {os.path.basename(path): file_digest(path) for path in generated_artifacts(paths)}
        write_json(state_path, state)

    write_json(override_path, overrides)
    if DEPLOYMENT_PHASES.index(phase) < DEPLOYMENT_PHASES.index('deployed'):
        if not os.path.isfile(pjoin(paths.contracts_dir, 'node_modules', '.bin', 'hardhat')):
            run_command(['yarn', 'install', '--frozen-lockfile'], env=env, cwd=paths.contracts_dir)
        run_command(['yarn', 'hardhat', 'compile'], env=env, cwd=paths.contracts_dir)
        run_command([
            'npx', 'hardhat', 'deploy', '--network', 'l1', '--storagepath', paths.deployment_dir,
        ], env=env, cwd=paths.contracts_dir)
        deployment_addresses(paths.deployment_dir)
        completed('deployed')
    genesis_command = [
            'bash', pjoin(paths.L2_dir, 'devnet-l2genesis.sh'),
            '--deployment-file', paths.deployment_dir,
            '--deploy-config', pjoin(paths.devnet_dir, 'genesis-input.json'),
            '--output-dir', paths.devnet_dir, '--l1-rpc', 'http://localhost:9545',
    ]
    if os.path.isfile(pjoin(paths.devnet_dir, 'genesis.done')):
        run_command([*genesis_command, '--verify-existing'], cwd=paths.L2_dir)
    elif DEPLOYMENT_PHASES.index(phase) >= DEPLOYMENT_PHASES.index('genesis'):
        raise RuntimeError('genesis.done is missing from an initialized deployment; restore original artifacts before retrying')
    else:
        run_command(genesis_command, cwd=paths.L2_dir)
    if DEPLOYMENT_PHASES.index(phase) < DEPLOYMENT_PHASES.index('genesis'):
        state['artifact_hashes'] = {os.path.basename(path): file_digest(path) for path in generated_artifacts(paths)}
        completed('genesis')
    else:
        # initialize appends implementation records; genesis verification checks
        # the original records. All other saved artifacts must remain unchanged.
        saved = state.get('artifact_hashes', {})
        for path in generated_artifacts(paths):
            if (path != paths.deployment_dir or DEPLOYMENT_PHASES.index(phase) >= DEPLOYMENT_PHASES.index('initialized')) and saved.get(os.path.basename(path)) != file_digest(path):
                raise RuntimeError(f'Existing deployment artifact changed: {path}; restore the original file before retrying')
    genesis_overrides = read_json(pjoin(paths.devnet_dir, 'deployment-config.json'))
    batch_header = genesis_overrides.get('batchHeader')
    if not isinstance(batch_header, str) or not re.fullmatch(r'0x(?:[0-9a-fA-F]{2})+', batch_header):
        raise RuntimeError('Generated deployment-config.json has no valid batchHeader')
    overrides['batchHeader'] = batch_header
    write_json(override_path, overrides)
    if DEPLOYMENT_PHASES.index(phase) < DEPLOYMENT_PHASES.index('initialized'):
        run_command([
            'npx', 'hardhat', 'initialize', '--network', 'l1', '--storagepath', paths.deployment_dir,
        ], env=env, cwd=paths.contracts_dir)
        completed('initialized')
    if DEPLOYMENT_PHASES.index(phase) < DEPLOYMENT_PHASES.index('registered'):
        run_command(['npx', 'hardhat', 'fund', '--network', 'l1'], env=env, cwd=paths.contracts_dir)
        run_command([
            'npx', 'hardhat', 'register', '--network', 'l1', '--storagepath', paths.deployment_dir,
        ], env=env, cwd=paths.contracts_dir)
        completed('registered')
    run_command([
        'npx', 'hardhat', 'verify-deployment', '--network', 'l1', '--storagepath', paths.deployment_dir,
        *(['--runtime'] if phase == 'complete' else []),
    ], env=env, cwd=paths.contracts_dir)
    addresses = deployment_addresses(paths.deployment_dir)
    verify_l1_contracts(paths, roles, addresses)
    if 'sequencer_upgrade_time' not in state:
        state['sequencer_upgrade_time'] = int(time.time() * 1000)
        write_json(state_path, state)
    start_l2(paths, args, deploy_config, addresses, state['sequencer_upgrade_time'])
    state['artifact_hashes'] = {os.path.basename(path): file_digest(path) for path in generated_artifacts(paths)}
    completed('complete')
    write_json(pjoin(paths.devnet_dir, 'done'), {'phase': 'complete'})


def rpc_result(url, method, params):
    conn = http.client.HTTPConnection(url, timeout=5)
    try:
        conn.request('POST', '/', json.dumps({
            'id': 1, 'jsonrpc': '2.0', 'method': method, 'params': params,
        }), {'Content-type': 'application/json'})
        response = conn.getresponse()
        if response.status != 200:
            raise RuntimeError(f'{method} returned HTTP {response.status}')
        data = json.loads(response.read())
        if 'error' in data or 'result' not in data:
            raise RuntimeError(f'{method} did not return a JSON-RPC result')
        return data['result']
    finally:
        conn.close()


def wait_for_rpc_server(url, expected_chain_id=900, retries=120, wait_secs=1):
    """Wait for the expected chain ID with a bounded number of RPC attempts."""
    for _ in range(retries):
        try:
            result = rpc_result(url, 'eth_chainId', [])
            chain_id = int(result, 16)
        except (OSError, http.client.HTTPException, ValueError, TypeError, RuntimeError):
            time.sleep(wait_secs)
            continue
        if chain_id != expected_chain_id:
            raise RuntimeError(f'RPC at {url} reports an unexpected chain ID')
        return
    raise RuntimeError(f'Timeout waiting for JSON-RPC at {url}')


def wait_for_l1_finalized(min_block, retries=120, wait_secs=3):
    """Wait until the local L1 finalized tag reaches min_block."""
    for _ in range(retries):
        finalized = eth_block_by_number('127.0.0.1:9545', 'finalized')
        if finalized is not None and finalized >= min_block:
            log.info(f'L1 finalized block {finalized} reached target {min_block}')
            return
        log.info(f'Waiting for L1 finalized block >= {min_block} (current: {finalized})')
        time.sleep(wait_secs)

    raise RuntimeError(f'Timeout waiting for L1 finalized block >= {min_block}')


def run_command(args, check=True, shell=False, cwd=None, env=None, output=None):
    """Report subprocess failures without exposing arguments that may contain private keys."""
    try:
        return subprocess.run(
            args, check=check, shell=shell, env={**os.environ, **(env or {})}, cwd=cwd,
            stdout=subprocess.PIPE if output else None,
            stderr=subprocess.PIPE if output else None, text=True,
        )
    except subprocess.CalledProcessError as error:
        program = os.path.basename(args[0]) if isinstance(args, (list, tuple)) else 'command'
        raise RuntimeError(f'{program} failed with exit code {error.returncode}; existing deployment records were preserved') from None


def run_command_capture_output(args, check=True, shell=False, cwd=None, env=None):
    return run_command(args, check=check, shell=shell, cwd=cwd, env=env, output=True)


def wait_up(port, retries=10, wait_secs=1):
    """Poll a TCP port on 127.0.0.1 until it accepts a connection or retries are exhausted."""
    for _ in range(retries):
        if test_port(port):
            return True
        time.sleep(wait_secs)

    raise RuntimeError(f'Timed out waiting for port {port}.')


def test_port(port):
    """Return True if a TCP connection to 127.0.0.1:port succeeds, False otherwise."""
    log.info(f'Testing 127.0.0.1:{port}')
    try:
        with socket.create_connection(('127.0.0.1', int(port)), timeout=1):
            log.info(f'Connected 127.0.0.1:{port}')
            return True
    except OSError:
        return False


def write_json(path, data):
    """Serialize data to path as indented JSON."""
    temporary = f'{path}.tmp'
    with open(temporary, 'w') as f:
        json.dump(data, f, indent='  ')
        f.write('\n')
    os.replace(temporary, path)


def read_json(path):
    """Load and return the JSON document stored at path."""
    with open(path, 'r') as f:
        return json.load(f)


def eth_blockNumber(url):
    """
    Call eth_blockNumber JSON-RPC method to get the current block number.
    Returns the block number as an integer, or None on error.
    """
    try:
        return int(rpc_result(url, 'eth_blockNumber', []), 16)
    except (OSError, http.client.HTTPException, ValueError, TypeError, RuntimeError) as e:
        log.debug(f'Error calling eth_blockNumber: {e}')
        return None


def eth_block_by_number(url, tag):
    """
    Call eth_getBlockByNumber JSON-RPC method for a tag and return the block number.
    Returns the block number as an integer, or None when the tag is unavailable.
    """
    try:
        block = rpc_result(url, 'eth_getBlockByNumber', [tag, False])
        if isinstance(block, dict) and block.get('number'):
            return int(block['number'], 16)
        return None
    except (OSError, http.client.HTTPException, ValueError, TypeError, RuntimeError) as e:
        log.debug(f'Error calling eth_getBlockByNumber({tag}): {e}')
        return None
