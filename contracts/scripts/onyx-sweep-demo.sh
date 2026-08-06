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
#   2. deploy a MockERC20, whitelist it
#   3. the controller sets its single destination pointer (route), then a source
#      EOA signs an EIP-712 SweepAuthorization binding itself to that controller,
#      and the controller registers it
#   4. transfer whitelisted tokens INTO the source -> the EL auto-sweeps them
#      to the controller's current destination and appends a Swept log
#   5. assert: Swept log present, source balance == 0
#
# Controller model: the source's signature binds a *controller*, not a concrete
# address. The controller owns the destination pointer (setSweepDestination) and
# may move it without new source signatures; here ACCT0 doubles as owner,
# destination and controller for a compact demo. Override CONTROLLER to exercise
# the separated-keys layout.
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
ACCT0=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266   # owner / destination / gas provider
ACCT0_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
DEPLOYER=0x70997970C51812dc3A010C7d01b50e0d17dc79C8  # acct1: Registry deployer
DEPLOYER_KEY=0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
# Source EOA, generated fresh on every run (see onyx_new_source) so the demo can be
# re-run against the same chain: a source can be registered exactly once per address
# and disableSweep is terminal. Override to use a specific still-unused EOA:
#   SOURCE=0x.. SOURCE_KEY=0x.. bash scripts/onyx-sweep-demo.sh
SOURCE="${SOURCE:-}"
SOURCE_KEY="${SOURCE_KEY:-}"

DESTINATION=$ACCT0
# Route controller: holds the destination pointer and is who the source signs
# for. Defaults to ACCT0 (owner/destination) for a compact demo; override to
# exercise separated keys (a production controller must be secured like the
# destination itself — the Onyx spec §11.2).
CONTROLLER="${CONTROLLER:-$ACCT0}"
CONTROLLER_KEY="${CONTROLLER_KEY:-$ACCT0_KEY}"
OWNER="${REGISTRY_OWNER:-$ACCT0}"
AMOUNT=1000000000000000000  # 1e18

send0() { cast send --rpc-url "$L2_RPC" --private-key "$ACCT0_KEY" "$@"; }
send_controller() { cast send --rpc-url "$L2_RPC" --private-key "$CONTROLLER_KEY" "$@"; }

onyx_require_tools

# After onyx_require_tools so a missing cast/jq is reported as such rather than as an
# empty keypair.
[ -n "$SOURCE" ] && [ -n "$SOURCE_KEY" ] || read -r SOURCE SOURCE_KEY < <(onyx_new_source)

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

pre=$(cast call --rpc-url "$L2_RPC" "$ONYX_REGISTRY" "resolveSweep(address,address)(address)" "$TOKEN" "$SOURCE")
echo "    resolveSweep(before register) = $pre"

# ---- 4. controller points its route at the destination ----------------------
echo "==> [4] controller $CONTROLLER sets destination $DESTINATION"
send_controller "$ONYX_REGISTRY" "setSweepDestination(address)" "$DESTINATION" >/dev/null

# ---- 5. source signs EIP-712 authorization, controller registers it --------
echo "==> [5] EIP-712 register source $SOURCE -> controller $CONTROLLER (dest $DESTINATION)"
DEADLINE=$(( $(date +%s) + 31536000 ))
TYPED=$(mktemp)
onyx_typed_data "$L2_CHAIN_ID" "$ONYX_REGISTRY" "$SOURCE" "$CONTROLLER" "$DEADLINE" > "$TYPED"
SIG=$(cast wallet sign --private-key "$SOURCE_KEY" --data --from-file "$TYPED")
rm -f "$TYPED"
send_controller "$ONYX_REGISTRY" "registerSweep(address,address,uint64,bytes)" \
  "$SOURCE" "$CONTROLLER" "$DEADLINE" "$SIG" >/dev/null
post=$(cast call --rpc-url "$L2_RPC" "$ONYX_REGISTRY" "resolveSweep(address,address)(address)" "$TOKEN" "$SOURCE")
echo "    resolveSweep(after register) = $post"

# ---- 6. transfer whitelisted tokens INTO the source -> EL auto-sweeps ------
echo "==> [6] transfer $AMOUNT of $TOKEN into source -> expect sweep"
TXHASH=$(send0 "$TOKEN" "transfer(address,uint256)" "$SOURCE" "$AMOUNT" --json | jq -r .transactionHash)
echo "    tx = $TXHASH"

# ---- 7. assertions ----------------------------------------------------------
echo "==> [7] assert"
LOGS=$(cast receipt --rpc-url "$L2_RPC" "$TXHASH" --json | jq -r '.logs[].topics[0]')
if echo "$LOGS" | grep -qi "${ONYX_SWEPT_TOPIC#0x}"; then
  echo "    OK: Swept log present"
else
  echo "    FAIL: no Swept log in receipt"; echo "$LOGS"; exit 1
fi
dep_bal=$(cast call --rpc-url "$L2_RPC" "$TOKEN" "balanceOf(address)(uint256)" "$SOURCE")
destination_bal=$(cast call --rpc-url "$L2_RPC" "$TOKEN" "balanceOf(address)(uint256)" "$DESTINATION")
echo "    source balance = $dep_bal (expect 0)"
echo "    destination  balance = $destination_bal"
[ "${dep_bal%% *}" = "0" ] || { echo "    FAIL: source not drained"; exit 1; }

echo ""
echo "SUCCESS: Onyx sweep worked end-to-end on devnet (CREATE2 deployment)."
