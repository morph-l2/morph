start-server:
	docker compose -f docker-compose-app.yml up -d

stop-server:
	docker compose -f docker-compose-app.yml down


build-shadow-prove:
	cd bin/shadow-prove && cargo build --release

build-challenge:
	cd bin/challenge && cargo build --release
