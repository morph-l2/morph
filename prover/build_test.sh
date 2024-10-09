#!/bin/bash

nohup docker build -f Dockerfile.sp1-test -t sp1-test:latest . > build_test.log 2>&1 &  
