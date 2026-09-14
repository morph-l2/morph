#!/bin/bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
compose=(docker compose --env-file "${DEVNET_RUNTIME_ENV:-/dev/null}" -f "$PROJECT_DIR/../docker-compose-devnet.yml")

if [ -d "$PROJECT_DIR/.genesis.lock" ]; then
    echo "Genesis generation is locked; wait for it to finish before cleaning L1 data." >&2
    exit 1
fi

# Preserve all files when Docker cannot stop/remove this project's L1 services.
"${compose[@]}" stop layer1-el layer1-cl layer1-vc
"${compose[@]}" rm -f layer1-el layer1-cl layer1-vc
for volume in layer1-el-data layer1-cl-data layer1-vc-data; do
    volumes=$(docker volume ls -q \
        --filter "label=com.docker.compose.project=${COMPOSE_PROJECT_NAME:-docker}" \
        --filter "label=com.docker.compose.volume=$volume")
    while IFS= read -r volume_name; do
        [ -n "$volume_name" ] || continue
        docker volume rm "$volume_name"
    done <<< "$volumes"
done

rm -rf "$PROJECT_DIR/genesis"
rm -f "$PROJECT_DIR/configs/values.env" "$PROJECT_DIR/configs/additional-contracts.json"

# The validator keys remain available for the replacement local chain.
if [ -d "$PROJECT_DIR/keystores/layer1" ]; then
    find "$PROJECT_DIR/keystores/layer1" -type f \
        \( -name '*.db' -o -name '*.sqlite' -o -name '*.sqlite3' \) -exec rm -f {} +
    find "$PROJECT_DIR/keystores/layer1" -type d -name slashing_protection -prune -exec rm -rf {} +
fi

echo "Removed this Compose project's L1 containers, volumes and generated genesis/configuration files."
echo "Preserved configuration templates, JWT secret, validator keys and L2 data."
echo "Run make -f ops/docker/Makefile.layer1 generate start from the repository root to create and start another local L1 chain."
