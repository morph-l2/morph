pub mod types;
mod verifier;
use sbv_primitives::{TxTrace, B256};
use sbv_utils::dev_info;
use types::input::ClientInput;
use verifier::{blob_verifier::BlobVerifier, evm_verifier::EVMVerifier};

pub fn verify(input: &ClientInput) -> Result<B256, anyhow::Error> {
    // Verify DA
    let (versioned_hash, batch_data) = BlobVerifier::verify(&input.blob_info).unwrap();

    println!("cycle-tracker-start: traces-to-data");
    let mut tx_bytes: Vec<u8> = vec![];
    for trace in &input.l2_traces {
        let x = trace
            .transactions
            .iter()
            .filter(|tx| !tx.is_l1_tx())
            .flat_map(|tx| tx.try_build_typed_tx().unwrap().rlp_da())
            .collect::<Vec<u8>>();
        tx_bytes.extend(x);
    }
    println!("cycle-tracker-end: traces-to-data");
    println!("cycle-tracker-start: batch-check");
    assert_eq!(batch_data, tx_bytes, "blob data mismatch!");
    println!("cycle-tracker-end: batch-check");

    // Verify EVM exec.
    let batch_info = EVMVerifier::verify(&input.l2_traces).unwrap();

    // Calc public input hash.
    println!("cycle-tracker-start: cacl_public_input_hash");
    let public_input_hash = batch_info.public_input_hash(&versioned_hash);
    println!("cycle-tracker-end: cacl_public_input_hash");
    dev_info!("public input hash: {:?}", public_input_hash);
    Ok(B256::from_slice(public_input_hash.as_slice()))
}
