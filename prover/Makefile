
save-plonk-circuits:
	rm -rf snark-circuits && \
	mkdir -p snark-circuits && \
	docker create --name sp1-plonk-temp sp1-plonk:latest
    docker cp -r sp1-plonk-temp:/app/sp1/crates/prover/build ./app
	tar --exclude='./app/srs.bin' --exclude='./app/srs_lagrange.bin' -czvf "../morph-dev.tar.gz" .
	docker rm sp1-plonk-temp

save-plonk-verifier:
	rm -rf plonk-verifier && \
	mkdir -p plonk-verifier && \
	docker create --name sp1-app-temp sp1-app:latest
    docker cp -r sp1-app-temp:/app/prover/build ./app
	docker rm sp1-app-temp