#!/bin/bash

RUST_LOG=debug RUST_BACKTRACE=full nohup ./target/release/shadow-proving >>shadow-proving.log 2>&1 &