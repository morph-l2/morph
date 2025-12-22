#!/bin/bash
set -e

# =============================================================================
# MPT Switch Local Test Script
# 
# Architecture:
# - Sequencer Node: Initially connects to ZK Geth, switches to MPT Geth after upgrade
# - Sentry Node: Initially connects to MPT Geth, switches to ZK Geth after upgrade
#
# Before upgrade:
#   Sequencer Node ──► ZK Geth (:8545)
#   Sentry Node ────► MPT Geth (:9545)
#
# After upgrade (swap):
#   Sequencer Node ──► MPT Geth (:9545)
#   Sentry Node ────► ZK Geth (:8545)
#
# Test flow:
# 1. Sequencer produces blocks using ZK Geth, Sentry syncs from Sequencer
# 2. After mptTime, both nodes swap their Geth connections
# =============================================================================

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

# 路径配置
MORPH_ROOT="${SCRIPT_DIR}/../.."
BIN_DIR="${SCRIPT_DIR}/bin"

# 所有二进制文件都放在 bin 目录下
GETH_BIN="${BIN_DIR}/geth"
NODE_BIN="${BIN_DIR}/morphnode"
TENDERMINT_BIN="${BIN_DIR}/tendermint"

# 测试数据目录
TEST_DATA_DIR="${SCRIPT_DIR}/.testdata"
ZK_GETH_DIR="${TEST_DATA_DIR}/zk-geth"
MPT_GETH_DIR="${TEST_DATA_DIR}/mpt-geth"
SEQUENCER_NODE_DIR="${TEST_DATA_DIR}/sequencer-node"
SENTRY_NODE_DIR="${TEST_DATA_DIR}/sentry-node"

# 端口配置
# ZK Geth (used by Sequencer before upgrade, by Sentry after upgrade)
ZK_GETH_HTTP_PORT=8545
ZK_GETH_WS_PORT=8546
ZK_GETH_AUTH_PORT=8551
ZK_GETH_P2P_PORT=30303

# MPT Geth (used by Sentry before upgrade, by Sequencer after upgrade)
MPT_GETH_HTTP_PORT=9545
MPT_GETH_WS_PORT=9546
MPT_GETH_AUTH_PORT=9551
MPT_GETH_P2P_PORT=30304

# Sequencer Node
SEQ_NODE_P2P_PORT=26656
SEQ_NODE_RPC_PORT=26657

# Sentry Node
SENTRY_NODE_P2P_PORT=26756
SENTRY_NODE_RPC_PORT=26757

# PID 文件
ZK_GETH_PID="${TEST_DATA_DIR}/zk-geth.pid"
MPT_GETH_PID="${TEST_DATA_DIR}/mpt-geth.pid"
SEQUENCER_NODE_PID="${TEST_DATA_DIR}/sequencer-node.pid"
SENTRY_NODE_PID="${TEST_DATA_DIR}/sentry-node.pid"

# 颜色输出
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 检查依赖
check_dependencies() {
    log_info "Checking dependencies in ${BIN_DIR}..."
    
    local missing=0
    
    if [ ! -f "$GETH_BIN" ]; then
        log_error "Missing: bin/geth"
        missing=1
    fi
    
    if [ ! -f "$NODE_BIN" ]; then
        log_error "Missing: bin/morphnode"
        missing=1
    fi
    
    if [ ! -f "$TENDERMINT_BIN" ]; then
        log_error "Missing: bin/tendermint"
        missing=1
    fi
    
    if [ $missing -eq 1 ]; then
        echo ""
        log_info "Please copy the required binaries to: ${BIN_DIR}/"
        exit 1
    fi
    
    log_success "All binaries found"
}

# 准备配置文件
prepare_configs() {
    log_info "Preparing configuration..."
    
    # 创建数据目录
    mkdir -p "$ZK_GETH_DIR"
    mkdir -p "$MPT_GETH_DIR"
    mkdir -p "$SEQUENCER_NODE_DIR"
    mkdir -p "$SENTRY_NODE_DIR"
    
    # 检查 genesis 文件 (两个模式需要不同的 genesis)
    local genesis_zk="${SCRIPT_DIR}/genesis-zk.json"
    local genesis_mpt="${SCRIPT_DIR}/genesis-mpt.json"
    
    if [ ! -f "$genesis_zk" ]; then
        log_error "genesis-zk.json not found at: ${genesis_zk}"
        exit 1
    fi
    
    if [ ! -f "$genesis_mpt" ]; then
        log_error "genesis-mpt.json not found at: ${genesis_mpt}"
        exit 1
    fi
    
    # 创建 JWT secret
    local jwt_file="${TEST_DATA_DIR}/jwt-secret.txt"
    if [ ! -f "$jwt_file" ]; then
        openssl rand -hex 32 > "$jwt_file"
        log_success "Generated JWT secret"
    fi
    
    # 初始化 geth 数据目录 (两个模式使用不同的 genesis)
    if [ ! -d "${ZK_GETH_DIR}/geth/chaindata" ]; then
        log_info "Initializing zk-geth with genesis-zk.json (useZktrie: true)..."
        "$GETH_BIN" init --datadir="$ZK_GETH_DIR" "$genesis_zk"
    fi
    
    if [ ! -d "${MPT_GETH_DIR}/geth/chaindata" ]; then
        # MPT Geth initialized with genesis-mpt.json (useZktrie: false)
        # But --genesisStateRoot is passed at startup to keep same initial state as ZK Geth
        log_info "Initializing mpt-geth with genesis-mpt.json (useZktrie: false)..."
        "$GETH_BIN" init --datadir="$MPT_GETH_DIR" "$genesis_mpt"
    fi
    
    # Initialize tendermint configs 
    # - Sequencer: validator (has voting power, produces blocks)
    # - Sentry: non-validator (no voting power, only syncs blocks)
    if [ ! -d "${SEQUENCER_NODE_DIR}/config" ]; then
        log_info "Initializing tendermint configs (1 validator + 1 non-validator)..."
        
        local temp_dir="${TEST_DATA_DIR}/tendermint-temp"
        
        # 创建 1 个 validator (sequencer) + 1 个 non-validator (sentry)
        "$TENDERMINT_BIN" testnet --v 1 --n 1 --o "$temp_dir" --populate-persistent-peers --hostname-prefix node-
        
        # node0 是 validator (sequencer), node1 是 non-validator (sentry)
        # 确保目标目录不存在，这样 mv 会重命名而不是移动到目录内
        rm -rf "${SEQUENCER_NODE_DIR}" "${SENTRY_NODE_DIR}"
        mv "${temp_dir}/node0" "${SEQUENCER_NODE_DIR}"
        mv "${temp_dir}/node1" "${SENTRY_NODE_DIR}"
        
        # 重要: sentry 必须使用与 sequencer 相同的 genesis.json (包含相同的 validator set)
        cp "${SEQUENCER_NODE_DIR}/config/genesis.json" "${SENTRY_NODE_DIR}/config/genesis.json"
        
        rm -rf "$temp_dir"
        
        # 修改 node 配置
        # 注意: testnet 命令生成的目录结构是 node0, node1
        #       但我们重命名后直接放在 SEQUENCER_NODE_DIR 和 SENTRY_NODE_DIR 下
        local seq_config="${SEQUENCER_NODE_DIR}/config/config.toml"
        local sentry_config="${SENTRY_NODE_DIR}/config/config.toml"
        
        # 获取 sequencer 的 node ID 用于 sentry 连接
        local seq_node_id=$("$TENDERMINT_BIN" show-node-id --home "${SEQUENCER_NODE_DIR}")
        
        if [[ "$OSTYPE" == "darwin"* ]]; then
            # Sequencer 配置
            sed -i '' 's#create_empty_blocks_interval = "0s"#create_empty_blocks_interval = "5s"#g' "$seq_config"
            sed -i '' 's#prometheus = false#prometheus = true#g' "$seq_config"
            sed -i '' "s#laddr = \"tcp://0.0.0.0:26656\"#laddr = \"tcp://0.0.0.0:${SEQ_NODE_P2P_PORT}\"#g" "$seq_config"
            sed -i '' "s#laddr = \"tcp://127.0.0.1:26657\"#laddr = \"tcp://127.0.0.1:${SEQ_NODE_RPC_PORT}\"#g" "$seq_config"
            
            # Sentry 配置 (连接到 sequencer)
            sed -i '' 's#create_empty_blocks_interval = "0s"#create_empty_blocks_interval = "5s"#g' "$sentry_config"
            sed -i '' 's#prometheus = false#prometheus = true#g' "$sentry_config"
            sed -i '' "s#laddr = \"tcp://0.0.0.0:26656\"#laddr = \"tcp://0.0.0.0:${SENTRY_NODE_P2P_PORT}\"#g" "$sentry_config"
            sed -i '' "s#laddr = \"tcp://127.0.0.1:26657\"#laddr = \"tcp://127.0.0.1:${SENTRY_NODE_RPC_PORT}\"#g" "$sentry_config"
            # 设置 persistent_peers 指向 sequencer (替换任何现有值)
            sed -i '' "s#persistent_peers = \".*\"#persistent_peers = \"${seq_node_id}@127.0.0.1:${SEQ_NODE_P2P_PORT}\"#" "$sentry_config"
        else
            # Linux - Sequencer 配置
            sed -i 's#create_empty_blocks_interval = "0s"#create_empty_blocks_interval = "5s"#g' "$seq_config"
            sed -i 's#prometheus = false#prometheus = true#g' "$seq_config"
            sed -i "s#laddr = \"tcp://0.0.0.0:26656\"#laddr = \"tcp://0.0.0.0:${SEQ_NODE_P2P_PORT}\"#g" "$seq_config"
            sed -i "s#laddr = \"tcp://127.0.0.1:26657\"#laddr = \"tcp://127.0.0.1:${SEQ_NODE_RPC_PORT}\"#g" "$seq_config"
            
            # Linux - Sentry 配置
            sed -i 's#create_empty_blocks_interval = "0s"#create_empty_blocks_interval = "5s"#g' "$sentry_config"
            sed -i 's#prometheus = false#prometheus = true#g' "$sentry_config"
            sed -i "s#laddr = \"tcp://0.0.0.0:26656\"#laddr = \"tcp://0.0.0.0:${SENTRY_NODE_P2P_PORT}\"#g" "$sentry_config"
            sed -i "s#laddr = \"tcp://127.0.0.1:26657\"#laddr = \"tcp://127.0.0.1:${SENTRY_NODE_RPC_PORT}\"#g" "$sentry_config"
            sed -i "s#persistent_peers = \".*\"#persistent_peers = \"${seq_node_id}@127.0.0.1:${SEQ_NODE_P2P_PORT}\"#" "$sentry_config"
        fi
        
        log_success "Tendermint configs initialized"
    fi
    
    log_success "Configuration ready"
}

# Start Geth generic function
start_geth() {
    local name=$1
    local geth_bin=$2
    local datadir=$3
    local http_port=$4
    local ws_port=$5
    local auth_port=$6
    local p2p_port=$7
    local pid_file=$8
    local extra_args=${9:-""}
    
    log_info "Starting ${name} on ports HTTP:${http_port} AUTH:${auth_port} P2P:${p2p_port}..."
    log_info "  Binary: $(basename $geth_bin)"
    
    local jwt_file="${TEST_DATA_DIR}/jwt-secret.txt"
    
    "$geth_bin" \
        --datadir="$datadir" \
        --networkid=53077 \
        --http \
        --http.addr=0.0.0.0 \
        --http.port=$http_port \
        --http.api=web3,debug,eth,txpool,net,morph,engine,admin \
        --http.corsdomain="*" \
        --http.vhosts="*" \
        --ws \
        --ws.addr=0.0.0.0 \
        --ws.port=$ws_port \
        --ws.api=web3,debug,eth,txpool,net,morph,engine,admin \
        --ws.origins="*" \
        --authrpc.addr=0.0.0.0 \
        --authrpc.port=$auth_port \
        --authrpc.vhosts="*" \
        --authrpc.jwtsecret="$jwt_file" \
        --port=$p2p_port \
        --nodiscover \
        --gcmode=archive \
        --miner.gasprice=0 \
        --verbosity=3 \
        $extra_args \
        > "${TEST_DATA_DIR}/${name}.log" 2>&1 &
    
    echo $! > "$pid_file"
    log_success "${name} started (PID: $(cat $pid_file))"
}

# Start Sequencer Node
# Before upgrade: connects to ZK Geth
# After upgrade: switches to MPT Geth
start_sequencer_node() {
    local mpt_time=$1
    
    log_info "Starting sequencer-node with MPT_TIME=${mpt_time}..."
    log_info "  Before upgrade: ZK Geth (:${ZK_GETH_HTTP_PORT})"
    log_info "  After upgrade:  MPT Geth (:${MPT_GETH_HTTP_PORT})"
    
    local jwt_file="${TEST_DATA_DIR}/jwt-secret.txt"
    local node_home="${SEQUENCER_NODE_DIR}"
    
    # Sequencer: legacy=ZK Geth, mpt=MPT Geth
    export MORPH_NODE_L2_LEGACY_ETH_RPC="http://127.0.0.1:${ZK_GETH_HTTP_PORT}"
    export MORPH_NODE_L2_LEGACY_ENGINE_RPC="http://127.0.0.1:${ZK_GETH_AUTH_PORT}"
    export MORPH_NODE_L2_ETH_RPC="http://127.0.0.1:${MPT_GETH_HTTP_PORT}"
    export MORPH_NODE_L2_ENGINE_RPC="http://127.0.0.1:${MPT_GETH_AUTH_PORT}"
    export MORPH_NODE_L2_ENGINE_AUTH="$jwt_file"
    export MORPH_NODE_MPT_TIME="$mpt_time"
    export MORPH_NODE_L1_ETH_RPC="${L1_ETH_RPC:-http://127.0.0.1:9545}"
    export MORPH_NODE_SYNC_DEPOSIT_CONTRACT_ADDRESS="0x6900000000000000000000000000000000000001"
    export MORPH_NODE_L1_CONFIRMATIONS=0
    export MORPH_NODE_LOG_LEVEL=debug
    
    "$NODE_BIN" \
        --dev-sequencer \
        --home "$node_home" \
        > "${TEST_DATA_DIR}/sequencer-node.log" 2>&1 &
    
    echo $! > "$SEQUENCER_NODE_PID"
    log_success "sequencer-node started (PID: $(cat $SEQUENCER_NODE_PID))"
}

# Start Sentry Node (non-validator, follower node)
# Before upgrade: connects to MPT Geth (syncs ZK blocks from sequencer)
# After upgrade: switches to ZK Geth (swaps with Sequencer)
# Note: Sentry has no block production rights, only syncs blocks from sequencer
start_sentry_node() {
    local mpt_time=$1
    
    log_info "Starting sentry-node with MPT_TIME=${mpt_time}..."
    log_info "  Before upgrade: MPT Geth (:${MPT_GETH_HTTP_PORT}) - syncs ZK blocks from sequencer"
    log_info "  After upgrade:  ZK Geth (:${ZK_GETH_HTTP_PORT}) - swaps with Sequencer"
    
    local jwt_file="${TEST_DATA_DIR}/jwt-secret.txt"
    local node_home="${SENTRY_NODE_DIR}"
    
    # Sentry: legacy=MPT Geth, mpt=ZK Geth (opposite to Sequencer, swaps on upgrade)
    export MORPH_NODE_L2_LEGACY_ETH_RPC="http://127.0.0.1:${MPT_GETH_HTTP_PORT}"
    export MORPH_NODE_L2_LEGACY_ENGINE_RPC="http://127.0.0.1:${MPT_GETH_AUTH_PORT}"
    export MORPH_NODE_L2_ETH_RPC="http://127.0.0.1:${ZK_GETH_HTTP_PORT}"
    export MORPH_NODE_L2_ENGINE_RPC="http://127.0.0.1:${ZK_GETH_AUTH_PORT}"
    export MORPH_NODE_L2_ENGINE_AUTH="$jwt_file"
    export MORPH_NODE_MPT_TIME="$mpt_time"
    export MORPH_NODE_L1_ETH_RPC="${L1_ETH_RPC:-http://127.0.0.1:9545}"
    export MORPH_NODE_SYNC_DEPOSIT_CONTRACT_ADDRESS="0x6900000000000000000000000000000000000001"
    export MORPH_NODE_L1_CONFIRMATIONS=0
    export MORPH_NODE_LOG_LEVEL=debug
    
    "$NODE_BIN" \
        --dev-sequencer \
        --home "$node_home" \
        > "${TEST_DATA_DIR}/sentry-node.log" 2>&1 &
    
    echo $! > "$SENTRY_NODE_PID"
    log_success "sentry-node started (PID: $(cat $SENTRY_NODE_PID))"
}

# 等待 Geth 就绪
wait_for_geth() {
    local port=$1
    local name=$2
    local max_wait=30
    local waited=0
    
    log_info "Waiting for ${name} to be ready..."
    
    while ! curl -s "http://127.0.0.1:${port}" > /dev/null 2>&1; do
        sleep 1
        waited=$((waited + 1))
        if [ $waited -ge $max_wait ]; then
            log_error "Timeout waiting for ${name}"
            return 1
        fi
    done
    
    log_success "${name} is ready"
}

# 停止所有服务
stop_all() {
    log_info "Stopping all services..."
    
    for pid_file in "$SENTRY_NODE_PID" "$SEQUENCER_NODE_PID" "$MPT_GETH_PID" "$ZK_GETH_PID"; do
        if [ -f "$pid_file" ]; then
            local pid=$(cat "$pid_file")
            local name=$(basename "$pid_file" .pid)
            if kill -0 "$pid" 2>/dev/null; then
                kill "$pid" 2>/dev/null || true
                log_info "Stopped ${name} (PID: $pid)"
            fi
            rm -f "$pid_file"
        fi
    done
    
    log_success "All services stopped"
}

# 清理数据
clean_data() {
    log_info "Cleaning test data..."
    rm -rf "$TEST_DATA_DIR"
    rm -f "${SCRIPT_DIR}/genesis-l2.json"
    log_success "Test data cleaned"
}

# Monitor logs
monitor_logs() {
    local target=${1:-all}
    
    log_info "Monitoring logs for MPT switch events..."
    log_info "Looking for:"
    echo "  - 'MPT switch time reached, MUST wait for MPT node to sync'"
    echo "  - 'Waiting for MPT node to sync...'"
    echo "  - 'Successfully switched to MPT client'"
    echo ""
    log_info "Press Ctrl+C to stop"
    echo ""
    
    case $target in
        sequencer)
            tail -f "${TEST_DATA_DIR}/sequencer-node.log" 2>/dev/null | grep --line-buffered -E "(MPT|switch|mpt|block)" || true
            ;;
        sentry)
            tail -f "${TEST_DATA_DIR}/sentry-node.log" 2>/dev/null | grep --line-buffered -E "(MPT|switch|mpt|block|error|Error)" || true
            ;;
        *)
            tail -f "${TEST_DATA_DIR}/sequencer-node.log" "${TEST_DATA_DIR}/sentry-node.log" 2>/dev/null | grep --line-buffered -E "(MPT|switch|mpt)" || true
            ;;
    esac
}

# View logs
view_logs() {
    local service=${1:-sequencer}
    
    case $service in
        sequencer|seq)
            tail -f "${TEST_DATA_DIR}/sequencer-node.log"
            ;;
        sentry)
            tail -f "${TEST_DATA_DIR}/sentry-node.log"
            ;;
        zk|zk-geth)
            tail -f "${TEST_DATA_DIR}/zk-geth.log"
            ;;
        mpt|mpt-geth)
            tail -f "${TEST_DATA_DIR}/mpt-geth.log"
            ;;
        all)
            tail -f "${TEST_DATA_DIR}"/*.log
            ;;
        *)
            log_error "Unknown service: $service"
            echo "Available: sequencer, sentry, zk-geth, mpt-geth, all"
            ;;
    esac
}

# 检查状态
check_status() {
    echo ""
    log_info "=== Service Status ==="
    
    # 进程状态
    for pid_file in "$ZK_GETH_PID" "$MPT_GETH_PID" "$SEQUENCER_NODE_PID" "$SENTRY_NODE_PID"; do
        local name=$(basename "$pid_file" .pid)
        if [ -f "$pid_file" ] && kill -0 "$(cat $pid_file)" 2>/dev/null; then
            printf "%-20s ${GREEN}Running${NC} (PID: %s)\n" "$name:" "$(cat $pid_file)"
        else
            printf "%-20s ${RED}Stopped${NC}\n" "$name:"
        fi
    done
    
    echo ""
    log_info "=== Block Heights ==="
    
    # ZK Geth
    local zk_block=$(curl -s -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
        "http://127.0.0.1:${ZK_GETH_HTTP_PORT}" 2>/dev/null | grep -o '"result":"[^"]*"' | cut -d'"' -f4)
    printf "%-20s %s\n" "zk-geth:" "${zk_block:-N/A}"
    
    # MPT Geth
    local mpt_block=$(curl -s -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
        "http://127.0.0.1:${MPT_GETH_HTTP_PORT}" 2>/dev/null | grep -o '"result":"[^"]*"' | cut -d'"' -f4)
    printf "%-20s %s\n" "mpt-geth:" "${mpt_block:-N/A}"
    
    echo ""
    log_info "=== Ports ==="
    echo "zk-geth:        HTTP=${ZK_GETH_HTTP_PORT}  AUTH=${ZK_GETH_AUTH_PORT}"
    echo "mpt-geth:       HTTP=${MPT_GETH_HTTP_PORT}  AUTH=${MPT_GETH_AUTH_PORT}"
    echo "sequencer-node: P2P=${SEQ_NODE_P2P_PORT}  RPC=${SEQ_NODE_RPC_PORT}"
    echo "sentry-node:    P2P=${SENTRY_NODE_P2P_PORT}  RPC=${SENTRY_NODE_RPC_PORT}"
    echo ""
}

# 显示帮助
show_help() {
    echo "MPT Switch Local Test Script"
    echo ""
    echo "Architecture:"
    echo "  Before upgrade:"
    echo "    Sequencer Node ──► ZK Geth  (:8545)"
    echo "    Sentry Node ────► MPT Geth (:9545)"
    echo "  After upgrade (swap):"
    echo "    Sequencer Node ──► MPT Geth (:9545)"
    echo "    Sentry Node ────► ZK Geth  (:8545)"
    echo ""
    echo "Required binaries (place in bin/ directory):"
    echo "  bin/geth"
    echo "  bin/morphnode"
    echo "  bin/tendermint"
    echo ""
    echo "Usage: $0 <command> [options]"
    echo ""
    echo "Commands:"
    echo "  start [delay]      Start test environment, delay is MPT switch delay in seconds (default 60)"
    echo "  stop               Stop all services"
    echo "  clean              Clean all test data"
    echo "  status             View service status and block height"
    echo "  monitor [target]   Monitor logs (sequencer/sentry/all)"
    echo "  logs [service]     View logs (sequencer/sentry/legacy-geth/mpt-geth/sentry-geth/all)"
    echo "  help               Show this help"
    echo ""
    echo "Examples:"
    echo "  $0 start 30           # Start, trigger switch after 30 seconds"
    echo "  $0 monitor sequencer  # Monitor sequencer switch events"
    echo "  $0 monitor sentry     # Monitor sentry node"
    echo "  $0 status             # View status"
    echo "  $0 stop               # Stop services"
}

# Main function
main() {
    local command=${1:-help}
    
    case $command in
        start)
            local mpt_delay=${2:-60}
            local mpt_time=$(($(date +%s) + mpt_delay))
            local mpt_time_readable=$(date -r $mpt_time 2>/dev/null || date -d @$mpt_time 2>/dev/null)
            
            echo ""
            log_info "=========================================="
            log_info "       MPT Switch Test Environment"
            log_info "=========================================="
            echo ""
            log_info "MPT switch time: ${mpt_time_readable}"
            log_info "That's ${mpt_delay} seconds from now"
            echo ""
            log_info "Architecture:"
            echo "  Sequencer: validator (block producer)"
            echo "  Sentry:    non-validator (follower, syncs blocks from sequencer)"
            echo ""
            echo "  Before upgrade:"
            echo "    Sequencer ──► ZK Geth  (:${ZK_GETH_HTTP_PORT}) - produces blocks"
            echo "    Sentry ────► MPT Geth (:${MPT_GETH_HTTP_PORT}) - syncs ZK blocks"
            echo "  After upgrade (swap):"
            echo "    Sequencer ──► MPT Geth (:${MPT_GETH_HTTP_PORT})"
            echo "    Sentry ────► ZK Geth  (:${ZK_GETH_HTTP_PORT})"
            echo ""
            
            check_dependencies
            prepare_configs
            
            # Start Geth (only 2 instances)
            # MPT Geth's genesis-mpt.json is configured with genesisStateRoot matching ZK Geth
            start_geth "zk-geth" "$GETH_BIN" "$ZK_GETH_DIR" $ZK_GETH_HTTP_PORT $ZK_GETH_WS_PORT $ZK_GETH_AUTH_PORT $ZK_GETH_P2P_PORT "$ZK_GETH_PID"
            start_geth "mpt-geth" "$GETH_BIN" "$MPT_GETH_DIR" $MPT_GETH_HTTP_PORT $MPT_GETH_WS_PORT $MPT_GETH_AUTH_PORT $MPT_GETH_P2P_PORT "$MPT_GETH_PID"
            
            # Wait for Geth to be ready
            wait_for_geth $ZK_GETH_HTTP_PORT "zk-geth"
            wait_for_geth $MPT_GETH_HTTP_PORT "mpt-geth"
            
            sleep 2
            
            # Start Nodes (both configured with the same mpt_time)
            start_sequencer_node $mpt_time
            sleep 2
            start_sentry_node $mpt_time
            
            echo ""
            log_success "=========================================="
            log_success "    Test environment is ready!"
            log_success "=========================================="
            echo ""
            log_info "Commands:"
            echo "  $0 monitor sequencer  - Watch sequencer MPT switch"
            echo "  $0 monitor sentry     - Watch sentry node (will stop at upgrade)"
            echo "  $0 status             - Check service status"
            echo "  $0 stop               - Stop all services"
            echo ""
            ;;
        stop)
            stop_all
            ;;
        clean)
            stop_all
            clean_data
            ;;
        status)
            check_status
            ;;
        monitor)
            monitor_logs ${2:-all}
            ;;
        logs)
            view_logs ${2:-sequencer}
            ;;
        help|*)
            show_help
            ;;
    esac
}

main "$@"
