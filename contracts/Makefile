.PHONY: test lint lint-go lint-sol

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic -p 1 ./...

lint: lint-sol lint-go

# npm install --global --save-dev solhint
lint-sol:
	find . -name '*.sol' | xargs solhint

lint-go:
	golangci-lint run --out-format=tab --timeout 10m
