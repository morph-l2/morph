use alloy::rlp::Decodable;
use anyhow::anyhow;
use ruzstd::StreamingDecoder;
use sbv_primitives::types::TypedTransaction;
use std::io::Read;

/// This magic number is included at the start of a single Zstandard frame
pub const MAGIC_NUM: u32 = 0xFD2F_B528;

/// The number of coefficients (BLS12-381 scalars) to represent the blob polynomial in evaluationform.
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
    println!("cycle-tracker-start: decompress_batch");
    let mut content = MAGIC_NUM.to_le_bytes().to_vec();
    content.append(&mut compressed_batch.to_vec());
    let mut x = content.as_slice();

    let mut decoder = StreamingDecoder::new(&mut x)?;
    let mut result = Vec::new();
    decoder.read_to_end(&mut result).unwrap();
    println!("cycle-tracker-end: decompress_batch");
    println!("decompress_batch: {:?}", result.len());
    Ok(result)
}

pub fn decode_raw_tx_payload(origin_batch: Vec<u8>) -> Result<Vec<u8>, anyhow::Error> {
    let len_bytes: [u8; 32] = origin_batch[0..32]
        .try_into()
        .expect("Slice with incorrect length");
    let mut restored_bytes: [u8; 8] = [0; 8];
    restored_bytes.copy_from_slice(&len_bytes[0..8]);
    let tx_data_len = u64::from_be_bytes(restored_bytes) as usize;
    Ok(origin_batch[32..tx_data_len + 32].to_vec())
}

pub fn decode_transactions(bs: &[u8]) -> Vec<TypedTransaction> {
    decode_transactions_from_blob(bs)
}

fn decode_transactions_from_blob(bs: &[u8]) -> Vec<TypedTransaction> {
    let mut txs_decoded: Vec<TypedTransaction> = Vec::new();

    let mut offset: usize = 0;
    while offset < bs.len() {
        let tx_len = *bs.get(offset).unwrap() as usize;
        if offset + 2 + tx_len > bs.len() {
            break;
        }
        let tx_bytes = bs[offset..offset + 2 + tx_len].to_vec();
        let tx_decoded: TypedTransaction = TypedTransaction::decode(&mut tx_bytes.as_slice())
            .inspect_err(|e| {
                println!("decode_transaction error: {e:?}");
            })
            .unwrap();

        txs_decoded.push(tx_decoded);
        offset += tx_len + 2
    }
    txs_decoded
}
