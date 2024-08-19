################## update dependencies ####################

ETHEREUM_TARGET_VERSION := v0.5.0
TENDERMINT_TARGET_VERSION := v0.3.0

ETHEREUM_MODULE_NAME := github.com/morph-l2/go-ethereum
TENDERMINT_MODULE_NAME := github.com/morph-l2/tendermint

.PHONY: update_mod
update_mod:
	@echo "Updating go.mod in $(MODULE)..."
	
	@if grep -q '$(ETHEREUM_MODULE_NAME)' $(MODULE)/go.mod; then \
		sed -i '' -e "s|$(ETHEREUM_MODULE_NAME) v[0-9][^[:space:]]*|$(ETHEREUM_MODULE_NAME) $(ETHEREUM_TARGET_VERSION)|" $(MODULE)/go.mod; \
	fi

	@if grep -q '$(TENDERMINT_MODULE_NAME)' $(MODULE)/go.mod; then \
		sed -i '' -e "s|$(TENDERMINT_MODULE_NAME) v[0-9][^[:space:]]*|$(TENDERMINT_MODULE_NAME) $(TENDERMINT_TARGET_VERSION)|" $(MODULE)/go.mod; \
	fi

	@cd $(MODULE) && go mod tidy
	@echo "go.mod in $(MODULE) updated and cleaned."

.PHONY: update_all_mod
update_all_mod:
	@$(MAKE) update_mod MODULE=bindings
	@$(MAKE) update_mod MODULE=contracts
	@$(MAKE) update_mod MODULE=node
	@$(MAKE) update_mod MODULE=ops/l2-genesis
	@$(MAKE) update_mod MODULE=ops/tools
	@$(MAKE) update_mod MODULE=oracle
	@$(MAKE) update_mod MODULE=tx-submitter


update:
	go work sync
	@$(MAKE) update_all_mod
.PHONY: update

submodules:
	git submodule update --init
	git submodule update --remote 
.PHONY: submodules

################## bindings ####################

bindings:
	make -C bindings all
.PHONY: bindings

################## lint code ####################

lint: lint-sol lint-go
.PHONY: lint

lint-sol:
	make -C contracts lint-sol
.PHONY: lint-sol

lint-go:
	make -C bindings lint
	make -C contracts lint-go
	make -C node lint
	make -C ops/l2-genesis lint
	make -C ops/tools lint
	make -C oracle lint
	make -C tx-submitter lint
.PHONY: lint-go

################## format code ####################

fmt: fmt-sol fmt-go
.PHONY: fmt

# npm install --global --save-dev prettier-plugin-solidity
fmt-sol:
	find ./contracts/ -name '*.sol' -type f -not -path "**/node_modules*" | xargs misspell -w
	cd $(PWD)/contracts/ && yarn prettier --write --plugin=prettier-plugin-solidity './contracts/**/*.sol'
.PHONY: fmt-sol

# go get -u github.com/client9/misspell/cmd/misspell
fmt-go:
	go work sync
	cd $(PWD)/bindings/ && go mod tidy
	cd $(PWD)/contracts/ && go mod tidy
	cd $(PWD)/node/ && go mod tidy
	cd $(PWD)/ops/l2-genesis/ && go mod tidy
	cd $(PWD)/ops/tools/ && go mod tidy
	cd $(PWD)/oracle/ && go mod tidy
	cd $(PWD)/tx-submitter/ && go mod tidy
	find . -name '*.go' -type f -not -path "./go-ethereum*" -not -name '*.pb.go' | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./go-ethereum*" -not -name '*.pb.go' | xargs misspell -w
	find . -name '*.go' -type f -not -path "./go-ethereum*" -not -name '*.pb.go' | xargs goimports -w -local $(PWD)/
.PHONY: fmt-go

################## docker build ####################

docker-build:
	cd ops/docker && docker compose build
.PHONY: docker-build

go-rust-builder:
	@if [ -z "$(shell docker images -q morph/go-rust-builder 2> /dev/null)" ]; then \
		echo "Docker image morph/go-rust-builder does not exist. Building..."; \
		cd ops/docker/intermediate && docker build -t morph/go-rust-builder:go-1.22-rust-nightly-2023-12-03 . -f go-rust-builder.Dockerfile; \
	else \
		echo "Docker image morph/go-rust-builder already exists."; \
	fi
.PHONY: go-rust-builder

go-rust-alpine-builder:
	@if [ -z "$(shell docker images -q morph/go-rust-alpine-builder 2> /dev/null)" ]; then \
		echo "Docker image morph/go-rust-alpine-builder does not exist. Building..."; \
		cd ops/docker/intermediate && docker build -t morph/go-rust-alpine-builder:go-1.22-rust-nightly-2023-12-03 . -f go-rust-alpine-builder.Dockerfile; \
	else \
		echo "Docker image morph/go-rust-alpine-builder already exists."; \
	fi
.PHONY: go-rust-alpine-builder

go-ubuntu-builder:
	@if [ -z "$(shell docker images -q morph/go-ubuntu-builder 2> /dev/null)" ]; then \
		echo "Docker image morph/go-ubuntu-builder does not exist. Building..."; \
		cd ops/docker/intermediate && docker build -t morph/go-ubuntu-builder:go-1.22-ubuntu . -f go-ubuntu-builder.Dockerfile; \
	else \
		echo "Docker image morph/go-ubuntu-builder already exists."; \
	fi
.PHONY: go-ubuntu-builder

################## devnet 4 nodes ####################

devnet-up: submodules go-ubuntu-builder go-rust-builder
	python3 ops/devnet-morph/main.py --polyrepo-dir=.
.PHONY: devnet-up

devnet-up-mockccc:
	python3 ops/devnet-morph/main.py --polyrepo-dir=. --mockccc
.PHONY: devnet-up-mockccc

devnet-up-debugccc:
	python3 ops/devnet-morph/main.py --polyrepo-dir=. --debugccc
.PHONY: devnet-up-debugccc

devnet-down:
	cd ops/docker && docker compose -f docker-compose-4nodes.yml down
.PHONY: devnet-down

devnet-clean-build: devnet-down
	docker volume ls --filter name=docker-* --format='{{.Name}}' | xargs -r docker volume rm
	rm -rf ops/l2-genesis/.devnet
	rm -rf ops/docker/consensus/beacondata ops/docker/consensus/validatordata ops/docker/consensus/genesis.ssz
	rm -rf ops/docker/execution/geth
.PHONY: devnet-clean-build

devnet-clean: devnet-clean-build
	docker image ls '*morph*' --format='{{.Repository}}' | xargs -r docker rmi
	docker image ls '*sentry-*' --format='{{.Repository}}' | xargs -r docker rmi
.PHONY: devnet-clean

devnet-l1:
	python3 ops/devnet-morph/main.py --polyrepo-dir=. --only-l1

devnet-logs:
	@(cd ops/docker && docker-compose logs -f)
.PHONY: devnet-logs

# tx-submitter
SUBMITTERS := $(shell grep -o 'tx-submitter-[0-9]*[^:]' ops/docker/docker-compose-4nodes.yml | sort | uniq)
rebuild-all-tx-submitter:
	@for submitter in $(SUBMITTERS); do \
		docker compose -f ./ops/docker/docker-compose-4nodes.yml up -d --build $$submitter --no-deps; \
	done
stop-all-tx-submitter:
	@for submitter in $(SUBMITTERS); do \
		docker compose -f ./ops/docker-compose-4nodes.yml stop $$submitter; \
	done
start-all-tx-submitter:
	@for submitter in $(SUBMITTERS); do \
		docker compose -f ./ops/docker-compose-4nodes.yml start $$submitter; \
	done

# build geth
nccc_geth: submodules
	cd go-ethereum && env GO111MODULE=on GOWORK=off go run build/ci.go install ./cmd/geth
