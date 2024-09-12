#![no_main]

use bls12_381::{G1Projective, Scalar};
use sha2::{Digest as _, Sha256};
sp1_zkvm::entrypoint!(main);

pub fn main() {
    let x = sp1_zkvm::io::read_vec();

    let mut evaluation: [u8; 32] = Sha256::digest(&x).into();
    // let evaluation: [u8; 32] = [0; 32];
    evaluation[0] = 0;
    evaluation[1] = 0;
    evaluation[evaluation.len() - 1] = 0;
    evaluation[evaluation.len() - 2] = 0;
    let y: Scalar = Scalar::from_bytes(&evaluation).unwrap();
    let y: G1Projective = G1Projective::generator() * y;

    println!("Successfully G1Projective::generator(), y.x= {:?}", y.x);

    // Commit to the public values of the program. The final proof will have a commitment to all the
    // bytes that were committed to.
    sp1_zkvm::io::commit(&x);
}
