use anyhow::anyhow;
use kzg_rs::{get_kzg_settings, Blob as KzgRsBlob, Bytes48};

use sbv_primitives::B256;
use sha2::{Digest as _, Sha256};

use crate::types::{blob::get_origin_batch, input::BlobInfo};

// use Verifier;
pub struct BlobVerifier;

impl BlobVerifier {
    pub fn verify(blob_info: &BlobInfo) -> Result<(B256, Vec<u8>), anyhow::Error> {
        // decode
        let origin_batch = get_origin_batch(&blob_info.blob_data).unwrap();
        cfg_if::cfg_if! {
            if #[cfg(not(target_os = "zkvm"))] {
                let tx_list = crate::types::blob::decode_transactions(origin_batch.as_slice());
                println!("decoded tx_list_len: {:?}", tx_list.len());
            }
        }

        // verify kzg
        let versioned_hash = kzg_to_versioned_hash(&blob_info.commitment);
        let blob = KzgRsBlob::from_slice(&blob_info.blob_data).unwrap();
        let commitent = Bytes48::from_slice(&blob_info.commitment).unwrap();
        let proof = Bytes48::from_slice(&blob_info.proof).unwrap();

        println!("cycle-tracker-start: verify_blob_kzg_proof");
        let verify_result =
            kzg_rs::KzgProof::verify_blob_kzg_proof(blob, &commitent, &proof, &get_kzg_settings())
                .map_err(|e| anyhow!("blob verification failed, kzg err: {:?}", e))?;
        if !verify_result {
            return Err(anyhow!("The blob kzg verification result is Failed"));
        }
        println!("cycle-tracker-end: verify_blob_kzg_proof");
        println!(
            "verify_blob_kzg_proof successfully, versioned_hash: {:?}",
            B256::from_slice(&versioned_hash)
        );

        Ok((B256::from_slice(&versioned_hash), origin_batch))
    }
}

pub fn kzg_to_versioned_hash(commitment: &[u8]) -> Vec<u8> {
    let mut hashed_bytes = Sha256::digest(commitment);
    hashed_bytes[0] = 0x01;
    hashed_bytes.to_vec()
}
