.PHONY: build test lint clean run docker-build

GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')
VERSION := v0.1.0

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGSSTRING +=-X main.Version=$(VERSION)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

build:
	GOWORK=off GOSUMDB=off go build $(LDFLAGS) -o bin/gas-price-oracle ./cmd

run: build
	./bin/gas-price-oracle

test:
	go test -v ./...

lint:
	golangci-lint run ./...

clean:
	rm -rf bin/

docker-build:
	docker build -t morph/gas-price-oracle:latest .

help:
	@echo "Available targets:"
	@echo "  build              - Build the gas-price-oracle binary"
	@echo "  run                - Build and run the service"
	@echo "  test               - Run tests"
	@echo "  lint               - Run linter"
	@echo "  clean              - Clean build artifacts"
	@echo "  docker-build       - Build Docker image"

