#!/usr/bin/env bash
#
# deploy-l1.sh — 在 L1(公共 Hoodi,或其 anvil 分叉)上部署 Morph 合约,并生成 L2 创世。
#
# 对应上游 ops/devnet-morph/devnet/__init__.py 的 devnet_deploy(),但:
#   - 不调用 docker compose,不改写仓库里被 git 跟踪的 .env / l1.ts(batchHeader 走环境变量 MORPH_BATCH_HEADER)
#   - 不调用 `hardhat fund`(它会转 100 ETH);提交者资金由 SUBMITTER_FUND_ETH 控制,默认不转
#   - 每一步幂等:成功后在 $RUN_DIR/state/ 落标记文件,重跑会跳过已完成的步骤
#   - 私钥只通过环境变量传递,不出现在命令行参数或输出里
#
# 必需环境变量:
#   RPC_URL   L1 执行层 RPC(Hoodi 或 anvil 分叉),链 ID 必须是 560048
#   RUN_DIR   本次运行的输出目录(部署产物、创世文件、日志、状态)
# 可选:
#   SECRETS            密钥文件,默认 env/secrets.env
#   WAIT_FINALIZED     1(默认)= 部署完等 L1 finalized 越过当前块;分叉干跑设 0
#   FINALIZED_TIMEOUT_SEC  等待 finalized 的超时,默认 1800(Hoodi 约需 13 分钟)
#   SUBMITTER_FUND_ETH 向 batch submitter 转多少 ETH(默认 0 = 不转)
#   STOP_AFTER         执行到指定步骤后停止,如 s21_batchheader(用于分段验证)
#   USE_CONSISTENCY_PROXY  auto(默认:RPC 是 https 就启用)/ 1 / 0。启用时在本机起 rpc-consistency-proxy.js,
#                      所有读都经它转发,解决公共 RPC 多节点造成的"刚上链就读不到"(见 notes/issues.md I-018、I-019)
#   PROXY_PORT         代理端口,默认 18546
#   RPC_FALLBACK_URLS  代理的备用上游(逗号分隔),主上游限流或不可达时自动切换。
#                      默认 0xrpc.io 与 publicnode;设为空字符串则不用备用
#
# 退出码:0 成功;1 某步骤失败(日志路径会打印)。

set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MORPH="${MORPH:-$(cd "${ROOT}/../.." && pwd)}"
: "${RPC_URL:?set RPC_URL (Hoodi EL RPC or an anvil fork of it)}"
: "${RUN_DIR:?set RUN_DIR}"
SECRETS="${SECRETS:-${ROOT}/env/secrets.env}"
WAIT_FINALIZED="${WAIT_FINALIZED:-1}"
FINALIZED_TIMEOUT_SEC="${FINALIZED_TIMEOUT_SEC:-1800}"
SUBMITTER_FUND_ETH="${SUBMITTER_FUND_ETH:-0}"
STOP_AFTER="${STOP_AFTER:-}"
EXPECT_L1_CHAIN_ID=560048

STATE="${RUN_DIR}/state"
LOGS="${RUN_DIR}/logs"
DEPLOY_DIR="${RUN_DIR}/deployment"
GENESIS_DIR="${RUN_DIR}/genesis"
DEPLOY_JSON="${DEPLOY_DIR}/hoodidev.json"
GEN_CONFIG="${RUN_DIR}/hoodidev-deploy-config.generated.json"
mkdir -p "${STATE}" "${LOGS}" "${DEPLOY_DIR}" "${GENESIS_DIR}"

# shellcheck disable=SC1090
source "${ROOT}/scripts/lib/env.sh"
[[ -f "${SECRETS}" ]] || { echo "missing ${SECRETS}" >&2; exit 1; }
set -a
# shellcheck disable=SC1090
source "${SECRETS}"
set +a
# Public RPCs sit behind several nodes; a read right after a mined write can hit a node that has not seen the block.
# contracts/src/deploy-utils.ts retries its assertion reads when these are set (default there: off).
export MORPH_CONTRACTS="${MORPH}/contracts"   # send-tx.js loads ethers from here
export MORPH_READ_RETRIES="${MORPH_READ_RETRIES:-15}"
export MORPH_READ_RETRY_DELAY_MS="${MORPH_READ_RETRY_DELAY_MS:-2000}"

# ---- read-after-write consistency proxy for public RPCs ----
USE_CONSISTENCY_PROXY="${USE_CONSISTENCY_PROXY:-auto}"
PROXY_PORT="${PROXY_PORT:-18546}"
UPSTREAM_RPC_URL="${RPC_URL}"
if [[ "${USE_CONSISTENCY_PROXY}" == "1" || ( "${USE_CONSISTENCY_PROXY}" == "auto" && "${RPC_URL}" == https://* ) ]]; then
  if (exec 3<>"/dev/tcp/127.0.0.1/${PROXY_PORT}") 2>/dev/null; then
    echo "port ${PROXY_PORT} is already in use; set PROXY_PORT or stop what is listening" >&2; exit 1
  fi
  # fallbacks are used when the primary is rate-limited or unreachable (public RPCs do rate-limit; see notes/issues.md I-020)
  RPC_FALLBACK_URLS="${RPC_FALLBACK_URLS-https://0xrpc.io/hoodi,https://ethereum-hoodi-rpc.publicnode.com}"
  node "${ROOT}/scripts/rpc-consistency-proxy.js" "${UPSTREAM_RPC_URL}${RPC_FALLBACK_URLS:+,${RPC_FALLBACK_URLS}}" "${PROXY_PORT}" >"${LOGS}/rpc-proxy.log" 2>&1 &
  PROXY_PID=$!
  trap 'kill "${PROXY_PID}" 2>/dev/null || true' EXIT
  for _ in $(seq 1 50); do (exec 3<>"/dev/tcp/127.0.0.1/${PROXY_PORT}") 2>/dev/null && break; sleep 0.1; done
  RPC_URL="http://127.0.0.1:${PROXY_PORT}"
  echo "$(date +%H:%M:%S) consistency proxy ${UPSTREAM_RPC_URL} -> ${RPC_URL} (pid ${PROXY_PID})"
fi

export HOODIDEV_RPC_URL="${RPC_URL}"
export ETH_RPC_URL="${RPC_URL}"
for v in DEPLOYER_PRIVATE_KEY SEQUENCER_PRIVATE_KEY BATCH_SUBMITTER_PRIVATE_KEY; do
  [[ "${!v:-}" =~ ^0x[0-9a-fA-F]{64}$ ]] || { echo "${v} is not set to a 32-byte hex key in ${SECRETS}" >&2; exit 1; }
done
# the challenger key is optional (older secrets files do not have one); if set it must be well-formed
[[ -z "${CHALLENGER_PRIVATE_KEY:-}" || "${CHALLENGER_PRIVATE_KEY}" =~ ^0x[0-9a-fA-F]{64}$ ]] || { echo "CHALLENGER_PRIVATE_KEY in ${SECRETS} is not a 32-byte hex key" >&2; exit 1; }

# addr_of ENVVAR: derive an address from the key held in that variable without putting the key on a command line
# (ethers from the contracts' node_modules reads the key from the environment; `cast wallet address` only takes it as an argument)
addr_of() { ( cd "${MORPH}/contracts" && KEY="${!1}" node -e 'console.log(new (require("ethers").Wallet)(process.env.KEY).address)' ); }
DEPLOYER="$(addr_of DEPLOYER_PRIVATE_KEY)"
SEQUENCER="$(addr_of SEQUENCER_PRIVATE_KEY)"
SUBMITTER="$(addr_of BATCH_SUBMITTER_PRIVATE_KEY)"

say() { printf '%s %s\n' "$(date +%H:%M:%S)" "$*"; }

# step NAME FUNCTION: run once; skip if its marker exists; log to $LOGS/NAME.log.
# The function runs in a subshell with `set -e` ON. (Calling it from an `if` would silently disable errexit and let a failed
# command in the middle of a step go unnoticed -- exactly how a failed `cast send` once showed up as "ok".)
step() {
  local name="$1" fn="$2" rc
  if [[ -f "${STATE}/${name}.done" ]]; then say "[skip] ${name} (already done)"; else
    say "[run ] ${name}"
    set +e
    ( set -e; "${fn}" ) >"${LOGS}/${name}.log" 2>&1
    rc=$?
    set -e
    if [[ "${rc}" -ne 0 ]]; then
      say "[FAIL] ${name} (exit ${rc}) -- last lines of ${LOGS}/${name}.log:"
      tail -n 15 "${LOGS}/${name}.log" | sed 's/^/        /' >&2
      exit 1
    fi
    touch "${STATE}/${name}.done"
    say "[ ok ] ${name}"
  fi
  if [[ "${STOP_AFTER}" == "${name}" ]]; then say "STOP_AFTER=${name}: stopping here"; exit 0; fi
}

hardhat() { ( cd "${MORPH}/contracts" && yarn -s hardhat "$@" ); }
wei_to_eth() { cast from-wei "$1" 2>/dev/null || echo "$1"; }

# ---------------------------------------------------------------- steps
s00_preflight() {
  local id bal start_dec
  id="$(cast chain-id --rpc-url "${RPC_URL}")"
  [[ "${id}" == "${EXPECT_L1_CHAIN_ID}" ]] || { echo "L1 chain id is ${id}, expected ${EXPECT_L1_CHAIN_ID}"; return 1; }
  bal="$(cast balance "${DEPLOYER}" --rpc-url "${RPC_URL}")"
  [[ "${bal}" != "0" ]] || { echo "deployer ${DEPLOYER} has no balance on this L1"; return 1; }
  echo "${bal}" >"${STATE}/deployer-balance-before.wei"
  cast balance "${SUBMITTER}" --rpc-url "${RPC_URL}" >"${STATE}/submitter-balance-before.wei"
  cast nonce "${DEPLOYER}" --rpc-url "${RPC_URL}" >"${STATE}/deployer-nonce-before"
  start_dec="$(cast block-number --rpc-url "${RPC_URL}")"
  echo "${start_dec}" >"${STATE}/start-block.dec"
  printf '0x%x\n' "${start_dec}" >"${STATE}/start-block.hex"
  # generated copy of the genesis deploy config with the real L1 starting block (the tracked file keeps a placeholder)
  node -e '
    const fs = require("fs");
    const [src, dst, tag] = process.argv.slice(1);
    const c = JSON.parse(fs.readFileSync(src, "utf8"));
    c.l1StartingBlockTag = tag;
    fs.writeFileSync(dst, JSON.stringify(c, null, 2) + "\n");
  ' "${MORPH}/ops/l2-genesis/deploy-config/hoodidev-deploy-config.json" "${GEN_CONFIG}" "$(cat "${STATE}/start-block.hex")"
  echo "chain ${id}; deployer ${DEPLOYER} balance $(wei_to_eth "${bal}") ETH; L1 start block ${start_dec}"
}

s10_deploy() {
  # hardhat's --storagepath appends to an existing file; start from nothing by moving any leftover of a failed run aside
  if [[ -e "${DEPLOY_JSON}" ]]; then mv "${DEPLOY_JSON}" "${DEPLOY_JSON}.prev-$(date +%s)"; fi
  hardhat deploy --network hoodidev --storagepath "${DEPLOY_JSON}" --concurrent false
}

s11_verify_deploy() {
  local n=0 bad=0 name addr
  while IFS=$'\t' read -r name addr; do
    n=$((n + 1))
    if [[ "$(cast code "${addr}" --rpc-url "${RPC_URL}")" == "0x" ]]; then echo "NO CODE: ${name} ${addr}"; bad=$((bad + 1)); fi
  done < <(node -e 'for (const d of JSON.parse(require("fs").readFileSync(process.argv[1], "utf8"))) console.log(d.name + "\t" + d.address)' "${DEPLOY_JSON}")
  echo "${n} deployed entries, ${bad} without code"
  [[ "${n}" -gt 0 && "${bad}" -eq 0 ]]
}

s20_genesis() {
  ( cd "${MORPH}/ops/l2-genesis" && env CGO_ENABLED=1 CGO_LDFLAGS="-ldl" go run cmd/main.go genesis l2 \
      --l1-rpc "${RPC_URL}" \
      --deploy-config "${GEN_CONFIG}" \
      --deployment-dir "${DEPLOY_JSON}" \
      --outfile.l2 "${GENESIS_DIR}/genesis-l2.json" \
      --outfile.genbatchheader "${GENESIS_DIR}/genesis-batch-header.json" \
      --outfile.rollup "${GENESIS_DIR}/rollup.json" )
  [[ -s "${GENESIS_DIR}/genesis-l2.json" && -s "${GENESIS_DIR}/rollup.json" ]]
}

s21_batchheader() {
  node -e '
    const r = JSON.parse(require("fs").readFileSync(process.argv[1], "utf8"));
    if (!r.genesis_batch_header) { console.error("no genesis_batch_header in rollup.json"); process.exit(1); }
    require("fs").writeFileSync(process.argv[2], r.genesis_batch_header);
    console.log("genesis_batch_header length", r.genesis_batch_header.length, "l2_genesis_state_root", r.l2_genesis_state_root);
  ' "${GENESIS_DIR}/rollup.json" "${STATE}/batch-header.hex"
}

s30_initialize() {
  [[ -s "${STATE}/batch-header.hex" ]] || { echo "batch header missing: run s21_batchheader first"; return 1; }
  MORPH_BATCH_HEADER="$(cat "${STATE}/batch-header.hex")" \
    hardhat initialize --network hoodidev --storagepath "${DEPLOY_JSON}" --concurrent false
}

s31_register() {
  # register: owner adds the submitter and STAKES submitterMinimumStake on its behalf (paid by the deployer).
  # NOT `hardhat fund`, which would send 100 ETH to the submitter.
  batchSubmitterPks="[\"${BATCH_SUBMITTER_PRIVATE_KEY}\"]" \
    hardhat register --network hoodidev --storagepath "${DEPLOY_JSON}"
}

s32_first_sequencer() {
  local seq_proxy owner cur data
  seq_proxy="$(node -e 'const d=JSON.parse(require("fs").readFileSync(process.argv[1],"utf8")).find(x=>x.name==="Proxy__L1Sequencer"); if(!d){process.exit(1)} console.log(d.address)' "${DEPLOY_JSON}")"
  owner="$(cast call "${seq_proxy}" 'owner()(address)' --rpc-url "${RPC_URL}")"
  [[ "${owner,,}" == "${DEPLOYER,,}" ]] || { echo "L1Sequencer owner ${owner} != deployer ${DEPLOYER}"; return 1; }
  cur="$(cast call "${seq_proxy}" 'getSequencer()(address)' --rpc-url "${RPC_URL}" 2>/dev/null || true)"
  if [[ "${cur,,}" == "${SEQUENCER,,}" ]]; then echo "first sequencer already ${SEQUENCER}"; return 0; fi
  data="$(cast calldata 'setFirstSequencer(address)' "${SEQUENCER}")"
  KEY="${DEPLOYER_PRIVATE_KEY}" node "${ROOT}/scripts/send-tx.js" "${RPC_URL}" "${seq_proxy}" --data "${data}"
  # do not trust the exit code alone: read the result back
  cur="$(cast call "${seq_proxy}" 'getSequencer()(address)' --rpc-url "${RPC_URL}")"
  [[ "${cur,,}" == "${SEQUENCER,,}" ]] || { echo "getSequencer() = ${cur}, expected ${SEQUENCER}"; return 1; }
  echo "L1Sequencer ${seq_proxy}: first sequencer is ${SEQUENCER}"
}

s33_fund_submitter() {
  if [[ "${SUBMITTER_FUND_ETH}" == "0" ]]; then echo "SUBMITTER_FUND_ETH=0: not funding the submitter (do it by hand)"; return 0; fi
  local before after want
  before="$(cast balance "${SUBMITTER}" --rpc-url "${RPC_URL}")"
  KEY="${DEPLOYER_PRIVATE_KEY}" node "${ROOT}/scripts/send-tx.js" "${RPC_URL}" "${SUBMITTER}" --value-ether "${SUBMITTER_FUND_ETH}"
  after="$(cast balance "${SUBMITTER}" --rpc-url "${RPC_URL}")"
  want="$(cast to-wei "${SUBMITTER_FUND_ETH}" ether)"
  node -e 'process.exit(BigInt(process.argv[2]) - BigInt(process.argv[1]) === BigInt(process.argv[3]) ? 0 : 1)' "${before}" "${after}" "${want}" ||
    { echo "submitter balance ${before} -> ${after}, expected +${want}"; return 1; }
  echo "submitter funded: $(wei_to_eth "${before}") -> $(wei_to_eth "${after}") ETH"
}

s34_register_challenger() {
  if [[ -z "${CHALLENGER_PRIVATE_KEY:-}" ]]; then echo "no CHALLENGER_PRIVATE_KEY in ${SECRETS}: not registering a challenger"; return 0; fi
  # same RPC as this script (the consistency proxy when it is on); the deployment file is the one this run wrote
  RPC_URL="${RPC_URL}" L1_RUN="${RUN_DIR}" SECRETS="${SECRETS}" MORPH="${MORPH}" "${ROOT}/scripts/register-challenger.sh"
}

s40_wait_finalized() {
  if [[ "${WAIT_FINALIZED}" != "1" ]]; then echo "WAIT_FINALIZED=${WAIT_FINALIZED}: skipped"; return 0; fi
  local target end fin
  target="$(cast block-number --rpc-url "${RPC_URL}")"
  end=$(( $(date +%s) + FINALIZED_TIMEOUT_SEC ))
  while true; do
    fin="$(cast block finalized --field number --rpc-url "${RPC_URL}" 2>/dev/null || echo 0)"
    echo "finalized=${fin} target=${target}"
    [[ "${fin}" -ge "${target}" ]] && return 0
    [[ "$(date +%s)" -lt "${end}" ]] || { echo "timeout after ${FINALIZED_TIMEOUT_SEC}s"; return 1; }
    sleep 20
  done
}

s90_summary() {
  local b0 b1 s0 s1 n0 n1 spent
  b0="$(cat "${STATE}/deployer-balance-before.wei")"; b1="$(cast balance "${DEPLOYER}" --rpc-url "${RPC_URL}")"
  s0="$(cat "${STATE}/submitter-balance-before.wei")"; s1="$(cast balance "${SUBMITTER}" --rpc-url "${RPC_URL}")"
  n0="$(cat "${STATE}/deployer-nonce-before")"; n1="$(cast nonce "${DEPLOYER}" --rpc-url "${RPC_URL}")"
  spent="$(node -e 'console.log((BigInt(process.argv[1]) - BigInt(process.argv[2])).toString())' "${b0}" "${b1}")"
  {
    echo "deployer  ${DEPLOYER}"
    echo "  balance before: $(wei_to_eth "${b0}") ETH"
    echo "  balance after : $(wei_to_eth "${b1}") ETH"
    echo "  spent         : $(wei_to_eth "${spent}") ETH  (${spent} wei)"
    echo "  transactions  : $((n1 - n0))"
    echo "submitter ${SUBMITTER}"
    echo "  balance before: $(wei_to_eth "${s0}") ETH   after: $(wei_to_eth "${s1}") ETH"
    if [[ -n "${CHALLENGER_PRIVATE_KEY:-}" ]]; then
      CHA="$(addr_of CHALLENGER_PRIVATE_KEY)"
      echo "challenger ${CHA}  (Rollup.isChallenger; balance $(wei_to_eth "$(cast balance "${CHA}" --rpc-url "${RPC_URL}")") ETH)"
    fi
    echo "L1 start block  : $(cat "${STATE}/start-block.dec") ($(cat "${STATE}/start-block.hex"))"
  } | tee "${STATE}/summary.txt"
}

say "RPC ${RPC_URL}   RUN_DIR ${RUN_DIR}"
step s00_preflight s00_preflight
step s10_deploy s10_deploy
step s11_verify_deploy s11_verify_deploy
step s20_genesis s20_genesis
step s21_batchheader s21_batchheader
step s30_initialize s30_initialize
step s31_register s31_register
step s32_first_sequencer s32_first_sequencer
step s33_fund_submitter s33_fund_submitter
step s34_register_challenger s34_register_challenger
step s40_wait_finalized s40_wait_finalized
s90_summary
if [[ -n "${PROXY_PID:-}" ]]; then say "proxy stats: $(curl -s "http://127.0.0.1:${PROXY_PORT}/stats")"; fi
