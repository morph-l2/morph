use once_cell::sync::Lazy;
use prometheus::{Gauge, IntGauge, Registry};

pub struct Metrics {
    pub shadow_batch_index: IntGauge,
    pub shadow_blocks_len: IntGauge,
    pub shadow_txn_len: IntGauge,
    pub shadow_verify_result: IntGauge,
    pub shadow_wallet_balance: Gauge,
}

pub static REGISTRY: Lazy<Registry> = Lazy::new(Registry::new);
pub static METRICS: Lazy<Metrics> = Lazy::new(|| Metrics {
    shadow_batch_index: IntGauge::new("shadow_batch_index", "shadow batch index")
        .expect("shadow_batch_index metric can be created"),
    shadow_blocks_len: IntGauge::new("shadow_blocks_len", "shadow blocks len")
        .expect("shadow_blocks_len metric can be created"),
    shadow_txn_len: IntGauge::new("shadow_txn_len", "shadow txn len")
        .expect("shadow_txn_len metric can be created"),
    shadow_verify_result: IntGauge::new("shadow_verify_result", "shadow verify result")
        .expect("shadow_verify_result metric can be created"),
    shadow_wallet_balance: Gauge::new("shadow_wallet_balance", "shadow wallet balance")
        .expect("shadow_wallet_balance metric can be created"),
});
