#!/bin/bash

# Local development startup script for token-price-oracle

./build/bin/token-price-oracle \
  --l2-eth-rpc http://localhost:8545 \
  --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
  --price-update-interval 30s \
  --price-threshold 100 \
  --price-feed-priority bitget \
  --token-mapping-bitget "1:BGBUSDT,2:BTCUSDT,3:\$1.0" \
  --bitget-api-base-url https://api.bitget.com \
  --log-level info \
  --metrics-server-enable

# Price threshold examples (in basis points):
# 1 bps = 0.01%, 10 bps = 0.1%, 100 bps = 1%, 500 bps = 5%, 1000 bps = 10%

# Token mapping format:
#   - Regular tokens: tokenID:SYMBOL (e.g., 1:BGBUSDT, 2:BTCUSDT)
#   - Stablecoins:    tokenID:$PRICE (e.g., 3:$1.0 for USDT pegged to $1 USD)
# Note: Use \$ in bash to escape the dollar sign

# Chainlink example:
#   --price-feed-priority chainlink,pyth,bitget,binance,okx \
#   --chainlink-rpc https://ethereum-rpc.publicnode.com \
#   --chainlink-eth-usd-feed 0x... \
#   --chainlink-max-staleness 1h \
#   --token-mapping-chainlink "1:0x...,2:0x..." \
#   --pyth-hermes-base-url https://hermes.pyth.network \
#   --pyth-api-key "$PYTH_API_KEY" \
#   --pyth-eth-usd-price-id 0xff61491a931112ddf1bd8147cd1b641375f79f5825126d665480874634fd0ace \
#   --pyth-max-staleness 1h \
#   --pyth-max-confidence-bps 0 \
#   --token-mapping-pyth "1:0x...,2:0x..." \
#   --token-mapping-binance "1:BTCUSDT,2:ETHUSDT" \
#   --token-mapping-okx "1:BTC-USDT,2:ETH-USDT" \

