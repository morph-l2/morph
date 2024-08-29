PWD := $(shell pwd)

# gas-oracle
build-bk-prod-morph-prod-mainnet-to-morph-gas-oracle:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/gas-oracle/app && cargo build --release
	cp gas-oracle/app/target/release/app dist/

start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export GAS_ORACLE_EXTERNAL_SIGN=true
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export GAS_ORACLE_EXTERNAL_SIGN_ADDRESS=0x0000000000000000000000000000000000000000
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export GAS_ORACLE_EXTERNAL_SIGN_APPID=xxxx
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export GAS_ORACLE_EXTERNAL_SIGN_CHAIN=MAINNET-L1
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export GAS_ORACLE_EXTERNAL_SIGN_URL=http://127.0.0.1:8080/v1/sign/tx_sign
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export GAS_ORACLE_EXTERNAL_SIGN_RSA_PRIV=xxxx
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export GAS_ORACLE_L1_RPC=http://morph-l1-geth:8545
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export GAS_ORACLE_L2_RPC=http://morph-geth-0:8545
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export GAS_ORACLE_L1_BEACON_RPC=$(L1_BEACON_RPC)
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export GAS_THRESHOLD=5
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export INTERVAL=10000
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export L2_GAS_PRICE_ORACLE=0x530000000000000000000000000000000000000F
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export L2_GAS_ORACLE_PRIVATE_KEY=0x1
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export OVERHEAD_SWITCH=true
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export OVERHEAD_THRESHOLD=200
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export OVERHEAD_INTERVAL=5
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export MAX_OVERHEAD=200000
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export TXN_PER_BLOCK=1
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export TXN_PER_BATCH=50
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export L1_ROLLUP=$(Rollup)
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export GAS_ORACLE_L1_BASE_FEE_BUFFER=100000000
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle: export GAS_ORACLE_COMMIT_SCALAR_BUFFER=20000000000
start-bk-prod-morph-prod-mainnet-to-morph-gas-oracle:
	app


# prover
build-bk-prod-morph-prod-mainnet-to-morph-prover:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover && cargo build --release
	cp prover/target/release/prover_server dist/
	cp -r prover/configs dist/

start-bk-prod-morph-prod-mainnet-to-morph-prover: export PROVER_L1_RPC=$(L1_RPC)
start-bk-prod-morph-prod-mainnet-to-morph-prover: export PROVER_L2_RPC=$(L2_RPC)
start-bk-prod-morph-prod-mainnet-to-morph-prover: export COINBASE=0x530000000000000000000000000000000000000a
start-bk-prod-morph-prod-mainnet-to-morph-prover: export MORPH_MAINNET_CURIE_BLOCK = 0
start-bk-prod-morph-prod-mainnet-to-morph-prover: export SCROLL_PROVER_ASSETS_DIR=/data/morph-prover/make-conf/configs
start-bk-prod-morph-prod-mainnet-to-morph-prover: export PROVER_PARAMS_DIR=/data/morph-prover-data/prove_params
start-bk-prod-morph-prod-mainnet-to-morph-prover: export PROVER_PROOF_DIR=/data/morph-prover-data/proof
start-bk-prod-morph-prod-mainnet-to-morph-prover: export GENERATE_EVM_VERIFIER=false
start-bk-prod-morph-prod-mainnet-to-morph-prover: export CHAIN_ID=$(L2_CHAIN_ID)
start-bk-prod-morph-prod-mainnet-to-morph-prover:
	prover_server


# challenge-handler
build-bk-prod-morph-prod-mainnet-to-morph-challenge-handler:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/challenge-handler && cargo build --release
	cp prover/challenge-handler/target/release/challenge-handler dist/

start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export HANDLER_EXTERNAL_SIGN=true
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export HANDLER_EXTERNAL_SIGN_ADDRESS=0x0000000000000000000000000000000000000000
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export HANDLER_EXTERNAL_SIGN_APPID=xxxx
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export HANDLER_EXTERNAL_SIGN_CHAIN=MAINNET-L1
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export HANDLER_EXTERNAL_SIGN_URL=http://127.0.0.1:8080/v1/sign/tx_sign
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export HANDLER_EXTERNAL_SIGN_RSA_PRIV=xxxxx
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export HANDLER_L1_RPC=$(L1_RPC)
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export HANDLER_L2_RPC=$(L2_RPC)
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export HANDLER_PROVER_RPC=http://morph-prover:3030
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export HANDLER_L1_ROLLUP=$(Rollup)
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export CHALLENGE_HANDLER_PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler:
	challenge-handler


# shadow-proving
build-bk-prod-morph-prod-mainnet-to-morph-shadow-proving:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/shadow-proving && cargo build --release
	cp prover/shadow-proving/target/release/shadow-proving dist/

start-bk-prod-morph-prod-mainnet-to-morph-shadow-proving: export SHADOW_PROVING_L1_RPC=$(L1_RPC)
start-bk-prod-morph-prod-mainnet-to-morph-shadow-proving: export SHADOW_PROVING_L2_RPC=$(L2_RPC)
start-bk-prod-morph-prod-mainnet-to-morph-shadow-proving: export SHADOW_PROVING_PROVER_RPC=http://morph-prover:3030
start-bk-prod-morph-prod-mainnet-to-morph-shadow-proving: export SHADOW_PROVING_L1_ROLLUP=$(Rollup)
start-bk-prod-morph-prod-mainnet-to-morph-shadow-proving: export SHADOW_PROVING_L1_SHADOW_ROLLUP=0x0000000000000000000000000000000000000000
start-bk-prod-morph-prod-mainnet-to-morph-shadow-proving: export SHADOW_PROVING_PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
start-bk-prod-morph-prod-mainnet-to-morph-shadow-proving:
	shadow-proving