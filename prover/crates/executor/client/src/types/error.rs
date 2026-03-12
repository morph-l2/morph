use alloy_primitives::{Address, FixedBytes, B256};
use prover_mpt::Error as MptError;

#[derive(Debug, thiserror::Error)]
pub enum ClientError {
    #[error("Failed to recover senders from signatures")]
    SignatureRecoveryFailed,
    #[error("Block header state root error")]
    InvalidHeaderStateRoot,
    #[error("Block state root error")]
    DiscontinuousStateRoot,
    #[error("root_after in trace doesn't match with root_after in revm: block number = {block_num} root_trace = {root_trace}, root_revm = {root_revm}")]
    MismatchedStateRoot {
        block_num: u64,
        /// Root after in trace
        root_trace: B256,
        /// Root after in revm
        root_revm: B256,
    },
    #[error("Mismatched storage root after executing the block")]
    MismatchedStorageRoot,
    #[error("unknown chain ID: {}", .0)]
    UnknownChainId(u64),
    #[error("Missing bytecode for account {}", .0)]
    MissingBytecode(Address),
    #[error("Missing trie for address {}", .0)]
    MissingTrie(Address),
    #[error("Invalid block number found in headers \n expected: {} found: {}", .0, .1)]
    InvalidHeaderBlockNumber(u64, u64),
    #[error("Invalid parent header found for block \n expected: {}, found: {}", .0, .1)]
    InvalidHeaderParentHash(FixedBytes<32>, FixedBytes<32>),
    #[error("Failed to validate post exectution state {}", 0)]
    PostExecutionError(String),
    #[error("Block Execution Failed: {}", .0)]
    BlockExecutionError(String),
    #[error("Mpt Error: {}", .0)]
    MptError(#[from] MptError),
}
