use lazy_static::lazy_static;
use std::fmt;
use tokio::sync::Mutex;

use ethers::{
    core::k256::sha2::{Digest, Sha256},
    types::H256,
};

lazy_static! {
    pub static ref PREV_ROLLUP_L1_BLOCK: Mutex<u64> = Mutex::new(0);
}

const MAX_BLOB_TX_PAYLOAD_SIZE: usize = 131072; // 131072 = 4096 * 32 = 1024 * 4 * 32 = 128kb

#[derive(Debug, Clone)]
pub struct Blob(pub [u8; MAX_BLOB_TX_PAYLOAD_SIZE]);

impl Blob {
    /// blob_bytes =
    ///     be_bytes(num_valid_chunks as u16) ||
    ///     be_bytes(chunks[0].chunk_size as u32) || ...
    ///     be_bytes(chunks[MAX_AGG_SNARKS-1].chunk_size as u32)||all_l2_tx_signed_rlp_in_batch
    pub fn decode_raw_tx_payload(&self) -> Result<Vec<u8>, BlobError> {
        let mut data = vec![0u8; MAX_BLOB_TX_PAYLOAD_SIZE];
        for i in 0..4096 {
            if self.0[i * 32] != 0 {
                return Err(BlobError::InvalidBlob {
                    high_order_byte: self.0[i * 32],
                    field_element: i,
                });
            }
            data[i * 31..i * 31 + 31].copy_from_slice(&self.0[i * 32 + 1..i * 32 + 32]);
        }

        let chunks_len = u16::from_be_bytes(data[0..2].try_into().unwrap()); // size of num_valid_chunks is 2bytes.
        if chunks_len > 15 {
            return Err(BlobError::InvalidData {
                num_valid_chunks: chunks_len,
            });
        }

        let mut data_size: u32 = 0;
        for i in 0..15 {
            let slice = &data[i * 4 + 2..i * 4 + 6]; // size of chunk_size is 4bytes.
            data_size += u32::from_be_bytes(slice.try_into().unwrap());
        }
        let tx_payload_start: usize = 2 + 4 * 15;
        let tx_payload = &data[tx_payload_start..tx_payload_start + data_size as usize];
        Ok(tx_payload.to_vec())
    }
}

#[derive(Debug)]
pub enum BlobError {
    InvalidBlob {
        high_order_byte: u8,
        field_element: usize,
    },
    InvalidData {
        num_valid_chunks: u16,
    },
}

impl fmt::Display for BlobError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            BlobError::InvalidBlob {
                high_order_byte,
                field_element,
            } => write!(
                f,
                "Invalid blob, found non-zero high order byte {:x} of field element {}",
                high_order_byte, field_element
            ),
            BlobError::InvalidData { num_valid_chunks } => write!(
                f,
                "Invalid blob data: num_valid_chunks bigger than 15. num_valid_chunks: {}",
                num_valid_chunks
            ),
        }
    }
}

impl std::error::Error for BlobError {}

pub fn kzg_to_versioned_hash(commitment: &[u8]) -> H256 {
    let mut hasher = Sha256::new();
    hasher.update(commitment);
    let mut hash = hasher.finalize();
    hash[0] = 0x01;
    H256::from_slice(&hash)
}

#[test]
fn test_decode_raw_tx_payload_success() {
    let mut raw_data = vec![];
    // Construct an effective Blob data
    let payload =
        br#"EIP-4844 introduces a new kind of transaction type to Ethereum which accepts "blobs"
        of data to be persisted in the beacon node for a short period of time. These changes are
        forwards compatible with Ethereum's scaling roadmap, and blobs are small enough to keep disk use manageable."#;
    let chunks: Vec<&[u8]> = payload.chunks(123).collect();
    let chunks_len = chunks.len();
    println!("chunks_len {:?}", chunks_len);

    raw_data.extend_from_slice(&(chunks_len as u16).to_be_bytes());

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

    let mut blob_data = [0u8; MAX_BLOB_TX_PAYLOAD_SIZE];
    for (i, &byte) in raw_data.iter().enumerate() {
        blob_data[1 + (i % 31) + 32 * (i / 31)] = byte;
    }

    let blob = Blob(blob_data);
    // Test the decoderaw_tx_payload method
    let result = blob.decode_raw_tx_payload();

    assert!(result.is_ok(), "{}", result.err().unwrap());
    assert_eq!(result.unwrap(), payload);
}
