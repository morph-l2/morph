use anyhow::anyhow;
use morph_executor_client::{
    types::input::{ClientInput, ShardInfo},
    verify,
};
use sbv_primitives::{alloy_primitives::keccak256, types::BlockTrace};
use sp1_sdk::{HashableKey, ProverClient, SP1ProofWithPublicValues, SP1Stdin};
use std::time::Instant;

/// The ELF (executable and linkable format) file for the Succinct RISC-V zkVM.
pub const SHARD_VERIFIER_ELF: &[u8] =
    include_bytes!("../../client/shard/elf/riscv32im-succinct-zkvm-elf");

const MAX_PROVE_BLOCKS: usize = 4096;

pub fn prove(
    blocks: &mut Vec<BlockTrace>,
    prove: bool,
) -> Result<Option<SP1ProofWithPublicValues>, anyhow::Error> {
    // Setup the logger.
    sp1_sdk::utils::setup_logger();
    let program_hash = keccak256(SHARD_VERIFIER_ELF);
    println!("Program Hash [view on Explorer]:");
    println!("0x{}", hex::encode(program_hash));

    if blocks.len() > MAX_PROVE_BLOCKS {
        return Err(anyhow!(format!(
            "check block_tracs, blocks len = {:?} exceeds MAX_PROVE_BLOCKS = {:?}",
            blocks.len(),
            MAX_PROVE_BLOCKS
        )))
    }

    // Prepare input.
    // Convert the traces' format to reduce conversion costs in the client.
    blocks.iter_mut().for_each(|blobk| blobk.flatten());
    let client_input = ClientInput { l2_traces: blocks.clone() };

    // Execute the program in native
    let native_rt =
        verify(&client_input).map_err(|e| anyhow!(format!("Native execution err: {:?}", e)))?;

    // Execute the program in sp1-vm
    let mut stdin = SP1Stdin::new();
    stdin.write(&serde_json::to_string(&client_input).unwrap());
    let client = ProverClient::new();
    let (mut public_values, execution_report) = client
        .execute(SHARD_VERIFIER_ELF, stdin.clone())
        .run()
        .map_err(|e| anyhow!(format!("Program execution err: {:?}", e)))?;

    println!(
        "Program executed successfully, Number of cycles: {:?}",
        execution_report.total_instruction_count()
    );
    let vm_rt = public_values.read::<ShardInfo>();
    assert_eq!(vm_rt.post_state_root, native_rt.post_state_root, "exec post_state_root verify");
    println!("Values are correct!");

    if !prove {
        println!("Execution completed, No prove request, skipping...");
        return Ok(None);
    }
    println!("Start proving...");
    // Setup the program for proving.
    let (pk, vk) = client.setup(SHARD_VERIFIER_ELF);
    println!("Batch ELF Verification Key: {:?}", vk.vk.bytes32());

    // Generate the proof
    let start = Instant::now();
    let proof = client.prove(&pk, stdin).compressed().run().expect("proving failed");

    let duration_mins = start.elapsed().as_secs() / 60;
    println!(
        "Successfully generated shard compressed proof!, time use: {:?} minutes",
        duration_mins
    );

    // Verify the proof.
    client.verify(&proof, &vk).expect("failed to verify proof");
    println!("Successfully verified proof!");

    Ok(Some(proof))
}
