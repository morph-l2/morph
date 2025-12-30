//! Storage utilities.

/// `revm` database adapters backed by storage/witness data.
pub mod db;

pub use db::TrieDB;

use alloy_primitives::Bytes;
use alloy_rpc_types_debug::ExecutionWitness;
use prover_primitives::{types::BlockTrace, Block};
use std::collections::BTreeSet;

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

#[cfg(test)]
mod tests {
    use prover_primitives::{types::BlockTrace, Block};
    use rsp_mpt::EthereumState;
    use std::fs::File;
    use std::io::BufReader;

    use crate::trace_to_execution_witness;

    #[test]
    fn test_trace_to_execution_witness() {
        let block_trace = load_trace("../../testdata/mpt/local_transfer_eth.json");
        println!("loaded {} blocks", block_trace.len());
        let witness = trace_to_execution_witness(&block_trace[0]).unwrap();

        assert!(!witness.state.is_empty());
        assert!(!witness.codes.is_empty());
        let state = EthereumState::from_execution_witness(&witness, block_trace[0].root_before());

        // Check mpt state root equals to block trace root_before
        println!("built ethereum state from witness: {:?}", state.state_root());
        assert_eq!(state.state_root(), block_trace[0].root_before(), "state root mismatch");

        // Check number of accounts and storage tries
        let mut account_count = 0;
        state.state_trie.for_each_leaves(|_, _| account_count += 1);
        println!(
            "built state trie with {} accounts and {} storage tries",
            account_count,
            state.storage_tries.len()
        );
        assert_eq!(
            account_count,
            block_trace[0].storage_trace.proofs.as_ref().unwrap_or(&Default::default()).len(),
            "account_trie_count not expected"
        );
        assert_eq!(
            state.storage_tries.len(),
            block_trace[0].storage_trace.storage_proofs.len(),
            "storage_trie_count not expected"
        );
    }

    fn load_trace(file_path: &str) -> Vec<BlockTrace> {
        let file = File::open(file_path).unwrap();
        let reader = BufReader::new(file);
        serde_json::from_reader(reader).unwrap()
    }
}
