use crate::types::batch::BatchInfo;
use crate::types::error::ClientError;
use crate::types::input::BlockInput;
use prover_executor_core::MorphExecutor;
use prover_primitives::predeployed::l2_to_l1_message::{
    SEQUENCER_ROOT_ADDRESS, SEQUENCER_ROOT_SLOT, WITHDRAW_ROOT_ADDRESS, WITHDRAW_ROOT_SLOT,
};
use reth_trie::{HashedPostState, KeccakKeyHasher};
use revm::context::BlockEnv;
use revm::database::State;

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
            .any(|w| w[0].current_block.post_state_root != w[1].current_block.prev_state_root)
        {
            return Err(ClientError::DiscontinuousStateRoot);
        }
        execute(blocks)
    }
}

fn execute(mut block_inputs: Vec<BlockInput>) -> Result<BatchInfo, ClientError> {
    // Execute each block sequentially.
    block_inputs.iter_mut().try_for_each(execute_block)?;

    // Find the last post_state with non-empty transactions, or fall back to the last one
    let latest_block = block_inputs.last().expect("block_inputs is non-empty");
    if latest_block.current_block.transactions.is_empty() {
        // If the latest block contains no transactions, verify the MPT state here;
        // otherwise, verify it during transaction execution.
        if latest_block.current_block.post_state_root != latest_block.parent_state.state_root() {
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
        latest_block.current_block.post_state_root,
        post_withdraw_root.into(),
        post_sequencer_root.into(),
    ))
}

fn execute_block(block_input: &mut BlockInput) -> Result<(), ClientError> {
    let block = &block_input.current_block;

    if block.transactions.is_empty() {
        if block.prev_state_root != block.post_state_root {
            // For empty blocks, EVM execution is skipped, but the post root is constrained to equal the previous root.
            return Err(ClientError::MismatchedStateRoot {
                block_num: block.header.number.to::<u64>(),
                root_trace: block.prev_state_root,
                root_revm: block.post_state_root,
            });
        }
        return Ok(());
    }
    let header = &block.header;
    let chain_id = block.chain_id;
    let tx_count = block.transactions.len();
    let block_num = header.number.to::<u64>();

    // Build DB, this will internally verify the correctness of mpt.
    let witness_block = block_input.clone();
    let trie_db = witness_block.witness_db()?;
    // Build evm state from the execution witness.
    let state = State::builder()
        .with_database_ref(&trie_db)
        .with_bundle_update()
        .without_state_clear()
        .build();

    // Build EVM block environment.
    let block_env = BlockEnv {
        number: header.number,
        timestamp: header.timestamp,
        basefee: header.base_fee_per_gas.unwrap_or_default().to::<u64>(),
        gas_limit: header.gas_limit.to::<u64>(),
        beneficiary: block.coinbase,
        ..Default::default()
    };

    let mut core_executor = MorphExecutor::with_hardfork(state, block_env, chain_id);
    // Execute block.
    let bundle_state = core_executor
        .execute_block(&block.transactions)
        .map_err(|e| ClientError::BlockExecutionError(format!("{e:#}")))?;
    // Verify the post-state root by applying the block's transition set to the parent (pre-block) state.
    let computed_state_root = {
        let hashed_post_state =
            HashedPostState::from_bundle_state::<KeccakKeyHasher>(&bundle_state.state);
        block_input.parent_state.update(&hashed_post_state);
        block_input.parent_state.state_root()
    };
    if computed_state_root != block.post_state_root {
        return Err(ClientError::MismatchedStateRoot {
            block_num: header.number.to::<u64>(),
            root_trace: block.post_state_root,
            root_revm: computed_state_root,
        });
    }
    println!("====success execute block_{block_num} in client, txns.len: {tx_count}====");

    Ok(())
}
