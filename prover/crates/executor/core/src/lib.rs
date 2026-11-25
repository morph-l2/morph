pub mod error;
pub mod executor;
mod validation;

pub use executor::{MorphExecutor, chain_spec_by_chain_id};
pub use validation::validate_block_pre_execution;

/// Chain IDs for different Morph networks.
pub const MAINNET_CHAIN_ID: u64 = 2818;
pub const TESTNET_CHAIN_ID: u64 = 2910;
pub const DEVNET_CHAIN_ID: u64 = 53077;
