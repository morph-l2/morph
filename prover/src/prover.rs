use crate::utils::{
    get_block_traces_by_number, kzg_to_versioned_hash, GENERATE_EVM_VERIFIER, MAINNET_KZG_TRUSTED_SETUP, PROVER_L2_RPC,
    PROVER_PARAMS_DIR, PROVER_PROOF_DIR, PROVE_RESULT, PROVE_TIME, SCROLL_PROVER_ASSETS_DIR,
};
use c_kzg::{Blob, KzgCommitment, KzgProof};
use ethers::abi::AbiEncode;
use ethers::providers::Provider;
use ethers::types::U256;
use ethers::utils::hex;
use halo2_proofs::halo2curves::bn256::Fr;
use prover::aggregator::Prover as BatchProver;
use prover::config::{LayerId, LAYER4_DEGREE};
use prover::utils::chunk_trace_to_witness_block;
use prover::zkevm::Prover as ChunkProver;
use prover::{BatchHash, BlockTrace, ChunkHash, ChunkProof, CompressionCircuit, MAX_AGG_SNARKS};
use serde::{Deserialize, Serialize};
use std::fs;
use std::fs::File;
use std::io::{BufWriter, Write};
use std::iter::repeat;
use std::time::{Duration, Instant};
use std::{sync::Arc, thread};
use tokio::sync::Mutex;
use zkevm_circuits::witness::Block;

const BLOB_DATA_SIZE: usize = 4096 * 32;

// proveRequest
#[derive(Serialize, Deserialize, Debug)]
pub struct ProveRequest {
    pub batch_index: u64,
    pub chunks: Vec<Vec<u64>>,
    pub rpc: String,
}

/// Processes prove requests from a queue.
pub async fn prove_for_queue(prove_queue: Arc<Mutex<Vec<ProveRequest>>>) {
    let mut chunk_prover = ChunkProver::from_dirs(PROVER_PARAMS_DIR.as_str(), SCROLL_PROVER_ASSETS_DIR.as_str());
    log::info!("Waiting for prove request");
    loop {
        thread::sleep(Duration::from_millis(12000));

        // Step1. Get request from queue
        let (batch_index, chunks) = match prove_queue.lock().await.pop() {
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
        let provider = match Provider::try_from(PROVER_L2_RPC.as_str()) {
            Ok(provider) => provider,
            Err(e) => {
                log::error!("Failed to init provider: {:#?}", e);
                PROVE_RESULT.set(2);
                continue;
            }
        };
        log::info!("Requesting trace of batch: {:#?}", batch_index);
        let chunk_traces = match get_chunk_traces(batch_index, chunks, provider).await {
            Some(chunk_traces) => chunk_traces,
            None => vec![],
        };
        if chunk_traces.is_empty() {
            PROVE_RESULT.set(2);
            continue;
        }

        // Step3. Generate evm proof
        generate_proof(batch_index, chunk_traces, &mut chunk_prover).await;
    }
}

/// Generate EVM Proof for block trace.
async fn generate_proof(batch_index: u64, chunk_traces: Vec<Vec<BlockTrace>>, chunk_prover: &mut ChunkProver) {
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

    let mut chunk_hashes_proofs: Vec<(ChunkHash, ChunkProof)> = vec![];
    for (index, chunk_trace) in chunk_traces.iter().enumerate() {
        let chunk_witness = match chunk_trace_to_witness_block(chunk_trace.to_vec()) {
            Ok(_witness) => _witness,
            Err(e) => {
                log::error!("convert trace to witness of batch = {:#?} error: {:#?}", batch_index, e);
                PROVE_RESULT.set(2);
                return;
            }
        };
        let chunk_hash = ChunkHash::from_witness_block(&chunk_witness, false);

        log::info!(
            ">>Starting chunk prove, batchIndex = {:#?}, chunkIndex = {:#?}",
            batch_index,
            index
        );
        // Start chunk prove
        let chunk_proof: ChunkProof =
            match chunk_prover.gen_chunk_proof(chunk_trace.to_vec(), None, None, Some(proof_path.as_str())) {
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

        chunk_hashes_proofs.push((chunk_hash, chunk_proof));
    }
    if chunk_hashes_proofs.len() != chunk_traces.len() {
        log::error!("chunk proofs len err, batchIndex = {:#?} ", batch_index);
        return;
    }

    log::info!(">>Starting batch prove, batch index = {:#?}", batch_index);
    let mut batch_prover = BatchProver::from_dirs(PROVER_PARAMS_DIR.as_str(), SCROLL_PROVER_ASSETS_DIR.as_str());
    let batch_proof = batch_prover.gen_agg_evm_proof(chunk_hashes_proofs, None, Some(proof_path.clone().as_str()));

    // Start batch prove
    match batch_proof {
        Ok(proof) => {
            log::info!(">>batch prove complate, batch index = {:#?}", batch_index);
            PROVE_RESULT.set(1);
            // let params: ParamsKZG<Bn256> = prover::utils::load_params("params_dir", 26, None).unwrap();
            if GENERATE_EVM_VERIFIER.to_owned() {
                generate_evm_verifier(batch_prover, proof);
            }
        }
        Err(e) => {
            PROVE_RESULT.set(2);
            log::error!("batch_{:#?} prove err: {:#?}", batch_index, e);
        }
    }
    let duration = start.elapsed();
    let minutes = duration.as_secs() / 60;
    PROVE_TIME.set(minutes.try_into().unwrap());
    return;
}

fn compute_and_save_kzg(
    chunk_traces: &Vec<Vec<BlockTrace>>,
    batch_index: u64,
    proof_path: &String,
) -> Result<(), String> {
    // Sequencer trace to witness.
    let mut blocks: Vec<Block<Fr>> = vec![];
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
    // Witness to chunkhash
    let mut chunk_hashes: Vec<ChunkHash> = blocks
        .iter()
        .map(|block| ChunkHash::from_witness_block(&block, false))
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
        "lastest_hash_withdraw_root: {:?}",
        chunk_hashes[MAX_AGG_SNARKS - 1].withdraw_root
    );

    log::debug!(
        "prev_state_root of batch_{:?} = {:#?}",
        batch_index,
        hex::encode(&chunk_hashes[0].prev_state_root)
    );

    let blob = BatchHash::construct(&chunk_hashes).blob_assignments();
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
        hex::encode(&versioned_hash)
    );
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

    let mut params_file = File::create(format!("{}/blob_kzg.data", proof_path.as_str())).unwrap();
    match params_file.write_all(&blob_kzg[..]) {
        Ok(()) => (),
        Err(e) => {
            return Err(format!("save kzg proof of batch = {:#?} error: {:#?}", batch_index, e));
        }
    };

    Ok(())
}

fn generate_evm_verifier(mut batch_prover: BatchProver, proof: prover::BatchProof) {
    log::info!("Starting generate evm verifier");
    let verifier = prover::common::Verifier::<CompressionCircuit>::from_params(
        batch_prover.inner.params(*LAYER4_DEGREE).clone(),
        &batch_prover.get_vk().unwrap(),
    );

    let instances = proof.clone().proof_to_verify().instances();
    let num_instances: Vec<usize> = instances.iter().map(|l| l.len()).collect();

    let evm_proof = prover::EvmProof::new(
        proof.clone().proof_to_verify().proof().to_vec(),
        &proof.proof_to_verify().instances(),
        num_instances,
        batch_prover.inner.pk(LayerId::Layer4.id()),
    );
    fs::create_dir_all("evm_verifier").unwrap();
    verifier.evm_verify(&evm_proof.unwrap(), Some("evm_verifier"));
    log::info!("generate evm verifier complate");
}

async fn get_chunk_traces(
    batch_index: u64,
    chunks: Vec<Vec<u64>>,
    provider: Provider<ethers::providers::Http>,
) -> Option<Vec<Vec<BlockTrace>>> {
    let mut chunk_traces: Vec<Vec<BlockTrace>> = vec![];
    for chunk in chunks {
        let chunk_trace = match get_block_traces_by_number(&provider, &chunk).await {
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
    return chunk_traces;
}

#[tokio::test]
async fn test_generate_proof() {
    use dotenv::dotenv;

    dotenv().ok();
    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("debug")).init();

    let mut chunk_prover = ChunkProver::from_dirs(PROVER_PARAMS_DIR.as_str(), SCROLL_PROVER_ASSETS_DIR.as_str());
    log::info!("Chunk_prover initialized");

    let chunk_traces = load_trace(17);
    log::info!("Loading traces from file successful");

    log::info!("Starting generate proof");
    generate_proof(17, chunk_traces, &mut chunk_prover).await;
}
