pub mod types;
mod verifier;
use sbv_primitives::{TxTrace, B256};
use sha2::{Digest as _, Sha256};
use types::input::{AggregationInput, ClientInput, ShardInfo};
pub use verifier::{blob_verifier::BlobVerifier, evm_verifier::EVMVerifier};

pub fn verify(input: &ClientInput) -> Result<ShardInfo, anyhow::Error> {
    // Verify DA
    let (versioned_hash, batch_data) = BlobVerifier::verify(&input.blob_info).unwrap();

    println!("cycle-tracker-start: traces-to-data");
    let mut tx_bytes: Vec<u8> = vec![];
    for trace in &input.l2_traces {
        let x = trace
            .transactions
            .iter()
            .filter(|tx| !tx.is_l1_tx())
            .flat_map(|tx| tx.try_build_typed_tx().unwrap().rlp())
            .collect::<Vec<u8>>();
        tx_bytes.extend(x);
    }
    println!("cycle-tracker-end: traces-to-data");
    //TODO
    assert_eq!(batch_data, tx_bytes, "blob data mismatch!");

    // Verify EVM exec.
    println!("cycle-tracker-start: evm-verify");
    let batch_info = EVMVerifier::verify(&input.l2_traces).unwrap();
    println!("cycle-tracker-end: evm-verify");

    let shard_info: ShardInfo = ShardInfo {
        chain_id: batch_info.chain_id(),
        prev_state_root: batch_info.prev_state_root(),
        post_state_root: batch_info.post_state_root(),
        withdraw_root: batch_info.withdraw_root(),
        sequencer_root: batch_info.sequencer_root(),
        versioned_hash,
    };
    Ok(shard_info)
}

pub fn verify_agg(agg_input: AggregationInput) -> Result<B256, anyhow::Error> {
    agg_input.shard_infos.windows(2).for_each(|pair| {
        let (prev_boot_info, boot_info) = (&pair[0], &pair[1]);

        // The claimed block of the previous boot info must be the L2 output root of the current
        // boot.
        assert_eq!(prev_boot_info.post_state_root, boot_info.prev_state_root);

        // The chain ID must be the same for all the boot infos, to ensure they're
        // from the same chain and span batch range.
        assert_eq!(prev_boot_info.chain_id, boot_info.chain_id);
    });

    // Verify each range program proof.
    agg_input.shard_infos.iter().for_each(|boot_info| {
        // In the range program, the public values digest is just the hash of the ABI encoded
        // boot info.
        // let serialized_boot_info = bincode::serialize(&boot_info).unwrap();
        let mut shard_bytes: Vec<u8> = vec![];
        shard_bytes.extend(boot_info.chain_id.to_be_bytes());
        shard_bytes.extend(boot_info.prev_state_root.as_slice());
        shard_bytes.extend(boot_info.post_state_root.as_slice());
        shard_bytes.extend(boot_info.withdraw_root.as_slice());
        shard_bytes.extend(boot_info.sequencer_root.as_slice());
        shard_bytes.extend(boot_info.versioned_hash.as_slice());
        let pv_digest = Sha256::digest(shard_bytes);

        if cfg!(target_os = "zkvm") {
            sp1_zkvm::lib::verify::verify_sp1_proof(&agg_input.shard_vkey, &pv_digest.into());
        }
    });

    Ok(B256::default())
}
