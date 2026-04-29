use anyhow::{anyhow, Context};
use ruzstd::StreamingDecoder;
use std::io::Read;

/// This magic number is included at the start of a single Zstandard frame
pub const MAGIC_NUM: u32 = 0xFD2F_B528;

/// The number of coefficients (BLS12-381 scalars) to represent the blob polynomial in
/// evaluationform.
pub const BLOB_WIDTH: usize = 4096;

const MAX_BLOB_TX_PAYLOAD_SIZE: usize = 131072; // 131072 = 4096 * 32 = 1024 * 4 * 32 = 128kb

#[derive(Clone, Debug)]
pub struct BlobData {}

pub fn get_origin_batch(blob_data: &[u8]) -> Result<Vec<u8>, anyhow::Error> {
    // Decode blob, recovering BLS12-381 scalars.
    let mut batch_data = vec![0u8; MAX_BLOB_TX_PAYLOAD_SIZE];
    for i in 0..4096 {
        if blob_data[i * 32] != 0 {
            return Err(anyhow!(format!(
                "Invalid blob, found non-zero high order byte {:x} of field element {}",
                blob_data[i * 32],
                i
            )));
        }
        batch_data[i * 31..i * 31 + 31].copy_from_slice(&blob_data[i * 32 + 1..i * 32 + 32]);
    }
    decompress_batch(&batch_data)
}

pub fn decompress_batch(compressed_batch: &[u8]) -> Result<Vec<u8>, anyhow::Error> {
    if compressed_batch.iter().all(|&x| x == 0) {
        // empty batch
        return Ok(Vec::new());
    }

    let mut content = MAGIC_NUM.to_le_bytes().to_vec();
    content.append(&mut compressed_batch.to_vec());
    let mut x = content.as_slice();

    let mut decoder = StreamingDecoder::new(&mut x)?;
    let mut result = Vec::new();
    decoder.read_to_end(&mut result).context("Failed to decompress batch")?;
    #[cfg(not(target_os = "zkvm"))]
    log::info!("decompressed_batch: {:?}", result.len());
    Ok(result)
}
