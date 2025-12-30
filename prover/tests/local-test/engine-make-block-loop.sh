#!/bin/sh
# Loop to call engine-make-block.sh periodically (default: every 200ms)
# - Run from local-test directory
# - Respects ENV vars used by engine-make-block.sh (HTTP_RPC, AUTH_RPC, JWT_SECRET_PATH)

set -e

DIR=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
INTERVAL="${INTERVAL:-3}"

echo "[info] starting loop: calling engine-make-block.sh every ${INTERVAL}s (Ctrl-C to stop)" >&2

while true; do
  # Don't exit the whole loop if a single run fails
  set +e
  bash "$DIR/engine-make-block.sh"
  rc=$?
  set -e
  # Timestamp + exit code for observability
  date "+[%F %T] engine-make-block.sh exit code: ${rc}"
  sleep "$INTERVAL"
done
