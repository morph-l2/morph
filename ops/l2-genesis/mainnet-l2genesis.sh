#!/bin/sh
set -eu

source ../../contracts/.env
MAINNET=".mainnet"
CONTRACT_CONFIG="../../contracts/src/deploy-config/mainnet.ts"
echo "Regenerating genesis files"
# Check if the folder exists
if [ ! -d "$MAINNET" ]; then
  echo "Folder not exists, mkdir $MAINNET"
  mkdir "$MAINNET"
fi

#if [[ -z "$RPC_URL" ]]; then
  # the environment variable is missing, set a default value
#  echo "RPC_URL is missing, using default value: http://qanet-l1"
  RPC_URL="https://mainnet.infura.io/v3/b6bf7d3508c941499b10025c0776eaf8"
#fi
echo "RPC_URL is $RPC_URL"
cat "deploy-config/mainnet-deploy-config.json" > $MAINNET/mainnet-deploy-config.json
(
go run cmd/main.go genesis l2 \
--l1-rpc $RPC_URL \
--deploy-config $MAINNET/mainnet-deploy-config.json \
--deployment-dir "$PWD/../../contracts/mainnet.json" \
--outfile.l2 $MAINNET/genesis-l2.json \
--outfile.rollup $MAINNET/rollup.json \
--outfile.genbatchheader $MAINNET/genesis-batch-header.json
touch "$MAINNET/done"
)
