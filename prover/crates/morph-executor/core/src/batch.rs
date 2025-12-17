use prover_primitives::Block;
use revm::primitives::B256;
use std::rc::Rc;
use tiny_keccak::{Hasher, Keccak};

/// A chunk is a set of continuous blocks.
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
    post_state_root: B256,
    /// withdraw_root
    pub withdraw_root: Option<B256>,
    /// sequencer_root
    pub sequencer_root: Option<B256>,
    data_hash: B256,
}

impl BatchInfo {
    /// Construct by block traces
    pub fn from_block_traces<T: Block>(traces: &[T]) -> Self {
        let chain_id = traces.first().unwrap().chain_id();
        let prev_state_root = traces.first().expect("at least 1 block needed").root_before();
        let post_state_root = traces.last().expect("at least 1 block needed").root_after();

        let mut data_hasher = Keccak::v256();
        data_hasher.update(&traces.last().unwrap().number().to_be_bytes());
        let num_l1_txs: u16 = traces.iter().map(|x| x.num_l1_txs()).sum::<u64>() as u16;
        data_hasher.update(&num_l1_txs.to_be_bytes());

        for trace in traces.iter() {
            trace.hash_l1_msg(&mut data_hasher);
        }
        let mut data_hash = B256::ZERO;
        data_hasher.finalize(&mut data_hash.0);

        let info = BatchInfo {
            chain_id,
            prev_state_root,
            post_state_root,
            withdraw_root: None,
            sequencer_root: None,
            data_hash,
        };

        info
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
        let mut hasher = Keccak::v256();

        hasher.update(&self.chain_id.to_be_bytes());
        hasher.update(self.prev_state_root.as_slice());
        hasher.update(self.post_state_root.as_slice());
        hasher.update(self.withdraw_root.unwrap().as_slice());
        hasher.update(self.sequencer_root.unwrap().as_slice());
        hasher.update(self.data_hash.as_slice());
        hasher.update(versioned_hash.as_slice());

        let mut public_input_hash = B256::ZERO;
        hasher.finalize(&mut public_input_hash.0);
        public_input_hash
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
