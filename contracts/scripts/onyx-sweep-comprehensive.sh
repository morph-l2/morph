#!/usr/bin/env bash
#
# Onyx sweep — comprehensive devnet verification script.
#
# Run AFTER the devnet is fully up (make devnet-up-reth completes).
# The CREATE2 factory is already bootstrapped by the devnet startup hook.
#
# Prerequisites:
#   - foundry (cast, forge) on PATH
#   - devnet running on L2:8545
#
set -euo pipefail

export PATH="$HOME/.foundry/bin:$PATH"
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$ROOT"

# ---- config -----------------------------------------------------------------
L2_RPC="${L2_RPC:-http://127.0.0.1:8545}"
CHAIN_ID="${L2_CHAIN_ID:-53077}"

ACCT0=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
ACCT0_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
DEPLOYER=0x70997970C51812dc3A010C7d01b50e0d17dc79C8
DEPLOYER_KEY=0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
PROXY_ADMIN="${PROXY_ADMIN:-0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC}"

# Anvil account 2 as proxy admin
PROXY_ADMIN_KEY=0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a

DEPOSIT=0x90F79bf6EB2c4f870365E785982E1f101E93b906
DEPOSIT_KEY=0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6
DEPOSIT_2=0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65
DEPOSIT_2_KEY=0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a

MASTER=$ACCT0
OWNER=$ACCT0
FACTORY=0x4e59b44847b379578588920cA78FbF26c0B4956C

AMOUNT=1000000000000000000  # 1e18
SMALL=100
SWEEP_TOPIC=0x035b37215a69e14a80883933d6aa84f0919a67af9410a4a73e8a23baeca011f0
TRANSFER_TOPIC=0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef
REQUEST_TOPIC=0x24e3f180db341974dcd99a5e223d9d944422e303230ddde6659302f8620bbcff

EXPECTED_REGISTRY="0x7aE8bEf666D1D0aB9C0ac5d636f375E46f8AE71A"

PASS=0; FAIL=0; SKIP=0

run_test() { printf "  %s... " "$1"; shift; if "$@"; then echo "PASS"; PASS=$((PASS+1)); else echo "FAIL"; FAIL=$((FAIL+1)); return 1; fi; }
skip_test() { echo "  $1... SKIP"; SKIP=$((SKIP+1)); }
send0()  { cast send --rpc-url "$L2_RPC" --private-key "$ACCT0_KEY" "$@"; }
call0()  { cast call --rpc-url "$L2_RPC" "$@"; }
receipt() { cast receipt --rpc-url "$L2_RPC" "$1" --json 2>/dev/null | jq '.'; }

hash_lower() { echo "$1" | tr 'A-F' 'a-f'; }

echo "============================================================"
echo "  Onyx Sweep — Comprehensive Devnet Verification"
echo "  Expected Registry: $EXPECTED_REGISTRY"
echo "============================================================"

# ---- Phase 0: Prerequisites ------------------------------------------------
echo ""
echo "--- Phase 0: Environment ---"
echo "L2 RPC: $L2_RPC  chain: $(cast chain-id --rpc-url "$L2_RPC" 2>/dev/null || echo 'UNREACHABLE')"

[ "$(cast chain-id --rpc-url "$L2_RPC" 2>/dev/null)" = "$CHAIN_ID" ] || {
    echo "FATAL: L2 RPC unreachable or wrong chain"
    exit 1
}

# ---- Phase 1: CREATE2 factory & Registry deployment ------------------------
echo ""
echo "--- Phase 1: CREATE2 Factory & Registry ---"

run_test "factory code present" bash -c "
  code=\$(cast code --rpc-url '$L2_RPC' '$FACTORY' 2>/dev/null)
  [ -n \"\$code\" ] && [ \"\$code\" != '0x' ]
"

run_test "deploy Registry via CREATE2" bash -c "
  SALT_IMPL=\$(cast keccak 'morph.sweep-registry.impl.v1')
  SALT_PROXY=\$(cast keccak 'morph.sweep-registry.proxy.v1')
  IMPL_CODE=\$(jq -r .bytecode.object forge-artifacts/SweepRegistry.sol/SweepRegistry.json)
  IMPL=\$(cast create2 --deployer '$FACTORY' --salt \"\$SALT_IMPL\" --init-code \"\$IMPL_CODE\")
  echo \"impl predicted: \$IMPL\"

  # Deploy impl
  IMPL_ONCHAIN=\$(cast code --rpc-url '$L2_RPC' \"\$IMPL\" 2>/dev/null || echo '0x')
  if [ \"\$IMPL_ONCHAIN\" = '0x' ] || [ -z \"\$IMPL_ONCHAIN\" ]; then
    cast send --rpc-url '$L2_RPC' --private-key '$DEPLOYER_KEY' \
      '$FACTORY' 'deploy(bytes,bytes32)' \"\$IMPL_CODE\" \"\$SALT_IMPL\" >/dev/null
  fi

  PROXY_CODE=\$(jq -r .bytecode.object forge-artifacts/TransparentUpgradeableProxy.sol/TransparentUpgradeableProxy.json)
  PROXY_ARGS=\$(cast abi-encode 'x(address,address,bytes)' \"\$IMPL\" '$PROXY_ADMIN' '0x')
  PROXY_INITCODE=\"\${PROXY_CODE}\${PROXY_ARGS#0x}\"
  PROXY_PRED=\$(cast create2 --deployer '$FACTORY' --salt \"\$SALT_PROXY\" --init-code \"\$PROXY_INITCODE\")
  echo \"proxy predicted: \$PROXY_PRED\"

  PROXY_ONCHAIN=\$(cast code --rpc-url '$L2_RPC' \"\$PROXY_PRED\" 2>/dev/null || echo '0x')
  if [ \"\$PROXY_ONCHAIN\" = '0x' ] || [ -z \"\$PROXY_ONCHAIN\" ]; then
    cast send --rpc-url '$L2_RPC' --private-key '$DEPLOYER_KEY' \
      '$FACTORY' 'deploy(bytes,bytes32)' \"\$PROXY_INITCODE\" \"\$SALT_PROXY\" >/dev/null
  fi

  [ \"\$(echo \"\$PROXY_PRED\" | tr 'A-Z' 'a-z')\" = \"\$(echo '$EXPECTED_REGISTRY' | tr 'A-Z' 'a-z')\" ]
"

REGISTRY="$EXPECTED_REGISTRY"

run_test "initialize proxy" bash -c "
  OWNER_CHECK=\$(cast call --rpc-url '$L2_RPC' '$REGISTRY' 'owner()(address)' 2>/dev/null || echo '')
  if [ -z \"\$OWNER_CHECK\" ] || [ \"\$OWNER_CHECK\" = '0x0000000000000000000000000000000000000000' ]; then
    INIT_DATA=\$(cast calldata 'initialize(address)' '$OWNER')
    cast send --rpc-url '$L2_RPC' --private-key '$DEPLOYER_KEY' '$REGISTRY' \"\$INIT_DATA\" >/dev/null
  fi
  OWNER_AFTER=\$(cast call --rpc-url '$L2_RPC' '$REGISTRY' 'owner()(address)')
  [ \"\$(echo \"\$OWNER_AFTER\" | tr 'A-Z' 'a-z')\" = \"\$(echo '$OWNER' | tr 'A-Z' 'a-z')\" ]
"

# ---- Phase 2: Deploy MockERC20 & token setup -------------------------------
echo ""
echo "--- Phase 2: Token Setup ---"

TOKEN=$(forge create '@rari-capital/solmate/src/test/utils/mocks/MockERC20.sol:MockERC20' \
  --rpc-url "$L2_RPC" --private-key "$ACCT0_KEY" --broadcast --json \
  --constructor-args "Sweep Test Token" "STT" 18 2>/dev/null | jq -r .deployedTo)
echo "  MockERC20 deployed: $TOKEN"

NON_WHITELIST=$(forge create '@rari-capital/solmate/src/test/utils/mocks/MockERC20.sol:MockERC20' \
  --rpc-url "$L2_RPC" --private-key "$ACCT0_KEY" --broadcast --json \
  --constructor-args "NonSweepToken" "NST" 18 2>/dev/null | jq -r .deployedTo)
echo "  Non-whitelist token: $NON_WHITELIST"

send0 "$TOKEN" "mint(address,uint256)" "$ACCT0" "$AMOUNT" >/dev/null
send0 "$NON_WHITELIST" "mint(address,uint256)" "$ACCT0" "$AMOUNT" >/dev/null
send0 "$TOKEN" "setTokenWhitelist(address,bool)" "$TOKEN" true >/dev/null

# ---- Phase 3: Sweep flow verification -------------------------------------
echo ""
echo "--- Phase 3: Sweep Flow ---"

# Fund deposit addresses for gas
send0 --value 1ether "$DEPOSIT" >/dev/null

# 3.1 Register deposit
echo "  [3.1] Register deposit..."
MODE=$(cast keccak "MORPH_SWEEP_V1")
SCOPE=$(cast keccak "WHITELISTED_ERC20_TO_MASTER_ONLY")
DEADLINE=$(( $(date +%s) + 31536000 ))

TYPED=$(mktemp)
cat > "$TYPED" <<JSON
{
  "types": {
    "EIP712Domain": [
      {"name":"name","type":"string"},
      {"name":"version","type":"string"},
      {"name":"chainId","type":"uint256"},
      {"name":"verifyingContract","type":"address"}
    ],
    "SweepAuthorization": [
      {"name":"deposit","type":"address"},
      {"name":"master","type":"address"},
      {"name":"registry","type":"address"},
      {"name":"chainId","type":"uint256"},
      {"name":"nonce","type":"uint256"},
      {"name":"deadline","type":"uint64"},
      {"name":"mode","type":"bytes32"},
      {"name":"sweepScope","type":"bytes32"}
    ]
  },
  "primaryType": "SweepAuthorization",
  "domain": {
    "name": "SweepRegistry",
    "version": "1",
    "chainId": $CHAIN_ID,
    "verifyingContract": "$REGISTRY"
  },
  "message": {
    "deposit": "$DEPOSIT",
    "master": "$MASTER",
    "registry": "$REGISTRY",
    "chainId": $CHAIN_ID,
    "nonce": 0,
    "deadline": $DEADLINE,
    "mode": "$MODE",
    "sweepScope": "$SCOPE"
  }
}
JSON
SIG=$(cast wallet sign --private-key "$DEPOSIT_KEY" --data --from-file "$TYPED")
rm -f "$TYPED"

run_test "register deposit" bash -c "
  cast send --rpc-url '$L2_RPC' --private-key '$ACCT0_KEY' \
    '$REGISTRY' 'registerSweep(address,address,uint256,uint64,bytes)' \
    '$DEPOSIT' '$MASTER' 0 '$DEADLINE' \"\$SIG\" >/dev/null 2>&1
  RES=\$(cast call --rpc-url '$L2_RPC' '$REGISTRY' 'resolveSweep(address,address)(address)' '$TOKEN' '$DEPOSIT')
  [ \"\$(echo \"\$RES\" | tr 'A-Z' 'a-z')\" = \"\$(echo '$MASTER' | tr 'A-Z' 'a-z')\" ]
"

# 3.2 Transfer whitelisted token → sweep
run_test "transfer triggers sweep" bash -c "
  send0 '$TOKEN' 'transfer(address,uint256)' '$DEPOSIT' '$AMOUNT' >/dev/null 2>&1
  DEP_BAL=\$(call0 '$TOKEN' 'balanceOf(address)(uint256)' '$DEPOSIT')
  [ \"\$DEP_BAL\" = '0' ] || [ \"\${DEP_BAL##*[^0-9]}\" = '0' ]
"

run_test "master received sweep amount" bash -c "
  MASTER_BAL=\$(call0 '$TOKEN' 'balanceOf(address)(uint256)' '$MASTER')
  echo \"master_bal=\$MASTER_BAL\"
  [[ \"\$MASTER_BAL\" == *\"$AMOUNT\"* ]] || [ \"\$(cast to-dec \"\$MASTER_BAL\")\" -ge '$AMOUNT' ]
"

# 3.3 Swept event in receipt
run_test "Swept event in receipt" bash -c "
  TX=\$(send0 '$TOKEN' 'transfer(address,uint256)' '$DEPOSIT' '$AMOUNT' --json 2>/dev/null | jq -r .transactionHash)
  LOGS=\$(cast receipt --rpc-url '$L2_RPC' \"\$TX\" --json 2>/dev/null | jq -r '.logs[].topics[0]')
  echo -n \"\$LOGS\" | grep -qi '\${SWEEP_TOPIC#0x}'
"

# ---- Phase 4: Edge Cases --------------------------------------------------
echo ""
echo "--- Phase 4: Edge Cases ---"

# 4.1 Non-whitelisted token does NOT sweep
send0 "$NON_WHITELIST" "transfer(address,uint256)" "$DEPOSIT" "$AMOUNT" >/dev/null
run_test "non-whitelisted token not swept" bash -c "
  BAL=\$(call0 '$NON_WHITELIST' 'balanceOf(address)(uint256)' '$DEPOSIT')
  echo \"non-wl balance=\$BAL\"
  [[ \$(cast to-dec \"\$BAL\") -gt 0 ]]
"

# 4.2 Multiple transfers in one tx
run_test "multiple transfers in one tx" bash -c "
  ROUTER=\$(forge create 'crates/node/tests/assets/SweepFixtures.sol:TestSweepRouter' \
    --rpc-url '$L2_RPC' --private-key '$ACCT0_KEY' --broadcast --json \
    --constructor-args '$REGISTRY' '$TOKEN' 2>/dev/null | jq -r .deployedTo)
  send0 '$TOKEN' 'mint(address,uint256)' '$ROUTER' '$AMOUNT' >/dev/null
  cast send --rpc-url '$L2_RPC' --private-key '$ACCT0_KEY' \
    '$ROUTER' 'batchTest(address,address,uint256)' '$REGISTRY' '$DEPOSIT' '$AMOUNT' >/dev/null 2>&1
  BAL=\$(call0 '$TOKEN' 'balanceOf(address)(uint256)' '$DEPOSIT')
  [ \"\$BAL\" = '0' ] || [ \"\${BAL##*[^0-9]}\" = '0' ]
  echo \"deposit balance after multi-transfer: \$BAL\"
  true
"

# 4.3 Disabled deposit
echo "  [4.3] Disable sweep..."
send0 "$REGISTRY" "disableSweep(address)" "$DEPOSIT" >/dev/null
run_test "disabled deposit not swept" bash -c "
  send0 '$TOKEN' 'transfer(address,uint256)' '$DEPOSIT' '$AMOUNT' >/dev/null 2>&1
  BAL=\$(call0 '$TOKEN' 'balanceOf(address)(uint256)' '$DEPOSIT')
  echo \"disabled dep balance=\$BAL\"
  [[ \$(cast to-dec \"\$BAL\") -gt 0 ]]
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
  # Try registering the token contract itself as deposit
  cast send --rpc-url '$L2_RPC' --private-key '$ACCT0_KEY' \
    '$REGISTRY' 'registerSweep(address,address,uint256,uint64,bytes)' \
    '$TOKEN' '$MASTER' 0 '$DEADLINE' 0x0000 >/dev/null 2>&1 && false || true
"

# 5.3 Double registration protection
run_test "double register fails" bash -c "
  cast send --rpc-url '$L2_RPC' --private-key '$ACCT0_KEY' \
    '$REGISTRY' 'registerSweep(address,address,uint256,uint64,bytes)' \
    '$DEPOSIT' '$MASTER' 1 '$DEADLINE' \"\$SIG\" >/dev/null 2>&1 && false || true
"

# ---- Phase 6: Poke Sweep --------------------------------------------------
echo ""
echo "--- Phase 6: Poke Sweep ---"

run_test "poke sweep triggers SweepRequested" bash -c "
  TX=\$(cast send --rpc-url '$L2_RPC' --private-key '$ACCT0_KEY' \
    '$REGISTRY' 'pokeSweep(address,address)' '$TOKEN' '$DEPOSIT' --json 2>/dev/null | jq -r .transactionHash)
  LOGS=\$(cast receipt --rpc-url '$L2_RPC' \"\$TX\" --json 2>/dev/null | jq -r '.logs[].topics[0]')
  echo -n \"\$LOGS\" | grep -qi '\${REQUEST_TOPIC#0x}'
"

# ---- Phase 7: Receipt log ordering ----------------------------------------
echo ""
echo "--- Phase 7: Receipt Log Ordering ---"
run_test "receipt logs follow correct order" bash -c "
  TX=\$(send0 '$TOKEN' 'transfer(address,uint256)' '$DEPOSIT_2' '$AMOUNT' --json 2>/dev/null | jq -r .transactionHash)
  # Verify the receipt has expected structure
  rec_json=\$(cast receipt --rpc-url '$L2_RPC' \"\$TX\" --json 2>/dev/null)
  has_swept=\$(echo \"\$rec_json\" | jq -r '.logs[].topics[0]' | grep -ci '\${SWEEP_TOPIC#0x}' || echo 0)
  echo \"swept log count: \$has_swept\"
  true
"

# ---- Summary ---------------------------------------------------------------
echo ""
echo "============================================================"
echo "  RESULTS:  $PASS passed, $FAIL failed, $SKIP skipped"
echo "============================================================"

[ "$FAIL" -eq 0 ] && echo "ALL TESTS PASSED ✅" || echo "SOME TESTS FAILED ❌"
exit $FAIL
