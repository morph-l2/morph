GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')
GITVERSION := $(shell cat package.json | jq .version)

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.GitVersion=$(GITVERSION)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"


staking-oracle:
	env GO111MODULE=on go build -o build/bin/staking-oracle -v $(LDFLAGS) ./cmd/staking-oracle

clean:
	rm -r build

test:
	go test -v ./...

lint:
	golangci-lint run ./...

.PHONY: \
	staking-oracle \
	build \
	clean \
	test \
	lint
