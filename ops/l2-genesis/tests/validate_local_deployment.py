#!/usr/bin/env python3
"""Check deployment preflight on Anvil, optionally exercising a synthetic legacy fixture."""

import argparse
import json
import os
from pathlib import Path
import socket
import subprocess
import sys
import tempfile
import time
import urllib.request

ROOT = Path(__file__).resolve().parents[3]


def rpc(url, method, params):
    request = urllib.request.Request(url, json.dumps({"jsonrpc": "2.0", "id": 1, "method": method, "params": params}).encode(),
                                     {"Content-Type": "application/json"})
    with urllib.request.urlopen(request, timeout=5) as response:
        data = json.load(response)
    if "error" in data:
        raise RuntimeError(data["error"])
    return data["result"]


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--legacy-staking-fixture', action='store_true',
                        help='Exercise the deployment pipeline using synthetic code on temporary Anvil only; this does not validate a production legacy staking contract.')
    arguments = parser.parse_args()
    directory = Path(tempfile.mkdtemp(prefix="morph-deployment-validation-"))
    print(f"Validation output: {directory}", flush=True)
    with socket.socket() as sock:
        sock.bind(("127.0.0.1", 0))
        port = sock.getsockname()[1]
    url = f"http://127.0.0.1:{port}"
    env = dict(os.environ)
    for name in ("DEPLOY_CONFIG_OVERRIDE", "DEPLOYER_PRIVATE_KEY", "SUBMITTER_OWNER_PRIVATE_KEY", "batchSubmitterPks",
                 "firstSequencerAddress", "QA_ROLLUP_DELAY_PERIOD", "DEPLOY_ROLLUP_DELAY_PERIOD",
                 "TX_SUBMITTER_BATCH_V2_UPGRADE_TIME", "TX_SUBMITTER_BATCH_BLOCK_INTERVAL", "TX_SUBMITTER_BATCH_TIMEOUT",
                 "LEGACY_L1_DEPLOYMENT_FILE"):
        env.pop(name, None)
    env["DOTENV_CONFIG_PATH"] = os.devnull
    # Use only Anvil's default test account; do not read workspace private keys.
    key_script = "const {Wallet}=require('ethers'); const w=Wallet.fromMnemonic('test test test test test test test test test test test junk'); process.stdout.write(JSON.stringify({address:w.address,key:w.privateKey}));"
    account = json.loads(subprocess.check_output(["node", "-e", key_script], cwd=ROOT / 'contracts', env=env, text=True))
    env["DEPLOYER_PRIVATE_KEY"] = account["key"]
    with (directory / 'anvil.log').open('w') as log:
        chain = subprocess.Popen(["anvil", "--silent", "--host", "127.0.0.1", "--port", str(port), "--chain-id", "900",
                                  "--hardfork", "cancun"], stdout=log, stderr=subprocess.STDOUT, env=env)
    try:
        for _ in range(100):
            if chain.poll() is not None:
                raise RuntimeError('Temporary Anvil process failed to start')
            try:
                rpc(url, 'eth_chainId', [])
                break
            except OSError:
                time.sleep(0.1)
        else:
            raise RuntimeError('Temporary Anvil RPC did not become ready')
        results = {'mode': 'synthetic-legacy-fixture' if arguments.legacy_staking_fixture else 'preflight-rejections',
                   'legacyContractValidation': 'Not performed. No production legacy staking contract is deployed or approved by this harness.'}
        fixture_address = '0x1000000000000000000000000000000000000001'
        fixture = directory / 'synthetic-legacy-l1-deployments.json'
        if arguments.legacy_staking_fixture:
            # This reverting code tests the script input boundary, not legacy staking behavior.
            rpc(url, 'anvil_setCode', [fixture_address, '0x60006000fd'])
            fixture.write_text(json.dumps([{'name': 'Proxy__L1Staking', 'address': fixture_address,
                                            'number': int(rpc(url, 'eth_blockNumber', []), 16)}]))
            print('Synthetic Anvil fixture enabled: pipeline checks do not prove a valid production legacy staking deployment.', flush=True)
        for network in ('devnet', 'qanet'):
            override = directory / f'{network}-override.json'
            override.write_text(json.dumps({'contractAdmin': account['address'], 'submitterOwner': account['address'],
                                            'batchSubmitterAddresses': [account['address']]}))
            output = directory / network
            command = ['sh', str(ROOT / 'contracts/scripts/localDeploy.sh'), '--network', network,
                       '--output-dir', str(output), '--config-override', str(override), '--l1-rpc', url,
                       '--sequencer-address', account['address'], '--rollup-delay-period', '86400',
                       '--batch-block-interval', '100', '--batch-timeout', '10']
            before = rpc(url, 'eth_getTransactionCount', [account['address'], 'latest'])
            rejected_log = directory / f'{network}-missing-legacy.log'
            with rejected_log.open('w') as stream:
                rejected = subprocess.run(command, env=env, cwd=directory, stdout=stream, stderr=subprocess.STDOUT)
            if rejected.returncode == 0 or 'existing confirmed Proxy__L1Staking record' not in rejected_log.read_text():
                raise RuntimeError(f'{network} did not reject missing legacy deployment input; log: {rejected_log}')
            if before != rpc(url, 'eth_getTransactionCount', [account['address'], 'latest']):
                raise RuntimeError(f'{network} missing-legacy preflight sent transactions')
            if any((output / name).exists() for name in ('deployment-identity.json', f'{network}L1.json', 'done')):
                raise RuntimeError(f'{network} missing-legacy preflight persisted deployment artifacts')
            results[network] = {'missingLegacyRejected': True, 'rejectedPreflightTransactions': 0}
            print(f'{network}: missing legacy deployment record rejected; transactions: 0', flush=True)

            for condition, record, expected in (
                ('no-code', {'name': 'Proxy__L1Staking', 'address': account['address'], 'number': 0}, 'has no code on the selected L1'),
                ('future-block', {'name': 'Proxy__L1Staking', 'address': fixture_address,
                                  'number': int(rpc(url, 'eth_blockNumber', []), 16) + 1}, 'deployment block is not confirmed'),
            ):
                invalid_record = directory / f'{network}-{condition}-legacy.json'
                invalid_record.write_text(json.dumps([record]))
                logfile = directory / f'{network}-{condition}.log'
                before = rpc(url, 'eth_getTransactionCount', [account['address'], 'latest'])
                with logfile.open('w') as stream:
                    rejected = subprocess.run(command + ['--legacy-l1-deployment-file', str(invalid_record)],
                                              env=env, cwd=directory, stdout=stream, stderr=subprocess.STDOUT)
                if rejected.returncode == 0 or expected not in logfile.read_text():
                    raise RuntimeError(f'{network} did not reject {condition} legacy deployment input; log: {logfile}')
                if before != rpc(url, 'eth_getTransactionCount', [account['address'], 'latest']):
                    raise RuntimeError(f'{network} {condition} preflight sent transactions')
                if any((output / name).exists() for name in ('deployment-identity.json', f'{network}L1.json', 'done')):
                    raise RuntimeError(f'{network} {condition} preflight persisted deployment artifacts')
                results[network][condition + 'Rejected'] = True
            if not arguments.legacy_staking_fixture:
                continue

            command += ['--legacy-l1-deployment-file', str(fixture)]
            for attempt in (1, 2):
                logfile = directory / f'{network}-{attempt}.log'
                before = rpc(url, 'eth_getTransactionCount', [account['address'], 'latest'])
                with logfile.open('w') as stream:
                    status = subprocess.run(command, env=env, cwd=directory, stdout=stream, stderr=subprocess.STDOUT)
                if status.returncode:
                    raise RuntimeError(f'{network} attempt {attempt} failed; log: {logfile}')
                after = rpc(url, 'eth_getTransactionCount', [account['address'], 'latest'])
                if attempt == 2 and before != after:
                    raise RuntimeError(f'{network} repeat execution sent additional transactions')
                if not (output / 'done').exists():
                    raise RuntimeError(f'{network} is missing its completion marker')
            identity = json.loads((output / 'deployment-identity.json').read_text())
            if identity['batchParameters'] != {'batchBlockInterval': '100', 'batchTimeout': '10'}:
                raise RuntimeError(f'{network} did not persist explicit batch parameters')
            runtime = (output / 'runtime.env').read_text()
            if 'TX_SUBMITTER_BATCH_BLOCK_INTERVAL=100\n' not in runtime or 'TX_SUBMITTER_BATCH_TIMEOUT=10\n' not in runtime:
                raise RuntimeError(f'{network} did not export explicit batch parameters')
            results[network].update({'syntheticPipelineVerified': True, 'repeatTransactions': 0,
                                     'deploymentRecords': len(json.loads((output / f'{network}L1.json').read_text()))})
            print(f'{network}: synthetic legacy fixture pipeline verified; repeat transactions: 0. Production legacy staking behavior is unverified.', flush=True)
        (directory / 'result.json').write_text(json.dumps(results, indent=2) + '\n')
        return 0
    finally:
        chain.terminate()
        try:
            chain.wait(timeout=10)
        except subprocess.TimeoutExpired:
            chain.kill()
            chain.wait()


if __name__ == '__main__':
    try:
        sys.exit(main())
    except Exception as error:
        print(str(error), file=sys.stderr)
        sys.exit(1)
