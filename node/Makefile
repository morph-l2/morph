GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

morphnode:
	if [ ! -d build/bin ]; then mkdir -p build/bin; fi
	go mod download
	env GO111MODULE=on CGO_ENABLED=1 CGO_LDFLAGS="-ldl" go build -o build/bin/morphnode -v $(LDFLAGS) ./cmd/node
.PHONY: morphnode

tendermint:
	if [ ! -d build/bin ]; then mkdir -p build/bin; fi
	go mod download
	env GO111MODULE=on CGO_ENABLED=1 go build -o build/bin/tendermint -v $(LDFLAGS) ./cmd/tendermint
.PHONY: tendermint

build: morphnode tendermint
.PHONY: build

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







