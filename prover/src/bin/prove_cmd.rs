use env_logger::Env;
use ethers::providers::{Http, Provider};
use prover::zkevm::Prover;
use std::env::var;
use zkevm_prover::utils::{get_block_traces_by_number, FS_PROVE_PARAMS, FS_PROVE_SEED};

// Used to generate poof for specified blocks in the development environment.
// It will read the cmd parameters and execute the following process.
// Main flow:
// 1. Init env
// 2. Fetch block traces
// 3. Create prover
// 4. Generate proof
// 5. Write proof & verifier
#[tokio::main]
async fn main() {
    // // Step1. prepare param
    // env_logger::Builder::from_env(Env::default().default_filter_or("info")).init();
    // let block_num: u64 = var("PROVERD_BLOCK_NUM")
    //     .expect("PROVERD_BLOCK_NUM env var")
    //     .parse()
    //     .expect("Cannot parse PROVERD_BLOCK_NUM env var");
    // let rpc_url: String = var("PROVERD_RPC_URL")
    //     .expect("PROVERD_RPC_URL env var")
    //     .parse()
    //     .expect("Cannot parse PROVERD_RPC_URL env var");
    // let params_path: String = var("PARAMS_PATH")
    //     .expect("PARAMS_PATH env var")
    //     .parse()
    //     .expect("Cannot parse PARAMS_PATH env var");

    // let provider =
    //     Provider::<Http>::try_from(rpc_url).expect("failed to initialize ethers Provider");

    // // Step 2. fetch block trace
    // let block_traces = get_block_traces_by_number(&provider, block_num, block_num + 1)
    //     .await
    //     .unwrap();

    // log::info!("block_traces_len is: {:#?}", block_traces.len());

    // log::info!("block_traces_chain_id is: {:#?}", block_traces[0].chain_id);

    // // Step 3. create prover
    // let mut prover = create_prover(params_path);

    // // Step 4. start prove
    // let result = prover.gen_chunk_proof(block_traces, None, None, Some("./proof"));
    // match result {
    //     Ok(proof) => {
    //         log::info!("prove result is: {:#?}", proof);
    //     }
    //     Err(e) => {
    //         log::info!("prove err: {:#?}", e);
    //     }
    // };
}

/**
 * Create prover of zkevm
 */
fn create_prover(params_path: String) -> Prover {
    let mut prover = Prover::from_dirs(FS_PROVE_PARAMS, FS_PROVE_SEED);
    prover
}
