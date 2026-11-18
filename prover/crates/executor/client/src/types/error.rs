use alloy_primitives::{Address, B256, FixedBytes, U256};
use prover_mpt::Error as MptError;

#[derive(Debug, thiserror::Error)]
pub enum ClientError {
    #[error("Failed to verify raw blob: {0}")]
    VerifyRawError(#[source] anyhow::Error),
    #[error("Failed to recover senders from signatures")]
    SignatureRecoveryFailed,
    #[error("Block header state root error")]
    InvalidHeaderStateRoot,
    #[error("ChainId not match")]
    ChainIdNotMatch,
    #[error("Block state root error")]
    DiscontinuousStateRoot,
    #[error(
        "root_after in trace doesn't match with root_after in revm: block number = {block_num} root_trace = {root_trace}, root_revm = {root_revm}"
    )]
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
    #[error("Invalid block timestamp for block: {}", .0)]
    InvalidBlockTimestamp(u64),
    #[error("Invalid base fee found in header \n expected: {} found: {}", .0, .1)]
    InvalidHeaderBaseFee(u64, u64),
    #[error("Invalid parent header found for block \n expected: {}, found: {}", .0, .1)]
    InvalidHeaderParentHash(FixedBytes<32>, FixedBytes<32>),
    #[error("Invalid mix hash in header: block number = {block_num}, mix hash = {mix_hash}")]
    InvalidHeaderMixHash { block_num: u64, mix_hash: B256 },
    #[error("Invalid difficulty in header: block number = {block_num}, difficulty = {difficulty}")]
    InvalidHeaderDifficulty { block_num: u64, difficulty: U256 },
    #[error(
        "L1 message follows an L2 transaction: block number = {block_num}, tx index = {tx_index}"
    )]
    L1MessageAfterL2Tx { block_num: u64, tx_index: usize },
    #[error(
        "L1 message queue index not consecutive: block number = {block_num}, expected = {expected}, actual = {actual}"
    )]
    L1MessageNotConsecutive { block_num: u64, expected: u64, actual: u64 },
    #[error("Malformed L1 message: block number = {block_num}, tx index = {tx_index}")]
    MalformedL1Message { block_num: u64, tx_index: usize },
    #[error("Failed to validate post exectution state {}", 0)]
    PostExecutionError(String),
    #[error("Block Execution Failed: {}", .0)]
    BlockExecutionError(String),
    #[error("Mpt Error: {}", .0)]
    MptError(#[from] MptError),
}
