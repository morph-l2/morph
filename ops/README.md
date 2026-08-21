# Running a devnet

All commands run from the repo root. Everything is driven by the `Makefile`; the
compose files live in `ops/docker/`.

## Launch modes

| Command | L2 nodes | Who produces blocks |
| --- | --- | --- |
| `make devnet-up` | `node-0`, `node-1` | `node-0` alone, as single sequencer |
| `make devnet-up-cluster` | above **+** `ha-node-0/1/2` | the raft leader among `ha-node-0/1/2` |

Reth instead of geth: `make devnet-up-reth` / `make devnet-up-cluster-reth`.
L1 only: `make devnet-l1`.

Both bring up L1 (`layer1-el` / `layer1-cl` / `layer1-vc`), deploy the L1
contracts, generate the L2 genesis, and start `tx-submitter-0` plus
`gas-price-oracle`. First run builds images and takes a while; later runs reuse
them.

Both also start the same way: PBFT runs for a block or two, then the chain
upgrades to single-sequencer mode. The switch is driven by a timestamp —
`DEVNET_SEQUENCER_UPGRADE_OFFSET_SECONDS` (default `0`) sets it to "now" at setup
time, and setup takes minutes, so by the time block 1 is produced the boundary is
already in the past and the upgrade happens immediately.

> Do not raise `DEVNET_SEQUENCER_UPGRADE_OFFSET_SECONDS`. If the upgrade has not
> happened by block 3, `updateSequencerSet` replaces the single-validator set with
> the L1-designated sequencer address, `node-0` loses its vote, and the chain
> deadlocks at `RoundStepPropose` permanently. `0` is the only value that is
> reliably safe.

After the upgrade:

- `make devnet-up` — `node-0` produces every block. `node-1` follows by replaying
  L1 batches (`MORPH_NODE_DERIVATION_VERIFY_MODE=layer1`, so it does not run
  tendermint).
- `make devnet-up-cluster` — the three `ha-node`s form a hashicorp/raft cluster
  and its leader produces. `node-0` has no sequencer key in this mode
  (`ACTIVE_SEQUENCER_PRIVATE_KEY` is emptied), so it becomes a follower.

## Skipping PBFT — QA environments only

`--startInSequencerMode` / `MORPH_NODE_START_IN_SEQUENCER_MODE` makes a node boot
straight into sequencer mode with no PBFT phase at all, by pre-setting the
consensus upgrade block height to 0. It exists so a QA environment can stand up
the post-upgrade shape without waiting for, or configuring, a PBFT phase.

It is **test-only and off by default**. The devnet does not use it — both modes
above go through PBFT — and the node refuses to start if it is set on a
production network.

Enabling it is not sufficient on its own. Four things must hold, and the first
one's default points the wrong way:

1. **`block_sync = true` in `config.toml`.** Tendermint defaults it to `false`
   and morph never overrides it (this devnet only works because
   `setup_nodes.py` rewrites it). With block sync off, the hand-over that starts
   the sequencer routines never runs.
2. **The sequencer must not be the only genesis validator.** Either don't give it
   the genesis `priv_validator_key.json` — tendermint generates a non-genesis key
   when the file is absent — or put two or more validators in genesis. A node
   holding the sole genesis validator key gets block sync disabled and never
   starts producing. This one is checked at startup and refuses to boot.
3. **At least one other node running tendermint**, i.e. **two nodes minimum**. The
   hand-over waits for the block pool to report caught up, and a pool with no
   peers never does. A node with `MORPH_NODE_DERIVATION_VERIFY_MODE=layer1` does
   not start tendermint and does not count.
4. **Only one node may hold the sequencer signing key, unless HA is enabled.**
   Block production is gated on the L1 sequencer contract plus, in HA mode, raft
   leadership — not on tendermint consensus. Two nodes with the same
   `MORPH_NODE_SEQUENCER_PRIVATE_KEY` and no HA will both produce and fork the
   chain. This is easy to hit when satisfying #3 by copying a node's config.

Only #2 is detected at startup. #1, #3 and #4 fail silently — for #1 and #3 the
containers stay up, RPC answers, nothing is logged and the height stays at 0.

## Endpoints

| | RPC | WS | Tendermint RPC |
| --- | --- | --- | --- |
| L1 | `9545` | `9546` | beacon `4000` |
| `morph-el-0` | `8545` | `8546` | `node-0` → `26657` |
| `morph-el-1` | `8645` | `8646` | — |
| `ha-geth-0/1/2` | `9145` / `9245` / `9345` | `9146` / `9246` / `9346` | `27657` / `27757` / `27857` |

`ha-node` admin API: `9501` / `9601` / `9701`. Each geth serves metrics on `6060`
inside its container; add `--metrics.expensive` if you need the
`chain/account/*` and `chain/storage/*` timers, which are otherwise zero.

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
