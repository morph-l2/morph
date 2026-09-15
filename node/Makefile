GITCOMMIT := $(shell git rev-parse --short HEAD)
GITDATE := $(shell git show -s --format='%ct')
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

# Set by the morphnode-repro target below. Strips everything that varies between
# two builds of the same source tree, so the binary is byte-identical:
#   - the build timestamp, which has no meaning for a reproducible artifact
#   - the Go build ID, which depends on the state of the build cache
#   - absolute checkout paths, via -trimpath
#   - the VCS stamps, which depend on whether .git is present and whether the
#     tree is dirty; Version and GitCommit already carry that information
#   - the random LC_UUID that Apple's linker stamps into every Mach-O binary
#     (on Linux the equivalent GNU build-id is derived from content already)
ifdef REPRO
BUILD_TIME := reproducible
BUILDFLAGS += -trimpath -buildvcs=false
LDFLAGSSTRING +=-buildid=
ifeq ($(shell uname -s),Darwin)
LDFLAGSSTRING +=-extldflags=-Wl,-no_uuid
endif
endif

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.Version=$(VERSION)
LDFLAGSSTRING +=-X main.BuildTime=$(BUILD_TIME)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"
BUILDFLAGS += $(LDFLAGS)

morphnode:
	if [ ! -d build/bin ]; then mkdir -p build/bin; fi
	go mod download
	env GO111MODULE=on CGO_ENABLED=1 CGO_LDFLAGS="-ldl" go build -o build/bin/morphnode -v $(BUILDFLAGS) ./cmd/node
.PHONY: morphnode

tendermint:
	if [ ! -d build/bin ]; then mkdir -p build/bin; fi
	go mod download
	env GO111MODULE=on CGO_ENABLED=1 go build -o build/bin/tendermint -v $(BUILDFLAGS) ./cmd/tendermint
.PHONY: tendermint

build: morphnode tendermint
.PHONY: build

# Same binary as `make morphnode`, but byte-identical across builds of an
# identical source tree. This still requires a fixed toolchain: morphnode is
# built with cgo, so the Go version and the C toolchain must match too. Build
# inside ops/docker/Dockerfile.l2-node to guarantee that across machines.
morphnode-repro:
	$(MAKE) morphnode REPRO=1
.PHONY: morphnode-repro

init:
	if [ -d build/config ]; then exit 0; fi
	if [ ! -d build ]; then mkdir -p build; fi
	./build/bin/tendermint init --home build
.PHONY: init

run: build init
	cd ops-morph && sh run.sh

clean:
	rm -r build
.PHONY: clean

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic -p 1 ./...
.PHONY: test

lint:
	GOBIN=$(PWD)/build/bin go run ../build/lint.go
.PHONY: lint

devnet-up:
	cd ops-morph && docker compose up -d sequencer_node
.PHONY: dev-up

devnet-validator-up:
	cd ops-morph && docker compose -f docker-compose-validator.yml up -d
.PHONY: dev-validator-up

devnet-validator-build:
	cd ops-morph && docker compose -f docker-compose-validator.yml build
.PHONY: dev-validator-build

devnet-down:
	cd ops-morph && docker compose down
.PHONY: dev-down

devnet-validator-down:
	cd ops-morph && docker compose -f docker-compose-validator.yml down
.PHONY: dev-validator-down

devnet-clean:
	cd ops-morph && docker compose down
	docker image ls '*morph*' --format='{{.Repository}}' | xargs -r docker rmi
	docker volume ls --filter name=ops-morph* --format='{{.Name}}' | xargs -r docker volume rm
.PHONY: devnet-clean

devnet-reset:
	cd ops-morph && docker compose down
	docker volume ls --filter name=ops-morph* --format='{{.Name}}' | xargs -r docker volume rm
.PHONY: devnet-reset

testnet-up: build
	sh ./ops-morph/testnet/tendermint-setup.sh
	cd ops-morph/testnet && docker compose up -d
.PHONY: testnet-up

testnet-down:
	cd ops-morph/testnet && docker compose down
.PHONY: testnet-down

testnet-clean: testnet-down
	docker volume ls --filter "name=morph_data*" -q | xargs -r docker volume rm
	docker volume ls --filter "name=node_data*" -q | xargs -r docker volume rm
	rm -rf ./mytestnet
.PHONY: testnet-clean

install-tendermint:
	if [ ! -d build/bin ]; then mkdir -p build/bin; fi
	go mod download
	env GO111MODULE=on CGO_ENABLED=1 go install -v $(BUILDFLAGS) ./cmd/tendermint
.PHONY: tendermint







