use anyhow::anyhow;
use sha2::{Digest as _, Sha256};
use std::path::Path;
use std::sync::Arc;

use c_kzg::KzgSettings;
use c_kzg::{Blob, KzgCommitment, KzgProof};
use morph_executor_client::types::input::BlobInfo;
use once_cell::sync::Lazy;
use sbv_primitives::types::BlockTrace;

const BLOB_DATA_SIZE: usize = 4096 * 32;

/// 4844 trusted setup config
pub static MAINNET_KZG_TRUSTED_SETUP: Lazy<Arc<KzgSettings>> =
    Lazy::new(|| Arc::new(load_trusted_setup()));

/// Loads the trusted setup parameters from the given bytes and returns the [KzgSettings].
pub fn load_trusted_setup() -> KzgSettings {
    let setup_config = "configs/4844_trusted_setup.txt";
    let trusted_setup_file = Path::new(setup_config);
    assert!(trusted_setup_file.exists());
    let kzg_settings = KzgSettings::load_trusted_setup_file(trusted_setup_file).unwrap();
    return kzg_settings;
}

pub fn kzg_to_versioned_hash(commitment: &[u8]) -> Vec<u8> {
    let mut hashed_bytes = Sha256::digest(commitment);
    hashed_bytes[0] = 0x01;
    hashed_bytes.to_vec()
}

pub fn populate_kzg(blob_bytes: &[u8; BLOB_DATA_SIZE]) -> Result<BlobInfo, anyhow::Error> {
    let kzg_settings: Arc<c_kzg::KzgSettings> = Arc::clone(&MAINNET_KZG_TRUSTED_SETUP);
    let commitment = KzgCommitment::blob_to_kzg_commitment(
        &Blob::from_bytes(blob_bytes).unwrap(),
        &kzg_settings,
    )
    .map_err(|e| anyhow!(format!("generate KzgCommitment error: {:?}", e)))?;

    // let versioned_hash = kzg_to_versioned_hash(commitment.to_bytes().to_vec().as_slice());

    let mut z = [0u8; 32];
    // challenge.to_big_endian(&mut z);
    let (proof, y) = KzgProof::compute_kzg_proof(
        &Blob::from_bytes(blob_bytes).unwrap(),
        &z.into(),
        &kzg_settings,
    )
    .map_err(|e| anyhow!(format!("generate KzgCommitment error: {:?}", e)))?;

    let blob_info = BlobInfo {
        blob_data: blob_bytes.to_vec(),
        commitment: commitment.to_vec(),
        proof: proof.as_slice().to_vec(),
    };
    Ok(blob_info)
}
