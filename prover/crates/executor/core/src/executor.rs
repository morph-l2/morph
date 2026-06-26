use anyhow::Result;
use morph_chainspec::{hardfork::MorphHardfork, MorphChainSpec, MORPH_HOODI, MORPH_MAINNET};
use morph_evm::MorphEvmConfig;
use morph_primitives::{Block as MorphBlock, MorphReceipt};
use reth_evm::execute::{BasicBlockExecutor, Executor};
use reth_execution_types::BlockExecutionResult;
use reth_primitives_traits::RecoveredBlock;
use revm::database::{BundleState, State, WrapDatabaseRef};
use revm::Database;

use std::sync::Arc;

use crate::{DEVNET_CHAIN_ID, MAINNET_CHAIN_ID, TESTNET_CHAIN_ID};

/// An Morph executor wrapper based on reth's block executor.
pub struct MorphExecutor<DB: Database> {
    inner: BasicBlockExecutor<MorphEvmConfig, DB>,
}

impl<DB> MorphExecutor<DB>
where
    DB: Database + std::fmt::Debug,
{
    /// Create a new [`MorphExecutor`] instance.
    pub fn new(db: DB, chain_id: u64) -> Self {
        let evm_config = MorphEvmConfig::new_with_default_factory(chain_spec_by_chain_id(chain_id));
        let inner = BasicBlockExecutor::new(evm_config, db);
        Self { inner }
    }

    /// Execute a full recovered block via [`BasicBlockExecutor::execute`].
    pub fn execute_recovered_block(
        self,
        block: &RecoveredBlock<MorphBlock>,
    ) -> Result<(BundleState, BlockExecutionResult<MorphReceipt>)> {
        let execution_output = self.inner.execute(block)?;
        let mut bundle_state = execution_output.state;
        retain_changed_accounts(&mut bundle_state);
        Ok((bundle_state, execution_output.result))
    }

    /// Execute a full block by first recovering transaction senders.
    pub fn execute_block(self, block: MorphBlock) -> Result<BundleState> {
        let recovered_block = RecoveredBlock::try_recover(block)?;
        let (bundle_state, _) = self.execute_recovered_block(&recovered_block)?;
        Ok(bundle_state)
    }
}

impl<DB> MorphExecutor<State<DB>>
where
    DB: Database + std::fmt::Debug,
{
    /// Compatibility constructor for callers that still build a [`State`] manually.
    pub fn with_hardfork(
        db: State<DB>,
        _block_env: revm::context::BlockEnv,
        chain_id: u64,
    ) -> Self {
        Self::new(db, chain_id)
    }
}

impl<DB> MorphExecutor<WrapDatabaseRef<DB>>
where
    DB: revm::DatabaseRef + std::fmt::Debug,
{
    /// Create a new executor from a read-only database reference.
    pub fn new_ref(db: DB, chain_id: u64) -> Self {
        Self::new(WrapDatabaseRef(db), chain_id)
    }
}

fn chain_spec_by_chain_id(chain_id: u64) -> Arc<MorphChainSpec> {
    match chain_id {
        MAINNET_CHAIN_ID => MORPH_MAINNET.clone(),
        TESTNET_CHAIN_ID => MORPH_HOODI.clone(),
        DEVNET_CHAIN_ID => devnet_chain_spec(chain_id),
        _ => devnet_chain_spec(chain_id),
    }
}

fn devnet_chain_spec(chain_id: u64) -> Arc<MorphChainSpec> {
    let mut spec = MORPH_HOODI.clone();
    let spec_mut = Arc::make_mut(&mut spec);
    spec_mut.inner.chain = chain_id.into();
    spec_mut.info.morph_chain_info.fee_vault_address = None;
    spec_mut.set_hardfork(MorphHardfork::Morph203, 0);
    spec_mut.set_hardfork(MorphHardfork::Viridian, 0);
    spec_mut.set_hardfork(MorphHardfork::Emerald, 0);
    spec
}

fn retain_changed_accounts(bundle_state: &mut BundleState) {
    // If current account_info does not exist and original_info is None, it means that it has not
    // been changed and the account can be filtered out.
    bundle_state.state.retain(|_, acc| {
        let exists = acc.info.as_ref().map(|info| info.exists()).unwrap_or(false);
        exists || acc.original_info.is_some()
    });
}

#[cfg(test)]
mod tests {
    use super::*;
    use revm::database::{CacheDB, EmptyDB};

    #[test]
    fn test_create_executor() {
        let _evm = MorphExecutor::new(CacheDB::<EmptyDB>::default(), TESTNET_CHAIN_ID);
    }
}
