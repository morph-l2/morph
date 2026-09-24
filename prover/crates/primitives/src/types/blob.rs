use alloy_consensus::BlockHeader;
use alloy_primitives::U256;
use morph_primitives::Block as MorphBlock;

/// Get blob data from L2 blocks
pub fn get_blob_data_from_blocks(block_inputs: &Vec<MorphBlock>) -> Vec<u8> {
    let num_blocks = block_inputs.len();
    let mut batch_data: Vec<u8> = Vec::with_capacity(num_blocks * 60);
    let mut tx_bytes: Vec<u8> = vec![];
    for block in block_inputs {
        // BlockContext
        // https://github.com/morph-l2/morph/blob/main/contracts/contracts/libraries/codec/BatchHeaderCodecV1.sol
        let mut block_ctx: Vec<u8> = Vec::with_capacity(60);
        block_ctx.extend_from_slice(&block.header.number().to_be_bytes());
        block_ctx.extend_from_slice(&block.header.timestamp().to_be_bytes());
        block_ctx.extend_from_slice(
            &U256::from(block.header.base_fee_per_gas().unwrap_or_default()).to_be_bytes::<32>(),
        );
        block_ctx.extend_from_slice(&block.header.gas_limit().to_be_bytes());
        let transactions = &block.body.transactions;
        block_ctx.extend_from_slice(&(transactions.len() as u16).to_be_bytes());
        let num_l1_txs = transactions.iter().filter(|tx| tx.is_l1_msg()).collect::<Vec<_>>().len();
        block_ctx.extend_from_slice(&(num_l1_txs as u16).to_be_bytes());
        batch_data.extend(block_ctx);

        // Collect txns
        let x = transactions
            .iter()
            .filter(|tx| !tx.is_l1_msg())
            .flat_map(|tx| tx.rlp())
            .collect::<Vec<u8>>();
        tx_bytes.extend(x);
    }
    batch_data.extend(tx_bytes);
    batch_data
}
