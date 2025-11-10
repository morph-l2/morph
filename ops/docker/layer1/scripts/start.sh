#!/bin/bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"

# Check if genesis files exist
if [ ! -f "$PROJECT_DIR/genesis/genesis.json" ]; then
    echo "Error: genesis.json does not exist, please run ./scripts/generate-genesis.sh first"
    exit 1
fi

if [ ! -f "$PROJECT_DIR/genesis/genesis.ssz" ]; then
    echo "Error: genesis.ssz does not exist, please run ./scripts/generate-genesis.sh first"
    exit 1
fi

if [ ! -f "$PROJECT_DIR/jwt/jwtsecret" ]; then
    echo "Error: jwtsecret does not exist, please run ./scripts/generate-genesis.sh first"
    exit 1
fi

echo "=== Starting Ethereum Network ==="

cd "$PROJECT_DIR/.."

# Start all services
docker compose -f docker-compose-4nodes.yml up -d layer1-el layer1-cl layer1-vc

echo ""
echo "Waiting for containers to start..."
sleep 15

# Wait for EL node initialization to complete
echo "Waiting for EL node initialization..."
for i in {1..30}; do
    if docker exec layer1-el geth --exec "eth.blockNumber" attach http://localhost:8545 > /dev/null 2>&1; then
        echo "Layer1 EL is ready"
        break
    fi
    sleep 2
done

echo ""
echo "=== Network Started Successfully ==="
echo ""
echo "Access ports:"
echo "  Layer1 EL RPC: http://localhost:9545"
echo "  Layer1 EL WebSocket: ws://localhost:9546"
echo "  Layer1 CL HTTP API: http://localhost:4000"
echo "  Layer1 CL Metrics: http://localhost:5054"
echo ""
echo "View logs:"
echo "  docker compose -f docker-compose-4nodes.yml logs -f layer1-el layer1-cl layer1-vc"
echo ""
echo "Stop network:"
echo "  docker compose -f docker-compose-4nodes.yml down"
