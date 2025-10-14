use crate::error::ZkTrieError;
use crate::{executor::hooks::ExecuteHooks, EvmExecutor, HardforkConfig, ReadOnlyDB};
use core::fmt;
use revm::db::CacheDB;
use sbv_primitives::{zk_trie::ZkMemoryDb, Block};
use std::rc::Rc;

/// Builder for EVM executor.
pub struct EvmExecutorBuilder<'e, H> {
    hardfork_config: H,
    execute_hooks: ExecuteHooks<'e>,
    zktrie_db: Rc<ZkMemoryDb>,
}

impl fmt::Debug for EvmExecutorBuilder<'_, ()> {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        f.debug_struct("EvmExecutorBuilder")
            .field("hardfork_config", &self.hardfork_config)
            .field("execute_hooks", &self.execute_hooks)
            .field("zktrie_db", &"...")
            .finish()
    }
}

impl<'e> EvmExecutorBuilder<'e, ()> {
    /// Create a new builder.
    pub fn new(zktrie_db: Rc<ZkMemoryDb>) -> Self {
        Self {
            hardfork_config: (),
            execute_hooks: ExecuteHooks::default(),
            zktrie_db,
        }
    }
}

impl<'e, H> EvmExecutorBuilder<'e, H> {
    /// Set hardfork config.
    pub fn hardfork_config<H1>(self, hardfork_config: H1) -> EvmExecutorBuilder<'e, H1> {
        EvmExecutorBuilder {
            hardfork_config,
            execute_hooks: self.execute_hooks,
            zktrie_db: self.zktrie_db,
        }
    }

    /// Modify execute hooks.
    pub fn with_execute_hooks(mut self, modify: impl FnOnce(&mut ExecuteHooks<'e>)) -> Self {
        modify(&mut self.execute_hooks);
        self
    }

    /// Set zktrie state.
    pub fn zktrie_db(self, zktrie_db: Rc<ZkMemoryDb>) -> EvmExecutorBuilder<'e, H> {
        EvmExecutorBuilder { zktrie_db, ..self }
    }
}

impl<'e> EvmExecutorBuilder<'e, HardforkConfig> {
    /// Initialize an EVM executor from a block trace as the initial state.
    pub fn build<T: Block>(self, l2_trace: &T) -> Result<EvmExecutor<'e>, ZkTrieError> {
        let block_number = l2_trace.number();
        let spec_id = self.hardfork_config.get_spec_id(block_number);

        dev_trace!("use spec id {:?}", spec_id);

        let db = cycle_track!(
            CacheDB::new(ReadOnlyDB::new(l2_trace, &self.zktrie_db)?),
            "build ReadOnlyDB"
        );

        Ok(EvmExecutor {
            hardfork_config: self.hardfork_config,
            db,
            spec_id,
            hooks: self.execute_hooks,
        })
    }
}
