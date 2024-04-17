build-prover:
	cargo build --release
	
build-challenge-handler:
	cd challenge-handler && cargo build --release

build-shadow-proving:
	cd shadow-proving && cargo build --release

all: build-prover build-challenge-handler build-shadow-proving
	if [ ! -d make-bin ]; then mkdir -p make-bin; fi
	cp -f `find ./target/release/ | grep libzktrie.so` make-bin/
	cp -f target/release/prover_server make-bin/
	cp -f challenge-handler/target/release/challenge-handler make-bin/
	cp -f shadow-proving/target/release/shadow-proving make-bin/

build: build-prover build-challenge-handler build-shadow-proving

run:all
	./start.sh