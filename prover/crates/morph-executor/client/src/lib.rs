pub mod types;
mod verifier;
use alloy::primitives::keccak256;
use sbv_core::VerificationError;
use sbv_primitives::B256;
use types::input::ClientInput;
use verifier::{blob_verifier::BlobVerifier, evm_verifier::EVMVerifier};

pub fn verify(input: &ClientInput) -> Result<B256, VerificationError> {
    // Verify DA
    let (versioned_hash, txns) = BlobVerifier::verify(&input.blob_info).unwrap();

    // Verify EVM exec.
    // &input.l2_trace.transactions(txns);
    let state_root = EVMVerifier::verify(&input.l2_trace).unwrap();

    // calc public input hash.
    let pi_hash = keccak256([versioned_hash, state_root].concat());

    Ok(B256::from_slice(pi_hash.as_slice()))
}
