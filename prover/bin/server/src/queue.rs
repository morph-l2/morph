use std::{sync::Arc, thread, time::Duration};

use crate::{PROVER_L2_RPC, PROVE_RESULT, PROVE_TIME};
use alloy::providers::Provider;
use alloy::{
    providers::{ProviderBuilder, ReqwestProvider, RootProvider},
    transports::http::reqwest,
};
use anyhow::anyhow;
use morph_prove::prove;
use sbv_primitives::types::BlockTrace;
use serde::{Deserialize, Serialize};
use tokio::sync::Mutex;

// proveRequest
#[derive(Serialize, Deserialize, Debug)]
pub struct ProveRequest {
    pub batch_index: u64,
    pub blocks: Vec<u64>,
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
        log::info!("Waiting for prove request");
        loop {
            thread::sleep(Duration::from_millis(12000));

            // Step1. Get request from queue
            let (batch_index, blocks) = match self.prove_queue.lock().await.pop() {
                Some(req) => {
                    log::info!(
                        "received prove request, batch index = {:#?}, blocks len = {:#?}",
                        req.batch_index,
                        req.blocks.len()
                    );
                    log::debug!(">>blocks details = {:#?}", req.blocks);
                    (req.batch_index, req.blocks.clone())
                }
                None => {
                    log::info!("no prove request");
                    continue;
                }
            };

            // Step2. Fetch trace
            log::info!("Requesting trace of batch: {:#?}", batch_index);
            let res_provider = &mut get_block_traces(batch_index, blocks, &self.provider).await;
            let block_traces = match res_provider {
                Some(block_traces) => block_traces,
                None => {
                    PROVE_RESULT.set(2);
                    continue;
                }
            };

            // Step3. Generate evm proof
            prove(block_traces, true);
        }
    }
}

// Fetches block traces by provider
async fn get_block_traces(batch_index: u64, blocks: Vec<u64>, provider: &ReqwestProvider) -> Option<Vec<BlockTrace>> {
    let mut block_traces: Vec<BlockTrace> = Vec::new();
    for block_num in &blocks {
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
    if blocks.len() != block_traces.len() {
        log::error!("block_traces.len not expect, batch index = {:#?}", batch_index);
        return None;
    }
    Some(block_traces)
}
