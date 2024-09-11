use sbv_primitives::{types::BlockTrace, Transaction, TxTrace, U256};

/// The number of coefficients (BLS12-381 scalars) to represent the blob polynomial in evaluation
/// form.
pub const BLOB_WIDTH: usize = 4096;

/// The number of bytes to represent an unsigned 256 bit number.
pub const N_BYTES_U256: usize = 32;

#[derive(Clone, Debug)]
pub struct BlobData {}

pub fn get_blob_data(l2_trace: &BlockTrace) -> [U256; BLOB_WIDTH] {
    let mut coefficients = [[0u8; N_BYTES_U256]; BLOB_WIDTH];

    let tx_bytes = l2_trace
        .transactions
        .iter()
        .filter(|tx| !tx.is_l1_tx())
        .flat_map(|tx| tx.try_build_typed_tx())
        .collect::<Vec<u8>>();

    // We only consider the data from `valid` chunks and ignore the padded chunks.
    let metadata_bytes = self.to_metadata_bytes();
    let blob_bytes = metadata_bytes.iter().chain(
        self.chunk_data
            .iter()
            .take(self.num_valid_chunks as usize)
            .flatten(),
    );

    for (i, &byte) in blob_bytes.enumerate() {
        coefficients[i / 31][1 + (i % 31)] = byte;
    }

    coefficients.map(|coeff| U256::from_big_endian(&coeff))
}
