use crate::types::batch::BatchInfo;
use crate::types::error::ClientError;
use crate::types::input::BlockInput;
use alloy_consensus::Transaction;
use alloy_primitives::Address;
use alloy_primitives::{ruint::aliases::U256, uint};
use prover_executor_core::MorphExecutor;
use reth_trie::{HashedPostState, KeccakKeyHasher};
use revm::context::BlockEnv;
use revm::database::{BundleState, State};
use revm::primitives::address;
use revm::ExecuteEvm;

// use Verifier;
pub struct EVMVerifier;

const WITHDRAW_ROOT_ADDRESS: Address = address!("0x5300000000000000000000000000000000000001");
const WITHDRAW_ROOT_SLOT: U256 = uint!(33_U256);

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
    let bundle_state = BundleState::default();
    let state = State::builder()
        .with_database_ref(&trie_db)
        .with_bundle_prestate(bundle_state)
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

    let mut evm = MorphExecutor::with_hardfork(state, block_env);
    // Execute transactions in block.
    let block = &block_input.current_block;
    for tx in &block.transactions {
        let recovered_from =
            tx.get_or_recover_signer().map_err(|_| ClientError::SignatureRecoveryFailed)?;
        let tx_env = revm::context::TxEnv {
            caller: recovered_from,
            nonce: tx.nonce(),
            gas_price: tx.gas_price().unwrap_or_default(),
            gas_limit: tx.gas_limit(),
            kind: tx.kind(),
            value: tx.value(),
            data: revm::primitives::Bytes::from(tx.input().to_vec()),
            ..Default::default()
        };
        let _rt = evm
            .inner
            .transact(tx_env)
            .map_err(|e| ClientError::BlockExecutionError(e.to_string()))?;
    }
    let bundle_state = evm.inner.ctx.journaled_state.database.take_bundle();
    drop(evm);
    drop(trie_db);
    // Verify post state root.
    let hashed_post_state =
        HashedPostState::from_bundle_state::<KeccakKeyHasher>(&bundle_state.state);
    block_input.parent_state.update(&hashed_post_state);
    if block_input.parent_state.state_root() != block_input.current_block.post_state_root {
        return Err(ClientError::MismatchedStateRoot(
            block_input.current_block.header.number.to::<u64>(),
        ));
    };
    Ok(())
}
