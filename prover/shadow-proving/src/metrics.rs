use lazy_static::lazy_static;
use prometheus::{Gauge, IntGauge, Registry};

pub struct Metrics {
    pub shadow_batch_index: IntGauge,
    pub shadow_chunks_len: IntGauge,
    pub shadow_txn_len: IntGauge,
    pub shadow_verify_result: IntGauge,
    pub shadow_wallet_balance: Gauge,
}

lazy_static! {
    pub static ref REGISTRY: Registry = Registry::new();
    pub static ref METRICS: Metrics = Metrics {
        shadow_batch_index: IntGauge::new("shadow_batch_index", "shadow batch index").expect("shadow_batch_index metric can be created"),
        shadow_chunks_len: IntGauge::new("shadow_chunks_len", "shadow chunks len").expect("shadow_chunks_len metric can be created"),
        shadow_txn_len: IntGauge::new("shadow_txn_len", "shadow txn len").expect("shadow_txn_len metric can be created"),
        shadow_verify_result: IntGauge::new("shadow_verify_result", "shadow verify result").expect("shadow_verify_result metric can be created"),
        shadow_wallet_balance: Gauge::new("shadow_wallet_balance", "shadow wallet balance").expect("shadow_wallet_balance metric can be created"),
    };
}
