#!/usr/bin/env bash
#
# Builds morphnode with the fast-sync patches applied to its dependencies. The patched sources go
# through a Go build overlay, so go.mod and the regular `make morphnode` build stay untouched.
#
# Usage: build.sh <output> <ldflags>

set -euo pipefail

FAST_SYNC_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
OUTPUT="$1"
LDFLAGS="$2"

cd "${FAST_SYNC_DIR}/.."

WORK_DIR="$(mktemp -d)"
trap 'rm -rf "${WORK_DIR}"' EXIT

TENDERMINT_DIR="$(go list -m -f '{{.Dir}}' github.com/tendermint/tendermint)"
TM_DB_DIR="$(go list -m -f '{{.Dir}}' github.com/tendermint/tm-db)"

mkdir -p "${WORK_DIR}/tendermint/blocksync" "${WORK_DIR}/tm-db"
cp "${TENDERMINT_DIR}/blocksync/pool.go" "${WORK_DIR}/tendermint/blocksync/"
cp "${TM_DB_DIR}/goleveldb.go" "${TM_DB_DIR}/goleveldb_batch.go" "${WORK_DIR}/tm-db/"
chmod -R u+w "${WORK_DIR}"

patch -s -p1 -d "${WORK_DIR}/tendermint" < "${FAST_SYNC_DIR}/tendermint-blocksync.patch"
patch -s -p1 -d "${WORK_DIR}/tm-db" < "${FAST_SYNC_DIR}/tm-db-no-sync.patch"

cat > "${WORK_DIR}/overlay.json" <<JSON
{
  "Replace": {
    "${TENDERMINT_DIR}/blocksync/pool.go": "${WORK_DIR}/tendermint/blocksync/pool.go",
    "${TM_DB_DIR}/goleveldb.go": "${WORK_DIR}/tm-db/goleveldb.go",
    "${TM_DB_DIR}/goleveldb_batch.go": "${WORK_DIR}/tm-db/goleveldb_batch.go"
  }
}
JSON

env GO111MODULE=on CGO_ENABLED=1 CGO_LDFLAGS="-ldl" \
  go build -overlay "${WORK_DIR}/overlay.json" -o "${OUTPUT}" -v -ldflags "${LDFLAGS}" ./cmd/node
