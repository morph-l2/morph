//! ABIs
//!
//! Contract ABIs are refactored into their own module to gracefully deal with allowing missing docs on the abigen macro.
#![allow(missing_docs)]

use ethers::{abi::AbiDecode, contract::EthCall, types::Bytes};
use eyre::{eyre, Result};

pub mod rollup_abi {
    use ethers::prelude::abigen;
    abigen!(Rollup, "src/abi/Rollup.json");
}

pub mod submitter_abi {
    use ethers::prelude::abigen;
    abigen!(Submitter, "src/abi/Submitter.json");
}

pub mod pre_submitter_rollup_abi {
    use ethers::prelude::abigen;
    abigen!(PreSubmitterRollup, "src/abi/pre_submitter/Rollup.json");
}

#[derive(Clone, Debug, Eq, PartialEq)]
pub struct CommitBatchInput {
    pub version: u8,
    pub parent_batch_header: Bytes,
    pub last_block_number: u64,
    pub num_l1_messages: u16,
    pub prev_state_root: [u8; 32],
    pub post_state_root: [u8; 32],
    pub withdrawal_root: [u8; 32],
}

macro_rules! commit_batch_input {
    ($input:expr) => {{
        let input = $input;
        CommitBatchInput {
            version: input.version,
            parent_batch_header: input.parent_batch_header,
            last_block_number: input.last_block_number,
            num_l1_messages: input.num_l1_messages,
            prev_state_root: input.prev_state_root,
            post_state_root: input.post_state_root,
            withdrawal_root: input.withdrawal_root,
        }
    }};
}

/// Decode only the pre- or post-upgrade commitBatch method supported by challenge.
pub fn decode_commit_batch(data: &[u8]) -> Result<CommitBatchInput> {
    use pre_submitter_rollup_abi::CommitBatchCall as PreSubmitterCommitBatchCall;
    use rollup_abi::CommitBatchCall;

    let selector = data.get(..4).ok_or_else(|| eyre!("rollup calldata is shorter than a selector"))?;
    if selector == CommitBatchCall::selector() {
        let call = CommitBatchCall::decode(data)?;
        return Ok(commit_batch_input!(call.batch_data_input));
    }
    if selector == PreSubmitterCommitBatchCall::selector() {
        let call = PreSubmitterCommitBatchCall::decode(data)?;
        return Ok(commit_batch_input!(call.batch_data_input));
    }

    Err(eyre!("unsupported rollup calldata selector"))
}

#[cfg(test)]
mod tests {
    use super::*;
    use pre_submitter_rollup_abi::CommitBatchCall as PreSubmitterCommitBatchCall;
    use rollup_abi::CommitBatchCall;

    const PRE_SUBMITTER_COMMIT_BATCH: &str = "0x428868b500000000000000000000000000000000000000000000000000000000000000400000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000030390000000000000000000000000000000000000000000000000000000000000007111111111111111111111111111111111111111111111111111111111111111122222222222222222222222222222222222222222222222222222222222222223333333333333333333333333333333333333333333333333333333333333333000000000000000000000000000000000000000000000000000000000000000301020300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000009000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000002aabb0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000002ccdd000000000000000000000000000000000000000000000000000000000000";
    const POST_SUBMITTER_COMMIT_BATCH: &str = "0x41f756da0000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000003039000000000000000000000000000000000000000000000000000000000000000711111111111111111111111111111111111111111111111111111111111111112222222222222222222222222222222222222222222222222222222222222222333333333333333333333333333333333333333333333333333333333333333300000000000000000000000000000000000000000000000000000000000000030102030000000000000000000000000000000000000000000000000000000000";

    fn calldata(value: &str) -> Vec<u8> {
        ethers::utils::hex::decode(value.trim_start_matches("0x")).unwrap()
    }

    #[test]
    fn decodes_fixed_pre_and_post_upgrade_commit_batch_calldata() {
        let pre = decode_commit_batch(&calldata(PRE_SUBMITTER_COMMIT_BATCH)).unwrap();
        let post = decode_commit_batch(&calldata(POST_SUBMITTER_COMMIT_BATCH)).unwrap();

        assert_eq!(pre, post);
        assert_eq!(post.version, 2);
        assert_eq!(post.last_block_number, 12_345);
        assert_eq!(post.parent_batch_header.as_ref(), [1, 2, 3]);
        assert_eq!(CommitBatchCall::selector(), [0x41, 0xf7, 0x56, 0xda]);
        assert_eq!(PreSubmitterCommitBatchCall::selector(), [0x42, 0x88, 0x68, 0xb5]);
    }

    #[test]
    fn rejects_non_commit_batch_calldata() {
        assert!(decode_commit_batch(&[]).is_err());
        assert!(decode_commit_batch(&[0xde, 0xad, 0xbe, 0xef]).is_err());
        assert!(decode_commit_batch(&[0x67, 0xca, 0xa3, 0x7a]).is_err());
        assert!(decode_commit_batch(&[0x15, 0x44, 0xba, 0x3a]).is_err());
        assert!(decode_commit_batch(&[0x4e, 0x8f, 0x1d, 0x67]).is_err());
    }
}
