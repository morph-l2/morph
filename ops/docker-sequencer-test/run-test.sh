#!/bin/bash
# Sequencer Upgrade Test Runner
# Reuses devnet-morph logic but with test-specific docker images

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MORPH_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"
BITGET_ROOT="$(cd "$MORPH_ROOT/.." && pwd)"
OPS_DIR="$MORPH_ROOT/ops"
DOCKER_DIR="$OPS_DIR/docker"
DEVNET_DIR="$OPS_DIR/devnet-morph"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

log_info() { echo -e "${BLUE}[INFO]${NC} $1"; }
log_success() { echo -e "${GREEN}[SUCCESS]${NC} $1"; }
log_warn() { echo -e "${YELLOW}[WARN]${NC} $1"; }
log_error() { echo -e "${RED}[ERROR]${NC} $1"; }

# Configuration
UPGRADE_HEIGHT=${UPGRADE_HEIGHT:-10}
L2_RPC="http://127.0.0.1:8545"
L2_RPC_NODE1="http://127.0.0.1:8645"

# ========== Helper Functions ==========

wait_for_rpc() {
    local rpc_url="$1"
    local max_retries=${2:-60}
    local retry=0
    
    log_info "Waiting for RPC at $rpc_url..."
    while [ $retry -lt $max_retries ]; do
        if curl -s -X POST -H "Content-Type: application/json" \
            --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
            "$rpc_url" 2>/dev/null | grep -q "result"; then
            log_success "RPC is ready!"
            return 0
        fi
        retry=$((retry + 1))
        sleep 2
    done
    log_error "Timeout waiting for RPC"
    return 1
}

get_block_number() {
    local rpc_url="${1:-$L2_RPC}"
    local result
    result=$(curl -s -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
        "$rpc_url" 2>/dev/null)
    echo "$result" | grep -o '"result":"[^"]*"' | cut -d'"' -f4 | xargs printf "%d" 2>/dev/null || echo "0"
}

wait_for_block() {
    local target_height=$1
    local rpc_url="${2:-$L2_RPC}"
    
    log_info "Waiting for block $target_height..."
    while true; do
        local current=$(get_block_number "$rpc_url")
        if [ "$current" -ge "$target_height" ]; then
            log_success "Reached block $current"
            return 0
        fi
        echo -ne "\r  Current block: $current / $target_height"
        sleep 2
    done
}

# ========== Setup Functions ==========

# Export consensus switch height as environment variable for Docker containers
# The morphnode binary reads MORPH_NODE_CONSENSUS_SWITCH_HEIGHT at runtime
set_upgrade_height() {
    local height=$1
    log_info "Setting consensus switch height to $height (via CONSENSUS_SWITCH_HEIGHT env)..."
    export CONSENSUS_SWITCH_HEIGHT="$height"
    log_success "CONSENSUS_SWITCH_HEIGHT=$height (will be passed to containers)"
}

# Build test images (with -test suffix)
# Uses bitget/ as build context to access local go-ethereum and tendermint
build_test_images() {
    log_info "Building test Docker images..."
    log_info "Using build context: $BITGET_ROOT"
    
    # Build go-ubuntu-builder if needed
    cd "$MORPH_ROOT"
    make go-ubuntu-builder
    
    # Build from bitget/ directory to access all repos
    cd "$BITGET_ROOT"
    
    # # Copy go module cache to avoid network downloads
    # if [ -d "$HOME/go/pkg/mod" ]; then
    #     log_info "Copying go module cache to build context..."
    #     rm -rf .gomodcache
    #     cp -r "$HOME/go/pkg/mod" .gomodcache
    # else
    #     log_warn "Go module cache not found at $HOME/go/pkg/mod"
    #     log_warn "Build may fail due to network issues"
    # fi
    
    # Build test geth image
    log_info "Building morph-geth-test (using local go-ethereum)..."
    docker build -t morph-geth-test:latest \
        -f morph/ops/docker-sequencer-test/Dockerfile.l2-geth-test .
    
    # Build test node image
    log_info "Building morph-node-test (using local go-ethereum + tendermint)..."
    docker build -t morph-node-test:latest \
        -f morph/ops/docker-sequencer-test/Dockerfile.l2-node-test .
    
    # # Cleanup go module cache copy
    # rm -rf .gomodcache
    
    log_success "Test images built!"
}

# Run full devnet setup (reusing existing logic, but skip L2 startup)
setup_devnet() {
    log_info "Running devnet setup..."
    cd "$MORPH_ROOT"
    
    # Note: upgrade height should already be set before build_test_images
    
    # Step 1: Start L1 and setup tendermint nodes
    # Note: main.py calls setup_devnet_nodes() before devnet.main()
    log_info "Step 1: Starting L1 and setting up tendermint nodes..."
    python3 "$DEVNET_DIR/main.py" --polyrepo-dir="$MORPH_ROOT" --only-l1
    
    # Step 2: Deploy contracts and generate L2 genesis (without starting L2)
    log_info "Step 2: Deploying contracts and generating L2 genesis..."
    python3 -c "
import sys
import os
import time
import re
import fileinput
sys.path.insert(0, '$DEVNET_DIR')
import devnet
from devnet import run_command, read_json, write_json, test_port, log

pjoin = os.path.join
polyrepo_dir = '$MORPH_ROOT'
L2_dir = pjoin(polyrepo_dir, 'ops', 'l2-genesis')
devnet_dir = pjoin(polyrepo_dir, 'ops', 'l2-genesis', '.devnet')
ops_dir = pjoin(polyrepo_dir, 'ops', 'docker')
contracts_dir = pjoin(polyrepo_dir, 'contracts')

os.makedirs(devnet_dir, exist_ok=True)

# Generate network config
devnet_cfg_orig = pjoin(L2_dir, 'deploy-config', 'devnet-deploy-config.json')
deploy_config = read_json(devnet_cfg_orig)
deploy_config['l1GenesisBlockTimestamp'] = '0x{:x}'.format(int(time.time()))
deploy_config['l1StartingBlockTag'] = 'earliest'
temp_deploy_config = pjoin(devnet_dir, 'deploy-config.json')
write_json(temp_deploy_config, deploy_config)

# Deploy L1 contracts
deployment_dir = pjoin(devnet_dir, 'devnetL1.json')
run_command(['rm', '-f', deployment_dir], env={}, cwd=contracts_dir)
log.info('Deploying L1 Proxy contracts...')
run_command(['yarn', 'build'], env={}, cwd=contracts_dir)
run_command(['npx', 'hardhat', 'deploy', '--network', 'l1', '--storagepath', deployment_dir, '--concurrent', 'true'], env={}, cwd=contracts_dir)

# Generate L2 genesis
log.info('Generating L2 genesis and rollup configs...')
run_command([
    'env', 'CGO_ENABLED=1', 'CGO_LDFLAGS=\"-ldl\"',
    'go', 'run', 'cmd/main.go', 'genesis', 'l2',
    '--l1-rpc', 'http://localhost:9545',
    '--deploy-config', temp_deploy_config,
    '--deployment-dir', deployment_dir,
    '--outfile.l2', pjoin(devnet_dir, 'genesis-l2.json'),
    '--outfile.genbatchheader', pjoin(devnet_dir, 'genesis-batch-header.json'),
    '--outfile.rollup', pjoin(devnet_dir, 'rollup.json')
], cwd=L2_dir)

# Initialize contracts
log.info('Deploying L1 Impl contracts and initialize...')
rollup_cfg = read_json(pjoin(devnet_dir, 'rollup.json'))
genesis_batch_header = rollup_cfg['genesis_batch_header']
contracts_config = pjoin(contracts_dir, 'src', 'deploy-config', 'l1.ts')
pattern3 = re.compile(\"batchHeader: '.*'\")
for line in fileinput.input(contracts_config, inplace=True):
    modified_line = re.sub(pattern3, f\"batchHeader: '{genesis_batch_header}'\", line)
    print(modified_line, end='')
run_command(['npx', 'hardhat', 'initialize', '--network', 'l1', '--storagepath', deployment_dir, '--concurrent', 'true'], env={}, cwd=contracts_dir)

# Staking
log.info('Staking sequencers...')
addresses = {}
deployment = read_json(deployment_dir)
for d in deployment:
    addresses[d['name']] = d['address']
for i in range(4):
    run_command(['cast', 'send', addresses['Proxy__L1Staking'],
                 'register(bytes32,bytes memory)',
                 deploy_config['l2StakingTmKeys'][i],
                 deploy_config['l2StakingBlsKeys'][i],
                 '--rpc-url', 'http://127.0.0.1:9545',
                 '--value', '1ether',
                 '--private-key', deploy_config['l2StakingPks'][i]
                 ])

# Update .env file
log.info('Updating .env file...')
env_file = pjoin(ops_dir, '.env')
env_data = {}
with open(env_file, 'r+') as envfile:
    env_content = envfile.readlines()
    for line in env_content:
        line = line.strip()
        if line and not line.startswith('#'):
            parts = line.split('=', 1)
            if len(parts) == 2:
                env_data[parts[0].strip()] = parts[1].strip()
    env_data['L1_CROSS_DOMAIN_MESSENGER'] = addresses['Proxy__L1CrossDomainMessenger']
    env_data['MORPH_PORTAL'] = addresses['Proxy__L1MessageQueueWithGasPriceOracle']
    env_data['MORPH_ROLLUP'] = addresses['Proxy__Rollup']
    env_data['MORPH_L1STAKING'] = addresses['Proxy__L1Staking']
    env_data['L1_SEQUENCER_CONTRACT'] = addresses.get('Proxy__L1Sequencer', '')
    envfile.seek(0)
    for key, value in env_data.items():
        envfile.write(f'{key}={value}\n')
    envfile.truncate()

log.info('Contract deployment and genesis generation complete!')
log.info('Skipping L2 startup - will be done with test images.')
"
    
    log_success "Devnet setup complete (L2 not started yet)"
}

# Docker compose command with override file
# Note: -f must explicitly include override file when using non-default compose file name
COMPOSE_CMD="docker compose -f docker-compose-4nodes.yml -f docker-compose.override.yml"
COMPOSE_CMD_NO_OVERRIDE="docker compose -f docker-compose-4nodes.yml"

# Copy override file to use test images
setup_override() {
    log_info "Setting up docker-compose override for test images..."
    cp "$SCRIPT_DIR/docker-compose.override.yml" "$DOCKER_DIR/docker-compose.override.yml"
    log_success "Override file copied to $DOCKER_DIR/"
}

# Remove override file
remove_override() {
    rm -f "$DOCKER_DIR/docker-compose.override.yml"
}

# Start L2 with test images
start_l2_test() {
    log_info "Starting L2 with test images..."
    cd "$DOCKER_DIR"
    
    # Setup override file
    setup_override
    
    # Read the .env file to get contract addresses
    source .env 2>/dev/null || true
    
    # Set sequencer private key
    export SEQUENCER_PRIVATE_KEY="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
    
    # Stop any existing L2 containers
    $COMPOSE_CMD stop \
        morph-geth-0 morph-geth-1 morph-geth-2 morph-geth-3 \
        node-0 node-1 node-2 node-3 2>/dev/null || true
    
    # Note: Test images should already be built by build_test_images()
    # Uncomment below if you need to rebuild during start
    # log_info "Building L2 containers with test images..."
    # $COMPOSE_CMD build morph-geth-0 node-0
    
    # Start L2 geth nodes
    log_info "Starting L2 geth nodes..."
    $COMPOSE_CMD up -d morph-geth-0 morph-geth-1 morph-geth-2 morph-geth-3
    
    sleep 5
    
    # Start L2 tendermint nodes
    log_info "Starting L2 tendermint nodes..."
    $COMPOSE_CMD up -d node-0 node-1 node-2 node-3
    
    wait_for_rpc "$L2_RPC"
    log_success "L2 is running with test images!"
}

# ========== Test Functions ==========

test_pbft_mode() {
    log_info "========== Phase 1: Testing PBFT Mode =========="
    
    local start_block=$(get_block_number)
    log_info "Starting block: $start_block"
    
    # Wait for some blocks
    local target=$((start_block + 10))
    wait_for_block $target
    
    # Verify nodes in sync
    local block0=$(get_block_number "$L2_RPC")
    local block1=$(get_block_number "$L2_RPC_NODE1")
    
    local diff=$((block0 - block1))
    if [ ${diff#-} -le 2 ]; then
        log_success "Nodes in sync (node0: $block0, node1: $block1)"
    else
        log_error "Nodes out of sync!"
        return 1
    fi
}

test_upgrade() {
    log_info "========== Phase 2: Waiting for Upgrade =========="
    log_info "Upgrade height: $UPGRADE_HEIGHT"
    
    wait_for_block $UPGRADE_HEIGHT
    sleep 10
    
    # Verify network continues
    local post_upgrade=$(get_block_number)
    wait_for_block $((post_upgrade + 5))
    
    log_success "Upgrade completed! Network continues producing blocks."
}

test_sequencer_mode() {
    log_info "========== Phase 3: Testing Sequencer Mode =========="
    
    local start_block=$(get_block_number)
    wait_for_block $((start_block + 20))
    
    local block0=$(get_block_number "$L2_RPC")
    local block1=$(get_block_number "$L2_RPC_NODE1")
    
    local diff=$((block0 - block1))
    if [ ${diff#-} -le 2 ]; then
        log_success "Nodes in sync after upgrade (node0: $block0, node1: $block1)"
    else
        log_error "Nodes out of sync after upgrade!"
        return 1
    fi
}

test_fullnode_sync() {
    log_info "========== Phase 4: Testing Fullnode Sync =========="
    
    local current_height=$(get_block_number)
    log_info "Current height: $current_height"
    
    cd "$DOCKER_DIR"
    
    # Start sentry node (fullnode)
    log_info "Starting fullnode (sentry-node-0)..."
    $COMPOSE_CMD up -d sentry-geth-0 sentry-node-0
    
    sleep 10
    wait_for_rpc "http://127.0.0.1:8945"
    
    # Wait for sync
    local target_sync=$((current_height - 5))
    local max_wait=300
    local waited=0
    
    while [ $waited -lt $max_wait ]; do
        local fn_block=$(get_block_number "http://127.0.0.1:8945")
        if [ "$fn_block" -ge "$target_sync" ]; then
            log_success "Fullnode synced to block $fn_block"
            return 0
        fi
        echo -ne "\r  Fullnode: $fn_block / $target_sync"
        sleep 5
        waited=$((waited + 5))
    done
    
    log_error "Fullnode sync timeout"
    return 1
}

# ========== Transaction Generator ==========

start_tx_generator() {
    log_info "Starting transaction generator..."
    
    # Simple tx generator using cast
    (
        while true; do
            RANDOM_ADDR="0x$(openssl rand -hex 20)"
            cast send --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 \
                --rpc-url "$L2_RPC" \
                --value 1wei \
                "$RANDOM_ADDR" 2>/dev/null || true
            sleep ${TX_INTERVAL:-5}
        done
    ) &
    TX_GEN_PID=$!
    log_info "TX generator started (PID: $TX_GEN_PID)"
}

stop_tx_generator() {
    if [ -n "$TX_GEN_PID" ]; then
        kill $TX_GEN_PID 2>/dev/null || true
        log_info "TX generator stopped"
    fi
}

# ========== Cleanup ==========

cleanup() {
    log_info "Cleaning up..."
    stop_tx_generator
    cd "$DOCKER_DIR"
    $COMPOSE_CMD_NO_OVERRIDE down -v 2>/dev/null || true
    remove_override
}

# ========== Main Commands ==========

run_full_test() {
    log_info "=========================================="
    log_info "  Sequencer Upgrade Test"
    log_info "  Upgrade Height: $UPGRADE_HEIGHT"
    log_info "=========================================="
    
    trap cleanup EXIT
    
    # Set upgrade height BEFORE building (so it's compiled into the binary)
    set_upgrade_height "$UPGRADE_HEIGHT"
    
    # Build test images (now with correct upgrade height)
    build_test_images
    
    # Setup devnet (L1 + contracts + L2 genesis)
    setup_devnet
    
    # Start L2 with test images
    start_l2_test
    
    # Start tx generator
    start_tx_generator
    
    # Run tests
    test_pbft_mode
    test_upgrade
    test_sequencer_mode
    test_fullnode_sync
    
    stop_tx_generator
    
    log_success "=========================================="
    log_success "  ALL TESTS PASSED!"
    log_success "=========================================="
}

show_status() {
    echo "Node 1:         Block $(get_block_number http://127.0.0.1:8645)"
    echo "Node 2:         Block $(get_block_number http://127.0.0.1:8745)"
    echo "Node 3:         Block $(get_block_number http://127.0.0.1:8845)"
     echo "Node 0 (seq-0): Block $(get_block_number http://127.0.0.1:8545)"
    echo "Sentry:         Block $(get_block_number http://127.0.0.1:8945 2>/dev/null || echo 'N/A')"
}

show_logs() {
    cd "$DOCKER_DIR"
    $COMPOSE_CMD_NO_OVERRIDE logs -f "$@"
}

# ========== Command Parsing ==========

case "${1:-}" in
    build)
        build_test_images
        ;;
    setup)
        setup_devnet
        ;;
    start)
        start_l2_test
        ;;
    stop)
        cd "$DOCKER_DIR"
        $COMPOSE_CMD_NO_OVERRIDE down
        ;;
    clean)
        cleanup
        # Also clean L2 genesis
        rm -rf "$OPS_DIR/l2-genesis/.devnet"
        rm -rf "$DOCKER_DIR/.devnet"
        ;;
    logs)
        shift
        show_logs "$@"
        ;;
    test)
        run_full_test
        ;;
    tx)
        start_tx_generator
        wait
        ;;
    status)
        show_status
        ;;
    upgrade-height)
        set_upgrade_height "${2:-50}"
        ;;
    *)
        echo "Sequencer Upgrade Test Runner"
        echo ""
        echo "Usage: $0 {build|setup|start|stop|clean|logs|test|tx|status|upgrade-height}"
        echo ""
        echo "Commands:"
        echo "  build           - Build test Docker images (morph-geth-test, morph-node-test)"
        echo "  setup           - Run full devnet setup (L1 + contracts + L2 genesis)"
        echo "  start           - Start L2 nodes with test images"
        echo "  stop            - Stop all containers"
        echo "  clean           - Stop and remove all containers and data"
        echo "  logs [service]  - Show container logs"
        echo "  test            - Run full upgrade test"
        echo "  tx              - Start transaction generator"
        echo "  status          - Show current block numbers"
        echo "  upgrade-height N - Set upgrade height to N"
        echo ""
    echo "Environment Variables:"
    echo "  UPGRADE_HEIGHT  - Block height for consensus switch (default: 10)"
        echo "  TX_INTERVAL     - Seconds between txs (default: 5)"
        echo ""
        echo "Test Flow:"
        echo "  1. build        - Build test images"
        echo "  2. setup        - Deploy L1, contracts, generate L2 genesis"
        echo "  3. start        - Start L2 with test images"
        echo "  4. test         - Run PBFT -> Upgrade -> Sequencer -> Fullnode tests"
        echo ""
        echo "Quick Start:"
        echo "  UPGRADE_HEIGHT=10 $0 test"
        ;;
esac
