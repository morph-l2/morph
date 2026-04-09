#!/bin/bash
# ============================================================
# HA Block Apply Performance Test
# ============================================================
# Measures Raft consensus overhead during block apply.
#
# Usage:
#   ./run-perf-test.sh build     - Rebuild test images with perf instrumentation
#   ./run-perf-test.sh setup     - Deploy L1 + contracts + L2 genesis
#   ./run-perf-test.sh start     - Start HA cluster
#   ./run-perf-test.sh load      - Start TX load generator
#   ./run-perf-test.sh run       - Full test: start + load + wait + analyze
#   ./run-perf-test.sh analyze   - Parse [PERF] logs and print summary
#   ./run-perf-test.sh all       - build + setup + start + run
#   ./run-perf-test.sh stop      - Stop everything
#   ./run-perf-test.sh clean     - Full cleanup
#
# Environment:
#   PERF_DURATION  - How long to collect data in seconds (default: 120)
#   TX_RATE        - Target TXs per second (default: 10)
#   UPGRADE_HEIGHT - Consensus switch height (default: 10)

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MORPH_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"
BITGET_ROOT="$(cd "$MORPH_ROOT/.." && pwd)"
OPS_DIR="$MORPH_ROOT/ops"
DOCKER_DIR="$OPS_DIR/docker"

# ── Config ────────────────────────────────────────────────────────────────────
PERF_DURATION=${PERF_DURATION:-120}
TX_RATE=${TX_RATE:-10}
UPGRADE_HEIGHT=${UPGRADE_HEIGHT:-10}
PRIVATE_KEY="0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
FROM_ADDR="0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"

L2_RPC="http://127.0.0.1:8545"
HA_RPC_NODE0="http://127.0.0.1:9501"

COMPOSE_BASE="docker compose -f docker-compose-4nodes.yml"
COMPOSE_OVERRIDE="docker compose -f docker-compose-4nodes.yml -f docker-compose.override.yml"
COMPOSE_HA="docker compose -f docker-compose-4nodes.yml -f docker-compose.override.yml -f docker-compose.ha-override.yml"

# ── Colors ────────────────────────────────────────────────────────────────────
RED='\033[0;31m'; GREEN='\033[0;32m'; YELLOW='\033[1;33m'
BLUE='\033[0;34m'; CYAN='\033[0;36m'; BOLD='\033[1m'; NC='\033[0m'

log_info()  { echo -e "${BLUE}[INFO]${NC} $1"; }
log_ok()    { echo -e "${GREEN}[ OK ]${NC} $1"; }
log_warn()  { echo -e "${YELLOW}[WARN]${NC} $1"; }
log_err()   { echo -e "${RED}[ERR ]${NC} $1"; }
log_section() { echo -e "\n${BOLD}${CYAN}── $1 ──${NC}"; }

# ── Helpers ───────────────────────────────────────────────────────────────────

wait_for_rpc() {
    local url="$1" max=${2:-60} i=0
    while [ $i -lt $max ]; do
        if curl -sf -X POST -H "Content-Type: application/json" \
            --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
            "$url" | grep -q "result"; then
            return 0
        fi
        sleep 2; i=$((i + 2))
    done
    return 1
}

get_block_number() {
    local url="${1:-$L2_RPC}"
    local hex
    hex=$(curl -sf -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
        "$url" | grep -o '"result":"0x[^"]*"' | cut -d'"' -f4)
    printf "%d" "$hex" 2>/dev/null || echo 0
}

wait_for_block() {
    local target=$1 url="${2:-$L2_RPC}" max=${3:-300} waited=0
    while [ $waited -lt $max ]; do
        local cur=$(get_block_number "$url")
        if [ "$cur" -ge "$target" ]; then return 0; fi
        echo -ne "\r  block: $cur / $target"
        sleep 3; waited=$((waited + 3))
    done
    echo ""; return 1
}

wait_for_ha_leader() {
    local max=${1:-60} waited=0
    while [ $waited -lt $max ]; do
        for rpc in http://127.0.0.1:9501 http://127.0.0.1:9601 http://127.0.0.1:9701; do
            local resp
            resp=$(curl -sf -X POST -H "Content-Type: application/json" \
                --data '{"jsonrpc":"2.0","method":"hakeeper_leader","params":[],"id":1}' \
                "$rpc" 2>/dev/null || true)
            if echo "$resp" | grep -q '"result":true'; then
                log_ok "HA leader found at $rpc"
                return 0
            fi
        done
        sleep 3; waited=$((waited + 3))
    done
    log_err "No HA leader found within ${max}s"
    return 1
}

# ── Build ─────────────────────────────────────────────────────────────────────

do_build() {
    log_section "Building test images with perf instrumentation"

    cd "$MORPH_ROOT"
    make go-ubuntu-builder

    cd "$BITGET_ROOT"
    log_info "Building morph-geth-test..."
    docker build -t morph-geth-test:latest \
        -f morph/ops/docker-sequencer-test/Dockerfile.l2-geth-test .

    log_info "Building morph-node-test..."
    docker build -t morph-node-test:latest \
        -f morph/ops/docker-sequencer-test/Dockerfile.l2-node-test .

    log_ok "Test images built"
}

# ── Setup ─────────────────────────────────────────────────────────────────────

do_setup() {
    log_section "Setting up devnet (L1 + contracts + L2 genesis)"
    cd "$SCRIPT_DIR"
    ./run-test.sh clean || true
    ./run-test.sh setup
    log_ok "Setup complete"
}

# ── Start HA cluster ──────────────────────────────────────────────────────────

do_start() {
    log_section "Starting HA cluster"
    cd "$DOCKER_DIR"

    # Copy override files
    cp "$SCRIPT_DIR/docker-compose.override.yml" .
    cp "$SCRIPT_DIR/docker-compose.ha-override.yml" .
    source .env 2>/dev/null || true

    # Wait for L1 finalized
    log_info "Waiting for L1 to finalize..."
    local l1_latest
    l1_latest=$(curl -sf -X POST -H "Content-Type: application/json" \
        --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
        http://127.0.0.1:9545 2>/dev/null | grep -o '"result":"0x[^"]*"' | cut -d'"' -f4)
    l1_latest=$(printf "%d" "$l1_latest" 2>/dev/null || echo 1)

    local waited=0
    while [ $waited -lt 120 ]; do
        local fin
        fin=$(curl -sf -X POST -H "Content-Type: application/json" \
            --data '{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["finalized",false],"id":1}' \
            http://127.0.0.1:9545 2>/dev/null | grep -o '"number":"0x[^"]*"' | head -1 | cut -d'"' -f4)
        local fin_dec=$(printf "%d" "$fin" 2>/dev/null || echo 0)
        if [ "$fin_dec" -ge "$l1_latest" ]; then
            log_ok "L1 finalized at $fin_dec"
            break
        fi
        echo -ne "\r  L1 finalized: $fin_dec / $l1_latest"
        sleep 3; waited=$((waited + 3))
    done

    # Stop any existing
    $COMPOSE_HA stop morph-geth-0 morph-geth-1 morph-geth-2 morph-geth-3 \
        node-0 node-1 node-2 node-3 2>/dev/null || true

    # Clean Raft state for fresh cluster
    rm -rf .devnet/node0/raft .devnet/node1/raft .devnet/node2/raft 2>/dev/null || true

    # Start geth nodes
    log_info "Starting geth nodes..."
    $COMPOSE_HA up -d morph-geth-0 morph-geth-1 morph-geth-2 morph-geth-3
    sleep 5

    # Start tendermint nodes
    log_info "Starting tendermint nodes (node-0: bootstrap, node-1/2: join, node-3: plain)..."
    $COMPOSE_HA up -d node-0 node-1 node-2 node-3

    log_info "Waiting for L2 RPC..."
    wait_for_rpc "$L2_RPC" 60 || { log_err "L2 RPC not ready"; return 1; }

    # Wait for upgrade height (PBFT → V2 switch)
    log_info "Waiting for upgrade height ($UPGRADE_HEIGHT)..."
    wait_for_block $UPGRADE_HEIGHT "$L2_RPC" 300 || { log_err "Upgrade height not reached"; return 1; }
    echo ""

    # Wait for HA leader
    log_info "Waiting for HA cluster formation..."
    sleep 10
    wait_for_ha_leader 60 || { log_warn "HA leader not found, checking logs..."; }

    log_ok "HA cluster running"
}

# ── TX Load Generator ────────────────────────────────────────────────────────

TX_GEN_PIDS=()
TXFLOOD_BIN="${SCRIPT_DIR}/txflood/txflood"

start_tx_load() {
    local num_senders=${TX_SENDERS:-5}
    local dur="${PERF_DURATION:-120}s"

    # Build txflood if missing or stale
    if [ ! -f "$TXFLOOD_BIN" ] || [ "$SCRIPT_DIR/txflood/main.go" -nt "$TXFLOOD_BIN" ]; then
        log_info "Building txflood..."
        (cd "$MORPH_ROOT" && go build -o "$TXFLOOD_BIN" ./ops/docker-sequencer-test/txflood/main.go)
        log_ok "txflood built"
    fi

    log_section "Starting TX load (Go txflood, ${num_senders} senders, ~${dur})"

    RPC_URL="$L2_RPC" SENDERS="$num_senders" DURATION="$dur" "$TXFLOOD_BIN" &
    TX_GEN_PIDS+=($!)

    log_ok "txflood started (PID: ${TX_GEN_PIDS[*]})"
}

stop_tx_load() {
    if [ ${#TX_GEN_PIDS[@]} -gt 0 ]; then
        for pid in "${TX_GEN_PIDS[@]}"; do
            kill "$pid" 2>/dev/null || true
        done
        for pid in "${TX_GEN_PIDS[@]}"; do
            wait "$pid" 2>/dev/null || true
        done
        TX_GEN_PIDS=()
        log_info "txflood stopped"
    fi
}

# ── Log Analysis ──────────────────────────────────────────────────────────────

do_analyze() {
    log_section "Collecting and analyzing [PERF] logs"
    cd "$DOCKER_DIR"

    local tmpdir=$(mktemp -d)
    local since="${PERF_LOG_SINCE:-}"

    # Collect logs from all nodes
    for node in node-0 node-1 node-2; do
        if [ -n "$since" ]; then
            docker logs --since "$since" "$node" 2>&1 | grep '\[PERF\]' > "$tmpdir/$node.log" 2>/dev/null || true
        else
            docker logs "$node" 2>&1 | grep '\[PERF\]' > "$tmpdir/$node.log" 2>/dev/null || true
        fi
    done

    # ── Summary per node ──
    for node in node-0 node-1 node-2; do
        local logfile="$tmpdir/$node.log"
        local count=$(wc -l < "$logfile" | tr -d ' ')

        if [ "$count" -eq 0 ]; then
            log_warn "$node: no [PERF] entries found"
            continue
        fi

        echo ""
        echo -e "${BOLD}═══ $node ($count entries) ═══${NC}"

        # produceBlock (only on leader = node-0 typically)
        local produce_count; produce_count=$(grep -c 'produceBlock' "$logfile" 2>/dev/null || true); produce_count=${produce_count:-0}
        if [ "${produce_count}" -gt 0 ] 2>/dev/null; then
            echo -e "\n${CYAN}[produceBlock] ($produce_count blocks)${NC}"
            grep 'produceBlock' "$logfile" | awk '
            {
                build=0; sign=0; commit=0; total=0; tx=0; gas=0
                for(i=1;i<=NF;i++) {
                    if($i ~ /build_ms=/)  { split($i,a,"="); build=a[2]+0 }
                    if($i ~ /sign_ms=/)   { split($i,a,"="); sign=a[2]+0 }
                    if($i ~ /raft_commit_ms=/) { split($i,a,"="); commit=a[2]+0 }
                    if($i ~ /apply_ms=/)  { split($i,a,"="); commit=a[2]+0 }
                    if($i ~ /total_ms=/)  { split($i,a,"="); total=a[2]+0 }
                    if($i ~ /txCount=/)   { split($i,a,"="); tx=a[2]+0 }
                    if($i ~ /gasUsed=/)   { split($i,a,"="); gas=a[2]+0 }
                }
                n++; s_build+=build; s_sign+=sign; s_commit+=commit; s_total+=total; s_tx+=tx; s_gas+=gas
                if(build>max_build) max_build=build
                if(commit>max_commit) max_commit=commit
                if(total>max_total) max_total=total
                if(n==1 || build<min_build) min_build=build
                if(n==1 || commit<min_commit) min_commit=commit
                if(n==1 || total<min_total) min_total=total
            }
            END {
                if(n>0) {
                    printf "  %-18s avg=%-10.2f min=%-10.2f max=%.2f\n", "build_ms:", s_build/n, min_build, max_build
                    printf "  %-18s avg=%-10.2f min=%-10.2f max=%.2f\n", "sign_ms:", s_sign/n, 0, 0
                    printf "  %-18s avg=%-10.2f min=%-10.2f max=%.2f\n", "raft_commit_ms:", s_commit/n, min_commit, max_commit
                    printf "  %-18s avg=%-10.2f min=%-10.2f max=%.2f\n", "total_ms:", s_total/n, min_total, max_total
                    printf "  %-18s avg=%.1f\n", "txCount:", s_tx/n
                    printf "  %-18s avg=%.0f\n", "gasUsed:", s_gas/n
                }
            }'
        fi

        # HAService.Commit (only on leader)
        local commit_count; commit_count=$(grep -c 'HAService.Commit' "$logfile" 2>/dev/null || true); commit_count=${commit_count:-0}
        if [ "${commit_count}" -gt 0 ] 2>/dev/null; then
            echo -e "\n${CYAN}[HAService.Commit] ($commit_count entries)${NC}"
            grep 'HAService.Commit' "$logfile" | awk '
            {
                enc=0; raft=0; total=0; bytes=0
                for(i=1;i<=NF;i++) {
                    if($i ~ /encode_ms=/) { split($i,a,"="); enc=a[2]+0 }
                    if($i ~ /raft_ms=/)   { split($i,a,"="); raft=a[2]+0 }
                    if($i ~ /total_ms=/)  { split($i,a,"="); total=a[2]+0 }
                    if($i ~ /dataBytes=/) { split($i,a,"="); bytes=a[2]+0 }
                }
                n++; s_enc+=enc; s_raft+=raft; s_total+=total; s_bytes+=bytes
                if(raft>max_raft) max_raft=raft
                if(n==1 || raft<min_raft) min_raft=raft
            }
            END {
                if(n>0) {
                    printf "  %-18s avg=%-10.2f\n", "encode_ms:", s_enc/n
                    printf "  %-18s avg=%-10.2f min=%-10.2f max=%.2f\n", "raft_ms:", s_raft/n, min_raft, max_raft
                    printf "  %-18s avg=%-10.2f\n", "total_ms:", s_total/n
                    printf "  %-18s avg=%.0f\n", "dataBytes:", s_bytes/n
                }
            }'
        fi

        # BlockFSM.Apply (on all HA nodes)
        local fsm_count=$(grep -c 'BlockFSM.Apply' "$logfile" 2>/dev/null || echo 0)
        if [ "$fsm_count" -gt 0 ]; then
            echo -e "\n${CYAN}[BlockFSM.Apply] ($fsm_count entries)${NC}"
            grep 'BlockFSM.Apply' "$logfile" | awk '
            {
                dec=0; applied=0; total=0
                for(i=1;i<=NF;i++) {
                    if($i ~ /decode_ms=/)    { split($i,a,"="); dec=a[2]+0 }
                    if($i ~ /onApplied_ms=/) { split($i,a,"="); applied=a[2]+0 }
                    if($i ~ /total_ms=/)     { split($i,a,"="); total=a[2]+0 }
                }
                n++; s_dec+=dec; s_applied+=applied; s_total+=total
                if(applied>max_applied) max_applied=applied
                if(total>max_total) max_total=total
                if(n==1 || applied<min_applied) min_applied=applied
                if(n==1 || total<min_total) min_total=total
            }
            END {
                if(n>0) {
                    printf "  %-18s avg=%-10.2f\n", "decode_ms:", s_dec/n
                    printf "  %-18s avg=%-10.2f min=%-10.2f max=%.2f\n", "onApplied_ms:", s_applied/n, min_applied, max_applied
                    printf "  %-18s avg=%-10.2f min=%-10.2f max=%.2f\n", "total_ms:", s_total/n, min_total, max_total
                }
            }'
        fi

        # ApplyBlock (on all HA nodes)
        local apply_count=$(grep -c 'ApplyBlock' "$logfile" | head -1 2>/dev/null || echo 0)
        # Exclude produceBlock lines
        local pure_apply=$(grep 'ApplyBlock' "$logfile" | grep -cv 'produceBlock' 2>/dev/null || echo 0)
        if [ "$pure_apply" -gt 0 ]; then
            echo -e "\n${CYAN}[ApplyBlock] ($pure_apply entries)${NC}"
            grep 'ApplyBlock' "$logfile" | grep -v 'produceBlock' | awk '
            {
                geth=0; sig=0; total=0
                for(i=1;i<=NF;i++) {
                    if($i ~ /geth_ms=/)    { split($i,a,"="); geth=a[2]+0 }
                    if($i ~ /sigSave_ms=/) { split($i,a,"="); sig=a[2]+0 }
                    if($i ~ /total_ms=/)   { split($i,a,"="); total=a[2]+0 }
                }
                n++; s_geth+=geth; s_sig+=sig; s_total+=total
                if(geth>max_geth) max_geth=geth
                if(n==1 || geth<min_geth) min_geth=geth
            }
            END {
                if(n>0) {
                    printf "  %-18s avg=%-10.2f min=%-10.2f max=%.2f\n", "geth_ms:", s_geth/n, min_geth, max_geth
                    printf "  %-18s avg=%-10.2f\n", "sigSave_ms:", s_sig/n
                    printf "  %-18s avg=%-10.2f\n", "total_ms:", s_total/n
                }
            }'
        fi
    done

    # ── Raft overhead summary ──
    echo ""
    log_section "Raft Overhead Summary"

    local leader_raft_avg leader_fsm_avg
    leader_raft_avg=$(grep 'HAService.Commit' "$tmpdir/node-0.log" 2>/dev/null | awk '
    { for(i=1;i<=NF;i++) if($i ~ /raft_ms=/) { split($i,a,"="); s+=a[2]+0; n++ } }
    END { if(n>0) printf "%.2f", s/n; else print "N/A" }')

    leader_fsm_avg=$(grep 'BlockFSM.Apply' "$tmpdir/node-0.log" 2>/dev/null | awk '
    { for(i=1;i<=NF;i++) if($i ~ /onApplied_ms=/) { split($i,a,"="); s+=a[2]+0; n++ } }
    END { if(n>0) printf "%.2f", s/n; else print "N/A" }')

    echo -e "  Leader raft_ms avg:      ${BOLD}${leader_raft_avg}${NC} ms"
    echo -e "  Leader onApplied_ms avg: ${BOLD}${leader_fsm_avg}${NC} ms"

    if [[ "$leader_raft_avg" != "N/A" && "$leader_fsm_avg" != "N/A" ]]; then
        local overhead
        overhead=$(awk "BEGIN { printf \"%.2f\", $leader_raft_avg - $leader_fsm_avg }")
        echo -e "  ${BOLD}Pure Raft overhead:        ${RED}${overhead}${NC} ms${NC} (network + quorum + log write)"
    fi

    # Follower comparison
    for node in node-1 node-2; do
        local f_avg
        f_avg=$(grep 'BlockFSM.Apply' "$tmpdir/$node.log" 2>/dev/null | awk '
        { for(i=1;i<=NF;i++) if($i ~ /onApplied_ms=/) { split($i,a,"="); s+=a[2]+0; n++ } }
        END { if(n>0) printf "%.2f", s/n; else print "N/A" }')
        echo -e "  $node onApplied_ms avg:  ${BOLD}${f_avg}${NC} ms"
    done

    rm -rf "$tmpdir"
    echo ""
}

# ── Run (full test cycle) ────────────────────────────────────────────────────

do_run() {
    log_section "Running HA performance test (${PERF_DURATION}s)"

    local start_block=$(get_block_number "$L2_RPC")
    log_info "Starting at block $start_block"

    start_tx_load

    local start_ts=$(date -u +%Y-%m-%dT%H:%M:%SZ)

    log_info "Collecting data for ${PERF_DURATION}s (txflood running)..."
    # Wait for txflood to finish (it runs for PERF_DURATION then exits)
    for pid in "${TX_GEN_PIDS[@]}"; do
        wait "$pid" 2>/dev/null || true
    done
    TX_GEN_PIDS=()

    local end_block=$(get_block_number "$L2_RPC")
    local blocks=$((end_block - start_block))
    log_ok "Collected $blocks blocks ($start_block → $end_block)"

    PERF_LOG_SINCE="$start_ts" do_analyze
}

# ── Stop ──────────────────────────────────────────────────────────────────────

do_stop() {
    log_section "Stopping all containers"
    stop_tx_load
    cd "$DOCKER_DIR"
    $COMPOSE_HA stop morph-geth-0 morph-geth-1 morph-geth-2 morph-geth-3 \
        node-0 node-1 node-2 node-3 2>/dev/null || true
    log_ok "Stopped"
}

# ── Clean ─────────────────────────────────────────────────────────────────────

do_clean() {
    log_section "Full cleanup"

    # 1. Clean L2 containers + data
    cd "$SCRIPT_DIR"
    ./run-test.sh clean || true

    # 2. Clean L1 volumes + genesis (MUST do this, otherwise beacon chain gets
    #    stuck at head_slot=0 with stale genesis on next setup)
    cd "$DOCKER_DIR"
    $COMPOSE_BASE down -v 2>/dev/null || true
    bash "$OPS_DIR/docker/layer1/scripts/clean.sh" 2>/dev/null || true

    # 3. Clean tendermint + L2 genesis state
    rm -rf "$DOCKER_DIR/.devnet" "$OPS_DIR/l2-genesis/.devnet" 2>/dev/null || true

    log_ok "Cleaned"
}

# ── Main ──────────────────────────────────────────────────────────────────────

case "${1:-help}" in
    build)   do_build ;;
    setup)   do_setup ;;
    start)   do_start ;;
    load)    start_tx_load; echo "Press Ctrl+C to stop"; wait ;;
    run)     do_run ;;
    analyze) do_analyze ;;
    all)
        do_build
        do_setup
        do_start
        do_run
        ;;
    stop)    do_stop ;;
    clean)   do_clean ;;
    *)
        echo "Usage: $0 {build|setup|start|load|run|analyze|all|stop|clean}"
        echo ""
        echo "  build   - Rebuild test images with perf instrumentation"
        echo "  setup   - Deploy L1 + contracts + L2 genesis"
        echo "  start   - Start HA cluster (waits for upgrade + cluster formation)"
        echo "  load    - Start TX load generator (interactive)"
        echo "  run     - Start load + collect ${PERF_DURATION}s + analyze"
        echo "  analyze - Parse existing [PERF] logs and print summary"
        echo "  all     - build + setup + start + run"
        echo "  stop    - Stop L2 containers"
        echo "  clean   - Full cleanup (L1 + L2 + data)"
        ;;
esac
