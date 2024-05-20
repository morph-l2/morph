.PHONY: staking-oracle build clean test lint

GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.GitVersion=$(GITVERSION)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

staking-oracle:
	env GO111MODULE=on go build -o build/bin/staking-oracle -v $(LDFLAGS) ./cmd/staking-oracle

clean:
	rm -r build

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic -p 1 ./...

lint:
	golangci-lint run --out-format=tab --timeout 10m
