use anyhow::anyhow;
use kzg_rs::{get_kzg_settings, Blob as KzgRsBlob, Bytes48};

use sbv_primitives::{types::TypedTransaction, B256};
use sha2::{Digest as _, Sha256};

use crate::types::{
    blob::{decode_raw_tx_payload, decode_transactions, get_origin_batch},
    input::BlobInfo,
};

// use Verifier;
pub struct BlobVerifier;

impl BlobVerifier {
    pub fn verify(blob_info: &BlobInfo) -> Result<(B256, Vec<TypedTransaction>), anyhow::Error> {
        // decode
        let origin_batch = get_origin_batch(&blob_info.blob_data).unwrap();
        let raw_tx_payload = decode_raw_tx_payload(origin_batch).unwrap();

        let tx_list: Vec<TypedTransaction> = decode_transactions(raw_tx_payload.as_slice());
        println!("decoded tx_list_len: {:?}", tx_list.len());

        // verify kzg
        let versioned_hash = kzg_to_versioned_hash(&blob_info.commitment);

        let blob = KzgRsBlob::from_slice(&blob_info.blob_data).unwrap();
        let commitent = Bytes48::from_slice(&blob_info.commitment).unwrap();
        let proof = Bytes48::from_slice(&blob_info.proof).unwrap();

        kzg_rs::KzgProof::verify_blob_kzg_proof(blob, &commitent, &proof, &get_kzg_settings())
            .map_err(|e| anyhow!("blob verification failed: {:?}", e))?;
        println!(
            "verify_blob_kzg_proof successfully, versioned_hash: {:?}",
            B256::from_slice(&versioned_hash)
        );

        Ok((B256::from_slice(&versioned_hash), tx_list))
    }
}

pub fn kzg_to_versioned_hash(commitment: &[u8]) -> Vec<u8> {
    let mut hashed_bytes = Sha256::digest(commitment);
    hashed_bytes[0] = 0x01;
    hashed_bytes.to_vec()
}
