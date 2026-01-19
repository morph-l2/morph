pub mod types;
mod verifier;
use alloy_primitives::hex;
use prover_primitives::{types::blob::get_blob_data_from_blocks, B256};
use types::input::ExecutorInput;
pub use verifier::{blob_verifier::BlobVerifier, evm_verifier::EVMVerifier};

pub const EVM_VERIFY: &str = "evm verify";

pub fn verify(input: ExecutorInput) -> Result<B256, anyhow::Error> {
    // Verify DA
    let num_blocks = input.block_inputs.len();
    let (versioned_hash, batch_data_from_blob) =
        BlobVerifier::verify(&input.blob_info, num_blocks).unwrap();
    let batch_data_from_blocks = get_blob_data_from_blocks(
        &input.block_inputs.iter().map(|input| input.current_block.clone()).collect::<Vec<_>>(),
    );
    assert_eq!(batch_data_from_blob, batch_data_from_blocks, "blob data mismatch!");

    // Verify EVM exec.
    let batch_info = profile_report!(EVM_VERIFY, { EVMVerifier::verify(input.block_inputs) })?;

    // Calc public input hash.
    println!(
        "cacl pi hash, prevStateRoot = {:?}, postStateRoot = {:?}, withdrawalRoot = {:?},
        dataHash = {:?}, blobVersionedHash = {:?}, sequencerSetVerifyHash = {:?}",
        hex::encode(batch_info.prev_state_root().as_slice()),
        hex::encode(batch_info.post_state_root().as_slice()),
        hex::encode(batch_info.withdraw_root().as_slice()),
        hex::encode(batch_info.data_hash().as_slice()),
        hex::encode(versioned_hash.as_slice()),
        hex::encode(batch_info.sequencer_root().as_slice()),
    );
    let public_input_hash = batch_info.public_input_hash(&versioned_hash);
    println!("public input hash: {:?}", public_input_hash);
    Ok(B256::from_slice(public_input_hash.as_slice()))
}
