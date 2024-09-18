#![no_main]

// use alloy_primitives::keccak256;
sp1_zkvm::entrypoint!(main);

pub fn main() {
    let x = sp1_zkvm::io::read::<u32>();

    let y = x + 1;
    // Commit to the public values of the program. The final proof will have a commitment to all the
    // bytes that were committed to.
    sp1_zkvm::io::commit(&y);
}
