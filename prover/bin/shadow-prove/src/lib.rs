use std::env::var;

use abi::{Rollup, ShadowRollup};
use once_cell::sync::Lazy;

use crate::util::read_env_var;

pub mod abi;
pub mod execute;
pub mod metrics;
pub mod shadow_prove;
pub mod shadow_rollup;
pub mod util;

#[derive(Clone, Debug)]
pub struct BatchInfo {
    pub batch_index: u64,
    pub start_block: u64,
    pub end_block: u64,
    pub total_txn: u64,
}
pub static SHADOW_EXECUTE: Lazy<bool> = Lazy::new(|| read_env_var("SHADOW_PROVING_EXECUTE", true));

#[derive(Clone, Copy, Debug, PartialEq, Eq)]
pub enum ShadowExecuteMode {
    Rpc = 0,
    Witness = 1,
    Both = 2,
}

impl std::str::FromStr for ShadowExecuteMode {
    type Err = &'static str;

    fn from_str(value: &str) -> Result<Self, Self::Err> {
        match value {
            "0" => Ok(Self::Rpc),
            "1" => Ok(Self::Witness),
            "2" => Ok(Self::Both),
            _ => Err("expected 0 (RPC), 1 (witness), or 2 (both)"),
        }
    }
}

/// `SHADOW_PROVING_EXECUTE_MODE`: 0 = RPC (default), 1 = witness, 2 = both.
pub static SHADOW_EXECUTE_MODE: Lazy<ShadowExecuteMode> =
    Lazy::new(|| read_env_var("SHADOW_PROVING_EXECUTE_MODE", ShadowExecuteMode::Rpc));

pub static SHADOW_PROVING_MAX_BLOCK: Lazy<u64> =
    Lazy::new(|| read_env_var("SHADOW_PROVING_MAX_BLOCK", 600));

pub static SHADOW_PROVING_MAX_TXN: Lazy<u64> =
    Lazy::new(|| read_env_var("SHADOW_PROVING_MAX_TXN", 200));

pub static SHADOW_PROVING_BLOCKS_RANGE: Lazy<u64> =
    Lazy::new(|| read_env_var("SHADOW_PROVING_BLOCKS_RANGE", 600));

pub static SHADOW_PROVING_PROVER_RPC: Lazy<String> =
    Lazy::new(|| var("SHADOW_PROVING_PROVER_RPC").expect("Cannot detect PROVER_RPC env var"));

pub static SHADOW_PROVING_BATCH_INTERVAL: Lazy<u64> =
    Lazy::new(|| read_env_var("SHADOW_PROVING_BATCH_INTERVAL", 0));
