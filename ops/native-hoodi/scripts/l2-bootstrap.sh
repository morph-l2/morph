#!/usr/bin/env bash
#
# l2-bootstrap.sh — one-time setup on the (local, private) L2 so that the gas-price-oracle can work:
#   1. give the deployer (owner of the L2 contracts) and the oracle key some L2 balance, from a dev account that the L2
#      genesis pre-funds (`fundDevAccounts`). That account's key is the well-known Anvil #0 key: acceptable ONLY because
#      this L2 listens on 127.0.0.1 and holds nothing of value.
#   2. let the deployer (GasPriceOracle owner) put the oracle key on the GasPriceOracle allow list
#      (the bot's setters are `onlyAllowed`; the owner is the user's deployer, not the oracle key).
# Idempotent: every step checks the current state first and reads the result back.
#
# Usage: L2_RPC=http://127.0.0.1:8545 scripts/l2-bootstrap.sh
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MORPH="${MORPH:-$(cd "${ROOT}/../.." && pwd)}"
L2_RPC="${L2_RPC:-http://127.0.0.1:8545}"
ORACLE_CONTRACT=0x530000000000000000000000000000000000000F
DEV_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80   # Anvil #0: public, local-only L2
FUND_DEPLOYER_ETH="${FUND_DEPLOYER_ETH:-100}"
FUND_ORACLE_ETH="${FUND_ORACLE_ETH:-10}"
# shellcheck disable=SC1090
source "${ROOT}/scripts/lib/env.sh"
set -a; source "${ROOT}/env/secrets.env"; set +a
export MORPH_CONTRACTS="${MORPH}/contracts"

addr_of() { ( cd "${MORPH}/contracts" && KEY="${!1}" node -e 'console.log(new (require("ethers").Wallet)(process.env.KEY).address)' ); }
DEPLOYER="$(addr_of DEPLOYER_PRIVATE_KEY)"; ORACLE="$(addr_of L2_GAS_ORACLE_PRIVATE_KEY)"
[[ "$(cast chain-id --rpc-url "${L2_RPC}")" == "53077" ]] || { echo "L2 at ${L2_RPC} is not chain 53077" >&2; exit 1; }
say() { printf '%s %s\n' "$(date +%H:%M:%S)" "$*"; }

fund() { # fund NAME ADDRESS ETH
  local bal; bal="$(cast balance "$2" --rpc-url "${L2_RPC}")"
  if [[ "${bal}" != "0" ]]; then say "$1 $2 already has L2 balance $(cast from-wei "${bal}") ETH"; return 0; fi
  cast send "$2" --value "$3ether" --private-key "${DEV_KEY}" --rpc-url "${L2_RPC}" >/dev/null
  bal="$(cast balance "$2" --rpc-url "${L2_RPC}")"
  [[ "${bal}" != "0" ]] || { echo "funding $1 failed" >&2; exit 1; }
  say "$1 $2 funded on L2: $(cast from-wei "${bal}") ETH"
}
fund deployer "${DEPLOYER}" "${FUND_DEPLOYER_ETH}"
fund oracle "${ORACLE}" "${FUND_ORACLE_ETH}"

owner="$(cast call "${ORACLE_CONTRACT}" 'owner()(address)' --rpc-url "${L2_RPC}")"
[[ "${owner,,}" == "${DEPLOYER,,}" ]] || { echo "GasPriceOracle owner ${owner} is not the deployer ${DEPLOYER}" >&2; exit 1; }
if [[ "$(cast call "${ORACLE_CONTRACT}" 'isAllowed(address)(bool)' "${ORACLE}" --rpc-url "${L2_RPC}")" == "true" ]]; then
  say "oracle ${ORACLE} is already on the allow list"
else
  data="$(cast calldata 'setAllowList(address[],bool[])' "[${ORACLE}]" "[true]")"
  KEY="${DEPLOYER_PRIVATE_KEY}" node "${ROOT}/scripts/send-tx.js" "${L2_RPC}" "${ORACLE_CONTRACT}" --data "${data}"
  [[ "$(cast call "${ORACLE_CONTRACT}" 'isAllowed(address)(bool)' "${ORACLE}" --rpc-url "${L2_RPC}")" == "true" ]] ||
    { echo "allow list still does not contain ${ORACLE}" >&2; exit 1; }
  say "oracle ${ORACLE} added to the allow list (read back: true)"
fi
say "done: allowListEnabled=$(cast call "${ORACLE_CONTRACT}" 'allowListEnabled()(bool)' --rpc-url "${L2_RPC}")"
