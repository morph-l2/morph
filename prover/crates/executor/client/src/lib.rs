pub mod types;
mod verifier;
use alloy_primitives::B256;
#[cfg(not(target_os = "zkvm"))]
use alloy_primitives::hex;
use prover_primitives::types::blob::get_blob_data_from_blocks;
use types::{batch::BATCH_VERSION, input::ExecutorInput};
pub use verifier::{blob_verifier::BlobVerifier, evm_verifier::EVMVerifier};

use crate::types::blob::decompress_batch;

pub fn verify(input: ExecutorInput) -> Result<B256, anyhow::Error> {
    // Verify basicInfo
    if input.batch_version != BATCH_VERSION {
        return Err(anyhow::anyhow!("unsupported batch version: {}", input.batch_version));
    }
    if input.block_inputs.is_empty() {
        return Err(anyhow::anyhow!("empty batch: no block inputs provided"));
    }

    // Verify DA
    let (versioned_hashes, batch_bytes) = BlobVerifier::verify_blobs(&input.blob_infos)?;
    if versioned_hashes.len() != input.blob_infos.len() {
        return Err(anyhow::anyhow!("versioned_hashes.len not equals blob_infos.len",));
    }
    let batch_data_from_blob = decompress_batch(&batch_bytes)?;

    let batch_data_from_blocks = get_blob_data_from_blocks(
        &input.block_inputs.iter().map(|input| input.current_block.clone()).collect::<Vec<_>>(),
    );
    if batch_data_from_blob != batch_data_from_blocks {
        return Err(anyhow::anyhow!("blob data mismatch!"));
    }

    // Verify EVM exec.
    let batch_info = EVMVerifier::verify(input.block_inputs)?;

    #[cfg(not(target_os = "zkvm"))]
    log::info!(
        "cacl pi hash, prevStateRoot = {:?}, postStateRoot = {:?}, withdrawalRoot = {:?},
        dataHash = {:?}, blobVersionedHashes = {:?}, batch_version = {}",
        hex::encode(batch_info.prev_state_root().as_slice()),
        hex::encode(batch_info.post_state_root().as_slice()),
        hex::encode(batch_info.withdraw_root().as_slice()),
        hex::encode(batch_info.data_hash().as_slice()),
        versioned_hashes.iter().map(|h| hex::encode(h.as_slice())).collect::<Vec<_>>(),
        input.batch_version,
    );
    // Calc public input hash.
    let public_input_hash = batch_info.public_input_hash(&versioned_hashes);

    #[cfg(not(target_os = "zkvm"))]
    log::info!("public input hash: {public_input_hash:?}");
    Ok(B256::from_slice(public_input_hash.as_slice()))
}
