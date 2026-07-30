#!/usr/bin/env bash
#
# Onyx sweep end-to-end demo against a running `make devnet-up-reth`.
#
# Prerequisites:
#   - devnet is up with the Onyx genesis patch applied, i.e. started with:
#       DEVNET_ONYX=1 MORPH_RETH_BUILD_FROM_SOURCE=true make devnet-up-reth
#     (morph-reth must be on branch feat/onyx-recoverable-sweep)
#   - foundry (cast, forge) + jq on PATH
#
# What it does (all on L2):
#   1. deploy impl + TransparentUpgradeableProxy via the deterministic CREATE2
#      factory (network-identical address) then call initialize()
#   2. deploy a MockERC20, whitelist it, register a deposit EOA via EIP-712
#   3. transfer whitelisted tokens INTO the deposit -> the EL auto-sweeps them
#      to master and appends a Swept log
#   4. assert: Swept log present, deposit balance == 0
#
# Deterministic-deployment plumbing (factory raw tx, salts, proxy admin, topics,
# EXPECTED_REGISTRY) lives in scripts/lib/onyx-sweep-common.sh — single source of
# truth shared with onyx-sweep-comprehensive.sh and the devnet startup hook.
set -euo pipefail

export PATH="$HOME/.foundry/bin:$PATH"
cd "$(dirname "${BASH_SOURCE[0]}")/.."   # morph/contracts (foundry root)
# shellcheck source=scripts/lib/onyx-sweep-common.sh
source scripts/lib/onyx-sweep-common.sh

# ---- config -----------------------------------------------------------------
L2_RPC="${L2_RPC:-http://127.0.0.1:8545}"
L2_CHAIN_ID="${L2_CHAIN_ID:-53077}"

# Anvil dev accounts (well-known keys, safe for a local devnet only)
ACCT0=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266   # owner / master / gas provider
ACCT0_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
DEPLOYER=0x70997970C51812dc3A010C7d01b50e0d17dc79C8  # acct1: Registry deployer
DEPLOYER_KEY=0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
# Deposit EOA, generated fresh on every run (see onyx_new_deposit) so the demo can be
# re-run against the same chain: a deposit can be registered exactly once per address
# and disableSweep is terminal. Override to use a specific still-unused EOA:
#   DEPOSIT=0x.. DEPOSIT_KEY=0x.. bash scripts/onyx-sweep-demo.sh
DEPOSIT="${DEPOSIT:-}"
DEPOSIT_KEY="${DEPOSIT_KEY:-}"

MASTER=$ACCT0
OWNER="${REGISTRY_OWNER:-$ACCT0}"
AMOUNT=1000000000000000000  # 1e18

send0() { cast send --rpc-url "$L2_RPC" --private-key "$ACCT0_KEY" "$@"; }

onyx_require_tools

# After onyx_require_tools so a missing cast/jq is reported as such rather than as an
# empty keypair.
[ -n "$DEPOSIT" ] && [ -n "$DEPOSIT_KEY" ] || read -r DEPOSIT DEPOSIT_KEY < <(onyx_new_deposit)

echo "==> L2 RPC $L2_RPC (chainId $L2_CHAIN_ID)"
[ "$(cast chain-id --rpc-url "$L2_RPC" 2>/dev/null)" = "$L2_CHAIN_ID" ] || {
  echo "L2 RPC unreachable or wrong chain (expected $L2_CHAIN_ID)"; exit 1;
}

# ---- 0. precompute CREATE2 addresses (asserts proxy == EXPECTED_REGISTRY) ----
echo "==> [0] precompute CREATE2 addresses"
onyx_precompute_addresses
echo "    (registry $ONYX_REGISTRY == morph-reth hardcoded constant)"

# ---- 1. fund deployer + ensure the CREATE2 factory exists -------------------
echo "==> [1] fund deployer $DEPLOYER + ensure CREATE2 factory"
send0 --value 10ether "$DEPLOYER" >/dev/null
onyx_ensure_create2_factory "$L2_RPC" "$ACCT0_KEY"

# ---- 2. deploy impl + proxy via CREATE2, then initialize --------------------
echo "==> [2] deploy Registry via CREATE2 factory $ONYX_FACTORY"
onyx_deploy_registry "$L2_RPC" "$DEPLOYER_KEY"
onyx_initialize_registry "$L2_RPC" "$DEPLOYER_KEY" "$OWNER"

# ---- 3. deploy MockERC20, whitelist it -------------------------------------
echo "==> [3] deploy MockERC20 + whitelist"
TOKEN=$(onyx_forge_create "$L2_RPC" "$ACCT0_KEY" "$ONYX_MOCK_ERC20" \
  "Onyx Demo" "ODM" 18)
echo "    token = $TOKEN"
send0 "$TOKEN" "mint(address,uint256)" "$ACCT0" "$AMOUNT" >/dev/null
send0 "$ONYX_REGISTRY" "setTokenWhitelist(address,bool)" "$TOKEN" true >/dev/null

pre=$(cast call --rpc-url "$L2_RPC" "$ONYX_REGISTRY" "resolveSweep(address,address)(address)" "$TOKEN" "$DEPOSIT")
echo "    resolveSweep(before register) = $pre"

# ---- 4. deposit signs EIP-712 authorization, owner registers it ------------
echo "==> [4] EIP-712 register deposit $DEPOSIT -> master $MASTER"
DEADLINE=$(( $(date +%s) + 31536000 ))
TYPED=$(mktemp)
onyx_typed_data "$L2_CHAIN_ID" "$ONYX_REGISTRY" "$DEPOSIT" "$MASTER" 0 "$DEADLINE" > "$TYPED"
SIG=$(cast wallet sign --private-key "$DEPOSIT_KEY" --data --from-file "$TYPED")
rm -f "$TYPED"
send0 "$ONYX_REGISTRY" "registerSweep(address,address,uint256,uint64,bytes)" \
  "$DEPOSIT" "$MASTER" 0 "$DEADLINE" "$SIG" >/dev/null
post=$(cast call --rpc-url "$L2_RPC" "$ONYX_REGISTRY" "resolveSweep(address,address)(address)" "$TOKEN" "$DEPOSIT")
echo "    resolveSweep(after register) = $post"

# ---- 5. transfer whitelisted tokens INTO the deposit -> EL auto-sweeps ------
echo "==> [5] transfer $AMOUNT of $TOKEN into deposit -> expect sweep"
TXHASH=$(send0 "$TOKEN" "transfer(address,uint256)" "$DEPOSIT" "$AMOUNT" --json | jq -r .transactionHash)
echo "    tx = $TXHASH"

# ---- 6. assertions ----------------------------------------------------------
echo "==> [6] assert"
LOGS=$(cast receipt --rpc-url "$L2_RPC" "$TXHASH" --json | jq -r '.logs[].topics[0]')
if echo "$LOGS" | grep -qi "${ONYX_SWEPT_TOPIC#0x}"; then
  echo "    OK: Swept log present"
else
  echo "    FAIL: no Swept log in receipt"; echo "$LOGS"; exit 1
fi
dep_bal=$(cast call --rpc-url "$L2_RPC" "$TOKEN" "balanceOf(address)(uint256)" "$DEPOSIT")
master_bal=$(cast call --rpc-url "$L2_RPC" "$TOKEN" "balanceOf(address)(uint256)" "$MASTER")
echo "    deposit balance = $dep_bal (expect 0)"
echo "    master  balance = $master_bal"
[ "${dep_bal%% *}" = "0" ] || { echo "    FAIL: deposit not drained"; exit 1; }

echo ""
echo "SUCCESS: Onyx sweep worked end-to-end on devnet (CREATE2 deployment)."
