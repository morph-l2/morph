use alloy_consensus::{transaction::SignerRecoverable, BlockHeader, Transaction};
use alloy_network::{BlockResponse, Network};
use alloy_provider::Provider;
use anyhow::{bail, Context};
use morph_revm::MorphTxEnv;
use prover_executor_core::MorphExecutor;
use prover_primitives::{MorphTxEnvelope, B256};
use prover_storage::basic_rpc_db::{BasicRpcDb, RpcDb};
use reth_trie::{HashedPostState, KeccakKeyHasher};
use revm::{
    context::BlockEnv,
    database::{BundleState, State},
    ExecuteEvm,
};

/// An executor that fetches data from a [Provider] to execute blocks in the [ClientExecutor].
#[derive(Debug, Clone)]
pub struct HostExecutor;

impl HostExecutor {
    pub async fn execute_block<P, N>(block_number: u64, provider: &P) -> Result<(), anyhow::Error>
    where
        P: Provider<N> + Clone + std::fmt::Debug,
        N: Network,
        MorphTxEnvelope: From<<N as Network>::TransactionResponse>,
    {
        let block: <N as Network>::BlockResponse = provider
            .get_block_by_number(block_number.into())
            .await?
            .with_context(|| format!("block not found: {block_number}"))?;

        let network_txns: &[<N as Network>::TransactionResponse] = block
            .transactions()
            .as_transactions()
            .context("block response does not contain full transactions")?;

        let evm_txns: Vec<MorphTxEnvelope> = network_txns.iter().cloned().map(Into::into).collect();

        // Init db.
        let rpc_db = BasicRpcDb::new(provider.clone(), block_number, B256::default());
        let bundle_state = BundleState::default();
        let state = State::builder()
            .with_database_ref(&rpc_db)
            .with_bundle_prestate(bundle_state)
            .without_state_clear()
            .build();
        let block_env = BlockEnv::default();

        // Build EVM.
        let mut evm = MorphExecutor::with_hardfork(state, block_env);
        // Execute transactions in block.
        for tx in evm_txns {
            let recovered_from = tx.recover_signer().context("tx recover signer error")?;
            let tx = revm::context::TxEnv {
                caller: recovered_from,
                nonce: tx.nonce(),
                gas_price: tx.gas_price().unwrap_or_default(),
                gas_limit: tx.gas_limit(),
                kind: tx.kind(),
                value: tx.value(),
                data: revm::primitives::Bytes::from(tx.input().to_vec()),
                ..Default::default()
            };
            let morph_tx =
                MorphTxEnv { inner: tx, rlp_bytes: Default::default(), fee_token_id: 1u16 };

            let _rt = evm.inner.transact(morph_tx).context("tx transact error")?;
        }
        let bundle_state = evm.inner.ctx.journaled_state.database.take_bundle();
        let state = rpc_db.state(&bundle_state).await?;
        drop(evm);
        // drop(rpc_db);
        // Verify post state root.

        let hashed_post_state =
            HashedPostState::from_bundle_state::<KeccakKeyHasher>(&bundle_state.state);
        let mut mutated_state = state.clone();
        mutated_state.update(&hashed_post_state);

        let computed_state_root = mutated_state.state_root();
        let expected_state_root = block.header().state_root();
        if computed_state_root != expected_state_root {
            bail!(
                "Mismatched state root after executing the block, block number: {block_number}, expected: {expected_state_root:?}, got: {computed_state_root:?}"
            );
        }
        Ok(())
    }
}

fn _to_prover_tx(_txns: &Vec<alloy_rpc_types::Transaction>) -> Vec<MorphTxEnvelope> {
    vec![]
}
