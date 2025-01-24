use anyhow::anyhow;
use morph_executor_utils::read_env_var;
use sbv_primitives::{types::BlockTrace, Block, TxTrace};
use std::{io::Write, path::Path, sync::Arc};
use zstd_util::{init_zstd_encoder, N_BLOCK_SIZE_TARGET};

use c_kzg::{Blob, Bytes48, KzgCommitment, KzgProof, KzgSettings};
use morph_executor_client::types::input::BlobInfo;
use once_cell::sync::Lazy;

mod zstd_util;

/// The number of bytes to represent an unsigned 256 bit number.
const N_BYTES_U256: usize = 32;

/// The number of coefficients (BLS12-381 scalars) to represent the blob polynomial in evaluation
/// form.
const BLOB_WIDTH: usize = 4096;

/// The bytes len of one blob.
const BLOB_DATA_SIZE: usize = BLOB_WIDTH * N_BYTES_U256;

/// 4844 trusted setup config
pub static MAINNET_KZG_TRUSTED_SETUP: Lazy<Arc<KzgSettings>> =
    Lazy::new(|| Arc::new(load_trusted_setup()));

/// Loads the trusted setup parameters from the given bytes and returns the [KzgSettings].
pub fn load_trusted_setup() -> KzgSettings {
    let setup_config =
        read_env_var("TRUSTED_SETUP_4844", "configs/4844_trusted_setup.txt".to_string());
    let trusted_setup_file = Path::new(&setup_config);
    assert!(trusted_setup_file.exists());
    KzgSettings::load_trusted_setup_file(trusted_setup_file).unwrap()
}

pub fn get_blob_info(block_trace: &Vec<BlockTrace>) -> Result<BlobInfo, anyhow::Error> {
    let batch_info = get_blob_data(block_trace);
    populate_kzg(&batch_info)
}

pub fn get_blob_data(block_trace: &Vec<BlockTrace>) -> [u8; BLOB_DATA_SIZE] {
    let num_blocks = block_trace.len();
    let mut batch_from_trace: Vec<u8> = Vec::with_capacity(num_blocks * 60);
    let mut tx_bytes: Vec<u8> = vec![];
    for trace in block_trace {
        // BlockContext
        // https://github.com/morph-l2/morph/blob/main/contracts/contracts/libraries/codec/BatchCodecV0.sol
        let mut block_ctx: Vec<u8> = Vec::with_capacity(60);
        block_ctx.extend_from_slice(&trace.number().to_be_bytes());
        block_ctx.extend_from_slice(&trace.timestamp().to::<u64>().to_be_bytes());
        block_ctx
            .extend_from_slice(&trace.base_fee_per_gas().unwrap_or_default().to_be_bytes::<32>());
        block_ctx.extend_from_slice(&trace.gas_limit().to::<u64>().to_be_bytes());
        block_ctx.extend_from_slice(&(trace.transactions.len() as u16).to_be_bytes());
        block_ctx.extend_from_slice(&(trace.num_l1_txs() as u16).to_be_bytes());
        batch_from_trace.extend(block_ctx);

        // Collect txns
        let x = trace
            .transactions
            .iter()
            .filter(|tx| !tx.is_l1_tx())
            .flat_map(|tx| tx.try_build_typed_tx().unwrap().rlp())
            .collect::<Vec<u8>>();
        tx_bytes.extend(x);
    }
    batch_from_trace.extend(tx_bytes);
    encode_blob(batch_from_trace)
}

pub fn encode_blob(tx_bytes: Vec<u8>) -> [u8; 131072] {
    // zstd compresse
    if tx_bytes.is_empty() {
        return [0; 131072];
    }
    let compressed_batch = compresse_batch(tx_bytes.as_slice()).unwrap();

    let mut coefficients = [[0u8; N_BYTES_U256]; BLOB_WIDTH];
    // bls element convert
    for (i, byte) in compressed_batch.into_iter().enumerate() {
        coefficients[i / 31][1 + (i % 31)] = byte;
    }
    let mut blob_bytes = [0u8; BLOB_DATA_SIZE];
    for (index, value) in coefficients.iter().enumerate() {
        blob_bytes[index * 32..(index + 1) * 32].copy_from_slice(value.as_slice());
    }
    blob_bytes
}

/// Populate kzg info.
pub fn populate_kzg(blob_bytes: &[u8]) -> Result<BlobInfo, anyhow::Error> {
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

    let blob_info = BlobInfo {
        blob_data: blob_bytes.to_vec(),
        commitment: commitment.to_vec(),
        proof: proof.as_slice().to_vec(),
    };
    Ok(blob_info)
}

pub fn compresse_batch(batch: &[u8]) -> Result<Vec<u8>, anyhow::Error> {
    let mut encoder = init_zstd_encoder(N_BLOCK_SIZE_TARGET);
    encoder.set_pledged_src_size(Some(batch.len() as u64)).expect("infallible");
    encoder.write_all(batch).expect("infallible");

    let encoded_bytes: Vec<u8> = encoder.finish().expect("infallible");
    Ok(encoded_bytes)
}
