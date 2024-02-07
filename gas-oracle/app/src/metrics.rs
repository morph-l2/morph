use lazy_static::lazy_static;
use prometheus::{IntGauge, Gauge, Registry};


pub struct OracleServiceMetrics {
    pub l1_base_fee: Gauge,
    pub l1_base_fee_on_l2: Gauge,
    pub gas_oracle_owner_balance: Gauge,
    pub overhead: IntGauge,
    pub scalar_ratio: Gauge,
    pub txn_per_batch: Gauge,
    pub l1_rpc_status: IntGauge,
}

lazy_static! {
    pub static ref REGISTRY: Registry = Registry::new();

    pub static ref ORACLE_SERVICE_METRICS: OracleServiceMetrics = OracleServiceMetrics {
        l1_base_fee: Gauge::new("l1BaseFee", "l1 base fee").expect("metric can be created"),
        l1_base_fee_on_l2: Gauge::new("l1BaseFeeOnL2", "l1 base fee on l2").expect("metric can be created"),
        gas_oracle_owner_balance: Gauge::new("gasOracleOwnerBalance", "gas oracle owner balance").expect("metric can be created"),
        overhead: IntGauge::new("gasOracleOverhead", "gas oracle overhead").expect("metric can be created"),
        scalar_ratio: Gauge::new("scalarRatio", "gas oracle scalar").expect("metric can be created"),
        txn_per_batch: Gauge::new("txnPerBatch", "txn per batch").expect("metric can be created"),
        l1_rpc_status: IntGauge::new("l1RpcStatus", "l1 prc").expect("metric can be created"),
    };
}
