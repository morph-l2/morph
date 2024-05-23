.PHONY: test lint lint-go lint-sol

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic -p 1 ./...

lint: lint-sol lint-go

# npm install --global --save-dev solhint
lint-sol:
	solhint contracts/**/*.sol

lint-go:
	GOBIN=$(PWD)/build/bin go run ../build/lint.go

