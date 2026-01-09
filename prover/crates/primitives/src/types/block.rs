use alloy_consensus::{transaction::TxHashRef, Transaction};
use alloy_primitives::{Address, Keccak256, B256};
use morph_primitives::MorphTxEnvelope;
use serde::{Deserialize, Serialize};
use serde_with::serde_as;

use crate::{
    types::{BlockHeader, BlockTrace},
    Block,
};

/// L2 Block
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
    pub transactions: Vec<MorphTxEnvelope>,
    /// previous state root
    pub prev_state_root: B256,
    /// post state root
    pub post_state_root: B256,
    /// start l1 queue index
    pub start_l1_queue_index: u64,
}

impl L2Block {
    /// constructs an L2Block from a BlockTrace
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

    /// Returns the number of L1 transactions in the block.
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

    /// Hashes the L1 messages in the block using the provided hasher.
    pub fn hash_l1_msg(&self, hasher: &mut Keccak256) {
        for tx_hash in self.transactions.iter().filter(|tx| tx.is_l1_msg()).map(|tx| tx.tx_hash()) {
            hasher.update(tx_hash.as_slice())
        }
    }
}
