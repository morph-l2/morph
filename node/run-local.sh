#!/bin/bash

# Execute the command
go run cmd/node/main.go \
  --l2.eth=http://127.0.0.1:8045 \
  --l2.engine=http://127.0.0.1:8051 \
  --l2.jwt-secret=../ops/docker/jwt-secret.txt \
  --home=./.devnet \
  --log.level=main:debug,*:info