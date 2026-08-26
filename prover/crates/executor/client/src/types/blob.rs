use anyhow::anyhow;
use morph_da_decoder_core::decompress_morph_da_zstd;

/// The number of coefficients (BLS12-381 scalars) to represent the blob polynomial in
/// evaluationform.
pub const BLOB_WIDTH: usize = 4096;

#[derive(Clone, Debug)]
pub struct BlobData {}

/// Decode a single blob's BLS12-381 field elements into raw bytes (4096 x 31 bytes).
/// Does NOT decompress — call [`decompress_batch`] on the concatenated output of all blobs.
pub fn decode_blob_scalars(blob_data: &[u8]) -> Result<Vec<u8>, anyhow::Error> {
    if blob_data.len() != BLOB_WIDTH * 32 {
        return Err(anyhow!(
            "invalid blob length: expected {}, got {}",
            BLOB_WIDTH * 32,
            blob_data.len()
        ));
    }
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

/// Alias for [`decode_blob_scalars`] — kept for backward compatibility.
pub fn unpack_blob(blob_data: &[u8]) -> Result<Vec<u8>, anyhow::Error> {
    decode_blob_scalars(blob_data)
}

/// Decode a single blob's scalars and immediately decompress (single-blob / V0/V1 path).
pub fn get_origin_batch(blob_data: &[u8]) -> Result<Vec<u8>, anyhow::Error> {
    let chunk = decode_blob_scalars(blob_data)?;
    decompress_batch(&chunk)
}

pub fn decompress_batch(compressed_batch: &[u8]) -> Result<Vec<u8>, anyhow::Error> {
    let decoded = decompress_morph_da_zstd(compressed_batch)?;
    #[cfg(not(target_os = "zkvm"))]
    log::info!("decompressed_batch: {:?}", decoded.len());
    Ok(decoded)
}
