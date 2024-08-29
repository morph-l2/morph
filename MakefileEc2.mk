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
