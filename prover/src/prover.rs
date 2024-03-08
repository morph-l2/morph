use crate::utils::{
    get_block_traces_by_number, GENERATE_EVM_VERIFIER, MAINNET_KZG_TRUSTED_SETUP, PROVER_L2_RPC, PROVER_PARAMS_DIR,
    PROVER_PROOF_DIR, PROVE_RESULT, PROVE_TIME, SCROLL_PROVER_ASSETS_DIR,
};
use bls12_381::Scalar as Fp;
use c_kzg::{Blob, KzgCommitment, KzgProof};
use eth_types::{ToLittleEndian, U256};
use ethers::providers::Provider;
use ethers::utils::keccak256;
use prover::aggregator::Prover as BatchProver;
use prover::config::{LayerId, LAYER4_DEGREE};
use prover::utils::{chunk_trace_to_witness_block, chunk_trace_to_witness_block_with_index};
use prover::zkevm::Prover as ChunkProver;
use prover::{BlockTrace, ChunkHash, ChunkProof, CompressionCircuit};
use serde::{Deserialize, Serialize};
use std::fs;
use std::fs::File;
use std::io::Write;
use std::time::{Duration, Instant};
use std::{sync::Arc, thread};
use tokio::sync::Mutex;
use zkevm_circuits::blob_circuit::block_to_blob;
use zkevm_circuits::blob_circuit::util::{poly_eval_partial, FP_S};

const BLOB_DATA_SIZE: usize = 4096 * 32;

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

    // get batch_blob from chunks
    let mut batch_blob = [0u8; BLOB_DATA_SIZE];
    let mut offset = 0;
    let mut data_hash_preimage = Vec::<u8>::new();
    for chunk_trace in chunk_traces.iter() {
        match chunk_trace_to_witness_block(chunk_trace.to_vec()) {
            Ok(witness) => {
                let partial_result = block_to_blob(&witness).unwrap();
                batch_blob[offset..partial_result.len()].copy_from_slice(&partial_result);
                offset += partial_result.len();

                let chunk_hash = ChunkHash::from_witness_block(&witness, false);
                data_hash_preimage.extend_from_slice(chunk_hash.data_hash.as_bytes())
            }
            Err(e) => {
                log::error!("convert trace to witness of batch = {:#?} error: {:#?}", batch_index, e);
                PROVE_RESULT.set(2);
                return;
            }
        };
    }
    let batch_data_hash = keccak256(data_hash_preimage);

    let kzg_settings: Arc<c_kzg::KzgSettings> = Arc::clone(&MAINNET_KZG_TRUSTED_SETUP);
    let commitment = match KzgCommitment::blob_to_kzg_commitment(&Blob::from_bytes(&batch_blob).unwrap(), &kzg_settings)
    {
        Ok(c) => c,
        Err(e) => {
            log::error!("generate KzgCommitment of batch = {:#?} error: {:#?}", batch_index, e);
            PROVE_RESULT.set(2);
            return;
        }
    };

    let mut pre: Vec<u8> = vec![];
    pre.extend(commitment.to_bytes().to_vec());
    pre.extend(batch_data_hash);
    let challenge_point = U256::from_little_endian(keccak256(pre.as_slice()).as_ref());
    // let challenge_point = U256::from(128);

    let (proof, y) = match KzgProof::compute_kzg_proof(
        &Blob::from_bytes(&batch_blob).unwrap(),
        &challenge_point.to_le_bytes().into(),
        &kzg_settings,
    ) {
        Ok((proof, y)) => (proof, y),
        Err(e) => {
            log::error!("compute kzg proof of batch = {:#?} error: {:#?}", batch_index, e);
            PROVE_RESULT.set(2);
            return;
        }
    };

    // save 4844 kzg proof
    // kzgData: [y(32) | commitment(48) | proof(48)]
    // https://github.com/morph-l2/morph/blob/eip4844-contract-verify/contracts/contracts/L1/rollup/Rollup.sol
    let mut blob_kzg = Vec::<u8>::new();
    blob_kzg.extend_from_slice(y.as_slice());
    blob_kzg.extend_from_slice(commitment.as_slice());
    blob_kzg.extend_from_slice(proof.as_slice());
    let mut params_file = File::create(SCROLL_PROVER_ASSETS_DIR.to_string() + "/blob_kzg.data").unwrap();
    params_file.write_all(&blob_kzg[..]).unwrap();

    // todo: get batch_commit from eth trace
    let batch_commit: U256 = U256::from(0);
    let mut index = 0;
    for chunk_trace in chunk_traces.iter() {
        let mut partial_result: U256 = U256::from(0);
        let mut chunk_witness = match chunk_trace_to_witness_block_with_index(
            chunk_trace.to_vec(),
            batch_commit,
            challenge_point,
            index,
            partial_result,
        ) {
            Ok(_witness) => _witness,
            Err(e) => {
                log::error!("convert trace to witness of batch = {:#?} error: {:#?}", batch_index, e);
                PROVE_RESULT.set(2);
                return;
            }
        };

        let partial_result_bytes = block_to_blob(&chunk_witness).ok();

        // compute partial_result from witness block;
        let omega = Fp::from(123).pow(&[(FP_S - 12) as u64, 0, 0, 0]);
        match block_to_blob(&chunk_witness) {
            Ok(blob) => {
                let mut result: Vec<Fp> = Vec::new();
                for chunk in blob.chunks(32) {
                    let reverse: Vec<u8> = chunk.iter().rev().cloned().collect();
                    result.push(Fp::from_bytes(reverse.as_slice().try_into().unwrap()).unwrap());
                }
                partial_result = U256::from(
                    poly_eval_partial(
                        result,
                        Fp::from_bytes(&chunk_witness.challenge_point.to_le_bytes()).unwrap(),
                        omega,
                        index,
                    )
                    .to_bytes(),
                );
                chunk_witness.partial_result = partial_result;
            }
            Err(e) => {
                log::error!("chunk-hash: block_to_blob failed: {}", e);
            }
        }

        let chunk_hash = ChunkHash::from_witness_block(&chunk_witness, false);

        log::info!(
            ">>Starting chunk prove, batchIndex = {:#?}, chunkIndex = {:#?}",
            batch_index,
            index
        );
        log::info!(
            "=========gen_chunk_proof_with_index, batch_commit= {:#?}, challenge_point= {:#?}, index= {:#?} ",
            batch_commit,
            challenge_point,
            index
        );

        // Start chunk prove
        let chunk_proof: ChunkProof = match chunk_prover.gen_chunk_proof_with_index(
            chunk_trace.to_vec(),
            batch_commit,
            challenge_point,
            index,
            partial_result,
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

        chunk_proofs.push((chunk_hash, chunk_proof));
        index += partial_result_bytes.unwrap().len() / 32;
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
