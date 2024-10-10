use crate::{
    queue::{ProveRequest, Prover},
    read_env_var, PROVER_PROOF_DIR, PROVE_RESULT, PROVE_TIME, REGISTRY,
};
use axum::{
    routing::{get, post},
    Router,
};

use morph_prove::evm::EvmProofFixture;
use once_cell::sync::Lazy;
use prometheus::{Encoder, TextEncoder};
use serde::{Deserialize, Serialize};
use std::{
    fs,
    io::{BufReader, Read},
    sync::Arc,
    time::Duration,
};
use tokio::{sync::Mutex, time::timeout};
#[derive(Serialize, Deserialize, Debug, Default)]
pub struct ProveResult {
    pub error_msg: String,
    pub error_code: String,
    pub proof_data: Vec<u8>,
    pub pi_data: Vec<u8>,
    pub batch_header: Vec<u8>,
}

mod task_status {
    pub const STARTED: &str = "Started";
    pub const PROVING: &str = "Proving";
    pub const PROVED: &str = "Proved";
}

pub static MAX_PROVE_BLOCKS: Lazy<usize> = Lazy::new(|| read_env_var("MAX_PROVE_BLOCKS", 4096));

pub static PROVE_QUEUE: Lazy<Arc<Mutex<Vec<ProveRequest>>>> =
    Lazy::new(|| Arc::new(Mutex::new(vec![])));

// Main async function to start prover service.
// 1. Spawns prover mng.
// 2. Spawns metric mng.
// 3. Start the Prover on the main thread with shared queue.
// Server handles prove requests .
// Prover consumes requests and generates proofs and save.
pub async fn start() {
    // Step1. start prover management
    prover_mng().await;

    // Step2. start metric management
    metric_mng().await;

    // Step3. start prover
    start_prover().await;
}

async fn prover_mng() {
    tokio::spawn(async {
        let service = Router::new()
            .route("/prove_batch", post(add_pending_req))
            .route("/query_proof", post(query_prove_result))
            .route("/query_status", post(query_status));

        let mng_address = read_env_var("PROVER_MNG_ADDRESS", "0.0.0.0:3030".to_string());
        axum::Server::bind(&mng_address.parse().unwrap())
            .serve(service.into_make_service())
            .await
            .unwrap();
    });
}

async fn metric_mng() {
    prometheus::default_registry();
    REGISTRY.register(Box::new(PROVE_RESULT.clone())).unwrap();
    REGISTRY.register(Box::new(PROVE_TIME.clone())).unwrap();

    tokio::spawn(async move {
        let metrics = Router::new().route("/metrics", get(handle_metrics));
        let metric_address = read_env_var("PROVER_METRIC_ADDRESS", "0.0.0.0:6060".to_string());
        axum::Server::bind(&metric_address.parse().unwrap())
            .serve(metrics.into_make_service())
            .await
            .unwrap();
    });
}

async fn start_prover() {
    let mut prover = Prover::new(Arc::clone(&PROVE_QUEUE)).unwrap();
    prover.prove_for_queue().await;
}

async fn handle_metrics() -> String {
    let mut buffer = Vec::new();
    let encoder = TextEncoder::new();

    // Gather the metrics.
    let mut metric_families = REGISTRY.gather();
    metric_families.extend(prometheus::gather());

    // Encode metrics to send.
    match encoder.encode(&metric_families, &mut buffer) {
        Ok(()) => String::from_utf8(buffer.clone()).unwrap_or_default(),
        Err(e) => {
            log::error!("encode metrics error: {:#?}", e);
            String::from("")
        }
    }
}

// Add pending prove request to queue.
async fn add_pending_req(param: String) -> String {
    // Verify parameter is not empty
    if param.is_empty() {
        return String::from("request is empty");
    }

    // Deserialize parameter to ProveRequest type
    let prove_request: Result<ProveRequest, serde_json::Error> = serde_json::from_str(&param);

    // Handle deserialization result
    let prove_request = match prove_request {
        Ok(req) => req,
        Err(_) => return String::from("deserialize proveRequest failed"),
    };
    log::info!(
        "recived prove request of batch_index: {:#?}, shadow: {:#?}",
        prove_request.batch_index,
        prove_request.shadow.unwrap_or(false)
    );

    let blocks_len =
        prove_request.end_block.checked_sub(prove_request.start_block).unwrap_or_default();

    // Verify block number is greater than 0
    if prove_request.start_block == 0 || blocks_len == 0 {
        return String::from("blocks index invalid");
    }

    if blocks_len as usize > *MAX_PROVE_BLOCKS {
        return format!(
            "blocks len = {:?} exceeds MAX_PROVE_BLOCKS = {:?}",
            blocks_len, MAX_PROVE_BLOCKS
        );
    }

    // Verify RPC URL format
    if !prove_request.rpc.starts_with("http://") && !prove_request.rpc.starts_with("https://") {
        return String::from("invalid rpc url");
    }

    if let Some(value) = check_batch_status(&prove_request).await {
        return value;
    }

    let mut queue_lock = match timeout(Duration::from_secs(1), PROVE_QUEUE.lock()).await {
        Ok(queue_lock) => queue_lock,
        Err(_) => return String::from(task_status::PROVING),
    };

    if queue_lock.len() > 2 {
        return String::from("The task queue is full");
    }
    // Add request to queue
    log::info!("add pending req of batch: {:#?}", prove_request.batch_index);
    queue_lock.push(prove_request);
    String::from(task_status::STARTED)
}

// Async function to check status of a proof request for a batch.
// PROVED  -> there are already proven results.
async fn check_batch_status(prove_request: &ProveRequest) -> Option<String> {
    // Query proof data for the batch index.
    let proof = query_proof(prove_request.batch_index.to_string()).await;

    // If proof data is not empty, the batch has already been proven.
    if !proof.proof_data.is_empty() {
        log::warn!("there are already proven results: {:#?}", prove_request.batch_index);
        return Some(String::from(task_status::PROVED));
    }

    // Batch not found in any state.
    None
}

// Async function to query proof data for a given block number.
// It reads the proof directory and finds the file matching the block number.
// The file contents are read into a String which is returned.
async fn query_prove_result(batch_index: String) -> String {
    let result = query_proof(batch_index).await;
    serde_json::to_string(&result).unwrap()
}

async fn query_proof(batch_index: String) -> ProveResult {
    if batch_index.is_empty() {
        return ProveResult { error_msg: "batch_index is empty ".to_string(), ..Default::default() };
    }
    let proof_dir = match fs::read_dir(PROVER_PROOF_DIR.to_string()) {
        Ok(dir) => dir,
        Err(_) => {
            return ProveResult {
                error_msg: "Read proof dir error".to_string(),
                ..Default::default()
            }
        }
    };
    let mut result = ProveResult {
        error_msg: String::new(),
        error_code: String::new(),
        proof_data: Vec::new(),
        pi_data: Vec::new(),
        batch_header: Vec::new(),
    };
    log::info!("query proof of batch_index: {:#?}", batch_index);

    for entry in proof_dir {
        let path = match entry {
            Ok(entry) => entry.path(),
            Err(_) => {
                result.error_msg = String::from("Read proof dir error");
                return result;
            }
        };

        if path
            .to_str()
            .unwrap_or("nothing")
            .ends_with(format!("batch_{}", batch_index.trim()).as_str())
        {
            //pi_batch_agg.data
            let proof_path = path.join("plonk_proof.json");
            if !proof_path.exists() {
                result.error_msg = String::from("No plonk_proof file was found");
                return result;
            }
            // Proof
            let mut proof_data = Vec::new();
            match fs::File::open(proof_path) {
                Ok(file) => {
                    let reader = BufReader::new(file);
                    let proof: EvmProofFixture = serde_json::from_reader(reader).unwrap();
                    proof_data.extend(alloy::hex::decode(proof.proof).unwrap());
                }
                Err(e) => {
                    log::error!("Failed to load proof_data: {:#?}", e);
                    result.error_msg = String::from("Failed to load proof_data");
                }
            }
            result.proof_data = proof_data;

            //Batch header data
            let batch_header_path = path.join("batch_header.data");
            let mut batch_header = Vec::new();
            match fs::File::open(batch_header_path) {
                Ok(mut file) => {
                    file.read_to_end(&mut batch_header).unwrap();
                }
                Err(e) => {
                    log::warn!("Failed to load batch_header: {:#?}", e);
                    result.error_msg = String::from("Failed to load batch_header");
                }
            }
            result.batch_header = batch_header;
            break;
        }
    }
    if result.proof_data.is_empty() {
        result.error_msg = String::from("No proof was found");
    }
    result
}

// Async function to check queue status.
// Locks queue and returns length > 0 ? "not empty" : "empty"
async fn query_status() -> String {
    let queue_try_lock = PROVE_QUEUE.try_lock();
    if queue_try_lock.is_err() {
        return String::from("1");
    }
    let queue_lock = queue_try_lock.ok().unwrap();
    match queue_lock.len() {
        0 => String::from("0"),
        _ => String::from("1"),
    }
}
