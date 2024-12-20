build-plonk-circuits:
	docker build -f Dockerfile.sp1-plonk -t sp1-plonk:latest .

save-plonk-circuits:
	rm -rf plonk-circuits && \
	mkdir -p plonk-circuits && \
	docker create --name sp1-plonk-temp sp1-plonk:latest
    docker cp -r sp1-plonk-temp:/app/sp1/crates/prover/build ./plonk-circuits
	cd plonk-circuits/build
	tar --exclude='srs.bin' --exclude='srs_lagrange.bin' -czvf "../morph-dev.tar.gz" .
	docker rm sp1-plonk-temp

start-server:
	docker compose -f docker-compose-app.yml up -d

stop-server:
	docker compose -f docker-compose-app.yml down

build-shadow-prove:
	cd bin/shadow-prove
	cargo build --release

build-challenge:
	cd bin/challenge
	cargo build --release