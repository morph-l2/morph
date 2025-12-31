# Token Price Oracle

Token Price Oracle service monitors token prices and updates the price ratio between tokens and ETH to L2 on-chain contracts, enabling Alt Fee Token functionality.

## Features

- **Real-time Price Monitoring**: Fetches token USD prices from exchange APIs (Bitget)
- **Price Ratio Calculation**: Computes price ratio between tokens and ETH
- **Threshold-based Updates**: Only updates on-chain when price change exceeds threshold, saving Gas
- **Batch Updates**: Updates multiple token prices in a single `batchUpdatePrices` transaction
- **Fallback Mechanism**: Supports automatic switching between multiple data sources
- **Transaction Management**: Prevents nonce conflicts, supports local and external signing
- **Prometheus Monitoring**: Provides operational metrics

## Quick Start

### Environment Variables

```bash
# Required
export TOKEN_PRICE_ORACLE_L2_ETH_RPC="https://rpc.morphl2.io"
export TOKEN_PRICE_ORACLE_PRIVATE_KEY="0x..."
export TOKEN_PRICE_ORACLE_BITGET_API_BASE_URL="https://api.bitget.com"
export TOKEN_PRICE_ORACLE_TOKEN_MAPPING_BITGET="1:BTCUSDT,2:ETHUSDT"

# Optional
export TOKEN_PRICE_ORACLE_PRICE_UPDATE_INTERVAL="1m"
export TOKEN_PRICE_ORACLE_PRICE_THRESHOLD="100"  # 1% (100 bps)
export TOKEN_PRICE_ORACLE_METRICS_SERVER_ENABLE="true"
export TOKEN_PRICE_ORACLE_METRICS_PORT="6060"
export TOKEN_PRICE_ORACLE_LOG_LEVEL="info"
```

### Build and Run

```bash
# Build
make build

# Run
./build/bin/token-price-oracle

# Or use Docker
make docker-build
docker run -d \
  -e TOKEN_PRICE_ORACLE_L2_ETH_RPC="..." \
  -e TOKEN_PRICE_ORACLE_PRIVATE_KEY="..." \
  -e TOKEN_PRICE_ORACLE_BITGET_API_BASE_URL="..." \
  -e TOKEN_PRICE_ORACLE_TOKEN_MAPPING_BITGET="..." \
  morph/token-price-oracle:latest
```

## Configuration

### Required

| Environment Variable | Description |
|---------------------|-------------|
| `TOKEN_PRICE_ORACLE_L2_ETH_RPC` | L2 node RPC endpoint |
| `TOKEN_PRICE_ORACLE_PRIVATE_KEY` | Signing private key (local signing mode) |
| `TOKEN_PRICE_ORACLE_BITGET_API_BASE_URL` | Bitget API base URL |
| `TOKEN_PRICE_ORACLE_TOKEN_MAPPING_BITGET` | TokenID to trading pair mapping |

### Optional

| Environment Variable | Default | Description |
|---------------------|---------|-------------|
| `TOKEN_PRICE_ORACLE_PRICE_UPDATE_INTERVAL` | `1m` | Price update interval |
| `TOKEN_PRICE_ORACLE_PRICE_THRESHOLD` | `100` | Update threshold (basis points, 100=1%) |
| `TOKEN_PRICE_ORACLE_PRICE_FEED_PRIORITY` | `bitget` | Price feed priority |
| `TOKEN_PRICE_ORACLE_METRICS_SERVER_ENABLE` | `false` | Enable metrics server |
| `TOKEN_PRICE_ORACLE_METRICS_HOSTNAME` | `0.0.0.0` | Metrics server hostname |
| `TOKEN_PRICE_ORACLE_METRICS_PORT` | `6060` | Metrics server port |
| `TOKEN_PRICE_ORACLE_LOG_LEVEL` | `info` | Log level |
| `TOKEN_PRICE_ORACLE_LOG_FILENAME` | - | Log file path |

### External Signing (Recommended for Production)

| Environment Variable | Description |
|---------------------|-------------|
| `TOKEN_PRICE_ORACLE_EXTERNAL_SIGN` | Enable external signing (`true`/`false`) |
| `TOKEN_PRICE_ORACLE_EXTERNAL_SIGN_ADDRESS` | Signing account address |
| `TOKEN_PRICE_ORACLE_EXTERNAL_SIGN_APPID` | External signing service AppID |
| `TOKEN_PRICE_ORACLE_EXTERNAL_SIGN_CHAIN` | Chain identifier |
| `TOKEN_PRICE_ORACLE_EXTERNAL_SIGN_URL` | External signing service URL |
| `TOKEN_PRICE_ORACLE_EXTERNAL_SIGN_RSA_PRIV` | RSA private key (PEM format) |

## Price Calculation

### Price Ratio Formula

```
priceRatio = tokenScale × tokenPriceUSD × 10^(18 - tokenDecimals) / ethPriceUSD
```

### Threshold

Threshold is specified in basis points (bps):
- 1 bps = 0.01%
- 100 bps = 1%
- 10000 bps = 100%

On-chain prices are only updated when price change exceeds the threshold, avoiding unnecessary Gas costs.

## Monitoring

### Prometheus Metrics

When metrics server is enabled, access `http://<host>:<port>/metrics`:

| Metric | Type | Description |
|--------|------|-------------|
| `last_successful_update_timestamp` | Gauge | Last successful update timestamp |
| `updates_total{type="updated"}` | Counter | Actual update count |
| `updates_total{type="skipped"}` | Counter | Skipped update count |
| `update_errors_total{type="price"}` | Counter | Update error count |
| `account_balance_eth` | Gauge | Oracle account balance |

### Health Check

```bash
curl http://<host>:<port>/health
```

### Suggested Alert Rules

```yaml
# Price not updated for a long time
- alert: TokenPriceOracleStalled
  expr: time() - last_successful_update_timestamp > 300
  for: 1m
  labels:
    severity: critical

# Low account balance
- alert: TokenPriceOracleLowBalance
  expr: account_balance_eth < 0.1
  for: 5m
  labels:
    severity: warning
```

## Project Structure

```
token-price-oracle/
├── cmd/              # Entry point
├── flags/            # CLI flags definition
├── config/           # Configuration loading
├── client/           # Client wrappers
│   ├── l2_client.go  # L2 chain client
│   ├── price_feed.go # Price feed interface
│   ├── bitget_sdk.go # Bitget API client
│   └── sign.go       # External signing
├── updater/          # Update logic
│   ├── token_price.go # Price updater
│   ├── tx_manager.go  # Transaction manager
│   └── factory.go     # Factory methods
├── metrics/          # Prometheus metrics
└── README.md         # This document
```

## Development

```bash
# Run tests
make test

# Test Bitget price feed (requires network)
go test ./client -run TestBitgetPriceFeed -v

# Format code
go fmt ./...

# Local run
cp env.example .env
# Edit .env configuration
source .env && make run
```

## License

MIT
