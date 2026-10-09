import os
from pathlib import Path
import subprocess
import tarfile
import tempfile
import unittest


PUBLICNODE = Path(__file__).resolve().parents[1]


class PublicNodeCommandsTest(unittest.TestCase):
    def setUp(self):
        self.temp = tempfile.TemporaryDirectory()
        self.addCleanup(self.temp.cleanup)
        self.root = Path(self.temp.name)
        self.node = self.root / "node home"
        self.snapshot = self.root / "snapshot"
        self.command_env = dict(os.environ)
        self.env_file = self.root / "node.env"
        self.env_file.write_text(
            "NODE_HOME=" + str(self.node) + "\n"
            "GETH_DATA_DIR=${NODE_HOME}/geth-data\n"
            "NODE_DATA_DIR=${NODE_HOME}/node-data\n"
            "GETH_ENTRYPOINT_FILE=./holesky/entrypoint-geth.sh\n"
            "JWT_SECRET_FILE=${NODE_HOME}/jwt-secret.txt\n"
            "SNAPSHOT_NAME=snapshot\n"
        )

    def make(self, target, success=True):
        result = subprocess.run(
            ["make", "-C", str(PUBLICNODE), "ENV_FILE=" + str(self.env_file),
             "SNAPSHOT_DIR=" + str(self.snapshot), target],
            capture_output=True, text=True, env=self.command_env,
        )
        if success:
            self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        else:
            self.assertNotEqual(result.returncode, 0, result.stdout + result.stderr)
        return result

    def snapshot_fixture(self):
        (self.snapshot / "geth/chaindata").mkdir(parents=True)
        (self.snapshot / "geth/chaindata/CURRENT").write_text("fixture database")
        (self.snapshot / "data").mkdir()
        (self.snapshot / "data/.state").write_text("fixture state")

    def test_missing_snapshot_does_not_create_node_home(self):
        self.make("prepare-holesky-data", success=False)
        self.assertFalse(self.node.exists())

    def test_prepare_preserves_snapshot_layout_and_existing_data(self):
        self.snapshot_fixture()
        self.make("prepare-holesky-data")
        self.assertEqual((self.node / "geth-data/geth/chaindata/CURRENT").read_text(), "fixture database")
        self.assertEqual((self.node / "node-data/data/.state").read_text(), "fixture state")
        marker = self.node / "node-data/data/operator-marker"
        marker.write_text("keep")
        self.make("prepare-holesky-data")
        self.assertEqual(marker.read_text(), "keep")
        self.assertTrue(self.snapshot.exists())

    def test_incomplete_node_home_is_preserved(self):
        self.node.mkdir()
        (self.node / "operator-marker").write_text("keep")
        self.snapshot_fixture()
        self.make("prepare-holesky-data", success=False)
        self.assertEqual((self.node / "operator-marker").read_text(), "keep")
        self.assertFalse((self.node / "geth-data").exists())

    def test_jwt_is_private_and_reused(self):
        self.make("generate-jwt")
        jwt = self.node / "jwt-secret.txt"
        original = jwt.read_text()
        self.assertEqual(jwt.stat().st_mode & 0o777, 0o600)
        self.assertEqual(len(original.strip()), 64)
        self.make("generate-jwt")
        self.assertEqual(jwt.read_text(), original)
        jwt.write_text("invalid")
        self.make("generate-jwt", success=False)
        self.assertEqual(jwt.read_text(), "invalid")

    def test_download_validates_layout_before_publication(self):
        source = self.root / "archive-source/snapshot"
        (source / "geth/chaindata").mkdir(parents=True)
        (source / "data").mkdir()
        (source / "data/state.db").write_text("fixture state")
        archive = self.root / "fixture.tar.gz"
        with tarfile.open(archive, "w:gz") as output:
            output.add(source, arcname="snapshot")
        binaries = self.root / "bin"
        binaries.mkdir()
        wget = binaries / "wget"
        wget.write_text('#!/bin/sh\ncp "$SNAPSHOT_FIXTURE" "$2"\n')
        wget.chmod(0o700)
        self.command_env["PATH"] = str(binaries) + os.pathsep + os.environ["PATH"]
        self.command_env["SNAPSHOT_FIXTURE"] = str(archive)
        self.make("download-and-decompress-snapshot")
        self.assertEqual((self.snapshot / "data/state.db").read_text(), "fixture state")
        self.assertTrue(archive.exists())
        self.make("download-and-decompress-snapshot", success=False)
        self.assertEqual((self.snapshot / "data/state.db").read_text(), "fixture state")

    def test_geth_rejects_missing_network_before_accessing_data(self):
        env = dict(os.environ)
        env.pop("GETH_NETWORK_ID", None)
        result = subprocess.run(["sh", str(PUBLICNODE / "holesky/entrypoint-geth.sh")], env=env, capture_output=True, text=True)
        self.assertNotEqual(result.returncode, 0)
        self.assertIn("GETH_NETWORK_ID", result.stderr)


if __name__ == "__main__":
    unittest.main()
