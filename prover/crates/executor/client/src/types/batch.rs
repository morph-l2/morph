use crate::types::input::BlockInput;
use alloy_consensus::{BlockHeader, SignableTransaction};
use alloy_primitives::Keccak256;
use morph_primitives::MorphTxEnvelope;
use prover_primitives::predeployed::l2_to_l1_message::{WITHDRAW_ROOT_ADDRESS, WITHDRAW_ROOT_SLOT};
use revm::primitives::B256;

/// BatchInfo is metadata of chunk, with following fields:
/// - state root before this chunk
/// - state root after this chunk
/// - the withdraw root after this chunk
/// - the data hash of this chunk
/// - the tx data hash of this chunk
/// - flattened L2 tx bytes hash
#[derive(Debug)]
pub struct BatchInfo {
    pub chain_id: u64,
    pub prev_state_root: B256,
    pub post_state_root: B256,
    /// withdraw_root
    pub withdraw_root: Option<B256>,
    pub data_hash: B256,
}

impl BatchInfo {
    /// Construct by block inputs
    pub fn from_block_inputs(
        prev_state_root: B256,
        block_inputs: &[BlockInput],
    ) -> Result<Self, crate::types::error::ClientError> {
        let chain_id = block_inputs.first().map(|b| b.chain_id).unwrap_or(2818);
        let latest_block_input = block_inputs.last().unwrap();

        // The post-withdraw-root is required for public inputs and is derived from the state of
        // the last verified block.
        let post_withdraw_root =
            latest_block_input.get_storage_value(WITHDRAW_ROOT_ADDRESS, WITHDRAW_ROOT_SLOT)?;

        let post_state_root = latest_block_input.current_block.state_root();
        let data_hash = Self::calculate_data_hash(block_inputs);

        Ok(BatchInfo {
            chain_id,
            prev_state_root,
            post_state_root,
            withdraw_root: Some(post_withdraw_root.into()),
            data_hash,
        })
    }

    fn calculate_data_hash(block_inputs: &[BlockInput]) -> B256 {
        let last_block_input = block_inputs.last().expect("block inputs must not be empty");
        let total_num_l1_txs = block_inputs
            .iter()
            .flat_map(|input| &input.current_block.body.transactions)
            .filter(|tx| tx.is_l1_msg())
            .count();

        let mut data_hasher = Keccak256::new();
        data_hasher.update(last_block_input.current_block.header.number().to_be_bytes());
        data_hasher.update((total_num_l1_txs as u16).to_be_bytes());
        for input in block_inputs {
            Self::hash_l1_msg(&input.current_block.body.transactions, &mut data_hasher);
        }
        data_hasher.finalize()
    }

    /// Hashes the L1 messages in the block using the provided hasher.
    pub fn hash_l1_msg(transactions: &Vec<MorphTxEnvelope>, hasher: &mut Keccak256) {
        for tx in transactions {
            if let MorphTxEnvelope::L1Msg(l1) = tx {
                hasher.update(l1.signature_hash());
            }
        }
    }

    /// Public input hash for V0/V1, using the single versioned hash directly.
    pub fn public_input_hash(&self, versioned_hash: &B256) -> B256 {
        self.public_input_hash_with_blob_input(*versioned_hash)
    }

    /// Public input hash for V2, using keccak256(hash[0] || ... || hash[N-1]) as blob input.
    pub fn public_input_hash_v2(&self, blob_hashes: &[B256]) -> B256 {
        let mut blob_hasher = Keccak256::new();
        for h in blob_hashes {
            blob_hasher.update(h.as_slice());
        }
        let blob_hashes_hash: B256 = blob_hasher.finalize();

        self.public_input_hash_with_blob_input(blob_hashes_hash)
    }

    fn public_input_hash_with_blob_input(&self, blob_input: B256) -> B256 {
        let mut hasher = Keccak256::new();
        hasher.update(self.chain_id.to_be_bytes());
        hasher.update(self.prev_state_root.as_slice());
        hasher.update(self.post_state_root.as_slice());
        hasher.update(self.withdraw_root.unwrap().as_slice());
        hasher.update(B256::ZERO.as_slice());
        hasher.update(self.data_hash.as_slice());
        hasher.update(blob_input.as_slice());
        hasher.finalize()
    }

    /// Calculates the public input hash for a batch whose blob cannot be decoded or does not
    /// match the block data. Since the batch is not executed, all post roots remain unchanged.
    pub fn public_input_hash_for_invalid_blob(
        blob_hashes: &[B256],
        block_inputs: &[BlockInput],
        batch_version: u8,
    ) -> Result<B256, crate::types::error::ClientError> {
        let first_block_input = block_inputs.first().expect("block inputs must not be empty");
        let original_state_root: B256 = first_block_input.parent_state.state_root().into();
        let original_withdraw_root =
            first_block_input.get_storage_value(WITHDRAW_ROOT_ADDRESS, WITHDRAW_ROOT_SLOT)?;

        let data_hash = Self::calculate_data_hash(block_inputs);

        let batch_info = BatchInfo {
            chain_id: first_block_input.chain_id,
            prev_state_root: original_state_root,
            post_state_root: original_state_root,
            withdraw_root: Some(original_withdraw_root.into()),
            data_hash,
        };

        Ok(match batch_version {
            0 | 1 => batch_info.public_input_hash(&blob_hashes[0]),
            2 => batch_info.public_input_hash_v2(blob_hashes),
            _ => unreachable!("batch version is validated before calculating the hash"),
        })
    }

    /// Chain ID of this chunk
    pub fn chain_id(&self) -> u64 {
        self.chain_id
    }

    #[cfg(test)]
    fn test_instance(chain_id: u64) -> Self {
        BatchInfo {
            chain_id,
            prev_state_root: B256::ZERO,
            post_state_root: B256::ZERO,
            withdraw_root: Some(B256::ZERO),
            data_hash: B256::ZERO,
        }
    }

    /// State root before this chunk
    pub fn prev_state_root(&self) -> B256 {
        self.prev_state_root
    }

    /// State root after this chunk
    pub fn post_state_root(&self) -> B256 {
        self.post_state_root
    }

    /// Withdraw root after this chunk
    pub fn withdraw_root(&self) -> B256 {
        self.withdraw_root.expect("get withdraw_root")
    }

    /// Data hash of this chunk
    pub fn data_hash(&self) -> B256 {
        self.data_hash
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use alloy_primitives::keccak256;

    // LAYER_2_CHAIN_ID used in Rollup.sol test environment
    const TEST_CHAIN_ID: u64 = 53077;

    fn make_hash(val: u64) -> B256 {
        let mut b = [0u8; 32];
        b[24..].copy_from_slice(&val.to_be_bytes());
        B256::from(b)
    }

    /// A single blob is encoded differently by the legacy and aggregated formats.
    #[test]
    fn test_public_input_hash_v2_single_blob_differs_from_v1() {
        let batch = BatchInfo::test_instance(TEST_CHAIN_ID);
        let h0 = make_hash(0xBEEF);

        let v1_hash = batch.public_input_hash(&h0);
        let v2_hash = batch.public_input_hash_v2(&[h0]);

        assert_ne!(v1_hash, v2_hash, "V2 single-blob must differ from V1");
    }

    /// V2 aggregated hash for two blobs: keccak256(h0 || h1) matches contract formula.
    #[test]
    fn test_public_input_hash_v2_two_blobs_matches_contract() {
        let batch = BatchInfo::test_instance(TEST_CHAIN_ID);
        let h0 = make_hash(0xAAAA);
        let h1 = make_hash(0xBBBB);

        // Replicate contract formula: aggregatedBlobHash = keccak256(h0 || h1)
        let mut concat = [0u8; 64];
        concat[..32].copy_from_slice(h0.as_slice());
        concat[32..].copy_from_slice(h1.as_slice());
        let aggregated = keccak256(&concat);

        // V2 public input uses aggregated as blob input
        let mut hasher = Keccak256::new();
        hasher.update(TEST_CHAIN_ID.to_be_bytes());
        hasher.update(B256::ZERO.as_slice()); // prev_state_root
        hasher.update(B256::ZERO.as_slice()); // post_state_root
        hasher.update(B256::ZERO.as_slice()); // withdraw_root
        hasher.update(B256::ZERO.as_slice()); // sequencer_root
        hasher.update(B256::ZERO.as_slice()); // data_hash
        hasher.update(aggregated.as_slice());
        let expected: B256 = hasher.finalize();

        let result = batch.public_input_hash_v2(&[h0, h1]);
        assert_eq!(result, expected, "V2 two-blob hash must match contract formula");
    }

    /// V2 aggregated hash for three blobs: keccak256(h0 || h1 || h2).
    #[test]
    fn test_public_input_hash_v2_three_blobs() {
        let batch = BatchInfo::test_instance(TEST_CHAIN_ID);
        let h0 = make_hash(0xAAAA);
        let h1 = make_hash(0xBBBB);
        let h2 = make_hash(0xCCCC);

        let mut concat = [0u8; 96];
        concat[..32].copy_from_slice(h0.as_slice());
        concat[32..64].copy_from_slice(h1.as_slice());
        concat[64..].copy_from_slice(h2.as_slice());
        let aggregated = keccak256(&concat);

        let mut hasher = Keccak256::new();
        hasher.update(TEST_CHAIN_ID.to_be_bytes());
        hasher.update(B256::ZERO.as_slice());
        hasher.update(B256::ZERO.as_slice());
        hasher.update(B256::ZERO.as_slice());
        hasher.update(B256::ZERO.as_slice());
        hasher.update(B256::ZERO.as_slice());
        hasher.update(aggregated.as_slice());
        let expected: B256 = hasher.finalize();

        let result = batch.public_input_hash_v2(&[h0, h1, h2]);
        assert_eq!(result, expected, "V2 three-blob hash must match contract formula");
    }

    /// V2 aggregated hash is order-sensitive: (h0,h1) != (h1,h0).
    #[test]
    fn test_public_input_hash_v2_order_sensitive() {
        let batch = BatchInfo::test_instance(TEST_CHAIN_ID);
        let h0 = make_hash(0xAAAA);
        let h1 = make_hash(0xBBBB);

        let fwd = batch.public_input_hash_v2(&[h0, h1]);
        let rev = batch.public_input_hash_v2(&[h1, h0]);
        assert_ne!(fwd, rev, "V2 aggregated hash must be order-sensitive");
    }

    #[test]
    fn test_public_input_hash_v2_matches_v1_with_preaggregated_hash() {
        let batch = BatchInfo::test_instance(TEST_CHAIN_ID);
        let h0 = make_hash(0x1234);

        let v2 = batch.public_input_hash_v2(&[h0]);
        let aggregated = keccak256(h0.as_slice());
        let v1_with_aggregated_hash = batch.public_input_hash(&aggregated);

        assert_eq!(v1_with_aggregated_hash, v2);
    }
}
