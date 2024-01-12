.PHONY: update format

update: ## update the dependencies
	git submodule update --init --recursive
	go work sync
	cd $(PWD)/l2geth && go mod tidy
	cd $(PWD)/tendermint && go mod tidy

format: ## format the code
	go work sync
	goimports -local $(PWD)/bindings/ -w .
	goimports -local $(PWD)/contracts/ -w .
	goimports -local $(PWD)/node/ -w .
	goimports -local $(PWD)/tx-submitter/ -w .

