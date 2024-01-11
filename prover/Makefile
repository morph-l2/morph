build-prover:
	cargo build --release

build-challenge-handler:
	cd challenge-handler && cargo build --release

all: build-prover build-challenge-handler
	if [ ! -d make-bin ]; then mkdir -p make-bin; fi
	cp -f `find ./target/release/ | grep libzktrie.so` make-bin/
	cp -f target/release/prove_cmd make-bin/
	cp -f target/release/prover_server make-bin/
	cp -f target/release/setup_prover make-bin/
	cp -f challenge-handler/target/release/challenge-handler make-bin/
	cp -f challenge-handler/target/release/auto_challenge make-bin/

run:all
	./start.sh