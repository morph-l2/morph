use crate::utils::{
    get_block_traces_by_number, kzg_to_versioned_hash, MAINNET_KZG_TRUSTED_SETUP, PROVER_L2_RPC, PROVER_PARAMS_DIR,
    PROVER_PROOF_DIR, PROVE_RESULT, PROVE_TIME, SCROLL_PROVER_ASSETS_DIR,
};
use c_kzg::{Blob, KzgCommitment, KzgProof};
use ethers::abi::AbiEncode;
use ethers::providers::Provider;
use ethers::types::U256;
use ethers::utils::{hex, keccak256};
use prover::aggregator::Prover as BatchProver;
use prover::utils::chunk_trace_to_witness_block;
use prover::zkevm::Prover as ChunkProver;
use prover::{BatchHash, BatchProvingTask, BlockTrace, ChunkInfo, ChunkProof, ChunkProvingTask, MAX_AGG_SNARKS};
use serde::{Deserialize, Serialize};
use std::fs;
use std::fs::File;
use std::io::{BufWriter, Write};
use std::iter::repeat;
use std::time::{Duration, Instant};
use std::{sync::Arc, thread};
use tokio::sync::Mutex;
use tokio::time::timeout;
use zkevm_circuits::witness::Block;

const BLOB_DATA_SIZE: usize = 4096 * 32;

// proveRequest
#[derive(Serialize, Deserialize, Debug)]
pub struct ProveRequest {
    pub batch_index: u64,
    pub chunks: Vec<Vec<u64>>,
    pub rpc: String,
    pub shadow: Option<bool>,
}

pub struct Prover {
    prove_queue: Arc<Mutex<Vec<ProveRequest>>>,
    chunk_prover: ChunkProver,
    batch_prover: BatchProver,
    provider: Provider<ethers::providers::Http>,
}

impl Prover {
    pub fn new(prove_queue: Arc<Mutex<Vec<ProveRequest>>>) -> Self {
        let chunk_prover = ChunkProver::from_dirs(PROVER_PARAMS_DIR.as_str(), SCROLL_PROVER_ASSETS_DIR.as_str());
        let batch_prover = BatchProver::from_dirs(PROVER_PARAMS_DIR.as_str(), SCROLL_PROVER_ASSETS_DIR.as_str());
        let provider: Provider<ethers::providers::Http> = Provider::try_from(PROVER_L2_RPC.as_str()).unwrap();

        Self {
            prove_queue,
            chunk_prover,
            batch_prover,
            provider,
        }
    }

    /// Processes prove requests from a queue.
    pub async fn prove_for_queue(&mut self) {
        log::info!("Waiting for prove request");
        loop {
            thread::sleep(Duration::from_millis(12000));

            // Step1. Get request from queue
            let (batch_index, chunks) = match self.prove_queue.lock().await.pop() {
                Some(req) => {
                    log::info!(
                        "received prove request, batch index = {:#?}, chunks len = {:#?}",
                        req.batch_index,
                        req.chunks.len()
                    );
                    log::debug!(">>chunks details = {:#?}", req.chunks);
                    (req.batch_index, req.chunks.clone())
                }
                None => {
                    log::info!("no prove request");
                    continue;
                }
            };

            // Step2. Fetch trace
            log::info!("Requesting trace of batch: {:#?}", batch_index);
            let chunk_traces = match get_chunk_traces(batch_index, chunks, &self.provider).await {
                Some(chunk_traces) => chunk_traces,
                None => {
                    PROVE_RESULT.set(2);
                    continue;
                }
            };

            // Step3. Generate evm proof
            self.generate_proof(batch_index, chunk_traces).await;
        }
    }

    /// Generate EVM Proof for block trace.
    pub async fn generate_proof(&mut self, batch_index: u64, chunk_traces: Vec<Vec<BlockTrace>>) {
        let start = Instant::now();

        let proof_path = PROVER_PROOF_DIR.to_string() + format!("/batch_{}", batch_index).as_str();
        fs::create_dir_all(proof_path.clone()).unwrap();

        if let Err(err_str) = compute_and_save_kzg(&chunk_traces, batch_index, &proof_path) {
            log::error!(
                "compute_and_save_kzg of batch = {:#?} error: {:#?}",
                batch_index,
                err_str
            );
            PROVE_RESULT.set(2);
            return;
        }

        let mut chunk_proofs: Vec<ChunkProof> = vec![];
        for (index, chunk_trace) in chunk_traces.iter().enumerate() {
            log::info!(
                ">>Starting chunk prove, batchIndex = {:#?}, chunkIndex = {:#?}",
                batch_index,
                index
            );
            // Start chunk prove
            let chunk_proof: ChunkProof = match self.chunk_prover.gen_chunk_proof(
                ChunkProvingTask::from(chunk_trace.to_vec()),
                None,
                None,
                Some(proof_path.as_str()),
            ) {
                Ok(proof) => {
                    log::info!(">>chunk_{:#?} prove complate, batch index = {:#?}", index, batch_index);
                    proof
                }
                Err(e) => {
                    log::error!("chunk in batch_{:#?} prove err: {:#?}", batch_index, e);
                    PROVE_RESULT.set(2);
                    return;
                }
            };

            //save chunk.protocol
            let protocol = &chunk_proof.protocol;
            let mut params_file = File::create(SCROLL_PROVER_ASSETS_DIR.to_string() + "/chunk.protocol").unwrap();
            params_file.write_all(&protocol[..]).unwrap();
            chunk_proofs.push(chunk_proof);

            // Check high-priority request
            let queue_lock = match timeout(Duration::from_secs(2), self.prove_queue.lock()).await {
                Ok(queue_lock) => queue_lock,
                Err(_) => continue,
            };
            if queue_lock.last().is_some() && !queue_lock.last().unwrap().shadow.unwrap_or(false) {
                // First handle the high-priority request
                log::info!("Received high-priority request, End the current task");
                return;
            }
        }
        if chunk_proofs.len() != chunk_traces.len() {
            log::error!("chunk proofs len err, batchIndex = {:#?} ", batch_index);
            return;
        }

        log::info!(">>Starting batch prove, batch index = {:#?}", batch_index);
        let batch_task = BatchProvingTask { chunk_proofs };

        // Start batch prove
        let batch_proof = self
            .batch_prover
            .gen_agg_evm_proof(batch_task, None, Some(proof_path.clone().as_str()));

        match batch_proof {
            Ok(_) => {
                log::info!(">>batch prove complate, batch index = {:#?}", batch_index);
                PROVE_RESULT.set(1);
            }
            Err(e) => {
                PROVE_RESULT.set(2);
                log::error!("batch_{:#?} prove err: {:#?}", batch_index, e);
            }
        }
        let duration = start.elapsed();
        let minutes = duration.as_secs() / 60;
        PROVE_TIME.set(minutes.try_into().unwrap());
    }
}

fn compute_and_save_kzg(chunk_traces: &Vec<Vec<BlockTrace>>, batch_index: u64, proof_path: &str) -> Result<(), String> {
    // Sequencer trace to witness.
    let mut blocks: Vec<Block> = vec![];
    for trace in chunk_traces {
        let block = match chunk_trace_to_witness_block(trace.to_vec()) {
            Ok(b) => b,
            Err(e) => {
                log::error!("batch_{:#?} chunk_trace_to_witness_block err: {:#?}", batch_index, e);
                return Err("chunk trace to witness fail".to_string());
            }
        };
        blocks.push(block);
    }
    // Witness to ChunkInfo
    let mut chunk_hashes: Vec<ChunkInfo> = blocks
        .iter()
        .map(|block| ChunkInfo::from_witness_block(block, false))
        .collect();

    if chunk_hashes.is_empty() {
        return Err("chunk_hashes is empty".to_string());
    }

    // Padding
    let number_of_valid_chunks = chunk_hashes.len();
    if number_of_valid_chunks < MAX_AGG_SNARKS {
        let mut padding_chunk_hash = chunk_hashes.last().unwrap().clone();
        padding_chunk_hash.is_padding = true;
        chunk_hashes.extend(repeat(padding_chunk_hash).take(MAX_AGG_SNARKS - number_of_valid_chunks));
    }
    log::debug!(
        "withdraw_root of batch_{:?} = {:#?}",
        batch_index,
        hex::encode(chunk_hashes[MAX_AGG_SNARKS - 1].withdraw_root)
    );

    log::debug!(
        "prev_state_root of batch_{:?} = {:#?}",
        batch_index,
        hex::encode(chunk_hashes[0].prev_state_root)
    );
    log::debug!(
        "post_state_root of batch_{:?} = {:#?}",
        batch_index,
        hex::encode(chunk_hashes[MAX_AGG_SNARKS - 1].post_state_root)
    );

    let data_hash_preimage = chunk_hashes
        .iter()
        .take(number_of_valid_chunks)
        .flat_map(|chunk_info| chunk_info.data_hash.0.iter())
        .cloned()
        .collect::<Vec<_>>();
    let batch_data_hash = keccak256(data_hash_preimage);
    let sequencer_root = chunk_hashes[MAX_AGG_SNARKS - 1].sequencer_root.as_bytes();

    let blob = BatchHash::<MAX_AGG_SNARKS>::construct(&chunk_hashes).point_evaluation_assignments();
    let challenge = blob.challenge; // z
    let evaluation = blob.evaluation; // y
    let mut blob_bytes = [0u8; BLOB_DATA_SIZE];
    for (index, value) in blob.coefficients.iter().enumerate() {
        value.to_big_endian(&mut blob_bytes[index * 32..(index + 1) * 32]);
    }
    let kzg_settings: Arc<c_kzg::KzgSettings> = Arc::clone(&MAINNET_KZG_TRUSTED_SETUP);
    let commitment = match KzgCommitment::blob_to_kzg_commitment(&Blob::from_bytes(&blob_bytes).unwrap(), &kzg_settings)
    {
        Ok(c) => c,
        Err(e) => {
            return Err(format!(
                "generate KzgCommitment of batch = {:#?} error: {:#?}",
                batch_index, e
            ));
        }
    };
    let versioned_hash = kzg_to_versioned_hash(commitment.to_bytes().to_vec().as_slice());
    log::info!(
        "=========> blob_versioned_hash of batch_{:?} = {:#?}",
        batch_index,
        hex::encode(versioned_hash.clone())
    );

    // Save batch_header
    // | batch_data_hash | versioned_hash | sequencer_root |
    // |-----------------|----------------|----------------|
    // |     bytes32     |     bytes32    |     bytes32    |
    let mut batch_header: Vec<u8> = Vec::with_capacity(96);
    batch_header.extend_from_slice(&batch_data_hash);
    batch_header.extend_from_slice(&versioned_hash);
    batch_header.extend_from_slice(sequencer_root);
    let mut batch_file = File::create(format!("{}/batch_header.data", proof_path)).unwrap();
    match batch_file.write_all(&batch_header[..]) {
        Ok(()) => (),
        Err(e) => {
            return Err(format!(
                "save header_info of batch = {:#?} error: {:#?}",
                batch_index, e
            ));
        }
    };

    let mut z = [0u8; 32];
    challenge.to_big_endian(&mut z);
    let (proof, y) =
        match KzgProof::compute_kzg_proof(&Blob::from_bytes(&blob_bytes).unwrap(), &z.into(), &kzg_settings) {
            Ok((proof, y)) => (proof, y),
            Err(e) => {
                return Err(format!(
                    "compute kzg proof of batch = {:#?} error: {:#?}",
                    batch_index, e
                ));
            }
        };

    log::debug!(
        "=========> y_from_barycentric = {:#?}, y_from_compute_kzg_proof = {:#?}",
        hex::encode(evaluation.encode()),
        hex::encode(y.as_slice())
    );
    if evaluation != U256::from_big_endian(y.as_slice()) {
        return Err("y_from_barycentric != y_from_compute_kzg_proof".to_string());
    }

    // Save 4844 kzgData
    // | z       | y       | kzg_commitment | kzg_proof |
    // |---------|---------|----------------|-----------|
    // | bytes32 | bytes32 | bytes48        | bytes48   |
    let mut blob_kzg = Vec::with_capacity(160);
    blob_kzg.extend_from_slice(z.as_slice());
    blob_kzg.extend_from_slice(y.as_slice());
    blob_kzg.extend_from_slice(commitment.as_slice());
    blob_kzg.extend_from_slice(proof.as_slice());

    let mut params_file = File::create(format!("{}/blob_kzg.data", proof_path)).unwrap();
    match params_file.write_all(&blob_kzg[..]) {
        Ok(()) => (),
        Err(e) => {
            return Err(format!("save kzg proof of batch = {:#?} error: {:#?}", batch_index, e));
        }
    };

    Ok(())
}

async fn get_chunk_traces(
    batch_index: u64,
    chunks: Vec<Vec<u64>>,
    provider: &Provider<ethers::providers::Http>,
) -> Option<Vec<Vec<BlockTrace>>> {
    let mut chunk_traces: Vec<Vec<BlockTrace>> = vec![];
    for chunk in chunks {
        let chunk_trace = match get_block_traces_by_number(provider, &chunk).await {
            Some(traces) => traces,
            None => {
                log::error!("No trace obtained for batch: {:#?}", batch_index);
                return None;
            }
        };
        if chunk.len() != chunk_trace.len() {
            log::error!("chunk_trace.len not expect, batch index = {:#?}", batch_index);
            return None;
        }
        chunk_traces.push(chunk_trace)
    }
    Some(chunk_traces)
}

#[allow(dead_code)]
fn save_trace(batch_index: u64, chunk_traces: &Vec<Vec<BlockTrace>>) {
    let path = PROVER_PROOF_DIR.to_string() + format!("/batch_{}", batch_index).as_str();
    fs::create_dir_all(path.clone()).unwrap();
    let file = File::create(format!("{}/chunk_traces.json", path.as_str())).unwrap();
    let writer = BufWriter::new(file);

    serde_json::to_writer_pretty(writer, &chunk_traces).unwrap();

    log::info!("chunk_traces of batch_index = {:#?} saved", batch_index);
}

#[cfg(test)]
fn load_trace(batch_index: u64) -> Vec<Vec<BlockTrace>> {
    use std::io::BufReader;

    let path = PROVER_PROOF_DIR.to_string() + format!("/batch_{}", batch_index).as_str();
    let file = File::open(format!("{}/chunk_traces.json", path.as_str())).unwrap();
    let reader = BufReader::new(file);

    let chunk_traces: Vec<Vec<BlockTrace>> = serde_json::from_reader(reader).unwrap();
    chunk_traces
}

#[tokio::test]
async fn test_generate_proof() {
    use dotenv::dotenv;
    dotenv().ok();
    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("debug")).init();

    let queue: Arc<Mutex<Vec<ProveRequest>>> = Arc::new(Mutex::new(Vec::new()));
    let mut prover = Prover::new(queue);

    let chunk_traces = load_trace(17);
    log::info!("Loading traces from file successful");

    log::info!("Starting generate proof");
    prover.generate_proof(17, chunk_traces).await;
}
