use anyhow::{bail, Context};
pub mod evm;
pub mod execute;
use evm::{save_plonk_fixture, EvmProofFixture};
use prover_executor_client::{
    types::input::{BlockInput, ExecutorInput},
    verify,
};
use prover_executor_host::blob::get_blob_info_from_traces;
use prover_primitives::{alloy_primitives::keccak256, types::BlockTrace, B256};
use prover_utils::read_env_var;
use sp1_sdk::{HashableKey, ProverClient, SP1Stdin};
use std::time::Instant;

/// The ELF (executable and linkable format) file for the Succinct RISC-V zkVM.
pub const BATCH_VERIFIER_ELF: &[u8] =
    include_bytes!("../../client/elf/riscv32im-succinct-zkvm-elf");

const MAX_PROVE_BLOCKS: usize = 4096;

pub fn prove(
    blocks: &mut Vec<BlockTrace>,
    prove: bool,
) -> Result<Option<EvmProofFixture>, anyhow::Error> {
    let program_hash = keccak256(BATCH_VERIFIER_ELF);
    log::info!("Program Hash [view on Explorer]:");
    log::info!("{}", alloy::hex::encode_prefixed(program_hash));

    if blocks.len() > MAX_PROVE_BLOCKS {
        bail!(
            "check block_traces, blocks len = {} exceeds MAX_PROVE_BLOCKS = {}",
            blocks.len(),
            MAX_PROVE_BLOCKS
        );
    }

    // Prepare input.
    // Convert the traces' format to reduce conversion costs in the client.
    let (client_input, expected_hash) = execute_batch(blocks)?;
    log::info!(
        "pi_hash generated with native execution: {}",
        alloy::hex::encode_prefixed(expected_hash.as_slice())
    );

    // Execute the program in sp1-vm
    let mut stdin = SP1Stdin::new();
    stdin.write(&serde_json::to_string(&client_input)?);
    let client = ProverClient::from_env();

    if read_env_var("DEVNET", false) {
        let (mut public_values, execution_report) =
            client.execute(BATCH_VERIFIER_ELF, &stdin).run().context("sp1-vm execution failed")?;

        log::info!(
            "Program executed successfully, Number of cycles: {}",
            execution_report.total_instruction_count()
        );
        let pi_hash = public_values.read::<[u8; 32]>();
        let public_values = B256::from_slice(&pi_hash);

        log::info!(
            "pi_hash generated with sp1-vm execution: {}",
            alloy::hex::encode_prefixed(public_values.as_slice())
        );
        assert_eq!(pi_hash, expected_hash, "pi_hash mismatch with expected hash");
        log::info!("Values are correct!");
    }

    if !prove {
        log::info!("Execution completed, No prove request, skipping...");
        return Ok(None);
    }

    log::info!("Start proving...");
    let (pk, vk) = client.setup(BATCH_VERIFIER_ELF);
    log::info!("Batch ELF Verification Key: {:?}", vk.vk.bytes32());

    // Generate the proof
    let start = Instant::now();
    let mut proof = client.prove(&pk, &stdin).core().run().context("proving failed")?;

    let duration_mins = start.elapsed().as_secs() / 60;
    log::info!("Successfully generated proof!, time use: {} minutes", duration_mins);

    // Verify the proof.
    client.verify(&proof, &vk).context("failed to verify proof")?;
    log::info!("Successfully verified proof!");

    // Deserialize the public values.
    let pi_bytes = proof.public_values.read::<[u8; 32]>();
    log::info!("pi_hash generated with sp1-vm prove: {}", alloy::hex::encode_prefixed(pi_bytes));

    let fixture = EvmProofFixture {
        vkey: vk.bytes32().to_string(),
        public_values: B256::from_slice(&pi_bytes).to_string(),
        proof: alloy::hex::encode_prefixed(proof.bytes()),
    };

    if read_env_var("SAVE_FIXTURE", false) {
        save_plonk_fixture(&fixture);
    }
    Ok(Some(fixture))
}

pub fn execute_batch(
    blocks: &mut Vec<BlockTrace>,
) -> Result<(ExecutorInput, alloy::primitives::FixedBytes<32>), anyhow::Error> {
    let blocks_inputs =
        blocks.iter().map(|trace| BlockInput::from_trace(trace)).collect::<Vec<_>>();
    let client_input = ExecutorInput {
        block_inputs: blocks_inputs,
        blob_info: get_blob_info_from_traces(blocks)?,
    };
    let expected_hash = verify(client_input.clone()).context("native execution failed")?;
    Ok((client_input, expected_hash))
}

#[cfg(test)]
mod tests {
    use prover_executor_client::{
        types::{
            blob::{decode_transactions, get_origin_batch},
            input::BlobInfo,
        },
        BlobVerifier,
    };
    use prover_executor_host::blob::{encode_blob, populate_kzg};
    use prover_primitives::MorphTxEnvelope;
    #[test]
    fn test_blob() {
        //blob to txn
        let blob_bytes = load_zstd_blob();
        println!("blob_bytes len: {:?}", blob_bytes.len());

        let origin_batch = get_origin_batch(&blob_bytes).unwrap();
        println!("origin_batch len: {:?}", origin_batch.len());

        let mut block_contexts = origin_batch[0..600 * 60].to_vec();
        let txs_data = origin_batch[600 * 60..origin_batch.len()].to_vec();
        let tx_list: Vec<MorphTxEnvelope> = decode_transactions(txs_data.as_slice());
        println!("decoded tx_list_len: {:?}", tx_list.len());

        //txn to blob
        let mut tx_bytes: Vec<u8> = vec![];
        let x = tx_list.iter().flat_map(|tx| tx.rlp()).collect::<Vec<u8>>();
        tx_bytes.extend(x);
        assert!(tx_bytes == txs_data, "tx_bytes==txs_data");
        block_contexts.extend_from_slice(&tx_bytes);
        let blob = encode_blob(block_contexts).unwrap();
        let blob_info: BlobInfo = populate_kzg(&blob).unwrap();
        let (versioned_hash, batch_data) = BlobVerifier::verify(&blob_info, 600).unwrap();
        let versioned_hash_hex = alloy::hex::encode_prefixed(versioned_hash.as_slice());
        println!(
            "versioned_hash: {:?}, batch_data len: {:?}",
            versioned_hash_hex,
            batch_data.len()
        );
        assert!(
            versioned_hash_hex
                == "0x012bdf80720ba8d07c589d672e47d4b183ac861a2fcb6a5dad0e320a4f368f4f",
            "versioned_hash check"
        );

        assert!(batch_data.len() == origin_batch.len(), "batch_data.len() == origin_batch.len()");
    }

    pub fn load_zstd_blob() -> [u8; 131072] {
        use prover_primitives::alloy_primitives::hex;
        use std::{fs, path::Path};

        //https://etherscan.io/blob/0x012bdf80720ba8d07c589d672e47d4b183ac861a2fcb6a5dad0e320a4f368f4f?bid=6318849
        //https://explorer.morphl2.io/batches/47561
        let blob_data_path = Path::new("../../testdata/blob/mainnet_47561.data");
        let data = fs::read_to_string(blob_data_path).expect("Unable to read file");
        let hex_data: Vec<u8> = hex::decode(data.trim()).unwrap();
        let mut array = [0u8; 131072];
        array.copy_from_slice(&hex_data);
        array
    }
}
