use axum::extract::Extension;
use axum::routing::get;
use axum::{routing::post, Router};
use dotenv::dotenv;
use prometheus::{Encoder, TextEncoder};
use serde::{Deserialize, Serialize};
use std::fs;
use std::io::Read;
use std::sync::Arc;
use std::time::Duration;
use tokio::sync::Mutex;
use tokio::time::timeout;
use zkevm_prover::utils::{read_env_var, PROVE_RESULT, REGISTRY};
use zkevm_prover::utils::{PROVER_PROOF_DIR, PROVE_TIME};

use flexi_logger::{Cleanup, Criterion, Duplicate, FileSpec, Logger, Naming, WriteMode};
use log::Record;
use tower_http::add_extension::AddExtensionLayer;
use tower_http::cors::CorsLayer;
use tower_http::trace::TraceLayer;
use zkevm_prover::prover::{ProveRequest, Prover};
#[derive(Serialize, Deserialize, Debug, Default)]
pub struct ProveResult {
    pub error_msg: String,
    pub error_code: String,
    pub proof_data: Vec<u8>,
    pub pi_data: Vec<u8>,
    pub blob_kzg: Vec<u8>,
    pub batch_header: Vec<u8>,
}

mod task_status {
    pub const STARTED: &str = "Started";
    pub const PROVING: &str = "Proving";
    pub const PROVED: &str = "Proved";
}

// Main async function to start prover service.
// 1. Initializes environment.
// 2. Spawns prover mng.
// 3. Spawns metric mng.
// 4. Start the Prover on the main thread with shared queue.
// Server handles prove requests .
// Prover consumes requests and generates proofs and save.
#[tokio::main]
async fn main() {
    // Step1. prepare environment
    dotenv().ok();

    // Initialize logger.
    setup_logging();

    // Step2. start prover management
    let queue: Arc<Mutex<Vec<ProveRequest>>> = Arc::new(Mutex::new(Vec::new()));
    prover_mng(Arc::clone(&queue)).await;

    // Step3. start metric management
    metric_mng().await;

    // Step4. start prover
    start_prover(Arc::clone(&queue)).await;
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
        let metrics = Router::new()
            .route("/metrics", get(handle_metrics))
            .layer(TraceLayer::new_for_http());
        let metric_address = read_env_var("PROVER_METRIC_ADDRESS", "0.0.0.0:6060".to_string());
        axum::Server::bind(&metric_address.parse().unwrap())
            .serve(metrics.into_make_service())
            .await
            .unwrap();
    });
}

async fn start_prover(task_queue: Arc<Mutex<Vec<ProveRequest>>>) {
    let mut prover = Prover::new(task_queue);
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
    log::info!(
        "recived prove request of batch_index: {:#?}, shadow: {:#?}",
        prove_request.batch_index,
        prove_request.shadow.unwrap_or(false)
    );

    // Verify block number is greater than 0
    if prove_request.chunks.is_empty() {
        return String::from("chunks is empty");
    }
    for chunk in &prove_request.chunks {
        if chunk.is_empty() {
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
        blob_kzg: Vec::new(),
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
            .ends_with(format!("batch_{}", batch_index).as_str())
        {
            //pi_batch_agg.data
            let proof_path = path.join("proof_batch_agg.data");
            if !proof_path.exists() {
                result.error_msg = String::from("No proof_batch_agg file");
                return result;
            }
            // Proof
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

            // Public input
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

            // Eip4844 kzg data
            let blob_kzg_path = path.join("blob_kzg.data");
            let mut blob_kzg = Vec::new();
            match fs::File::open(blob_kzg_path) {
                Ok(mut file) => {
                    file.read_to_end(&mut blob_kzg).unwrap();
                }
                Err(e) => {
                    log::error!("Failed to load blob_kzg: {:#?}", e);
                    result.error_msg = String::from("Failed to load blob_kzg");
                }
            }
            result.blob_kzg = blob_kzg;

            // Batch header data
            let batch_header_path = path.join("batch_header.data");
            let mut batch_header = Vec::new();
            match fs::File::open(batch_header_path) {
                Ok(mut file) => {
                    file.read_to_end(&mut batch_header).unwrap();
                }
                Err(e) => {
                    log::error!("Failed to load batch_header: {:#?}", e);
                    result.error_msg = String::from("Failed to load batch_header");
                }
            }
            result.batch_header = batch_header;
            break;
        }
    }
    result
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

// Constants for configuration
const LOG_LEVEL: &str = "info";
const LOG_FILE_BASENAME: &str = "app_info";
const LOG_FILE_SIZE_LIMIT: u64 = 200 * 10u64.pow(6); // 200MB
                                                     // const LOG_FILE_SIZE_LIMIT: u64 = 10u64.pow(3); // 1kB
const LOG_FILES_TO_KEEP: usize = 3;
fn setup_logging() {
    //configure the logger
    Logger::try_with_env_or_str(LOG_LEVEL)
        .unwrap()
        .log_to_file(
            FileSpec::default()
                .directory(read_env_var("PROVER_LOG_DIR", String::from("/data/logs/morph-prover")))
                .basename(LOG_FILE_BASENAME),
        )
        .format(log_format)
        .duplicate_to_stdout(Duplicate::All)
        .rotate(
            Criterion::Size(LOG_FILE_SIZE_LIMIT), // Scroll when file size reaches 200MB
            Naming::TimestampsCustomFormat {
                current_infix: Some(""),
                format: "r%Y-%m-%d_%H-%M-%S",
            }, // Using timestamps as part of scrolling files
            Cleanup::KeepLogFiles(LOG_FILES_TO_KEEP), // Keep the latest 3 scrolling files
        )
        .write_mode(WriteMode::BufferAndFlush)
        .start()
        .unwrap();
}

fn log_format(
    w: &mut dyn std::io::Write,
    now: &mut flexi_logger::DeferredNow,
    record: &Record,
) -> Result<(), std::io::Error> {
    write!(
        w,
        "{} [{}] {} - {}",
        now.now().format("%Y-%m-%d %H:%M:%S"), // Custom time format
        record.level(),
        record.target(),
        record.args()
    )
}
