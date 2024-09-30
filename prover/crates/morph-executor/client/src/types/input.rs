use sbv_primitives::types::BlockTrace;
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
    pub blob_info: BlobInfo,
}
