use alloy_primitives::{map::HashMap, U256};
use prover_mpt::EthereumState;
use prover_primitives::{types::block::L2Block, Address, B256};
use prover_storage_witness::TrieDB;
use revm::{primitives::keccak256, state::Bytecode};
use serde::{Deserialize, Serialize};
use serde_with::serde_as;

use crate::types::error::ClientError;

#[serde_as]
#[derive(Clone, Debug, Default, Serialize, Deserialize)]
pub struct BlobInfo {
    pub blob_data: Vec<u8>,
    pub commitment: Vec<u8>,
    pub proof: Vec<u8>,
}

#[serde_as]
#[derive(Debug, Clone, Serialize, Deserialize, PartialEq, Eq)]
pub struct BlockInput {
    // l2 block info
    pub current_block: L2Block,

    /// state as of the parent block.
    pub parent_state: EthereumState,

    /// Account bytecodes.
    pub bytecodes: Vec<Bytecode>,
}

impl BlockInput {
    pub fn witness_db(&self) -> Result<TrieDB<'_>, ClientError> {
        // verify the state root
        if self.current_block.prev_state_root != self.parent_state.state_root() {
            return Err(ClientError::InvalidHeaderStateRoot);
        }

        let bytecodes_by_hash =
            self.bytecodes.iter().map(|code| (code.hash_slow(), code)).collect::<HashMap<_, _>>();
        let block_hashes: HashMap<u64, B256> = HashMap::with_hasher(Default::default());

        Ok(TrieDB::new(&self.parent_state, block_hashes, bytecodes_by_hash))
    }

    /// Get storage value of address at index.
    pub fn get_storage_value(&self, address: Address, index: U256) -> Result<U256, ClientError> {
        let hashed_address = keccak256(address);
        let hashed_address = hashed_address.as_slice();

        let storage_trie = self
            .parent_state
            .storage_tries
            .get(hashed_address)
            .expect("A storage trie must be provided for each account");

        Ok(storage_trie
            .get_rlp::<U256>(keccak256(index.to_be_bytes::<32>()).as_slice())
            .expect("Can get from MPT")
            .unwrap_or_default())
    }
}

#[serde_as]
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ExecutorInput {
    pub block_inputs: Vec<BlockInput>,
    pub blob_info: BlobInfo,
}

#[cfg(test)]
mod tests {
    use crate::types::input::L2Block;
    use alloy_consensus::Typed2718;
    use prover_primitives::types::BlockTrace;
    use std::fs::File;
    use std::io::BufReader;

    #[test]
    fn test_trace_to_execution_witness() {
        let block_trace = load_trace("../../../testdata/mpt/local_transfer_eth.json");
        println!("loaded {} blocks", block_trace.len());
        let blocks: Vec<L2Block> =
            block_trace.iter().map(|trace| L2Block::from_block_trace(trace)).collect();

        let first_block = blocks.first().unwrap();
        let txs = first_block.transactions.clone();
        let first_txn = txs.first().unwrap();

        println!("first_txn ty: {:?}", first_txn.ty());
    }

    fn load_trace(file_path: &str) -> Vec<BlockTrace> {
        let file = File::open(file_path).unwrap();
        let reader = BufReader::new(file);
        serde_json::from_reader(reader).unwrap()
    }
}
