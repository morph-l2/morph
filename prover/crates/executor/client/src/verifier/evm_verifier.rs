use crate::types::batch::BatchInfo;
use crate::types::error::ClientError;
use crate::types::input::BlockInput;
use alloy_consensus::transaction::SignerRecoverable;
use alloy_consensus::Transaction;
use morph_revm::MorphTxEnv;
use prover_executor_core::MorphExecutor;
use prover_primitives::predeployed::l2_to_l1_message::{WITHDRAW_ROOT_ADDRESS, WITHDRAW_ROOT_SLOT};
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
    // Verify that each block's prev_state_root matches the previous block's post_state_root
    if block_inputs
        .windows(2)
        .any(|w| w[0].current_block.post_state_root != w[1].current_block.prev_state_root)
    {
        return Err(ClientError::DiscontinuousStateRoot);
    }

    // Execute each block.
    for block_input in block_inputs.iter_mut() {
        execute_block(block_input)?;
    }

    // Get the post withdrawal root(required for public_input) from the last block.
    let last_block = block_inputs.last().unwrap();
    let post_withdraw_root =
        last_block.get_storage_value(WITHDRAW_ROOT_ADDRESS, WITHDRAW_ROOT_SLOT)?;

    // Construct BatchInfo.
    let batch_info = BatchInfo::from_block_inputs(
        &block_inputs,
        last_block.current_block.post_state_root,
        post_withdraw_root.into(),
    );

    Ok(batch_info)
}

fn execute_block(block_input: &mut BlockInput) -> Result<(), ClientError> {
    //Build db and verify state.
    let trie_db = block_input.witness_db()?;

    let state = State::builder()
        .with_database_ref(&trie_db)
        .with_bundle_update()
        .without_state_clear()
        .build();
    // Build EVM.
    let block_env = BlockEnv {
        number: block_input.current_block.header.number,
        timestamp: block_input.current_block.header.timestamp,
        basefee: block_input.current_block.header.base_fee_per_gas.unwrap_or_default().to::<u64>(),
        gas_limit: block_input.current_block.header.gas_limit.to::<u64>(),
        beneficiary: block_input.current_block.coinbase,
        ..Default::default()
    };

    let mut evm = MorphExecutor::with_hardfork(state, block_env.clone());
    // Execute transactions in block.
    let mut block_input_orgin = block_input.clone();
    let block = &block_input.current_block;
    for tx in &block.transactions {
        let recovered_from = SignerRecoverable::recover_signer(tx)
            .map_err(|_| ClientError::SignatureRecoveryFailed)?;
        let tx_env = revm::context::TxEnv {
            caller: recovered_from,
            nonce: tx.nonce(),
            gas_price: tx.effective_gas_price(Some(block_env.basefee)),
            gas_priority_fee: tx.max_priority_fee_per_gas(),
            gas_limit: tx.gas_limit(),
            kind: tx.kind(),
            value: tx.value(),
            data: revm::primitives::Bytes::from(tx.input().to_vec()),
            ..Default::default()
        };
        let morph_tx =
            MorphTxEnv { inner: tx_env, rlp_bytes: Some(tx.rlp()), ..Default::default() };

        let _rt = evm
            .inner
            .transact_commit(morph_tx)
            .map_err(|e| ClientError::BlockExecutionError(e.to_string()))?;
    }

    evm.inner.ctx.journaled_state.database.merge_transitions(BundleRetention::Reverts);
    let bundle_state = evm.inner.ctx.journaled_state.database.take_bundle();

    // Verify post state root.
    let hashed_post_state =
        HashedPostState::from_bundle_state::<KeccakKeyHasher>(&bundle_state.state);
    block_input_orgin.parent_state.update(&hashed_post_state);

    if block_input_orgin.parent_state.state_root()
        != block_input_orgin.current_block.post_state_root
    {
        return Err(ClientError::MismatchedStateRoot(
            block_input_orgin.current_block.header.number.to::<u64>(),
        ));
    };
    println!(
        "====success execute block_{:?} in client====",
        block_input.current_block.header.number.to::<u64>()
    );

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
