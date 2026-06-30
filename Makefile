################## update dependencies ####################
ETHEREUM_SUBMODULE_COMMIT_OR_TAG := efdd4dbee28b2fb4005267a7793d607111d9490b
ETHEREUM_TARGET_VERSION := v1.10.14-0.20260629084315-efdd4dbee28b
TENDERMINT_TARGET_VERSION := v0.3.8-0.20260625105428-a5de063445a3


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
	@$(MAKE) update_mod MODULE=common
	@$(MAKE) update_mod MODULE=contracts
	@$(MAKE) update_mod MODULE=node
	@$(MAKE) update_mod MODULE=ops/l2-genesis
	@$(MAKE) update_mod MODULE=ops/tools
	@$(MAKE) update_mod MODULE=tx-submitter
	@$(MAKE) update_mod MODULE=token-price-oracle


update:
	go work sync
	@$(MAKE) update_all_mod
.PHONY: update

submodules:
	git submodule update --init
	@if [ -d "go-ethereum" ]; then \
		echo "Updating go-ethereum submodule to tag $(ETHEREUM_SUBMODULE_COMMIT_OR_TAG)..."; \
		cd go-ethereum && \
		git fetch --tags && \
		git checkout $(ETHEREUM_SUBMODULE_COMMIT_OR_TAG) && \
		cd ..; \
	fi
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
	go work sync
	make -C bindings lint
	make -C contracts lint-go
	make -C node lint
	make -C ops/l2-genesis lint
	make -C ops/tools lint
##	make -C oracle lint
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
		cd ops/docker/intermediate && docker build -t morph/go-ubuntu-builder:go-1.24-ubuntu . -f go-ubuntu-builder.Dockerfile; \
	else \
		echo "Docker image morph/go-ubuntu-builder already exists."; \
	fi
.PHONY: go-ubuntu-builder

################## devnet 2 nodes ####################

EXECUTION_CLIENT ?= geth
DEVNET_CLUSTER ?= false
DEVNET_CLUSTER_ENABLED := $(filter true 1 yes,$(DEVNET_CLUSTER))
DEVNET_SEQUENCER_PRIVATE_KEY ?= 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
DEVNET_SEQUENCER_ADDRESS ?= 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
DEVNET_SEQUENCER_UPGRADE_OFFSET_SECONDS ?= 0
MORPH_RETH_BUILD_FROM_SOURCE ?= false
ifeq ($(MORPH_RETH_BUILD_FROM_SOURCE),true)
MORPH_RETH_IMAGE ?= morph-reth:latest
MORPH_RETH_ENTRYPOINT ?= /app/morph-reth
else
MORPH_RETH_IMAGE ?= ghcr.io/morph-l2/morph-reth:latest
MORPH_RETH_ENTRYPOINT ?= /usr/local/bin/morph-reth
endif
MORPH_RETH_DIR ?= ../morph-reth
MORPH_RETH_BUILD_PROFILE ?= release
MORPH_RETH_RUSTFLAGS ?=
MORPH_RETH_DOCKER_TARGET ?= builder
export MORPH_RETH_IMAGE
export MORPH_RETH_DIR
export MORPH_RETH_BUILD_PROFILE
export MORPH_RETH_RUSTFLAGS
export MORPH_RETH_DOCKER_TARGET
export MORPH_RETH_ENTRYPOINT
DEVNET_COMPOSE_FILES := -f docker-compose-devnet.yml
DEVNET_CLEAN_COMPOSE_FILES := -f docker-compose-devnet.yml -f docker-compose-reth.yml -f docker-compose-cluster.yml

ifeq ($(EXECUTION_CLIENT),geth)
DEVNET_EXECUTION_DEPS := submodules
else ifeq ($(EXECUTION_CLIENT),reth)
DEVNET_COMPOSE_FILES += -f docker-compose-reth.yml
ifeq ($(MORPH_RETH_BUILD_FROM_SOURCE),true)
DEVNET_EXECUTION_DEPS := reth
else
DEVNET_EXECUTION_DEPS := reth-image
endif
else
$(error unsupported EXECUTION_CLIENT "$(EXECUTION_CLIENT)", expected "geth" or "reth")
endif
ifneq ($(DEVNET_CLUSTER_ENABLED),)
DEVNET_COMPOSE_FILES += -f docker-compose-cluster.yml
endif

devnet-up: $(DEVNET_EXECUTION_DEPS) go-ubuntu-builder
	python3 ops/devnet-morph/main.py --polyrepo-dir=. --execution-client=$(EXECUTION_CLIENT) \
		$(if $(DEVNET_CLUSTER_ENABLED),--cluster,) \
		--sequencer-private-key=$(DEVNET_SEQUENCER_PRIVATE_KEY) \
		--sequencer-address=$(DEVNET_SEQUENCER_ADDRESS) \
		--sequencer-upgrade-offset-seconds=$(DEVNET_SEQUENCER_UPGRADE_OFFSET_SECONDS)
.PHONY: devnet-up

devnet-up-cluster:
	$(MAKE) devnet-up DEVNET_CLUSTER=true
.PHONY: devnet-up-cluster

devnet-up-reth:
	$(MAKE) devnet-up EXECUTION_CLIENT=reth
.PHONY: devnet-up-reth

devnet-up-cluster-reth:
	$(MAKE) devnet-up EXECUTION_CLIENT=reth DEVNET_CLUSTER=true
.PHONY: devnet-up-cluster-reth

devnet-up-debugccc: $(DEVNET_EXECUTION_DEPS) go-ubuntu-builder
	python3 ops/devnet-morph/main.py --polyrepo-dir=. --execution-client=$(EXECUTION_CLIENT) --debugccc \
		$(if $(DEVNET_CLUSTER_ENABLED),--cluster,) \
		--sequencer-private-key=$(DEVNET_SEQUENCER_PRIVATE_KEY) \
		--sequencer-address=$(DEVNET_SEQUENCER_ADDRESS) \
		--sequencer-upgrade-offset-seconds=$(DEVNET_SEQUENCER_UPGRADE_OFFSET_SECONDS)
.PHONY: devnet-up-debugccc

devnet-down:
	cd ops/docker && docker compose $(DEVNET_COMPOSE_FILES) down
.PHONY: devnet-down

devnet-down-reth:
	$(MAKE) devnet-down EXECUTION_CLIENT=reth
.PHONY: devnet-down-reth

devnet-clean-build: devnet-l1-clean
	cd ops/docker && docker compose $(DEVNET_CLEAN_COMPOSE_FILES) down --volumes --remove-orphans
	docker volume ls --filter label=com.docker.compose.project=docker --format='{{.Name}}' | xargs docker volume rm 2>/dev/null || true
	rm -rf ops/l2-genesis/.devnet
	rm -rf ops/docker/.devnet
	rm -rf ops/docker/consensus ops/docker/execution
.PHONY: devnet-clean-build

devnet-clean: devnet-clean-build
	docker image ls '*morph*' --format='{{.Repository}}' | xargs -r docker rmi
	docker image ls '*sentry-*' --format='{{.Repository}}' | xargs -r docker rmi
.PHONY: devnet-clean

devnet-clean-reth:
	$(MAKE) devnet-clean EXECUTION_CLIENT=reth
.PHONY: devnet-clean-reth

devnet-l1:
	python3 ops/devnet-morph/main.py --polyrepo-dir=. --only-l1

devnet-l1-clean:
	@cd ops/docker && ./layer1/scripts/clean.sh
.PHONY: devnet-l1-clean

devnet-logs:
	@(cd ops/docker && docker compose $(DEVNET_COMPOSE_FILES) logs -f)
.PHONY: devnet-logs

reth-image:
	docker pull "$(MORPH_RETH_IMAGE)"
.PHONY: reth-image

reth:
	@test -d "$(MORPH_RETH_DIR)" || (echo "morph-reth directory not found: $(MORPH_RETH_DIR)" && exit 1)
	docker build -t "$(MORPH_RETH_IMAGE)" --target "$(MORPH_RETH_DOCKER_TARGET)" --build-arg BUILD_PROFILE="$(MORPH_RETH_BUILD_PROFILE)" --build-arg RUSTFLAGS="$(MORPH_RETH_RUSTFLAGS)" "$(MORPH_RETH_DIR)"
.PHONY: reth

# tx-submitter
SUBMITTERS := $(shell grep -o 'tx-submitter-[0-9]*[^:]' ops/docker/docker-compose-devnet.yml | sort | uniq)
rebuild-all-tx-submitter:
	@for submitter in $(SUBMITTERS); do \
		docker compose -f ./ops/docker/docker-compose-devnet.yml up -d --build $$submitter --no-deps; \
	done
stop-all-tx-submitter:
	@for submitter in $(SUBMITTERS); do \
		docker compose -f ./ops/docker/docker-compose-devnet.yml stop $$submitter; \
	done
start-all-tx-submitter:
	@for submitter in $(SUBMITTERS); do \
		docker compose -f ./ops/docker/docker-compose-devnet.yml start $$submitter; \
	done

# build geth
geth: submodules
	cd go-ethereum && env GO111MODULE=on GOWORK=off go run build/ci.go install ./cmd/geth
