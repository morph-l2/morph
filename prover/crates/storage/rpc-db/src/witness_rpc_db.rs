use std::{
    marker::PhantomData,
    sync::{Arc, RwLock},
};

use alloy_consensus::Header;
use alloy_primitives::{map::HashMap, B256};
use alloy_provider::{Network, Provider};
use alloy_rlp::Decodable;
use alloy_rpc_types::BlockId;
use prover_mpt::EthereumState;

use revm::primitives::keccak256;
use revm::state::{AccountInfo, Bytecode};
use reth_storage_errors::{db::DatabaseError, provider::ProviderError};
use revm::database::DatabaseRef;
use revm::primitives::{Address, U256};

use crate::error::RpcDbError;

/// A database that fetches data via `debug_executionWitness` RPC method.
///
/// Unlike [`BasicRpcDb`](crate::basic_rpc_db::BasicRpcDb) which fetches data per-account
/// via `eth_getProof`, this struct pre-fetches all required state data in a single call,
/// making it much more efficient for full-block execution.
#[derive(Debug)]
pub struct ExecutionWitnessRpcDb<P, N> {
    /// The provider which fetches data.
    pub provider: P,
    /// Chain ID used by Morph's `BLOCKHASH` opcode behavior.
    pub chain_id: u64,
    /// The block to fetch data from (parent block number).
    pub block_number: u64,
    /// The state root to fetch data from.
    pub state_root: B256,
    /// The cached state built from execution witness.
    pub state: EthereumState,
    /// The cached bytecodes indexed by code hash.
    pub codes: HashMap<B256, Bytecode>,
    /// The cached ancestor headers indexed by block number.
    pub ancestor_headers: Arc<RwLock<HashMap<u64, Header>>>,

    phantom: PhantomData<N>,
}

impl<P: Provider<N> + Clone, N: Network> ExecutionWitnessRpcDb<P, N> {
    /// Create a new [`ExecutionWitnessRpcDb`].
    ///
    /// This fetches the `debug_executionWitness` data for the block
    /// at `block_number + 1` (i.e., the next block), and builds the
    /// [`EthereumState`] from the witness data.
    pub async fn new(
        provider: P,
        chain_id: u64,
        block_number: u64,
        state_root: B256,
    ) -> Result<Self, RpcDbError> {
        // Fetch the execution witness for the block at block_number + 1.
        // The witness contains all state trie nodes, codes, and ancestor headers
        // needed for executing the block.
        // Uses raw_request to call `debug_executionWitness` RPC method directly,
        // avoiding the need for the `debug-api` feature on alloy-provider.
        let execution_witness: alloy_rpc_types_debug::ExecutionWitness = provider
            .raw_request::<(BlockId,), alloy_rpc_types_debug::ExecutionWitness>(
                "debug_executionWitness".into(),
                ((block_number + 1).into(),),
            )
            .await
            .map_err(RpcDbError::Transport)?;

        // Build the EthereumState from the execution witness.
        let state = EthereumState::from_execution_witness(&execution_witness, state_root);

        // Decode and index bytecodes by their code hash.
        let codes: HashMap<B256, Bytecode> = execution_witness
            .codes
            .iter()
            .map(|encoded: &alloy_primitives::Bytes| {
                let bytecode = Bytecode::new_raw(encoded.clone());
                (keccak256(encoded), bytecode)
            })
            .collect();

        // Decode and index ancestor headers by their block number.
        let ancestor_headers: HashMap<u64, Header> = execution_witness
            .headers
            .iter()
            .filter_map(|encoded: &alloy_primitives::Bytes| {
                let header = Header::decode(&mut encoded.as_ref())
                    .expect("Valid RLP-encoded header in witness");
                Some((header.number, header))
            })
            .collect();

        let db = Self {
            provider,
            chain_id,
            block_number,
            state_root,
            state,
            codes,
            ancestor_headers: Arc::new(RwLock::new(ancestor_headers)),
            phantom: PhantomData,
        };

        Ok(db)
    }

    /// Returns all bytecodes indexed by code hash.
    pub fn bytecodes(&self) -> Vec<Bytecode> {
        self.codes.values().cloned().collect()
    }
}

impl<P: Provider<N> + Clone, N: Network> DatabaseRef for ExecutionWitnessRpcDb<P, N> {
    type Error = ProviderError;

    /// Get basic account information.
    ///
    /// The witness-backed DB is read-only: all state was pre-fetched from
    /// `debug_executionWitness`, so we look up the account directly in the
    /// in-memory [`EthereumState`] trie.
    fn basic_ref(&self, address: Address) -> Result<Option<AccountInfo>, Self::Error> {
        use alloy_primitives::keccak256 as alloy_keccak256;
        use reth_trie::TrieAccount;

        let hashed = alloy_keccak256(address);
        let trie_account = self
            .state
            .state_trie
            .get_rlp::<TrieAccount>(hashed.as_slice())
            .map_err(|e| ProviderError::Database(DatabaseError::Other(e.to_string())))?;

        let Some(account) = trie_account else {
            return Ok(None);
        };

        // Look up bytecode by code hash.
        let code = if account.code_hash == revm::primitives::KECCAK_EMPTY
            || account.code_hash == B256::ZERO
        {
            Some(Bytecode::new())
        } else {
            self.codes.get(&account.code_hash).cloned().map(Some).unwrap_or(Some(Bytecode::new()))
        };

        Ok(Some(AccountInfo {
            account_id: None,
            nonce: account.nonce,
            balance: account.balance,
            code_hash: account.code_hash,
            code,
        }))
    }

    /// Get account code by its hash.
    fn code_by_hash_ref(&self, code_hash: B256) -> Result<Bytecode, Self::Error> {
        self.codes
            .get(&code_hash)
            .cloned()
            .ok_or_else(|| ProviderError::Database(DatabaseError::Other(
                format!("bytecode not found for hash {code_hash:?}"),
            )))
    }

    /// Get storage value of address at index.
    fn storage_ref(&self, address: Address, index: U256) -> Result<U256, Self::Error> {
        use alloy_primitives::keccak256 as alloy_keccak256;

        let hashed_address = alloy_keccak256(address);
        let hashed_slot = alloy_keccak256(index.to_be_bytes::<32>());

        let storage_trie = match self.state.storage_tries.get(&hashed_address) {
            Some(trie) => trie,
            None => return Ok(U256::ZERO),
        };

        let value = storage_trie
            .get_rlp::<U256>(hashed_slot.as_slice())
            .map_err(|e| ProviderError::Database(DatabaseError::Other(e.to_string())))?
            .unwrap_or(U256::ZERO);

        Ok(value)
    }

    /// Get block hash by its number.
    ///
    /// Uses the same Morph-specific `keccak(chain_id || block_number)` semantics
    /// as [`BasicRpcDb`](crate::basic_rpc_db::BasicRpcDb).
    fn block_hash_ref(&self, number: u64) -> Result<B256, Self::Error> {
        let current_block_number = self.block_number + 1;
        if number >= current_block_number {
            return Ok(B256::ZERO);
        }
        if current_block_number.saturating_sub(number) > 256 {
            return Ok(B256::ZERO);
        }

        let mut buf = [0u8; 16];
        buf[..8].copy_from_slice(&self.chain_id.to_be_bytes());
        buf[8..].copy_from_slice(&number.to_be_bytes());
        Ok(alloy_primitives::utils::keccak256(buf))
    }
}
