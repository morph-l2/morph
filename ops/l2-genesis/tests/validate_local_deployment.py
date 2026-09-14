#!/usr/bin/env python3
"""Validate both deployment flows and transaction-free repeats on a temporary Anvil chain."""

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
    directory = Path(tempfile.mkdtemp(prefix="morph-deployment-validation-"))
    print(f"Validation output: {directory}", flush=True)
    with socket.socket() as sock:
        sock.bind(("127.0.0.1", 0))
        port = sock.getsockname()[1]
    url = f"http://127.0.0.1:{port}"
    env = dict(os.environ)
    for name in ("DEPLOY_CONFIG_OVERRIDE", "DEPLOYER_PRIVATE_KEY", "SUBMITTER_OWNER_PRIVATE_KEY", "batchSubmitterPks",
                 "firstSequencerAddress", "QA_ROLLUP_DELAY_PERIOD", "DEPLOY_ROLLUP_DELAY_PERIOD",
                 "TX_SUBMITTER_BATCH_V2_UPGRADE_TIME"):
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
        results = {}
        for network in ('devnet', 'qanet'):
            override = directory / f'{network}-override.json'
            override.write_text(json.dumps({'contractAdmin': account['address'], 'submitterOwner': account['address'],
                                            'batchSubmitterAddresses': [account['address']]}))
            output = directory / network
            command = ['sh', str(ROOT / 'contracts/scripts/localDeploy.sh'), '--network', network,
                       '--output-dir', str(output), '--config-override', str(override), '--l1-rpc', url,
                       '--sequencer-address', account['address'], '--rollup-delay-period', '86400']
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
            results[network] = {'complete': True, 'repeatTransactions': 0,
                                'contracts': len(json.loads((output / f'{network}L1.json').read_text()))}
            print(f'{network}: deployment, genesis and state checks passed; repeat transactions: 0', flush=True)
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
