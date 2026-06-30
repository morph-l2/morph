import importlib
import sys
import unittest
from pathlib import Path


REPO_ROOT = Path(__file__).resolve().parents[3]
DEVNET_PACKAGE = REPO_ROOT / "ops" / "devnet-morph"
DOCKER_DIR = REPO_ROOT / "ops" / "docker"


class DevnetConfigTest(unittest.TestCase):
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
            "DEVNET_CLEAN_COMPOSE_FILES := -f docker-compose-devnet.yml -f docker-compose-reth.yml -f docker-compose-cluster.yml",
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

    def test_cluster_compose_defines_ha_services(self):
        cluster_compose = DOCKER_DIR / "docker-compose-cluster.yml"

        self.assertTrue(cluster_compose.exists())
        compose = cluster_compose.read_text()
        for service in ("ha-geth-0:", "ha-geth-1:", "ha-geth-2:", "ha-node-0:", "ha-node-1:", "ha-node-2:"):
            self.assertIn(service, compose)
        self.assertIn("MORPH_NODE_HA_ENABLED=true", compose)
        self.assertIn("MORPH_NODE_HA_BOOTSTRAP=true", compose)
        self.assertIn("MORPH_NODE_HA_JOIN=ha-node-0:9401", compose)

    def test_compose_file_args_can_enable_cluster_mode(self):
        sys.path.insert(0, str(DEVNET_PACKAGE))
        try:
            devnet = importlib.import_module("devnet")
            importlib.reload(devnet)
        finally:
            sys.path.remove(str(DEVNET_PACKAGE))

        self.assertEqual(devnet.compose_file_args("geth"), ["-f", "docker-compose-devnet.yml"])
        self.assertEqual(
            devnet.compose_file_args("reth"),
            ["-f", "docker-compose-devnet.yml", "-f", "docker-compose-reth.yml"],
        )
        self.assertEqual(
            devnet.compose_file_args("geth", cluster=True),
            ["-f", "docker-compose-devnet.yml", "-f", "docker-compose-cluster.yml"],
        )
        self.assertEqual(
            devnet.compose_file_args("reth", cluster=True),
            [
                "-f",
                "docker-compose-devnet.yml",
                "-f",
                "docker-compose-reth.yml",
                "-f",
                "docker-compose-cluster.yml",
            ],
        )


if __name__ == "__main__":
    unittest.main()
