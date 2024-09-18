#![no_main]

use bls12_381::{G1Projective, Scalar};
sp1_zkvm::entrypoint!(main);

pub fn main() {
    let x = sp1_zkvm::io::read_vec();

    let mut evaluation: [u8; 32] = [0; 32];
    evaluation[0] = 0;
    evaluation[evaluation.len() - 1] = 0;

    let z: Scalar = Scalar::from_bytes(&evaluation).unwrap();
    let y: G1Projective = G1Projective::generator() * z;

    println!("Successfully G1Projective::generator(), y = {:?}", y);

    // Commit to the public values of the program. The final proof will have a commitment to all the
    // bytes that were committed to.
    sp1_zkvm::io::commit(&x);
}
