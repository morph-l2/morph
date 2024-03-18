use std::{path::Path, str::FromStr, sync::Arc};

use c_kzg::KzgSettings;
use ethers::providers::{Http, Provider};
use once_cell::sync::Lazy;
use prometheus::{IntGauge, Registry};
use prover::BlockTrace;

// environment variables
pub static PROVER_PROOF_DIR: Lazy<String> = Lazy::new(|| read_env_var("PROVER_PROOF_DIR", "./proof".to_string()));
pub static PROVER_PARAMS_DIR: Lazy<String> =
    Lazy::new(|| read_env_var("PROVER_PARAMS_DIR", "./prove_params".to_string()));
pub static SCROLL_PROVER_ASSETS_DIR: Lazy<String> =
    Lazy::new(|| read_env_var("SCROLL_PROVER_ASSETS_DIR", "./config".to_string()));
pub static PROVER_L2_RPC: Lazy<String> = Lazy::new(|| read_env_var("PROVER_L2_RPC", "localhost:8545".to_string()));
pub static GENERATE_EVM_VERIFIER: Lazy<bool> = Lazy::new(|| read_env_var("GENERATE_EVM_VERIFIER", false));

// metrics
pub static REGISTRY: Lazy<Registry> = Lazy::new(|| Registry::new());
pub static PROVE_RESULT: Lazy<IntGauge> =
    Lazy::new(|| IntGauge::new("prove_result", "prove result").expect("prove metric can be created"));
pub static PROVE_TIME: Lazy<IntGauge> =
    Lazy::new(|| IntGauge::new("prove_time", "prove time").expect("time metric can be created"));

/// 4844 trusted setup config
pub static MAINNET_KZG_TRUSTED_SETUP: Lazy<Arc<KzgSettings>> = Lazy::new(|| Arc::new(load_trusted_setup()));

/// Loads the trusted setup parameters from the given bytes and returns the [KzgSettings].
pub fn load_trusted_setup() -> KzgSettings {
    let trusted_setup_file = Path::new("./configs/4844_trusted_setup.txt");
    assert!(trusted_setup_file.exists());
    let kzg_settings = KzgSettings::load_trusted_setup_file(trusted_setup_file).unwrap();
    return kzg_settings;
}

// Fetches block traces by provider
pub async fn get_block_traces_by_number(provider: &Provider<Http>, block_nums: &Vec<u64>) -> Option<Vec<BlockTrace>> {
    let mut block_traces: Vec<BlockTrace> = Vec::new();
    for block_num in block_nums {
        log::debug!("zkevm-prover: requesting trace of block {block_num}");
        let result = provider
            .request("morph_getBlockTraceByNumberOrHash", [format!("{block_num:#x}")])
            .await;

        match result {
            Ok(trace) => block_traces.push(trace),
            Err(e) => {
                log::error!("zkevm-prover: requesting trace error: {e}");
                return None;
            }
        }
    }
    Some(block_traces)
}

pub fn read_env_var<T: Clone + FromStr>(var_name: &'static str, default: T) -> T {
    std::env::var(var_name)
        .map(|s| s.parse::<T>().unwrap_or_else(|_| default.clone()))
        .unwrap_or(default)
}
