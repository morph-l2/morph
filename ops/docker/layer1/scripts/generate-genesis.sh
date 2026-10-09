#!/bin/bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
GENESIS_DIR="$PROJECT_DIR/genesis"
CONFIGS_DIR="$PROJECT_DIR/configs"
VALUES_ENV_TEMPLATE="$CONFIGS_DIR/values.env.template"

for required in "$VALUES_ENV_TEMPLATE" "$PROJECT_DIR/jwt/jwtsecret" \
    "$PROJECT_DIR/keystores/layer1/keys/validator_definitions.yml"; do
    if [ ! -s "$required" ]; then
        echo "Required input is missing or empty: $required" >&2
        exit 1
    fi
done
if [ -d "$GENESIS_DIR" ] && [ -n "$(find "$GENESIS_DIR" -mindepth 1 -maxdepth 1 -print -quit)" ]; then
    echo "Genesis data already exists at $GENESIS_DIR; restore incomplete files or explicitly clean the network before generating another chain." >&2
    exit 1
fi
for volume in layer1-el-data layer1-cl-data layer1-vc-data; do
    existing=$(docker volume ls -q \
        --filter "label=com.docker.compose.project=${COMPOSE_PROJECT_NAME:-docker}" \
        --filter "label=com.docker.compose.volume=$volume")
    if [ -n "$existing" ]; then
        echo "L1 data volumes exist; restore their original genesis files before starting the network." >&2
        exit 1
    fi
done

lock_dir="$PROJECT_DIR/.genesis.lock"
if ! mkdir "$lock_dir"; then
    echo "Genesis generation is already locked at $lock_dir; remove the lock only after confirming no generator is running." >&2
    exit 1
fi
trap 'rmdir "$lock_dir"' EXIT
attempt_dir=$(mktemp -d "$PROJECT_DIR/genesis-attempt.XXXXXX")
mkdir "$attempt_dir/output"
echo "Generating L1 genesis in $attempt_dir"
GENESIS_TIMESTAMP=$(($(date -u +%s) + 20))
sed "s/{{GENESIS_TIMESTAMP}}/$GENESIS_TIMESTAMP/g" "$VALUES_ENV_TEMPLATE" > "$attempt_dir/values.env"
printf '[]\n' > "$attempt_dir/additional-contracts.json"

# Generate into a separate directory so a failure cannot leave usable-looking output.
docker run --rm \
    -v "$attempt_dir/values.env:/opt/values.env:ro" \
    -v "$attempt_dir/additional-contracts.json:/opt/additional-contracts.json:ro" \
    -v "$attempt_dir/output:/data" \
    --entrypoint="" \
    ethpandaops/ethereum-genesis-generator:5.1.0 \
    sh -ec 'cp /opt/values.env /config/values.env
        /work/entrypoint.sh all
        if [ -d /data/metadata ]; then
            for file in /data/metadata/*; do
                [ -e "$file" ] || continue
                mv "$file" /data/
            done
            rmdir /data/metadata
        fi'

for file in genesis.json genesis.ssz config.yaml deposit_contract_block.txt; do
    if [ ! -s "$attempt_dir/output/$file" ]; then
        echo "Generator did not produce $file; attempt data is preserved at $attempt_dir." >&2
        exit 1
    fi
done
# Remove only a pre-existing empty output directory after successful generation.
if [ -d "$GENESIS_DIR" ]; then
    rmdir "$GENESIS_DIR"
fi
mv "$attempt_dir/output" "$GENESIS_DIR"
mv "$attempt_dir/values.env" "$CONFIGS_DIR/values.env"
mv "$attempt_dir/additional-contracts.json" "$CONFIGS_DIR/additional-contracts.json"
rmdir "$attempt_dir"
echo "L1 genesis generated at $GENESIS_DIR (timestamp: $GENESIS_TIMESTAMP)."
