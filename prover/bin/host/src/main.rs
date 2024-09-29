<<<<<<< HEAD
use std::fs::File;
use std::io::BufReader;

use clap::Parser;
use morph_prove::prove;
use sbv_primitives::types::BlockTrace;
=======
use clap::Parser;
use morph_prover::prove;
>>>>>>> main

/// The arguments for the command.
#[derive(Parser, Debug)]
#[clap(author, version, about, long_about = None)]
struct Args {
    #[clap(long)]
    prove: bool,
<<<<<<< HEAD
    #[clap(long, default_value = "../../testdata/mainnet_batch_traces.json")]
    block_path: String,
=======
>>>>>>> main
}

fn main() {
    let args = Args::parse();
<<<<<<< HEAD
    let traces: &mut Vec<Vec<BlockTrace>> = &mut load_trace(&args.block_path);
    let block_traces: &mut Vec<BlockTrace> = &mut traces[0];

    prove(block_traces, args.prove);
}

fn load_trace(file_path: &str) -> Vec<Vec<BlockTrace>> {
    let file = File::open(file_path).unwrap();
    let reader = BufReader::new(file);
    serde_json::from_reader(reader).unwrap()
}



=======
    prove("../../testdata/devnet_batch_traces.json", args.prove);
}
>>>>>>> main
