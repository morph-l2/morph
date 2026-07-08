use std::{
    marker::PhantomData,
    sync::{Arc, RwLock},
};

use alloy_consensus::Header;
use alloy_primitives::{map::HashMap, B256};
use alloy_provider::{Network, Provider};
use alloy_rlp::Decodable;
use alloy_rpc_types::BlockId;
use morph_primitives::MorphHeader;
use prover_mpt::EthereumState;

use revm::primitives::keccak256;
use revm::state::{AccountInfo, Bytecode};
use reth_storage_errors::{db::DatabaseError, provider::ProviderError};
use revm::database::DatabaseRef;
use revm::primitives::{Address, U256};

use crate::account_proof::EIP1186AccountProofResponseCompat;
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
    ///
    /// `force_include` lists `(address, slot)` pairs whose storage trie must be present in the
    /// resulting state even if the block did not touch them. An execution witness only contains
    /// state that was accessed during block execution, so accounts read *outside* of execution
    /// (e.g. Morph predeploys read to derive batch public inputs) would otherwise be missing.
    /// For each pair we fetch an `eth_getProof` at `block_number` and merge its storage trie into
    /// the state.
    pub async fn new(
        provider: P,
        chain_id: u64,
        block_number: u64,
        state_root: B256,
        force_include: &[(Address, U256)],
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
        let mut state = EthereumState::from_execution_witness(&execution_witness, state_root);

        // Merge storage tries for accounts that must be present regardless of whether the block
        // touched them (see `force_include` docs). Grouping slots by address keeps this to one
        // `eth_getProof` per account.
        let mut slots_by_address: HashMap<Address, Vec<B256>> = HashMap::default();
        for (address, slot) in force_include {
            slots_by_address.entry(*address).or_default().push(B256::from(slot.to_be_bytes::<32>()));
        }
        for (address, slots) in slots_by_address {
            let proof = Self::eth_get_proof(&provider, address, slots, block_number).await?;
            state.insert_storage_trie_from_proof(&proof)?;
        }

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
        //
        // Morph-reth's `debug_executionWitness` RLP-encodes headers using the node's
        // native `MorphHeader` type, whose encoding differs from a standard Ethereum
        // header (it carries the extra `next_l1_msg_index` field). Decoding those bytes
        // as an `alloy_consensus::Header` fails with `UnexpectedLength`, so we decode the
        // `MorphHeader` wrapper and keep its inner Ethereum header.
        let mut ancestor_headers: HashMap<u64, Header> = HashMap::default();
        for encoded in execution_witness.headers.iter() {
            let header = MorphHeader::decode(&mut encoded.as_ref())
                .map_err(RpcDbError::HeaderDecodeError)?;
            ancestor_headers.insert(header.inner.number, header.inner);
        }

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

    /// Fetches an EIP-1186 proof for `address` and `slots` at `block_number`.
    async fn eth_get_proof(
        provider: &P,
        address: Address,
        slots: Vec<B256>,
        block_number: u64,
    ) -> Result<alloy_rpc_types::EIP1186AccountProofResponse, RpcDbError> {
        let compact_proof: EIP1186AccountProofResponseCompat = provider
            .raw_request::<(Address, Vec<B256>, BlockId), _>(
                "eth_getProof".into(),
                (address, slots, block_number.into()),
            )
            .await
            .map_err(|e| RpcDbError::GetProofError(address, e.to_string()))?;
        Ok(compact_proof.into())
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
