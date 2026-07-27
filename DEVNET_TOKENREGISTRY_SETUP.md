# Devnet TokenRegistry Setup

Developer genesis pre-registers three tokens in `L2TokenRegistry`. Registration
alone does not make the tokens immediately updateable: the contract owner must
activate them and allow the oracle signer before starting `token-price-oracle`.

## Pre-registered tokens

| Token ID | Symbol | Address | Decimals | Scale | Purpose |
|----------|--------|---------|----------|-------|---------|
| 1 | BTC | `0x0000000000000000000000000000000000000001` | 8 | 10^10 | High-value asset feed testing |
| 2 | ETH | `0x5300000000000000000000000000000000000011` | 18 | 1 | L2WETH benchmark |
| 3 | BGB | `0x0000000000000000000000000000000000000003` | 18 | 1 | CEX-specific feed testing |

BTC and BGB use mock addresses and are not deployed ERC-20 contracts.

## Storage initialization

`SetDevnetTestTokens` runs after the system contract implementations are
installed by `BuildL2DeveloperGenesis`. It initializes:

- `tokenRegistry[tokenID]`
- `tokenRegistration[tokenAddress]`
- `supportedTokenSet`

Tokens are deliberately initialized with `isActive=false`. A non-zero balance
slot is stored as `balanceSlot + 1`, matching the encoding used by
`L2TokenRegistry`; zero remains zero.

## Start and verify the devnet

```bash
make devnet-down
rm -rf ops/docker/.devnet
make devnet-up
```

Verify the registered IDs with the `getSupportedIDList()` selector or a contract
binding. The expected decoded result is `[1, 2, 3]`.

## Activate tokens and allow the oracle

Perform both owner operations before starting the oracle:

```bash
REGISTRY=0x5300000000000000000000000000000000000021
RPC_URL=http://localhost:8545
OWNER_PRIVATE_KEY="<DEVNET_OWNER_PRIVATE_KEY>"
ORACLE_ADDRESS="<ORACLE_SIGNER_ADDRESS>"

cast send "$REGISTRY" \
  "batchUpdateTokenStatus(uint16[],bool[])" \
  "[1,2,3]" "[true,true,true]" \
  --rpc-url "$RPC_URL" \
  --private-key "$OWNER_PRIVATE_KEY"

cast send "$REGISTRY" \
  "setAllowList(address[],bool[])" \
  "[$ORACLE_ADDRESS]" "[true]" \
  --rpc-url "$RPC_URL" \
  --private-key "$OWNER_PRIVATE_KEY"
```

The private key must belong only to an isolated local devnet account. Never use
it on a public network or commit it to the repository.

## Configure token-price-oracle

```bash
export TOKEN_PRICE_ORACLE_L2_ETH_RPC=http://localhost:8545
export TOKEN_PRICE_ORACLE_L2_TOKEN_REGISTRY_ADDRESS=0x5300000000000000000000000000000000000021
export TOKEN_PRICE_ORACLE_PRIVATE_KEY="<DEVNET_ORACLE_PRIVATE_KEY>"
export TOKEN_PRICE_ORACLE_TOKEN_IDS=1,2,3
export TOKEN_PRICE_ORACLE_PRICE_UPDATE_INTERVAL=30s
export TOKEN_PRICE_ORACLE_PRICE_THRESHOLD=100
export TOKEN_PRICE_ORACLE_PRICE_FEED_PRIORITY=chainlink,pyth,bitget,okx

export TOKEN_PRICE_ORACLE_TOKEN_MAPPING_BITGET="1:BTCUSDT,2:ETHUSDT,3:BGBUSDT"
export TOKEN_PRICE_ORACLE_BITGET_API_BASE_URL=https://api.bitget.com
export TOKEN_PRICE_ORACLE_TOKEN_MAPPING_OKX="1:BTC-USDT,2:ETH-USDT,3:BGB-USDT"
export TOKEN_PRICE_ORACLE_OKX_API_BASE_URL=https://www.okx.com

export TOKEN_PRICE_ORACLE_CHAINLINK_RPC=https://ethereum-rpc.publicnode.com
export TOKEN_PRICE_ORACLE_TOKEN_MAPPING_CHAINLINK="1:0xF4030086522a5bEEa4988F8cA5B36dbC97BeE88c,2:0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419"
export TOKEN_PRICE_ORACLE_CHAINLINK_ETH_USD_FEED=0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419
export TOKEN_PRICE_ORACLE_CHAINLINK_MAX_STALENESS=1h

export TOKEN_PRICE_ORACLE_PYTH_HERMES_BASE_URL=https://hermes.pyth.network
export TOKEN_PRICE_ORACLE_TOKEN_MAPPING_PYTH="1:0xe62df6c8b4a85fe1a67db44dc12de5db330f7ac66b72dc658afedf0f4a415b43,2:0xff61491a931112ddf1bd8147cd1b641375f79f5825126d665480874634fd0ace"
export TOKEN_PRICE_ORACLE_PYTH_ETH_USD_PRICE_ID=0xff61491a931112ddf1bd8147cd1b641375f79f5825126d665480874634fd0ace
export TOKEN_PRICE_ORACLE_PYTH_MAX_STALENESS=1m
export TOKEN_PRICE_ORACLE_PYTH_MAX_CONFIDENCE_BPS=500

export TOKEN_PRICE_ORACLE_METRICS_SERVER_ENABLE=true
export TOKEN_PRICE_ORACLE_METRICS_PORT=6060
```

Use an isolated devnet-only oracle private key. For production, use the external
signing mode described in `token-price-oracle/README.md`.

## Start the oracle

Run the binary directly:

```bash
cd token-price-oracle
./build/bin/token-price-oracle
```

Or run the container with an explicit name:

```bash
docker run -d \
  --name token-price-oracle \
  --network docker_default \
  --env-file devnet.env \
  morph/token-price-oracle:latest

docker logs -f token-price-oracle
```

## Verification checklist

These checks require a freshly generated devnet and are not implied by unit
tests:

- [ ] `getSupportedIDList()` returns `[1, 2, 3]`.
- [ ] `getTokenInfo()` returns the expected address, balance slot, decimals,
      scale, and inactive initial status.
- [ ] The owner activates token IDs 1, 2, and 3.
- [ ] The owner adds the oracle signer to the allowlist.
- [ ] The oracle fetches all configured prices.
- [ ] `batchUpdatePrices` succeeds on-chain.
- [ ] The metrics endpoint reports the successful update.

## Troubleshooting

- `No tokens to update`: regenerate the devnet genesis and verify
  `getSupportedIDList()`.
- `CallerNotAllowed`: add the oracle signer to the allowlist.
- Inactive-token errors: call `batchUpdateTokenStatus` before starting the
  oracle.
- All feeds fail: verify network access, endpoints, mappings, and API keys.
