# tx-submitter

### install dependency
`go mod tidy`
### build
`make build`
### run
`make run`

### commands useful for debugging
```bash
cast call $ROLLUP_ADDR "lastFinalizedBatchIndex()" -r  $L1RPC
cast call $ROLLUP_ADDR "lastCommittedBatchIndex()" -r  $L1RPC
cast call $SUBMITTER_ADDR "isActive(address)(bool)" $WALLET_ADDR -r $L1RPC
cast call $ROLLUP_ADDR "batchBlobVersionedHashes(uint256)(bytes32)" $BATCH_INDEX -r $L1RPC
```
