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

# Source EOAs are generated fresh on every run (see onyx_new_source) so this script
# can be re-run against the same chain: Phase 3.1 registers SOURCE and Phase 4.3
# disables it permanently, both of which are one-shot per address. Override to replay a
# specific run against a chain where those addresses are still unused:
#   SOURCE=0x.. SOURCE_KEY=0x.. SOURCE_2=0x.. SOURCE_2_KEY=0x.. bash scripts/…
SOURCE="${SOURCE:-}"
SOURCE_KEY="${SOURCE_KEY:-}"
SOURCE_2="${SOURCE_2:-}"
SOURCE_2_KEY="${SOURCE_2_KEY:-}"

DESTINATION=$ACCT0
OWNER=$ACCT0

AMOUNT=1000000000000000000  # 1e18

PASS=0; FAIL=0; SKIP=0

# No `return 1` on failure: this runs under `set -e`, so a non-zero return would abort
# the script at the first FAIL and the summary below could never report more than one
# failure (nor print at all). Failures are tallied in $FAIL and surfaced by the exit
# code at the end instead.
run_test() { printf "  %s... " "$1"; shift; if "$@"; then echo "PASS"; PASS=$((PASS+1)); else echo "FAIL"; FAIL=$((FAIL+1)); fi; }
skip_test() { echo "  $1... SKIP ($2)"; SKIP=$((SKIP+1)); }
call0()  { cast call --rpc-url "$L2_RPC" "$@"; }

# Send a tx from ACCT0 and echo its hash; fail loudly instead of silently.
#
# `cast send` has TWO silent-failure paths, and under a balance-only assertion both
# are indistinguishable from "the tx landed but the sweep did not run":
#   1. The tx never gets mined — gas estimation reverts, or the nonce is taken. The
#      devnet's gas-price-oracle signs L2 txs from ACCT0 too (it shares
#      L2_GAS_ORACLE_PRIVATE_KEY with $ACCT0_KEY) and fires one every ~35s for the
#      first minutes after startup while the L1 base fee decays, so a concurrent send
#      loses the nonce race and is rejected with "replacement transaction
#      underpriced". Exit code is 1 here.
#   2. The tx IS mined but reverts — exit code is still 0; only receipt.status is 0x0.
# The previous `cast send … >/dev/null 2>&1` swallowed both and then compared balances,
# so every assertion expecting "balance == 0" passed spuriously and only 4.3, which
# expects "balance > 0", failed — reporting a nonce race as broken disableSweep logic.
onyx_send_tx() {
  local out hash st
  if ! out=$(cast send --rpc-url "$L2_RPC" --private-key "$ACCT0_KEY" --json "$@" 2>&1); then
    echo "!! cast send failed, tx never mined: $(echo "$out" | tr '\n' ' ' | cut -c1-200)" >&2
    return 1
  fi
  hash=$(echo "$out" | jq -r '.transactionHash // empty')
  st=$(echo "$out" | jq -r '.status // empty')
  [ -n "$hash" ] || { echo "!! cast send returned no tx hash: $out" >&2; return 1; }
  [ "$st" = "0x1" ] || { echo "!! tx $hash reverted (status=$st)" >&2; return 1; }
  echo "$hash"
}
# Exported (with $L2_RPC/$ACCT0_KEY, which the function body resolves at call time)
# so the `run_test … bash -c "…"` subshells can use it. Guarded for the same reason as
# onyx_forge_create in onyx-sweep-common.sh: zsh has no `export -f` and would dump the
# function body to stdout instead of exporting it.
export L2_RPC ACCT0_KEY
[ -n "${BASH_VERSION:-}" ] && export -f onyx_send_tx

onyx_require_tools

# After onyx_require_tools so a missing cast/jq is reported as such rather than as an
# empty keypair.
[ -n "$SOURCE" ] && [ -n "$SOURCE_KEY" ] || read -r SOURCE SOURCE_KEY < <(onyx_new_source)
[ -n "$SOURCE_2" ] && [ -n "$SOURCE_2_KEY" ] || read -r SOURCE_2 SOURCE_2_KEY < <(onyx_new_source)

echo "============================================================"
echo "  Onyx Sweep — Comprehensive Devnet Verification"
echo "  Expected Registry: $ONYX_EXPECTED_REGISTRY"
echo "  Sources (this run): $SOURCE  /  $SOURCE_2"
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
onyx_send_tx --value 10ether "$DEPLOYER" >/dev/null

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

onyx_send_tx "$TOKEN" "mint(address,uint256)" "$ACCT0" "$AMOUNT" >/dev/null
onyx_send_tx "$NON_WHITELIST" "mint(address,uint256)" "$ACCT0" "$AMOUNT" >/dev/null
onyx_send_tx "$REGISTRY" "setTokenWhitelist(address,bool)" "$TOKEN" true >/dev/null

# ---- Phase 3: Sweep flow verification -------------------------------------
echo ""
echo "--- Phase 3: Sweep Flow ---"

# Fund source for gas headroom (sweep itself is gasless, this is for safety).
onyx_send_tx --value 1ether "$SOURCE" >/dev/null

# 3.1 Register source
echo "  [3.1] Register source..."
DEADLINE=$(( $(date +%s) + 31536000 ))
TYPED=$(mktemp)
onyx_typed_data "$CHAIN_ID" "$REGISTRY" "$SOURCE" "$DESTINATION" 0 "$DEADLINE" > "$TYPED"
SIG=$(cast wallet sign --private-key "$SOURCE_KEY" --data --from-file "$TYPED")
rm -f "$TYPED"

run_test "register source" bash -c "
  onyx_send_tx '$REGISTRY' 'registerSweep(address,address,uint256,uint64,bytes)' \
    '$SOURCE' '$DESTINATION' 0 '$DEADLINE' \"$SIG\" >/dev/null || exit 1
  RES=\$(cast call --rpc-url '$L2_RPC' '$REGISTRY' 'resolveSweep(address,address)(address)' '$TOKEN' '$SOURCE')
  [ \"\$(echo \"\$RES\" | tr 'A-Z' 'a-z')\" = \"\$(echo '$DESTINATION' | tr 'A-Z' 'a-z')\" ]
"

# 3.2 Transfer whitelisted token -> sweep. Baseline captured BEFORE the transfer: DESTINATION
# is also the sender here, so the round trip (DESTINATION -AMOUNT to the source, then sweep
# +AMOUNT back) must return it to exactly this value.
DESTINATION_BAL_BEFORE=$(cast to-dec "$(call0 "$TOKEN" 'balanceOf(address)(uint256)' "$DESTINATION" | cut -d' ' -f1)")

# 3.2 records its tx hash here so the destination-balance assertion below can tie itself to
# that specific transfer. Without the handoff the two run_test subshells are independent
# and the balance check cannot distinguish "swept back to destination" from "the transfer
# never happened, so the balance never moved" — both leave it at DESTINATION_BAL_BEFORE.
SWEEP_TX_FILE=$(mktemp)
trap 'rm -f "$SWEEP_TX_FILE"' EXIT

run_test "transfer triggers sweep" bash -c "
  TX=\$(onyx_send_tx '$TOKEN' 'transfer(address,uint256)' '$SOURCE' '$AMOUNT') || exit 1
  echo \"\$TX\" > '$SWEEP_TX_FILE'
  DEP_BAL=\$(cast call --rpc-url '$L2_RPC' '$TOKEN' 'balanceOf(address)(uint256)' '$SOURCE')
  [ \"\${DEP_BAL%% *}\" = '0' ]
"

# Balance is compared for EQUALITY against the pre-transfer baseline, not '>= AMOUNT':
# DESTINATION is ACCT0, the very account the tokens were sent from, so '>= AMOUNT' also holds
# when the tokens never moved — it cannot tell a completed sweep from a no-op. A sweep
# that did not run leaves the balance at baseline-AMOUNT. The Swept-log check pins the
# credit to 3.2's actual transaction rather than to any balance that merely looks right.
run_test "destination received sweep amount" bash -c "
  TX=\$(cat '$SWEEP_TX_FILE' 2>/dev/null)
  [ -n \"\$TX\" ] || { echo 'no sweep tx recorded — 3.2 never landed a transfer'; exit 1; }
  cast receipt --rpc-url '$L2_RPC' \"\$TX\" --json 2>/dev/null \
    | jq -r '.logs[].topics[0]' | grep -qi '${ONYX_SWEPT_TOPIC#0x}' \
    || { echo \"no Swept log in 3.2 tx \$TX\"; exit 1; }
  DESTINATION_BAL=\$(cast to-dec \"\$(cast call --rpc-url '$L2_RPC' '$TOKEN' 'balanceOf(address)(uint256)' '$DESTINATION' | cut -d' ' -f1)\")
  echo \"destination_bal=\$DESTINATION_BAL (baseline before sweep: $DESTINATION_BAL_BEFORE)\"
  [ \"\$DESTINATION_BAL\" = '$DESTINATION_BAL_BEFORE' ]
"

# 3.3 Swept event in receipt
run_test "Swept event in receipt" bash -c "
  TX=\$(onyx_send_tx '$TOKEN' 'transfer(address,uint256)' '$SOURCE' '$AMOUNT') || exit 1
  LOGS=\$(cast receipt --rpc-url '$L2_RPC' \"\$TX\" --json 2>/dev/null | jq -r '.logs[].topics[0]')
  echo -n \"\$LOGS\" | grep -qi '${ONYX_SWEPT_TOPIC#0x}'
"

# ---- Phase 4: Edge Cases --------------------------------------------------
echo ""
echo "--- Phase 4: Edge Cases ---"

# 4.1 Non-whitelisted token does NOT sweep
onyx_send_tx "$NON_WHITELIST" "transfer(address,uint256)" "$SOURCE" "$AMOUNT" >/dev/null
run_test "non-whitelisted token not swept" bash -c "
  BAL=\$(cast call --rpc-url '$L2_RPC' '$NON_WHITELIST' 'balanceOf(address)(uint256)' '$SOURCE')
  echo \"non-wl balance=\$BAL\"
  [ \"\$(cast to-dec \"\${BAL%% *}\")\" -gt 0 ]
"

# 4.2 Multiple transfers in one tx — the TestSweepRouter fixture lives in the
#     morph-reth repo (crates/node/tests/assets), not here, so skip unless present.
if [ -f 'crates/node/tests/assets/SweepFixtures.sol' ]; then
  run_test "multiple transfers in one tx" bash -c "
    ROUTER=\$(onyx_forge_create '$L2_RPC' '$ACCT0_KEY' \
      'crates/node/tests/assets/SweepFixtures.sol:TestSweepRouter' '$REGISTRY' '$TOKEN') || exit 1
    onyx_send_tx '$TOKEN' 'mint(address,uint256)' \"\$ROUTER\" '$AMOUNT' >/dev/null || exit 1
    onyx_send_tx \"\$ROUTER\" 'batchTest(address,address,uint256)' \
      '$REGISTRY' '$SOURCE' '$AMOUNT' >/dev/null || exit 1
    BAL=\$(cast call --rpc-url '$L2_RPC' '$TOKEN' 'balanceOf(address)(uint256)' '$SOURCE')
    [ \"\${BAL%% *}\" = '0' ]
  "
else
  skip_test "multiple transfers in one tx" "SweepFixtures.sol only in morph-reth repo"
fi

# 4.3 Disabled source
echo "  [4.3] Disable sweep..."
onyx_send_tx "$REGISTRY" "disableSweep(address)" "$SOURCE" >/dev/null
# This is the one assertion in the script that expects a NON-zero balance, which made it
# the sole canary for a swallowed send: any send that never landed left the balance at 0
# and got reported here as broken disableSweep logic. onyx_send_tx now separates the two.
run_test "disabled source not swept" bash -c "
  onyx_send_tx '$TOKEN' 'transfer(address,uint256)' '$SOURCE' '$AMOUNT' >/dev/null || exit 1
  BAL=\$(cast call --rpc-url '$L2_RPC' '$TOKEN' 'balanceOf(address)(uint256)' '$SOURCE')
  echo \"disabled dep balance=\$BAL\"
  [ \"\$(cast to-dec \"\${BAL%% *}\")\" -gt 0 ]
"

# ---- Phase 5: Security ----------------------------------------------------
echo ""
echo "--- Phase 5: Security ---"

# 5.1 Unauthorized registration (no source signature)
run_test "reject registration without source sig" bash -c "
  cast send --rpc-url '$L2_RPC' --private-key '$ACCT0_KEY' \
    '$REGISTRY' 'registerSweep(address,address,uint256,uint64,bytes)' \
    '$SOURCE_2' '$DESTINATION' 0 '$DEADLINE' 0x0000 >/dev/null 2>&1 && false || true
"

# 5.2 Code check: attempt to register a contract as source
run_test "reject contract-as-source" bash -c "
  cast send --rpc-url '$L2_RPC' --private-key '$ACCT0_KEY' \
    '$REGISTRY' 'registerSweep(address,address,uint256,uint64,bytes)' \
    '$TOKEN' '$DESTINATION' 0 '$DEADLINE' 0x0000 >/dev/null 2>&1 && false || true
"

# 5.3 Double registration protection
run_test "double register fails" bash -c "
  cast send --rpc-url '$L2_RPC' --private-key '$ACCT0_KEY' \
    '$REGISTRY' 'registerSweep(address,address,uint256,uint64,bytes)' \
    '$SOURCE' '$DESTINATION' 1 '$DEADLINE' \"$SIG\" >/dev/null 2>&1 && false || true
"

# ---- Phase 6: Poke Sweep --------------------------------------------------
echo ""
echo "--- Phase 6: Poke Sweep ---"

# Phase 4.3 disabled SOURCE permanently (disableSweep is irreversible), and pokeSweep
# reverts with SourceNotActive() on a disabled source — so register the fresh
# SOURCE_2 here and use it for this phase and Phase 7. ACCT0's TOKEN balance was
# drained into the disabled SOURCE in 4.3, so mint a fresh supply too. Registering
# only now keeps Phase 5.1's "reject registration without source sig" test meaningful:
# an already-registered SOURCE_2 would revert there for the wrong reason.
echo "  [6.0] Register SOURCE_2 + mint..."
onyx_send_tx --value 1ether "$SOURCE_2" >/dev/null   # gas headroom, same as SOURCE in 3.0
onyx_send_tx "$TOKEN" "mint(address,uint256)" "$ACCT0" "$AMOUNT" >/dev/null
TYPED2=$(mktemp)
onyx_typed_data "$CHAIN_ID" "$REGISTRY" "$SOURCE_2" "$DESTINATION" 0 "$DEADLINE" > "$TYPED2"
SIG2=$(cast wallet sign --private-key "$SOURCE_2_KEY" --data --from-file "$TYPED2")
rm -f "$TYPED2"
onyx_send_tx "$REGISTRY" "registerSweep(address,address,uint256,uint64,bytes)" \
  "$SOURCE_2" "$DESTINATION" 0 "$DEADLINE" "$SIG2" >/dev/null

run_test "poke sweep triggers SweepRequested" bash -c "
  TX=\$(onyx_send_tx '$REGISTRY' 'pokeSweep(address,address)' '$TOKEN' '$SOURCE_2') || exit 1
  LOGS=\$(cast receipt --rpc-url '$L2_RPC' \"\$TX\" --json 2>/dev/null | jq -r '.logs[].topics[0]')
  echo -n \"\$LOGS\" | grep -qi '${ONYX_REQUEST_TOPIC#0x}'
"

# ---- Phase 7: Fresh source end-to-end + receipt Swept log -----------------
echo ""
echo "--- Phase 7: Fresh Source + Receipt Swept Log ---"

# SOURCE_2 was registered and freshly minted in Phase 6.0; pokeSweep there was a
# no-op sweep (zero balance), so it still has an empty balance to sweep into here.
run_test "fresh source sweeps + Swept log in receipt" bash -c "
  TX=\$(onyx_send_tx '$TOKEN' 'transfer(address,uint256)' '$SOURCE_2' '$AMOUNT') || exit 1
  rec_json=\$(cast receipt --rpc-url '$L2_RPC' \"\$TX\" --json 2>/dev/null)
  has_swept=\$(echo \"\$rec_json\" | jq -r '.logs[].topics[0]' | grep -ci '${ONYX_SWEPT_TOPIC#0x}' || true)
  bal=\$(cast call --rpc-url '$L2_RPC' '$TOKEN' 'balanceOf(address)(uint256)' '$SOURCE_2')
  echo \"swept log count: \$has_swept  source2 balance: \${bal%% *}\"
  [ \"\$has_swept\" -ge 1 ] && [ \"\${bal%% *}\" = '0' ]
"

# ---- Phase 8: SweepFailed protocol log ------------------------------------
echo ""
echo "--- Phase 8: SweepFailed Protocol Log ---"

# A reportable sweep failure is hard to reach on a live chain: most classifications
# need a misbehaving token. FailOnSweepERC20 returns false from `transfer` for
# exactly one caller — the source — which is how the EL issues the sweep. So a
# plain mint into the source triggers a sweep that fails as `transfer_false`,
# producing the on-chain failure record while leaving the balance untouched.
read -r SOURCE_3 SOURCE_3_KEY < <(onyx_new_source)
FAIL_TOKEN_BC=$(jq -r '.bytecode.object' forge-artifacts/FailOnSweepERC20.sol/FailOnSweepERC20.json)
FAIL_TOKEN=$(cast send --rpc-url "$L2_RPC" --private-key "$ACCT0_KEY" \
  --create "${FAIL_TOKEN_BC}$(cast abi-encode 'x(address)' "$SOURCE_3" | cut -c3-)" --json \
  | jq -r .contractAddress)
echo "  fail-on-sweep token: $FAIL_TOKEN   source: $SOURCE_3"

onyx_send_tx "$REGISTRY" "setTokenWhitelist(address,bool)" "$FAIL_TOKEN" true >/dev/null
TYPED3=$(mktemp)
onyx_typed_data "$CHAIN_ID" "$REGISTRY" "$SOURCE_3" "$DESTINATION" 0 "$DEADLINE" > "$TYPED3"
SIG3=$(cast wallet sign --private-key "$SOURCE_3_KEY" --data --from-file "$TYPED3")
rm -f "$TYPED3"
onyx_send_tx "$REGISTRY" "registerSweep(address,address,uint256,uint64,bytes)" \
  "$SOURCE_3" "$DESTINATION" 0 "$DEADLINE" "$SIG3" >/dev/null

run_test "failed sweep appends SweepFailed with hashed reason" bash -c "
  TX=\$(onyx_send_tx '$FAIL_TOKEN' 'mint(address,uint256)' '$SOURCE_3' '$AMOUNT') || exit 1
  rec=\$(cast receipt --rpc-url '$L2_RPC' \"\$TX\" --json 2>/dev/null)
  log=\$(echo \"\$rec\" | jq -r '.logs[] | select(.topics[0] == \"$ONYX_SWEEP_FAILED_TOPIC\")')
  [ -n \"\$log\" ] || { echo '    no SweepFailed log in receipt'; exit 1; }

  got_emitter=\$(echo \"\$log\" | jq -r .address | tr 'A-Z' 'a-z')
  got_reason=\$(echo \"\$log\" | jq -r .data)
  bal=\$(cast call --rpc-url '$L2_RPC' '$FAIL_TOKEN' 'balanceOf(address)(uint256)' '$SOURCE_3')
  echo \"    emitter=\$got_emitter reason=\$got_reason source_bal=\${bal%% *}\"

  [ \"\$got_emitter\" = \"\$(echo $REGISTRY | tr 'A-Z' 'a-z')\" ] || { echo '    wrong emitter'; exit 1; }
  [ \"\$(echo \"\$log\" | jq -r '.topics[1]')\" = \"\$(onyx_address_topic '$FAIL_TOKEN')\" ] || { echo '    wrong token topic'; exit 1; }
  [ \"\$(echo \"\$log\" | jq -r '.topics[2]')\" = \"\$(onyx_address_topic '$SOURCE_3')\" ] || { echo '    wrong source topic'; exit 1; }
  [ \"\$(echo \"\$log\" | jq -r '.topics[3]')\" = \"\$(onyx_address_topic '$DESTINATION')\" ] || { echo '    wrong destination topic'; exit 1; }
  [ \"\$got_reason\" = \"\$(cast keccak 'transfer_false')\" ] || { echo '    reason must be keccak256(transfer_false)'; exit 1; }
  [ \"\${bal%% *}\" = '$AMOUNT' ] || { echo '    a failed sweep must not move tokens'; exit 1; }
"

# ---- Summary ---------------------------------------------------------------
echo ""
echo "============================================================"
echo "  RESULTS:  $PASS passed, $FAIL failed, $SKIP skipped"
echo "============================================================"

[ "$FAIL" -eq 0 ] && echo "ALL TESTS PASSED ✅" || echo "SOME TESTS FAILED ❌"
exit "$FAIL"
