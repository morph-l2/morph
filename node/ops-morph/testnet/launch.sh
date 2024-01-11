#!/bin/bash

export MORPH_NODE_L2_ENGINE_AUTH=ops-morph/jwt-secret.txt
export MORPH_NODE_L1_ETH_RPC=https://eth-goerli.g.alchemy.com/v2/hmGN_z77oy7LlFgPivFuKF893Iog5AOV
export MORPH_NODE_SYNC_DEPOSIT_CONTRACT_ADDRESS=0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9

nohup ./build/bin/morphnode --dev-sequencer --home ./mytestnet/node0 --l2.eth=http://127.0.0.1:8545 --l2.engine=http://127.0.0.1:8551 > ./mytestnet/node0.log 2>&1 &
nohup ./build/bin/morphnode --dev-sequencer --home ./mytestnet/node1 --l2.eth=http://127.0.0.1:8645 --l2.engine=http://127.0.0.1:8651 > ./mytestnet/node1.log 2>&1 &
nohup ./build/bin/morphnode --dev-sequencer --home ./mytestnet/node2 --l2.eth=http://127.0.0.1:8745 --l2.engine=http://127.0.0.1:8751 > ./mytestnet/node2.log 2>&1 &
nohup ./build/bin/morphnode --dev-sequencer --home ./mytestnet/node3 --l2.eth=http://127.0.0.1:8845 --l2.engine=http://127.0.0.1:8851 > ./mytestnet/node3.log 2>&1 &