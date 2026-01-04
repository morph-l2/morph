use alloy_evm::{revm::Context, Database, EvmEnv};
use morph_chainspec::hardfork::MorphHardfork;
use morph_evm::MorphBlockEnv;
use morph_revm::MorphEvm;

use revm::context::BlockEnv;
use revm::inspector::NoOpInspector;
use revm::MainContext;

/// An Morph executor wrapper based on `revm`.

pub struct MorphExecutor<DB: Database, I = NoOpInspector> {
    pub inner: morph_revm::MorphEvm<DB, I>,
}

impl<DB: Database> MorphExecutor<DB> {
    /// Create a new [`MorphEvm`] instance.
    pub fn new(db: DB, input: EvmEnv<MorphHardfork, MorphBlockEnv>) -> Self {
        let ctx = Context::mainnet()
            .with_db(db)
            .with_block(input.block_env)
            .with_cfg(input.cfg_env)
            .with_tx(Default::default());

        let evm = MorphEvm::new(ctx, NoOpInspector {});
        Self { inner: evm }
    }
    pub fn with_hardfork(db: DB, block_env: BlockEnv) -> Self {
        let mut env: EvmEnv<MorphHardfork, MorphBlockEnv> =
            EvmEnv::default().with_timestamp(block_env.timestamp);
        env.cfg_env = env.cfg_env.with_spec(MorphHardfork::Curie);
        env.block_env = MorphBlockEnv { inner: block_env };
        Self::new(db, env)
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use morph_revm::MorphTxEnv;
    use revm::context::TxEnv;
    use revm::database::State;
    use revm::primitives::{Address, U256};
    use revm::ExecuteEvm;

    #[test]
    fn test_main_context() {
        let state = State::builder().build();

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
