.PHONY: test int

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic -p 1 ./...

lint: ## Lint the files - used for CI
	golangci-lint run --out-format=tab --timeout 10m
