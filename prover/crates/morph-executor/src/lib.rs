use std::rc::Rc;

use alloy::primitives::keccak256;
use sbv_core::{EvmExecutorBuilder, HardforkConfig, VerificationError};
use sbv_primitives::{types::BlockTrace, zk_trie::ZkMemoryDb, Block, B256};
use sbv_utils::{cycle_track, dev_error, dev_info};

pub fn verify(l2_trace: &BlockTrace) -> Result<B256, VerificationError> {
    // Step1. verify blob
    // TODO
    let versioned_hash: Vec<u8> = vec![];

    // Step2. execute all transactions
    let root_before = l2_trace.root_before();
    let root_after = l2_trace.root_after();
    let withdraw_root = l2_trace.withdraw_root();
    let sequencer_set_root: Vec<u8> = vec![];

    let revm_root_after = execute(l2_trace)?;

    if root_after != revm_root_after {
        dev_error!(
            "Block #{}({:?}) root mismatch: root after in trace = {root_after:x}, root after in
    revm = {revm_root_after:x}",
            l2_trace.number(),
            l2_trace.block_hash(),
        );

        return Err(VerificationError::RootMismatch {
            root_trace: root_after,
            root_revm: revm_root_after,
        });
    }
    dev_info!(
        "Block #{}({}) verified successfully",
        l2_trace.number(),
        l2_trace.block_hash(),
    );

    // Step3. compute pi hash
    let pi_hash = keccak256(
        [
            versioned_hash,
            withdraw_root.to_vec(),
            sequencer_set_root,
            root_before.to_vec(),
            revm_root_after.to_vec(),
        ]
        .concat(),
    );

    Ok(B256::from_slice(pi_hash.as_slice()))
}

fn execute(l2_trace: &BlockTrace) -> Result<alloy::primitives::FixedBytes<32>, VerificationError> {
    let fork_config: HardforkConfig = HardforkConfig::default_from_chain_id(2818);
    let zktrie_db = cycle_track!(
        {
            let mut zktrie_db = ZkMemoryDb::new();
            l2_trace.build_zktrie_db(&mut zktrie_db);
            Rc::new(zktrie_db)
        },
        "build ZktrieState"
    );
    let mut executor = EvmExecutorBuilder::new(zktrie_db.clone())
        .hardfork_config(fork_config)
        .build(&l2_trace)?;
    #[allow(clippy::map_identity)]
    #[allow(clippy::manual_inspect)]
    executor.handle_block(&l2_trace).map_err(|e| {
        dev_error!(
            "Error occurs when executing block #{}({:?}): {e:?}",
            l2_trace.number(),
            l2_trace.block_hash()
        );
        e
    })?;
    let revm_root_after = executor.commit_changes(&zktrie_db);
    Ok(revm_root_after)
}
