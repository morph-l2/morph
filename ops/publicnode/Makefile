go-ubuntu-builder:
	@if [ -z "$(shell docker images -q morph/go-ubuntu-builder 2> /dev/null)" ]; then \
		echo "Docker image morph/go-ubuntu-builder does not exist. Building..."; \
		cd ../docker/intermediate && docker build -t morph/go-ubuntu-builder:go-1.22-ubuntu . -f go-ubuntu-builder.Dockerfile; \
	else \
		echo "Docker image morph/go-ubuntu-builder already exists."; \
	fi
.PHONY: go-ubuntu-builder


SECRET_FILE = holesky/jwt-secret.txt

jwt:
	@[ -f $(SECRET_FILE) ] || (echo "Generating $(SECRET_FILE)..." && openssl rand -hex 32 > $(SECRET_FILE) && echo "$(SECRET_FILE) created.")

clean-jwt:
	@rm -f $(SECRET_FILE)
	@echo "Cleaned up $(SECRET_FILE)."

run-holesky-node: go-ubuntu-builder jwt
	docker-compose up holesky-node 
	