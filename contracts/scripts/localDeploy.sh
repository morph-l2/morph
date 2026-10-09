#!/bin/sh
set -eu
umask 077

SCRIPT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
CONTRACTS_DIR=$(CDPATH= cd -- "$SCRIPT_DIR/.." && pwd)
GENESIS_DIR=$(CDPATH= cd -- "$CONTRACTS_DIR/../ops/l2-genesis" && pwd)
network=devnet
output_dir= deploy_config= l1_rpc= delay=
config_override=${DEPLOY_CONFIG_OVERRIDE:-}
legacy_l1_deployment_file=${LEGACY_L1_DEPLOYMENT_FILE:-}
batch_block_interval=${TX_SUBMITTER_BATCH_BLOCK_INTERVAL:-}
batch_timeout=${TX_SUBMITTER_BATCH_TIMEOUT:-}
fail() { printf '%s\n' "$*" >&2; exit 1; }
while [ "$#" -gt 0 ]; do
    case "$1" in
        --help|-h)
            printf '%s\n' 'Usage: localDeploy.sh [--network devnet|qanet] [--l1-rpc URL] [--output-dir DIR]' \
                '       [--deploy-config FILE] [--config-override FILE] [--sequencer-address ADDRESS]' \
                '       [--rollup-delay-period SECONDS] [--legacy-l1-deployment-file FILE]' \
                '       --batch-block-interval BLOCKS --batch-timeout SECONDS' \
                'Batch parameters may also use TX_SUBMITTER_BATCH_BLOCK_INTERVAL and TX_SUBMITTER_BATCH_TIMEOUT.' \
                'Provide a confirmed Proxy__L1Staking record using --legacy-l1-deployment-file or LEGACY_L1_DEPLOYMENT_FILE; resumes may use the saved output records.'
            exit 0 ;;
        --network|--output-dir|--deploy-config|--config-override|--l1-rpc|--sequencer-address|--rollup-delay-period|--batch-block-interval|--batch-timeout|--legacy-l1-deployment-file)
            [ "$#" -ge 2 ] || fail "Missing value for $1"
            case "$1" in
                --network) network=$2 ;; --output-dir) output_dir=$2 ;;
                --deploy-config) deploy_config=$2 ;; --config-override) config_override=$2 ;;
                --l1-rpc) l1_rpc=$2 ;; --sequencer-address) export firstSequencerAddress=$2 ;;
                --rollup-delay-period) delay=$2 ;;
                --batch-block-interval) batch_block_interval=$2 ;;
                --batch-timeout) batch_timeout=$2 ;;
                --legacy-l1-deployment-file) legacy_l1_deployment_file=$2 ;;
            esac
            shift 2 ;;
        *) fail "Unknown option: $1" ;;
    esac
done
case "$network" in
    devnet) hardhat_network=l1; l1_rpc=${l1_rpc:-${L1_RPC_URL:-}}; export L1_RPC_URL=$l1_rpc ;;
    qanet) hardhat_network=qanetl1; l1_rpc=${l1_rpc:-${QA_RPC_URL:-}}; export QA_RPC_URL=$l1_rpc ;;
    *) fail 'Network must be devnet or qanet' ;;
esac
[ -n "$l1_rpc" ] || fail 'Set L1_RPC_URL (devnet), QA_RPC_URL (qanet), or --l1-rpc'
for command in node jq; do command -v "$command" >/dev/null || fail "Required command is unavailable: $command"; done
[ -n "$batch_block_interval" ] || fail 'Set --batch-block-interval or TX_SUBMITTER_BATCH_BLOCK_INTERVAL explicitly'
[ -n "$batch_timeout" ] || fail 'Set --batch-timeout or TX_SUBMITTER_BATCH_TIMEOUT explicitly'
batch_parameters=$(jq -cen --arg interval "$batch_block_interval" --arg timeout "$batch_timeout" '
    def uint64($name):
        if test("\\A[0-9]+\\z") then sub("^0+"; "") | if . == "" then "0" else . end
        else error($name + " must be an unsigned decimal integer") end |
        if length < 20 or (length == 20 and . <= "18446744073709551615") then .
        else error($name + " exceeds uint64") end;
    {batchBlockInterval: ($interval | uint64("batch block interval")), batchTimeout: ($timeout | uint64("batch timeout"))} |
    if .batchBlockInterval == "0" and .batchTimeout == "0" then error("batch block interval and batch timeout cannot both be zero") else . end
') || fail 'Invalid explicit batch parameters'
absolute_file() { (CDPATH= cd -- "$(dirname -- "$1")" && printf '%s/%s' "$PWD" "$(basename -- "$1")"); }
deploy_config=$(absolute_file "${deploy_config:-$GENESIS_DIR/deploy-config/$network-deploy-config.json}")
[ -f "$deploy_config" ] || fail "Missing genesis source: $deploy_config"
if [ -n "$config_override" ]; then config_override=$(absolute_file "$config_override"); fi
if [ -n "$legacy_l1_deployment_file" ]; then
    legacy_l1_deployment_file=$(absolute_file "$legacy_l1_deployment_file")
    [ -f "$legacy_l1_deployment_file" ] || fail "Missing legacy L1 deployment file: $legacy_l1_deployment_file"
fi
output_dir=${output_dir:-$GENESIS_DIR/.$network}
mkdir -p "$output_dir"
output_dir=$(CDPATH= cd -- "$output_dir" && pwd)
lock=$output_dir/.deployment.lock.d
mkdir "$lock" 2>/dev/null || fail "Deployment is locked: $lock; confirm no process is running before removing a stale lock"
trap 'rm -rf "$lock"' EXIT
trap 'exit 1' HUP INT TERM
deployment_file=$output_dir/${network}L1.json
identity=$output_dir/deployment-identity.json
if [ ! -f "$identity" ]; then
    for file in "$deployment_file" "$output_dir/genesis.done" "$output_dir/genesis-l2.json" "$output_dir/done"; do
        [ ! -e "$file" ] || fail 'Existing deployment has no matching identity; preserve its records and inspect them before continuing'
    done
fi
if [ -f "$deployment_file" ]; then
    jq -e 'type == "array" and length > 0' "$deployment_file" >/dev/null || fail 'Saved deployment records must be a nonempty array; preserve the output and restore its original records'
    cp "$deployment_file" "$lock/existing-records.json"
else
    printf '[]\n' > "$lock/existing-records.json"
fi
legacy_l1_deployment_file=${legacy_l1_deployment_file:-$deployment_file}
[ -f "$legacy_l1_deployment_file" ] || fail 'Provide an existing confirmed Proxy__L1Staking record with --legacy-l1-deployment-file or LEGACY_L1_DEPLOYMENT_FILE before deploying; no legacy contract is created automatically'
jq -e --slurpfile existing "$lock/existing-records.json" '
    def staking:
        if type != "array" then error("Legacy L1 deployment records must be an array") else . end |
        [.[] | select(.name == "Proxy__L1Staking")] |
        if length != 1 then error("Expected exactly one confirmed Proxy__L1Staking record") else .[0] end |
        if (.address | type) != "string" or (.address | test("\\A0x[0-9a-fA-F]{40}\\z") | not) then error("Invalid Proxy__L1Staking address") else . end |
        .address |= ascii_downcase |
        if .address == "0x0000000000000000000000000000000000000000" or .address == "0x000000000000000000000000000000000000dead" then error("Proxy__L1Staking must be an existing contract, not zero or a placeholder") else . end |
        if (.number | type) != "number" or .number < 0 or .number > 9007199254740991 or .number != (.number | floor) then error("Proxy__L1Staking number must be a nonnegative safe integer") else . end |
        if has("pending") and .pending != false then error("Proxy__L1Staking deployment must be confirmed") else . end |
        {name, address, number};
    . as $source | staking as $selected |
    if ($existing[0] | type) != "array" then error("Saved deployment records must be an array") else . end |
    if ($existing[0] | length) > 0 and ($existing[0] | staking) != $selected then error("Selected Proxy__L1Staking differs from saved deployment records; restore the original record") else . end |
    if any(($source + $existing[0])[]; .name == "Proxy__Submitter" and ((.address | ascii_downcase) == $selected.address)) then error("Proxy__L1Staking must not use the Submitter address") else . end |
    $selected
' "$legacy_l1_deployment_file" > "$lock/legacy-l1-staking.json" || fail 'Invalid legacy L1 deployment record; preserve existing output and correct the selected source'
if [ -n "$config_override" ]; then cp "$config_override" "$lock/input.json"; else printf '{}\n' > "$lock/input.json"; fi
if [ -n "$delay" ]; then
    jq --arg delay "$delay" '.rollupDelayPeriod = ($delay | tonumber)' "$lock/input.json" > "$lock/delay.json"
    mv "$lock/delay.json" "$lock/input.json"
fi
export DOTENV_CONFIG_PATH=/dev/null HARDHAT_NETWORK=$hardhat_network DEPLOY_CONFIG_OVERRIDE=$lock/input.json
cd "$CONTRACTS_DIR"
# Reuse the Hardhat configuration and validators before sending any transaction.
# Only public configuration and account addresses are persisted.
node -r ts-node/register/transpile-only - "$deploy_config" "$lock/legacy-l1-staking.json" > "$lock/identity.json" <<'NODE'
const fs = require('fs');
const crypto = require('crypto');
const hre = require('hardhat');
const { validateDeploymentConfig, registrationAddresses, firstSequencerAddress } = require('./src/deployment-validation');
(async () => {
    const config = { ...hre.deployConfig, batchHeader: '' };
    config.firstSequencerAddress = firstSequencerAddress(config);
    config.batchSubmitterAddresses = registrationAddresses(config);
    const walletAddress = (key, label) => {
        try { return new hre.ethers.Wallet(key).address; }
        catch (_) { throw new Error(`${label} must contain a valid private key`); }
    };
    const deployer = walletAddress(process.env.DEPLOYER_PRIVATE_KEY, 'DEPLOYER_PRIVATE_KEY');
    const owner = walletAddress(process.env.SUBMITTER_OWNER_PRIVATE_KEY || process.env.DEPLOYER_PRIVATE_KEY, 'Registration signer');
    if (owner.toLowerCase() !== config.submitterOwner.toLowerCase()) throw new Error('Registration signer must match submitterOwner');
    if (!hre.ethers.utils.isAddress(config.firstSequencerAddress) || config.firstSequencerAddress === hre.ethers.constants.AddressZero) throw new Error('Set a nonzero firstSequencerAddress');
    const source = fs.readFileSync(process.argv[2]);
    const genesis = JSON.parse(source);
    const legacyL1StakingDeployment = JSON.parse(fs.readFileSync(process.argv[3]));
    if (config.l1ChainID !== genesis.l1ChainID || config.l2ChainID !== genesis.l2ChainID) throw new Error('L1 and L2 configurations must have matching chain IDs');
    await validateDeploymentConfig(hre, config, 'deploy');
    if (genesis.l1StakingProxy && genesis.l1StakingProxy.toLowerCase() !== hre.ethers.constants.AddressZero && genesis.l1StakingProxy.toLowerCase() !== legacyL1StakingDeployment.address) throw new Error('Genesis l1StakingProxy differs from the selected legacy deployment record');
    if (legacyL1StakingDeployment.number > await hre.ethers.provider.getBlockNumber()) throw new Error('Proxy__L1Staking deployment block is not confirmed on the selected L1');
    if (await hre.ethers.provider.getCode(legacyL1StakingDeployment.address) === '0x') throw new Error('Proxy__L1Staking has no code on the selected L1; provide a confirmed legacy contract');
    const l1GenesisHash = (await hre.ethers.provider.getBlock(0)).hash;
    process.stdout.write(JSON.stringify({ network: hre.network.name, config, deployer, owner, l1GenesisHash, legacyL1StakingDeployment,
        genesisConfigHash: crypto.createHash('sha256').update(source).digest('hex') }));
})().catch(error => { console.error(error.message); process.exitCode = 1; });
NODE
jq -S --argjson batchParameters "$batch_parameters" '. + {batchParameters:$batchParameters}' "$lock/identity.json" > "$lock/identity.sorted.json"
if [ -f "$identity" ]; then
    jq -e '.batchParameters | type == "object"' "$identity" >/dev/null || fail 'Existing deployment identity has no recorded batch parameters; preserve its output and inspect the original scripts and inputs. Automatic identity migration is not supported'
    jq -e '.legacyL1StakingDeployment | type == "object"' "$identity" >/dev/null || fail 'Existing deployment identity has no recorded legacy L1Staking deployment; preserve its output and inspect the original scripts and inputs. Automatic identity migration is not supported'
    jq -S . "$identity" > "$lock/previous.json"
    cmp -s "$lock/previous.json" "$lock/identity.sorted.json" || fail 'Saved chain, accounts or configuration differ, including explicit batch parameters; restore the original inputs and values. Preserve this directory and use a separate directory for a new deployment'
else
    mv "$lock/identity.sorted.json" "$identity"
fi
if [ ! -f "$deployment_file" ]; then
    jq '[.legacyL1StakingDeployment]' "$identity" > "$lock/seed.json"
    mv "$lock/seed.json" "$deployment_file"
fi
jq '.config' "$identity" > "$lock/config.json"
genesis() { sh "$GENESIS_DIR/$network-l2genesis.sh" --deploy-config "$deploy_config" --deployment-file "$deployment_file" --output-dir "$output_dir" --l1-rpc "$l1_rpc" "$@"; }
task() { ./node_modules/.bin/hardhat "$1" --network "$hardhat_network" --storagepath "$deployment_file"; }
if [ -f "$output_dir/genesis.done" ]; then
    genesis --verify-existing
    jq -s '.[0] * .[1]' "$lock/config.json" "$output_dir/deployment-config.json" > "$lock/merged.json"
    mv "$lock/merged.json" "$lock/config.json"
elif [ -f "$output_dir/done" ]; then
    fail 'Completed deployment is missing genesis.done; preserve the directory and inspect its artifacts'
fi
cp "$lock/config.json" "$output_dir/contract-config.json"
export DEPLOY_CONFIG_OVERRIDE=$output_dir/contract-config.json
if [ -f "$output_dir/done" ]; then
    ./node_modules/.bin/hardhat verify-deployment --network "$hardhat_network" --storagepath "$deployment_file" --runtime
else
    task deploy
    if [ ! -f "$output_dir/genesis.done" ]; then genesis; fi
    jq -s '.[0] * .[1]' "$lock/config.json" "$output_dir/deployment-config.json" > "$lock/merged.json"
    mv "$lock/merged.json" "$output_dir/contract-config.json"
    task initialize
    (export DEPLOYER_PRIVATE_KEY=${SUBMITTER_OWNER_PRIVATE_KEY:-$DEPLOYER_PRIVATE_KEY}; task register)
    task verify-deployment
fi
jq -er --slurpfile identity "$identity" '
    def address($name): [.[] | select(.name == $name) | .address] |
        if length == 1 and (.[0] | test("^0x[0-9a-fA-F]{40}$")) then .[0] else error("Missing or ambiguous address: " + $name) end;
    "MORPH_NODE_L1_SEQUENCER_CONTRACT=\(address("Proxy__L1Sequencer"))",
    "MORPH_NODE_SYNC_DEPOSIT_CONTRACT_ADDRESS=\(address("Proxy__L1MessageQueueWithGasPriceOracle"))",
    "MORPH_NODE_ROLLUP_ADDRESS=\(address("Proxy__Rollup"))",
    "TX_SUBMITTER_SUBMITTER_ADDRESS=\(address("Proxy__Submitter"))",
    "TX_SUBMITTER_ROLLUP_ADDRESS=\(address("Proxy__Rollup"))",
    "TX_SUBMITTER_BATCH_BLOCK_INTERVAL=\($identity[0].batchParameters.batchBlockInterval)",
    "TX_SUBMITTER_BATCH_TIMEOUT=\($identity[0].batchParameters.batchTimeout)",
    "TX_SUBMITTER_BATCH_V2_UPGRADE_TIME=0"
' "$deployment_file" > "$lock/runtime.env"
mv "$lock/runtime.env" "$output_dir/runtime.env"
printf 'L1 contracts and genesis verified; external services require separate startup and checks.\n' > "$output_dir/done"
printf 'Deployment verified: %s\nPublic service parameters: %s/runtime.env\n' "$output_dir" "$output_dir"
