use morph_executor_client::{types::input::ClientInput, verify};
use morph_executor_host::get_blob_info;
use sbv_primitives::{alloy_primitives::keccak256, types::BlockTrace, B256};
use sp1_sdk::{HashableKey, ProverClient, SP1Stdin};
use std::time::Instant;

/// The ELF (executable and linkable format) file for the Succinct RISC-V zkVM.
pub const BATCH_VERIFIER_ELF: &[u8] =
    include_bytes!("../../client/elf/riscv32im-succinct-zkvm-elf");

pub fn prove(blocks: &mut Vec<BlockTrace>, prove: bool) {
    // Setup the logger.
    sp1_sdk::utils::setup_logger();
    let program_hash = keccak256(BATCH_VERIFIER_ELF);
    println!("Program Hash [view on Explorer]:");
    println!("0x{}", hex::encode(program_hash));

    // Prepare input.
    // Convert the traces' format to reduce conversion costs in the client.
    blocks.iter_mut().for_each(|blobk| blobk.flatten());

    let client_input =
        ClientInput { l2_traces: blocks.clone(), blob_info: get_blob_info(blocks).unwrap() };

    // Execute the program in native
    let expected_hash = verify(&client_input).unwrap();
    println!("pi_hash generated with native execution: {}", hex::encode(expected_hash.as_slice()));

    // Execute the program in sp1-vm
    let mut stdin = SP1Stdin::new();
    stdin.write(&serde_json::to_string(&client_input).unwrap());
    let client = ProverClient::new();
    let (mut public_values, execution_report) =
        client.execute(BATCH_VERIFIER_ELF, stdin.clone()).run().unwrap();

    println!(
        "Program executed successfully, Number of cycles: {:?}",
        execution_report.total_instruction_count()
    );
    let pi_hash = public_values.read::<B256>();
    println!("pi_hash generated with sp1-vm execution: {}", hex::encode(pi_hash.as_slice()));
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
    println!("plonk proof: {:#?}", proof);

    let duration_mins = start.elapsed().as_secs() / 60;
    println!("Successfully generated proof!, time use: {:?} minutes", duration_mins);

    // Verify the proof.
    client.verify(&proof, &vk).expect("failed to verify proof");
    println!("Successfully verified proof!");
}

#[test]
fn test_blob() {
    use morph_executor_client::{
        types::{
            blob::{decode_transactions, get_origin_batch},
            input::BlobInfo,
        },
        BlobVerifier,
    };
    use morph_executor_host::encode_blob;
    use sbv_primitives::{alloy_primitives::hex, types::TypedTransaction};

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
    tx_bytes[121..1000].fill(0);

    //https://holesky.etherscan.io/blob/0x018494ae7657bebd9e590baf3736ac9207a5d2275ef98c025dad3232b7875278?bid=2391294
    //https://explorer-holesky.morphl2.io/batches/223946
    let blob = encode_blob(tx_bytes);
    let blob_info = BlobInfo {
        blob_data: blob.to_vec(),
        commitment: hex::decode(
            "0x83a5d1ffa11a6c5246a91002415ce9b815c31a8e204b09411ac478f0eef5f9b832ce8d0588ebf4e1000f6fa811372a72",
        )
        .unwrap(),
        proof: hex::decode(
            "0x8fb9142aa00410565ba9469f9fe7a56c0f6655966cddd0b9d51f6f91a4a88b050279b8d875f0da1e6b4694dfe33b0c90",
        )
        .unwrap(),
    };

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

    let blob_data_path = Path::new("../../testdata/blob/blob_with_zstd_batch_holesky.data");
    let data = fs::read_to_string(blob_data_path).expect("Unable to read file");
    let hex_data: Vec<u8> = hex::decode(data.trim()).unwrap();
    let mut array = [0u8; 131072];
    array.copy_from_slice(&hex_data);
    array
}
