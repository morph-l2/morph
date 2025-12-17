#!/bin/sh
# Single-node L2 execution engine (no block production)
# - Uses local-test/genesis.json and local-test/jwt-secret.txt by default
# - Initializes the datadir on first run
# - Does NOT mine; you can drive block production via Engine API yourself

set -ex

# Resolve script directory (absolute path)
SCRIPT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)

# Locate geth binary
GETH_BIN="${GETH_BIN:-geth}"
if ! command -v "$GETH_BIN" >/dev/null 2>&1; then
  if [ -x "$SCRIPT_DIR/../build/bin/geth" ]; then
    GETH_BIN="$SCRIPT_DIR/../build/bin/geth"
  else
    echo "ERROR: geth binary not found. Build with 'make geth' at repo root or set GETH_BIN to absolute path." >&2
    exit 1
  fi
fi

# Configurable via environment variables
GETH_DATA_DIR="${GETH_DATA_DIR:-$SCRIPT_DIR/geth-data}"
GENESIS_PATH="${GENESIS_PATH:-$SCRIPT_DIR/genesis.json}"
JWT_SECRET_PATH="${JWT_SECRET_PATH:-$SCRIPT_DIR/jwt-secret.txt}"
CHAIN_ID="${CHAIN_ID:-}"

# Derive CHAIN_ID from genesis if not provided (requires jq)
if [ -z "$CHAIN_ID" ]; then
  if command -v jq >/dev/null 2>&1; then
    CHAIN_ID=$(jq -r '.config.chainId' "$GENESIS_PATH")
  else
    CHAIN_ID=53077
    echo "WARN: jq not found; defaulting CHAIN_ID=$CHAIN_ID. Export CHAIN_ID to override." >&2
  fi
fi

# One-time init with provided genesis
if [ ! -d "$GETH_DATA_DIR/geth/chaindata" ]; then
  "$GETH_BIN" --datadir "$GETH_DATA_DIR" init "$GENESIS_PATH"
fi


# Start execution engine only (no mining, no discovery)
exec "$GETH_BIN" \
  --datadir "$GETH_DATA_DIR" \
  --networkid "$CHAIN_ID" \
  --verbosity 3 \
  --http \
  --http.addr 0.0.0.0 \
  --http.port 8545 \
  --http.corsdomain "*" \
  --http.vhosts "*" \
  --http.api web3,eth,txpool,net,morph,engine,admin,debug \
  --ws \
  --ws.addr 0.0.0.0 \
  --ws.port 8546 \
  --ws.origins "*" \
  --ws.api web3,eth,txpool,net,morph,engine,admin,debug \
  --authrpc.addr 0.0.0.0 \
  --authrpc.port 8551 \
  --authrpc.vhosts "*" \
  --authrpc.jwtsecret "$JWT_SECRET_PATH" \
  --gcmode archive \
  --rpc.allow-unprotected-txs \
  --nodiscover \
  --maxpeers 0


