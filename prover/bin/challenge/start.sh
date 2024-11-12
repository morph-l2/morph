#!/bin/bash

RUST_LOG=debug RUST_BACKTRACE=full nohup ./target/release/challenge-handler >>handler.log 2>&1 &