PWD := $(shell pwd)


build-bk-prod-morph-prod-mainnet-to-morph-gas-oracle:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/gas-oracle/app && cargo build --release
	cp gas-oracle/app/target/release/app dist/
	tar -czvf gas-oracle.tar.gz dist
	aws s3 cp gas-oracle.tar.gz s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/gas-oracle.tar.gz

build-bk-prod-morph-prod-mainnet-to-morph-prover:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover && cargo build --release
	cp prover/target/release/prover_server dist/
	cp -r prover/configs dist/
	tar -czvf prover.tar.gz dist
	aws s3 cp prover.tar.gz s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/prover.tar.gz

build-bk-prod-morph-prod-mainnet-to-morph-challenge-handler:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/challenge-handler && cargo build --release
	cp prover/challenge-handler/target/release/challenge-handler dist/
	tar -czvf challenge-handler.tar.gz dist
	aws s3 cp challenge-handler.tar.gz s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/challenge-handler.tar.gz

build-bk-prod-morph-prod-mainnet-to-morph-shadow-proving:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/prover/shadow-proving && cargo build --release
	cp prover/shadow-proving/target/release/shadow-proving dist/
	tar -czvf shadow-proving.tar.gz dist
	aws s3 cp shadow-proving.tar.gz s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/shadow-proving.tar.gz