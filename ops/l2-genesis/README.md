# L2 genesis generation

The existing `devnet-l2genesis.sh` implements generation and verification;
`qanet-l2genesis.sh` selects the QA configuration. These scripts do not initialize
contracts, register submitters or start services. For the complete flow, see
[deployment and recovery](../README.md#qa-deployment).

Install the repository's Go toolchain, `jq`, `curl` and `shasum`. L1 deployment
records must exist, with confirmed contracts on the selected L1. Provide its RPC
endpoint explicitly through `L1_RPC_URL` for devnet, `QA_RPC_URL` for qanet or
`--l1-rpc`. For example, from the repository root with `QA_RPC_URL` configured:

```sh
sh ops/l2-genesis/qanet-l2genesis.sh
sh ops/l2-genesis/qanet-l2genesis.sh --verify-existing
```

Default inputs are `deploy-config/devnet-deploy-config.json` or
`deploy-config/qanet-deploy-config.json`; deployment records are
`.devnet/devnetL1.json` or `.qanet/qanetL1.json`. Paths in this paragraph are relative
to this directory. Use `--deploy-config`, `--deployment-file`, `--output-dir` and
`--l1-rpc` to select explicit inputs and destinations.

The script checks configuration addresses against deployment records and verifies
L1 chain identity and deployed code. It generates `deploy-config.json`,
`genesis-l2.json`, `rollup.json`, `genesis-batch-header.json` and a
`deployment-config.json` containing `batchHeader`. It checks chain IDs, the genesis
header and roots, then writes `genesis.done` with input and artifact hashes. This
marker confirms genesis generation only; it is separate from the complete
flow's `done`.

On failure, stop initialization and service startup. Preserve existing files and
the reported `genesis-attempt-*` directory containing inputs and diagnostic logs.
Correct the cause and retry using the same inputs. `--verify-existing` checks saved
files and L1 identity without running Go; use it before reusing generated genesis.
`--overwrite` explicitly regenerates files and must never be used for a running
chain or an existing node database. Remove a stale `.genesis.lock.d` only after
confirming no generation or verification process is running for that directory.
