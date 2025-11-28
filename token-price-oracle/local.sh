#!/bin/bash

# Local development startup script for token-price-oracle

./build/bin/token-price-oracle \
  --l2-eth-rpc http://localhost:8545 \
  --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
  --price-update-interval 30s \
  --price-threshold 100 \
  --price-feed-priority bitget \
  --token-mapping-bitget "1:BGBUSDT,2:BTCUSDT" \
  --bitget-api-base-url https://api.bitget.com \
  --log-level info \
  --metrics-server-enable

# Price threshold examples (in basis points):
# 1 bps = 0.01%, 10 bps = 0.1%, 100 bps = 1%, 500 bps = 5%, 1000 bps = 10%

