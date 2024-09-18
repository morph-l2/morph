use crate::error::ZkTrieError;
use once_cell::sync::Lazy;
use revm::{
    db::{AccountState, DatabaseRef},
    primitives::{AccountInfo, Address, Bytecode, B256, U256},
};
use sbv_primitives::{
    init_hash_scheme,
    zk_trie::{SharedMemoryDb, ZkMemoryDb, ZkTrie},
    Block,
};
use std::rc::Rc;
use std::{cell::RefCell, collections::HashMap, convert::Infallible, fmt};

type Result<T, E = ZkTrieError> = std::result::Result<T, E>;

type StorageTrieLazyFn = Box<dyn FnOnce() -> ZkTrie<SharedMemoryDb>>;

/// A read-only in-memory database that consists of account and storage information.
pub struct ReadOnlyDB {
    /// In-memory map of code hash to bytecode.
    code_db: HashMap<B256, Bytecode>,
    /// The initial storage roots of accounts, used for after commit.
    /// Need to be updated after zkTrie commit.
    prev_storage_roots: RefCell<HashMap<Address, B256>>,
    /// Storage trie cache, avoid re-creating trie for the same account.
    /// Need to invalidate before `update`, otherwise the trie root may be outdated.
    storage_trie_refs: RefCell<HashMap<Address, Lazy<ZkTrie<SharedMemoryDb>, StorageTrieLazyFn>>>,
    /// Current uncommitted zkTrie root based on the block trace.
    committed_zktrie_root: B256,
    /// The underlying zkTrie database.
    zktrie_db: Rc<ZkMemoryDb>,
    /// Current view of zkTrie database.
    zktrie_db_ref: ZkTrie<SharedMemoryDb>,
}

impl fmt::Debug for ReadOnlyDB {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        f.debug_struct("ReadOnlyDB")
            .field("code_db", &self.code_db.len())
            .field("committed_zktrie_root", &self.committed_zktrie_root)
            .finish()
    }
}

impl ReadOnlyDB {
    /// Initialize an EVM database from a block trace.
    pub fn new<T: Block>(l2_trace: T, zktrie_db: &Rc<ZkMemoryDb>) -> Result<Self> {
        let size_hint = l2_trace.codes().len();
        Self::new_with_size_hint(l2_trace, zktrie_db, size_hint)
    }

    /// Initialize an EVM database from a block trace with size hint of code database.
    pub fn new_with_size_hint<T: Block>(
        l2_trace: T,
        zktrie_db: &Rc<ZkMemoryDb>,
        size_hint: usize,
    ) -> Result<Self> {
        init_hash_scheme();

        cycle_tracker_start!("insert CodeDB");
        let mut code_db = HashMap::with_capacity(size_hint);
        for code in l2_trace.codes() {
            let hash = revm::primitives::keccak256(code);
            code_db.entry(hash).or_insert_with(|| {
                dev_trace!("insert code {:?}", hash);
                Bytecode::new_raw(revm::primitives::Bytes::from(code.to_vec()))
            });
        }
        cycle_tracker_end!("insert CodeDB");

        let uncommitted_zktrie_root = l2_trace.root_before();

        Ok(ReadOnlyDB {
            code_db,
            prev_storage_roots: Default::default(),
            storage_trie_refs: Default::default(),
            committed_zktrie_root: uncommitted_zktrie_root,
            zktrie_db: zktrie_db.clone(),
            zktrie_db_ref: zktrie_db
                .new_ref_trie(&uncommitted_zktrie_root.0)
                .ok_or(ZkTrieError::ZkTrieRootNotFound)?,
        })
    }

    /// Set the previous storage root of an account.
    ///
    /// Should be updated after commit.
    #[inline]
    pub(crate) fn set_prev_storage_root(
        &self,
        address: Address,
        storage_root: B256,
    ) -> Option<B256> {
        self.prev_storage_roots
            .borrow_mut()
            .insert(address, storage_root)
    }

    /// Get the previous storage root of an account.
    #[inline]
    pub(crate) fn prev_storage_root(&self, address: &Address) -> B256 {
        self.prev_storage_roots
            .borrow()
            .get(address)
            .copied()
            .unwrap_or_default()
    }

    /// Get the zkTrie root.
    #[inline]
    pub(crate) fn committed_zktrie_root(&self) -> B256 {
        self.committed_zktrie_root
    }

    /// Get the zkTrie root.
    #[inline]
    pub(crate) fn updated_committed_zktrie_root(&mut self, new_root: B256) {
        self.committed_zktrie_root = new_root;
    }

    /// Update the database with a new block trace.
    pub fn update<T: Block>(&mut self, l2_trace: T) -> Result<()> {
        measure_duration_histogram!(update_db_duration_microseconds, self.update_inner(l2_trace))
    }

    fn update_inner<T: Block>(&mut self, l2_trace: T) -> Result<()> {
        cycle_tracker_start!("insert CodeDB");
        for code in l2_trace.codes() {
            let hash = revm::primitives::keccak256(code);
            self.code_db.entry(hash).or_insert_with(|| {
                dev_trace!("insert code {:?}", hash);
                Bytecode::new_raw(revm::primitives::Bytes::from(code.to_vec()))
            });
        }
        cycle_tracker_end!("insert CodeDB");

        self.zktrie_db_ref = self
            .zktrie_db
            .new_ref_trie(&l2_trace.root_before().0)
            .ok_or(ZkTrieError::ZkTrieRootNotFound)?;

        Ok(())
    }

    /// Invalidate internal cache for any account touched by EVM.
    pub(crate) fn invalidate_storage_root_caches(
        &mut self,
        account_states: impl Iterator<Item = (Address, AccountState)>,
    ) {
        let mut storage_trie_refs = self.storage_trie_refs.borrow_mut();
        for (address, account_state) in account_states {
            if account_state != AccountState::None {
                storage_trie_refs.remove(&address);
            }
        }
    }
}

impl DatabaseRef for ReadOnlyDB {
    type Error = Infallible;

    /// Get basic account information.
    fn basic_ref(&self, address: Address) -> Result<Option<AccountInfo>, Self::Error> {
        Ok(self
            .zktrie_db_ref
            .get_account(address.as_slice())
            .map(|account_data| {
                let code_size =
                    u64::from_be_bytes((&account_data[0][16..24]).try_into().unwrap()) as usize;
                let nonce = u64::from_be_bytes((&account_data[0][24..]).try_into().unwrap());
                let balance = U256::from_be_bytes(account_data[1]);
                let code_hash = B256::from(account_data[3]);
                let poseidon_code_hash = B256::from(account_data[4]);

                let storage_root = B256::from(account_data[2]);
                self.prev_storage_roots
                    .borrow_mut()
                    .entry(address)
                    .or_insert(storage_root.0.into());

                let zktrie_db = self.zktrie_db.clone();
                self.storage_trie_refs.borrow_mut().insert(
                    address,
                    Lazy::new(Box::new(move || {
                        zktrie_db
                            .new_ref_trie(&storage_root.0)
                            .expect("storage trie associated with account not found")
                    })),
                );
                AccountInfo {
                    balance,
                    nonce,
                    code_size,
                    code_hash,
                    poseidon_code_hash,
                    code: self.code_db.get(&code_hash).cloned(),
                }
            }))
    }

    /// Get account code by its code hash.
    fn code_by_hash_ref(&self, hash: B256) -> Result<Bytecode, Self::Error> {
        // Sometimes the code in previous account info is not contained,
        // and the CacheDB has already loaded the previous account info,
        // then the upcoming trace contains code (meaning the code is used in this new block),
        // we can't directly update the CacheDB, so we offer the code by hash here.
        // However, if the code still cannot be found, this is an error.
        self.code_db.get(&hash).cloned().ok_or_else(|| {
            unreachable!(
                "Code is either loaded or not needed (like EXTCODESIZE), code hash: {:?}",
                hash
            );
        })
    }

    /// Get storage value of address at index.
    fn storage_ref(&self, address: Address, index: U256) -> Result<U256, Self::Error> {
        let mut storage_trie_refs = self.storage_trie_refs.borrow_mut();
        let trie = storage_trie_refs
            .entry(address)
            .or_insert_with_key(|address| {
                let storage_root = self
                    .zktrie_db_ref
                    .get_account(address.as_slice())
                    .map(|account_data| B256::from(account_data[2]))
                    .unwrap_or_default();
                let zktrie_db = self.zktrie_db.clone();
                Lazy::new(Box::new(move || {
                    zktrie_db
                        .clone()
                        .new_ref_trie(&storage_root.0)
                        .expect("storage trie associated with account not found")
                }))
            });

        Ok(trie
            .get_store(&index.to_be_bytes::<32>())
            .map(|store_data| U256::from_be_bytes(store_data))
            .unwrap_or_default())
    }

    /// Get block hash by block number.
    fn block_hash_ref(&self, _: u64) -> Result<B256, Self::Error> {
        unreachable!("BLOCKHASH is disabled")
    }
}
