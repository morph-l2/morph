use crate::types::batch::BatchInfo;
use crate::types::error::ClientError;
use crate::types::input::BlockInput;
use alloy_consensus::transaction::SignerRecoverable;
use alloy_consensus::Transaction;
use alloy_primitives::Bytes;
use morph_revm::MorphTxEnv;
use prover_executor_core::MorphExecutor;
use prover_primitives::predeployed::l2_to_l1_message::{
    SEQUENCER_ROOT_ADDRESS, SEQUENCER_ROOT_SLOT, WITHDRAW_ROOT_ADDRESS, WITHDRAW_ROOT_SLOT,
};
use reth_trie::{HashedPostState, KeccakKeyHasher};
use revm::context::BlockEnv;
use revm::database::states::bundle_state::BundleRetention;
use revm::database::State;
use revm::ExecuteCommitEvm;

// use Verifier;
pub struct EVMVerifier;

impl EVMVerifier {
    pub fn verify(blocks: Vec<BlockInput>) -> Result<BatchInfo, ClientError> {
        let batch_info = execute(blocks)?;
        Ok(batch_info)
    }
}

fn execute(mut block_inputs: Vec<BlockInput>) -> Result<BatchInfo, ClientError> {
    // Edge case: nothing to execute.
    if block_inputs.is_empty() {
        return Err(ClientError::BlockExecutionError(
            "empty batch: no block inputs provided".to_string(),
        ));
    }

    // Verify that each block's `prev_state_root` matches the previous block's `post_state_root`.
    // This ensures the batch is contiguous.
    if block_inputs
        .windows(2)
        .any(|w| w[0].current_block.post_state_root != w[1].current_block.prev_state_root)
    {
        return Err(ClientError::DiscontinuousStateRoot);
    }

    // Execute each block sequentially.
    block_inputs.iter_mut().try_for_each(|block_input| execute_block(block_input))?;

    let last_input = block_inputs.last().expect("block_inputs is non-empty (checked above)");

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

    // Build DB, this will internally verify the correctness of mpt.
    let trie_db = block_input.witness_db()?;
    // Build evm state from the execution witness.
    let state = State::builder()
        .with_database_ref(&trie_db)
        .with_bundle_update()
        .without_state_clear()
        .build();

    // Build EVM block environment.
    // Note: base_fee_per_gas is optional (pre-EIP1559); defaulting to 0 keeps behavior stable.
    let block_env = BlockEnv {
        number: header.number,
        timestamp: header.timestamp,
        basefee: header.base_fee_per_gas.unwrap_or_default().to::<u64>(),
        gas_limit: header.gas_limit.to::<u64>(),
        beneficiary: block.coinbase,
        ..Default::default()
    };
    let basefee = block_env.basefee;

    let mut evm = MorphExecutor::with_hardfork(state, block_env, chain_id);

    // Execute transactions.
    for tx in &block.transactions {
        let caller = SignerRecoverable::recover_signer(tx)
            .map_err(|_| ClientError::SignatureRecoveryFailed)?;

        let tx_env = revm::context::TxEnv {
            caller,
            nonce: tx.nonce(),
            gas_price: tx.effective_gas_price(Some(basefee)),
            gas_priority_fee: tx.max_priority_fee_per_gas(),
            gas_limit: tx.gas_limit(),
            kind: tx.kind(),
            value: tx.value(),
            data: Bytes::copy_from_slice(tx.input().as_ref()),
            chain_id: Some(chain_id),
            ..Default::default()
        };
        let morph_tx =
            MorphTxEnv { inner: tx_env, rlp_bytes: Some(tx.rlp()), ..Default::default() };

        evm.inner
            .transact_commit(morph_tx)
            .map_err(|e| ClientError::BlockExecutionError(e.to_string()))?;
    }

    // Finalize the transition set into a bundle we can hash for state-root verification.
    evm.inner.ctx.journaled_state.database.merge_transitions(BundleRetention::Reverts);
    let bundle_state = evm.inner.ctx.journaled_state.database.take_bundle();

    // Verify the post-state root by applying the block's transition set to the parent (pre-block) state.
    let hashed_post_state =
        HashedPostState::from_bundle_state::<KeccakKeyHasher>(&bundle_state.state);
    state_for_root_verification.update(&hashed_post_state);
    if state_for_root_verification.state_root() != block.post_state_root {
        return Err(ClientError::MismatchedStateRoot(header.number.to::<u64>()));
    }

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
