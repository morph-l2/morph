#!/usr/bin/env bash
#
# deposit-test.sh — L1 -> L2 ETH deposit smoke test (P7-1).
# Deposits AMOUNT_ETH from the deployer on Hoodi into a brand-new random L2 address, then waits for the balance to
# appear on BOTH L2 nodes (node-0 produced the block; node-1 re-derives it from L1).
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MORPH="${MORPH:-$(cd "${ROOT}/../.." && pwd)}"
L1_RUN="${L1_RUN:-${ROOT}/run/l1}"
AMOUNT_ETH="${AMOUNT_ETH:-0.01}"
VALUE_ETH="${VALUE_ETH:-0.012}"     # amount + the L1 message fee (110000 gas * l2BaseFee is tiny; the rest is refunded)
GAS_LIMIT="${GAS_LIMIT:-400000}"   # 110000 (the repo example) ran out of gas on L2: see notes/issues.md I-027
L1_RPC="${L1_RPC:-http://127.0.0.1:18546}"   # the l1-proxy started by morph-native.sh
L2_A=http://127.0.0.1:8545; L2_B=http://127.0.0.1:8645
# shellcheck disable=SC1090
source "${ROOT}/scripts/lib/env.sh"
set -a; source "${ROOT}/env/secrets.env"; set +a
export MORPH_CONTRACTS="${MORPH}/contracts"
ROUTER="$(node -e 'console.log(JSON.parse(require("fs").readFileSync(process.argv[1],"utf8")).find(x=>x.name==="Proxy__L1GatewayRouter").address)' "${L1_RUN}/deployment/hoodidev.json")"
RECIPIENT="$(cd "${MORPH}/contracts" && node -e 'console.log(require("ethers").Wallet.createRandom().address)')"   # address only; nobody holds its key
say() { printf '%s %s\n' "$(date +%H:%M:%S)" "$*"; }

say "L1GatewayRouter ${ROUTER}"; say "fresh L2 recipient ${RECIPIENT} (balance now: A=$(cast balance "${RECIPIENT}" --rpc-url ${L2_A}) B=$(cast balance "${RECIPIENT}" --rpc-url ${L2_B}))"
data="$(cast calldata 'depositETH(address,uint256,uint256)' "${RECIPIENT}" "$(cast to-wei "${AMOUNT_ETH}" ether)" "${GAS_LIMIT}")"
KEY="${DEPLOYER_PRIVATE_KEY}" node "${ROOT}/scripts/send-tx.js" "${L1_RPC}" "${ROUTER}" --data "${data}" --value-ether "${VALUE_ETH}"
t0=$(date +%s); want="$(cast to-wei "${AMOUNT_ETH}" ether)"
while true; do
  a="$(cast balance "${RECIPIENT}" --rpc-url ${L2_A})"; b="$(cast balance "${RECIPIENT}" --rpc-url ${L2_B})"
  say "L2 balance: node-0=${a} node-1=${b}  (want ${want})"
  [[ "${a}" == "${want}" && "${b}" == "${want}" ]] && break
  [[ $(( $(date +%s) - t0 )) -lt "${DEPOSIT_TIMEOUT_SEC:-600}" ]] || { echo "timed out waiting for the deposit on L2" >&2; exit 1; }
  sleep 15
done
say "OK: ${AMOUNT_ETH} ETH deposited on Hoodi arrived on L2 at ${RECIPIENT} on both nodes after $(( $(date +%s) - t0 )) s"
