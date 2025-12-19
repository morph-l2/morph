use crate::types::batch::BatchInfo;
use crate::types::error::ClientError;
use crate::types::input::BlockInput;
use alloy::primitives::ruint::aliases::U256;
use alloy::uint;
use alloy_consensus::Transaction;
use morph_executor_core::EthEvm;
use prover_primitives::Address;
use reth_trie::{HashedPostState, KeccakKeyHasher};
use revm::context::{BlockEnv, CfgEnv};
use revm::database::State;
use revm::primitives::address;
use revm::ExecuteEvm;

// use Verifier;
pub struct EVMVerifier;

const WITHDRAW_ROOT_ADDRESS: Address = address!("0x5300000000000000000000000000000000000001");
const WITHDRAW_ROOT_SLOT: U256 = uint!(33_U256);

const SEQUENCER_ROOT_ADDRESS: Address = address!("0x5300000000000000000000000000000000000017");
const SEQUENCER_ROOT_SLOT: U256 = uint!(117_U256);
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

    for block_input in block_inputs.iter_mut() {
        //Verify the db's state root, and construct the account and storage values by reading from the state tries.
        let trie_db = block_input.witness_db().unwrap();
        let state = State::builder().with_database_ref(&trie_db).build();
        let mut evm = EthEvm::new(state, BlockEnv::default(), CfgEnv::default());

        // Execute all transactions in the block
        let block = &block_input.current_block;
        for tx in &block.transactions {
            let tx_env = revm::context::TxEnv {
                nonce: tx.nonce(),
                gas_price: tx.gas_price().unwrap_or_default(),
                gas_limit: tx.gas_limit(),
                kind: tx.kind(),
                value: tx.value(),
                data: revm::primitives::Bytes::from(tx.input().to_vec()),
                ..Default::default()
            };
            let _rt = evm.inner.transact(tx_env);
        }

        // Apply all state changes
        let bundle_state = evm.inner.ctx.journaled_state.database.take_bundle();
        let hashed_post_state =
            HashedPostState::from_bundle_state::<KeccakKeyHasher>(bundle_state.state());
        block_input.parent_state.update(&hashed_post_state);

        // Verify the post state root
        if block_input.parent_state.state_root() != block_input.current_block.post_state_root {
            return Err(ClientError::MismatchedStateRoot(
                block_input.current_block.header.number.to::<u64>(),
            ));
        }
    }

    let last_block = block_inputs.last().unwrap();
    let post_withdraw_root =
        last_block.get_storage_value(WITHDRAW_ROOT_ADDRESS, WITHDRAW_ROOT_SLOT)?;

    let batch_info = BatchInfo::from_block_inputs(
        &block_inputs,
        last_block.current_block.post_state_root,
        post_withdraw_root.into(),
    );

    Ok(batch_info)
}
