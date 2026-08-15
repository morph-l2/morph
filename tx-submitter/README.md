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
cast call $L2_SEQUENCER_ADDR "getSequencerSet2()(address[])" -r $L2RPC
cast call $L1_STAKING_ADDR "getActiveStakers()(address[])" -r $L1RPC
```