#![no_main]
sp1_zkvm::entrypoint!(main);

// use eth_types::l2_types::BlockTrace;
use morph_executor::verify;
use sbv_primitives::types::BlockTrace;

pub fn main() {
    // Read the input.
    let x = sp1_zkvm::io::read::<String>();
    let trace: BlockTrace = serde_json::from_str(&x).unwrap();

    // Execute the block.
    let pi_hash = verify(&trace).unwrap();

    // Commit to the public values of the program. The final proof will have a commitment to all the
    // bytes that were committed to.
    sp1_zkvm::io::commit(&pi_hash);
}
