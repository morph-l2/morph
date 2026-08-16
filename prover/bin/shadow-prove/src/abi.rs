use alloy_primitives::Bytes;
use alloy_sol_types::{sol, SolCall};
use anyhow::{anyhow, Result};

// Codegen from ABI file to interact with the contract.
sol!(
    #[sol(rpc)]
    Rollup,
    "abi/Rollup.json"
);

// Immutable Rollup ABI from immediately before BatchSignatureInput was removed.
pub mod pre_submitter {
    use alloy_sol_types::sol;

    sol!(PreSubmitterRollup, "abi/pre_submitter/Rollup.json");
}
use pre_submitter::PreSubmitterRollup;

#[derive(Clone, Debug, Eq, PartialEq)]
pub struct CommitBatchInput {
    pub parent_batch_header: Bytes,
    pub last_block_number: u64,
}

/// Derive the target batch's L2 block range from the previous and target commits.
pub fn target_batch_block_range(
    previous: &CommitBatchInput,
    target: &CommitBatchInput,
) -> (u64, u64) {
    (previous.last_block_number + 1, target.last_block_number)
}

/// A commit carries its parent's header, so the commit after the target is the source of the
/// target batch header.
pub fn target_batch_header_from_next_commit(next: &CommitBatchInput) -> Bytes {
    next.parent_batch_header.clone()
}

/// Decode the pre- or post-upgrade commitBatch calldata handled by shadow-prove.
pub fn decode_commit_batch(data: &[u8]) -> Result<CommitBatchInput> {
    let selector =
        data.get(..4).ok_or_else(|| anyhow!("rollup calldata is shorter than a selector"))?;

    if selector == Rollup::commitBatchCall::SELECTOR {
        let call = Rollup::commitBatchCall::abi_decode(data)?;
        return Ok(CommitBatchInput {
            parent_batch_header: call.batchDataInput.parentBatchHeader,
            last_block_number: call.batchDataInput.lastBlockNumber,
        });
    }
    if selector == PreSubmitterRollup::commitBatchCall::SELECTOR {
        let call = PreSubmitterRollup::commitBatchCall::abi_decode(data)?;
        return Ok(CommitBatchInput {
            parent_batch_header: call.batchDataInput.parentBatchHeader,
            last_block_number: call.batchDataInput.lastBlockNumber,
        });
    }

    Err(anyhow!("unsupported rollup calldata selector"))
}

sol!(
    #[sol(rpc)]
    ShadowRollup,
    "abi/ShadowRollup.json"
);

sol!(
    #[sol(rpc)]
    SP1Verifier,
    "abi/SP1Verifier.json"
);

#[cfg(test)]
mod tests {
    use super::*;
    use alloy_primitives::{B256, U256};

    const PRE_SUBMITTER_COMMIT_BATCH: &str = "0x428868b500000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000030390000000000000000000000000000000000000000000000000000000000000007111111111111111111111111111111111111111111111111111111111111111122222222222222222222222222222222222222222222222222222222222222223333333333333333333333333333333333333333333333333333333333333333000000000000000000000000000000000000000000000000000000000000000301020300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000009000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000002aabb0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002ccdd000000000000000000000000000000000000000000000000000000000000";
    const POST_SUBMITTER_COMMIT_BATCH: &str = "0x41f756da0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000003039000000000000000000000000000000000000000000000000000000000000000711111111111111111111111111111111111111111111111111111111111111112222222222222222222222222222222222222222222222222222222222222222333333333333333333333333333333333333333333333333333333333333333300000000000000000000000000000000000000000000000000000000000000030102030000000000000000000000000000000000000000000000000000000000";

    fn calldata(value: &str) -> Vec<u8> {
        alloy_primitives::hex::decode(value.trim_start_matches("0x")).unwrap()
    }

    fn post_submitter_commit_batch(parent_batch_header: &[u8], last_block_number: u64) -> Vec<u8> {
        Rollup::commitBatchCall {
            batchDataInput: IRollup::BatchDataInput {
                version: 2,
                parentBatchHeader: Bytes::copy_from_slice(parent_batch_header),
                lastBlockNumber: last_block_number,
                numL1Messages: 0,
                prevStateRoot: B256::ZERO,
                postStateRoot: B256::ZERO,
                withdrawalRoot: B256::ZERO,
            },
        }
        .abi_encode()
    }

    fn pre_submitter_commit_batch(parent_batch_header: &[u8], last_block_number: u64) -> Vec<u8> {
        PreSubmitterRollup::commitBatchCall {
            batchDataInput: pre_submitter::IRollup::BatchDataInput {
                version: 2,
                parentBatchHeader: Bytes::copy_from_slice(parent_batch_header),
                lastBlockNumber: last_block_number,
                numL1Messages: 0,
                prevStateRoot: B256::ZERO,
                postStateRoot: B256::ZERO,
                withdrawalRoot: B256::ZERO,
            },
            batchSignatureInput: pre_submitter::IRollup::BatchSignatureInput {
                signedSequencersBitmap: U256::ZERO,
                sequencerSets: Bytes::new(),
                signature: Bytes::new(),
            },
        }
        .abi_encode()
    }

    #[test]
    fn decodes_fixed_pre_and_post_upgrade_commit_batch_calldata() {
        let pre = decode_commit_batch(&calldata(PRE_SUBMITTER_COMMIT_BATCH)).unwrap();
        let post = decode_commit_batch(&calldata(POST_SUBMITTER_COMMIT_BATCH)).unwrap();

        assert_eq!(pre, post);
        assert_eq!(post.last_block_number, 12_345);
        assert_eq!(post.parent_batch_header.as_ref(), [1, 2, 3]);
        assert_eq!(Rollup::commitBatchCall::SELECTOR, [0x41, 0xf7, 0x56, 0xda]);
        assert_eq!(PreSubmitterRollup::commitBatchCall::SELECTOR, [0x42, 0x88, 0x68, 0xb5]);
    }

    #[test]
    fn rejects_non_commit_batch_calldata() {
        assert!(decode_commit_batch(&[]).is_err());
        assert!(decode_commit_batch(&[0xde, 0xad, 0xbe, 0xef]).is_err());
        assert!(decode_commit_batch(&[0x67, 0xca, 0xa3, 0x7a]).is_err());
        assert!(decode_commit_batch(&[0x15, 0x44, 0xba, 0x3a]).is_err());
        assert!(decode_commit_batch(&[0x4e, 0x8f, 0x1d, 0x67]).is_err());
    }

    #[test]
    fn decodes_and_assembles_old_new_new_cutover_window() {
        let old_calldata = pre_submitter_commit_batch(b"old-parent", 99);
        let target_calldata = post_submitter_commit_batch(b"target-parent", 109);
        let next_calldata = post_submitter_commit_batch(b"target-header", 119);
        assert_eq!(&old_calldata[..4], PreSubmitterRollup::commitBatchCall::SELECTOR);
        assert_eq!(&target_calldata[..4], Rollup::commitBatchCall::SELECTOR);
        assert_eq!(&next_calldata[..4], Rollup::commitBatchCall::SELECTOR);

        let old = decode_commit_batch(&old_calldata).unwrap();
        let target = decode_commit_batch(&target_calldata).unwrap();
        let next = decode_commit_batch(&next_calldata).unwrap();

        assert_eq!(target_batch_block_range(&old, &target), (100, 109));
        assert_eq!(target_batch_header_from_next_commit(&next).as_ref(), b"target-header");
        assert_ne!(target.parent_batch_header, target_batch_header_from_next_commit(&next));
    }
}
