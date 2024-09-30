pub mod types;
mod verifier;
use alloy::{hex, primitives::FixedBytes};
use sbv_core::BatchInfo;
use sbv_primitives::{TxTrace, B256};
use sbv_utils::dev_info;
use sha2::{Digest as _, Sha256};
use tiny_keccak::{Hasher, Keccak};
use types::input::{AggregationInput, ClientInput, ShardInfo};

pub use verifier::{blob_verifier::BlobVerifier, evm_verifier::EVMVerifier};

pub fn verify(input: &ClientInput) -> Result<ShardInfo, anyhow::Error> {
    // Shard blocks hash
    let (l1_data_hash, l2_data_hash) = types::shard::shard_hash(input).unwrap();

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
        l1_data_hash,
        l2_data_hash,
    };
    Ok(shard_info)
}

pub fn verify_agg(agg_input: AggregationInput) -> Result<B256, anyhow::Error> {
    // Verify Blob
    let (versioned_hash, tx_from_blob) = BlobVerifier::verify(&agg_input.blob_info).unwrap();

    let mut tx_from_traces: Vec<u8> = vec![];
    let mut shard_hash: Vec<(B256, B256)> = vec![];

    for shard in &agg_input.l2_traces {
        let mut l1_tx_bytes: Vec<u8> = vec![];
        let mut l2_tx_bytes: Vec<u8> = vec![];

        for trace in shard {
            for tx in &trace.transactions {
                let tx_data = tx.try_build_typed_tx().unwrap().rlp();
                if tx.is_l1_tx() {
                    l1_tx_bytes.extend(tx_data);
                } else {
                    l2_tx_bytes.extend(tx_data);
                }
            }
        }
        tx_from_traces.extend(l2_tx_bytes.clone());

        println!("cycle-tracker-start: agg_l1_data_hash");
        let mut l1_hasher = Keccak::v256();
        l1_hasher.update(&l1_tx_bytes);
        let mut l1_data_hash = B256::ZERO;
        l1_hasher.finalize(&mut l1_data_hash.0);
        println!("cycle-tracker-end: agg_l1_data_hash");

        println!("cycle-tracker-start: agg_l2_data_hash");
        let mut l2_hasher = Keccak::v256();
        l2_hasher.update(&l2_tx_bytes);
        let mut l2_data_hash = B256::ZERO;
        l2_hasher.finalize(&mut l2_data_hash.0);
        println!("cycle-tracker-start: agg_l2_data_hash");

        shard_hash.push((l1_data_hash, l2_data_hash));
    }

    // Constraint for tx_data_from_blob==tx_from_traces
    assert_eq!(tx_from_blob, tx_from_traces, "blob data mismatch!");

    // Constraint all_traces(equivalent to blob) contains shard_trace
    for (shard, shard_info) in shard_hash.iter().zip(agg_input.shard_infos.iter()) {
        assert!(shard.0 == shard_info.l1_data_hash, "shard_info.l1_data_hash");
        assert!(shard.1 == shard_info.l2_data_hash, "shard_info.l2_data_hash");
    }

    // Constraint for state transition
    agg_input.shard_infos.windows(2).for_each(|pair| {
        let (prev_boot_info, boot_info) = (&pair[0], &pair[1]);

        // The block of the previous boot info must be the L2 output root of the current
        // boot.
        assert_eq!(prev_boot_info.post_state_root, boot_info.prev_state_root);

        // The chain ID must be the same for all the block infos, to ensure they're
        // from the same chain and span batch range.
        assert_eq!(prev_boot_info.chain_id, boot_info.chain_id);
    });

    // Constraint for shard state transition.
    // Verify each range program proof.
    agg_input.shard_infos.iter().for_each(|boot_info| {
        let serialized_boot_info = bincode::serialize(&boot_info).unwrap();
        let pv_digest = Sha256::digest(serialized_boot_info);

        if cfg!(target_os = "zkvm") {
            sp1_zkvm::lib::verify::verify_sp1_proof(&agg_input.shard_vkey, &pv_digest.into());
        }
    });

    // public_values
    let batch_data_hash = BatchInfo::batch_data_hash(&agg_input.l2_traces);
    println!(
        "cacl pi hash, prevStateRoot = {:?}, postStateRoot = {:?}, withdrawalRoot = {:?},
        dataHash = {:?}, blobVersionedHash = {:?}, sequencerSetVerifyHash = {:?}",
        hex::encode(agg_input.shard_infos[0].prev_state_root.as_slice()),
        hex::encode(agg_input.shard_infos.last().unwrap().post_state_root.as_slice()),
        hex::encode(agg_input.shard_infos.last().unwrap().withdraw_root.as_slice()),
        hex::encode(batch_data_hash.as_slice()),
        hex::encode(versioned_hash.as_slice()),
        hex::encode(agg_input.shard_infos.last().unwrap().sequencer_root.as_slice()),
    );
    let public_input_hash = pi_hash(agg_input, batch_data_hash, versioned_hash);
    dev_info!("public input hash: {:?}", public_input_hash);
    Ok(public_input_hash)
}

fn pi_hash(
    agg_input: AggregationInput,
    batch_data_hash: FixedBytes<32>,
    versioned_hash: FixedBytes<32>,
) -> FixedBytes<32> {
    let mut hasher = Keccak::v256();
    hasher.update(agg_input.shard_infos[0].chain_id.to_be_bytes().as_slice());
    hasher.update(agg_input.shard_infos[0].prev_state_root.as_slice());
    hasher.update(agg_input.shard_infos.last().unwrap().post_state_root.as_slice());
    hasher.update(agg_input.shard_infos.last().unwrap().withdraw_root.as_slice());
    hasher.update(agg_input.shard_infos.last().unwrap().sequencer_root.as_slice());
    hasher.update(batch_data_hash.as_slice());
    hasher.update(versioned_hash.as_slice());

    let mut public_input_hash = B256::ZERO;
    hasher.finalize(&mut public_input_hash.0);
    public_input_hash
}
