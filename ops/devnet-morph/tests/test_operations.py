"""Validate tracked operational inputs and command failure boundaries without live services."""

import json
import os
from pathlib import Path
import shutil
import subprocess
import tempfile
import unittest

ROOT = Path(__file__).resolve().parents[3]


class OperationsTest(unittest.TestCase):
    def tracked_files(self):
        paths = subprocess.check_output(['git', 'ls-files', '-z', 'ops'], cwd=ROOT).decode().split('\0')
        return [ROOT / path for path in paths if path and (ROOT / path).is_file()]

    def test_tracked_shell_scripts_parse_with_their_declared_shell(self):
        scripts = [path for path in self.tracked_files() if path.suffix == '.sh']
        self.assertTrue(scripts)
        for path in scripts:
            with self.subTest(script=str(path.relative_to(ROOT))):
                shell = 'bash' if 'bash' in path.read_text().splitlines()[0] else 'sh'
                result = subprocess.run([shell, '-n', str(path)], text=True, capture_output=True)
                self.assertEqual(result.returncode, 0, result.stderr)

    def test_tracked_json_inputs_are_valid_without_printing_their_contents(self):
        for path in self.tracked_files():
            if path.suffix == '.json':
                with self.subTest(config=str(path.relative_to(ROOT))):
                    with path.open() as source:
                        self.assertIsNotNone(json.load(source))

    def test_root_docker_build_dry_run_uses_selected_compose_files(self):
        with tempfile.TemporaryDirectory() as directory:
            binary = Path(directory) / 'docker'
            called = Path(directory) / 'docker-called'
            binary.write_text('#!/bin/sh\ntouch "$DOCKER_CALLED"\nexit 99\n')
            binary.chmod(0o700)
            env = {**os.environ, 'PATH': f'{directory}:{os.environ["PATH"]}', 'DOCKER_CALLED': str(called)}
            result = subprocess.run(['make', '-n', 'docker-build', 'EXECUTION_CLIENT=reth', 'DEVNET_CLUSTER=true'],
                                    cwd=ROOT, env=env, text=True, capture_output=True)
            self.assertEqual(result.returncode, 0, result.stderr)
            self.assertFalse(called.exists(), 'Make dry runs must not invoke Docker during parsing')
            build = next(line for line in result.stdout.splitlines() if 'docker compose' in line)
            self.assertIn('--project-name docker', build)
            self.assertIn('docker-compose-cluster-reth.yml', build)
            self.assertLess(build.index('docker-compose-cluster.yml'), build.index('docker-compose-reth.yml'))

    def test_failed_cleanup_keeps_existing_l1_l2_and_deployment_files(self):
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            shutil.copyfile(ROOT / 'Makefile', root / 'Makefile')
            scripts = root / 'ops' / 'docker' / 'layer1' / 'scripts'
            scripts.mkdir(parents=True)
            shutil.copy2(ROOT / 'ops/docker/layer1/scripts/clean.sh', scripts / 'clean.sh')
            sentinels = []
            for relative in ('ops/docker/layer1/genesis/keep', 'ops/docker/.devnet/keep',
                             'ops/l2-genesis/.devnet/keep'):
                path = root / relative
                path.parent.mkdir(parents=True, exist_ok=True)
                path.write_text('preserve')
                sentinels.append(path)
            binary_dir = root / 'bin'
            binary_dir.mkdir()
            docker = binary_dir / 'docker'
            docker.write_text('#!/bin/sh\nexit 7\n')
            docker.chmod(0o700)
            env = {**os.environ, 'PATH': f'{binary_dir}:{os.environ["PATH"]}'}
            result = subprocess.run(['make', 'devnet-clean-build'], cwd=root, env=env,
                                    text=True, capture_output=True)
            self.assertNotEqual(result.returncode, 0)
            for path in sentinels:
                self.assertEqual(path.read_text(), 'preserve')


if __name__ == '__main__':
    unittest.main()
