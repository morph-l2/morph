#!/usr/bin/env python3
"""add-node2.py — add a second tendermint-running node (node-2 + morph-el-2) to an existing L2 run directory.

Why: after node-0 (the only tendermint node) is restarted it can never resume block production. Its block pool must
report "caught up" before it starts the sequencer routines, and blocksync/pool.go IsCaughtUp() returns false while the
pool has no peer. node-1 runs in layer1 verify mode and does not run tendermint, so it does not count. Morph's own
ops/README.md says the same: "At least one other node running tendermint, i.e. two nodes minimum".

node-2 is an ordinary full node: it runs tendermint, follows node-0 over P2P, executes blocks on its own morph-geth, and
holds NO sequencer key (only one node may hold it unless HA is enabled).

It leaves everything that exists untouched, except that node-0's persistent_peers gains node-2 (so a restarted node-0
dials it at once). It refuses to overwrite an existing node2 / el2.

Usage: add-node2.py --morph REPO --run-dir run/l2 --tendermint BIN --geth BIN
"""
import argparse, base64, hashlib, json, os, re, shutil, subprocess, sys, tempfile

N2_P2P, N2_RPC, N2_PROM = 26666, 26667, 26670     # node-2 tendermint ports (node-0 uses 26656/26657/26660)
EL2_P2P = 30305                                    # morph-el-2 p2p; the other ports follow morph-native.sh's formula


def node_id(path):
    with open(path) as f:
        priv = json.load(f)["priv_key"]["value"]
    return hashlib.sha256(base64.b64decode(priv)[32:]).hexdigest()[:40]


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--morph", required=True)
    ap.add_argument("--run-dir", required=True)
    ap.add_argument("--tendermint", required=True)
    ap.add_argument("--geth", required=True)
    ap.add_argument("--genesis", help="L2 genesis file (genesis-l2.json); default: run/l1/genesis/genesis-l2.json next to the run dir")
    a = ap.parse_args()
    docker_dir = os.path.join(a.morph, "ops", "docker")
    run_dir = os.path.abspath(a.run_dir)
    nodes = os.path.join(run_dir, "nodes")
    n2, el2 = os.path.join(nodes, "node2"), os.path.join(run_dir, "el2")
    for p in (n2, el2):
        if os.path.exists(p):
            sys.exit(f"{p} already exists; refusing to overwrite")
    for p in (os.path.join(nodes, "node0", "config", "config.toml"), os.path.join(run_dir, "el0")):
        if not os.path.exists(p):
            sys.exit(f"{p} not found: run setup first")

    # ---- node-2's tendermint home: take the skeleton of a 1-validator + 2-non-validator testnet, keep only node2
    tmp = tempfile.mkdtemp(prefix="node2-skel-")
    subprocess.run([a.tendermint, "testnet", "--v", "1", "--n", "2", "--o", tmp, "--populate-persistent-peers",
                    "--hostname", "127.0.0.1", "--hostname", "127.0.0.1", "--hostname", "127.0.0.1"],
                   check=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
    shutil.copytree(os.path.join(tmp, "node2"), n2)
    cfg_dir = os.path.join(n2, "config")
    shutil.copyfile(os.path.join(docker_dir, "node2", "node_key.json"), os.path.join(cfg_dir, "node_key.json"))
    shutil.copyfile(os.path.join(docker_dir, "tendermint-devnet-genesis.json"), os.path.join(cfg_dir, "genesis.json"))
    # not a genesis validator: it must not hold a validator key (a node holding the sole genesis key gets block sync disabled)
    for f in (os.path.join(cfg_dir, "priv_validator_key.json"), os.path.join(n2, "data", "priv_validator_state.json")):
        if os.path.exists(f):
            os.remove(f)

    id0 = node_id(os.path.join(nodes, "node0", "config", "node_key.json"))
    id2 = node_id(os.path.join(cfg_dir, "node_key.json"))
    path = os.path.join(cfg_dir, "config.toml")
    c = open(path).read()
    for old, new in [('create_empty_blocks_interval = "0s"', 'create_empty_blocks_interval = "5s"'),
                     ('peer_gossip_sleep_duration = "100ms"', 'peer_gossip_sleep_duration = "10ms"'),
                     ('flush_throttle_timeout = "100ms"', 'flush_throttle_timeout = "10ms"'),
                     ('max_packet_msg_payload_size = 1024', 'max_packet_msg_payload_size = 10485760'),
                     ('send_rate = 5120000', 'send_rate = 52428800'), ('recv_rate = 5120000', 'recv_rate = 102428800'),
                     ('block_sync = false', 'block_sync = true'), ('prometheus = false', 'prometheus = true'),
                     ('pex = true', 'pex = false')]:
        c = c.replace(old, new)
    c = re.sub(r'persistent_peers\s*=\s*".*?"', f'persistent_peers = "{id0}@127.0.0.1:26656"', c)
    c = re.sub(r'proxy_app\s*=\s*"[^"]*"', 'proxy_app = "tcp://127.0.0.1:26668"', c, count=1)
    c = re.sub(r'(\[rpc\][\s\S]*?laddr\s*=\s*)"[^"]*"', rf'\1"tcp://127.0.0.1:{N2_RPC}"', c, count=1)
    c = re.sub(r'(\[p2p\][\s\S]*?laddr\s*=\s*)"[^"]*"', rf'\1"tcp://127.0.0.1:{N2_P2P}"', c, count=1)
    c = re.sub(r'prometheus_listen_addr\s*=\s*"[^"]*"', f'prometheus_listen_addr = "127.0.0.1:{N2_PROM}"', c)
    open(path, "w").write(c)

    # ---- node-0 dials node-2 as soon as it starts
    p0 = os.path.join(nodes, "node0", "config", "config.toml")
    c0 = open(p0).read()
    c0, n = re.subn(r'persistent_peers\s*=\s*".*?"', f'persistent_peers = "{id2}@127.0.0.1:{N2_P2P}"', c0)
    assert n == 1
    open(p0, "w").write(c0)

    # ---- morph-el-2: same genesis as the others, nodekey2, static peers to el-0 / el-1 (loopback ports from morph-native.sh)
    # must be the genesis of the deployment this chain belongs to (L1_RUN), not whatever sits in run/l1: a different genesis hash
    # makes node-2 reject node-0's blocks ("wrong parent hash") and node-0 then has no tendermint peer
    genesis = a.genesis or os.path.join(run_dir, "..", "l1", "genesis", "genesis-l2.json")
    os.makedirs(os.path.join(el2, "geth"))
    shutil.copyfile(os.path.join(docker_dir, "nodekey2"), os.path.join(el2, "geth", "nodekey"))
    subprocess.run([a.geth, "--verbosity=3", "init", f"--datadir={el2}", os.path.abspath(genesis)], check=True,
                   stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
    shutil.copyfile(os.path.join(run_dir, "el1", "geth", "static-nodes.json"), os.path.join(el2, "geth", "static-nodes.json"))
    print("node-2 id :", id2)
    print("node-0 id :", id0)
    print("added     :", n2, "and", el2)


if __name__ == "__main__":
    main()
