DEVNET="$PWD/.devnet"
echo "Regenerating genesis files"
# Check if the folder exists
if [ ! -d "$DEVNET" ]; then
  echo "Folder not exists, mkdir $DEVNET"
  mkdir "$DEVNET"
fi

TIMESTAMP=$(date +%s | xargs printf '0x%x')
cat "deploy-config/devnet-deploy-config.json" | jq -r ".l1GenesisBlockTimestamp = \"$TIMESTAMP\"" > $DEVNET/deploy-config.json
(
go run cmd/main.go genesis l2 \
--l1-rpc http://localhost:9545 \
--deploy-config $DEVNET/deploy-config.json \
--deployment-dir $DEVNET/devnetL1.json \
--outfile.l2 $DEVNET/genesis-l2.json \
--outfile.rollup $DEVNET/rollup.json \
--outfile.genbatchheader $DEVNET/genesis-batch-header.json
touch "$DEVNET/done"
)