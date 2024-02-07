use axum::extract::Extension;
use axum::routing::get;
use axum::{routing::post, Router};
use dotenv::dotenv;
use env_logger::Env;
use prometheus::{Encoder, TextEncoder};
use serde::{Deserialize, Serialize};
use std::fs;
use std::io::Read;
use std::sync::Arc;
use std::time::Duration;
use tokio::sync::Mutex;
use tokio::time::timeout;
use zkevm_prover::utils::{PROVER_PROOF_DIR, PROVE_TIME};
use zkevm_prover::utils::{read_env_var, PROVE_RESULT, REGISTRY};

use tower_http::add_extension::AddExtensionLayer;
use tower_http::cors::CorsLayer;
use tower_http::trace::TraceLayer;
use zkevm_prover::prover::{prove_for_queue, ProveRequest};

#[derive(Serialize, Deserialize, Debug)]
pub struct ProveResult {
    pub error_msg: String,
    pub error_code: String,
    pub proof_data: Vec<u8>,
    pub pi_data: Vec<u8>,
}

mod task_status {
    pub const STARTED: &str = "Started";
    pub const PROVING: &str = "Proving";
    pub const PROVED: &str = "Proved";
}

// Main async function to start prover service.
// 1. Initializes environment.
// 2. Spawns management server.
// 3. Start the Prover on the main thread with shared queue.
// Server handles prove requests .
// Prover consumes requests and generates proofs and save.
#[tokio::main]
async fn main() {
    // Step1. prepare environment
    dotenv().ok();
    env_logger::Builder::from_env(Env::default().default_filter_or("info")).init();
    log::info!(
        "SCROLL_PROVER_ASSETS_DIR env: {:#?}",
        zkevm_prover::utils::SCROLL_PROVER_ASSETS_DIR.as_str()
    );

    // Step2. start prover management
    let queue: Arc<Mutex<Vec<ProveRequest>>> = Arc::new(Mutex::new(Vec::new()));
    prover_mng(Arc::clone(&queue)).await;

    // Step3. start metric management
    metric_mng().await;

    // Step4. start prover
    prove_for_queue(Arc::clone(&queue)).await;
}

async fn prover_mng(task_queue: Arc<Mutex<Vec<ProveRequest>>>) {
    tokio::spawn(async {
        let service = Router::new()
            .route("/prove_batch", post(add_pending_req))
            .route("/query_proof", post(query_prove_result))
            .route("/query_status", post(query_status))
            .layer(AddExtensionLayer::new(task_queue))
            .layer(CorsLayer::permissive())
            .layer(TraceLayer::new_for_http());

        axum::Server::bind(&"0.0.0.0:3030".parse().unwrap())
            .serve(service.into_make_service())
            .await
            .unwrap();
    });
}

async fn metric_mng() {
    prometheus::default_registry();
    REGISTRY.register(Box::new(PROVE_RESULT.clone())).unwrap();
    REGISTRY.register(Box::new(PROVE_TIME.clone())).unwrap();

    let metric_address = read_env_var("PROVER_METRIC_ADDRESS", "0.0.0.0:6060".to_string());
    tokio::spawn(async move {
        let metrics = Router::new()
            .route("/metrics", get(handle_metrics))
            .layer(TraceLayer::new_for_http());
        axum::Server::bind(&metric_address.parse().unwrap())
            .serve(metrics.into_make_service())
            .await
            .unwrap();
    });
}

async fn handle_metrics() -> String {
    let mut buffer = Vec::new();
    let encoder = TextEncoder::new();

    // Gather the metrics.
    let mut metric_families = REGISTRY.gather();
    metric_families.extend(prometheus::gather());

    // Encode metrics to send.
    match encoder.encode(&metric_families, &mut buffer) {
        Ok(()) => {
            let output = String::from_utf8(buffer.clone()).unwrap();
            return output;
        }
        Err(e) => {
            log::error!("encode metrics error: {:#?}", e);
            return String::from("");
        }
    }
}

// Add pending prove request to queue.
async fn add_pending_req(Extension(queue): Extension<Arc<Mutex<Vec<ProveRequest>>>>, param: String) -> String {
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
    log::info!("recived prove request of batch_index: {:#?}", prove_request.batch_index);

    // Verify block number is greater than 0
    if prove_request.chunks.len() == 0 {
        return String::from("chunks is empty");
    }
    for chunk in &prove_request.chunks {
        if chunk.len() == 0 {
            return String::from("blocks is empty");
        }
    }

    // Verify RPC URL format
    if !prove_request.rpc.starts_with("http://") && !prove_request.rpc.starts_with("https://") {
        return String::from("invalid rpc url");
    }

    if let Some(value) = check_batch_status(&prove_request).await {
        return value;
    }

    let mut queue_lock = match timeout(Duration::from_secs(1), queue.lock()).await {
        Ok(queue_lock) => queue_lock,
        Err(_) => return String::from(task_status::PROVING),
    };

    if queue_lock.len() > 0 {
        return String::from(task_status::PROVING);
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
    return serde_json::to_string(&result).unwrap();
}

async fn query_proof(batch_index: String) -> ProveResult {
    let proof_dir: Result<fs::ReadDir, std::io::Error> = fs::read_dir(PROVER_PROOF_DIR.to_string());
    let mut result = ProveResult {
        error_msg: String::new(),
        error_code: String::new(),
        proof_data: Vec::new(),
        pi_data: Vec::new(),
    };
    log::info!("query proof of batch_index: {:#?}", batch_index);

    if proof_dir.is_err() {
        result.error_msg = String::from("Read proof dir error");
        return result;
    }

    for entry in proof_dir.unwrap() {
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
            .ends_with(format!("batch_{}", batch_index).as_str())
        {
            //pi_batch_agg.data
            let proof_path = path.join("proof_batch_agg.data");
            if !proof_path.exists() {
                result.error_msg = String::from("No proof_batch_agg file");
                return result;
            }
            // let mut proof_data = String::new();
            let mut proof_data = Vec::new();

            match fs::File::open(proof_path) {
                Ok(mut file) => {
                    file.read_to_end(&mut proof_data).unwrap();
                }
                Err(e) => {
                    log::error!("Failed to load proof_data: {:#?}", e);
                    result.error_msg = String::from("Failed to load proof_data");
                }
            }
            result.proof_data = proof_data;

            let pi_path = path.join("pi_batch_agg.data");
            let mut pi_data = Vec::new();

            match fs::File::open(pi_path) {
                Ok(mut file) => {
                    file.read_to_end(&mut pi_data).unwrap();
                }
                Err(e) => {
                    log::error!("Failed to load pi_data: {:#?}", e);
                    result.error_msg = String::from("Failed to load pi_data");
                }
            }
            result.pi_data = pi_data;
            break;
        }
    }

    return result;
}

// Async function to check queue status.
// Locks queue and returns length > 0 ? "not empty" : "empty"
async fn query_status(Extension(queue): Extension<Arc<Mutex<Vec<ProveRequest>>>>) -> String {
    let queue_try_lock = queue.try_lock();
    if queue_try_lock.is_err() {
        return String::from("1");
    }
    let queue_lock = queue_try_lock.ok().unwrap();
    match queue_lock.len() {
        0 => String::from("0"),
        _ => String::from("1"),
    }
}

#[tokio::test]
async fn test_gen_request() {
    let request: ProveRequest = ProveRequest {
        batch_index: 4,
        chunks: vec![vec![1], vec![2, 3]],
        rpc: String::from("127.0.0.1:8569"),
    };
    let info = serde_json::to_string(&request);
    println!("{:?}", info.unwrap());
}
