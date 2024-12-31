use alloy::primitives::B256;
use poseidon_bn254::{hash_with_domain, Fr, PrimeField};
use sp1_sdk::{ProverClient, SP1Stdin};
use std::time::Instant;
fn main() {
    // Setup the logger.
    sp1_sdk::utils::setup_logger();

    let native_hash = cacl_poseidon_hash();
    println!("native_hash: {}", native_hash);
    println!("hash generated with native execution: {}", hex::encode(native_hash));

    // The ELF (executable and linkable format) file for the Succinct RISC-V zkVM.
    let dev_elf: &[u8] = include_bytes!("../../client/elf/riscv32im-succinct-zkvm-elf");

    // Setup the prover client.
    let client = ProverClient::new();

    // Setup the inputs.
    let mut stdin = SP1Stdin::new();

    let data = vec![1, 2];
    stdin.write(&data);

    // Execute the program in sp1-vm
    let (mut public_values, execution_report) =
        client.execute(dev_elf, stdin.clone()).run().unwrap();
    println!("Program executed successfully.");
    // Record the number of cycles executed.
    println!("Number of cycles: {}", execution_report.total_instruction_count());
    println!("public_values: {:?}", public_values);

    let rt_data = public_values.read::<[u8; 32]>();
    // let rt_data: [u8; 32] = public_values.to_vec()[8..40].try_into().unwrap();
    // println!("rt_data: {:?}", rt_data);

    println!("hash generated with sp1-vm execution: {}", hex::encode(B256::from_slice(&rt_data)));

    // let start = Instant::now();

    // // Setup the program for proving.
    // let (pk, vk) = client.setup(dev_elf);

    // // Generate the proof
    // let proof = client.prove(&pk, stdin).run().expect("failed to generate proof");

    // let duration_mins = start.elapsed().as_secs() / 60;
    // println!("Successfully generated proof!, time use: {:?} minutes", duration_mins);

    // // Verify the proof.
    // client.verify(&proof, &vk).expect("failed to verify proof");
    // println!("Successfully verified proof!");
}

fn cacl_poseidon_hash() -> B256 {
    let mut hash_bytes =
        hash_with_domain(&[Fr::from(1u64), Fr::from(2u64)], Fr::from(3u64)).to_repr();
    hash_bytes.reverse();
    B256::from_slice(&hash_bytes)
}
