#!/bin/bash

# Step 1: Execute npx hardhat compile
echo "Executing npx hardhat compile..."
npx hardhat compile

# Check if the previous command was executed successfully
if [ $? -ne 0 ]; then
    echo "Compilation failed, please check the error message."
    exit 1
fi

# Step 2: Check if the bytecode directory exists, if not, create it
if [ ! -d "./bytecode" ]; then
    echo "The bytecode directory does not exist, creating..."
    mkdir ./bytecode
fi

# Step 3: Copy ShadowRollup.json to the bytecode directory
echo "Copying ShadowRollup.json to the bytecode directory..."
cp artifacts/contracts/ShadowRollup.sol/ShadowRollup.json ./bytecode/

# Check if the copy operation was successful
if [ $? -ne 0 ]; then
    echo "Copy failed, please check the error message."
    exit 1
fi

echo "Operation successfully completed."
