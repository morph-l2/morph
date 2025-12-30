use crate::BatchInfo;
use alloy_provider::{DynProvider, Provider};
use anyhow::{anyhow, Context};
use prover_executor_client::{
    types::input::{BlockInput, ExecutorInput},
    verify,
};
use prover_executor_host::get_blob_info;
use prover_primitives::types::BlockTrace;
use serde::{Deserialize, Serialize};

#[derive(Serialize)]
pub struct ExecuteRequest {
    pub batch_index: u64,
    pub start_block: u64,
    pub end_block: u64,
    pub rpc: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct ExecuteResult {
    pub error_msg: String,
    pub error_code: String,
}

pub async fn try_execute_batch(
    batch: &BatchInfo,
    provider: &DynProvider,
) -> Result<(), anyhow::Error> {
    let block_traces =
        get_block_traces(batch.batch_index, batch.start_block, batch.end_block, provider)
            .await
            .ok_or_else(|| {
                anyhow!("get_block_traces failed for batch index = {:#?}", batch.batch_index)
            })?;

    let blocks_inputs =
        block_traces.iter().map(|trace| BlockInput::from_trace(trace)).collect::<Vec<_>>();
    let client_input =
        ExecutorInput { block_inputs: blocks_inputs, blob_info: get_blob_info(&block_traces)? };
    let result = verify(client_input.clone()).context("native execution failed");

    match result {
        Ok(_) => Ok(()),
        Err(e) => Err(anyhow::Error::from(e)),
    }
}

// Fetches block traces by provider
async fn get_block_traces(
    batch_index: u64,
    start_block: u64,
    end_block: u64,
    provider: &DynProvider,
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
