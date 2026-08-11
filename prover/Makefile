RUST_TOOLCHAIN ?= 1.91.0

start-server:
	docker compose -f docker-compose-app.yml up -d

stop-server:
	docker compose -f docker-compose-app.yml down


build-shadow-prove:
	cd bin/shadow-prove && cargo +$(RUST_TOOLCHAIN) build --release

build-challenge:
	cd bin/challenge && cargo +$(RUST_TOOLCHAIN) build --release
