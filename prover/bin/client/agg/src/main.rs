#![no_main]
sp1_zkvm::entrypoint!(main);

use morph_executor_client::{types::input::AggregationInput, verify_agg};

pub fn main() {
    // Read the input.
    let input = sp1_zkvm::io::read::<AggregationInput>();

    // Execute the block.
    let pi = verify_agg(input).unwrap();

    // Commit to the public values of the program. The final proof will have a commitment to all the
    // bytes that were committed to.
    sp1_zkvm::io::commit(&pi.0);
}
