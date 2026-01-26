use std::{
    collections::{BTreeMap, BTreeSet},
    marker::PhantomData,
    sync::{Arc, RwLock},
};

use crate::account_proof::{eip1186_proof_to_account_proof, EIP1186AccountProofResponseCompat};
use alloy_consensus::{BlockHeader, Header};
use alloy_primitives::{map::HashMap, StorageKey, U256};
use alloy_provider::{network::BlockResponse, Network, Provider};
use alloy_rpc_types::BlockId;
use async_trait::async_trait;
use prover_mpt::EthereumState;
use reth_storage_errors::{db::DatabaseError, provider::ProviderError};
use revm::database::BundleState;
use revm::database::DatabaseRef;
use revm::primitives::{Address, B256, KECCAK_EMPTY};
use revm::state::{AccountInfo, Bytecode};
use tracing::debug;

use crate::error::RpcDbError;

/// A database that fetches data from a [Provider] over a [Transport].
#[derive(Debug, Clone)]
pub struct BasicRpcDb<P, N> {
    /// The provider which fetches data.
    pub provider: P,
    /// Chain ID used by Morph's `BLOCKHASH` opcode behavior.
    pub chain_id: u64,
    /// The block to fetch data from.
    pub block_number: u64,
    ///The state root to fetch data from.
    pub state_root: B256,
    /// The cached accounts.
    pub accounts: Arc<RwLock<HashMap<Address, AccountInfo>>>,
    /// The cached storage values.
    pub storage: Arc<RwLock<HashMap<Address, HashMap<U256, U256>>>>,
    /// The oldest block whose header/hash has been requested.
    pub oldest_ancestor: Arc<RwLock<u64>>,

    phantom: PhantomData<N>,
}

impl<P: Provider<N> + Clone, N: Network> BasicRpcDb<P, N> {
    /// Create a new [`BasicRpcDb`].
    pub fn new(provider: P, chain_id: u64, block_number: u64, state_root: B256) -> Self {
        Self {
            provider,
            chain_id,
            block_number,
            state_root,
            accounts: Arc::new(RwLock::new(HashMap::with_hasher(Default::default()))),
            storage: Arc::new(RwLock::new(HashMap::with_hasher(Default::default()))),
            oldest_ancestor: Arc::new(RwLock::new(block_number)),
            phantom: PhantomData,
        }
    }

    /// Fetch the [AccountInfo] for an [Address].
    pub async fn fetch_account_info(&self, address: Address) -> Result<AccountInfo, RpcDbError> {
        debug!("fetching account info for address: {}", address);

        // Fetch the proof for the account.
        let proof = self
            .eth_get_proof(address, Vec::<alloy_primitives::StorageKey>::new(), self.block_number)
            .await?;

        // Fetch the code of the account.
        let code = self
            .provider
            .get_code_at(address)
            .number(self.block_number)
            .await
            .map_err(|e| RpcDbError::GetCodeError(address, e.to_string()))?;

        // Construct the account info & write it to the log.
        let bytecode = Bytecode::new_raw(code);

        // Normalize code_hash for REVM compatibility:
        // RPC response for getProof method for non-existing (unused) EOAs may contain B256::ZERO
        // for code_hash, but REVM expects KECCAK_EMPTY
        let code_hash = if proof.code_hash == B256::ZERO { KECCAK_EMPTY } else { proof.code_hash };

        let account_info = AccountInfo {
            nonce: proof.nonce,
            balance: proof.balance,
            code_hash,
            code: Some(bytecode.clone()),
        };

        // Record the account info to the state.
        self.accounts
            .write()
            .map_err(|_| RpcDbError::Poisoned)?
            .insert(address, account_info.clone());

        Ok(account_info)
    }

    async fn eth_get_proof(
        &self,
        address: Address,
        keys: Vec<StorageKey>,
        block_number: u64,
    ) -> Result<alloy_rpc_types::EIP1186AccountProofResponse, RpcDbError> {
        let compact_proof: EIP1186AccountProofResponseCompat = self
            .provider
            .raw_request::<(Address, Vec<alloy_primitives::StorageKey>, BlockId), _>(
                "eth_getProof".into(),
                (address, keys, block_number.into()),
            )
            .await
            .map_err(|e| RpcDbError::GetProofError(address, e.to_string()))?;
        let proof: alloy_rpc_types::EIP1186AccountProofResponse = compact_proof.into();
        Ok(proof)
    }

    /// Fetch the storage value at an [Address] and [U256] index.
    pub async fn fetch_storage_at(
        &self,
        address: Address,
        index: U256,
    ) -> Result<U256, RpcDbError> {
        debug!("fetching storage value at address: {}, index: {}", address, index);

        // Fetch the storage value.
        let value = self
            .provider
            .get_storage_at(address, index)
            .number(self.block_number)
            .await
            .map_err(|e| RpcDbError::GetStorageError(address, index, e.to_string()))?;

        // Record the storage value to the state.
        let mut storage_values = self.storage.write().map_err(|_| RpcDbError::Poisoned)?;
        let entry = storage_values.entry(address).or_default();
        entry.insert(index, value);

        Ok(value)
    }

    /// Gets all the state keys used. The client uses this to read the actual state data from tries.
    pub fn get_state_requests(&self) -> HashMap<Address, Vec<U256>> {
        let accounts = self.accounts.read().unwrap();
        let storage = self.storage.read().unwrap();

        accounts
            .keys()
            .chain(storage.keys())
            .map(|&address| {
                let storage_keys_for_address: BTreeSet<U256> = storage
                    .get(&address)
                    .map(|storage_map| storage_map.keys().cloned().collect())
                    .unwrap_or_default();

                (address, storage_keys_for_address.into_iter().collect())
            })
            .collect()
    }
}

impl<P: Provider<N> + Clone, N: Network> DatabaseRef for BasicRpcDb<P, N> {
    type Error = ProviderError;

    /// Get basic account information.
    fn basic_ref(&self, address: Address) -> Result<Option<AccountInfo>, Self::Error> {
        let handle = tokio::runtime::Handle::try_current().map_err(|_| {
            ProviderError::Database(DatabaseError::Other("no tokio runtime found".to_string()))
        })?;
        let result =
            tokio::task::block_in_place(|| handle.block_on(self.fetch_account_info(address)));

        let account_info =
            result.map_err(|e| ProviderError::Database(DatabaseError::Other(e.to_string())))?;
        if !account_info.exists() {
            Ok(None)
        } else {
            Ok(Some(account_info))
        }
    }

    /// Get account code by its hash.
    fn code_by_hash_ref(&self, _code_hash: B256) -> Result<Bytecode, Self::Error> {
        unimplemented!()
    }

    /// Get storage value of address at index.
    fn storage_ref(&self, address: Address, index: U256) -> Result<U256, Self::Error> {
        let handle = tokio::runtime::Handle::try_current().map_err(|_| {
            ProviderError::Database(DatabaseError::Other("no tokio runtime found".to_string()))
        })?;
        let result =
            tokio::task::block_in_place(|| handle.block_on(self.fetch_storage_at(address, index)));
        let value =
            result.map_err(|e| ProviderError::Database(DatabaseError::Other(e.to_string())))?;
        Ok(value)
    }

    /// Get block hash by its number.
    fn block_hash_ref(&self, number: u64) -> Result<B256, Self::Error> {
        // Morph EVM opcode difference:
        // `BLOCKHASH(n)` returns `keccak(chain_id || n)` for the last 256 blocks.
        // Keep Ethereum semantics for out-of-range queries (return zero).
        let current_block_number = self.block_number + 1;
        if number >= current_block_number {
            return Ok(B256::ZERO);
        }

        // NOTE: `BLOCKHASH` is only defined for the last 256 blocks.
        // I.e. for `n < current_number` and `current_number - n <= 256`.
        if current_block_number.saturating_sub(number) > 256 {
            return Ok(B256::ZERO);
        }

        let mut buf = [0u8; 16];
        buf[..8].copy_from_slice(&self.chain_id.to_be_bytes());
        buf[8..].copy_from_slice(&number.to_be_bytes());
        Ok(alloy_primitives::utils::keccak256(buf))
    }
}

#[async_trait]
pub trait RpcDb<N: Network>: DatabaseRef {
    async fn state(&self, bundle_state: &BundleState) -> Result<EthereumState, RpcDbError>;

    /// Gets all account bytecodes.
    fn bytecodes(&self) -> Vec<Bytecode>;

    // Fetches the parent headers needed to constrain the BLOCKHASH opcode.
    async fn ancestor_headers(&self) -> Result<Vec<Header>, RpcDbError>;
}

#[async_trait]
impl<P, N> RpcDb<N> for BasicRpcDb<P, N>
where
    P: Provider<N> + Clone,
    N: Network,
{
    /// Assemble the EthereumState from the current database state.
    async fn state(&self, bundle_state: &BundleState) -> Result<EthereumState, RpcDbError> {
        let state_requests = self.get_state_requests();

        // For every account we touched, fetch the storage proofs for all the slots we touched.
        tracing::info!("fetching storage proofs");
        let mut before_storage_proofs = Vec::new();
        let mut after_storage_proofs = Vec::new();

        for (address, used_keys) in state_requests.iter() {
            let modified_keys = bundle_state
                .state
                .get(address)
                .map(|account| {
                    account.storage.keys().map(|key| B256::from(*key)).collect::<BTreeSet<_>>()
                })
                .unwrap_or_default()
                .into_iter()
                .collect::<Vec<_>>();

            let keys = used_keys
                .iter()
                .map(|key| B256::from(*key))
                .chain(modified_keys.clone().into_iter())
                .collect::<BTreeSet<_>>()
                .into_iter()
                .collect::<Vec<_>>();

            let storage_proof =
                self.eth_get_proof(*address, keys.clone(), self.block_number).await?;
            before_storage_proofs.push(eip1186_proof_to_account_proof(storage_proof));

            let storage_proof =
                self.eth_get_proof(*address, modified_keys, self.block_number + 1).await?;
            after_storage_proofs.push(eip1186_proof_to_account_proof(storage_proof));
        }

        let state = EthereumState::from_transition_proofs(
            self.state_root,
            &before_storage_proofs.iter().map(|item| (item.address, item.clone())).collect(),
            &after_storage_proofs.iter().map(|item| (item.address, item.clone())).collect(),
        )?;

        Ok(state)
    }

    fn bytecodes(&self) -> Vec<Bytecode> {
        let accounts = self.accounts.read().unwrap();

        accounts
            .values()
            .flat_map(|account| account.code.clone())
            .map(|code| (code.hash_slow(), code))
            .collect::<BTreeMap<_, _>>()
            .into_values()
            .collect::<Vec<_>>()
    }

    async fn ancestor_headers(&self) -> Result<Vec<Header>, RpcDbError> {
        let oldest_ancestor = *self.oldest_ancestor.read().unwrap();
        let mut ancestor_headers = vec![];
        for height in (oldest_ancestor..=(self.block_number)).rev() {
            let block = self
                .provider
                .get_block_by_number(height.into())
                .await?
                .ok_or(RpcDbError::BlockNotFound(height))?;

            ancestor_headers.push(Header {
                parent_hash: block.header().parent_hash(),
                ommers_hash: block.header().ommers_hash(),
                beneficiary: block.header().beneficiary(),
                state_root: block.header().state_root(),
                transactions_root: block.header().transactions_root(),
                receipts_root: block.header().receipts_root(),
                logs_bloom: block.header().logs_bloom(),
                difficulty: block.header().difficulty(),
                number: block.header().number(),
                gas_limit: block.header().gas_limit(),
                gas_used: block.header().gas_used(),
                timestamp: block.header().timestamp(),
                extra_data: block.header().extra_data().clone(),
                mix_hash: block.header().mix_hash().unwrap_or_default(),
                nonce: block.header().nonce().unwrap_or_default(),
                base_fee_per_gas: block.header().base_fee_per_gas(),
                withdrawals_root: block.header().withdrawals_root(),
                blob_gas_used: block.header().blob_gas_used(),
                excess_blob_gas: block.header().excess_blob_gas(),
                parent_beacon_block_root: block.header().parent_beacon_block_root(),
                requests_hash: block.header().requests_hash(),
            });
        }

        Ok(ancestor_headers)
    }
}
