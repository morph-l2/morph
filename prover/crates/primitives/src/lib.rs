//! Stateless Block Verifier primitives library.

use crate::types::{tx_alt_fee::TxAltFee, TxL1Msg, TypedTransaction};
use alloy::{
    consensus::{SignableTransaction, TxEip1559, TxEip2930, TxEip7702, TxEnvelope, TxLegacy},
    eips::eip2930::AccessList,
    hex::{self, FromHex, ToHexExt},
    primitives::{keccak256, Bytes, ChainId, Signature, SignatureError, TxKind},
};
use std::{fmt::Debug, sync::Once};
use zktrie::{ZkMemoryDb, ZkTrieNode};

/// Predeployed contracts
pub mod predeployed;
/// Types definition
pub mod types;

pub use alloy::{
    consensus as alloy_consensus,
    consensus::Transaction,
    eips::eip7702::SignedAuthorization,
    primitives as alloy_primitives,
    primitives::{Address, B256, U256},
};
pub use zktrie as zk_trie;

/// Initialize the hash scheme for zkTrie.
pub fn init_hash_scheme() {
    static INIT: Once = Once::new();
    INIT.call_once(|| {
        zktrie::init_hash_scheme_simple(|a: &[u8; 32], b: &[u8; 32], domain: &[u8; 32]| {
            use poseidon_bn254::{hash_with_domain, Fr, PrimeField};
            let a = Fr::from_bytes(a).into_option()?;
            let b = Fr::from_bytes(b).into_option()?;
            let domain = Fr::from_bytes(domain).into_option()?;
            Some(hash_with_domain(&[a, b], domain).to_repr())
        });
    });
}

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
    fn build_zktrie_db(&self, zktrie_db: &mut ZkMemoryDb) {
        init_hash_scheme();
        for (_, bytes) in self.flatten_proofs() {
            zktrie_db.add_node_bytes(bytes, None).unwrap();
        }
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
        hasher
            .update(&self.base_fee_per_gas().unwrap_or_default().to_be_bytes::<{ U256::BYTES }>());
        hasher.update(&self.gas_limit().to::<u64>().to_be_bytes());
        hasher.update(&num_txs.to_be_bytes());
    }

    /// Hash the l1 messages of the block
    #[inline]
    fn hash_l1_msg(&self, hasher: &mut impl tiny_keccak::Hasher) {
        for tx_hash in self.transactions().filter(|tx| tx.is_l1_tx()).map(|tx| tx.tx_hash()) {
            hasher.update(tx_hash.as_slice())
        }
    }
}

#[test]
fn test_parse_node() {
    use zk_trie::ZkTrieNode;
    init_hash_scheme();
    // let k = Node::<H>::hash_bytes(key).unwrap();
    // let slot = "0xf6fe0582e0e323a551f456b66350a1ed549948a34fdd35be011d229e00171c1f";//trace base
    // let slot = "f6fe0582e0e323a551f456b66350a1ed549948a34fdd35be011d229e00171c22"; // trace scale
    // --token_registry_base 53bdca72fa8d2e145a1b3bd11cde5bd75428acd18eac3d6adf4e06e7e637706d
    let slot = "0x53bdca72fa8d2e145a1b3bd11cde5bd75428acd18eac3d6adf4e06e7e637706d"; //trace2 base
    
    let slot_bytes = Bytes::from_hex(slot).unwrap().0;
    let slot_vec = slot_bytes.to_vec();
    let keccak256_storage_key = keccak256(&slot_vec);      
    println!("keccak256_storage_key: {:?}", hex::encode(&keccak256_storage_key));
    // let s_key = Node::<HashImpl>::hash_bytes(key).map_err(|e: raw::ImplError| e.to_string())?;
    // slot_vec.reverse();
    let hash = ZkTrieNode::caculate_hash(&slot_vec);
    println!("slot.hash: {:?}", hash.encode_hex_with_prefix());

    // let data =
    // "0x092d3daf332cd7de6ab5b3c6a35a4b5cce72f3d27a06651eb978fd9388f01269e806fcbf20a7b857bd27beed745f7a8d78d88e5296073cd03f10374152757dc520"
    // ;
    // let data ="0x040883d4562dfe8ba7ad8176bc2844d35ecc486199d589a152eabcee8398f3aebb05080000000000000000000000000000000000000000000000000aec000000000000000000000000000000000000000000000000000000000000000000000000000000002210b66a78475855b4b7ce7919e4a755b8bd1db84d0bcb2080d24d08923fc51444107d7c70e67e04d35d77455d5755c94d4897196e4ea9a341c1a6d864a6d4d518d32c9790b7656cb1991e831f76e6d65151cbd1c05aef8cfa7d369ad2076f84205300000000000000000000000000000000000021000000000000000000000000";
    let data = "0x040aa21e0431b54bcb95ffe603681684b4e243ffa1b430a687f471298b5c5460260101000000000000000000000000000000000000000000000000000000000000000000012035fc5993f6158720375840ee7e788ea393d23dd7e7b4b98143d0d76537902793";
    // let s = data.trim_start_matches("0x").trim_start_matches("0X");

    let proof = Bytes::from_hex(data).unwrap().0;
    let proof = proof.to_vec();
    // proof.reverse();
    // let bytes = n.to_be_bytes(); // [0x09, 0x2d]

    let n = ZkTrieNode::parse(proof.as_ref()).unwrap();
    println!("node.node_hash: {:?}", n.node_hash().encode_hex_with_prefix());
    println!("node.value_hash: {:?}", n.value_hash().unwrap().encode_hex_with_prefix());
    println!("node.node_key: {:?}", n.node_key().encode_hex_with_prefix());

    let value = n.value();
    // println!("node.value: {:?}", value);
    println!("node.value_hex: {:?}", hex::encode(&value));
}

#[test]
fn test_keccak256_node() {
    let slot = "0xf6fe0582e0e323a551f456b66350a1ed549948a34fdd35be011d229e00171c1f";
    let storage_key = keccak256(&Bytes::from_hex(slot).unwrap());
    println!("--storage_key_hex: {:?}", hex::encode(&storage_key));
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
    fn fee_limit(&self) -> u64;

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
            0x04 => {
                let tx = TxEip7702 {
                    chain_id,
                    nonce: self.nonce(),
                    gas_limit: self.gas_limit(),
                    max_fee_per_gas: self.max_fee_per_gas(),
                    max_priority_fee_per_gas: self.max_priority_fee_per_gas(),
                    to: self.to(),
                    value: self.value(),
                    access_list: self.access_list(),
                    authorization_list: self.authorization_list(),
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
            0x7f => {
                let tx = TxAltFee {
                    chain_id,
                    nonce: self.nonce(),
                    gas_limit: self.gas_limit(),
                    max_fee_per_gas: self.max_fee_per_gas(),
                    max_priority_fee_per_gas: self.max_priority_fee_per_gas(),
                    to: self.to(),
                    value: self.value(),
                    access_list: self.access_list(),
                    input: self.data(),
                    fee_token_id: self.fee_token_id(),
                    fee_limit: self.fee_limit(),
                };
                println!("tx.self.fee_token_id(): {:?}", self.fee_token_id());

                TypedTransaction::AltFee(tx.into_signed(self.signature()?))
            }
            _ => unimplemented!("unsupported tx type: {}", self.ty()),
        };

        println!("----tx.tx_hash(): {:?}", tx.tx_hash());

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

    fn authorization_list(&self) -> Vec<SignedAuthorization> {
        (*self).authorization_list()
    }

    fn signature(&self) -> Result<Signature, SignatureError> {
        (*self).signature()
    }

    fn fee_token_id(&self) -> u16 {
        (*self).fee_token_id()
    }

    fn fee_limit(&self) -> u64 {
        (*self).fee_limit()
    }
}
