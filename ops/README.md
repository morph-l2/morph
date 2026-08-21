# Running a devnet

All commands run from the repo root. Everything is driven by the `Makefile`; the
compose files live in `ops/docker/`.

## Launch modes

| Command | L2 nodes | Consensus | Who produces blocks |
| --- | --- | --- | --- |
| `make devnet-up` | `node-0`, `node-1` | PBFT for a block or two, then upgrades | `node-0` alone, as single sequencer |
| `make devnet-up-cluster` | above **+** `ha-node-0/1/2` | **none — starts in sequencer mode** | the raft leader among `ha-node-0/1/2` |
| `START_IN_SEQUENCER_MODE=false make devnet-up-cluster` | same as above | PBFT for a block or two, then upgrades | the raft leader among `ha-node-0/1/2` |

Reth instead of geth: `make devnet-up-reth` / `make devnet-up-cluster-reth`.
L1 only: `make devnet-l1`.

Both non-cluster and cluster runs bring up L1 (`layer1-el` / `layer1-cl` /
`layer1-vc`), deploy the L1 contracts, generate the L2 genesis, and start
`tx-submitter-0` plus `gas-price-oracle`. First run builds images and takes a
while; later runs reuse them.

### `make devnet-up` — single sequencer

`node-0` is the only genesis validator, so PBFT runs with a validator set of one.
The upgrade to single-sequencer mode is driven by a timestamp:
`DEVNET_SEQUENCER_UPGRADE_OFFSET_SECONDS` (default `0`) sets it to "now" at setup
time, and setup takes minutes, so by the time block 1 is produced the boundary is
already in the past and the node switches immediately. `node-0` then produces
every block; `node-1` follows by replaying L1 batches (`verify_mode=layer1`, no
tendermint).

> Do not raise `DEVNET_SEQUENCER_UPGRADE_OFFSET_SECONDS`. If the upgrade has not
> happened by block 3, `updateSequencerSet` replaces the single-validator set with
> the L1-designated sequencer address, `node-0` loses its vote, and the chain
> deadlocks at `RoundStepPropose` forever. `0` is the only value that is reliably
> safe.

### `make devnet-up-cluster` — HA cluster, no PBFT

Adds three `ha-node` / `ha-geth` pairs that form a hashicorp/raft cluster; the
leader produces blocks. `node-0` keeps running but has no sequencer key in this
mode (`ACTIVE_SEQUENCER_PRIVATE_KEY` is emptied), so it is a plain follower — but
keep it up for the initial start: the `ha-node`s' `persistent_peers` list names
`node-0/1/2`, and `node-0` is the only one of those that runs tendermint, so it is
where they get the peer that the hand-over waits on. Once they are producing it no
longer matters.

PBFT is skipped entirely: the three `ha-node`s default to
`MORPH_NODE_START_IN_SEQUENCER_MODE=true`, which pre-sets the upgrade block
height to 0 so the consensus reactor never starts. Set
`START_IN_SEQUENCER_MODE=false` to run the normal PBFT → upgrade path instead.

`MORPH_NODE_START_IN_SEQUENCER_MODE` is **test-only**. The node refuses to start
if it is set on a production network, or on a node that would be unable to
produce blocks with it (see below).

## Skipping PBFT outside this devnet

For a QA environment that manages its own configuration, the switch alone is not
enough. Three things must hold, and the first one's default is wrong:

1. **`block_sync = true` in `config.toml`.** Tendermint defaults it to `false`
   and morph never overrides it (this devnet flips it in `setup_nodes.py`). With
   block sync off, the hand-over that starts the sequencer routines never runs.
2. **The sequencer must not be the only genesis validator.** Either don't give it
   the genesis `priv_validator_key.json` — tendermint generates a non-genesis key
   when the file is absent — or put two or more validators in genesis. A node
   holding the sole genesis validator key gets block sync disabled and never
   starts producing. The node detects this case and refuses to start.
3. **At least one other node running tendermint**, i.e. **two nodes minimum**. The
   hand-over waits for the block pool to catch up, and a pool with no peers never
   reports caught up. A node with `MORPH_NODE_DERIVATION_VERIFY_MODE=layer1` does
   not start tendermint and does not count.

Only #2 is checked at startup. #1 and #3 fail silently: containers stay up, RPC
answers, nothing is logged, and the block height stays at 0.

## Endpoints

| | RPC | WS | Tendermint RPC |
| --- | --- | --- | --- |
| L1 | `9545` | `9546` | beacon `4000` |
| `morph-el-0` | `8545` | `8546` | `node-0` → `26657` |
| `morph-el-1` | `8645` | `8646` | — |
| `ha-geth-0/1/2` | `9145` / `9245` / `9345` | `9146` / `9246` / `9346` | `27657` / `27757` / `27857` |

`ha-node` admin API: `9501` / `9601` / `9701`. Metrics are on `6060` inside each
geth container (`--metrics.expensive` is enabled, so the `chain/account/*` and
`chain/storage/*` timers are populated).

## Restarting, stopping, cleaning

```sh
make devnet-down          # stop containers, keep all data
make devnet-clean-build   # wipe L1 + L2 data and generated config, keep images
make devnet-clean         # the above, and delete the morph images too
make devnet-l1-clean      # wipe L1 only
make devnet-logs          # follow logs
```

To restart while keeping the chain, use `docker compose restart` or
`down` + `up -d` in `ops/docker/` — **not** `make devnet-up`, which re-runs the
whole bootstrap.

A few things worth knowing before you debug a failed clean:

- **A full clean must include L1.** `layer1/genesis/` is generated, and a stale
  copy leaves the beacon chain stuck at `head_slot=0`. `devnet-clean-build`
  already depends on `devnet-l1-clean`; if you clean by hand, do both.
- **Start Docker first.** Cleaning with the daemon down silently does nothing.
- **`make devnet-up` does not rebuild images.** After changing `go-ethereum` or
  `tendermint` sources, rebuild explicitly or you will keep running the old
  binary.
- **Stale `geth.ipc`.** If an execution client logs
  `IPC opening failed ... operation not supported` and then serves no RPC, delete
  `ops/docker/.devnet/el<N>/geth.ipc` and restart it. The socket cannot be
  re-bound over on a bind-mounted volume.
