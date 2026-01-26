use std::{fs::File, io::BufReader};

use alloy_provider::{Provider, ProviderBuilder};
use clap::Parser;
use morph_prove::{execute::execute_batch, prove, utils::command_args::parse_u64_auto_radix};
use prover_executor_client::types::input::ExecutorInput;
use prover_executor_host::{blob::get_blob_info_from_traces, trace::trace_to_input};
use prover_primitives::types::BlockTrace;

/// The arguments for the command.
#[derive(Parser, Debug)]
#[clap(author, version, about, long_about = None)]
struct Args {
    /// Whether to generate proof.
    #[clap(long)]
    prove: bool,
    /// Block trace file path (json).
    #[clap(long, default_value = "./testdata/mpt/mainnet_25215.json")]
    block_path: String,
    /// Whether to use RPC to fetch traces instead of local file.
    #[clap(long)]
    use_rpc_db: bool,
    /// Start L2 block number.
    #[clap(long = "start-block", default_value_t = 1, alias = "start", value_parser = parse_u64_auto_radix)]
    start_block: u64,
    /// End L2 block number.
    #[clap(long = "end-block", default_value_t = 10, alias = "end", value_parser = parse_u64_auto_radix)]
    end_block: u64,
    /// RPC endpoint
    #[clap(long, default_value = "http://localhost:8545")]
    rpc: String,
    /// Whether to save input.
    #[clap(long)]
    save_input: bool,
}

#[tokio::main]
async fn main() {
    dotenv::dotenv().ok();
    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();

    let args = Args::parse();
    let mut input = if args.use_rpc_db {
        // Use RPC to fetch state.
        let provider = ProviderBuilder::new().connect_http(args.rpc.parse().unwrap()).erased();
        execute_batch(1, args.start_block, args.end_block, &provider, true).await.unwrap()
    } else {
        // Use local traces file.
        let block_traces = &mut load_trace(&args.block_path);
        println!("block_traces.len: {:?}", block_traces.len());
        let blocks_inputs =
            block_traces.iter().map(trace_to_input).collect::<Vec<_>>();
        ExecutorInput {
            block_inputs: blocks_inputs,
            blob_info: get_blob_info_from_traces(block_traces).unwrap(),
        }
    };
    if args.save_input {
        let input_path =
            format!("proof/executor_input_{}_{}.json", args.start_block, args.end_block);
        let file = File::create(&input_path).unwrap();
        serde_json::to_writer(file, &input).unwrap();
        println!("Saved executor input to {input_path}");
    }

    let _ = prove(&mut input, args.prove).unwrap();
}

fn load_trace(file_path: &str) -> Vec<BlockTrace> {
    let file = File::open(file_path).unwrap();
    let reader = BufReader::new(file);
    serde_json::from_reader(reader).unwrap()
}
