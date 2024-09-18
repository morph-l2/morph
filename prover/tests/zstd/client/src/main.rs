#![no_main]

use std::io::Read;

use ruzstd::StreamingDecoder;

// use alloy_primitives::keccak256;
sp1_zkvm::entrypoint!(main);

pub fn main() {
    let x = sp1_zkvm::io::read_vec();
    let mut x = x.as_slice();
    let mut decoder = StreamingDecoder::new(&mut x).unwrap();

    let mut result = Vec::new();
    decoder.read_to_end(&mut result).unwrap();

    println!("decoded data: {:?}", String::from_utf8(result.clone()));
    assert_eq!("abcdefghijklmnopqrstuvwxyz".as_bytes(), result);

    // Commit to the public values of the program. The final proof will have a commitment to all the
    // bytes that were committed to.
    sp1_zkvm::io::commit(&vec![0u64]);
}
