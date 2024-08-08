COMMAND="geth \
--morph-holesky \
--datadir="./db" \
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
--authrpc.addr=0.0.0.0 \
--authrpc.port=8551 \
--authrpc.vhosts="*" \
--authrpc.jwtsecret="./jwt-secret.txt" \
--gcmode=archive \
--log.filename=./geth.log \
--metrics \
--metrics.addr=0.0.0.0"

$COMMAND
