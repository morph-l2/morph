"""Exercise Docker entrypoints and L1 lifecycle scripts without running a network."""

import json
import os
from pathlib import Path
import shutil
import subprocess
import tempfile
import unittest

DOCKER_DIR = Path(__file__).resolve().parents[2] / 'docker'


class DockerScriptTests(unittest.TestCase):
    def setUp(self):
        self.temp = tempfile.TemporaryDirectory()
        self.addCleanup(self.temp.cleanup)
        self.root = Path(self.temp.name)
        self.docker_dir = self.root / 'docker'
        self.layer1 = self.docker_dir / 'layer1'
        (self.layer1 / 'scripts').mkdir(parents=True)
        for script in (DOCKER_DIR / 'layer1/scripts').glob('*.sh'):
            shutil.copy(script, self.layer1 / 'scripts' / script.name)
        (self.docker_dir / 'docker-compose-devnet.yml').write_text('services: {}\n')
        for file in ('configs/values.env.template', 'jwt/jwtsecret',
                     'keystores/layer1/keys/validator_definitions.yml'):
            path = self.layer1 / file
            path.parent.mkdir(parents=True, exist_ok=True)
            path.write_text('test input {{GENESIS_TIMESTAMP}}\n')
        self.bin_dir = self.root / 'bin'
        self.bin_dir.mkdir()
        self.log = self.root / 'commands.jsonl'
        self.env = {**os.environ, 'PATH': str(self.bin_dir) + os.pathsep + os.environ['PATH'],
                    'TEST_LOG': str(self.log), 'COMPOSE_PROJECT_NAME': 'isolated-test',
                    'L1_START_ATTEMPTS': '1', 'L1_START_INTERVAL': '0'}
        self.install('docker', '''
import json, os, sys
from pathlib import Path
args = sys.argv[1:]
with open(os.environ['TEST_LOG'], 'a') as log:
    log.write(json.dumps(args) + '\\n')
if os.environ.get('FAIL_COMMAND') and os.environ['FAIL_COMMAND'] in args:
    sys.exit(17)
if args[:2] == ['volume', 'ls']:
    if os.environ.get('TEST_VOLUME'):
        print(os.environ['TEST_VOLUME'])
elif args[:1] == ['run']:
    target = next(value[:-len(':/data')] for value in args if value.endswith(':/data'))
    for name in ['genesis.json', 'genesis.ssz', 'config.yaml', 'deposit_contract_block.txt']:
        if name != os.environ.get('MISSING_OUTPUT'):
            (Path(target) / name).write_text('generated')
elif 'exec' in args:
    print(os.environ.get('TEST_BLOCK', '1'))
''')

    def install(self, name, source):
        path = self.bin_dir / name
        path.write_text('#!' + os.sys.executable + '\n' + source)
        path.chmod(0o755)

    def commands(self):
        return [json.loads(line) for line in self.log.read_text().splitlines()] if self.log.exists() else []

    def run_script(self, name, **env):
        return subprocess.run(['bash', str(self.layer1 / 'scripts' / name)],
                              cwd=self.root, env={**self.env, **env}, text=True,
                              stdout=subprocess.PIPE, stderr=subprocess.PIPE)

    def genesis(self):
        for name in ('genesis.json', 'genesis.ssz', 'config.yaml', 'deposit_contract_block.txt'):
            path = self.layer1 / 'genesis' / name
            path.parent.mkdir(exist_ok=True)
            path.write_text('original')

    def test_generator_publishes_all_required_files(self):
        result = self.run_script('generate-genesis.sh')
        self.assertEqual(result.returncode, 0, result.stderr)
        self.assertEqual((self.layer1 / 'genesis/genesis.json').read_text(), 'generated')
        self.assertFalse((self.layer1 / '.genesis.lock').exists())

    def test_generator_preserves_previous_genesis(self):
        self.genesis()
        result = self.run_script('generate-genesis.sh')
        self.assertNotEqual(result.returncode, 0)
        self.assertEqual((self.layer1 / 'genesis/genesis.json').read_text(), 'original')
        self.assertEqual(self.commands(), [])

    def test_generator_refuses_old_volumes_without_genesis(self):
        result = self.run_script('generate-genesis.sh', TEST_VOLUME='isolated-test_layer1-el-data')
        self.assertNotEqual(result.returncode, 0)
        self.assertFalse(any('run' in command for command in self.commands()))

    def test_generator_failure_does_not_publish(self):
        for env in ({'FAIL_COMMAND': 'run'}, {'MISSING_OUTPUT': 'genesis.ssz'}):
            with self.subTest(env=env):
                result = self.run_script('generate-genesis.sh', **env)
                self.assertNotEqual(result.returncode, 0)
                self.assertFalse((self.layer1 / 'genesis').exists())
                self.assertFalse((self.layer1 / '.genesis.lock').exists())
        self.assertTrue(list(self.layer1.glob('genesis-attempt.*')))

    def test_start_fails_if_chain_does_not_produce_a_block(self):
        self.genesis()
        result = self.run_script('start.sh', TEST_BLOCK='0')
        self.assertNotEqual(result.returncode, 0)
        self.assertIn('startup timeout', result.stderr)

    def test_start_checks_all_inputs_before_docker(self):
        self.genesis()
        (self.layer1 / 'genesis/config.yaml').unlink()
        result = self.run_script('start.sh')
        self.assertNotEqual(result.returncode, 0)
        self.assertEqual(self.commands(), [])

    def test_start_uses_compose_service_exec(self):
        self.genesis()
        result = self.run_script('start.sh')
        self.assertEqual(result.returncode, 0, result.stderr)
        self.assertTrue(any('exec' in cmd and cmd[0] == 'compose' for cmd in self.commands()))

    def test_clean_refuses_active_genesis_generation(self):
        self.genesis()
        (self.layer1 / ".genesis.lock").mkdir()
        result = self.run_script("clean.sh")
        self.assertNotEqual(result.returncode, 0)
        self.assertEqual(self.commands(), [])
        self.assertEqual((self.layer1 / "genesis/genesis.json").read_text(), "original")

    def test_clean_stops_before_deleting_on_docker_failure(self):
        self.genesis()
        result = self.run_script('clean.sh', FAIL_COMMAND='stop')
        self.assertNotEqual(result.returncode, 0)
        self.assertEqual((self.layer1 / 'genesis/genesis.json').read_text(), 'original')
        self.assertEqual(len(self.commands()), 1)

    def test_clean_preserves_genesis_if_volume_removal_fails(self):
        self.genesis()
        result = self.run_script('clean.sh', TEST_VOLUME='selected-volume', FAIL_COMMAND='selected-volume')
        self.assertNotEqual(result.returncode, 0)
        self.assertEqual((self.layer1 / 'genesis/genesis.json').read_text(), 'original')

    def test_clean_removes_only_project_l1_volumes_and_preserves_l2(self):
        self.genesis()
        l2_data = self.docker_dir / 'data'
        l2_data.mkdir()
        (l2_data / 'keep').write_text('L2')
        result = self.run_script('clean.sh', TEST_VOLUME='selected-volume')
        self.assertEqual(result.returncode, 0, result.stderr)
        volume_lists = [cmd for cmd in self.commands() if cmd[:2] == ['volume', 'ls']]
        self.assertEqual(len(volume_lists), 3)
        for command in volume_lists:
            self.assertIn('label=com.docker.compose.project=isolated-test', command)
            self.assertTrue(any(arg.startswith('label=com.docker.compose.volume=layer1-') for arg in command))
        self.assertEqual((l2_data / 'keep').read_text(), 'L2')
        self.assertTrue((self.layer1 / 'jwt/jwtsecret').is_file())
        self.assertFalse((self.layer1 / 'genesis').exists())

    def test_geth_arguments_preserve_spaces_and_exec_after_init(self):
        self.install('geth', '''
import json, os, sys
with open(os.environ['TEST_LOG'], 'a') as log:
    log.write(json.dumps(sys.argv[1:]) + '\\n')
if 'init' in sys.argv and os.environ.get('FAIL_INIT'):
    sys.exit(7)
''')
        env = {**self.env, 'GETH_DATA_DIR': str(self.root / 'db with spaces'),
               'GENESIS_FILE_PATH': str(self.root / 'genesis with spaces.json'),
               'JWT_SECRET_PATH': str(self.root / 'jwt with spaces'), 'BOOT_NODES': 'enode://a,enode://b'}
        result = subprocess.run(['sh', str(DOCKER_DIR / 'entrypoint-l2.sh'), '--extra=value with spaces'],
                                env=env, capture_output=True, text=True)
        self.assertEqual(result.returncode, 0, result.stderr)
        commands = self.commands()
        self.assertEqual(len(commands), 2)
        self.assertIn(env['GENESIS_FILE_PATH'], commands[0])
        self.assertIn('--authrpc.jwtsecret=' + env['JWT_SECRET_PATH'], commands[1])
        self.assertIn('--http.vhosts=*', commands[1])
        self.assertIn('--extra=value with spaces', commands[1])
        self.log.unlink()
        result = subprocess.run(['sh', str(DOCKER_DIR / 'entrypoint-l2.sh')],
                                env={**env, 'FAIL_INIT': '1'}, capture_output=True)
        self.assertEqual(result.returncode, 7)
        self.assertEqual(len(self.commands()), 1)


@unittest.skipUnless(shutil.which('docker'), 'Docker Compose is unavailable')
class ComposeConfigTests(unittest.TestCase):
    def test_all_modes_resolve_paths_without_pwd_or_dotenv(self):
        for cluster, reth in ((False, False), (True, False), (False, True), (True, True)):
            with self.subTest(cluster=cluster, reth=reth):
                command = ['docker', 'compose', '--env-file', os.devnull,
                           '-f', str(DOCKER_DIR / 'docker-compose-devnet.yml')]
                if cluster:
                    command += ['-f', str(DOCKER_DIR / 'docker-compose-cluster.yml')]
                if reth:
                    command += ['-f', str(DOCKER_DIR / 'docker-compose-reth.yml')]
                if cluster and reth:
                    command += ['-f', str(DOCKER_DIR / 'docker-compose-cluster-reth.yml')]
                result = subprocess.run(command + ['config', '--format', 'json'], cwd='/tmp',
                                        env={'PATH': os.environ['PATH'], 'PWD': '/unrelated'},
                                        text=True, capture_output=True)
                self.assertEqual(result.returncode, 0, result.stderr)
                self.assertEqual(result.stderr, '')
                config = json.loads(result.stdout)
                services = config['services']
                self.assertEqual(len(services), 15 if cluster else 9)
                submitter = services['tx-submitter-0']['environment']
                self.assertEqual(submitter['TX_SUBMITTER_BATCH_V2_UPGRADE_TIME'], '0')
                self.assertEqual(submitter['TX_SUBMITTER_L1_PRIVATE_KEY'], '')
                volumes = services['morph-el-0']['volumes']
                jwt = next(vol for vol in volumes if vol['target'] == '/jwt-secret.txt')
                self.assertEqual(jwt['source'], str(DOCKER_DIR / 'jwt-secret.txt'))
                if reth:
                    self.assertNotIn('build', services['morph-el-0'])


if __name__ == '__main__':
    unittest.main()
