use std::collections::BTreeMap;

use c_kzg::{ethereum_kzg_settings, Blob as KzgBlob, Bytes48};
use ethers::{
    prelude::*,
    utils::{hex, rlp},
};
use eyre::anyhow;
use serde_json::Value;

use super::{
    blob::{decompress_batch, kzg_to_versioned_hash, Blob as MorphBlob},
    error::ScalarError,
    typed_tx::TypedTransaction,
    MAX_BLOB_TX_PAYLOAD_SIZE,
};

fn sidecar_index(sidecar: &Value) -> Result<u64, ScalarError> {
    let raw = sidecar.get("index").and_then(Value::as_str).ok_or_else(|| {
        ScalarError::CalculateError(anyhow!("beacon blob sidecar index is missing or not a string"))
    })?;
    if raw.is_empty() || !raw.bytes().all(|byte| byte.is_ascii_digit()) {
        return Err(ScalarError::CalculateError(anyhow!(
            "beacon blob sidecar index is not an unsigned decimal integer"
        )));
    }
    raw.parse::<u64>().map_err(|e| {
        ScalarError::CalculateError(anyhow!(format!(
            "beacon blob sidecar index overflows u64: {e}"
        )))
    })
}

fn decode_beacon_hex(
    sidecar: &Value,
    field: &'static str,
    expected_len: usize,
) -> Result<Vec<u8>, ScalarError> {
    let encoded = sidecar.get(field).and_then(Value::as_str).ok_or_else(|| {
        ScalarError::CalculateError(anyhow!(format!(
            "beacon blob sidecar {field} is missing or not a string"
        )))
    })?;
    let payload = encoded.strip_prefix("0x").ok_or_else(|| {
        ScalarError::CalculateError(anyhow!(format!(
            "beacon blob sidecar {field} must be 0x-prefixed"
        )))
    })?;
    if payload.len() != expected_len.saturating_mul(2) {
        return Err(ScalarError::CalculateError(anyhow!(format!(
            "beacon blob sidecar {field} has invalid encoded length"
        ))));
    }
    let decoded = hex::decode(payload).map_err(|e| {
        ScalarError::CalculateError(anyhow!(format!(
            "beacon blob sidecar {field} is not valid hex: {e}"
        )))
    })?;
    if decoded.len() != expected_len {
        return Err(ScalarError::CalculateError(anyhow!(format!(
            "beacon blob sidecar {field} has invalid decoded length"
        ))));
    }
    Ok(decoded)
}

fn verify_blob_sidecar(
    sidecar: &Value,
    expected_versioned_hash: H256,
) -> Result<[u8; MAX_BLOB_TX_PAYLOAD_SIZE], ScalarError> {
    let commitment_bytes = decode_beacon_hex(sidecar, "kzg_commitment", 48)?;
    let actual_versioned_hash = kzg_to_versioned_hash(&commitment_bytes);
    if expected_versioned_hash != actual_versioned_hash {
        return Err(ScalarError::CalculateError(anyhow!(format!(
            "invalid blob versioned hash: expected {expected_versioned_hash:?}, got {actual_versioned_hash:?}"
        ))));
    }

    let blob_bytes = decode_beacon_hex(sidecar, "blob", MAX_BLOB_TX_PAYLOAD_SIZE)?;
    let proof_bytes = decode_beacon_hex(sidecar, "kzg_proof", 48)?;
    let blob = KzgBlob::from_bytes(&blob_bytes).map_err(|e| {
        ScalarError::CalculateError(anyhow!(format!("invalid KZG blob encoding: {e}")))
    })?;
    let commitment = Bytes48::from_bytes(&commitment_bytes).map_err(|e| {
        ScalarError::CalculateError(anyhow!(format!("invalid KZG commitment encoding: {e}")))
    })?;
    let proof = Bytes48::from_bytes(&proof_bytes).map_err(|e| {
        ScalarError::CalculateError(anyhow!(format!("invalid KZG proof encoding: {e}")))
    })?;
    let verified = ethereum_kzg_settings(0)
        .verify_blob_kzg_proof(&blob, &commitment, &proof)
        .map_err(|e| {
            ScalarError::CalculateError(anyhow!(format!("KZG sidecar verification failed: {e}")))
        })?;
    if !verified {
        return Err(ScalarError::CalculateError(anyhow!(
            "KZG sidecar proof does not match blob and commitment"
        )));
    }

    blob_bytes.try_into().map_err(|_| {
        ScalarError::CalculateError(anyhow!("verified blob has an unexpected byte length"))
    })
}

/// Extract the full batch data from multiple blobs.
/// All blobs belong to the same batch: the batch is compressed as a whole, split into
/// multiple segments (each at most 4096 * 31 bytes), and then encoded into BLS12-381
/// field elements stored in the blobs.
/// This function remains compatible with the single-blob case.
pub(super) fn extract_tx_payload(
    indexed_hashes: Vec<IndexedBlobHash>,
    sidecars: &[Value],
) -> Result<Vec<u8>, ScalarError> {
    let num_blobs = indexed_hashes.len();
    let mut sidecars_by_index = BTreeMap::new();
    for sidecar in sidecars {
        let index = sidecar_index(sidecar)?;
        if sidecars_by_index.insert(index, sidecar).is_some() {
            return Err(ScalarError::CalculateError(anyhow!(format!(
                "beacon response contains duplicate blob sidecar index {index}"
            ))));
        }
    }

    let mut combined_payload = Vec::<u8>::new();
    for i_h in indexed_hashes {
        let sidecar = sidecars_by_index.get(&i_h.index).ok_or_else(|| {
            ScalarError::CalculateError(anyhow!(format!(
                "no blob in response matches desired index: {}",
                i_h.index
            )))
        })?;
        let blob_array = verify_blob_sidecar(sidecar, i_h.hash)?;
        let blob_struct = MorphBlob(blob_array);
        // Extract the raw payload segment by removing only the BLS12-381 encoding,
        // without zstd decompression.
        let payload_bytes = blob_struct.get_payload_bytes().map_err(|e| {
            ScalarError::CalculateError(anyhow!(format!(
                "Failed to get payload bytes from blob, blob_hash: {:?}, err: {}",
                i_h.hash, e
            )))
        })?;
        combined_payload.extend_from_slice(&payload_bytes);
    }

    // After concatenation, use detect_zstd_compressed to trim the valid compressed payload
    // (excluding trailing zero padding), then decompress the batch as a whole.
    let compressed_data =
        MorphBlob::detect_zstd_compressed(combined_payload, num_blobs).map_err(|e| {
            ScalarError::CalculateError(anyhow!(format!(
                "Failed to detect zstd compressed data from combined blob payload: {}",
                e
            )))
        })?;
    decompress_batch(&compressed_data).map_err(|e| {
        ScalarError::CalculateError(anyhow!(format!(
            "Failed to decompress combined blob payload: {}",
            e
        )))
    })
}

pub fn extract_txn_count(origin_batch: &[u8], last_block_num: u64) -> Option<u64> {
    if origin_batch.is_empty() || origin_batch.len() < 8 {
        return None;
    }
    let first_block_num = u64::from_be_bytes(origin_batch[0..8].try_into().unwrap_or_default());
    let block_count = last_block_num.checked_sub(first_block_num)?.checked_add(1)?;
    let block_count = usize::try_from(block_count).ok()?;
    let required_len = 60usize.checked_mul(block_count)?;
    if origin_batch.len() < required_len {
        log::error!("invalid blob batch len");
        return None;
    }
    let mut txn_count_in_batch = 0u64;
    for i in 0..block_count {
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

#[test]
fn extract_txn_count_rejects_reversed_or_truncated_ranges() {
    let mut block = vec![0u8; 60];
    block[..8].copy_from_slice(&10u64.to_be_bytes());
    block[56..58].copy_from_slice(&2u16.to_be_bytes());
    assert_eq!(extract_txn_count(&block, 10), Some(2));
    assert_eq!(extract_txn_count(&block, 9), None);
    assert_eq!(extract_txn_count(&block[..59], 10), None);
}

#[test]
fn forged_sidecar_fails_kzg_verification() {
    let settings = ethereum_kzg_settings(0);
    let blob = KzgBlob::new([0u8; MAX_BLOB_TX_PAYLOAD_SIZE]);
    let commitment = settings.blob_to_kzg_commitment(&blob).unwrap().to_bytes();

    // The proof is valid, but for a different canonical blob/commitment pair.
    // This ensures the negative test reaches the KZG relationship check rather
    // than failing only on malformed point encodings.
    let mut other_blob_bytes = [0u8; MAX_BLOB_TX_PAYLOAD_SIZE];
    other_blob_bytes[31] = 1;
    let other_blob = KzgBlob::new(other_blob_bytes);
    let other_commitment = settings.blob_to_kzg_commitment(&other_blob).unwrap().to_bytes();
    let forged_proof = settings.compute_blob_kzg_proof(&other_blob, &other_commitment).unwrap();

    let sidecar = serde_json::json!({
        "index": "0",
        "blob": format!("0x{}", hex::encode(blob.as_ref())),
        "kzg_commitment": format!("0x{}", hex::encode(commitment.as_ref())),
        "kzg_proof": format!("0x{}", hex::encode(forged_proof.to_bytes().as_ref())),
    });
    let expected_hash = kzg_to_versioned_hash(commitment.as_ref());
    let error = verify_blob_sidecar(&sidecar, expected_hash).unwrap_err();
    assert!(error.to_string().contains("does not match"));
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
