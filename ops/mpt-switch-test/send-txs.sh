#!/bin/bash
# Low-frequency transaction sending script - ensures both empty and non-empty blocks

set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
GETH_BIN="${SCRIPT_DIR}/bin/geth"

# RPC endpoint: default 8545, pass "1" to use 9545
if [ "${USE_MPT:-0}" == "1" ]; then
    GETH_HTTP="http://127.0.0.1:9545"
else
    GETH_HTTP="http://127.0.0.1:8545"
fi

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

log_info() { echo -e "${BLUE}[INFO]${NC} $1"; }
log_success() { echo -e "${GREEN}[TX]${NC} $1"; }

# Test account (needs balance in genesis)
# This is a commonly used test private key, corresponding address: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
PRIVATE_KEY="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
FROM_ADDRESS="0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
TO_ADDRESS="0x70997970C51812dc3A010C7d01b50e0d17dc79C8"

# Send interval (seconds) - one transaction per second
MIN_INTERVAL=1
MAX_INTERVAL=1

# Check account balance
check_balance() {
    local balance=$(curl -s -X POST -H "Content-Type: application/json" \
        --data "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBalance\",\"params\":[\"${FROM_ADDRESS}\", \"latest\"],\"id\":1}" \
        "$GETH_HTTP" | grep -o '"result":"[^"]*"' | cut -d'"' -f4)
    echo "$balance"
}

# Get nonce
get_nonce() {
    local nonce=$(curl -s -X POST -H "Content-Type: application/json" \
        --data "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getTransactionCount\",\"params\":[\"${FROM_ADDRESS}\", \"latest\"],\"id\":1}" \
        "$GETH_HTTP" | grep -o '"result":"[^"]*"' | cut -d'"' -f4)
    echo "$nonce"
}

# Get current block
get_block() {
    local block=$(curl -s -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
        "$GETH_HTTP" | grep -o '"result":"[^"]*"' | cut -d'"' -f4)
    echo "$block"
}

# Send simple transfer transaction
send_tx() {
    local nonce=$1
    local value="0x1"  # 1 wei
    
    # Construct and send transaction (using personal_sendTransaction or send after signing)
    # Using simple eth_sendTransaction here (requires geth support)
    local result=$(curl -s -X POST -H "Content-Type: application/json" \
        --data "{
            \"jsonrpc\":\"2.0\",
            \"method\":\"eth_sendTransaction\",
            \"params\":[{
                \"from\": \"${FROM_ADDRESS}\",
                \"to\": \"${TO_ADDRESS}\",
                \"value\": \"${value}\",
                \"gas\": \"0x5208\",
                \"gasPrice\": \"0x0\"
            }],
            \"id\":1
        }" \
        "$GETH_HTTP" 2>/dev/null)
    
    echo "$result"
}

# Send transaction using cast (if foundry is installed)
send_tx_with_cast() {
    if command -v cast &> /dev/null; then
        cast send --private-key "$PRIVATE_KEY" \
            --rpc-url "$GETH_HTTP" \
            "$TO_ADDRESS" \
            --value 1wei \
            --gas-price 0 \
            2>/dev/null
        return $?
    fi
    return 1
}

# Main loop
main() {
    log_info "Starting low-frequency transaction sender..."
    log_info "From: $FROM_ADDRESS"
    log_info "To:   $TO_ADDRESS"
    log_info "RPC:  $GETH_HTTP"
    log_info "Interval: ${MIN_INTERVAL}-${MAX_INTERVAL} seconds"
    echo ""
    
    # Check balance
    local balance=$(check_balance)
    log_info "Account balance: $balance"
    
    if [ "$balance" == "0x0" ] || [ -z "$balance" ]; then
        echo ""
        echo "Warning: Account balance is 0 or cannot be retrieved!"
        echo "Please ensure the genesis file allocates balance to the following address:"
        echo "  $FROM_ADDRESS"
        echo ""
        echo "Or modify FROM_ADDRESS and PRIVATE_KEY in the script"
        exit 1
    fi
    
    local tx_count=0
    
    while true; do
        sleep 1
        
        local block=$(get_block)
        local nonce=$(get_nonce)
        
        log_info "Sending tx... block=$block nonce=$nonce"
        
        # Try to send transaction
        if command -v cast &> /dev/null; then
            # Using foundry cast (legacy transaction format, async without waiting for confirmation)
            # Using 1 gwei gas price
            local result=$(cast send --private-key "$PRIVATE_KEY" \
                --rpc-url "$GETH_HTTP" \
                "$TO_ADDRESS" \
                --value 1wei \
                --legacy \
                --gas-price 1000000000 \
                --async \
                2>&1)
            
            if [[ "$result" == 0x* ]]; then
                tx_count=$((tx_count + 1))
                log_success "#${tx_count} TxHash: ${result}"
            else
                log_info "Failed: ${result:0:80}"
            fi
        else
            log_info "cast not found, please install foundry"
            exit 1
        fi
    done
}

# Show help
show_help() {
    echo "Low-frequency transaction sending script"
    echo ""
    echo "Usage: [USE_MPT=1] $0 [start|stop|status]"
    echo ""
    echo "Commands:"
    echo "  start   - Start sending transactions (foreground)"
    echo "  bg      - Run in background"
    echo "  stop    - Stop background process"
    echo "  status  - Check status"
    echo ""
    echo "Environment Variables:"
    echo "  USE_MPT=1  - Use MPT Geth (9545), default is ZK Geth (8545)"
    echo ""
    echo "Examples:"
    echo "  $0 start           # Send to 8545 (before switch)"
    echo "  USE_MPT=1 $0 start # Send to 9545 (after switch)"
    echo ""
    echo "Note: Requires foundry (cast)"
}

PID_FILE="${SCRIPT_DIR}/.testdata/send-txs.pid"

case "${1:-start}" in
    start)
        main
        ;;
    bg)
        log_info "Starting in background..."
        nohup "$0" start > "${SCRIPT_DIR}/.testdata/send-txs.log" 2>&1 &
        echo $! > "$PID_FILE"
        log_success "Started (PID: $(cat $PID_FILE))"
        ;;
    stop)
        if [ -f "$PID_FILE" ]; then
            kill $(cat "$PID_FILE") 2>/dev/null && rm "$PID_FILE"
            log_success "Stopped"
        else
            log_info "Not running"
        fi
        ;;
    status)
        if [ -f "$PID_FILE" ] && kill -0 $(cat "$PID_FILE") 2>/dev/null; then
            log_success "Running (PID: $(cat $PID_FILE))"
        else
            log_info "Not running"
        fi
        ;;
    -h|--help|help)
        show_help
        ;;
    *)
        show_help
        exit 1
        ;;
esac
