.PHONY: build test lint

build:
	mkdir -p build/bin
	@set -e; for tool in batchparse gasinspect keygen multisend; do \
		go build -o "build/bin/$$tool" "./$$tool"; \
	done

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic -p 1 ./...

lint: ## Lint the files - used for CI
	GOBIN="$(CURDIR)/build/bin" go run ../../build/lint.go -d 2
