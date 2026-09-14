#!/bin/sh
set -eu
umask 077

SCRIPT_DIR=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
CONTRACTS_DIR=$(CDPATH= cd -- "$SCRIPT_DIR/.." && pwd)
GENESIS_DIR=$(CDPATH= cd -- "$CONTRACTS_DIR/../ops/l2-genesis" && pwd)
network=devnet
output_dir= deploy_config= l1_rpc= delay=
config_override=${DEPLOY_CONFIG_OVERRIDE:-}
fail() { printf '%s\n' "$*" >&2; exit 1; }
while [ "$#" -gt 0 ]; do
    case "$1" in
        --help|-h)
            printf '%s\n' 'Usage: localDeploy.sh [--network devnet|qanet] [--l1-rpc URL] [--output-dir DIR]' \
                '       [--deploy-config FILE] [--config-override FILE] [--sequencer-address ADDRESS]' \
                '       [--rollup-delay-period SECONDS]'
            exit 0 ;;
        --network|--output-dir|--deploy-config|--config-override|--l1-rpc|--sequencer-address|--rollup-delay-period)
            [ "$#" -ge 2 ] || fail "Missing value for $1"
            case "$1" in
                --network) network=$2 ;; --output-dir) output_dir=$2 ;;
                --deploy-config) deploy_config=$2 ;; --config-override) config_override=$2 ;;
                --l1-rpc) l1_rpc=$2 ;; --sequencer-address) export firstSequencerAddress=$2 ;;
                --rollup-delay-period) delay=$2 ;;
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
absolute_file() { (CDPATH= cd -- "$(dirname -- "$1")" && printf '%s/%s' "$PWD" "$(basename -- "$1")"); }
deploy_config=$(absolute_file "${deploy_config:-$GENESIS_DIR/deploy-config/$network-deploy-config.json}")
[ -f "$deploy_config" ] || fail "Missing genesis source: $deploy_config"
if [ -n "$config_override" ]; then config_override=$(absolute_file "$config_override"); fi
output_dir=${output_dir:-$GENESIS_DIR/.$network}
mkdir -p "$output_dir"
output_dir=$(CDPATH= cd -- "$output_dir" && pwd)
lock=$output_dir/.deployment.lock.d
mkdir "$lock" 2>/dev/null || fail "Deployment is locked: $lock; confirm no process is running before removing a stale lock"
trap 'rm -rf "$lock"' EXIT
trap 'exit 1' HUP INT TERM
deployment_file=$output_dir/${network}L1.json
if [ -n "$config_override" ]; then cp "$config_override" "$lock/input.json"; else printf '{}\n' > "$lock/input.json"; fi
if [ -n "$delay" ]; then
    jq --arg delay "$delay" '.rollupDelayPeriod = ($delay | tonumber)' "$lock/input.json" > "$lock/delay.json"
    mv "$lock/delay.json" "$lock/input.json"
fi
export DOTENV_CONFIG_PATH=/dev/null HARDHAT_NETWORK=$hardhat_network DEPLOY_CONFIG_OVERRIDE=$lock/input.json
cd "$CONTRACTS_DIR"
# Reuse the Hardhat configuration and validators before sending any transaction.
# Only public configuration and account addresses are persisted.
node -r ts-node/register/transpile-only - "$deploy_config" > "$lock/identity.json" <<'NODE'
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
    if (config.l1ChainID !== genesis.l1ChainID || config.l2ChainID !== genesis.l2ChainID) throw new Error('L1 and L2 configurations must have matching chain IDs');
    await validateDeploymentConfig(hre, config, 'deploy');
    const l1GenesisHash = (await hre.ethers.provider.getBlock(0)).hash;
    process.stdout.write(JSON.stringify({ network: hre.network.name, config, deployer, owner, l1GenesisHash,
        genesisConfigHash: crypto.createHash('sha256').update(source).digest('hex') }));
})().catch(error => { console.error(error.message); process.exitCode = 1; });
NODE
jq -S . "$lock/identity.json" > "$lock/identity.sorted.json"
identity=$output_dir/deployment-identity.json
if [ -f "$identity" ]; then
    jq -S . "$identity" > "$lock/previous.json"
    cmp -s "$lock/previous.json" "$lock/identity.sorted.json" || fail 'Saved chain, accounts or configuration differ; preserve this directory and use a separate directory for a new deployment'
else
    for file in "$deployment_file" "$output_dir/genesis.done" "$output_dir/genesis-l2.json" "$output_dir/done"; do
        [ ! -e "$file" ] || fail 'Existing deployment has no matching identity; preserve its records and inspect them before continuing'
    done
    mv "$lock/identity.sorted.json" "$identity"
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
jq -er --slurpfile config "$deploy_config" '
    def address($name): [.[] | select(.name == $name) | .address] |
        if length == 1 and (.[0] | test("^0x[0-9a-fA-F]{40}$")) then .[0] else error("Missing or ambiguous address: " + $name) end;
    "MORPH_NODE_L1_SEQUENCER_CONTRACT=\(address("Proxy__L1Sequencer"))",
    "MORPH_NODE_SYNC_DEPOSIT_CONTRACT_ADDRESS=\(address("Proxy__L1MessageQueueWithGasPriceOracle"))",
    "MORPH_NODE_ROLLUP_ADDRESS=\(address("Proxy__Rollup"))",
    "TX_SUBMITTER_SUBMITTER_ADDRESS=\(address("Proxy__Submitter"))",
    "TX_SUBMITTER_ROLLUP_ADDRESS=\(address("Proxy__Rollup"))",
    "TX_SUBMITTER_BATCH_BLOCK_INTERVAL=\($config[0].govBatchBlockInterval)",
    "TX_SUBMITTER_BATCH_TIMEOUT=\($config[0].govBatchTimeout)",
    "TX_SUBMITTER_BATCH_V2_UPGRADE_TIME=0"
' "$deployment_file" > "$lock/runtime.env"
mv "$lock/runtime.env" "$output_dir/runtime.env"
printf 'L1 contracts and genesis verified; external services require separate startup and checks.\n' > "$output_dir/done"
printf 'Deployment verified: %s\nPublic service parameters: %s/runtime.env\n' "$output_dir" "$output_dir"
