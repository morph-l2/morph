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
pub static SHADOW_EXECUTE: Lazy<bool> = Lazy::new(|| read_env_var("SHADOW_PROVING_EXECUTE", false));

pub static SHADOW_EXECUTE_USE_RPC_DB: Lazy<bool> =
    Lazy::new(|| read_env_var("SHADOW_PROVING_EXECUTE_USE_RPC_DB", true));

pub static SHADOW_PROVING_MAX_BLOCK: Lazy<u64> =
    Lazy::new(|| read_env_var("SHADOW_PROVING_MAX_BLOCK", 600));

pub static SHADOW_PROVING_MAX_TXN: Lazy<u64> =
    Lazy::new(|| read_env_var("SHADOW_PROVING_MAX_TXN", 200));

pub static SHADOW_PROVING_PROVER_RPC: Lazy<String> =
    Lazy::new(|| var("SHADOW_PROVING_PROVER_RPC").expect("Cannot detect PROVER_RPC env var"));
