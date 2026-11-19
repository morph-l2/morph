#!/usr/bin/env bash
#
# reset-env.sh — 把本机环境重置到"干净状态",并且可以原样恢复。从不删除任何东西,只 mv。
#
# 用法:
#   reset-env.sh plan    [--with-build] [--new-keys]      只打印会做什么(默认命令,不改任何东西)
#   reset-env.sh run     [--with-build] [--new-keys] --yes   真正执行
#   reset-env.sh restore <时间戳> --yes                    把某次 reset 挪走的东西原样挪回
#   reset-env.sh list                                      列出已有的归档
#
# run 做的事(按顺序):
#   1. 检查 submitter 在 L1 上的余额;高于 SWEEP_THRESHOLD_ETH(默认 0.02)就拒绝,先回收资金再来(--ignore-balance 可跳过)
#   2. morph-native.sh down(停全部进程)
#   3. 把 run/ 下的运行数据(l1、l2、失败尝试、测试输出、日志)挪到 run/archive/<时间戳>/
#   4. --with-build:把二进制和构建产物也挪进去(bin/、go-ethereum/build/bin、node/build/bin、tx-submitter 二进制、
#      gas-oracle/app/target、contracts 的 node_modules/artifacts/dist/typechain-types)。这些在 git 里都是被忽略的
#   5. --new-keys:旧 env/secrets.env 挪到 env/archive/secrets-<时间戳>.env,新建一份:deployer 和 RPC 配置原样保留,
#      sequencer / batch submitter / gas oracle 三个密钥重新生成(私钥不会被打印)
#   6. 写 run/archive/<时间戳>/MANIFEST.tsv(原路径<TAB>现路径),restore 就靠它
#
# 不碰:仓库里被 git 跟踪的文件、deployer 密钥、notes/ 与 docs、deployments/(部署记录)。
# 公共 Hoodi 上已部署的合约不受影响(链上的东西没法"清空")。
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
MORPH="${MORPH:-$(cd "${ROOT}/../.." && pwd)}"
CTL="${ROOT}/scripts/morph-native.sh"
SECRETS="${SECRETS:-${ROOT}/env/secrets.env}"
ARCH="${ROOT}/run/archive"
THRESH="${SWEEP_THRESHOLD_ETH:-0.02}"
# shellcheck disable=SC1090
source "${ROOT}/scripts/lib/env.sh"

die() { echo "error: $*" >&2; exit 1; }
say() { printf '%s %s\n' "$(date +%H:%M:%S)" "$*"; }

CMD="${1:-plan}"; [[ $# -gt 0 ]] && shift
WITH_BUILD=0; NEW_KEYS=0; YES=0; IGNORE_BAL=0; TS_ARG=""
for a in "$@"; do
  case "$a" in
    --with-build) WITH_BUILD=1 ;; --new-keys) NEW_KEYS=1 ;; --yes) YES=1 ;; --ignore-balance) IGNORE_BAL=1 ;;
    -*) die "unknown option $a" ;; *) TS_ARG="$a" ;;
  esac
done

data_paths() { # run data: everything under run/ except the archive itself
  local p
  for p in "${ROOT}"/run/*; do [[ "$p" == "${ARCH}" ]] && continue; [[ -e "$p" ]] && echo "$p"; done
}
build_paths() {
  local p
  for p in "${ROOT}/bin" "${MORPH}/go-ethereum/build/bin" "${MORPH}/node/build/bin" "${MORPH}/tx-submitter/tx-submitter" \
           "${MORPH}/gas-oracle/app/target" "${MORPH}/contracts/node_modules" "${MORPH}/contracts/artifacts" \
           "${MORPH}/contracts/dist" "${MORPH}/contracts/typechain-types"; do
    [[ -e "$p" ]] && echo "$p"
  done
}

alive_any() { local f; for f in "${ROOT}"/run/l2/pids/*.pid; do [[ -f "$f" ]] && kill -0 "$(cat "$f")" 2>/dev/null && return 0; done; return 1; }

# value of a variable in secrets.env without printing anything else
secret() { ( set -a; source "${SECRETS}"; set +a; printf '%s' "${!1:-}" ); }

submitter_balance_eth() {
  local key rpc addr
  key="$(secret BATCH_SUBMITTER_PRIVATE_KEY)"; rpc="${HOODI_RPC:-https://rpc.hoodi.ethpandaops.io}"
  [[ -n "${key}" ]] || { echo "?"; return; }
  addr="$(cast wallet address --private-key "${key}" 2>/dev/null)" || { echo "?"; return; }
  cast balance "${addr}" --ether --rpc-url "${rpc}" 2>/dev/null || echo "?"
}

gen_secrets() { # $1 = new file; copies the old one, replacing the three generated roles
  python3 -I - "${SECRETS}" "$1" <<'PY'
import json, re, subprocess, sys
src, dst = sys.argv[1:3]
roles = ["SEQUENCER_PRIVATE_KEY", "BATCH_SUBMITTER_PRIVATE_KEY", "L2_GAS_ORACLE_PRIVATE_KEY"]
def new_key():
    out = json.loads(subprocess.check_output(["cast", "wallet", "new", "--json"]))
    if isinstance(out, dict) and "data" in out: out = out["data"]
    if isinstance(out, list): out = out[0]
    return out["private_key"]
text = open(src).read()
for r in roles:
    pat = re.compile(r'^(\s*(?:export\s+)?' + r + r'\s*=\s*).*$', re.M)
    if not pat.search(text): sys.exit("role not found in secrets: " + r)
    k = new_key()
    text = pat.sub(lambda m: m.group(1) + k, text, count=1)
open(dst, "w").write(text)
PY
  chmod 600 "$1"
}

do_plan_or_run() {
  local ts items=() p
  ts="$(date +%Y%m%d-%H%M%S)"
  mapfile -t items < <(data_paths)
  [[ "${WITH_BUILD}" -eq 1 ]] && mapfile -t -O "${#items[@]}" items < <(build_paths)

  echo "== reset-env ${CMD} (timestamp ${ts})"
  echo "-- submitter balance on L1: $(submitter_balance_eth) ETH (threshold ${THRESH})"
  echo "-- processes: $(alive_any && echo 'RUNNING -> will be stopped with morph-native.sh down' || echo 'none running')"
  echo "-- will move to ${ARCH}/${ts}/ :"
  for p in "${items[@]}"; do printf '     %s  (%s)\n' "${p#"${ROOT}"/}" "$(du -sh "$p" 2>/dev/null | cut -f1)"; done
  [[ "${WITH_BUILD}" -eq 1 ]] || echo "     (build artifacts NOT included: add --with-build)"
  if [[ "${NEW_KEYS}" -eq 1 ]]; then
    echo "-- keys: env/secrets.env -> env/archive/secrets-${ts}.env ; new secrets.env keeps deployer + RPC settings, regenerates sequencer / submitter / gas-oracle keys"
  else
    echo "-- keys: unchanged (add --new-keys to regenerate sequencer / submitter / oracle)"
  fi
  [[ "${CMD}" == plan ]] && { echo "== plan only; nothing was changed. Re-run as: reset-env.sh run ... --yes"; return 0; }

  [[ "${YES}" -eq 1 ]] || die "refusing to run without --yes"
  if [[ "${IGNORE_BAL}" -ne 1 ]]; then
    local bal; bal="$(submitter_balance_eth)"
    [[ "${bal}" != "?" ]] || die "could not read the submitter balance; fix the RPC or pass --ignore-balance"
    awk -v b="${bal}" -v t="${THRESH}" 'BEGIN{exit !(b+0 > t+0)}' && die "submitter still holds ${bal} ETH (> ${THRESH}): sweep it to the deployer first, or pass --ignore-balance"
  fi

  alive_any && "${CTL}" down
  alive_any && die "processes still alive after down"
  local dest="${ARCH}/${ts}" man
  mkdir -p "${dest}"; man="${dest}/MANIFEST.tsv"; : >"${man}"
  local n=0
  for p in "${items[@]}"; do
    n=$((n + 1)); mkdir -p "${dest}/${n}"
    mv "$p" "${dest}/${n}/" && printf '%s\t%s\n' "$p" "${dest}/${n}/$(basename "$p")" >>"${man}"
    say "moved ${p#"${ROOT}"/}"
  done
  if [[ "${NEW_KEYS}" -eq 1 ]]; then
    mkdir -p "${ROOT}/env/archive"
    local old="${ROOT}/env/archive/secrets-${ts}.env"
    cp -p "${SECRETS}" "${old}.tmp" && gen_secrets "${SECRETS}.new" && mv "${old}.tmp" "${old}" && mv "${SECRETS}.new" "${SECRETS}"
    printf '%s\t%s\n' "${SECRETS}" "${old}" >>"${man}.keys"
    say "secrets regenerated (old file kept at env/archive/secrets-${ts}.env)"
  fi
  mkdir -p "${ROOT}/run"
  echo "== done. Restore with: reset-env.sh restore ${ts} --yes"
}

do_restore() {
  [[ -n "${TS_ARG}" ]] || die "usage: restore <timestamp> --yes"
  local man="${ARCH}/${TS_ARG}/MANIFEST.tsv" orig cur
  [[ -f "${man}" ]] || die "no manifest ${man}"
  echo "== restore ${TS_ARG}"
  while IFS=$'\t' read -r orig cur; do
    if [[ -e "${orig}" ]]; then echo "  CONFLICT (exists now): ${orig}"; CONFLICT=1; else echo "  ${cur#"${ROOT}"/} -> ${orig#"${ROOT}"/}"; fi
  done <"${man}"
  [[ "${CONFLICT:-0}" -eq 0 ]] || die "some targets exist; reset again or move them away first"
  [[ "${YES}" -eq 1 ]] || { echo "== preview only; add --yes"; return 0; }
  alive_any && die "stop the stack first (morph-native.sh down)"
  while IFS=$'\t' read -r orig cur; do mkdir -p "$(dirname "${orig}")"; mv "${cur}" "${orig}"; done <"${man}"
  if [[ -f "${man}.keys" ]]; then
    while IFS=$'\t' read -r orig cur; do
      echo "  note: the keys of that time are in ${cur#"${ROOT}"/}; current secrets.env was NOT overwritten (copy it yourself if wanted)"
    done <"${man}.keys"
  fi
  echo "== restored"
}

case "${CMD}" in
  plan|run) do_plan_or_run ;;
  restore)  do_restore ;;
  list)     ls -1 "${ARCH}" 2>/dev/null || echo "(no archives)" ;;
  *)        sed -n '2,25p' "${BASH_SOURCE[0]}" | sed 's/^# \{0,1\}//'; exit 1 ;;
esac
