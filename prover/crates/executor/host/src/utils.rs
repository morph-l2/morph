use alloy_primitives::hex::FromHex;
use alloy_primitives::{Address, Bytes, B256, U256, U64, U8};
use alloy_provider::{DynProvider, Provider};
use alloy_rpc_types::AccessList;
use anyhow::Context;
use prover_mpt::EthereumState;
use prover_primitives::types::{AuthorizationList, BlockHeader, TransactionTrace};
use revm::state::Bytecode;
use serde::{Deserialize, Serialize};
use serde_with::{serde_as, DefaultOnNull};

use std::collections::HashMap;
use std::sync::LazyLock;

pub static CHAIN_CONFIG: LazyLock<HashMap<u64, Address>> = LazyLock::new(|| {
    HashMap::from([
        (2818u64, Address::from_hex("0x530000000000000000000000000000000000000a").unwrap()),
        (53077u64, Address::from_hex("0xfabb0ac9d68b0b445fb7357272ff202c5651694a").unwrap()),
    ])
});

#[serde_as]
#[derive(serde::Serialize, serde::Deserialize, Default, Debug, Clone, Hash, PartialEq, Eq)]
pub struct ProverTx {
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
    /// value amount
    pub(crate) value: U256,
    /// call data
    pub(crate) input: Bytes,
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

#[derive(Serialize, Deserialize, Default, Debug, Clone)]
pub struct ProverBlock {
    /// block
    #[serde(flatten)]
    pub header: BlockHeader,
    /// txs
    pub transactions: Vec<TransactionTrace>,
}

#[derive(Serialize, Deserialize, Default, Debug, Clone)]
pub struct DiskRoot {
    #[serde(default, rename = "diskRoot")]
    pub disk_root: B256,
    #[serde(default, rename = "headerRoot")]
    pub header_root: B256,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HostExecutorOutput {
    pub chain_id: u64,
    pub beneficiary: Address,
    pub block: ProverBlock,
    pub state: EthereumState,
    pub codes: Vec<Bytecode>,
    /// prev state root
    pub prev_state_root: B256,
    /// post state root
    pub post_state_root: B256,
}

pub async fn query_state_root(
    block_number: u64,
    provider: &DynProvider,
) -> Result<DiskRoot, anyhow::Error> {
    provider
        .raw_request::<_, DiskRoot>("morph_diskRoot".into(), [format!("{block_number:#x}")])
        .await
        .context("morph_diskRoot error")
}

pub async fn query_block(
    block_number: u64,
    provider: &DynProvider,
) -> Result<ProverBlock, anyhow::Error> {
    provider
        .raw_request::<_, ProverBlock>(
            "eth_getBlockByNumber".into(),
            (format!("{block_number:#x}"), true),
        )
        .await
        .context("eth_getBlockByNumber error")
}

pub async fn query_chain_d(
    block_number: u64,
    provider: &DynProvider,
) -> Result<ProverBlock, anyhow::Error> {
    provider
        .raw_request::<_, ProverBlock>(
            "eth_getBlockByNumber".into(),
            [format!("{block_number:#x}")],
        )
        .await
        .context("eth_getBlockByNumber error")
}
