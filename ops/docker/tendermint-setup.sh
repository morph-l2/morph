#!/bin/bash

# Specify the F path to check
FOLDER="./.morphdevnet"

# Check if the F exists
if [ -d "$FOLDER" ]; then
  echo "Folder exists, returning directly"
  exit
fi

#./build/bin/tendermint testnet --v 4
../node/build/bin/tendermint testnet --v 4  --o $FOLDER --populate-persistent-peers --hostname-prefix node-

#sed -i '' 's#proxy_app = "tcp://127.0.0.1:26658"#proxy_app = "tcp://127.0.0.1:36658"#g' $FOLDER/node1/config/config.toml
#sed -i '' 's#laddr = "tcp://127.0.0.1:26657"#laddr = "tcp://127.0.0.1:36657"#g' $FOLDER/node1/config/config.toml
#sed -i '' 's#laddr = "tcp://0.0.0.0:26656"#laddr = "tcp://0.0.0.0:36656"#g' $FOLDER/node1/config/config.toml
#
#sed -i '' 's#proxy_app = "tcp://127.0.0.1:26658"#proxy_app = "tcp://127.0.0.1:46658"#g' $FOLDER/node2/config/config.toml
#sed -i '' 's#laddr = "tcp://127.0.0.1:26657"#laddr = "tcp://127.0.0.1:46657"#g' $FOLDER/node2/config/config.toml
#sed -i '' 's#laddr = "tcp://0.0.0.0:26656"#laddr = "tcp://0.0.0.0:46656"#g' $FOLDER/node2/config/config.toml
#
#sed -i '' 's#proxy_app = "tcp://127.0.0.1:26658"#proxy_app = "tcp://127.0.0.1:56658"#g' $FOLDER/node3/config/config.toml
#sed -i '' 's#laddr = "tcp://127.0.0.1:26657"#laddr = "tcp://127.0.0.1:56657"#g' $FOLDER/node3/config/config.toml
#sed -i '' 's#laddr = "tcp://0.0.0.0:26656"#laddr = "tcp://0.0.0.0:56656"#g' $FOLDER/node3/config/config.toml


#
#sed -i '' 's/@node0:26656/@127.0.0.1:26656/g' $FOLDER/node0/config/config.toml $FOLDER/node1/config/config.toml $FOLDER/node2/config/config.toml $FOLDER/node3/config/config.toml
#sed -i '' 's/@node1:26656/@127.0.0.1:36656/g' $FOLDER/node0/config/config.toml $FOLDER/node1/config/config.toml $FOLDER/node2/config/config.toml $FOLDER/node3/config/config.toml
#sed -i '' 's/@node2:26656/@127.0.0.1:46656/g' $FOLDER/node0/config/config.toml $FOLDER/node1/config/config.toml $FOLDER/node2/config/config.toml $FOLDER/node3/config/config.toml
#sed -i '' 's/@node3:26656/@127.0.0.1:56656/g' $FOLDER/node0/config/config.toml $FOLDER/node1/config/config.toml $FOLDER/node2/config/config.toml $FOLDER/node3/config/config.toml


sed -i '' 's#create_empty_blocks_interval = "0s"#create_empty_blocks_interval = "5s"#g' $FOLDER/node0/config/config.toml $FOLDER/node1/config/config.toml $FOLDER/node2/config/config.toml $FOLDER/node3/config/config.toml
#sed -i '' 's#timeout_commit = "1s"#timeout_commit = "3s"#g' $FOLDER/node0/config/config.toml $FOLDER/node1/config/config.toml $FOLDER/node2/config/config.toml $FOLDER/node3/config/config.toml
sed -i '' 's#batch_max_bytes = "8388608"#batch_max_bytes = "12492"#g' $FOLDER/node0/config/config.toml $FOLDER/node1/config/config.toml $FOLDER/node2/config/config.toml $FOLDER/node3/config/config.toml
sed -i '' 's#batch_blocks_interval = "10"#batch_blocks_interval = "10"#g' $FOLDER/node0/config/config.toml $FOLDER/node1/config/config.toml $FOLDER/node2/config/config.toml $FOLDER/node3/config/config.toml
sed -i '' 's#batch_timeout = "60s"#batch_timeout = "600s"#g' $FOLDER/node0/config/config.toml $FOLDER/node1/config/config.toml $FOLDER/node2/config/config.toml $FOLDER/node3/config/config.toml
sed -i '' 's#prometheus = false#prometheus = true#g' $FOLDER/node0/config/config.toml $FOLDER/node1/config/config.toml $FOLDER/node2/config/config.toml $FOLDER/node3/config/config.toml
