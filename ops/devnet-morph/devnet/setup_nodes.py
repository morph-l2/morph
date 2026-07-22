import os
import shutil
import subprocess
import sys
import re

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
        expected_paths = [
            os.path.join(devnet_dir, "node0"),
            os.path.join(devnet_dir, "node1"),
            os.path.join(devnet_dir, "node2"),
            os.path.join(devnet_dir, "ha-node0"),
            os.path.join(devnet_dir, "ha-node1"),
            os.path.join(devnet_dir, "ha-node2"),
        ]
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

    # Modify config.toml files.
    print("Modifying config.toml files...")
    config_files = [
        os.path.join(devnet_dir, node, "config", "config.toml")
        for node in ("node0", "node1", "node2", "ha-node0", "ha-node1", "ha-node2")
    ]

    persistent_peers_value = (
        "93e27ea2306e158a8146d5f44caaab97496797d2@node-0:26656,"
        "7f78b7d7a7e6bad4faf68d5731d437f4288d96d0@node-1:26656,"
        "06c699be2f9aeb9f7ec79f508a95ff80576deb12@node-2:26656"
    )

    for i, config_file in enumerate(config_files):
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
        content = re.sub(r'persistent_peers\s*=\s*".*?"', f'persistent_peers = "{persistent_peers_value}"', content)

        # Modify pex for the sequencer validator node.
        if i == 0:
            content = content.replace('pex = true', 'pex = false')

        # Enable prometheus metrics for all nodes
        content = content.replace('prometheus = false', 'prometheus = true')

        with open(config_file, "w") as f:
            f.write(content)

    print("All config.toml files have been updated successfully.")

    # Copy key files to devnet node directories
    print("Copying key files...")
    node_dirs = ["node0", "node1", "node2", "ha-node0", "ha-node1", "ha-node2"]

    for node in node_dirs:
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
    print("Devnet nodes setup completed successfully.")
