#!/bin/bash

# Check block status and stateRoot of both Geth nodes

ZK_GETH="http://127.0.0.1:8545"
MPT_GETH="http://127.0.0.1:9545"

echo "=========================================="
echo "       Node Status Check"
echo "=========================================="
echo ""

# Function to get block info
get_block_info() {
    local name=$1
    local url=$2
    
    echo "=== ${name} (${url}) ==="
    
    # Get latest block number
    local block_number=$(curl -s -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
        "$url" 2>/dev/null | jq -r '.result // "error"')
    
    if [ "$block_number" = "error" ] || [ -z "$block_number" ]; then
        echo "  Status: NOT RUNNING"
        echo ""
        return
    fi
    
    # Convert to decimal
    local block_dec=$((block_number))
    echo "  Block Number: ${block_number} (${block_dec})"
    
    # Get latest block details
    local block_info=$(curl -s -X POST -H "Content-Type: application/json" \
        --data "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBlockByNumber\",\"params\":[\"${block_number}\", false],\"id\":1}" \
        "$url" 2>/dev/null)
    
    local state_root=$(echo "$block_info" | jq -r '.result.stateRoot // "N/A"')
    local timestamp=$(echo "$block_info" | jq -r '.result.timestamp // "N/A"')
    local tx_count=$(echo "$block_info" | jq -r '.result.transactions | length // 0')
    
    # Convert timestamp
    if [ "$timestamp" != "N/A" ]; then
        local ts_dec=$((timestamp))
        local ts_readable=$(date -r $ts_dec 2>/dev/null || date -d @$ts_dec 2>/dev/null || echo "$ts_dec")
        echo "  Timestamp: ${ts_readable}"
    fi
    
    echo "  StateRoot: ${state_root}"
    echo "  Tx Count: ${tx_count}"
    echo ""
}

# Check both nodes
get_block_info "ZK Geth" "$ZK_GETH"
get_block_info "MPT Geth" "$MPT_GETH"

# Compare stateRoot
echo "=========================================="
echo "       StateRoot Comparison"
echo "=========================================="

zk_block=$(curl -s -X POST -H "Content-Type: application/json" \
    --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
    "$ZK_GETH" 2>/dev/null | jq -r '.result // ""')

mpt_block=$(curl -s -X POST -H "Content-Type: application/json" \
    --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
    "$MPT_GETH" 2>/dev/null | jq -r '.result // ""')

if [ -n "$zk_block" ] && [ -n "$mpt_block" ]; then
    # Use the smaller block number for comparison
    zk_dec=$((zk_block))
    mpt_dec=$((mpt_block))
    
    if [ $zk_dec -le $mpt_dec ]; then
        compare_block=$zk_block
    else
        compare_block=$mpt_block
    fi
    
    echo "Comparing at block: ${compare_block} ($((compare_block)))"
    echo ""
    
    zk_state=$(curl -s -X POST -H "Content-Type: application/json" \
        --data "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBlockByNumber\",\"params\":[\"${compare_block}\", false],\"id\":1}" \
        "$ZK_GETH" 2>/dev/null | jq -r '.result.stateRoot // "N/A"')
    
    mpt_state=$(curl -s -X POST -H "Content-Type: application/json" \
        --data "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBlockByNumber\",\"params\":[\"${compare_block}\", false],\"id\":1}" \
        "$MPT_GETH" 2>/dev/null | jq -r '.result.stateRoot // "N/A"')
    
    echo "ZK Geth  StateRoot: ${zk_state}"
    echo "MPT Geth StateRoot: ${mpt_state}"
    echo ""
    
    if [ "$zk_state" = "$mpt_state" ]; then
        echo "Result: SAME (Both nodes have identical stateRoot)"
    else
        echo "Result: DIFFERENT (Nodes have different stateRoot - expected for ZK vs MPT)"
    fi
else
    echo "Cannot compare - one or both nodes not running"
fi

echo ""


