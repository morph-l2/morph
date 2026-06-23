#!/bin/bash
# ============================================================
# Persistent Peer Whitelist Integration Test — Observer Script
# ============================================================
# Run AFTER `run-ha-test.sh start` and after the cluster has crossed
# the upgrade height (so node-0 is now a fullnode, ha-node-X is the
# producer, and node-0 is gossiping blocks to sentry-node-0).
#
# This script collects evidence from container logs and prints a
# structured pass/fail summary. It does NOT modify state.

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
MORPH_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"
DOCKER_DIR="$MORPH_ROOT/ops/docker"
COMPOSE_HA="docker compose -f docker-compose-4nodes.yml -f docker-compose.override.yml -f docker-compose.ha-override.yml -f docker-compose.whitelist-test.override.yml"

NODE0_NODEID="93e27ea2306e158a8146d5f44caaab97496797d2"
SENTRY_NODEID="dae813274913aaf39e7cd3226a0aa8bce00644e1"

LOG_TAIL=20000

cd "$DOCKER_DIR"

echo "=========================================="
echo "Whitelist Integration Test — Evidence"
echo "=========================================="

# Evidence 1: node-0 DID inject malicious blocks
echo ""
echo "--- Evidence 1: node-0 malicious-inject events ---"
INJECT_COUNT=$($COMPOSE_HA logs --tail $LOG_TAIL node-0 2>&1 | grep -c "MALICIOUS_INJECT" || true)
echo "node-0 [MALICIOUS_INJECT] events: $INJECT_COUNT"

# Evidence 2: sentry triggered the whitelist alarm
echo ""
echo "--- Evidence 2: sentry [WHITELIST_ALARM] events ---"
ALARM_COUNT=$($COMPOSE_HA logs --tail $LOG_TAIL sentry-node-0 2>&1 | grep -c "WHITELIST_ALARM" || true)
echo "sentry [WHITELIST_ALARM] events: $ALARM_COUNT"
$COMPOSE_HA logs --tail $LOG_TAIL sentry-node-0 2>&1 | grep "WHITELIST_ALARM" | tail -5

# Evidence 3: sentry stopped node-0 (disconnect happened)
echo ""
echo "--- Evidence 3: sentry stopped node-0 for error ---"
STOP_COUNT=$($COMPOSE_HA logs --tail $LOG_TAIL sentry-node-0 2>&1 | grep "Stopping peer for error" | grep -c "$NODE0_NODEID" || true)
echo "sentry stopped node-0 (by ID match): $STOP_COUNT"

# Evidence 4: sentry reconnected to node-0
echo ""
echo "--- Evidence 4: sentry reconnected to node-0 ---"
RECONNECT_COUNT=$($COMPOSE_HA logs --tail $LOG_TAIL sentry-node-0 2>&1 | grep -c "Reconnecting to peer" || true)
echo "sentry [Reconnecting to peer] events: $RECONNECT_COUNT"
$COMPOSE_HA logs --tail $LOG_TAIL sentry-node-0 2>&1 | grep "Reconnecting to peer" | tail -3

# Evidence 5: sentry did NOT add node-0 to bannedPeers
echo ""
echo "--- Evidence 5: sentry did NOT ban node-0 ---"
BAN_COUNT=$($COMPOSE_HA logs --tail $LOG_TAIL sentry-node-0 2>&1 | grep "Banning peer" | grep -c "$NODE0_NODEID" || true)
echo "sentry 'Banning peer' for node-0: $BAN_COUNT (expected 0)"

# Evidence 6: sentry continued syncing blocks after reconnect events
echo ""
echo "--- Evidence 6: sentry sync still progressing ---"
APPLY_COUNT=$($COMPOSE_HA logs --tail $LOG_TAIL sentry-node-0 2>&1 | grep -c "Applied block" || true)
echo "sentry 'Applied block' events: $APPLY_COUNT"

# Evidence 7: sentry block height
echo ""
echo "--- Evidence 7: current heights ---"
SENTRY_HEIGHT=$(curl -s http://127.0.0.1:8945 -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' 2>/dev/null | grep -o '"result":"[^"]*"' | cut -d'"' -f4 || echo "unreachable")
HA0_HEIGHT=$(curl -s http://127.0.0.1:9145 -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' 2>/dev/null | grep -o '"result":"[^"]*"' | cut -d'"' -f4 || echo "unreachable")
echo "sentry-el-0 height: $SENTRY_HEIGHT"
echo "ha-geth-0 height: $HA0_HEIGHT"

echo ""
echo "=========================================="
echo "Verdict"
echo "=========================================="
echo "PASS criteria:"
echo "  - INJECT > 0  (we did inject):                $([ $INJECT_COUNT -gt 0 ] && echo PASS || echo FAIL)"
echo "  - ALARM > 0   (whitelist triggered):          $([ $ALARM_COUNT -gt 0 ] && echo PASS || echo FAIL)"
echo "  - STOP > 0    (disconnect happened):          $([ $STOP_COUNT -gt 0 ] && echo PASS || echo FAIL)"
echo "  - RECONNECT > 0 (reconnect happened):         $([ $RECONNECT_COUNT -gt 0 ] && echo PASS || echo FAIL)"
echo "  - BAN == 0    (node-0 not in ban list):       $([ $BAN_COUNT -eq 0 ] && echo PASS || echo FAIL)"
echo "  - APPLY > 0   (sentry still syncing):         $([ $APPLY_COUNT -gt 0 ] && echo PASS || echo FAIL)"
