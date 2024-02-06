# Morph prover
Generate zk proof for the challenged batch.

## Functional Overview
In general design, for the challenge event of emit on the L1 chain, we have two processes, challenge-handler and prover-server, to handle it.

The prover-server receives the prove request and generates chunk proof and batch proof.

The hallenge-handler detects L1 challenge events, then sends a prove request to the prover, queries the prove results, and finally calls the L1 rollup contract to complete the verify.

## Requirements
* rust 1.70.0
* golang 1.20
* python 3.10.6

## download kzg param
`./download_params.sh`

## build
`git clone https://github.com/morphism-labs/prover.git`  
`cd prover`  
`make build`  

## run
`make run`
