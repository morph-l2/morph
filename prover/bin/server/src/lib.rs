pub mod server;

use std::str::FromStr;
pub mod queue;

use once_cell::sync::Lazy;
use prometheus::{IntGauge, Registry};

// environment variables
pub static PROVER_PROOF_DIR: Lazy<String> = Lazy::new(|| read_env_var("PROVER_PROOF_DIR", "./proof".to_string()));
pub static PROVER_L2_RPC: Lazy<String> = Lazy::new(|| read_env_var("PROVER_L2_RPC", "localhost:8545".to_string()));

// metrics
pub static REGISTRY: Lazy<Registry> = Lazy::new(Registry::new);
pub static PROVE_RESULT: Lazy<IntGauge> =
    Lazy::new(|| IntGauge::new("prove_result", "prove result").expect("prove metric can be created")); // 1 = Ok, 2 = Fail
pub static PROVE_TIME: Lazy<IntGauge> =
    Lazy::new(|| IntGauge::new("prove_time", "prove time").expect("time metric can be created"));

pub fn read_env_var<T: Clone + FromStr>(var_name: &'static str, default: T) -> T {
    std::env::var(var_name)
        .map(|s| s.parse::<T>().unwrap_or_else(|_| default.clone()))
        .unwrap_or(default)
}
