include .env

.PHONY: prepare-holesky-data download-and-decompress-snapshot

go-ubuntu-builder:
	@if [ -z "$(shell docker images -q morph/go-ubuntu-builder 2> /dev/null)" ]; then \
		echo "Docker image morph/go-ubuntu-builder does not exist. Building..."; \
		cd ../docker/intermediate && docker build -t morph/go-ubuntu-builder:go-1.22-ubuntu . -f go-ubuntu-builder.Dockerfile; \
	else \
		echo "Docker image morph/go-ubuntu-builder already exists."; \
	fi
.PHONY: go-ubuntu-builder

generate-jwt:
	@[ -f $(JWT_SECRET_FILE) ] || (echo "Generating $(JWT_SECRET_FILE)..." && openssl rand -hex 32 > $(JWT_SECRET_FILE) && echo "$(JWT_SECRET_FILE) created.")


submodules:
	git submodule update --init
	git submodule update --remote
.PHONY: submodules

######################start commands for managing holesky node##########################

prepare-holesky-data:
	@if [ ! -d "$(NODE_HOME)" ]; then \
  		echo "Creating node home directory..."; \
		mkdir -p $(NODE_HOME); \
		echo "Copying holesky config files to target node home..."; \
		cp -r holesky/* $(NODE_HOME); \
		echo "Node home directory and config files are ready."; \
		$(MAKE) download-and-decompress-snapshot; \
		mv ${SNAPSHOT_NAME}/geth $(NODE_HOME)/geth-data; \
		mv ${SNAPSHOT_NAME}/data/* $(NODE_HOME)/node-data/data; \
		echo "Snapshot data is successfully set up."; \
	fi 

download-and-decompress-snapshot:
	@echo "Checking for required tools (wget, tar)..."; \
	if ! command -v wget > /dev/null; then \
		echo "Error: wget is not installed."; \
		exit 1; \
	fi; \
	if ! command -v tar > /dev/null; then \
		echo "Error: tar is not installed."; \
		exit 1; \
	fi; \
	echo "Tools are available."; \
	echo "Downloading ${SNAPSHOT_NAME}.tar.gz from https://snapshot.morphl2.io/holesky..."; \
	if wget -q --show-progress -O ${SNAPSHOT_NAME}.tar.gz https://snapshot.morphl2.io/holesky/${SNAPSHOT_NAME}.tar.gz; then \
		echo "${SNAPSHOT_NAME}.tar.gz downloaded successfully."; \
	else \
		echo "Error downloading ${SNAPSHOT_NAME}.tar.gz."; \
		exit 1; \
	fi; \
	echo "Decompressing ${SNAPSHOT_NAME}.tar.gz..."; \
	if tar -xzvf ${SNAPSHOT_NAME}.tar.gz; then \
		echo "${SNAPSHOT_NAME}.tar.gz decompressed successfully."; \
	else \
		echo "Error decompressing ${SNAPSHOT_NAME}.tar.gz."; \
		exit 1; \
	fi; \
	rm -rf ${SNAPSHOT_NAME}.tar.gz; 


run-holesky-node: go-ubuntu-builder submodules prepare-holesky-data generate-jwt
	docker-compose up holesky-node &

stop-holesky-node:
	docker stop morph-node-holesky morph-geth-holesky

rm-holesky-node:
	docker rm morph-node-holesky morph-geth-holesky

######################end commands for managing holesky node##########################
	