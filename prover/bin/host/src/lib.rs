use std::{fs::File, time::Instant};

use clap::Parser;
use morph_executor_client::{types::input::ClientInput, verify};
use morph_executor_host::get_blob_info;
use sbv_primitives::{types::BlockTrace, Block, B256};
use sp1_sdk::{ProverClient, SP1Stdin};
use std::io::BufReader;

/// The ELF (executable and linkable format) file for the Succinct RISC-V zkVM.
pub const VERIFIER_ELF: &[u8] = include_bytes!("../../client/elf/riscv32im-succinct-zkvm-elf");

/// The arguments for the command.
#[derive(Parser, Debug)]
#[clap(author, version, about, long_about = None)]
struct Args {
    #[clap(long)]
    prove: bool,
}

fn load_trace(file_path: &str) -> Vec<Vec<BlockTrace>> {
    let file = File::open(file_path).unwrap();
    let reader = BufReader::new(file);
    serde_json::from_reader(reader).unwrap()
}

pub fn prove_for_queue() {}

pub fn prove(trace_path: &str) {
    // Setup the logger.
    sp1_sdk::utils::setup_logger();
    let args = Args::parse();
    let mut traces: Vec<Vec<BlockTrace>> = load_trace(trace_path);
    let block_trace: &mut BlockTrace = &mut traces[0][1];
    println!(
        "traces' post_state_root: {:?}, transactions.len: {:?}",
        block_trace.root_after(),
        block_trace.transactions.len()
    );

    // Convert the traces' format to reduce conversion costs in the client.
    block_trace.flatten();
    let client_input = ClientInput {
        l2_trace: block_trace.clone(),
        blob_info: get_blob_info(block_trace).unwrap(),
    };

    // Execute the program in native
    let expected_hash = verify(&client_input).unwrap();
    println!(
        "pi_hash generated with native execution: {}",
        hex::encode(expected_hash.as_slice())
    );

    let mut stdin = SP1Stdin::new();
    let trace_str = serde_json::to_string(&client_input).unwrap();
    stdin.write(&trace_str);

    // Execute the program in sp1-vm
    let client = ProverClient::new();
    let (mut public_values, execution_report) =
        client.execute(VERIFIER_ELF, stdin.clone()).run().unwrap();
    println!(
        "Program executed successfully, Number of cycles: {:?}",
        execution_report.total_instruction_count()
    );
    let pi_hash = public_values.read::<B256>();
    println!(
        "pi_hash generated with sp1-vm execution: {}",
        hex::encode(pi_hash.as_slice())
    );
    assert_eq!(pi_hash, expected_hash);
    println!("Values are correct!");

    if args.prove {
        let start = Instant::now();

        // Setup the program for proving.
        let (pk, vk) = client.setup(VERIFIER_ELF);

        // Generate the proof
        let proof = client
            .prove(&pk, stdin)
            .run()
            .expect("failed to generate proof");

        let duration_mins = start.elapsed().as_secs() / 60;
        println!(
            "Successfully generated proof!, time use: {:?} minutes",
            duration_mins
        );

        // Verify the proof.
        client.verify(&proof, &vk).expect("failed to verify proof");
        println!("Successfully verified proof!");
    }
}
