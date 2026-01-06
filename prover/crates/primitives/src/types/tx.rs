use crate::{types::AuthorizationList, TxTrace};

use alloy_eips::{eip2930::AccessList, eip7702::SignedAuthorization};
use alloy_primitives::{
    normalize_v, Address, Bytes, ChainId, Signature, SignatureError, TxKind, B256, U256, U64, U8,
};
use serde_with::{serde_as, DefaultOnNull};

/// Layer1 Message Transaction
// #[derive(Clone, Debug, Default, PartialEq, Eq, Hash, Serialize, Deserialize)]
// pub struct TxL1Msg {
//     /// The 32-byte hash of the transaction.
//     pub tx_hash: B256,
//     /// The 160-bit address of the message call’s sender.
//     pub from: Address,
//     /// A scalar value equal to the number of transactions sent by the sender; formally Tn.
//     pub nonce: u64,
//     /// A scalar value equal to the maximum
//     /// amount of gas that should be used in executing
//     /// this transaction. This is paid up-front, before any
//     /// computation is done and may not be increased
//     /// later; formally Tg.
//     pub gas_limit: u64,
//     /// The 160-bit address of the message call’s recipient or, for a contract creation
//     /// transaction, ∅, used here to denote the only member of B0 ; formally Tt.
//     pub to: TxKind,
//     /// A scalar value equal to the number of Wei to
//     /// be transferred to the message call’s recipient or,
//     /// in the case of contract creation, as an endowment
//     /// to the newly created account; formally Tv.
//     pub value: U256,
//     /// Input has two uses depending if transaction is Create or Call (if `to` field is None or
//     /// Some). pub init: An unlimited size byte array specifying the
//     /// EVM-code for the account initialisation procedure CREATE,
//     /// data: An unlimited size byte array specifying the
//     /// input data of the message call, formally Td.
//     pub input: Bytes,
// }

fn default_chain_id() -> U64 {
    U64::from(53077)
}

/// Transaction Trace
#[serde_as]
#[derive(serde::Serialize, serde::Deserialize, Default, Debug, Clone, Hash, PartialEq, Eq)]
pub struct TransactionTrace {
    /// tx hash
    #[serde(default, rename = "txHash")]
    pub(crate) tx_hash: B256,
    /// tx type (in raw from)
    #[serde(rename = "type")]
    pub(crate) ty: U8,
    /// nonce
    pub(crate) nonce: U64,
    /// gas limit
    pub(crate) gas: U64,
    #[serde(rename = "gasPrice")]
    /// gas price
    pub(crate) gas_price: U256,
    #[serde(rename = "gasTipCap", alias = "maxPriorityFeePerGas")]
    /// gas tip cap
    pub(crate) gas_tip_cap: Option<U256>,
    #[serde(rename = "gasFeeCap", alias = "maxFeePerGas")]
    /// gas fee cap
    pub(crate) gas_fee_cap: Option<U256>,
    /// from
    pub(crate) from: Address,
    /// to, NONE for creation (0 addr)
    pub(crate) to: Option<Address>,
    /// chain id
    #[serde(rename = "chainId")]
    #[serde(default = "default_chain_id")]
    pub(crate) chain_id: U64,
    /// value amount
    pub(crate) value: U256,
    /// call data
    #[serde(alias = "input")]
    pub(crate) data: Bytes,
    /// is creation
    #[serde(rename = "isCreate")]
    #[serde(default)]
    pub(crate) is_create: bool,
    /// access list
    #[serde(rename = "accessList")]
    #[serde(default)]
    #[serde_as(as = "DefaultOnNull")]
    pub(crate) access_list: AccessList,
    /// authorization list
    #[serde(rename = "authorizationList")]
    #[serde(default)]
    #[serde_as(as = "DefaultOnNull")]
    pub(crate) authorization_list: AuthorizationList,
    /// For AltFeeType
    #[serde(rename = "feeTokenID")]
    pub(crate) fee_token_id: Option<u16>,
    /// For AltFeeType
    #[serde(rename = "feeLimit")]
    pub(crate) fee_limit: Option<U256>,
    /// signature v
    pub(crate) v: U64,
    /// signature r
    pub(crate) r: U256,
    /// signature s
    pub(crate) s: U256,
}

impl TxTrace for TransactionTrace {
    fn tx_hash(&self) -> B256 {
        self.tx_hash
    }

    fn ty(&self) -> u8 {
        self.ty.to::<u8>()
    }

    fn nonce(&self) -> u64 {
        self.nonce.to::<u64>()
    }

    fn gas_limit(&self) -> u64 {
        self.gas.to::<u64>()
    }

    fn gas_price(&self) -> u128 {
        self.gas_price.to()
    }

    fn max_fee_per_gas(&self) -> u128 {
        self.gas_fee_cap.map(|v| v.to()).unwrap_or_default()
    }

    fn max_priority_fee_per_gas(&self) -> u128 {
        self.gas_tip_cap.map(|v| v.to()).unwrap_or_default()
    }

    unsafe fn get_from_unchecked(&self) -> Address {
        self.from
    }

    fn to(&self) -> TxKind {
        if self.is_create {
            TxKind::Create
        } else {
            debug_assert!(self.to.map(|a| !a.is_zero()).unwrap_or(false));
            TxKind::Call(self.to.expect("to address must be present"))
        }
    }

    fn chain_id(&self) -> ChainId {
        self.chain_id.to()
    }

    fn value(&self) -> U256 {
        self.value
    }

    fn data(&self) -> Bytes {
        self.data.clone()
    }

    fn access_list(&self) -> AccessList {
        self.access_list.clone()
    }

    fn authorization_list(&self) -> Vec<SignedAuthorization> {
        self.authorization_list.clone().into()
    }

    fn signature(&self) -> Result<Signature, SignatureError> {
        Ok(Signature::from_scalars_and_parity(
            self.r.into(),
            self.s.into(),
            normalize_v(self.v.to::<u64>()).unwrap_or_default(),
        ))
    }

    fn fee_token_id(&self) -> u16 {
        self.fee_token_id.unwrap_or_default()
    }

    fn fee_limit(&self) -> U256 {
        self.fee_limit.unwrap_or_default().to()
    }

    fn sig_v(&self) -> u64 {
        self.v.to::<u64>()
    }
}

// impl Typed2718 for TypedTransaction {
//     fn ty(&self) -> u8 {
//         0x7e
//     }
// }

// impl Transaction for TypedTransaction {
//     fn chain_id(&self) -> Option<ChainId> {
//         match self {
//             TypedTransaction::Enveloped(tx) => tx.chain_id(),
//             TypedTransaction::L1Msg(tx) => tx.chain_id(),
//             TypedTransaction::AltFee(tx) => Some(tx.tx().chain_id),
//         }
//     }

//     fn nonce(&self) -> u64 {
//         match self {
//             TypedTransaction::Enveloped(tx) => tx.nonce(),
//             TypedTransaction::L1Msg(tx) => tx.nonce(),
//             TypedTransaction::AltFee(tx) => tx.tx().nonce(),
//         }
//     }

//     fn gas_limit(&self) -> u64 {
//         match self {
//             TypedTransaction::Enveloped(tx) => tx.gas_limit(),
//             TypedTransaction::L1Msg(tx) => tx.gas_limit(),
//             TypedTransaction::AltFee(tx) => tx.tx().gas_limit(),
//         }
//     }

//     fn gas_price(&self) -> Option<u128> {
//         match self {
//             TypedTransaction::Enveloped(tx) => tx.gas_price(),
//             TypedTransaction::L1Msg(tx) => tx.gas_price(),
//             TypedTransaction::AltFee(tx) => tx.tx().gas_price(),
//         }
//     }

//     fn max_fee_per_gas(&self) -> u128 {
//         match self {
//             TypedTransaction::Enveloped(tx) => tx.max_fee_per_gas(),
//             TypedTransaction::L1Msg(tx) => tx.max_fee_per_gas(),
//             TypedTransaction::AltFee(tx) => tx.tx().max_fee_per_gas(),
//         }
//     }

//     fn max_priority_fee_per_gas(&self) -> Option<u128> {
//         match self {
//             TypedTransaction::Enveloped(tx) => tx.max_priority_fee_per_gas(),
//             TypedTransaction::L1Msg(tx) => tx.max_priority_fee_per_gas(),
//             TypedTransaction::AltFee(tx) => tx.tx().max_priority_fee_per_gas(),
//         }
//     }

//     fn max_fee_per_blob_gas(&self) -> Option<u128> {
//         match self {
//             TypedTransaction::Enveloped(tx) => tx.max_fee_per_blob_gas(),
//             TypedTransaction::L1Msg(tx) => tx.max_fee_per_blob_gas(),
//             TypedTransaction::AltFee(tx) => tx.tx().max_fee_per_blob_gas(),
//         }
//     }

//     fn priority_fee_or_price(&self) -> u128 {
//         match self {
//             TypedTransaction::Enveloped(tx) => tx.priority_fee_or_price(),
//             TypedTransaction::L1Msg(tx) => tx.priority_fee_or_price(),
//             TypedTransaction::AltFee(tx) => tx.tx().priority_fee_or_price(),
//         }
//     }

//     fn to(&self) -> Option<Address> {
//         match self {
//             TypedTransaction::Enveloped(tx) => tx.to(),
//             TypedTransaction::L1Msg(tx) => tx.to(),
//             TypedTransaction::AltFee(tx) => tx.tx().to(),
//         }
//     }

//     fn value(&self) -> U256 {
//         match self {
//             TypedTransaction::Enveloped(tx) => tx.value(),
//             TypedTransaction::L1Msg(tx) => tx.value(),
//             TypedTransaction::AltFee(tx) => tx.tx().value(),
//         }
//     }

//     fn input(&self) -> &alloy_primitives::Bytes {
//         match self {
//             TypedTransaction::Enveloped(tx) => tx.input(),
//             TypedTransaction::L1Msg(tx) => tx.input(),
//             TypedTransaction::AltFee(tx) => tx.tx().input(),
//         }
//     }

//     // fn ty(&self) -> u8 {
//     //     match self {
//     //         TypedTransaction::Enveloped(tx) => tx.ty(),
//     //         TypedTransaction::L1Msg(tx) => tx.ty(),
//     //         TypedTransaction::AltFee(tx) => tx.tx().ty(),
//     //     }
//     // }

//     fn access_list(&self) -> Option<&AccessList> {
//         match self {
//             TypedTransaction::Enveloped(tx) => tx.access_list(),
//             TypedTransaction::L1Msg(tx) => tx.access_list(),
//             TypedTransaction::AltFee(tx) => tx.tx().access_list(),
//         }
//     }

//     fn blob_versioned_hashes(&self) -> Option<&[B256]> {
//         match self {
//             TypedTransaction::Enveloped(tx) => tx.blob_versioned_hashes(),
//             TypedTransaction::L1Msg(tx) => tx.blob_versioned_hashes(),
//             TypedTransaction::AltFee(tx) => tx.tx().blob_versioned_hashes(),
//         }
//     }

//     fn authorization_list(&self) -> Option<&[SignedAuthorization]> {
//         match self {
//             TypedTransaction::Enveloped(tx) => tx.authorization_list(),
//             TypedTransaction::L1Msg(_) => None,
//             TypedTransaction::AltFee(_) => None,
//         }
//     }

//     fn effective_gas_price(&self, _base_fee: Option<u64>) -> u128 {
//         self.gas_price().unwrap_or_default()
//     }

//     fn is_dynamic_fee(&self) -> bool {
//         self.max_fee_per_gas() != 0
//     }

//     fn kind(&self) -> TxKind {
//         self.to().into()
//     }

//     fn is_create(&self) -> bool {
//         self.kind().is_create()
//     }

//     fn effective_tip_per_gas(&self, base_fee: u64) -> Option<u128> {
//         let base_fee = base_fee as u128;
//         let max_fee_per_gas = self.max_fee_per_gas();
//         if max_fee_per_gas < base_fee {
//             return None;
//         }
//         let fee = max_fee_per_gas - base_fee;
//         self.max_priority_fee_per_gas()
//             .map_or(Some(fee), |priority_fee| Some(fee.min(priority_fee)))
//     }

//     fn function_selector(&self) -> Option<&Selector> {
//         if self.kind().is_call() {
//             self.input().get(..4).and_then(|s| TryFrom::try_from(s).ok())
//         } else {
//             None
//         }
//     }

//     fn blob_count(&self) -> Option<u64> {
//         self.blob_versioned_hashes().map(|h| h.len() as u64)
//     }

//     #[inline]
//     fn blob_gas_used(&self) -> Option<u64> {
//         self.blob_count().map(|blobs| blobs * DATA_GAS_PER_BLOB)
//     }

//     fn authorization_count(&self) -> Option<u64> {
//         self.authorization_list().map(|auths| auths.len() as u64)
//     }
// }

// impl TxL1Msg {
//     /// Outputs the length of the transaction's fields.
//     #[doc(hidden)]
//     pub fn fields_len(&self) -> usize {
//         let mut len = 0;
//         len += self.nonce.length();
//         len += self.gas_limit.length();
//         len += self.to.length();
//         len += self.value.length();
//         len += self.input.0.length();
//         len += self.from.length();
//         len
//     }
// }

// impl Transaction for TxL1Msg {
//     fn chain_id(&self) -> Option<ChainId> {
//         None
//     }

//     fn nonce(&self) -> u64 {
//         self.nonce
//     }

//     fn gas_limit(&self) -> u64 {
//         self.gas_limit as u64
//     }

//     fn gas_price(&self) -> Option<u128> {
//         Some(0)
//     }

//     fn max_fee_per_gas(&self) -> u128 {
//         0
//     }

//     fn max_priority_fee_per_gas(&self) -> Option<u128> {
//         None
//     }

//     fn max_fee_per_blob_gas(&self) -> Option<u128> {
//         None
//     }

//     fn priority_fee_or_price(&self) -> u128 {
//         0
//     }

//     fn to(&self) -> Option<Address> {
//         self.to.to().copied()
//     }

//     fn value(&self) -> U256 {
//         self.value
//     }

//     fn input(&self) -> &alloy_primitives::Bytes {
//         &self.input
//     }

//     // fn ty(&self) -> u8 {
//     //     0x7e
//     // }

//     fn access_list(&self) -> Option<&AccessList> {
//         None
//     }

//     fn blob_versioned_hashes(&self) -> Option<&[B256]> {
//         None
//     }

//     fn authorization_list(&self) -> Option<&[SignedAuthorization]> {
//         None
//     }

//     fn effective_tip_per_gas(&self, base_fee: u64) -> Option<u128> {
//         let base_fee = base_fee as u128;
//         let max_fee_per_gas = self.max_fee_per_gas();
//         if max_fee_per_gas < base_fee {
//             return None;
//         }
//         let fee = max_fee_per_gas - base_fee;
//         self.max_priority_fee_per_gas()
//             .map_or(Some(fee), |priority_fee| Some(fee.min(priority_fee)))
//     }

//     fn function_selector(&self) -> Option<&Selector> {
//         if self.kind().is_call() {
//             self.input().get(..4).and_then(|s| TryFrom::try_from(s).ok())
//         } else {
//             None
//         }
//     }

//     fn blob_count(&self) -> Option<u64> {
//         self.blob_versioned_hashes().map(|h| h.len() as u64)
//     }

//     #[inline]
//     fn blob_gas_used(&self) -> Option<u64> {
//         self.blob_count().map(|blobs| blobs * DATA_GAS_PER_BLOB)
//     }

//     fn authorization_count(&self) -> Option<u64> {
//         self.authorization_list().map(|auths| auths.len() as u64)
//     }

//     fn effective_gas_price(&self, _base_fee: Option<u64>) -> u128 {
//         todo!()
//     }

//     fn is_dynamic_fee(&self) -> bool {
//         todo!()
//     }

//     fn kind(&self) -> TxKind {
//         todo!()
//     }

//     fn is_create(&self) -> bool {
//         todo!()
//     }
// }

// impl Encodable for TxL1Msg {
//     fn encode(&self, out: &mut dyn BufMut) {
//         self.nonce.encode(out);
//         self.gas_limit.encode(out);
//         self.to.encode(out);
//         self.value.encode(out);
//         self.input.0.encode(out);
//         self.from.encode(out);
//     }
// }

// impl Typed2718 for TxL1Msg {
//     fn ty(&self) -> u8 {
//         0x7e
//     }
// }

// impl Encodable2718 for TxL1Msg {
//     fn type_flag(&self) -> Option<u8> {
//         Some(0x7e)
//     }

//     fn encode_2718_len(&self) -> usize {
//         let payload_length = self.fields_len();
//         1 + Header { list: true, payload_length }.length() + payload_length
//     }

//     fn encode_2718(&self, out: &mut dyn BufMut) {
//         0x7eu8.encode(out);
//         let header = Header { list: true, payload_length: self.fields_len() };
//         header.encode(out);
//         self.encode(out)
//     }

//     fn encoded_2718(&self) -> Vec<u8> {
//         let mut out = Vec::with_capacity(self.encode_2718_len());
//         self.encode_2718(&mut out);
//         out
//     }
// }

// impl TypedTransaction {

//     /// Return the hash of the inner transaction.
//     pub fn tx_hash(&self) -> B256 {
//         match self {
//             TypedTransaction::Enveloped(tx) => *tx.tx_hash(),
//             TypedTransaction::L1Msg(tx) => tx.tx_hash,
//             TypedTransaction::AltFee(tx) => {
//                 let mut bytes = BytesMut::new();
//                 tx.tx().encode_2718(tx.signature(), &mut bytes);
//                 keccak256(&bytes)
//             }
//         }
//     }

//     /// Get the caller of the transaction, recover the signer if the transaction is enveloped.
//     ///
//     /// Fails if the transaction is enveloped and recovering the signer fails.
//     pub fn get_or_recover_signer(&self) -> Result<Address, SignatureError> {
//         match self {
//             TypedTransaction::Enveloped(tx) => {
//                 tx.recover_signer().map_err(|_| SignatureError::InvalidParity(0))
//             }
//             TypedTransaction::L1Msg(tx) => Ok(tx.from),
//             TypedTransaction::AltFee(tx) => tx.recover_signer(),
//         }
//     }

//     /// Get the effective gas price of the transaction.
//     pub fn effective_gas_price(&self, base_fee_per_gas: u64) -> Option<u128> {
//         match self {
//             TypedTransaction::Enveloped(TxEnvelope::Eip1559(ref tx)) => {
//                 let priority_fee_per_gas = tx.tx().effective_tip_per_gas(base_fee_per_gas)?;
//                 Some(priority_fee_per_gas + base_fee_per_gas as u128)
//             }
//             TypedTransaction::Enveloped(TxEnvelope::Eip7702(ref tx)) => {
//                 let priority_fee_per_gas = tx.tx().effective_tip_per_gas(base_fee_per_gas)?;
//                 Some(priority_fee_per_gas + base_fee_per_gas as u128)
//             }
//             TypedTransaction::AltFee(tx) => {
//                 let priority_fee_per_gas = tx.tx().effective_tip_per_gas(base_fee_per_gas)?;
//                 Some(priority_fee_per_gas + base_fee_per_gas as u128)
//             }
//             _ => self.gas_price(),
//         }
//     }

//     /// Encode the transaction according to [EIP-2718] rules. First a 1-byte
//     /// type flag in the range 0x0-0x7f, then the body of the transaction.
//     pub fn rlp(&self) -> Bytes {
//         let mut bytes = BytesMut::new();
//         match self {
//             TypedTransaction::Enveloped(tx) => tx.encode_2718(&mut bytes),
//             TypedTransaction::L1Msg(tx) => tx.encode_2718(&mut bytes),
//             TypedTransaction::AltFee(tx) => tx.tx().encode_2718(tx.signature(), &mut bytes),
//         }
//         Bytes(bytes.freeze())
//     }

//     /// Calculate the signing hash for the transaction.
//     pub fn signature_hash(&self) -> B256 {
//         match self {
//             TypedTransaction::Enveloped(tx) => tx.signature_hash(),
//             TypedTransaction::L1Msg(_) => keccak256(self.rlp()),
//             TypedTransaction::AltFee(tx) => tx.signature_hash(),
//         }
//     }

//     /// Get `data`
//     pub fn data(&self) -> Bytes {
//         match self {
//             TypedTransaction::Enveloped(tx) => match tx.tx_type() {
//                 TxType::Legacy => tx.as_legacy().unwrap().tx().input.clone(),
//                 TxType::Eip1559 => tx.as_eip1559().unwrap().tx().input.clone(),
//                 TxType::Eip2930 => tx.as_eip2930().unwrap().tx().input.clone(),
//                 TxType::Eip7702 => tx.as_eip7702().unwrap().tx().input.clone(),
//                 _ => unimplemented!("unsupported tx type {:?}", tx.tx_type()),
//             },
//             TypedTransaction::L1Msg(tx) => tx.input.clone(),
//             TypedTransaction::AltFee(tx) => tx.tx().input.clone(),
//         }
//     }

//     /// Check if the transaction is an L1 transaction
//     pub fn is_l1_msg(&self) -> bool {
//         matches!(self, TypedTransaction::L1Msg(_))
//     }

//     /// Returns the fee token ID if this is an AltFee transaction, otherwise None.
//     pub fn fee_token_id(&self) -> Option<u16> {
//         match self {
//             TypedTransaction::Enveloped(_) => None,
//             TypedTransaction::L1Msg(_) => None,
//             TypedTransaction::AltFee(tx) => Some(tx.tx().fee_token_id),
//         }
//     }

//     /// Returns the fee limit if this is an AltFee transaction, otherwise None.
//     pub fn fee_limit(&self) -> Option<U256> {
//         match self {
//             TypedTransaction::Enveloped(_) => None,
//             TypedTransaction::L1Msg(_) => None,
//             TypedTransaction::AltFee(tx) => Some(tx.tx().fee_limit),
//         }
//     }
// }

// /// Get a TypedTransaction directly from a rlp encoded byte stream
// impl Decodable for TypedTransaction {
//     fn decode(buf: &mut &[u8]) -> alloy_rlp::Result<Self> {
//         if buf.is_empty() {
//             return Err(alloy_rlp::Error::InputTooShort);
//         }
//         let tx_type = *buf.first().unwrap_or(&0u8);
//         match tx_type {
//             0x7f => {
//                 return Ok(TypedTransaction::AltFee(
//                     TxAltFee::decode_signed_fields(&mut &buf[1..])
//                         .map_err(|_| alloy_rlp::Error::Custom("decode TxAltFee error"))?,
//                 ))
//             }
//             _ => return Ok(TypedTransaction::Enveloped(TxEnvelope::decode_2718(buf).unwrap())),
//         };
//     }
// }
