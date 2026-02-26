use std::{
    fs::{self, File},
    io::{BufReader, BufWriter, Write},
    path::PathBuf,
    sync::Arc,
    time::{Duration, Instant},
};

use crate::{PROVER_L2_RPC, PROVER_PROOF_DIR, PROVER_USE_RPC_DB, PROVE_RESULT, PROVE_TIME};
use alloy_provider::{DynProvider, Provider, ProviderBuilder};
use morph_prove::{evm::EvmProofFixture, execute::execute_batch, prove};
use prover_executor_client::{types::input::ExecutorInput, BlobVerifier, EVMVerifier};

use prover_primitives::types::BlockTrace;
use serde::{Deserialize, Serialize};
use tokio::sync::Mutex;

// proveRequest
#[derive(Serialize, Deserialize, Debug)]
pub struct ProveRequest {
    pub batch_index: u64,
    pub start_block: u64,
    pub end_block: u64,
    pub rpc: String,
    pub shadow: Option<bool>,
}

/// The prover that processes prove requests from a queue.
pub struct Prover {
    pub prove_queue: Arc<Mutex<Vec<ProveRequest>>>,
    provider: DynProvider,
}

/// Implementation of the Prover.
impl Prover {
    // Create a new Prover instance.
    pub fn new(prove_queue: Arc<Mutex<Vec<ProveRequest>>>) -> Result<Self, anyhow::Error> {
        let rpc_url = PROVER_L2_RPC.parse()?;
        let provider = ProviderBuilder::new().connect_http(rpc_url).erased();

        Ok(Self { prove_queue, provider })
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
                .pop()
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

            // Step2. Generate ExecutorInput
            let mut input =
                match gen_client_input(batch_index, start_block, end_block, &self.provider).await {
                    Ok(input) => input,
                    Err(e) => {
                        log::error!(
                            "Generate ExecutorInput error for batch-{:?}, error: {:?}",
                            batch_index,
                            e
                        );
                        PROVE_RESULT.set(2);
                        continue;
                    }
                };

            // Step3. Generate evm proof
            log::info!("Generate evm proof");
            let start = Instant::now();
            let prove_rt = prove(&mut input, !shadow);

            match prove_rt {
                Ok(Some(proof)) => {
                    save_proof(batch_index, proof);
                    PROVE_RESULT.set(1);
                    let duration_mins = start.elapsed().as_secs() / 60;
                    PROVE_TIME.set(duration_mins.try_into().unwrap_or_default());
                }
                Ok(None) => {
                    PROVE_RESULT.set(2);
                    log::error!("Gen proof of batch-{:?} is none", batch_index)
                }
                Err(e) => {
                    PROVE_RESULT.set(2);
                    log::error!("Gen proof of batch-{:?} error: {:?}", batch_index, e)
                }
            }
        }
    }
}

/// Generate ExecutorInput for prover client.
async fn gen_client_input(
    batch_index: u64,
    start_block: u64,
    end_block: u64,
    provider: &DynProvider,
) -> Result<ExecutorInput, anyhow::Error> {
    // Step1. Get ExecutorInput
    let executor_input =
        execute_batch(batch_index, start_block, end_block, provider, *PROVER_USE_RPC_DB).await?;
    let proof_dir =
        PathBuf::from(PROVER_PROOF_DIR.to_string()).join(format!("batch_{batch_index}"));
    std::fs::create_dir_all(&proof_dir).expect("failed to create proof path");

    // Step2. Get BatchInfo by EVM Verify.
    let verify_result = EVMVerifier::verify(executor_input.block_inputs.clone());

    // Step3. Save batch header or error info.
    if let Ok(batch_info) = verify_result {
        let (versioned_hash, _) =
            BlobVerifier::verify(&executor_input.blob_info, executor_input.block_inputs.len())?;
        // Save batch_header
        // | batch_data_hash | versioned_hash | sequencer_root |
        // |-----------------|----------------|----------------|
        // |     bytes32     |     bytes32    |     bytes32    |
        let mut batch_header: Vec<u8> = Vec::with_capacity(96);
        batch_header.extend_from_slice(&batch_info.data_hash().0);
        batch_header.extend_from_slice(&versioned_hash.0);
        batch_header.extend_from_slice(&batch_info.sequencer_root().0);
        batch_header.extend_from_slice(&batch_info.sequencer_root().0);
        let mut batch_file = File::create(proof_dir.join("batch_header.data"))?;
        batch_file.write_all(&batch_header[..]).expect("failed to batch_header");
    } else {
        let err = verify_result.unwrap_err();
        let error_data = serde_json::json!({
            "error_code": "EVM_EXECUTE_NOT_EXPECTED",
            "error_msg": err.to_string()
        });
        let mut batch_file = File::create(proof_dir.join("execute_result.json"))?;
        batch_file
            .write_all(serde_json::to_string_pretty(&error_data)?.as_bytes())
            .expect("failed to write error");
        log::error!("EVM verification failed for batch {}: {}", batch_index, err);
    }
    Ok(executor_input)
}

/// Save evm proof to file.
fn save_proof(batch_index: u64, proof: EvmProofFixture) {
    let batch_dir =
        PathBuf::from(PROVER_PROOF_DIR.to_string()).join(format!("batch_{batch_index}"));
    std::fs::create_dir_all(&batch_dir).expect("failed to create proof path");
    std::fs::write(
        batch_dir.join("plonk_proof.json"),
        serde_json::to_string_pretty(&proof).unwrap(),
    )
    .expect("failed to write proof");
    log::info!("Successfully save evm proof of batch-{:?}", batch_index);
}

#[allow(dead_code)]
fn load_trace(file_path: &str) -> Vec<Vec<BlockTrace>> {
    let file = File::open(file_path).unwrap();
    let reader = BufReader::new(file);
    serde_json::from_reader(reader).unwrap()
}

#[allow(dead_code)]
fn save_trace(batch_index: u64, chunk_traces: &Vec<BlockTrace>) {
    let path = PathBuf::from(PROVER_PROOF_DIR.to_string()).join(format!("batch_{batch_index}"));
    fs::create_dir_all(&path).unwrap();
    let file = File::create(path.join("block_traces.json")).unwrap();
    let writer = BufWriter::new(file);

    serde_json::to_writer_pretty(writer, &chunk_traces).unwrap();
    log::info!("chunk_traces of batch_index = {:#?} saved", batch_index);
}
