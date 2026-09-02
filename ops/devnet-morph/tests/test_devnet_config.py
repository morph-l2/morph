import importlib
import sys
import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parents[3]
DEVNET_PACKAGE = REPO_ROOT / "ops" / "devnet-morph"
DOCKER_DIR = REPO_ROOT / "ops" / "docker"


class DevnetConfigTest(unittest.TestCase):
    def test_root_dockerignore_excludes_generated_build_outputs(self):
        dockerignore = (REPO_ROOT / ".dockerignore").read_text().splitlines()

        for generated_path in (
            "gas-oracle/app/target/",
            "node/build/",
            "ops/docker/.devnet/",
        ):
            self.assertIn(generated_path, dockerignore)

    def test_node_dockerfile_caches_go_dependencies_before_source_copy(self):
        dockerfile = (DOCKER_DIR / "Dockerfile.l2-node").read_text()

        dependency_layer = dockerfile.index("RUN go mod download")
        source_layer = dockerfile.index("COPY . /morph")

        self.assertLess(dependency_layer, source_layer)
        self.assertIn("COPY go.work go.work.sum /morph/", dockerfile)
        self.assertIn("COPY node/go.mod node/go.sum /morph/node/", dockerfile)

    def test_geth_dockerfile_caches_go_dependencies_before_source_copy(self):
        dockerfile = (DOCKER_DIR / "Dockerfile.l2-geth").read_text()

        dependency_layer = dockerfile.index("RUN go mod download")
        source_layer = dockerfile.index("COPY go-ethereum /go-ethereum")

        self.assertLess(dependency_layer, source_layer)
        self.assertIn("COPY go-ethereum/go.mod go-ethereum/go.sum /go-ethereum/", dockerfile)

    def test_devnet_clean_removes_compose_project_volumes(self):
        makefile = (REPO_ROOT / "Makefile").read_text()

        self.assertIn(
            "DEVNET_CLEAN_COMPOSE_FILES := -f docker-compose-devnet.yml "
            "-f docker-compose-cluster.yml -f docker-compose-reth.yml "
            "-f docker-compose-cluster-reth.yml",
            makefile,
        )
        self.assertIn("docker compose $(DEVNET_CLEAN_COMPOSE_FILES) down --volumes --remove-orphans", makefile)
        self.assertIn("--filter label=com.docker.compose.project=docker", makefile)
        self.assertNotIn("docker_morph_data_0 docker_morph_data_1", makefile)
        self.assertNotIn("devnet-clean-build-reth", makefile)
        self.assertNotIn("devnet-clean-reth", makefile)

    def test_default_compose_includes_layer1_derivation_node(self):
        compose = (DOCKER_DIR / "docker-compose-devnet.yml").read_text()

        self.assertIn("node-1:", compose)
        self.assertIn("MORPH_NODE_L2_ETH_RPC=http://morph-el-1:8545", compose)
        self.assertIn("MORPH_NODE_DERIVATION_VERIFY_MODE=layer1", compose)
        self.assertNotIn("morph-el-2:", compose)
        self.assertNotIn("node-2:", compose)
        self.assertNotIn("morph-el-2:8545", compose)

    def test_tx_submitter_uses_submitter_contract_and_explicit_batch_settings(self):
        compose = (DOCKER_DIR / "docker-compose-devnet.yml").read_text()

        self.assertIn("TX_SUBMITTER_SUBMITTER_ADDRESS=${MORPH_SUBMITTER}", compose)
        self.assertIn("TX_SUBMITTER_BATCH_BLOCK_INTERVAL=${BATCH_BLOCK_INTERVAL}", compose)
        self.assertIn("TX_SUBMITTER_BATCH_TIMEOUT=${BATCH_TIMEOUT}", compose)
        self.assertIn("TX_SUBMITTER_L1_PRIVATE_KEY=${BATCH_SUBMITTER_PRIVATE_KEY}", compose)
        self.assertIn("until (true > /dev/tcp/morph-el-0/8545)", compose)
        for removed_setting in (
            "TX_SUBMITTER_PRIORITY_ROLLUP",
            "TX_SUBMITTER_L1_STAKING_ADDRESS",
            "TX_SUBMITTER_L1_STAKING_DEPLOYED_BLOCKNUM",
        ):
            self.assertNotIn(removed_setting, compose)

        launcher = (DEVNET_PACKAGE / "devnet" / "__init__.py").read_text()
        self.assertIn("addresses['Proxy__Submitter']", launcher)
        self.assertIn("deploy_config['govBatchBlockInterval']", launcher)
        self.assertIn("deploy_config['govBatchTimeout']", launcher)
        self.assertIn("batchSubmitterPks", launcher)
        self.assertIn("args.batch_submitter_private_key", launcher)
        fund_command = "'npx', 'hardhat', 'fund', '--network', 'l1'"
        register_command = "'npx', 'hardhat', 'register', '--network', 'l1'"
        self.assertIn(fund_command, launcher)
        self.assertLess(launcher.index(fund_command), launcher.index(register_command))
        self.assertIn(
            "deploy_config['l1StakingProxy'] = LEGACY_GENESIS_L1_STAKING_PROXY",
            launcher,
        )
        self.assertIn(
            "LEGACY_GENESIS_L1_STAKING_PROXY = '0x000000000000000000000000000000000000dEaD'",
            launcher,
        )
        self.assertNotIn("addresses['Proxy__L1Staking']", launcher)
        self.assertIn("env_data.pop('Proxy__L1Staking', None)", launcher)
        self.assertIn("env_data.pop('MORPH_L1STAKING', None)", launcher)
        self.assertNotIn("Proxy__L1Staking", (DOCKER_DIR / ".env").read_text())
        self.assertNotIn(
            "'layer1-el', 'layer1-cl', 'layer1-vc'], check=False",
            launcher,
        )
        self.assertNotIn(
            "'up', '-d'], check=False",
            launcher,
        )

        deploy_task = (REPO_ROOT / "contracts" / "tasks" / "deploy.ts").read_text()
        register_task = deploy_task[deploy_task.index('task("register")'):]
        self.assertIn('JSON.parse(process.env.batchSubmitterPks || "[]")', register_task)
        self.assertIn("new ethers.Wallet(privateKey).address", register_task)

    def test_cluster_compose_defines_ha_services(self):
        cluster_compose = DOCKER_DIR / "docker-compose-cluster.yml"

        self.assertTrue(cluster_compose.exists())
        compose = cluster_compose.read_text()
        for service in ("ha-el-0:", "ha-el-1:", "ha-el-2:", "ha-node-0:", "ha-node-1:", "ha-node-2:"):
            self.assertIn(service, compose)
        self.assertNotIn("ha-geth", compose)
        self.assertIn("MORPH_NODE_HA_ENABLED=true", compose)
        self.assertIn("MORPH_NODE_HA_BOOTSTRAP=true", compose)
        self.assertIn("MORPH_NODE_HA_JOIN=ha-node-0:9401", compose)
        for index in (0, 1, 2):
            self.assertIn(f"MORPH_NODE_L2_ETH_RPC=http://ha-el-{index}:8545", compose)
            self.assertIn(f"MORPH_NODE_L2_ENGINE_RPC=http://ha-el-{index}:8551", compose)

    def test_cluster_geth_nodes_peer_with_each_other(self):
        """The shared static-nodes.json lists only morph-el-*, so the cluster
        needs its own file or the ha-el nodes never dial each other."""
        cluster_static_nodes = (DOCKER_DIR / "static-nodes-cluster.json").read_text()

        for host in ("ha-el-0", "ha-el-1", "ha-el-2", "morph-el-0", "morph-el-1"):
            self.assertIn(f"@{host}:30303", cluster_static_nodes)

        compose = (DOCKER_DIR / "docker-compose-cluster.yml").read_text()
        # geth only reads the fixed filename, so the mount has to be renamed on
        # the way in.
        self.assertIn(
            '"${PWD}/static-nodes-cluster.json:/db/geth/static-nodes.json"',
            compose,
        )

    def test_reth_compose_configures_deterministic_peering(self):
        compose = (DOCKER_DIR / "docker-compose-reth.yml").read_text()

        # reth cannot read geth's static-nodes.json, so peers are passed as
        # flags, and a fixed key file is what makes the enodes predictable.
        for service, key in (
            ("morph-el-0", "nodekey0"),
            ("morph-el-1", "nodekey1"),
        ):
            self.assertIn(f"{service}:", compose)
            self.assertIn(f'"${{PWD}}/{key}:/p2p-secret.key"', compose)
        self.assertEqual(compose.count("--p2p-secret-key=/p2p-secret.key"), 2)
        self.assertEqual(compose.count("--trusted-peers="), 2)

    def test_reth_compose_leaves_cluster_services_to_the_cluster_reth_file(self):
        """compose starts every service a later -f file introduces, even one no
        earlier file declared. ha-el-* overrides parked in the shared reth file
        therefore came up in the non-cluster devnet as well, missing the
        /genesis.json and /jwt-secret.txt mounts that only
        docker-compose-cluster.yml supplies, and exited with
        "Invalid value '/genesis.json' for --chain"."""
        compose = (DOCKER_DIR / "docker-compose-reth.yml").read_text()

        for service in ("ha-el-0", "ha-el-1", "ha-el-2"):
            self.assertNotIn(f"{service}:", compose)
        for key in ("ha-nodekey0", "ha-nodekey1", "ha-nodekey2"):
            self.assertNotIn(key, compose)

    def test_cluster_reth_compose_overrides_the_ha_execution_clients(self):
        cluster_reth = DOCKER_DIR / "docker-compose-cluster-reth.yml"

        self.assertTrue(cluster_reth.exists())
        compose = cluster_reth.read_text()

        self.assertIn("${MORPH_RETH_IMAGE:-ghcr.io/morph-l2/morph-reth:latest}", compose)
        self.assertIn("${MORPH_RETH_ENTRYPOINT:-/usr/local/bin/morph-reth}", compose)
        for service, key in (
            ("ha-el-0", "ha-nodekey0"),
            ("ha-el-1", "ha-nodekey1"),
            ("ha-el-2", "ha-nodekey2"),
        ):
            self.assertIn(f"{service}:", compose)
            self.assertIn(f'"${{PWD}}/{key}:/p2p-secret.key"', compose)
        self.assertEqual(compose.count("--p2p-secret-key=/p2p-secret.key"), 3)
        self.assertEqual(compose.count("--trusted-peers="), 3)
        # Nothing here may redefine the plain devnet's execution clients, or
        # they would be reconfigured only in cluster mode.
        for service in ("morph-el-0:", "morph-el-1:"):
            self.assertNotIn(f"  {service}", compose)

    def test_execution_client_keys_have_no_trailing_newline(self):
        """reth rejects a key file with a trailing newline ("malformed or
        out-of-range secret key"); geth accepts it either way."""
        for key in ("nodekey0", "nodekey1", "nodekey2",
                    "ha-nodekey0", "ha-nodekey1", "ha-nodekey2"):
            contents = (DOCKER_DIR / key).read_bytes()
            self.assertEqual(len(contents), 64, f"{key} should be 64 hex characters")
            self.assertFalse(contents.endswith(b"\n"), f"{key} must not end with a newline")

    def test_tendermint_peers_are_derived_from_installed_node_keys(self):
        sys.path.insert(0, str(DEVNET_PACKAGE))
        try:
            setup_nodes = importlib.import_module("devnet.setup_nodes")
            importlib.reload(setup_nodes)
        finally:
            sys.path.remove(str(DEVNET_PACKAGE))

        # Anchored on the ID that used to be hardcoded in persistent_peers.
        self.assertEqual(
            setup_nodes.tendermint_node_id(DOCKER_DIR / "node0" / "node_key.json"),
            "93e27ea2306e158a8146d5f44caaab97496797d2",
        )

        # node-1 never starts tendermint and node-2 has no compose service, so
        # neither may appear as a peer.
        self.assertEqual(
            setup_nodes.TENDERMINT_PEERS,
            ("node0", "ha-node0", "ha-node1", "ha-node2"),
        )

        source = (DEVNET_PACKAGE / "devnet" / "setup_nodes.py").read_text()
        # Node IDs come from the keys that end up installed, so the copy step
        # must run before the peer list is built.
        self.assertLess(
            source.index("copy_key_files(docker_dir, devnet_dir)"),
            source.index("persistent_peers = build_persistent_peers(devnet_dir)"),
        )
        self.assertIn('laddr = "tcp://0.0.0.0:26657"', source)

    def test_compose_file_args_can_enable_cluster_mode(self):
        sys.path.insert(0, str(DEVNET_PACKAGE))
        try:
            devnet = importlib.import_module("devnet")
            importlib.reload(devnet)
        finally:
            sys.path.remove(str(DEVNET_PACKAGE))

        self.assertEqual(devnet.compose_file_args("geth"), ["-f", "docker-compose-devnet.yml"])
        # The ha-el-* reth overrides must not reach the non-cluster devnet:
        # compose would start those services without the cluster's mounts.
        self.assertEqual(
            devnet.compose_file_args("reth"),
            ["-f", "docker-compose-devnet.yml", "-f", "docker-compose-reth.yml"],
        )
        self.assertEqual(
            devnet.compose_file_args("geth", cluster=True),
            ["-f", "docker-compose-devnet.yml", "-f", "docker-compose-cluster.yml"],
        )
        # The cluster file must come first so the reth override wins on the
        # ha-el-* image, entrypoint and command; reversed, the cluster silently
        # stays on geth.
        self.assertEqual(
            devnet.compose_file_args("reth", cluster=True),
            [
                "-f",
                "docker-compose-devnet.yml",
                "-f",
                "docker-compose-cluster.yml",
                "-f",
                "docker-compose-reth.yml",
                "-f",
                "docker-compose-cluster-reth.yml",
            ],
        )


if __name__ == "__main__":
    unittest.main()
