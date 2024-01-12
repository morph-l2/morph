#!/bin/bash

# Specify the folder path to check
folder="./mytestnet"

# Check if the folder exists
if [ -d "$folder" ]; then
  echo "Folder exists, returning directly"
  exit
fi

#./build/bin/tendermint testnet --v 4
./build/bin/tendermint testnet --v 4 --populate-persistent-peers --hostname-prefix node-

#sed -i '' 's#proxy_app = "tcp://127.0.0.1:26658"#proxy_app = "tcp://127.0.0.1:36658"#g' ./mytestnet/node1/config/config.toml
#sed -i '' 's#laddr = "tcp://127.0.0.1:26657"#laddr = "tcp://127.0.0.1:36657"#g' ./mytestnet/node1/config/config.toml
#sed -i '' 's#laddr = "tcp://0.0.0.0:26656"#laddr = "tcp://0.0.0.0:36656"#g' ./mytestnet/node1/config/config.toml
#
#sed -i '' 's#proxy_app = "tcp://127.0.0.1:26658"#proxy_app = "tcp://127.0.0.1:46658"#g' ./mytestnet/node2/config/config.toml
#sed -i '' 's#laddr = "tcp://127.0.0.1:26657"#laddr = "tcp://127.0.0.1:46657"#g' ./mytestnet/node2/config/config.toml
#sed -i '' 's#laddr = "tcp://0.0.0.0:26656"#laddr = "tcp://0.0.0.0:46656"#g' ./mytestnet/node2/config/config.toml
#
#sed -i '' 's#proxy_app = "tcp://127.0.0.1:26658"#proxy_app = "tcp://127.0.0.1:56658"#g' ./mytestnet/node3/config/config.toml
#sed -i '' 's#laddr = "tcp://127.0.0.1:26657"#laddr = "tcp://127.0.0.1:56657"#g' ./mytestnet/node3/config/config.toml
#sed -i '' 's#laddr = "tcp://0.0.0.0:26656"#laddr = "tcp://0.0.0.0:56656"#g' ./mytestnet/node3/config/config.toml


#
#sed -i '' 's/@node0:26656/@127.0.0.1:26656/g' ./mytestnet/node0/config/config.toml ./mytestnet/node1/config/config.toml ./mytestnet/node2/config/config.toml ./mytestnet/node3/config/config.toml
#sed -i '' 's/@node1:26656/@127.0.0.1:36656/g' ./mytestnet/node0/config/config.toml ./mytestnet/node1/config/config.toml ./mytestnet/node2/config/config.toml ./mytestnet/node3/config/config.toml
#sed -i '' 's/@node2:26656/@127.0.0.1:46656/g' ./mytestnet/node0/config/config.toml ./mytestnet/node1/config/config.toml ./mytestnet/node2/config/config.toml ./mytestnet/node3/config/config.toml
#sed -i '' 's/@node3:26656/@127.0.0.1:56656/g' ./mytestnet/node0/config/config.toml ./mytestnet/node1/config/config.toml ./mytestnet/node2/config/config.toml ./mytestnet/node3/config/config.toml


sed -i '' 's#create_empty_blocks_interval = "0s"#create_empty_blocks_interval = "5s"#g' ./mytestnet/node0/config/config.toml ./mytestnet/node1/config/config.toml ./mytestnet/node2/config/config.toml ./mytestnet/node3/config/config.toml
sed -i '' 's#timeout_commit = "1s"#timeout_commit = "3s"#g' ./mytestnet/node0/config/config.toml ./mytestnet/node1/config/config.toml ./mytestnet/node2/config/config.toml ./mytestnet/node3/config/config.toml