#!/usr/bin/env bash
#
# onyx-sweep-common.sh — shared constants + helpers for the Onyx sweep devnet
# scripts. SOURCE this (do not execute it); the caller must already have foundry
# (cast/forge) + jq on PATH and be cd'd into the foundry root (morph/contracts).
#
# This is the single source of truth for the deterministic CREATE2 deployment of
# the Onyx SweepRegistry, so onyx-sweep-demo.sh, onyx-sweep-comprehensive.sh and
# the devnet startup hook stay byte-for-byte consistent with the morph-reth
# SWEEP_REGISTRY_ADDRESS constant and scripts/deploy-sweep-registry.ts.

# ---- deterministic deployment inputs ---------------------------------------

# Solady deterministic-deployment-proxy — identical address on every EVM chain.
ONYX_FACTORY="0x4e59b44847b379578588920cA78FbF26c0B4956C"

# Keyless deployer of that factory (Nick's method). Its nonce-0 contract-creation
# lands the factory at ONYX_FACTORY; we just have to fund it for gas.
ONYX_FACTORY_SENDER="0x3fab184622dc19b6109349b94811493bf2a45362"

# Canonical PRE-EIP155 (chain-id-agnostic) factory bootstrap transaction.
# v=0x1b, r=s=0x2222…; recovers to ONYX_FACTORY_SENDER on ANY chain id. This is
# the same tx that deployed the factory on Morph mainnet and Hoodi.
# Do NOT swap in a chain-id-bound (EIP-155) variant — that only decodes on one
# chain id and is rejected everywhere else (e.g. on the 53077 devnet).
ONYX_FACTORY_RAW_TX="0xf8a58085174876e800830186a08080b853604580600e600039806000f350fe7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe03601600081602082378035828234f58015156039578182fd5b8082525050506014600cf31ba02222222222222222222222222222222222222222222222222222222222222222a02222222222222222222222222222222222222222222222222222222222222222"

# ProxyAdmin predeploy. This exact value is baked into the deterministic proxy
# address — changing it changes ONYX_EXPECTED_REGISTRY. Must match the default in
# scripts/deploy-sweep-registry.ts (DEFAULT_PROXY_ADMIN). Override via $PROXY_ADMIN
# only if you know what you are doing (it will break the EXPECTED_REGISTRY check).
ONYX_PROXY_ADMIN="${PROXY_ADMIN:-0x530000000000000000000000000000000000000b}"

# Fixed, versioned CREATE2 salts.
ONYX_SALT_IMPL_STR="morph.sweep-registry.impl.v1"
ONYX_SALT_PROXY_STR="morph.sweep-registry.proxy.v1"

# Deterministic proxy address given ONYX_PROXY_ADMIN=0x53..000b. MUST equal the
# morph-reth SWEEP_REGISTRY_ADDRESS constant (crates/chainspec/src/constants.rs).
ONYX_EXPECTED_REGISTRY="0x7aE8bEf666D1D0aB9C0ac5d636f375E46f8AE71A"

# EIP-712 / event topics — consensus values shared with morph-reth.
ONYX_MODE_STR="MORPH_SWEEP_V1"
ONYX_SCOPE_STR="WHITELISTED_ERC20_TO_MASTER_ONLY"
# keccak256("Swept(address,address,address,uint256,uint32)")   (sweep.rs SWEEP_TOPIC)
ONYX_SWEPT_TOPIC="0x035b37215a69e14a80883933d6aa84f0919a67af9410a4a73e8a23baeca011f0"
# keccak256("Transfer(address,address,uint256)")
ONYX_TRANSFER_TOPIC="0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
# keccak256("SweepRequested(address,address)")
ONYX_REQUEST_TOPIC="0x24e3f180db341974dcd99a5e223d9d944422e303230ddde6659302f8620bbcff"

# forge artifact paths (relative to the foundry root, foundry.toml out=forge-artifacts).
ONYX_IMPL_ARTIFACT="forge-artifacts/SweepRegistry.sol/SweepRegistry.json"
ONYX_PROXY_ARTIFACT="forge-artifacts/TransparentUpgradeableProxy.sol/TransparentUpgradeableProxy.json"

# ---- derived globals (populated by onyx_precompute_addresses) --------------
ONYX_SALT_IMPL=""
ONYX_SALT_PROXY=""
ONYX_IMPL=""              # predicted impl address
ONYX_REGISTRY=""          # predicted proxy (registry) address
ONYX_IMPL_INITCODE=""
ONYX_PROXY_INITCODE=""

onyx_require_tools() {
  local t
  for t in cast forge jq; do
    command -v "$t" >/dev/null 2>&1 || { echo "!! required tool not on PATH: $t" >&2; return 1; }
  done
}

onyx_code_present() {
  local rpc="$1" addr="$2" code
  code=$(cast code --rpc-url "$rpc" "$addr" 2>/dev/null || echo "0x")
  [ -n "$code" ] && [ "$code" != "0x" ]
}

# Compute impl + proxy CREATE2 addresses from the on-disk forge artifacts and
# assert the proxy matches ONYX_EXPECTED_REGISTRY (hence the morph-reth constant).
# Populates ONYX_IMPL / ONYX_REGISTRY / ONYX_*_INITCODE / ONYX_SALT_*.
onyx_precompute_addresses() {
  ONYX_SALT_IMPL=$(cast keccak "$ONYX_SALT_IMPL_STR")
  ONYX_SALT_PROXY=$(cast keccak "$ONYX_SALT_PROXY_STR")

  ONYX_IMPL_INITCODE=$(jq -r .bytecode.object "$ONYX_IMPL_ARTIFACT" 2>/dev/null || echo "")
  if [ "${ONYX_IMPL_INITCODE:0:2}" != "0x" ]; then
    echo "!! missing SweepRegistry bytecode ($ONYX_IMPL_ARTIFACT); run 'forge build' first" >&2
    return 1
  fi
  ONYX_IMPL=$(cast create2 --deployer "$ONYX_FACTORY" --salt "$ONYX_SALT_IMPL" --init-code "$ONYX_IMPL_INITCODE")

  local proxy_code proxy_args
  proxy_code=$(jq -r .bytecode.object "$ONYX_PROXY_ARTIFACT" 2>/dev/null || echo "")
  if [ "${proxy_code:0:2}" != "0x" ]; then
    echo "!! missing TransparentUpgradeableProxy bytecode ($ONYX_PROXY_ARTIFACT); run 'forge build' first" >&2
    return 1
  fi
  proxy_args=$(cast abi-encode 'x(address,address,bytes)' "$ONYX_IMPL" "$ONYX_PROXY_ADMIN" "0x")
  ONYX_PROXY_INITCODE="${proxy_code}${proxy_args#0x}"
  ONYX_REGISTRY=$(cast create2 --deployer "$ONYX_FACTORY" --salt "$ONYX_SALT_PROXY" --init-code "$ONYX_PROXY_INITCODE")

  echo "    impl  (predicted) = $ONYX_IMPL"
  echo "    proxy (predicted) = $ONYX_REGISTRY   (proxyAdmin=$ONYX_PROXY_ADMIN)"

  if [ "$(echo "$ONYX_REGISTRY" | tr 'A-Z' 'a-z')" != "$(echo "$ONYX_EXPECTED_REGISTRY" | tr 'A-Z' 'a-z')" ]; then
    echo "!! FATAL: predicted registry $ONYX_REGISTRY != expected $ONYX_EXPECTED_REGISTRY" >&2
    echo "   morph-reth hardcodes $ONYX_EXPECTED_REGISTRY; a mismatch means the EL will" >&2
    echo "   never find the registry. Check proxyAdmin / salts / SweepRegistry bytecode." >&2
    return 1
  fi
}

# Ensure the deterministic CREATE2 factory exists on the L2 (idempotent).
# $2 is a funded private key used only to pay the presigned sender's gas.
onyx_ensure_create2_factory() {
  local rpc="$1" funder_key="$2" i
  if onyx_code_present "$rpc" "$ONYX_FACTORY"; then
    echo "    CREATE2 factory already present"
    return 0
  fi
  echo "    CREATE2 factory absent — bootstrapping via presigned tx"
  cast send --rpc-url "$rpc" --private-key "$funder_key" --value 0.1ether "$ONYX_FACTORY_SENDER" >/dev/null
  cast publish --rpc-url "$rpc" "$ONYX_FACTORY_RAW_TX" >/dev/null 2>&1 || true
  for i in $(seq 1 30); do
    if onyx_code_present "$rpc" "$ONYX_FACTORY"; then
      echo "    CREATE2 factory deployed at $ONYX_FACTORY"
      return 0
    fi
    sleep 1
  done
  echo "!! failed to bootstrap CREATE2 factory $ONYX_FACTORY" >&2
  return 1
}

# Deploy impl + proxy via the factory (idempotent). Requires onyx_precompute_addresses.
onyx_deploy_registry() {
  local rpc="$1" deployer_key="$2"
  if onyx_code_present "$rpc" "$ONYX_IMPL"; then
    echo "    impl already deployed at $ONYX_IMPL"
  else
    echo "    deploying impl via CREATE2 …"
    cast send --rpc-url "$rpc" --private-key "$deployer_key" \
      "$ONYX_FACTORY" "deploy(bytes,bytes32)" "$ONYX_IMPL_INITCODE" "$ONYX_SALT_IMPL" >/dev/null
  fi
  if onyx_code_present "$rpc" "$ONYX_REGISTRY"; then
    echo "    proxy already deployed at $ONYX_REGISTRY"
  else
    echo "    deploying proxy via CREATE2 …"
    cast send --rpc-url "$rpc" --private-key "$deployer_key" \
      "$ONYX_FACTORY" "deploy(bytes,bytes32)" "$ONYX_PROXY_INITCODE" "$ONYX_SALT_PROXY" >/dev/null
  fi
}

# Initialize the proxy (idempotent: skips when owner() is already non-zero).
onyx_initialize_registry() {
  local rpc="$1" deployer_key="$2" owner="$3" cur
  cur=$(cast call --rpc-url "$rpc" "$ONYX_REGISTRY" "owner()(address)" 2>/dev/null || echo "")
  if [ -z "$cur" ] || [ "$cur" = "0x0000000000000000000000000000000000000000" ]; then
    echo "    initializing proxy (owner=$owner) …"
    cast send --rpc-url "$rpc" --private-key "$deployer_key" \
      "$ONYX_REGISTRY" "$(cast calldata 'initialize(address)' "$owner")" >/dev/null
  else
    echo "    proxy already initialized (owner=$cur)"
  fi
}

# Emit the EIP-712 SweepAuthorization typed-data JSON to stdout (for cast wallet sign).
onyx_typed_data() {
  local chain_id="$1" registry="$2" deposit="$3" master="$4" nonce="$5" deadline="$6"
  local mode scope
  mode=$(cast keccak "$ONYX_MODE_STR")
  scope=$(cast keccak "$ONYX_SCOPE_STR")
  cat <<JSON
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
    "chainId": $chain_id,
    "verifyingContract": "$registry"
  },
  "message": {
    "deposit": "$deposit",
    "master": "$master",
    "registry": "$registry",
    "chainId": $chain_id,
    "nonce": $nonce,
    "deadline": $deadline,
    "mode": "$mode",
    "sweepScope": "$scope"
  }
}
JSON
}
