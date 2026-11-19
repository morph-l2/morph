#!/usr/bin/env bash
#
# verify-l1.sh — read the deployed Morph L1 state back from the chain and compare it with what the deployment should have
# produced. It ignores every exit code and log of the deploy steps (hardhat tasks often print an error and still exit 0).
#
# Usage:  RPC_URL=... [L1_RUN=run/l1] scripts/verify-l1.sh
# Exit 0 only if every check passes.

set -uo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
L1_RUN="${L1_RUN:-${ROOT}/run/l1}"
MORPH="${MORPH:-$(cd "${ROOT}/../.." && pwd)}"
: "${RPC_URL:?set RPC_URL}"
SECRETS="${SECRETS:-${ROOT}/env/secrets.env}"
DEPLOY_JSON="${L1_RUN}/deployment/hoodidev.json"
CFG="${MORPH}/ops/l2-genesis/deploy-config/hoodidev-deploy-config.json"
# shellcheck disable=SC1090
source "${ROOT}/scripts/lib/env.sh"

# addresses only, derived from the keys without putting a key on a command line
addr_of() { ( set -a; source "${SECRETS}"; set +a; cd "${MORPH}/contracts" && KEY="${!1}" node -e 'console.log(new (require("ethers").Wallet)(process.env.KEY).address)' ); }
DEPLOYER="$(addr_of DEPLOYER_PRIVATE_KEY)"; SEQUENCER="$(addr_of SEQUENCER_PRIVATE_KEY)"; SUBMITTER="$(addr_of BATCH_SUBMITTER_PRIVATE_KEY)"
L2_CHAIN_ID="$(node -e 'console.log(require(process.argv[1]).l2ChainID)' "${CFG}")"
MIN_STAKE_WEI="$(cast to-wei 0.01 ether)"

dep() { node -e 'const d=JSON.parse(require("fs").readFileSync(process.argv[1],"utf8")).find(x=>x.name===process.argv[2]); if(!d)process.exit(1); console.log(d.address)' "${DEPLOY_JSON}" "$1"; }
call() { cast call "$1" "$2" "${@:3}" --rpc-url "${RPC_URL}" 2>&1 | awk 'NR==1{print $1}'; }
lc() { printf '%s' "${1,,}"; }

PASS=0; FAIL=0
check() { # check NAME ACTUAL EXPECTED
  if [[ "$(lc "$2")" == "$(lc "$3")" ]]; then PASS=$((PASS + 1)); printf '  PASS  %-52s %s\n' "$1" "$2"
  else FAIL=$((FAIL + 1)); printf '  FAIL  %-52s got %s, want %s\n' "$1" "$2" "$3"; fi
}
nonzero() { # nonzero NAME VALUE  (a 32-byte value that is not all zeros / not an error)
  if [[ "$2" =~ ^0x[0-9a-fA-F]{64}$ && ! "$2" =~ ^0x0{64}$ ]]; then PASS=$((PASS + 1)); printf '  PASS  %-52s %s\n' "$1" "${2:0:18}…"
  else FAIL=$((FAIL + 1)); printf '  FAIL  %-52s got %s\n' "$1" "$2"; fi
}

ROLLUP="$(dep Proxy__Rollup)"; MQ="$(dep Proxy__L1MessageQueueWithGasPriceOracle)"; MSGR="$(dep Proxy__L1CrossDomainMessenger)"
SUBM="$(dep Proxy__Submitter)"; SEQ="$(dep Proxy__L1Sequencer)"; PADMIN="$(dep Impl__ProxyAdmin)"
echo "L1 chain $(cast chain-id --rpc-url "${RPC_URL}"), deployer ${DEPLOYER}"
echo "Rollup ${ROLLUP}  MessageQueue ${MQ}  Messenger ${MSGR}  Submitter ${SUBM}  L1Sequencer ${SEQ}  ProxyAdmin ${PADMIN}"

echo; echo "== Rollup"
check "owner" "$(call "${ROLLUP}" 'owner()(address)')" "${DEPLOYER}"
check "LAYER_2_CHAIN_ID" "$(call "${ROLLUP}" 'LAYER_2_CHAIN_ID()(uint64)')" "${L2_CHAIN_ID}"
check "messageQueue" "$(call "${ROLLUP}" 'messageQueue()(address)')" "${MQ}"
check "submitterContract" "$(call "${ROLLUP}" 'submitterContract()(address)')" "${SUBM}"
nonzero "genesis batch imported: committedBatches(0)" "$(call "${ROLLUP}" 'committedBatches(uint256)(bytes32)' 0)"
check "lastCommittedBatchIndex" "$(call "${ROLLUP}" 'lastCommittedBatchIndex()(uint256)')" "0"
check "paused" "$(call "${ROLLUP}" 'paused()(bool)')" "false"

echo; echo "== L1CrossDomainMessenger"
check "owner" "$(call "${MSGR}" 'owner()(address)')" "${DEPLOYER}"
check "rollup" "$(call "${MSGR}" 'rollup()(address)')" "${ROLLUP}"
check "messageQueue" "$(call "${MSGR}" 'messageQueue()(address)')" "${MQ}"
check "counterpart (L2 messenger)" "$(call "${MSGR}" 'counterpart()(address)')" "0x5300000000000000000000000000000000000007"

echo; echo "== L1MessageQueueWithGasPriceOracle"
check "owner" "$(call "${MQ}" 'owner()(address)')" "${DEPLOYER}"
check "MESSENGER" "$(call "${MQ}" 'MESSENGER()(address)')" "${MSGR}"
check "ROLLUP_CONTRACT" "$(call "${MQ}" 'ROLLUP_CONTRACT()(address)')" "${ROLLUP}"

echo; echo "== Submitter"
check "owner" "$(call "${SUBM}" 'owner()(address)')" "${DEPLOYER}"
check "rollupContract" "$(call "${SUBM}" 'rollupContract()(address)')" "${ROLLUP}"
check "minimumStake" "$(call "${SUBM}" 'minimumStake()(uint256)')" "${MIN_STAKE_WEI}"
check "registered(submitter)" "$(call "${SUBM}" 'registered(address)(bool)' "${SUBMITTER}")" "true"
check "stakeOf(submitter)" "$(call "${SUBM}" 'stakeOf(address)(uint256)' "${SUBMITTER}")" "${MIN_STAKE_WEI}"
check "isActive(submitter)" "$(call "${SUBM}" 'isActive(address)(bool)' "${SUBMITTER}")" "true"

echo; echo "== L1Sequencer"
check "owner" "$(call "${SEQ}" 'owner()(address)')" "${DEPLOYER}"
check "getSequencer() is the sequencer we generated" "$(call "${SEQ}" 'getSequencer()(address)')" "${SEQUENCER}"

echo; echo "== every proxy: has code, implementation slot = recorded Impl__, admin slot = ProxyAdmin"
IMPL_SLOT=0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc
ADMIN_SLOT=0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103
while IFS=$'\t' read -r name addr; do
  impl_name="Impl__${name#Proxy__}"
  want_impl="$(dep "${impl_name}" 2>/dev/null || true)"
  code="$(cast code "${addr}" --rpc-url "${RPC_URL}")"
  if [[ "${code}" == "0x" ]]; then FAIL=$((FAIL + 1)); printf '  FAIL  %-52s no code\n' "${name}"; continue; fi
  got_impl="$(cast parse-bytes32-address "$(cast storage "${addr}" "${IMPL_SLOT}" --rpc-url "${RPC_URL}")" 2>/dev/null)"
  got_admin="$(cast parse-bytes32-address "$(cast storage "${addr}" "${ADMIN_SLOT}" --rpc-url "${RPC_URL}")" 2>/dev/null)"
  if [[ -n "${want_impl}" ]]; then check "${name} implementation" "${got_impl}" "${want_impl}"; else printf '  ----  %-52s no recorded Impl__ (implementation %s)\n' "${name}" "${got_impl}"; fi
  check "${name} admin" "${got_admin}" "${PADMIN}"
done < <(node -e 'for (const d of JSON.parse(require("fs").readFileSync(process.argv[1],"utf8"))) if (d.name.startsWith("Proxy__")) console.log(d.name+"\t"+d.address)' "${DEPLOY_JSON}")

echo; echo "== result: ${PASS} passed, ${FAIL} failed"
[[ "${FAIL}" -eq 0 ]]
