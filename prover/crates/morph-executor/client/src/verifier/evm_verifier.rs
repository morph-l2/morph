use sbv_core::{BatchInfo, EvmExecutorBuilder, HardforkConfig, VerificationError};
use sbv_primitives::types::BlockTrace;
use sbv_utils::dev_error;

// use Verifier;
pub struct EVMVerifier;

impl EVMVerifier {
    pub fn verify(l2_traces: &Vec<BlockTrace>) -> Result<BatchInfo, VerificationError> {
        let batch_info = execute(l2_traces)?;
        Ok(batch_info)
    }
}

fn execute(traces: &Vec<BlockTrace>) -> Result<BatchInfo, VerificationError> {
    let (batch_info, zktrie_db) = BatchInfo::from_block_traces(&traces);

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
        "root mismatch: root after in trace = {trace_root_after:x}, root after in revm = {revm_root_after:x}" );
        return Err(VerificationError::RootMismatch {
            root_trace: trace_root_after,
            root_revm: revm_root_after,
        });
    }

    drop(executor);
    Ok(batch_info)
}
