use lazy_static::lazy_static;
use prometheus::{Gauge, IntGauge, Registry};

pub struct OracleServiceMetrics {
    pub l1_base_fee: Gauge,
    pub l1_base_fee_on_l2: Gauge,
    pub l1_blob_base_fee_on_l2: Gauge,
    pub gas_oracle_owner_balance: Gauge,
    pub base_fee_scalar: Gauge,
    pub commit_scalar: IntGauge,
    pub blob_scalar: Gauge,
    pub txn_per_batch: Gauge,
    pub l1_rpc_status: IntGauge,
}

lazy_static! {
    pub static ref REGISTRY: Registry = Registry::new();
    pub static ref ORACLE_SERVICE_METRICS: OracleServiceMetrics = OracleServiceMetrics {
        l1_base_fee: Gauge::new("l1_base_fee", "l1 base fee").expect("metric can't  be created"),
        l1_base_fee_on_l2: Gauge::new("l1_base_fee_on_l2", "l1 base fee on l2")
        .expect("metric can't  be created"),
        l1_blob_base_fee_on_l2: Gauge::new("l1_blob_base_fee_on_l2", "l1 blob base fee on l2")
            .expect("metric can't  be created"),
        base_fee_scalar: Gauge::new("base_fee_scalar", "base fee scalar")
            .expect("metric can't  be created"),
        commit_scalar: IntGauge::new("commit_scalar", "commit scalar")
            .expect("metric can't  be created"),
        blob_scalar: Gauge::new("blob_scalar", "blob scalar").expect("metric can't  be created"),
        txn_per_batch: Gauge::new("txn_per_batch", "txn per batch").expect("metric can't  be created"),
        gas_oracle_owner_balance: Gauge::new(
            "gas_oracle_owner_balance",
            "gas oracle owner balance"
        )
        .expect("metric can't  be created"),
        l1_rpc_status: IntGauge::new("l1_rpc_status", "l1 prc").expect("metric can't  be created"),
    };
}
