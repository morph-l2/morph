use morph_executor_client::{types::input::ClientInput, verify};
use morph_executor_host::get_blob_info;
use sbv_primitives::alloy_primitives::keccak256;
use sbv_primitives::{types::BlockTrace, B256};
use sp1_sdk::{HashableKey, ProverClient, SP1Stdin};
use std::time::Instant;

/// The ELF (executable and linkable format) file for the Succinct RISC-V zkVM.
pub const BATCH_VERIFIER_ELF: &[u8] = include_bytes!("../../client/elf/riscv32im-succinct-zkvm-elf");

pub fn prove_for_queue() {}

pub fn prove(blocks: &mut Vec<BlockTrace>, prove: bool) {
    // Setup the logger.
    sp1_sdk::utils::setup_logger();
    let program_hash = keccak256(BATCH_VERIFIER_ELF);
    println!("Program Hash [view on Explorer]:");
    println!("0x{}", hex::encode(program_hash));

    // Prepare input.
    // Convert the traces' format to reduce conversion costs in the client.
    blocks.iter_mut().for_each(|blobk| blobk.flatten());

    let client_input = ClientInput {
        l2_traces: blocks.clone(),
        blob_info: get_blob_info(blocks).unwrap(),
    };

    // Execute the program in native
    let expected_hash = verify(&client_input).unwrap();
    println!(
        "pi_hash generated with native execution: {}",
        hex::encode(expected_hash.as_slice())
    );

    // Execute the program in sp1-vm
    let mut stdin = SP1Stdin::new();
    stdin.write(&serde_json::to_string(&client_input).unwrap());
    let client = ProverClient::new();
    let (mut public_values, execution_report) = client.execute(BATCH_VERIFIER_ELF, stdin.clone()).run().unwrap();

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

    if !prove {
        println!("Execution completed, No prove request, skipping...");
        return;
    }
    println!("Start proving...");
    // Setup the program for proving.
    let (pk, vk) = client.setup(BATCH_VERIFIER_ELF);
    println!("Batch ELF Verification Key: {:?}", vk.vk.bytes32());

    // Generate the proof
    let start = Instant::now();
    let proof = client.prove(&pk, stdin).plonk().run().expect("proving failed");

    let duration_mins = start.elapsed().as_secs() / 60;
    println!("Successfully generated proof!, time use: {:?} minutes", duration_mins);

    // Verify the proof.
    client.verify(&proof, &vk).expect("failed to verify proof");
    println!("Successfully verified proof!");
}
