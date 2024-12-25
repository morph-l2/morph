use crate::TxTrace;
use alloy::{
    consensus::{Transaction, TxEnvelope, TxType, Typed2718},
    eips::{
        eip2718::{Decodable2718, Encodable2718},
        eip2930::AccessList,
        eip7702::SignedAuthorization,
    },
    primitives::{
        keccak256, normalize_v, Address, Bytes, ChainId, PrimitiveSignature, SignatureError,
        TxKind, B256, U256, U64,
    },
    rlp::{BufMut, BytesMut, Decodable, Encodable, Header},
};
use serde::{Deserialize, Serialize};
use serde_with::{serde_as, DefaultOnNull};

/// Wrapped Ethereum Transaction
#[derive(Clone, Debug, PartialEq, Eq)]
pub enum TypedTransaction {
    /// Normal enveloped ethereum transaction
    Enveloped(TxEnvelope),
    /// Layer1 Message Transaction
    L1Msg(TxL1Msg),
}

/// Layer1 Message Transaction
#[derive(Clone, Debug, Default, PartialEq, Eq, Hash)]
pub struct TxL1Msg {
    /// The 32-byte hash of the transaction.
    pub tx_hash: B256,
    /// The 160-bit address of the message call’s sender.
    pub from: Address,
    /// A scalar value equal to the number of transactions sent by the sender; formally Tn.
    pub nonce: u64,
    /// A scalar value equal to the maximum
    /// amount of gas that should be used in executing
    /// this transaction. This is paid up-front, before any
    /// computation is done and may not be increased
    /// later; formally Tg.
    pub gas_limit: u128,
    /// The 160-bit address of the message call’s recipient or, for a contract creation
    /// transaction, ∅, used here to denote the only member of B0 ; formally Tt.
    pub to: TxKind,
    /// A scalar value equal to the number of Wei to
    /// be transferred to the message call’s recipient or,
    /// in the case of contract creation, as an endowment
    /// to the newly created account; formally Tv.
    pub value: U256,
    /// Input has two uses depending if transaction is Create or Call (if `to` field is None or
    /// Some). pub init: An unlimited size byte array specifying the
    /// EVM-code for the account initialisation procedure CREATE,
    /// data: An unlimited size byte array specifying the
    /// input data of the message call, formally Td.
    pub input: Bytes,
}

/// Transaction Trace
#[serde_as]
#[derive(Serialize, Deserialize, Debug, Clone, Hash, PartialEq, Eq)]
pub struct TransactionTrace {
    /// tx hash
    #[serde(default, rename = "txHash")]
    pub(crate) tx_hash: B256,
    /// tx type (in raw from)
    #[serde(rename = "type")]
    pub(crate) ty: u8,
    /// nonce
    pub(crate) nonce: u64,
    /// gas limit
    pub(crate) gas: u64,
    #[serde(rename = "gasPrice")]
    /// gas price
    pub(crate) gas_price: U256,
    #[serde(rename = "gasTipCap")]
    /// gas tip cap
    pub(crate) gas_tip_cap: Option<U256>,
    #[serde(rename = "gasFeeCap")]
    /// gas fee cap
    pub(crate) gas_fee_cap: Option<U256>,
    /// from
    pub(crate) from: Address,
    /// to, NONE for creation (0 addr)
    pub(crate) to: Option<Address>,
    /// chain id
    #[serde(rename = "chainId")]
    pub(crate) chain_id: U64,
    /// value amount
    pub(crate) value: U256,
    /// call data
    pub(crate) data: Bytes,
    /// is creation
    #[serde(rename = "isCreate")]
    pub(crate) is_create: bool,
    /// access list
    #[serde(rename = "accessList")]
    #[serde_as(as = "DefaultOnNull")]
    pub(crate) access_list: AccessList,
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
        self.ty
    }

    fn nonce(&self) -> u64 {
        self.nonce
    }

    fn gas_limit(&self) -> u128 {
        self.gas as u128
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

    fn signature(&self) -> Result<PrimitiveSignature, SignatureError> {
        Ok(PrimitiveSignature::from_scalars_and_parity(
            self.r.into(),
            self.s.into(),
            normalize_v(self.v.to::<u64>()).unwrap_or_default(),
        ))
    }
}

impl Typed2718 for TypedTransaction {
    #[doc = " Returns the EIP-2718 type flag."]
    fn ty(&self) -> u8 {
        todo!()
    }

    #[doc = " Returns true if the type matches the given type."]
    fn is_type(&self, ty: u8) -> bool {
        self.ty() == ty
    }

    #[doc = " Returns true if the type is a legacy transaction."]
    fn is_legacy(&self) -> bool {
        self.ty() == 0
    }

    #[doc = " Returns true if the type is an EIP-2930 transaction."]
    fn is_eip2930(&self) -> bool {
        self.ty() == 1
    }

    #[doc = " Returns true if the type is an EIP-1559 transaction."]
    fn is_eip1559(&self) -> bool {
        self.ty() == 2
    }

    #[doc = " Returns true if the type is an EIP-4844 transaction."]
    fn is_eip4844(&self) -> bool {
        self.ty() == 3
    }

    #[doc = " Returns true if the type is an EIP-7702 transaction."]
    fn is_eip7702(&self) -> bool {
        self.ty() == 4
    }
}

impl Transaction for TypedTransaction {
    fn chain_id(&self) -> Option<ChainId> {
        match self {
            TypedTransaction::Enveloped(tx) => tx.chain_id(),
            TypedTransaction::L1Msg(tx) => tx.chain_id(),
        }
    }

    fn nonce(&self) -> u64 {
        match self {
            TypedTransaction::Enveloped(tx) => tx.nonce(),
            TypedTransaction::L1Msg(tx) => tx.nonce(),
        }
    }

    fn gas_limit(&self) -> u64 {
        match self {
            TypedTransaction::Enveloped(tx) => tx.gas_limit(),
            TypedTransaction::L1Msg(tx) => tx.gas_limit(),
        }
    }

    fn gas_price(&self) -> Option<u128> {
        match self {
            TypedTransaction::Enveloped(tx) => tx.gas_price(),
            TypedTransaction::L1Msg(tx) => tx.gas_price(),
        }
    }

    fn max_fee_per_gas(&self) -> u128 {
        match self {
            TypedTransaction::Enveloped(tx) => tx.max_fee_per_gas(),
            TypedTransaction::L1Msg(tx) => tx.max_fee_per_gas(),
        }
    }

    fn max_priority_fee_per_gas(&self) -> Option<u128> {
        match self {
            TypedTransaction::Enveloped(tx) => tx.max_priority_fee_per_gas(),
            TypedTransaction::L1Msg(tx) => tx.max_priority_fee_per_gas(),
        }
    }

    fn max_fee_per_blob_gas(&self) -> Option<u128> {
        match self {
            TypedTransaction::Enveloped(tx) => tx.max_fee_per_blob_gas(),
            TypedTransaction::L1Msg(tx) => tx.max_fee_per_blob_gas(),
        }
    }

    fn priority_fee_or_price(&self) -> u128 {
        match self {
            TypedTransaction::Enveloped(tx) => tx.priority_fee_or_price(),
            TypedTransaction::L1Msg(tx) => tx.priority_fee_or_price(),
        }
    }

    fn to(&self) -> Option<Address> {
        match self {
            TypedTransaction::Enveloped(tx) => tx.to(),
            TypedTransaction::L1Msg(tx) => tx.to(),
        }
    }

    fn value(&self) -> U256 {
        match self {
            TypedTransaction::Enveloped(tx) => tx.value(),
            TypedTransaction::L1Msg(tx) => tx.value(),
        }
    }

    fn input(&self) -> &Bytes {
        match self {
            TypedTransaction::Enveloped(tx) => tx.input(),
            TypedTransaction::L1Msg(tx) => tx.input(),
        }
    }

    fn access_list(&self) -> Option<&AccessList> {
        match self {
            TypedTransaction::Enveloped(tx) => tx.access_list(),
            TypedTransaction::L1Msg(tx) => tx.access_list(),
        }
    }

    fn blob_versioned_hashes(&self) -> Option<&[B256]> {
        match self {
            TypedTransaction::Enveloped(tx) => tx.blob_versioned_hashes(),
            TypedTransaction::L1Msg(tx) => tx.blob_versioned_hashes(),
        }
    }

    fn authorization_list(&self) -> Option<&[SignedAuthorization]> {
        match self {
            TypedTransaction::Enveloped(tx) => tx.authorization_list(),
            TypedTransaction::L1Msg(tx) => tx.authorization_list(),
        }
    }

    #[doc = " Returns the effective gas price for the given base fee."]
    #[doc = ""]
    #[doc = " If the transaction is a legacy or EIP2930 transaction, the gas price is returned."]
    fn effective_gas_price(&self, base_fee: Option<u64>) -> u128 {
        todo!()
    }

    #[doc = " Returns `true` if the transaction supports dynamic fees."]
    fn is_dynamic_fee(&self) -> bool {
        todo!()
    }

    #[doc = " Returns the transaction kind."]
    fn kind(&self) -> TxKind {
        todo!()
    }

    #[doc = " Returns true if the transaction is a contract creation."]
    #[doc = " We don\'t provide a default implementation via `kind` as it copies the 21-byte"]
    #[doc = " [`TxKind`] for this simple check. A proper implementation shouldn\'t allocate."]
    fn is_create(&self) -> bool {
        todo!()
    }

    #[doc = " Returns the effective tip for this transaction."]
    #[doc = ""]
    #[doc = " For EIP-1559 transactions: `min(max_fee_per_gas - base_fee, max_priority_fee_per_gas)`."]
    #[doc = " For legacy transactions: `gas_price - base_fee`."]
    fn effective_tip_per_gas(&self, base_fee: u64) -> Option<u128> {
        let base_fee = base_fee as u128;
        let max_fee_per_gas = self.max_fee_per_gas();
        if max_fee_per_gas < base_fee {
            return None;
        }
        let fee = max_fee_per_gas - base_fee;
        self.max_priority_fee_per_gas()
            .map_or(Some(fee), |priority_fee| Some(fee.min(priority_fee)))
    }

    #[doc = " Returns the total gas for all blobs in this transaction."]
    #[doc = ""]
    #[doc = " Returns `None` for non-eip4844 transactions."]
    #[inline]
    fn blob_gas_used(&self) -> Option<u64> {
        // self.blob_versioned_hashes().map(|blobs| blobs.len() as u64 * DATA_GAS_PER_BLOB)
        todo!()
    }
}

impl TxL1Msg {
    /// Outputs the length of the transaction's fields.
    #[doc(hidden)]
    pub fn fields_len(&self) -> usize {
        let mut len = 0;
        len += self.nonce.length();
        len += self.gas_limit.length();
        len += self.to.length();
        len += self.value.length();
        len += self.input.0.length();
        len += self.from.length();
        len
    }
}

impl Typed2718 for TxL1Msg {
    #[doc = " Returns the EIP-2718 type flag."]
    fn ty(&self) -> u8 {
        todo!()
    }

    #[doc = " Returns true if the type matches the given type."]
    fn is_type(&self, ty: u8) -> bool {
        self.ty() == ty
    }

    #[doc = " Returns true if the type is a legacy transaction."]
    fn is_legacy(&self) -> bool {
        self.ty() == 0
    }

    #[doc = " Returns true if the type is an EIP-2930 transaction."]
    fn is_eip2930(&self) -> bool {
        self.ty() == 1
    }

    #[doc = " Returns true if the type is an EIP-1559 transaction."]
    fn is_eip1559(&self) -> bool {
        self.ty() == 2
    }

    #[doc = " Returns true if the type is an EIP-4844 transaction."]
    fn is_eip4844(&self) -> bool {
        self.ty() == 3
    }

    #[doc = " Returns true if the type is an EIP-7702 transaction."]
    fn is_eip7702(&self) -> bool {
        self.ty() == 4
    }
}

impl Transaction for TxL1Msg {
    fn chain_id(&self) -> Option<ChainId> {
        None
    }

    fn nonce(&self) -> u64 {
        self.nonce
    }

    fn gas_limit(&self) -> u64 {
        self.gas_limit as u64
    }

    fn gas_price(&self) -> Option<u128> {
        Some(0)
    }

    fn max_fee_per_gas(&self) -> u128 {
        0
    }

    fn max_priority_fee_per_gas(&self) -> Option<u128> {
        None
    }

    fn max_fee_per_blob_gas(&self) -> Option<u128> {
        None
    }

    fn priority_fee_or_price(&self) -> u128 {
        0
    }

    fn to(&self) -> Option<Address> {
        self.to.to().cloned()
    }

    fn value(&self) -> U256 {
        self.value
    }

    fn input(&self) -> &Bytes {
        &self.input
    }

    fn access_list(&self) -> Option<&AccessList> {
        None
    }

    fn blob_versioned_hashes(&self) -> Option<&[B256]> {
        None
    }

    fn authorization_list(&self) -> Option<&[SignedAuthorization]> {
        None
    }

    #[doc = " Returns the effective gas price for the given base fee."]
    #[doc = ""]
    #[doc = " If the transaction is a legacy or EIP2930 transaction, the gas price is returned."]
    fn effective_gas_price(&self, base_fee: Option<u64>) -> u128 {
        todo!()
    }

    #[doc = " Returns `true` if the transaction supports dynamic fees."]
    fn is_dynamic_fee(&self) -> bool {
        todo!()
    }

    #[doc = " Returns the transaction kind."]
    fn kind(&self) -> TxKind {
        todo!()
    }

    #[doc = " Returns true if the transaction is a contract creation."]
    #[doc = " We don\'t provide a default implementation via `kind` as it copies the 21-byte"]
    #[doc = " [`TxKind`] for this simple check. A proper implementation shouldn\'t allocate."]
    fn is_create(&self) -> bool {
        todo!()
    }

    #[doc = " Returns the effective tip for this transaction."]
    #[doc = ""]
    #[doc = " For EIP-1559 transactions: `min(max_fee_per_gas - base_fee, max_priority_fee_per_gas)`."]
    #[doc = " For legacy transactions: `gas_price - base_fee`."]
    fn effective_tip_per_gas(&self, base_fee: u64) -> Option<u128> {
        todo!()
    }

    #[doc = " Returns the total gas for all blobs in this transaction."]
    #[doc = ""]
    #[doc = " Returns `None` for non-eip4844 transactions."]
    #[inline]
    fn blob_gas_used(&self) -> Option<u64> {
        todo!()
    }
}

impl Encodable for TxL1Msg {
    fn encode(&self, out: &mut dyn BufMut) {
        self.nonce.encode(out);
        self.gas_limit.encode(out);
        self.to.encode(out);
        self.value.encode(out);
        self.input.0.encode(out);
        self.from.encode(out);
    }
}

impl Encodable2718 for TxL1Msg {
    fn type_flag(&self) -> Option<u8> {
        Some(0x7e)
    }

    fn encode_2718_len(&self) -> usize {
        let payload_length = self.fields_len();
        1 + Header { list: true, payload_length }.length() + payload_length
    }

    fn encode_2718(&self, out: &mut dyn BufMut) {
        0x7eu8.encode(out);
        let header = Header { list: true, payload_length: self.fields_len() };
        header.encode(out);
        self.encode(out)
    }
}

impl TypedTransaction {
    /// Return the hash of the inner transaction.
    pub fn tx_hash(&self) -> &B256 {
        match self {
            TypedTransaction::Enveloped(tx) => tx.tx_hash(),
            TypedTransaction::L1Msg(tx) => &tx.tx_hash,
        }
    }

    /// Get the caller of the transaction, recover the signer if the transaction is enveloped.
    ///
    /// Fails if the transaction is enveloped and recovering the signer fails.
    pub fn get_or_recover_signer(&self) -> Result<Address, SignatureError> {
        match self {
            TypedTransaction::Enveloped(tx) => tx.recover_signer(),
            TypedTransaction::L1Msg(tx) => Ok(tx.from),
        }
    }

    /// Get the effective gas price of the transaction.
    pub fn effective_gas_price(&self, base_fee_per_gas: u64) -> Option<u128> {
        match self {
            TypedTransaction::Enveloped(TxEnvelope::Eip1559(ref tx)) => {
                let priority_fee_per_gas = tx.tx().effective_tip_per_gas(base_fee_per_gas)?;
                Some(priority_fee_per_gas + base_fee_per_gas as u128)
            }
            _ => self.gas_price(),
        }
    }

    /// Encode the transaction according to [EIP-2718] rules. First a 1-byte
    /// type flag in the range 0x0-0x7f, then the body of the transaction.
    pub fn rlp(&self) -> Bytes {
        let mut bytes = BytesMut::new();
        match self {
            TypedTransaction::Enveloped(tx) => tx.encode_2718(&mut bytes),
            TypedTransaction::L1Msg(tx) => tx.encode_2718(&mut bytes),
        }
        Bytes(bytes.freeze())
    }

    /// Calculate the signing hash for the transaction.
    pub fn signature_hash(&self) -> B256 {
        match self {
            TypedTransaction::Enveloped(tx) => tx.signature_hash(),
            TypedTransaction::L1Msg(_) => keccak256(self.rlp()),
        }
    }

    /// Get `data`
    pub fn data(&self) -> Bytes {
        match self {
            TypedTransaction::Enveloped(tx) => match tx.tx_type() {
                TxType::Legacy => tx.as_legacy().unwrap().tx().input.clone(),
                TxType::Eip1559 => tx.as_eip1559().unwrap().tx().input.clone(),
                TxType::Eip2930 => tx.as_eip2930().unwrap().tx().input.clone(),
                _ => unimplemented!("unsupported tx type {:?}", tx.tx_type()),
            },
            TypedTransaction::L1Msg(tx) => tx.input.clone(),
        }
    }

    /// Check if the transaction is an L1 transaction
    pub fn is_l1_msg(&self) -> bool {
        matches!(self, TypedTransaction::L1Msg(_))
    }
}

/// Get a TypedTransaction directly from a rlp encoded byte stream
impl Decodable for TypedTransaction {
    fn decode(buf: &mut &[u8]) -> alloy::rlp::Result<Self> {
        if buf.is_empty() {
            return Err(alloy::rlp::Error::InputTooShort);
        }
        Ok(TypedTransaction::Enveloped(TxEnvelope::decode_2718(buf).unwrap()))
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    const TRACE: &str = include_str!("../../../../testdata/dev.json");

    #[test]
    fn test_transaction_trace_deserialize() {
        let trace = serde_json::from_str::<serde_json::Value>(TRACE).unwrap()["result"].clone();
        let txs = trace["transactions"].clone();
        for tx in txs.as_array().unwrap() {
            let tx: TransactionTrace = serde_json::from_value(tx.clone()).unwrap();
            let _ = tx.try_build_typed_tx().unwrap();
        }
    }
}
