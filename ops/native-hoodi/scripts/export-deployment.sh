#!/usr/bin/env bash
# export-deployment.sh — save a record of a deployment under deployments/<NAME>/ (no private keys).
# Usage: export-deployment.sh NAME [L1_RUN_DIR=run/l1]
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
NAME="${1:?usage: export-deployment.sh NAME [L1_RUN_DIR]}"
L1="${2:-${ROOT}/run/l1}"
D="${ROOT}/deployments/${NAME}"
source "${ROOT}/scripts/lib/env.sh"
[[ -f "${L1}/deployment/hoodidev.json" ]] || { echo "no deployment in ${L1}" >&2; exit 1; }
[[ ! -e "${D}" ]] || { echo "${D} exists; refusing to overwrite" >&2; exit 1; }
mkdir -p "${D}/genesis" "${D}/verification"
cp "${L1}/deployment/hoodidev.json" "${D}/contracts.json"
cp "${L1}"/genesis/*.json "${D}/genesis/" 2>/dev/null || true
[[ -f "${L1}/state/start-block.dec" ]] && cp "${L1}/state/start-block.dec" "${D}/l1-start-block.txt"
cp "${ROOT}/notes/versions.md" "${D}/versions.md"
for f in "${ROOT}"/run/*test*.out "${ROOT}"/run/verify*.out "${L1}"/../verify*.out; do [[ -f "$f" ]] && cp "$f" "${D}/verification/"; done 2>/dev/null || true
set -a; source "${ROOT}/env/secrets.env"; set +a
DEP="$(cast wallet address --private-key "${DEPLOYER_PRIVATE_KEY}")"
SEQ="$(cast wallet address --private-key "${SEQUENCER_PRIVATE_KEY}")"
SUB="$(cast wallet address --private-key "${BATCH_SUBMITTER_PRIVATE_KEY}")"
ORA="$(cast wallet address --private-key "${L2_GAS_ORACLE_PRIVATE_KEY}")"
{
  echo "# Deployment ${NAME}"; echo
  echo "- exported: $(date -Is)"
  echo "- L1: Ethereum Hoodi (chain 560048); L2 chain id 53077"
  echo "- L1 start block: $(cat "${D}/l1-start-block.txt" 2>/dev/null || echo '?')"
  echo "- L2 genesis hash: $(node -e 'try{const j=require(process.argv[1]);console.log(j.genesisHash||j.hash||"see genesis/")}catch(e){console.log("see genesis/")}' "${D}/genesis/rollup.json" 2>/dev/null || echo 'see genesis/')"
  echo; echo "## Accounts (addresses only)"; echo
  echo "| role | address |"; echo "|---|---|"
  echo "| deployer | ${DEP} |"; echo "| sequencer | ${SEQ} |"; echo "| batch submitter | ${SUB} |"; echo "| gas oracle | ${ORA} |"
  echo; echo "## Contracts (L1)"; echo; echo "| name | address |"; echo "|---|---|"
  node -e 'for (const x of JSON.parse(require("fs").readFileSync(process.argv[1],"utf8"))) console.log("| "+x.name+" | "+x.address+" |")' "${D}/contracts.json"
  echo; echo "Files: contracts.json, genesis/, versions.md, verification/ (test outputs)."
} >"${D}/README.md"
echo "exported to ${D}"
