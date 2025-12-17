#!/bin/bash
# Send a real transaction to the mempool
# Uses Hardhat test account #1 which has funds in genesis
#
# Usage: ./send-tx.sh [rpc_url] [to_address] [value_in_ether]
# Example: ./send-tx.sh http://127.0.0.1:8545 0xdead 1.5

set -e

cd "$(dirname "$0")"

# Default values
RPC="${1:-http://127.0.0.1:8545}"
TO="${2:-0x000000000000000000000000000000000000dEaD}"
VALUE="${3:-1}"

cd /Users/xx/workspace/go-ethereum
go run ./local-test/send-tx.go "$RPC" "$TO" "$VALUE"

