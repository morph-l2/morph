#!/usr/bin/env bash
# build-all.sh — build every binary and the contracts toolchain into bin/ (steps from TODO.md P1). Logs: run/build-logs/.
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MORPH="${MORPH:-$(cd "${ROOT}/../.." && pwd)}"
source "${ROOT}/scripts/lib/env.sh"
require_build_tools
L="${ROOT}/run/build-logs"; mkdir -p "${L}" "${ROOT}"/bin/{morph-geth,morphnode,tx-submitter,gas-oracle}
step() { echo "$(date +%H:%M:%S) build: $1"; }
step morph-geth;   ( cd "${MORPH}/go-ethereum" && env GO111MODULE=on GOWORK=off go run build/ci.go install ./cmd/geth ) >"${L}/p1-2-morph-geth.log" 2>&1
cp "${MORPH}/go-ethereum/build/bin/geth" "${ROOT}/bin/morph-geth/geth"
step node;         ( cd "${MORPH}/node" && make build ) >"${L}/p1-3-node.log" 2>&1
cp "${MORPH}/node/build/bin/morphnode" "${MORPH}/node/build/bin/tendermint" "${ROOT}/bin/morphnode/"
step tx-submitter; ( cd "${MORPH}/tx-submitter" && make build ) >"${L}/p1-4-tx-submitter.log" 2>&1
cp "${MORPH}/tx-submitter/tx-submitter" "${ROOT}/bin/tx-submitter/"
step gas-oracle;   ( cd "${MORPH}/gas-oracle" && make build ) >"${L}/p1-5-gas-oracle.log" 2>&1
cp "${MORPH}/gas-oracle/app/target/release/app" "${ROOT}/bin/gas-oracle/"
step contracts;    ( cd "${MORPH}/contracts" && yarn install --frozen-lockfile && yarn build ) >"${L}/p1-6-contracts.log" 2>&1
step done
"${ROOT}/bin/morph-geth/geth" version | sed -n 1,2p; "${ROOT}/bin/morphnode/tendermint" version
ls -la "${ROOT}"/bin/*/ | grep -v '^total'
git -C "${MORPH}" status --short
