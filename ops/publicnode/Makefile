include .env

go-ubuntu-builder:
	@if [ -z "$(shell docker images -q morph/go-ubuntu-builder 2> /dev/null)" ]; then \
		echo "Docker image morph/go-ubuntu-builder does not exist. Building..."; \
		cd ../docker/intermediate && docker build -t morph/go-ubuntu-builder:go-1.22-ubuntu . -f go-ubuntu-builder.Dockerfile; \
	else \
		echo "Docker image morph/go-ubuntu-builder already exists."; \
	fi
.PHONY: go-ubuntu-builder


######################start commands for managing holesky node##########################

holesky-jwt:
	@[ -f $(JWT_SECRET_FILE) ] || (echo "Generating $(JWT_SECRET_FILE)..." && openssl rand -hex 32 > $(JWT_SECRET_FILE) && echo "$(JWT_SECRET_FILE) created.")

copy-holesky-config:
	@if [ ! -d "$(NODE_HOME)" ]; then \
		mkdir -p $(NODE_HOME); \
		echo "Copying holesky config files to your target node home..."; \
		cp -r holesky/* $(NODE_HOME); \
		echo "${NODE_HOME} created."; \
	fi

run-holesky-node: go-ubuntu-builder copy-holesky-config holesky-jwt
	docker-compose up holesky-node 

stop-holesky-node:
	docker stop morph-node-holesky morph-geth-holesky

rm-holesky-node:
	docker rm morph-node-holesky morph-geth-holesky

######################end commands for managing holesky node##########################
	