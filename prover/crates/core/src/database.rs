use crate::error::DatabaseError;
use once_cell::sync::Lazy;
use revm::{
    db::{AccountState, DatabaseRef},
    primitives::{AccountInfo, Address, Bytecode, B256, U256},
};
use sbv_primitives::{
    zk_trie::{
        db::{KVDatabase, KVDatabaseItem},
        hash::{key_hasher::NoCacheHasher, poseidon::Poseidon},
        scroll_types::Account,
        trie::ZkTrie,
    },
    Block,
};
use std::{cell::RefCell, collections::HashMap, fmt};

type Result<T, E = DatabaseError> = std::result::Result<T, E>;

type StorageTrieLazyFn<Db> = Box<dyn FnOnce() -> ZkTrie<Poseidon, Db>>;
type LazyStorageTrie<Db> = Lazy<ZkTrie<Poseidon, Db>, StorageTrieLazyFn<Db>>;

/// A database that consists of account and storage information.
pub struct EvmDatabase<CodeDb, ZkDb> {
    /// Map of code hash to bytecode.
    code_db: CodeDb,
    /// The initial storage roots of accounts, used for after commit.
    /// Need to be updated after zkTrie commit.
    prev_storage_roots: RefCell<HashMap<Address, B256>>,
    /// Storage trie cache, avoid re-creating trie for the same account.
    /// Need to invalidate before `update`, otherwise the trie root may be outdated.
    storage_trie_refs: RefCell<HashMap<Address, LazyStorageTrie<ZkDb>>>,
    /// Current uncommitted zkTrie root based on the block trace.
    committed_zktrie_root: B256,
    /// The underlying zkTrie database.
    zktrie_db: ZkDb,
    /// Current view of zkTrie database.
    zktrie: ZkTrie<Poseidon, ZkDb>,
}

impl<CodeDb, Db> fmt::Debug for EvmDatabase<CodeDb, Db> {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        f.debug_struct("EvmDatabase")
            .field("committed_zktrie_root", &self.committed_zktrie_root)
            .finish()
    }
}

impl<CodeDb: KVDatabase, ZkDb: KVDatabase + Clone + 'static> EvmDatabase<CodeDb, ZkDb> {
    /// Initialize an EVM database from a block trace.
    pub fn new<T: Block>(l2_trace: T, mut code_db: CodeDb, zktrie_db: ZkDb) -> Result<Self> {
        cycle_tracker_start!("insert CodeDB");
        for code in l2_trace.codes() {
            let hash = revm::primitives::keccak256(code);
            code_db
                .or_put(hash.as_slice(), code)
                .map_err(DatabaseError::code_db)?;
        }
        cycle_tracker_end!("insert CodeDB");

        let committed_zktrie_root = l2_trace.root_before();

        let zktrie = ZkTrie::new_with_root(zktrie_db.clone(), NoCacheHasher, committed_zktrie_root)
            .map_err(DatabaseError::zk_trie)?;

        Ok(EvmDatabase {
            code_db,
            prev_storage_roots: Default::default(),
            storage_trie_refs: Default::default(),
            committed_zktrie_root,
            zktrie_db,
            zktrie,
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

    /// Get the committed zkTrie root.
    #[inline]
    pub(crate) fn committed_zktrie_root(&self) -> B256 {
        self.committed_zktrie_root
    }

    /// Set the committed zkTrie root.
    #[inline]
    pub(crate) fn updated_committed_zktrie_root(&mut self, new_root: B256) {
        self.committed_zktrie_root = new_root;
    }

    /// Update the database with a new block trace.
    pub fn update<T: Block>(&mut self, l2_trace: T) -> Result<()> {
        measure_duration_millis!(update_db_duration_milliseconds, self.update_inner(l2_trace))
    }

    fn update_inner<T: Block>(&mut self, l2_trace: T) -> Result<()> {
        cycle_tracker_start!("insert CodeDB");
        for code in l2_trace.codes() {
            let hash = revm::primitives::keccak256(code);
            self.code_db
                .or_put(hash.as_slice(), code)
                .map_err(DatabaseError::code_db)?;
        }
        cycle_tracker_end!("insert CodeDB");

        self.zktrie = ZkTrie::new_with_root(
            self.zktrie_db.clone(),
            NoCacheHasher,
            l2_trace.root_before(),
        )
        .map_err(DatabaseError::zk_trie)?;

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

impl<CodeDb: KVDatabase, ZkDb: KVDatabase + Clone + 'static> DatabaseRef
    for EvmDatabase<CodeDb, ZkDb>
{
    type Error = DatabaseError;

    /// Get basic account information.
    fn basic_ref(&self, address: Address) -> Result<Option<AccountInfo>> {
        let Some(account) = measure_duration_micros!(
            zktrie_get_duration_microseconds,
            self.zktrie.get::<Account, _>(address)
        )
        .map_err(DatabaseError::zk_trie)?
        else {
            return Ok(None);
        };

        self.prev_storage_roots
            .borrow_mut()
            .entry(address)
            .or_insert(account.storage_root);
        let zktrie_db = self.zktrie_db.clone();
        self.storage_trie_refs
            .borrow_mut()
            .entry(address)
            .or_insert_with(|| {
                Lazy::new(Box::new(move || {
                    ZkTrie::new_with_root(zktrie_db.clone(), NoCacheHasher, account.storage_root)
                        .expect("storage trie associated with account not found")
                }))
            });

        let mut info = AccountInfo::from(account);
        info.code = self
            .code_db
            .get(&account.code_hash)
            .map_err(DatabaseError::code_db)?
            .map(|v| Bytecode::new_legacy(v.into_bytes().into()));

        Ok(Some(info))
    }

    /// Get account code by its code hash.
    fn code_by_hash_ref(&self, hash: B256) -> Result<Bytecode> {
        // Sometimes the code in previous account info is not contained,
        // and the CacheDB has already loaded the previous account info,
        // then the upcoming trace contains code (meaning the code is used in this new block),
        // we can't directly update the CacheDB, so we offer the code by hash here.
        // However, if the code still cannot be found, this is an error.
        self.code_db
            .get(&hash)
            .map_err(DatabaseError::code_db)?
            .map(|v| Bytecode::new_legacy(v.into_bytes().into()))
            .ok_or_else(|| {
                unreachable!(
                    "Code is either loaded or not needed (like EXTCODESIZE), code hash: {:?}",
                    hash
                );
            })
    }

    /// Get storage value of address at index.
    fn storage_ref(&self, address: Address, index: U256) -> Result<U256> {
        dev_trace!("get storage of {:?} at index {:?}", address, index);
        let mut storage_trie_refs = self.storage_trie_refs.borrow_mut();
        let trie = storage_trie_refs
            .entry(address)
            .or_insert_with_key(|address| {
                let storage_root = measure_duration_micros!(
                    zktrie_get_duration_microseconds,
                    self.zktrie.get::<Account, _>(address)
                )
                .expect("unexpected zktrie error")
                .map(|acc| acc.storage_root)
                .unwrap_or_default();
                dev_debug!("storage root of {:?} is {:?}", address, storage_root);

                let zktrie_db = self.zktrie_db.clone();
                Lazy::new(Box::new(move || {
                    ZkTrie::new_with_root(zktrie_db.clone(), NoCacheHasher, storage_root)
                        .expect("storage trie associated with account not found")
                }))
            });

        #[cfg(debug_assertions)]
        {
            let current_root = trie.root().unwrap_ref();
            let expected_root = self
                .zktrie
                .get::<Account, _>(address)
                .expect("unexpected zktrie error")
                .map(|acc| acc.storage_root)
                .unwrap_or_default();
            assert_eq!(*current_root, expected_root);
        }

        Ok(measure_duration_micros!(
            zktrie_get_duration_microseconds,
            trie.get::<U256, _>(index.to_be_bytes::<32>())
        )
        .map_err(DatabaseError::zk_trie)?
        .unwrap_or_default())
    }

    /// Get block hash by block number.
    fn block_hash_ref(&self, _: u64) -> Result<B256> {
        unreachable!("BLOCKHASH is disabled")
    }
}
