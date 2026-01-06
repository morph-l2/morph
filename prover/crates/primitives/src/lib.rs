//! Stateless Block Verifier primitives library.

use alloy_consensus::{SignableTransaction, TxEip1559, TxEip2930, TxEip7702, TxLegacy};
use alloy_eips::eip2930::AccessList;
use alloy_primitives::{Bytes, ChainId, Signature, SignatureError, TxKind};
use morph_primitives::{TxAltFee, TxL1Msg};
use std::fmt::Debug;

/// Predeployed contracts
pub mod predeployed;
/// Types definition
pub mod types;

pub use alloy_consensus;
pub use alloy_consensus::Transaction;
pub use alloy_eips::eip7702::SignedAuthorization;
pub use alloy_primitives;
pub use alloy_primitives::{Address, B256, U256};
pub use morph_primitives::MorphTxEnvelope;

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
    // #[inline]
    // fn build_zktrie_db(&self, zktrie_db: &mut ZkMemoryDb) {
    //     init_hash_scheme();
    //     for (_, bytes) in self.flatten_proofs() {
    //         zktrie_db.add_node_bytes(bytes, None).unwrap();
    //     }
    // }

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
    fn gas_limit(&self) -> u64;

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

    /// Get `authorization_list`.
    fn authorization_list(&self) -> Vec<SignedAuthorization>;

    /// Get `signature`.
    fn signature(&self) -> Result<Signature, SignatureError>;

    /// Check if the transaction is an L1 transaction
    fn is_l1_tx(&self) -> bool {
        self.ty() == 0x7e
    }

    /// Get `fee_token_id`.
    fn fee_token_id(&self) -> u16;

    /// Get `fee_limit`.
    fn fee_limit(&self) -> U256;

    /// Get `sig_v`.
    fn sig_v(&self) -> u64;

    /// Try to build a envelope tx
    fn try_build_tx_envelope(&self) -> Result<MorphTxEnvelope, SignatureError> {
        let chain_id = self.chain_id();

        let tx = match self.ty() {
            0x0 => {
                fn chain_id_from_v_eip155(v: u64) -> Option<u64> {
                    if v >= 35 {
                        Some((v - 35) / 2)
                    } else {
                        None
                    }
                }
                let tx = TxLegacy {
                    chain_id: chain_id_from_v_eip155(self.sig_v()),
                    nonce: self.nonce(),
                    gas_price: self.gas_price(),
                    gas_limit: self.gas_limit(),
                    to: self.to(),
                    value: self.value(),
                    input: self.data(),
                };

                MorphTxEnvelope::Legacy(tx.into_signed(self.signature()?))
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

                MorphTxEnvelope::Eip2930(tx.into_signed(self.signature()?))
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

                MorphTxEnvelope::Eip1559(tx.into_signed(self.signature()?))
            }
            0x04 => {
                let tx = TxEip7702 {
                    chain_id,
                    nonce: self.nonce(),
                    gas_limit: self.gas_limit(),
                    max_fee_per_gas: self.max_fee_per_gas(),
                    max_priority_fee_per_gas: self.max_priority_fee_per_gas(),
                    to: *self.to().to().expect("EIP-7702 transaction must have a recipient"),
                    value: self.value(),
                    access_list: self.access_list(),
                    authorization_list: self.authorization_list(),
                    input: self.data(),
                };
                MorphTxEnvelope::Eip7702(tx.into_signed(self.signature()?))
            }
            0x7e => {
                let tx = TxL1Msg {
                    tx_hash: self.tx_hash(),
                    from: unsafe { self.get_from_unchecked() },
                    nonce: self.nonce(),
                    gas_limit: self.gas_limit() as u128,
                    to: self.to(),
                    value: self.value(),
                    input: self.data(),
                };

                MorphTxEnvelope::L1Msg(tx.into_signed(self.signature()?))
            }
            0x7f => {
                let tx = TxAltFee {
                    chain_id,
                    nonce: self.nonce(),
                    gas_limit: self.gas_limit() as u128,
                    max_fee_per_gas: self.max_fee_per_gas(),
                    max_priority_fee_per_gas: self.max_priority_fee_per_gas(),
                    to: self.to(),
                    value: self.value(),
                    access_list: self.access_list(),
                    input: self.data(),
                    fee_token_id: self.fee_token_id(),
                    fee_limit: self.fee_limit(),
                };
                MorphTxEnvelope::AltFee(tx.into_signed(self.signature()?))
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

    fn gas_limit(&self) -> u64 {
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

    fn authorization_list(&self) -> Vec<SignedAuthorization> {
        (*self).authorization_list()
    }

    fn signature(&self) -> Result<Signature, SignatureError> {
        (*self).signature()
    }

    fn fee_token_id(&self) -> u16 {
        (*self).fee_token_id()
    }

    fn fee_limit(&self) -> U256 {
        (*self).fee_limit()
    }

    fn sig_v(&self) -> u64 {
        (*self).sig_v()
    }
}
