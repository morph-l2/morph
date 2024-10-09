use anyhow::anyhow;
pub mod evm;
use evm::{save_plonk_fixture, EvmProofFixture};
use morph_executor_client::{types::input::ClientInput, verify};
use morph_executor_host::get_blob_info;
use morph_executor_utils::read_env_var;
use sbv_primitives::{alloy_primitives::keccak256, types::BlockTrace, B256};
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
    // Setup the logger.
    sp1_sdk::utils::setup_logger();
    let program_hash = keccak256(BATCH_VERIFIER_ELF);
    println!("Program Hash [view on Explorer]:");
    println!("{}", alloy::hex::encode_prefixed(program_hash));

    if blocks.len() > MAX_PROVE_BLOCKS {
        return Err(anyhow!(format!(
            "check block_tracs, blocks len = {:?} exceeds MAX_PROVE_BLOCKS = {:?}",
            blocks.len(),
            MAX_PROVE_BLOCKS
        )));
    }

    // Prepare input.
    // Convert the traces' format to reduce conversion costs in the client.
    blocks.iter_mut().for_each(|blobk| blobk.flatten());

    let client_input =
        ClientInput { l2_traces: blocks.clone(), blob_info: get_blob_info(blocks).unwrap() };

    // Execute the program in native
    let expected_hash =
        verify(&client_input).map_err(|e| anyhow!(format!("native execution err: {:?}", e)))?;
    println!(
        "pi_hash generated with native execution: {}",
        alloy::hex::encode_prefixed(expected_hash.as_slice())
    );

    // Execute the program in sp1-vm
    let mut stdin = SP1Stdin::new();
    stdin.write(&serde_json::to_string(&client_input).unwrap());
    let client = ProverClient::new();
    let (mut public_values, execution_report) = client
        .execute(BATCH_VERIFIER_ELF, stdin.clone())
        .run()
        .map_err(|e| anyhow!(format!("native execution err: {:?}", e)))?;

    println!(
        "Program executed successfully, Number of cycles: {:?}",
        execution_report.total_instruction_count()
    );
    let pi_hash = public_values.read::<[u8; 32]>();
    let public_values = B256::from_slice(&pi_hash);

    println!(
        "pi_hash generated with sp1-vm execution: {}",
        alloy::hex::encode_prefixed(public_values.as_slice())
    );
    assert_eq!(pi_hash, expected_hash, "pi_hash == expected_pi_hash ");
    println!("Values are correct!");

    if !prove {
        println!("Execution completed, No prove request, skipping...");
        return Ok(None);
    }
    println!("Start proving...");
    // Setup the program for proving.
    let (pk, vk) = client.setup(BATCH_VERIFIER_ELF);
    println!("Batch ELF Verification Key: {:?}", vk.vk.bytes32());

    // Generate the proof
    let start = Instant::now();
    let mut proof = client.prove(&pk, stdin).plonk().run().expect("proving failed");

    let duration_mins = start.elapsed().as_secs() / 60;
    println!("Successfully generated proof!, time use: {:?} minutes", duration_mins);

    // Verify the proof.
    client.verify(&proof, &vk).expect("failed to verify proof");
    println!("Successfully verified proof!");

    // Deserialize the public values.
    let pi_bytes = proof.public_values.read::<[u8; 32]>();
    println!("pi_hash generated with sp1-vm prove: {}", alloy::hex::encode_prefixed(pi_bytes));
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

#[cfg(test)]
mod tests {
    use morph_executor_client::{
        types::{
            blob::{decode_transactions, get_origin_batch},
            input::BlobInfo,
        },
        BlobVerifier,
    };
    use morph_executor_host::{encode_blob, populate_kzg};
    use sbv_primitives::{alloy_primitives::hex, types::TypedTransaction};
    #[test]
    fn test_blob() {
        //blob to txn
        let blob_bytes = load_zstd_blob();
        println!("blob_bytes len: {:?}", blob_bytes.len());

        let origin_batch = get_origin_batch(&blob_bytes).unwrap();

        let origin_tx_bytes = decode_raw_tx_payload(origin_batch);
        println!("origin_tx_bytes len: {:?}", origin_tx_bytes.len());

        let tx_list: Vec<TypedTransaction> = decode_transactions(origin_tx_bytes.as_slice());
        println!("decoded tx_list_len: {:?}", tx_list.len());

        //txn to blob
        let mut tx_bytes: Vec<u8> = vec![];
        let x = tx_list.iter().flat_map(|tx| tx.rlp()).collect::<Vec<u8>>();
        tx_bytes.extend(x);
        assert!(tx_bytes == origin_tx_bytes, "tx_bytes==origin_tx_bytes");
        // tx_bytes[121..1000].fill(0);
        let blob = encode_blob(tx_bytes);

        std::env::set_var("TRUSTED_SETUP_4844", "../../configs/4844_trusted_setup.txt");
        let blob_info: BlobInfo = populate_kzg(&blob).unwrap();

        let (versioned_hash, batch_data) = BlobVerifier::verify(&blob_info).unwrap();
        println!(
            "versioned_hash: {:?}, batch_data len: {:?}",
            hex::encode(versioned_hash.as_slice()),
            batch_data.len()
        );
    }

    fn decode_raw_tx_payload(origin_batch: Vec<u8>) -> Vec<u8> {
        assert!(origin_batch.len() > 182, "batch.len need less than METADATA_LENGTH");

        let num_valid_chunks = u16::from_be_bytes(origin_batch[0..2].try_into().unwrap()); // size of num_valid_chunks is 2bytes.
        assert!(num_valid_chunks as usize <= 45, "Exceeded MAX_AGG_SNARKS");

        let data_size: u64 = origin_batch[2..2 + 4 * num_valid_chunks as usize]
            .chunks_exact(4)
            .map(|chunk| u32::from_be_bytes(chunk.try_into().unwrap()) as u64)
            .sum();

        let tx_payload_end = 182 + data_size as usize;
        assert!(
            origin_batch.len() >= tx_payload_end,
            "The batch does not contain the complete tx_payload"
        );

        origin_batch[182..tx_payload_end].to_vec()
    }

    pub fn load_zstd_blob() -> [u8; 131072] {
        use sbv_primitives::alloy_primitives::hex;
        use std::{fs, path::Path};

        //https://holesky.etherscan.io/blob/0x018494ae7657bebd9e590baf3736ac9207a5d2275ef98c025dad3232b7875278?bid=2391294
        //https://explorer-holesky.morphl2.io/batches/223946
        let blob_data_path = Path::new("../../testdata/blob/blob_with_zstd_batch_holesky.data");
        let data = fs::read_to_string(blob_data_path).expect("Unable to read file");
        let hex_data: Vec<u8> = hex::decode(data.trim()).unwrap();
        let mut array = [0u8; 131072];
        array.copy_from_slice(&hex_data);
        array
    }
}
