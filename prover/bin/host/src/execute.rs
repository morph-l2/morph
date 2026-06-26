use alloy_provider::DynProvider;
use prover_executor_client::{types::input::ExecutorInput, EVMVerifier};
use prover_executor_host::{
    blob::get_blob_infos_from_blocks,
    execute::HostExecutor,
    utils::{assemble_block_input, HostExecutorOutput},
    ClientBlockInput,
};

/// Data-source strategy for building [`ClientBlockInput`].
#[derive(Debug, Clone, Copy, PartialEq, Eq, Default)]
pub enum InputSource {
    /// Fetch state via per-account `eth_getProof` calls (original behaviour).
    #[default]
    RpcDb,
    /// Fetch state via a single `debug_executionWitness` call (reth).
    Witness,
}

/// Execute a single block using per-account `eth_getProof` (original RPC-DB path).
pub async fn execute(
    block_number: u64,
    provider: &DynProvider,
) -> Result<ClientBlockInput, anyhow::Error> {
    let output: HostExecutorOutput = HostExecutor::execute_block(block_number, provider).await?;

    // let prev_block = query_block(block_number.saturating_sub(1), provider).await?;
    let block_input = assemble_block_input(output);
    let _ = EVMVerifier::verify(vec![block_input.clone()])?;
    Ok(block_input)
}

/// Execute a single block using a single `debug_executionWitness` RPC call.
///
/// This is more efficient than [`execute`] for nodes that support the
/// `debug_executionWitness` endpoint (e.g. reth, recent geth).
pub async fn execute_with_witness(
    block_number: u64,
    provider: &DynProvider,
) -> Result<ClientBlockInput, anyhow::Error> {
    let output: HostExecutorOutput =
        HostExecutor::execute_block_with_witness(block_number, provider).await?;

    // let prev_block = query_block(block_number.saturating_sub(1), provider).await?;
    let block_input = assemble_block_input(output);
    let _ = EVMVerifier::verify(vec![block_input.clone()])?;
    Ok(block_input)
}

/// Execute a batch of blocks (inclusive).
pub async fn execute_batch(
    batch_index: u64,
    start_block: u64,
    end_block: u64,
    provider: &DynProvider,
    source: InputSource,
    batch_version: u8,
) -> Result<ExecutorInput, anyhow::Error> {
    assert!(
        end_block >= start_block,
        "end_block ({end_block}) must be >= start_block ({start_block})"
    );
    log::info!(
        "Executing batch {}, blocks {} to {}, source = {:?}",
        batch_index,
        start_block,
        end_block,
        source,
    );
    let executor_input = match source {
        InputSource::RpcDb => {
            // Use per-account eth_getProof RPC calls.
            let mut block_inputs = vec![];
            for block_number in start_block..=end_block {
                block_inputs.push(execute(block_number, provider).await?);
            }
            ExecutorInput {
                block_inputs: block_inputs.clone(),
                blob_infos: get_blob_infos_from_blocks(
                    &block_inputs
                        .iter()
                        .map(|input| input.current_block.clone())
                        .collect::<Vec<_>>(),
                )?,
                batch_version,
            }
        }
        InputSource::Witness => {
            // Use a single debug_executionWitness call per block (reth / geth).
            let mut block_inputs = vec![];
            for block_number in start_block..=end_block {
                block_inputs.push(execute_with_witness(block_number, provider).await?);
            }
            ExecutorInput {
                block_inputs: block_inputs.clone(),
                blob_infos: get_blob_infos_from_blocks(
                    &block_inputs
                        .iter()
                        .map(|input| input.current_block.clone())
                        .collect::<Vec<_>>(),
                )?,
                batch_version,
            }
        }
    };

    Ok(executor_input)
}

/// Execute a range of blocks (inclusive).
pub async fn execute_range(start_block: u64, end_block: u64, provider: &DynProvider) {
    assert!(
        end_block >= start_block,
        "end_block ({end_block}) must be >= start_block ({start_block})"
    );
    for block_number in start_block..=end_block {
        execute(block_number, provider).await.unwrap();
    }
}

/// Execute blocks continuously starting from `start_block`.
///
/// Note: In tests we bound the execution by `max_blocks` to avoid infinite loops.
pub async fn execute_continuous(start_block: u64, max_blocks: u64, provider: &DynProvider) {
    for offset in 0..max_blocks {
        let block_number = match start_block.checked_add(offset) {
            Some(n) => n,
            None => break,
        };
        execute(block_number, provider).await.unwrap();
    }
}

#[cfg(test)]
mod tests {
    use crate::{
        execute::{execute, execute_continuous, execute_range},
        utils::command_args,
    };

    use alloy_provider::{Provider, ProviderBuilder};
    use prover_executor_host::trace::trace_to_input;
    use prover_primitives::types::BlockTrace;
    use std::{
        fs::{self, File},
        io::BufReader,
        path::{Path, PathBuf},
    };

    // cargo test -p morph-prove --lib -- execute::tests::test_execute --exact --nocapture -- --block-number 19997 --rpc http://127.0.0.1:9545
    #[test]
    fn test_execute() {
        env_logger::Builder::new().filter_level(log::LevelFilter::Info).format_target(false).init();

        let rt = tokio::runtime::Runtime::new().unwrap();
        let (block_number, rpc) = command_args::read_execute_args_from_argv();
        let provider = ProviderBuilder::new().connect_http(rpc.parse().unwrap()).erased();

        rt.block_on(execute(block_number, &provider)).unwrap();
    }

    // cargo test -p morph-prove --lib -- execute::tests::test_execute_range --exact --nocapture -- --start-block 0x35 --end-block 0x36 --rpc http://127.0.0.1:9545
    #[test]
    fn test_execute_range() {
        env_logger::Builder::new().filter_level(log::LevelFilter::Info).format_target(false).init();

        let rt = tokio::runtime::Runtime::new().unwrap();
        let (start_block, end_block, rpc) = command_args::read_execute_range_args_from_argv();
        let provider = ProviderBuilder::new().connect_http(rpc.parse().unwrap()).erased();
        rt.block_on(execute_range(start_block, end_block, &provider));
    }

    // cargo test -p morph-prove --lib -- execute::tests::test_execute_continuous --exact --nocapture -- --start-block 0x35 --max-blocks 2 --rpc http://127.0.0.1:9545
    #[test]
    fn test_execute_continuous() {
        env_logger::Builder::new().filter_level(log::LevelFilter::Info).format_target(false).init();

        let rt = tokio::runtime::Runtime::new().unwrap();
        let (start_block, max_blocks, rpc) = command_args::read_execute_continuous_args_from_argv();
        let provider = ProviderBuilder::new().connect_http(rpc.parse().unwrap()).erased();
        rt.block_on(execute_continuous(start_block, max_blocks, &provider));
    }

    // Examples:
    //   cargo test -p morph-prove --lib -- execute::tests::test_execute_local_traces --exact --nocapture
    //   cargo test -p morph-prove --lib -- execute::tests::test_execute_local_traces --exact --nocapture -- --trace ../../testdata/mpt/mainnet_19720219.json
    //   cargo test -p morph-prove --lib -- execute::tests::test_execute_local_traces --exact --nocapture -- --trace ../../testdata/mpt
    #[test]
    fn test_execute_local_traces() {
        use prover_executor_client::EVMVerifier;
        env_logger::Builder::new()
            .filter_level(log::LevelFilter::Debug)
            .format_target(false)
            .init();

        let provided = command_args::read_execute_local_traces_paths_from_argv();
        let files = resolve_trace_files(&provided);
        assert!(!files.is_empty(), "no trace files found");

        for file in files {
            let file_str = file.to_string_lossy();
            let block_traces = &mut load_trace(&file_str);

            let block_inputs = block_traces.iter().map(trace_to_input).collect::<Vec<_>>();

            let _ = EVMVerifier::verify(block_inputs).map_err(|e| {
                println!("execute_local_traces verify error for file {file_str}: {e:?}");
            });
        }
    }

    fn resolve_trace_files(paths: &[String]) -> Vec<PathBuf> {
        // Default: run all *.json under testdata/mpt/
        if paths.is_empty() {
            let dir = default_mpt_trace_dir();
            return list_json_files(&dir);
        }

        let mut out = Vec::new();
        for p in paths {
            let pb = PathBuf::from(p);
            if pb.is_dir() {
                out.extend(list_json_files(&pb));
            } else {
                out.push(pb);
            }
        }

        out.sort();
        out
    }

    fn default_mpt_trace_dir() -> PathBuf {
        // bin/host (manifest dir) -> repo_root/testdata/mpt
        PathBuf::from(env!("CARGO_MANIFEST_DIR")).join("../../testdata/mpt")
    }

    fn list_json_files(dir: &Path) -> Vec<PathBuf> {
        let mut files = Vec::new();
        if let Ok(rd) = fs::read_dir(dir) {
            for entry in rd.flatten() {
                let path = entry.path();
                if path
                    .extension()
                    .and_then(|e| e.to_str())
                    .is_some_and(|e| e.eq_ignore_ascii_case("json"))
                {
                    files.push(path);
                }
            }
        }
        files.sort();
        files
    }

    fn load_trace(file_path: &str) -> Vec<BlockTrace> {
        let file = File::open(file_path).unwrap();
        let reader = BufReader::new(file);
        serde_json::from_reader(reader).unwrap()
    }
}
