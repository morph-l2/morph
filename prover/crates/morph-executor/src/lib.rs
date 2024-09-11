mod types;
mod verifier;
use alloy::primitives::keccak256;
use sbv_core::{EvmExecutorBuilder, HardforkConfig, VerificationError};
use sbv_primitives::{types::BlockTrace, zk_trie::ZkMemoryDb, Block, B256};
use sbv_utils::{cycle_track, dev_error, dev_info};
use std::rc::Rc;
use verifier::{evm::EVMVerifier, Verifier};

pub fn verify(l2_trace: &BlockTrace) -> Result<B256, VerificationError> {
    let rt = EVMVerifier::verify(l2_trace);
    rt
}
