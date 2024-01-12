use axum::extract::Extension;
use axum::{routing::post, Router};
use env_logger::Env;
use serde::{Deserialize, Serialize};
use std::fs;
use std::io::Read;
use std::sync::Arc;
use tokio::sync::Mutex;
use dotenv::dotenv;
use std::env::var;

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
    // dotenv().ok();
    env_logger::Builder::from_env(Env::default().default_filter_or("debug")).init();

    // fs::create_dir_all(FS_PROOF).unwrap();
    let queue: Arc<Mutex<Vec<ProveRequest>>> = Arc::new(Mutex::new(Vec::new()));

    // Step2. start mng
    let task_queue: Arc<Mutex<Vec<ProveRequest>>> = Arc::clone(&queue);
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

    // Step3. start prover
    let prove_queue: Arc<Mutex<Vec<ProveRequest>>> = Arc::clone(&queue);
    prove_for_queue(prove_queue).await;
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

    // Verify RPC URL format
    if !prove_request.rpc.starts_with("http://") && !prove_request.rpc.starts_with("https://") {
        return String::from("invalid rpc url");
    }

    let mut queue_lock = queue.lock().await;
    if queue_lock.len() > 0 {
        return String::from("add prove batch fail, please waiting for the pending task to complete");
    }
    // if let Some(value) = check_batch_status(&prove_request).await {
    //     return value;
    // }
    
    dotenv().ok();
    let prover_proof = var("PROVER_PROOF_DIR").expect("PROVER_PROOF env var");
    let proof_path = prover_proof + format!("/batch_{}", &prove_request.batch_index).as_str();
    fs::create_dir_all(proof_path.clone()).unwrap();
    log::info!("add pending req of batch: {:#?}", prove_request.batch_index);

    // Add request to queue
    queue_lock.push(prove_request);
    String::from(task_status::STARTED)
}

// Async function to check status of a proof request for a batch.
// PROVING -> prover is proving this batch.
// PROVED  -> there are already proven results.
async fn check_batch_status(prove_request: &ProveRequest) -> Option<String> {
    // Query proof data for the batch index.
    let proof = query_proof(prove_request.batch_index.to_string()).await;

    // If proof data is not empty, the batch has already been proven.
    if !proof.proof_data.is_empty() || !proof.pi_data.is_empty() {
        log::warn!("there are already proven results: {:#?}", prove_request.batch_index);
        return Some(String::from(task_status::PROVED));
    }

    // Read proof directory.
    dotenv().ok();
    let prover_proof = var("PROVER_PROOF_DIR").expect("PROVER_PROOF env var");
    let proof_dir: Result<fs::ReadDir, std::io::Error> = fs::read_dir(prover_proof);
    let entries = match proof_dir {
        Ok(entries) => entries,
        Err(_) => return Some(String::from("Read proof dir error")),
    };

    for entry in entries {
        let path = match entry {
            Ok(entry) => entry.path(),
            Err(_) => return Some(String::from("Read proof dir error")),
        };

        // If entry contains this batch index, it is currently being proven.
        if path
            .to_str()
            .unwrap_or("nothing")
            .ends_with(format!("batch_{}", prove_request.batch_index).as_str())
        {
            log::warn!("Prover is proving this batch: {:#?}", prove_request.batch_index);
            return Some(String::from(task_status::PROVING));
        }
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
    dotenv().ok();
    let prover_proof = var("PROVER_PROOF_DIR").expect("PROVER_PROOF env var");

    let proof_dir: Result<fs::ReadDir, std::io::Error> = fs::read_dir(prover_proof);
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
    match queue.lock().await.len() {
        0 => String::from("0"),
        _ => String::from("1"),
    }
}

#[tokio::test]
async fn test_query_proof() {
    use std::path::Path;
    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("debug")).init();

    let proof = query_prove_result("1".to_string()).await;
    let prove_result: ProveResult = match serde_json::from_str(proof.as_str()) {
        Ok(pr) => pr,
        Err(_) => {
            log::error!("deserialize prove_result failed, batch index = {:#?}", 1);
            return;
        }
    };
    use ethers::abi::AbiDecode;
    use std::str::FromStr;

    let aggr_proof = ethers::types::Bytes::from(prove_result.proof_data);
    //     Ok(ap) => ap,
    //     Err(e) => {
    //         log::error!("decode proof_data failed, error = {:#?}", e);
    //         return;
    //     }
    // };
    let mut proof_data1 = aggr_proof.to_vec();

    // println!("{:?}", aggr_proof);

    // let proof_path = Path("proof_batch_agg.data");
    let proof_path = Path::new("proof/batch_1/proof_batch_agg.data");

    let mut proof_data = Vec::new();
    // let mut proof_data = ethers::types::Bytes::new();

    match fs::File::open(proof_path) {
        Ok(mut file) => {
            file.read_to_end(&mut proof_data).unwrap();
        }
        Err(e) => {
            log::error!("Failed to load proof_data: {:#?}", e);
        }
    }
    println!("{:?}", aggr_proof);
}

#[tokio::test]
async fn test() {
    let request = ProveRequest {
        batch_index: 4,
        chunks: vec![vec![1], vec![2, 3]],
        rpc: String::from("127.0.0.1:8569"),
    };
    let info = serde_json::to_string(&request);
    println!("{:?}", info.unwrap());
}
