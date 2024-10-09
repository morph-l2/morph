//! An end-to-end example of using the SP1 SDK to generate a proof of a program that can have an
//! EVM-Compatible proof generated which can be verified on-chain.
//!
//! You can run this script using the following command:
//! ```shell
//! RUST_LOG=info cargo run --release --bin evm
//! ```

use morph_executor_utils::read_env_var;
use serde::{Deserialize, Serialize};
use std::path::PathBuf;

/// The ELF (executable and linkable format) file for the Succinct RISC-V zkVM.
pub const STATELESS_VERIFIER_ELF: &[u8] =
    include_bytes!("../../client/elf/riscv32im-succinct-zkvm-elf");

/// A fixture that can be used to test the verification of SP1 zkVM proofs inside Solidity.
#[derive(Debug, Clone, Serialize, Deserialize, Default)]
#[serde(rename_all = "camelCase")]
pub struct EvmProofFixture {
    pub vkey: String,
    pub public_values: String,
    pub proof: String,
}

/// Save a fixture for the given proof.
pub fn save_plonk_fixture(fixture: &EvmProofFixture) {
    // Save the fixture to a file.
    let proof_dir: String = read_env_var("PROVER_PROOF_DIR", "/data/proof".to_string());
    let fixture_path = PathBuf::from(proof_dir).join("./contracts/src/fixtures");
    std::fs::create_dir_all(&fixture_path).expect("failed to create fixture path");
    std::fs::write(
        fixture_path.join("plonk-fixture.json"),
        serde_json::to_string_pretty(&fixture).unwrap(),
    )
    .expect("failed to write fixture");
}
