use alloy_provider::{Provider, ProviderBuilder};
use prover_executor_client::{
    types::input::{BlockInput, L2Block},
    EVMVerifier,
};
use prover_executor_host::{
    execute::HostExecutor,
    utils::{query_block, HostExecutorOutput, ProverBlock},
};
use prover_primitives::TxTrace;

pub async fn execute(block_number: u64, rpc: &str) {
    let provider = ProviderBuilder::new().connect_http(rpc.parse().unwrap()).erased();
    let prev_block = query_block(block_number, &provider).await.unwrap();

    let output = HostExecutor::execute_block(block_number, &provider).await.unwrap();
    let block_input = to_block_input(output, prev_block);
    let _batch_info = EVMVerifier::verify(vec![block_input]).unwrap();
}

fn to_block_input(output: HostExecutorOutput, prev_block: ProverBlock) -> BlockInput {
    let block = output.block;
    let state = output.state;
    let codes = output.codes;

    let l2_block = L2Block {
        chain_id: output.chain_id,
        coinbase: output.beneficiary,
        header: block.header,
        transactions: block
            .transactions
            .iter()
            .map(|tx_trace| tx_trace.try_build_tx_envelope().unwrap())
            .collect(),
        prev_state_root: output.prev_state_root,
        post_state_root: output.post_state_root,
        start_l1_queue_index: prev_block.header.next_l1_msg_index.to::<u64>(),
    };

    let block_input = BlockInput { current_block: l2_block, parent_state: state, bytecodes: codes };

    block_input
}

// cargo test -p morph-prove test_execute -- --nocapture -- --block-number 0x35 --rpc http://127.0.0.1:9545
#[test]
fn test_execute() {
    let rt = tokio::runtime::Runtime::new().unwrap();
    let (block_number, rpc) = test_args::read_execute_args_from_argv();
    rt.block_on(execute(block_number, &rpc));
}

#[cfg(test)]
mod test_args {
    const DEFAULT_BLOCK_NUMBER: u64 = 0x477;
    const DEFAULT_RPC: &str = "http://127.0.0.1:9545";
    pub(super) fn read_execute_args_from_argv() -> (u64, String) {
        let mut block_number = DEFAULT_BLOCK_NUMBER;
        let mut rpc = DEFAULT_RPC.to_string();

        let mut it = std::env::args().skip(1);
        while let Some(arg) = it.next() {
            match arg.as_str() {
                "--block-number" | "--block" => {
                    if let Some(v) = it.next() {
                        if let Some(parsed) = parse_u64_auto_radix(&v) {
                            block_number = parsed;
                        }
                    }
                }
                "--rpc" => {
                    if let Some(v) = it.next() {
                        rpc = v;
                    }
                }
                _ => {
                    // ignore unknown args
                }
            }
        }

        (block_number, rpc)
    }

    fn parse_u64_auto_radix(s: &str) -> Option<u64> {
        let s = s.trim();
        if let Some(hex) = s.strip_prefix("0x").or_else(|| s.strip_prefix("0X")) {
            u64::from_str_radix(hex, 16).ok()
        } else {
            s.parse::<u64>().ok()
        }
    }
}
