#!/usr/bin/env bash
#
# Onyx sweep — comprehensive devnet verification script.
#
# Run AFTER the devnet is fully up (make devnet-up-reth completes).
# The CREATE2 factory is already bootstrapped by the devnet startup hook, but this
# script also self-heals it (and deploys the Registry) if missing.
#
# Prerequisites:
#   - foundry (cast, forge) + jq on PATH
#   - devnet running on L2:8545
#
# Deterministic-deployment plumbing lives in scripts/lib/onyx-sweep-common.sh.
set -euo pipefail

export PATH="$HOME/.foundry/bin:$PATH"
cd "$(dirname "${BASH_SOURCE[0]}")/.."   # morph/contracts (foundry root)
# shellcheck source=scripts/lib/onyx-sweep-common.sh
source scripts/lib/onyx-sweep-common.sh

# ---- config -----------------------------------------------------------------
L2_RPC="${L2_RPC:-http://127.0.0.1:8545}"
CHAIN_ID="${L2_CHAIN_ID:-53077}"

ACCT0=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
ACCT0_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
DEPLOYER=0x70997970C51812dc3A010C7d01b50e0d17dc79C8
DEPLOYER_KEY=0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d

# Deposit EOAs are generated fresh on every run (see onyx_new_deposit) so this script
# can be re-run against the same chain: Phase 3.1 registers DEPOSIT and Phase 4.3
# disables it permanently, both of which are one-shot per address. Override to replay a
# specific run against a chain where those addresses are still unused:
#   DEPOSIT=0x.. DEPOSIT_KEY=0x.. DEPOSIT_2=0x.. DEPOSIT_2_KEY=0x.. bash scripts/…
DEPOSIT="${DEPOSIT:-}"
DEPOSIT_KEY="${DEPOSIT_KEY:-}"
DEPOSIT_2="${DEPOSIT_2:-}"
DEPOSIT_2_KEY="${DEPOSIT_2_KEY:-}"

MASTER=$ACCT0
OWNER=$ACCT0

AMOUNT=1000000000000000000  # 1e18

PASS=0; FAIL=0; SKIP=0

# No `return 1` on failure: this runs under `set -e`, so a non-zero return would abort
# the script at the first FAIL and the summary below could never report more than one
# failure (nor print at all). Failures are tallied in $FAIL and surfaced by the exit
# code at the end instead.
run_test() { printf "  %s... " "$1"; shift; if "$@"; then echo "PASS"; PASS=$((PASS+1)); else echo "FAIL"; FAIL=$((FAIL+1)); fi; }
skip_test() { echo "  $1... SKIP ($2)"; SKIP=$((SKIP+1)); }
send0()  { cast send --rpc-url "$L2_RPC" --private-key "$ACCT0_KEY" "$@"; }
call0()  { cast call --rpc-url "$L2_RPC" "$@"; }

onyx_require_tools

# After onyx_require_tools so a missing cast/jq is reported as such rather than as an
# empty keypair.
[ -n "$DEPOSIT" ] && [ -n "$DEPOSIT_KEY" ] || read -r DEPOSIT DEPOSIT_KEY < <(onyx_new_deposit)
[ -n "$DEPOSIT_2" ] && [ -n "$DEPOSIT_2_KEY" ] || read -r DEPOSIT_2 DEPOSIT_2_KEY < <(onyx_new_deposit)

echo "============================================================"
echo "  Onyx Sweep — Comprehensive Devnet Verification"
echo "  Expected Registry: $ONYX_EXPECTED_REGISTRY"
echo "  Deposits (this run): $DEPOSIT  /  $DEPOSIT_2"
echo "============================================================"

# ---- Phase 0: Prerequisites ------------------------------------------------
echo ""
echo "--- Phase 0: Environment ---"
echo "L2 RPC: $L2_RPC  chain: $(cast chain-id --rpc-url "$L2_RPC" 2>/dev/null || echo 'UNREACHABLE')"
[ "$(cast chain-id --rpc-url "$L2_RPC" 2>/dev/null)" = "$CHAIN_ID" ] || {
    echo "FATAL: L2 RPC unreachable or wrong chain (expected $CHAIN_ID)"
    exit 1
}

# ---- Phase 1: CREATE2 factory & Registry deployment ------------------------
echo ""
echo "--- Phase 1: CREATE2 Factory & Registry ---"

# Fund the Registry deployer.
send0 --value 10ether "$DEPLOYER" >/dev/null

run_test "CREATE2 factory present" onyx_ensure_create2_factory "$L2_RPC" "$ACCT0_KEY"

# Fatal precondition: the predicted proxy MUST equal the morph-reth constant, or
# the EL will never find the Registry. Sets ONYX_REGISTRY / initcodes / salts.
onyx_precompute_addresses
REGISTRY="$ONYX_REGISTRY"

run_test "deploy Registry via CREATE2" onyx_deploy_registry "$L2_RPC" "$DEPLOYER_KEY"
run_test "registry code present at expected address" onyx_code_present "$L2_RPC" "$REGISTRY"
run_test "initialize proxy" onyx_initialize_registry "$L2_RPC" "$DEPLOYER_KEY" "$OWNER"
run_test "owner is set correctly" bash -c "
  [ \"\$(cast call --rpc-url '$L2_RPC' '$REGISTRY' 'owner()(address)' | tr 'A-Z' 'a-z')\" = \"\$(echo '$OWNER' | tr 'A-Z' 'a-z')\" ]
"

# ---- Phase 2: Deploy MockERC20 & token setup -------------------------------
echo ""
echo "--- Phase 2: Token Setup ---"

TOKEN=$(onyx_forge_create "$L2_RPC" "$ACCT0_KEY" "$ONYX_MOCK_ERC20" \
  "Sweep Test Token" "STT" 18)
echo "  MockERC20 deployed: $TOKEN"

NON_WHITELIST=$(onyx_forge_create "$L2_RPC" "$ACCT0_KEY" "$ONYX_MOCK_ERC20" \
  "NonSweepToken" "NST" 18)
echo "  Non-whitelist token: $NON_WHITELIST"

send0 "$TOKEN" "mint(address,uint256)" "$ACCT0" "$AMOUNT" >/dev/null
send0 "$NON_WHITELIST" "mint(address,uint256)" "$ACCT0" "$AMOUNT" >/dev/null
send0 "$REGISTRY" "setTokenWhitelist(address,bool)" "$TOKEN" true >/dev/null

# ---- Phase 3: Sweep flow verification -------------------------------------
echo ""
echo "--- Phase 3: Sweep Flow ---"

# Fund deposit for gas headroom (sweep itself is gasless, this is for safety).
send0 --value 1ether "$DEPOSIT" >/dev/null

# 3.1 Register deposit
echo "  [3.1] Register deposit..."
DEADLINE=$(( $(date +%s) + 31536000 ))
TYPED=$(mktemp)
onyx_typed_data "$CHAIN_ID" "$REGISTRY" "$DEPOSIT" "$MASTER" 0 "$DEADLINE" > "$TYPED"
SIG=$(cast wallet sign --private-key "$DEPOSIT_KEY" --data --from-file "$TYPED")
rm -f "$TYPED"

run_test "register deposit" bash -c "
  cast send --rpc-url '$L2_RPC' --private-key '$ACCT0_KEY' \
    '$REGISTRY' 'registerSweep(address,address,uint256,uint64,bytes)' \
    '$DEPOSIT' '$MASTER' 0 '$DEADLINE' \"$SIG\" >/dev/null 2>&1
  RES=\$(cast call --rpc-url '$L2_RPC' '$REGISTRY' 'resolveSweep(address,address)(address)' '$TOKEN' '$DEPOSIT')
  [ \"\$(echo \"\$RES\" | tr 'A-Z' 'a-z')\" = \"\$(echo '$MASTER' | tr 'A-Z' 'a-z')\" ]
"

# 3.2 Transfer whitelisted token -> sweep
run_test "transfer triggers sweep" bash -c "
  send0() { cast send --rpc-url '$L2_RPC' --private-key '$ACCT0_KEY' \"\$@\"; }
  send0 '$TOKEN' 'transfer(address,uint256)' '$DEPOSIT' '$AMOUNT' >/dev/null 2>&1
  DEP_BAL=\$(cast call --rpc-url '$L2_RPC' '$TOKEN' 'balanceOf(address)(uint256)' '$DEPOSIT')
  [ \"\${DEP_BAL%% *}\" = '0' ]
"

run_test "master received sweep amount" bash -c "
  MASTER_BAL=\$(cast call --rpc-url '$L2_RPC' '$TOKEN' 'balanceOf(address)(uint256)' '$MASTER')
  echo \"master_bal=\$MASTER_BAL\"
  [ \"\$(cast to-dec \"\${MASTER_BAL%% *}\")\" -ge '$AMOUNT' ]
"

# 3.3 Swept event in receipt
run_test "Swept event in receipt" bash -c "
  TX=\$(cast send --rpc-url '$L2_RPC' --private-key '$ACCT0_KEY' '$TOKEN' 'transfer(address,uint256)' '$DEPOSIT' '$AMOUNT' --json 2>/dev/null | jq -r .transactionHash)
  LOGS=\$(cast receipt --rpc-url '$L2_RPC' \"\$TX\" --json 2>/dev/null | jq -r '.logs[].topics[0]')
  echo -n \"\$LOGS\" | grep -qi '${ONYX_SWEPT_TOPIC#0x}'
"

# ---- Phase 4: Edge Cases --------------------------------------------------
echo ""
echo "--- Phase 4: Edge Cases ---"

# 4.1 Non-whitelisted token does NOT sweep
send0 "$NON_WHITELIST" "transfer(address,uint256)" "$DEPOSIT" "$AMOUNT" >/dev/null
run_test "non-whitelisted token not swept" bash -c "
  BAL=\$(cast call --rpc-url '$L2_RPC' '$NON_WHITELIST' 'balanceOf(address)(uint256)' '$DEPOSIT')
  echo \"non-wl balance=\$BAL\"
  [ \"\$(cast to-dec \"\${BAL%% *}\")\" -gt 0 ]
"

# 4.2 Multiple transfers in one tx — the TestSweepRouter fixture lives in the
#     morph-reth repo (crates/node/tests/assets), not here, so skip unless present.
if [ -f 'crates/node/tests/assets/SweepFixtures.sol' ]; then
  run_test "multiple transfers in one tx" bash -c "
    ROUTER=\$(onyx_forge_create '$L2_RPC' '$ACCT0_KEY' \
      'crates/node/tests/assets/SweepFixtures.sol:TestSweepRouter' '$REGISTRY' '$TOKEN')
    cast send --rpc-url '$L2_RPC' --private-key '$ACCT0_KEY' '$TOKEN' 'mint(address,uint256)' \"\$ROUTER\" '$AMOUNT' >/dev/null
    cast send --rpc-url '$L2_RPC' --private-key '$ACCT0_KEY' \
      \"\$ROUTER\" 'batchTest(address,address,uint256)' '$REGISTRY' '$DEPOSIT' '$AMOUNT' >/dev/null 2>&1
    BAL=\$(cast call --rpc-url '$L2_RPC' '$TOKEN' 'balanceOf(address)(uint256)' '$DEPOSIT')
    [ \"\${BAL%% *}\" = '0' ]
  "
else
  skip_test "multiple transfers in one tx" "SweepFixtures.sol only in morph-reth repo"
fi

# 4.3 Disabled deposit
echo "  [4.3] Disable sweep..."
send0 "$REGISTRY" "disableSweep(address)" "$DEPOSIT" >/dev/null
run_test "disabled deposit not swept" bash -c "
  send0() { cast send --rpc-url '$L2_RPC' --private-key '$ACCT0_KEY' \"\$@\"; }
  send0 '$TOKEN' 'transfer(address,uint256)' '$DEPOSIT' '$AMOUNT' >/dev/null 2>&1
  BAL=\$(cast call --rpc-url '$L2_RPC' '$TOKEN' 'balanceOf(address)(uint256)' '$DEPOSIT')
  echo \"disabled dep balance=\$BAL\"
  [ \"\$(cast to-dec \"\${BAL%% *}\")\" -gt 0 ]
"

# ---- Phase 5: Security ----------------------------------------------------
echo ""
echo "--- Phase 5: Security ---"

# 5.1 Unauthorized registration (no deposit signature)
run_test "reject registration without deposit sig" bash -c "
  cast send --rpc-url '$L2_RPC' --private-key '$ACCT0_KEY' \
    '$REGISTRY' 'registerSweep(address,address,uint256,uint64,bytes)' \
    '$DEPOSIT_2' '$MASTER' 0 '$DEADLINE' 0x0000 >/dev/null 2>&1 && false || true
"

# 5.2 Code check: attempt to register a contract as deposit
run_test "reject contract-as-deposit" bash -c "
  cast send --rpc-url '$L2_RPC' --private-key '$ACCT0_KEY' \
    '$REGISTRY' 'registerSweep(address,address,uint256,uint64,bytes)' \
    '$TOKEN' '$MASTER' 0 '$DEADLINE' 0x0000 >/dev/null 2>&1 && false || true
"

# 5.3 Double registration protection
run_test "double register fails" bash -c "
  cast send --rpc-url '$L2_RPC' --private-key '$ACCT0_KEY' \
    '$REGISTRY' 'registerSweep(address,address,uint256,uint64,bytes)' \
    '$DEPOSIT' '$MASTER' 1 '$DEADLINE' \"$SIG\" >/dev/null 2>&1 && false || true
"

# ---- Phase 6: Poke Sweep --------------------------------------------------
echo ""
echo "--- Phase 6: Poke Sweep ---"

# Phase 4.3 disabled DEPOSIT permanently (disableSweep is irreversible), and pokeSweep
# reverts with DepositNotActive() on a disabled deposit — so register the fresh
# DEPOSIT_2 here and use it for this phase and Phase 7. ACCT0's TOKEN balance was
# drained into the disabled DEPOSIT in 4.3, so mint a fresh supply too. Registering
# only now keeps Phase 5.1's "reject registration without deposit sig" test meaningful:
# an already-registered DEPOSIT_2 would revert there for the wrong reason.
echo "  [6.0] Register DEPOSIT_2 + mint..."
send0 --value 1ether "$DEPOSIT_2" >/dev/null   # gas headroom, same as DEPOSIT in 3.0
send0 "$TOKEN" "mint(address,uint256)" "$ACCT0" "$AMOUNT" >/dev/null
TYPED2=$(mktemp)
onyx_typed_data "$CHAIN_ID" "$REGISTRY" "$DEPOSIT_2" "$MASTER" 0 "$DEADLINE" > "$TYPED2"
SIG2=$(cast wallet sign --private-key "$DEPOSIT_2_KEY" --data --from-file "$TYPED2")
rm -f "$TYPED2"
send0 "$REGISTRY" "registerSweep(address,address,uint256,uint64,bytes)" \
  "$DEPOSIT_2" "$MASTER" 0 "$DEADLINE" "$SIG2" >/dev/null

run_test "poke sweep triggers SweepRequested" bash -c "
  TX=\$(cast send --rpc-url '$L2_RPC' --private-key '$ACCT0_KEY' \
    '$REGISTRY' 'pokeSweep(address,address)' '$TOKEN' '$DEPOSIT_2' --json 2>/dev/null | jq -r .transactionHash)
  LOGS=\$(cast receipt --rpc-url '$L2_RPC' \"\$TX\" --json 2>/dev/null | jq -r '.logs[].topics[0]')
  echo -n \"\$LOGS\" | grep -qi '${ONYX_REQUEST_TOPIC#0x}'
"

# ---- Phase 7: Fresh deposit end-to-end + receipt Swept log -----------------
echo ""
echo "--- Phase 7: Fresh Deposit + Receipt Swept Log ---"

# DEPOSIT_2 was registered and freshly minted in Phase 6.0; pokeSweep there was a
# no-op sweep (zero balance), so it still has an empty balance to sweep into here.
run_test "fresh deposit sweeps + Swept log in receipt" bash -c "
  TX=\$(cast send --rpc-url '$L2_RPC' --private-key '$ACCT0_KEY' '$TOKEN' 'transfer(address,uint256)' '$DEPOSIT_2' '$AMOUNT' --json 2>/dev/null | jq -r .transactionHash)
  rec_json=\$(cast receipt --rpc-url '$L2_RPC' \"\$TX\" --json 2>/dev/null)
  has_swept=\$(echo \"\$rec_json\" | jq -r '.logs[].topics[0]' | grep -ci '${ONYX_SWEPT_TOPIC#0x}' || true)
  bal=\$(cast call --rpc-url '$L2_RPC' '$TOKEN' 'balanceOf(address)(uint256)' '$DEPOSIT_2')
  echo \"swept log count: \$has_swept  deposit2 balance: \${bal%% *}\"
  [ \"\$has_swept\" -ge 1 ] && [ \"\${bal%% *}\" = '0' ]
"

# ---- Summary ---------------------------------------------------------------
echo ""
echo "============================================================"
echo "  RESULTS:  $PASS passed, $FAIL failed, $SKIP skipped"
echo "============================================================"

[ "$FAIL" -eq 0 ] && echo "ALL TESTS PASSED ✅" || echo "SOME TESTS FAILED ❌"
exit "$FAIL"
