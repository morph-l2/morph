use alloy::rlp::Decodable;
use anyhow::{anyhow, Ok};
use sbv_primitives::{
    types::{BlockTrace, TransactionTrace, TypedTransaction},
    Transaction, TxTrace, U256,
};

/// The number of coefficients (BLS12-381 scalars) to represent the blob polynomial in evaluation
/// form.
pub const BLOB_WIDTH: usize = 4096;

/// The number of bytes to represent an unsigned 256 bit number.
pub const N_BYTES_U256: usize = 32;

const BLOB_DATA_SIZE: usize = 4096 * 32;

const MAX_BLOB_TX_PAYLOAD_SIZE: usize = 131072; // 131072 = 4096 * 32 = 1024 * 4 * 32 = 128kb

#[derive(Clone, Debug)]
pub struct BlobData {}

pub fn get_blob_data(l2_trace: &BlockTrace) -> [u8; BLOB_DATA_SIZE] {
    let mut coefficients = [[0u8; N_BYTES_U256]; BLOB_WIDTH];

    let tx_bytes = l2_trace
        .transactions
        .iter()
        .filter(|tx| !tx.is_l1_tx())
        .flat_map(|tx| tx.try_build_typed_tx().unwrap().rlp())
        .collect::<Vec<u8>>();

    for (i, byte) in tx_bytes.into_iter().enumerate() {
        coefficients[i / 31][1 + (i % 31)] = byte;
    }

    let mut blob_bytes = [0u8; BLOB_DATA_SIZE];
    for (index, value) in coefficients.iter().enumerate() {
        blob_bytes[index * 32..(index + 1) * 32].copy_from_slice(value.as_slice());
    }
    blob_bytes
}

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
    Vec::<TypedTransaction>::decode(&mut bs).unwrap_or_else(|e| {
        // If decoding fails we need to make an empty block
        println!("error, decode_transactions not successful: {e:?}, use empty tx_list");
        vec![]
    })
}
