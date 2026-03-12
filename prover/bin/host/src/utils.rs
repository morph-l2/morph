pub mod command_args {
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

    pub fn read_execute_args_from_argv() -> (u64, String) {
        let filtered = filter_argv(&["--block-number", "--block", "--rpc"]);

        let args = ExecuteArgs::parse_from(filtered);
        (args.block_number, args.rpc)
    }

    pub fn read_execute_local_traces_paths_from_argv() -> Vec<String> {
        let filtered = filter_argv(&["--trace", "--trace-path"]);
        let args = ExecuteLocalTracesArgs::parse_from(filtered);
        args.traces
    }

    pub fn read_execute_range_args_from_argv() -> (u64, u64, String) {
        let filtered = filter_argv(&["--start-block", "--start", "--end-block", "--end", "--rpc"]);
        let args = ExecuteRangeArgs::parse_from(filtered);
        (args.start_block, args.end_block, args.rpc)
    }

    pub fn read_execute_continuous_args_from_argv() -> (u64, u64, String) {
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

    pub fn parse_u64_auto_radix(s: &str) -> Result<u64, String> {
        let s = s.trim();
        if let Some(hex) = s.strip_prefix("0x").or_else(|| s.strip_prefix("0X")) {
            u64::from_str_radix(hex, 16).map_err(|e| e.to_string())
        } else {
            s.parse::<u64>().map_err(|e| e.to_string())
        }
    }
}
