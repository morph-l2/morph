use std::collections::BTreeSet;

use crate::ClientBlockInput;
use alloy_primitives::Bytes;
use alloy_rpc_types_debug::ExecutionWitness;
use prover_mpt::EthereumState;
use prover_primitives::{
    types::{block::L2Block, BlockTrace},
    Block,
};
use revm::state::Bytecode;

pub fn trace_to_input(trace: &BlockTrace) -> ClientBlockInput {
    let witness = trace_to_execution_witness(trace).unwrap();
    let state = EthereumState::from_execution_witness(&witness, trace.root_before());
    let bytecodes = witness.codes.into_iter().map(Bytecode::new_raw).collect::<Vec<_>>();
    ClientBlockInput {
        current_block: L2Block::from_block_trace(trace),
        parent_state: state,
        bytecodes,
    }
}

/// Converts a block trace to an execution witness.
pub fn trace_to_execution_witness(trace: &BlockTrace) -> anyhow::Result<ExecutionWitness> {
    let storage_trace = &trace.storage_trace;
    let mut state_nodes = BTreeSet::new();

    if let Some(proofs) = &storage_trace.proofs {
        for (_, proof) in proofs {
            for node in proof {
                state_nodes.insert(node.clone());
            }
        }
    }

    for (_, storage_proof) in &storage_trace.storage_proofs {
        for (_, proof) in storage_proof {
            for node in proof {
                state_nodes.insert(node.clone());
            }
        }
    }

    let state: Vec<_> = state_nodes.into_iter().collect();
    let codes: Vec<_> = trace.codes().map(Bytes::copy_from_slice).collect();

    Ok(ExecutionWitness { state, codes, keys: Default::default(), headers: Default::default() })
}
