use ethers::{
    types::{Eip1559TransactionRequest, Eip2930TransactionRequest, TransactionRequest, U64},
    utils::rlp::{self, Decodable},
};
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Clone, PartialEq, Eq, Debug)]
#[cfg_attr(not(feature = "legacy"), serde(tag = "type"))]
#[cfg_attr(feature = "legacy", serde(untagged))]
pub enum TypedTransaction {
    // 0x00
    #[serde(rename = "0x00", alias = "0x0")]
    Legacy(TransactionRequest),
    // 0x01
    #[serde(rename = "0x01", alias = "0x1")]
    Eip2930(Eip2930TransactionRequest),
    // 0x02
    #[serde(rename = "0x02", alias = "0x2")]
    Eip1559(Eip1559TransactionRequest),
    // 0x7E for Morphism
    // Just calculate the number of transactions, use Legacy instead
    #[serde(rename = "0x7E")]
    L1MessageTx(TransactionRequest),
}

/// Get a TypedTransaction directly from a rlp encoded byte stream
impl Decodable for TypedTransaction {
    fn decode(rlp: &rlp::Rlp) -> Result<Self, rlp::DecoderError> {
        let mut tx_type: Option<U64> = None;
        let empty = vec![];
        let mut rest = rlp::Rlp::new(&empty);
        if !rlp.is_list() {
            // if it is not enveloped then we need to use rlp.as_raw instead of rlp.data
            let first_byte = rlp.as_raw()[0];
            let (first, data) = if first_byte <= 0x7f {
                (first_byte, rlp.as_raw())
            } else {
                let data = rlp.data()?;
                let first = *data
                    .first()
                    .ok_or(rlp::DecoderError::Custom("empty slice"))?;
                (first, data)
            };

            let bytes = data
                .get(1..)
                .ok_or(rlp::DecoderError::Custom("no tx body"))?;
            rest = rlp::Rlp::new(bytes);

            tx_type = Some(U64::from_big_endian(&vec![first]));
        }

        match tx_type {
            Some(x) if x == U64::from(1) => {
                // EIP-2930 (0x01)
                Ok(Self::Eip2930(Eip2930TransactionRequest::decode(&rest)?))
            }
            Some(x) if x == U64::from(2) => {
                // EIP-1559 (0x02)
                Ok(Self::Eip1559(Eip1559TransactionRequest::decode(&rest)?))
            }
            Some(x) if x == U64::from(0x7E) => {
                // L1 tx_type, Just calculate the number of transactions, use Legacy instead
                Ok(Self::L1MessageTx(TransactionRequest::default()))
            }
            _ => {
                // Legacy (0x00)
                // use the original rlp
                Ok(Self::Legacy(TransactionRequest::decode(rlp)?))
            }
        }
    }
}
