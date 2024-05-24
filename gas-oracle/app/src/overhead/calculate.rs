use ethers::{
    prelude::*,
    utils::{hex, rlp},
};
use lazy_static::lazy_static;
use serde::{Deserialize, Serialize};
use serde_json::Value;
use std::{ops::Mul, str::FromStr};

use crate::read_env_var;

use super::{
    blob::{kzg_to_versioned_hash, Blob},
    typed_tx::TypedTransaction,
    MAX_BLOB_TX_PAYLOAD_SIZE,
};

lazy_static! {
    pub static ref TXN_PER_BATCH: u64 = read_env_var("TXN_PER_BATCH", 50);
}

/// Calculate the transaction overhead for L2 tx.
///
/// # Returns
/// The calculated overhead as a u128 value.
pub(super) fn calc_tx_overhead(
    rollup_gas_used: u128,
    blob_gas_used: f64,
    x: f64,
    l2_data_gas: u64,
    l2_txn: u64,
) -> u128 {
    let mut sys_gas: u128 = rollup_gas_used + 156400 + (blob_gas_used * x).ceil() as u128;
    sys_gas = if sys_gas < l2_data_gas as u128 {
        // Log the difference if system gas is less than L2 data gas
        log::info!("sys_gas - l2_data_gas: {:?}", sys_gas.abs_diff(l2_data_gas.into()));
        0u128
    } else {
        sys_gas - l2_data_gas as u128
    };

    sys_gas / l2_txn.max(*TXN_PER_BATCH) as u128
}

pub(super) fn extract_tx_payload(
    indexed_hashes: Vec<IndexedBlobHash>,
    sidecars: &Vec<Value>,
) -> Result<Vec<u8>, String> {
    let mut tx_payload = Vec::<u8>::new();
    for i_h in indexed_hashes {
        if let Some(sidecar) = sidecars.iter().find(|sidecar| {
            sidecar["index"].as_str().unwrap_or("1000").parse::<u64>().unwrap_or(1000) == i_h.index
        }) {
            let kzg_commitment = sidecar["kzg_commitment"]
                .as_str()
                .ok_or_else(|| "Failed to fetch kzg commitment from blob".to_string())?;
            let decoded_commitment: Vec<u8> =
                hex::decode(kzg_commitment).map_err(|e| e.to_string())?;
            let actual_versioned_hash = kzg_to_versioned_hash(&decoded_commitment);

            if i_h.hash != actual_versioned_hash {
                log::error!(
                    "expected hash {:?} for blob at index {:?} but got {:?}",
                    i_h.hash,
                    i_h.index,
                    actual_versioned_hash
                );
                return Err("Invalid versionedHash for Blob".to_string());
            }

            let encoded_blob = sidecar["blob"]
                .as_str()
                .ok_or_else(|| format!("Missing blob value in blob_hash: {:?}", i_h.hash))?;
            let decoded_blob = hex::decode(encoded_blob).map_err(|e| {
                format!("Failed to decode blob, blob_hash: {:?}, err: {}", i_h.hash, e)
            })?;

            if decoded_blob.len() != MAX_BLOB_TX_PAYLOAD_SIZE {
                return Err("Invalid length for Blob".to_string());
            }

            let blob_array: [u8; MAX_BLOB_TX_PAYLOAD_SIZE] = decoded_blob.try_into().unwrap();
            let blob_struct = Blob(blob_array);
            let mut decoded_payload = blob_struct
                .decode_raw_tx_payload()
                .map_err(|e| format!("Failed to decode blob tx payload: {}", e))?;
            tx_payload.append(&mut decoded_payload);
        } else {
            return Err(format!("no blob in response matches desired index: {}", i_h.index));
        }
    }
    Ok(tx_payload)
}

pub(super) fn extract_txn_num(chunks: Vec<Bytes>) -> Option<u64> {
    if chunks.is_empty() {
        return None;
    }

    let mut txn_in_batch = 0;
    let mut l1_txn_in_batch = 0;
    for chunk in chunks.iter() {
        let mut chunk_bn: Vec<u64> = vec![];
        let bs: &[u8] = chunk;
        // decode blockcontext from chunk
        // |   1 byte   | 60 bytes | ... | 60 bytes |
        // | num blocks |  block 1 | ... |  block n |
        let num_blocks = U256::from_big_endian(bs.get(..1)?);
        for i in 0..num_blocks.as_usize() {
            let block_num = U256::from_big_endian(bs.get((60.mul(i) + 1)..(60.mul(i) + 1 + 8))?);
            let txs_num =
                U256::from_big_endian(bs.get((60.mul(i) + 1 + 56)..(60.mul(i) + 1 + 58))?);
            let l1_txs_num =
                U256::from_big_endian(bs.get((60.mul(i) + 1 + 58)..(60.mul(i) + 1 + 60))?);
            txn_in_batch += txs_num.as_u32();
            l1_txn_in_batch += l1_txs_num.as_u32();
            chunk_bn.push(block_num.as_u64());
        }
    }
    log::debug!("total_txn_in_batch: {:#?}, l1_txn_in_batch: {:#?}", txn_in_batch, l1_txn_in_batch);
    if txn_in_batch < l1_txn_in_batch {
        log::error!("total_txn_in_batch < l1_txn_in_batch");
        return None;
    }
    return Some((txn_in_batch - l1_txn_in_batch) as u64);
}

pub(super) fn data_gas_cost(data: &[u8]) -> u64 {
    if data.len() == 0 {
        return 0;
    }
    let (zeroes, ones) = zeroes_and_ones(data);
    // 4 Paid for every zero byte of data or code for a transaction.
    // 16 Paid for every non-zero byte of data or code for a transaction.
    let zeroes_gas = zeroes * 4;
    let ones_gas = ones * 16;
    zeroes_gas + ones_gas
}

fn zeroes_and_ones(data: &[u8]) -> (u64, u64) {
    let mut zeroes = 0;
    let mut ones = 0;

    for &byt in data {
        if byt == 0 {
            zeroes += 1;
        } else {
            ones += 1;
        }
    }
    (zeroes, ones)
}

#[derive(Debug, Serialize, Deserialize)]
struct BlockInfo {
    block_number: U256,
    timestamp: u64,
    base_fee: U256,
    gas_limit: u64,
    num_txs: u64,
}

#[derive(Debug, Clone)]
pub(super) struct IndexedBlobHash {
    pub(super) index: u64,
    pub(super) hash: H256,
}

pub(super) fn data_and_hashes_from_txs(txs: &[Value], target_tx: &Value) -> Vec<IndexedBlobHash> {
    let mut hashes = Vec::new();
    let mut blob_index = 0u64; // index of each blob in the block's blob sidecar

    for tx in txs {
        // skip any non-batcher transactions
        if tx["hash"] != target_tx["hash"] {
            if let Some(blob_hashes) = tx["blobVersionedHashes"].as_array() {
                blob_index += blob_hashes.len() as u64;
            }
            continue;
        }
        if let Some(blob_hashes) = tx["blobVersionedHashes"].as_array() {
            for h in blob_hashes {
                let idh = IndexedBlobHash {
                    index: blob_index,
                    hash: H256::from_str(h.as_str().unwrap()).unwrap(),
                };
                hashes.push(idh);
                blob_index += 1;
            }
        }
    }
    hashes
}

/// Minimum gas price for data blobs.
pub const MIN_BLOB_GASPRICE: u64 = 1;

/// Controls the maximum rate of change for blob gas price.
pub const BLOB_GASPRICE_UPDATE_FRACTION: u64 = 3338477;

pub(super) fn calc_blob_gasprice(excess_blob_gas: u64) -> u128 {
    fake_exponential(MIN_BLOB_GASPRICE, excess_blob_gas, BLOB_GASPRICE_UPDATE_FRACTION)
}

fn fake_exponential(factor: u64, numerator: u64, denominator: u64) -> u128 {
    assert_ne!(denominator, 0, "attempt to divide by zero");
    let factor = factor as u128;
    let numerator = numerator as u128;
    let denominator = denominator as u128;

    let mut i = 1;
    let mut output = 0;
    let mut numerator_accum = factor * denominator;
    while numerator_accum > 0 {
        output += numerator_accum;

        // Denominator is asserted as not zero at the start of the function.
        numerator_accum = (numerator_accum * numerator) / (denominator * i);
        i += 1;
    }
    output / denominator
}

#[allow(dead_code)]
pub(super) fn decode_transactions_from_blob(bs: &[u8]) -> Vec<TypedTransaction> {
    let mut txs_decoded: Vec<TypedTransaction> = Vec::new();

    let mut offset: usize = 0;
    while offset < bs.len() {
        let tx_len = *bs.get(offset + 1).unwrap() as usize;
        if tx_len == 0 {
            break;
        }
        let tx_bytes = bs[offset..offset + 2 + tx_len].to_vec();
        let tx_decoded: TypedTransaction = match rlp::decode(&tx_bytes) {
            Ok(tx) => tx,
            Err(e) => {
                log::error!("decode_transactions_from_blob error: {:?}", e);
                return vec![];
            }
        };

        txs_decoded.push(tx_decoded);
        offset += tx_len + 2
    }
    txs_decoded
}

#[tokio::test]
async fn test_decode_transactions_from_blob() {
    use ethers::{prelude::*, types::transaction::eip2718::TypedTransaction, utils::to_checksum};

    let wallet: LocalWallet =
        "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80".parse().unwrap();

    let addresses = vec![
        "0x4e6bA705D14b2237374cF3a308ec466cAb24cA6a",
        "0x0425266311AA5858625cD399EADBBfab183494f7",
        "0x1f68c776FBe7285eBe02111F0A982D1640b0a483",
    ];

    let txs: Vec<TypedTransaction> = addresses
        .iter()
        .map(|to| {
            let req = TransactionRequest::new()
                .to(*to)
                .value(1000000000000000000u64)
                .gas(21000)
                .chain_id(1u64);
            TypedTransaction::Legacy(req)
        })
        .collect();

    let mut txs_bytes: Vec<u8> = Vec::new();
    for tx in txs {
        let sig = wallet.sign_transaction(&tx).await.unwrap();
        txs_bytes.extend_from_slice(&tx.rlp_signed(&sig));
    }

    let txs_decoded: Vec<super::typed_tx::TypedTransaction> =
        decode_transactions_from_blob(txs_bytes.as_slice());

    for (tx, address_str) in txs_decoded.iter().zip(addresses) {
        if let super::typed_tx::TypedTransaction::Legacy(tr) = tx.clone() {
            let address_to = tr.to.unwrap();
            let to_tx = to_checksum(address_to.as_address().unwrap(), None);
            assert_eq!(to_tx.as_str(), address_str);
        };
    }
}
