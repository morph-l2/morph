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
        data_hasher.update(blocks.last().unwrap().header.number.to_be_bytes::<32>());
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

    /// Chain ID of this chunk
    pub fn chain_id(&self) -> u64 {
        self.chain_id
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
