#!/bin/bash
set -e

GETH_DATA_DIR=${GETH_DATA_DIR:-/db}
GENESIS_FILE_PATH=${GENESIS_FILE_PATH:-/genesis.json}
JWT_SECRET_PATH=${JWT_SECRET_PATH:-/jwt-secret.txt}

# Initialize geth if not already done
if [ ! -d "$GETH_DATA_DIR/geth/chaindata" ]; then
    echo "Initializing geth with genesis file..."
    geth init --datadir "$GETH_DATA_DIR" "$GENESIS_FILE_PATH"
fi

echo "Starting geth..."
exec geth \
    --datadir "$GETH_DATA_DIR" \
    --http \
    --http.addr "0.0.0.0" \
    --http.port 8545 \
    --http.api "eth,net,web3,debug,txpool,engine" \
    --http.corsdomain "*" \
    --http.vhosts "*" \
    --ws \
    --ws.addr "0.0.0.0" \
    --ws.port 8546 \
    --ws.api "eth,net,web3,debug,txpool,engine" \
    --ws.origins "*" \
    --authrpc.addr "0.0.0.0" \
    --authrpc.port 8551 \
    --authrpc.vhosts "*" \
    --authrpc.jwtsecret "$JWT_SECRET_PATH" \
    --networkid 53077 \
    --nodiscover \
    --syncmode full \
    --gcmode archive \
    --metrics \
    --metrics.addr "0.0.0.0" \
    --pprof \
    --pprof.addr "0.0.0.0" \
    --verbosity 3 \
    "$@"

