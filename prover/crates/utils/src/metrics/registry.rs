use prometheus_client::{
    metrics::{
        counter::Counter,
        gauge::Gauge,
        histogram::{linear_buckets, Histogram},
    },
    registry,
};

#[derive(Debug)]
pub struct Registry {
    pub registry: registry::Registry,

    pub block_counter: Counter,
    pub fetched_rpc_block_height: Gauge,
    pub latest_rpc_block_height: Gauge,

    pub verification_error: Counter,

    pub build_zktrie_db_duration_microseconds: Histogram,
    pub update_db_duration_microseconds: Histogram,
    pub handle_block_duration_microseconds: Histogram,
    pub commit_changes_duration_microseconds: Histogram,
    pub total_block_verification_duration_microseconds: Histogram,
}

pub(super) fn init() -> Registry {
    let mut registry = registry::Registry::default();

    let block_counter = Counter::default();
    registry.register(
        "block_counter",
        "Number of blocks processed",
        block_counter.clone(),
    );

    let fetched_rpc_block_height = Gauge::default();
    registry.register(
        "fetched_rpc_block_height",
        "Fetched RPC block height",
        fetched_rpc_block_height.clone(),
    );

    let latest_rpc_block_height = Gauge::default();
    registry.register(
        "latest_rpc_block_height",
        "Latest RPC block height",
        latest_rpc_block_height.clone(),
    );

    let verification_error = Counter::default();
    registry.register(
        "verification_error",
        "Number of verification errors",
        verification_error.clone(),
    );

    let build_zktrie_db_duration_microseconds = Histogram::new(linear_buckets(50.0, 50.0, 10));
    registry.register(
        "build_zktrie_db_duration",
        "Duration of build_zktrie_db_duration in microseconds",
        build_zktrie_db_duration_microseconds.clone(),
    );

    let update_db_duration_microseconds = Histogram::new(linear_buckets(2.0, 4.0, 10));
    registry.register(
        "update_db_duration",
        "Duration of update_db in microseconds",
        update_db_duration_microseconds.clone(),
    );

    let handle_block_duration_microseconds = Histogram::new(linear_buckets(1.0, 5.0, 10));
    registry.register(
        "handle_block_duration",
        "Duration of handle_block in microseconds",
        handle_block_duration_microseconds.clone(),
    );

    let commit_changes_duration_microseconds = Histogram::new(linear_buckets(25.0, 50.0, 10));
    registry.register(
        "commit_changes_duration",
        "Duration of commit_changes in microseconds",
        commit_changes_duration_microseconds.clone(),
    );

    let total_block_verification_duration_microseconds =
        Histogram::new(linear_buckets(50.0, 50.0, 15));
    registry.register(
        "total_block_verification_duration",
        "Total block verification duration in microseconds",
        total_block_verification_duration_microseconds.clone(),
    );

    Registry {
        registry,

        block_counter,
        fetched_rpc_block_height,
        latest_rpc_block_height,

        verification_error,

        build_zktrie_db_duration_microseconds,
        update_db_duration_microseconds,
        handle_block_duration_microseconds,
        commit_changes_duration_microseconds,
        total_block_verification_duration_microseconds,
    }
}
