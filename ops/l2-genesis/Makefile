.PHONY: test test-scripts lint

PYTHON ?= python3

test: test-scripts
	go test -v -race -coverprofile=coverage.txt -covermode=atomic -p 1 ./...

test-scripts:
	$(PYTHON) -m unittest discover -s tests -v

lint: ## Lint the files - used for CI
	GOBIN="$(CURDIR)/build/bin" go run ../../build/lint.go -d 2
