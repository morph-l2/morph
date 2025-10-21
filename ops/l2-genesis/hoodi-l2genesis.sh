#!/bin/sh
set -eu

source ../../contracts/.env
FOLDER=".hoodi"
CONTRACT_CONFIG="../../contracts/src/deploy-config/hoodi.ts"
echo "Regenerating genesis files"
# Check if the folder exists
if [ ! -d "$FOLDER" ]; then
  echo "Folder not exists, mkdir $FOLDER"
  mkdir "$FOLDER"
fi


RPC_URL="https://rpc.hoodi.ethpandaops.io"

echo "RPC_URL is $RPC_URL"
cat "deploy-config/hoodi-deploy-config.json" > $FOLDER/hoodi-deploy-config.json
(
go run cmd/main.go genesis l2 \
--l1-rpc $RPC_URL \
--deploy-config $FOLDER/hoodi-deploy-config.json \
--deployment-dir "$PWD/../../contracts/hoodi.json" \
--outfile.l2 $FOLDER/genesis-l2.json \
--outfile.rollup $FOLDER/rollup.json \
--outfile.genbatchheader $FOLDER/genesis-batch-header.json
touch "$FOLDER/done"
)
