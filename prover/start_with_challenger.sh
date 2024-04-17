#!/bin/bash

pkill -9 prover_server
pkill -9 challenge

# Start prover
RUST_LOG=info RUST_BACKTRACE=full nohup ./target/release/prover_server >>prover.log 2>&1 &

# Start handler
RUST_LOG=info RUST_BACKTRACE=full nohup ./challenge-handler/target/release/challenge-handler >>handler.log 2>&1 &

# Start challenger
RUST_LOG=info RUST_BACKTRACE=full nohup ./challenge-handler/target/release/auto_challenge >>challenge.log 2>&1 &