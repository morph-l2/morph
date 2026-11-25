# Morph L2 on native processes, with public Ethereum Hoodi as L1

Run a private Morph L2 (chain id 53077) **without Docker**: every component is a native process, and the L1 is the
public Ethereum **Hoodi** testnet (chain id 560048), so you do not run an L1 yourself. The Rollup contracts are
deployed on Hoodi, and L2 batches are submitted there as EIP-4844 blob transactions.

```
 public Hoodi: execution RPC (ethpandaops, fallbacks 0xrpc.io / publicnode) + beacon RPC
        ▲
   l1-proxy (127.0.0.1:18546, scripts/rpc-consistency-proxy.js): every process reaches L1 through it
        ▲
 ┌──────┴───────────────────────────────────────────────────────────────────────┐
 │ morph-el-0 ◀▶ node-0   block producer (tendermint, single validator, sequencer key)
 │ morph-el-1 ◀▶ node-1   "layer1" verify mode: re-derives the chain from L1 blobs, no tendermint
 │ morph-el-2 ◀▶ node-2   full node with tendermint, no sequencer key; required so node-0 can restart
 │ tx-submitter           submits batches to the Rollup on Hoodi and finalizes them
 │ gas-price-oracle       writes L1 fees into the L2 GasPriceOracle
 │ challenger (optional)  HTTP service: compares batch roots with el-0 and sends Rollup.challengeState on a mismatch
 └──────────────────────────────────────────────────────────────────────────────┘
```

Why it is shaped like this, and every problem met on the way: `notes/issues.md`. How to operate it: `RUNBOOK.md`.

## 1. What this changes in the Morph repo

Four small, already-committed changes on this branch (everything else is in this directory):

- `contracts/hardhat.config.ts`: a `hoodidev` network (chain 560048, RPC from `HOODIDEV_RPC_URL`, key from `DEPLOYER_PRIVATE_KEY`)
- `contracts/src/deploy-config/hoodidev.ts`: deploy config for it
- `ops/l2-genesis/deploy-config/hoodidev-deploy-config.json`: L2 genesis config for it

No owner, admin, fee-recipient or challenger address is fixed in these two files. `deploy-l1.sh` derives them from the keys in
`env/secrets.env` and passes them in (`MORPH_DEPLOYER_ADDRESS`, `MORPH_CHALLENGER_ADDRESS`, `MORPH_BATCH_SUBMITTER_ADDRESS` for `hoodidev.ts`;
a `REPLACE_WITH_DEPLOYER_ADDRESS` placeholder in the genesis config, filled into `run/l1/hoodidev-deploy-config.generated.json`).
Owner, admin and fee recipient are **the deployer**, so whoever runs the deployment holds them; the challenger, sequencer, batch
submitter and gas oracle are separate generated keys. The sequencer set, Tendermint and BLS keys in the configs are fixed.
- `contracts/src/deploy-utils.ts`: optional read retries (`MORPH_READ_RETRIES`), used because public RPCs are read-after-write inconsistent

## 2. What you need

| | |
|---|---|
| OS / CPU / disk | Linux x86_64; ~10 GB free for builds (Rust target dir ~1.3 GB, node_modules ~0.7 GB) and chain data |
| Build tools | Go ≥ 1.24 (tested 1.25.7), Rust (tested 1.95; needs OpenSSL headers + pkg-config), Node + yarn 1.x (tested Node 24, yarn 1.22), make, gcc |
| Runtime tools | Node ≥ 18, Python ≥ 3.9, curl, Foundry `cast` and `anvil` (tested 1.9.0) |
| Network | outbound HTTPS to the Hoodi RPC and beacon endpoints (no inbound ports) |
| **Hoodi ETH** | **about 1.5 ETH in your deployer account**: ~0.07 for gas, 1 ETH is passed on to the batch submitter (which then spends ~0.09 ETH/day), 0.2 ETH to the challenger (`CHALLENGER_FUND_ETH`). Faucet: https://hoodi.ethpandaops.io |
| A deployer key | one private key you control and fund. Other keys are generated for you |

If your tools are not on `PATH` by default, put the `PATH`/build settings in a file and `export MORPH_ENV_FILE=/path/to/that/file`;
every script sources it first (`scripts/lib/env.sh`) and prints exactly what is missing if a tool is absent.

Ports used on 127.0.0.1 (stop other local devnets first): 18546, 8545/8546/8551, 8645/8646/8651, 8745/8746/8751,
30303-30305, 26656/26657/26660, 26666/26667/26670, 27660, 6160-6162, 6070, 8060, and 8090 / 6065 if you run the challenger
(`CHALLENGER_HTTP_PORT` / `CHALLENGER_METRICS_PORT`; not its defaults 8080 / 6060, which are commonly taken).

## 3. From zero to a running chain

All commands from this directory (`ops/native-hoodi/`).

```bash
# 1. build all binaries into ./bin (about 5-10 minutes the first time)
scripts/build-all.sh

# 2. keys: your funded deployer key (it becomes the owner of everything; the public Anvil key is refused) + four generated ones (sequencer, batch submitter, gas oracle, challenger). Prints
#    addresses only. Then fund the DEPLOYER address.
DEPLOYER_PRIVATE_KEY=0x... scripts/init-secrets.sh

# 3. (optional, free) rehearse the whole L1 deployment on a local anvil fork of Hoodi (~5 minutes when the RPC behaves).
#    It reads state from the public RPC, so a hiccup there can make it hang or fail to start: stop it (Ctrl-C) and run it
#    again, or point it at another node with FORK_URL=https://...
scripts/dryrun-fork.sh

# 4. deploy the L1 contracts to Hoodi and generate the L2 genesis. IRREVERSIBLE, spends testnet ETH. About 20-45 minutes
#    (public RPCs are slow; it is resumable: re-running skips finished steps)
source env/secrets.env
RPC_URL="$L1_ETH_RPC" RUN_DIR="$PWD/run/l1" WAIT_FINALIZED=1 SUBMITTER_FUND_ETH=1 scripts/deploy-l1.sh
RPC_URL="$L1_ETH_RPC" scripts/verify-l1.sh        # reads the chain: expect "53 passed, 0 failed"
#    deploy-l1.sh also runs step s34: it adds the challenger to Rollup.isChallenger and tops it up to 0.2 ETH. The step is
#    skipped when env/secrets.env has no CHALLENGER_PRIVATE_KEY; for an older deployment run scripts/register-challenger.sh.

# 5. create the node directories, start everything, bootstrap L2 balances and the gas-oracle permission
scripts/morph-native.sh setup
scripts/morph-native.sh up
scripts/l2-bootstrap.sh
scripts/morph-native.sh status

# 6. (optional) the challenger, see section 4a. It is not started by a plain `up` unless CHALLENGER_ENABLE=1 is set in env/secrets.env
scripts/morph-native.sh up challenger
```

Do not trust a hardhat task's exit code: they often print an error and exit 0. `verify-l1.sh` reads the actual on-chain state.

## 4. Check that it works

```bash
scripts/morph-native.sh status                                  # 9 processes running; el-0 and el-2 heights equal and rising
cast block-number --rpc-url http://127.0.0.1:8545               # run twice a few seconds apart: it grows
H=$(( $(cast block-number --rpc-url http://127.0.0.1:8645) - 1 ))   # node-1 lags up to one batch; compare below its tip
for p in 8545 8645 8745; do cast block $H --field hash --rpc-url http://127.0.0.1:$p; done   # three equal hashes
# batches on L1 (ROLLUP = Proxy__Rollup address in run/l1/deployment/hoodidev.json); both numbers keep rising
cast call $ROLLUP 'lastCommittedBatchIndex()(uint256)' --rpc-url "$L1_ETH_RPC"
cast call $ROLLUP 'lastFinalizedBatchIndex()(uint256)' --rpc-url "$L1_ETH_RPC"
```

### 4a. The challenger (optional)

The challenger is the `challenger/` service of this repo (built into `bin/challenger/challenger` by `scripts/build-all.sh`).
It is an HTTP service, not a watcher: it acts only when it is asked.

- `GET /batch/{index}` reads the batch's commit transaction on L1 and the same block on el-0 and compares the state root and
  the withdraw root. `"verdict":"match"` while the chain is honest; `mismatch` or `inconclusive` otherwise. It sends nothing.
- `POST /challenge` with `{"batchIndex":N,"batchHash":"0x..."}` sends `Rollup.challengeState` **only** when that exact batch
  hash is the one committed on L1 and the roots differ. Everything else is refused with a 4xx and no transaction.
- It listens on `127.0.0.1:8090` (metrics `127.0.0.1:6065`) with auth off, reads the L1 through the l1-proxy and the L2 from
  el-0. el-0 is an archive node (`--gcmode=archive`), which the withdraw-root read needs.
- The account is `CHALLENGER_PRIVATE_KEY` from `env/secrets.env`, a new key per `init-secrets.sh` run (it is not the deployer).
  `deploy-l1.sh` (step s34) or `scripts/register-challenger.sh` adds it to `Rollup.isChallenger` and funds it.

```bash
scripts/morph-native.sh up challenger        # needs bin/challenger/challenger and the key above
scripts/morph-native.sh status               # lists challenger as [optional]
curl -s http://127.0.0.1:8090/health         # {"status":"ok"}
FIN=$(cast call $ROLLUP 'lastFinalizedBatchIndex()(uint256)' --rpc-url "$L1_ETH_RPC")
curl -s http://127.0.0.1:8090/batch/$FIN          # 409 "batch already finalized": finalized batches are refused
curl -s http://127.0.0.1:8090/batch/$((FIN + 1))  # "verdict":"match" for a committed, not yet finalized batch (404 if there is none yet)
```

A successful challenge is a real transaction that locks the challenge deposit and stops finalization until the Rollup owner
calls `revertBatch`; once `proofWindow` has passed, the submitter that committed the batch can be slashed. Run it only on
purpose. All options are in `challenger/env.example`.

End-to-end tests (each uses a few mETH; read the header of each script):

| script | what it proves | time |
|---|---|---|
| `scripts/deposit-test.sh` | L1 → L2 deposit arrives on both L2 nodes | ~8 min |
| `scripts/withdraw-test.sh` | L2 → L1 withdrawal: Merkle proof built locally, relayed on L1 | ~10 min |
| `scripts/restart-test.sh` (then `restart-test.sh T7`) | crash/restart recovery of every component | ~35 min |
| `scripts/rpc-consistency-proxy.selftest.js`, `scripts/withdraw-proof.selftest.js` | offline unit checks | seconds |

## 5. Daily use, reset, export

| | |
|---|---|
| `scripts/morph-native.sh status|up|down|logs NAME|kill NAME|start NAME` | operate the stack; `down` keeps chain data |
| `scripts/reset-env.sh plan|run|restore|list` | move run data (and with `--with-build` the binaries, with `--new-keys` the keys) into `run/archive/`; nothing is deleted. A new L2 genesis needs a new L1 deployment |
| `scripts/register-challenger.sh` | add `CHALLENGER_PRIVATE_KEY`'s address to `Rollup.isChallenger` and top it up (idempotent; `deploy-l1.sh` runs it as step s34) |
| `scripts/sweep-submitter.sh` | send the submitter's remaining Hoodi ETH back to the deployer (stack must be down) |
| `scripts/export-deployment.sh NAME` | save contract addresses, genesis files and test outputs (no private keys) to `deployments/NAME/` |
| `scripts/check-no-secrets.sh` | run before committing: fails if anything looks like a private key |

`bin/`, `run/`, `env/secrets.env`, `env/archive/` and `deployments/` are git-ignored. **Never commit or paste private keys.**

## 6. Limits worth knowing

- Do not remove node-2: after a restart node-0 only resumes block production if it has a tendermint peer (`notes/issues.md` I-028).
- Resetting the L2 means redeploying L1: the Rollup on Hoodi is bound to this chain's batches. Old deployments stay on Hoodi as orphan contracts.
- Deposit gas limits must be generous (use ≥ 400000), see `RUNBOOK.md` section 7.
- The challenger does not poll: nothing challenges automatically. Something has to `POST /challenge`.
- Not covered: HA sequencer cluster, reth, multi-day soak, several withdrawals in one batch, the challenger with an external signer.
- L2 chain id 53077 is shared with other Morph-style devnets; do not join them.
