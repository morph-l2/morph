use alloy_primitives::{Address, FixedBytes};
use rsp_mpt::Error as MptError;

#[derive(Debug, thiserror::Error)]
pub enum ClientError {
    #[error("Failed to recover senders from signatures")]
    SignatureRecoveryFailed,
    #[error("Block header state root error")]
    InvalidHeaderStateRoot,
    #[error("Block state root error")]
    DiscontinuousStateRoot,
    #[error("Mismatched state root after executing the block, block number: {}", .0)]
    MismatchedStateRoot(u64),
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
