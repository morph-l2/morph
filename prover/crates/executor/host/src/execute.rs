use alloy_consensus::{transaction::SignerRecoverable, Transaction};
use alloy_provider::{DynProvider, Provider};
use anyhow::{bail, Context};
use morph_revm::MorphTxEnv;
use prover_executor_core::MorphExecutor;
use prover_primitives::{
    predeployed::l2_to_l1_message::{WITHDRAW_ROOT_ADDRESS, WITHDRAW_ROOT_SLOT},
    TxTrace,
};
use prover_storage::basic_rpc_db::{BasicRpcDb, RpcDb};
use reth_trie::{HashedPostState, KeccakKeyHasher};
use revm::{
    context::BlockEnv,
    database::{states::bundle_state::BundleRetention, State},
    ExecuteCommitEvm,
};

use crate::utils::{query_block, query_state_root, HostExecutorOutput, CHAIN_CONFIG};

/// An executor that fetches data from a [Provider] to execute blocks in the [ClientExecutor].
#[derive(Debug, Clone)]
pub struct HostExecutor;

impl HostExecutor {
    /// Executes `block_number` by:
    /// 1) fetching block + previous state root
    /// 2) executing all txs against an RPC-backed DB
    /// 3) verifying the computed post-state root
    pub async fn execute_block(
        block_number: u64,
        provider: &DynProvider,
    ) -> Result<HostExecutorOutput, anyhow::Error> {
        // We need a previous block root to initialize the RPC-backed DB.
        // Using `checked_sub` avoids underflow panics for genesis blocks.
        let prev_block_number = block_number
            .checked_sub(1)
            .context("HostExecutor::execute_block requires block_number > 0 (needs prev state)")?;

        let block = query_block(block_number, provider)
            .await
            .with_context(|| format!("query_block failed for block {block_number}"))?;

        let chain_id =
            provider.get_chain_id().await.context("failed to fetch chain_id from provider")?;

        let beneficiary = *CHAIN_CONFIG
            .get(&chain_id)
            .with_context(|| format!("chain_id {chain_id} not found in CHAIN_CONFIG"))?;

        let disk_root = query_state_root(block_number, provider)
            .await
            .with_context(|| format!("query_state_root failed for block {block_number}"))?;
        let prev_disk_root =
            query_state_root(prev_block_number, provider).await.with_context(|| {
                format!("query_state_root failed for prev block {prev_block_number}")
            })?;

        let tx_count = block.transactions.len();
        let header_number_u64 = block.header.number.to::<u64>();
        println!("start execute block_{header_number_u64} in host, txns.len: {tx_count}");

        // Init DB (RPC-backed, rooted at previous block).
        let rpc_db = BasicRpcDb::new(provider.clone(), prev_block_number, prev_disk_root.disk_root);

        // Warm up predeployed contract info (improves determinism and reduces latency spikes).
        rpc_db
            .fetch_account_info(WITHDRAW_ROOT_ADDRESS)
            .await
            .context("failed to fetch WITHDRAW_ROOT_ADDRESS account info")?;
        rpc_db
            .fetch_storage_at(WITHDRAW_ROOT_ADDRESS, WITHDRAW_ROOT_SLOT)
            .await
            .context("failed to fetch WITHDRAW_ROOT_ADDRESS storage slot")?;

        let state = State::builder()
            .with_database_ref(&rpc_db)
            .with_bundle_update()
            .without_state_clear()
            .build();

        let basefee_u64 = block.header.base_fee_per_gas.unwrap_or_default().to::<u64>();
        let block_env = BlockEnv {
            number: block.header.number,
            timestamp: block.header.timestamp,
            basefee: basefee_u64,
            gas_limit: block.header.gas_limit.to::<u64>(),
            beneficiary,
            ..Default::default()
        };

        // Build EVM.
        let mut evm = MorphExecutor::with_hardfork(state, block_env, chain_id);

        // Execute transactions in block.
        for (tx_index, tx_trace) in block.transactions.iter().enumerate() {
            let tx = tx_trace
                .try_build_tx_envelope()
                .with_context(|| format!("failed to build tx envelope at index {tx_index}"))?;

            let caller = tx
                .recover_signer()
                .with_context(|| format!("tx[{tx_index}] recover signer error"))?;

            let tx_env = revm::context::TxEnv {
                caller,
                nonce: tx.nonce(),
                gas_price: tx.effective_gas_price(Some(basefee_u64)),
                gas_priority_fee: tx.max_priority_fee_per_gas(),
                gas_limit: tx.gas_limit(),
                kind: tx.kind(),
                value: tx.value(),
                data: revm::primitives::Bytes::from(tx.input().to_vec()),
                chain_id: Some(chain_id),
                ..Default::default()
            };

            let morph_tx =
                MorphTxEnv { inner: tx_env, rlp_bytes: Some(tx.rlp()), ..Default::default() };

            evm.inner
                .transact_commit(morph_tx)
                .with_context(|| format!("tx[{tx_index}] transact_commit error"))?;
        }

        // Merge transitions and build hashed post-state.
        evm.inner.ctx.journaled_state.database.merge_transitions(BundleRetention::Reverts);
        let bundle_state = evm.inner.ctx.journaled_state.database.take_bundle();
        let hashed_post_state =
            HashedPostState::from_bundle_state::<KeccakKeyHasher>(&bundle_state.state);

        // Populate state by fetching missing trie nodes/accounts from provider.
        let state = rpc_db
            .state(&bundle_state)
            .await
            .context("failed to populate post-state from RPC DB")?;

        // Verify post state root.
        let expected_state_root = disk_root.disk_root;
        let mut state_for_verification = state.clone();
        state_for_verification.update(&hashed_post_state);
        let computed_state_root = state_for_verification.state_root();

        if computed_state_root != expected_state_root {
            bail!(
                "Mismatched state root after executing block {block_number}: expected {expected_state_root:?}, got {computed_state_root:?}"
            );
        }

        println!("====success execute block_{header_number_u64} in host====");

        Ok(HostExecutorOutput {
            chain_id,
            beneficiary,
            block,
            state,
            codes: rpc_db.bytecodes(),
            prev_state_root: prev_disk_root.disk_root,
            post_state_root: disk_root.disk_root,
        })
    }
}

#[cfg(test)]
mod tests {
    use crate::{execute::HostExecutor, utils::ProverBlock};
    use alloy_provider::{Provider, ProviderBuilder};

    #[tokio::test(flavor = "multi_thread", worker_threads = 4)]
    async fn test_execute_host() {
        let provider =
            ProviderBuilder::new().connect_http("http://127.0.0.1:9545".parse().unwrap()).erased();

        // let block_number = 53;
        let block_number = 0x477;

        HostExecutor::execute_block(block_number, &provider).await.unwrap();
    }

    #[tokio::test]
    async fn test_prover_block() {
        let provider =
            ProviderBuilder::new().connect_http("http://127.0.0.1:9545".parse().unwrap()).erased();

        // let block_number = 53;
        let block_number = 0x477;

        let result: ProverBlock = provider
            .raw_request("eth_getBlockByNumber".into(), (format!("{block_number:#x}"), true))
            .await
            .unwrap();
        println!("result: {:?}", result);
    }
}
