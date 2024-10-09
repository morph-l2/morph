use alloy::primitives::ruint::aliases::U256;
use sbv_core::{BatchInfo, EvmExecutorBuilder, HardforkConfig, VerificationError};
use sbv_primitives::{types::BlockTrace, Address};
use sbv_utils::dev_error;
use std::str::FromStr;

// use Verifier;
pub struct EVMVerifier;

impl EVMVerifier {
    pub fn verify(l2_traces: &[BlockTrace]) -> Result<BatchInfo, VerificationError> {
        let batch_info = execute(l2_traces)?;
        Ok(batch_info)
    }
}

fn execute(traces: &[BlockTrace]) -> Result<BatchInfo, VerificationError> {
    println!("cycle-tracker-start: zktrie_db");
    let (mut batch_info, zktrie_db) = BatchInfo::from_block_traces(traces);
    println!("cycle-tracker-end: zktrie_db");

    let fork_config: HardforkConfig = HardforkConfig::default_from_chain_id(2818);
    let mut executor = EvmExecutorBuilder::new(zktrie_db.clone())
        .hardfork_config(fork_config)
        .build(&traces[0])?;
    #[allow(clippy::map_identity)]
    #[allow(clippy::manual_inspect)]
    executor.handle_block(&traces[0])?;
    for trace in traces[1..].iter() {
        executor.update_db(trace)?;
        executor.handle_block(trace)?;
    }

    let trace_root_after = batch_info.post_state_root();
    let revm_root_after = executor.commit_changes(&zktrie_db);
    if revm_root_after != batch_info.post_state_root() {
        dev_error!(
            "root mismatch: root after in trace = {trace_root_after:x}, root after in revm = {revm_root_after:x}"
        );
        return Err(VerificationError::RootMismatch {
            root_trace: trace_root_after,
            root_revm: revm_root_after,
        });
    }

    // post_sequencer_root;
    let withdraw_root = executor.get_storage_value(
        Address::from_str("0x5300000000000000000000000000000000000001").unwrap(),
        U256::from(33),
    );
    // post_sequencer_root;
    let sequencer_root = executor.get_storage_value(
        Address::from_str("0x5300000000000000000000000000000000000017").unwrap(),
        U256::from(101),
    );

    batch_info.withdraw_root = Some(withdraw_root.into());
    batch_info.sequencer_root = Some(sequencer_root.into());

    drop(executor);
    Ok(batch_info)
}
