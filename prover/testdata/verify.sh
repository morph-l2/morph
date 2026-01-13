#!/bin/bash

# Script to fetch block traces and verify them with Rust prover
# Usage: ./verify.sh <start_block> <end_block> [rpc_url]
# Example: ./verify.sh 10 20 http://localhost:8545

set -e

# Check if required commands exist
command -v cast >/dev/null 2>&1 || { echo "Error: 'cast' command not found. Please install foundry."; exit 1; }
command -v jq >/dev/null 2>&1 || { echo "Error: 'jq' command not found. Please install jq."; exit 1; }

# Parse arguments
if [ $# -lt 2 ]; then
    echo "Usage: $0 <start_block> <end_block> [rpc_url]"
    echo "Example: $0 10 20 http://localhost:8545"
    exit 1
fi

START_BLOCK=$1
END_BLOCK=$2
RPC_URL=${3:-"http://localhost:8545"}

# Validate block numbers
if ! [[ "$START_BLOCK" =~ ^[0-9]+$ ]] || ! [[ "$END_BLOCK" =~ ^[0-9]+$ ]]; then
    echo "Error: Block numbers must be integers"
    exit 1
fi

if [ "$START_BLOCK" -gt "$END_BLOCK" ]; then
    echo "Error: Start block must be less than or equal to end block"
    exit 1
fi

# Create output directory if it doesn't exist
OUTPUT_DIR="./generated"
mkdir -p "$OUTPUT_DIR"

OUTPUT_FILE="$OUTPUT_DIR/block_traces_${START_BLOCK}_${END_BLOCK}.json"
TEMP_FILE="$OUTPUT_DIR/temp_traces.json"

echo "Fetching block traces from block $START_BLOCK to $END_BLOCK..."
echo "RPC URL: $RPC_URL"
echo "Output file: $OUTPUT_FILE"
echo ""

# Initialize the JSON array with outer and inner array
echo "[" > "$TEMP_FILE"
echo "  [" >> "$TEMP_FILE"

FIRST=true
SUCCESS_COUNT=0
FAIL_COUNT=0

# Iterate through blocks
for ((block=$START_BLOCK; block<=$END_BLOCK; block++)); do
    # Convert block number to hex
    BLOCK_HEX=$(printf "0x%x" $block)
    
    echo -n "Fetching block $block (${BLOCK_HEX})... "
    
    # Fetch the block trace
    TRACE=$(cast rpc morph_getBlockTraceByNumberOrHash "$BLOCK_HEX" --rpc-url "$RPC_URL" 2>/dev/null)
    
    if [ $? -eq 0 ] && [ -n "$TRACE" ] && [ "$TRACE" != "null" ]; then
        # Add comma separator if not first element
        if [ "$FIRST" = false ]; then
            echo "," >> "$TEMP_FILE"
        fi
        FIRST=false
        
        # Append the trace directly (without wrapping in individual array)
        echo -n "    $TRACE" >> "$TEMP_FILE"
        
        echo "✓ Success"
        SUCCESS_COUNT=$((SUCCESS_COUNT + 1))
    else
        echo "✗ Failed (empty or null response)"
        FAIL_COUNT=$((FAIL_COUNT + 1))
    fi
done

# Close the inner and outer JSON arrays
echo "" >> "$TEMP_FILE"
echo "  ]" >> "$TEMP_FILE"
echo "]" >> "$TEMP_FILE"

# Validate and format JSON
echo ""
echo "Validating and formatting JSON..."
if jq empty "$TEMP_FILE" 2>/dev/null; then
    jq '.' "$TEMP_FILE" > "$OUTPUT_FILE"
    rm "$TEMP_FILE"
    echo "✓ JSON is valid and formatted"
else
    echo "✗ Error: Invalid JSON generated"
    echo "Temp file saved at: $TEMP_FILE"
    exit 1
fi

echo ""
echo "Summary:"
echo "  Total blocks: $((END_BLOCK - START_BLOCK + 1))"
echo "  Successfully fetched: $SUCCESS_COUNT"
echo "  Failed: $FAIL_COUNT"
echo "  Output file (absolute): $(cd "$(dirname "$OUTPUT_FILE")" && pwd)/$(basename "$OUTPUT_FILE")"
echo ""

# If all blocks failed, exit
if [ "$SUCCESS_COUNT" -eq 0 ]; then
    echo "Error: No blocks were successfully fetched"
    exit 1
fi

# Run Rust verification
echo "================================================"
echo "Running Rust verification..."
echo "================================================"
echo ""

cd "$(dirname "$0")/.."

# Update OUTPUT_FILE path to be relative to prover directory
OUTPUT_FILE_FOR_RUST="./testdata/$OUTPUT_FILE"

RUST_LOG=info TRUSTED_SETUP_4844=./configs/4844_trusted_setup.txt cargo run --release -- --block-path "$OUTPUT_FILE_FOR_RUST"

RUST_EXIT_CODE=$?

echo ""
echo "================================================"
if [ $RUST_EXIT_CODE -eq 0 ]; then
    echo "✓ Verification completed successfully!"
else
    echo "✗ Verification failed with exit code $RUST_EXIT_CODE"
fi
echo "================================================"

exit $RUST_EXIT_CODE
