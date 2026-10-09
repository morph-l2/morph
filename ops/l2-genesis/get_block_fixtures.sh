#!/bin/sh
# Fetch a block, its prestate trace and the preceding 256 block hashes.
set -eu
umask 077
fail() { printf '%s\n' "$*" >&2; exit 1; }
usage() { printf '%s\n' 'Usage: get_block_fixtures.sh RPC_URL DECIMAL_BLOCK_NUMBER [OUTPUT_DIRECTORY]'; }
if [ "${1:-}" = --help ] || [ "${1:-}" = -h ]; then usage; exit 0; fi
[ "$#" -ge 2 ] && [ "$#" -le 3 ] || { usage >&2; exit 1; }
url=$1
case "$url" in http://?*|https://?*) ;; *) fail 'RPC_URL must be an explicit HTTP(S) endpoint' ;; esac
case "$2" in ''|*[!0-9]*) fail 'BLOCK_NUMBER must be a nonnegative decimal integer' ;; esac
for command in curl jq; do command -v "$command" >/dev/null || fail "Required command is unavailable: $command"; done
block_number=$(jq -nr --arg value "$2" '$value | tonumber | if . <= 9007199254740991 then . else error("BLOCK_NUMBER exceeds the exact integer range") end')
output_dir=${3:-.}
mkdir -p "$output_dir"
output_dir=$(CDPATH= cd -- "$output_dir" && pwd)
lock=$output_dir/.block-fixtures.lock.d
mkdir "$lock" 2>/dev/null || fail "Fixture download is already locked: $lock"
attempt=
cleanup() {
  status=$?
  trap - EXIT HUP INT TERM
  rmdir "$lock"
  if [ "$status" -ne 0 ] && [ -n "$attempt" ]; then printf 'Download failed; incomplete fixtures retained at %s\n' "$attempt" >&2; fi
  exit "$status"
}
trap cleanup EXIT
trap 'exit 1' HUP INT TERM
for file in block.json prestate.json block_hashes.json; do
  [ ! -e "$output_dir/$file" ] || fail "Existing fixture is preserved: $output_dir/$file; choose another output directory"
done
attempt=$(mktemp -d "$output_dir/fixtures-attempt-XXXXXX")
rpc() {
  request=$(jq -cn --arg method "$1" --argjson params "$2" '{jsonrpc:"2.0",id:1,method:$method,params:$params}')
  response=$(curl --fail --silent --max-time 60 -H 'Content-Type: application/json' --data "$request" "$url") || fail "RPC request failed: $1"
  printf '%s' "$response" | jq -e 'if type == "object" and has("result") and .result != null and (has("error") | not) then .result else error("RPC returned an error or no result") end'
}
block_hex=$(printf '0x%x' "$block_number")
rpc eth_getBlockByNumber "[\"$block_hex\",true]" | jq -e --arg number "$block_hex" 'if type == "object" and .number == $number and (.transactions | type) == "array" and (.hash | type) == "string" and (.hash | test("^0x[0-9a-fA-F]{64}$")) then . else error("Invalid block response") end' > "$attempt/block.json"
rpc debug_traceBlockByNumber "[\"$block_hex\",{\"tracer\":\"prestateTracer\"}]" | jq -e --slurpfile block "$attempt/block.json" 'if type == "array" and length == ($block[0].transactions | length) and all(.[]; type == "object" and (has("error") | not) and (.result | type) == "object") then . else error("Invalid prestate trace") end' > "$attempt/prestate.json"
: > "$attempt/hashes.jsonl"
number=$((block_number - 256))
while [ "$number" -lt "$block_number" ]; do
  if [ "$number" -lt 0 ]; then
    printf '"0x%064d"\n' 0 >> "$attempt/hashes.jsonl"
  else
    block_hex=$(printf '0x%x' "$number")
    rpc eth_getBlockByNumber "[\"$block_hex\",false]" | jq -e --arg number "$block_hex" 'if .number == $number and (.hash | type) == "string" and (.hash | test("^0x[0-9a-fA-F]{64}$")) then .hash else error("Invalid block hash response") end' >> "$attempt/hashes.jsonl"
  fi
  number=$((number + 1))
done
jq -s 'if length == 256 then . else error("Expected 256 preceding block hashes") end' "$attempt/hashes.jsonl" > "$attempt/block_hashes.json"
for file in block.json prestate.json block_hashes.json; do mv "$attempt/$file" "$output_dir/$file"; done
rm "$attempt/hashes.jsonl"
rmdir "$attempt"
attempt=
printf 'Block fixtures saved: %s\n' "$output_dir"
