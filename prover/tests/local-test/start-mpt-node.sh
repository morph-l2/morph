#!/bin/bash
# Start MPT node (syncs from zkTrie node)

set -ex

SCRIPT_DIR=$(cd "$(dirname "$0")" && pwd)
GETH_BIN="${GETH_BIN:-$SCRIPT_DIR/../build/bin/geth}"

# MPT node configuration
GETH_DATA_DIR="${GETH_DATA_DIR:-$SCRIPT_DIR/mpt-data}"
GENESIS_MPT_PATH="$SCRIPT_DIR/genesis-mpt.json"
JWT_SECRET_PATH="$SCRIPT_DIR/jwt-secret.txt"

# Initialize if needed
if [ ! -d "$GETH_DATA_DIR/geth/chaindata" ]; then
  echo "Initializing MPT node with genesis..."
  "$GETH_BIN" --datadir "$GETH_DATA_DIR" init "$GENESIS_MPT_PATH"
fi

# Start MPT node
# IMPORTANT: 
# - --morph-mpt: Enable MPT format
# - --cache.snapshot=0: Disable snapshot for cross-format genesis sync
exec "$GETH_BIN" \
  --datadir "$GETH_DATA_DIR" \
  --networkid 53077 \
  --verbosity 3 \
  --morph-mpt \
  --cache.snapshot=0 \
  --http \
  --http.addr 0.0.0.0 \
  --http.port 9545 \
  --http.corsdomain "*" \
  --http.vhosts "*" \
  --http.api web3,eth,txpool,net,morph,engine,admin,debug \
  --ws \
  --ws.addr 0.0.0.0 \
  --ws.port 9546 \
  --ws.origins "*" \
  --ws.api web3,eth,txpool,net,morph,engine,admin,debug \
  --authrpc.addr 0.0.0.0 \
  --authrpc.port 9551 \
  --authrpc.vhosts "*" \
  --authrpc.jwtsecret "$JWT_SECRET_PATH" \
  --gcmode archive \
  --rpc.allow-unprotected-txs \
  --port 30304 \
  --nodiscover

# To connect to zkTrie node after both are running:
# ./connect-nodes.sh
