use alloy_consensus::BlockHeader;
use alloy_primitives::B256;
use morph_primitives::Block;
use prover_executor_core::{MorphExecutor, validate_next_l1_msg_index};
use reth_trie::{HashedPostState, KeccakKeyHasher};

use crate::types::{batch::BatchInfo, error::ClientError, input::BlockInput};

// use Verifier;
pub struct EVMVerifier;

const MAX_L2_BASE_FEE: u64 = 10_000_000_000;

impl EVMVerifier {
    pub fn verify(blocks: Vec<BlockInput>) -> Result<BatchInfo, ClientError> {
        // Edge case: nothing to execute.
        if blocks.is_empty() {
            return Err(ClientError::BlockExecutionError(
                "empty batch: no block inputs provided".to_owned(),
            ));
        }
        // Verify block numbers and state roots are consecutive within the batch.
        for window in blocks.windows(2) {
            // ChainId check
            if window[0].chain_id != window[1].chain_id {
                return Err(ClientError::ChainIdNotMatch);
            }

            // Blocknum check
            let previous = window[0].current_block.number();
            let current = window[1].current_block.number();
            let expected = previous
                .checked_add(1)
                .ok_or(ClientError::InvalidHeaderBlockNumber(previous, current))?;
            if current != expected {
                return Err(ClientError::InvalidHeaderBlockNumber(expected, current));
            }

            // Block timestamp check
            if window[0].current_block.timestamp() > window[1].current_block.timestamp() {
                return Err(ClientError::InvalidBlockTimestamp(current));
            }

            // StateRoot continuous check
            if window[0].current_block.state_root() != window[1].parent_state.state_root() {
                return Err(ClientError::DiscontinuousStateRoot);
            }
        }

        for input in &blocks {
            validate_header_constants(&input.current_block)?;
            validate_base_fee(&input.current_block)?;
        }

        // The DA commitments bind only the L1-only projection (`data_hash`) and the
        // L2-only projection (blob) of each block's transactions, never their
        // interleaving. Enforce the canonical order before execution so that one
        // committed batch admits exactly one transaction stream.
        validate_canonical_order(blocks.iter().map(|input| &input.current_block))?;
        for (index, input) in blocks.iter().enumerate() {
            let parent_header =
                index.checked_sub(1).map(|parent_index| &blocks[parent_index].current_block.header);
            validate_next_l1_msg_index(
                input.chain_id,
                parent_header,
                &input.current_block.header,
                &input.current_block.body.transactions,
            )
            .map_err(|reason| ClientError::InvalidNextL1MessageIndex {
                block_num: input.current_block.number(),
                reason: reason.to_string(),
            })?;
        }

        execute(blocks)
    }
}

/// Validates header fields that feed the EVM block environment but are not committed by the
/// blob. `mix_hash` becomes PREVRANDAO and `difficulty` becomes DIFFICULTY (pre-Merge spec);
/// morph-geth always returns 0 for opcode 0x44 and builds L2 headers with both set to zero, so
/// any other value would yield a state transition the L2 network never executed.
fn validate_header_constants(block: &Block) -> Result<(), ClientError> {
    let block_num = block.header.number();
    let mix_hash = block.header.mix_hash().unwrap_or_default();
    if mix_hash != B256::ZERO {
        return Err(ClientError::InvalidHeaderMixHash { block_num, mix_hash });
    }
    let difficulty = block.header.difficulty();
    if !difficulty.is_zero() {
        return Err(ClientError::InvalidHeaderDifficulty { block_num, difficulty });
    }
    Ok(())
}

fn validate_base_fee(block: &Block) -> Result<(), ClientError> {
    let base_fee = block
        .header
        .base_fee_per_gas()
        .ok_or(ClientError::MissingHeaderBaseFee(block.header.number()))?;
    if base_fee > MAX_L2_BASE_FEE {
        return Err(ClientError::InvalidHeaderBaseFee(MAX_L2_BASE_FEE, base_fee));
    }
    Ok(())
}

/// Validates the L1 message ordering of a batch, mirroring the node's
/// `validateL1Messages` (node/core/l1_message.go):
///
/// - within a block, every L1 message precedes every L2 transaction (constraint 5);
/// - L1 message queue indices are strictly consecutive (constraint 3), here checked across the
///   whole batch, which `Rollup._commitBatch` already requires through `data_hash`.
///
/// Together with the per-block `num_l1_txs` in the blob's block context, this makes the
/// transaction stream of every block a function of the committed batch data.
pub(crate) fn validate_canonical_order<'a>(
    blocks: impl IntoIterator<Item = &'a Block>,
) -> Result<(), ClientError> {
    let mut expected_queue_index: Option<u64> = None;
    for block in blocks {
        let block_num = block.header.number();
        let mut l1_section_over = false;
        for (tx_index, tx) in block.body.transactions.iter().enumerate() {
            if !tx.is_l1_msg() {
                l1_section_over = true;
                continue;
            }
            if l1_section_over {
                return Err(ClientError::L1MessageAfterL2Tx { block_num, tx_index });
            }
            let queue_index =
                tx.queue_index().ok_or(ClientError::MalformedL1Message { block_num, tx_index })?;
            if let Some(expected) = expected_queue_index
                && queue_index != expected
            {
                return Err(ClientError::L1MessageNotConsecutive {
                    block_num,
                    expected,
                    actual: queue_index,
                });
            }
            expected_queue_index = Some(
                queue_index
                    .checked_add(1)
                    .ok_or(ClientError::MalformedL1Message { block_num, tx_index })?,
            );
        }
    }
    Ok(())
}

fn execute(mut block_inputs: Vec<BlockInput>) -> Result<BatchInfo, ClientError> {
    // fetch & calculate initial state root
    let prev_state_root = block_inputs.first().unwrap().parent_state.state_root();
    // Execute each block sequentially.
    block_inputs.iter_mut().try_for_each(execute_block)?;

    BatchInfo::from_block_inputs(prev_state_root, &block_inputs)
}

fn execute_block(block_input: &mut BlockInput) -> Result<(), ClientError> {
    let block = &block_input.current_block;

    if block.body.transactions.is_empty() {
        if block.state_root() != block_input.parent_state.state_root() {
            // For empty blocks, EVM execution is skipped, but the post root is constrained to equal
            // the previous root.
            return Err(ClientError::MismatchedStateRoot {
                block_num: block.header.number(),
                root_trace: block.state_root(),
                root_revm: block_input.parent_state.state_root(),
            });
        }
        return Ok(());
    }
    let header = &block.header;
    let chain_id = block_input.chain_id;
    let _block_num = block.number();
    let _txn_count = block.body.transactions.len();

    // Build DB, this will internally verify the correctness of mpt.
    let trie_db = block_input.witness_db()?;

    // Execute the whole block via reth's BasicBlockExecutor.
    let core_executor = MorphExecutor::new_ref(trie_db, chain_id);
    let bundle_state = core_executor
        .execute_block(block.clone())
        .map_err(|e| ClientError::BlockExecutionError(format!("{e:#}")))?;
    // Verify the post-state root by applying the block's transition set to the parent (pre-block)
    // state.
    let computed_state_root = {
        let hashed_post_state =
            HashedPostState::from_bundle_state::<KeccakKeyHasher>(&bundle_state.state);
        block_input.parent_state.update(&hashed_post_state);
        block_input.parent_state.state_root()
    };
    if computed_state_root != block.state_root() {
        return Err(ClientError::MismatchedStateRoot {
            block_num: header.number(),
            root_trace: block.state_root(),
            root_revm: computed_state_root,
        });
    }
    #[cfg(not(target_os = "zkvm"))]
    log::debug!("success execute block_{_block_num} in client, txns.len: {_txn_count}");

    Ok(())
}

#[cfg(test)]
mod tests {
    use alloy_consensus::{BlockBody, Sealed, Signed, TxLegacy};
    use alloy_primitives::{Signature, U256};
    use morph_primitives::{MorphHeader, MorphTxEnvelope, TxL1Msg};

    use super::*;

    fn l1(queue_index: u64) -> MorphTxEnvelope {
        MorphTxEnvelope::L1Msg(Sealed::new(TxL1Msg { queue_index, ..Default::default() }))
    }

    fn l2(nonce: u64) -> MorphTxEnvelope {
        MorphTxEnvelope::Legacy(Signed::new_unchecked(
            TxLegacy { nonce, ..Default::default() },
            Signature::test_signature(),
            B256::with_last_byte(nonce as u8),
        ))
    }

    fn block(number: u64, transactions: Vec<MorphTxEnvelope>) -> Block {
        let mut header = MorphHeader::default();
        header.inner.number = number;
        Block { header, body: BlockBody { transactions, ..Default::default() } }
    }

    #[test]
    fn accepts_canonical_order() {
        let blocks = [
            block(1, vec![l1(0), l1(1), l2(0), l2(1)]),
            block(2, vec![l2(2)]),
            block(3, vec![]),
            block(4, vec![l1(2)]),
            block(5, vec![l1(3), l2(3)]),
        ];
        validate_canonical_order(&blocks).unwrap();
    }

    #[test]
    fn rejects_l1_message_after_l2_tx() {
        for txs in [vec![l2(0), l2(1), l1(0), l1(1)], vec![l1(0), l2(0), l1(1)], vec![l2(0), l1(0)]]
        {
            let err = validate_canonical_order(&[block(7, txs)]).unwrap_err();
            assert!(matches!(err, ClientError::L1MessageAfterL2Tx { block_num: 7, .. }), "{err}");
        }
    }

    #[test]
    fn rejects_non_consecutive_queue_index() {
        let err = validate_canonical_order(&[block(1, vec![l1(0), l1(2)])]).unwrap_err();
        assert!(matches!(
            err,
            ClientError::L1MessageNotConsecutive { block_num: 1, expected: 1, actual: 2 }
        ));

        let err =
            validate_canonical_order(&[block(1, vec![l1(5)]), block(2, vec![l1(5)])]).unwrap_err();
        assert!(matches!(
            err,
            ClientError::L1MessageNotConsecutive { block_num: 2, expected: 6, actual: 5 }
        ));
    }

    #[test]
    fn validates_base_fee_range() {
        let mut b = block(9, vec![]);
        b.header.inner.base_fee_per_gas = Some(0);
        validate_base_fee(&b).unwrap();

        b.header.inner.base_fee_per_gas = Some(MAX_L2_BASE_FEE);
        validate_base_fee(&b).unwrap();

        b.header.inner.base_fee_per_gas = None;
        assert!(matches!(validate_base_fee(&b), Err(ClientError::MissingHeaderBaseFee(9))));
    }

    #[test]
    fn rejects_excessive_base_fee_for_an_empty_block() {
        let mut input = BlockInput {
            current_block: block(9, vec![]),
            parent_state: Default::default(),
            bytecodes: vec![],
            chain_id: 0,
        };
        input.current_block.header.inner.state_root = input.parent_state.state_root();
        input.current_block.header.inner.base_fee_per_gas = Some(MAX_L2_BASE_FEE + 1);

        let err = EVMVerifier::verify(vec![input]).unwrap_err();
        assert!(matches!(
            err,
            ClientError::InvalidHeaderBaseFee(MAX_L2_BASE_FEE, actual)
                if actual == MAX_L2_BASE_FEE + 1
        ));
    }

    #[test]
    fn rejects_non_zero_mix_hash_and_difficulty() {
        let mut b = block(9, vec![]);
        validate_header_constants(&b).unwrap();

        b.header.inner.mix_hash = B256::with_last_byte(1);
        let err = validate_header_constants(&b).unwrap_err();
        assert!(matches!(err, ClientError::InvalidHeaderMixHash { block_num: 9, .. }));

        b.header.inner.mix_hash = B256::ZERO;
        b.header.inner.difficulty = U256::from(1);
        let err = validate_header_constants(&b).unwrap_err();
        assert!(matches!(err, ClientError::InvalidHeaderDifficulty { block_num: 9, .. }));
    }

    /// Real mainnet/devnet block inputs must still pass full verification.
    #[test]
    fn testdata_block_inputs_verify() {
        let dir = concat!(env!("CARGO_MANIFEST_DIR"), "/../../../testdata/block_inputs");
        let mut paths: Vec<_> = std::fs::read_dir(dir)
            .unwrap()
            .map(|entry| entry.unwrap().path())
            .filter(|path| path.extension().is_some_and(|ext| ext == "json"))
            .collect();
        paths.sort();
        assert!(!paths.is_empty());

        let mut num_l1_msgs = 0;
        for path in paths {
            let inputs: Vec<BlockInput> =
                serde_json::from_reader(std::fs::File::open(&path).unwrap()).unwrap();
            num_l1_msgs += inputs
                .iter()
                .flat_map(|input| &input.current_block.body.transactions)
                .filter(|tx| tx.is_l1_msg())
                .count();
            EVMVerifier::verify(inputs).unwrap_or_else(|e| panic!("{}: {e}", path.display()));
        }
        assert!(num_l1_msgs > 0, "testdata should cover L1 messages");
    }
}
