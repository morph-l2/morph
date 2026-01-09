use crate::zstd_util::{init_zstd_encoder, N_BLOCK_SIZE_TARGET};
use prover_executor_client::types::input::BlobInfo;
use prover_primitives::types::blob::{get_blob_data_from_blocks, get_blob_data_from_traces};
use prover_primitives::types::block::L2Block;
use prover_primitives::types::BlockTrace;
use std::{io::Write, sync::Arc};

/// The number of bytes to represent an unsigned 256 bit number.
const N_BYTES_U256: usize = 32;

/// The number of coefficients (BLS12-381 scalars) to represent the blob polynomial in evaluation
/// form.
const BLOB_WIDTH: usize = 4096;

/// The bytes len of one blob.
const BLOB_DATA_SIZE: usize = BLOB_WIDTH * N_BYTES_U256;

pub fn get_blob_info_from_blocks(blocks: &Vec<L2Block>) -> Result<BlobInfo, anyhow::Error> {
    // Assemble batch data from block header and transactions.
    let batch_data = get_blob_data_from_blocks(blocks);

    // Compress batch data and encode into blob format.
    let blob_data = encode_blob(batch_data);

    // Populate kzg commitment & proof.
    populate_kzg(&blob_data)
}

pub fn get_blob_info_from_traces(
    block_traces: &Vec<BlockTrace>,
) -> Result<BlobInfo, anyhow::Error> {
    // Assemble batch data from block header and transactions.

    let batch_data = get_blob_data_from_traces(block_traces);
    // Compress batch data and encode into blob format.

    let blob_data = encode_blob(batch_data);
    // Populate kzg commitment & proof.
    populate_kzg(&blob_data)
}

pub fn encode_blob(tx_bytes: Vec<u8>) -> [u8; 131072] {
    if tx_bytes.is_empty() {
        return [0; 131072];
    }
    // zstd compresse
    let compressed_batch = compresse_batch(tx_bytes.as_slice()).unwrap();

    let mut coefficients = [[0u8; N_BYTES_U256]; BLOB_WIDTH];
    // bls element convert
    for (i, byte) in compressed_batch.into_iter().enumerate() {
        coefficients[i / 31][1 + (i % 31)] = byte;
    }
    let mut blob_bytes = [0u8; BLOB_DATA_SIZE];
    for (index, value) in coefficients.iter().enumerate() {
        blob_bytes[index * 32..(index + 1) * 32].copy_from_slice(value.as_slice());
    }
    blob_bytes
}

/// Populate kzg info.
pub fn populate_kzg(blob_bytes: &[u8]) -> Result<BlobInfo, anyhow::Error> {
    let kzg_settings: Arc<c_kzg::KzgSettings> = c_kzg::ethereum_kzg_settings_arc(8);
    let blob = c_kzg::Blob::from_bytes(blob_bytes)?;

    let commitment = kzg_settings.blob_to_kzg_commitment(&blob)?;
    let proof = kzg_settings.compute_blob_kzg_proof(&blob, &commitment.to_bytes())?;
    let blob_info = BlobInfo {
        blob_data: blob_bytes.to_vec(),
        commitment: commitment.to_vec(),
        proof: proof.as_slice().to_vec(),
    };

    Ok(blob_info)
}

/// zstd compress batch data
pub fn compresse_batch(batch: &[u8]) -> Result<Vec<u8>, anyhow::Error> {
    let mut encoder = init_zstd_encoder(N_BLOCK_SIZE_TARGET);
    encoder.set_pledged_src_size(Some(batch.len() as u64)).expect("infallible");
    encoder.write_all(batch).expect("infallible");

    let encoded_bytes: Vec<u8> = encoder.finish().expect("infallible");
    Ok(encoded_bytes)
}
