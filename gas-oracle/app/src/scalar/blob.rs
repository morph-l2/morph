use eyre::anyhow;
use std::{error::Error, io::Cursor, vec};

use ethers::{
    core::k256::sha2::{Digest, Sha256},
    types::H256,
};

use super::zstd_util::init_zstd_decoder;

const MAX_BLOB_TX_PAYLOAD_SIZE: usize = 131072; // 131072 = 4096 * 32 = 1024 * 4 * 32 = 128kb

#[derive(Debug, Clone)]
pub struct Blob(pub [u8; MAX_BLOB_TX_PAYLOAD_SIZE]);

impl Blob {
    pub fn get_compressed_batch(&self) -> Result<Vec<u8>, BlobError> {
        // Decode blob
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

    pub fn detect_zstd_compressed(decoded_blob: Vec<u8>) -> Result<Vec<u8>, BlobError> {
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

        let (_last_block, _block_type, block_size) =
            parse_block_header(&decoded_blob.to_vec(), fcs_field_size);

        // compressed_data = framheader + contentsize + zstd_block
        let compressed_data =
            decoded_blob[..block_size as usize + 1 + fcs_field_size as usize + 3].to_vec();

        // check data
        Self::check_batch(&compressed_data, &decoded_blob, fcs_field_size)?;

        Ok(compressed_data)
    }

    fn check_batch(
        compressed_batch: &Vec<u8>,
        decoded_blob: &[u8],
        fcs_field_size: usize,
    ) -> Result<(), BlobError> {
        let origin_batch = decompress_batch(compressed_batch)?;

        let mut buffer = [0u8; 8];
        buffer[..fcs_field_size].copy_from_slice(&decoded_blob[1..1 + fcs_field_size]);
        let orgin_content_size =
            u64::from_le_bytes(buffer) + if fcs_field_size == 2 { 256 } else { 0 };
        println!("orgin_content_size: {}", orgin_content_size);

        if origin_batch.len() != orgin_content_size as usize {
            return Err(BlobError::Error(anyhow!(
                "decompress batch error: origin_batch_len is not equal to zstd_orgin_content_size"
            )))
        }
        Self::decode_raw_tx_payload(origin_batch)?;
        Ok(())
    }

    // The format of batch is as follows:
    // origin_batch = be_bytes(num_valid_chunks as u16) || be_bytes(chunks[0].chunk_size as u32) ||
    // ...be_bytes(chunks[MAX_AGG_SNARKS-1].chunk_size as u32)||all_l2_tx_signed_rlp_in_batch
    pub fn decode_raw_tx_payload(origin_batch: Vec<u8>) -> Result<Vec<u8>, BlobError> {
        let chunks_len = u16::from_be_bytes(origin_batch[0..2].try_into().unwrap()); // size of num_valid_chunks is 2bytes.
        if chunks_len > 15 {
            return Err(BlobError::InvalidData(anyhow!(format!(
                "Invalid blob data: num_valid_chunks bigger than 15. num_valid_chunks: {}",
                chunks_len
            ))));
        }

        let mut data_size: u32 = 0;
        for i in 0..15 {
            let slice = &origin_batch[i * 4 + 2..i * 4 + 6]; // size of chunk_size is 4bytes.
            data_size += u32::from_be_bytes(slice.try_into().unwrap());
        }
        let tx_payload_start: usize = 2 + 4 * 15;
        let tx_payload = &origin_batch[tx_payload_start..tx_payload_start + data_size as usize];
        Ok(tx_payload.to_vec())
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

fn parse_frame_header_descriptor(compressed_batch: &[u8]) -> Result<u8, Box<dyn Error>> {
    if compressed_batch.is_empty() {
        return Err("Compressed data is empty".into());
    }

    let descriptor = compressed_batch[0];

    // resolve Frame_Content_Size_flag (2 bits)
    let frame_content_size_flag = (descriptor >> 6) & 0b11;

    // resolve Single_Segment_flag (1 bit)
    let single_segment_flag = (descriptor >> 5) & 0b1;

    // resolve Unused_bit (1 bit)
    let unused_bit = (descriptor >> 4) & 0b1;

    // resolve Reserved_bit (1 bit)
    let reserved_bit = (descriptor >> 3) & 0b1;

    // resolve Content_Checksum_flag (1 bit)
    let content_checksum_flag = (descriptor >> 2) & 0b1;

    // resolve Dictionary_ID_flag (2 bits)
    let dictionary_id_flag = descriptor & 0b11;

    println!("Frame_Content_Size_flag: {}", frame_content_size_flag);
    println!("Single_Segment_flag: {}", single_segment_flag);
    println!("Unused_bit: {}", unused_bit);
    println!("Reserved_bit: {}", reserved_bit);
    println!("Content_Checksum_flag: {}", content_checksum_flag);
    println!("Dictionary_ID_flag: {}", dictionary_id_flag);

    Ok(frame_content_size_flag)
}

fn parse_block_header(decoded_blob: &[u8], fcs_field_size: usize) -> (bool, u8, u32) {
    // Make sure we have enough data to parse
    if decoded_blob.len() < 7 {
        // 4 (starting point) + 3 (header size)
        panic!("Compressed batch is too small to contain a valid block header");
    }

    // Extract the 3-byte header
    // data_block_start_index = fcs_field_size + 1(frame block size);
    let header = &decoded_blob[1 + fcs_field_size..1 + fcs_field_size + 3];

    // resolve Last_Block (1 bit)
    let last_block = (header[0] & 0x01) == 1;

    // resolve Block_Type (2 bits)
    let block_type = (header[0] >> 1) & 0x03;

    // resolve Block_Size (21 bits)
    let block_size =
        ((header[0] as u32 >> 3) | ((header[1] as u32) << 5) | ((header[2] as u32) << 13)) &
            0x1FFFFF;

    (last_block, block_type, block_size)
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
mod test {

    use super::*;
    use crate::scalar::zstd_util::{init_zstd_encoder, N_BLOCK_SIZE_TARGET};
    use std::io::Write;
    pub const N_BLOB_BYTES: usize = 4096 * 31;

    fn get_test_batch_data_bytes(payload: &[u8]) -> Vec<u8> {
        let chunks: Vec<&[u8]> = payload.chunks(123).collect();
        let chunks_len = chunks.len();
        println!("chunks_len {:?}", chunks_len);

        let mut raw_data = vec![];
        raw_data.extend_from_slice(&(chunks_len as u16).to_be_bytes());

        #[allow(clippy::needless_range_loop)]
        for i in 0..15 {
            if i < chunks_len {
                raw_data.extend_from_slice(&(chunks[i].len() as u32).to_be_bytes());
            } else {
                raw_data.extend_from_slice(&(0u32).to_be_bytes());
            }
        }
        for chunk in chunks {
            raw_data.extend_from_slice(chunk);
        }
        raw_data
    }

    #[test]
    #[allow(clippy::needless_range_loop)]
    fn test_decode_zstd_blob() {
        let payload =
        br#"EIP-4844 introduces a new kind of transaction type to Ethereum which accepts "blobs"
        of data to be persisted in the beacon node for a short period of time. These changes are
        forwards compatible with Ethereum's scaling roadmap, and blobs are small enough to keep disk use manageable. and blobs are small enough to keep disk use manageable."#;
        let batch_data_bytes = get_test_batch_data_bytes(payload);

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
        println!("origin_batch: {:?}", String::from_utf8_lossy(&origin_batch));
    }
}
