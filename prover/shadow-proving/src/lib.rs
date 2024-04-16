use abi::{rollup_abi::Rollup, shadow_rollup_abi::ShadowRollup};
use ethers::{
    middleware::SignerMiddleware,
    providers::{Http, Provider},
    signers::LocalWallet,
};

pub mod abi;
pub mod metrics;
pub mod shadow_prove;
pub mod shadow_rollup;
pub mod util;

#[derive(Clone, Debug)]
pub struct BatchInfo {
    pub batch_index: u64,
    pub chunks: Vec<Vec<u64>>,
}

type RollupType = Rollup<SignerMiddleware<Provider<Http>, LocalWallet>>;
type ShadowRollupType = ShadowRollup<SignerMiddleware<Provider<Http>, LocalWallet>>;
