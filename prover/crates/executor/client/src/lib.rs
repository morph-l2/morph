pub mod types;
mod verifier;
#[cfg(not(target_os = "zkvm"))]
use alloy_primitives::hex;
use prover_primitives::{types::blob::get_blob_data_from_blocks, B256};
use types::input::ExecutorInput;
pub use verifier::{blob_verifier::BlobVerifier, evm_verifier::EVMVerifier};

use crate::types::{batch::BatchInfo, blob::decompress_batch};

pub const EVM_VERIFY: &str = "evm verify";

pub fn verify(input: ExecutorInput) -> Result<B256, anyhow::Error> {
    if input.block_inputs.is_empty() {
        return Err(anyhow::anyhow!("empty batch: no block inputs provided"));
    }

    // Verify DA
    let (versioned_hashes, batch_bytes) = BlobVerifier::verify_blobs(&input.blob_infos)?;
    match input.batch_version {
        0..=2 if versioned_hashes.is_empty() => {
            return Err(anyhow::anyhow!(
                "batch version {} requires at least 1 versioned hash",
                input.batch_version
            ));
        }
        0..=2 => {}
        v => return Err(anyhow::anyhow!("unsupported batch version: {}", v)),
    }
    let default_pi_hash = BatchInfo::public_input_hash_for_invalid_blob(
        &versioned_hashes,
        &input.block_inputs,
        input.batch_version,
    )?;
    let batch_data_from_blob = match decompress_batch(&batch_bytes) {
        Ok(data) => data,
        Err(_) => {
            // Since the batch is not executed, all post roots remain unchanged
            return Ok(default_pi_hash);
        }
    };
    let batch_data_from_blocks = get_blob_data_from_blocks(
        &input.block_inputs.iter().map(|input| input.current_block.clone()).collect::<Vec<_>>(),
    );
    if batch_data_from_blob != batch_data_from_blocks {
        return Err(anyhow::anyhow!("blob data mismatch!"));
    }

    // Verify EVM exec.
    let batch_info = match profile_report!(EVM_VERIFY, { EVMVerifier::verify(input.block_inputs) })
    {
        Ok(batch_info) => batch_info,
        Err(_) => return Ok(default_pi_hash),
    };

    // Calc public input hash based on version.
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
    let public_input_hash = match input.batch_version {
        0 | 1 => batch_info.public_input_hash(&versioned_hashes[0]),
        2 => batch_info.public_input_hash_v2(&versioned_hashes),
        _ => unreachable!("batch version is validated before calculating the hash"),
    };

    #[cfg(not(target_os = "zkvm"))]
    log::info!("public input hash: {public_input_hash:?}");
    Ok(B256::from_slice(public_input_hash.as_slice()))
}
