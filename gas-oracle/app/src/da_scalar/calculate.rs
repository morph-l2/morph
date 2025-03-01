use ethers::{
    prelude::*,
    utils::{hex, rlp},
};
use eyre::anyhow;
use serde::{Deserialize, Serialize};
use serde_json::Value;

use super::{
    blob::{kzg_to_versioned_hash, Blob},
    error::ScalarError,
    typed_tx::TypedTransaction,
    MAX_BLOB_TX_PAYLOAD_SIZE,
};

pub(super) fn extract_tx_payload(
    indexed_hashes: Vec<IndexedBlobHash>,
    sidecars: &[Value],
) -> Result<Vec<Vec<u8>>, ScalarError> {
    let mut batch_bytes = Vec::<Vec<u8>>::new();
    for i_h in indexed_hashes {
        if let Some(sidecar) = sidecars.iter().find(|sidecar| {
            sidecar["index"].as_str().unwrap_or("1000").parse::<u64>().unwrap_or(1000) == i_h.index
        }) {
            let kzg_commitment = sidecar["kzg_commitment"].as_str().ok_or_else(|| {
                ScalarError::CalculateError(anyhow!("Failed to fetch kzg commitment from blob"))
            })?;
            let decoded_commitment: Vec<u8> =
                hex::decode(kzg_commitment).map_err(|e| ScalarError::CalculateError(e.into()))?;
            let actual_versioned_hash = kzg_to_versioned_hash(&decoded_commitment);

            if i_h.hash != actual_versioned_hash {
                log::error!(
                    "expected hash {:?} for blob at index {:?} but got {:?}",
                    i_h.hash,
                    i_h.index,
                    actual_versioned_hash
                );

                return Err(ScalarError::CalculateError(anyhow!(format!(
                    "Invalid versionedHash for Blob, expected hash {:?} for blob at index {:?} but got {:?}",
                    i_h.hash, i_h.index, actual_versioned_hash
                ))));
            }

            let encoded_blob = sidecar["blob"].as_str().ok_or_else(|| {
                ScalarError::CalculateError(anyhow!(format!(
                    "Missing blob value in blob_hash: {:?}",
                    i_h.hash
                )))
            })?;
            let decoded_blob = hex::decode(encoded_blob).map_err(|e| {
                ScalarError::CalculateError(anyhow!(format!(
                    "Failed to decode blob, blob_hash: {:?}, err: {}",
                    i_h.hash, e
                )))
            })?;

            if decoded_blob.len() != MAX_BLOB_TX_PAYLOAD_SIZE {
                return Err(ScalarError::CalculateError(anyhow!("Invalid length for Blob")));
            }

            let blob_array: [u8; MAX_BLOB_TX_PAYLOAD_SIZE] = decoded_blob.try_into().unwrap();
            let blob_struct = Blob(blob_array);
            let origin_batch = blob_struct.get_origin_batch().map_err(|e| {
                ScalarError::CalculateError(anyhow!(format!(
                    "Failed to decode blob tx payload: {}",
                    e
                )))
            })?;

            batch_bytes.push(origin_batch);
        } else {
            return Err(ScalarError::CalculateError(anyhow!(format!(
                "no blob in response matches desired index: {}",
                i_h.index
            ))));
        }
    }
    Ok(batch_bytes)
}

pub fn extract_txn_count(origin_batch: &Vec<u8>, last_block_num: u64) -> Option<u64> {
    if origin_batch.is_empty() || origin_batch.len() < 8 {
        return None;
    }
    let first_block_num = u64::from_be_bytes(origin_batch[0..8].try_into().unwrap_or_default());
    let block_count = last_block_num - first_block_num + 1;
    if origin_batch.len() < 60 * block_count as usize {
        log::error!("invalid blob batch len");
        return None;
    }
    let mut txn_count_in_batch = 0u64;
    for i in 0..block_count as usize {
        let bys = &origin_batch[60 * i + 56..60 * i + 58];
        let num_txn = u16::from_be_bytes(bys.try_into().unwrap_or_default());

        let bys = &origin_batch[60 * i + 58..60 * i + 60];
        let num_l1_messages = u16::from_be_bytes(bys.try_into().unwrap_or_default());
        if num_txn < num_l1_messages {
            log::error!("total_txn_in_batch < l1_txn_in_batch");
            return None;
        }

        txn_count_in_batch += (num_txn - num_l1_messages) as u64;
    }

    Some(txn_count_in_batch)
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

pub(super) fn data_and_hashes_from_txs(
    txs: &[Transaction],
    target_tx: &Transaction,
) -> Vec<IndexedBlobHash> {
    let mut hashes = Vec::new();
    let mut blob_index = 0u64; // index of each blob in the block's blob sidecar

    for tx in txs {
        let tx_blob_versioned_hashes = tx
            .other
            .get_with("blobVersionedHashes", serde_json::from_value::<Vec<H256>>)
            .unwrap_or(Ok(Vec::<H256>::new()))
            .unwrap_or_default();

        // skip any non-batcher transactions
        if tx.hash != target_tx.hash {
            blob_index += tx_blob_versioned_hashes.len() as u64;
            continue;
        }

        for h in tx_blob_versioned_hashes {
            let idh = IndexedBlobHash { index: blob_index, hash: h };
            hashes.push(idh);
            blob_index += 1;
        }
    }
    hashes
}

#[allow(dead_code)]
pub(super) fn decode_transactions_from_blob(bs: &[u8]) -> Vec<TypedTransaction> {
    let mut txs_decoded: Vec<TypedTransaction> = Vec::new();

    let mut offset: usize = 0;
    while offset < bs.len() {
        if *bs.get(offset).unwrap() < 0xf7 {
            break;
        };
        let tx_len_size = *bs.get(offset).unwrap() as usize - 0xf7;

        let mut tx_len_bytes = [0u8; 4];
        tx_len_bytes[4 - tx_len_size..]
            .copy_from_slice(bs.get(offset + 1..offset + tx_len_size + 1).unwrap_or_default());

        let tx_len = u32::from_be_bytes(tx_len_bytes) as usize;

        let tx_bytes = bs[offset..offset + tx_len_size + tx_len + 1].to_vec();
        let tx_decoded: TypedTransaction = match rlp::decode(&tx_bytes) {
            Ok(tx) => tx,
            Err(e) => {
                log::error!("decode_transactions_from_blob error: {:?}", e);
                return vec![];
            }
        };
        txs_decoded.push(tx_decoded);
        offset += tx_len_size + tx_len + 1
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
