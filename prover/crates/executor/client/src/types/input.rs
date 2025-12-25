use alloy_consensus::Transaction;
use alloy_primitives::{map::HashMap, Keccak256, U256};
use prover_storage::TrieDB;
use prover_primitives::{
    types::{BlockHeader, BlockTrace, TypedTransaction},
    Address, Block, B256,
};
use prover_storage::trace_to_execution_witness;
use revm::{primitives::keccak256, state::Bytecode};
use rsp_mpt::EthereumState;
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
pub struct L2Block {
    /// chain id
    pub chain_id: u64,
    /// coinbase
    pub coinbase: Address,
    /// block
    pub header: BlockHeader,
    /// transactions
    pub transactions: Vec<TypedTransaction>,
    /// previous state root
    pub prev_state_root: B256,
    /// post state root
    pub post_state_root: B256,
    /// start l1 queue index
    pub start_l1_queue_index: u64,
}

impl L2Block {
    pub fn from_block_trace(trace: &BlockTrace) -> Self {
        L2Block {
            chain_id: trace.chain_id(),
            coinbase: trace.coinbase(),
            header: trace.header.clone(),
            transactions: trace.typed_transactions(),
            prev_state_root: trace.root_before(),
            post_state_root: trace.root_after(),
            start_l1_queue_index: trace.start_l1_queue_index(),
        }
    }
    pub fn num_l1_txs(&self) -> u64 {
        // 0x7e is l1 tx
        match self
            .transactions
            .iter()
            .filter(|tx| tx.is_l1_msg())
            // tx.nonce for l1 tx is the l1 queue index, which is a globally index,
            // not per user as suggested by the name...
            .map(|tx| tx.nonce())
            .max()
        {
            None => 0, // not l1 tx in this block
            Some(end_l1_queue_index) => end_l1_queue_index - self.start_l1_queue_index + 1,
        }
    }

    pub fn hash_l1_msg(&self, hasher: &mut Keccak256) {
        for tx_hash in self.transactions.iter().filter(|tx| tx.is_l1_msg()).map(|tx| tx.tx_hash()) {
            hasher.update(tx_hash.as_slice())
        }
    }
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
    pub fn from_trace(trace: &BlockTrace) -> Self {
        let witness = trace_to_execution_witness(trace).unwrap();
        let state = EthereumState::from_execution_witness(&witness, trace.root_before());
        let bytecodes =
            witness.codes.into_iter().map(|code| Bytecode::new_raw(code)).collect::<Vec<_>>();
        BlockInput {
            current_block: L2Block::from_block_trace(trace),
            parent_state: state,
            bytecodes,
        }
    }
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
