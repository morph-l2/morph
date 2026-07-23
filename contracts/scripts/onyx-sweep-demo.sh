#!/usr/bin/env bash
#
# Onyx recoverable-sweep end-to-end demo against a running `make devnet-up-reth`.
#
# Prerequisites:
#   - devnet is up with the Onyx genesis patch applied, i.e. started with:
#       DEVNET_ONYX=1 MORPH_RETH_BUILD_FROM_SOURCE=true make devnet-up-reth
#     (morph-reth must be on branch feat/onyx-recoverable-sweep before build)
#   - foundry (cast, forge) + jq on PATH
#
# What it does (all on L2):
#   1. fund the demo deployer (acct1) from acct0
#   2. acct1 deploys the Registry impl (nonce 0) + TransparentUpgradeableProxy
#      (nonce 1) -> lands at the CREATE-predicted address the genesis config
#      was patched with (REGISTRY below)
#   3. deploy a MockERC20, whitelist it, register a deposit EOA via EIP-712
#   4. transfer whitelisted tokens INTO the deposit -> the EL auto-sweeps them
#      to master and appends a RecoverableSweep log
#   5. assert: RecoverableSweep log present, deposit balance == 0
#
set -euo pipefail

export PATH="$HOME/.foundry/bin:$PATH"
cd "$(dirname "${BASH_SOURCE[0]}")/.."   # morph/contracts (foundry root)

# ---- config -----------------------------------------------------------------
L2_RPC="${L2_RPC:-http://127.0.0.1:8545}"
L2_CHAIN_ID="${L2_CHAIN_ID:-53077}"

# Anvil dev accounts (well-known keys, safe for a local devnet only)
ACCT0=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266   # owner / master / gas provider
ACCT0_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
DEPLOYER=0x70997970C51812dc3A010C7d01b50e0d17dc79C8  # acct1: Registry deployer (nonce 0,1)
DEPLOYER_KEY=0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
PROXY_ADMIN=0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC # acct2: proxy admin (never calls impl)
DEPOSIT=0x90F79bf6EB2c4f870365E785982E1f101E93b906    # acct3: deposit EOA (plain, no code)
DEPOSIT_KEY=0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6

MASTER=$ACCT0
OWNER=$ACCT0
# CREATE-predicted proxy address = keccak(rlp(DEPLOYER, nonce=1)); must match the
# recoverableDepositRegistryAddress patched into the genesis config.
REGISTRY="${REGISTRY:-0x71C95911E9a5D330f4D621842EC243EE1343292e}"

SWEEP_TOPIC=0x4cb65f464c97b7cae979110960f2dba5a9447c795638563ad5f1e2b52c6f37dd
AMOUNT=1000000000000000000  # 1e18

send0() { cast send --rpc-url "$L2_RPC" --private-key "$ACCT0_KEY" "$@"; }

echo "==> L2 RPC $L2_RPC (chainId $L2_CHAIN_ID); expected Registry $REGISTRY"
cast chain-id --rpc-url "$L2_RPC" >/dev/null || { echo "L2 RPC unreachable"; exit 1; }

# 1. fund the deployer
echo "==> [1] fund deployer $DEPLOYER"
send0 --value 10ether "$DEPLOYER" >/dev/null

# 2. deploy impl (nonce 0) then proxy (nonce 1)
echo "==> [2] deploy Registry impl + proxy from $DEPLOYER"
IMPL=$(forge create contracts/l2/system/RecoverableDepositRegistry.sol:RecoverableDepositRegistry \
  --rpc-url "$L2_RPC" --private-key "$DEPLOYER_KEY" --broadcast --json | jq -r .deployedTo)
echo "    impl  = $IMPL"
INIT_DATA=$(cast calldata "initialize(address)" "$OWNER")
PROXY=$(forge create \
  '@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol:TransparentUpgradeableProxy' \
  --rpc-url "$L2_RPC" --private-key "$DEPLOYER_KEY" --broadcast --json \
  --constructor-args "$IMPL" "$PROXY_ADMIN" "$INIT_DATA" | jq -r .deployedTo)
echo "    proxy = $PROXY"
if [ "$(echo "$PROXY" | tr 'A-Z' 'a-z')" != "$(echo "$REGISTRY" | tr 'A-Z' 'a-z')" ]; then
  echo "!! proxy $PROXY != predicted $REGISTRY — genesis config address won't match; abort"
  exit 1
fi

# 3. deploy MockERC20, mint to acct0, whitelist it
echo "==> [3] deploy MockERC20 + whitelist"
TOKEN=$(forge create '@rari-capital/solmate/src/test/utils/mocks/MockERC20.sol:MockERC20' \
  --rpc-url "$L2_RPC" --private-key "$ACCT0_KEY" --broadcast --json \
  --constructor-args "Onyx Demo" "ODM" 18 | jq -r .deployedTo)
echo "    token = $TOKEN"
send0 "$TOKEN" "mint(address,uint256)" "$ACCT0" "$AMOUNT" >/dev/null
send0 "$PROXY" "setTokenWhitelist(address,bool)" "$TOKEN" true >/dev/null

# resolveSweep is zero before registration
pre=$(cast call --rpc-url "$L2_RPC" "$PROXY" "resolveSweep(address,address)(address)" "$TOKEN" "$DEPOSIT")
echo "    resolveSweep(before register) = $pre"

# 4. deposit signs EIP-712 authorization, owner registers it
echo "==> [4] EIP-712 register deposit $DEPOSIT -> master $MASTER"
MODE=$(cast keccak "MORPH_RECOVERABLE_DEPOSIT_V1")
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
    "RecoverableDepositAuthorization": [
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
  "primaryType": "RecoverableDepositAuthorization",
  "domain": {
    "name": "RecoverableDepositRegistry",
    "version": "1",
    "chainId": $L2_CHAIN_ID,
    "verifyingContract": "$PROXY"
  },
  "message": {
    "deposit": "$DEPOSIT",
    "master": "$MASTER",
    "registry": "$PROXY",
    "chainId": $L2_CHAIN_ID,
    "nonce": 0,
    "deadline": $DEADLINE,
    "mode": "$MODE",
    "sweepScope": "$SCOPE"
  }
}
JSON
SIG=$(cast wallet sign --private-key "$DEPOSIT_KEY" --data --from-file "$TYPED")
rm -f "$TYPED"
send0 "$PROXY" "registerRecoverableDeposit(address,address,uint256,uint64,bytes)" \
  "$DEPOSIT" "$MASTER" 0 "$DEADLINE" "$SIG" >/dev/null
post=$(cast call --rpc-url "$L2_RPC" "$PROXY" "resolveSweep(address,address)(address)" "$TOKEN" "$DEPOSIT")
echo "    resolveSweep(after register) = $post"

# 5. transfer whitelisted tokens INTO the deposit -> EL auto-sweeps
echo "==> [5] transfer $AMOUNT of $TOKEN into deposit -> expect sweep"
TXHASH=$(send0 "$TOKEN" "transfer(address,uint256)" "$DEPOSIT" "$AMOUNT" --json | jq -r .transactionHash)
echo "    tx = $TXHASH"

# 6. assertions
echo "==> [6] assert"
LOGS=$(cast receipt --rpc-url "$L2_RPC" "$TXHASH" --json | jq -r '.logs[].topics[0]')
if echo "$LOGS" | grep -qi "${SWEEP_TOPIC#0x}"; then
  echo "    OK: RecoverableSweep log present"
else
  echo "    FAIL: no RecoverableSweep log in receipt"; echo "$LOGS"; exit 1
fi
dep_bal=$(cast call --rpc-url "$L2_RPC" "$TOKEN" "balanceOf(address)(uint256)" "$DEPOSIT")
master_bal=$(cast call --rpc-url "$L2_RPC" "$TOKEN" "balanceOf(address)(uint256)" "$MASTER")
echo "    deposit balance = $dep_bal (expect 0)"
echo "    master  balance = $master_bal"
[ "${dep_bal%% *}" = "0" ] || { echo "    FAIL: deposit not drained"; exit 1; }

echo ""
echo "SUCCESS: Onyx recoverable-sweep worked end-to-end on devnet."
