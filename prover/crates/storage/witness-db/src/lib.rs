use alloy_consensus::TrieAccount;
use alloy_primitives::map::HashMap;
use anyhow::anyhow;
use prover_mpt::EthereumState;
use revm::primitives::{keccak256, Address, B256, U256};
use revm::{
    state::{AccountInfo, Bytecode},
    DatabaseRef,
};

/// A read-only `revm::DatabaseRef` backed by an [`rsp_mpt::EthereumState`].
///
/// This is mainly used to execute EVM transactions against a witness-derived MPT state.
#[derive(Debug)]
pub struct TrieDB<'a> {
    inner: &'a EthereumState,
    bytecode_by_hash: HashMap<B256, &'a Bytecode>,
    chain_id: u64,
    block_number: u64,
}

impl<'a> TrieDB<'a> {
    /// Create a new [`TrieDB`].
    ///
    /// - `inner`: MPT-backed ethereum state.
    /// - `bytecode_by_hash`: bytecode lookup by code hash.
    pub fn new(
        inner: &'a EthereumState,
        bytecode_by_hash: HashMap<B256, &'a Bytecode>,
        chain_id: u64,
        block_number: u64,
    ) -> Self {
        Self { inner, bytecode_by_hash, chain_id, block_number }
    }

    /// Convenience helper to read a storage slot value.
    pub fn get_storage_value(&self, address: Address, index: U256) -> Result<U256, anyhow::Error> {
        self.storage_ref(address, index).map_err(|e| anyhow!("storage_ref error: {e:?}"))
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
                eprintln!(
                    "get account of {:?}, hashed_address: {:?} from trie error: {:?}",
                    address,
                    alloy_primitives::hex::encode_prefixed(hashed_address),
                    e
                );
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
        Ok(self.bytecode_by_hash.get(&hash).map(|code| (*code).clone()).unwrap())
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
        // Morph EVM opcode difference:
        // `BLOCKHASH(n)` returns `keccak(chain_id || n)` for the last 256 blocks.
        // Keep Ethereum semantics for out-of-range queries (return zero).
        if number > self.block_number {
            return Ok(B256::ZERO);
        }

        // NOTE: `BLOCKHASH` is only defined for the last 256 blocks.
        // I.e. for `n < current_number` and `current_number - n <= 256`.
        if self.block_number.saturating_sub(number) >= 256 {
            return Ok(B256::ZERO);
        }

        let mut buf = [0u8; 16];
        buf[..8].copy_from_slice(&self.chain_id.to_be_bytes());
        buf[8..].copy_from_slice(&number.to_be_bytes());
        Ok(alloy_primitives::utils::keccak256(buf))
    }
}
