#!/usr/bin/env bash
# sweep-submitter.sh — send the batch submitter's L1 balance (minus a reserve) to the deployer. Stop the stack first.
# Usage: sweep-submitter.sh [RESERVE_ETH=0.01]
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
source "${ROOT}/scripts/lib/env.sh"
set -a; source "${ROOT}/env/secrets.env"; set +a
RPC="${HOODI_RPC:-https://rpc.hoodi.ethpandaops.io}"
RESERVE="${1:-0.01}"
export MORPH_CONTRACTS="$(cd "${ROOT}/../../contracts" && pwd)"
SUB="$(cast wallet address --private-key "${BATCH_SUBMITTER_PRIVATE_KEY}")"
DEP="$(cast wallet address --private-key "${DEPLOYER_PRIVATE_KEY}")"
BAL="$(cast balance "${SUB}" --ether --rpc-url "${RPC}")"
AMT="$(awk -v b="${BAL}" -v r="${RESERVE}" 'BEGIN{printf "%.6f", b-r}')"
echo "submitter ${SUB}: ${BAL} ETH -> sending ${AMT} to deployer ${DEP}"
echo "deployer before: $(cast balance "${DEP}" --ether --rpc-url "${RPC}")"
KEY="${BATCH_SUBMITTER_PRIVATE_KEY}" node "${ROOT}/scripts/send-tx.js" "${RPC}" "${DEP}" --value-ether "${AMT}"
sleep 20
echo "deployer after : $(cast balance "${DEP}" --ether --rpc-url "${RPC}")"
echo "submitter after: $(cast balance "${SUB}" --ether --rpc-url "${RPC}")"
