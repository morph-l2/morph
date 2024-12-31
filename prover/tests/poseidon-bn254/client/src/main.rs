#![no_main]

use alloy::primitives::B256;
use poseidon_bn254::{hash_with_domain, Fr, PrimeField};

sp1_zkvm::entrypoint!(main);

fn cacl_poseidon_hash() -> B256 {
    // let mut domain_byte32 = [0u8; 32];

    let mut hash_bytes =
        hash_with_domain(&[Fr::from(1u64), Fr::from(2u64)], Fr::from(3u64)).to_repr();
    hash_bytes.reverse();
    B256::from_slice(&hash_bytes)
}

pub fn main() {
    let _x = sp1_zkvm::io::read_vec();

    let hash_bytes = cacl_poseidon_hash();
    // Commit to the public values of the program. The final proof will have a commitment to all the
    // bytes that were committed to.
    sp1_zkvm::io::commit(&hash_bytes);
}
