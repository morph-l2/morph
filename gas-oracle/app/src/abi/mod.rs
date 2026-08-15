//! ABIs
//!
//! Contract ABIs are refactored into their own module to gracefully deal with allowing missing docs
//! on the abigen macro.
#![allow(missing_docs)]

use ethers::{abi::AbiDecode, contract::EthCall};
use eyre::{eyre, Result};

pub mod gas_price_oracle_abi {
    use ethers::prelude::abigen;
    abigen!(GasPriceOracle, "src/abi/GasPriceOracle.json");
}

pub mod rollup_abi {
    use ethers::prelude::abigen;
    abigen!(Rollup, "src/abi/Rollup.json");
}

pub mod pre_submitter_rollup_abi {
    use ethers::prelude::abigen;
    abigen!(PreSubmitterRollup, "src/abi/pre_submitter/Rollup.json");
}

/// Decode the last block number from the pre- or post-upgrade commitBatch calldata.
pub fn decode_commit_batch_last_block_number(data: &[u8]) -> Result<u64> {
    use pre_submitter_rollup_abi::CommitBatchCall as PreSubmitterCommitBatchCall;
    use rollup_abi::CommitBatchCall;

    let selector =
        data.get(..4).ok_or_else(|| eyre!("rollup calldata is shorter than a selector"))?;
    if selector == CommitBatchCall::selector() {
        return Ok(CommitBatchCall::decode(data)?.batch_data_input.last_block_number);
    }
    if selector == PreSubmitterCommitBatchCall::selector() {
        return Ok(PreSubmitterCommitBatchCall::decode(data)?.batch_data_input.last_block_number);
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
        assert_eq!(
            decode_commit_batch_last_block_number(&calldata(PRE_SUBMITTER_COMMIT_BATCH)).unwrap(),
            12_345
        );
        assert_eq!(
            decode_commit_batch_last_block_number(&calldata(POST_SUBMITTER_COMMIT_BATCH)).unwrap(),
            12_345
        );
        assert_eq!(CommitBatchCall::selector(), [0x41, 0xf7, 0x56, 0xda]);
        assert_eq!(PreSubmitterCommitBatchCall::selector(), [0x42, 0x88, 0x68, 0xb5]);
    }

    #[test]
    fn rejects_non_commit_batch_calldata() {
        assert!(decode_commit_batch_last_block_number(&[]).is_err());
        assert!(decode_commit_batch_last_block_number(&[0xde, 0xad, 0xbe, 0xef]).is_err());
        assert!(decode_commit_batch_last_block_number(&[0x67, 0xca, 0xa3, 0x7a]).is_err());
        assert!(decode_commit_batch_last_block_number(&[0x15, 0x44, 0xba, 0x3a]).is_err());
        assert!(decode_commit_batch_last_block_number(&[0x4e, 0x8f, 0x1d, 0x67]).is_err());
    }
}
