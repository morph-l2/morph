use std::{
    fs::{self, File},
    io::BufWriter,
    sync::Arc,
    thread,
    time::Duration,
};

use crate::{read_env_var, PROVER_L2_RPC, PROVER_PROOF_DIR};
use alloy::providers::{ProviderBuilder, RootProvider};
use anyhow::anyhow;
use morph_prove::prove;
use prover_primitives::types::BlockTrace;
use serde::{Deserialize, Serialize};
use tokio::sync::Mutex;

// proveRequest
#[derive(Serialize, Deserialize, Debug)]
pub struct ExecuteRequest {
    pub batch_index: u64,
    pub start_block: u64,
    pub end_block: u64,
    pub rpc: String,
}

pub struct Executor {
    execute_queue: Arc<Mutex<Vec<ExecuteRequest>>>,
    provider: RootProvider,
}

impl Executor {
    pub fn new(execute_queue: Arc<Mutex<Vec<ExecuteRequest>>>) -> Result<Self, anyhow::Error> {
        let url = reqwest::Url::parse(PROVER_L2_RPC.as_str())
            .map_err(|_| anyhow!("Invalid L2 RPC URL"))?;
        let provider = ProviderBuilder::new().on_provider(RootProvider::new_http(url));

        Ok(Self { execute_queue, provider })
    }

    /// Processes execute requests from a queue asynchronously.
    pub async fn execute_for_queue(&mut self) {
        log::info!("Waiting for execute request");
        loop {
            thread::sleep(Duration::from_millis(12000));

            // Step1. Get request from queue
            let req = match self.execute_queue.lock().await.pop() {
                Some(req) => {
                    log::info!(
                        "start execute batch, batch index = {:#?}, blocks len = {:#?}, start_block = {:#?}, end_block = {:#?}",
                        req.batch_index,
                        req.end_block - req.start_block + 1,
                        req.start_block,
                        req.end_block,
                    );
                    req
                }
                None => {
                    log::info!("no execute request");
                    continue;
                }
            };

            // Spawn async task to handle the execution
            let provider = self.provider.clone();
            execute_batch(req, provider).await;
        }
    }
}

/// Executes a batch asynchronously.
async fn execute_batch(req: ExecuteRequest, provider: RootProvider) {
    // Step1. Fetch trace
    log::info!("Requesting trace of batch-{:#?} ...", req.batch_index);
    let mut block_traces =
        match get_block_traces(req.batch_index, req.start_block, req.end_block, &provider).await {
            Some(traces) => traces,
            None => {
                log::error!("Failed to get block traces for batch {}", req.batch_index);
                return;
            }
        };

    if read_env_var("SAVE_TRACE", false) {
        save_trace(req.batch_index, &block_traces);
    }

    // Step2. Execute batch (without generating proof)
    log::info!("Executing evm proof for batch-{}", req.batch_index);
    let prove_rt = prove(&mut block_traces, false);
    match prove_rt {
        Ok(_) => {
            log::info!("Successfully executed batch-{}", req.batch_index,);
        }
        Err(e) => {
            log::error!("Execution of batch-{} error: {:?}", req.batch_index, e);
        }
    }
}

// Fetches block traces by provider
async fn get_block_traces(
    batch_index: u64,
    start_block: u64,
    end_block: u64,
    provider: &RootProvider,
) -> Option<Vec<BlockTrace>> {
    let mut block_traces: Vec<BlockTrace> = Vec::new();
    for block_num in start_block..end_block + 1 {
        log::debug!("requesting trace of block {block_num}");
        let result = provider
            .raw_request("morph_getBlockTraceByNumberOrHash".into(), [format!("{block_num:#x}")])
            .await;

        match result {
            Ok(trace) => block_traces.push(trace),
            Err(e) => {
                log::error!("requesting trace error: {e}");
                return None;
            }
        }
    }
    if (end_block + 1 - start_block) as usize != block_traces.len() {
        log::error!("block_traces.len not expected, batch index = {:#?}", batch_index);
        return None;
    }
    Some(block_traces)
}

#[allow(dead_code)]
fn save_trace(batch_index: u64, batch_traces: &Vec<BlockTrace>) {
    let path = PROVER_PROOF_DIR.to_string() + format!("/batch_{}", batch_index).as_str();
    fs::create_dir_all(path.clone()).unwrap();
    let file = File::create(format!("{}/block_traces.json", path.as_str())).unwrap();
    let writer = BufWriter::new(file);

    serde_json::to_writer_pretty(writer, &batch_traces).unwrap();
    log::info!("batch_traces of batch_index = {:#?} saved", batch_index);
}
