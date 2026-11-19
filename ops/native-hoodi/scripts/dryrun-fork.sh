#!/usr/bin/env bash
#
# dryrun-fork.sh — P2-5:在 Hoodi 的 anvil 分叉上把 deploy-l1.sh 完整干跑一遍,量出真实的 gas / ETH 需求。
#
# anvil 从公共 Hoodi 分叉出一个本地链(链 ID 560048),给 deployer / submitter 凭空设置余额;
# 部署只发生在本地分叉里,**不会向公共 Hoodi 发送任何交易,也不花任何测试币**。
#
# 用法:dryrun-fork.sh [RUN_DIR]     (默认 run/dryrun-<时间戳>)
# 环境:FORK_URL 默认 https://rpc.hoodi.ethpandaops.io;PORT 默认 18545
#
# 退出时会停掉自己启动的 anvil(按 pid 文件)。

set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
FORK_URL="${FORK_URL:-https://rpc.hoodi.ethpandaops.io}"
PORT="${PORT:-18545}"
RUN_DIR="${1:-${ROOT}/run/dryrun-$(date +%Y%m%d-%H%M%S)}"
mkdir -p "${RUN_DIR}/logs"
# shellcheck disable=SC1090
source "${ROOT}/scripts/lib/env.sh"
command -v anvil >/dev/null 2>&1 || { echo "error: anvil (Foundry) is not on PATH" >&2; exit 1; }

PIDFILE="${RUN_DIR}/anvil.pid"
cleanup() { if [[ -f "${PIDFILE}" ]]; then kill "$(cat "${PIDFILE}")" 2>/dev/null || true; fi; }
trap cleanup EXIT

if (exec 3<>"/dev/tcp/127.0.0.1/${PORT}") 2>/dev/null; then echo "port ${PORT} is already in use" >&2; exit 1; fi

echo "starting anvil fork of ${FORK_URL} on 127.0.0.1:${PORT} (chain-id 560048)"
anvil --fork-url "${FORK_URL}" --chain-id 560048 --port "${PORT}" --host 127.0.0.1 >"${RUN_DIR}/logs/anvil.log" 2>&1 &
echo $! >"${PIDFILE}"
for _ in $(seq 1 60); do
  cast chain-id --rpc-url "http://127.0.0.1:${PORT}" >/dev/null 2>&1 && break
  sleep 1
done
RPC="http://127.0.0.1:${PORT}"
[[ "$(cast chain-id --rpc-url "${RPC}")" == "560048" ]] || { echo "anvil fork did not come up; see ${RUN_DIR}/logs/anvil.log" >&2; exit 1; }
echo "fork head: $(cast block-number --rpc-url "${RPC}")"

# give the two paying accounts a generous balance on the FORK ONLY (nothing touches public Hoodi)
MORPH="$(cd "${ROOT}/../.." && pwd)"
set -a; source "${ROOT}/env/secrets.env"; set +a
for v in DEPLOYER_PRIVATE_KEY BATCH_SUBMITTER_PRIVATE_KEY; do
  a="$( cd "${MORPH}/contracts" && KEY="${!v}" node -e 'console.log(new (require("ethers").Wallet)(process.env.KEY).address)' )"
  cast rpc anvil_setBalance "${a}" "$(cast to-hex 100000000000000000000)" --rpc-url "${RPC}" >/dev/null   # 100 ETH, fork only
done

RPC_URL="${RPC}" RUN_DIR="${RUN_DIR}" WAIT_FINALIZED=0 SUBMITTER_FUND_ETH="${SUBMITTER_FUND_ETH:-0}" \
  "${ROOT}/scripts/deploy-l1.sh"
