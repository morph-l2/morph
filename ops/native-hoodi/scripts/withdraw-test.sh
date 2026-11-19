#!/usr/bin/env bash
# withdraw-test.sh — run scripts/withdraw-test.js with the deployer key taken from env/secrets.env (environment only).
# Needs the whole stack up (scripts/morph-native.sh status). Waits for the withdrawal batch to be finalized on L1.
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MORPH="${MORPH:-$(cd "${ROOT}/../.." && pwd)}"
# shellcheck disable=SC1090
source "${ROOT}/scripts/lib/env.sh"
set -a; source "${ROOT}/env/secrets.env"; set +a
export KEY="${DEPLOYER_PRIVATE_KEY}" MORPH_CONTRACTS="${MORPH}/contracts" L1_RUN="${L1_RUN:-${ROOT}/run/l1}"
export L1_RPC="${L1_RPC:-http://127.0.0.1:18546}" L2_RPC="${L2_RPC:-http://127.0.0.1:8545}"
exec node "${ROOT}/scripts/withdraw-test.js"
