use crate::utils::{beneficiary_by_chain_id, query_block, HostExecutorOutput};
use alloy_provider::{DynProvider, Provider};
use anyhow::{bail, Context};
use prover_executor_core::MorphExecutor;
use prover_primitives::{
    predeployed::l2_to_l1_message::{
        SEQUENCER_ROOT_ADDRESS, SEQUENCER_ROOT_SLOT, WITHDRAW_ROOT_ADDRESS, WITHDRAW_ROOT_SLOT,
    },
    TxTrace,
};
use prover_storage_rpc::{
    basic_rpc_db::{BasicRpcDb, RpcDb},
    witness_rpc_db::ExecutionWitnessRpcDb,
};
use reth_trie::{HashedPostState, KeccakKeyHasher};
use revm::{context::BlockEnv, database::State};

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
        // Fetch block.
        let block = query_block(block_number, provider)
            .await
            .with_context(|| format!("query_block failed for block {block_number}"))?;
        // Post-MPT migration (PR #886), the L2 state trie is a standard Ethereum MPT, so
        // `header.state_root` is authoritative and replaces the previous custom `morph_diskRoot`
        // RPC field. The root is re-derived locally below and checked against this value.
        let post_state_root = block.header.state_root;

        // layer2 chain id
        let chain_id =
            provider.get_chain_id().await.context("failed to fetch chain_id from provider")?;

        // beneficiary(coinbase)
        // In Clique consensus, header.Coinbase is always 0x0000...0000.
        // The actual beneficiary is the signer recovered from extraData.
        // We use a per-chain hardcoded address as the sequencer/beneficiary.
        let beneficiary = beneficiary_by_chain_id(chain_id);

        // We need a previous block root to initialize the RPC-backed DB.
        let prev_block_number = block_number
            .checked_sub(1)
            .context("HostExecutor::execute_block requires block_number > 0 (needs prev state)")?;

        let prev_block = query_block(prev_block_number, provider)
            .await
            .with_context(|| format!("query_block failed for prev block {prev_block_number}"))?;
        // Same rationale as `post_state_root`: header state_root is the MPT root.
        let prev_state_root = prev_block.header.state_root;

        let tx_count = block.transactions.len();
        let block_num = block.header.number.to::<u64>();

        // Init DB (RPC-backed, rooted at previous block).
        let rpc_db =
            BasicRpcDb::new(provider.clone(), chain_id, prev_block_number, prev_state_root);

        // Warm up predeployed contract info.
        load_predeployed_contracts(&rpc_db).await?;
        // Build state.
        let state = State::builder().with_database_ref(&rpc_db).with_bundle_update().build();

        let basefee = block.header.base_fee_per_gas.unwrap_or_default().to::<u64>();
        let block_env = BlockEnv {
            number: block.header.number,
            timestamp: block.header.timestamp,
            basefee,
            gas_limit: block.header.gas_limit.to::<u64>(),
            beneficiary,
            ..Default::default()
        };

        let txns = block
            .transactions
            .iter()
            .map(|tx_trace| tx_trace.try_build_tx_envelope())
            .collect::<Result<Vec<_>, _>>()?;

        // Build EVM.
        let mut core_executor = MorphExecutor::with_hardfork(state, block_env, chain_id);
        // Execute block.
        let bundle_state = core_executor
            .execute_block(&txns)
            .with_context(|| format!("failed to execute block {block_number}"))?;

        // Populate state by fetching missing trie nodes/accounts from provider.
        let state = rpc_db.state(&bundle_state).await.context("failed to populate post-state")?;

        // Verify post state root.
        let computed_state_root = {
            let mut state_for_verification = state.clone();
            state_for_verification.update(&HashedPostState::from_bundle_state::<KeccakKeyHasher>(
                &bundle_state.state,
            ));
            state_for_verification.state_root()
        };
        let expected_state_root = block.header.state_root;
        if computed_state_root != expected_state_root {
            bail!(
                "Mismatched state root after executing block {block_number}: expected {expected_state_root:?}, got {computed_state_root:?}"
            );
        }

        log::debug!("success execute block_{block_num} in host, txns.len: {tx_count}");
        Ok(HostExecutorOutput {
            chain_id,
            beneficiary,
            block,
            state,
            codes: rpc_db.bytecodes(),
            prev_state_root,
            post_state_root,
        })
    }

    /// Executes `block_number` using a pre-fetched `debug_executionWitness` instead of
    /// per-account `eth_getProof` calls.
    ///
    /// This is significantly more efficient than [`execute_block`] because it fetches all
    /// required state data in a single RPC call rather than one call per account.
    ///
    /// # Flow
    /// 1. Fetch block header + transactions via `eth_getBlockByNumber`.
    /// 2. Call `debug_executionWitness` for the block to obtain the pre-state trie nodes,
    ///    bytecodes, and ancestor headers.
    /// 3. Execute all transactions against the witness-backed DB.
    /// 4. Verify the computed post-state root matches the block header.
    pub async fn execute_block_with_witness(
        block_number: u64,
        provider: &DynProvider,
    ) -> Result<HostExecutorOutput, anyhow::Error> {
        // Fetch block.
        let block = query_block(block_number, provider)
            .await
            .with_context(|| format!("query_block failed for block {block_number}"))?;
        let post_state_root = block.header.state_root;

        // Fetch chain id and beneficiary.
        let chain_id =
            provider.get_chain_id().await.context("failed to fetch chain_id from provider")?;
        let beneficiary = beneficiary_by_chain_id(chain_id);

        // Fetch previous block to obtain prev_state_root.
        let prev_block_number = block_number
            .checked_sub(1)
            .context("execute_block_with_witness requires block_number > 0")?;
        let prev_block = query_block(prev_block_number, provider)
            .await
            .with_context(|| format!("query_block failed for prev block {prev_block_number}"))?;
        let prev_state_root = prev_block.header.state_root;

        let tx_count = block.transactions.len();
        let block_num = block.header.number.to::<u64>();

        // Build the witness-backed DB.  This issues a single `debug_executionWitness` RPC call
        // and pre-populates the entire pre-state trie in memory.
        let witness_db =
            ExecutionWitnessRpcDb::new(provider.clone(), chain_id, prev_block_number, prev_state_root)
                .await
                .context("failed to build ExecutionWitnessRpcDb")?;

        // Build state on top of the witness DB.
        let state = State::builder().with_database_ref(&witness_db).with_bundle_update().build();

        let basefee = block.header.base_fee_per_gas.unwrap_or_default().to::<u64>();
        let block_env = BlockEnv {
            number: block.header.number,
            timestamp: block.header.timestamp,
            basefee,
            gas_limit: block.header.gas_limit.to::<u64>(),
            beneficiary,
            ..Default::default()
        };

        let txns = block
            .transactions
            .iter()
            .map(|tx_trace| tx_trace.try_build_tx_envelope())
            .collect::<Result<Vec<_>, _>>()?;

        // Execute block.
        let mut core_executor = MorphExecutor::with_hardfork(state, block_env, chain_id);
        let bundle_state = core_executor
            .execute_block(&txns)
            .with_context(|| format!("failed to execute block {block_number}"))?;

        // Verify post state root by applying the bundle diff to the pre-state.
        let computed_state_root = {
            let mut state_for_verification = witness_db.state.clone();
            state_for_verification.update(&HashedPostState::from_bundle_state::<KeccakKeyHasher>(
                &bundle_state.state,
            ));
            state_for_verification.state_root()
        };
        if computed_state_root != post_state_root {
            bail!(
                "Mismatched state root after executing block {block_number} (witness): \
                 expected {post_state_root:?}, got {computed_state_root:?}"
            );
        }

        log::debug!(
            "success execute block_{block_num} in host (witness), txns.len: {tx_count}"
        );

        // Return the pre-state (parent_state) so the client can re-execute and verify.
        // Extract codes before consuming witness_db.state.
        let codes = witness_db.bytecodes();
        Ok(HostExecutorOutput {
            chain_id,
            beneficiary,
            block,
            state: witness_db.state,
            codes,
            prev_state_root,
            post_state_root,
        })
    }
}

async fn load_predeployed_contracts(
    rpc_db: &BasicRpcDb<DynProvider, alloy_network::Ethereum>,
) -> Result<(), anyhow::Error> {
    rpc_db
        .fetch_account_info(WITHDRAW_ROOT_ADDRESS)
        .await
        .context("failed to fetch WITHDRAW_ROOT_ADDRESS account info")?;
    rpc_db
        .fetch_storage_at(WITHDRAW_ROOT_ADDRESS, WITHDRAW_ROOT_SLOT)
        .await
        .context("failed to fetch WITHDRAW_ROOT_ADDRESS storage slot")?;
    rpc_db
        .fetch_account_info(SEQUENCER_ROOT_ADDRESS)
        .await
        .context("failed to fetch SEQUENCER_ROOT_ADDRESS account info")?;
    rpc_db
        .fetch_storage_at(SEQUENCER_ROOT_ADDRESS, SEQUENCER_ROOT_SLOT)
        .await
        .context("failed to fetch SEQUENCER_ROOT_ADDRESS storage slot")?;
    Ok(())
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
        println!("result: {result:?}");
    }
}
