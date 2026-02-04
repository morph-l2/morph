#!/bin/bash

RUST_LOG=info RUST_BACKTRACE=full nohup ../../target/release/shadow-proving >>shadow-proving.log 2>&1 &