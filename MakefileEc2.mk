TIMESTAMP := $(shell date +%s | xargs printf '0x%x')
PWD := $(shell pwd)
GORUN = env GO111MODULE=on go run
GITVERSION := v1.0.0
LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.GitVersion=$(GITVERSION)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

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

build-bk-prod-morph-prod-mainnet-to-morph-gas-oracle:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/gas-oracle/app && cargo build --release
	cp gas-oracle/app/target/release/app dist/
	tar -czvf gas-oracle.tar.gz dist
	aws s3 cp gas-oracle.tar.gz s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/gas-oracle.tar.gz

build-bk-prod-morph-prod-mainnet-to-morph-prover:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover && cargo build --release
	cp prover/target/release/prover_server dist/
	cp -r prover/configs dist/
	tar -czvf prover.tar.gz dist
	aws s3 cp prover.tar.gz s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/prover.tar.gz

build-bk-prod-morph-prod-mainnet-to-morph-challenge-handler:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/challenge-handler && cargo build --release
	cp prover/challenge-handler/target/release/challenge-handler dist/
	tar -czvf challenge-handler.tar.gz dist
	aws s3 cp challenge-handler.tar.gz s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/challenge-handler.tar.gz

build-bk-prod-morph-prod-mainnet-to-morph-shadow-proving:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/shadow-proving && cargo build --release
	cp prover/shadow-proving/target/release/shadow-proving dist/
	tar -czvf shadow-proving.tar.gz dist
	aws s3 cp shadow-proving.tar.gz s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/shadow-proving.tar.gz