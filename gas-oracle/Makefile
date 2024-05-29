# Run cargo build with release tag
build:
	cd app && cargo build --release

# Generate bin file and run
run:build
	./app/target/release/app

# Run the formatter but don't change any files.
fmt:
	cd app && cargo +nightly fmt --all -- --check

# Run the formatter and update files with style changes
lint:
	cd app && cargo +nightly fmt --all

# Run clippy
clippy:
	cd app && cargo +nightly clippy --all-targets -- -D warnings

# Run all tests
test:
	cd app && cargo test
.PHONY: test

# Clean the builds
clean:
	cd app && cargo clean

# Generate the help message
help:
	@echo "usage: make [target]"
	@echo ""
	@echo "fmt: Run the formatter but don't change any files."
	@echo "lint: Run the formatter and update files with style changes"
	@echo "clippy: Run clippy"
	@echo "test: Run all tests"
	@echo "clean: Clean the builds"
	@echo "help: Show this help message"
	@echo ""