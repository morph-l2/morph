use crate::ClientBlockInput;
use alloy_primitives::hex::FromHex;
use alloy_primitives::{Address, B256};
use alloy_provider::{DynProvider, Provider};
use anyhow::Context;
use prover_mpt::EthereumState;
use prover_primitives::types::block::L2Block;
use prover_primitives::types::{BlockHeader, TransactionTrace};
use prover_primitives::TxTrace;
use revm::state::Bytecode;
use serde::{Deserialize, Serialize};
use std::collections::HashMap;
use std::sync::LazyLock;

pub static CHAIN_CONFIG: LazyLock<HashMap<u64, Address>> = LazyLock::new(|| {
    HashMap::from([
        (2818u64, Address::from_hex("0x530000000000000000000000000000000000000a").unwrap()),
        (53077u64, Address::from_hex("0xfabb0ac9d68b0b445fb7357272ff202c5651694a").unwrap()),
    ])
});

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

    let block_input =
        ClientBlockInput { current_block: l2_block, parent_state: state, bytecodes: codes };

    block_input
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
