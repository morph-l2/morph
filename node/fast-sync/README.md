# morphnode fast sync build

A `morphnode` build for syncing a full node from genesis as quickly as possible. It only changes how blocks are downloaded and how Tendermint persists them. Commit verification and block execution are the same as in the regular build.

Do not use it on sequencers or validators, and switch back to the regular build once the node has caught up.

## Changes

| Change | Why |
|---|---|
| Tendermint database writes no longer fsync | Tendermint syncs at least three writes per block, which dominates block time on disks with slow fsync, macOS in particular |
| Concurrent block requests per peer are sized from the recent block size to keep about 2 MiB of blocks in flight, between 20 and 300. The regular build allows 20 | Keeps enough requests in flight on high-latency links, and falls back to the regular 20 once blocks are around 100 KB, so large blocks do not queue behind one connection for longer than the 30 s retry |
| Blocks are requested at most 1.5 times the total per-peer allowance ahead of the current height, up to 900 instead of 600 | Requests are handed out in no particular height order, so a much wider window can leave the block the sync is waiting for unrequested |
| While blocks are under 16 KiB, the request for the height the sync is blocked on is also sent to another peer after 700 ms, up to twice | A single slow connection no longer stalls every block behind it. With larger blocks the link is the limit, and a second request mostly downloads the same block twice |
| A block that arrives after its request was retried elsewhere is used, or dropped if a copy is already in, instead of disconnecting the peer. A block that has arrived is not requested again when the 30 s retry fires, and a retry prefers a peer that has not been asked yet | The regular build disconnects the peer for such a block, and throws away a block still waiting to be processed when its retry fires. Once requests wait longer than 30 s, every peer gets disconnected in turn and the sync stops making progress |

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
- On empty blocks each peer receives up to 15 times the usual number of concurrent block requests.

## Results

On a Hoodi sync from genesis on macOS, over the initial empty blocks, the sync rate went from about 95 to about 620 blocks per second. The gain depends on the machine and the network, and blocks with transactions leave less to gain because execution takes a larger share.
