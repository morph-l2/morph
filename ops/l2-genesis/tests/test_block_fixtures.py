"""Check fixture downloads without contacting an external RPC endpoint."""

import json
import os
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest

SCRIPT = Path(__file__).resolve().parents[1] / "get_block_fixtures.sh"
STUB = r'''
import json, os, sys
from pathlib import Path
args = sys.argv[1:]
request = json.loads(args[args.index("--data") + 1])
root = Path(os.environ["FIXTURE_TEST_DIR"])
with (root / "calls.jsonl").open("a") as stream:
    stream.write(json.dumps(request) + "\n")
if os.environ.get("FIXTURE_FAILURE") == request["method"]:
    print(json.dumps({"error": {"code": -32601, "message": "Unavailable"}}))
    sys.exit(0)
number = request["params"][0]
result = [] if request["method"] == "debug_traceBlockByNumber" else {
    "number": number, "hash": "0x" + format(int(number, 16) + 1, "064x"), "transactions": []}
print(json.dumps({"jsonrpc": "2.0", "id": 1, "result": result}))
'''


class BlockFixtureTests(unittest.TestCase):
    def setUp(self):
        temporary = tempfile.TemporaryDirectory(prefix="block-fixtures-")
        self.addCleanup(temporary.cleanup)
        self.root = Path(temporary.name)
        self.output = self.root / "output with spaces"
        binary = self.root / "curl"
        binary.write_text("#!" + sys.executable + "\n" + STUB)
        binary.chmod(0o700)
        self.env = dict(os.environ, FIXTURE_TEST_DIR=str(self.root),
                        PATH=str(self.root) + os.pathsep + os.environ["PATH"])

    def run_script(self, block="0", success=True):
        result = subprocess.run(["sh", str(SCRIPT), "http://rpc.invalid", block, str(self.output)],
                                env=self.env, cwd=self.root, capture_output=True, text=True)
        self.assertEqual(result.returncode == 0, success, result.stderr)
        return result

    def test_genesis_has_256_zero_preceding_hashes(self):
        self.run_script()
        hashes = json.loads((self.output / "block_hashes.json").read_text())
        self.assertEqual(hashes, ["0x" + "00" * 32] * 256)
        self.assertEqual(json.loads((self.output / "prestate.json").read_text()), [])
        self.assertFalse(list(self.output.glob("fixtures-attempt-*")))

    def test_early_block_preserves_chronological_hash_order(self):
        self.run_script("002")
        hashes = json.loads((self.output / "block_hashes.json").read_text())
        self.assertEqual(len(hashes), 256)
        self.assertEqual(hashes[-2:], ["0x" + format(value, "064x") for value in (1, 2)])
        calls = [json.loads(line) for line in (self.root / "calls.jsonl").read_text().splitlines()]
        self.assertTrue(all(call["method"] in ("eth_getBlockByNumber", "debug_traceBlockByNumber") for call in calls))

    def test_rpc_error_never_publishes_partial_files(self):
        self.env["FIXTURE_FAILURE"] = "debug_traceBlockByNumber"
        self.run_script(success=False)
        self.assertFalse((self.output / "block.json").exists())
        self.assertTrue(list(self.output.glob("fixtures-attempt-*/block.json")))
        self.assertFalse((self.output / ".block-fixtures.lock.d").exists())

    def test_existing_fixtures_are_preserved(self):
        self.run_script()
        original = (self.output / "block.json").read_bytes()
        self.assertIn("Existing fixture is preserved", self.run_script(success=False).stderr)
        self.assertEqual((self.output / "block.json").read_bytes(), original)

    def test_existing_lock_stops_download(self):
        self.output.mkdir()
        (self.output / ".block-fixtures.lock.d").mkdir()
        self.assertIn("already locked", self.run_script(success=False).stderr)
        self.assertFalse((self.root / "calls.jsonl").exists())

    def test_invalid_block_numbers_stop_before_rpc(self):
        for block in ("-1", "0x1", "1;echo", "9007199254740992"):
            with self.subTest(block=block):
                self.run_script(block, success=False)
        self.assertFalse((self.root / "calls.jsonl").exists())


if __name__ == "__main__":
    unittest.main()
