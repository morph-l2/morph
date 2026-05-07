import unittest
from pathlib import Path
import sys

sys.path.insert(0, str(Path(__file__).resolve().parent))
from devnet import compose_file_args


class ExecutionClientComposeArgsTest(unittest.TestCase):
    def test_geth_uses_base_compose_file(self):
        self.assertEqual(compose_file_args("geth"), ["-f", "docker-compose-4nodes.yml"])

    def test_reth_adds_reth_override_file(self):
        self.assertEqual(
            compose_file_args("reth"),
            ["-f", "docker-compose-4nodes.yml", "-f", "docker-compose-reth.yml"],
        )


if __name__ == "__main__":
    unittest.main()
