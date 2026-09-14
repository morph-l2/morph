# Operational commands

This guide covers devnet and qanet deployment, standalone L1 commands, genesis
generation, existing public nodes and diagnostic tools using the existing scripts.
For QA contracts and services, see [QA deployment](#qa-deployment). For standalone
genesis generation, see the [L2 genesis guide](l2-genesis/README.md).

Examples run from the repository root unless a different directory is stated.
Devnet commands use the root `Makefile`; Compose files live in `ops/docker/`.

Full devnet/QA deployment requires a confirmed existing `Proxy__L1Staking` on the
intended L1. Provide its deployment records through `LEGACY_L1_DEPLOYMENT_FILE`.
The centralization design retains this L2 genesis dependency; the scripts no
longer replace it with a placeholder. On an empty L1, full deployment stops until
that dependency is resolved through an approved legacy source or a separate
genesis change. L1-only commands remain available.

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

Both devnet modes start with PBFT before switching to a single sequencer.
`DEVNET_SEQUENCER_UPGRADE_OFFSET_SECONDS` must be `0`; the launcher rejects any
other value before deployment. It saves the upgrade time before starting L2 and
reuses that value on subsequent attempts.

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
   starts producing.
3. **At least one other node running tendermint**, i.e. **two nodes minimum**. The
   hand-over waits for the block pool to report caught up, and a pool with no
   peers never does. A node with `MORPH_NODE_DERIVATION_VERIFY_MODE=layer1` does
   not start tendermint and does not count.
4. **Only one node may hold the sequencer signing key, unless HA is enabled.**
   Block production is gated on the L1 sequencer contract plus, in HA mode, raft
   leadership — not on tendermint consensus. Two nodes with the same
   `MORPH_NODE_SEQUENCER_PRIVATE_KEY` and no HA will both produce and fork the
   chain. This is easy to hit when satisfying #3 by copying a node's config.

None of the four is detected at startup, and #1 to #3 fail silently in the same
way: the process stays up, RPC answers, nothing is logged, and the block height
stays at 0. If a node in this mode produces nothing, work down this list before
looking anywhere else.

## Endpoints

| | RPC | WS | Tendermint RPC |
| --- | --- | --- | --- |
| L1 | `9545` | `9546` | beacon `4000` |
| `morph-el-0` | `8545` | `8546` | `node-0` → `26657` |
| `morph-el-1` | `8645` | `8646` | — |
| `ha-el-0/1/2` | `9145` / `9245` / `9345` | `9146` / `9246` / `9346` | `27657` / `27757` / `27857` |

`ha-node` admin API: `9501` / `9601` / `9701`.

The execution-layer services are named `el` rather than `geth` because either
client can back them: `docker-compose-cluster.yml` defines `ha-el-*` as geth,
and `docker-compose-cluster-reth.yml` overrides them to reth. That override only
works because the cluster file is layered *before* the reth files — later `-f`
files win, so the reverse order silently leaves the cluster on geth.

The `ha-el-*` reth overrides need their own file rather than a section of
`docker-compose-reth.yml`: compose starts every service a later `-f` file
introduces, whether or not an earlier file declared it, so overrides living in
the shared reth file also came up in the non-cluster devnet — without the
`/genesis.json` and `/jwt-secret.txt` mounts that only the cluster file
provides, which made them exit with `Invalid value '/genesis.json' for --chain`.

## Execution-layer peering

Discovery is off everywhere (`--nodiscover` / `--disable-discovery`), so peers
are configured explicitly and the topology is fixed:

- geth reads `static-nodes.json` (mounted into `morph-el-1`) and, for the
  cluster, `static-nodes-cluster.json` (mounted into all three `ha-el-*`).
- reth ignores those files and takes `--trusted-peers` on the command line.

Both clients derive their identity from the same `nodekey*` / `ha-nodekey*`
files, so a node's enode is the same whichever client is running. reth needs
`--p2p-secret-key` for this; without it, it invents a random identity per
datadir and no peer list can be written in advance. The key files must not have
a trailing newline — reth rejects those with `malformed or out-of-range secret
key`, while geth tolerates them either way.

`morph-el-0` and `morph-el-1` only know each other. The `ha-el-*` nodes dial
both of those plus each other, which keeps `ha-el-*` names out of the
non-cluster setup, where they would not resolve.

## Consensus-layer peering

`setup_nodes.py` writes `persistent_peers` for every tendermint home, deriving
each node ID from the `node_key.json` that ends up installed — which is why the
key files are copied before the peer list is built. Overwriting a
`node_key.json` changes the node's identity, so a hardcoded ID silently goes
stale.

In cluster mode, the list contains `node-0` and the three `ha-node-*`. The ordinary
devnet excludes HA hostnames because those services are not started. `node-1` runs with
`MORPH_NODE_DERIVATION_VERIFY_MODE=layer1` and never starts tendermint, and
`node-2` has no compose service at all; listing either just produces endless
reconnect and DNS failures.

The `ha-node-*` reaching each other matters: the sequencer hand-over waits for
the block pool to report caught up, and a pool whose only peers are unreachable
never does. That is the silent-stall-at-height-0 failure described above.

RPC is served on `0.0.0.0:26657` inside each container so the published ports in
the table above are actually reachable from the host.

Each geth serves metrics on `6060` inside its container, with
`--metrics.expensive` on. Without that flag every counter behind
`metrics.EnabledExpensive` stays zero, which blanks `chain/account/*` and
`chain/storage/*` and also makes `chain/execution` wrong: it is computed as
processing time minus trie time, so with the trie terms at zero it reports all
processing as EVM execution.

Two traps if you set the flag yourself:

- It only works as a bare flag. `--metrics.expensive=true` is silently ignored,
  because `metrics.init()` string-compares `os.Args` before flag parsing. For the
  same reason a TOML config file cannot enable it.
- Its "Enabling expensive metrics collection" log line never appears — `init()`
  runs before the log handler is configured. Check a metric value, not the log.

## Restarting, stopping, cleaning

```sh
make devnet-down          # stop containers, keep all data
make devnet-clean-build   # wipe L1 + L2 data and generated config, keep images
make devnet-clean         # the above, and delete the morph images too
make devnet-l1-clean      # wipe L1 only
make devnet-logs          # follow logs
```

To restart the same chain, rerun the original `make devnet-up` command with the
same client, topology, signing identities and explicit batch parameters. The launcher verifies saved files
and resumes incomplete stages. It reuses completed contracts and verifies genesis
files without regenerating them. It writes `ops/l2-genesis/.devnet/done` only after
L2 reports chain ID `53077` and a block number of at least `1`. An existing `done`
records a previous successful start; it does not establish current service health.

Deployment state, generated contract overrides and public Compose parameters are
kept in `ops/l2-genesis/.devnet/`; node identities and databases are in
`ops/docker/.devnet/`. Private keys are excluded from the state and `runtime.env`.
The launcher passes keys to Hardhat and Compose through the child environment.
The gas oracle uses `L2_GAS_ORACLE_PRIVATE_KEY` or `--gas-oracle-private-key`; when
neither is set, it uses the deployer key only if that address matches
`gasPriceOracleOwner` in the L2 configuration. A mismatch stops deployment before
transactions. No oracle signing key is written to Compose files or `runtime.env`.

Submitter sealing uses `--batch-block-interval` and `--batch-timeout`, or
`TX_SUBMITTER_BATCH_BLOCK_INTERVAL` and `TX_SUBMITTER_BATCH_TIMEOUT`. The devnet
launcher has test defaults of `200` blocks and `600` seconds. It records those
values in `deployment-state.json`; service-management commands reuse the saved
values. It does not read `govBatchBlockInterval` or `govBatchTimeout` for runtime
configuration. These two JSON fields remain inputs to the retained Gov predeploy.

Both the devnet launcher and the QA shell require the independently confirmed
legacy `Proxy__L1Staking` record. Supply `--legacy-l1-deployment-file` or
`LEGACY_L1_DEPLOYMENT_FILE`; later runs can reuse the saved record. The record
must identify an existing contract on the intended L1, not `Submitter` or a
placeholder. Only the selected legacy record is imported; its administration and
operations are not transferred to this deployment. A new empty L1 cannot satisfy
this requirement under the current centralization specification. L1-only startup
remains available through `--only-l1`.

`nodes.done` records the generated node configuration. New node homes are prepared
in `ops/docker/.devnet-setup-*` and published as `.devnet` only after all files pass
validation. Failed attempt directories remain available for inspection; correct the
cause and rerun the same command. Missing state in an existing `.devnet`, changed
identities or changed genesis files stop the operation while preserving existing
data. Restore the original configuration and missing files before retrying.
Cleaning destroys chain data and must not be used to recover an interrupted stage.
Removing ignored fields from a tracked source also changes its saved digest.
For deployments created before configuration cleanup, follow
[input recovery](l2-genesis/README.md#existing-deployments-after-input-cleanup)
before rerunning deployment or starting or rebuilding the submitter service.

`make stop-all-tx-submitter`, `make start-all-tx-submitter` and
`make rebuild-all-tx-submitter` use the existing launcher to operate only
`tx-submitter-0` in the `docker` Compose project. Stop requires no deployment metadata,
generated artifacts, signing key or L1 connection, and does not rewrite `runtime.env`.
It remains available after an interrupted deployment, once the other launcher
operation has released the deployment lock. Start and rebuild use the saved client
and topology and require `BATCH_SUBMITTER_PRIVATE_KEY` to match the
saved submitter identity and verify its active L1 registration. They use
`docker compose up -d --no-deps`; rebuild also adds `--build`. These actions do not
update `done` or verify new batches. The operator must check service logs and
confirmed batch submissions afterward.

A few things worth knowing before you debug a failed clean:

- **A full clean must include L1.** `layer1/genesis/` is generated, and a stale
  copy leaves the beacon chain stuck at `head_slot=0`. `devnet-clean-build`
  already depends on `devnet-l1-clean`; if you clean by hand, do both.
- **Start Docker first.** A Docker stop or removal failure stops cleanup before
  deleting genesis files or node data. Correct the Docker failure and rerun the
  same cleanup command. Only declared devnet volumes and explicitly named local
  service images are removed; unrelated volumes and images remain untouched.
- **`make devnet-up` does not rebuild images.** After changing `go-ethereum` or
  `tendermint` sources, rebuild explicitly or you will keep running the old
  binary.
- **Stale `geth.ipc`.** If an execution client logs
  `IPC opening failed ... operation not supported` and then serves no RPC, delete
  `ops/docker/.devnet/el<N>/geth.ipc` and restart it. The socket cannot be
  re-bound over on a bind-mounted volume.


## QA deployment

The operator must prepare an L1 RPC endpoint reporting chain ID `900`, funded
deployment accounts and the repository's Node.js, Yarn and Go tools, plus `jq`,
`curl` and `shasum`. Install contract dependencies with `yarn install --frozen-lockfile`
and compile with `yarn hardhat compile` in `contracts/` before running the shell
flow. The devnet container flow additionally needs Docker and Foundry's `cast`.

L1 settings come from `contracts/src/deploy-config/l1.ts` or `qanetl1.ts`; L2
settings come from `ops/l2-genesis/deploy-config/devnet-deploy-config.json` or
`qanet-deploy-config.json`. Both must agree on L1 and L2 chain IDs, currently `900`
and `53077`. The scripts use generated JSON overrides instead of editing tracked
TypeScript configuration. This deployment flow supports devnet and qanet;
Solidity contracts are unchanged.
The active L2 JSON inputs omit obsolete fields. Existing deployments still require
their original input hashes; pass the original source through `--deploy-config`
when resuming the shell flow. See
[configuration consumers and recovery](l2-genesis/README.md#configuration-consumers-and-field-status).

Before QA deployment, provide these inputs in the process environment or through
the corresponding shell option:

| Input | Requirement |
| --- | --- |
| `QA_RPC_URL` or `--l1-rpc` | Explicit endpoint for the intended L1 chain. |
| `DEPLOYER_PRIVATE_KEY` | Funded signer for deployment and initialization. |
| `firstSequencerAddress` or `--sequencer-address` | Nonzero address matching the actual block signer. |
| `QA_ROLLUP_DELAY_PERIOD` or `--rollup-delay-period` | Explicit positive integer number of seconds. |
| `TX_SUBMITTER_BATCH_BLOCK_INTERVAL` or `--batch-block-interval` | Explicit `uint64` L2 block count; zero disables this trigger. |
| `TX_SUBMITTER_BATCH_TIMEOUT` or `--batch-timeout` | Explicit `uint64` difference in seconds between the first and last L2 block timestamps in a batch; zero disables this trigger. |
| `LEGACY_L1_DEPLOYMENT_FILE` or `--legacy-l1-deployment-file` | Existing confirmed `Proxy__L1Staking` deployment record on this L1; required unless the output already contains it. |
| `SUBMITTER_OWNER_PRIVATE_KEY` | Required when `submitterOwner` differs from the deployer; must match `Submitter.owner()`. |
| `--config-override FILE` | Optional JSON containing existing L1 configuration fields, including role or submitter addresses. |

Private keys must remain outside JSON overrides, deployment records and
`runtime.env`. The complete shell flow sets `DOTENV_CONFIG_PATH=/dev/null` and does
not obtain deployment inputs from `contracts/.env`. When calling individual
Hardhat tasks, the operator must supply the same endpoint and configuration.

The shell requires both batch values explicitly and rejects them when both are
zero. It never derives them from Gov's genesis values. For an already deployed
network, the release operator must obtain both values from the same L2 block
accepted by the existing release process, and retain its number and hash. The
existing Hardhat task file provides a read-only snapshot command; from `contracts/`,
with `L2_RPC_URL` set to the intended L2 and `L2_SNAPSHOT_BLOCK` set to that decimal
block number. This example uses the devnet/QA L2 chain ID `53077`; another network
requires a Hardhat network configuration with its matching chain ID:

```sh
yarn hardhat read-batch-parameters --network l2 --block-number "$L2_SNAPSHOT_BLOCK"
```

The command returns both values, L2 chain ID, block number and block hash. It
rejects a changed block or values the submitter cannot accept. Record the result
with the deployment inputs; before stopping the old Gov refresh path, repeat the
read at the block selected by the release process and update both working and
backup submitter settings if either value changed. After the transition, running
submitters use only the recorded explicit settings.

From the repository root, run:

```sh
sh contracts/scripts/localDeploy.sh --network qanet
```

The equivalent command in `contracts/` is `yarn deploy:qanet`. The existing
`localDeploy.sh` also accepts `--network devnet` (the default), `--output-dir`,
`--deploy-config`, `--config-override`, `--l1-rpc`, `--sequencer-address` and
`--rollup-delay-period`, plus the batch and legacy-record options above.
Its devnet mode uses `L1_RPC_URL` and requires an already
running L1; use `make devnet-up` to start the local container network. Default shell
outputs are `ops/l2-genesis/.devnet/` or `.qanet/`.

The deployment operator runs these stages through that command:

1. Check the configuration, signer identities and L1 chain identity before sending
   transactions. Save the public deployment parameters in the output directory.
2. Deploy L1 contracts and record each deployment transaction. Generate and verify
   L2 genesis files using the existing network shell script.
3. Initialize proxies, install the first sequencer, import the genesis batch and
   transfer proxy administration to `ProxyAdmin`. This does not transfer every
   contract's ownership to a different QA account.
4. Use the configured `submitterOwner` signer to add submitters and supply their
   minimum stake. Submitter accounts need separate funds for later batch transactions.
5. Verify contracts, then write public `runtime.env` settings and `done`.

On failure, stop subsequent stages and preserve the output directory and chain
state. Resume with the same command, parameters and directory after correcting the
cause. Do not delete deployment records or replace genesis to force progress. For
an unconfirmed deployment, inspect the recorded sender, nonce, expected address and
transaction receipt first. If a broadcast response was lost and no hash was saved,
identify the original transaction and restore its hash before retrying; do not
send another deployment blindly. Remove a stale shell lock only after confirming
that no process is operating on that directory.

A completed rerun verifies the saved genesis and runtime contract state without
sending new deployment transactions. `verify-deployment --runtime` allows
`l2BaseFee` to differ from its initialization value because the oracle updates it;
other configuration and active submitter checks still apply. `done` records the
last successful contract and genesis checks; an existing marker remains after a
failed verification and does not establish current service health. Require a
successful verification command before starting services. QA service startup,
block production and proofs require the separate checks below.

### Starting existing QA services

Before starting services, rerun the completed deployment command with its original
parameters so genesis and contract verification both pass. The operator must use
`genesis-l2.json` for the L2 execution clients and prepare node databases, engine
connections, peer configuration and signing keys. Load the generated public
settings from the repository root:

```sh
set -a
. ops/l2-genesis/.qanet/runtime.env
set +a
export MORPH_NODE_L1_ETH_RPC="$QA_RPC_URL"
export TX_SUBMITTER_L1_ETH_RPC="$QA_RPC_URL"
```

Use the corresponding output path if `--output-dir` was supplied. `QA_RPC_URL` must
be the verified L1 endpoint, including when deployment used `--l1-rpc`. The file
sets the node's sequencer, message-queue and rollup addresses, the submitter's
Submitter and Rollup addresses, batch interval, timeout and V2 activation time.
The operator must additionally provide `MORPH_NODE_L2_ETH_RPC`,
`TX_SUBMITTER_L2_ETH_RPCS` and the existing service-specific connection settings.
The block signer must match `firstSequencerAddress`; `TX_SUBMITTER_L1_PRIVATE_KEY`
must belong to an active configured submitter. Start the existing node and
submitter commands with this environment and do not override the verified
addresses or batch settings.

If using `MORPH_NODE_START_IN_SEQUENCER_MODE`, satisfy the
[QA consensus requirements](#skipping-pbft--qa-environments-only) above. After
startup, check that L2 block height continues increasing, submitter transactions
are confirmed and `Rollup.lastCommittedBatchIndex()` increases. If these checks
fail, stop the affected service, preserve its database and genesis files, correct
connections or signing configuration and restart the same service. Deployment does
not start or validate external `Challenger` or `finalizer` services.

The deployed runtime verifier supports batch version `1`.
`TX_SUBMITTER_BATCH_V2_UPGRADE_TIME=0` leaves runtime V2 disabled; the V2 genesis
header is a separate format. Install and validate the corresponding verifier
before enabling runtime V2 batches.

## Other operational commands

| Scope | Command | Preconditions and result |
| --- | --- | --- |
| Devnet images | `make docker-build` | Docker must be available. Builds the local Go base image and the Compose services selected by `EXECUTION_CLIENT` and `DEVNET_CLUSTER`. `GO_BUILDER_IMAGE` selects the Go builder used by node and submitter builds. |
| Standalone L1 | `make -f ops/docker/Makefile.layer1 start` | Generates missing L1 genesis only when no previous L1 data exists, then starts L1 and waits for a produced block. |
| L1 logs and stop | `make -f ops/docker/Makefile.layer1 logs` or `stop` | Operates on `layer1-el`, `layer1-cl` and `layer1-vc`; does not stop L2 services. |
| L1 replacement | `make -f ops/docker/Makefile.layer1 clean` | Deletes the selected Compose project's L1 containers, declared L1 volumes and generated genesis. Validator key inputs and L2 data remain. A running L2 cannot continue against a replacement L1; preserve its original L1 or deliberately create a separate deployment. |
| Genesis | `sh ops/l2-genesis/<network>-l2genesis.sh` | Requires explicit RPC and confirmed matching L1 deployment records. See the [genesis guide](l2-genesis/README.md) for exact network inputs and verification rules. |
| Diagnostic binaries | `make -C ops/tools build` | Builds `batchparse`, `gasinspect`, `keygen` and `multisend` under `ops/tools/build/bin/`. |

L1 generation requires the checked-in validator definitions and keys. Before
startup, the scripts require nonempty `genesis.json`, `genesis.ssz`, `config.yaml`,
`deposit_contract_block.txt` and the JWT secret. Incomplete existing files stop the
operation; restore the original files instead of regenerating an existing chain.
The unused Clique stack formerly under `ops/l2-genesis/docker-compose/l1` was
removed. Use the standalone L1 commands above for the current execution and
consensus clients.

The `testnet` and `holesky` genesis configurations are historical references.
Their shell entrypoints reject new genesis generation, including `--overwrite`
and attempts to supply another network through the wrapper. `--verify-existing`
only checks existing artifacts that satisfy the current manifest format; it does
not migrate older configurations or create a deployment. Historical configuration
values remain unchanged.

## Running an existing public node

`ops/publicnode` restores an existing snapshot; it does not deploy a new Holesky
network. The checked-in Holesky genesis and peer identities remain historical data.
The operator must establish that the snapshot, current binaries and supplied L1
configuration belong to the same compatible chain before starting a node.

1. Select an environment file with `ENV_FILE`, or use `ops/publicnode/.env`. Set
   `GETH_NETWORK_ID`, `MORPH_NODE_L1_ETH_RPC`, `MORPH_NODE_L1_ETH_BEACON_RPC`,
   `MORPH_NODE_SYNC_DEPOSIT_CONTRACT_ADDRESS`, `MORPH_NODE_L1_SEQUENCER_CONTRACT`,
   `MORPH_NODE_ROLLUP_ADDRESS` and `MORPH_NODE_SEQUENCER_UPGRADE_TIME` using that
   chain's actual parameters. No removed network preset supplies these values.
2. Run `make -C ops/publicnode download-and-decompress-snapshot`. `SNAPSHOT_URL`,
   `SNAPSHOT_NAME` and `SNAPSHOT_DIR` select the archive and destination. Existing
   destinations are preserved; failed download or extraction directories remain
   for inspection. Verify the archive's origin and chain identity before use.
3. Run `make -C ops/publicnode check-config`, then
   `make -C ops/publicnode run-holesky-node`. The latter validates the restored
   data, preserves or creates the JWT secret, builds the services and starts the
   node with its execution client.
4. Inspect service logs, L1/L2 chain IDs, synchronization height and peer
   connectivity. If the chain identity or binary compatibility is wrong, stop the
   services and preserve the snapshot and node data before correcting the inputs.

Use `make -C ops/publicnode stop-holesky-node` to stop both services.
`rm-holesky-node` removes their stopped containers; bind-mounted chain data remains.
Supply the same `ENV_FILE` and path overrides on every command. Offline checks do
not establish that the historical public snapshot endpoint is available or that
the snapshot is compatible with the current binaries.

## Diagnostic and transaction tools

All RPC endpoints are explicit command arguments or documented environment
variables. `gasinspect` and `batchparse` read chain data:

```sh
ops/tools/build/bin/gasinspect -rpc "$L1_RPC_URL" -tx "$TRANSACTION_HASH"
ops/tools/build/bin/batchparse -rpc "$L2_RPC_URL" -batch "$BATCH_INDEX"
```

`gasinspect` decodes the current Rollup method and reports EIP-2028 calldata cost;
that value is not a measurement of contract execution gas. `batchparse` reports a
missing batch as an error. `keygen` creates local key material and prints it to
standard output; keep that output in private storage rather than shared logs.

`multisend` sends transactions. The operator must select the intended test chain,
provide `FUNDING_PRIVATE_KEY` through the environment and choose a private
`-accounts` file before running it. For example, with those inputs already set:

```sh
ops/tools/build/bin/multisend -rpc "$L2_RPC_URL" -chain-id "$L2_CHAIN_ID" \
  -accounts "$SENDER_ACCOUNTS_FILE"
```

Defaults are ten sender accounts, a minimum balance of 10 ETH per sender and one
minute of transfers; inspect `-help` to set `-senders`, `-fund-value` in wei and
`-duration` explicitly. The tool verifies chain ID and genesis across endpoints,
saves new sender keys with private file permissions before funding and waits for
each sender's transaction receipt before reusing its nonce. Keep the account file
and printed transaction hashes. On timeout or interruption, stop creating new
transactions, inspect those transactions and resolve outstanding nonces before
retrying with the same accounts. Do not delete funded sender keys to restart.

## Validation

From the repository root, run the combined operational checks:

```sh
make ops-check
```

This runs tracked Shell syntax and JSON parsing, devnet/L1 command regressions,
all four Compose combinations, publicnode command tests, genesis script tests and
the two operations Go modules with race detection. Compose checks require the
Docker Compose plugin and are reported as skipped if Docker is not installed. Tests
use temporary files and simulated subprocesses or RPC servers; they do not invoke
cleanup on existing deployments or broadcast transactions to public networks.
The individual script suites remain available:

```sh
python3 -m unittest discover -s ops/l2-genesis/tests -v
python3 -m unittest discover -s ops/devnet-morph/tests -v
python3 -m unittest discover -s ops/publicnode/tests -v
```

In `contracts/`, run `yarn typecheck:deployment` and `yarn test:deployment`.
With contracts compiled and Anvil installed, run
`python3 ops/l2-genesis/tests/validate_local_deployment.py` from the root to verify
both networks reject missing or invalid legacy contract records with zero
transactions. Add `--legacy-staking-fixture` to exercise contract deployment,
real Go genesis generation and repeated execution using a synthetic reverting
contract installed only on the temporary Anvil chain. The fixture does not verify
a real legacy L1Staking implementation or establish that a new network meets the
design prerequisites. The harness stops its Anvil process and prints the retained
log directory. These checks do not validate a complete Docker cluster or
external QA services, sustained block production or proof submission.
