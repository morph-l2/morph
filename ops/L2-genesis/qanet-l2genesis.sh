#!/bin/sh

QANET=".qanet"
CONTRACT_CONFIG="../contracts/src/deploy-config/qanetl1.ts"
echo "Regenerating genesis files"
# Check if the folder exists
if [ ! -d "$QANET" ]; then
  echo "Folder not exists, mkdir $QANET"
  mkdir "$QANET"
fi

if [[ -z "$L1_RPC" ]]; then
  # the environment variable is missing, set a default value
  echo "L1_RPC is missing, using default value: http://l2-qa-morph-l1-geth.bitkeep.tools"
  L1_RPC="http://l2-qa-morph-l1-geth.bitkeep.tools"
fi

cat "deploy-config/qanet-deploy-config.json" > $QANET/qanet-deploy-config.json
(
go run cmd/main.go genesis l2 \
--l1-rpc $L1_RPC \
--deploy-config $QANET/qanet-deploy-config.json \
--deployment-dir "$PWD/../contracts/qanetL1.json" \
--outfile.l2 $QANET/genesis-l2.json \
--outfile.rollup $QANET/rollup.json \
--outfile.genbatchheader $QANET/genesis-batch-header.json
touch "$DEVNET/done"
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
