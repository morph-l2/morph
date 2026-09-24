"""Exercise deployment ordering and retries through the existing shell entry point."""

import json
import os
from pathlib import Path
import shutil
import subprocess
import sys
import tempfile
import unittest


ROOT = Path(__file__).resolve().parents[3]
SCRIPT = ROOT / "contracts/scripts/localDeploy.sh"

STUB = r'''
import hashlib, json, os, sys
from pathlib import Path

root = Path(os.environ["LOCAL_DEPLOY_TEST_DIR"])
kind = Path(sys.argv[0]).name
args = sys.argv[1:]
event = {"kind": kind, "args": args}
if kind == "hardhat":
    event["config"] = json.loads(Path(os.environ["DEPLOY_CONFIG_OVERRIDE"]).read_text())
with (root / "calls.jsonl").open("a") as stream:
    stream.write(json.dumps(event) + "\n")

def option(name):
    return args[args.index(name) + 1]

if kind == "node":
    sys.stdin.read()
    rpc = os.environ.get("QA_RPC_URL") if os.environ["HARDHAT_NETWORK"] == "qanetl1" else os.environ.get("L1_RPC_URL")
    if rpc == "http://wrong-chain.invalid":
        print("l1ChainID does not match RPC chainId", file=sys.stderr)
        sys.exit(1)
    source = Path(args[-2]).read_bytes()
    legacy = json.loads(Path(args[-1]).read_text())
    if (root / "missing-legacy-code").exists():
        print("Proxy__L1Staking has no code on the selected L1", file=sys.stderr)
        sys.exit(1)
    config = {"l1ChainID": 900, "l2ChainID": 53077, "batchHeader": "", "rollupDelayPeriod": 600,
              "firstSequencerAddress": "0x" + "aa" * 20,
              "submitterOwner": "0x" + "bb" * 20}
    config.update(json.loads(Path(os.environ["DEPLOY_CONFIG_OVERRIDE"]).read_text()))
    print(json.dumps({"network": os.environ["HARDHAT_NETWORK"], "config": config,
                      "deployer": "0x" + "bb" * 20, "owner": config["submitterOwner"],
                      "l1GenesisHash": "0x" + "cc" * 32,
                      "legacyL1StakingDeployment": legacy,
                      "genesisConfigHash": hashlib.sha256(source).hexdigest()}))
elif kind == "hardhat":
    task = args[0]
    if task == "deploy":
        destination = Path(option("--storagepath"))
        records = json.loads(destination.read_text()) if destination.exists() else []
        assert records[0]["name"] == "Proxy__L1Staking"
        names = ("Proxy__L1Sequencer", "Proxy__L1MessageQueueWithGasPriceOracle",
                 "Proxy__Rollup", "Proxy__Submitter")
        for index, name in enumerate(names, 1):
            if not any(record["name"] == name for record in records):
                records.append({"name": name, "address": "0x" + format(index, "040x"), "number": 1})
        destination.write_text(json.dumps(records))
    if task in ("initialize", "register", "verify-deployment"):
        assert event["config"]["batchHeader"] == "0x1234"
    if task == "initialize" and (root / "fail-initialize-once").exists():
        (root / "fail-initialize-once").unlink()
        print("simulated initialization interruption", file=sys.stderr)
        sys.exit(1)
elif kind == "genesis":
    output = Path(option("--output-dir"))
    if "--verify-existing" in args:
        assert (output / "genesis.done").exists()
        assert json.loads((output / "deployment-config.json").read_text())["batchHeader"] == "0x1234"
    elif (root / "fail-genesis").exists():
        print("simulated genesis generation failure", file=sys.stderr)
        sys.exit(1)
    else:
        (output / "deployment-config.json").write_text(json.dumps({"batchHeader": "0x1234"}))
        (output / "genesis-l2.json").write_text("{}")
        (output / "genesis.done").write_text("{}")
else:
    raise AssertionError(kind)
'''


@unittest.skipUnless(shutil.which("jq"), "jq is required by the deployment shell script")
class LocalDeployShellTests(unittest.TestCase):
    def setUp(self):
        temporary = tempfile.TemporaryDirectory(prefix="local-deploy-test-")
        self.addCleanup(temporary.cleanup)
        self.directory = Path(temporary.name)
        self.contracts = self.directory / "contracts"
        self.genesis = self.directory / "ops/l2-genesis"
        self.script = self.contracts / "scripts/localDeploy.sh"
        self.script.parent.mkdir(parents=True)
        shutil.copyfile(SCRIPT, self.script)
        (self.genesis / "deploy-config").mkdir(parents=True)
        self.source = {"l1ChainID": 900, "l2ChainID": 53077,
                       "govBatchBlockInterval": 100, "govBatchTimeout": 10}
        for network in ("devnet", "qanet"):
            (self.genesis / f"deploy-config/{network}-deploy-config.json").write_text(json.dumps(self.source))
            (self.genesis / f"{network}-l2genesis.sh").write_text('#!/bin/sh\nexec genesis "$@"\n')
        bin_directory = self.directory / "bin"
        bin_directory.mkdir()
        for destination in (bin_directory / "node", bin_directory / "genesis",
                            self.contracts / "node_modules/.bin/hardhat"):
            destination.parent.mkdir(parents=True, exist_ok=True)
            destination.write_text("#!" + sys.executable + "\n" + STUB)
            destination.chmod(0o700)
        self.env = dict(os.environ, LOCAL_DEPLOY_TEST_DIR=str(self.directory),
                        PATH=str(bin_directory) + os.pathsep + os.environ["PATH"])
        for name in ("L1_RPC_URL", "QA_RPC_URL", "DEPLOY_CONFIG_OVERRIDE", "firstSequencerAddress",
                     "SUBMITTER_OWNER_PRIVATE_KEY", "QA_ROLLUP_DELAY_PERIOD", "batchSubmitterPks",
                     "TX_SUBMITTER_BATCH_BLOCK_INTERVAL", "TX_SUBMITTER_BATCH_TIMEOUT", "LEGACY_L1_DEPLOYMENT_FILE"):
            self.env.pop(name, None)
        self.env["DEPLOYER_PRIVATE_KEY"] = "unused-test-placeholder"
        self.output = self.directory / "output"
        self.legacy_record = {"name": "Proxy__L1Staking", "address": "0x" + "dd" * 20, "number": 1}
        self.legacy_file = self.directory / "legacy-l1-deployments.json"
        self.legacy_file.write_text(json.dumps([self.legacy_record]))

    def run_deploy(self, *options, network="devnet", rpc=True, batch=True, legacy=True, success=True):
        command = ["/bin/sh", str(self.script), "--network", network, "--output-dir", str(self.output)]
        if rpc:
            command.extend(["--l1-rpc", "http://rpc.invalid"])
        if batch:
            command.extend(["--batch-block-interval", "100", "--batch-timeout", "10"])
        if legacy:
            command.extend(["--legacy-l1-deployment-file", str(self.legacy_file)])
        result = subprocess.run(command + list(options), cwd=self.directory, env=self.env,
                                capture_output=True, text=True)
        if success:
            self.assertEqual(result.returncode, 0, result.stderr)
        else:
            self.assertNotEqual(result.returncode, 0, result.stdout)
        return result

    def events(self):
        filename = self.directory / "calls.jsonl"
        return [json.loads(line) for line in filename.read_text().splitlines()] if filename.exists() else []

    def clear_events(self):
        (self.directory / "calls.jsonl").unlink(missing_ok=True)

    def tasks(self):
        return [event["args"][0] for event in self.events() if event["kind"] == "hardhat"]

    def test_both_networks_generate_before_initialize_and_publish_public_parameters(self):
        for network, hardhat_network in (("devnet", "l1"), ("qanet", "qanetl1")):
            with self.subTest(network=network):
                self.output = self.directory / network
                self.clear_events()
                self.run_deploy(network=network)
                events = self.events()
                self.assertEqual([event["kind"] for event in events],
                                 ["node", "hardhat", "genesis", "hardhat", "hardhat", "hardhat"])
                self.assertEqual(self.tasks(), ["deploy", "initialize", "register", "verify-deployment"])
                for event in events:
                    if event["kind"] == "hardhat":
                        self.assertEqual(event["args"][event["args"].index("--network") + 1], hardhat_network)
                self.assertTrue((self.output / "done").exists())
                runtime = (self.output / "runtime.env").read_text()
                self.assertIn("MORPH_NODE_SYNC_DEPOSIT_CONTRACT_ADDRESS=0x" + format(2, "040x"), runtime)
                self.assertIn("TX_SUBMITTER_BATCH_TIMEOUT=10", runtime)
                self.assertNotIn("PRIVATE_KEY", runtime)
                self.assertFalse((self.output / ".deployment.lock.d").exists())

    def test_genesis_failure_stops_initialization_and_preserves_records(self):
        (self.directory / "fail-genesis").touch()
        self.assertIn("genesis generation failure", self.run_deploy(success=False).stderr)
        self.assertEqual(self.tasks(), ["deploy"])
        self.assertTrue((self.output / "devnetL1.json").exists())
        self.assertTrue((self.output / "deployment-identity.json").exists())
        self.assertFalse((self.output / "done").exists())
        self.assertFalse((self.output / ".deployment.lock.d").exists())

    def test_interrupted_initialization_reuses_verified_genesis(self):
        (self.directory / "fail-initialize-once").touch()
        self.run_deploy(success=False)
        original_header = (self.output / "deployment-config.json").read_bytes()
        original_records = (self.output / "devnetL1.json").read_bytes()
        self.assertFalse((self.output / "done").exists())
        self.clear_events()
        self.run_deploy()
        genesis_calls = [event for event in self.events() if event["kind"] == "genesis"]
        self.assertEqual(len(genesis_calls), 1)
        self.assertIn("--verify-existing", genesis_calls[0]["args"])
        self.assertEqual(self.tasks(), ["deploy", "initialize", "register", "verify-deployment"])
        self.assertEqual((self.output / "deployment-config.json").read_bytes(), original_header)
        self.assertEqual((self.output / "devnetL1.json").read_bytes(), original_records)
        self.assertTrue((self.output / "done").exists())

    def test_completed_repeat_runs_only_runtime_verification(self):
        self.run_deploy()
        self.clear_events()
        self.run_deploy()
        self.assertEqual(self.tasks(), ["verify-deployment"])
        task = next(event for event in self.events() if event["kind"] == "hardhat")
        self.assertIn("--runtime", task["args"])
        genesis = next(event for event in self.events() if event["kind"] == "genesis")
        self.assertIn("--verify-existing", genesis["args"])

    def test_changed_public_identity_stops_before_genesis_or_hardhat(self):
        self.run_deploy()
        identity = (self.output / "deployment-identity.json").read_bytes()
        self.clear_events()
        self.source["govBatchTimeout"] = 20
        (self.genesis / "deploy-config/devnet-deploy-config.json").write_text(json.dumps(self.source))
        self.assertIn("configuration differ", self.run_deploy(success=False).stderr)
        self.assertEqual([event["kind"] for event in self.events()], ["node"])
        self.assertEqual((self.output / "deployment-identity.json").read_bytes(), identity)
        self.assertFalse((self.output / ".deployment.lock.d").exists())

    def test_removed_ignored_key_requires_exact_original_source_for_resume(self):
        for network in ("devnet", "qanet"):
            with self.subTest(network=network):
                self.output = self.directory / network
                source_path = self.genesis / f"deploy-config/{network}-deploy-config.json"
                legacy_source = dict(self.source, maxTxPerBlock=100)
                original_bytes = (json.dumps(legacy_source, indent=2) + "\n").encode()
                source_path.write_bytes(original_bytes)
                self.run_deploy(network=network)
                artifacts = {path.name: path.read_bytes() for path in self.output.iterdir() if path.is_file()}

                source_path.write_text(json.dumps(self.source))
                self.clear_events()
                result = self.run_deploy(network=network, success=False)
                self.assertIn("configuration differ", result.stderr)
                self.assertEqual([event["kind"] for event in self.events()], ["node"])
                for name, data in artifacts.items():
                    self.assertEqual((self.output / name).read_bytes(), data)

                original_path = self.directory / f"{network}-original-source.json"
                original_path.write_text(json.dumps(legacy_source))
                self.clear_events()
                self.assertIn("configuration differ", self.run_deploy(
                    "--deploy-config", str(original_path), network=network, success=False).stderr)
                self.assertEqual(self.tasks(), [])

                original_path.write_bytes(original_bytes)
                self.clear_events()
                self.run_deploy("--deploy-config", str(original_path), network=network)
                self.assertEqual(self.tasks(), ["verify-deployment"])
                genesis = next(event for event in self.events() if event["kind"] == "genesis")
                self.assertIn("--verify-existing", genesis["args"])
                for name, data in artifacts.items():
                    self.assertEqual((self.output / name).read_bytes(), data)

    def test_explicit_batch_parameters_are_independent_of_genesis_gov_values(self):
        for network in ("devnet", "qanet"):
            with self.subTest(network=network):
                self.output = self.directory / network
                source = dict(self.source, govBatchBlockInterval=777, govBatchTimeout=888)
                (self.genesis / f"deploy-config/{network}-deploy-config.json").write_text(json.dumps(source))
                self.run_deploy("--batch-block-interval", "7", "--batch-timeout", "90", network=network)
                identity = json.loads((self.output / "deployment-identity.json").read_text())
                self.assertEqual(identity["batchParameters"], {"batchBlockInterval": "7", "batchTimeout": "90"})
                runtime = (self.output / "runtime.env").read_text()
                self.assertIn("TX_SUBMITTER_BATCH_BLOCK_INTERVAL=7\n", runtime)
                self.assertIn("TX_SUBMITTER_BATCH_TIMEOUT=90\n", runtime)

    def test_batch_parameters_accept_environment_and_uint64_without_precision_loss(self):
        self.env["TX_SUBMITTER_BATCH_BLOCK_INTERVAL"] = "18446744073709551615"
        self.env["TX_SUBMITTER_BATCH_TIMEOUT"] = "0"
        self.run_deploy(batch=False)
        runtime = (self.output / "runtime.env").read_text()
        self.assertIn("TX_SUBMITTER_BATCH_BLOCK_INTERVAL=18446744073709551615\n", runtime)
        self.assertIn("TX_SUBMITTER_BATCH_TIMEOUT=0\n", runtime)
        self.output = self.directory / "zero-interval"
        self.run_deploy("--batch-block-interval", "000", "--batch-timeout", "001")
        parameters = json.loads((self.output / "deployment-identity.json").read_text())["batchParameters"]
        self.assertEqual(parameters, {"batchBlockInterval": "0", "batchTimeout": "1"})

    def test_invalid_or_missing_batch_parameters_stop_before_preflight(self):
        invalid = [(), ("--batch-block-interval", "1"), ("--batch-timeout", "1")]
        invalid += [("--batch-block-interval", interval, "--batch-timeout", timeout) for interval, timeout in (
            ("0", "0"), ("-1", "1"), ("1", "-1"), ("1.5", "1"), ("1", "1e2"),
            ("18446744073709551616", "1"), ("1", "18446744073709551616"), ("", "1"),
            ("1\nOTHER=2", "1"), ("1", "1\n"))]
        for network in ("devnet", "qanet"):
            for options in invalid:
                with self.subTest(network=network, options=options):
                    self.clear_events()
                    self.run_deploy(*options, network=network, batch=False, success=False)
                    self.assertEqual(self.events(), [])
                    self.assertFalse(self.output.exists())

    def test_changed_batch_parameters_preserve_completed_identity(self):
        self.run_deploy()
        original = {path.name: path.read_bytes() for path in self.output.iterdir() if path.is_file()}
        self.clear_events()
        result = self.run_deploy("--batch-timeout", "11", success=False)
        self.assertIn("explicit batch parameters", result.stderr)
        self.assertEqual([event["kind"] for event in self.events()], ["node"])
        for name, data in original.items():
            self.assertEqual((self.output / name).read_bytes(), data)
        self.clear_events()
        self.run_deploy()
        self.assertEqual(self.tasks(), ["verify-deployment"])

    def test_legacy_identity_without_batch_parameters_is_not_migrated(self):
        self.run_deploy()
        identity_path = self.output / "deployment-identity.json"
        identity = json.loads(identity_path.read_text())
        del identity["batchParameters"]
        original = json.dumps(identity)
        identity_path.write_text(original)
        self.clear_events()
        result = self.run_deploy(success=False)
        self.assertIn("no recorded batch parameters", result.stderr)
        self.assertIn("original scripts and inputs", result.stderr)
        self.assertEqual([event["kind"] for event in self.events()], ["node"])
        self.assertEqual(identity_path.read_text(), original)

    def test_legacy_record_is_persisted_and_resume_uses_output_records(self):
        self.env["LEGACY_L1_DEPLOYMENT_FILE"] = str(self.legacy_file)
        self.run_deploy(legacy=False)
        identity = json.loads((self.output / "deployment-identity.json").read_text())
        self.assertEqual(identity["legacyL1StakingDeployment"], self.legacy_record)
        records_path = self.output / "devnetL1.json"
        records = json.loads(records_path.read_text())
        self.assertEqual(records[0], self.legacy_record)
        self.assertEqual(len(records), 5)
        original = records_path.read_bytes()
        self.env.pop("LEGACY_L1_DEPLOYMENT_FILE")
        self.legacy_file.unlink()
        self.clear_events()
        self.run_deploy(legacy=False)
        self.assertEqual(self.tasks(), ["verify-deployment"])
        self.assertEqual(records_path.read_bytes(), original)

    def test_missing_legacy_record_stops_both_networks_before_preflight(self):
        for network in ("devnet", "qanet"):
            with self.subTest(network=network):
                self.output = self.directory / network
                result = self.run_deploy(network=network, legacy=False, success=False)
                self.assertIn("existing confirmed Proxy__L1Staking record", result.stderr)
                self.assertEqual(self.events(), [])
                self.assertFalse((self.output / "deployment-identity.json").exists())
                self.assertFalse((self.output / f"{network}L1.json").exists())
                self.assertFalse((self.output / ".deployment.lock.d").exists())

    def test_invalid_legacy_records_stop_before_preflight_and_do_not_seed_output(self):
        invalid = [[], {}, [self.legacy_record, self.legacy_record],
                   [dict(self.legacy_record, name="Proxy__Submitter")]]
        invalid += [[dict(self.legacy_record, **changes)] for changes in (
            {"address": "0x" + "0" * 40}, {"address": "0x" + "0" * 36 + "dEaD"},
            {"address": "0x" + "dd" * 20 + "\n"}, {"address": "invalid"},
            {"number": -1}, {"number": 1.5}, {"number": "1"}, {"number": 9007199254740992},
            {"pending": True}, {"pending": "false"}, {"pending": None})]
        invalid.append([self.legacy_record, dict(self.legacy_record, name="Proxy__Submitter")])
        for records in invalid:
            with self.subTest(records=records):
                self.legacy_file.write_text(json.dumps(records))
                self.clear_events()
                result = self.run_deploy(success=False)
                self.assertIn("Invalid legacy L1 deployment record", result.stderr)
                self.assertEqual(self.events(), [])
                self.assertFalse((self.output / "deployment-identity.json").exists())
                self.assertFalse((self.output / "devnetL1.json").exists())

    def test_missing_legacy_code_does_not_persist_identity_or_records(self):
        (self.directory / "missing-legacy-code").touch()
        result = self.run_deploy(success=False)
        self.assertIn("no code on the selected L1", result.stderr)
        self.assertEqual([event["kind"] for event in self.events()], ["node"])
        self.assertFalse((self.output / "deployment-identity.json").exists())
        self.assertFalse((self.output / "devnetL1.json").exists())

    def test_changed_legacy_record_cannot_overwrite_saved_records(self):
        self.run_deploy()
        original = {path.name: path.read_bytes() for path in self.output.iterdir() if path.is_file()}
        self.legacy_file.write_text(json.dumps([dict(self.legacy_record, address="0x" + "ee" * 20)]))
        self.clear_events()
        result = self.run_deploy(success=False)
        self.assertIn("differs from saved deployment records", result.stderr)
        self.assertEqual(self.events(), [])
        for name, data in original.items():
            self.assertEqual((self.output / name).read_bytes(), data)

    def test_changed_legacy_identity_cannot_seed_missing_output_records(self):
        self.run_deploy()
        identity = (self.output / "deployment-identity.json").read_bytes()
        records = self.output / "devnetL1.json"
        records.unlink()
        self.legacy_file.write_text(json.dumps([dict(self.legacy_record, number=2)]))
        self.clear_events()
        result = self.run_deploy(success=False)
        self.assertIn("configuration differ", result.stderr)
        self.assertEqual([event["kind"] for event in self.events()], ["node"])
        self.assertEqual((self.output / "deployment-identity.json").read_bytes(), identity)
        self.assertFalse(records.exists())

    def test_empty_saved_records_are_not_seeded_or_deployed_again(self):
        self.run_deploy()
        records = self.output / "devnetL1.json"
        records.write_text("[]")
        self.clear_events()
        result = self.run_deploy(success=False)
        self.assertIn("Saved deployment records must be a nonempty array", result.stderr)
        self.assertEqual(self.events(), [])
        self.assertEqual(records.read_text(), "[]")

    def test_identity_without_legacy_record_is_not_migrated(self):
        self.run_deploy()
        identity_path = self.output / "deployment-identity.json"
        identity = json.loads(identity_path.read_text())
        del identity["legacyL1StakingDeployment"]
        original = json.dumps(identity)
        identity_path.write_text(original)
        self.clear_events()
        result = self.run_deploy(success=False)
        self.assertIn("no recorded legacy L1Staking deployment", result.stderr)
        self.assertEqual(self.tasks(), [])
        self.assertEqual(identity_path.read_text(), original)

    def test_missing_or_wrong_rpc_stops_before_transaction_tasks(self):
        for network in ("devnet", "qanet"):
            with self.subTest(network=network):
                self.assertIn("Set L1_RPC_URL", self.run_deploy(network=network, rpc=False, success=False).stderr)
                self.assertEqual(self.events(), [])
        self.assertIn("does not match RPC", self.run_deploy("--l1-rpc", "http://wrong-chain.invalid", success=False).stderr)
        self.assertEqual(self.tasks(), [])
        self.assertFalse((self.output / "deployment-identity.json").exists())
        self.assertFalse((self.output / ".deployment.lock.d").exists())

    def test_existing_lock_is_preserved_and_prevents_preflight(self):
        lock = self.output / ".deployment.lock.d"
        lock.mkdir(parents=True)
        (lock / "owner").write_text("another invocation")
        self.assertIn("Deployment is locked", self.run_deploy(success=False).stderr)
        self.assertEqual(self.events(), [])
        self.assertEqual((lock / "owner").read_text(), "another invocation")

    def test_missing_genesis_marker_on_completed_deployment_blocks_transaction_tasks(self):
        self.run_deploy()
        (self.output / "genesis.done").unlink()
        self.clear_events()
        self.assertIn("missing genesis.done", self.run_deploy(success=False).stderr)
        self.assertEqual(self.tasks(), [])
        self.assertTrue((self.output / "done").exists())

    def test_unidentified_existing_deployment_is_not_overwritten(self):
        self.output.mkdir()
        records = self.output / "devnetL1.json"
        records.write_text("original deployment records")
        self.assertIn("no matching identity", self.run_deploy(success=False).stderr)
        self.assertEqual(records.read_text(), "original deployment records")
        self.assertEqual(self.tasks(), [])


if __name__ == "__main__":
    unittest.main()
