#!/bin/bash

pkill -9 prover_server
pkill -9 challenge
pkill -9 shadow

# Start prover
RUST_LOG=debug RUST_BACKTRACE=full nohup ./target/release/prover_server >>prover.log 2>&1 &

# Start handler
RUST_LOG=debug RUST_BACKTRACE=full nohup ./challenge-handler/target/release/challenge-handler >>handler.log 2>&1 &

# Start shadow-proving
RUST_LOG=debug RUST_BACKTRACE=full nohup ./shadow-proving/target/release/shadow-proving >>shadow_proving.log 2>&1 &