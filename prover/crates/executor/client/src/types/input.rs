use alloy_consensus::BlockHeader;
use alloy_primitives::{map::HashMap, U256};
use morph_primitives::Block;
use prover_mpt::EthereumState;
use prover_primitives::Address;
use prover_storage_witness::TrieDB;
use reth_trie::{TrieAccount, EMPTY_ROOT_HASH};
use revm::{primitives::keccak256, state::Bytecode};
use serde::{Deserialize, Serialize};
use serde_with::serde_as;

use crate::types::error::ClientError;

#[serde_as]
#[derive(Clone, Debug, Default, Serialize, Deserialize)]
pub struct BlobInfo {
    pub blob_data: Vec<u8>,
    pub commitment: Vec<u8>,
    pub proof: Vec<u8>,
}

#[serde_as]
#[derive(Debug, Clone, Serialize, Deserialize, PartialEq, Eq)]
pub struct BlockInput {
    /// l2 block info
    pub current_block: Block,

    /// state as of the parent block.
    pub parent_state: EthereumState,

    /// Account bytecodes.
    pub bytecodes: Vec<Bytecode>,

    /// ChainId
    pub chain_id: u64,
}

impl BlockInput {
    fn validate_parent_state(&self) -> Result<(), ClientError> {
        for (hashed_address, storage_trie) in &self.parent_state.storage_tries {
            let account =
                self.parent_state.state_trie.get_rlp::<TrieAccount>(hashed_address.as_slice())?;
            let storage_root = account.map_or(EMPTY_ROOT_HASH, |account| account.storage_root);
            if storage_trie.hash() != storage_root {
                return Err(ClientError::MismatchedStorageRoot);
            }
        }

        Ok(())
    }

    pub fn witness_db(&self) -> Result<TrieDB<'_>, ClientError> {
        self.validate_parent_state()?;

        let bytecodes_by_hash =
            self.bytecodes.iter().map(|code| (code.hash_slow(), code)).collect::<HashMap<_, _>>();

        Ok(TrieDB::new(
            &self.parent_state,
            bytecodes_by_hash,
            u64::default(),
            self.current_block.header.number(),
        ))
    }

    /// Get storage value of address at index.
    pub fn get_storage_value(&self, address: Address, index: U256) -> Result<U256, ClientError> {
        let hashed_address = keccak256(address);

        let storage_trie = self
            .parent_state
            .storage_tries
            .get(&hashed_address)
            .ok_or(ClientError::MissingTrie(address))?;
        let account =
            self.parent_state.state_trie.get_rlp::<TrieAccount>(hashed_address.as_slice())?;
        let storage_root = account.map_or(EMPTY_ROOT_HASH, |account| account.storage_root);
        if storage_trie.hash() != storage_root {
            return Err(ClientError::MismatchedStorageRoot);
        }

        Ok(storage_trie
            .get_rlp::<U256>(keccak256(index.to_be_bytes::<32>()).as_slice())?
            .unwrap_or_default())
    }
}

#[serde_as]
#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct ExecutorInput {
    pub block_inputs: Vec<BlockInput>,
    pub blob_infos: Vec<BlobInfo>,
    #[serde(default)]
    pub batch_version: u8,
}
