GETH_DATA_DIR="${GETH_DATA_DIR:-/db}"
GETH_CHAINDATA_DIR="$GETH_DATA_DIR/geth/chaindata"
GENESIS_FILE_PATH="${GENESIS_FILE_PATH:-/genesis.json}"
JWT_SECRET_PATH="${JWT_SECRET_PATH:-/jwt-secret.txt}"
DEFAULE_MINER_ETHERBASE="0x0e87cd091e091562F25CB1cf4641065dA2C049F5"

if [[ ! -e "$GETH_CHAINDATA_DIR" ]]; then
  echo "$GETH_CHAINDATA_DIR missing, running init"
  echo "Initializing genesis."
  geth --verbosity=3 init --datadir="$GETH_DATA_DIR"  "$GENESIS_FILE_PATH"
else
	echo "$GETH_KEYSTORE_DIR exists."
fi

if [[ -z "$MINER_ETHERBASE" ]]; then
  # the environment variable is missing, set a default value
  MINER_ETHERBASE=$DEFAULE_MINER_ETHERBASE
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
--nodiscover
--mine \
--miner.gasprice="100000000" \
--miner.etherbase=$MINER_ETHERBASE \
--miner.gaslimit=10000000 \
--metrics \
--metrics.addr=0.0.0.0 $optional_bootnodes"


$COMMAND
