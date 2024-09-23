//! Stateless Block Verifier primitives library.

use crate::types::{TxL1Msg, TypedTransaction};
use alloy::{
    consensus::{SignableTransaction, TxEip1559, TxEip2930, TxEnvelope, TxLegacy},
    eips::eip2930::AccessList,
    primitives::{Bytes, ChainId, Signature, SignatureError, TxKind},
};
use std::fmt::Debug;

use zktrie_ng::{
    db::KVDatabase,
    hash::poseidon::Poseidon,
    trie::{Node, MAGIC_NODE_BYTES},
};

/// Predeployed contracts
pub mod predeployed;
/// Types definition
pub mod types;

pub use alloy::consensus as alloy_consensus;
pub use alloy::consensus::Transaction;
pub use alloy::primitives as alloy_primitives;
pub use alloy::primitives::{Address, B256, U256};
pub use zktrie_ng as zk_trie;

/// Blanket trait for block trace extensions.
pub trait Block: Debug {
    /// transaction type
    type Tx: TxTrace;

    /// Get block number
    fn number(&self) -> u64;

    /// Get block hash
    fn block_hash(&self) -> B256;

    /// Get chain id
    fn chain_id(&self) -> u64;

    /// Get coinbase
    fn coinbase(&self) -> Address;

    /// Get timestamp
    fn timestamp(&self) -> U256;

    /// Get gas limit
    fn gas_limit(&self) -> U256;

    /// Get base fee per gas
    fn base_fee_per_gas(&self) -> Option<U256>;

    /// Get difficulty
    fn difficulty(&self) -> U256;

    /// Get prevrandao
    fn prevrandao(&self) -> Option<B256>;

    /// transactions
    fn transactions(&self) -> impl Iterator<Item = &Self::Tx>;

    /// root before
    fn root_before(&self) -> B256;
    /// root after
    fn root_after(&self) -> B256;
    /// codes
    fn codes(&self) -> impl ExactSizeIterator<Item = &[u8]>;
    /// start l1 queue index
    fn start_l1_queue_index(&self) -> u64;

    /// flatten proofs
    fn flatten_proofs(&self) -> impl Iterator<Item = (&B256, &[u8])>;

    /// Update zktrie state from trace
    #[inline]
    fn build_zktrie_db<Db: KVDatabase>(&self, db: &mut Db) -> Result<(), Db::Error> {
        for (k, bytes) in self.flatten_proofs() {
            if bytes == MAGIC_NODE_BYTES {
                continue;
            }
            let node = Node::<Poseidon>::try_from(bytes).expect("invalid node");
            let node_hash = node.get_or_calculate_node_hash().expect("infallible");
            debug_assert_eq!(k.as_slice(), node_hash.as_slice());
            db.put_owned(node_hash.as_slice(), node.canonical_value(false))?;
        }
        Ok(())
    }

    /// Number of l1 transactions
    #[inline]
    fn num_l1_txs(&self) -> u64 {
        // 0x7e is l1 tx
        match self
            .transactions()
            .filter(|tx| tx.is_l1_tx())
            // tx.nonce for l1 tx is the l1 queue index, which is a globally index,
            // not per user as suggested by the name...
            .map(|tx| tx.nonce())
            .max()
        {
            None => 0, // not l1 tx in this block
            Some(end_l1_queue_index) => end_l1_queue_index - self.start_l1_queue_index() + 1,
        }
    }

    /// Number of l2 transactions
    #[inline]
    fn num_l2_txs(&self) -> u64 {
        // 0x7e is l1 tx
        self.transactions().filter(|tx| !tx.is_l1_tx()).count() as u64
    }

    /// Hash the header of the block
    #[inline]
    fn hash_da_header(&self, hasher: &mut impl tiny_keccak::Hasher) {
        let num_txs = (self.num_l1_txs() + self.num_l2_txs()) as u16;
        hasher.update(&self.number().to_be_bytes());
        hasher.update(&self.timestamp().to::<u64>().to_be_bytes());
        hasher.update(
            &self
                .base_fee_per_gas()
                .unwrap_or_default()
                .to_be_bytes::<{ U256::BYTES }>(),
        );
        hasher.update(&self.gas_limit().to::<u64>().to_be_bytes());
        hasher.update(&num_txs.to_be_bytes());
    }

    /// Hash the l1 messages of the block
    #[inline]
    fn hash_l1_msg(&self, hasher: &mut impl tiny_keccak::Hasher) {
        for tx_hash in self
            .transactions()
            .filter(|tx| tx.is_l1_tx())
            .map(|tx| tx.tx_hash())
        {
            hasher.update(tx_hash.as_slice())
        }
    }
}

/// Utility trait for transaction trace
pub trait TxTrace {
    /// Return the hash of the transaction
    fn tx_hash(&self) -> B256;

    /// Returns the transaction type
    fn ty(&self) -> u8;

    /// Get `nonce`.
    fn nonce(&self) -> u64;

    /// Get `gas_limit`.
    fn gas_limit(&self) -> u128;

    /// Get `gas_price`
    fn gas_price(&self) -> u128;

    /// Get `max_fee_per_gas`
    fn max_fee_per_gas(&self) -> u128;

    /// Get `max_priority_fee_per_gas`
    fn max_priority_fee_per_gas(&self) -> u128;

    /// Get `from` without checking
    ///
    /// # Safety
    ///
    /// Can only be used when the transaction is known to be an L1 transaction
    unsafe fn get_from_unchecked(&self) -> Address;

    /// Get `to`.
    fn to(&self) -> TxKind;

    /// Get `chain_id`.
    fn chain_id(&self) -> ChainId;

    /// Get `value`.
    fn value(&self) -> U256;

    /// Get `data`.
    fn data(&self) -> Bytes;

    /// Get `access_list`.
    fn access_list(&self) -> AccessList;

    /// Get `signature`.
    fn signature(&self) -> Result<Signature, SignatureError>;

    /// Check if the transaction is an L1 transaction
    fn is_l1_tx(&self) -> bool {
        self.ty() == 0x7e
    }

    /// Try to build a typed transaction
    fn try_build_typed_tx(&self) -> Result<TypedTransaction, SignatureError> {
        let chain_id = self.chain_id();

        let tx = match self.ty() {
            0x0 => {
                let tx = TxLegacy {
                    chain_id: if chain_id >= 35 { Some(chain_id) } else { None },
                    nonce: self.nonce(),
                    gas_price: self.gas_price(),
                    gas_limit: self.gas_limit(),
                    to: self.to(),
                    value: self.value(),
                    input: self.data(),
                };

                TypedTransaction::Enveloped(TxEnvelope::from(tx.into_signed(self.signature()?)))
            }
            0x1 => {
                let tx = TxEip2930 {
                    chain_id,
                    nonce: self.nonce(),
                    gas_price: self.gas_price(),
                    gas_limit: self.gas_limit(),
                    to: self.to(),
                    value: self.value(),
                    access_list: self.access_list(),
                    input: self.data(),
                };

                TypedTransaction::Enveloped(TxEnvelope::from(tx.into_signed(self.signature()?)))
            }
            0x02 => {
                let tx = TxEip1559 {
                    chain_id,
                    nonce: self.nonce(),
                    max_fee_per_gas: self.max_fee_per_gas(),
                    max_priority_fee_per_gas: self.max_priority_fee_per_gas(),
                    gas_limit: self.gas_limit(),
                    to: self.to(),
                    value: self.value(),
                    access_list: self.access_list(),
                    input: self.data(),
                };

                TypedTransaction::Enveloped(TxEnvelope::from(tx.into_signed(self.signature()?)))
            }
            0x7e => {
                let tx = TxL1Msg {
                    tx_hash: self.tx_hash(),
                    from: unsafe { self.get_from_unchecked() },
                    nonce: self.nonce(),
                    gas_limit: self.gas_limit(),
                    to: self.to(),
                    value: self.value(),
                    input: self.data(),
                };

                TypedTransaction::L1Msg(tx)
            }
            _ => unimplemented!("unsupported tx type: {}", self.ty()),
        };

        Ok(tx)
    }
}

impl<T: Block> Block for &T {
    type Tx = <T as Block>::Tx;

    fn number(&self) -> u64 {
        (*self).number()
    }

    fn block_hash(&self) -> B256 {
        (*self).block_hash()
    }

    fn chain_id(&self) -> u64 {
        (*self).chain_id()
    }

    fn coinbase(&self) -> Address {
        (*self).coinbase()
    }

    fn timestamp(&self) -> U256 {
        (*self).timestamp()
    }

    fn gas_limit(&self) -> U256 {
        (*self).gas_limit()
    }

    fn base_fee_per_gas(&self) -> Option<U256> {
        (*self).base_fee_per_gas()
    }

    fn difficulty(&self) -> U256 {
        (*self).difficulty()
    }

    fn prevrandao(&self) -> Option<B256> {
        (*self).prevrandao()
    }

    fn transactions(&self) -> impl Iterator<Item = &Self::Tx> {
        (*self).transactions()
    }

    fn root_before(&self) -> B256 {
        (*self).root_before()
    }

    fn root_after(&self) -> B256 {
        (*self).root_after()
    }

    fn codes(&self) -> impl ExactSizeIterator<Item = &[u8]> {
        (*self).codes()
    }

    fn start_l1_queue_index(&self) -> u64 {
        (*self).start_l1_queue_index()
    }

    fn flatten_proofs(&self) -> impl Iterator<Item = (&B256, &[u8])> {
        (*self).flatten_proofs()
    }
}

impl<T: TxTrace> TxTrace for &T {
    fn tx_hash(&self) -> B256 {
        (*self).tx_hash()
    }

    fn ty(&self) -> u8 {
        (*self).ty()
    }

    fn nonce(&self) -> u64 {
        (*self).nonce()
    }

    fn gas_limit(&self) -> u128 {
        (*self).gas_limit()
    }

    fn gas_price(&self) -> u128 {
        (*self).gas_price()
    }

    fn max_fee_per_gas(&self) -> u128 {
        (*self).max_fee_per_gas()
    }

    fn max_priority_fee_per_gas(&self) -> u128 {
        (*self).max_priority_fee_per_gas()
    }

    unsafe fn get_from_unchecked(&self) -> Address {
        (*self).get_from_unchecked()
    }

    fn to(&self) -> TxKind {
        (*self).to()
    }

    fn chain_id(&self) -> ChainId {
        (*self).chain_id()
    }

    fn value(&self) -> U256 {
        (*self).value()
    }

    fn data(&self) -> Bytes {
        (*self).data()
    }

    fn access_list(&self) -> AccessList {
        (*self).access_list()
    }

    fn signature(&self) -> Result<Signature, SignatureError> {
        (*self).signature()
    }
}
