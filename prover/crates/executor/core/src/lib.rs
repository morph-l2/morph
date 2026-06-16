pub mod error;
pub mod executor;

pub use executor::{build_morph_block, MorphExecutor};

/// Chain IDs for different Morph networks.
pub const MAINNET_CHAIN_ID: u64 = 2818;
pub const TESTNET_CHAIN_ID: u64 = 2910;
pub const DEVNET_CHAIN_ID: u64 = 53077;
