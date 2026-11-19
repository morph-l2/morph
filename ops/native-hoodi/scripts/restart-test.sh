#!/usr/bin/env bash
#
# restart-test.sh — does the Morph L2 stack recover from processes being killed or restarted?
#
# Scenarios T1..T8 (see below). Before each one the stack must be healthy (every process up, both L2 heights moving); a
# scenario passes only if its stated condition holds after the action. Nothing is assumed from exit codes: heights and
# block hashes are read from the nodes.
#
# Usage: restart-test.sh [T1 T2 ...]      (default: all except T7, which waits for the next batch: add T7 explicitly)
set -uo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
CTL="${ROOT}/scripts/morph-native.sh"
L2_RUN="${L2_RUN:-${ROOT}/run/l2}"
L1_URL="${HOODI_RPC:-https://rpc.hoodi.ethpandaops.io}"
A=http://127.0.0.1:8545; B=http://127.0.0.1:8645
# shellcheck disable=SC1090
source "${ROOT}/scripts/lib/env.sh"
ALL_PROCS=(l1-proxy morph-el-0 node-0 morph-el-1 node-1 morph-el-2 node-2 tx-submitter gas-price-oracle)

say() { printf '%s %s\n' "$(date +%H:%M:%S)" "$*"; }
ha() { cast block-number --rpc-url "${A}" 2>/dev/null || echo 0; }
hb() { cast block-number --rpc-url "${B}" 2>/dev/null || echo 0; }
hash_at() { cast block "$2" --field hash --rpc-url "$1" 2>/dev/null; }
alive() { [[ -f "${L2_RUN}/pids/$1.pid" ]] && kill -0 "$(cat "${L2_RUN}/pids/$1.pid")" 2>/dev/null; }
wait_until() { # wait_until SECONDS DESCRIPTION CMD...
  local end=$(( $(date +%s) + $1 )) what="$2"; shift 2
  until "$@"; do [[ "$(date +%s)" -lt "${end}" ]] || { say "   timed out: ${what}"; return 1; }; sleep 3; done
}
PASS=0; FAIL=0
verdict() { # verdict NAME RC [DETAIL]
  if [[ "$2" -eq 0 ]]; then PASS=$((PASS + 1)); say "PASS  $1  ${3:-}"; else FAIL=$((FAIL + 1)); say "FAIL  $1  ${3:-}"; fi
}

all_up() { local p; for p in "${ALL_PROCS[@]}"; do alive "${p}" || return 1; done; }
a_moves() { local h; h="$(ha)"; sleep 10; [[ "$(ha)" -gt "${h}" ]]; }
b_moves() { local h; h="$(hb)"; sleep 30; [[ "$(hb)" -gt "${h}" ]]; }   # node-1 advances in bursts, one batch at a time
tip_hashes_agree() { # hashes at node-1's tip-1 equal on both nodes
  local t; t="$(( $(hb) - 1 ))"; [[ "$(hash_at "${A}" "${t}")" == "$(hash_at "${B}" "${t}")" && -n "$(hash_at "${A}" "${t}")" ]]
}
healthy() { all_up && a_moves; }
preflight() { say "-- preflight: all processes up and node-0 producing"; healthy || { say "   stack is not healthy before ${1}; aborting"; exit 2; }; }

# ---------------------------------------------------------------- scenarios
T1() { # crash node-1 with kill -9 (the supervisor should restart it by itself)
  preflight T1
  local sup child h0; sup="$(cat "${L2_RUN}/pids/node-1.pid")"; child="$(pgrep -P "${sup}" | head -1)"; h0="$(hb)"
  say "T1: kill -9 node-1 (pid ${child}, supervisor ${sup}); node-1 height before: ${h0}"
  kill -9 "${child}"; sleep 5
  local newchild; newchild="$(pgrep -P "${sup}" | head -1)"
  [[ -n "${newchild}" && "${newchild}" != "${child}" ]]; verdict "T1a supervisor restarted node-1 on its own" $? "(new pid ${newchild:-none})"
  wait_until 900 "node-1 moves past ${h0}" bash -c "[[ \$(cast block-number --rpc-url ${B}) -gt ${h0} ]]"; verdict "T1b node-1 keeps deriving after the crash" $? "(height $(hb))"
  tip_hashes_agree; verdict "T1c node-1 tip hash == node-0" $?
}
T2() { # stop/start morph-el-1
  preflight T2
  local h0; h0="$(hb)"; say "T2: stop morph-el-1 then start it; node-1 height before: ${h0}"
  "${CTL}" kill morph-el-1 >/dev/null; sleep 5; "${CTL}" start morph-el-1 >/dev/null
  wait_until 900 "node-1 moves past ${h0}" bash -c "[[ \$(cast block-number --rpc-url ${B}) -gt ${h0} ]]"; verdict "T2a node-1 resumes after its execution client restarted" $? "(height $(hb))"
  tip_hashes_agree; verdict "T2b node-1 tip hash == node-0" $?
}
T3() { # stop/start node-1
  preflight T3
  local h0; h0="$(hb)"; say "T3: stop node-1 then start it; node-1 height before: ${h0}"
  "${CTL}" kill node-1 >/dev/null; sleep 5; "${CTL}" start node-1 >/dev/null
  wait_until 900 "node-1 moves past ${h0}" bash -c "[[ \$(cast block-number --rpc-url ${B}) -gt ${h0} ]]"; verdict "T3a node-1 resumes after a clean restart" $? "(height $(hb))"
  tip_hashes_agree; verdict "T3b node-1 tip hash == node-0" $?
}
T4() { # stop/start the L1 proxy
  preflight T4
  local l0; l0="$(curl -s -m 5 localhost:26660/metrics | awk '/^morphnode_syncer_l1height /{printf "%d", $2}')"
  say "T4: stop l1-proxy for 20 s then start it; node-0 L1 syncer height before: ${l0}"
  "${CTL}" kill l1-proxy >/dev/null; sleep 20; "${CTL}" start l1-proxy >/dev/null
  wait_until 180 "node-0 syncer advances past ${l0}" bash -c "[[ \$(curl -s -m 5 localhost:26660/metrics | awk '/^morphnode_syncer_l1height /{printf \"%d\", \$2}') -gt ${l0} ]]"
  verdict "T4a node-0 L1 syncer recovers after the proxy was gone" $? "(now $(curl -s -m 5 localhost:26660/metrics | awk '/^morphnode_syncer_l1height /{printf "%d", $2}'))"
  a_moves; verdict "T4b node-0 never stopped producing" $?
  all_up; verdict "T4c every process is still running (nothing exited because of the outage)" $?
}
T5() { # stop/start node-0, the block producer
  preflight T5
  local h0 hash0 mid; h0="$(ha)"; mid="$(( h0 - 5 ))"; hash0="$(hash_at "${A}" "${mid}")"
  say "T5: stop node-0 for 15 s then start it; node-0 height before: ${h0}; remembering the hash of block ${mid}"
  "${CTL}" kill node-0 >/dev/null; sleep 15; "${CTL}" start node-0 >/dev/null
  wait_until 180 "node-0 produces past ${h0}" bash -c "[[ \$(cast block-number --rpc-url ${A}) -gt ${h0} ]]"; verdict "T5a node-0 resumes producing after a restart" $? "(height $(ha))"
  [[ "$(hash_at "${A}" "${mid}")" == "${hash0}" ]]; verdict "T5b block ${mid} unchanged after the restart (no reorg)" $?
  a_moves; verdict "T5c node-0 keeps producing" $?
}
T6() { # stop/start the execution client of the producer
  preflight T6
  local h0; h0="$(ha)"; say "T6: stop morph-el-0 for 15 s then start it; node-0 height before: ${h0}"
  "${CTL}" kill morph-el-0 >/dev/null; sleep 15; "${CTL}" start morph-el-0 >/dev/null
  wait_until 240 "node-0 produces past ${h0}" bash -c "[[ \$(cast block-number --rpc-url ${A}) -gt ${h0} ]]"; verdict "T6a node-0 produces again after its execution client came back" $? "(height $(ha))"
  alive node-0; verdict "T6b node-0 process survived" $?
}
T7() { # stop/start tx-submitter, then wait for the next batch
  preflight T7
  local R b0; R="$(node -e 'console.log(JSON.parse(require("fs").readFileSync(process.argv[1],"utf8")).find(x=>x.name==="Proxy__Rollup").address)' "${ROOT}/run/l1/deployment/hoodidev.json")"
  b0="$(cast call "${R}" 'lastCommittedBatchIndex()(uint256)' --rpc-url ${L1_URL} | awk '{print $1}')"
  say "T7: stop tx-submitter for 10 s then start it; committed batches before: ${b0}; waiting up to 15 min for the next one"
  "${CTL}" kill tx-submitter >/dev/null; sleep 10; "${CTL}" start tx-submitter >/dev/null
  wait_until 900 "a new batch is committed" bash -c "[[ \$(cast call ${R} 'lastCommittedBatchIndex()(uint256)' --rpc-url ${L1_URL} | awk '{print \$1}') -gt ${b0} ]]"
  verdict "T7a submitter commits the next batch after a restart" $? "(now $(cast call "${R}" 'lastCommittedBatchIndex()(uint256)' --rpc-url ${L1_URL} | awk '{print $1}'))"
}
T8() { # everything down, everything up
  preflight T8
  local h0 mid hash0; h0="$(ha)"; mid="$(( h0 - 5 ))"; hash0="$(hash_at "${A}" "${mid}")"
  say "T8: down ALL, wait 15 s, up ALL; node-0 height before: ${h0}; remembering the hash of block ${mid}"
  "${CTL}" down >/dev/null 2>&1; sleep 15; "${CTL}" up >/dev/null 2>&1
  wait_until 300 "node-0 serves RPC again" bash -c "cast block-number --rpc-url ${A} >/dev/null 2>&1"
  wait_until 300 "node-0 produces past ${h0}" bash -c "[[ \$(cast block-number --rpc-url ${A}) -gt ${h0} ]]"; verdict "T8a node-0 continues from its old height, not from 0" $? "(before ${h0}, now $(ha))"
  [[ "$(hash_at "${A}" "${mid}")" == "${hash0}" ]]; verdict "T8b history unchanged after a full stop/start" $?
  all_up; verdict "T8c all processes are running" $?
  wait_until 900 "node-1 catches up" bash -c "[[ \$(cast block-number --rpc-url ${B}) -gt 0 ]]"; tip_hashes_agree; verdict "T8d node-1 tip hash == node-0" $?
}

SEL=("$@"); [[ ${#SEL[@]} -gt 0 ]] || SEL=(T1 T2 T3 T4 T5 T6 T8)
say "== restart test: ${SEL[*]}"
for s in "${SEL[@]}"; do "${s}"; done
say "== result: ${PASS} passed, ${FAIL} failed"
[[ "${FAIL}" -eq 0 ]]
