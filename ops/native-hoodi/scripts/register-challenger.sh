#!/usr/bin/env bash
#
# register-challenger.sh — make the challenger account usable on an existing deployment (idempotent):
#   1. the deployer (Rollup owner) adds CHALLENGER_PRIVATE_KEY's address to Rollup.isChallenger
#   2. the deployer tops the challenger up to CHALLENGER_FUND_ETH so it can pay the challenge deposit and gas
# Each result is read back from the chain; a zero exit code from a transaction is not trusted (notes/issues.md I-022).
#
# deploy-l1.sh runs this as step s34 when CHALLENGER_PRIVATE_KEY is set. Run it by hand for a deployment made without it.
#
# Environment (all optional):
#   RPC_URL             L1 endpoint. Default: the local l1-proxy if it is listening, else L1_ETH_RPC from secrets
#   L1_RUN              deploy-l1.sh output dir (default run/l1); the Rollup address comes from its deployment/hoodidev.json
#   CHALLENGER_FUND_ETH balance to keep on the challenger, in ETH (default 0.2; challengeDeposit is 0.01 on this deployment)
#   SECRETS             default env/secrets.env
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MORPH="${MORPH:-$(cd "${ROOT}/../.." && pwd)}"
SECRETS="${SECRETS:-${ROOT}/env/secrets.env}"
L1_RUN="${L1_RUN:-${ROOT}/run/l1}"
FUND_ETH="${CHALLENGER_FUND_ETH:-0.2}"
# shellcheck disable=SC1090
source "${ROOT}/scripts/lib/env.sh"
[[ -f "${SECRETS}" ]] || { echo "missing ${SECRETS}" >&2; exit 1; }
set -a
# shellcheck disable=SC1090
source "${SECRETS}"
set +a
for v in DEPLOYER_PRIVATE_KEY CHALLENGER_PRIVATE_KEY; do
  [[ "${!v:-}" =~ ^0x[0-9a-fA-F]{64}$ ]] || {
    echo "${v} is not set to a 32-byte hex key in ${SECRETS}" >&2
    [[ "${v}" == "CHALLENGER_PRIVATE_KEY" ]] && echo "  add one (cast wallet new), or generate a full set with scripts/init-secrets.sh on a fresh checkout" >&2
    exit 1
  }
done
export MORPH_CONTRACTS="${MORPH}/contracts"   # send-tx.js loads ethers from here

if [[ -z "${RPC_URL:-}" ]]; then
  PROXY_PORT="${PROXY_PORT:-18546}"
  if (exec 3<>"/dev/tcp/127.0.0.1/${PROXY_PORT}") 2>/dev/null; then RPC_URL="http://127.0.0.1:${PROXY_PORT}"; else RPC_URL="${L1_ETH_RPC:?set RPC_URL or L1_ETH_RPC}"; fi
fi
DEPLOY_JSON="${L1_RUN}/deployment/hoodidev.json"
[[ -f "${DEPLOY_JSON}" ]] || { echo "no deployment at ${DEPLOY_JSON}" >&2; exit 1; }
ROLLUP="$(node -e 'const d=JSON.parse(require("fs").readFileSync(process.argv[1],"utf8")).find(x=>x.name==="Proxy__Rollup"); if(!d){process.exit(1)} console.log(d.address)' "${DEPLOY_JSON}")"

addr_of() { ( cd "${MORPH}/contracts" && KEY="${!1}" node -e 'console.log(new (require("ethers").Wallet)(process.env.KEY).address)' ); }
DEPLOYER="$(addr_of DEPLOYER_PRIVATE_KEY)"
CHALLENGER="$(addr_of CHALLENGER_PRIVATE_KEY)"
say() { printf '%s %s\n' "$(date +%H:%M:%S)" "$*"; }
# a read right after a mined write can hit a public-RPC node that has not seen the block: poll briefly
wait_for() { local i; for i in $(seq 1 30); do "$@" && return 0; sleep 2; done; return 1; }
is_challenger() { [[ "$(cast call "${ROLLUP}" 'isChallenger(address)(bool)' "${CHALLENGER}" --rpc-url "${RPC_URL}")" == "true" ]]; }

owner="$(cast call "${ROLLUP}" 'owner()(address)' --rpc-url "${RPC_URL}")"
[[ "${owner,,}" == "${DEPLOYER,,}" ]] || { echo "Rollup owner ${owner} is not the deployer ${DEPLOYER}" >&2; exit 1; }

if is_challenger; then
  say "challenger ${CHALLENGER} is already on Rollup ${ROLLUP}"
else
  data="$(cast calldata 'addChallenger(address)' "${CHALLENGER}")"
  KEY="${DEPLOYER_PRIVATE_KEY}" node "${ROOT}/scripts/send-tx.js" "${RPC_URL}" "${ROLLUP}" --data "${data}"
  wait_for is_challenger || { echo "isChallenger(${CHALLENGER}) is still false after addChallenger" >&2; exit 1; }
  say "added ${CHALLENGER} to the challenger list of Rollup ${ROLLUP}"
fi

want="$(cast to-wei "${FUND_ETH}" ether)"
has() { node -e 'process.exit(BigInt(process.argv[1]) >= BigInt(process.argv[2]) ? 0 : 1)' "$(cast balance "${CHALLENGER}" --rpc-url "${RPC_URL}")" "${want}"; }
if has; then
  say "challenger balance $(cast balance "${CHALLENGER}" --ether --rpc-url "${RPC_URL}") ETH (>= ${FUND_ETH}): not funding"
else
  KEY="${DEPLOYER_PRIVATE_KEY}" node "${ROOT}/scripts/send-tx.js" "${RPC_URL}" "${CHALLENGER}" --value-ether "${FUND_ETH}"
  wait_for has || { echo "challenger balance is still below ${FUND_ETH} ETH after funding" >&2; exit 1; }
  say "challenger funded: $(cast balance "${CHALLENGER}" --ether --rpc-url "${RPC_URL}") ETH"
fi
echo "challenger ${CHALLENGER}"
