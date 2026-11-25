#!/usr/bin/env bash
#
# morph-native.sh — 用原生进程运行 Morph L2(L1 = 公共 Hoodi),代替上游的 docker compose。
#
# 进程(启动顺序):
#   l1-proxy → morph-el-0 → node-0 → morph-el-1 → node-1 → morph-el-2 → node-2 → tx-submitter → gas-price-oracle
#   node-2 exists only so that node-0 can be restarted: see notes/issues.md I-028
#   l1-proxy 是本机的 rpc-consistency-proxy.js:所有进程访问 L1 都走它(限流时切换备用端点、读写一致性),见 notes/issues.md I-018~I-021
#   node-0 是 tendermint 验证者(出块);node-1 以 layer1 校验模式运行,不启动 tendermint。
#   可选进程 challenger(不在上面的默认列表里):比对 L1/L2 的 batch 根并在不一致时发起 challengeState。
#   用 `up challenger` 启动,或在 env/secrets.env 里设 CHALLENGER_ENABLE=1 让 `up` 一并启动。仅监听 127.0.0.1。
#
# 用法:
#   morph-native.sh setup            生成 tendermint / geth 运行目录(用 deploy-l1.sh 产出的创世文件)
#   morph-native.sh up [NAME...]     启动全部或指定进程(challenger 需显式指定或设 CHALLENGER_ENABLE=1)
#   morph-native.sh down             停止全部进程(保留链数据)
#   morph-native.sh status           进程、L2 高度、L1 同步高度、Rollup 批次
#   morph-native.sh logs NAME        跟随某个进程的日志
#   morph-native.sh clean logs       把日志挪到 run/archive/(需先 down)
#   morph-native.sh clean l2         把 L2 链数据挪到 run/archive/(需先 down)。**会让 L2 与 L1 上的 Rollup 对不上**,见下
#   morph-native.sh kill NAME        停掉某个进程(不会被自动重启)
#   morph-native.sh start NAME       启动某个进程
#
# 路径(可用环境变量改):
#   L1_RUN  deploy-l1.sh 的输出目录,默认 run/l1(其中 deployment/hoodidev.json、genesis/、state/start-block.dec)
#   L2_RUN  本脚本的运行目录,默认 run/l2
#
# 密钥取自 env/secrets.env,只经环境变量传给子进程,不出现在命令行参数里。

set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MORPH="${MORPH:-$(cd "${ROOT}/../.." && pwd)}"
BIN="${ROOT}/bin"
L1_RUN="${L1_RUN:-${ROOT}/run/l1}"
L2_RUN="${L2_RUN:-${ROOT}/run/l2}"
PIDS="${L2_RUN}/pids"
LOGS="${L2_RUN}/logs"
STATE="${L2_RUN}/state"
SECRETS="${SECRETS:-${ROOT}/env/secrets.env}"

ALL=(l1-proxy morph-el-0 node-0 morph-el-1 node-1 morph-el-2 node-2 tx-submitter gas-price-oracle)
OPTIONAL=(challenger)   # not started by a plain `up` unless CHALLENGER_ENABLE=1
CHAIN_ID=53077
JWT="${L2_RUN}/jwt-secret.txt"

# shellcheck disable=SC1090
source "${ROOT}/scripts/lib/env.sh"

die() { echo "error: $*" >&2; exit 1; }
say() { printf '%s %s\n' "$(date +%H:%M:%S)" "$*"; }

load_secrets() {
  [[ -f "${SECRETS}" ]] || die "missing ${SECRETS}"
  set -a
  # shellcheck disable=SC1090
  source "${SECRETS}"
  set +a
  : "${L1_ETH_RPC:?L1_ETH_RPC not set in secrets}" "${L1_BEACON_RPC:?L1_BEACON_RPC not set in secrets}"
  # every process talks to L1 through the local consistency proxy; the real endpoints become its upstreams
  L1_UPSTREAM="${L1_ETH_RPC}"
  L1_FALLBACKS="${RPC_FALLBACK_URLS-https://0xrpc.io/hoodi,https://ethereum-hoodi-rpc.publicnode.com}"
  L1_PROXY_PORT="${PROXY_PORT:-18546}"
  L1_ETH_RPC="http://127.0.0.1:${L1_PROXY_PORT}"
}

# deployed address by name from the deployment file written by deploy-l1.sh
addr() {
  node -e '
    const d = JSON.parse(require("fs").readFileSync(process.argv[1], "utf8")).find(x => x.name === process.argv[2]);
    if (!d) { console.error("not in deployment: " + process.argv[2]); process.exit(1); }
    console.log(d.address);' "${L1_RUN}/deployment/hoodidev.json" "$1"
}

# ---------------------------------------------------------------- ports (PLAN.md §6)
idx_of() { case "$1" in morph-el-0|node-0) echo 0;; morph-el-1|node-1) echo 1;; morph-el-2|node-2) echo 2;; *) echo 0;; esac; }
el_http() { echo $((8545 + 100 * $1)); }
el_ws() { echo $((8546 + 100 * $1)); }
el_auth() { echo $((8551 + 100 * $1)); }
el_p2p() { echo $((30303 + $1)); }
el_metrics() { echo $((6160 + $1)); }   # not 6060: the Coder workspace agent listens there

# ---------------------------------------------------------------- process control
alive() { [[ -f "${PIDS}/$1.pid" ]] && kill -0 "$(cat "${PIDS}/$1.pid")" 2>/dev/null; }

# spawn NAME ENV... -- CMD...   : run under a tiny restart supervisor in its own process group
spawn() {
  local name="$1"; shift
  local envs=()
  while [[ "$1" != "--" ]]; do envs+=("$1"); shift; done
  shift
  mkdir -p "${PIDS}" "${LOGS}" "${STATE}"
  if alive "${name}"; then say "${name}: already running (pid $(cat "${PIDS}/${name}.pid"))"; return 0; fi
  # cwd = state dir, so programs that write relative files (tx-submitter: journal.rlp, leveldb) stay inside run/
  cd "${STATE}"
  setsid bash -c '
    echo $$ > "$1"; shift
    while :; do
      "$@"; rc=$?
      echo "--- $(date +%FT%T) exited with status ${rc}; restarting in 2s"
      sleep 2
    done' _ "${PIDS}/${name}.pid" env "${envs[@]}" "$@" >>"${LOGS}/${name}.log" 2>&1 &
  for _ in $(seq 1 20); do [[ -s "${PIDS}/${name}.pid" ]] && break; sleep 0.1; done
  say "${name}: started (pid $(cat "${PIDS}/${name}.pid" 2>/dev/null || echo ?))"
}

stop_one() {
  local name="$1"
  if alive "${name}"; then
    local pid; pid="$(cat "${PIDS}/${name}.pid")"
    kill -TERM -- "-${pid}" 2>/dev/null || kill -TERM "${pid}" 2>/dev/null || true
    for _ in $(seq 1 50); do kill -0 "${pid}" 2>/dev/null || break; sleep 0.2; done
    kill -KILL -- "-${pid}" 2>/dev/null || true
    say "${name}: stopped"
  else
    say "${name}: not running"
  fi
  mv -f "${PIDS}/${name}.pid" "${PIDS}/${name}.pid.stopped" 2>/dev/null || true
}

wait_port() { # wait_port PORT [SECONDS]
  local end=$(( $(date +%s) + ${2:-60} ))
  until (exec 3<>"/dev/tcp/127.0.0.1/$1") 2>/dev/null; do
    [[ "$(date +%s)" -lt "${end}" ]] || die "port $1 did not open in time"
    sleep 1
  done
}

upgrade_time_ms() { # fixed on first start so restarts reuse it (see morph ops/README: never push this into the future)
  mkdir -p "${STATE}"
  [[ -s "${STATE}/sequencer-upgrade-time.ms" ]] || echo $(( $(date +%s) * 1000 )) >"${STATE}/sequencer-upgrade-time.ms"
  cat "${STATE}/sequencer-upgrade-time.ms"
}

# ---------------------------------------------------------------- per-process launchers
launch() {
  local name="$1" i; i="$(idx_of "${name}")"
  case "${name}" in
    l1-proxy)
      spawn l1-proxy -- node "${ROOT}/scripts/rpc-consistency-proxy.js" "${L1_UPSTREAM}${L1_FALLBACKS:+,${L1_FALLBACKS}}" "${L1_PROXY_PORT}"
      wait_port "${L1_PROXY_PORT}" 15
      ;;
    morph-el-0|morph-el-1|morph-el-2)
      spawn "${name}" -- "${BIN}/morph-geth/geth" --datadir "${L2_RUN}/el${i}" --verbosity=3 \
        --http --http.corsdomain='*' --http.vhosts='*' --http.addr=127.0.0.1 --http.port="$(el_http "${i}")" \
        --http.api=web3,debug,eth,txpool,net,morph,engine,admin \
        --ws --ws.addr=127.0.0.1 --ws.port="$(el_ws "${i}")" --ws.origins='*' --ws.api=web3,debug,eth,txpool,net,morph,engine,admin \
        --networkid="${CHAIN_ID}" \
        --authrpc.addr=127.0.0.1 --authrpc.port="$(el_auth "${i}")" --authrpc.vhosts='*' --authrpc.jwtsecret="${JWT}" \
        --gcmode=archive --nodiscover --miner.gasprice=0 --port="$(el_p2p "${i}")" \
        --metrics --metrics.expensive --metrics.addr=127.0.0.1 --metrics.port="$(el_metrics "${i}")"
      wait_port "$(el_http "${i}")" 60
      ;;
    node-0|node-1|node-2)
      local envs=(
        "MORPH_NODE_L2_ETH_RPC=http://127.0.0.1:$(el_http "${i}")"
        "MORPH_NODE_L2_ENGINE_RPC=http://127.0.0.1:$(el_auth "${i}")"
        "MORPH_NODE_L2_ENGINE_AUTH=${JWT}"
        "MORPH_NODE_L1_ETH_RPC=${L1_ETH_RPC}"
        "MORPH_NODE_L1_ETH_BEACON_RPC=${L1_BEACON_RPC}"
        "MORPH_NODE_SYNC_DEPOSIT_CONTRACT_ADDRESS=$(addr Proxy__L1MessageQueueWithGasPriceOracle)"
        "MORPH_NODE_L1_CONFIRMATIONS=0"
        "MORPH_NODE_ROLLUP_ADDRESS=$(addr Proxy__Rollup)"
        "MORPH_NODE_SYNC_START_HEIGHT=$(cat "${L1_RUN}/state/start-block.dec")"
        "MORPH_NODE_UPGRADE_BATCH_TIME=0"
        "MORPH_NODE_L1_SEQUENCER_CONTRACT=$(addr Proxy__L1Sequencer)"
        "MORPH_NODE_SEQUENCER_UPGRADE_TIME=$(upgrade_time_ms)"
      )
      if [[ "${name}" == node-0 ]]; then
        envs+=("MORPH_NODE_SEQUENCER_PRIVATE_KEY=${SEQUENCER_PRIVATE_KEY}")
      elif [[ "${name}" == node-2 ]]; then
        # an ordinary full node: runs tendermint (so node-0's block pool always has a peer after a restart), follows node-0,
        # and holds NO sequencer key (only one node may, unless HA is on)
        :
      else
        # layer1 verify mode: no tendermint. Its metrics port defaults to 26660, the same as node-0's tendermint prometheus
        envs+=("EMPTY_BLOCK_DELAY=true" "MORPH_NODE_DERIVATION_VERIFY_MODE=layer1"
               "MORPH_NODE_METRICS_SERVER_ENABLE=true" "MORPH_NODE_METRICS_PORT=27660")
      fi
      spawn "${name}" "${envs[@]}" -- "${BIN}/morphnode/morphnode" --home "${L2_RUN}/nodes/node${i}"
      ;;
    tx-submitter)
      spawn tx-submitter \
        "TX_SUBMITTER_BUILD_ENV=dev" \
        "TX_SUBMITTER_L1_ETH_RPC=${L1_ETH_RPC}" \
        "TX_SUBMITTER_L1_PRIVATE_KEY=${BATCH_SUBMITTER_PRIVATE_KEY}" \
        "TX_SUBMITTER_L2_ETH_RPCS=http://127.0.0.1:$(el_http 0),http://127.0.0.1:$(el_http 1)" \
        "TX_SUBMITTER_TX_TIMEOUT=60s" \
        "TX_SUBMITTER_ROLLUP_ADDRESS=$(addr Proxy__Rollup)" \
        "TX_SUBMITTER_SUBMITTER_ADDRESS=$(addr Proxy__Submitter)" \
        "TX_SUBMITTER_BATCH_BLOCK_INTERVAL=$(node -e 'console.log(require(process.argv[1]).govBatchBlockInterval)' "${MORPH}/ops/l2-genesis/deploy-config/hoodidev-deploy-config.json")" \
        "TX_SUBMITTER_BATCH_TIMEOUT=$(node -e 'console.log(require(process.argv[1]).govBatchTimeout)' "${MORPH}/ops/l2-genesis/deploy-config/hoodidev-deploy-config.json")" \
        "TX_SUBMITTER_FINALIZE=true" "TX_SUBMITTER_SEAL_BATCH=true" \
        "TX_SUBMITTER_TX_FEE_LIMIT=${TX_FEE_LIMIT_WEI:-50000000000000000}" \
        "TX_SUBMITTER_BATCH_V2_UPGRADE_TIME=1777533291" \
        "TX_SUBMITTER_METRICS_SERVER_ENABLE=true" "TX_SUBMITTER_METRICS_HOSTNAME=127.0.0.1" "TX_SUBMITTER_METRICS_PORT=8060" \
        -- "${BIN}/tx-submitter/tx-submitter"
      ;;
    gas-price-oracle)
      mkdir -p "${LOGS}/gas-oracle-files"
      spawn gas-price-oracle \
        "GAS_ORACLE_L1_RPC=${L1_ETH_RPC}" \
        "GAS_ORACLE_L1_BEACON_RPC=${L1_BEACON_RPC%%,*}" \
        "GAS_ORACLE_L2_RPC=http://127.0.0.1:$(el_http 0)" \
        "GAS_THRESHOLD=5" "INTERVAL=28000" "OVERHEAD_INTERVAL=10" \
        "L2_GAS_PRICE_ORACLE=0x530000000000000000000000000000000000000F" \
        "L2_GAS_ORACLE_PRIVATE_KEY=${L2_GAS_ORACLE_PRIVATE_KEY}" \
        "L1_ROLLUP=$(addr Proxy__Rollup)" \
        "GAS_ORACLE_METRIC_ADDRESS=127.0.0.1:6070" \
        "GAS_ORACLE_LOG_DIR=${LOGS}/gas-oracle-files" \
        -- "${BIN}/gas-oracle/app"
      ;;
    challenger)
      [[ -n "${CHALLENGER_PRIVATE_KEY:-}" ]] || die "CHALLENGER_PRIVATE_KEY is not set in ${SECRETS} (see scripts/register-challenger.sh)"
      [[ -x "${BIN}/challenger/challenger" ]] || die "missing ${BIN}/challenger/challenger: run scripts/build-all.sh"
      # The service reads archive state from el-0 and sends challengeState through the l1-proxy. It is open (no auth) by
      # default, so it only listens on 127.0.0.1. Ports 8080/6060 are avoided on purpose (see the port table in RUNBOOK.md).
      spawn challenger \
        "CHALLENGER_L1_ETH_RPC=${L1_ETH_RPC}" \
        "CHALLENGER_L2_ETH_RPC=http://127.0.0.1:$(el_http 0)" \
        "CHALLENGER_ROLLUP_ADDRESS=$(addr Proxy__Rollup)" \
        "CHALLENGER_PRIVATE_KEY=${CHALLENGER_PRIVATE_KEY}" \
        "CHALLENGER_HTTP_HOSTNAME=127.0.0.1" "CHALLENGER_HTTP_PORT=${CHALLENGER_HTTP_PORT:-8090}" \
        "CHALLENGER_METRICS_SERVER_ENABLE=true" "CHALLENGER_METRICS_HOSTNAME=127.0.0.1" "CHALLENGER_METRICS_PORT=${CHALLENGER_METRICS_PORT:-6065}" \
        "CHALLENGER_LOG_TERMINAL=true" \
        -- "${BIN}/challenger/challenger"
      ;;
    *) die "unknown process: ${name} (one of: ${ALL[*]} ${OPTIONAL[*]})" ;;
  esac
}

# ---------------------------------------------------------------- commands
cmd_setup() {
  [[ -s "${L1_RUN}/genesis/genesis-l2.json" ]] || die "no ${L1_RUN}/genesis/genesis-l2.json — run deploy-l1.sh first"
  python3 -I "${ROOT}/scripts/setup-nodes.py" --morph "${MORPH}" --run-dir "${L2_RUN}" \
    --genesis "${L1_RUN}/genesis/genesis-l2.json" \
    --tendermint "${BIN}/morphnode/tendermint" --geth "${BIN}/morph-geth/geth"
  # node-0 cannot resume block production after a restart unless another tendermint node is its peer (notes/issues.md I-028)
  python3 -I "${ROOT}/scripts/add-node2.py" --morph "${MORPH}" --run-dir "${L2_RUN}" \
    --tendermint "${BIN}/morphnode/tendermint" --geth "${BIN}/morph-geth/geth"
}

# clean: never deletes anything -- it MOVES data into run/archive/<what>-<timestamp>/ so a mistake can be undone.
cmd_clean() {
  local what="${1:-}" p ts dest
  for p in "${ALL[@]}" "${OPTIONAL[@]}"; do alive "${p}" && die "${p} is still running: run 'down' first"; done
  ts="$(date +%Y%m%d-%H%M%S)"; dest="${ROOT}/run/archive/${what}-${ts}"
  case "${what}" in
    logs)
      [[ -d "${LOGS}" ]] || die "no logs at ${LOGS}"
      mkdir -p "${dest}" && mv "${LOGS}" "${dest}/logs" && say "logs moved to ${dest}" ;;
    l2)
      [[ -d "${L2_RUN}" ]] || die "nothing at ${L2_RUN}"
      cat >&2 <<WARN
WARNING: this moves the whole L2 chain (execution-client data, tendermint data, pids, logs) out of the way.
The L1 Rollup on Hoodi still holds every batch committed for THIS chain. A fresh L2 starts from genesis again, its blocks
will not match those batches, and tx-submitter cannot continue from them. A fresh L2 therefore needs a fresh L1
deployment: run scripts/deploy-l1.sh with a NEW RUN_DIR, point L1_RUN at it, then 'morph-native.sh setup' and 'up'.
Nothing is deleted: the data goes to ${dest}.
WARN
      mkdir -p "${dest}" && mv "${L2_RUN}" "${dest}/l2" && say "L2 data moved to ${dest}" ;;
    *) die "usage: clean logs|l2" ;;
  esac
}

cmd_up() {
  [[ -f "${L1_RUN}/deployment/hoodidev.json" ]] || die "no L1 deployment at ${L1_RUN}/deployment/hoodidev.json"
  [[ -d "${L2_RUN}/nodes/node0" ]] || die "no ${L2_RUN}/nodes — run: morph-native.sh setup"
  load_secrets
  local names=("$@"); [[ ${#names[@]} -gt 0 ]] || { names=("${ALL[@]}"); [[ "${CHALLENGER_ENABLE:-0}" == "1" ]] && names+=("${OPTIONAL[@]}"); }
  for n in "${names[@]}"; do launch "${n}"; done
}

cmd_down() {
  local i p
  for p in "${OPTIONAL[@]}"; do alive "${p}" && stop_one "${p}"; done
  for ((i = ${#ALL[@]} - 1; i >= 0; i--)); do stop_one "${ALL[i]}"; done
}

rpc_height() { cast block-number --rpc-url "http://127.0.0.1:$1" 2>/dev/null || echo "-"; }

cmd_status() {
  local n
  for n in "${ALL[@]}"; do
    if alive "${n}"; then printf '  %-18s running (pid %s)\n' "${n}" "$(cat "${PIDS}/${n}.pid")"; else printf '  %-18s stopped\n' "${n}"; fi
  done
  for n in "${OPTIONAL[@]}"; do
    if alive "${n}"; then printf '  %-18s running (pid %s)  [optional]\n' "${n}" "$(cat "${PIDS}/${n}.pid")"; else printf '  %-18s stopped  [optional]\n' "${n}"; fi
  done
  echo ">>> L2 heights:  el-0 :$(el_http 0) -> $(rpc_height "$(el_http 0)")    el-1 :$(el_http 1) -> $(rpc_height "$(el_http 1)")    el-2 :$(el_http 2) -> $(rpc_height "$(el_http 2)")"
  if [[ -f "${SECRETS}" ]]; then
    load_secrets
    echo ">>> L1 (Hoodi) head: $(cast block-number --rpc-url "${L1_UPSTREAM}" 2>/dev/null || echo '-')"
    if alive l1-proxy; then echo ">>> l1-proxy stats: $(curl -s -m 3 "http://127.0.0.1:${L1_PROXY_PORT}/stats" | cut -c1-200)"; fi
    if [[ -f "${L1_RUN}/deployment/hoodidev.json" ]]; then
      local rollup; rollup="$(addr Proxy__Rollup)"
      echo ">>> Rollup ${rollup}: lastCommittedBatchIndex = $(cast call "${rollup}" 'lastCommittedBatchIndex()(uint256)' --rpc-url "${L1_UPSTREAM}" 2>/dev/null || echo '?')"
    fi
  fi
}

case "${1:-}" in
  setup)  cmd_setup ;;
  clean)  shift; cmd_clean "$@" ;;
  up)     shift; cmd_up "$@" ;;
  down)   cmd_down ;;
  status) cmd_status ;;
  logs)   [[ -n "${2:-}" ]] || die "usage: logs NAME"; tail -n 100 -F "${LOGS}/$2.log" ;;
  kill)   [[ -n "${2:-}" ]] || die "usage: kill NAME"; stop_one "$2" ;;
  start)  [[ -n "${2:-}" ]] || die "usage: start NAME"; cmd_up "$2" ;;
  *)      sed -n '2,30p' "${BASH_SOURCE[0]}" | sed 's/^# \{0,1\}//'; exit 1 ;;
esac
