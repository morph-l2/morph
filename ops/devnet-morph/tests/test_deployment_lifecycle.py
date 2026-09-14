import importlib
import contextlib
import io
import json
import subprocess
import sys
import tempfile
import unittest
from pathlib import Path
from types import SimpleNamespace
from unittest.mock import patch


PACKAGE = Path(__file__).resolve().parents[1]
sys.path.insert(0, str(PACKAGE))
devnet = importlib.import_module('devnet')
setup_nodes = importlib.import_module('devnet.setup_nodes')
sys.path.remove(str(PACKAGE))


class DeploymentLifecycleTest(unittest.TestCase):
    def setUp(self):
        self.directory = tempfile.TemporaryDirectory()
        self.addCleanup(self.directory.cleanup)
        root = Path(self.directory.name)
        output = root / 'output'
        contracts = root / 'contracts'
        docker = root / 'ops' / 'docker'
        configs = root / 'configs'
        for directory in (output, contracts, docker, configs):
            directory.mkdir(parents=True)
        self.paths = devnet.Bunch(
            polyrepo_dir=str(root), devnet_dir=str(output), ops_dir=str(docker),
            L2_dir=str(root / 'genesis'), contracts_dir=str(contracts),
            contracts_config=str(contracts / 'l1.ts'), deploy_config_dir=str(configs),
            deployment_dir=str(output / 'devnetL1.json'),
            genesis_l2_path=str(output / 'genesis-l2.json'),
            rollup_config_path=str(output / 'rollup.json'),
            env_file=str(output / 'runtime.env'),
        )
        Path(self.paths.contracts_config).write_text('export default {batchHeader: "unchanged"};\n')
        self.config = {'l1ChainID': 900, 'l2ChainID': 53077,
                       'govBatchBlockInterval': 200, 'govBatchTimeout': 600,
                       'gasPriceOracleOwner': '0x' + '22' * 20}
        (configs / 'devnet-deploy-config.json').write_text(json.dumps(self.config))
        self.legacy_record = {'name': 'Proxy__L1Staking', 'address': '0x' + '88' * 20, 'number': 1}
        legacy_file = root / 'legacy-l1.json'
        legacy_file.write_text(json.dumps([self.legacy_record]))
        self.roles = {'sequencer': '0x' + '11' * 20, 'deployer': '0x' + '22' * 20,
                      'batch_submitter': '0x' + '33' * 20}
        # Use invalid placeholders created by this test, without reading local identity files.
        self.args = SimpleNamespace(
            execution_client='geth', cluster=False, debugccc=False,
            sequencer_address=self.roles['sequencer'], sequencer_upgrade_offset_seconds=0,
            sequencer_private_key='test-sequencer', deployer_private_key='test-deployer',
            batch_submitter_private_key='test-submitter',
            gas_oracle_private_key='', batch_block_interval=200, batch_timeout=600,
            legacy_l1_deployment_file=str(legacy_file),
        )
        self.commands = []
        self.fail = None
        self.real_start_l2 = devnet.start_l2
        self.patches = [
            patch.object(devnet, 'validate_parameters', return_value=self.roles),
            patch.object(devnet, 'setup_devnet_nodes'),
            patch.object(devnet, 'test_port', return_value=True),
            patch.object(devnet, 'l1_identity', return_value='0x' + '44' * 32),
            patch.object(devnet, 'verify_l1_contracts'),
            patch.object(devnet, 'start_l2'),
            patch.object(devnet, 'run_command', side_effect=self.command),
            patch.object(devnet, 'rpc_result', side_effect=lambda url, method, params:
                         '0x10' if method == 'eth_blockNumber' else '0x6000'),
        ]
        self.mocks = [item.start() for item in self.patches]
        for item in self.patches:
            self.addCleanup(item.stop)

    def command(self, command, **kwargs):
        self.commands.append((command, kwargs))
        action = command[2] if command[:2] == ['npx', 'hardhat'] else None
        output = Path(self.paths.devnet_dir)
        if action == 'deploy':
            rows = [dict(self.legacy_record) if name == 'Proxy__L1Staking'
                    else {'name': name, 'address': '0x' + f'{index + 1:040x}'}
                    for index, name in enumerate(devnet.REQUIRED_DEPLOYMENTS)]
            Path(self.paths.deployment_dir).write_text(json.dumps(rows))
        if len(command) > 1 and command[1].endswith('devnet-l2genesis.sh') and '--verify-existing' not in command:
            source = Path(command[command.index('--deploy-config') + 1])
            normalized = {**json.loads(source.read_text()), 'l1WETH': '0x' + 'ab' * 20}
            (output / 'deploy-config.json').write_text(json.dumps(normalized))
            for name in ('genesis-l2.json', 'rollup.json', 'genesis-batch-header.json', 'genesis.done'):
                (output / name).write_text('{}')
            (output / 'deployment-config.json').write_text(json.dumps({'batchHeader': '0x0102'}))
        if action == 'initialize':
            rows = json.loads(Path(self.paths.deployment_dir).read_text())
            if not any(row['name'] == 'Impl__Rollup' for row in rows):
                rows.append({'name': 'Impl__Rollup', 'address': '0x' + 'ab' * 20})
                Path(self.paths.deployment_dir).write_text(json.dumps(rows))
        if self.fail is not None and action == self.fail:
            raise RuntimeError('simulated interruption')
        return SimpleNamespace(stdout='')

    def state(self):
        return json.loads((Path(self.paths.devnet_dir) / 'deployment-state.json').read_text())

    def test_complete_run_uses_override_and_repeated_run_does_not_redeploy(self):
        original = Path(self.paths.contracts_config).read_bytes()
        devnet.devnet_deploy(self.paths, self.args)
        self.assertEqual(self.state()['phase'], 'complete')
        self.assertTrue((Path(self.paths.devnet_dir) / 'done').exists())
        self.assertEqual(Path(self.paths.contracts_config).read_bytes(), original)
        actions = [command[2] for command, _ in self.commands if command[:2] == ['npx', 'hardhat']]
        self.assertEqual(actions, ['deploy', 'initialize', 'fund', 'register', 'verify-deployment'])
        for command, options in self.commands:
            if command[:2] == ['npx', 'hardhat']:
                self.assertEqual(options['env']['DEPLOYER_PRIVATE_KEY'], self.args.deployer_private_key)
                self.assertEqual(options['env']['L1_RPC_URL'], 'http://localhost:9545')
                self.assertEqual(options['env']['DOTENV_CONFIG_PATH'], '/dev/null')
                self.assertEqual(options['env']['firstSequencerAddress'], self.roles['sequencer'])
                self.assertTrue(options['env']['DEPLOY_CONFIG_OVERRIDE'].endswith('devnet-contract-config.json'))
        overrides = json.loads((Path(self.paths.devnet_dir) / 'devnet-contract-config.json').read_text())
        self.assertEqual(overrides['batchHeader'], '0x0102')
        self.assertEqual(overrides['submitterOwner'], self.roles['deployer'])
        saved_time = self.state()['sequencer_upgrade_time']
        self.commands.clear()
        devnet.devnet_deploy(self.paths, self.args)
        self.assertEqual([command[2] for command, _ in self.commands if command[:2] == ['npx', 'hardhat']], ['verify-deployment'])
        self.assertTrue(any('--verify-existing' in command for command, _ in self.commands))
        self.assertTrue(any('--runtime' in command for command, _ in self.commands))
        self.assertEqual(self.state()['sequencer_upgrade_time'], saved_time)

    def test_missing_legacy_contract_stops_before_node_setup_or_deployment(self):
        self.args.legacy_l1_deployment_file = None
        with self.assertRaisesRegex(RuntimeError, 'existing Proxy__L1Staking record is required'):
            devnet.devnet_deploy(self.paths, self.args)
        self.mocks[1].assert_not_called()
        self.mocks[7].assert_not_called()
        self.assertFalse(self.commands)
        self.assertFalse((Path(self.paths.devnet_dir) / 'deployment-state.json').exists())
        self.assertFalse(Path(self.paths.deployment_dir).exists())

    def test_legacy_contract_is_imported_and_reused_without_substitution(self):
        devnet.devnet_deploy(self.paths, self.args)
        self.assertEqual(self.state()['request']['legacy_l1_staking'],
                         {'address': self.legacy_record['address'], 'number': 1})
        genesis_input = json.loads((Path(self.paths.devnet_dir) / 'genesis-input.json').read_text())
        self.assertEqual(genesis_input['l1StakingProxy'], self.legacy_record['address'])
        self.mocks[7].assert_any_call('127.0.0.1:9545', 'eth_getCode',
                                     [self.legacy_record['address'], '0x10'])
        self.args.legacy_l1_deployment_file = None
        self.commands.clear()
        devnet.devnet_deploy(self.paths, self.args)
        self.assertEqual([command[2] for command, _ in self.commands
                          if command[:2] == ['npx', 'hardhat']], ['verify-deployment'])

    def test_legacy_contract_code_is_required_before_hardhat_transactions(self):
        self.mocks[7].side_effect = lambda url, method, params: '0x10' if method == 'eth_blockNumber' else '0x'
        with self.assertRaisesRegex(RuntimeError, 'has no valid nonzero bytecode'):
            devnet.devnet_deploy(self.paths, self.args)
        self.assertFalse(self.commands)
        self.assertEqual(self.state()['phase'], 'prepared')
        self.assertEqual(json.loads(Path(self.paths.deployment_dir).read_text()), [self.legacy_record])

    def test_legacy_contract_future_block_is_rejected_before_reading_code(self):
        self.legacy_record['number'] = 17
        Path(self.args.legacy_l1_deployment_file).write_text(json.dumps([self.legacy_record]))
        with self.assertRaisesRegex(RuntimeError, 'newer than the connected L1 head'):
            devnet.devnet_deploy(self.paths, self.args)
        self.assertEqual(self.mocks[7].call_args_list, [unittest.mock.call(
            '127.0.0.1:9545', 'eth_blockNumber', [])])
        self.assertFalse(self.commands)

    def test_legacy_contract_head_and_bytecode_must_use_valid_hex_encoding(self):
        for head in ('0x', '0x01', '-0x1', '16', 16, None):
            with self.subTest(head=head):
                self.mocks[7].side_effect = lambda url, method, params: head
                with self.assertRaisesRegex(RuntimeError, 'canonical hexadecimal block number'):
                    devnet.devnet_deploy(self.paths, self.args)
                self.assertFalse(self.commands)
        for code in ('0x0', '0x1', '0x00', '0x0000', '0xgg', None):
            with self.subTest(code=code):
                self.mocks[7].side_effect = lambda url, method, params: '0x10' if method == 'eth_blockNumber' else code
                with self.assertRaisesRegex(RuntimeError, 'valid nonzero bytecode'):
                    devnet.devnet_deploy(self.paths, self.args)
                self.assertFalse(self.commands)

    def test_configured_legacy_address_must_match_the_imported_record(self):
        source = Path(self.paths.deploy_config_dir) / 'devnet-deploy-config.json'
        source.write_text(json.dumps({**self.config, 'l1StakingProxy': '0x' + '99' * 20}))
        with self.assertRaisesRegex(RuntimeError, 'must match the confirmed'):
            devnet.devnet_deploy(self.paths, self.args)
        self.mocks[1].assert_not_called()
        self.assertFalse(self.commands)
        self.assertFalse((Path(self.paths.devnet_dir) / 'deployment-state.json').exists())

    def test_explicit_batch_parameters_are_independent_of_genesis_gov_values(self):
        self.args.batch_block_interval = 7
        self.args.batch_timeout = 11
        devnet.devnet_deploy(self.paths, self.args)
        self.assertEqual(self.state()['request']['batch_parameters'],
                         {'batchBlockInterval': 7, 'batchTimeout': 11})
        addresses = {name: '0x' + '12' * 20 for name in devnet.REQUIRED_DEPLOYMENTS}
        for config in (self.config, {}):
            environment = devnet.runtime_environment(self.paths, self.args, config, addresses, 1234)
            self.assertEqual(environment['BATCH_BLOCK_INTERVAL'], '7')
            self.assertEqual(environment['BATCH_TIMEOUT'], '11')
        genesis_input = json.loads((Path(self.paths.devnet_dir) / 'genesis-input.json').read_text())
        self.assertEqual(genesis_input['govBatchBlockInterval'], 200)
        self.assertEqual(genesis_input['govBatchTimeout'], 600)

    def test_changed_explicit_batch_parameters_reject_resume_before_commands(self):
        devnet.devnet_deploy(self.paths, self.args)
        self.commands.clear()
        self.mocks[5].reset_mock()
        self.args.batch_timeout += 1
        with self.assertRaisesRegex(RuntimeError, 'different configuration'):
            devnet.devnet_deploy(self.paths, self.args)
        self.assertFalse(self.commands)
        self.mocks[5].assert_not_called()

    def test_old_deployment_without_explicit_batch_parameters_preserves_all_files(self):
        devnet.devnet_deploy(self.paths, self.args)
        output = Path(self.paths.devnet_dir)
        state = self.state()
        del state['request']['batch_parameters']
        (output / 'deployment-state.json').write_text(json.dumps(state))
        Path(self.paths.env_file).write_text('EXISTING_RUNTIME=preserved\n')
        before = {str(path.relative_to(output)): path.read_bytes()
                  for path in output.rglob('*') if path.is_file()}
        self.commands.clear()
        self.mocks[5].reset_mock()
        with self.assertRaisesRegex(RuntimeError, 'no recorded explicit batch parameters'):
            devnet.devnet_deploy(self.paths, self.args)
        for action in ('start', 'rebuild'):
            self.args.service_action = action
            with self.assertRaisesRegex(RuntimeError, 'no recorded explicit batch parameters'):
                devnet.devnet_service_action(self.paths, self.args)
        self.assertFalse(self.commands)
        self.mocks[5].assert_not_called()
        self.assertEqual(before, {str(path.relative_to(output)): path.read_bytes()
                                 for path in output.rglob('*') if path.is_file()})

    def test_interrupted_initialize_preserves_genesis_and_resumes(self):
        self.fail = 'initialize'
        with self.assertRaisesRegex(RuntimeError, 'interruption'):
            devnet.devnet_deploy(self.paths, self.args)
        self.assertEqual(self.state()['phase'], 'genesis')
        self.assertFalse((Path(self.paths.devnet_dir) / 'done').exists())
        original = Path(self.paths.genesis_l2_path).read_bytes()
        self.fail = None
        self.commands.clear()
        devnet.devnet_deploy(self.paths, self.args)
        actions = [command[2] for command, _ in self.commands if command[:2] == ['npx', 'hardhat']]
        self.assertEqual(actions, ['initialize', 'fund', 'register', 'verify-deployment'])
        self.assertTrue(any('--verify-existing' in command for command, _ in self.commands))
        self.assertEqual(Path(self.paths.genesis_l2_path).read_bytes(), original)

    def test_failed_l2_start_has_no_done_and_retry_keeps_upgrade_time(self):
        self.mocks[5].side_effect = RuntimeError('L2 unavailable')
        with self.assertRaisesRegex(RuntimeError, 'unavailable'):
            devnet.devnet_deploy(self.paths, self.args)
        self.assertEqual(self.state()['phase'], 'registered')
        self.assertFalse((Path(self.paths.devnet_dir) / 'done').exists())
        upgrade_time = self.state()['sequencer_upgrade_time']
        self.mocks[5].side_effect = None
        self.commands.clear()
        devnet.devnet_deploy(self.paths, self.args)
        self.assertEqual([command[2] for command, _ in self.commands if command[:2] == ['npx', 'hardhat']], ['verify-deployment'])
        self.assertEqual(self.state()['sequencer_upgrade_time'], upgrade_time)

    def test_untracked_existing_output_is_preserved_and_rejected(self):
        sentinel = Path(self.paths.devnet_dir) / 'devnetL1.json'
        original = json.dumps([self.legacy_record])
        sentinel.write_text(original)
        with self.assertRaisesRegex(RuntimeError, 'no deployment-state.json'):
            devnet.devnet_deploy(self.paths, self.args)
        self.assertEqual(sentinel.read_text(), original)
        self.mocks[1].assert_not_called()
        self.assertFalse(self.commands)

    def test_changed_identity_or_artifact_is_rejected(self):
        devnet.devnet_deploy(self.paths, self.args)
        self.roles['sequencer'] = '0x' + '55' * 20
        with self.assertRaisesRegex(RuntimeError, 'different configuration'):
            devnet.devnet_deploy(self.paths, self.args)
        self.roles['sequencer'] = self.args.sequencer_address
        Path(self.paths.genesis_l2_path).write_text('changed')
        with self.assertRaisesRegex(RuntimeError, 'artifact changed'):
            devnet.devnet_deploy(self.paths, self.args)

    def test_removed_ignored_source_fields_require_original_source_to_resume(self):
        source = Path(self.paths.deploy_config_dir) / 'devnet-deploy-config.json'
        original = {
            **self.config,
            'BLOCK_SIGNER_ADDRESS': '0x' + '66' * 20,
            'maxTxPerBlock': 1000,
            'morphTokenName': 'Morph Token',
            'morphTokenSymbol': 'Morph',
            'morphTokenOwner': self.roles['deployer'],
            'morphTokenInitialSupply': 1000000000,
            'morphTokenDailyInflationRate': 1,
            'BLOCK_SIGNER_PRIVATE_KEY': 'unused-test-key',
            'l2StakingPks': ['unused-test-staking-key'],
        }
        original_bytes = json.dumps(original).encode()
        source.write_bytes(original_bytes)
        devnet.devnet_deploy(self.paths, self.args)
        Path(self.paths.env_file).write_text('EXISTING_RUNTIME=preserved\n')
        output = Path(self.paths.devnet_dir)
        before = {str(path.relative_to(output)): path.read_bytes()
                  for path in output.rglob('*') if path.is_file()}
        self.commands.clear()
        for mock in self.mocks[1:]:
            mock.reset_mock()

        # Go ignores these removed fields, but the existing deployment records
        # still identify the original source configuration by its exact hash.
        source.write_text(json.dumps(self.config))
        with self.assertRaisesRegex(RuntimeError, 'different configuration'):
            devnet.devnet_deploy(self.paths, self.args)
        self.assertFalse(self.commands)
        self.mocks[1].assert_not_called()  # Node setup.
        self.mocks[3].assert_not_called()  # L1 identity query.
        self.mocks[5].assert_not_called()  # Runtime generation and service startup.
        self.assertEqual(before, {str(path.relative_to(output)): path.read_bytes()
                                 for path in output.rglob('*') if path.is_file()})

        source.write_bytes(original_bytes)
        devnet.devnet_deploy(self.paths, self.args)
        actions = [command[2] for command, _ in self.commands
                   if command[:2] == ['npx', 'hardhat']]
        self.assertEqual(actions, ['verify-deployment'])
        self.assertTrue(any('--verify-existing' in command for command, _ in self.commands))
        self.assertTrue(any('--runtime' in command for command, _ in self.commands))
        self.mocks[5].assert_called_once()
        for name in ('genesis-l2.json', 'genesis-input.json', 'genesis.done', 'devnetL1.json'):
            self.assertEqual((output / name).read_bytes(), before[name])

    def test_runtime_file_has_only_public_parameters(self):
        environment = devnet.runtime_environment(self.paths, self.args, self.config,
            {name: '0x' + '12' * 20 for name in devnet.REQUIRED_DEPLOYMENTS}, 1234)
        self.assertEqual(environment['MORPH_SUBMITTER'], '0x' + '12' * 20)
        self.assertEqual(environment['BATCH_BLOCK_INTERVAL'], '200')
        self.assertFalse(any('PRIVATE_KEY' in key for key in environment))
        self.assertFalse(any(value in environment.values() for value in
                             (self.args.sequencer_private_key, self.args.deployer_private_key,
                              self.args.batch_submitter_private_key)))
        with patch.object(devnet, 'wait_up'), patch.object(devnet, 'wait_for_rpc_server'), \
                patch.object(devnet, 'eth_blockNumber', return_value=1):
            self.args.cluster = True
            self.real_start_l2(self.paths, self.args, self.config,
                {name: '0x' + '12' * 20 for name in devnet.REQUIRED_DEPLOYMENTS}, 1234)
        contents = Path(self.paths.env_file).read_text()
        self.assertNotIn('PRIVATE_KEY', contents)
        self.assertNotIn(self.args.sequencer_private_key, contents)
        for command, options in self.commands:
            self.assertIn('--env-file', command)
            self.assertIn('docker-compose-cluster.yml', command)
            self.assertEqual(options['env']['ACTIVE_SEQUENCER_PRIVATE_KEY'], '')
            self.assertEqual(options['env']['SEQUENCER_PRIVATE_KEY'], self.args.sequencer_private_key)

    def test_original_genesis_script_receives_saved_input_and_verifies_it_on_reuse(self):
        devnet.devnet_deploy(self.paths, self.args)
        input_path = Path(self.paths.devnet_dir) / 'genesis-input.json'
        input_before = input_path.read_bytes()
        genesis_commands = [command for command, _ in self.commands
                            if command[:2] == ['bash', str(Path(self.paths.L2_dir) / 'devnet-l2genesis.sh')]]
        self.assertEqual(len(genesis_commands), 1)
        command = genesis_commands[0]
        for flag, expected in (
                ('--deployment-file', self.paths.deployment_dir),
                ('--deploy-config', str(input_path)),
                ('--output-dir', self.paths.devnet_dir),
                ('--l1-rpc', 'http://localhost:9545')):
            self.assertEqual(command[command.index(flag) + 1], expected)
        self.assertNotIn('--verify-existing', command)
        self.commands.clear()
        devnet.devnet_deploy(self.paths, self.args)
        self.assertIn(([*command, '--verify-existing'], {'cwd': self.paths.L2_dir}), self.commands)
        self.assertEqual(input_path.read_bytes(), input_before)
        self.assertEqual(self.state()['phase'], 'complete')

    def test_service_start_and_rebuild_use_saved_topology_and_matching_key(self):
        self.args.batch_block_interval = 7
        self.args.batch_timeout = 11
        self.args.cluster = True
        self.args.execution_client = 'reth'
        devnet.devnet_deploy(self.paths, self.args)
        self.args.cluster = False
        self.args.execution_client = 'geth'
        self.args.batch_submitter_private_key = '04' * 32
        self.args.batch_block_interval = 77
        self.args.batch_timeout = 88
        for action in ('start', 'rebuild'):
            self.commands.clear()
            self.args.service_action = action
            with patch.object(devnet, 'run_command_capture_output', return_value=SimpleNamespace(stdout=self.roles['batch_submitter'])):
                devnet.devnet_service_action(self.paths, self.args)
            self.assertEqual(len(self.commands), 2)
            command, options = self.commands[-1]
            self.assertIn('--env-file', command)
            self.assertIn('docker-compose-cluster-reth.yml', command)
            self.assertIn('--no-deps', command)
            self.assertEqual(command[-1], 'tx-submitter-0')
            self.assertEqual('--build' in command, action == 'rebuild')
            self.assertEqual(options['env']['BATCH_SUBMITTER_PRIVATE_KEY'], self.args.batch_submitter_private_key)
            self.assertEqual(options['env']['SEQUENCER_PRIVATE_KEY'], '')
            self.assertEqual(options['env']['BATCH_BLOCK_INTERVAL'], '7')
            self.assertEqual(options['env']['BATCH_TIMEOUT'], '11')
            self.assertNotIn('PRIVATE_KEY', Path(self.paths.env_file).read_text())
            self.assertEqual(self.state()['phase'], 'complete')

    def assert_service_stop_preserves_files(self):
        output = Path(self.paths.devnet_dir)
        def saved_files():
            return {str(path.relative_to(output)): path.read_bytes()
                    for path in output.rglob('*')
                    if path.is_file() and path.name != '.deployment.lock'}
        before = saved_files()
        self.commands.clear()
        for index in (3, 4, 7):
            self.mocks[index].reset_mock()
        with patch.object(devnet, 'run_command_capture_output') as cast, \
                patch.object(devnet, 'read_json') as read_json, \
                patch.object(devnet, 'compose_runtime') as compose_runtime:
            devnet.devnet_service_action(self.paths, SimpleNamespace(service_action='stop'))
            cast.assert_not_called()
            read_json.assert_not_called()
            compose_runtime.assert_not_called()
        for index in (3, 4, 7):
            self.mocks[index].assert_not_called()
        self.assertEqual(self.commands, [([
            'docker', 'compose', '--project-name', 'docker', '--env-file', devnet.os.devnull,
            '-f', 'docker-compose-devnet.yml', 'stop', 'tx-submitter-0',
        ], {'cwd': self.paths.ops_dir,
            'env': {'NODE_DATA_DIR': '/data', 'JWT_SECRET_PATH': '/jwt-secret.txt'}})])
        self.assertEqual(saved_files(), before)

    def test_service_stop_needs_no_signing_key_or_network(self):
        devnet.devnet_deploy(self.paths, self.args)
        Path(self.paths.env_file).write_text('EXISTING_RUNTIME=preserved\n')
        self.assert_service_stop_preserves_files()

    def test_service_stop_after_containers_start_but_deployment_is_interrupted(self):
        self.mocks[5].side_effect = self.real_start_l2
        with patch.object(devnet, 'wait_up', side_effect=RuntimeError('simulated startup timeout')):
            with self.assertRaisesRegex(RuntimeError, 'simulated startup timeout'):
                devnet.devnet_deploy(self.paths, self.args)
        self.assertTrue(any(command[-2:] == ['up', '-d'] for command, _ in self.commands))
        self.assertEqual(self.state()['phase'], 'registered')
        self.assertFalse((Path(self.paths.devnet_dir) / 'done').exists())
        self.commands.clear()
        for action in ('start', 'rebuild'):
            self.args.service_action = action
            with self.assertRaisesRegex(RuntimeError, 'Finish the existing devnet deployment'):
                devnet.devnet_service_action(self.paths, self.args)
        self.assertFalse(self.commands)
        self.assert_service_stop_preserves_files()

    def test_service_stop_without_output_directory_or_source_files(self):
        output = Path(self.paths.devnet_dir)
        output.rmdir()
        (Path(self.paths.deploy_config_dir) / 'devnet-deploy-config.json').unlink()
        Path(self.paths.contracts_config).unlink()
        self.assert_service_stop_preserves_files()
        self.assertEqual([path.name for path in output.iterdir()], ['.deployment.lock'])

    def test_service_stop_preserves_corrupt_metadata_and_runtime_file(self):
        output = Path(self.paths.devnet_dir)
        for name in ('deployment-state.json', 'done', 'runtime.env'):
            (output / name).write_text('invalid and incomplete content\n')
        self.assert_service_stop_preserves_files()

    def test_service_stop_retains_deployment_lock(self):
        self.args.service_action = 'stop'
        with open(Path(self.paths.devnet_dir) / '.deployment.lock', 'a') as lock:
            devnet.fcntl.flock(lock, devnet.fcntl.LOCK_EX | devnet.fcntl.LOCK_NB)
            with self.assertRaisesRegex(RuntimeError, 'Another devnet operation is running'):
                devnet.devnet_service_action(self.paths, self.args)
        self.assertFalse(self.commands)

    def test_only_stop_allows_changed_config_and_missing_artifacts(self):
        devnet.devnet_deploy(self.paths, self.args)
        config = Path(self.paths.contracts_config)
        original = config.read_bytes()
        config.write_text('changed source configuration\n')
        self.commands.clear()
        for action in ('start', 'rebuild'):
            self.args.service_action = action
            with self.assertRaisesRegex(RuntimeError, 'source configuration changed'):
                devnet.devnet_service_action(self.paths, self.args)
        self.assertFalse(self.commands)
        self.assert_service_stop_preserves_files()
        config.write_bytes(original)
        Path(self.paths.genesis_l2_path).unlink()
        self.commands.clear()
        for action in ('start', 'rebuild'):
            self.args.service_action = action
            with self.assertRaisesRegex(RuntimeError, 'artifact changed or is missing'):
                devnet.devnet_service_action(self.paths, self.args)
        self.assertFalse(self.commands)
        self.assert_service_stop_preserves_files()

    def test_service_rejects_incomplete_state_and_wrong_signer(self):
        for action in ('start', 'rebuild'):
            self.args.service_action = action
            with self.assertRaisesRegex(RuntimeError, 'completed devnet'):
                devnet.devnet_service_action(self.paths, self.args)
        devnet.devnet_deploy(self.paths, self.args)
        self.commands.clear()
        self.args.batch_submitter_private_key = '04' * 32
        for action in ('start', 'rebuild'):
            self.args.service_action = action
            with patch.object(devnet, 'run_command_capture_output', return_value=SimpleNamespace(stdout='0x' + '77' * 20)):
                with self.assertRaisesRegex(RuntimeError, 'does not match'):
                    devnet.devnet_service_action(self.paths, self.args)
        self.assertFalse(self.commands)


class ExistingNodeDataTest(unittest.TestCase):
    def test_incomplete_node_setup_is_preserved(self):
        with tempfile.TemporaryDirectory() as directory:
            data = Path(directory) / 'ops' / 'docker' / '.devnet' / 'node3' / 'data'
            data.mkdir(parents=True)
            sentinel = data / 'consensus-state'
            sentinel.write_text('preserve')
            with patch.object(setup_nodes.subprocess, 'call') as command:
                with self.assertRaisesRegex(RuntimeError, 'no nodes.done'):
                    setup_nodes.setup_devnet_nodes(directory)
                command.assert_not_called()
            self.assertEqual(sentinel.read_text(), 'preserve')

    def test_complete_node_setup_is_reused_and_changed_identity_is_rejected(self):
        with tempfile.TemporaryDirectory() as directory:
            output = Path(directory) / 'ops' / 'docker' / '.devnet'
            for node in setup_nodes.NODE_DIRS:
                config = output / node / 'config'
                config.mkdir(parents=True)
                (config / 'config.toml').write_text('block_sync = true\n')
                for name in ('node_key.json', 'genesis.json', 'priv_validator_key.json'):
                    (config / name).write_text('{}')
            (output / 'nodes.done').write_text(json.dumps({
                'version': 1, 'files': setup_nodes.node_file_hashes(output),
            }))
            with patch.object(setup_nodes.subprocess, 'call') as command, \
                    contextlib.redirect_stdout(io.StringIO()):
                setup_nodes.setup_devnet_nodes(directory)
                command.assert_not_called()
                key = output / 'node0' / 'config' / 'node_key.json'
                key.write_text('changed')
                with self.assertRaisesRegex(RuntimeError, 'differs from nodes.done'):
                    setup_nodes.setup_devnet_nodes(directory)
                self.assertEqual(key.read_text(), 'changed')


class SubprocessAndReadinessTest(unittest.TestCase):
    def test_legacy_contract_records_reject_placeholders_pending_and_submitter_aliases(self):
        with tempfile.TemporaryDirectory() as directory:
            path = Path(directory) / 'legacy.json'
            good = {'name': 'Proxy__L1Staking', 'address': '0x' + '88' * 20, 'number': 1}
            invalid = [[], [good, good], [{**good, 'pending': True}],
                       [{**good, 'pending': None}], [{**good, 'pending': 0}],
                       [{**good, 'pending': ''}], [{**good, 'pending': 'false'}],
                       [{**good, 'number': -1}], [{**good, 'number': True}], [{**good, 'number': 2**53}],
                       [{**good, 'address': '0x' + '00' * 20}],
                       [{**good, 'address': '0x000000000000000000000000000000000000dEaD'}],
                       [good, {'name': 'Proxy__Submitter', 'address': good['address']}]]
            for records in invalid:
                with self.subTest(records=records):
                    path.write_text(json.dumps(records))
                    with self.assertRaises(RuntimeError):
                        devnet.legacy_l1_staking_record(str(path))
            path.write_text(json.dumps([good]))
            self.assertEqual(devnet.legacy_l1_staking_record(str(path)),
                             {'address': good['address'], 'number': 1})

    def test_explicit_batch_parameters_validate_uint64_and_independent_zero_values(self):
        for interval, timeout in ((0, 1), (1, 0), (2**64 - 1, 2**64 - 1)):
            self.assertEqual(devnet.validate_batch_parameters(interval, timeout),
                             {'batchBlockInterval': interval, 'batchTimeout': timeout})
        for interval, timeout in ((0, 0), (-1, 1), (2**64, 1), (1, 2**64),
                                  (True, 1), (1, False), ('1', 1), (1, 1.5)):
            with self.subTest(interval=interval, timeout=timeout):
                with self.assertRaises(RuntimeError):
                    devnet.validate_batch_parameters(interval, timeout)
        args = devnet.parser.parse_args(['--batch-block-interval', '0', '--batch-timeout', '19'])
        self.assertEqual(args.batch_block_interval, 0)
        self.assertEqual(args.batch_timeout, 19)

    def test_only_l1_does_not_require_legacy_contract_records(self):
        with tempfile.TemporaryDirectory() as directory:
            args = SimpleNamespace(polyrepo_dir=directory, only_l1=True, service_action=None)
            with patch.object(devnet.parser, 'parse_args', return_value=args), \
                    patch.object(devnet, 'devnet_l1') as start_l1, \
                    patch.object(devnet, 'requested_legacy_l1_staking') as legacy:
                self.assertTrue(devnet.main())
            start_l1.assert_called_once()
            legacy.assert_not_called()

    def test_incomplete_l1_inputs_stop_before_container_start(self):
        with tempfile.TemporaryDirectory() as directory:
            layer1 = Path(directory) / 'layer1'
            genesis = layer1 / 'genesis'
            genesis.mkdir(parents=True)
            for name in ('genesis.json', 'genesis.ssz'):
                (genesis / name).write_text('existing')
            (layer1 / 'jwt').mkdir()
            (layer1 / 'jwt' / 'jwtsecret').write_text('existing')
            with patch.object(devnet, 'run_command') as command:
                with self.assertRaisesRegex(RuntimeError, 'incomplete'):
                    devnet.devnet_l1(SimpleNamespace(ops_dir=directory))
                command.assert_not_called()
            self.assertEqual((genesis / 'genesis.json').read_text(), 'existing')

    def test_parameter_validation_rejects_wrong_chain_and_signer_before_deployment(self):
        args = SimpleNamespace(sequencer_upgrade_offset_seconds=0,
            sequencer_private_key='01' * 32, deployer_private_key='02' * 32,
            batch_submitter_private_key='03' * 32, gas_oracle_private_key='',
            batch_block_interval=200, batch_timeout=600,
            sequencer_address='0x' + '11' * 20)
        paths = SimpleNamespace(contracts_dir='unused')
        config = {'l1ChainID': 900, 'l2ChainID': 53077,
                  'govBatchBlockInterval': 200, 'govBatchTimeout': 600}
        with patch.object(devnet, 'run_command_capture_output') as cast:
            with self.assertRaisesRegex(RuntimeError, 'l1ChainID=900'):
                devnet.validate_parameters(paths, args, {**config, 'l1ChainID': 1})
            cast.assert_not_called()
            cast.return_value = SimpleNamespace(stdout='0x' + '22' * 20)
            with self.assertRaisesRegex(RuntimeError, 'does not match'):
                devnet.validate_parameters(paths, args, config)

    def test_subprocess_failure_does_not_include_secret_arguments(self):
        command = ['cast', 'wallet', 'address', '--private-key', 'private-test-value']
        with patch.object(devnet.subprocess, 'run', side_effect=subprocess.CalledProcessError(7, command)):
            with self.assertRaises(RuntimeError) as caught:
                devnet.run_command_capture_output(command)
        self.assertNotIn('private-test-value', str(caught.exception))
        self.assertIn('exit code 7', str(caught.exception))

    def test_rpc_wait_checks_chain_id_and_has_a_retry_limit(self):
        with patch.object(devnet, 'rpc_result', return_value='0x1'):
            with self.assertRaisesRegex(RuntimeError, 'unexpected chain ID'):
                devnet.wait_for_rpc_server('test', retries=1, wait_secs=0)
        with patch.object(devnet, 'rpc_result', side_effect=OSError('not ready')) as rpc:
            with self.assertRaisesRegex(RuntimeError, 'Timeout'):
                devnet.wait_for_rpc_server('test', retries=2, wait_secs=0)
            self.assertEqual(rpc.call_count, 2)

    def test_oracle_signer_must_match_l2_owner(self):
        args = SimpleNamespace(sequencer_upgrade_offset_seconds=0,
            sequencer_private_key='01' * 32, deployer_private_key='02' * 32,
            batch_submitter_private_key='03' * 32, gas_oracle_private_key='',
            batch_block_interval=200, batch_timeout=600,
            sequencer_address='0x' + '11' * 20)
        config = {'l1ChainID': 900, 'l2ChainID': 53077,
                  'govBatchBlockInterval': 200, 'govBatchTimeout': 600,
                  'gasPriceOracleOwner': '0x' + '44' * 20}
        derived = [SimpleNamespace(stdout='0x' + byte * 20) for byte in ('11', '22', '33', '22')]
        with patch.object(devnet, 'run_command_capture_output', side_effect=derived):
            with self.assertRaisesRegex(RuntimeError, 'gasPriceOracleOwner'):
                devnet.validate_parameters(SimpleNamespace(contracts_dir='unused'), args, config)
        args.gas_oracle_private_key = '04' * 32
        derived[-1] = SimpleNamespace(stdout=config['gasPriceOracleOwner'])
        with patch.object(devnet, 'run_command_capture_output', side_effect=derived) as cast:
            roles = devnet.validate_parameters(SimpleNamespace(contracts_dir='unused'), args, config)
        self.assertEqual(set(roles), {'sequencer', 'deployer', 'batch_submitter'})
        self.assertEqual(cast.call_args.args[0][-1], args.gas_oracle_private_key)

    def test_tcp_readiness_closes_connections_and_limits_attempts(self):
        with patch.object(devnet.socket, 'create_connection') as connect:
            self.assertTrue(devnet.test_port(1234))
            connect.assert_called_once_with(('127.0.0.1', 1234), timeout=1)
            connect.return_value.__exit__.assert_called_once()
        with patch.object(devnet.socket, 'create_connection', side_effect=OSError('offline')) as connect:
            with self.assertRaisesRegex(RuntimeError, 'Timed out'):
                devnet.wait_up(1234, retries=2, wait_secs=0)
            self.assertEqual(connect.call_count, 2)

    def test_rpc_readiness_closes_failed_http_connections(self):
        with patch.object(devnet.http.client, 'HTTPConnection') as connection:
            connection.return_value.getresponse.side_effect = OSError('disconnected')
            self.assertIsNone(devnet.eth_blockNumber('unused'))
            connection.return_value.close.assert_called_once()

    def test_duplicate_or_pending_deployment_records_are_rejected(self):
        with tempfile.TemporaryDirectory() as directory:
            filename = Path(directory) / 'contracts.json'
            records = [{'name': name, 'address': '0x' + '11' * 20}
                       for name in devnet.REQUIRED_DEPLOYMENTS]
            filename.write_text(json.dumps(records + [records[0]]))
            with self.assertRaisesRegex(RuntimeError, 'duplicate'):
                devnet.deployment_addresses(filename)
            records[0]['pending'] = True
            filename.write_text(json.dumps(records))
            with self.assertRaisesRegex(RuntimeError, 'not been confirmed'):
                devnet.deployment_addresses(filename)


class NodeGenerationTest(unittest.TestCase):
    def test_missing_path_binary_uses_build_output_and_publishes_only_complete_nodes(self):
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            docker = root / 'ops' / 'docker'
            docker.mkdir(parents=True)
            (root / 'node').mkdir()
            binary = root / 'node' / 'build' / 'bin' / 'tendermint'
            calls = []

            def run(command, **kwargs):
                calls.append(command)
                self.assertEqual(command, ['make', 'tendermint'])
                binary.parent.mkdir(parents=True)
                binary.write_text('#!/bin/sh\nexit 0\n')
                binary.chmod(0o700)

            def generate(source, output, executable, cluster):
                self.assertEqual(Path(executable), binary)
                self.assertTrue(cluster)
                self.assertFalse((docker / '.devnet').exists())
                (Path(output) / 'nodes.done').write_text('complete')

            with patch.object(setup_nodes.shutil, 'which', return_value=None), \
                    patch.object(setup_nodes.subprocess, 'run', side_effect=run), \
                    patch.object(setup_nodes, 'generate_node_files', side_effect=generate):
                setup_nodes.setup_devnet_nodes(directory, cluster=True)
            self.assertEqual((docker / '.devnet' / 'nodes.done').read_text(), 'complete')
            self.assertEqual(len(calls), 1)

    def test_failed_generation_preserves_attempt_and_allows_retry(self):
        with tempfile.TemporaryDirectory() as directory:
            docker = Path(directory) / 'ops' / 'docker'
            docker.mkdir(parents=True)

            def fail(source, output, executable, cluster):
                (Path(output) / 'partial-key').write_text('retain')
                raise RuntimeError('generation failed')

            with patch.object(setup_nodes.shutil, 'which', return_value='/test/tendermint'), \
                    patch.object(setup_nodes, 'generate_node_files', side_effect=fail):
                with self.assertRaisesRegex(RuntimeError, 'generation failed'):
                    setup_nodes.setup_devnet_nodes(directory)
            self.assertFalse((docker / '.devnet').exists())
            attempt = next(docker.glob('.devnet-setup-*'))
            self.assertEqual((attempt / 'partial-key').read_text(), 'retain')
            with patch.object(setup_nodes.shutil, 'which', return_value='/test/tendermint'), \
                    patch.object(setup_nodes, 'generate_node_files') as generate:
                setup_nodes.setup_devnet_nodes(directory)
                generate.assert_called_once()
            self.assertTrue(attempt.exists())
            self.assertTrue((docker / '.devnet').exists())

    def test_noncluster_peers_exclude_unstarted_ha_services(self):
        with patch.object(setup_nodes, 'tendermint_node_id', return_value='test-node'):
            peers = setup_nodes.build_persistent_peers('unused', cluster=False)
            self.assertEqual(peers['node0'], '')
            self.assertEqual(peers['node1'], 'test-node@node-0:26656')
            self.assertFalse(any('ha-node' in value for value in peers.values()))
            peers = setup_nodes.build_persistent_peers('unused', cluster=True)
            self.assertIn('ha-node-0:26656', peers['node0'])


if __name__ == '__main__':
    unittest.main()
