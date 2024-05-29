# gas oracle

## Development

### Prerequisites

Make sure you have Rust installed. We recommend using [rustup](https://rustup.rs/). On top of the stable version of Rust, you will also need to install the nightly version of Rust, since the nightly version is used for some of the build and test scripts.

```bash
rustup toolchain install nightly
```

### Makefile

The Makefile contains various commands for building, testing, and linting the files. A list of the commands can be found by running `make help`.

### Testing

To run all tests, run `make test`.

### Linting

To run a format check but not actually format the code, run `make fmt`. To run a format check and format the code, run `make lint`. To run a clippy check, run `make clippy`.

### Build

`make build`

### Run

`make run`
