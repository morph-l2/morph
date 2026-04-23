use crate::types::input::BlockInput;
use alloy_primitives::Keccak256;
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
    chain_id: u64,
    prev_state_root: B256,
    pub post_state_root: B256,
    /// withdraw_root
    pub withdraw_root: Option<B256>,
    /// sequencer_root
    pub sequencer_root: Option<B256>,
    data_hash: B256,
}

impl BatchInfo {
    /// Construct by block inputs
    pub fn from_block_inputs(
        block_inputs: &[BlockInput],
        post_state_root: B256,
        withdraw_root: B256,
        sequencer_root: B256,
    ) -> Self {
        let blocks = block_inputs.iter().map(|x| x.current_block.clone()).collect::<Vec<_>>();
        let chain_id = blocks.first().unwrap().chain_id;
        let prev_state_root = blocks.first().unwrap().prev_state_root;

        let mut data_hasher = Keccak256::new();
        data_hasher.update(blocks.last().unwrap().header.number.to::<u64>().to_be_bytes());
        let num_l1_txs: u16 = blocks.iter().map(|x| x.num_l1_txs()).sum::<u64>() as u16;
        data_hasher.update(num_l1_txs.to_be_bytes());

        for block in blocks.iter() {
            block.hash_l1_msg(&mut data_hasher);
        }
        let l1_data_hash = data_hasher.finalize();

        BatchInfo {
            chain_id,
            prev_state_root,
            post_state_root,
            withdraw_root: Some(withdraw_root),
            sequencer_root: Some(sequencer_root),
            data_hash: l1_data_hash,
        }
    }

    /// Public input hash for a given batch is defined as
    /// keccak(
    ///     chain id ||
    ///     prev state root ||
    ///     post state root ||
    ///     withdraw root ||
    ///     sequencer root ||
    ///     txdata hash ||
    ///     blob versioned hash
    /// )
    pub fn public_input_hash(&self, versioned_hash: &B256) -> B256 {
        let mut hasher = Keccak256::new();

        hasher.update(self.chain_id.to_be_bytes());
        hasher.update(self.prev_state_root.as_slice());
        hasher.update(self.post_state_root.as_slice());
        hasher.update(self.withdraw_root.unwrap().as_slice());
        hasher.update(self.sequencer_root.unwrap().as_slice());
        hasher.update(self.data_hash.as_slice());
        hasher.update(versioned_hash.as_slice());

        hasher.finalize()
    }

    /// V2 public input hash: uses keccak256(hash[0] || ... || hash[N-1]) as blob input
    pub fn public_input_hash_v2(&self, blob_hashes: &[B256]) -> B256 {
        let mut blob_hasher = Keccak256::new();
        for h in blob_hashes {
            blob_hasher.update(h.as_slice());
        }
        let blob_hashes_hash: B256 = blob_hasher.finalize();

        let mut hasher = Keccak256::new();
        hasher.update(self.chain_id.to_be_bytes());
        hasher.update(self.prev_state_root.as_slice());
        hasher.update(self.post_state_root.as_slice());
        hasher.update(self.withdraw_root.unwrap().as_slice());
        hasher.update(self.sequencer_root.unwrap().as_slice());
        hasher.update(self.data_hash.as_slice());
        hasher.update(blob_hashes_hash.as_slice());
        hasher.finalize()
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
            sequencer_root: Some(B256::ZERO),
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

    /// Sequencer root after this chunk
    pub fn sequencer_root(&self) -> B256 {
        self.sequencer_root.expect("get sequencer_root")
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

    /// V2 aggregated hash for a single blob: keccak256(h0) != h0 (not backward-compatible with V1).
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

    /// V2 and V1 produce the same result only when blob_hashes_hash accidentally equals
    /// the raw versioned hash — which should never happen in practice.
    /// This test confirms the structural difference by construction.
    #[test]
    fn test_public_input_hash_v2_vs_v1_structural_difference() {
        let batch = BatchInfo::test_instance(TEST_CHAIN_ID);
        let h0 = make_hash(0x1234);

        // V1: uses h0 directly as blob input
        let v1 = batch.public_input_hash(&h0);
        // V2: uses keccak256(h0) as blob input — structurally different
        let v2 = batch.public_input_hash_v2(&[h0]);
        assert_ne!(v1, v2);

        // Confirm: if we manually pass keccak256(h0) into V1, it matches V2
        let agg = keccak256(h0.as_slice());
        let v1_with_agg = batch.public_input_hash(&agg);
        assert_eq!(v1_with_agg, v2, "V2 is equivalent to V1 with pre-aggregated hash");
    }
}
