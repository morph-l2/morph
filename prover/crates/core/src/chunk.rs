use revm::primitives::B256;
use sbv_primitives::{zk_trie::ZkMemoryDb, Block};
use std::rc::Rc;
use tiny_keccak::{Hasher, Keccak};

/// A chunk is a set of continuous blocks.
/// ChunkInfo is metadata of chunk, with following fields:
/// - state root before this chunk
/// - state root after this chunk
/// - the withdraw root after this chunk
/// - the data hash of this chunk
/// - the tx data hash of this chunk
/// - flattened L2 tx bytes hash
#[derive(Debug)]
pub struct ChunkInfo {
    chain_id: u64,
    prev_state_root: B256,
    post_state_root: B256,
    withdraw_root: B256,
    data_hash: B256,
}

impl ChunkInfo {
    /// Construct by block traces
    pub fn from_block_traces<T: Block>(traces: &[T]) -> (Self, Rc<ZkMemoryDb>) {
        let chain_id = traces.first().unwrap().chain_id();
        let prev_state_root = traces.first().expect("at least 1 block needed").root_before();
        let post_state_root = traces.last().expect("at least 1 block needed").root_after();
        let withdraw_root = traces.last().unwrap().withdraw_root();

        let mut data_hasher = Keccak::v256();
        for trace in traces.iter() {
            trace.hash_da_header(&mut data_hasher);
        }
        for trace in traces.iter() {
            trace.hash_l1_msg(&mut data_hasher);
        }
        let mut data_hash = B256::ZERO;
        data_hasher.finalize(&mut data_hash.0);

        let mut zktrie_db = ZkMemoryDb::new();
        for trace in traces.iter() {
            measure_duration_histogram!(
                build_zktrie_db_duration_microseconds,
                trace.build_zktrie_db(&mut zktrie_db)
            );
        }
        let zktrie_db = Rc::new(zktrie_db);

        let info =
            ChunkInfo { chain_id, prev_state_root, post_state_root, withdraw_root, data_hash };

        (info, zktrie_db)
    }

    /// Public input hash for a given chunk is defined as
    /// keccak(
    ///     chain id ||
    ///     prev state root ||
    ///     post state root ||
    ///     withdraw root ||
    ///     chunk data hash ||
    ///     chunk txdata hash
    /// )
    pub fn public_input_hash(&self, tx_bytes_hash: &B256) -> B256 {
        let mut hasher = Keccak::v256();

        hasher.update(&self.chain_id.to_be_bytes());
        hasher.update(self.prev_state_root.as_ref());
        hasher.update(self.post_state_root.as_slice());
        hasher.update(self.withdraw_root.as_slice());
        hasher.update(self.data_hash.as_slice());
        hasher.update(tx_bytes_hash.as_slice());

        let mut public_input_hash = B256::ZERO;
        hasher.finalize(&mut public_input_hash.0);
        public_input_hash
    }

    /// Chain ID of this chunk
    pub fn chain_id(&self) -> u64 {
        self.chain_id
    }

    /// State root before this chunk
    pub fn prev_state_root(&self) -> B256 {
        self.prev_state_root
    }

    /// State root after this chunk
    pub fn post_state_root(&self) -> B256 {
        self.post_state_root
    }

    /// Withdraw root after this chunk
    pub fn withdraw_root(&self) -> B256 {
        self.withdraw_root
    }

    /// Data hash of this chunk
    pub fn data_hash(&self) -> B256 {
        self.data_hash
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::{EvmExecutorBuilder, HardforkConfig};
    use revm::primitives::b256;
    use sbv_primitives::types::BlockTrace;
    use std::cell::RefCell;

    const TRACES_STR: [&str; 4] = [
        include_str!("../../../testdata/dev.json"),
        include_str!("../../../testdata/dev.json"),
        include_str!("../../../testdata/dev.json"),
        include_str!("../../../testdata/dev.json"),
    ];

    #[test]
    fn test_public_input_hash() {
        let traces: [BlockTrace; 4] = TRACES_STR.map(|s| {
            #[derive(serde::Deserialize)]
            pub struct BlockTraceJsonRpcResult {
                pub result: BlockTrace,
            }
            serde_json::from_str::<BlockTraceJsonRpcResult>(s).unwrap().result
        });

        let fork_config = HardforkConfig::default_from_chain_id(traces[0].chain_id());
        let (chunk_info, zktrie_db) = ChunkInfo::from_block_traces(&traces);

        let tx_bytes_hasher = RefCell::new(Keccak::v256());

        let mut executor = EvmExecutorBuilder::new(zktrie_db.clone())
            .hardfork_config(fork_config)
            .with_execute_hooks(|hooks| {
                hooks.add_tx_rlp_handler(|_, rlp| {
                    tx_bytes_hasher.borrow_mut().update(rlp);
                });
            })
            .build(&traces[0])
            .unwrap();
        executor.handle_block(&traces[0]).unwrap();

        for trace in traces[1..].iter() {
            executor.update_db(trace).unwrap();
            executor.handle_block(trace).unwrap();
        }

        let post_state_root = executor.commit_changes(&zktrie_db);
        assert_eq!(post_state_root, chunk_info.post_state_root);
        drop(executor); // drop executor to release Rc<Keccek>

        let mut tx_bytes_hash = B256::ZERO;
        tx_bytes_hasher.into_inner().finalize(&mut tx_bytes_hash.0);
        let public_input_hash = chunk_info.public_input_hash(&tx_bytes_hash);
        assert_eq!(
            public_input_hash,
            b256!("764bffabc9fd4227d447a46d8bb04e5448ed64d89d6e5f4215fcf3593e00f109")
        );
    }
}
