use std::{fs::File, io::BufReader};

use clap::Parser;
use morph_prove::prove;
use prover_primitives::types::BlockTrace;

/// The arguments for the command.
#[derive(Parser, Debug)]
#[clap(author, version, about, long_about = None)]
struct Args {
    #[clap(long)]
    prove: bool,
    #[clap(long, default_value = "./testdata/altfeetx/trace.json")]
    block_path: String,
}

fn main() {
    dotenv::dotenv().ok();
    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();

    let args = Args::parse();
    let traces: &mut Vec<Vec<BlockTrace>> = &mut load_trace(&args.block_path);
    let block_traces: &mut Vec<BlockTrace> = &mut traces[0];
    println!("block_traces.len: {:?}", block_traces.len());

    let _ = prove(block_traces, args.prove).unwrap();
}

fn load_trace(file_path: &str) -> Vec<Vec<BlockTrace>> {
    let file = File::open(file_path).unwrap();
    let reader = BufReader::new(file);
    serde_json::from_reader(reader).unwrap()
}
