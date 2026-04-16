use crate::types::{blob::get_origin_batch, input::BlobInfo};
use anyhow::anyhow;
use kzg_rs::{get_kzg_settings, Blob as KzgRsBlob, Bytes48};
use prover_primitives::B256;
use sha2::{Digest as _, Sha256};

// use Verifier;
pub struct BlobVerifier;

impl BlobVerifier {
    /// Verify multiple blobs, return (Vec<versioned_hash>, concatenated decoded data).
    pub fn verify_blobs(blob_infos: &[BlobInfo]) -> Result<(Vec<B256>, Vec<u8>), anyhow::Error> {
        let mut hashes = Vec::new();
        let mut all_data = Vec::new();
        for info in blob_infos {
            let (hash, data) = Self::verify(info)?;
            hashes.push(hash);
            all_data.extend(data);
        }
        Ok((hashes, all_data))
    }

    pub fn verify(blob_info: &BlobInfo) -> Result<(B256, Vec<u8>), anyhow::Error> {
        // decode
        let origin_batch = get_origin_batch(&blob_info.blob_data)?;

        // verify kzg
        let versioned_hash = kzg_to_versioned_hash(&blob_info.commitment);
        let blob = KzgRsBlob::from_slice(&blob_info.blob_data).unwrap();
        let commitment = Bytes48::from_slice(&blob_info.commitment).unwrap();
        let proof = Bytes48::from_slice(&blob_info.proof).unwrap();

        let verify_result =
            kzg_rs::KzgProof::verify_blob_kzg_proof(blob, &commitment, &proof, &get_kzg_settings())
                .map_err(|e| anyhow!("blob verification failed, kzg err: {e:?}"))?;
        if !verify_result {
            return Err(anyhow!("The blob kzg verification result is Failed"));
        }
        #[cfg(not(target_os = "zkvm"))]
        log::info!(
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
