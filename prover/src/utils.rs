use std::{path::Path, str::FromStr, sync::Arc};

use c_kzg::KzgSettings;
use ethers::{
    core::k256::sha2::{Digest, Sha256},
    providers::{Http, Provider},
};
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
pub static REGISTRY: Lazy<Registry> = Lazy::new(Registry::new);
pub static PROVE_RESULT: Lazy<IntGauge> =
    Lazy::new(|| IntGauge::new("prove_result", "prove result").expect("prove metric can be created")); // 1 = Ok, 2 = Fail
pub static PROVE_TIME: Lazy<IntGauge> =
    Lazy::new(|| IntGauge::new("prove_time", "prove time").expect("time metric can be created"));

/// 4844 trusted setup config
pub static MAINNET_KZG_TRUSTED_SETUP: Lazy<Arc<KzgSettings>> = Lazy::new(|| Arc::new(load_trusted_setup()));

/// Loads the trusted setup parameters from the given bytes and returns the [KzgSettings].
pub fn load_trusted_setup() -> KzgSettings {
    let setup_config = SCROLL_PROVER_ASSETS_DIR.to_string() + "/4844_trusted_setup.txt";
    let trusted_setup_file = Path::new(setup_config.as_str());
    assert!(trusted_setup_file.exists());
    KzgSettings::load_trusted_setup_file(trusted_setup_file).unwrap()
}

pub fn kzg_to_versioned_hash(commitment: &[u8]) -> Vec<u8> {
    let mut hasher = Sha256::new();
    hasher.update(commitment);
    let hash = hasher.finalize();
    let mut hashed_bytes = hash.as_slice().to_vec();
    hashed_bytes[0] = 0x01;
    hashed_bytes
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

#[test]
fn test_kzg() {
    use c_kzg::{Blob, KzgCommitment};

    let mut raw_data = [0u8; 131072];
    // Construct an effective Blob data
    let payload = br#"EIP-4844 introduces a new kind of transaction type to Ethereum which accepts "blobs"
        of data to be persisted in the beacon node for a short period of time. These changes are
        forwards compatible with Ethereum's scaling roadmap, and blobs are small enough to keep disk use manageable."#;

    let mut offset = 0;
    for chunk in payload.chunks(27) {
        let chunk_len = chunk.len() as u32;
        raw_data[offset + 1..offset + 5].copy_from_slice(&chunk_len.to_le_bytes());
        raw_data[offset + 5..offset + 5 + chunk_len as usize].copy_from_slice(chunk);
        offset += 5 + chunk_len as usize;
    }

    let kzg_settings: Arc<c_kzg::KzgSettings> = Arc::clone(&MAINNET_KZG_TRUSTED_SETUP);
    let commitment = match KzgCommitment::blob_to_kzg_commitment(&Blob::from_bytes(&raw_data).unwrap(), &kzg_settings) {
        Ok(c) => c,
        Err(e) => {
            log::error!("generate KzgCommitment error: {:#?}", e);
            return;
        }
    };

    let versioned_hash = kzg_to_versioned_hash(commitment.to_bytes().to_vec().as_slice());
    log::info!("versioned_hash_Hex= {:#?}", ethers::utils::hex::encode(versioned_hash));
}
