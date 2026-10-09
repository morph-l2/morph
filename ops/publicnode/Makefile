ENV_FILE ?= .env
include $(ENV_FILE)

export NODE_HOME GETH_DATA_DIR NODE_DATA_DIR GETH_ENTRYPOINT_FILE JWT_SECRET_FILE
export GETH_NETWORK_ID MORPH_NODE_L1_ETH_RPC MORPH_NODE_SYNC_DEPOSIT_CONTRACT_ADDRESS
export MORPH_NODE_L1_SEQUENCER_CONTRACT MORPH_NODE_ROLLUP_ADDRESS MORPH_NODE_SEQUENCER_UPGRADE_TIME
export MORPH_NODE_L1_ETH_BEACON_RPC
COMPOSE := docker compose --env-file "$(ENV_FILE)"
SNAPSHOT_DIR ?= $(CURDIR)/$(SNAPSHOT_NAME)
SNAPSHOT_URL ?= https://snapshot.morphl2.io/holesky/$(SNAPSHOT_NAME).tar.gz

.PHONY: go-ubuntu-builder generate-jwt submodules check-config prepare-holesky-data
.PHONY: download-and-decompress-snapshot run-holesky-node stop-holesky-node rm-holesky-node

# Runtime checks must not run while Make parses the file or during a dry run.
go-ubuntu-builder:
	@if ! docker image inspect morph/go-ubuntu-builder:go-1.24-ubuntu >/dev/null 2>&1; then \
		docker build -t morph/go-ubuntu-builder:go-1.24-ubuntu ../docker/intermediate -f ../docker/intermediate/go-ubuntu-builder.Dockerfile; \
	fi

generate-jwt:
	@set -eu; \
	if [ -e "$(JWT_SECRET_FILE)" ]; then \
		[ -f "$(JWT_SECRET_FILE)" ] || { echo "JWT_SECRET_FILE must be a file" >&2; exit 1; }; \
		jwt=$$(cat "$(JWT_SECRET_FILE)"); \
		[ $${#jwt} -eq 64 ] && ! printf '%s' "$$jwt" | LC_ALL=C grep -q '[^0-9a-fA-F]' || { echo "Invalid JWT_SECRET_FILE" >&2; exit 1; }; \
	else \
		umask 077; mkdir -p "$$(dirname "$(JWT_SECRET_FILE)")"; \
		tmp=$$(mktemp "$(JWT_SECRET_FILE).XXXXXX"); \
		openssl rand -hex 32 > "$$tmp"; mv "$$tmp" "$(JWT_SECRET_FILE)"; \
	fi

submodules:
	git submodule update --init --recursive

# Legacy Holesky files do not provide defaults for the current node binary.
check-config:
	@set -eu; \
	for name in GETH_NETWORK_ID MORPH_NODE_L1_ETH_RPC MORPH_NODE_L1_ETH_BEACON_RPC MORPH_NODE_SYNC_DEPOSIT_CONTRACT_ADDRESS MORPH_NODE_L1_SEQUENCER_CONTRACT MORPH_NODE_ROLLUP_ADDRESS MORPH_NODE_SEQUENCER_UPGRADE_TIME; do \
		printenv "$$name" | LC_ALL=C grep -q '[^[:space:]]' || { echo "Set $$name in $(ENV_FILE) or the environment" >&2; exit 1; }; \
	done; \
	printf '%s\n' "$$GETH_NETWORK_ID" | LC_ALL=C grep -Eq '^[1-9][0-9]*$$' || { echo "GETH_NETWORK_ID must be positive" >&2; exit 1; }; \
	printf '%s\n' "$$MORPH_NODE_SEQUENCER_UPGRADE_TIME" | LC_ALL=C grep -Eq '^[0-9]+$$' || { echo "MORPH_NODE_SEQUENCER_UPGRADE_TIME must be Unix milliseconds" >&2; exit 1; }; \
	for name in MORPH_NODE_SYNC_DEPOSIT_CONTRACT_ADDRESS MORPH_NODE_L1_SEQUENCER_CONTRACT MORPH_NODE_ROLLUP_ADDRESS; do \
		address=$$(printenv "$$name"); \
		printf '%s\n' "$$address" | LC_ALL=C grep -Eq '^0x[0-9a-fA-F]{40}$$' && [ "$$address" != 0x0000000000000000000000000000000000000000 ] || { echo "$$name must be a nonzero contract address" >&2; exit 1; }; \
	done; \
	$(COMPOSE) config --quiet

prepare-holesky-data:
	@set -eu; \
	if [ -e "$(NODE_HOME)" ]; then \
		[ -d "$(GETH_DATA_DIR)/geth/chaindata" ] && [ -d "$(NODE_DATA_DIR)/data" ] && [ -f "$(NODE_DATA_DIR)/config/genesis.json" ] || { echo "Existing NODE_HOME is incomplete; preserve it and restore a matching snapshot before retrying" >&2; exit 1; }; \
	else \
		[ -d "$(SNAPSHOT_DIR)/geth/chaindata" ] && [ -d "$(SNAPSHOT_DIR)/data" ] || { echo "Run make download-and-decompress-snapshot first" >&2; exit 1; }; \
		mkdir -p "$$(dirname "$(NODE_HOME)")"; \
		stage=$$(mktemp -d "$(NODE_HOME).prepare.XXXXXX"); \
		cp -R holesky/. "$$stage/"; \
		cp -R "$(SNAPSHOT_DIR)/geth" "$$stage/geth-data/geth"; \
		cp -R "$(SNAPSHOT_DIR)/data/." "$$stage/node-data/data/"; \
		mv "$$stage" "$(NODE_HOME)"; \
	fi

download-and-decompress-snapshot:
	@set -eu; \
	command -v wget >/dev/null; command -v tar >/dev/null; \
	[ ! -e "$(SNAPSHOT_DIR)" ] || { echo "SNAPSHOT_DIR already exists; inspect or choose another directory" >&2; exit 1; }; \
	mkdir -p "$$(dirname "$(SNAPSHOT_DIR)")"; \
	stage=$$(mktemp -d "$(SNAPSHOT_DIR).download.XXXXXX"); \
	echo "Downloading snapshot into $$stage"; \
	wget -O "$$stage/snapshot.tar.gz" "$(SNAPSHOT_URL)"; \
	tar -tzf "$$stage/snapshot.tar.gz" > "$$stage/entries.txt"; \
	awk '/^\/|(^|\/)\.\.(\/|$$)/ { exit 1 }' "$$stage/entries.txt"; \
	tar -xzf "$$stage/snapshot.tar.gz" -C "$$stage"; \
	[ -d "$$stage/$(SNAPSHOT_NAME)/geth/chaindata" ] && [ -d "$$stage/$(SNAPSHOT_NAME)/data" ]; \
	mv "$$stage/$(SNAPSHOT_NAME)" "$(SNAPSHOT_DIR)"; \
	rm "$$stage/snapshot.tar.gz" "$$stage/entries.txt"; rmdir "$$stage"

# Keep preparation ordered even when the caller uses make -j.
run-holesky-node: check-config
	$(MAKE) go-ubuntu-builder
	$(MAKE) submodules
	$(MAKE) prepare-holesky-data
	$(MAKE) generate-jwt
	$(COMPOSE) up -d --build holesky-node

stop-holesky-node:
	$(COMPOSE) stop holesky-node holesky-geth

rm-holesky-node:
	$(COMPOSE) rm -f holesky-node holesky-geth
