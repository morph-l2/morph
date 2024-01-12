# tx-submitter

### install dependency
`go mod tidy`
### build
`make build`
### run
`make run`

### commands useful for debugging
```bash
# stake
cast send  0x6900000000000000000000000000000000000010 "stake()"  --rpc-url http://127.0.0.1:9545  --private-key 59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d --value 2ether# query submitter
cast call  0x6900000000000000000000000000000000000010 "submitter()" --rpc-url http://127.0.0.1:9545  --private-key 59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
# query contract map
cast call  0x6900000000000000000000000000000000000010 "storageBatchs(uint64)" 1 --rpc-url http://127.0.0.1:9545  --private-key 59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
cast call  0x6900000000000000000000000000000000000010 "lastL2BlockNumber()" -r  http://127.0.0.1:9545
cast call  0xb7f8bc63bbcad18155201308c8f3540b07f84f5e "lastFinalizedBatchIndex()" -r  http://127.0.0.1:9545
cast call  0xb7f8bc63bbcad18155201308c8f3540b07f84f5e "lastCommittedBatchIndex()" -r  http://127.0.0.1:9545
```