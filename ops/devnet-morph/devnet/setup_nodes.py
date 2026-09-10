import base64
import hashlib
import json
import os
import shutil
import subprocess
import sys
import re

# Directories that hold a node's tendermint home, in the order their config
# files are processed. The first entry is the genesis validator.
NODE_DIRS = ("node0", "node1", "node2", "ha-node0", "ha-node1", "ha-node2")

# Directory name -> compose service hostname, used to build peer addresses.
SERVICE_HOSTNAMES = {
    "node0": "node-0",
    "node1": "node-1",
    "node2": "node-2",
    "ha-node0": "ha-node-0",
    "ha-node1": "ha-node-1",
    "ha-node2": "ha-node-2",
}

# Nodes that actually run tendermint, and so can be dialed as peers. node-1
# runs with MORPH_NODE_DERIVATION_VERIFY_MODE=layer1 and never starts
# tendermint; node-2 has no compose service at all. Listing either as a peer
# only produces endless reconnect and DNS lookup failures.
TENDERMINT_PEERS = ("node0", "ha-node0", "ha-node1", "ha-node2")


def tendermint_node_id(node_key_path):
    """Derive a tendermint node ID from a node_key.json file.

    The ID is the hex encoding of the first 20 bytes of sha256(pubkey). An
    ed25519 private key is stored as seed(32) || pubkey(32), so the public half
    is the tail of the decoded value.

    IDs must be derived from the key files that are actually in place, which is
    why this runs after the key files have been copied: overwriting a
    node_key.json changes the node's identity.
    """
    with open(node_key_path) as f:
        priv_key = json.load(f)["priv_key"]["value"]
    pubkey = base64.b64decode(priv_key)[32:]
    return hashlib.sha256(pubkey).hexdigest()[:40]


def copy_key_files(docker_dir, devnet_dir):
    """Install the fixed node keys and the shared genesis into each node home.

    Only node0 gets a genesis validator key. The others must not have one: a
    node holding the sole genesis validator key gets block sync disabled and
    never hands over to the sequencer routines.
    """
    print("Copying key files...")

    for node in NODE_DIRS:
        source_dir = os.path.join(docker_dir, node)
        dest_dir = os.path.join(devnet_dir, node, "config")

        if not os.path.isdir(dest_dir):
            print(f"Error: Missing destination directory for {node}. Exiting.")
            sys.exit(1)

        if os.path.isdir(source_dir):
            shutil.copyfile(os.path.join(source_dir, "node_key.json"), os.path.join(dest_dir, "node_key.json"))

        if node == "node0" and os.path.isdir(source_dir):
            shutil.copyfile(os.path.join(source_dir, "priv_validator_key.json"), os.path.join(dest_dir, "priv_validator_key.json"))
        else:
            priv_validator_key = os.path.join(dest_dir, "priv_validator_key.json")
            priv_validator_state = os.path.join(devnet_dir, node, "data", "priv_validator_state.json")
            for validator_file in (priv_validator_key, priv_validator_state):
                if os.path.exists(validator_file):
                    os.remove(validator_file)

        # Copy and rename genesis file
        shutil.copyfile(os.path.join(docker_dir, "tendermint-devnet-genesis.json"), os.path.join(dest_dir, "genesis.json"))

        print(f"Files copied successfully for {node}.")

    print("All key files have been copied successfully.")


def build_persistent_peers(devnet_dir):
    """Map each node directory to the peer list it should dial.

    Every tendermint-running node is given all the others, so the HA nodes
    reach each other rather than only node-0. Without that the HA block pool
    never reports caught up, the sequencer hand-over never runs, and the
    cluster silently stalls at height 0.
    """
    addresses = {}
    for node in TENDERMINT_PEERS:
        node_key = os.path.join(devnet_dir, node, "config", "node_key.json")
        node_id = tendermint_node_id(node_key)
        addresses[node] = f"{node_id}@{SERVICE_HOSTNAMES[node]}:26656"

    peers = {}
    for node in NODE_DIRS:
        peers[node] = ",".join(
            address for peer, address in addresses.items() if peer != node
        )
    return peers


def setup_devnet_nodes():
    """
    Set up the devnet nodes, modify configuration files using toml library, and copy key files.
    """
    root_dir = subprocess.check_output(["git", "rev-parse", "--show-toplevel"], text=True).strip()
    # Check if Tendermint is installed
    if shutil.which("tendermint") is None:
        print("Tendermint is not installed. Starting the build process...")
        node_dir = os.path.join(root_dir, "node")
        ops_dir = os.path.join(root_dir, "ops", "docker")

        if not os.path.isdir(node_dir):
            print(f"Error: Node directory not found at {node_dir}. Exiting.")
            sys.exit(1)

        os.chdir(node_dir)
        print(f"Building Tendermint in {node_dir}...")
        if subprocess.call(["make", "install-tendermint"]) != 0:
            print("Error: Failed to build Tendermint. Exiting.")
            sys.exit(1)

        os.chdir(root_dir)
        print("Tendermint build process completed.")

    # Check if .devnet directory already exists
    docker_dir = os.path.join(root_dir, "ops", "docker")
    devnet_dir = os.path.join(docker_dir, ".devnet")
    if os.path.exists(devnet_dir):
        old_topology_paths = [os.path.join(devnet_dir, f"node{i}") for i in range(3, 6)]
        expected_paths = [os.path.join(devnet_dir, node) for node in NODE_DIRS]
        if any(os.path.exists(path) for path in old_topology_paths) or any(
                not os.path.exists(path) for path in expected_paths):
            print("Existing stale devnet detected. Regenerating single-sequencer config.")
            shutil.rmtree(devnet_dir)
        else:
            print(".devnet directory already exists. Devnet nodes setup has already been completed. Exiting.")
            return

    # Run the Tendermint testnet command
    print("Setting up the devnet...")
    command = [
        "tendermint", "testnet", "--v", "1", "--n", "5", "--o", devnet_dir,
        "--populate-persistent-peers",
        "--hostname", "node-0",
        "--hostname", "node-1",
        "--hostname", "node-2",
        "--hostname", "ha-node-0",
        "--hostname", "ha-node-1",
        "--hostname", "ha-node-2",
    ]

    if subprocess.call(command) != 0:
        print("Failed to set up devnet.")
        sys.exit(1)

    # Rename generated non-validator directories to match the compose service names.
    for generated, desired in (("node3", "ha-node0"), ("node4", "ha-node1"), ("node5", "ha-node2")):
        generated_path = os.path.join(devnet_dir, generated)
        desired_path = os.path.join(devnet_dir, desired)
        if os.path.exists(generated_path):
            os.rename(generated_path, desired_path)

    # Install the key files first: node IDs are derived from node_key.json, so
    # the peer addresses below must be computed from the final keys.
    copy_key_files(docker_dir, devnet_dir)

    persistent_peers = build_persistent_peers(devnet_dir)

    # Modify config.toml files.
    print("Modifying config.toml files...")

    for i, node in enumerate(NODE_DIRS):
        config_file = os.path.join(devnet_dir, node, "config", "config.toml")
        if not os.path.isfile(config_file):
            print(f"Error: {config_file} not found. Exiting.")
            sys.exit(1)

        with open(config_file, "r") as f:
            content = f.read()

        # Replace the required fields
        content = content.replace('create_empty_blocks_interval = "0s"', 'create_empty_blocks_interval = "5s"')
        content = content.replace('peer_gossip_sleep_duration = "100ms"', 'peer_gossip_sleep_duration = "10ms"')
        content = content.replace('flush_throttle_timeout = "100ms"', 'flush_throttle_timeout = "10ms"')
        content = content.replace('max_packet_msg_payload_size = 1024', 'max_packet_msg_payload_size = 10485760')
        content = content.replace('send_rate = 5120000', 'send_rate = 52428800')
        content = content.replace('recv_rate = 5120000', 'recv_rate = 102428800')
        content = content.replace('block_sync = false', 'block_sync = true')
        content = re.sub(r'persistent_peers\s*=\s*".*?"', f'persistent_peers = "{persistent_peers[node]}"', content)

        # Serve the RPC on all interfaces so the published container ports
        # (26657, 27657, 27757, 27857) are reachable from the host.
        content = content.replace('laddr = "tcp://127.0.0.1:26657"', 'laddr = "tcp://0.0.0.0:26657"')

        # Modify pex for the sequencer validator node.
        if i == 0:
            content = content.replace('pex = true', 'pex = false')

        # Enable prometheus metrics for all nodes
        content = content.replace('prometheus = false', 'prometheus = true')

        with open(config_file, "w") as f:
            f.write(content)

    print("All config.toml files have been updated successfully.")
    print("Devnet nodes setup completed successfully.")
