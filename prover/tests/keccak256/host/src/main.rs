use sp1_sdk::{Elf, ProverClient, ProvingKey, SP1Stdin};
use std::time::Instant;

#[tokio::main]
async fn main() {
    // Setup the logger.
    sp1_sdk::utils::setup_logger();

    // The ELF (executable and linkable format) file for the Succinct RISC-V zkVM.
    let dev_elf: &[u8] = include_bytes!("../../client/elf/riscv32im-succinct-zkvm-elf");

    // Setup the prover client.
    let client = ProverClient::from_env().await;

    // Setup the inputs.
    let mut stdin = SP1Stdin::new();

    let data = vec![1, 2];
    stdin.write(&data);

    // Execute the program in sp1-vm
    let (public_values, execution_report) =
        client.execute(Elf::Static(dev_elf), stdin.clone()).await.unwrap();
    println!("Program executed successfully.");
    // Record the number of cycles executed.
    println!("Number of cycles: {}", execution_report.total_instruction_count());

    let rt_data = public_values.as_slice();
    println!("pi_hash generated with sp1-vm execution: {}", hex::encode(rt_data));

    let start = Instant::now();

    // Setup the program for proving.
    let pk = client.setup(Elf::Static(dev_elf)).await.unwrap();

    // Generate the proof
    let proof = client.prove(&pk, stdin).await.expect("failed to generate proof");

    let duration_mins = start.elapsed().as_secs() / 60;
    println!("Successfully generated proof!, time use: {:?} minutes", duration_mins);

    // Verify the proof.
    client.verify(&proof, pk.verifying_key(), None).expect("failed to verify proof");
    println!("Successfully verified proof!");
}
