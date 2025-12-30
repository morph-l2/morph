use alloy_consensus::TrieAccount;
use alloy_primitives::map::HashMap;
use anyhow::anyhow;
use revm::primitives::{keccak256, Address, B256, U256};
use revm::{
    state::{AccountInfo, Bytecode},
    DatabaseRef,
};
use rsp_mpt::EthereumState;

/// A read-only `revm::DatabaseRef` backed by an [`rsp_mpt::EthereumState`].
///
/// This is mainly used to execute EVM transactions against a witness-derived MPT state.
#[derive(Debug)]
pub struct TrieDB<'a> {
    inner: &'a EthereumState,
    block_hashes: HashMap<u64, B256>,
    bytecode_by_hash: HashMap<B256, &'a Bytecode>,
}

impl<'a> TrieDB<'a> {
    /// Create a new [`TrieDB`].
    ///
    /// - `inner`: MPT-backed ethereum state.
    /// - `block_hashes`: optional block hash mapping (used by `BLOCKHASH`).
    /// - `bytecode_by_hash`: bytecode lookup by code hash.
    pub fn new(
        inner: &'a EthereumState,
        block_hashes: HashMap<u64, B256>,
        bytecode_by_hash: HashMap<B256, &'a Bytecode>,
    ) -> Self {
        Self { inner, block_hashes, bytecode_by_hash }
    }

    /// Convenience helper to read a storage slot value.
    pub fn get_storage_value(&self, address: Address, index: U256) -> Result<U256, anyhow::Error> {
        self.storage_ref(address, index).map_err(|e| anyhow!("storage_ref error: {:?}", e))
    }
}

impl DatabaseRef for TrieDB<'_> {
    /// The database error type.
    type Error = core::convert::Infallible;

    /// Get basic account information.
    fn basic_ref(&self, address: Address) -> Result<Option<AccountInfo>, Self::Error> {
        let hashed_address = keccak256(address);
        let hashed_address = hashed_address.as_slice();

        let account_in_trie = self
            .inner
            .state_trie
            .get_rlp::<TrieAccount>(hashed_address)
            .map_err(|e| {
                // keep behavior non-panicking (consistent with original code's debug printing)
                eprintln!("get account of {:?} from trie error: {:?}", address, e);
            })
            .unwrap();

        let account = account_in_trie.map(|account_in_trie| AccountInfo {
            balance: account_in_trie.balance,
            nonce: account_in_trie.nonce,
            code_hash: account_in_trie.code_hash,
            code: None,
        });

        Ok(account)
    }

    /// Get account code by its hash.
    fn code_by_hash_ref(&self, hash: B256) -> Result<Bytecode, Self::Error> {
        Ok(self
            .bytecode_by_hash
            .get(&hash)
            .map(|code| (*code).clone())
            .unwrap())
    }

    /// Get storage value of address at index.
    fn storage_ref(&self, address: Address, index: U256) -> Result<U256, Self::Error> {
        let hashed_address = keccak256(address);
        let hashed_address = hashed_address.as_slice();

        let storage_trie = self
            .inner
            .storage_tries
            .get(hashed_address)
            .expect("A storage trie must be provided for each account");

        Ok(storage_trie
            .get_rlp::<U256>(keccak256(index.to_be_bytes::<32>()).as_slice())
            .expect("Can get from MPT")
            .unwrap_or_default())
    }

    /// Get block hash by block number.
    fn block_hash_ref(&self, number: u64) -> Result<B256, Self::Error> {
        Ok(*self
            .block_hashes
            .get(&number)
            .expect("A block hash must be provided for each block number"))
    }
}
