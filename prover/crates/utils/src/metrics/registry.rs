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

    // database metrics
    pub build_zktrie_db_duration_milliseconds: Histogram,
    pub update_db_duration_milliseconds: Histogram,
    pub zktrie_get_duration_microseconds: Histogram,
    pub zktrie_update_duration_microseconds: Histogram,
    pub zktrie_delete_duration_microseconds: Histogram,
    pub zktrie_commit_duration_microseconds: Histogram,

    // executor metrics
    pub transact_commit_duration_milliseconds: Histogram,
    pub handle_block_duration_milliseconds: Histogram,
    pub commit_changes_duration_milliseconds: Histogram,
    pub total_block_verification_duration_milliseconds: Histogram,
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

    let build_zktrie_db_duration_milliseconds = Histogram::new(linear_buckets(50.0, 50.0, 10));
    registry.register(
        "build_zktrie_db_duration",
        "Duration of build_zktrie_db_duration in milliseconds",
        build_zktrie_db_duration_milliseconds.clone(),
    );

    let update_db_duration_milliseconds = Histogram::new(linear_buckets(2.0, 4.0, 10));
    registry.register(
        "update_db_duration",
        "Duration of update_db in milliseconds",
        update_db_duration_milliseconds.clone(),
    );

    let zktrie_get_duration_microseconds = Histogram::new(linear_buckets(50.0, 500.0, 10));
    registry.register(
        "zktrie_get_duration",
        "Duration of zktrie_get in microseconds",
        zktrie_get_duration_microseconds.clone(),
    );

    let zktrie_update_duration_microseconds = Histogram::new(linear_buckets(50.0, 500.0, 10));
    registry.register(
        "zktrie_update_duration",
        "Duration of zktrie_update in microseconds",
        zktrie_update_duration_microseconds.clone(),
    );

    let zktrie_delete_duration_microseconds = Histogram::new(linear_buckets(50.0, 500.0, 10));
    registry.register(
        "zktrie_delete_duration",
        "Duration of zktrie_delete in microseconds",
        zktrie_delete_duration_microseconds.clone(),
    );

    let zktrie_commit_duration_microseconds = Histogram::new(linear_buckets(100.0, 2000.0, 10));
    registry.register(
        "zktrie_commit_duration",
        "Duration of zktrie_commit in microseconds",
        zktrie_commit_duration_microseconds.clone(),
    );

    let transact_commit_duration_milliseconds = Histogram::new(linear_buckets(0.1, 15.0, 10));
    registry.register(
        "transact_commit_duration",
        "Duration of transact_commit in milliseconds",
        transact_commit_duration_milliseconds.clone(),
    );

    let handle_block_duration_milliseconds = Histogram::new(linear_buckets(1.0, 5.0, 10));
    registry.register(
        "handle_block_duration",
        "Duration of handle_block in milliseconds",
        handle_block_duration_milliseconds.clone(),
    );

    let commit_changes_duration_milliseconds = Histogram::new(linear_buckets(25.0, 50.0, 10));
    registry.register(
        "commit_changes_duration",
        "Duration of commit_changes in milliseconds",
        commit_changes_duration_milliseconds.clone(),
    );

    let total_block_verification_duration_milliseconds =
        Histogram::new(linear_buckets(50.0, 50.0, 15));
    registry.register(
        "total_block_verification_duration",
        "Total block verification duration in milliseconds",
        total_block_verification_duration_milliseconds.clone(),
    );

    Registry {
        registry,

        block_counter,
        fetched_rpc_block_height,
        latest_rpc_block_height,

        verification_error,

        build_zktrie_db_duration_milliseconds,
        update_db_duration_milliseconds,
        zktrie_get_duration_microseconds,
        zktrie_update_duration_microseconds,
        zktrie_delete_duration_microseconds,
        zktrie_commit_duration_microseconds,

        handle_block_duration_milliseconds,
        transact_commit_duration_milliseconds,
        commit_changes_duration_milliseconds,
        total_block_verification_duration_milliseconds,
    }
}
