use crate::{BatchInfo, SHADOW_EXECUTE_WITH_WITNESS};
use alloy_provider::{DynProvider, Provider};
use anyhow::{anyhow, Context};
use prover_executor_client::{
    types::input::{BlockInput, ExecutorInput},
    verify,
};
use prover_executor_host::{
    blob::get_blob_info_from_blocks,
    blob::get_blob_info_from_traces,
    execute::HostExecutor,
    utils::{assemble_block_input, query_block, HostExecutorOutput},
};
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
        get_block_traces(batch.batch_index, batch.start_block, batch.end_block, &provider.clone())
            .await
            .ok_or_else(|| {
                anyhow!("get_block_traces failed for batch index = {:#?}", batch.batch_index)
            })?;

    let client_input = if *SHADOW_EXECUTE_WITH_WITNESS {
        let blocks_inputs =
            block_traces.iter().map(|trace| BlockInput::from_trace(trace)).collect::<Vec<_>>();
        ExecutorInput {
            block_inputs: blocks_inputs,
            blob_info: get_blob_info_from_traces(&block_traces)?,
        }
    } else {
        let start_block = batch.start_block;
        let end_block = batch.end_block;
        let provider = provider.clone();
        
        let runtime = tokio::runtime::Builder::new_current_thread()
            .enable_all()
            .build()
            .context("Failed to build tokio runtime for shadow exec host")?;
        let blocks_inputs =
            runtime.block_on(async { execute_host_range(start_block, end_block, &provider).await });

        ExecutorInput {
            block_inputs: blocks_inputs.clone(),
            blob_info: get_blob_info_from_blocks(
                &blocks_inputs.iter().map(|input| input.current_block.clone()).collect::<Vec<_>>(),
            )?,
        }
    };

    let result = verify(client_input.clone()).context("native execution failed");
    match result {
        Ok(_) => Ok(()),
        Err(e) => Err(anyhow::Error::from(e)),
    }
}

/// Execute a range of blocks (inclusive).
pub async fn execute_host_range(
    start_block: u64,
    end_block: u64,
    provider: &DynProvider,
) -> Vec<BlockInput> {
    assert!(
        end_block >= start_block,
        "end_block ({end_block}) must be >= start_block ({start_block})"
    );
    let mut block_inputs = Vec::new();
    for block_number in start_block..=end_block {
        let block_input = execute_host(block_number, provider).await;
        block_inputs.push(block_input);
    }
    block_inputs
}

pub async fn execute_host(block_number: u64, provider: &DynProvider) -> BlockInput {
    let prev_block = query_block(block_number.saturating_sub(1), &provider).await.unwrap();
    let output: HostExecutorOutput =
        HostExecutor::execute_block(block_number, &provider).await.unwrap();

    assemble_block_input(output, prev_block)
    // let _batch_info = EVMVerifier::verify(vec![block_input]).unwrap();
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

#[cfg(test)]
mod tests {
    use alloy_provider::{Provider, ProviderBuilder};
    use prover_executor_client::{types::input::BlockInput, EVMVerifier};
    use prover_primitives::types::BlockTrace;
    use prover_utils::provider::get_block_trace;

    use crate::{
        execute::{execute_host_range, test_args, try_execute_batch},
        BatchInfo,
    };

    // cargo test -p shadow-proving --lib -- execute::tests::test_execute_range --exact --nocapture -- --start-block 0x35 --end-block 0x36 --rpc http://127.0.0.1:9545
    #[test]
    fn test_execute_range() {
        let (start_block, end_block, rpc) = test_args::read_execute_range_args_from_argv();
        let provider = ProviderBuilder::new().connect_http(rpc.parse().unwrap()).erased();

        let rt = tokio::runtime::Runtime::new().unwrap();
        let block_inputs = rt.block_on(execute_host_range(start_block, end_block, &provider));
        let _batch_info = EVMVerifier::verify(block_inputs).unwrap();
    }

    #[tokio::test]
    async fn test_execute_batch() {
        let handle = tokio::spawn(async move {
            let provider = ProviderBuilder::new()
                .connect_http("http://127.0.0.1:9545".parse().unwrap())
                .erased();

            try_execute_batch(
                &BatchInfo { batch_index: 1, start_block: 53, end_block: 54, total_txn: 1 },
                &provider,
            )
            .await
            .unwrap();
        });

        handle.await.unwrap();
    }
    #[tokio::test]
    async fn test_execute_remote() {
        let provider =
            ProviderBuilder::new().connect_http("http://127.0.0.1:9545".parse().unwrap()).erased();

        let block_trace = get_block_trace::<BlockTrace>(53, &provider).await.unwrap();
        println!("loaded block_{} traces", block_trace.header.number);
        let block_input: BlockInput = BlockInput::from_trace(&block_trace);

        let batch_info = EVMVerifier::verify(vec![block_input]).unwrap();
        println!("batch_info.post_state_root: {:?}", batch_info.post_state_root);
    }
}

#[cfg(test)]
mod test_args {
    use clap::Parser;

    const DEFAULT_BLOCK_NUMBER: u64 = 0x477;
    const DEFAULT_START_BLOCK: u64 = DEFAULT_BLOCK_NUMBER;
    const DEFAULT_END_BLOCK: u64 = DEFAULT_BLOCK_NUMBER;
    const DEFAULT_RPC: &str = "http://127.0.0.1:9545";

    /// Range execute parameters.
    #[derive(Parser, Debug)]
    #[command(author, version, about, long_about = None, disable_help_flag = true)]
    struct ExecuteRangeArgs {
        /// Start L2 block number (inclusive).
        #[arg(long = "start-block", alias = "start", default_value_t = DEFAULT_START_BLOCK, value_parser = parse_u64_auto_radix)]
        start_block: u64,

        /// End L2 block number (inclusive).
        #[arg(long = "end-block", alias = "end", default_value_t = DEFAULT_END_BLOCK, value_parser = parse_u64_auto_radix)]
        end_block: u64,

        /// RPC endpoint。
        #[arg(long, default_value = DEFAULT_RPC)]
        rpc: String,
    }

    pub(super) fn read_execute_range_args_from_argv() -> (u64, u64, String) {
        let filtered = filter_argv(&["--start-block", "--start", "--end-block", "--end", "--rpc"]);
        let args = ExecuteRangeArgs::parse_from(filtered);
        (args.start_block, args.end_block, args.rpc)
    }

    fn filter_argv(allowed_flags: &[&str]) -> Vec<String> {
        let argv: Vec<String> = std::env::args().skip(1).collect();
        let mut filtered: Vec<String> = Vec::with_capacity(argv.len() + 1);
        // clap expects argv[0] to be the binary name, so we use a placeholder.
        filtered.push("execute_test".to_string());

        let mut it = argv.into_iter();
        while let Some(arg) = it.next() {
            if allowed_flags.iter().any(|f| *f == arg) {
                filtered.push(arg);
                if let Some(v) = it.next() {
                    filtered.push(v);
                }
            } else {
                // ignore unknown args
            }
        }

        filtered
    }

    fn parse_u64_auto_radix(s: &str) -> Result<u64, String> {
        let s = s.trim();
        if let Some(hex) = s.strip_prefix("0x").or_else(|| s.strip_prefix("0X")) {
            u64::from_str_radix(hex, 16).map_err(|e| e.to_string())
        } else {
            s.parse::<u64>().map_err(|e| e.to_string())
        }
    }
}
