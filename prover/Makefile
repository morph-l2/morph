
save-plonk-circuits:
	rm -rf plonk-circuits && \
	mkdir -p plonk-circuits && \
	docker create --name sp1-plonk-temp sp1-plonk:latest
    docker cp -r sp1-plonk-temp:/app/sp1/crates/prover/build ./plonk-circuits
	cd plonk-circuits/build
	tar --exclude='srs.bin' --exclude='srs_lagrange.bin' -czvf "../morph-dev.tar.gz" .
	docker rm sp1-plonk-temp
