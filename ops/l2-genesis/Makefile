.PHONY: test lint

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic -p 1 ./...

lint: ## Lint the files - used for CI
	GOBIN=$(PWD)/build/bin go run ../../build/lint.go -d 2
