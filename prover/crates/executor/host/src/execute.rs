use crate::utils::{beneficiary_by_chain_id, query_block, query_state_root, HostExecutorOutput};
use alloy_provider::{DynProvider, Provider};
use anyhow::{bail, Context};
use prover_executor_core::MorphExecutor;
use prover_primitives::{
    predeployed::l2_to_l1_message::{
        SEQUENCER_ROOT_ADDRESS, SEQUENCER_ROOT_SLOT, WITHDRAW_ROOT_ADDRESS, WITHDRAW_ROOT_SLOT,
    },
    TxTrace,
};
use prover_storage_rpc::basic_rpc_db::{BasicRpcDb, RpcDb};
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

        // layer2 chain id
        let chain_id =
            provider.get_chain_id().await.context("failed to fetch chain_id from provider")?;

        // beneficiary(coinbase)
        let beneficiary = beneficiary_by_chain_id(chain_id);

        // mpt root at this block
        let disk_root = query_state_root(block_number, provider)
            .await
            .with_context(|| format!("query_state_root failed for block {block_number}"))?;

        // We need a previous block root to initialize the RPC-backed DB.
        let prev_block_number = block_number
            .checked_sub(1)
            .context("HostExecutor::execute_block requires block_number > 0 (needs prev state)")?;
        let prev_disk_root =
            query_state_root(prev_block_number, provider).await.with_context(|| {
                format!("query_state_root failed for prev block {prev_block_number}")
            })?;

        let tx_count = block.transactions.len();
        let block_num = block.header.number.to::<u64>();

        // Init DB (RPC-backed, rooted at previous block).
        let rpc_db = BasicRpcDb::new(
            provider.clone(),
            chain_id,
            prev_block_number,
            prev_disk_root.disk_root,
        );

        // Warm up predeployed contract info.
        load_predeployed_contracts(&rpc_db).await?;
        // Build state.
        let state = State::builder()
            .with_database_ref(&rpc_db)
            .with_bundle_update()
            .without_state_clear()
            .build();

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
        let expected_state_root = disk_root.disk_root;
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
            prev_state_root: prev_disk_root.disk_root,
            post_state_root: disk_root.disk_root,
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
