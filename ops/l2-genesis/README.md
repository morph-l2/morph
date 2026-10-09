# L2 genesis generation

The existing network scripts generate and verify L2 genesis artifacts. They do
not initialize contracts, register submitters or start services. Operators must
provide confirmed L1 deployment records and the configuration approved for the
target network. For the complete QA flow, see
[deployment and recovery](../README.md#qa-deployment).

The [centralization specification](../../.vscode/doing/centralization-cleanup-spec.md)
retains L2 predeploys and their state. Generation requires a confirmed
`Proxy__L1Staking` deployment on the selected L1, including devnet and qanet.
The scripts do not deploy that retired contract, substitute `Submitter`, or fill
its address with a placeholder. An empty new L1 therefore does not satisfy the
current genesis prerequisites; a separate L2 genesis design is required to remove
that dependency.

Install the repository's Go toolchain, `jq`, `curl` and `shasum`. Every wrapper
works from any directory and uses the shared implementation in
`devnet-l2genesis.sh`. RPC means the JSON-RPC endpoint of the target L1. Set the
network's environment variable or pass `--l1-rpc`; scripts never load
`contracts/.env` or select an external endpoint automatically.

| Network wrapper | RPC environment variable | Default deployment records |
| --- | --- | --- |
| `devnet-l2genesis.sh` | `L1_RPC_URL` | `.devnet/devnetL1.json` |
| `qanet-l2genesis.sh` | `QA_RPC_URL` | `.qanet/qanetL1.json` |
| `testnet-l2genesis.sh` | `SEPOLIA_RPC_URL` | `../../contracts/sepolia.json` |
| `holesky-l2genesis.sh` | `HOLESKY_RPC_URL` | `../../contracts/holesky.json` |
| `hoodi-l2genesis.sh` | `HOODI_RPC_URL` | `../../contracts/hoodi.json` |
| `mainnet-l2genesis.sh` | `MAINNET_RPC_URL` | `../../contracts/mainnet.json` |

Paths in the table are relative to this directory. Each network defaults to
`deploy-config/<network>-deploy-config.json` and writes to `.<network>`. Use
`--deploy-config`, `--deployment-file`, `--output-dir` and `--l1-rpc` to select
explicit inputs and destinations. For example, from the repository root with
`QA_RPC_URL` configured:

```sh
sh ops/l2-genesis/qanet-l2genesis.sh
sh ops/l2-genesis/qanet-l2genesis.sh --verify-existing
```

The `testnet` and `holesky` wrappers are retained for historical reference.
New genesis generation is disabled for these two networks, including
`--overwrite` and custom `--deploy-config` inputs. Rejection happens before RPC
requests or output changes. `--verify-existing` remains a read-only check for
artifacts using the current `genesis.done` manifest format and matching normalized
inputs; old `done` files do not satisfy that verification.

Their historical configurations contain `l2Sequencer*` fields from an earlier
contract layout. They do not contain the current
`recordOracleAddress`, `recordNextBatchSubmissionIndex`,
`l2StakingSequencerMaxSize`, `l2StakingRewardStartTime`,
`l2StakingUnDelegatedLockEpochs` or `l2StakingAddresses`/`l2StakingTmKeys`/
`l2StakingBlsKeys` fields; `testnet` also lacks `finalSystemOwner`. Their recorded
addresses and keys are retained as historical inputs. These inputs must not be used for a new deployment. Do not copy role addresses,
reward times or staking parameters from another network to make them executable. The current generator
must not be used to replace the genesis of an existing chain.

The devnet, qanet, hoodi and mainnet inputs contain only fields recognized by the
current Go `DeployConfig`. Cleanup removes 20 ignored field occurrences while
preserving every recognized value. The testnet and holesky files remain unchanged
as historical references. File hashes for the four cleaned inputs change; operators
resuming an existing deployment must follow the input recovery instructions below.

## Configuration consumers and field status

These JSON files describe L2 genesis and initial contract state. They do not
generate the L1 chain and are not read directly by running node or geth processes.
The current paths are:

1. `make devnet-up` calls the existing Python launcher, which reads the devnet
   JSON, checks signing roles and saves `genesis-input.json`. The launcher then
   calls `devnet-l2genesis.sh`. Service-management commands also read the source
   JSON to check that it matches the saved deployment.
2. `yarn deploy:qanet` calls `contracts/scripts/localDeploy.sh --network qanet`.
   That script reads the QA JSON, checks its chain IDs against the separate
   Hardhat configuration and calls `qanet-l2genesis.sh`. The shell also supports
   devnet. Submitter batch interval and timeout come from separate explicit
   runtime inputs, not from the `govBatch*` fields in these JSON files.
3. Each network wrapper selects `deploy-config/<network>-deploy-config.json`.
   The shared shell fills L1 addresses from confirmed deployment records, then
   passes the normalized copy through `--deploy-config` to the Go command.
   `NewDeployConfig` in `morph-chain-ops/genesis/config.go` reads that JSON.
4. The Go generator writes `genesis-l2.json`, `rollup.json` and the genesis batch
   header. Devnet geth initializes its database from `genesis-l2.json`; the shell
   exports the header for Hardhat initialization. Current Compose services do
   not mount `rollup.json`; scripts retain and verify it as deployment metadata.
5. A separate consumer exists in `contracts/tasks/proxy_upgrade.ts`:
   `gasOracleProxy-upgrade --l2config FILE` reads the supplied JSON's
   `gasPriceOracleOwner`. Its example names the QA JSON.

`deploy:qanetL1` and `deploy:devnetL1` do not automatically load these JSON files.
Their Hardhat configurations are `contracts/src/deploy-config/qanetl1.ts` and
`contracts/src/deploy-config/l1.ts`, with an optional `DEPLOY_CONFIG_OVERRIDE`.
Publicnode restores a database and uses its separate Tendermint genesis under
`ops/publicnode/holesky`; it does not read this directory's Holesky configuration.

The following counts compare top-level fields in the six tracked JSON inputs
with the current `DeployConfig` structure. "Unrecognized" means no active Go
JSON field exists. Recognition alone does not establish that an address, timestamp
or parameter is approved for a new public network deployment.

| Input | Total fields | Recognized by Go | Unrecognized by Go | Current role |
| --- | ---: | ---: | ---: | --- |
| `devnet-deploy-config.json` | 28 | 28 | 0 | Devnet deployment, services and genesis |
| `qanet-deploy-config.json` | 29 | 29 | 0 | QA deployment and genesis; optional gas oracle upgrade input |
| `testnet-deploy-config.json` | 28 | 16 | 12 | Historical reference; new generation disabled |
| `holesky-deploy-config.json` | 20 | 16 | 4 | Historical reference; new generation disabled |
| `hoodi-deploy-config.json` | 25 | 25 | 0 | Default input for explicit Hoodi genesis generation |
| `mainnet-deploy-config.json` | 23 | 23 | 0 | Default input for explicit mainnet genesis generation |

The removed fields were:

- All four active inputs: `maxTxPerBlock`. The active payload setting is
  `maxTxPayloadBytesPerBlock`: devnet and QA retain `737280`; hoodi and mainnet
  continue to omit it and use the Go default. Transaction count and payload byte
  limits are different quantities, so the old count is not copied into the active
  setting.
- Devnet and QA: `BLOCK_SIGNER_PRIVATE_KEY`, `BLOCK_SIGNER_ADDRESS`,
  `morphTokenName`, `morphTokenSymbol`, `morphTokenOwner`,
  `morphTokenInitialSupply`, `morphTokenDailyInflationRate`. These token fields have
  no `DeployConfig` definition, and `SetImplementations` skips `MorphToken`.
  Current signing identities come from command arguments or environment variables.
- Devnet additionally: `l2StakingPks`. This is distinct from the active public
  staking addresses and keys. The shell removes it and `BLOCK_SIGNER_PRIVATE_KEY`
  from legacy input copies instead of using them to sign transactions.
- QA additionally: `useMPT`. `NewL2Genesis` directly sets `UseZktrie: false`;
  changing this legacy JSON field cannot switch trie implementations.

The remaining 16 unrecognized field occurrences are confined to the two historical
files. Both retain `maxTxPerBlock`; the other names are:

- Testnet: `baseFeeVaultRecipient`, `governanceTokenSymbol`,
  `governanceTokenName`, `governanceTokenOwner`, `stakingSequencerSize`,
  `stakingLockNumber`, `govBatchMaxBytes`, `govBatchMaxChunks` and
  `l2SequencerAddresses`, `l2SequencerTmKeys`, `l2SequencerBlsKeys`.
- Holesky: `l2SequencerAddresses`, `l2SequencerTmKeys`,
  `l2SequencerBlsKeys`.

`NewDeployConfig` retains ordinary `json.Unmarshal` so saved legacy inputs remain
readable. Unknown fields are ignored; the old names are not aliases. The Go tests
strictly decode the four active tracked files and reject unknown fields there,
without changing how previously generated input copies are read.

### Preserved predeploy state and submitter settings

Sections 1.2, 5.3 and 9.2 of the centralization specification preserve the L2
contract implementations, storage and node dependencies. Removing the submitter's
runtime dependency on Gov does not remove `Gov.sol` or its genesis allocation.
The same distinction applies to `Record`, `Sequencer` and `L2Staking`.

| Configuration | Actual consumer | Role after centralization |
| --- | --- | --- |
| `govVotingDuration`, `govBatchBlockInterval`, `govBatchTimeout`, `govRollupEpoch` | `NewL2StorageConfig` writes the corresponding `Gov` storage | Retained genesis values; never used to derive submitter runtime settings |
| `recordOracleAddress`, `recordNextBatchSubmissionIndex` | `NewL2StorageConfig` writes `Record.oracle` and `Record.nextBatchSubmissionIndex` | Retained Record permission and initial index; not the gas oracle signer or submitter owner |
| `l2StakingAddresses` | `Sequencer.initialize` and `L2Staking.initialize` | Initial L2 sequencer/staker state; not L1 submitter registration |
| `l2StakingTmKeys`, `l2StakingBlsKeys` | `L2Staking.initialize`; node reads `getStakesInfo` | Node startup and historical validator reads still require these keys |
| `l2StakingSequencerMaxSize`, `l2StakingUnDelegatedLockEpochs`, `l2StakingRewardStartTime` | `L2Staking.initialize` | Retained contract state; not `Submitter.minimumStake` or withdrawal conditions |
| `l1StakingProxy` | `L2Staking.OTHER_STAKING` immutable | Must match the confirmed legacy L1 deployment, independently of `Proxy__Submitter` |
| `batchSubmitterAddresses` in the L1 TypeScript configuration, or `batchSubmitterPks` | L1 `Submitter` registration and stake tasks | Current batch submitter identities; never copied into L2 staking inputs |
| `--batch-block-interval`, `--batch-timeout` or `TX_SUBMITTER_BATCH_BLOCK_INTERVAL`, `TX_SUBMITTER_BATCH_TIMEOUT` | Deployment identity and submitter environment | Explicit runtime sealing settings, independent of the genesis JSON |

Normal node startup still calls `updateSequencerSet`, reads all three L2 sequencer
sets and retrieves staking information before using the Tendermint keys. The BLS
validation boundary remains in `node/core/sequencers.go`. V2 block authority comes
from L1Sequencer, while historical batch reconstruction also needs the old L2
sequencer state. Renaming `l2StakingAddresses` to submitter addresses would mix
distinct roles and would not remove these code dependencies.

Genesis validation requires nonempty, equal-length staking address/TM/BLS arrays,
and a positive reward time aligned to the contract's 86400-second reward epoch.
It does not rewrite historical key material. Gov and runtime batch validation
allow either interval or timeout to be zero, but reject both being zero. The
runtime values must fit `uint64`; the interval counts L2 blocks and timeout counts
seconds between the first and last L2 block timestamps in a batch.

The role, fee and network fields also remain explicit genesis inputs. The three
fee-recipient fields configure different destinations even when their values
match. Missing L1 addresses are resolved from deployment records.

The launcher no longer writes the unrecognized `l1GenesisBlockTimestamp` into
new devnet input copies. Existing saved inputs remain unchanged. L2 genesis time
comes from `l2GenesisBlockTimestamp`, or the Go generator's current-time default
when that value is absent or zero.
`rollup.json` now records the timestamp and gas limit of the actual generated L2
block, including defaults. It previously used the L1 timestamp and could record a
zero gas limit when the L2 input omitted it. This metadata correction does not
change predeploy storage or rewrite existing artifacts.

### Existing deployments after input cleanup

Cleanup changes source hashes even though the removed fields do not affect Go
genesis generation. The scripts retain their exact input and artifact checks;
they do not migrate saved identities or regenerate existing genesis automatically.
The deployment operator must preserve the output directory and node databases.
If a source hash differs, stop deployment, initialization and service-management
commands until the original input is restored and verification succeeds.

Current deployment identities also record explicit batch parameters and the
selected legacy L1Staking record. Identities created before those fields were
recorded are rejected; restoring the JSON source alone does not migrate them.
Preserve the saved runtime settings, original inputs and artifacts for review.
Do not add fields to an old identity or recompute its hashes manually. Existing
outputs created with the former staking placeholder do not meet the current
specification and must not be used to continue initialization or start a network.

Verification also rejects older `rollup.json` files whose `genesis.l2_time` or
`genesis.system_config.gasLimit` differs from the generated L2 block, even when
the original inputs and stored hashes match. Preserve every original artifact.
Use the recorded generator version for historical inspection and resolve metadata
compatibility separately before reuse. Do not edit the metadata or manifest
hashes, or regenerate an existing chain's genesis to make verification pass.

1. For `make devnet-up` and submitter service commands, restore the source JSON
   used by that deployment at `deploy-config/devnet-deploy-config.json`, using a
   deployment backup or its recorded repository revision. Keep the saved
   `genesis-input.json`, `deployment-state.json` and artifacts unchanged. The
   launcher compares the source's public-field digest; it has no source override
   option. Restoring the removed private-key fields is unnecessary because that
   digest excludes them. Keep the original roles, client, topology and L1
   TypeScript configuration. Do not copy `genesis-input.json` over the tracked
   source: the launcher adds or replaces fields in that generated copy.
2. For `localDeploy.sh`, pass `--deploy-config` with the original source file and
   keep the original output directory, L1 endpoint, accounts and L1 configuration.
   `deployment-identity.json` hashes the source bytes, including formatting. A
   generated `deploy-config.json` is not a substitute for those original bytes.
3. For standalone genesis verification, use the saved normalized input with
   `--verify-existing`. For example, from the repository root with `QA_RPC_URL`
   pointing to the original L1:

   ```sh
   sh ops/l2-genesis/qanet-l2genesis.sh \
     --deploy-config ops/l2-genesis/.qanet/deploy-config.json \
     --deployment-file ops/l2-genesis/.qanet/qanetL1.json \
     --output-dir ops/l2-genesis/.qanet \
     --verify-existing
   ```

The standalone command checks the manifest, artifacts, L1 identity and deployed
code without running Go or writing chain transactions. It does not replace the
full deployment's identity and contract-state checks. If the original input or a
required artifact cannot be recovered, retain the data and investigate the
deployment records. Do not edit saved hashes or use `--overwrite` to bypass a
mismatch.

## Generation and recovery

The scripts check configuration addresses against deployment records and verify
L1 chain identity and deployed code. Every network's records must include the
original `Proxy__L1Staking`; the source must be confirmed independently by the
deployment operator. The full devnet/QA drivers import only that record from
`--legacy-l1-deployment-file` or `LEGACY_L1_DEPLOYMENT_FILE`, or reuse the saved
record on a later run. They do not deploy, administer or register through the
legacy contract. They generate `deploy-config.json`,
`genesis-l2.json`, `rollup.json`, `genesis-batch-header.json` and a
`deployment-config.json` containing `batchHeader`. After checking chain IDs,
header fields and roots, they write `genesis.done` with input and artifact
hashes. This confirms genesis generation only and is separate from the complete
flow's `done`. Tracked TypeScript deployment configurations are never rewritten;
pass the generated override through `DEPLOY_CONFIG_OVERRIDE` when initializing
contracts with Hardhat.

On failure, stop initialization and service startup. Preserve existing files and
the reported `genesis-attempt-*` directory containing inputs and diagnostic logs.
Correct the cause and retry using the same inputs. `--verify-existing` checks saved
files and L1 identity without running Go; use it before reusing generated genesis.
`--overwrite` explicitly regenerates files and must never be used for a running
chain or an existing node database. Remove a stale `.genesis.lock.d` only after
confirming no generation or verification process is running for that directory.
The direct Go command also refuses existing output paths; use the wrappers for
staging and explicit replacement.

The unused Clique L1 Compose example was removed because its mining and account
unlock options do not match the current Ethereum client. Use the maintained
[L1 scripts](../docker/layer1/scripts/start.sh) with the generated execution,
consensus and validator inputs described in the [operations guide](../README.md).
Existing Docker containers, volumes and generated data are not removed by this
repository change.

## Block fixtures

`get_block_fixtures.sh RPC_URL DECIMAL_BLOCK_NUMBER [OUTPUT_DIRECTORY]` downloads
`block.json`, `prestate.json` and the preceding 256 hashes in `block_hashes.json`.
The RPC must support `debug_traceBlockByNumber` with `prestateTracer`. Output
defaults to the current directory. The script rejects existing fixture files,
invalid block responses and RPC errors; incomplete downloads remain in the
reported `fixtures-attempt-*` directory. After resolving the error, retry with
an output directory containing none of the three published fixture files. Remove
a stale `.block-fixtures.lock.d` only after confirming its download process has
stopped.

## Offline validation

Run `make -C ops/l2-genesis test` from the repository root to execute the Go tests
and shell regression tests. The shell tests use simulated RPC responses and Go
output in temporary directories, including all six network wrappers and fixture
failure/retry behavior. They also check that removing legacy fields cannot bypass
saved input hashes and that verification with the original normalized input
preserves artifacts. The Go tests check active configuration field names and
compare actual genesis, rollup and batch-header output with and without legacy
ignored fields using fixed L1 data and a fixed L2 timestamp. Additional tests call
the generated proxy getters to verify the retained Gov, Record, Sequencer and
L2Staking state, and check that rollup metadata matches the actual L2 block.
They do not prove
that a public RPC or deployed network
matches the supplied configuration. `make -C ops/l2-genesis test-scripts` runs
only those script tests. The optional
`tests/validate_local_deployment.py` uses a temporary local Anvil chain to verify
that the shell's devnet and qanet flows reject missing or invalid legacy records without
sending transactions. Its explicit `--legacy-staking-fixture` option also runs
the deployment/genesis/repeat pipeline with synthetic reverting code installed
only in that temporary chain. That fixture tests script integration, not a real
legacy L1Staking implementation or approval for a new network. It never connects
to a public network.
