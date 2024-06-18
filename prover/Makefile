# Run cargo build with release tag for prover
build-prover:
	cargo build --release

# Run cargo build with release tag for challenge-handler
build-challenge-handler:
	cd challenge-handler && cargo build --release

# Run cargo build with release tag for shadow-proving
build-shadow-proving:
	cd shadow-proving && cargo build --release

# Build all bins and prepare the deployment directory
all: build-prover build-challenge-handler build-shadow-proving
	if [ ! -d make-bin ]; then mkdir -p make-bin; fi
	cp -f `find ./target/release/ | grep libzktrie.so` make-bin/
	cp -f target/release/prover_server make-bin/
	cp -f challenge-handler/target/release/challenge-handler make-bin/
	cp -f shadow-proving/target/release/shadow-proving make-bin/

# Build all bins
build: build-prover build-challenge-handler build-shadow-proving

# Run all bins in the background
run:all
	./start.sh

# Run the formatter but don't change any files.
fmt:
	cargo +nightly fmt --all -- --check

# Run the formatter and update files with style changes
lint:
	cargo +nightly fmt --all

# Run clippy
clippy:
	cargo clippy --all-targets -- -D warnings