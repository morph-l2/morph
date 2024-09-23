use crate::{
    database::EvmDatabase, error::DatabaseError, error::VerificationError, HardforkConfig,
};
use revm::db::AccountState;
use revm::primitives::{BlockEnv, TxEnv, U256};
use revm::Database;
use revm::{
    db::CacheDB,
    primitives::{AccountInfo, Env, SpecId, B256, KECCAK_EMPTY, POSEIDON_EMPTY},
};
use sbv_primitives::Address;
use sbv_primitives::{
    zk_trie::{
        db::KVDatabase,
        hash::{
            key_hasher::NoCacheHasher,
            poseidon::{Poseidon, PoseidonError},
        },
        trie::{ZkTrie, ZkTrieError},
    },
    Block, Transaction, TxTrace,
};
use sbv_utils::{cycle_tracker_start, dev_trace, measure_duration_micros};
use std::fmt::Debug;

mod builder;
pub use builder::EvmExecutorBuilder;
use sbv_primitives::zk_trie::scroll_types::Account;

/// Execute hooks
pub mod hooks;

/// EVM executor that handles the block.
pub struct EvmExecutor<'a, CodeDb, ZkDb> {
    hardfork_config: HardforkConfig,
    db: CacheDB<EvmDatabase<CodeDb, ZkDb>>,
    spec_id: SpecId,
    hooks: hooks::ExecuteHooks<'a, CodeDb, ZkDb>,
}

impl<CodeDb: KVDatabase, ZkDb: KVDatabase + Clone + 'static> EvmExecutor<'_, CodeDb, ZkDb> {
    /// Get reference to the DB
    pub fn db(&self) -> &CacheDB<EvmDatabase<CodeDb, ZkDb>> {
        &self.db
    }

    /// Get slot
    pub fn get_storage_value(&mut self, address: Address, index: U256) -> U256 {
        self.db.storage(address, index).expect("load storage error")
    }

    /// Update the DB
    pub fn update_db<T: Block>(&mut self, l2_trace: &T) -> Result<(), DatabaseError> {
        self.db.db.invalidate_storage_root_caches(
            self.db
                .accounts
                .iter()
                .map(|(addr, acc)| (*addr, acc.account_state.clone())),
        );

        self.db.db.update(l2_trace)
    }

    /// Handle a block.
    pub fn handle_block<T: Block>(&mut self, l2_trace: &T) -> Result<(), VerificationError> {
        measure_duration_millis!(
            handle_block_duration_milliseconds,
            self.handle_block_inner(l2_trace)
        )?;

        #[cfg(feature = "metrics")]
        sbv_utils::metrics::REGISTRY.block_counter.inc();

        Ok(())
    }

    #[inline(always)]
    fn handle_block_inner<T: Block>(&mut self, l2_trace: &T) -> Result<(), VerificationError> {
        self.hardfork_config
            .migrate(l2_trace.number(), &mut self.db)
            .unwrap();

        dev_debug!("handle block {:?}", l2_trace.number());
        let mut env = Box::<Env>::default();
        env.cfg.chain_id = l2_trace.chain_id();
        env.block = BlockEnv {
            number: U256::from_limbs([l2_trace.number(), 0, 0, 0]),
            coinbase: l2_trace.coinbase(),
            timestamp: l2_trace.timestamp(),
            gas_limit: l2_trace.gas_limit(),
            basefee: l2_trace.base_fee_per_gas().unwrap_or_default(),
            difficulty: l2_trace.difficulty(),
            prevrandao: l2_trace.prevrandao(),
            blob_excess_gas_and_price: None,
        };

        for (idx, tx) in l2_trace.transactions().enumerate() {
            cycle_tracker_start!("handle tx {}", idx);

            dev_trace!("handle {idx}th tx");

            let tx = tx
                .try_build_typed_tx()
                .map_err(|e| VerificationError::InvalidSignature {
                    tx_hash: tx.tx_hash(),
                    source: e,
                })?;

            dev_trace!("{tx:#?}");
            let mut env = env.clone();
            env.tx = TxEnv {
                caller: tx.get_or_recover_signer().map_err(|e| {
                    VerificationError::InvalidSignature {
                        tx_hash: *tx.tx_hash(),
                        source: e,
                    }
                })?,
                gas_limit: tx.gas_limit() as u64,
                gas_price: tx
                    .effective_gas_price(l2_trace.base_fee_per_gas().unwrap_or_default().to())
                    .map(U256::from)
                    .ok_or_else(|| VerificationError::InvalidGasPrice {
                        tx_hash: *tx.tx_hash(),
                    })?,
                transact_to: tx.to(),
                value: tx.value(),
                data: tx.data(),
                nonce: if !tx.is_l1_msg() {
                    Some(tx.nonce())
                } else {
                    None
                },
                chain_id: tx.chain_id(),
                access_list: tx.access_list().cloned().unwrap_or_default().0,
                gas_priority_fee: tx.max_priority_fee_per_gas().map(U256::from),
                ..Default::default()
            };

            if tx.is_l1_msg() {
                env.cfg.disable_base_fee = true; // disable base fee for l1 msg
            }
            env.tx.morph.is_l1_msg = tx.is_l1_msg();
            let rlp_bytes = tx.rlp();
            self.hooks.tx_rlp(self, &rlp_bytes);
            env.tx.morph.rlp_bytes = Some(rlp_bytes);

            dev_trace!("{env:#?}");
            {
                let mut revm = cycle_track!(
                    revm::Evm::builder()
                        .with_spec_id(self.spec_id)
                        .with_db(&mut self.db)
                        .with_env(env)
                        // .with_external_context(CustomPrintTracer::default())
                        // .append_handler_register(inspector_handle_register)
                        .build(),
                    "build Evm"
                );

                dev_trace!("handler cfg: {:?}", revm.handler.cfg);

                let _result = measure_duration_millis!(
                    transact_commit_duration_milliseconds,
                    cycle_track!(revm.transact_commit(), "transact_commit").map_err(|e| {
                        VerificationError::EvmExecution {
                            tx_hash: *tx.tx_hash(),
                            source: e,
                        }
                    })?
                );

                dev_trace!("{_result:#?}");
            }
            self.hooks.post_tx_execution(self, idx);

            dev_debug!("handle {idx}th tx done");
            cycle_tracker_end!("handle tx {}", idx);
        }
        Ok(())
    }

    /// Commit pending changes in cache db to zktrie
    pub fn commit_changes(&mut self, zktrie_db: ZkDb) -> Result<B256, DatabaseError> {
        measure_duration_millis!(
            commit_changes_duration_milliseconds,
            cycle_track!(
                self.commit_changes_inner(zktrie_db)
                    .map_err(DatabaseError::zk_trie),
                "commit_changes"
            )
        )
    }

    fn commit_changes_inner(
        &mut self,
        zktrie_db: ZkDb,
    ) -> Result<B256, ZkTrieError<PoseidonError, ZkDb::Error>> {
        let mut zktrie = ZkTrie::<Poseidon, ZkDb>::new_with_root(
            zktrie_db.clone(),
            NoCacheHasher,
            self.db.db.committed_zktrie_root(),
        )
        .expect("infallible");

        #[cfg(any(feature = "debug-account", feature = "debug-storage"))]
        let mut debug_recorder = sbv_utils::DebugRecorder::new();

        for (addr, db_acc) in self.db.accounts.iter() {
            // If EVM didn't touch the account, we don't need to update it
            if db_acc.account_state == AccountState::None {
                continue;
            }
            let Some(mut info): Option<AccountInfo> = db_acc.info() else {
                continue;
            };
            if info.is_empty() {
                continue;
            }

            dev_trace!("committing {addr}, {:?} {db_acc:?}", db_acc.account_state);
            cycle_tracker_start!("commit account {}", addr);

            let mut storage_root = self.db.db.prev_storage_root(addr);

            if !db_acc.storage.is_empty() {
                // get current storage root
                let storage_root_before = storage_root;
                // get storage tire
                cycle_tracker_start!("update storage_tire");
                let mut storage_trie = ZkTrie::<Poseidon, ZkDb>::new_with_root(
                    zktrie_db.clone(),
                    NoCacheHasher,
                    storage_root_before,
                )
                .expect("unable to get storage trie");
                for (key, value) in db_acc.storage.iter() {
                    if !value.is_zero() {
                        measure_duration_micros!(
                            zktrie_update_duration_microseconds,
                            cycle_track!(
                                storage_trie.update(key.to_be_bytes::<32>(), value)?,
                                "Zktrie::update_store"
                            )
                        );
                    } else {
                        measure_duration_micros!(
                            zktrie_delete_duration_microseconds,
                            cycle_track!(
                                storage_trie.delete(key.to_be_bytes::<32>())?,
                                "Zktrie::delete"
                            )
                        );
                    }

                    #[cfg(feature = "debug-storage")]
                    debug_recorder.record_storage(*addr, *key, *value);
                }

                measure_duration_micros!(
                    zktrie_commit_duration_microseconds,
                    storage_trie.commit()?
                );

                cycle_tracker_end!("update storage_tire");
                storage_root = *storage_trie.root().unwrap_ref();

                #[cfg(feature = "debug-storage")]
                debug_recorder.record_storage_root(*addr, storage_root);

                self.db.db.set_prev_storage_root(*addr, storage_root);
            }
            if !info.is_empty() {
                // if account not exist, all fields will be zero.
                // but if account exist, code_hash will be empty hash if code is empty
                if info.is_empty_code_hash() {
                    info.code_hash = KECCAK_EMPTY.0.into();
                    info.poseidon_code_hash = POSEIDON_EMPTY.0.into();
                } else {
                    assert_ne!(
                        info.poseidon_code_hash,
                        B256::ZERO,
                        "revm didn't update poseidon_code_hash, revm: {info:?}",
                    );
                }
            } else {
                info.code_hash = B256::ZERO;
                info.poseidon_code_hash = B256::ZERO;
            }

            #[cfg(feature = "debug-account")]
            debug_recorder.record_account(*addr, info.clone(), storage_root);

            let acc_data = Account::from_revm_account_with_storage_root(info, storage_root);
            measure_duration_micros!(
                zktrie_update_duration_microseconds,
                cycle_track!(
                    zktrie
                        .update(addr, acc_data)
                        .expect("failed to update account"),
                    "Zktrie::update_account"
                )
            );

            cycle_tracker_end!("commit account {}", addr);
        }

        measure_duration_micros!(zktrie_commit_duration_microseconds, zktrie.commit()?);

        let root_after = *zktrie.root().unwrap_ref();

        self.db.db.updated_committed_zktrie_root(root_after);

        Ok(B256::from(root_after))
    }
}

impl<CodeDb, ZkDb> Debug for EvmExecutor<'_, CodeDb, ZkDb> {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("EvmExecutor")
            .field("db", &self.db)
            .field("spec_id", &self.spec_id)
            .finish()
    }
}
