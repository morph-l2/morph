#![no_main]

use std::io::{BufRead, BufReader, Cursor};

use zstd::stream::Decoder;
// re-export zstd
pub use zstd;
// use alloy_primitives::keccak256;
sp1_zkvm::entrypoint!(main);

pub fn main() {
    let x = sp1_zkvm::io::read_vec();

    let cursor = Cursor::new(x);
    let mut decoder = init_zstd_decoder(cursor).unwrap();
    let mut decompressed_batch = Vec::new();
    std::io::copy(&mut decoder, &mut decompressed_batch).unwrap();

    // Commit to the public values of the program. The final proof will have a commitment to all the
    // bytes that were committed to.
    sp1_zkvm::io::commit(&vec![0u64]);
}

// we use offset window no more than = 17
// TODO: use for multi-block zstd.
#[allow(dead_code)]
pub const CL_WINDOW_LIMIT: usize = 17;

#[allow(dead_code)]
/// zstd block size target.
pub const N_BLOCK_SIZE_TARGET: u32 = 124 * 1024;

#[allow(dead_code)]
/// Maximum number of blocks that we can expect in the encoded data.
pub const N_MAX_BLOCKS: u64 = 10;

pub fn init_zstd_decoder<R: BufRead>(
    reader: R,
) -> Result<Decoder<'static, BufReader<R>>, anyhow::Error> {
    let mut decoder = Decoder::new(reader)?;
    decoder.include_magicbytes(false)?;
    Ok(decoder)
}
