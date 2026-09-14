#!/bin/bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
compose=(docker compose --env-file "${DEVNET_RUNTIME_ENV:-/dev/null}" -f "$PROJECT_DIR/../docker-compose-devnet.yml")

for file in genesis/genesis.json genesis/genesis.ssz genesis/config.yaml \
    genesis/deposit_contract_block.txt jwt/jwtsecret keystores/layer1/keys/validator_definitions.yml; do
    if [ ! -s "$PROJECT_DIR/$file" ]; then
        echo "Required L1 input is missing or empty: $PROJECT_DIR/$file. Restore the original file, or generate genesis if no previous chain exists." >&2
        exit 1
    fi
done

attempts=${L1_START_ATTEMPTS:-60}
interval=${L1_START_INTERVAL:-2}
if [[ ! "$attempts" =~ ^[1-9][0-9]*$ ]] || [[ ! "$interval" =~ ^[0-9]+$ ]]; then
    echo "L1_START_ATTEMPTS must be positive and L1_START_INTERVAL must be a nonnegative integer." >&2
    exit 1
fi
"${compose[@]}" up -d layer1-el layer1-cl layer1-vc
for ((i = 0; i < attempts; i++)); do
    block=$("${compose[@]}" exec -T layer1-el geth --exec 'eth.blockNumber' attach http://localhost:8545 2>/dev/null || true)
    if [[ "$block" =~ ^[0-9]+$ ]] && [ "$block" -ge 1 ]; then
        echo "L1 produced block $block. RPC: http://localhost:9545; beacon API: http://localhost:4000."
        exit 0
    fi
    sleep "$interval"
done

echo "L1 did not produce a block before the startup timeout; containers and data were preserved. Inspect the L1 service logs before retrying." >&2
exit 1
