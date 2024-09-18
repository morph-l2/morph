pub mod types;
mod verifier;
use alloy::primitives::keccak256;
use sbv_core::VerificationError;
use sbv_primitives::B256;
use sbv_utils::dev_info;
use types::input::ClientInput;
use verifier::{blob_verifier::BlobVerifier, evm_verifier::EVMVerifier};

pub fn verify(input: &ClientInput) -> Result<B256, VerificationError> {
    // Verify DA
    let (versioned_hash, txns) = BlobVerifier::verify(&input.blob_info).unwrap();

    for tx in txns {}

    // Verify EVM exec.
    let batch_info = EVMVerifier::verify(&input.l2_traces).unwrap();

    // Calc public input hash.
    let public_input_hash = batch_info.public_input_hash(&versioned_hash);

    dev_info!("public input hash: {:?}", public_input_hash);
    Ok(B256::from_slice(public_input_hash.as_slice()))
}
