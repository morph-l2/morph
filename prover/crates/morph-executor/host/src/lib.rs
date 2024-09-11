use anyhow::anyhow;
use sbv_primitives::types::BlockTrace;
use sbv_primitives::TxTrace;
use std::path::Path;
use std::sync::Arc;

use c_kzg::{Blob, KzgCommitment, KzgProof};
use c_kzg::{Bytes48, KzgSettings};
use morph_executor_client::types::input::BlobInfo;
use once_cell::sync::Lazy;

/// The number of bytes to represent an unsigned 256 bit number.
const N_BYTES_U256: usize = 32;

/// The number of coefficients (BLS12-381 scalars) to represent the blob polynomial in evaluation form.
const BLOB_WIDTH: usize = 4096;

/// The bytes len of one blob.
const BLOB_DATA_SIZE: usize = BLOB_WIDTH * N_BYTES_U256;

/// 4844 trusted setup config
pub static MAINNET_KZG_TRUSTED_SETUP: Lazy<Arc<KzgSettings>> =
    Lazy::new(|| Arc::new(load_trusted_setup()));

/// Loads the trusted setup parameters from the given bytes and returns the [KzgSettings].
pub fn load_trusted_setup() -> KzgSettings {
    let setup_config = "../../configs/4844_trusted_setup.txt";
    let trusted_setup_file = Path::new(setup_config);
    assert!(trusted_setup_file.exists());
    KzgSettings::load_trusted_setup_file(trusted_setup_file).unwrap()
}

pub fn get_blob_info(block_trace: &BlockTrace) -> Result<BlobInfo, anyhow::Error> {
    populate_kzg(&get_blob_data(block_trace))
}

pub fn get_blob_data(block_trace: &BlockTrace) -> [u8; BLOB_DATA_SIZE] {
    let mut coefficients = [[0u8; N_BYTES_U256]; BLOB_WIDTH];

    let tx_bytes = block_trace
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

/// Populate kzg info.
pub fn populate_kzg(blob_bytes: &[u8; BLOB_DATA_SIZE]) -> Result<BlobInfo, anyhow::Error> {
    let kzg_settings: Arc<c_kzg::KzgSettings> = Arc::clone(&MAINNET_KZG_TRUSTED_SETUP);
    let commitment = KzgCommitment::blob_to_kzg_commitment(
        &Blob::from_bytes(blob_bytes).unwrap(),
        &kzg_settings,
    )
    .map_err(|e| anyhow!(format!("generate KzgCommitment error: {:?}", e)))?;

    let proof = KzgProof::compute_blob_kzg_proof(
        &Blob::from_bytes(blob_bytes).unwrap(),
        &Bytes48::from_bytes(commitment.as_slice()).unwrap(),
        &kzg_settings,
    )
    .map_err(|e| anyhow!(format!("generate KzgCommitment error: {:?}", e)))?;
    println!("populate kzg commitment successfully");

    let blob_info = BlobInfo {
        blob_data: blob_bytes.to_vec(),
        commitment: commitment.to_vec(),
        proof: proof.as_slice().to_vec(),
    };
    Ok(blob_info)
}
