use crate::utils::{
    get_block_traces_by_number, GENERATE_EVM_VERIFIER, PROVER_L2_RPC, PROVER_PARAMS_DIR, PROVER_PROOF_DIR, PROVE_RESULT, PROVE_TIME, SCROLL_PROVER_ASSETS_DIR
};
use ethers::providers::Provider;
use prover::aggregator::Prover as BatchProver;
use prover::config::{LayerId, LAYER4_DEGREE};
use prover::utils::chunk_trace_to_witness_block;
use prover::zkevm::Prover as ChunkProver;
use prover::{BlockTrace, ChunkHash, ChunkProof, CompressionCircuit};
use serde::{Deserialize, Serialize};
use std::fs;
use std::fs::File;
use std::io::Write;
use std::time::{Duration, Instant};
use std::{sync::Arc, thread};
use tokio::sync::Mutex;

// proveRequest
#[derive(Serialize, Deserialize, Debug)]
pub struct ProveRequest {
    pub batch_index: u64,
    pub chunks: Vec<Vec<u64>>,
    pub rpc: String,
}

/// Generate EVM Proof for block trace.
pub async fn prove_for_queue(prove_queue: Arc<Mutex<Vec<ProveRequest>>>) {
    let mut chunk_prover = ChunkProver::from_dirs(PROVER_PARAMS_DIR.as_str(), SCROLL_PROVER_ASSETS_DIR.as_str());
    log::info!("Waiting for prove request");
    loop {
        thread::sleep(Duration::from_millis(12000));

        // Step1. Get request from queue
        let (batch_index, chunks) = {
            let queue_lock = prove_queue.lock().await;
            let req = match queue_lock.first() {
                Some(req) => {
                    log::info!(
                        ">>Received prove request, batch index = {:#?}, chunks len = {:#?}",
                        req.batch_index,
                        req.chunks.len()
                    );
                    req
                }
                None => {
                    log::debug!("no prove request");
                    continue;
                }
            };
            (req.batch_index, req.chunks.clone())
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
        log::info!("requesting trace of batch: {:#?}", batch_index);
        let chunk_traces = match get_chunk_traces(batch_index, chunks, provider).await {
            Some(chunk_traces) => chunk_traces,
            None => vec![],
        };
        if chunk_traces.is_empty() {
            prove_queue.lock().await.pop();
            PROVE_RESULT.set(2);
            continue;
        }

        // Step3. Generate evm proof
        generate_proof(batch_index, chunk_traces, &mut chunk_prover).await;
        prove_queue.lock().await.pop();
    }
}

async fn generate_proof(batch_index: u64, chunk_traces: Vec<Vec<BlockTrace>>, chunk_prover: &mut ChunkProver) {
    let start = Instant::now();

    let proof_path = PROVER_PROOF_DIR.to_string() + format!("/batch_{}", batch_index).as_str();
    fs::create_dir_all(proof_path.clone()).unwrap();
    let mut chunk_proofs: Vec<(ChunkHash, ChunkProof)> = vec![];
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

        chunk_proofs.push((chunk_hash, chunk_proof));
    }
    if chunk_proofs.len() != chunk_traces.len() {
        log::error!("chunk proofs len err, batchIndex = {:#?} ", batch_index);
        return;
    }
    log::info!(">>Starting batch prove, batch index = {:#?}", batch_index);
    let mut batch_prover = BatchProver::from_dirs(PROVER_PARAMS_DIR.as_str(), SCROLL_PROVER_ASSETS_DIR.as_str());
    let batch_proof = batch_prover.gen_agg_evm_proof(chunk_proofs, None, Some(proof_path.clone().as_str()));

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

#[tokio::test]
async fn test() {
    use std::fs::File;
    use std::io::Write;

    let protocol: Vec<u8> = vec![1, 2, 3, 4];
    std::fs::create_dir_all("configs").unwrap();
    let mut params_file = File::create("configs/chunk.protocol").unwrap();
    params_file.write_all(&protocol[..]).unwrap();
}

