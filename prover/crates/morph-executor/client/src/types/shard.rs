use sbv_primitives::{types::BlockTrace, Block, TxTrace, B256};
use tiny_keccak::{Hasher, Keccak};

pub fn shard_bytes_hash(l2_traces: &Vec<BlockTrace>) -> Result<(B256, Vec<u8>), anyhow::Error> {
    println!("cycle-tracker-start: shard_data_hash");
    let mut data_hasher = Keccak::v256();
    let mut l2_tx_bytes: Vec<u8> = vec![];
    for trace in l2_traces {
        trace.hash_da_header(&mut data_hasher);
        for tx in &trace.transactions {
            let tx_data = tx.try_build_typed_tx().unwrap().rlp();
            data_hasher.update(&tx_data);
            if !tx.is_l1_tx() {
                l2_tx_bytes.extend(&tx_data);
            }
        }
    }
    let mut shard_data_hash = B256::ZERO;
    data_hasher.finalize(&mut shard_data_hash.0);
    println!("cycle-tracker-end: shard_data_hash");

    Ok((shard_data_hash, l2_tx_bytes))
}
