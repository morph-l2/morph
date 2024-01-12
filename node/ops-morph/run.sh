#!/bin/bash

export MORPH_NODE_L2_ETH_RPC=http://127.0.0.1:8545
export MORPH_NODE_L2_ENGINE_RPC=http://127.0.0.1:8551
export MORPH_NODE_L2_ENGINE_AUTH=jwt-secret.txt
export MORPH_NODE_L1_ETH_RPC=https://eth-goerli.g.alchemy.com/v2/hmGN_z77oy7LlFgPivFuKF893Iog5AOV
export MORPH_NODE_SYNC_DEPOSIT_CONTRACT_ADDRESS=0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9
## export MORPH_NODE_SYNC_START_HEIGHT=88854536

../build/bin/morphnode --sequencer --home ../build