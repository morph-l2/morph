# morph-deployer

## Compile Smart Contract

### step 1 build solidity files

checkout the contract repo: ``git clone https://github.com/morph-l2/morph``

```cd contracts```

Install Foundry with a specific version.

```foundryup -C da2392e58bb8a7fefeba46b40c4df1afad8ccd22```

Install node modules with yarn (v1) and Node.js (16+):

Build Contract

```yarn install```

Make sure artifacts are generated in ./artifacts directory

### step 2 generate go-bindings

To be developed standalone project bindings  
checkout this repo ```git clone https://github.com/morph-l2/morph.git```.  

```shell
cd morph/bindings
make all
```

## Prepare Genesis File

Run script `devnet-gensis.sh`

```shell
DEVNET="$PWD/.devnet"
echo "Regenerating genesis files"
TIMESTAMP=$(date +%s | xargs printf '0x%x')
cat "deploy-config/devnetL1.json" | jq -r ".l1GenesisBlockTimestamp = \"$TIMESTAMP\"" > /tmp/bedrock-devnet-deploy-config.json
(
go run cmd/main.go genesis devnet \
--deploy-config /tmp/bedrock-devnet-deploy-config.json \
--outfile.l1 $DEVNET/genesis-l1.json \
--outfile.l2 $DEVNET/genesis-l2.json \
--outfile.rollup $DEVNET/rollup.json
touch "$DEVNET/done"
)
```

Make sure `genesis-l1.json` and `genesis-l2.json` are generated correctly.
