#!/bin/bash
# 低频发送交易脚本 - 保证有空块也有非空块

set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
GETH_BIN="${SCRIPT_DIR}/bin/geth"
ZK_GETH_HTTP="http://127.0.0.1:8545"

# 颜色
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

log_info() { echo -e "${BLUE}[INFO]${NC} $1"; }
log_success() { echo -e "${GREEN}[TX]${NC} $1"; }

# 测试账户 (genesis 中需要有余额)
# 这是一个常用的测试私钥，对应地址: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
PRIVATE_KEY="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
FROM_ADDRESS="0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
TO_ADDRESS="0x70997970C51812dc3A010C7d01b50e0d17dc79C8"

# 发送间隔 (秒) - 每秒发送一笔
MIN_INTERVAL=1
MAX_INTERVAL=1

# 检查账户余额
check_balance() {
    local balance=$(curl -s -X POST -H "Content-Type: application/json" \
        --data "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getBalance\",\"params\":[\"${FROM_ADDRESS}\", \"latest\"],\"id\":1}" \
        "$ZK_GETH_HTTP" | grep -o '"result":"[^"]*"' | cut -d'"' -f4)
    echo "$balance"
}

# Get nonce
get_nonce() {
    local nonce=$(curl -s -X POST -H "Content-Type: application/json" \
        --data "{\"jsonrpc\":\"2.0\",\"method\":\"eth_getTransactionCount\",\"params\":[\"${FROM_ADDRESS}\", \"latest\"],\"id\":1}" \
        "$ZK_GETH_HTTP" | grep -o '"result":"[^"]*"' | cut -d'"' -f4)
    echo "$nonce"
}

# Get current block
get_block() {
    local block=$(curl -s -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
        "$ZK_GETH_HTTP" | grep -o '"result":"[^"]*"' | cut -d'"' -f4)
    echo "$block"
}

# 发送简单转账交易
send_tx() {
    local nonce=$1
    local value="0x1"  # 1 wei
    
    # 构造并发送交易 (使用 personal_sendTransaction 或签名后发送)
    # 这里使用简单的 eth_sendTransaction (需要 geth 支持)
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
        "$ZK_GETH_HTTP" 2>/dev/null)
    
    echo "$result"
}

# 使用 cast 发送交易 (如果安装了 foundry)
send_tx_with_cast() {
    if command -v cast &> /dev/null; then
        cast send --private-key "$PRIVATE_KEY" \
            --rpc-url "$ZK_GETH_HTTP" \
            "$TO_ADDRESS" \
            --value 1wei \
            --gas-price 0 \
            2>/dev/null
        return $?
    fi
    return 1
}

# 主循环
main() {
    log_info "Starting low-frequency transaction sender..."
    log_info "From: $FROM_ADDRESS"
    log_info "To:   $TO_ADDRESS"
    log_info "Interval: ${MIN_INTERVAL}-${MAX_INTERVAL} seconds"
    echo ""
    
    # 检查余额
    local balance=$(check_balance)
    log_info "Account balance: $balance"
    
    if [ "$balance" == "0x0" ] || [ -z "$balance" ]; then
        echo ""
        echo "Warning: Account balance is 0 or cannot be retrieved!"
        echo "请确保 genesis 文件中为以下地址分配了余额："
        echo "  $FROM_ADDRESS"
        echo ""
        echo "或者修改脚本中的 FROM_ADDRESS 和 PRIVATE_KEY"
        exit 1
    fi
    
    local tx_count=0
    
    while true; do
        sleep 1
        
        local block=$(get_block)
        local nonce=$(get_nonce)
        
        log_info "Sending tx... block=$block nonce=$nonce"
        
        # 尝试发送交易
        if command -v cast &> /dev/null; then
            # 使用 foundry cast (legacy 交易格式, async 不等待确认)
            # 使用 1 gwei gas price
            local result=$(cast send --private-key "$PRIVATE_KEY" \
                --rpc-url "$ZK_GETH_HTTP" \
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

# 显示帮助
show_help() {
    echo "低频交易发送脚本"
    echo ""
    echo "用法: $0 [start|stop|status]"
    echo ""
    echo "命令:"
    echo "  start   - 开始发送交易 (前台运行)"
    echo "  bg      - 后台运行"
    echo "  stop    - 停止后台进程"
    echo "  status  - 查看状态"
    echo ""
    echo "配置:"
    echo "  发送间隔: ${MIN_INTERVAL}-${MAX_INTERVAL} 秒"
    echo "  发送账户: $FROM_ADDRESS"
    echo ""
    echo "注意: 需要安装 foundry (cast) 或确保 geth 开启了 personal API"
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

