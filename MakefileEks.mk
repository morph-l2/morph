PWD := $(shell pwd)

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.GitVersion=$(GITVERSION)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

# gas-oracle
# mainnet
build-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/gas-oracle/app && cargo build --release
	cp gas-oracle/app/target/release/app dist/
	aws s3 cp s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz


start-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle:
	/data/secret-manager-wrapper ./app


# qanet
build-bk-test-morph-test-qanet-to-morph-gas-price-oracle-qanet:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/gas-oracle/app && cargo build --release
	cp gas-oracle/app/target/release/app dist/
	aws s3 cp s3://morph-7637-morph-technical-department-qanet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz


start-bk-test-morph-test-qanet-to-morph-gas-price-oracle-qanet:
	/data/secret-manager-wrapper ./app


# prover
# mainnet
build-bk-prod-morph-prod-mainnet-to-morph-prover:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/server && CARGO_NET_GIT_FETCH_WITH_CLI=true cargo build --release
	cp prover/target/release/prover-server dist/
	aws s3 cp s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

start-bk-prod-morph-prod-mainnet-to-morph-prover:
	/data/secret-manager-wrapper ./prover-server

# testnet
build-bk-prod-morph-prod-testnet-to-morph-prover-hoodi:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/server && CARGO_NET_GIT_FETCH_WITH_CLI=true cargo build --release
	cp prover/target/release/prover-server dist/
	aws s3 cp s3://morph-0582-morph-technical-department-testnet-data/testnet/hoodi/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

start-bk-prod-morph-prod-testnet-to-morph-prover-hoodi:
	/data/secret-manager-wrapper ./prover-server

# qanet
build-bk-test-morph-test-qanet-to-morph-prover:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/server && CARGO_NET_GIT_FETCH_WITH_CLI=true cargo build --release
	cp prover/target/release/prover-server dist/
	aws s3 cp s3://morph-7637-morph-technical-department-qanet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

start-bk-test-morph-test-qanet-to-morph-prover:
	/data/secret-manager-wrapper ./prover-server

# challenge-handler
# mainnet
build-bk-prod-morph-prod-mainnet-to-morph-challenge-handler:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/challenge && CARGO_NET_GIT_FETCH_WITH_CLI=true cargo build --release
	cp prover/bin/challenge/target/release/challenge-handler dist/
	aws s3 cp s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz


start-bk-prod-morph-prod-mainnet-to-morph-challenge-handler:
	/data/secret-manager-wrapper ./challenge-handler

# testnet
build-bk-prod-morph-prod-testnet-to-morph-challenge-handler-hoodi:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/challenge && CARGO_NET_GIT_FETCH_WITH_CLI=true cargo build --release
	cp prover/bin/challenge/target/release/challenge-handler dist/
	aws s3 cp s3://morph-0582-morph-technical-department-testnet-data/testnet/hoodi/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

start-bk-prod-morph-prod-testnet-to-morph-challenge-handler-hoodi:
	/data/secret-manager-wrapper ./challenge-handler

# qanet
build-bk-test-morph-test-qanet-to-morph-challenge-handler:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/challenge && CARGO_NET_GIT_FETCH_WITH_CLI=true cargo build --release
	cp prover/bin/challenge/target/release/challenge-handler dist/
	aws s3 cp s3://morph-7637-morph-technical-department-qanet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

start-bk-test-morph-test-qanet-to-morph-challenge-handler:
	/data/secret-manager-wrapper ./challenge-handler

# shadow-proving
# mainnet
build-bk-prod-morph-prod-mainnet-to-morph-shadow-proving:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/shadow-prove && CARGO_NET_GIT_FETCH_WITH_CLI=true cargo build --release
	cp prover/target/release/shadow-proving dist/
	aws s3 cp s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

start-bk-prod-morph-prod-mainnet-to-morph-shadow-proving:
	/data/secret-manager-wrapper  ./shadow-proving

# testnet
build-bk-prod-morph-prod-testnet-to-morph-shadow-proving-hoodi:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/shadow-prove && CARGO_NET_GIT_FETCH_WITH_CLI=true cargo build --release
	cp prover/target/release/shadow-proving dist/
	aws s3 cp s3://morph-0582-morph-technical-department-testnet-data/testnet/hoodi/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

start-bk-prod-morph-prod-testnet-to-morph-shadow-proving-hoodi:
	/data/secret-manager-wrapper ./shadow-proving

# qanet
build-bk-test-morph-test-qanet-to-morph-shadow-proving:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/shadow-prove && CARGO_NET_GIT_FETCH_WITH_CLI=true cargo build --release
	cp prover/target/release/shadow-proving dist/
	aws s3 cp s3://morph-7637-morph-technical-department-qanet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

start-bk-test-morph-test-qanet-to-morph-shadow-proving:
	/data/secret-manager-wrapper ./shadow-proving


# staking-oracle
# mainnet
build-bk-prod-morph-prod-mainnet-to-morph-staking-oracle:
	if [ ! -d dist ]; then mkdir -p dist; fi
	env GO111MODULE=on CGO_LDFLAGS="-ldl" CGO_ENABLED=1 go build -v $(LDFLAGS) -o oracle/staking-oracle ./oracle/cmd/staking-oracle
	cp oracle/staking-oracle dist/
	aws s3 cp s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

start-bk-prod-morph-prod-mainnet-to-morph-staking-oracle:
	/data/secret-manager-wrapper  ./staking-oracle

# qanet
build-bk-test-morph-test-qanet-to-morph-staking-oracle-qanet:
	if [ ! -d dist ]; then mkdir -p dist; fi
	env GO111MODULE=on CGO_LDFLAGS="-ldl" CGO_ENABLED=1 go build -v $(LDFLAGS) -o oracle/staking-oracle ./oracle/cmd/staking-oracle
	cp oracle/staking-oracle dist/
	aws s3 cp s3://morph-7637-morph-technical-department-qanet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

start-bk-test-morph-test-qanet-to-morph-staking-oracle-qanet:
	/data/secret-manager-wrapper  ./staking-oracle

# qanet
build-bk-test-morph-test-qanet-to-morph-token-price-oracle:
	if [ ! -d dist ]; then mkdir -p dist; fi
	env GO111MODULE=on CGO_LDFLAGS="-ldl" CGO_ENABLED=1 go build -v $(LDFLAGS) -o token-price-oracle/token-price-oracle ./token-price-oracle/cmd
	cp token-price-oracle/token-price-oracle dist/
	aws s3 cp s3://morph-7637-morph-technical-department-qanet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

start-bk-test-morph-test-qanet-to-morph-token-price-oracle:
	 /data/secret-manager-wrapper  ./token-price-oracle

# hoodi token price oracle
build-bk-prod-morph-prod-testnet-to-morph-token-price-oracle-hoodi:
	if [ ! -d dist ]; then mkdir -p dist; fi
	env GO111MODULE=on CGO_LDFLAGS="-ldl" CGO_ENABLED=1 go build -v $(LDFLAGS) -o token-price-oracle/token-price-oracle ./token-price-oracle/cmd
	cp token-price-oracle/token-price-oracle dist/
	aws s3 cp s3://morph-0582-morph-technical-department-testnet-data/testnet/hoodi/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz


start-bk-prod-morph-prod-testnet-to-morph-token-price-oracle-hoodi:
	/data/secret-manager-wrapper ./token-price-oracle


build-bk-prod-morph-prod-mainnet-to-morph-token-price-oracle:
	if [ ! -d dist ]; then mkdir -p dist; fi
	env GO111MODULE=on CGO_LDFLAGS="-ldl" CGO_ENABLED=1 go build -v $(LDFLAGS) -o token-price-oracle/token-price-oracle ./token-price-oracle/cmd
	cp token-price-oracle/token-price-oracle dist/
	aws s3 cp s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz


start-bk-prod-morph-prod-mainnet-to-morph-token-price-oracle:
	/data/secret-manager-wrapper ./token-price-oracle


# gas-oracle
# hoodi
build-bk-prod-morph-prod-testnet-to-morph-gas-price-oracle-hoodi:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/gas-oracle/app && cargo build --release
	cp gas-oracle/app/target/release/app dist/
	aws s3 cp s3://morph-0582-morph-technical-department-testnet-data/testnet/hoodi/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz


start-bk-prod-morph-prod-testnet-to-morph-gas-price-oracle-hoodi:
	/data/secret-manager-wrapper ./app


# staking-oracle
# testnet
build-bk-prod-morph-prod-testnet-to-morph-staking-oracle-holesky:
	if [ ! -d dist ]; then mkdir -p dist; fi
	env GO111MODULE=on CGO_LDFLAGS="-ldl" CGO_ENABLED=1 go build -v $(LDFLAGS) -o oracle/staking-oracle ./oracle/cmd/staking-oracle
	cp oracle/staking-oracle dist/
	aws s3 cp s3://morph-0582-morph-technical-department-testnet-data/testnet/holesky/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

start-bk-prod-morph-prod-testnet-to-morph-staking-oracle-holesky:
	/data/secret-manager-wrapper  ./staking-oracle