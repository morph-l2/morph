use alloy_provider::{DynProvider, Provider};
use anyhow::{Context, Result};

const METHOD_GET_BLOCK_TRACE: &str = "morph_getBlockTraceByNumberOrHash";

/// Fetch block trace for a single block.
pub async fn get_block_trace<T: alloy_json_rpc::RpcRecv>(
    block_num: u64,
    provider: &DynProvider,
) -> Result<T> {
    request_block_trace(block_num, provider).await
}

/// Fetch block traces for the inclusive range `[start_block, end_block]`.
///
/// `batch_index` is only used to enrich error context.
pub async fn get_block_traces<T: alloy_json_rpc::RpcRecv>(
    batch_index: u64,
    start_block: u64,
    end_block: u64,
    provider: &DynProvider,
) -> Result<Vec<T>> {
    if end_block < start_block {
        return Err(anyhow::anyhow!(
            "invalid block range for batch {batch_index}: start_block={start_block}, end_block={end_block}"
        ));
    }

    let mut block_traces = Vec::with_capacity(end_block.saturating_sub(start_block) as usize + 1);
    for block_num in start_block..=end_block {
        let trace = request_block_trace(block_num, provider)
            .await
            .with_context(|| format!("batch {batch_index}: fetching block {block_num} trace"))?;
        block_traces.push(trace);
    }
    Ok(block_traces)
}

// Internal helper to request a block trace.
async fn request_block_trace<T: alloy_json_rpc::RpcRecv>(
    block_num: u64,
    provider: &DynProvider,
) -> Result<T> {
    log::debug!("requesting trace of block {block_num}");

    provider
        .raw_request(METHOD_GET_BLOCK_TRACE.into(), [block_num_to_hex_arg(block_num)])
        .await
        .with_context(|| format!("{METHOD_GET_BLOCK_TRACE} failed for block {block_num}"))
}

// Convert a block number to a hex-encoded string argument for JSON-RPC calls.
fn block_num_to_hex_arg(block_num: u64) -> String {
    // JSON-RPC APIs conventionally accept block numbers as hex-encoded quantities.
    format!("{block_num:#x}")
}
