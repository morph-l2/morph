use abi::{Rollup, ShadowRollup};

pub mod abi;
pub mod metrics;
pub mod shadow_prove;
pub mod shadow_rollup;
pub mod util;

#[derive(Clone, Debug)]
pub struct BatchInfo {
    pub batch_index: u64,
    pub start_block: u64,
    pub end_block: u64,
}
