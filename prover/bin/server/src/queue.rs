use std::{
    fs::{self, File},
    io::{BufReader, BufWriter, Write},
    path::PathBuf,
    sync::Arc,
    time::{Duration, Instant},
};

use crate::{read_env_var, PROVER_L2_RPC, PROVER_PROOF_DIR, PROVE_RESULT, PROVE_TIME};
use alloy_provider::{DynProvider, Provider, ProviderBuilder};
use morph_prove::{evm::EvmProofFixture, prove};
use prover_executor_client::{BlobVerifier, EVMVerifier};
use prover_executor_host::trace_to_input;
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

pub struct Prover {
    pub prove_queue: Arc<Mutex<Vec<ProveRequest>>>,
    provider: DynProvider,
}

impl Prover {
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
            let (batch_index, start_block, end_block) = match self.prove_queue.lock().await.pop() {
                Some(req) => {
                    log::info!(
                        "received prove request, batch index = {:#?}, blocks len = {:#?}, start_block = {:#?}, end_block = {:#?}",
                        req.batch_index,
                        req.end_block - req.start_block + 1,
                        req.start_block,
                        req.end_block,
                    );
                    (req.batch_index, req.start_block, req.end_block)
                }
                None => {
                    log::info!("no prove request");
                    continue;
                }
            };

            // let traces: &mut Vec<Vec<BlockTrace>> =
            //     &mut load_trace("testdata/mainnet_batch_traces.json");
            // let block_traces: &mut Vec<BlockTrace> = &mut traces[0];

            // Step2. Fetch trace
            log::info!("Requesting trace of batch-{:#?} ...", batch_index);
            let res_provider =
                &mut get_block_traces(batch_index, start_block, end_block, &self.provider).await;
            let block_traces = match res_provider {
                Some(block_traces) => block_traces,
                None => {
                    PROVE_RESULT.set(2);
                    continue;
                }
            };

            if read_env_var("SAVE_TRACE", false) {
                save_trace(batch_index, block_traces);
            }
            if !save_batch_header(block_traces, batch_index) {
                save_trace(batch_index, block_traces);
                continue;
            }

            // Step3. Generate evm proof
            log::info!("Generate evm proof");
            let start = Instant::now();
            let prove_rt = prove(block_traces, true);

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

fn save_batch_header(blocks: &mut Vec<BlockTrace>, batch_index: u64) -> bool {
    let proof_dir =
        PathBuf::from(PROVER_PROOF_DIR.to_string()).join(format!("batch_{}", batch_index));
    std::fs::create_dir_all(&proof_dir).expect("failed to create proof path");

    let blocks_inputs = blocks.iter().map(|trace| trace_to_input(trace)).collect::<Vec<_>>();
    let verify_result = EVMVerifier::verify(blocks_inputs);

    if let Ok(batch_info) = verify_result {
        let blob_info = prover_executor_host::blob::get_blob_info_from_traces(blocks).unwrap();
        let (versioned_hash, _) = BlobVerifier::verify(&blob_info, blocks.len()).unwrap();

        // Save batch_header
        // | batch_data_hash | versioned_hash | sequencer_root |
        // |-----------------|----------------|----------------|
        // |     bytes32     |     bytes32    |     bytes32    |
        let mut batch_header: Vec<u8> = Vec::with_capacity(96);
        batch_header.extend_from_slice(&batch_info.data_hash().0);
        batch_header.extend_from_slice(&versioned_hash.0);
        batch_header.extend_from_slice(&batch_info.sequencer_root().0);
        batch_header.extend_from_slice(&batch_info.sequencer_root().0);
        let mut batch_file = File::create(proof_dir.join("batch_header.data")).unwrap();
        batch_file.write_all(&batch_header[..]).expect("failed to batch_header");
        true
    } else {
        let e = verify_result.unwrap_err();
        let error_data = serde_json::json!({
            "error_code": "EVM_EXECUTE_NOT_EXPECTED",
            "error_msg": e.to_string()
        });
        let mut batch_file = File::create(proof_dir.join("execute_result.json")).unwrap();
        batch_file
            .write_all(serde_json::to_string_pretty(&error_data).unwrap().as_bytes())
            .expect("failed to write error");
        log::error!("EVM verification failed for batch {}: {}", batch_index, e);
        false
    }
}

fn save_proof(batch_index: u64, proof: EvmProofFixture) {
    let batch_dir =
        PathBuf::from(PROVER_PROOF_DIR.to_string()).join(format!("batch_{}", batch_index));
    std::fs::create_dir_all(&batch_dir).expect("failed to create proof path");
    std::fs::write(
        batch_dir.join("plonk_proof.json"),
        serde_json::to_string_pretty(&proof).unwrap(),
    )
    .expect("failed to write proof");
    log::info!("Successfully save evm proof of batch-{:?}", batch_index);
}

// Fetches block traces by provider
async fn get_block_traces(
    batch_index: u64,
    start_block: u64,
    end_block: u64,
    provider: &DynProvider,
) -> Option<Vec<BlockTrace>> {
    let mut block_traces: Vec<BlockTrace> = Vec::new();
    for block_num in start_block..end_block + 1 {
        log::debug!("requesting trace of block {block_num}");
        let result = provider
            .raw_request("morph_getBlockTraceByNumberOrHash".into(), [format!("{block_num:#x}")])
            .await;

        match result {
            Ok(trace) => block_traces.push(trace),
            Err(e) => {
                log::error!("requesting trace error: {e}");
                return None;
            }
        }
    }
    if (end_block + 1 - start_block) as usize != block_traces.len() {
        log::error!("block_traces.len not expected, batch index = {:#?}", batch_index);
        return None;
    }
    Some(block_traces)
}

#[allow(dead_code)]
fn load_trace(file_path: &str) -> Vec<Vec<BlockTrace>> {
    let file = File::open(file_path).unwrap();
    let reader = BufReader::new(file);
    serde_json::from_reader(reader).unwrap()
}

#[allow(dead_code)]
fn save_trace(batch_index: u64, chunk_traces: &Vec<BlockTrace>) {
    let path = PathBuf::from(PROVER_PROOF_DIR.to_string()).join(format!("batch_{}", batch_index));
    fs::create_dir_all(&path).unwrap();
    let file = File::create(path.join("block_traces.json")).unwrap();
    let writer = BufWriter::new(file);

    serde_json::to_writer_pretty(writer, &chunk_traces).unwrap();
    log::info!("chunk_traces of batch_index = {:#?} saved", batch_index);
}

#[test]
fn test_save_execute() {
    let batch_index = 102u64;

    let mut blocks = load_trace("../../testdata/viridian/eip7702_traces.json");
    println!("blocks.len(): {:?}", blocks.len());
    let traces = blocks.first_mut().unwrap();

    if !save_batch_header(traces, batch_index) {
        save_trace(batch_index, traces);
        println!("save_batch_header error");
    } else {
        println!("save_batch_header success");
    }
}
