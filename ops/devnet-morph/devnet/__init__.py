import argparse
import logging
import os
import subprocess
import json
import socket
import calendar
import datetime
import time
import fileinput
import re
import platform
import shutil
import http.client
import devnet.log_setup

# from devnet.genesis import GENESIS_TMPL

pjoin = os.path.join

parser = argparse.ArgumentParser(description='devnet launcher')
parser.add_argument('--polyrepo-dir', help='Directory of the polyrepo', default=os.getcwd())
parser.add_argument('--only-l1', help='Only bootstrap l1 geth', action="store_true")
# parser.add_argument('--deploy', help='Whether the contracts should be predeployed or deployed', action="store_true")
parser.add_argument('--mockccc', help='Whether the use mockccc way to build sequencer geth', action="store_true")
parser.add_argument('--debugccc', help='Whether set the debug log level for ccc', action="store_true")

log = logging.getLogger()

GWEI = 1e9
ETH = GWEI * GWEI


class Bunch:
    def __init__(self, **kwds):
        self.__dict__.update(kwds)


def main():
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
        env_file=pjoin(ops_dir, '.env'),
        genesis_l1_path=pjoin(devnet_dir, 'genesis-l1.json'),
        genesis_l2_path=pjoin(devnet_dir, 'genesis-l2.json'),
        rollup_config_path=pjoin(devnet_dir, 'rollup.json'),
        deployment_dir=pjoin(devnet_dir, 'devnetL1.json'),
        contracts_dir=pjoin(polyrepo_dir, 'contracts'),
        contracts_config=pjoin(contracts_dir, 'src', 'deploy-config', 'l1.ts'),
        bindings_dir=pjoin(polyrepo_dir, 'morphism-bindings')
    )

    os.makedirs(devnet_dir, exist_ok=True)
    if args.only_l1:
        devnet_l1(paths)
        return True

    # log.info(f'Building docker images')
    # devnet_build(paths)
    log.info('Devnet with upcoming smart contract deployments')
    devnet_deploy(paths, args)


def devnet_l1(paths, result=None):
    log.info('Starting L1.')
    run_command(['docker', 'compose', '-f', 'docker-compose-4nodes.yml', 'build', '--no-cache', 'l1'], check=False,
                cwd=paths.ops_dir, env={
            'PWD': paths.ops_dir
        })
    run_command(['docker', 'compose', '-f', 'docker-compose-4nodes.yml', 'up', '-d', 'l1'], check=False,
                cwd=paths.ops_dir, env={
            'PWD': paths.ops_dir
        })
    wait_up(9545)
    wait_for_rpc_server('127.0.0.1:9545')
    log.info('Sleep another 20s...')
    time.sleep(20)
    res = eth_accounts('127.0.0.1:9545')
    response = json.loads(res)
    account = response['result'][0]

    devnet_cfg_orig = pjoin(paths.deploy_config_dir, 'devnet-deploy-config.json')
    deploy_config = read_json(devnet_cfg_orig)
    for sequencer in deploy_config['l2StakingAddresses']:
        result = run_command_capture_output(
            ['cast', 'balance', sequencer, '--rpc-url', 'http://127.0.0.1:9545'])
        log.info(f"Account {sequencer}, Balance: {result.stdout}", )

        if int(result.stdout) < 5 * ETH:
            log.info(f'Insufficient Sequencer: {sequencer}, Founding with account: {account}')
            run_command([
                'cast', 'send', '--private-key', deploy_config['BLOCK_SIGNER_PRIVATE_KEY'],
                '--rpc-url', 'http://127.0.0.1:9545',
                '--value', '10ether',
                sequencer,
            ], env={}, cwd=paths.contracts_dir)


def devnet_build(paths):
    run_command(['docker', 'compose', '-f', 'docker-compose-4nodes.yml', 'build'], cwd=paths.ops_dir, env={
        'PWD': paths.ops_dir,
        'DOCKER_BUILDKIT': '1',  # (should be available by default in later versions, but explicitly enable it anyway)
        'COMPOSE_DOCKER_CLI_BUILD': '1'  # use the docker cache
    })


# Bring up the devnet where the contracts are deployed to L1
def devnet_deploy(paths, args):
    if not test_port(9545):
        devnet_l1(paths)
    done_file = pjoin(paths.devnet_dir, 'done')
    log.info('Generating network config.')
    devnet_cfg_orig = pjoin(paths.deploy_config_dir, 'devnet-deploy-config.json')
    deploy_config = read_json(devnet_cfg_orig)
    deploy_config['l1GenesisBlockTimestamp'] = "0x{:x}".format(int(time.time()))
    deploy_config['l1StartingBlockTag'] = 'earliest'
    temp_deploy_config = pjoin(paths.devnet_dir, 'deploy-config.json')
    write_json(temp_deploy_config, deploy_config)

    # private_key = '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80'
    log.info(f'Removing contracts deployment file: {paths.deployment_dir}...')
    run_command(['rm', '-f', paths.deployment_dir], env={}, cwd=paths.contracts_dir)
    log.info('Deploying L1 Proxy contracts...')
    run_command([
        'yarn', 'build'
    ], env={}, cwd=paths.contracts_dir)
    run_command([
        'npx', 'hardhat', 'deploy', '--network', 'l1', '--storagepath', paths.deployment_dir
    ], env={}, cwd=paths.contracts_dir)

    log.info('Generating L2 genesis and rollup configs.')
    run_command([
        'env', 'CGO_ENABLED=1', 'CGO_LDFLAGS="-ldl"',
        'go', 'run', 'cmd/main.go', 'genesis', 'l2',
        '--l1-rpc', 'http://localhost:9545',
        '--deploy-config', temp_deploy_config,
        '--deployment-dir', paths.deployment_dir,
        '--outfile.l2', pjoin(paths.devnet_dir, 'genesis-l2.json'),
        '--outfile.genbatchheader', pjoin(paths.devnet_dir, 'genesis-batch-header.json'),
        '--outfile.rollup', pjoin(paths.devnet_dir, 'rollup.json')
    ], cwd=paths.L2_dir)
    write_json(done_file, {})

    log.info('Deploying L1 Impl contracts and initialize contracts...')
    rollup_cfg = read_json(paths.rollup_config_path)
    l2_genesis_state_root = rollup_cfg['l2_genesis_state_root']
    withdraw_root = rollup_cfg['withdraw_root']
    genesis_batch_header = rollup_cfg['genesis_batch_header']
#     Do not need genesis root and withdraw root
#     pattern1 = re.compile("rollupGenesisStateRoot: '.*'")
#     pattern2 = re.compile("withdrawRoot: '.*'")
    pattern3 = re.compile("batchHeader: '.*'")
#     for line in fileinput.input(paths.contracts_config, inplace=True):
#         modified_line = re.sub(pattern1, f"rollupGenesisStateRoot: '{l2_genesis_state_root}'", line)
#         print(modified_line, end='')
#     for line in fileinput.input(paths.contracts_config, inplace=True):
#         modified_line = re.sub(pattern2, f"withdrawRoot: '{withdraw_root}'", line)
#         print(modified_line, end='')

    for line in fileinput.input(paths.contracts_config, inplace=True):
        modified_line = re.sub(pattern3, f"batchHeader: '{genesis_batch_header}'", line)
        print(modified_line, end='')
    run_command([
        'npx', 'hardhat', 'initialize', '--network', 'l1', '--storagepath', paths.deployment_dir
    ], env={}, cwd=paths.contracts_dir)

    # run_command([
    #     'npx', 'hardhat', 'staking', '--network', 'l1', '--storagepath', paths.deployment_dir
    # ], env={}, cwd=paths.contracts_dir)

    log.info('Parser L1 contracts...')
    addresses = {}
    deployment = read_json(paths.deployment_dir)
    for d in deployment:
        addresses[d['name']] = d['address']
    log.info('Passing L1 contracts address:', addresses)

    log.info('Do Staking Sequencer...')
    deploy_config['l2StakingAddresses']
    deploy_config['l2StakingPks']
    deploy_config['l2StakingTmKeys']
    deploy_config['l2StakingBlsKeys']
    for i in range(4):
        run_command(['cast', 'send', addresses['Proxy__L1Staking'],
                     'register(bytes32,bytes memory)',
                     deploy_config['l2StakingTmKeys'][i],
                     deploy_config['l2StakingBlsKeys'][i],
                     '--rpc-url', 'http://127.0.0.1:9545',
                     '--value', '1ether',
                     '--private-key', deploy_config['l2StakingPks'][i]
                     ])

    build_geth_target = 'l2-geth'
    if platform.system().lower() != 'darwin':
        build_geth_target = 'l2-geth-x86'
    if args.mockccc:
        build_geth_target = 'l2-geth-mockccc'

    rust_log_level = 'info'
    if args.debugccc:
        rust_log_level = 'debug'

    log.info(f'Starting modify env file:{paths.env_file}')
    env_data = {}
    with open(paths.env_file, 'r+') as envfile:
        env_content = envfile.readlines()
        for line in env_content:
            line = line.strip()
            if line and not line.startswith('#'):
                key, value = line.split('=')
                env_data[key.strip()] = value.strip()
        env_data['L1_CROSS_DOMAIN_MESSENGER'] = addresses['Proxy__L1CrossDomainMessenger']
        env_data['MORPH_PORTAL'] = addresses['Proxy__L1MessageQueueWithGasPriceOracle']
        env_data['MORPH_ROLLUP'] = addresses['Proxy__Rollup']
        env_data['BUILD_GETH'] = build_geth_target
        env_data['RUST_LOG'] = rust_log_level
        env_data['Proxy__L1Staking'] = addresses['Proxy__L1Staking']
        envfile.seek(0)
        for key, value in env_data.items():
            envfile.write(f'{key}={value}\n')
        envfile.truncate()
        envfile.close()

    log.info('Bringing up L2.')



    run_command(['docker', 'compose', '-f', 'docker-compose-4nodes.yml', 'up',
                 '-d'], check=False, cwd=paths.ops_dir,
                env={
                    'MORPH_PORTAL': addresses['Proxy__L1MessageQueueWithGasPriceOracle'],
                    'MORPH_ROLLUP': addresses['Proxy__Rollup'],
                    'MORPH_L1STAKING': addresses['Proxy__L1Staking'],
                    'PWD': paths.ops_dir,
                    'NODE_DATA_DIR': '/data',
                    'GETH_DATA_DIR': '/db',
                    'GENESIS_FILE_PATH': '/genesis.json',
                    'L1_ETH_RPC': 'http://l1:8545',
                    'L1_BEACON_CHAIN_RPC': 'http://beacon-chain:3500',
                    'BUILD_GETH': build_geth_target,  
                })
    wait_up(8545)
    wait_for_rpc_server('127.0.0.1:8545')


def wait_for_rpc_server(url):
    log.info(f'Waiting for RPC server at {url}')

    conn = http.client.HTTPConnection(url)
    headers = {'Content-type': 'application/json'}
    body = '{"id":1, "jsonrpc":"2.0", "method": "eth_chainId", "params":[]}'

    while True:
        try:
            conn.request('POST', '/', body, headers)
            response = conn.getresponse()
            conn.close()
            if response.status < 300:
                log.info(f'RPC server at {url} ready')
                return
        except Exception as e:
            log.info(f'Waiting for RPC server at {url}')
            time.sleep(1)


def run_command(args, check=True, shell=False, cwd=None, env=None, output=None):
    env = env if env else {}
    return subprocess.run(
        args,
        check=check,
        shell=shell,
        env={
            **os.environ,
            **env
        },
        cwd=cwd,
        stdout=subprocess.PIPE if output else None,
        stderr=subprocess.PIPE if output else None,
        text=True
    )


def run_command_capture_output(args, check=True, shell=False, cwd=None, env=None):
    env = env if env else {}
    return subprocess.run(
        args,
        check=check,
        shell=shell,
        env={
            **os.environ,
            **env
        },
        cwd=cwd,
        capture_output=True,
        text=True
    )


def wait_up(port, retries=10, wait_secs=1):
    for i in range(0, retries):
        log.info(f'Trying 127.0.0.1:{port}')
        s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        try:
            s.connect(('127.0.0.1', int(port)))
            s.shutdown(2)
            log.info(f'Connected 127.0.0.1:{port}')
            return True
        except Exception:
            time.sleep(wait_secs)

    raise Exception(f'Timed out waiting for port {port}.')


def test_port(port):
    log.info(f'Testing 127.0.0.1:{port}')
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    try:
        s.connect(('127.0.0.1', int(port)))
        s.shutdown(2)
        log.info(f'Connected 127.0.0.1:{port}')
        return True
    except Exception:
        return False


def write_json(path, data):
    with open(path, 'w+') as f:
        json.dump(data, f, indent='  ')


def read_json(path):
    with open(path, 'r') as f:
        return json.load(f)


def eth_accounts(url):
    log.info(f'Fetch eth_accounts {url}')
    conn = http.client.HTTPConnection(url)
    headers = {'Content-type': 'application/json'}
    body = '{"id":2, "jsonrpc":"2.0", "method": "eth_accounts", "params":[]}'
    conn.request('POST', '/', body, headers)
    response = conn.getresponse()
    data = response.read().decode()
    conn.close()
    return data
