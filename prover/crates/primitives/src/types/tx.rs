use crate::{types::AuthorizationList, TxTrace};

use alloy_eips::{eip2930::AccessList, eip7702::SignedAuthorization};
use alloy_primitives::{
    normalize_v, Address, Bytes, ChainId, Signature, SignatureError, TxKind, B256, U16, U256, U64,
    U8,
};
use serde_with::{serde_as, DefaultOnNull};

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
    pub(crate) fee_token_id: Option<U16>,
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
        self.fee_token_id.unwrap_or_default().to::<u16>()
    }

    fn fee_limit(&self) -> U256 {
        self.fee_limit.unwrap_or_default().to()
    }

    fn sig_v(&self) -> u64 {
        self.v.to::<u64>()
    }
}
