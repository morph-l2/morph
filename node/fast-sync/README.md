# morphnode fast sync build

A `morphnode` build for syncing a full node from genesis as quickly as possible. It only changes how blocks are downloaded and how Tendermint persists them. Commit verification and block execution are the same as in the regular build.

Do not use it on sequencers or validators, and switch back to the regular build once the node has caught up.

## Changes

| Change | Why |
|---|---|
| Tendermint database writes no longer fsync | Tendermint syncs at least three writes per block, which dominates block time on disks with slow fsync, macOS in particular |
| Up to 300 concurrent block requests per peer and 900 in total, instead of 20 and 600 | Keeps enough requests in flight on high-latency links |
| The request for the height the pool is blocked on is also sent to another peer after 700 ms, up to twice | A single slow connection no longer stalls every block behind it |

## Build

```bash
cd node
make morphnode-fast-sync
```

The binary is `build/bin/morphnode`, the same path as the regular build, so whichever build ran last is the one in place. The patches in this directory are applied to copies of the dependency sources through a Go build overlay, so `go.mod` and the regular `make morphnode` build are unchanged.

## Run

Use the same flags as the regular `morphnode`, and set this in `config.toml`:

```toml
flush_throttle_timeout = "10ms"
```

## Risks

- A power loss or OS crash can corrupt the Tendermint database, which then has to be synced again. A crash of the process alone is safe.
- Each peer receives up to 15 times the usual number of concurrent block requests.

## Results

On a Hoodi sync from genesis on macOS, over the initial empty blocks, the sync rate went from about 95 to about 620 blocks per second. The gain depends on the machine and the network, and blocks with transactions leave less to gain because execution takes a larger share.
