use crate::{
    database::ReadOnlyDB,
    error::{VerificationError, ZkTrieError},
    HardforkConfig,
};
use revm::{
    db::{AccountState, CacheDB},
    primitives::{
        AccountInfo, BlockEnv, Env, SpecId, TxEnv, B256, KECCAK_EMPTY, POSEIDON_EMPTY, U256,
    },
};
use sbv_primitives::{zk_trie::ZkMemoryDb, Block, Transaction, TxTrace};
use std::{fmt::Debug, rc::Rc};

mod builder;
pub use builder::EvmExecutorBuilder;

/// Execute hooks
pub mod hooks;

/// EVM executor that handles the block.
pub struct EvmExecutor<'a> {
    hardfork_config: HardforkConfig,
    db: CacheDB<ReadOnlyDB>,
    spec_id: SpecId,
    hooks: hooks::ExecuteHooks<'a>,
}

impl EvmExecutor<'_> {
    /// Get reference to the DB
    pub fn db(&self) -> &CacheDB<ReadOnlyDB> {
        &self.db
    }

    /// Update the DB
    pub fn update_db<T: Block>(&mut self, l2_trace: &T) -> Result<(), ZkTrieError> {
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
        measure_duration_histogram!(
            handle_block_duration_microseconds,
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

                let _result =
                    cycle_track!(revm.transact_commit(), "transact_commit").map_err(|e| {
                        VerificationError::EvmExecution {
                            tx_hash: *tx.tx_hash(),
                            source: e,
                        }
                    })?;

                dev_trace!("{_result:#?}");
            }
            self.hooks.post_tx_execution(self, idx);

            dev_debug!("handle {idx}th tx done");
            cycle_tracker_end!("handle tx {}", idx);
        }
        Ok(())
    }

    /// Commit pending changes in cache db to zktrie
    pub fn commit_changes(&mut self, zktrie_db: &Rc<ZkMemoryDb>) -> B256 {
        measure_duration_histogram!(
            commit_changes_duration_microseconds,
            cycle_track!(self.commit_changes_inner(zktrie_db), "commit_changes")
        )
    }

    fn commit_changes_inner(&mut self, zktrie_db: &Rc<ZkMemoryDb>) -> B256 {
        let mut zktrie = zktrie_db
            .new_trie(&self.db.db.committed_zktrie_root())
            .expect("infallible");

        #[cfg(any(feature = "debug-account", feature = "debug-storage"))]
        let mut debug_recorder = sbv_utils::DebugRecorder::new();

        for (addr, db_acc) in self.db.accounts.iter() {
            // If EVM didn't touch the account, we don't need to update it
            if db_acc.account_state == AccountState::None {
                continue;
            }
            let Some(info): Option<AccountInfo> = db_acc.info() else {
                continue;
            };
            if info.is_empty() {
                continue;
            }

            dev_trace!("committing {addr}, {:?} {db_acc:?}", db_acc.account_state);
            cycle_tracker_start!("commit account {}", addr);

            let mut code_size = 0;
            let mut storage_root = self.db.db.prev_storage_root(addr);
            let mut code_hash = B256::ZERO;
            let mut poseidon_code_hash = B256::ZERO;

            if !db_acc.storage.is_empty() {
                // get current storage root
                let storage_root_before = storage_root;
                // get storage tire
                cycle_tracker_start!("update storage_tire");
                let mut storage_trie = zktrie_db
                    .new_trie(storage_root_before.as_ref())
                    .expect("unable to get storage trie");
                for (key, value) in db_acc.storage.iter() {
                    if !value.is_zero() {
                        cycle_track!(
                            storage_trie
                                .update_store(&key.to_be_bytes::<32>(), &value.to_be_bytes())
                                .expect("failed to update storage"),
                            "Zktrie::update_store"
                        );
                    } else {
                        cycle_track!(
                            storage_trie.delete(&key.to_be_bytes::<32>()),
                            "Zktrie::delete"
                        );
                    }

                    #[cfg(feature = "debug-storage")]
                    debug_recorder.record_storage(*addr, *key, *value);
                }

                if storage_trie.is_trie_dirty() {
                    storage_trie.prepare_root();
                }

                cycle_tracker_end!("update storage_tire");
                storage_root = storage_trie.root().into();

                #[cfg(feature = "debug-storage")]
                debug_recorder.record_storage_root(*addr, storage_root);

                self.db.db.set_prev_storage_root(*addr, storage_root);
            }
            if !info.is_empty() {
                // if account not exist, all fields will be zero.
                // but if account exist, code_hash will be empty hash if code is empty
                if info.is_empty_code_hash() {
                    code_hash = KECCAK_EMPTY.0.into();
                    poseidon_code_hash = POSEIDON_EMPTY.0.into();
                } else {
                    assert_ne!(
                        info.poseidon_code_hash,
                        B256::ZERO,
                        "revm didn't update poseidon_code_hash, revm: {info:?}",
                    );
                    code_size = info.code_size as u64;
                    code_hash = info.code_hash.0.into();
                    poseidon_code_hash = info.poseidon_code_hash.0.into();
                }
            }

            #[cfg(feature = "debug-account")]
            debug_recorder.record_account(
                *addr,
                info.nonce,
                info.balance,
                code_hash,
                poseidon_code_hash,
                code_size,
                storage_root,
            );

            let acc_data = [
                U256::from_limbs([info.nonce, code_size, 0, 0]).to_be_bytes(),
                info.balance.to_be_bytes(),
                storage_root.0,
                code_hash.0,
                poseidon_code_hash.0,
            ];
            cycle_track!(
                zktrie
                    .update_account(addr.as_slice(), &acc_data)
                    .expect("failed to update account"),
                "Zktrie::update_account"
            );

            cycle_tracker_end!("commit account {}", addr);
        }

        if zktrie.is_trie_dirty() {
            zktrie.prepare_root();
        }

        let root_after = zktrie.root();

        self.db.db.updated_committed_zktrie_root(root_after.into());

        B256::from(root_after)
    }
}

impl Debug for EvmExecutor<'_> {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("EvmExecutor")
            .field("db", &self.db)
            .field("spec_id", &self.spec_id)
            .finish()
    }
}
