use eyre::anyhow;
use std::{error::Error, io::Cursor, vec};

use ethers::{
    core::k256::sha2::{Digest, Sha256},
    types::H256,
};

use super::zstd_util::init_zstd_decoder;

const MAX_BLOB_TX_PAYLOAD_SIZE: usize = 131072; // 131072 = 4096 * 32 = 1024 * 4 * 32 = 128kb
const MAX_AGG_SNARKS: usize = 45;
const METADATA_LENGTH: usize = 2 + 4 * MAX_AGG_SNARKS;

#[derive(Debug, Clone)]
pub struct Blob(pub [u8; MAX_BLOB_TX_PAYLOAD_SIZE]);

impl Blob {
    pub fn get_origin_batch(&self) -> Result<Vec<u8>, BlobError> {
        let compressed_data = self.get_compressed_batch()?;
        decompress_batch(&compressed_data)
    }
    
    pub fn get_compressed_batch(&self) -> Result<Vec<u8>, BlobError> {
        // Decode blob, recovering BLS12-381 scalars.
        let mut data = vec![0u8; MAX_BLOB_TX_PAYLOAD_SIZE];
        for i in 0..4096 {
            if self.0[i * 32] != 0 {
                return Err(BlobError::InvalidBlob(anyhow!(format!(
                    "Invalid blob, found non-zero high order byte {:x} of field element {}",
                    self.0[i * 32],
                    i
                ))));
            }
            data[i * 31..i * 31 + 31].copy_from_slice(&self.0[i * 32 + 1..i * 32 + 32]);
        }

        // detect_zstd_compressed
        let compressed_batch = Self::detect_zstd_compressed(data)?;

        Ok(compressed_batch)
    }

    fn detect_zstd_compressed(decoded_blob: Vec<u8>) -> Result<Vec<u8>, BlobError> {
        // The format of zstd_compression is shown in the following link:
        // https://github.com/facebook/zstd/blob/dev/doc/zstd_compression_format.md#frame_header
        let fcs_field_size = match parse_frame_header_descriptor(&decoded_blob) {
            Ok(0) => 1,
            Ok(1) => 2,
            Ok(2) => 4,
            Ok(3) => 8,
            Ok(_) => {
                return Err(BlobError::Error(anyhow!(
                    "detect_zstd_compressed: unexpected fcs_field_size"
                )))
            } /* unexpected value */
            Err(e) => {
                return Err(BlobError::Error(anyhow!(format!(
                    "parse_frame_header_descriptor error: {:#?}",
                    e
                ))))
            }
        };

        fn get_blocks_size(decoded_blob: &Vec<u8>, start: usize) -> Result<usize, BlobError> {
            let (last_block, _block_type, block_size) =
                parse_block_header(&decoded_blob.to_vec(), start).map_err(|e| {
                    BlobError::Error(anyhow!(format!("parse_block_header error: {:#?}", e)))
                })?;

            if !last_block {
                return get_blocks_size(decoded_blob, block_size as usize + 3 + start);
            }
            // block = Block_Header(3 bytes) + Block_Content
            Ok(block_size as usize + 3 + start)
        }

        // compressed_data = frame_header + frame_content_field_size + zstd_blocks
        let compressed_len = get_blocks_size(&decoded_blob, fcs_field_size)? + 1;

        if compressed_len as usize > MAX_BLOB_TX_PAYLOAD_SIZE - 4096 {
            return Err(BlobError::Error(anyhow!("oversized batch payload")))
        }
        let compressed_batch = decoded_blob[..compressed_len].to_vec();

        // check data
        Self::check_data(&compressed_batch, &decoded_blob, fcs_field_size)?;

        Ok(compressed_batch)
    }

    fn check_data(
        compressed_data: &Vec<u8>,
        decoded_blob: &[u8],
        fcs_field_size: usize,
    ) -> Result<(), BlobError> {
        let origin_batch = decompress_batch(compressed_data)?;

        let mut buffer = [0u8; 8];
        buffer[..fcs_field_size].copy_from_slice(&decoded_blob[1..1 + fcs_field_size]);
        let orgin_content_size =
            u64::from_le_bytes(buffer) + if fcs_field_size == 2 { 256 } else { 0 };

        if origin_batch.len() != orgin_content_size as usize {
            return Err(BlobError::Error(anyhow!(
                "decompress batch error: origin_batch_len is not equal to zstd_orgin_content_size"
            )))
        }

        log::info!(
            "check_blob_data, blob usage {:.3}, batch_compression_ratio: {:.3}",
            compressed_data.len() as f32 / MAX_BLOB_TX_PAYLOAD_SIZE as f32,
            orgin_content_size as f32 / compressed_data.len() as f32
        );

        Self::decode_raw_tx_payload(origin_batch)?;
        Ok(())
    }

    // The format of batch is as follows:
    // origin_batch = be_bytes(num_valid_chunks as u16) || be_bytes(chunks[0].chunk_size as u32) ||
    // ...be_bytes(chunks[MAX_AGG_SNARKS-1].chunk_size as u32)||all_l2_tx_signed_rlp_in_batch
    pub fn decode_raw_tx_payload(origin_batch: Vec<u8>) -> Result<Vec<u8>, BlobError> {
        if origin_batch.len() < METADATA_LENGTH {
            log::warn!("batch.len < METADATA_LENGTH ");
            return Ok(Vec::new());
        }
        let num_valid_chunks = u16::from_be_bytes(origin_batch[0..2].try_into().unwrap()); // size of num_valid_chunks is 2bytes.
        if num_valid_chunks as usize > MAX_AGG_SNARKS {
            return Err(BlobError::InvalidData(anyhow!(format!(
                "Invalid blob data: num_valid_chunks bigger than MAX_AGG_SNARKS. parsed num_valid_chunks: {}",
                num_valid_chunks
            ))));
        }

        let data_size: u64 = origin_batch[2..2 + 4 * num_valid_chunks as usize]
            .chunks_exact(4)
            .map(|chunk| u32::from_be_bytes(chunk.try_into().unwrap()) as u64)
            .sum();

        let tx_payload_end = METADATA_LENGTH + data_size as usize;
        if origin_batch.len() < tx_payload_end {
            return Err(BlobError::InvalidData(anyhow!(
                "The batch does not contain the complete tx_payload"
            )));
        }
        Ok(origin_batch[METADATA_LENGTH..tx_payload_end].to_vec())
    }
}

pub fn decompress_batch(compressed_batch: &Vec<u8>) -> Result<Vec<u8>, BlobError> {
    let cursor = Cursor::new(compressed_batch);
    let mut decoder = init_zstd_decoder(cursor)
        .map_err(|_| BlobError::Error(anyhow!("init_zstd_decoder error")))?;
    let mut decompressed_batch = Vec::new();
    std::io::copy(&mut decoder, &mut decompressed_batch).map_err(|e| {
        BlobError::Error(anyhow!(format!("Failed to decompress data, error:{:?}", e)))
    })?;

    Ok(decompressed_batch)
}

fn parse_frame_header_descriptor(compressed_data: &[u8]) -> Result<u8, Box<dyn Error>> {
    if compressed_data.is_empty() {
        return Err("Compressed data is empty".into());
    }
    let descriptor = compressed_data[0];

    // resolve Frame_Content_Size_flag (2 bits)
    let frame_content_size_flag = (descriptor >> 6) & 0b11;

    Ok(frame_content_size_flag)
}

fn parse_block_header(
    compressed_data: &[u8],
    fcs_field_size: usize,
) -> Result<(bool, u8, u32), Box<dyn Error>> {
    // Make sure we have enough data to parse
    if compressed_data.len() < 1 + fcs_field_size + 3 {
        // 2 (minimum starting point) + 3 (block header size)
        return Err("Compressed batch is too small to contain a valid block header".into());
    }

    // Make sure we have enough data to parse
    if compressed_data.len() < 1 + fcs_field_size + 3 {
        // 2 (minimum starting point) + 3 (block header size)
        return Err("Compressed batch is too small to contain a valid block header".into());
    }

    // Extract the 3-byte header
    // data_block_start_index = fcs_field_size + 1(frame block size);
    let header = &compressed_data[1 + fcs_field_size..1 + fcs_field_size + 3];

    // resolve Last_Block (1 bit)
    let last_block = (header[0] & 0x01) == 1;

    // resolve Block_Type (2 bits)
    let block_type = (header[0] >> 1) & 0x03;

    // resolve Block_Size (21 bits)
    let block_size =
        ((header[0] as u32 >> 3) | ((header[1] as u32) << 5) | ((header[2] as u32) << 13)) &
            0x1FFFFF;

    Ok((last_block, block_type, block_size))
}

#[derive(Debug, thiserror::Error)]
pub enum BlobError {
    #[error("{0}")]
    Error(eyre::Error),
    #[error("{0}")]
    InvalidBlob(eyre::Error),
    #[error("{0}")]
    InvalidData(eyre::Error),
}

pub fn kzg_to_versioned_hash(commitment: &[u8]) -> H256 {
    let mut hasher = Sha256::new();
    hasher.update(commitment);
    let mut hash = hasher.finalize();
    hash[0] = 0x01;
    H256::from_slice(&hash)
}

#[cfg(test)]
mod tests {

    use super::*;
    use crate::da_scalar::zstd_util::{init_zstd_encoder, N_BLOCK_SIZE_TARGET};
    use ethers::utils::hex;
    use std::{fs, io::Write, path::Path};
    pub const N_BLOB_BYTES: usize = 4096 * 31;

    #[test]
    fn test_decode_blob_with_zstd_batch() {
        use crate::da_scalar::{
            calculate::decode_transactions_from_blob, typed_tx::TypedTransaction,
        };

        let blob_bytes = load_zstd_blob();
        let blob = Blob(blob_bytes);

        let result = blob.get_compressed_batch();
        assert!(result.is_ok(), "{}", result.err().unwrap());

        let compressed_batch: Vec<u8> = result.unwrap();
        assert_eq!(compressed_batch.len(), 60576);

        let origin_batch = super::decompress_batch(&compressed_batch).unwrap();
        assert_eq!(origin_batch.len(), 125091);

        let chunks_len = u16::from_be_bytes(origin_batch[0..2].try_into().expect("chunks_len"));
        // size of num_valid_chunks is 2bytes.
        assert_eq!(chunks_len, 11);

        let tx_payload =
            super::Blob::decode_raw_tx_payload(origin_batch).expect("decode_raw_tx_payload");
        assert!(tx_payload.len() == 124909, "tx_payload.len()");

        let txs_decoded: Vec<TypedTransaction> =
            decode_transactions_from_blob(tx_payload.as_slice());
        assert!(txs_decoded.len() == 200, "txs_decoded.len()");
    }

    #[test]
    #[allow(clippy::needless_range_loop)]
    fn test_decode_zstd_working_example() {
        for i in 1..12 {
            let random_batch = gen_batch_data(51200 * i);
            decode_batch_test(&random_batch);
        }
    }

    fn decode_batch_test(random_batch: &[u8]) {
        let batch_data_bytes: Vec<u8> = encode_test_batch_data_bytes(random_batch);

        let encoded_bytes = {
            // Compress batch
            let mut encoder = init_zstd_encoder(N_BLOCK_SIZE_TARGET);
            encoder.set_pledged_src_size(Some(batch_data_bytes.len() as u64)).expect("infallible");
            encoder.write_all(&batch_data_bytes).expect("infallible");

            let encoded_bytes: Vec<u8> = encoder.finish().expect("infallible");
            println!(
                "compress batch data from {} to {}, compression ratio {:.2}, blob usage {:.3}",
                batch_data_bytes.len(),
                encoded_bytes.len(),
                batch_data_bytes.len() as f32 / encoded_bytes.len() as f32,
                encoded_bytes.len() as f32 / N_BLOB_BYTES as f32
            );
            encoded_bytes
        };

        let origin_batch = decompress_batch(&encoded_bytes).unwrap();
        println!(
            "=======origin_batch_len: {:?}, batch_data_bytes_len: {:?}",
            origin_batch.len(),
            batch_data_bytes.len()
        );

        // Encode to blob
        let mut blob_data = [0u8; MAX_BLOB_TX_PAYLOAD_SIZE];
        for (i, &byte) in encoded_bytes.iter().enumerate() {
            blob_data[1 + (i % 31) + 32 * (i / 31)] = byte;
        }
        let blob = Blob(blob_data);

        // Test compressed_batch from blob
        let result = blob.get_compressed_batch();
        assert!(result.is_ok(), "{}", result.err().unwrap());

        let compressed_batch: Vec<u8> = result.unwrap();
        println!("encoded_bytes_len: {:?}", encoded_bytes.len());
        assert_eq!(compressed_batch.len(), encoded_bytes.len());
        assert_eq!(compressed_batch, encoded_bytes);

        let origin_batch = super::decompress_batch(&compressed_batch).unwrap();
        assert_eq!(origin_batch, batch_data_bytes);
    }

    pub fn gen_batch_data(batch_len: usize) -> Vec<u8> {
        let mut batch: Vec<u8> = vec![0; batch_len];
        #[allow(clippy::needless_range_loop)]
        for i in 0..batch_len {
            if 1 & i == 0 {
                batch[i] = (i % 16 + i % 128) as u8 + ((i / 64) % 96) as u8
            } else {
                batch[i] = (i % 121) as u8 + 1
            }
        }
        batch
    }
    pub fn load_zstd_blob() -> [u8; 131072] {
        let blob_data_path = Path::new("data/blob_with_zstd_batch.data");
        let data = fs::read_to_string(blob_data_path).expect("Unable to read file");
        let hex_data: Vec<u8> = hex::decode(data.trim()).unwrap();

        let mut array = [0u8; 131072];
        array.copy_from_slice(&hex_data);
        array
    }

    fn encode_test_batch_data_bytes(payload: &[u8]) -> Vec<u8> {
        let chunk_size = (payload.len() / 45) + 1;
        let chunks: Vec<&[u8]> = payload.chunks(chunk_size).collect();
        let chunks_len = chunks.len();
        println!("chunks_len {:?}", chunks_len);

        let mut raw_data = vec![];
        raw_data.extend_from_slice(&(chunks_len as u16).to_be_bytes());

        #[allow(clippy::needless_range_loop)]
        for i in 0..chunks_len {
            raw_data.extend_from_slice(&(chunks[i].len() as u32).to_be_bytes());
        }
        for chunk in chunks {
            raw_data.extend_from_slice(chunk);
        }
        raw_data
    }
}
