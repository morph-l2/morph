#!/bin/bash

nohup docker build -f Dockerfile.sp1-app -t sp1-evm:latest . > build.log 2>&1 &  
