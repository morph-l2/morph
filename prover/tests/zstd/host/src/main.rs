use sp1_sdk::{ProverClient, SP1Stdin};
use std::io::Read;
use std::{fs::File, path::Path, time::Instant};

fn main() {
    // Setup the logger.
    sp1_sdk::utils::setup_logger();

    // The ELF (executable and linkable format) file for the Succinct RISC-V zkVM.
    let dev_elf: &[u8] = include_bytes!("../../client/elf/riscv32im-succinct-zkvm-elf");

    let data_path = Path::new("./abc.txt.zst");
    let mut file = File::open(data_path).unwrap();
    let mut data = Vec::new();
    file.read_to_end(&mut data).unwrap();

    // Setup the inputs.
    let mut stdin = SP1Stdin::new();

    stdin.write_slice(&data);

    let client = ProverClient::new();
    // Execute the program in sp1-vm
    let (public_values, execution_report) = client.execute(dev_elf, stdin.clone()).run().unwrap();
    println!("Program executed successfully.");
    // Record the number of cycles executed.
    println!(
        "Number of cycles: {}",
        execution_report.total_instruction_count()
    );

    let rt_data = public_values.as_slice();
    println!(
        "pi_hash generated with sp1-vm execution: {}",
        hex::encode(rt_data)
    );

    let start = Instant::now();

    // Setup the program for proving.
    let (pk, vk) = client.setup(dev_elf);

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

#[test]
fn test() {
    use ruzstd::StreamingDecoder;
    use std::io::Read;

    let data_path = Path::new("./decoded_blob_with_compressed_batch.data");
    let mut file = File::open(data_path).unwrap();
    let mut data = Vec::new();
    file.read_to_end(&mut data).unwrap();

    // This magic number is included at the start of a single Zstandard frame
    let mut content = 0xFD2F_B528u32.to_le_bytes().to_vec();
    content.append(&mut data);
    let mut x = content.as_slice();

    let mut decoder = StreamingDecoder::new(&mut x).unwrap();
    let mut result = Vec::new();
    decoder.read_to_end(&mut result).unwrap();
    println!("decoded data len: {:?}", result.len());
    assert_eq!(result.len(), 125091);
}
