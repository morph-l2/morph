.PHONY: all build test lint clean run help

all: build

GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')
VERSION := v0.1.0

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.GitVersion=$(VERSION)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

build:
	if [ ! -d build/bin ]; then mkdir -p build/bin; fi
	go mod download
	env GO111MODULE=on CGO_ENABLED=1 CGO_LDFLAGS="-ldl" go build -o build/bin/challenger -v $(LDFLAGS) ./cmd

test:
	CGO_ENABLED=1 go test ./...

clean:
	rm -rf build/

help:
	@echo "Available targets:"
	@echo "  build  - Build the challenger binary"
	@echo "  test   - Run tests"
	@echo "  clean  - Remove build artifacts"
