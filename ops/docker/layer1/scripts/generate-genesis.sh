#!/bin/bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"

GENESIS_DIR="$PROJECT_DIR/genesis"
KEYSTORES_DIR="$PROJECT_DIR/keystores"
JWT_DIR="$PROJECT_DIR/jwt"
CONFIGS_DIR="$PROJECT_DIR/configs"

# Create necessary directories
mkdir -p "$GENESIS_DIR" "$CONFIGS_DIR"

echo "=== Generating Genesis Data ==="

# Calculate genesis timestamp (current time + 20 seconds)
GENESIS_TIMESTAMP=$(date -u +%s)
GENESIS_TIMESTAMP=$((GENESIS_TIMESTAMP + 20))

# Use ethereum-genesis-generator to generate genesis
echo "Generating genesis data..."

# Create values.env file from template
VALUES_ENV_TEMPLATE="$SCRIPT_DIR/../configs/values.env.template"
VALUES_ENV="$CONFIGS_DIR/values.env"

if [ ! -f "$VALUES_ENV_TEMPLATE" ]; then
    echo "Error: Template file not found: $VALUES_ENV_TEMPLATE"
    exit 1
fi

# Replace placeholders in template
sed -e "s/{{GENESIS_TIMESTAMP}}/$GENESIS_TIMESTAMP/g" \
    "$VALUES_ENV_TEMPLATE" > "$VALUES_ENV"

# Create additional-contracts.json
ADDITIONAL_CONTRACTS="$CONFIGS_DIR/additional-contracts.json"
echo "[]" > "$ADDITIONAL_CONTRACTS"

# Run genesis generator
# The generator creates files in /data/metadata/ which need to be moved to /data/ root
# Since /data is mounted to genesis/ on the host, moving to /data automatically puts files in genesis/
# IMPORTANT: additional-contracts.json must be at /opt/additional-contracts.json as specified in values.env
echo "Running genesis generator..."
docker run --rm \
    -v "$CONFIGS_DIR/values.env:/opt/values.env" \
    -v "$CONFIGS_DIR/additional-contracts.json:/opt/additional-contracts.json" \
    -v "$GENESIS_DIR:/data" \
    --entrypoint="" \
    ethpandaops/ethereum-genesis-generator:5.1.0 \
    sh -c "cp /opt/values.env /config/values.env && /work/entrypoint.sh all && mv /data/metadata/* /data/ 2>/dev/null || true && [ -d /data/parsed ] && mv /data/parsed /data/ 2>/dev/null || true && rmdir /data/metadata 2>/dev/null || true"

# Verify that files were generated (they should be directly in genesis/ now)
if [ ! -f "$GENESIS_DIR/genesis.json" ]; then
    echo "❌ Error: genesis.json not found after generation"
    exit 1
fi

# Check output
if [ -f "$GENESIS_DIR/genesis.json" ]; then
    echo "✓ genesis.json generated successfully"
else
    echo "❌ Error: genesis.json not generated"
    exit 1
fi

if [ -f "$GENESIS_DIR/genesis.ssz" ]; then
    echo "✓ genesis.ssz generated successfully"
else
    echo "⚠ Warning: genesis.ssz not generated"
fi

if [ -f "$GENESIS_DIR/config.yaml" ]; then
    echo "✓ config.yaml generated successfully"
else
    echo "⚠ Warning: config.yaml not generated"
fi

if [ -f "$GENESIS_DIR/deposit_contract_block.txt" ]; then
    echo "✓ deposit_contract_block.txt generated successfully"
else
    echo "⚠ Warning: deposit_contract_block.txt not generated"
fi

echo ""
echo "=== Genesis Data Generation Complete ==="
echo "Genesis files location: $GENESIS_DIR"
echo "Keystores location: $KEYSTORES_DIR"
echo "JWT secret location: $JWT_DIR/jwtsecret"
echo ""
echo "Genesis timestamp: $GENESIS_TIMESTAMP"
echo "Validator count: 32"
echo ""
echo "Lighthouse testnet-dir files:"
ls -1 "$GENESIS_DIR"/*.{yaml,ssz,txt} 2>/dev/null | xargs -n 1 basename | head -10
