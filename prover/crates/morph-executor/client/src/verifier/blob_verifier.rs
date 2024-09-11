use kzg_rs::{Bytes32, Bytes48, KzgSettings};
use sbv_core::VerificationError;
use sbv_primitives::{types::TypedTransaction, B256};
use sha2::{Digest as _, Sha256};

use crate::types::{
    blob::{decode_raw_tx_payload, decode_transactions, get_origin_batch},
    input::BlobInfo,
};

// use Verifier;
pub struct BlobVerifier;

impl BlobVerifier {
    pub fn verify(blob_info: &BlobInfo) -> Result<B256, VerificationError> {
        // decode
        let origin_batch = get_origin_batch(&blob_info.blob_data).unwrap();
        let raw_tx_payload = decode_raw_tx_payload(origin_batch).unwrap();

        // verify kzg
        let versioned_hash = kzg_to_versioned_hash(&blob_info.commitment);
        // TODO
        let z = &versioned_hash[0..31];
        let y = &versioned_hash[32..63];
        kzg_rs::KzgProof::verify_kzg_proof(
            &Bytes48::from_slice(&blob_info.commitment).unwrap(),
            &Bytes32::from_slice(z).unwrap(),
            &Bytes32::from_slice(y).unwrap(),
            &Bytes48::from_slice(&blob_info.proof).unwrap(),
            &KzgSettings::load_trusted_setup_file().unwrap(),
        )
        .expect("verify_kzg_proof");

        // rlp decode
        let tx_list: Vec<TypedTransaction> = decode_transactions(raw_tx_payload.as_slice());
        // set tx_list to evm trace
        // TODO

        Ok(B256::from_slice(&versioned_hash))
    }
}

pub fn kzg_to_versioned_hash(commitment: &[u8]) -> Vec<u8> {
    let mut hashed_bytes = Sha256::digest(commitment);
    hashed_bytes[0] = 0x01;
    hashed_bytes.to_vec()
}
