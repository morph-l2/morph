pub mod types;
mod verifier;
#[cfg(not(target_os = "zkvm"))]
use alloy_primitives::hex;
use prover_primitives::{types::blob::get_blob_data_from_blocks, B256};
use types::input::ExecutorInput;
pub use verifier::{blob_verifier::BlobVerifier, evm_verifier::EVMVerifier};

pub const EVM_VERIFY: &str = "evm verify";

pub fn verify(input: ExecutorInput) -> Result<B256, anyhow::Error> {
    // Verify DA
    let (versioned_hashes, batch_data_from_blob) = BlobVerifier::verify_blobs(&input.blob_infos)?;
    let batch_data_from_blocks = get_blob_data_from_blocks(
        &input.block_inputs.iter().map(|input| input.current_block.clone()).collect::<Vec<_>>(),
    );
    if batch_data_from_blob != batch_data_from_blocks {
        return Err(anyhow::anyhow!("blob data mismatch!"));
    }

    // Verify EVM exec.
    let batch_info = profile_report!(EVM_VERIFY, { EVMVerifier::verify(input.block_inputs) })?;

    // Calc public input hash based on version.
    #[cfg(not(target_os = "zkvm"))]
    log::info!(
        "cacl pi hash, prevStateRoot = {:?}, postStateRoot = {:?}, withdrawalRoot = {:?},
        dataHash = {:?}, blobVersionedHashes = {:?}, sequencerSetVerifyHash = {:?}, batch_version = {}",
        hex::encode(batch_info.prev_state_root().as_slice()),
        hex::encode(batch_info.post_state_root().as_slice()),
        hex::encode(batch_info.withdraw_root().as_slice()),
        hex::encode(batch_info.data_hash().as_slice()),
        versioned_hashes.iter().map(|h| hex::encode(h.as_slice())).collect::<Vec<_>>(),
        hex::encode(batch_info.sequencer_root().as_slice()),
        input.batch_version,
    );
    let public_input_hash = if input.batch_version >= 2 {
        batch_info.public_input_hash_v2(&versioned_hashes)
    } else {
        batch_info.public_input_hash(&versioned_hashes[0])
    };
    #[cfg(not(target_os = "zkvm"))]
    log::info!("public input hash: {public_input_hash:?}");
    Ok(B256::from_slice(public_input_hash.as_slice()))
}
