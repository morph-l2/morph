build-bk-prod-morph-prod-mainnet-to-morph-node:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/node && make build
	cp node/build/bin/morphnode dist/
	cp node/build/bin/tendermint dist/
	tar -czvf morph-node.tar.gz dist
	aws s3 cp morph-node.tar.gz s3://morph-0582-morph-technical-department-mainnet-data/morph-setup/morph-node.tar.gz
