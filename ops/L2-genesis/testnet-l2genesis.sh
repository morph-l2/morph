#!/bin/sh

TESTNET=".testnet"
CONTRACT_CONFIG="../contracts/src/deploy-config/sepolia.ts"
echo "Regenerating genesis files"
# Check if the folder exists
if [ ! -d "$TESTNET" ]; then
  echo "Folder not exists, mkdir $TESTNET"
  mkdir "$TESTNET"
fi

if [[ -z "$L1_RPC" ]]; then
  # the environment variable is missing, set a default value
  echo "L1_RPC is missing, using default value: http://10.11.63.110:8545"
  L1_RPC="http://10.11.63.110:8545"
fi

cat "deploy-config/testnet-deploy-config.json" > $TESTNET/testnet-deploy-config.json
(
go run cmd/main.go genesis l2 \
--l1-rpc $L1_RPC \
--deploy-config $TESTNET/testnet-deploy-config.json \
--deployment-dir "$PWD/../contracts/sepolia.json" \
--outfile.l2 $TESTNET/genesis-l2.json \
--outfile.rollup $TESTNET/rollup.json \
--outfile.genbatchheader $TESTNET/genesis-batch-header.json
touch "$DEVNET/done"
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
