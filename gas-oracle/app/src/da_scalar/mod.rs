mod blob;
mod blob_client;
mod calculate;
mod error;
#[allow(clippy::module_inception)]
pub mod l1_scalar;
mod typed_tx;
mod zstd_util;

const MAX_BLOB_TX_PAYLOAD_SIZE: usize = 131072; // 131072 = 4096 * 32 = 1024 * 4 * 32 = 128kb
