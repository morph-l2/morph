#!/usr/bin/env bash

# Run `morph-prove` continuous execution test in background.
#
# Example (foreground):
#   ./script/execute.sh run --start-block 0x35 --max-blocks 2 --rpc http://127.0.0.1:9545
#
# Example (background):
#   ./script/execute.sh start --start-block 0x35 --max-blocks 1000 --rpc http://127.0.0.1:9545
#   ./script/execute.sh status
#   ./script/execute.sh stop

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"

DEFAULT_START_BLOCK="0x35"
DEFAULT_MAX_BLOCKS="1000"
DEFAULT_RPC="http://127.0.0.1:9545"

LOG_DIR_DEFAULT="${REPO_ROOT}/logs"
PID_FILE_DEFAULT="${REPO_ROOT}/.execute_continuous.pid"
LOG_FILE_DEFAULT="${LOG_DIR_DEFAULT}/execute_continuous.log"

usage() {
  cat <<'EOF'
Usage:
  script/execute.sh <start|stop|status|run> [options]

Commands:
  start    Run in background (nohup) and write PID/log.
  stop     Stop the background process using PID file.
  status   Show whether the background process is running.
  run      Run in foreground.

Options:
  --start-block <u64>   Start L2 block number (decimal or 0x...).
  --max-blocks <u64>    Max blocks to execute.
  --rpc <url>           RPC endpoint.
  --pid-file <path>     PID file path.
  --log-file <path>     Log file path.
  --rust-log <level>    Set RUST_LOG (default: info).
  --backtrace <mode>    Set RUST_BACKTRACE (e.g. 1|full).
  --no-nocapture        Do NOT pass --nocapture to cargo test.

Environment:
  If <repo>/.env exists, it will be sourced when running.

Underlying test:
  cargo test -p morph-prove --lib -- execute::tests::test_execute_continuous --exact --nocapture -- \
    --start-block <...> --max-blocks <...> --rpc <...>
EOF
}

is_running_pid() {
  local pid="$1"
  [[ -n "${pid}" ]] || return 1
  kill -0 "${pid}" >/dev/null 2>&1
}

read_pid() {
  local pid_file="$1"
  [[ -f "${pid_file}" ]] || return 1
  tr -d '[:space:]' <"${pid_file}" | head -n 1
}

cmd="${1:-}"
shift || true

start_block="${DEFAULT_START_BLOCK}"
max_blocks="${DEFAULT_MAX_BLOCKS}"
rpc="${DEFAULT_RPC}"
pid_file="${PID_FILE_DEFAULT}"
log_file="${LOG_FILE_DEFAULT}"
rust_log="info"
backtrace=""
nocapture_flag="--nocapture"

while [[ $# -gt 0 ]]; do
  case "$1" in
    --start-block)
      start_block="$2"; shift 2 ;;
    --max-blocks)
      max_blocks="$2"; shift 2 ;;
    --rpc)
      rpc="$2"; shift 2 ;;
    --pid-file)
      pid_file="$2"; shift 2 ;;
    --log-file)
      log_file="$2"; shift 2 ;;
    --rust-log)
      rust_log="$2"; shift 2 ;;
    --backtrace)
      backtrace="$2"; shift 2 ;;
    --no-nocapture)
      nocapture_flag=""; shift 1 ;;
    -h|--help)
      usage; exit 0 ;;
    *)
      echo "Unknown argument: $1" >&2
      usage
      exit 2
      ;;
  esac
done

test_cmd=(
  cargo test -p morph-prove --lib -- execute::tests::test_execute_continuous --exact
)
if [[ -n "${nocapture_flag}" ]]; then
  test_cmd+=("${nocapture_flag}")
fi
test_cmd+=(-- --start-block "${start_block}" --max-blocks "${max_blocks}" --rpc "${rpc}")

run_foreground() {
  cd "${REPO_ROOT}"

  # Load env vars if present.
  if [[ -f "${REPO_ROOT}/.env" ]]; then
    # shellcheck disable=SC1091
    set -a
    source "${REPO_ROOT}/.env"
    set +a
  fi

  export RUST_LOG="${RUST_LOG:-${rust_log}}"
  if [[ -n "${backtrace}" ]]; then
    export RUST_BACKTRACE="${backtrace}"
  fi

  echo "[execute.sh] repo_root=${REPO_ROOT}"
  echo "[execute.sh] cmd: ${test_cmd[*]}"
  "${test_cmd[@]}"
}

start_background() {
  mkdir -p "$(dirname "${pid_file}")" "$(dirname "${log_file}")" "${LOG_DIR_DEFAULT}" || true

  if [[ -f "${pid_file}" ]]; then
    local old_pid
    old_pid="$(read_pid "${pid_file}" || true)"
    if [[ -n "${old_pid}" ]] && is_running_pid "${old_pid}"; then
      echo "[execute.sh] already running: pid=${old_pid} (pid_file=${pid_file})"
      echo "[execute.sh] log_file=${log_file}"
      return 0
    fi
    rm -f "${pid_file}" || true
  fi

  # Use a subshell so we can source .env before running cargo.
  # Use nohup + background, and store PID.
  (
    cd "${REPO_ROOT}"
    if [[ -f "${REPO_ROOT}/.env" ]]; then
      # shellcheck disable=SC1091
      set -a
      source "${REPO_ROOT}/.env"
      set +a
    fi
    export RUST_LOG="${RUST_LOG:-${rust_log}}"
    if [[ -n "${backtrace}" ]]; then
      export RUST_BACKTRACE="${backtrace}"
    fi
    nohup "${test_cmd[@]}" >>"${log_file}" 2>&1 &
    echo $! >"${pid_file}"
  )

  local pid
  pid="$(read_pid "${pid_file}" || true)"
  echo "[execute.sh] started: pid=${pid}"
  echo "[execute.sh] pid_file=${pid_file}"
  echo "[execute.sh] log_file=${log_file}"
}

stop_background() {
  local pid
  pid="$(read_pid "${pid_file}" || true)"
  if [[ -z "${pid}" ]]; then
    echo "[execute.sh] not running (pid file missing/empty): ${pid_file}"
    return 0
  fi

  if ! is_running_pid "${pid}"; then
    echo "[execute.sh] not running (stale pid=${pid}), removing pid file"
    rm -f "${pid_file}" || true
    return 0
  fi

  echo "[execute.sh] stopping pid=${pid}"
  kill "${pid}" || true

  # Wait a bit then force kill if still alive.
  for _ in {1..30}; do
    if ! is_running_pid "${pid}"; then
      rm -f "${pid_file}" || true
      echo "[execute.sh] stopped"
      return 0
    fi
    sleep 1
  done

  echo "[execute.sh] force killing pid=${pid}"
  kill -9 "${pid}" || true
  rm -f "${pid_file}" || true
  echo "[execute.sh] stopped"
}

status_background() {
  local pid
  pid="$(read_pid "${pid_file}" || true)"
  if [[ -z "${pid}" ]]; then
    echo "[execute.sh] status: not running (pid file missing/empty): ${pid_file}"
    return 1
  fi
  if is_running_pid "${pid}"; then
    echo "[execute.sh] status: running pid=${pid}"
    echo "[execute.sh] pid_file=${pid_file}"
    echo "[execute.sh] log_file=${log_file}"
    return 0
  fi
  echo "[execute.sh] status: not running (stale pid=${pid})"
  return 1
}

case "${cmd}" in
  start)
    start_background
    ;;
  stop)
    stop_background
    ;;
  status)
    status_background
    ;;
  run)
    run_foreground
    ;;
  "")
    usage
    exit 2
    ;;
  *)
    echo "Unknown command: ${cmd}" >&2
    usage
    exit 2
    ;;
esac
