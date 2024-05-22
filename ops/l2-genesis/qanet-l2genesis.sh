#!/bin/sh
set -eu

source ../../contracts/.env
QANET=".qanet"
CONTRACT_CONFIG="../../contracts/src/deploy-config/qanetl1.ts"
echo "Regenerating genesis files"
# Check if the folder exists
if [ ! -d "$QANET" ]; then
  echo "Folder not exists, mkdir $QANET"
  mkdir "$QANET"
fi

if [[ -z "$QA_RPC_URL" ]]; then
  # the environment variable is missing, set a default value
  echo "QA_RPC_URL is missing, using default value: http://qanet-l1"
  QA_RPC_URL="http://qanet-l1"
fi
echo "QA_RPC_URL is $QA_RPC_URL"
cat "deploy-config/qanet-deploy-config.json" > $QANET/qanet-deploy-config.json
(
go run cmd/main.go genesis l2 \
--l1-rpc $QA_RPC_URL \
--deploy-config $QANET/qanet-deploy-config.json \
--deployment-dir "$PWD/../../contracts/qanetL1.json" \
--outfile.l2 $QANET/genesis-l2.json \
--outfile.rollup $QANET/rollup.json \
--outfile.genbatchheader $QANET/genesis-batch-header.json
touch "$QANET/done"
)
l2_genesis_state_root=$(cat $QANET/rollup.json | jq -r .l2_genesis_state_root)
withdraw_root=$(cat $QANET/rollup.json | jq -r .withdraw_root)
genesis_batch_header=$(cat $QANET/rollup.json | jq -r .genesis_batch_header)

echo "Replacing l2_genesis_state_root: $l2_genesis_state_root"
sed -i "s#rollupGenesisStateRoot: '.*'#rollupGenesisStateRoot: '$l2_genesis_state_root'#g" $CONTRACT_CONFIG

echo "Replacing withdraw_root: $withdraw_root"
sed -i "s#withdrawRoot: '.*'#withdrawRoot: '$withdraw_root'#g" $CONTRACT_CONFIG

echo "Replacing  genesis_batch_header: $genesis_batch_header"
sed -i "s#batchHeader: '.*'#batchHeader: '$genesis_batch_header'#g" $CONTRACT_CONFIG
