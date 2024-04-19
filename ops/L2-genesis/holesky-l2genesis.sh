#!/bin/sh
set -eu

source ../../contracts/.env
TESTNET=".holesky"
CONTRACT_CONFIG="../../contracts/src/deploy-config/holesky.ts"
echo "Regenerating genesis files"
# Check if the folder exists
if [ ! -d "$TESTNET" ]; then
  echo "Folder not exists, mkdir $TESTNET"
  mkdir "$TESTNET"
fi

if [[ -z "$HOLESKY_RPC_URL" ]]; then
  # the environment variable is missing, set a default value
  echo "L1_RPC is missing, using default value: http://holesky-testnet"
  HOLESKY_RPC_URL="http://holesky-testnet:8545"
fi
echo "HOLESKY_RPC_URL is $HOLESKY_RPC_URL"
cat "deploy-config/holesky-deploy-config.json" > $TESTNET/holesky-deploy-config.json
(
go run cmd/main.go genesis l2 \
--l1-rpc $HOLESKY_RPC_URL \
--deploy-config $TESTNET/holesky-deploy-config.json \
--deployment-dir "$PWD/../../contracts/holesky.json" \
--outfile.l2 $TESTNET/genesis-l2.json \
--outfile.rollup $TESTNET/rollup.json \
--outfile.genbatchheader $TESTNET/genesis-batch-header.json
touch "$TESTNET/done"
)
l2_genesis_state_root=$(cat $TESTNET/rollup.json | jq -r .l2_genesis_state_root)
withdraw_root=$(cat $TESTNET/rollup.json | jq -r .withdraw_root)
genesis_batch_header=$(cat $TESTNET/rollup.json | jq -r .genesis_batch_header)

echo "Replacing l2_genesis_state_root: $l2_genesis_state_root"
sed -i "s#rollupGenesisStateRoot: '.*'#rollupGenesisStateRoot: '$l2_genesis_state_root'#g" $CONTRACT_CONFIG

echo "Replacing withdraw_root: $withdraw_root"
sed -i "s#withdrawRoot: '.*'#withdrawRoot: '$withdraw_root'#g" $CONTRACT_CONFIG

echo "Replacing  genesis_batch_header: $genesis_batch_header"
sed -i "s#batchHeader: '.*'#batchHeader: '$genesis_batch_header'#g" $CONTRACT_CONFIG
