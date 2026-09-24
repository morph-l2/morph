# New release system: MakefileBuild.mk only produces build artifacts.
# Service start / healthcheck commands live in MakefileEc2.mk / MakefileEks.mk.

TIMESTAMP := $(shell date +%s | xargs printf '0x%x')
PWD := $(shell pwd)
GORUN = env GO111MODULE=on go run
GITVERSION := v1.0.0
LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.GitVersion=$(GITVERSION)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

# ─── morph-node / tx-submitter (formerly MakefileEc2.mk) ─────────────────────

build-bk-prod-morph-prod-mainnet-to-morph-node:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/node && make build
	cp node/build/bin/morphnode dist/
	cp node/build/bin/tendermint dist/
	tar -czvf morph-node.tar.gz dist
	aws s3 cp morph-node.tar.gz s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/morph-node.tar.gz

build-bk-prod-morph-prod-mainnet-to-morph-tx-submitter:
	if [ ! -d dist ]; then mkdir -p dist; fi
	env GO111MODULE=on CGO_LDFLAGS="-ldl" CGO_ENABLED=1 go build -v $(LDFLAGS) -o tx-submitter/tx-submitter ./tx-submitter/cmd
	cp tx-submitter/tx-submitter dist/
	tar -czvf tx-submitter.tar.gz dist
	aws s3 cp tx-submitter.tar.gz s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/tx-submitter.tar.gz

# build for qanet
build-bk-test-morph-test-qanet-to-morph-node-qanet:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/node && make build
	cp node/build/bin/morphnode dist/
	cp node/build/bin/tendermint dist/
	tar -czvf morph-node.tar.gz dist
	aws s3 cp morph-node.tar.gz s3://morph-7637-morph-technical-department-qanet-data/morph-setup/morph-node.tar.gz

build-bk-test-morph-test-qanet-to-morph-tx-submitter-qanet:
	if [ ! -d dist ]; then mkdir -p dist; fi
	env GO111MODULE=on CGO_LDFLAGS="-ldl" CGO_ENABLED=1 go build -v $(LDFLAGS) -o tx-submitter/tx-submitter ./tx-submitter/cmd
	cp tx-submitter/tx-submitter dist/
	tar -czvf tx-submitter.tar.gz dist
	aws s3 cp tx-submitter.tar.gz s3://morph-7637-morph-technical-department-qanet-data/morph-setup/tx-submitter.tar.gz

 # build for hoodi
build-bk-prod-morph-prod-testnet-to-morph-node-hoodi:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/node && make build
	cp node/build/bin/morphnode dist/
	cp node/build/bin/tendermint dist/
	tar -czvf morph-node.tar.gz dist
	aws s3 cp morph-node.tar.gz s3://morph-0582-morph-technical-department-testnet-data/testnet/hoodi/morph-setup/morph-node.tar.gz

build-bk-prod-morph-prod-testnet-to-morph-tx-submitter-hoodi:
	if [ ! -d dist ]; then mkdir -p dist; fi
	env GO111MODULE=on CGO_LDFLAGS="-ldl" CGO_ENABLED=1 go build -v $(LDFLAGS) -o tx-submitter/tx-submitter ./tx-submitter/cmd
	cp tx-submitter/tx-submitter dist/
	tar -czvf tx-submitter.tar.gz dist
	aws s3 cp tx-submitter.tar.gz s3://morph-0582-morph-technical-department-testnet-data/testnet/hoodi/morph-setup/tx-submitter.tar.gz

# ─── gas-oracle / prover / challenge-handler / shadow-proving /
#     token-price-oracle (formerly MakefileEks.mk build targets) ─────────────

# gas-oracle
# mainnet
build-bk-prod-morph-prod-mainnet-to-morph-gas-price-oracle:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/gas-oracle/app && cargo build --release
	cp gas-oracle/app/target/release/app dist/
	aws s3 cp s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

# qanet
build-bk-test-morph-test-qanet-to-morph-gas-price-oracle-qanet:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/gas-oracle/app && cargo build --release
	cp gas-oracle/app/target/release/app dist/
	aws s3 cp s3://morph-7637-morph-technical-department-qanet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

# gas-oracle
# hoodi
build-bk-prod-morph-prod-testnet-to-morph-gas-price-oracle-hoodi:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/gas-oracle/app && cargo build --release
	cp gas-oracle/app/target/release/app dist/
	aws s3 cp s3://morph-0582-morph-technical-department-testnet-data/testnet/hoodi/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

# prover
# mainnet
build-bk-prod-morph-prod-mainnet-to-morph-prover:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/server && CARGO_NET_GIT_FETCH_WITH_CLI=true cargo build --release
	cp prover/target/release/prover-server dist/
	aws s3 cp s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

# testnet
build-bk-prod-morph-prod-testnet-to-morph-prover-hoodi:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/server && CARGO_NET_GIT_FETCH_WITH_CLI=true cargo build --release
	cp prover/target/release/prover-server dist/
	aws s3 cp s3://morph-0582-morph-technical-department-testnet-data/testnet/hoodi/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

# qanet
build-bk-test-morph-test-qanet-to-morph-prover:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/server && CARGO_NET_GIT_FETCH_WITH_CLI=true cargo build --release
	cp prover/target/release/prover-server dist/
	aws s3 cp s3://morph-7637-morph-technical-department-qanet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

# challenge-handler
# mainnet
build-bk-prod-morph-prod-mainnet-to-morph-challenge-handler:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/challenge && CARGO_NET_GIT_FETCH_WITH_CLI=true cargo build --release
	cp prover/bin/challenge/target/release/challenge-handler dist/
	aws s3 cp s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

# testnet
build-bk-prod-morph-prod-testnet-to-morph-challenge-handler-hoodi:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/challenge && CARGO_NET_GIT_FETCH_WITH_CLI=true cargo build --release
	cp prover/bin/challenge/target/release/challenge-handler dist/
	aws s3 cp s3://morph-0582-morph-technical-department-testnet-data/testnet/hoodi/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

# qanet
build-bk-test-morph-test-qanet-to-morph-challenge-handler:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/challenge && CARGO_NET_GIT_FETCH_WITH_CLI=true cargo build --release
	cp prover/bin/challenge/target/release/challenge-handler dist/
	aws s3 cp s3://morph-7637-morph-technical-department-qanet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

# shadow-proving
# mainnet
build-bk-prod-morph-prod-mainnet-to-morph-shadow-proving:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/shadow-prove && CARGO_NET_GIT_FETCH_WITH_CLI=true cargo build --release
	cp prover/target/release/shadow-proving dist/
	aws s3 cp s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

# testnet
build-bk-prod-morph-prod-testnet-to-morph-shadow-proving-hoodi:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/shadow-prove && CARGO_NET_GIT_FETCH_WITH_CLI=true cargo build --release
	cp prover/target/release/shadow-proving dist/
	aws s3 cp s3://morph-0582-morph-technical-department-testnet-data/testnet/hoodi/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

# qanet
build-bk-test-morph-test-qanet-to-morph-shadow-proving:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/bin/shadow-prove && CARGO_NET_GIT_FETCH_WITH_CLI=true cargo build --release
	cp prover/target/release/shadow-proving dist/
	aws s3 cp s3://morph-7637-morph-technical-department-qanet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

# token-price-oracle
# qanet
build-bk-test-morph-test-qanet-to-morph-token-price-oracle:
	if [ ! -d dist ]; then mkdir -p dist; fi
	env GO111MODULE=on CGO_LDFLAGS="-ldl" CGO_ENABLED=1 go build -v $(LDFLAGS) -o token-price-oracle/token-price-oracle ./token-price-oracle/cmd
	cp token-price-oracle/token-price-oracle dist/
	aws s3 cp s3://morph-7637-morph-technical-department-qanet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

# hoodi token price oracle
build-bk-prod-morph-prod-testnet-to-morph-token-price-oracle-hoodi:
	if [ ! -d dist ]; then mkdir -p dist; fi
	env GO111MODULE=on CGO_LDFLAGS="-ldl" CGO_ENABLED=1 go build -v $(LDFLAGS) -o token-price-oracle/token-price-oracle ./token-price-oracle/cmd
	cp token-price-oracle/token-price-oracle dist/
	aws s3 cp s3://morph-0582-morph-technical-department-testnet-data/testnet/hoodi/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz

# mainnet token price oracle
build-bk-prod-morph-prod-mainnet-to-morph-token-price-oracle:
	if [ ! -d dist ]; then mkdir -p dist; fi
	env GO111MODULE=on CGO_LDFLAGS="-ldl" CGO_ENABLED=1 go build -v $(LDFLAGS) -o token-price-oracle/token-price-oracle ./token-price-oracle/cmd
	cp token-price-oracle/token-price-oracle dist/
	aws s3 cp s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/secret-manager-wrapper.tar.gz ./
	tar -xvzf secret-manager-wrapper.tar.gz
