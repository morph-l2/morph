use std::{fs::File, io::BufReader};

use clap::Parser;
use morph_prove::prove;
use sbv_primitives::types::BlockTrace;

/// The arguments for the command.
#[derive(Parser, Debug)]
#[clap(author, version, about, long_about = None)]
struct Args {
    #[clap(long)]
    prove: bool,
    #[clap(long, default_value = "../../testdata/mainnet_batch_traces.json")]
    block_path: String,
}

fn main() {
    let args = Args::parse();
    let traces: &mut Vec<Vec<BlockTrace>> = &mut load_trace(&args.block_path);
    let block_traces: &mut Vec<BlockTrace> = &mut traces[0];

    let _ = prove(block_traces, args.prove);
}

fn load_trace(file_path: &str) -> Vec<Vec<BlockTrace>> {
    let file = File::open(file_path).unwrap();
    let reader = BufReader::new(file);
    serde_json::from_reader(reader).unwrap()
}
