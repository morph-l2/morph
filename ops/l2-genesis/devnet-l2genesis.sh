#!/bin/sh
# Generate devnet genesis; qanet-l2genesis.sh selects the same flow for qanet.
set -eu
umask 077
SCRIPT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
network=devnet
output_dir= deployment_file= deploy_config= l1_rpc=
verify_existing=false overwrite=false
fail() { printf '%s\n' "$*" >&2; exit 1; }
usage() {
  cat <<'HELP'
Usage: devnet-l2genesis.sh [--network devnet|qanet] [options]
  --deploy-config FILE     L2 genesis source configuration
  --deployment-file FILE   Confirmed L1 deployment records
  --output-dir DIRECTORY  Generated artifacts and genesis.done
  --l1-rpc URL            Explicit L1 endpoint (or L1_RPC_URL / QA_RPC_URL)
  --verify-existing       Verify existing artifacts without running Go
  --overwrite             Explicitly regenerate existing artifacts
HELP
}
while [ "$#" -gt 0 ]; do
  case "$1" in
    --help|-h) usage; exit 0 ;;
    --verify-existing) verify_existing=true; shift ;;
    --overwrite) overwrite=true; shift ;;
    --network|--deploy-config|--deployment-file|--output-dir|--l1-rpc)
      [ "$#" -ge 2 ] || fail "Missing value for $1"
      case "$1" in
        --network) network=$2 ;; --deploy-config) deploy_config=$2 ;;
        --deployment-file) deployment_file=$2 ;; --output-dir) output_dir=$2 ;;
        --l1-rpc) l1_rpc=$2 ;;
      esac
      shift 2 ;;
    *) fail "Unknown option: $1" ;;
  esac
done
case "$network" in
  devnet) l1_rpc=${l1_rpc:-${L1_RPC_URL:-}} ;;
  qanet) l1_rpc=${l1_rpc:-${QA_RPC_URL:-}} ;;
  *) fail 'Network must be devnet or qanet' ;;
esac
[ "$verify_existing:$overwrite" != true:true ] || fail '--verify-existing and --overwrite are mutually exclusive'
case "$l1_rpc" in http://?*|https://?*) ;; *) fail 'Provide an explicit HTTP(S) L1 RPC with --l1-rpc, L1_RPC_URL or QA_RPC_URL' ;; esac
for command in jq curl shasum; do command -v "$command" >/dev/null || fail "Required command is unavailable: $command"; done
output_dir=${output_dir:-$SCRIPT_DIR/.$network}
deploy_config=${deploy_config:-$SCRIPT_DIR/deploy-config/$network-deploy-config.json}
deployment_file=${deployment_file:-$SCRIPT_DIR/.$network/${network}L1.json}
[ -f "$deploy_config" ] || fail "Missing deploy config: $deploy_config"
[ -f "$deployment_file" ] || fail "Missing deployment records: $deployment_file"
deploy_config=$(CDPATH= cd -- "$(dirname -- "$deploy_config")" && printf '%s/%s' "$PWD" "$(basename -- "$deploy_config")")
deployment_file=$(CDPATH= cd -- "$(dirname -- "$deployment_file")" && printf '%s/%s' "$PWD" "$(basename -- "$deployment_file")")
if "$verify_existing"; then [ -d "$output_dir" ] || fail "Missing artifact directory: $output_dir"; fi
mkdir -p "$output_dir"
output_dir=$(CDPATH= cd -- "$output_dir" && pwd)
# mkdir provides a portable exclusive lock. A terminated process may leave it behind.
lock=$output_dir/.genesis.lock.d
mkdir "$lock" 2>/dev/null || fail "Genesis generation or verification is already locked: $lock; confirm no process is running before removing a stale lock"
attempt=
cleanup() {
  status=$?
  trap - EXIT HUP INT TERM
  rmdir "$lock"
  if [ "$status" -ne 0 ] && [ -n "$attempt" ]; then printf 'Genesis generation did not complete; diagnostics: %s\n' "$attempt" >&2; fi
  exit "$status"
}
trap cleanup EXIT
trap 'exit 1' HUP INT TERM
files='deploy-config.json genesis-l2.json rollup.json genesis-batch-header.json deployment-config.json'
if ! "$verify_existing"; then
  if ! "$overwrite"; then
    for file in genesis.done genesis-l2.json rollup.json genesis-batch-header.json; do
      [ ! -e "$output_dir/$file" ] || fail 'Existing genesis artifacts are preserved; use --verify-existing or explicitly --overwrite'
    done
  fi
fi
attempt=$(mktemp -d "$output_dir/genesis-attempt-XXXXXX")
# Genesis still uses the retired staking address; Submitter is a different role.
mapping='{"l1CrossDomainMessengerProxy":"Proxy__L1CrossDomainMessenger","RollupProxy":"Proxy__Rollup","l1GatewayRouterProxy":"Proxy__L1GatewayRouter","l1StandardERC20GatewayProxy":"Proxy__L1StandardERC20Gateway","l1CustomERC20GatewayProxy":"Proxy__L1CustomERC20Gateway","l1ReverseCustomGatewayProxy":"Proxy__L1ReverseCustomGateway","l1ETHGatewayProxy":"Proxy__L1ETHGateway","l1ERC721GatewayProxy":"Proxy__L1ERC721Gateway","l1ERC1155GatewayProxy":"Proxy__L1ERC1155Gateway","l1WETHGatewayProxy":"Proxy__L1WETHGateway","l1WETH":"Impl__WETH","l1WithdrawLockERC20Gateway":"Proxy__L1WithdrawLockERC20Gateway"}'
jq_common='def uint: if type == "number" and . >= 0 and floor == . then . elif type == "string" and test("^0[xX][0-9a-fA-F]+$") then ascii_downcase | .[2:] | explode | reduce .[] as $c (0; . * 16 + (if $c >= 97 then $c - 87 else $c - 48 end)) elif type == "string" and test("^[0-9]+$") then tonumber else error("Expected an unsigned integer") end;
def address: type == "string" and test("^0x[0-9a-fA-F]{40}$") and . != "0x0000000000000000000000000000000000000000";
def hash: type == "string" and test("^0x[0-9a-fA-F]{64}$");'
jq -e --argjson mapping "$mapping" --slurpfile records "$deployment_file" "$jq_common"'
  if type != "object" or (.l1ChainID | uint) == 0 or (.l2ChainID | uint) == 0 then error("Invalid genesis configuration") else . end |
  $records[0] as $r |
  if ($r | type) != "array" or ($r | length) == 0 then error("Deployment records must be a nonempty array") else . end |
  if any($r[]; type != "object" or (.name | type) != "string" or (.address | address | not) or (.number | type) != "number" or .number < 0 or .number != (.number | floor)) then error("Invalid deployment name, address or number") else . end |
  if ($r | map(.name) | unique | length) != ($r | length) then error("Duplicate deployment records") else . end |
  ($r | map({key:.name,value:.}) | from_entries) as $by_name |
  if any($mapping[]; . as $name | $by_name[$name].pending == true) then error("A required deployment is still pending") else . end |
  .l1StakingProxy //= "0x000000000000000000000000000000000000dEaD" |
  if (.l1StakingProxy | address | not) or (.l1StakingProxy | ascii_downcase) == ($by_name["Proxy__Submitter"].address // "" | ascii_downcase) then error("l1StakingProxy must be a nonzero address distinct from Proxy__Submitter") else . end |
  reduce ($mapping | to_entries[]) as $item (.;
    $by_name[$item.value].address as $deployed |
    if $deployed == null then error("Missing deployment record: " + $item.value)
    elif .[$item.key] == null or .[$item.key] == "0x0000000000000000000000000000000000000000" then .[$item.key] = $deployed
    elif (.[$item.key] | address | not) or (.[$item.key] | ascii_downcase) != ($deployed | ascii_downcase) then error("Configured address differs from deployment: " + $item.key)
    else . end) |
  del(.BLOCK_SIGNER_PRIVATE_KEY, .l2StakingPks)
' "$deploy_config" > "$attempt/deploy-config.json"
cp "$deployment_file" "$attempt/l1-deployments.json"
config=$attempt/deploy-config.json
json_hash() { printf '%s' "$(jq -S -c . "$1")" | shasum -a 256 | awk '{print $1}'; }
config_hash=$(json_hash "$config")
# initialize appends implementations; only these records were consumed by genesis.
jq --argjson mapping "$mapping" '[.[] | select(.name as $n | $mapping | any(. == $n)) | {name,address:(.address | ascii_downcase),number}] | sort_by(.name)' "$deployment_file" > "$attempt/genesis-deployments.json"
deployments_hash=$(json_hash "$attempt/genesis-deployments.json")
rpc() {
  request=$(jq -cn --arg method "$1" --argjson params "$2" '{jsonrpc:"2.0",id:1,method:$method,params:$params}')
  response=$(curl --fail --silent --max-time 30 -H 'Content-Type: application/json' --data "$request" "$l1_rpc") || fail "L1 RPC failed: $1"
  printf '%s' "$response" | jq -e 'if type == "object" and has("result") and (has("error") | not) then .result else error("Invalid RPC response") end'
}
chain_id=$(rpc eth_chainId '[]' | jq -e "$jq_common uint")
[ "$chain_id" = "$(jq -r "$jq_common .l1ChainID | uint" "$config")" ] || fail 'L1 RPC chain ID differs from deploy-config.l1ChainID'
genesis_hash=$(rpc eth_getBlockByNumber '["0x0",false]' | jq -er "$jq_common .hash | if hash then ascii_downcase else error(\"Invalid L1 genesis hash\") end")
for address in $(jq -r --argjson mapping "$mapping" '. as $config | $mapping | keys[] | $config[.]' "$config"); do
  rpc eth_getCode "[\"$address\",\"latest\"]" | jq -e 'type == "string" and test("^0x([0-9a-fA-F]{2})+$") and test("[1-9a-fA-F]")' >/dev/null || fail "No L1 contract code at $address"
done
identity=$(jq -cn --argjson chainId "$chain_id" --arg genesisHash "$genesis_hash" '{chainId:$chainId,genesisHash:$genesisHash}')
validate_outputs() {
  jq -en --slurpfile config "$config" --slurpfile genesis "$1/genesis-l2.json" --slurpfile rollup "$1/rollup.json" --slurpfile header "$1/genesis-batch-header.json" "$jq_common"'
    $config[0] as $c | $genesis[0] as $g | $rollup[0] as $r | $header[0] as $h |
    if ($g | type) != "object" or ($g.config | type) != "object" or ($g.alloc | type) != "object" or ($g.alloc | length) == 0 then error("Genesis requires config and nonempty alloc") else . end |
    if ($g.config.chainId | uint) != ($c.l2ChainID | uint) or ($r.l1_chain_id | uint) != ($c.l1ChainID | uint) or ($r.l2_chain_id | uint) != ($c.l2ChainID | uint) then error("Generated chain IDs do not match configuration") else . end |
    if ($r.l2_genesis_state_root | hash | not) or ($r.withdraw_root | hash | not) or ($h | type) != "string" or ($h | test("^0x[0-9a-fA-F]{514}$") | not) then error("Invalid genesis roots or 257-byte batch header") else . end |
    ($h | ascii_downcase) as $hex |
    if $hex[2:4] != "02" or ($hex[4:52] | test("^0+$") | not) or ($hex[52:116] | test("^0+$")) or ($hex[244:308] | test("^0+$")) or $hex[116:180] != "010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014" then error("Invalid V2 genesis batch header fields") else . end |
    if $hex[244:308] != ($r.l2_genesis_state_root[2:] | ascii_downcase) or $hex != ($r.genesis_batch_header | ascii_downcase) then error("Batch postStateRoot or header differs from rollup.json") else . end |
    if ($g.number | uint) != 0 or ($hex[500:516] | test("^0+$") | not) or ($g.timestamp | uint) == 0 then error("Genesis number must be zero and timestamp positive") else . end |
    if any($r.genesis.l1, $r.genesis.l2; type != "object" or (.hash | hash | not) or (.number | uint) < 0) or ($r.genesis.l2.number | uint) != 0 then error("Invalid rollup genesis block identity") else . end |
    {batchHeader:$h}
  '
}
file_hashes() {
  : > "$attempt/artifact-hashes.jsonl"
  for file in $files; do
    [ -f "$output_dir/$file" ] || fail "Missing artifact: $file"
    hash=$(shasum -a 256 "$output_dir/$file") || fail "Cannot hash artifact: $file"
    hash=${hash%% *}
    jq -cn --arg name "$file" --arg hash "$hash" '{key:$name,value:$hash}' >> "$attempt/artifact-hashes.jsonl"
  done
  jq -s 'from_entries' "$attempt/artifact-hashes.jsonl"
}
if "$verify_existing"; then
  hashes=$(file_hashes)
  jq -e --arg network "$network" --arg configHash "$config_hash" --arg deploymentsHash "$deployments_hash" --argjson identity "$identity" --argjson files "$hashes" '.network == $network and .configHash == $configHash and .deploymentsHash == $deploymentsHash and .l1Identity == $identity and .files == $files' "$output_dir/genesis.done" >/dev/null || fail 'Existing artifacts, inputs or L1 identity differ from genesis.done'
  override=$(validate_outputs "$output_dir")
  jq -e --argjson override "$override" '. == $override' "$output_dir/deployment-config.json" >/dev/null || fail 'deployment-config.json differs from the genesis header'
  rm -r "$attempt"
  attempt=
  printf 'Existing genesis verified without regeneration: %s\n' "$output_dir"
  exit 0
fi
# Invalidate the marker only after input and L1 checks succeed; retain old artifacts on failure.
rm -f "$output_dir/genesis.done"
status=0
(cd "$SCRIPT_DIR" && go run cmd/main.go genesis l2 --l1-rpc "$l1_rpc" --deploy-config "$config" --deployment-dir "$attempt/l1-deployments.json" --outfile.l2 "$attempt/genesis-l2.json" --outfile.rollup "$attempt/rollup.json" --outfile.genbatchheader "$attempt/genesis-batch-header.json") > "$attempt/genesis.raw.log" 2>&1 || status=$?
sed -E 's@https?://[^[:space:]"<>]+@<L1_RPC>@g' "$attempt/genesis.raw.log" > "$attempt/genesis.log"
rm "$attempt/genesis.raw.log"
[ "$status" -eq 0 ] || fail "Go genesis generation failed with exit code $status; log: $attempt/genesis.log"
validate_outputs "$attempt" > "$attempt/deployment-config.json"
for file in $files; do mv "$attempt/$file" "$output_dir/$file"; done
hashes=$(file_hashes)
jq -n --arg network "$network" --arg configHash "$config_hash" --arg deploymentsHash "$deployments_hash" --argjson l1Identity "$identity" --argjson files "$hashes" '{network:$network,configHash:$configHash,deploymentsHash:$deploymentsHash,l1Identity:$l1Identity,files:$files}' > "$attempt/genesis.done"
mv "$attempt/genesis.done" "$output_dir/genesis.done"
printf 'Genesis files generated: %s\nL1 initialization override: %s/deployment-config.json\n' "$output_dir" "$output_dir"
