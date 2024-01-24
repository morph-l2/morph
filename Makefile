.PHONY: update format

update: ## update the dependencies
	go work sync

format: ## format the code
	go work sync
	goimports -local $(PWD)/bindings/ -w .
	goimports -local $(PWD)/contracts/ -w .
	goimports -local $(PWD)/node/ -w .
	goimports -local $(PWD)/tx-submitter/ -w .

