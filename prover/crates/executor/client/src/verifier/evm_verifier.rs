use crate::types::batch::BatchInfo;
use crate::types::error::ClientError;
use crate::types::input::BlockInput;
use alloy_consensus::BlockHeader;
use prover_executor_core::MorphExecutor;
use prover_primitives::predeployed::l2_to_l1_message::{
    SEQUENCER_ROOT_ADDRESS, SEQUENCER_ROOT_SLOT, WITHDRAW_ROOT_ADDRESS, WITHDRAW_ROOT_SLOT,
};
use reth_trie::{HashedPostState, KeccakKeyHasher};

// use Verifier;
pub struct EVMVerifier;

impl EVMVerifier {
    pub fn verify(blocks: Vec<BlockInput>) -> Result<BatchInfo, ClientError> {
        // Edge case: nothing to execute.
        if blocks.is_empty() {
            return Err(ClientError::BlockExecutionError(
                "empty batch: no block inputs provided".to_owned(),
            ));
        }
        // Verify that each block's `prev_state_root` matches the previous block's `post_state_root`.
        // This ensures the batch is contiguous.
        if blocks
            .windows(2)
            .any(|w| w[0].current_block.state_root() != w[1].parent_state.state_root())
        {
            return Err(ClientError::DiscontinuousStateRoot);
        }
        execute(blocks)
    }
}

fn execute(mut block_inputs: Vec<BlockInput>) -> Result<BatchInfo, ClientError> {
    // fetch & calculate initial state root
    let prev_state_root = block_inputs.first().unwrap().parent_state.state_root();
    // Execute each block sequentially.
    block_inputs.iter_mut().try_for_each(execute_block)?;

    // Find the last post_state with non-empty transactions, or fall back to the last one
    let latest_block = block_inputs.last().expect("block_inputs is non-empty");
    if latest_block.current_block.body.transactions().collect::<Vec<_>>().is_empty() {
        // If the latest block contains no transactions, verify the MPT state here;
        // otherwise, verify it during transaction execution.
        if latest_block.current_block.state_root() != latest_block.parent_state.state_root() {
            return Err(ClientError::InvalidHeaderStateRoot);
        }
    }

    // The post-withdraw-root & post-sequencer-root is required for public inputs.
    // Tt is derived from the state of the last verified block.
    let post_withdraw_root =
        latest_block.get_storage_value(WITHDRAW_ROOT_ADDRESS, WITHDRAW_ROOT_SLOT)?;
    let post_sequencer_root =
        latest_block.get_storage_value(SEQUENCER_ROOT_ADDRESS, SEQUENCER_ROOT_SLOT)?;

    Ok(BatchInfo::from_block_inputs(
        &block_inputs,
        prev_state_root,
        latest_block.current_block.state_root(),
        post_withdraw_root.into(),
        post_sequencer_root.into(),
    ))
}

fn execute_block(block_input: &mut BlockInput) -> Result<(), ClientError> {
    let block = &block_input.current_block;

    if block.body.transactions.is_empty() {
        if block.state_root() != block_input.parent_state.state_root() {
            // For empty blocks, EVM execution is skipped, but the post root is constrained to equal the previous root.
            return Err(ClientError::MismatchedStateRoot {
                block_num: block.header.number(),
                root_trace: block.state_root(),
                root_revm: block_input.parent_state.state_root(),
            });
        }
        return Ok(());
    }
    let header = &block.header;
    let chain_id = 2818;
    let block_num = block.number();
    let txn_count = block.body.transactions.len();

    // Build DB, this will internally verify the correctness of mpt.
    let witness_block = block_input.clone();
    let trie_db = witness_block.witness_db()?;

    // Execute the whole block via reth's BasicBlockExecutor.
    let core_executor = MorphExecutor::new_ref(trie_db, chain_id);
    let bundle_state = core_executor
        .execute_block(block.clone())
        .map_err(|e| ClientError::BlockExecutionError(format!("{e:#}")))?;
    // Verify the post-state root by applying the block's transition set to the parent (pre-block) state.
    let computed_state_root = {
        let hashed_post_state =
            HashedPostState::from_bundle_state::<KeccakKeyHasher>(&bundle_state.state);
        block_input.parent_state.update(&hashed_post_state);
        block_input.parent_state.state_root()
    };
    if computed_state_root != block.state_root() {
        return Err(ClientError::MismatchedStateRoot {
            block_num: header.number(),
            root_trace: block.state_root(),
            root_revm: computed_state_root,
        });
    }
    #[cfg(not(target_os = "zkvm"))]
    log::debug!("success execute block_{block_num} in client, txns.len: {txn_count}");

    Ok(())
}
