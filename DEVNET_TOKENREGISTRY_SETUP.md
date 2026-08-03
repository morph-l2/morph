# Devnet TokenRegistry Setup

Developer genesis pre-registers three tokens in `L2TokenRegistry`. Registration
alone does not make the tokens immediately updateable: the contract owner must
activate them and allow the oracle signer before starting `token-price-oracle`.

## Pre-registered tokens

| Token ID | Symbol | Address | Decimals | Scale | Purpose |
|----------|--------|---------|----------|-------|---------|
| 1 | BTC | `0x1111111111111111111111111111111111111111` | 8 | 10^8 | High-value asset feed testing |
| 2 | ETH | `0x5300000000000000000000000000000000000011` | 18 | 10^18 | L2WETH benchmark |
| 3 | BGB | `0x3333333333333333333333333333333333333333` | 18 | 10^18 | CEX-specific feed testing |

BTC and BGB use placeholder addresses and are not deployed ERC-20 contracts.
They are kept clear of the `0x01`-`0x0a` precompile range so that a call to
`balanceOf` cannot silently dispatch to a precompile.

Scale is `10^decimals`, the same convention `L2TokenRegistry`'s tests use for
USDC and DAI. It is not the decimals adjustment: the oracle already multiplies by
`10^(18-decimals)` on its own, and scale cancels out of `calculateTokenAmount`
entirely. What it does control is how much of `priceRatio` survives truncation to
a uint256, and `10^decimals` is what makes the stored ratio equal
`10^18 * tokenPrice/ethPrice` for every token regardless of its decimals. Setting
it to `10^(18-decimals)` instead yields a scale of 1 for an 18-decimal token,
which truncates the ratio of anything cheaper than ETH to zero and leaves the
token permanently unpriceable.

## Storage initialization

`SetDevnetTestTokens` runs after the system contract implementations are
installed by `BuildL2DeveloperGenesis`, and only when `fundDevAccounts` is set,
so non-developer genesis is unaffected. It initializes:

- `tokenRegistry[tokenID]`
- `tokenRegistration[tokenAddress]`
- `supportedTokenSet`

Tokens are deliberately initialized with `isActive=false`. Balance slots follow
the `L2TokenRegistry` encoding: a token that needs one stores `balanceSlot + 1`
and a token that does not stores zero. The `NeedBalanceSlot` flag carries that
distinction, because slot 0 is a real balance slot (it is where `WrappedEther`
keeps `_balances`) and cannot be represented by a zero value alone.

## Start and verify the devnet

```bash
make devnet-down
rm -rf ops/docker/.devnet
make devnet-up
```

This is enough to pick up a new L2 genesis, which is regenerated on every run.
It does not reset L1: `make devnet-down` removes the containers and the network
but not the named volumes, so `layer1-el-data` survives and the L1 chain resumes
where it left off. The L1 contracts are then redeployed from the same account at
higher nonces, so every L1 address in `ops/docker/.env` shifts between runs. That
is expected — those are plain `CREATE` addresses derived from the deployer nonce,
unlike the L2 predeploys, which are fixed by genesis. To start L1 from scratch as
well:

```bash
cd ops/docker && docker compose -f docker-compose-devnet.yml down --volumes
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
export TOKEN_PRICE_ORACLE_PRIVATE_KEY="<DEVNET_ORACLE_PRIVATE_KEY>"
export TOKEN_PRICE_ORACLE_PRICE_UPDATE_INTERVAL=30s
export TOKEN_PRICE_ORACLE_PRICE_THRESHOLD=100
export TOKEN_PRICE_ORACLE_PRICE_FEED_PRIORITY=chainlink,pyth,bitget,okx

export TOKEN_PRICE_ORACLE_TOKEN_MAPPING_BITGET="1:BTCUSDT,2:ETHUSDT,3:BGBUSDT"
export TOKEN_PRICE_ORACLE_BITGET_API_BASE_URL=https://api.bitget.com
export TOKEN_PRICE_ORACLE_TOKEN_MAPPING_OKX="1:BTC-USDT,2:ETH-USDT"
export TOKEN_PRICE_ORACLE_OKX_API_BASE_URL=https://www.okx.com

export TOKEN_PRICE_ORACLE_CHAINLINK_RPC=https://ethereum-rpc.publicnode.com
export TOKEN_PRICE_ORACLE_TOKEN_MAPPING_CHAINLINK="1:0xF4030086522a5bEEa4988F8cA5B36dbC97BeE88c,2:0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419"
export TOKEN_PRICE_ORACLE_CHAINLINK_ETH_USD_FEED=0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419
export TOKEN_PRICE_ORACLE_CHAINLINK_MAX_STALENESS=1h

export TOKEN_PRICE_ORACLE_PYTH_HERMES_BASE_URL=https://hermes.pyth.network
export TOKEN_PRICE_ORACLE_PYTH_API_KEY="<PYTH_API_KEY>"
export TOKEN_PRICE_ORACLE_TOKEN_MAPPING_PYTH="1:0xe62df6c8b4a85fe1a67db44dc12de5db330f7ac66b72dc658afedf0f4a415b43,2:0xff61491a931112ddf1bd8147cd1b641375f79f5825126d665480874634fd0ace"
export TOKEN_PRICE_ORACLE_PYTH_ETH_USD_PRICE_ID=0xff61491a931112ddf1bd8147cd1b641375f79f5825126d665480874634fd0ace
export TOKEN_PRICE_ORACLE_PYTH_MAX_STALENESS=1m
export TOKEN_PRICE_ORACLE_PYTH_MAX_CONFIDENCE_BPS=500

export TOKEN_PRICE_ORACLE_METRICS_SERVER_ENABLE=true
export TOKEN_PRICE_ORACLE_METRICS_PORT=6060
```

BGB is only listed on Bitget, so token 3 is deliberately absent from the Chainlink,
Pyth, and OKX mappings. Feeds omit tokens they cannot map rather than failing the
batch, so the higher-priority feeds resolve tokens 1 and 2 and Bitget resolves
token 3.

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
- [ ] `priceRatio()` is non-zero for all three token IDs, and the oracle logs no
      `Skipping zero price`. A ratio that truncates to zero is dropped with only a
      warning, so a scale mistake surfaces here rather than as a failed update.
- [ ] The metrics endpoint reports the successful update and `unresolved_tokens`
      is 0.

## Troubleshooting

- `No tokens to update`: regenerate the devnet genesis and verify
  `getSupportedIDList()`.
- `CallerNotAllowed`: add the oracle signer to the allowlist.
- Inactive-token errors: call `batchUpdateTokenStatus` before starting the
  oracle.
- All feeds fail: verify network access, endpoints, mappings, and API keys.
