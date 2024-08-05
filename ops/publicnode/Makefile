go-ubuntu-builder:
	@if [ -z "$(shell docker images -q morph/go-ubuntu-builder 2> /dev/null)" ]; then \
		echo "Docker image morph/go-ubuntu-builder does not exist. Building..."; \
		cd ../docker/intermediate && docker build -t morph/go-ubuntu-builder:go-1.22-ubuntu . -f go-ubuntu-builder.Dockerfile; \
	else \
		echo "Docker image morph/go-ubuntu-builder already exists."; \
	fi
.PHONY: go-ubuntu-builder

run-holesky-node: go-ubuntu-builder 
	docker-compose up holesky-node 
	