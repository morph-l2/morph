use crate::EvmExecutor;
use std::fmt::{Debug, Formatter};

/// Transaction RLP handler.
pub type TxRLPHandler<'a, CodeDb, ZkDb> = dyn Fn(&EvmExecutor<CodeDb, ZkDb>, &[u8]) + 'a;
/// Post transaction execution handler.
pub type PostTxExecutionHandler<'a, CodeDb, ZkDb> = dyn Fn(&EvmExecutor<CodeDb, ZkDb>, usize) + 'a;

/// Hooks for the EVM executor.
pub struct ExecuteHooks<'a, CodeDb, ZkDb> {
    tx_rlp_handlers: Vec<Box<TxRLPHandler<'a, CodeDb, ZkDb>>>,
    post_tx_execution_handlers: Vec<Box<PostTxExecutionHandler<'a, CodeDb, ZkDb>>>,
}

impl<'a, CodeDb, ZkDb> Default for ExecuteHooks<'a, CodeDb, ZkDb> {
    fn default() -> Self {
        Self {
            tx_rlp_handlers: Vec::new(),
            post_tx_execution_handlers: Vec::new(),
        }
    }
}

impl<'a, CodeDb, ZkDb> ExecuteHooks<'a, CodeDb, ZkDb> {
    /// Create a new hooks.
    pub fn new() -> Self {
        Self::default()
    }

    /// Add a transaction RLP handler.
    pub fn add_tx_rlp_handler<F>(&mut self, handler: F)
    where
        F: Fn(&EvmExecutor<CodeDb, ZkDb>, &[u8]) + 'a,
    {
        self.tx_rlp_handlers.push(Box::new(handler));
    }

    /// Add a post transaction execution handler.
    pub fn add_post_tx_execution_handler<F>(&mut self, handler: F)
    where
        F: Fn(&EvmExecutor<CodeDb, ZkDb>, usize) + 'a,
    {
        self.post_tx_execution_handlers.push(Box::new(handler));
    }

    /// Execute transaction RLP handlers.
    pub(crate) fn tx_rlp(&self, executor: &EvmExecutor<CodeDb, ZkDb>, rlp: &[u8]) {
        for handler in &self.tx_rlp_handlers {
            handler(executor, rlp);
        }
    }

    pub(crate) fn post_tx_execution(&self, executor: &EvmExecutor<CodeDb, ZkDb>, tx_index: usize) {
        for handler in &self.post_tx_execution_handlers {
            handler(executor, tx_index);
        }
    }
}

impl<CodeDb, ZkDb> Debug for ExecuteHooks<'_, CodeDb, ZkDb> {
    fn fmt(&self, f: &mut Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("ExecuteHooks")
            .field("tx_rlp_handlers", &self.tx_rlp_handlers.len())
            .field(
                "post_tx_execution_handlers",
                &self.post_tx_execution_handlers.len(),
            )
            .finish()
    }
}
