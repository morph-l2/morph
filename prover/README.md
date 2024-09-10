# Morph prover
Generate zk proof for the l2 batch.

## Requirements

- [Rust](https://rustup.rs/)
- [SP1](https://succinctlabs.github.io/sp1/getting-started/install.html)


### Build the Program


To build the program to risc-v bin:

```sh
cd bin/client
cargo prove build
```

This will output the compiled ELF to the file program/elf/riscv32im-succinct-zkvm-elf.

### Execute the Program


To run the program without generating a proof:

```sh
cd bin/host
RUST_BACKTRACE=full cargo run --release -- --execute
```

This will execute the program and display the output.

### Generate a Core Proof

To generate a core proof for your program:

```sh
cd bin/host
RUST_BACKTRACE=full cargo run --release -- --prove
```

### Generate an EVM-Compatible (PLONK) Proof

> [!WARNING]
> You will need at least 128GB RAM to generate the PLONK proof.

To generate a PLONK proof that is small enough to be verified on-chain and verifiable by the EVM:

```sh
cd script
cargo run --release --bin evm
```

This command also generates a fixture that can be used to test the verification of SP1 zkVM proofs
inside Solidity.

### Retrieve the Verification Key

To retrieve your `programVKey` for your on-chain contract, run the following command:

```sh
cargo run --release --bin vkey
```

