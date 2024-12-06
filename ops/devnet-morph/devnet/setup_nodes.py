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

        # back to the root directory
        os.chdir(root_dir)
        print("Tendermint build process completed.")

    # Check if .devnet directory already exists
    docker_dir = os.path.join(root_dir, "ops", "docker")
    devnet_dir = os.path.join(docker_dir, ".devnet")
    if os.path.exists(devnet_dir):
        print(".devnet directory already exists. Devnet nodes setup has already been completed. Exiting.")
        return

    # Run the Tendermint testnet command
    print("Setting up the devnet...")
    command = [
        "tendermint", "testnet", "--v", "4", "--n", "1", "--o", devnet_dir,
        "--populate-persistent-peers",
        "--hostname", "node-0",
        "--hostname", "node-1",
        "--hostname", "node-2",
        "--hostname", "node-3",
        "--hostname", "sentry-node-0"
    ]

    if subprocess.call(command) != 0:
        print("Failed to set up devnet.")
        sys.exit(1)

    # Modify config.toml files using toml library
    print("Modifying config.toml files...")
    config_files = [
        os.path.join(devnet_dir, f"node{i}/config/config.toml") for i in range(5)
    ]

    persistent_peers_value = (
        "93e27ea2306e158a8146d5f44caaab97496797d2@node-0:26656,"
        "7f78b7d7a7e6bad4faf68d5731d437f4288d96d0@node-1:26656,"
        "06c699be2f9aeb9f7ec79f508a95ff80576deb12@node-2:26656,"
        "b1a131f40d5d3abefe0dd787513c936ef62ac2d6@node-3:26656,"
        "dae813274913aaf39e7cd3226a0aa8bce00644e1@sentry-node-0:26656"
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

        # Modify pex for nodes 0 to 3
        if i < 4:
            content = content.replace('pex = true', 'pex = false')

        with open(config_file, "w") as f:
            f.write(content)

    print("All config.toml files have been updated successfully.")

    # Copy key files to devnet node directories
    print("Copying key files...")
    node_dirs = [f"node{i}" for i in range(5)]

    for node in node_dirs:
        source_dir = os.path.join(docker_dir, node)
        dest_dir = os.path.join(devnet_dir, node, "config")

        if not os.path.isdir(source_dir) or not os.path.isdir(dest_dir):
            print(f"Error: Missing source or destination directory for {node}. Exiting.")
            sys.exit(1)

        # Copy specific files
        shutil.copyfile(os.path.join(source_dir, "node_key.json"), os.path.join(dest_dir, "node_key.json"))

        # Skip copying for node4 bls_key.json and priv_validator_key.json
        if node != "node4":
            shutil.copyfile(os.path.join(source_dir, "bls_key.json"), os.path.join(dest_dir, "bls_key.json"))
            shutil.copyfile(os.path.join(source_dir, "priv_validator_key.json"), os.path.join(dest_dir, "priv_validator_key.json"))

        # Copy and rename genesis file
        shutil.copyfile(os.path.join(docker_dir, "tendermint-devnet-genesis.json"), os.path.join(dest_dir, "genesis.json"))

        print(f"Files copied successfully for {node}.")

    print("All key files have been copied successfully.")
    print("Devnet nodes setup completed successfully.")
