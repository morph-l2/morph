# Gas Price Oracle

Gas Price Oracle service monitors L1 gas prices and updates the GasPriceOracle contract on L2.

## Features

-   **L1 Base Fee Update**: Monitors L1 base fee and blob base fee, updates to L2
-   **Scalar Update**: Calculates and updates commit scalar and blob scalar
-   **Transaction Manager**: Serializes all contract updates to avoid nonce conflicts
-   **Metrics Monitoring**: Exposes Prometheus metrics
-   **Flags Configuration**: Uses `urfave/cli` for configuration management (supports both CLI flags and environment variables)

## Configuration

The service uses flags that can be set either via command line or environment variables (with `GAS_ORACLE_` prefix).

### Required Flags

| Flag                  | Env Var                     | Description                     |
| --------------------- | --------------------------- | ------------------------------- |
| `--l1-eth-rpc`        | `GAS_ORACLE_L1_ETH_RPC`     | L1 RPC endpoint                 |
| `--l2-eth-rpc`        | `GAS_ORACLE_L2_ETH_RPC`     | L2 RPC endpoint                 |
| `--l1-beacon-rpc`     | `GAS_ORACLE_L1_BEACON_RPC`  | L1 Beacon Chain API endpoint    |
| `--l1-rollup-address` | `GAS_ORACLE_L1_ROLLUP`      | L1 Rollup contract address      |
| `--private-key`       | `GAS_ORACLE_L2_PRIVATE_KEY` | Private key for L2 transactions |

### Optional Flags

| Flag                            | Env Var                            | Default         | Description                 |
| ------------------------------- | ---------------------------------- | --------------- | --------------------------- |
| `--l2-gas-price-oracle-address` | `GAS_ORACLE_L2_GAS_PRICE_ORACLE`   | `0x5300...0002` | L2 GasPriceOracle contract  |
| `--gas-threshold`               | `GAS_ORACLE_GAS_THRESHOLD`         | `10`            | Update threshold percentage |
| `--interval`                    | `GAS_ORACLE_INTERVAL`              | `6s`            | Base fee update interval    |
| `--overhead-interval`           | `GAS_ORACLE_OVERHEAD_INTERVAL`     | `10`            | Scalar update frequency     |
| `--txn-per-batch`               | `GAS_ORACLE_TXN_PER_BATCH`         | `50`            | Expected txs per batch      |
| `--log-level`                   | `GAS_ORACLE_LOG_LEVEL`             | `info`          | Log level                   |
| `--log-filename`                | `GAS_ORACLE_LOG_FILENAME`          | -               | Log file path               |
| `--metrics-server-enable`       | `GAS_ORACLE_METRICS_SERVER_ENABLE` | `false`         | Enable metrics server       |
| `--metrics-hostname`            | `GAS_ORACLE_METRICS_HOSTNAME`      | `0.0.0.0`       | Metrics server host         |
| `--metrics-port`                | `GAS_ORACLE_METRICS_PORT`          | `6060`          | Metrics server port         |

## Usage

### Command Line

```bash
./bin/token-price-oracle \
  --l1-eth-rpc https://ethereum-rpc.com \
  --l2-eth-rpc https://morph-l2-rpc.com \
  --l1-beacon-rpc https://beacon-api.com \
  --l1-rollup-address 0x... \
  --private-key 0x... \
  --metrics-server-enable \
  --log-level debug
```

### Environment Variables

```bash
export GAS_ORACLE_L1_ETH_RPC="https://ethereum-rpc.com"
export GAS_ORACLE_L2_ETH_RPC="https://morph-l2-rpc.com"
export GAS_ORACLE_L1_BEACON_RPC="https://beacon-api.com"
export GAS_ORACLE_L1_ROLLUP="0x..."
export GAS_ORACLE_L2_PRIVATE_KEY="0x..."
export GAS_ORACLE_METRICS_SERVER_ENABLE=true
export GAS_ORACLE_LOG_LEVEL=info

./bin/token-price-oracle
```

## Build and Run

**Note**: This project uses Go workspace and depends on `../bindings` module.

```bash
# Build
make build

# Run
make run

# Test
make test

# Test Bitget price feed (requires network)
go test ./client -run TestBitgetPriceFeed -v

# Docker
make docker-build
docker run -d \
  -e GAS_ORACLE_L1_ETH_RPC="..." \
  -e GAS_ORACLE_L2_ETH_RPC="..." \
  -e GAS_ORACLE_L1_BEACON_RPC="..." \
  -e GAS_ORACLE_L1_ROLLUP="0x..." \
  -e GAS_ORACLE_L2_PRIVATE_KEY="0x..." \
  morph/token-price-oracle:latest
```

## Monitoring

When metrics server is enabled, it exposes metrics at `<hostname>:<port>/metrics`:

-   `l1_base_fee` - L1 base fee (Gwei)
-   `l1_base_fee_on_l2` - L1 base fee on L2
-   `l1_blob_base_fee_on_l2` - L1 blob base fee on L2
-   `commit_scalar` - Commit scalar value
-   `blob_scalar` - Blob scalar value
-   `txn_per_batch` - Transactions per batch
-   `gas_oracle_owner_balance` - Oracle account balance
-   `base_fee_update_count` - Total base fee updates
-   `scalar_update_count` - Total scalar updates
-   `update_errors_total` - Update errors by type

Health check endpoint: `<hostname>:<port>/health`

## Architecture

```
gas-price-oracle/
├── cmd/              # Main entry point
├── flags/            # CLI flags definitions
├── config/           # Configuration from flags
├── updater/          # Update implementations
│   ├── basefee.go    # Base fee updater
│   ├── scalar.go     # Scalar updater
│   └── tx_manager.go # Transaction manager (prevents nonce conflicts)
├── client/           # Client wrappers
├── calc/             # Calculation logic
└── metrics/          # Prometheus metrics

Uses: ../bindings/bindings (project root contract bindings)
```

## Key Components

### Transaction Manager

All contract updates are serialized through `TxManager` to prevent nonce conflicts:

-   Holds a mutex to ensure only one transaction is sent at a time
-   Manages nonce retrieval and transaction confirmation
-   Used by both `BaseFeeUpdater` and `ScalarUpdater`

### Base Fee Updater

-   Runs on a fixed interval (default 6s)
-   Fetches L1 base fee and blob base fee
-   Updates L2 contract when threshold is exceeded

### Scalar Updater

-   Runs every N base fee update cycles (default 10)
-   Reads `CommitBatch` events from L1 Rollup
-   Calculates commit and blob scalars
-   Updates L2 contract when necessary

### Blob Processing

Blob data processing is partially implemented (interface defined in `calc/blob.go`). The actual blob parsing and L2 transaction extraction is deferred for future implementation.

## Testing

```bash
# Run all tests
go test ./...

# Test Bitget price feed (requires network)
go test ./client -run TestBitgetPriceFeed -v

# Skip integration tests
go test ./... -short
```

## Documentation

-   [PRICE_UPDATE.md](./PRICE_UPDATE.md) - Token price update 功能
-   [BITGET_PRICE_FEED.md](./BITGET_PRICE_FEED.md) - Bitget 价格源使用指南
