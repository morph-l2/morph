GETH_DATA_DIR="${GETH_DATA_DIR:-/db}"
GETH_CHAINDATA_DIR="$GETH_DATA_DIR/geth/chaindata"
GENESIS_FILE_PATH="${GENESIS_FILE_PATH:-/genesis.json}"
JWT_SECRET_PATH="${JWT_SECRET_PATH:-/jwt-secret.txt}"

if [[ ! -e "$GETH_CHAINDATA_DIR" ]]; then
  echo "$GETH_CHAINDATA_DIR missing, running init"
  echo "Initializing genesis."
  geth --verbosity=3 init --datadir="$GETH_DATA_DIR"  "$GENESIS_FILE_PATH"
else
	echo "$GETH_KEYSTORE_DIR exists."
fi

optional_bootnodes=${BOOT_NODES:+"--bootnodes=$BOOT_NODES"}

# shellcheck disable=SC2125
COMMAND="geth \
--datadir="$GETH_DATA_DIR" \
--verbosity=3 \
--http \
--http.corsdomain="*" \
--http.vhosts="*" \
--http.addr=0.0.0.0 \
--http.port=8545 \
--http.api=web3,debug,eth,txpool,net,morph,engine,admin \
--ws \
--ws.addr=0.0.0.0 \
--ws.port=8546 \
--ws.origins="*" \
--ws.api=web3,debug,eth,txpool,net,morph,engine,admin \
--networkid=53077 \
--authrpc.addr="0.0.0.0" \
--authrpc.port="8551" \
--authrpc.vhosts="*" \
--authrpc.jwtsecret=$JWT_SECRET_PATH \
--gcmode=archive \
--nodiscover \
--miner.gasprice="0" \
--metrics \
--metrics.addr=0.0.0.0 $optional_bootnodes"


$COMMAND
