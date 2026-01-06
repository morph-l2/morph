use std::ops::Sub;

use alloy_consensus::{transaction::SignerRecoverable, Transaction};
use alloy_provider::{DynProvider, Provider};
use anyhow::{bail, Context};
use morph_revm::MorphTxEnv;
use prover_executor_core::MorphExecutor;
use prover_primitives::{
    predeployed::l2_to_l1_message::{WITHDRAW_ROOT_ADDRESS, WITHDRAW_ROOT_SLOT},
    MorphTxEnvelope, TxTrace,
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
    pub async fn execute_block(
        block_number: u64,
        provider: &DynProvider,
    ) -> Result<HostExecutorOutput, anyhow::Error> {
        let block = query_block(block_number, provider).await?;
        let chain_id = provider.get_chain_id().await?;
        let beneficiary = *CHAIN_CONFIG
            .get(&chain_id)
            .ok_or(anyhow::anyhow!("Chain ID not found in CHAIN_CONFIG"))?;

        let disk_root = query_state_root(block_number, provider).await?;
        let prev_disk_root = query_state_root(block_number.sub(1), provider).await?;

        let evm_txns: Vec<MorphTxEnvelope> = block
            .transactions
            .iter()
            .map(|tx_trace| tx_trace.try_build_tx_envelope().unwrap())
            .collect();
        println!(
            "start execute block_{:?} in host, txns.len: {:?}",
            block.header.number.to::<u64>(),
            evm_txns.len()
        );

        // Init db.
        let rpc_db =
            BasicRpcDb::new(provider.clone(), block_number.sub(1), prev_disk_root.disk_root);
        // Fetch predeployed contract info.
        let _ = rpc_db.fetch_account_info(WITHDRAW_ROOT_ADDRESS).await?;
        let _ = rpc_db.fetch_storage_at(WITHDRAW_ROOT_ADDRESS, WITHDRAW_ROOT_SLOT).await?;
        let state = State::builder()
            .with_database_ref(&rpc_db)
            .with_bundle_update()
            .without_state_clear()
            .build();

        let block_env = BlockEnv {
            number: block.header.number,
            timestamp: block.header.timestamp,
            basefee: block.header.base_fee_per_gas.unwrap_or_default().to::<u64>(),
            gas_limit: block.header.gas_limit.to::<u64>(),
            beneficiary,
            ..Default::default()
        };
        // Build EVM.
        let mut evm = MorphExecutor::with_hardfork(state, block_env.clone());
        // Execute transactions in block.
        for tx in evm_txns {
            let recovered_from = tx.recover_signer().context("tx recover signer error")?;
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
            let _rt = evm.inner.transact_commit(morph_tx).context("tx transact error")?;
        }
        evm.inner.ctx.journaled_state.database.merge_transitions(BundleRetention::Reverts);
        let bundle_state = evm.inner.ctx.journaled_state.database.take_bundle();

        let state = rpc_db.state(&bundle_state).await?;

        let hashed_post_state =
            HashedPostState::from_bundle_state::<KeccakKeyHasher>(&bundle_state.state);

        let mut mut_state = state.clone();
        // Verify post state root.
        mut_state.update(&hashed_post_state);

        let computed_state_root = mut_state.state_root();
        let expected_state_root = disk_root.disk_root;
        if computed_state_root != expected_state_root {
            bail!(
                "Mismatched state root after executing the block, block number: {block_number}, expected: {expected_state_root:?}, got: {computed_state_root:?}"
            );
        }

        println!("====success execute block_{:?} in host====", block.header.number.to::<u64>());

        let out_put = HostExecutorOutput {
            chain_id,
            beneficiary,
            block,
            state,
            codes: rpc_db.bytecodes(),
            prev_state_root: prev_disk_root.disk_root,
            post_state_root: disk_root.disk_root,
        };
        Ok(out_put)
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
