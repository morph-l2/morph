use ethers::providers::{Http, Provider};
use prover::BlockTrace;

pub const FS_PROVE_PARAMS: &'static str = "prove_params";
pub const FS_PROVE_SEED: &'static str = "prove_seed";
pub const FS_PROOF: &'static str = "proof";

// Fetches block traces by provider
pub async fn get_block_traces_by_number(
    provider: &Provider<Http>,
    block_nums: &Vec<u64>,
) -> Option<Vec<BlockTrace>> {
    let mut block_traces: Vec<BlockTrace> = Vec::new();
    for block_num in block_nums {
        log::info!("zkevm-prover: requesting trace of block {block_num}");
        let result = provider
            .request(
                "morph_getBlockTraceByNumberOrHash",
                [format!("{block_num:#x}")],
            )
            .await;

        match result {
            Ok(trace) => block_traces.push(trace),
            Err(e) => {
                log::error!("zkevm-prover: requesting trace error: {e}");
                return None;
            }
        }
    }
    Some(block_traces)
}
