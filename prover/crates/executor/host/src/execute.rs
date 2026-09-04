use crate::utils::{beneficiary_by_chain_id, query_morph_rpc_block, HostExecutorOutput};
use alloy_primitives::{Address, B256, U256};
use alloy_provider::{DynProvider, Provider};
use anyhow::{bail, Context};
use morph_primitives::{MorphHeader, MorphTxEnvelope};
use prover_executor_core::MorphExecutor;
use prover_mpt::EthereumState;
use prover_primitives::{
    alloy_consensus::BlockHeader,
    predeployed::l2_to_l1_message::{WITHDRAW_ROOT_ADDRESS, WITHDRAW_ROOT_SLOT},
};
use prover_storage_rpc::basic_rpc_db::{BasicRpcDb, RpcDb};
use prover_storage_rpc::witness_rpc_db::ExecutionWitnessRpcDb;
use reth_trie::{HashedPostState, KeccakKeyHasher};
use revm::database::BundleState;

/// The consensus block type both execution paths operate on.
type MorphConsensusBlock = alloy_consensus::Block<MorphTxEnvelope, MorphHeader>;

/// Predeployed contract (address, slot) pairs that must always be present in the state,
/// regardless of whether the block touches them.
const PREDEPLOYED_FORCE_INCLUDE: [(Address, U256); 1] =
    [(WITHDRAW_ROOT_ADDRESS, WITHDRAW_ROOT_SLOT)];

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
        let ctx =
            ExecutionContext::fetch(block_number, provider, "HostExecutor::execute_block").await?;

        // Init DB (RPC-backed, rooted at previous block).
        let rpc_db = BasicRpcDb::new(
            provider.clone(),
            ctx.chain_id,
            ctx.prev_block_number,
            ctx.prev_state_root,
        );

        // Warm up predeployed contract info.
        load_predeployed_contracts(&rpc_db).await?;

        // Execute the whole block via reth's BasicBlockExecutor.
        let core_executor = MorphExecutor::new_ref(&rpc_db, ctx.chain_id);
        let bundle_state = core_executor
            .execute_block(ctx.morph_block.clone())
            .with_context(|| format!("failed to execute block {block_number}"))?;

        // Populate state by fetching missing trie nodes/accounts from provider.
        let state = rpc_db.state(&bundle_state).await.context("failed to populate post-state")?;

        verify_post_state_root(&state, &bundle_state, ctx.post_state_root, block_number, "")?;

        log::debug!("success execute block_{} in host, txns.len: {}", ctx.block_num, ctx.tx_count);

        Ok(HostExecutorOutput {
            chain_id: ctx.chain_id,
            beneficiary: ctx.beneficiary,
            block: ctx.morph_block,
            state,
            codes: rpc_db.bytecodes(),
            prev_state_root: ctx.prev_state_root,
            post_state_root: ctx.post_state_root,
        })
    }

    /// Executes `block_number` using a pre-fetched `debug_executionWitness` instead of
    /// per-account `eth_getProof` calls.
    ///
    /// This is significantly more efficient than [`execute_block`] because it fetches all
    /// required state data in a single RPC call rather than one call per account.
    pub async fn execute_block_with_witness(
        block_number: u64,
        provider: &DynProvider,
    ) -> Result<HostExecutorOutput, anyhow::Error> {
        let ctx =
            ExecutionContext::fetch(block_number, provider, "execute_block_with_witness").await?;

        // Build the witness-backed DB.  This issues a single `debug_executionWitness` RPC call
        // and pre-populates the entire pre-state trie in memory.
        // Force-include the Morph predeploys whose storage the client reads to derive batch
        // public inputs (withdraw root). A block that does not touch it would
        // otherwise leave their storage tries out of the execution witness.
        let witness_db = ExecutionWitnessRpcDb::new(
            provider.clone(),
            ctx.chain_id,
            ctx.prev_block_number,
            ctx.prev_state_root,
            &PREDEPLOYED_FORCE_INCLUDE,
        )
        .await
        .context("failed to build ExecutionWitnessRpcDb")?;

        // Execute the whole block via reth's BasicBlockExecutor.
        let core_executor = MorphExecutor::new_ref(&witness_db, ctx.chain_id);
        let bundle_state = core_executor
            .execute_block(ctx.morph_block.clone())
            .with_context(|| format!("failed to execute block {block_number}"))?;

        // Verify post state root by applying the bundle diff to the pre-state.
        verify_post_state_root(
            &witness_db.state,
            &bundle_state,
            ctx.post_state_root,
            block_number,
            " (witness)",
        )?;

        log::debug!(
            "success execute block_{} in host (witness), txns.len: {}",
            ctx.block_num,
            ctx.tx_count
        );

        // Return the pre-state (parent_state) so the client can re-execute and verify.
        // Extract codes before consuming witness_db.state.
        let codes = witness_db.bytecodes();
        Ok(HostExecutorOutput {
            chain_id: ctx.chain_id,
            beneficiary: ctx.beneficiary,
            block: ctx.morph_block,
            state: witness_db.state,
            codes,
            prev_state_root: ctx.prev_state_root,
            post_state_root: ctx.post_state_root,
        })
    }
}

/// Provider-fetched data shared by both execution paths before a block is run.
struct ExecutionContext {
    /// The block, converted to the consensus representation the executor operates on.
    morph_block: MorphConsensusBlock,
    /// Layer-2 chain id.
    chain_id: u64,
    /// Sequencer/beneficiary (coinbase) resolved per chain — see [`beneficiary_by_chain_id`].
    beneficiary: Address,
    /// Number of the parent block, used to root the state-backing DB.
    prev_block_number: u64,
    /// State root of the parent block (the MPT root the DB is initialized against).
    prev_state_root: B256,
    /// State root expected after executing the block (from the block header).
    post_state_root: B256,
    tx_count: usize,
    block_num: u64,
}

impl ExecutionContext {
    /// Fetches the block and its parent, resolves chain metadata, and converts the block into
    /// the consensus representation both execution paths consume.
    ///
    /// `ctx_label` names the caller so the "requires block_number > 0" error is attributable.
    async fn fetch(
        block_number: u64,
        provider: &DynProvider,
        ctx_label: &str,
    ) -> Result<Self, anyhow::Error> {
        // Fetch block.
        let block = query_morph_rpc_block(block_number, provider)
            .await
            .with_context(|| format!("query_block failed for block {block_number}"))?;
        let post_state_root = block.header.state_root();

        // layer2 chain id
        let chain_id =
            provider.get_chain_id().await.context("failed to fetch chain_id from provider")?;

        // beneficiary(coinbase)
        // In Clique consensus, header.Coinbase is always 0x0000...0000.
        // The actual beneficiary is the signer recovered from extraData.
        // We use a per-chain hardcoded address as the sequencer/beneficiary.
        let beneficiary = beneficiary_by_chain_id(chain_id);

        // We need a previous block root to initialize the state-backing DB.
        let prev_block_number = block_number
            .checked_sub(1)
            .with_context(|| format!("{ctx_label} requires block_number > 0 (needs prev state)"))?;
        let prev_block = query_morph_rpc_block(prev_block_number, provider)
            .await
            .with_context(|| format!("query_block failed for prev block {prev_block_number}"))?;
        // Same rationale as `post_state_root`: header state_root is the MPT root.
        let prev_state_root = prev_block.header.state_root();

        let tx_count = block.transactions.len();
        let block_num = block.header.number();

        let mut morph_block = block
            .into_consensus_block()
            .map_header(|header| header.into_consensus())
            .map_transactions(|tx| tx.into_inner());
        // Morph's Clique-style RPC header may expose a zero coinbase. Execution must use the
        // configured sequencer address, matching the pre-reth executor path on main.
        morph_block.header.inner.beneficiary = beneficiary;

        Ok(Self {
            morph_block,
            chain_id,
            beneficiary,
            prev_block_number,
            prev_state_root,
            post_state_root,
            tx_count,
            block_num,
        })
    }
}

/// Applies the executed `bundle_state` diff to `state` and bails if the resulting root does not
/// match `expected_state_root`. `label` distinguishes execution paths in the error message.
fn verify_post_state_root(
    state: &EthereumState,
    bundle_state: &BundleState,
    expected_state_root: B256,
    block_number: u64,
    label: &str,
) -> Result<(), anyhow::Error> {
    let mut state_for_verification = state.clone();
    state_for_verification
        .update(&HashedPostState::from_bundle_state::<KeccakKeyHasher>(&bundle_state.state));
    let computed_state_root = state_for_verification.state_root();
    if computed_state_root != expected_state_root {
        bail!(
            "Mismatched state root after executing block {block_number}{label}: \
             expected {expected_state_root:?}, got {computed_state_root:?}"
        );
    }
    Ok(())
}

async fn load_predeployed_contracts(
    rpc_db: &BasicRpcDb<DynProvider, alloy_network::Ethereum>,
) -> Result<(), anyhow::Error> {
    for (address, slot) in PREDEPLOYED_FORCE_INCLUDE {
        rpc_db
            .fetch_account_info(address)
            .await
            .with_context(|| format!("failed to fetch account info for {address}"))?;
        rpc_db
            .fetch_storage_at(address, slot)
            .await
            .with_context(|| format!("failed to fetch storage at {address} slot {slot}"))?;
    }
    Ok(())
}
