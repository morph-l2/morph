use morph_prove::prove;
use std::fs::{self, File};
use std::io::{BufReader, BufWriter};
use std::{sync::Arc, thread, time::Duration};

use crate::{read_env_var, PROVER_L2_RPC, PROVER_PROOF_DIR, PROVE_RESULT, PROVE_TIME};
use alloy::providers::Provider;
use alloy::{
    providers::{ProviderBuilder, ReqwestProvider, RootProvider},
    transports::http::reqwest,
};
use anyhow::anyhow;
use sbv_primitives::types::BlockTrace;
use serde::{Deserialize, Serialize};
use tokio::sync::Mutex;

// proveRequest
#[derive(Serialize, Deserialize, Debug)]
pub struct ProveRequest {
    pub batch_index: u64,
    pub start_block: u64,
    pub end_block: u64,
    pub rpc: String,
    pub shadow: Option<bool>,
}

pub struct Prover {
    prove_queue: Arc<Mutex<Vec<ProveRequest>>>,
    provider: ReqwestProvider,
}

impl Prover {
    pub fn new(prove_queue: Arc<Mutex<Vec<ProveRequest>>>) -> Result<Self, anyhow::Error> {
        let url = reqwest::Url::parse(PROVER_L2_RPC.as_str()).map_err(|_| anyhow!("Invalid L2 RPC URL"))?;
        let provider = ProviderBuilder::new().on_provider(RootProvider::new_http(url));

        Ok(Self { prove_queue, provider })
    }

    /// Processes prove requests from a queue.
    pub async fn prove_for_queue(&mut self) {
        println!("Waiting for prove request");
        log::info!("Waiting for prove request");
        loop {
            thread::sleep(Duration::from_millis(12000));

            // Step1. Get request from queue
            let (batch_index, start_block, end_block) = match self.prove_queue.lock().await.pop() {
                Some(req) => {
                    log::info!(
                        "received prove request, batch index = {:#?}, blocks len = {:#?}, start_block = {:#?}, end_block = {:#?}",
                        req.batch_index,
                        req.end_block - req.start_block + 1,
                        req.start_block,
                        req.end_block,
                    );
                    (req.batch_index, req.start_block, req.end_block)
                }
                None => {
                    log::info!("no prove request");
                    continue;
                }
            };

            // let traces: &mut Vec<Vec<BlockTrace>> = &mut load_trace("testdata/mainnet_batch_traces.json");
            // let block_traces: &mut Vec<BlockTrace> = &mut traces[0];

            // Step2. Fetch trace
            log::info!("Requesting trace of batch: {:#?}", batch_index);
            let res_provider = &mut get_block_traces(batch_index, start_block, end_block, &self.provider).await;
            let block_traces = match res_provider {
                Some(block_traces) => block_traces,
                None => {
                    PROVE_RESULT.set(2);
                    continue;
                }
            };

            if read_env_var("SAVE_TRACE", false) {
                save_trace(batch_index, block_traces);
            }

            // Step3. Generate evm proof
            prove(block_traces, true);
        }
    }
}

// Fetches block traces by provider
async fn get_block_traces(
    batch_index: u64,
    start_block: u64,
    end_block: u64,
    provider: &ReqwestProvider,
) -> Option<Vec<BlockTrace>> {
    let mut block_traces: Vec<BlockTrace> = Vec::new();
    for block_num in start_block..end_block + 1 {
        log::debug!("zkevm-prover: requesting trace of block {block_num}");
        let result = provider
            .raw_request("morph_getBlockTraceByNumberOrHash".into(), [format!("{block_num:#x}")])
            .await;

        match result {
            Ok(trace) => block_traces.push(trace),
            Err(e) => {
                log::error!("zkevm-prover: requesting trace error: {e}");
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
fn load_trace(file_path: &str) -> Vec<Vec<BlockTrace>> {
    let file = File::open(file_path).unwrap();
    let reader = BufReader::new(file);
    serde_json::from_reader(reader).unwrap()
}

#[allow(dead_code)]
fn save_trace(batch_index: u64, chunk_traces: &Vec<BlockTrace>) {
    let path = PROVER_PROOF_DIR.to_string() + format!("/batch_{}", batch_index).as_str();
    fs::create_dir_all(path.clone()).unwrap();
    let file = File::create(format!("{}/block_traces.json", path.as_str())).unwrap();
    let writer = BufWriter::new(file);

    serde_json::to_writer_pretty(writer, &chunk_traces).unwrap();
    log::info!("chunk_traces of batch_index = {:#?} saved", batch_index);
}
