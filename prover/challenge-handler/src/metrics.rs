use lazy_static::lazy_static;
use prometheus::{Gauge, IntGauge, Registry};

pub struct Metrics {
    pub detected_batch_index: IntGauge,
    pub chunks_len: IntGauge,
    pub txn_len: IntGauge,
    pub verify_result: IntGauge,
    pub wallet_balance: Gauge,
}

lazy_static! {
    pub static ref REGISTRY: Registry = Registry::new();
    pub static ref METRICS: Metrics = Metrics {
        detected_batch_index: IntGauge::new("detected_batch_index", "detected batch index").expect("detected metric can be created"),
        chunks_len: IntGauge::new("chunks_len", "chunks len").expect("chunks_len metric can be created"),
        txn_len: IntGauge::new("txn_len", "txn len").expect("txn_len metric can be created"),
        verify_result: IntGauge::new("verify_result", "verify result").expect("verify metric can be created"),
        wallet_balance: Gauge::new("handler_wallet_balance", "handler wallet balance").expect("wallet metric can be created"),
    };
}
