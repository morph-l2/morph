#!/bin/sh
set -eu

: "${GETH_NETWORK_ID:?Set GETH_NETWORK_ID for the restored snapshot}"
case "$GETH_NETWORK_ID" in
    ''|*[!0-9]*) echo "GETH_NETWORK_ID must be a positive integer" >&2; exit 1 ;;
esac
[ "$GETH_NETWORK_ID" -gt 0 ] || { echo "GETH_NETWORK_ID must be positive" >&2; exit 1; }
[ -d /db/geth/chaindata ] || { echo "Restore an initialized snapshot in /db before starting geth" >&2; exit 1; }
[ -f /jwt-secret.txt ] || { echo "JWT secret file is missing" >&2; exit 1; }

exec geth \
    --networkid "$GETH_NETWORK_ID" \
    --datadir /db \
    --verbosity 3 \
    --http \
    --http.corsdomain '*' \
    --http.vhosts '*' \
    --http.addr 0.0.0.0 \
    --http.port 8545 \
    --http.api web3,debug,eth,txpool,net,morph \
    --ws \
    --ws.addr 0.0.0.0 \
    --ws.port 8546 \
    --ws.origins '*' \
    --ws.api web3,debug,eth,txpool,net,morph \
    --authrpc.addr 0.0.0.0 \
    --authrpc.port 8551 \
    --authrpc.vhosts '*' \
    --authrpc.jwtsecret /jwt-secret.txt \
    --gcmode archive \
    --log.filename /db/geth.log \
    --metrics \
    --metrics.addr 0.0.0.0
