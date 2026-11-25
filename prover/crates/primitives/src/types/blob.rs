use alloy_consensus::BlockHeader;
use alloy_eips::eip2718::Encodable2718;
use alloy_primitives::U256;
use morph_primitives::Block as MorphBlock;

/// A block's total transaction count cannot fit in its blob context.
#[derive(Debug, PartialEq, Eq, thiserror::Error)]
#[error("block {block_number} has {count} transactions, exceeding the blob limit of {}", u16::MAX)]
pub struct BlockTransactionCountOverflow {
    /// The block whose transaction count overflowed.
    pub block_number: u64,
    /// The actual number of transactions in the block.
    pub count: usize,
}

/// Encodes the block contexts followed by the concatenated L2 transactions.
pub fn get_blob_data_from_blocks<'a>(
    blocks: impl IntoIterator<Item = &'a MorphBlock>,
) -> Result<Vec<u8>, BlockTransactionCountOverflow> {
    let blocks = blocks.into_iter();
    let mut batch_data = Vec::with_capacity(blocks.size_hint().0 * 60);
    let mut tx_bytes = Vec::new();
    for block in blocks {
        let block_number = block.header.number();
        let transactions = &block.body.transactions;
        let count = transactions.len();
        let tx_count = u16::try_from(count)
            .map_err(|_| BlockTransactionCountOverflow { block_number, count })?;

        // L1 messages are a subset of the checked transaction count.
        let mut l1_count = 0u16;
        for tx in transactions {
            if tx.is_l1_msg() {
                l1_count += 1;
            } else {
                tx.encode_2718(&mut tx_bytes);
            }
        }

        batch_data.extend_from_slice(&block_number.to_be_bytes());
        batch_data.extend_from_slice(&block.header.timestamp().to_be_bytes());
        batch_data.extend_from_slice(
            &U256::from(block.header.base_fee_per_gas().unwrap_or_default()).to_be_bytes::<32>(),
        );
        batch_data.extend_from_slice(&block.header.gas_limit().to_be_bytes());
        batch_data.extend_from_slice(&tx_count.to_be_bytes());
        batch_data.extend_from_slice(&l1_count.to_be_bytes());
    }
    batch_data.extend_from_slice(&tx_bytes);
    Ok(batch_data)
}

#[cfg(test)]
mod tests {
    use alloy_consensus::{Sealed, Signed, TxLegacy};
    use alloy_primitives::{B256, Signature, TxKind, hex};
    use morph_primitives::{MorphTxEnvelope, TxL1Msg};

    use super::*;

    #[test]
    fn encodes_contexts_before_ordered_l2_transactions() {
        let l2 = |nonce| {
            MorphTxEnvelope::Legacy(Signed::new_unchecked(
                TxLegacy { nonce, to: TxKind::Create, ..Default::default() },
                Signature::new(U256::from(1), U256::from(1), false),
                B256::ZERO,
            ))
        };
        let mut first = MorphBlock::default();
        first.header.inner.number = 1;
        first.header.inner.timestamp = 2;
        first.header.inner.base_fee_per_gas = Some(3);
        first.header.inner.gas_limit = 4;
        first.body.transactions =
            vec![MorphTxEnvelope::L1Msg(Sealed::new(TxL1Msg::default())), l2(1)];
        let mut second = MorphBlock::default();
        second.header.inner.number = 2;
        second.body.transactions.push(l2(2));

        let expected: &[&[u8]] = &[
            &hex!("00000000000000010000000000000002"),
            &hex!("0000000000000000000000000000000000000000000000000000000000000003"),
            &hex!("000000000000000400020001"),
            &hex!("0000000000000002"),
            &[0; 48],
            &hex!("00010000"),
            &hex!("c90180808080801b0101c90280808080801b0101"),
        ];
        assert_eq!(get_blob_data_from_blocks([&first, &second]).unwrap(), expected.concat());
    }

    #[test]
    fn checks_block_transaction_count_limit() {
        let mut block = MorphBlock::default();
        block.header.inner.number = 42;
        let l1 = MorphTxEnvelope::L1Msg(Sealed::new(TxL1Msg::default()));
        block.body.transactions = vec![l1.clone(); usize::from(u16::MAX)];

        let encoded = get_blob_data_from_blocks([&block]).unwrap();
        assert_eq!(&encoded[56..], &[0xff; 4]);

        block.body.transactions.push(l1);
        assert_eq!(
            get_blob_data_from_blocks([&block]),
            Err(BlockTransactionCountOverflow { block_number: 42, count: 65_536 })
        );
    }
}
