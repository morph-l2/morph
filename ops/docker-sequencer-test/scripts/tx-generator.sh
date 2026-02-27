#!/bin/sh
# Transaction Generator for Sequencer Test
# Sends random transactions to keep the network active

set -e

L2_RPC="${L2_RPC:-http://morph-geth-0:8545}"
INTERVAL="${TX_INTERVAL:-5}"  # seconds between txs
PRIVATE_KEY="${PRIVATE_KEY:-0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80}"

# Wait for L2 to be ready
echo "Waiting for L2 RPC to be ready..."
while true; do
    if curl -s -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
        "$L2_RPC" | grep -q "result"; then
        echo "L2 RPC is ready!"
        break
    fi
    echo "Waiting..."
    sleep 2
done

# Get initial nonce
get_nonce() {
    curl -s -X POST -H "Content-Type: application/json" \
        --data "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getTransactionCount\",\"params\":[\"0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266\",\"latest\"],\"id\":1}" \
        "$L2_RPC" | grep -o '"result":"[^"]*"' | cut -d'"' -f4
}

# Get current block number
get_block_number() {
    curl -s -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
        "$L2_RPC" | grep -o '"result":"[^"]*"' | cut -d'"' -f4
}

echo "Starting transaction generator..."
echo "L2 RPC: $L2_RPC"
echo "Interval: ${INTERVAL}s"

NONCE_HEX=$(get_nonce)
NONCE=$((NONCE_HEX))
TX_COUNT=0

while true; do
    BLOCK=$(get_block_number)
    BLOCK_DEC=$((BLOCK))
    
    # Generate random recipient address
    RANDOM_SUFFIX=$(od -An -N4 -tx1 /dev/urandom | tr -d ' ')
    TO_ADDR="0x000000000000000000000000000000${RANDOM_SUFFIX}"
    
    # Create and send transaction
    NONCE_HEX=$(printf "0x%x" $NONCE)
    TX_DATA="{\"jsonrpc\":\"2.0\",\"method\":\"eth_sendTransaction\",\"params\":[{\"from\":\"0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266\",\"to\":\"${TO_ADDR}\",\"value\":\"0x1\",\"nonce\":\"${NONCE_HEX}\"}],\"id\":1}"
    
    RESULT=$(curl -s -X POST -H "Content-Type: application/json" --data "$TX_DATA" "$L2_RPC")
    
    if echo "$RESULT" | grep -q "result"; then
        TX_HASH=$(echo "$RESULT" | grep -o '"result":"[^"]*"' | cut -d'"' -f4)
        echo "[Block $BLOCK_DEC] TX #$TX_COUNT sent: $TX_HASH"
        NONCE=$((NONCE + 1))
        TX_COUNT=$((TX_COUNT + 1))
    else
        echo "[Block $BLOCK_DEC] TX failed: $RESULT"
        # Refresh nonce in case of error
        NONCE_HEX=$(get_nonce)
        NONCE=$((NONCE_HEX))
    fi
    
    sleep "$INTERVAL"
done

