use alloy_consensus::{transaction::SignerRecoverable, BlockHeader, Transaction};
use alloy_network::{BlockResponse, Ethereum, Network};
use alloy_primitives::map::HashMap;
use alloy_provider::{DynProvider, Provider, ProviderBuilder};
use anyhow::{bail, Context};
use morph_revm::MorphTxEnv;
use prover_executor_core::MorphExecutor;
use prover_mpt::EthereumState;
use prover_primitives::{types::BlockTrace, Block, MorphTxEnvelope, B256};
use prover_storage::{
    basic_rpc_db::{BasicRpcDb, RpcDb},
    trace_to_execution_witness, TrieDB,
};
use prover_utils::provider::{get_block_trace, get_block_traces};
use reth_trie::{HashedPostState, KeccakKeyHasher};
use revm::{
    context::BlockEnv,
    database::{states::bundle_state::BundleRetention, BundleState, State},
    state::Bytecode,
    ExecuteCommitEvm, ExecuteEvm,
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

        // let (trie_db, mut orgin_state) = build_trie_db(block_number, provider).await.unwrap();
        let state = State::builder()
            .with_database_ref(&rpc_db)
            .with_bundle_update()
            .without_state_clear()
            .build();
        let block_env = BlockEnv::default();

        // Build EVM.
        let mut evm = MorphExecutor::with_hardfork(state, block_env);
        // Execute transactions in block.
        for tx in evm_txns {
            let recovered_from = tx.recover_signer().context("tx recover signer error")?;
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
            let morph_tx =
                MorphTxEnv { inner: tx_env, rlp_bytes: Some(tx.rlp()), ..Default::default() };

            let _rt = evm.inner.transact_commit(morph_tx).context("tx transact error")?;
        }
        evm.inner.ctx.journaled_state.database.merge_transitions(BundleRetention::Reverts);
        let bundle_state = evm.inner.ctx.journaled_state.database.take_bundle();
        let mut state = rpc_db.state(&bundle_state).await?;
        // Verify post state root.
        let hashed_post_state =
            HashedPostState::from_bundle_state::<KeccakKeyHasher>(&bundle_state.state);
        state.update(&hashed_post_state);

        let computed_state_root = state.state_root();
        let expected_state_root = block.header().state_root();
        if computed_state_root != expected_state_root {
            bail!(
                "Mismatched state root after executing the block, block number: {block_number}, expected: {expected_state_root:?}, got: {computed_state_root:?}"
            );
        }
        Ok(())
    }
}

// async fn build_trie_db<P, N>(
//     block_number: u64,
//     provider: &P,
// ) -> Result<(TrieDB<'_>, EthereumState), String>
// where
//     P: Provider<N> + Clone + std::fmt::Debug,
//     N: Network,
// {
//     let trace = get_block_trace::<BlockTrace, P, N>(block_number, &provider).await.unwrap();
//     let witness = trace_to_execution_witness(&trace).unwrap();
//     let state = EthereumState::from_execution_witness(&witness, trace.root_before());
//     let bytecodes =
//         witness.codes.into_iter().map(|code| Bytecode::new_raw(code)).collect::<Vec<_>>();
//     let bytecodes_by_hash =
//         bytecodes.iter().map(|code| (code.hash_slow(), code)).collect::<HashMap<_, _>>();
//     let block_hashes: HashMap<u64, B256> = HashMap::with_hasher(Default::default());
//     let statea= state.clone();
//     Ok((TrieDB::new(&state.clone(), block_hashes, bytecodes_by_hash), statea))
// }

fn _to_prover_tx(_txns: &Vec<alloy_rpc_types::Transaction>) -> Vec<MorphTxEnvelope> {
    vec![]
}

#[test]
fn test_execute_host() {
    let provider =
        ProviderBuilder::new().connect_http("http://127.0.0.1:8545".parse().unwrap()).erased();
}
