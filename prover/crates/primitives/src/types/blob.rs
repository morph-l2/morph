use crate::{
    types::{block::L2Block, BlockTrace},
    Block, TxTrace,
};

/// Get blob data from L2 blocks
pub fn get_blob_data_from_blocks(block_inputs: &Vec<L2Block>) -> Vec<u8> {
    let num_blocks = block_inputs.len();
    let mut batch_data: Vec<u8> = Vec::with_capacity(num_blocks * 60);
    let mut tx_bytes: Vec<u8> = vec![];
    for block in block_inputs {
        // BlockContext
        // https://github.com/morph-l2/morph/blob/main/contracts/contracts/libraries/codec/BatchCodecV0.sol
        let mut block_ctx: Vec<u8> = Vec::with_capacity(60);
        block_ctx.extend_from_slice(&block.header.number.to::<u64>().to_be_bytes());
        block_ctx.extend_from_slice(&block.header.timestamp.to::<u64>().to_be_bytes());
        block_ctx.extend_from_slice(
            &block.header.base_fee_per_gas.unwrap_or_default().to_be_bytes::<32>(),
        );
        block_ctx.extend_from_slice(&block.header.gas_limit.to::<u64>().to_be_bytes());
        block_ctx.extend_from_slice(&(block.transactions.len() as u16).to_be_bytes());
        block_ctx.extend_from_slice(&(block.num_l1_txs() as u16).to_be_bytes());
        batch_data.extend(block_ctx);

        // Collect txns
        let x = block
            .transactions
            .iter()
            .filter(|tx| !tx.is_l1_msg())
            .flat_map(|tx| tx.rlp())
            .collect::<Vec<u8>>();
        tx_bytes.extend(x);
    }
    batch_data.extend(tx_bytes);
    batch_data
}

/// Get blob data from BlockTraces
pub fn get_blob_data_from_traces(block_trace: &Vec<BlockTrace>) -> Vec<u8> {
    let num_blocks = block_trace.len();
    let mut batch_data: Vec<u8> = Vec::with_capacity(num_blocks * 60);
    let mut tx_bytes: Vec<u8> = vec![];
    for trace in block_trace {
        // BlockContext
        // https://github.com/morph-l2/morph/blob/main/contracts/contracts/libraries/codec/BatchCodecV0.sol
        let mut block_ctx: Vec<u8> = Vec::with_capacity(60);
        block_ctx.extend_from_slice(&trace.number().to_be_bytes());
        block_ctx.extend_from_slice(&trace.timestamp().to::<u64>().to_be_bytes());
        block_ctx
            .extend_from_slice(&trace.base_fee_per_gas().unwrap_or_default().to_be_bytes::<32>());
        block_ctx.extend_from_slice(&trace.gas_limit().to::<u64>().to_be_bytes());
        block_ctx.extend_from_slice(&(trace.transactions.len() as u16).to_be_bytes());
        block_ctx.extend_from_slice(&(trace.num_l1_txs() as u16).to_be_bytes());
        batch_data.extend(block_ctx);

        // Collect txns
        let x = trace
            .transactions
            .iter()
            .filter(|tx| !tx.is_l1_tx())
            .flat_map(|tx| tx.try_build_tx_envelope().unwrap().rlp())
            .collect::<Vec<u8>>();
        tx_bytes.extend(x);
    }
    batch_data.extend(tx_bytes);
    batch_data
}
