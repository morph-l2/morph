use std::rc::Rc;

use alloy::primitives::keccak256;
use sbv_core::{EvmExecutorBuilder, HardforkConfig, VerificationError};
use sbv_primitives::{types::BlockTrace, zk_trie::ZkMemoryDb, Block, B256};
use sbv_utils::{cycle_track, dev_error, dev_info};

use super::Verifier;
// use Verifier;
pub struct EVMVerifier;

impl Verifier for EVMVerifier {
    fn verify(l2_trace: &BlockTrace) -> Result<B256, VerificationError> {
        // let versioned_hash = commitment_to_version_hash(&commitment.clone().try_into().unwrap());
        // let (x, y) = eip4844::proof_of_equivalence(&input.taiko.tx_data, &versioned_hash)?;
        // ct.end();
        // let verified = eip4844::verify_kzg_proof_impl(
        //     commitment.clone().try_into().unwrap(),
        //     x,
        //     y,
        //     input
        //         .taiko
        //         .blob_proof
        //         .clone()
        //         .map(|p| TryInto::<[u8; 48]>::try_into(p).unwrap())
        //         .unwrap(),
        // )?;

        Ok(B256::default())
    }
}
