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

        let mut offset: usize = 0;
        let mut chunk_index: u16 = 0;
        let mut payload = Vec::new();
        while offset < MAX_BLOB_TX_PAYLOAD_SIZE {
            let data_len =
                u32::from_le_bytes(data[offset..offset + 4].try_into().unwrap()) as usize;
            if data_len == 0 {
                break;
            }
            let remaining_len = MAX_BLOB_TX_PAYLOAD_SIZE - offset - 4;
            if data_len > remaining_len {
                return Err(BlobError::DecodeError {
                    chunk_index,
                    data_len,
                    remaining_len,
                });
            }
            payload.extend_from_slice(&data[offset + 4..offset + 4 + data_len]);

            let ret = (4 + data_len) / 31;
            let remainder = (4 + data_len) % 31;
            offset += if remainder > 0 { ret + 1 } else { ret } * 31;
            chunk_index += 1;
        }
        log::info!("blob chunk_index = {:?}", chunk_index);
        Ok(payload)
    }
}

#[derive(Debug)]
pub enum BlobError {
    InvalidBlob {
        high_order_byte: u8,
        field_element: usize,
    },
    DecodeError {
        chunk_index: u16,
        data_len: usize,
        remaining_len: usize,
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
            BlobError::DecodeError {
                chunk_index,
                data_len,
                remaining_len,
            } => write!(
                f,
                "Decode error: dataLen is bigger than remainingLen. chunkIndex: {}, dataLen: {}, remainingLen: {}",
                chunk_index, data_len, remaining_len
            ),
        }
    }
}

impl std::error::Error for BlobError {}

pub fn kzg_to_versioned_hash(commitment: &[u8]) -> H256 {
    let mut hasher = Sha256::new();
    hasher.update(commitment);
    let hash = hasher.finalize();
    H256::from_slice(&hash)
}

#[test]
fn test_decode_raw_tx_payload_success() {
    let mut raw_data = [0u8; MAX_BLOB_TX_PAYLOAD_SIZE];
    // Construct an effective Blob data
    let payload =
        br#"EIP-4844 introduces a new kind of transaction type to Ethereum which accepts "blobs"
        of data to be persisted in the beacon node for a short period of time. These changes are
        forwards compatible with Ethereum's scaling roadmap, and blobs are small enough to keep disk use manageable."#;

    let mut offset = 0;
    for chunk in payload.chunks(27) {
        let chunk_len = chunk.len() as u32;
        raw_data[offset + 1..offset + 5].copy_from_slice(&chunk_len.to_le_bytes());
        raw_data[offset + 5..offset + 5 + chunk_len as usize].copy_from_slice(chunk);
        offset += 5 + chunk_len as usize;
    }
    let blob = Blob(raw_data);

    // Test the decoderaw_tx_payload method
    let result = blob.decode_raw_tx_payload();
    assert!(result.is_ok());
    assert_eq!(result.unwrap(), payload);
}
