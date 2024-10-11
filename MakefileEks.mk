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


start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle:
	/data/secret-manager-wrapper ./app


# prover
build-bk-prod-morph-prod-mainnet-to-morph-prover:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/server && RUSTFLAGS="-C target-cpu=native -C target-feature=+avx512f" cargo build --release
	cp prover/target/release/prover_server dist/
	cp -r prover/configs dist/

start-bk-prod-morph-prod-mainnet-to-morph-prover:
	#if [ ! -d morph-prover-data/circuits ]; then aws s3 cp s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/sp1-circuits morph-prover-data/; fi
	./prover_server

# challenge-handler
build-bk-prod-morph-prod-mainnet-to-morph-challenge-handler:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/challenge && cargo build --release
	cp prover/bin/challenge/target/release/challenge-handler dist/
	aws s3 cp s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz


start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler:
	/data/secret-manager-wrapper ./challenge-handler


# shadow-proving
build-bk-prod-morph-prod-mainnet-to-morph-shadow-proving:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/shadow-prove && cargo build --release
	cp prover/bin/shadow-prove/target/release/shadow-proving dist/
	aws s3 cp s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

start-bk-prod-morph-prod-mainnet-to-morph-shadow-proving:
	/data/secret-manager-wrapper  ./shadow-proving

build-bk-prod-morph-prod-mainnet-to-morph-staking-oracle:
	if [ ! -d dist ]; then mkdir -p dist; fi
	env GO111MODULE=on CGO_LDFLAGS="-ldl" CGO_ENABLED=1 go build -v $(LDFLAGS) -o oracle/staking-oracle ./oracle/cmd/staking-oracle
	cp oracle/staking-oracle dist/
	aws s3 cp s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

start-bk-prod-morph-prod-mainnet-to-morph-staking-oracle:
	/data/secret-manager-wrapper  ./staking-oracle