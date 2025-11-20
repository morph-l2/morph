.PHONY: build test lint clean run docker-build

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
	env GO111MODULE=on CGO_ENABLED=1 CGO_LDFLAGS="-ldl" go build -o build/bin/token-price-oracle -v $(LDFLAGS) ./cmd
run: build
	./build/bin/token-price-oracle

test:
	go test -v ./...

lint:
	golangci-lint run ./...

clean:
	rm -rf build/

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

