#![no_main]
sp1_zkvm::entrypoint!(main);

use prover_executor_client::{types::input::ExecutorInput, verify};

pub fn main() {
    // Read the input.
    let x = sp1_zkvm::io::read::<String>();
    let input: ExecutorInput = serde_json::from_str(&x).unwrap();

    // Execute the block.
    let pi_hash = verify(input).unwrap();

    // Commit to the public values of the program. The final proof will have a commitment to all the
    // bytes that were committed to.
    sp1_zkvm::io::commit(&pi_hash.0);
}
