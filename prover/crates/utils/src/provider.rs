use alloy_network::Network;
use alloy_provider::{DynProvider, Provider};

pub async fn get_block_trace<T: alloy_json_rpc::RpcRecv>(
    block_num: u64,
    provider: &DynProvider,
) -> Option<T> {
    log::debug!("requesting trace of block {block_num}");
    let result = provider
        .raw_request("morph_getBlockTraceByNumberOrHash".into(), [format!("{block_num:#x}")])
        .await;

    match result {
        Ok(trace) => Some(trace),
        Err(e) => {
            log::error!("requesting trace error: {e}");
            return None;
        }
    }
}

pub async fn get_block_traces<T: alloy_json_rpc::RpcRecv, N: Network>(
    batch_index: u64,
    start_block: u64,
    end_block: u64,
    provider: &DynProvider<N>,
) -> Option<Vec<T>> {
    let mut block_traces: Vec<T> = Vec::new();
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
