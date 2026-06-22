use crate::ClientBlockInput;
use alloy_primitives::hex::FromHex;
use alloy_primitives::{Address, B256};
use alloy_provider::{DynProvider, EthGetBlock, Provider};
use alloy_rpc_types::{BlockNumberOrTag, Header as RpcHeader};
use anyhow::Context;
use prover_mpt::EthereumState;
use prover_primitives::types::block::L2Block;
use prover_primitives::types::{BlockHeader, TransactionTrace};
use prover_primitives::TxTrace;
use revm::state::Bytecode;
use serde::{Deserialize, Serialize};
use std::collections::HashMap;
use std::sync::LazyLock;

/// Mapping from chain ID to default coinbase address.
pub static CHAIN_COINBASE: LazyLock<HashMap<u64, Address>> = LazyLock::new(|| {
    HashMap::from([
        (2818u64, Address::from_hex("0x530000000000000000000000000000000000000a").unwrap()),
        (2910u64, Address::from_hex("0x29107CB79Ef8f69fE1587F77e283d47E84c5202f").unwrap()),
        (53077u64, Address::from_hex("0x716170d0687c3d31cc10debe0daa1ddd3fe3d792").unwrap()),
    ])
});

// Default coinbase address if chain ID is not found in CHAIN_COINBASE.
// 0xfabb0ac9d68b0b445fb7357272ff202c5651694a
pub static DEFAULT_COINBASE: LazyLock<Address> =
    LazyLock::new(|| Address::from_hex("0x716170d0687c3d31cc10debe0daa1ddd3fe3d792").unwrap());

/// Returns the beneficiary(coinbase) address for a given chain ID.
pub fn beneficiary_by_chain_id(chain_id: u64) -> Address {
    *CHAIN_COINBASE.get(&chain_id).unwrap_or(&DEFAULT_COINBASE)
}

/// Morph-compatible RPC block returned by `eth_getBlockByNumber`.
///
/// The default Alloy Ethereum network binds block transactions to Ethereum's
/// `TxEnvelope`, which cannot deserialize Morph's custom transaction types.
/// This alias keeps the normal RPC block shape, but swaps in Morph's header and
/// transaction envelope types from `morph-reth`.
pub type MorphRpcBlock = alloy_rpc_types::Block<
    morph_rpc::MorphRpcTransaction,
    RpcHeader<morph_primitives::MorphHeader>,
>;

/// Block structure returned by the prover RPC.
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
    /// disk root(mpt root)
    #[serde(default, rename = "diskRoot")]
    pub disk_root: B256,
    /// header root(zktrie root for headers)
    #[serde(default, rename = "headerRoot")]
    pub header_root: B256,
}

#[derive(Serialize, Deserialize, Debug, Clone)]
pub struct HostExecutorOutput {
    /// chain id
    pub chain_id: u64,
    /// beneficiary(coinbase)
    pub beneficiary: Address,
    /// executed block
    pub block: ProverBlock,
    /// State containing sparse MPT nodes.
    pub state: EthereumState,
    /// executed codes
    pub codes: Vec<Bytecode>,
    /// prev state root
    pub prev_state_root: B256,
    /// post state root
    pub post_state_root: B256,
}

/// Assembles a [ClientBlockInput] from the [HostExecutorOutput] and previous [ProverBlock].
pub fn assemble_block_input(
    output: HostExecutorOutput,
    prev_block: ProverBlock,
) -> ClientBlockInput {
    let block = output.block;
    let state = output.state;
    let codes = output.codes;

    let l2_block = L2Block {
        chain_id: output.chain_id,
        coinbase: output.beneficiary,
        header: block.header,
        transactions: block
            .transactions
            .iter()
            .map(|tx_trace| tx_trace.try_build_tx_envelope().unwrap())
            .collect(),
        prev_state_root: output.prev_state_root,
        post_state_root: output.post_state_root,
        start_l1_queue_index: prev_block.header.next_l1_msg_index.to::<u64>(),
    };
    ClientBlockInput { current_block: l2_block, parent_state: state, bytecodes: codes }
}

/// Queries the Morph RPC block at a given block number.
///
/// This intentionally uses [`MorphRpcBlock`] as the response type instead of the
/// provider network's default block type. For example,
/// `ProviderBuilder::new().connect_http(...).erased()` creates an Ethereum
/// provider whose `get_block(...).full()` decodes transactions as Ethereum
/// envelopes and fails on Morph-specific transaction types. `EthGetBlock` lets
/// us reuse the provider's RPC client while choosing a Morph-aware response
/// type.
pub async fn query_morph_rpc_block(
    block_number: u64,
    provider: &DynProvider,
) -> Result<MorphRpcBlock, anyhow::Error> {
    EthGetBlock::<MorphRpcBlock>::by_number(
        BlockNumberOrTag::Number(block_number),
        provider.client(),
    )
    .full()
    .await
    .with_context(|| format!("eth_getBlockByNumber failed for block {block_number}"))?
    .with_context(|| format!("block {block_number} not found"))
}

/// Queries the block at a given block number.
pub async fn query_block(
    block_number: u64,
    provider: &DynProvider,
) -> Result<ProverBlock, anyhow::Error> {
    let block = query_morph_rpc_block(block_number, provider).await?;
    let value = serde_json::to_value(block)
        .with_context(|| format!("failed to serialize Morph RPC block {block_number}"))?;
    serde_json::from_value(value).with_context(|| {
        format!("failed to convert Morph RPC block {block_number} into ProverBlock")
    })
}

#[cfg(test)]
mod tests {
    use super::query_morph_rpc_block;
    use alloy_provider::{Provider, ProviderBuilder};

    #[tokio::test]
    #[ignore = "requires public Morph RPC access"]
    async fn query_morph_block_test() {
        let rpc = "https://rpc-quicknode.morphl2.io";

        // Provider construction is independent from the response type used for
        // `eth_getBlockByNumber`: `query_morph_rpc_block` reuses the provider
        // client but decodes into `MorphRpcBlock`.
        // let root_provider: RootProvider<_> = RootProvider::new_http(rpc.parse().unwrap());
        // let block = query_morph_rpc_block(19720290, &root_provider).await.unwrap();
        // println!("block via RootProvider: {:?}", serde_json::to_string(&block));

        // The same Morph-aware request also works with a ProviderBuilder-created
        // provider; avoid `.get_block(...).full()` here because that method is
        // bound to the provider network's default Ethereum block response type.
        let builder_provider =
            ProviderBuilder::default().connect_http(rpc.parse().unwrap()).erased();
        let block = query_morph_rpc_block(19720290, &builder_provider).await.unwrap();
        println!("block via ProviderBuilder: {:?}", &block);
    }
}
