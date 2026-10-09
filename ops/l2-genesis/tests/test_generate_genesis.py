"""Exercise the existing shell scripts with temporary files and simulated RPC/Go commands."""

import json
import os
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest

SCRIPT_DIR = Path(__file__).resolve().parents[1]
SCRIPT = SCRIPT_DIR / "devnet-l2genesis.sh"
DEPLOYED_NAMES = (
    "Proxy__L1CrossDomainMessenger", "Proxy__Rollup", "Proxy__L1GatewayRouter",
    "Proxy__L1StandardERC20Gateway", "Proxy__L1CustomERC20Gateway", "Proxy__L1ReverseCustomGateway",
    "Proxy__L1ETHGateway", "Proxy__L1ERC721Gateway", "Proxy__L1ERC1155Gateway",
    "Proxy__L1WETHGateway", "Impl__WETH", "Proxy__L1WithdrawLockERC20Gateway", "Proxy__L1Staking",
)
PUBLISHED_FILES = ("deploy-config.json", "genesis-l2.json", "rollup.json", "genesis-batch-header.json", "deployment-config.json")
NETWORKS = ("devnet", "qanet", "testnet", "holesky", "hoodi", "mainnet")
EMPTY_BLOB_HASH = "010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014"

STUB = r'''
import json, os, sys
from pathlib import Path
root = Path(os.environ["GENESIS_TEST_DIR"])
state = json.loads((root / "state.json").read_text())
args = sys.argv[1:]
name = Path(sys.argv[0]).name
with (root / (name + ".calls")).open("a") as stream:
    stream.write(json.dumps(args) + "\n")
if name == "curl":
    request = json.loads(args[args.index("--data") + 1])
    method = request["method"]
    result = {"eth_chainId": state.get("chainId", "0x384"),
              "eth_blockNumber": state.get("blockNumber", "0x100"),
              "eth_getBlockByNumber": {"hash": state.get("genesisHash", "0x" + "11" * 32)},
              "eth_getCode": state.get("code", "0x60006000")}[method]
    print(json.dumps({"jsonrpc": "2.0", "id": 1, "result": result}))
    sys.exit(0)
if state.get("goFailure"):
    print("error calling " + args[args.index("--l1-rpc") + 1])
    sys.exit(1)
config_path = Path(args[args.index("--deploy-config") + 1])
assert config_path.is_absolute()
config = json.loads(config_path.read_text())
assert config["l1StakingProxy"].lower() == state.get("stakingAddress", "0x000000000000000000000000000000000000000d").lower()
assert "BLOCK_SIGNER_PRIVATE_KEY" not in config and "l2StakingPks" not in config
assert "l2GenesisBlockTimestamp" not in config
assert str(Path.cwd()) == os.environ["GENESIS_SCRIPT_DIR"]
header = bytearray(257)
header[0] = 2
header[25:57] = bytes.fromhex("55" * 32)
header[57:89] = bytes.fromhex("010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014")
header[121:153] = bytes.fromhex("22" * 32)
encoded = "0x" + header.hex()
genesis = {"config": {"chainId": 53077}, "alloc": {"123": {"balance": "0x0"}}, "number": "0x0", "timestamp": "0x1234", "gasLimit": "0x1c9c380"}
rollup = {"l1_chain_id": 900, "l2_chain_id": 53077,
          "l2_genesis_state_root": "0x" + ("44" if state.get("badRoot") else "22") * 32,
          "withdraw_root": "0x" + "00" * 32, "genesis_batch_header": encoded,
          "genesis": {"l2_time": 1 if state.get("badTimestamp") else 4660,
                      "system_config": {"gasLimit": 0 if state.get("badGasLimit") else 30000000},
                      "l1": {"hash": "0x" + "11" * 32, "number": 0},
                      "l2": {"hash": "0x" + "33" * 32, "number": 0}}}
for option, value in (("--outfile.l2", genesis), ("--outfile.rollup", rollup), ("--outfile.genbatchheader", encoded)):
    Path(args[args.index(option) + 1]).write_text(json.dumps(value))
print("generated")
'''


class GenesisShellTests(unittest.TestCase):
    def setUp(self):
        self.temp = tempfile.TemporaryDirectory(prefix="genesis-test-")
        self.addCleanup(self.temp.cleanup)
        self.directory = Path(self.temp.name)
        self.config = self.directory / "input.json"
        self.deployment = self.directory / "l1.json"
        self.output = self.directory / "outputs"
        self.config_data = {"l1ChainID": 900, "l2ChainID": 53077, "l1StartingBlockTag": "earliest",
                            "BLOCK_SIGNER_PRIVATE_KEY": "unused-test-placeholder", "l2StakingPks": ["unused"]}
        self.records = [{"name": name, "address": "0x" + format(index, "040x"), "number": 1}
                        for index, name in enumerate(DEPLOYED_NAMES, 1)]
        self.records.append({"name": "Proxy__Submitter", "address": "0x" + "ab" * 20, "number": 1})
        self.write(self.config, self.config_data)
        self.write(self.deployment, self.records)
        self.set_state()
        bin_dir = self.directory / "bin"
        bin_dir.mkdir()
        for name in ("curl", "go"):
            path = bin_dir / name
            path.write_text("#!" + sys.executable + "\n" + STUB)
            path.chmod(0o700)
        self.env = dict(os.environ, GENESIS_TEST_DIR=str(self.directory), GENESIS_SCRIPT_DIR=str(SCRIPT_DIR),
                        PATH=str(bin_dir) + os.pathsep + os.environ["PATH"])
        self.rpc_url = "http://rpc.invalid"

    @staticmethod
    def write(path, value):
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_text(json.dumps(value))

    def set_state(self, **state):
        self.write(self.directory / "state.json", state)

    def calls(self, name):
        path = self.directory / (name + ".calls")
        return len(path.read_text().splitlines()) if path.exists() else 0

    def generate(self, *options, success=True, network="devnet"):
        result = subprocess.run(
            ["/bin/sh", str(SCRIPT_DIR / (network + "-l2genesis.sh")), "--deploy-config", str(self.config),
             "--deployment-file", str(self.deployment), "--output-dir", str(self.output),
             "--l1-rpc", self.rpc_url, *options], cwd=self.directory, env=self.env, capture_output=True, text=True)
        if success:
            self.assertEqual(result.returncode, 0, result.stderr)
        else:
            self.assertNotEqual(result.returncode, 0, result.stdout)
        return result

    def test_supported_networks_publish_validated_artifacts_and_batch_override(self):
        for network in ("devnet", "qanet", "hoodi", "mainnet"):
            with self.subTest(network=network):
                self.output = self.directory / network
                self.generate(network=network)
                marker = json.loads((self.output / "genesis.done").read_text())
                self.assertEqual(marker["network"], network)
                self.assertEqual(set(marker["files"]), set(PUBLISHED_FILES))
                override = json.loads((self.output / "deployment-config.json").read_text())
                self.assertEqual(list(override), ["batchHeader"])
                self.assertEqual(len(bytes.fromhex(override["batchHeader"][2:])), 257)
                self.assertEqual(json.loads(self.config.read_text()), self.config_data)
                self.assertFalse((self.output / "done").exists())
        self.assertEqual(self.calls("go"), 4)

    def test_all_networks_require_original_staking_deployment(self):
        self.write(self.deployment, [row for row in self.records if row["name"] != "Proxy__L1Staking"])
        for network in ("devnet", "qanet", "hoodi", "mainnet"):
            with self.subTest(network=network):
                result = self.generate(network=network, success=False)
                self.assertIn("Missing deployment record: Proxy__L1Staking", result.stderr)
        self.assertEqual(self.calls("curl"), 0)
        self.assertEqual(self.calls("go"), 0)

    def test_staking_placeholder_cannot_be_used_even_with_a_deployment_record(self):
        staking = next(row for row in self.records if row["name"] == "Proxy__L1Staking")
        staking["address"] = "0x000000000000000000000000000000000000dEaD"
        self.write(self.deployment, self.records)
        self.assertIn("not a placeholder", self.generate(success=False).stderr)
        self.assertEqual(self.calls("curl"), 0)
        self.assertEqual(self.calls("go"), 0)

    def test_unconfirmed_deployment_height_cannot_generate_artifacts(self):
        self.set_state(blockNumber="0x0")
        self.assertIn("not confirmed", self.generate(success=False).stderr)
        self.assertEqual(self.calls("go"), 0)
        self.assertFalse((self.output / "genesis.done").exists())

    def test_historical_configuration_stops_before_rpc_with_actionable_error(self):
        for network in ("mainnet", "hoodi"):
            with self.subTest(network=network):
                self.write(self.config, {**self.config_data, "l2SequencerAddresses": []})
                result = self.generate(network=network, success=False)
                self.assertIn("supply --deploy-config", result.stderr)
        self.assertEqual(self.calls("curl"), 0)

    def test_historical_networks_cannot_generate_or_overwrite_with_custom_inputs(self):
        for network in ("testnet", "holesky"):
            for options in ((), ("--overwrite",), ("--network", "devnet")):
                with self.subTest(network=network, options=options):
                    result = self.generate(*options, network=network, success=False)
                    self.assertIn("new genesis generation is disabled", result.stderr)
                    self.assertFalse(self.output.exists())
        self.assertEqual(self.calls("curl"), 0)
        self.assertEqual(self.calls("go"), 0)

    def test_historical_network_rejection_preserves_existing_data(self):
        self.output.mkdir()
        marker = self.output / "genesis.done"
        marker.write_text("preserved historical marker")
        self.generate("--overwrite", network="holesky", success=False)
        self.assertEqual(list(self.output.iterdir()), [marker])
        self.assertEqual(marker.read_text(), "preserved historical marker")

    def test_missing_deployment_stops_before_rpc_or_go(self):
        self.write(self.deployment, self.records[1:])
        self.assertIn("Missing deployment record", self.generate(success=False).stderr)
        self.assertEqual(self.calls("curl"), 0)
        self.assertEqual(self.calls("go"), 0)

    def test_submitter_cannot_be_the_retired_staking_address(self):
        self.config_data["l1StakingProxy"] = self.records[-1]["address"]
        self.write(self.config, self.config_data)
        staking = next(row for row in self.records if row["name"] == "Proxy__L1Staking")
        staking["address"] = self.records[-1]["address"]
        self.write(self.deployment, self.records)
        self.assertIn("distinct from Proxy__Submitter", self.generate(success=False).stderr)
        self.assertEqual(self.calls("curl"), 0)

    def test_wrong_chain_or_empty_code_stops_before_go(self):
        for state in ({"chainId": "0x1"}, {"code": "0x"}):
            with self.subTest(state=state):
                self.set_state(**state)
                self.generate(success=False)
                self.assertFalse((self.output / "genesis.done").exists())
                self.assertEqual(self.calls("go"), 0)

    def test_failed_go_keeps_diagnostics_and_redacts_rpc(self):
        self.rpc_url = "https://user:secret@rpc.invalid/key?token=secret"
        self.set_state(goFailure=True)
        self.assertIn("exit code 1", self.generate(success=False).stderr)
        attempt = next(self.output.glob("genesis-attempt-*"))
        self.assertTrue((attempt / "deploy-config.json").exists())
        self.assertTrue((attempt / "l1-deployments.json").exists())
        log = (attempt / "genesis.log").read_text()
        self.assertIn("<L1_RPC>", log)
        self.assertNotIn("secret", log)
        self.assertFalse((attempt / "genesis.raw.log").exists())
        self.assertFalse((self.output / "genesis.done").exists())

    def test_invalid_outputs_are_not_published(self):
        self.set_state(badRoot=True)
        self.assertIn("postStateRoot", self.generate(success=False).stderr)
        self.assertFalse((self.output / "genesis.done").exists())
        self.assertFalse((self.output / "genesis-l2.json").exists())
        self.assertTrue(list(self.output.glob("genesis-attempt-*/genesis-l2.json")))

    def test_rollup_metadata_must_match_the_generated_genesis(self):
        for state in ({"badTimestamp": True}, {"badGasLimit": True}):
            with self.subTest(state=state):
                self.set_state(**state)
                self.assertIn("timestamp or gas limit differs", self.generate(success=False).stderr)
                self.assertFalse((self.output / "genesis.done").exists())
                self.assertFalse((self.output / "genesis-l2.json").exists())

    def test_existing_outputs_require_explicit_overwrite(self):
        self.generate()
        self.assertIn("--overwrite", self.generate(success=False).stderr)
        self.assertEqual(self.calls("go"), 1)
        self.assertTrue((self.output / "genesis.done").exists())

    def test_exclusive_lock_preserves_existing_marker_for_all_modes(self):
        self.generate()
        marker = (self.output / "genesis.done").read_bytes()
        calls = self.calls("curl")
        lock = self.output / ".genesis.lock.d"
        lock.mkdir()
        for options in ((), ("--overwrite",), ("--verify-existing",)):
            with self.subTest(options=options):
                self.assertIn("already locked", self.generate(*options, success=False).stderr)
                self.assertEqual((self.output / "genesis.done").read_bytes(), marker)
        self.assertEqual(self.calls("curl"), calls)
        self.assertEqual(self.calls("go"), 1)

    def test_failed_generation_releases_lock_for_retry(self):
        self.set_state(goFailure=True)
        self.generate(success=False)
        self.assertFalse((self.output / ".genesis.lock.d").exists())
        self.set_state()
        self.generate()

    def test_failed_overwrite_invalidates_only_genesis_marker(self):
        self.generate()
        (self.output / "done").write_text("network deployment marker")
        self.set_state(goFailure=True)
        self.generate("--overwrite", success=False)
        self.assertFalse((self.output / "genesis.done").exists())
        self.assertTrue((self.output / "done").exists())
        self.assertTrue((self.output / "genesis-l2.json").exists())

    def test_verification_allows_appended_implementation_records(self):
        self.generate()
        self.write(self.deployment, [*self.records, {"name": "Impl__Submitter", "address": "0x" + "cd" * 20,
                                                    "number": 0, "pending": True}])
        self.generate("--verify-existing")
        self.assertEqual(self.calls("go"), 1)

    def test_pending_required_deployment_is_rejected(self):
        self.records[0].update(number=0, pending=True)
        self.write(self.deployment, self.records)
        self.assertIn("still pending", self.generate(success=False).stderr)
        self.assertEqual(self.calls("curl"), 0)

    def test_verification_rejects_changed_config_artifacts_and_chain_identity(self):
        self.generate()
        self.write(self.config, {**self.config_data, "l2ChainID": 2})
        self.assertIn("genesis.done", self.generate("--verify-existing", success=False).stderr)
        self.write(self.config, self.config_data)
        self.set_state(genesisHash="0x" + "ff" * 32)
        self.assertIn("genesis.done", self.generate("--verify-existing", success=False).stderr)
        self.set_state()
        (self.output / "genesis-l2.json").write_text("{}")
        self.assertIn("genesis.done", self.generate("--verify-existing", success=False).stderr)
        self.assertEqual(self.calls("go"), 1)

    def test_verification_rejects_missing_artifact(self):
        self.generate()
        (self.output / "rollup.json").unlink()
        self.assertIn("Missing artifact: rollup.json", self.generate("--verify-existing", success=False).stderr)
        self.assertEqual(self.calls("go"), 1)

    def test_shared_source_and_output_config_remains_verifiable(self):
        self.config = self.output / "deploy-config.json"
        self.write(self.config, self.config_data)
        self.generate()
        self.generate("--verify-existing")
        self.assertEqual(self.calls("go"), 1)

    def test_legacy_field_cleanup_preserves_original_artifacts_and_verification(self):
        legacy_config = {**self.config_data, "maxTxPerBlock": 100,
                         "BLOCK_SIGNER_ADDRESS": "0x" + "ab" * 20,
                         "morphTokenName": "Legacy token", "useMPT": True}
        self.write(self.config, legacy_config)
        self.generate()
        saved = {name: (self.output / name).read_bytes()
                 for name in (*PUBLISHED_FILES, "genesis.done")}

        self.write(self.config, self.config_data)
        result = self.generate("--verify-existing", success=False)
        self.assertIn("differ from genesis.done", result.stderr)
        self.assertEqual(self.calls("go"), 1)

        # The saved normalized input remains usable for read-only verification.
        source = self.config
        self.config = self.output / "deploy-config.json"
        self.generate("--verify-existing")
        self.config = source
        self.write(self.config, legacy_config)
        self.generate("--verify-existing")
        self.assertEqual(self.calls("go"), 1)
        self.assertEqual(saved, {name: (self.output / name).read_bytes() for name in saved})

    def test_wrappers_work_outside_repository_and_require_explicit_rpc(self):
        env = dict(self.env)
        for variable in ("L1_RPC_URL", "QA_RPC_URL", "MAINNET_RPC_URL", "HOLESKY_RPC_URL", "HOODI_RPC_URL", "SEPOLIA_RPC_URL"):
            env.pop(variable, None)
        for network in NETWORKS:
            script = SCRIPT_DIR / (network + "-l2genesis.sh")
            result = subprocess.run(["/bin/sh", str(script), "--help"], cwd=self.directory,
                                    env=env, capture_output=True, text=True)
            self.assertEqual(result.returncode, 0, result.stderr)
            self.assertIn("--verify-existing", result.stdout)
            result = subprocess.run(["/bin/sh", str(script)], cwd=self.directory,
                                    env=env, capture_output=True, text=True)
            self.assertNotEqual(result.returncode, 0)
            expected = "new genesis generation is disabled" if network in ("testnet", "holesky") else "explicit HTTP(S) L1 RPC"
            self.assertIn(expected, result.stderr)


if __name__ == "__main__":
    unittest.main()
