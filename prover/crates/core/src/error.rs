use revm::primitives::alloy_primitives::SignatureError;
use revm::primitives::{EVMError, B256};
use std::error::Error;

/// Error variants encountered during manipulation of a zkTrie.
#[derive(Debug, thiserror::Error)]
pub enum DatabaseError {
    #[error("error encountered from code db: {0}")]
    CodeDb(Box<dyn Error + Send + Sync>),
    #[error("error encountered from zkTrie: {0}")]
    ZkTrie(Box<dyn Error + Send + Sync>),
}

/// Error variants encountered during verification of transactions in a L2 block.
#[derive(Debug, thiserror::Error)]
pub enum VerificationError {
    /// Error while operating on the database.
    #[error("database error: {0}")]
    Database(#[from] DatabaseError),
    /// Error while recovering signer from an ECDSA signature.
    #[error("invalid signature for tx_hash={tx_hash}: {source}")]
    InvalidSignature {
        /// The tx hash.
        tx_hash: B256,
        /// The source error that occurred while recovering signer.
        source: SignatureError,
    },
    /// Invalid gas price
    #[error("invalid gas price for tx_hash={tx_hash}")]
    InvalidGasPrice {
        /// The tx hash.
        tx_hash: B256,
    },
    /// Error encountered from [`revm`].
    #[error("error encountered from revm for tx_hash={tx_hash}: {source}")]
    EvmExecution {
        /// The tx hash.
        tx_hash: B256,
        /// The source error originating in [`revm`].
        source: EVMError<DatabaseError>,
    },
    /// Root mismatch error
    #[error("root_after in trace doesn't match with root_after in revm: root_trace={root_trace}, root_revm={root_revm}")]
    RootMismatch {
        /// Root after in trace
        root_trace: B256,
        /// Root after in revm
        root_revm: B256,
    },
}

impl DatabaseError {
    pub(crate) fn code_db<E: Error + Send + Sync + 'static>(err: E) -> Self {
        dev_error!(
            "code_db error {err} occurred in:\n{}",
            std::backtrace::Backtrace::force_capture()
        );
        DatabaseError::CodeDb(Box::new(err))
    }

    pub(crate) fn zk_trie<E: Error + Send + Sync + 'static>(err: E) -> Self {
        dev_error!(
            "zk_trie error {err} occurred in:\n{}",
            std::backtrace::Backtrace::force_capture()
        );
        DatabaseError::ZkTrie(Box::new(err))
    }
}
