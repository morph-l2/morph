#!/usr/bin/env bash
# probe-hoodi.sh — 对 Hoodi 的公共 EL / 信标端点做温和的限流与能力探测(只读,不发交易)。
#
# 用法:probe-hoodi.sh [并发数,默认 20] [轮数,默认 3]
# 输出:每个端点的 HTTP 状态分布、平均延迟,以及几项 Morph 实际会用到的接口。
# 注意:对公共服务只做小规模并发,不是压测。

set -uo pipefail
CONC="${1:-20}"
ROUNDS="${2:-3}"

EL_ENDPOINTS=("https://0xrpc.io/hoodi" "https://rpc.hoodi.ethpandaops.io")
BEACON_ENDPOINTS=("https://beacon.hoodi.ethpandaops.io" "https://ethereum-hoodi-beacon-api.publicnode.com")

tmp="$(mktemp -d)"
trap 'rm -rf "${tmp}"' EXIT

one_el() { # one_el URL METHOD PARAMS -> "HTTP_CODE TIME"
  curl -s -o /dev/null -w '%{http_code} %{time_total}\n' -m 20 -X POST "$1" \
    -H 'content-type: application/json' \
    --data "{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"$2\",\"params\":$3}"
}
export -f one_el

summarize() { # summarize FILE -> "codes: 200x20 429x0 ...; avg=..s max=..s"
  awk '{c[$1]++; t+=$2; if($2>m)m=$2; n++} END{
    s=""; for(k in c) s=s k "x" c[k] " "; printf "%s avg=%.2fs max=%.2fs n=%d\n", s, (n?t/n:0), m, n}' "$1"
}

echo "== 并发 ${CONC} × ${ROUNDS} 轮"
for u in "${EL_ENDPOINTS[@]}"; do
  : >"${tmp}/el.txt"
  for _ in $(seq 1 "${ROUNDS}"); do
    seq 1 "${CONC}" | xargs -P "${CONC}" -I{} bash -c "one_el '${u}' eth_blockNumber '[]'" >>"${tmp}/el.txt"
    sleep 1
  done
  printf 'EL      %-44s ' "${u}"; summarize "${tmp}/el.txt"
done

for u in "${BEACON_ENDPOINTS[@]}"; do
  : >"${tmp}/bn.txt"
  for _ in $(seq 1 "${ROUNDS}"); do
    seq 1 "${CONC}" | xargs -P "${CONC}" -I{} curl -s -o /dev/null -w '%{http_code} %{time_total}\n' -m 20 "${u}/eth/v1/beacon/headers/head" >>"${tmp}/bn.txt"
    sleep 1
  done
  printf 'BEACON  %-44s ' "${u}"; summarize "${tmp}/bn.txt"
done

echo
echo "== eth_getLogs:区间上限(Morph 的 L1 同步会用)"
for u in "${EL_ENDPOINTS[@]}"; do
  head_hex="$(curl -s -m 15 -X POST "${u}" -H 'content-type: application/json' --data '{"jsonrpc":"2.0","id":1,"method":"eth_blockNumber","params":[]}' | sed -E 's/.*"result":"(0x[0-9a-f]+)".*/\1/')"
  head=$((head_hex))
  for span in 1000 5000 10000 50000; do
    from=$(printf '0x%x' $((head - span)))
    out="$(curl -s -m 30 -X POST "${u}" -H 'content-type: application/json' \
      --data "{\"jsonrpc\":\"2.0\",\"id\":1,\"method\":\"eth_getLogs\",\"params\":[{\"fromBlock\":\"${from}\",\"toBlock\":\"latest\",\"address\":\"0x00000000219ab540356cBB839Cbe05303d7705Fa\"}]}")"
    if [[ "${out}" == *'"error"'* ]]; then msg="$(printf '%s' "${out}" | sed -E 's/.*"message":"([^"]{0,80}).*/\1/')"; else msg="ok"; fi
    printf '  %-40s span=%-6s %s\n' "${u}" "${span}" "${msg}"
  done
done

echo
echo "== 信标:近期 slot 的 blob_sidecars 连续取 5 个"
for u in "${BEACON_ENDPOINTS[@]}"; do
  head_slot="$(curl -s -m 15 "${u}/eth/v1/beacon/headers/head" | sed -E 's/.*"slot":"([0-9]+)".*/\1/')"
  ok=0; fail=0
  for i in 1 2 3 4 5; do
    code="$(curl -s -o /dev/null -w '%{http_code}' -m 20 "${u}/eth/v1/beacon/blob_sidecars/$((head_slot - 8 * i))")"
    [[ "${code}" == 200 ]] && ok=$((ok + 1)) || fail=$((fail + 1))
  done
  printf '  %-52s head_slot=%s ok=%d fail=%d\n' "${u}" "${head_slot}" "${ok}" "${fail}"
done
