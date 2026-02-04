# Morph prover
Generate zk proof for the l2 batch.

## Requirements

- [Rust](https://rustup.rs/)
- [SP1](https://docs.succinct.xyz/docs/sp1/getting-started/install)


### Build the Program


To build the program to risc-v bin:

```sh
cd bin/client
cargo prove build
```
or use docker(reproducible compilation):
```sh
cargo prove build  --docker --tag v5.2.4
```

This will output the compiled ELF to the file client/elf/verifier-client.

### Execute the Program


To run the program without generating a proof:

```sh
cd bin/host
RUST_BACKTRACE=full cargo run --release
```

This will execute the program and display the output.


### Generate an EVM-Compatible (PLONK) Proof

> [!WARNING]
> If you are generating the proof locally, you will need at least 128GB of RAM to generate the PLONK proof.

To generate a PLONK proof that is small enough to be verified on-chain and verifiable by the EVM:

```sh
// use network prover
cargo run --release --package morph-prove --bin prove  -- --block-path ./testdata/mpt/mainnet_25215.json --prove
// use local cpu prover
cargo run --release --no-default-features --features local --package morph-prove --bin prove  -- --block-path ./testdata/mpt/mainnet_25215.json --prove
// prove specified block range
cargo run --release --package morph-prove --bin prove  -- --start-block 0x35 --end-block 0x37 --rpc http://127.0.0.1:8545 --use-rpc-db --prove
```

This command also generates a fixture that can be used to test the verification of SP1 zkVM proofs
inside Solidity.

### Retrieve the Verification Key

To retrieve your `programVKey` for on-chain contract, run the following command:

```sh
cargo run --release --bin vkey
```

The programVKey generated above must be identical to the configuration in the L1 Verifier contract `morph-repo/contracts/src/deploy-config`.
This key serves as a cryptographic commitment to the L2 State Transition Function (STF) logic, ensuring that the off-chain computation has not been tampered with.