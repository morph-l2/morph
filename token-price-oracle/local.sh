#!/bin/bash

# Local development startup script for token-price-oracle

./build/bin/token-price-oracle \
  --l2-eth-rpc http://localhost:8545 \
  --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
  --price-update-interval 30s \
  --price-threshold 5 \
  --price-feed-priority bitget \
  --token-ids "1,2" \
  --token-mapping-bitget "1:BGBUSDT,2:BTCUSDT" \
  --log-level info \
  --metrics-server-enable

