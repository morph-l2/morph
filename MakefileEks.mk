PWD := $(shell pwd)
LAYER1_RPC := http://ethereum.upex30.net:8545
LAYER1_BEACON_RPC := http://ethereum.upex30.net:3500
LAYER2_RPC := http://morph-geth-0:8545
L2_TENDERMINT_RPC := http://morph-node-0:26657
LAYER2_CHAIN_ID := 2818
ROLLUP := 0x1dc010026af6fd4a6dc3686446c752094bda8d4d
START_HEIGHT := 20626146
EXTERNAL_SIGN_URL:= http://morph-proxy-backend:8080/v1/sign/tx_sign

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.GitVersion=$(GITVERSION)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

# gas-oracle
build-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/gas-oracle/app && cargo build --release
	cp gas-oracle/app/target/release/app dist/
	aws s3 cp s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export GAS_ORACLE_EXTERNAL_SIGN=true
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export GAS_ORACLE_EXTERNAL_SIGN_ADDRESS=0xa478a762c79cd3245da45e2a0169810136a697fd
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export GAS_ORACLE_EXTERNAL_SIGN_APPID=gas-oracle-EA07BFE6-5138-4ABA-A78A-16CDAA251D7A
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export GAS_ORACLE_EXTERNAL_SIGN_CHAIN=MORPH
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export GAS_ORACLE_EXTERNAL_SIGN_URL=$(EXTERNAL_SIGN_URL)
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export GAS_ORACLE_L1_RPC=$(LAYER1_RPC)
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export GAS_ORACLE_L2_RPC=$(LAYER2_RPC)
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export GAS_ORACLE_L1_BEACON_RPC=$(LAYER1_BEACON_RPC)
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export GAS_THRESHOLD=5
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export INTERVAL=10000
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export L2_GAS_PRICE_ORACLE=0x530000000000000000000000000000000000000F
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export L2_GAS_ORACLE_PRIVATE_KEY=0x1
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export OVERHEAD_SWITCH=true
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export OVERHEAD_THRESHOLD=200
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export OVERHEAD_INTERVAL=5
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export MAX_OVERHEAD=200000
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export TXN_PER_BLOCK=1
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export TXN_PER_BATCH=50
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export L1_ROLLUP=$(ROLLUP)
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export GAS_ORACLE_L1_BASE_FEE_BUFFER=100000000
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle: export GAS_ORACLE_COMMIT_SCALAR_BUFFER=20000000000
start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle:
	/data/secret-manager-wrapper ./app


# prover
build-bk-prod-morph-prod-mainnet-to-morph-prover:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover && cargo build --release
	cp prover/target/release/prover_server dist/
	cp -r prover/configs dist/
	#aws s3 cp --recursive s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/prove_params/ dist/morph-prover-data/

start-bk-prod-morph-prod-mainnet-to-morph-prover: export PROVER_L1_RPC=$(LAYER1_RPC)
start-bk-prod-morph-prod-mainnet-to-morph-prover: export PROVER_L2_RPC=$(LAYER2_RPC)
start-bk-prod-morph-prod-mainnet-to-morph-prover: export COINBASE=0x530000000000000000000000000000000000000a
start-bk-prod-morph-prod-mainnet-to-morph-prover: export MORPH_MAINNET_CURIE_BLOCK = 0
start-bk-prod-morph-prod-mainnet-to-morph-prover: export SCROLL_PROVER_ASSETS_DIR=/data/configs
start-bk-prod-morph-prod-mainnet-to-morph-prover: export PROVER_PARAMS_DIR=/data/morph-prover-data/prove_params
start-bk-prod-morph-prod-mainnet-to-morph-prover: export PROVER_PROOF_DIR=/data/morph-prover-data/proof
start-bk-prod-morph-prod-mainnet-to-morph-prover: export GENERATE_EVM_VERIFIER=false
start-bk-prod-morph-prod-mainnet-to-morph-prover: export CHAIN_ID=$(LAYER2_CHAIN_ID)
start-bk-prod-morph-prod-mainnet-to-morph-prover:
	#if [ ! -d morph-prover-data/prove_params ]; then aws s3 cp s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/prove_params morph-prover-data/; fi
	#./prover_server
	sleep infinity

# challenge-handler
build-bk-prod-morph-prod-mainnet-to-morph-challenge-handler:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/challenge-handler && cargo build --release
	cp prover/challenge-handler/target/release/challenge-handler dist/
	aws s3 cp s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export HANDLER_EXTERNAL_SIGN=true
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export HANDLER_EXTERNAL_SIGN_ADDRESS=0xe00a22c0a262df40e7a65441441bebc08e248469
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export HANDLER_EXTERNAL_SIGN_APPID=handler-3A43346E-25F3-458B-9E93-88B92ACAA963
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export HANDLER_EXTERNAL_SIGN_CHAIN=ETH
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export HANDLER_EXTERNAL_SIGN_URL=$(EXTERNAL_SIGN_URL)
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export HANDLER_L1_RPC=$(LAYER1_RPC)
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export HANDLER_L2_RPC=$(LAYER2_RPC)
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export HANDLER_PROVER_RPC=http://morph-prover:3030
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export HANDLER_L1_ROLLUP=$(ROLLUP)
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler: export CHALLENGE_HANDLER_PRIVATE_KEY=0x1
start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler:
	/data/secret-manager-wrapper ./challenge-handler


# shadow-proving
build-bk-prod-morph-prod-mainnet-to-morph-shadow-proving:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/shadow-proving && cargo build --release
	cp prover/shadow-proving/target/release/shadow-proving dist/

start-bk-prod-morph-prod-mainnet-to-morph-shadow-proving: export SHADOW_PROVING_L1_RPC=$(LAYER1_RPC)
start-bk-prod-morph-prod-mainnet-to-morph-shadow-proving: export SHADOW_PROVING_VERIFY_L1_RPC=https://ethereum-holesky.publicnode.com
start-bk-prod-morph-prod-mainnet-to-morph-shadow-proving: export SHADOW_PROVING_L2_RPC=$(LAYER2_RPC)
start-bk-prod-morph-prod-mainnet-to-morph-shadow-proving: export SHADOW_PROVING_PROVER_RPC=http://morph-prover:3030
start-bk-prod-morph-prod-mainnet-to-morph-shadow-proving: export SHADOW_PROVING_L1_ROLLUP=$(ROLLUP)
start-bk-prod-morph-prod-mainnet-to-morph-shadow-proving: export SHADOW_PROVING_L1_SHADOW_ROLLUP=0x21c7FCE94d71aFC4e9787362C8c02Ea182520A22
# start-bk-prod-morph-prod-mainnet-to-morph-shadow-proving: export SHADOW_PROVING_PRIVATE_KEY=0x1
start-bk-prod-morph-prod-mainnet-to-morph-shadow-proving:
	/data/secret-manager-wrapper  ./shadow-proving

build-bk-prod-morph-prod-mainnet-to-morph-staking-oracle:
	if [ ! -d dist ]; then mkdir -p dist; fi
	env GO111MODULE=on CGO_LDFLAGS="-ldl" CGO_ENABLED=1 go build -v $(LDFLAGS) -o oracle/staking-oracle ./oracle/cmd/staking-oracle
	cp oracle/staking-oracle dist/
	aws s3 cp s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz


start-bk-prod-morph-prod-mainnet-to-morph-staking-oracle: export STAKING_ORACLE_BUILD_ENV=mainnet
start-bk-prod-morph-prod-mainnet-to-morph-staking-oracle: export STAKING_ORACLE_L1_ETH_RPC=$(LAYER1_RPC)
start-bk-prod-morph-prod-mainnet-to-morph-staking-oracle: export STAKING_ORACLE_L2_ETH_RPC=$(LAYER2_RPC)
start-bk-prod-morph-prod-mainnet-to-morph-staking-oracle: export STAKING_ORACLE_L2_TENDERMINT_RPC=$(L2_TENDERMINT_RPC)
start-bk-prod-morph-prod-mainnet-to-morph-staking-oracle: export STAKING_ORACLE_ROLLUP=$(ROLLUP)
start-bk-prod-morph-prod-mainnet-to-morph-staking-oracle: export STAKING_ORACLE_START_HEIGHT=$(START_HEIGHT)
start-bk-prod-morph-prod-mainnet-to-morph-staking-oracle: export STAKING_ORACLE_LOG_FILENAME=/data/logs/morph-staking-oracle/app_info.log
start-bk-prod-morph-prod-mainnet-to-morph-staking-oracle: export STAKING_ORACLE_LOG_FILE_MAX_SIZE=200
start-bk-prod-morph-prod-mainnet-to-morph-staking-oracle: export STAKING_ORACLE_LOG_FILE_MAX_AGE=30
start-bk-prod-morph-prod-mainnet-to-morph-staking-oracle: export STAKING_ORACLE_LOG_COMPRESS=true
start-bk-prod-morph-prod-mainnet-to-morph-staking-oracle: export STAKING_ORACLE_METRICS_SERVER_ENABLE=true
start-bk-prod-morph-prod-mainnet-to-morph-staking-oracle: export STAKING_ORACLE_EXTERNAL_SIGN=true
start-bk-prod-morph-prod-mainnet-to-morph-staking-oracle: export STAKING_ORACLE_EXTERNAL_SIGN_ADDRESS=0xf720410eeb471e97000afdcfb4e6c515ecf5fa98
start-bk-prod-morph-prod-mainnet-to-morph-staking-oracle: export STAKING_ORACLE_EXTERNAL_SIGN_APPID=staking-oracle-83276E92-0198-4CE0-A990-4835743403CC
start-bk-prod-morph-prod-mainnet-to-morph-staking-oracle: export STAKING_ORACLE_EXTERNAL_SIGN_CHAIN=MORPH
start-bk-prod-morph-prod-mainnet-to-morph-staking-oracle: export STAKING_ORACLE_EXTERNAL_SIGN_URL=$(EXTERNAL_SIGN_URL)
start-bk-prod-morph-prod-mainnet-to-morph-staking-oracle:
	/data/secret-manager-wrapper  ./staking-oracle