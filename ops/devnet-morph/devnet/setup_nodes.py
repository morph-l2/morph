import base64
import hashlib
import json
import os
import shutil
import subprocess
import sys
import re
import tempfile

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
    key = base64.b64decode(priv_key, validate=True)
    if len(key) != 64:
        raise ValueError('Tendermint ed25519 private keys must contain 64 bytes')
    pubkey = key[32:]
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


def build_persistent_peers(devnet_dir, cluster=False):
    """Map each node directory to the peer list it should dial.

    Every tendermint-running node is given all the others, so the HA nodes
    reach each other rather than only node-0. Without that the HA block pool
    never reports caught up, the sequencer hand-over never runs, and the
    cluster silently stalls at height 0.
    """
    addresses = {}
    for node in TENDERMINT_PEERS if cluster else ('node0',):
        node_key = os.path.join(devnet_dir, node, "config", "node_key.json")
        node_id = tendermint_node_id(node_key)
        addresses[node] = f"{node_id}@{SERVICE_HOSTNAMES[node]}:26656"

    peers = {}
    for node in NODE_DIRS:
        peers[node] = ",".join(
            address for peer, address in addresses.items() if peer != node
        )
    return peers


def validate_node_files(devnet_dir):
    """Check existing node configuration without replacing identity files or consensus state."""
    for node in NODE_DIRS:
        for name in ('config.toml', 'genesis.json', 'node_key.json'):
            path = os.path.join(devnet_dir, node, 'config', name)
            if not os.path.isfile(path):
                raise RuntimeError(f'Missing existing node configuration: {path}; restore the original file before retrying')
        with open(os.path.join(devnet_dir, node, 'config', 'config.toml')) as source:
            if not re.search(r'^block_sync\s*=\s*true\s*$', source.read(), re.MULTILINE):
                raise RuntimeError(f'{node} requires block_sync = true; existing configuration was preserved')


def node_file_hashes(devnet_dir):
    result = {}
    for node in NODE_DIRS:
        names = ['config.toml', 'genesis.json', 'node_key.json']
        if node == 'node0':
            names.append('priv_validator_key.json')
        for name in names:
            relative = os.path.join(node, 'config', name)
            with open(os.path.join(devnet_dir, relative), 'rb') as source:
                result[relative] = hashlib.sha256(source.read()).hexdigest()
    return result


def setup_devnet_nodes(root_dir=None, cluster=False):
    """Verify existing node homes or publish a complete new set after generation succeeds."""
    root_dir = root_dir or subprocess.check_output(["git", "rev-parse", "--show-toplevel"], text=True).strip()
    docker_dir = os.path.join(root_dir, 'ops', 'docker')
    devnet_dir = os.path.join(docker_dir, '.devnet')
    marker = os.path.join(devnet_dir, 'nodes.done')
    if os.path.exists(devnet_dir):
        if not os.path.isfile(marker):
            raise RuntimeError(f'Existing node directory has no nodes.done: {devnet_dir}; preserve its keys and data and inspect the interrupted setup')
        validate_node_files(devnet_dir)
        with open(marker) as source:
            saved = json.load(source)
        if saved != {'version': 1, 'files': node_file_hashes(devnet_dir)}:
            raise RuntimeError('Existing node configuration differs from nodes.done; preserve node data and restore the original configuration before retrying')
        print('Existing devnet nodes verified; preserving their keys and data.')
        return
    tendermint = shutil.which("tendermint")
    if tendermint is None:
        node_dir = os.path.join(root_dir, "node")
        if not os.path.isdir(node_dir):
            raise RuntimeError(f'Node directory not found: {node_dir}')
        tendermint = os.path.join(node_dir, 'build', 'bin', 'tendermint')
        print(f"Building Tendermint in {node_dir}...")
        subprocess.run(["make", "tendermint"], cwd=node_dir, check=True)
        if not os.path.isfile(tendermint) or not os.access(tendermint, os.X_OK):
            raise RuntimeError(f'Tendermint build did not produce an executable: {tendermint}')

    attempt = tempfile.mkdtemp(prefix='.devnet-setup-', dir=docker_dir)
    try:
        generate_node_files(docker_dir, attempt, tendermint, cluster)
        os.rename(attempt, devnet_dir)
    except (Exception, SystemExit):
        print(f'Node setup failed; existing data was preserved. Inspect generated files in {attempt}')
        raise


def generate_node_files(docker_dir, devnet_dir, tendermint, cluster):
    """Generate node files in a separate directory before publishing nodes.done."""
    marker = os.path.join(devnet_dir, 'nodes.done')

    # Run the Tendermint testnet command
    print("Setting up the devnet...")
    command = [
        tendermint, "testnet", "--v", "1", "--n", "5", "--o", devnet_dir,
        "--populate-persistent-peers",
        "--hostname", "node-0",
        "--hostname", "node-1",
        "--hostname", "node-2",
        "--hostname", "ha-node-0",
        "--hostname", "ha-node-1",
        "--hostname", "ha-node-2",
    ]

    subprocess.run(command, check=True)

    # Rename generated non-validator directories to match the compose service names.
    for generated, desired in (("node3", "ha-node0"), ("node4", "ha-node1"), ("node5", "ha-node2")):
        generated_path = os.path.join(devnet_dir, generated)
        desired_path = os.path.join(devnet_dir, desired)
        if os.path.exists(generated_path):
            os.rename(generated_path, desired_path)

    # Install the key files first: node IDs are derived from node_key.json, so
    # the peer addresses below must be computed from the final keys.
    copy_key_files(docker_dir, devnet_dir)

    persistent_peers = build_persistent_peers(devnet_dir, cluster=cluster)

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
    validate_node_files(devnet_dir)
    with open(marker, 'x') as target:
        json.dump({'version': 1, 'files': node_file_hashes(devnet_dir)}, target, indent=2)
        target.write('\n')
    print("Devnet nodes setup completed successfully.")
