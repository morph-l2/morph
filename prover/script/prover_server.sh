#!/bin/bash

RUST_LOG=info RUST_BACKTRACE=full nohup ./target/release/prover-server >>prover-server.log 2>&1 &