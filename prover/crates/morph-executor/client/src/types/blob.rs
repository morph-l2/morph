use alloy::rlp::Decodable;
use anyhow::anyhow;
use sbv_primitives::{
    types::{BlockTrace, TransactionTrace, TypedTransaction},
    Transaction, TxTrace, U256,
};

/// The number of coefficients (BLS12-381 scalars) to represent the blob polynomial in evaluationform.
pub const BLOB_WIDTH: usize = 4096;

const MAX_BLOB_TX_PAYLOAD_SIZE: usize = 131072; // 131072 = 4096 * 32 = 1024 * 4 * 32 = 128kb

#[derive(Clone, Debug)]
pub struct BlobData {}

pub fn get_origin_batch(blob_data: &Vec<u8>) -> Result<Vec<u8>, anyhow::Error> {
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

pub fn decompress_batch(compressed_batch: &Vec<u8>) -> Result<Vec<u8>, anyhow::Error> {
    // let mut decoder = init_zstd_decoder(cursor);
    Ok(compressed_batch.to_vec())
}

pub fn decode_raw_tx_payload(origin_batch: Vec<u8>) -> Result<Vec<u8>, anyhow::Error> {
    Ok(origin_batch)
}

pub fn decode_transactions(mut bs: &[u8]) -> Vec<TypedTransaction> {
    // Vec::<TypedTransaction>::decode(&mut bs).unwrap_or_else(|e| {
    //     // If decoding fails we need to make an empty block
    //     println!("decode_transactions error: {e:?}, use empty tx_list");
    // })
    vec![]
}
