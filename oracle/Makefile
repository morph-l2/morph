.PHONY: staking-oracle build clean test lint

GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

build:
	env GO111MODULE=on CGO_LDFLAGS="-ldl" go build -o build/bin/staking-oracle -v $(LDFLAGS) ./cmd/staking-oracle

clean:
	rm -r build

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic -p 1 ./...

lint:
	GOBIN=$(PWD)/build/bin go run ../build/lint.go
