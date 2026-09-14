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
    source = Path(args[-1]).read_bytes()
    config = {"l1ChainID": 900, "l2ChainID": 53077, "batchHeader": "", "rollupDelayPeriod": 600,
              "firstSequencerAddress": "0x" + "aa" * 20,
              "submitterOwner": "0x" + "bb" * 20}
    config.update(json.loads(Path(os.environ["DEPLOY_CONFIG_OVERRIDE"]).read_text()))
    print(json.dumps({"network": os.environ["HARDHAT_NETWORK"], "config": config,
                      "deployer": "0x" + "bb" * 20, "owner": config["submitterOwner"],
                      "l1GenesisHash": "0x" + "cc" * 32,
                      "genesisConfigHash": hashlib.sha256(source).hexdigest()}))
elif kind == "hardhat":
    task = args[0]
    if task == "deploy":
        destination = Path(option("--storagepath"))
        if not destination.exists():
            names = ("Proxy__L1Sequencer", "Proxy__L1MessageQueueWithGasPriceOracle",
                     "Proxy__Rollup", "Proxy__Submitter")
            destination.write_text(json.dumps([
                {"name": name, "address": "0x" + format(index, "040x"), "number": 1}
                for index, name in enumerate(names, 1)]))
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
                     "SUBMITTER_OWNER_PRIVATE_KEY", "QA_ROLLUP_DELAY_PERIOD", "batchSubmitterPks"):
            self.env.pop(name, None)
        self.env["DEPLOYER_PRIVATE_KEY"] = "unused-test-placeholder"
        self.output = self.directory / "output"

    def run_deploy(self, *options, network="devnet", rpc=True, success=True):
        command = ["/bin/sh", str(self.script), "--network", network, "--output-dir", str(self.output)]
        if rpc:
            command.extend(["--l1-rpc", "http://rpc.invalid"])
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
