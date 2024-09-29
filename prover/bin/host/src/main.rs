use clap::Parser;
use morph_prover::prove;

/// The arguments for the command.
#[derive(Parser, Debug)]
#[clap(author, version, about, long_about = None)]
struct Args {
    #[clap(long)]
    prove: bool,
}

fn main() {
    let args = Args::parse();
    prove("../../testdata/devnet_batch_traces.json", args.prove);
}
