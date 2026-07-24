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
#   1. fund the demo deployer (acct1) from acct0
#   2. deploy impl + TransparentUpgradeableProxy via the deterministic
#      CREATE2 factory (address is network-identical) then call initialize()
#   3. deploy a MockERC20, whitelist it, register a deposit EOA via EIP-712
#   4. transfer whitelisted tokens INTO the deposit -> the EL auto-sweeps them
#      to master and appends a Swept log
#   5. assert: Swept log present, deposit balance == 0
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
DEPLOYER=0x70997970C51812dc3A010C7d01b50e0d17dc79C8  # acct1: Registry deployer
DEPLOYER_KEY=0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d
PROXY_ADMIN="${PROXY_ADMIN:-0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC}" # acct2
DEPOSIT=0x90F79bf6EB2c4f870365E785982E1f101E93b906    # acct3: deposit EOA (plain, no code)
DEPOSIT_KEY=0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6

MASTER=$ACCT0
OWNER="${REGISTRY_OWNER:-$ACCT0}"

# CREATE2
FACTORY=0x4e59b44847b379578588920cA78FbF26c0B4956C
SALT_IMPL=$(cast keccak "morph.sweep-registry.impl.v1")
SALT_PROXY=$(cast keccak "morph.sweep-registry.proxy.v1")

SWEEP_TOPIC=0x4cb65f464c97b7cae979110960f2dba5a9447c795638563ad5f1e2b52c6f37dd
AMOUNT=1000000000000000000  # 1e18

send0() { cast send --rpc-url "$L2_RPC" --private-key "$ACCT0_KEY" "$@"; }

echo "==> L2 RPC $L2_RPC (chainId $L2_CHAIN_ID)"
cast chain-id --rpc-url "$L2_RPC" >/dev/null || { echo "L2 RPC unreachable"; exit 1; }

# ---- precompute CREATE2 addresses -------------------------------------------
echo "==> [0] precompute CREATE2 addresses"

IMPL_CODE=$(jq -r .bytecode.object forge-artifacts/SweepRegistry.sol/SweepRegistry.json)
if [ "${IMPL_CODE:0:2}" != "0x" ]; then
  echo "!! missing SweepRegistry bytecode; run 'forge build' first"
  exit 1
fi
IMPL=$(cast create2 --deployer "$FACTORY" --salt "$SALT_IMPL" --init-code "$IMPL_CODE")
echo "    impl  (predicted) = $IMPL"

PROXY_CODE=$(jq -r .bytecode.object \
  forge-artifacts/TransparentUpgradeableProxy.sol/TransparentUpgradeableProxy.json)
if [ "${PROXY_CODE:0:2}" != "0x" ]; then
  echo "!! missing TransparentUpgradeableProxy bytecode; run 'forge build' first"
  exit 1
fi
PROXY_ARGS=$(cast abi-encode 'x(address,address,bytes)' "$IMPL" "$PROXY_ADMIN" "0x")
PROXY_INITCODE="${PROXY_CODE}${PROXY_ARGS#0x}"
REGISTRY=$(cast create2 --deployer "$FACTORY" --salt "$SALT_PROXY" --init-code "$PROXY_INITCODE")
echo "    proxy (predicted) = $REGISTRY"
echo "    (paste this into genesis config sweepRegistryAddress)"

# ---- 1. fund the deployer + ensure factory exists ---------------------------
echo "==> [1] fund deployer $DEPLOYER"
send0 --value 10ether "$DEPLOYER" >/dev/null

FACTORY_CODE=$(cast code --rpc-url "$L2_RPC" "$FACTORY" 2>/dev/null || echo "0x")
if [ "$FACTORY_CODE" = "0x" ] || [ -z "$FACTORY_CODE" ]; then
  echo "    CREATE2 factory not present — deploying …"
  # The deterministic-deployment-proxy one-shot transaction (KeylessDeployer):
  # 1. Fund the presigned sender so the transaction can pay for gas.
  SIGNER=0x3fab184622dc19b6109349b94811493bf2a45362
  send0 --value 0.5ether "$SIGNER" >/dev/null
  # 2. Broadcast the well-known raw transaction.
  cast publish --rpc-url "$L2_RPC" \
    0xf8a58085174876e800830186a08080b853604580600e600039806000f350fe7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe03601600081602082378035828234f58015156039578182fd5b8082525050506014600cf3820e1da053edb6323539302ee2e91844c4a7d8c59303dec73b2b12692d4a52805bddb018a00a8cb751043801076f50e7841e6789b6d644ed66471ab701755c373fdef8f020 2>/dev/null || true
  sleep 1
  FACTORY_CODE_AFTER=$(cast code --rpc-url "$L2_RPC" "$FACTORY")
  if [ "$FACTORY_CODE_AFTER" = "0x" ] || [ -z "$FACTORY_CODE_AFTER" ]; then
    echo "!! failed to deploy CREATE2 factory; you may need to do it manually"
    echo "   cast send --rpc-url $L2_RPC --private-key $ACCT0_KEY --value 0.5ether $SIGNER"
    echo "   cast publish --rpc-url $L2_RPC <raw-tx>"
    exit 1
  fi
  echo "    factory deployed successfully"
fi

# ---- 2. deploy impl + proxy via CREATE2 ------------------------------------
echo "==> [2] deploy via CREATE2 factory $FACTORY"

IMPL_ONCHAIN=$(cast code --rpc-url "$L2_RPC" "$IMPL" 2>/dev/null || echo "0x")
if [ "$IMPL_ONCHAIN" = "0x" ] || [ -z "$IMPL_ONCHAIN" ]; then
  echo "    deploying impl …"
  cast send --rpc-url "$L2_RPC" --private-key "$DEPLOYER_KEY" \
    "$FACTORY" "deploy(bytes,bytes32)" "$IMPL_CODE" "$SALT_IMPL" >/dev/null
  echo "    impl  = $IMPL"
else
  echo "    impl already deployed at $IMPL — skipping"
fi

PROXY_ONCHAIN=$(cast code --rpc-url "$L2_RPC" "$REGISTRY" 2>/dev/null || echo "0x")
if [ "$PROXY_ONCHAIN" = "0x" ] || [ -z "$PROXY_ONCHAIN" ]; then
  echo "    deploying proxy …"
  cast send --rpc-url "$L2_RPC" --private-key "$DEPLOYER_KEY" \
    "$FACTORY" "deploy(bytes,bytes32)" "$PROXY_INITCODE" "$SALT_PROXY" >/dev/null
  echo "    proxy = $REGISTRY"
else
  echo "    proxy already deployed at $REGISTRY — skipping"
fi

# ---- 2b. initialize the proxy (separate from deployment) --------------------
INIT_DATA=$(cast calldata "initialize(address)" "$OWNER")
# Check if already initialized: call owner()
INIT_CHECK=$(cast call --rpc-url "$L2_RPC" "$REGISTRY" "owner()(address)" 2>/dev/null || echo "")
if [ -z "$INIT_CHECK" ] || [ "$INIT_CHECK" = "0x0000000000000000000000000000000000000000" ]; then
  echo "    initializing proxy (owner = $OWNER) …"
  cast send --rpc-url "$L2_RPC" --private-key "$DEPLOYER_KEY" \
    "$REGISTRY" "$INIT_DATA" >/dev/null
  echo "    initialized"
else
  echo "    proxy already initialized, owner = $INIT_CHECK"
fi

# ---- 3. deploy MockERC20, whitelist it -------------------------------------
echo "==> [3] deploy MockERC20 + whitelist"
TOKEN=$(forge create '@rari-capital/solmate/src/test/utils/mocks/MockERC20.sol:MockERC20' \
  --rpc-url "$L2_RPC" --private-key "$ACCT0_KEY" --broadcast --json \
  --constructor-args "Onyx Demo" "ODM" 18 | jq -r .deployedTo)
echo "    token = $TOKEN"
send0 "$TOKEN" "mint(address,uint256)" "$ACCT0" "$AMOUNT" >/dev/null
send0 "$REGISTRY" "setTokenWhitelist(address,bool)" "$TOKEN" true >/dev/null

# resolveSweep is zero before registration
pre=$(cast call --rpc-url "$L2_RPC" "$REGISTRY" "resolveSweep(address,address)(address)" "$TOKEN" "$DEPOSIT")
echo "    resolveSweep(before register) = $pre"

# ---- 4. deposit signs EIP-712 authorization, owner registers it ------------
echo "==> [4] EIP-712 register deposit $DEPOSIT -> master $MASTER"
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
    "chainId": $L2_CHAIN_ID,
    "verifyingContract": "$REGISTRY"
  },
  "message": {
    "deposit": "$DEPOSIT",
    "master": "$MASTER",
    "registry": "$REGISTRY",
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
send0 "$REGISTRY" "registerSweep(address,address,uint256,uint64,bytes)" \
  "$DEPOSIT" "$MASTER" 0 "$DEADLINE" "$SIG" >/dev/null
post=$(cast call --rpc-url "$L2_RPC" "$REGISTRY" "resolveSweep(address,address)(address)" "$TOKEN" "$DEPOSIT")
echo "    resolveSweep(after register) = $post"

# ---- 5. transfer whitelisted tokens INTO the deposit -> EL auto-sweeps ------
echo "==> [5] transfer $AMOUNT of $TOKEN into deposit -> expect sweep"
TXHASH=$(send0 "$TOKEN" "transfer(address,uint256)" "$DEPOSIT" "$AMOUNT" --json | jq -r .transactionHash)
echo "    tx = $TXHASH"

# ---- 6. assertions ----------------------------------------------------------
echo "==> [6] assert"
LOGS=$(cast receipt --rpc-url "$L2_RPC" "$TXHASH" --json | jq -r '.logs[].topics[0]')
if echo "$LOGS" | grep -qi "${SWEEP_TOPIC#0x}"; then
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
