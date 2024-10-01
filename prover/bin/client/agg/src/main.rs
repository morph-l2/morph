#![no_main]
sp1_zkvm::entrypoint!(main);

use morph_executor_client::{types::input::AggregationInput, verify_agg};
use morph_executor_utils::u32_to_u8;
use sbv_primitives::B256;

pub fn main() {
    // Read the input.
    let input = sp1_zkvm::io::read::<AggregationInput>();

    // Execute the block.
    let agg_output = verify_agg(&input).unwrap();

    // Convert the shard vkey to a B256.
    let shard_blocks_vkey_b256 = B256::from(u32_to_u8(&input.shard_vkey));

    let mut preimage: Vec<u8> = Vec::with_capacity(64);
    preimage.extend(&agg_output.0);
    preimage.extend(&shard_blocks_vkey_b256.0);

    // Commit to the public values of the program. The final proof will have a commitment to all the
    // bytes that were committed to.
    sp1_zkvm::io::commit(&preimage);
}
