use std::{
    collections::VecDeque,
    fs::File,
    io::Write,
    path::{Path, PathBuf},
    sync::Arc,
    time::{Duration, Instant},
};

use alloy_provider::{DynProvider, Provider, ProviderBuilder};
use anyhow::Context;
use morph_prove::{
    BatchProver, DefaultClient, PreparedBatch,
    evm::EvmProofFixture,
    execute::{InputSource, execute_batch},
};
use prover_executor_client::EvmVerificationError;
use serde::{Deserialize, Serialize};
use tokio::sync::Mutex;

use crate::{PROVE_RESULT, PROVE_TIME, PROVER_L2_RPC, PROVER_PROOF_DIR, PROVER_USE_RPC_DB};

// proveRequest
#[derive(Serialize, Deserialize, Debug)]
pub struct ProveRequest {
    pub batch_index: u64,
    pub start_block: u64,
    pub end_block: u64,
    pub rpc: String,
    pub shadow: Option<bool>,
    #[serde(default = "default_batch_version")]
    pub batch_version: u8,
}

fn default_batch_version() -> u8 {
    2
}

/// The prover that processes prove requests from a queue.
pub struct Prover {
    pub prove_queue: Arc<Mutex<VecDeque<ProveRequest>>>,
    batch_prover: BatchProver<DefaultClient>,
    provider: DynProvider,
}

/// Implementation of the Prover.
impl Prover {
    // Create a new Prover instance.
    pub async fn new(
        prove_queue: Arc<Mutex<VecDeque<ProveRequest>>>,
    ) -> Result<Self, anyhow::Error> {
        let rpc_url = PROVER_L2_RPC.parse()?;
        let provider = ProviderBuilder::new().connect_http(rpc_url).erased();
        let batch_prover = BatchProver::new().await?;

        Ok(Self { prove_queue, batch_prover, provider })
    }

    /// Processes prove requests from a queue.
    pub async fn prove_for_queue(&mut self) {
        log::info!("Waiting for prove request");
        loop {
            tokio::time::sleep(Duration::from_millis(12000)).await;

            // Step1. Get request from queue
            let (batch_index, start_block, end_block, shadow) = match self
                .prove_queue
                .lock()
                .await
                .pop_front()
            {
                Some(req) => {
                    log::info!(
                        "received prove request, batch index = {:#?}, blocks len = {:#?}, start_block = {:#?}, shadow = {:#?}",
                        req.batch_index,
                        req.end_block - req.start_block + 1,
                        req.start_block,
                        req.shadow,
                    );
                    (
                        req.batch_index,
                        req.start_block,
                        req.end_block,
                        req.shadow.unwrap_or_default(),
                    )
                }
                None => {
                    log::info!("no prove request");
                    continue;
                }
            };

            // Step2. Prepare and verify the batch
            let prepared = match prepare_batch_for_proving(
                batch_index,
                start_block,
                end_block,
                &self.provider,
                !shadow,
            )
            .await
            {
                Ok(input) => input,
                Err(e) => {
                    log::error!("Prepare input error for batch-{:?}, error: {:?}", batch_index, e);
                    PROVE_RESULT.set(2);
                    continue;
                }
            };

            // Step3. Generate evm proof
            log::info!("Generate evm proof");
            let start = Instant::now();
            let prove_rt = self.batch_prover.prove_prepared(prepared).await;

            match prove_rt {
                Ok(Some(proof)) => match save_proof(batch_index, proof) {
                    Ok(()) => {
                        PROVE_RESULT.set(1);
                        let duration_mins = start.elapsed().as_secs() / 60;
                        PROVE_TIME.set(duration_mins.try_into().unwrap_or_default());
                    }
                    Err(e) => {
                        PROVE_RESULT.set(2);
                        log::error!("Save evm proof of batch-{:?} error: {:?}", batch_index, e);
                    }
                },
                Ok(None) if shadow => {
                    PROVE_RESULT.set(1);
                    log::info!(
                        "Shadow execution of batch-{:?} completed successfully",
                        batch_index
                    );
                }
                Ok(None) => {
                    PROVE_RESULT.set(2);
                    log::error!("Gen proof of batch-{:?} is none", batch_index);
                }
                Err(e) => {
                    PROVE_RESULT.set(2);
                    log::error!("Gen proof of batch-{:?} error: {:?}", batch_index, e)
                }
            }
        }
    }
}

/// Generate and verify the prover input before publishing the batch header.
async fn prepare_batch_for_proving(
    batch_index: u64,
    start_block: u64,
    end_block: u64,
    provider: &DynProvider,
    prove: bool,
) -> Result<PreparedBatch, anyhow::Error> {
    // Step1. Get ExecutorInput
    let (input_source, fallback_source) = if *PROVER_USE_RPC_DB {
        (InputSource::Basic, InputSource::ExecutionWitness)
    } else {
        (InputSource::ExecutionWitness, InputSource::Basic)
    };
    log::info!("Prover input source: {:?}", input_source);
    let executor_input =
        match execute_batch(batch_index, start_block, end_block, provider, input_source).await {
            Ok(input) => input,
            Err(primary_error) => {
                log::warn!(
                    "Failed to generate ExecutorInput with {:?}: {:?}; retrying with {:?}",
                    input_source,
                    primary_error,
                    fallback_source,
                );
                execute_batch(batch_index, start_block, end_block, provider, fallback_source)
                    .await
                    .map_err(|fallback_error| {
                        anyhow::anyhow!(
                            "failed to generate ExecutorInput with both {input_source:?} \
                             ({primary_error:?}) and {fallback_source:?} ({fallback_error:?})"
                        )
                    })?
            }
        };
    let proof_dir =
        PathBuf::from(PROVER_PROOF_DIR.to_string()).join(format!("batch_{batch_index}"));
    std::fs::create_dir_all(&proof_dir)
        .with_context(|| format!("failed to create proof directory {}", proof_dir.display()))?;

    let error_path = proof_dir.join("execute_result.json");
    let prepared = match BatchProver::prepare_batch(executor_input, prove) {
        Ok(prepared) => prepared,
        Err(err) => {
            if let Some(evm_err) = err.downcast_ref::<EvmVerificationError>() {
                let error_data = serde_json::json!({
                    "error_code": "EVM_EXECUTE_NOT_EXPECTED",
                    "error_msg": evm_err.to_string()
                });
                let error_json = serde_json::to_vec_pretty(&error_data)?;
                write_atomically(&error_path, &error_json).with_context(|| {
                    format!(
                        "failed to record EVM verification error for batch {batch_index}: {evm_err}"
                    )
                })?;
            }
            return Err(err);
        }
    };

    write_atomically(&proof_dir.join("batch_header.data"), prepared.batch_header())?;

    match std::fs::remove_file(&error_path) {
        Ok(()) => log::info!("Removed stale error result for batch-{batch_index:?}"),
        Err(err) if err.kind() == std::io::ErrorKind::NotFound => {}
        Err(err) => {
            return Err(err).with_context(|| {
                format!("failed to remove stale error result {}", error_path.display())
            });
        }
    }

    Ok(prepared)
}

/// Save evm proof to file.
fn save_proof(batch_index: u64, proof: EvmProofFixture) -> Result<(), anyhow::Error> {
    let batch_dir =
        PathBuf::from(PROVER_PROOF_DIR.to_string()).join(format!("batch_{batch_index}"));
    std::fs::create_dir_all(&batch_dir)
        .with_context(|| format!("failed to create proof directory {}", batch_dir.display()))?;
    let proof_json = serde_json::to_vec_pretty(&proof)?;
    write_atomically(&batch_dir.join("plonk_proof.json"), &proof_json)?;
    log::info!("Successfully save evm proof of batch-{:?}", batch_index);
    Ok(())
}

fn write_atomically(path: &Path, contents: &[u8]) -> Result<(), anyhow::Error> {
    let file_name = path.file_name().context("atomic write target must have a file name")?;
    let temp_path = path.with_file_name(format!(".{}.tmp", file_name.to_string_lossy()));

    let write_result = (|| {
        let mut file = File::create(&temp_path)
            .with_context(|| format!("failed to create temporary file {}", temp_path.display()))?;
        file.write_all(contents)
            .with_context(|| format!("failed to write temporary file {}", temp_path.display()))?;
        file.sync_all()
            .with_context(|| format!("failed to sync temporary file {}", temp_path.display()))?;
        drop(file);
        std::fs::rename(&temp_path, path).with_context(|| {
            format!(
                "failed to replace {} with temporary file {}",
                path.display(),
                temp_path.display()
            )
        })?;
        Ok(())
    })();

    if write_result.is_err() {
        match std::fs::remove_file(&temp_path) {
            Ok(()) => {}
            Err(err) if err.kind() == std::io::ErrorKind::NotFound => {}
            Err(err) => {
                log::warn!("Failed to clean up temporary file {}: {}", temp_path.display(), err)
            }
        }
    }

    write_result
}
