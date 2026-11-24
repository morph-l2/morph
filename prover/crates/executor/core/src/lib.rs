pub mod error;
pub mod executor;
mod validation;

pub use executor::MorphExecutor;
pub use validation::validate_next_l1_msg_index;

/// Chain IDs for different Morph networks.
pub const MAINNET_CHAIN_ID: u64 = 2818;
pub const TESTNET_CHAIN_ID: u64 = 2910;
pub const DEVNET_CHAIN_ID: u64 = 53077;
