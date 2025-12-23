pub mod types;
mod verifier;
use alloy_primitives::hex;
use prover_primitives::B256;
use prover_utils::profile_report;
use types::input::ExecutorInput;
pub use verifier::{blob_verifier::BlobVerifier, evm_verifier::EVMVerifier};

pub const EVM_VERIFY: &str = "evm verify";

pub fn verify(input: ExecutorInput) -> Result<B256, anyhow::Error> {
    // Verify DA
    let num_blocks = input.block_inputs.len();
    let (versioned_hash, batch_data) = BlobVerifier::verify(&input.blob_info, num_blocks).unwrap();
    println!("cycle-tracker-start: traces-to-data");
    let mut batch_from_trace: Vec<u8> = Vec::with_capacity(num_blocks * 60);
    let mut tx_bytes: Vec<u8> = vec![];
    for trace in &input.block_inputs {
        // BlockContext
        let mut block_ctx: Vec<u8> = Vec::with_capacity(60);
        block_ctx.extend_from_slice(&trace.current_block.header.number.to::<u64>().to_be_bytes());
        block_ctx
            .extend_from_slice(&trace.current_block.header.timestamp.to::<u64>().to_be_bytes());
        block_ctx.extend_from_slice(
            &trace.current_block.header.base_fee_per_gas.unwrap_or_default().to_be_bytes::<32>(),
        );
        block_ctx
            .extend_from_slice(&trace.current_block.header.gas_limit.to::<u64>().to_be_bytes());
        block_ctx.extend_from_slice(&(trace.current_block.transactions.len() as u16).to_be_bytes());
        block_ctx.extend_from_slice(&(trace.current_block.num_l1_txs() as u16).to_be_bytes());
        batch_from_trace.extend(block_ctx);

        // Collect txns
        let x = trace
            .current_block
            .transactions
            .iter()
            .filter(|tx| !tx.is_l1_msg())
            .flat_map(|tx| tx.rlp())
            .collect::<Vec<u8>>();
        tx_bytes.extend(x);
    }
    batch_from_trace.extend(tx_bytes);

    println!("cycle-tracker-end: traces-to-data");
    assert_eq!(batch_data, batch_from_trace, "blob data mismatch!");

    // Verify EVM exec.
    let batch_info = profile_report!(EVM_VERIFY, { EVMVerifier::verify(input.block_inputs) })?;

    // Calc public input hash.
    println!("cycle-tracker-start: cacl_public_input_hash");
    println!(
        "cacl pi hash, prevStateRoot = {:?}, postStateRoot = {:?}, withdrawalRoot = {:?},
        dataHash = {:?}, blobVersionedHash = {:?}, sequencerSetVerifyHash = {:?}",
        hex::encode(batch_info.prev_state_root().as_slice()),
        hex::encode(batch_info.post_state_root().as_slice()),
        hex::encode(batch_info.withdraw_root().as_slice()),
        hex::encode(batch_info.data_hash().as_slice()),
        hex::encode(versioned_hash.as_slice()),
        hex::encode(batch_info.sequencer_root().as_slice()),
    );
    let public_input_hash = batch_info.public_input_hash(&versioned_hash);
    println!("cycle-tracker-end: cacl_public_input_hash");
    println!("public input hash: {:?}", public_input_hash);
    Ok(B256::from_slice(public_input_hash.as_slice()))
}
