use super::input::ClientInput;
use sbv_primitives::{TxTrace, B256};
use tiny_keccak::{Hasher, Keccak};

pub fn shard_hash(input: &ClientInput) -> Result<(B256, B256), anyhow::Error> {
    // Verify DA
    let mut l2_tx_bytes: Vec<u8> = vec![];
    let mut l1_tx_bytes: Vec<u8> = vec![];

    for trace in &input.l2_traces {
        for tx in &trace.transactions {
            let tx_data = tx.try_build_typed_tx().unwrap().rlp();
            if tx.is_l1_tx() {
                l1_tx_bytes.extend(tx_data);
            } else {
                l2_tx_bytes.extend(tx_data);
            }
        }
    }

    println!("cycle-tracker-start: l1_data_hash");
    let mut l1_hasher = Keccak::v256();
    l1_hasher.update(&l1_tx_bytes);
    let mut l1_data_hash = B256::ZERO;
    l1_hasher.finalize(&mut l1_data_hash.0);
    println!("cycle-tracker-end: l1_data_hash");

    println!("cycle-tracker-start: l2_data_hash");
    let mut l2_hasher = Keccak::v256();
    l2_hasher.update(&l2_tx_bytes);
    let mut l2_data_hash = B256::ZERO;
    l2_hasher.finalize(&mut l2_data_hash.0);
    println!("cycle-tracker-start: l2_data_hash");

    Ok((l1_data_hash, l2_data_hash))
}
