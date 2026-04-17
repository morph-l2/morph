use anyhow::{anyhow, Context};
use ruzstd::StreamingDecoder;
use std::io::Read;

/// This magic number is included at the start of a single Zstandard frame
pub const MAGIC_NUM: u32 = 0xFD2F_B528;

/// The number of coefficients (BLS12-381 scalars) to represent the blob polynomial in
/// evaluationform.
pub const BLOB_WIDTH: usize = 4096;

#[derive(Clone, Debug)]
pub struct BlobData {}

/// Unpack a single blob's field elements into its raw compressed-data bytes (4096 × 31 bytes).
/// Does NOT decompress — call [`decompress_batch`] on the concatenated output of all blobs.
pub fn unpack_blob(blob_data: &[u8]) -> Result<Vec<u8>, anyhow::Error> {
    let mut chunk = vec![0u8; BLOB_WIDTH * 31];
    for i in 0..BLOB_WIDTH {
        if blob_data[i * 32] != 0 {
            return Err(anyhow!(
                "Invalid blob, found non-zero high order byte {:x} of field element {}",
                blob_data[i * 32],
                i
            ));
        }
        chunk[i * 31..i * 31 + 31].copy_from_slice(&blob_data[i * 32 + 1..i * 32 + 32]);
    }
    Ok(chunk)
}

/// Unpack a single blob and immediately decompress it (single-blob / V0/V1 path).
pub fn get_origin_batch(blob_data: &[u8]) -> Result<Vec<u8>, anyhow::Error> {
    let chunk = unpack_blob(blob_data)?;
    decompress_batch(&chunk)
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
