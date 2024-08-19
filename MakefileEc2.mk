build-bk-prod-morph-prod-mainnet-to-morph-node:
	if [ ! -d dist ]; then mkdir -p dist; fi
	cd $(PWD)/node && make build
	cp node/build/bin/morphnode dist/
	cp node/build/bin/tendermint dist/