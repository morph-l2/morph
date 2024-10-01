use sbv_primitives::{types::BlockTrace, B256};
use serde::{Deserialize, Serialize};
use serde_with::serde_as;

#[serde_as]
#[derive(Clone, Debug, Default, Serialize, Deserialize)]
pub struct BlobInfo {
    pub blob_data: Vec<u8>,
    pub commitment: Vec<u8>,
    pub proof: Vec<u8>,
}

#[serde_as]
#[derive(Clone, Debug, Default, Serialize, Deserialize)]
pub struct ClientInput {
    pub l2_traces: Vec<BlockTrace>,
}

#[serde_as]
#[derive(Clone, Debug, Default, Serialize, Deserialize)]
pub struct ShardInfo {
    pub chain_id: u64,
    pub prev_state_root: B256,
    pub post_state_root: B256,
    /// withdraw_root
    pub withdraw_root: B256,
    /// sequencer_root
    pub sequencer_root: B256,
    pub shard_data_hash: B256,
}

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct AggregationInput {
    pub shard_infos: Vec<ShardInfo>,
    pub shard_vkey: [u32; 8],
    pub l2_traces: Vec<Vec<BlockTrace>>,
    pub blob_info: BlobInfo,
}
