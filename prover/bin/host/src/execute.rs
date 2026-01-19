use alloy_provider::{Provider, ProviderBuilder};
use prover_executor_client::EVMVerifier;
use prover_executor_host::{
    execute::HostExecutor,
    utils::{assemble_block_input, query_block, HostExecutorOutput},
};

pub async fn execute(block_number: u64, rpc: &str) {
    let provider = ProviderBuilder::new().connect_http(rpc.parse().unwrap()).erased();
    let output: HostExecutorOutput =
        HostExecutor::execute_block(block_number, &provider).await.unwrap();

    let prev_block = query_block(block_number.saturating_sub(1), &provider).await.unwrap();
    let block_input = assemble_block_input(output, prev_block);
    let _batch_info = EVMVerifier::verify(vec![block_input]).unwrap();
}

/// Execute a range of blocks (inclusive).
pub async fn execute_range(start_block: u64, end_block: u64, rpc: &str) {
    assert!(
        end_block >= start_block,
        "end_block ({end_block}) must be >= start_block ({start_block})"
    );
    for block_number in start_block..=end_block {
        execute(block_number, rpc).await;
    }
}

/// Execute blocks continuously starting from `start_block`.
///
/// Note: In tests we bound the execution by `max_blocks` to avoid infinite loops.
pub async fn execute_continuous(start_block: u64, max_blocks: u64, rpc: &str) {
    for offset in 0..max_blocks {
        let block_number = match start_block.checked_add(offset) {
            Some(n) => n,
            None => break,
        };
        execute(block_number, rpc).await;
    }
}

#[cfg(test)]
mod tests {
    use crate::execute::{execute, execute_continuous, execute_range, test_args};
    use prover_executor_host::trace_to_input;
    use prover_primitives::types::BlockTrace;
    use std::{
        fs::{self, File},
        io::BufReader,
        path::{Path, PathBuf},
    };

    // cargo test -p morph-prove --lib -- execute::tests::test_execute --exact --nocapture -- --block-number 0x35 --rpc http://127.0.0.1:9545
    #[test]
    fn test_execute() {
        let rt = tokio::runtime::Runtime::new().unwrap();
        let (block_number, rpc) = test_args::read_execute_args_from_argv();
        rt.block_on(execute(block_number, &rpc));
    }

    // cargo test -p morph-prove --lib -- execute::tests::test_execute_range --exact --nocapture -- --start-block 0x35 --end-block 0x36 --rpc http://127.0.0.1:9545
    #[test]
    fn test_execute_range() {
        let rt = tokio::runtime::Runtime::new().unwrap();
        let (start_block, end_block, rpc) = test_args::read_execute_range_args_from_argv();
        rt.block_on(execute_range(start_block, end_block, &rpc));
    }

    // cargo test -p morph-prove --lib -- execute::tests::test_execute_continuous --exact --nocapture -- --start-block 0x35 --max-blocks 2 --rpc http://127.0.0.1:9545
    #[test]
    fn test_execute_continuous() {
        let rt = tokio::runtime::Runtime::new().unwrap();
        let (start_block, max_blocks, rpc) = test_args::read_execute_continuous_args_from_argv();
        rt.block_on(execute_continuous(start_block, max_blocks, &rpc));
    }

    // Examples:
    //   cargo test -p morph-prove --lib -- execute::tests::test_execute_local_traces --exact --nocapture
    //   cargo test -p morph-prove --lib -- execute::tests::test_execute_local_traces --exact --nocapture -- --trace ../../testdata/mpt/mainnet_19720219.json
    //   cargo test -p morph-prove --lib -- execute::tests::test_execute_local_traces --exact --nocapture -- --trace ../../testdata/mpt
    #[test]
    fn test_execute_local_traces() {
        use prover_executor_client::EVMVerifier;

        let provided = test_args::read_execute_local_traces_paths_from_argv();
        let files = resolve_trace_files(&provided);
        assert!(!files.is_empty(), "no trace files found");

        for file in files {
            let file_str = file.to_string_lossy();
            let block_traces = &mut load_trace(&file_str);

            let block_inputs =
                block_traces.iter().map(|trace| trace_to_input(trace)).collect::<Vec<_>>();

            let _ = EVMVerifier::verify(block_inputs).map_err(|e| {
                println!("execute_local_traces verify error for file {file_str}: {:?}", e);
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

#[cfg(test)]
mod test_args {
    use clap::{ArgAction, Parser};
    const DEFAULT_BLOCK_NUMBER: u64 = 0x477;
    const DEFAULT_START_BLOCK: u64 = DEFAULT_BLOCK_NUMBER;
    const DEFAULT_END_BLOCK: u64 = DEFAULT_BLOCK_NUMBER;
    const DEFAULT_MAX_BLOCKS: u64 = 1000;
    const DEFAULT_RPC: &str = "http://127.0.0.1:9545";

    /// `execute.rs` Test parameters (supports passing in from `cargo test ... -- <args>`).
    #[derive(Parser, Debug)]
    #[command(author, version, about, long_about = None, disable_help_flag = true)]
    struct ExecuteArgs {
        /// L2 block number (supports decimal or 0x prefix hexadecimal).
        #[arg(long = "block-number", alias = "block", default_value_t = DEFAULT_BLOCK_NUMBER, value_parser = parse_u64_auto_radix)]
        block_number: u64,

        /// RPC endpoint。
        #[arg(long, default_value = DEFAULT_RPC)]
        rpc: String,
    }

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

    /// Continuous execute parameters.
    #[derive(Parser, Debug)]
    #[command(author, version, about, long_about = None, disable_help_flag = true)]
    struct ExecuteContinuousArgs {
        /// Start L2 block number.
        #[arg(long = "start-block", alias = "start", default_value_t = DEFAULT_START_BLOCK, value_parser = parse_u64_auto_radix)]
        start_block: u64,

        /// Max blocks to execute (to avoid infinite loop in tests).
        #[arg(long = "max-blocks", default_value_t = DEFAULT_MAX_BLOCKS, value_parser = parse_u64_auto_radix)]
        max_blocks: u64,

        /// RPC endpoint。
        #[arg(long, default_value = DEFAULT_RPC)]
        rpc: String,
    }

    /// Local traces execute parameters.
    #[derive(Parser, Debug)]
    #[command(author, version, about, long_about = None, disable_help_flag = true)]
    struct ExecuteLocalTracesArgs {
        /// Trace file path (json) or directory path. Can be specified multiple times.
        #[arg(long = "trace", alias = "trace-path", value_name = "FILE_OR_DIR", action = ArgAction::Append)]
        traces: Vec<String>,
    }

    pub(super) fn read_execute_args_from_argv() -> (u64, String) {
        let filtered = filter_argv(&["--block-number", "--block", "--rpc"]);

        let args = ExecuteArgs::parse_from(filtered);
        (args.block_number, args.rpc)
    }

    pub(super) fn read_execute_local_traces_paths_from_argv() -> Vec<String> {
        let filtered = filter_argv(&["--trace", "--trace-path"]);
        let args = ExecuteLocalTracesArgs::parse_from(filtered);
        args.traces
    }

    pub(super) fn read_execute_range_args_from_argv() -> (u64, u64, String) {
        let filtered = filter_argv(&["--start-block", "--start", "--end-block", "--end", "--rpc"]);
        let args = ExecuteRangeArgs::parse_from(filtered);
        (args.start_block, args.end_block, args.rpc)
    }

    pub(super) fn read_execute_continuous_args_from_argv() -> (u64, u64, String) {
        let filtered = filter_argv(&["--start-block", "--start", "--max-blocks", "--rpc"]);
        let args = ExecuteContinuousArgs::parse_from(filtered);
        (args.start_block, args.max_blocks, args.rpc)
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

