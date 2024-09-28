use crate::evm::{save_plonk_fixture, EvmProofFixture};
use anyhow::anyhow;
use morph_executor_client::{
    types::input::{AggregationInput, ShardInfo},
    verify_agg,
};
use morph_executor_host::get_blob_info;
use morph_executor_utils::read_env_var;
use sbv_primitives::{alloy_primitives::keccak256, types::BlockTrace, B256};
use sp1_sdk::{HashableKey, ProverClient, SP1Proof, SP1ProofWithPublicValues, SP1Stdin};

use std::time::Instant;

/// The ELF (executable and linkable format) file for the Succinct RISC-V zkVM.
pub const SHARD_VERIFIER_ELF: &[u8] =
    include_bytes!("../../client/shard/elf/riscv32im-succinct-zkvm-elf");

pub const AGG_VERIFIER_ELF: &[u8] =
    include_bytes!("../../client/agg/elf/riscv32im-succinct-zkvm-elf");

pub fn prove(
    shard_proofs: Vec<SP1ProofWithPublicValues>,
    blocks: Vec<Vec<BlockTrace>>,
    prove: bool,
) -> Result<Option<EvmProofFixture>, anyhow::Error> {
    // Setup the logger.
    // sp1_sdk::utils::setup_logger();
    let client = ProverClient::new();
    let program_hash = keccak256(AGG_VERIFIER_ELF);
    println!("Program Hash [view on Explorer]:");
    println!("0x{}", hex::encode(program_hash));

    let mut proofs = Vec::with_capacity(shard_proofs.len());
    let mut shard_infos: Vec<ShardInfo> = Vec::with_capacity(shard_proofs.len());

    for proof in shard_proofs {
        proofs.push(proof.proof.clone());
        shard_infos.push(proof.public_values.clone().read());
    }

    let (_, vkey) = client.setup(SHARD_VERIFIER_ELF);
    let stdin = get_agg_proof_stdin(blocks, shard_infos, proofs, &vkey).unwrap();

    // // Execute the program in native
    // let native_rt =
    //     verify_agg(&stdin).map_err(|e| anyhow!(format!("native execution err: {:?}", e)))?;

    let (mut public_values, execution_report) = client
        .execute(AGG_VERIFIER_ELF, stdin.clone())
        .run()
        .map_err(|e| anyhow!(format!("native execution err: {:?}", e)))?;

    println!(
        "Program executed successfully, Number of cycles: {:?}",
        execution_report.total_instruction_count()
    );
    let vm_rt = public_values.read::<B256>();
    println!("pi_hash generated with sp1-vm execution: {}", hex::encode(vm_rt.as_slice()));
    // assert_eq!(vm_rt, native_rt, "vm_rt == native_rt ");
    // println!("Values are correct!");

    if !prove {
        println!("Execution completed, No prove request, skipping...");
        return Ok(None);
    }
    println!("Start proving...");
    // Setup the program for proving.
    let (pk, vk) = client.setup(AGG_VERIFIER_ELF);
    println!("Batch ELF Verification Key: {:?}", vk.vk.bytes32());

    // Generate the proof
    let start = Instant::now();
    let proof = client.prove(&pk, stdin).plonk().run().expect("proving failed");

    let duration_mins = start.elapsed().as_secs() / 60;
    println!("Successfully generated proof!, time use: {:?} minutes", duration_mins);

    // Verify the proof.
    client.verify(&proof, &vk).expect("failed to verify proof");
    println!("Successfully verified proof!");

    // Deserialize the public values.
    let pi_bytes = proof.public_values.as_slice();

    let fixture = EvmProofFixture {
        vkey: vk.bytes32().to_string(),
        public_values: format!("0x{}", hex::encode(pi_bytes)),
        proof: format!("0x{}", hex::encode(proof.bytes())),
    };

    if read_env_var("SAVE_FIXTURE", false) {
        save_plonk_fixture(&fixture);
    }
    Ok(Some(fixture))
}

/// Get the stdin for the aggregation proof.
fn get_agg_proof_stdin(
    blocks: Vec<Vec<BlockTrace>>,
    shard_infos: Vec<ShardInfo>,
    proofs: Vec<SP1Proof>,
    shard_vkey: &sp1_sdk::SP1VerifyingKey,
) -> Result<SP1Stdin, anyhow::Error> {
    let mut stdin = SP1Stdin::new();
    for proof in proofs {
        let SP1Proof::Compressed(compressed_proof) = proof else {
            panic!();
        };
        stdin.write_proof(*compressed_proof, shard_vkey.vk.clone());
    }

    // Write the aggregation inputs to the stdin.
    stdin.write(&AggregationInput {
        shard_infos,
        shard_vkey: shard_vkey.hash_u32(),
        // TODO
        blob_info: get_blob_info(&blocks[0]).unwrap(),
        l2_traces: blocks,
    });

    Ok(stdin)
}
