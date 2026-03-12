.PHONY: build clean test lint

GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')
GITVERSION := v1.0.0

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.GitVersion=$(GITVERSION)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

PLATFORM ?= pc

ifeq (${PLATFORM},server)
    export GOOS=linux
    export GOARCH=386
endif

build:
	env CGO_LDFLAGS="-ldl" go build -v $(LDFLAGS) -o tx-submitter ./cmd

clean:
	rm tx-submitter

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic -p 1 $$(go list ./... | grep -v '/batch')

lint:
	GOBIN=$(PWD)/build/bin go run ../build/lint.go

run: build
	./tx-submitter \
	--build-env=dev \
    --l1-eth-rpc=http://127.0.0.1:9545 \
    --private-key="your_private_key" \
    --l2-eth-rpc=http://127.0.0.1:8545 \
    --poll-interval=3s \
    --safe-minimum-ether-balance=1 \
    --rollup-address="0x6900000000000000000000000000000000000010" \
    --network-timeout=10s \
    --tx-timeout=15s \
    --max-tx-size=125952 \
    --batch-build-timeout=60s \
    --log-level=debug
