use std::io::{BufRead, BufReader};

use zstd::{
    stream::{Decoder, Encoder},
    zstd_safe::{CParameter, ParamSwitch},
};

// re-export zstd
pub use zstd;

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

#[allow(dead_code)]
/// Zstd encoder configuration
pub fn init_zstd_encoder(target_block_size: u32) -> Encoder<'static, Vec<u8>> {
    let mut encoder = Encoder::new(Vec::new(), 0).expect("infallible");

    // disable compression of literals, i.e. literals will be raw bytes.
    encoder
        .set_parameter(CParameter::LiteralCompressionMode(ParamSwitch::Disable))
        .expect("infallible");
    // with a hack in zstd we can set window log <= 17 with single segment kept
    encoder.set_parameter(CParameter::WindowLog(17)).expect("infallible");
    // set target block size to fit within a single block.
    encoder.set_parameter(CParameter::TargetCBlockSize(target_block_size)).expect("infallible");
    // do not include the checksum at the end of the encoded data.
    encoder.include_checksum(false).expect("infallible");
    // do not include magic bytes at the start of the frame since we will have a single
    // frame.
    encoder.include_magicbytes(false).expect("infallible");
    // do not include dictionary id so we have more simple content
    encoder.include_dictid(false).expect("infallible");
    // include the content size to know at decode time the expected size of decoded
    // data.
    encoder.include_contentsize(true).expect("infallible");

    encoder
}

pub fn init_zstd_decoder<R: BufRead>(
    reader: R,
) -> Result<Decoder<'static, BufReader<R>>, eyre::Error> {
    let mut decoder = Decoder::new(reader)?;
    decoder.include_magicbytes(false)?;
    Ok(decoder)
}
