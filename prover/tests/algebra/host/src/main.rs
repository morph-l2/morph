use sp1_sdk::{ProverClient, SP1Stdin};
use std::time::Instant;

fn main() {
    // Setup the logger.
    sp1_sdk::utils::setup_logger();

    // The ELF (executable and linkable format) file for the Succinct RISC-V zkVM.
    let dev_elf: &[u8] = include_bytes!("../../client/elf/riscv32im-succinct-zkvm-elf");

    // Setup the prover client.
    let client = ProverClient::new();

    // Setup the inputs.
    let mut stdin = SP1Stdin::new();

    let data = 123u32;
    stdin.write(&data);

    // Execute the program in sp1-vm
    let (public_values, execution_report) = client.execute(dev_elf, stdin.clone()).run().unwrap();
    println!("Program executed successfully.");
    // Record the number of cycles executed.
    println!(
        "Number of cycles: {}",
        execution_report.total_instruction_count()
    );

    println!(
        "pi_hash generated with sp1-vm execution: {:?}",
        public_values
    );

    println!("Start proving...");
    let start = Instant::now();

    // Setup the program for proving.
    let (pk, vk) = client.setup(dev_elf);

    // Generate the proof
    let proof = client
        .prove(&pk, stdin)
        .plonk()
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
