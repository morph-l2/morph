use crate::types::{
    blob::{decompress_batch, get_origin_batch, unpack_blob},
    input::BlobInfo,
};
use anyhow::anyhow;
use kzg_rs::{get_kzg_settings, Blob as KzgRsBlob, Bytes48};
use prover_primitives::B256;
use sha2::{Digest as _, Sha256};

pub struct BlobVerifier;

impl BlobVerifier {
    /// Verify multiple blobs:
    /// - KZG-verify each blob independently
    /// - Unpack each blob's field elements into its compressed-data chunk
    /// - Concatenate all chunks, then decompress once
    ///
    /// Returns `(versioned_hashes, decompressed_batch_data)`.
    pub fn verify_blobs(blob_infos: &[BlobInfo]) -> Result<(Vec<B256>, Vec<u8>), anyhow::Error> {
        let mut hashes = Vec::new();
        let mut compressed = Vec::new();
        for info in blob_infos {
            let hash = Self::verify_kzg(info)?;
            hashes.push(hash);
            compressed.extend(unpack_blob(&info.blob_data)?);
        }
        let batch_data = decompress_batch(&compressed)?;
        Ok((hashes, batch_data))
    }

    /// KZG-verify a single blob and unpack + decompress its payload (V0/V1 single-blob path).
    pub fn verify(blob_info: &BlobInfo) -> Result<(B256, Vec<u8>), anyhow::Error> {
        let hash = Self::verify_kzg(blob_info)?;
        let origin_batch = get_origin_batch(&blob_info.blob_data)?;
        Ok((hash, origin_batch))
    }

    /// KZG-verify a blob's commitment/proof and return its versioned hash.
    fn verify_kzg(blob_info: &BlobInfo) -> Result<B256, anyhow::Error> {
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
        Ok(B256::from_slice(&versioned_hash))
    }
}

pub fn kzg_to_versioned_hash(commitment: &[u8]) -> Vec<u8> {
    let mut hashed_bytes = Sha256::digest(commitment);
    hashed_bytes[0] = 0x01;
    hashed_bytes.to_vec()
}
