use alloy::{
    consensus::{EncodableSignature, SignableTransaction, Signed, Transaction, TxType},
    eips::{eip2930::AccessList, eip7702::SignedAuthorization},
    primitives::{keccak256, Bytes, ChainId, Signature, TxKind, B256, U256},
    rlp::{BufMut, Decodable, Encodable, Header},
};
use core::mem;
use serde_with::serde_as;

// #[cfg(not(feature = "std"))]
// use alloc::vec::Vec;

#[derive(Clone, Debug, Default, PartialEq, Eq, serde::Serialize, serde::Deserialize, Hash)]
#[serde_as]
pub struct TxAltFee {
    /// EIP-155: Simple replay attack protection
    pub chain_id: ChainId,
    /// A scalar value equal to the number of transactions sent by the sender; formally Tn.
    pub nonce: u64,
    /// A scalar value equal to the maximum
    /// amount of gas that should be used in executing
    /// this transaction. This is paid up-front, before any
    /// computation is done and may not be increased
    /// later; formally Tg.
    pub gas_limit: u128,
    /// A scalar value equal to the maximum
    /// amount of gas that should be used in executing
    /// this transaction. This is paid up-front, before any
    /// computation is done and may not be increased
    /// later; formally Tg.
    ///
    /// As ethereum circulation is around 120mil eth as of 2022 that is around
    /// 120000000000000000000000000 wei we are safe to use u128 as its max number is:
    /// 340282366920938463463374607431768211455
    ///
    /// This is also known as `GasFeeCap`
    pub max_fee_per_gas: u128,
    /// Max Priority fee that transaction is paying
    ///
    /// As ethereum circulation is around 120mil eth as of 2022 that is around
    /// 120000000000000000000000000 wei we are safe to use u128 as its max number is:
    /// 340282366920938463463374607431768211455
    ///
    /// This is also known as `GasTipCap`
    pub max_priority_fee_per_gas: u128,
    /// The 160-bit address of the message call’s recipient or, for a contract creation
    /// transaction, ∅, used here to denote the only member of B0 ; formally Tt.
    pub to: TxKind,
    /// A scalar value equal to the number of Wei to
    /// be transferred to the message call’s recipient or,
    /// in the case of contract creation, as an endowment
    /// to the newly created account; formally Tv.
    pub value: U256,
    /// The accessList specifies a list of addresses and storage keys;
    /// these addresses and storage keys are added into the `accessed_addresses`
    /// and `accessed_storage_keys` global sets (introduced in EIP-2929).
    /// A gas cost is charged, though at a discount relative to the cost of
    /// accessing outside the list.
    pub access_list: AccessList,
    /// Input has two uses depending if transaction is Create or Call (if `to` field is None or
    /// Some). pub init: An unlimited size byte array specifying the
    /// EVM-code for the account initialisation procedure CREATE,
    /// data: An unlimited size byte array specifying the
    /// input data of the message call, formally Td.
    pub input: Bytes,
    /// For AltFeeType
    pub fee_token_id: u16,
    /// For AltFeeType
    pub fee_limit: u64,
}

impl TxAltFee {
    /// Returns the effective gas price for the given `base_fee`.
    pub const fn effective_gas_price(&self, base_fee: Option<u64>) -> u128 {
        match base_fee {
            None => self.max_fee_per_gas,
            Some(base_fee) => {
                // if the tip is greater than the max priority fee per gas, set it to the max
                // priority fee per gas + base fee
                let tip = self.max_fee_per_gas.saturating_sub(base_fee as u128);
                if tip > self.max_priority_fee_per_gas {
                    self.max_priority_fee_per_gas + base_fee as u128
                } else {
                    // otherwise return the max fee per gas
                    self.max_fee_per_gas
                }
            }
        }
    }

    /// Decodes the inner [TxEip1559] fields from RLP bytes.
    ///
    /// NOTE: This assumes a RLP header has already been decoded, and _just_ decodes the following
    /// RLP fields in the following order:
    ///
    /// - `chain_id`
    /// - `nonce`
    /// - `max_priority_fee_per_gas`
    /// - `max_fee_per_gas`
    /// - `gas_limit`
    /// - `to`
    /// - `value`
    /// - `data` (`input`)
    /// - `access_list`
    /// - `fee_token_id`
    /// - `fee_limit`
    pub fn decode_fields(buf: &mut &[u8]) -> alloy::rlp::Result<Self> {
        Ok(Self {
            chain_id: Decodable::decode(buf)?,
            nonce: Decodable::decode(buf)?,
            max_priority_fee_per_gas: Decodable::decode(buf)?,
            max_fee_per_gas: Decodable::decode(buf)?,
            gas_limit: Decodable::decode(buf)?,
            to: Decodable::decode(buf)?,
            value: Decodable::decode(buf)?,
            input: Decodable::decode(buf)?,
            access_list: Decodable::decode(buf)?,
            fee_token_id: Decodable::decode(buf)?,
            fee_limit: Decodable::decode(buf)?,
        })
    }

    /// Outputs the length of the transaction's fields, without a RLP header.
    #[doc(hidden)]
    pub fn fields_len(&self) -> usize {
        let mut len = 0;
        len += self.chain_id.length();
        len += self.nonce.length();
        len += self.max_priority_fee_per_gas.length();
        len += self.max_fee_per_gas.length();
        len += self.gas_limit.length();
        len += self.to.length();
        len += self.value.length();
        len += self.input.0.length();
        len += self.access_list.length();
        len += self.fee_token_id.length();
        len += self.fee_limit.length();
        len
    }

    /// Encodes only the transaction's fields into the desired buffer, without a RLP header.
    pub(crate) fn encode_fields(&self, out: &mut dyn alloy::rlp::BufMut) {
        self.chain_id.encode(out);
        self.nonce.encode(out);
        self.max_priority_fee_per_gas.encode(out);
        self.max_fee_per_gas.encode(out);
        self.gas_limit.encode(out);
        self.to.encode(out);
        self.value.encode(out);
        self.input.0.encode(out);
        self.access_list.encode(out);
        self.fee_token_id.encode(out);
        self.fee_limit.encode(out);
    }

    /// Returns what the encoded length should be, if the transaction were RLP encoded with the
    /// given signature, depending on the value of `with_header`.
    ///
    /// If `with_header` is `true`, the payload length will include the RLP header length.
    /// If `with_header` is `false`, the payload length will not include the RLP header length.
    pub fn encoded_len_with_signature<S>(&self, signature: &S, with_header: bool) -> usize
    where
        S: EncodableSignature,
    {
        // this counts the tx fields and signature fields
        let payload_length = self.fields_len() + signature.rlp_vrs_len();

        // this counts:
        // * tx type byte
        // * inner header length
        // * inner payload length
        let inner_payload_length =
            1 + Header { list: true, payload_length }.length() + payload_length;

        if with_header {
            // header length plus length of the above, wrapped with a string header
            Header { list: false, payload_length: inner_payload_length }.length() +
                inner_payload_length
        } else {
            inner_payload_length
        }
    }

    /// Inner encoding function that is used for both rlp [`Encodable`] trait and for calculating
    /// hash that for eip2718 does not require a rlp header.
    #[doc(hidden)]
    pub fn encode_with_signature<S>(&self, signature: &S, out: &mut dyn BufMut, with_header: bool)
    where
        S: EncodableSignature,
    {
        let payload_length = self.fields_len() + signature.rlp_vrs_len();
        if with_header {
            Header {
                list: false,
                payload_length: 1 + Header { list: true, payload_length }.length() + payload_length,
            }
            .encode(out);
        }
        out.put_u8(self.tx_type() as u8);
        self.encode_with_signature_fields(signature, out);
    }

    /// Decodes the transaction from RLP bytes, including the signature.
    ///
    /// This __does not__ expect the bytes to start with a transaction type byte or string
    /// header.
    ///
    /// This __does__ expect the bytes to start with a list header and include a signature.
    #[doc(hidden)]
    pub fn decode_signed_fields(buf: &mut &[u8]) -> alloy::rlp::Result<Signed<Self>> {
        let header = Header::decode(buf)?;
        if !header.list {
            return Err(alloy::rlp::Error::UnexpectedString);
        }

        // record original length so we can check encoding
        let original_len = buf.len();

        let tx = Self::decode_fields(buf)?;
        let signature = Signature::decode_rlp_vrs(buf)?;

        let signed = tx.into_signed(signature);
        if buf.len() + header.payload_length != original_len {
            return Err(alloy::rlp::Error::ListLengthMismatch {
                expected: header.payload_length,
                got: original_len - buf.len(),
            });
        }

        Ok(signed)
    }

    /// Encodes the transaction from RLP bytes, including the signature. This __does not__ encode a
    /// tx type byte or string header.
    ///
    /// This __does__ encode a list header and include a signature.
    pub fn encode_with_signature_fields<S>(&self, signature: &S, out: &mut dyn BufMut)
    where
        S: EncodableSignature,
    {
        let payload_length = self.fields_len() + signature.rlp_vrs_len();
        let header = Header { list: true, payload_length };
        header.encode(out);
        self.encode_fields(out);
        signature.write_rlp_vrs(out);
    }

    /// Get transaction type
    #[doc(alias = "transaction_type")]
    pub(crate) const fn tx_type(&self) -> TxType {
        TxType::Eip1559
    }

    /// Calculates a heuristic for the in-memory size of the [TxEip1559] transaction.
    #[inline]
    pub fn size(&self) -> usize {
        mem::size_of::<ChainId>() + // chain_id
        mem::size_of::<u64>() + // nonce
        mem::size_of::<u64>() + // gas_limit
        mem::size_of::<u128>() + // max_fee_per_gas
        mem::size_of::<u128>() + // max_priority_fee_per_gas
        self.to.size() + // to
        mem::size_of::<U256>() + // value
        self.access_list.size() + // access_list
        self.input.len() + // input
        mem::size_of::<u16>()+ // fee_token_id
        mem::size_of::<u64>() // fee_limit
    }
}

impl Transaction for TxAltFee {
    fn chain_id(&self) -> Option<ChainId> {
        Some(self.chain_id)
    }

    fn nonce(&self) -> u64 {
        self.nonce
    }

    fn gas_limit(&self) -> u128 {
        self.gas_limit
    }

    fn gas_price(&self) -> Option<u128> {
        None
    }

    fn max_fee_per_gas(&self) -> u128 {
        self.max_fee_per_gas
    }

    fn max_priority_fee_per_gas(&self) -> Option<u128> {
        Some(self.max_priority_fee_per_gas)
    }

    fn priority_fee_or_price(&self) -> u128 {
        self.max_priority_fee_per_gas
    }

    fn max_fee_per_blob_gas(&self) -> Option<u128> {
        None
    }

    fn to(&self) -> TxKind {
        self.to
    }

    fn value(&self) -> U256 {
        self.value
    }

    fn input(&self) -> &[u8] {
        &self.input
    }

    fn ty(&self) -> u8 {
        0x7f
    }

    fn access_list(&self) -> Option<&AccessList> {
        Some(&self.access_list)
    }

    fn blob_versioned_hashes(&self) -> Option<&[B256]> {
        None
    }

    fn authorization_list(&self) -> Option<&[SignedAuthorization]> {
        None
    }
    // fn fee_token_id(&self) -> u16 {
    //     0x7f
    // }
    // fn fee_limit(&self) -> u64 {
    //     0x7f
    // }
}

impl SignableTransaction<Signature> for TxAltFee {
    fn set_chain_id(&mut self, chain_id: ChainId) {
        self.chain_id = chain_id;
    }

    fn encode_for_signing(&self, out: &mut dyn alloy::rlp::BufMut) {
        out.put_u8(self.tx_type() as u8);
        self.encode(out)
    }

    fn payload_len_for_signature(&self) -> usize {
        self.length() + 1
    }

    fn into_signed(self, signature: Signature) -> Signed<Self> {
        // Drop any v chain id value to ensure the signature format is correct at the time of
        // combination for an EIP-1559 transaction. V should indicate the y-parity of the
        // signature.
        let signature = signature.with_parity_bool();

        let mut buf = Vec::with_capacity(self.encoded_len_with_signature(&signature, false));
        self.encode_with_signature(&signature, &mut buf, false);
        let hash = keccak256(&buf);

        Signed::new_unchecked(self, signature, hash)
    }
}

impl Encodable for TxAltFee {
    fn encode(&self, out: &mut dyn BufMut) {
        Header { list: true, payload_length: self.fields_len() }.encode(out);
        self.encode_fields(out);
    }

    fn length(&self) -> usize {
        let payload_length = self.fields_len();
        Header { list: true, payload_length }.length() + payload_length
    }
}

impl Decodable for TxAltFee {
    fn decode(data: &mut &[u8]) -> alloy::rlp::Result<Self> {
        let header = Header::decode(data)?;
        let remaining_len = data.len();

        if header.payload_length > remaining_len {
            return Err(alloy::rlp::Error::InputTooShort);
        }

        Self::decode_fields(data)
    }
}
