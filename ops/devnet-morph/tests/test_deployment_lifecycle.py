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
                       'govBatchBlockInterval': 200, 'govBatchTimeout': 600}
        (configs / 'devnet-deploy-config.json').write_text(json.dumps(self.config))
        self.roles = {'sequencer': '0x' + '11' * 20, 'deployer': '0x' + '22' * 20,
                      'batch_submitter': '0x' + '33' * 20}
        # Use invalid placeholders created by this test, without reading local identity files.
        self.args = SimpleNamespace(
            execution_client='geth', cluster=False, debugccc=False,
            sequencer_address=self.roles['sequencer'], sequencer_upgrade_offset_seconds=0,
            sequencer_private_key='test-sequencer', deployer_private_key='test-deployer',
            batch_submitter_private_key='test-submitter',
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
        ]
        self.mocks = [item.start() for item in self.patches]
        for item in self.patches:
            self.addCleanup(item.stop)

    def command(self, command, **kwargs):
        self.commands.append((command, kwargs))
        action = command[2] if command[:2] == ['npx', 'hardhat'] else None
        output = Path(self.paths.devnet_dir)
        if action == 'deploy':
            rows = [{'name': name, 'address': '0x' + f'{index + 1:040x}'}
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
        sentinel.write_text('old deployment')
        with self.assertRaisesRegex(RuntimeError, 'no deployment-state.json'):
            devnet.devnet_deploy(self.paths, self.args)
        self.assertEqual(sentinel.read_text(), 'old deployment')
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
        self.args.cluster = True
        self.args.execution_client = 'reth'
        devnet.devnet_deploy(self.paths, self.args)
        self.args.cluster = False
        self.args.execution_client = 'geth'
        self.args.batch_submitter_private_key = '04' * 32
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
            self.assertNotIn('PRIVATE_KEY', Path(self.paths.env_file).read_text())
            self.assertEqual(self.state()['phase'], 'complete')

    def test_service_stop_needs_no_signing_key_or_network(self):
        devnet.devnet_deploy(self.paths, self.args)
        self.commands.clear()
        self.mocks[3].reset_mock()
        self.mocks[4].reset_mock()
        self.args.service_action = 'stop'
        self.args.batch_submitter_private_key = None
        with patch.object(devnet, 'run_command_capture_output') as cast:
            devnet.devnet_service_action(self.paths, self.args)
            cast.assert_not_called()
        self.mocks[3].assert_not_called()
        self.mocks[4].assert_not_called()
        self.assertEqual(self.commands[-1][0][-2:], ['stop', 'tx-submitter-0'])

    def test_service_rejects_incomplete_state_and_wrong_signer(self):
        self.args.service_action = 'start'
        with self.assertRaisesRegex(RuntimeError, 'completed devnet'):
            devnet.devnet_service_action(self.paths, self.args)
        devnet.devnet_deploy(self.paths, self.args)
        self.commands.clear()
        self.args.batch_submitter_private_key = '04' * 32
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
    def test_parameter_validation_rejects_wrong_chain_and_signer_before_deployment(self):
        args = SimpleNamespace(sequencer_upgrade_offset_seconds=0,
            sequencer_private_key='01' * 32, deployer_private_key='02' * 32,
            batch_submitter_private_key='03' * 32, sequencer_address='0x' + '11' * 20)
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


if __name__ == '__main__':
    unittest.main()
