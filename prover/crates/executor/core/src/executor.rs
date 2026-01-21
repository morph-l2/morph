use alloy_consensus::transaction::SignerRecoverable;
use alloy_consensus::Transaction;
use alloy_evm::{revm::Context as EvmContext, Database, EvmEnv};
use anyhow::Context;
use anyhow::Result;
use morph_chainspec::hardfork::MorphHardfork;
use morph_chainspec::hardfork::MorphHardforks;
use morph_chainspec::MORPH_MAINNET;
use morph_evm::MorphBlockEnv;
use morph_primitives::MorphTxEnvelope;
use morph_revm::MorphEvm;

use morph_revm::MorphTxEnv;
use revm::context::BlockEnv;
use revm::database::BundleState;
use revm::inspector::NoOpInspector;
use revm::MainContext;
use revm::{
    database::{states::bundle_state::BundleRetention, State},
    ExecuteCommitEvm,
};
/// An Morph executor wrapper based on `revm`.
pub struct MorphExecutor<DB: Database, I = NoOpInspector> {
    pub inner: morph_revm::MorphEvm<State<DB>, I>,
}

impl<DB: Database> MorphExecutor<DB> {
    /// Create a new [`MorphEvm`] instance.
    pub fn new(db: State<DB>, input: EvmEnv<MorphHardfork, MorphBlockEnv>) -> Self {
        let ctx = EvmContext::mainnet()
            .with_db(db)
            .with_block(input.block_env)
            .with_cfg(input.cfg_env)
            .with_tx(Default::default());

        let evm = MorphEvm::new(ctx, NoOpInspector {});
        Self { inner: evm }
    }
    pub fn with_hardfork(db: State<DB>, block_env: BlockEnv, chain_id: u64) -> Self {
        let mut env: EvmEnv<MorphHardfork, MorphBlockEnv> =
            EvmEnv::default().with_timestamp(block_env.timestamp);

        let spec = &MORPH_MAINNET;
        let hardfork =
            spec.morph_hardfork_at(block_env.number.to::<u64>(), block_env.timestamp.to::<u64>());
        env.cfg_env = env.cfg_env.with_spec(hardfork);
        env.cfg_env.chain_id = chain_id;
        env.cfg_env.tx_gas_limit_cap = Some(block_env.gas_limit);
        env.cfg_env.disable_eip7623 = true;
        env.block_env = MorphBlockEnv { inner: block_env };
        Self::new(db, env)
    }

    pub fn execute_block(&mut self, txns: &Vec<MorphTxEnvelope>) -> Result<BundleState> {
        // Execute transactions in block.
        for (tx_index, tx) in txns.iter().enumerate() {
            let basefee = self.inner.ctx.block.basefee;
            let caller = SignerRecoverable::recover_signer(tx)?;

            let mut morph_tx = MorphTxEnv::from_recovered_tx(tx, caller);
            morph_tx.gas_price = tx.effective_gas_price(Some(basefee));

            self.inner
                .transact_commit(morph_tx)
                .with_context(|| format!("tx[{tx_index}] transact_commit error"))?;
        }
        // Merge transitions and build hashed post-state.
        self.inner.ctx.journaled_state.database.merge_transitions(BundleRetention::Reverts);
        // Collect values that got changed.
        let mut bundle_state = self.inner.ctx.journaled_state.database.take_bundle();
        // If current account_info is not exists and original_info is None,
        // It means that it has not been changed, account can be filtered out
        bundle_state.state.retain(|_, acc| {
            let exists = acc.info.as_ref().map(|info| info.exists()).unwrap_or(false);
            exists || acc.original_info.is_some()
        });
        Ok(bundle_state)
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use morph_revm::MorphTxEnv;
    use revm::context::TxEnv;
    use revm::database::{CacheDB, EmptyDB, State};
    use revm::primitives::{Address, U256};
    use revm::ExecuteEvm;

    #[test]
    fn test_main_context() {
        let state: State<CacheDB<EmptyDB>> =
            State::builder().with_database(CacheDB::default()).build();

        let mut env = EvmEnv::default().with_timestamp(U256::ZERO);
        env.cfg_env = env.cfg_env.with_spec(MorphHardfork::Viridian);

        let mut evm = MorphExecutor::new(state, env);
        let tx = TxEnv {
            nonce: 0,
            gas_price: 1_000_000_000,
            gas_limit: 21_000,
            kind: revm::primitives::TxKind::Call(Address::default()),
            value: U256::from(1_000_000u64),
            data: revm::primitives::Bytes::new(),
            ..Default::default()
        };
        let morph_tx =
            MorphTxEnv { inner: tx, rlp_bytes: Default::default(), ..Default::default() };
        let _rt = evm.inner.transact_one(morph_tx);
        let _state = evm.inner.finalize();
        let _db = evm.inner.journaled_state.database.take_bundle();
    }
}
