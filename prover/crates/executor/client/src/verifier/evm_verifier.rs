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
                "empty batch: no block inputs provided".to_string(),
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
        let batch_info = execute(blocks)?;
        Ok(batch_info)
    }
}

fn execute(mut block_inputs: Vec<BlockInput>) -> Result<BatchInfo, ClientError> {
    // Execute each block sequentially.
    block_inputs
        .iter_mut()
        .filter(|block_input| !block_input.current_block.transactions.is_empty())
        .try_for_each(|block_input| execute_block(block_input))?;

    // Find the last block_input with non-empty transactions, or fall back to the last one
    let last_input = block_inputs
        .iter()
        .rev()
        .find(|block_input| !block_input.current_block.transactions.is_empty())
        .unwrap_or_else(|| block_inputs.last().expect("block_inputs is non-empty (checked above)"));

    // The post-withdraw-root & post-sequencer-root is required for public inputs.
    // Tt is derived from the state of the last verified block.
    let post_withdraw_root =
        last_input.get_storage_value(WITHDRAW_ROOT_ADDRESS, WITHDRAW_ROOT_SLOT)?;
    let post_sequencer_root =
        last_input.get_storage_value(SEQUENCER_ROOT_ADDRESS, SEQUENCER_ROOT_SLOT)?;

    Ok(BatchInfo::from_block_inputs(
        &block_inputs,
        last_input.current_block.post_state_root,
        post_withdraw_root.into(),
        post_sequencer_root.into(),
    ))
}

fn execute_block(block_input: &mut BlockInput) -> Result<(), ClientError> {
    // Clone the parent (pre-block) state and mutate it with the block's transition set to obtain
    // the locally computed post-execution state used for state-root verification.
    let mut state_for_root_verification = block_input.parent_state.clone();

    let block = &block_input.current_block;
    let header = &block.header;
    let chain_id = block.chain_id;
    let tx_count = block.transactions.len();
    let block_num = header.number.to::<u64>();

    // Build DB, this will internally verify the correctness of mpt.
    let trie_db = block_input.witness_db()?;
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
        .map_err(|e| ClientError::BlockExecutionError(e.to_string()))?;

    // Verify the post-state root by applying the block's transition set to the parent (pre-block) state.
    let computed_state_root = {
        let hashed_post_state =
            HashedPostState::from_bundle_state::<KeccakKeyHasher>(&bundle_state.state);
        state_for_root_verification.update(&hashed_post_state);
        state_for_root_verification.state_root()
    };
    if computed_state_root != block.post_state_root {
        return Err(ClientError::MismatchedStateRoot(header.number.to::<u64>()));
    }
    println!("====success execute block_{block_num} in client, txns.len: {tx_count}====");

    Ok(())
}

#[cfg(test)]
mod tests {
    use crate::types::input::BlockInput;
    use crate::verifier::evm_verifier::execute;
    use prover_primitives::types::BlockTrace;
    use std::fs::File;
    use std::io::BufReader;

    #[test]
    fn test_execute_local() {
        // local_transfer_eth
        // mainnet_809
        let block_trace = load_trace("../../../testdata/mpt/local_transfer_eth.json");
        println!("loaded {} block_traces", block_trace.len());
        let blocks: Vec<BlockInput> =
            block_trace.iter().map(|trace| BlockInput::from_trace(trace)).collect();
        println!("blocks len: {:?}", blocks.len());

        let _rt = execute(blocks).unwrap();
    }

    fn load_trace(file_path: &str) -> Vec<BlockTrace> {
        let file = File::open(file_path).unwrap();
        let reader = BufReader::new(file);
        serde_json::from_reader(reader).unwrap()
    }
}
