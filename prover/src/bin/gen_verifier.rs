use std::{fs::File, sync::Arc};

use dotenv::dotenv;
use prover::BlockTrace;
use tokio::sync::Mutex;
use zkevm_prover::{
    prover::{ProveRequest, Prover},
    utils::{read_env_var, PROVER_PROOF_DIR},
};

#[tokio::main]
async fn main() {
    dotenv().ok();
    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("debug")).init();

    let batch_index = read_env_var("TRACES_INDEX", 1);
    let queue: Arc<Mutex<Vec<ProveRequest>>> = Arc::new(Mutex::new(Vec::new()));
    let mut prover = Prover::new(queue);

    prover.generate_proof(batch_index, load_trace(batch_index)).await;
}

fn load_trace(batch_index: u64) -> Vec<Vec<BlockTrace>> {
    use std::io::BufReader;

    let path = PROVER_PROOF_DIR.to_string() + format!("/batch_{}", batch_index).as_str();
    let file = File::open(format!("{}/chunk_traces.json", path.as_str())).unwrap();
    let reader = BufReader::new(file);

    let chunk_traces: Vec<Vec<BlockTrace>> = serde_json::from_reader(reader).unwrap();
    log::info!(
        "Load traces from file successful, block num = {:?}",
        chunk_traces.len()
    );

    chunk_traces
}
