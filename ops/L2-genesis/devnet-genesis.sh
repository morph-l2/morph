DEVNET="$PWD/.devnet"
echo "Regenerating genesis files"
# Check if the folder exists
if [ ! -d "$DEVNET" ]; then
  echo "Folder not exists, mkdir $DEVNET"
  mkdir "$DEVNET"
fi

TIMESTAMP=$(date +%s | xargs printf '0x%x')
cat "deploy-config/devnet-deploy-config.json" | jq -r ".l1GenesisBlockTimestamp = \"$TIMESTAMP\"" > $DEVNET/devnet-deploy-config.json
(
go run cmd/main.go genesis devnet \
--deploy-config $DEVNET/devnet-deploy-config.json \
--outfile.l1 $DEVNET/genesis-l1.json \
--outfile.l2 $DEVNET/genesis-l2.json \
--outfile.rollup $DEVNET/rollup.json
touch "$DEVNET/done"
)