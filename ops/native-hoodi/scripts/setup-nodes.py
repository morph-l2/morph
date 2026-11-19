#!/usr/bin/env python3
"""setup-nodes.py — 生成 2 节点(单序列器)Morph L2 的本地运行目录,不依赖 Docker。

改编自上游 ops/devnet-morph/devnet/setup_nodes.py,差异:
  - 只做 node-0(tendermint 验证者,出块)和 node-1(layer1 校验模式,不启动 tendermint)两个节点
  - 所有路径落在 --run-dir 下;所有主机名换成 127.0.0.1;端口按 PLAN.md 第 6 节分开
  - 不依赖 PATH 里的 tendermint,用 --tendermint 指定全路径
  - 同时初始化两个 morph-geth 的数据目录(geth init)并放好 nodekey / static-nodes.json

输出(都在 --run-dir 下):
  nodes/node0, nodes/node1      tendermint / morphnode 的 home
  el0, el1                      morph-geth 的 datadir(已 init)
  jwt-secret.txt                引擎 API 的 JWT
幂等:目标目录已存在则拒绝覆盖(不删除任何东西,换一个 --run-dir 或自行处理)。
"""
import argparse
import base64
import hashlib
import json
import os
import re
import shutil
import subprocess
import sys

# 端口规划(PLAN.md 第 6 节)
EL = {
    "el0": dict(http=8545, ws=8546, authrpc=8551, p2p=30303, metrics=6060),
    "el1": dict(http=8645, ws=8646, authrpc=8651, p2p=30304, metrics=6061),
}
NODE0 = dict(p2p=26656, rpc=26657, prometheus=26660)


def node_id(node_key_path):
    """tendermint 节点 ID = sha256(ed25519 公钥)前 20 字节;私钥 = seed(32) || pubkey(32)。"""
    with open(node_key_path) as f:
        priv = json.load(f)["priv_key"]["value"]
    return hashlib.sha256(base64.b64decode(priv)[32:]).hexdigest()[:40]


def run(cmd, **kw):
    print("+", " ".join(cmd))
    subprocess.run(cmd, check=True, **kw)


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--morph", required=True, help="Morph 仓库根目录")
    ap.add_argument("--run-dir", required=True, help="输出目录(必须不存在或为空)")
    ap.add_argument("--genesis", required=True, help="L2 创世文件 genesis-l2.json")
    ap.add_argument("--tendermint", required=True, help="tendermint 二进制的全路径")
    ap.add_argument("--geth", required=True, help="morph-geth 二进制的全路径")
    a = ap.parse_args()

    docker_dir = os.path.join(a.morph, "ops", "docker")
    run_dir = os.path.abspath(a.run_dir)
    if os.path.exists(run_dir) and os.listdir(run_dir):
        sys.exit(f"{run_dir} already exists and is not empty; refusing to overwrite")
    os.makedirs(run_dir, exist_ok=True)
    nodes_dir = os.path.join(run_dir, "nodes")

    # ---- tendermint homes: 1 validator (node0) + 1 non-validator (node1) ----
    run([a.tendermint, "testnet", "--v", "1", "--n", "1", "--o", nodes_dir,
         "--populate-persistent-peers", "--hostname", "127.0.0.1", "--hostname", "127.0.0.1"])

    for node in ("node0", "node1"):
        cfg_dir = os.path.join(nodes_dir, node, "config")
        shutil.copyfile(os.path.join(docker_dir, node, "node_key.json"), os.path.join(cfg_dir, "node_key.json"))
        shutil.copyfile(os.path.join(docker_dir, "tendermint-devnet-genesis.json"), os.path.join(cfg_dir, "genesis.json"))
        if node == "node0":
            shutil.copyfile(os.path.join(docker_dir, node, "priv_validator_key.json"),
                            os.path.join(cfg_dir, "priv_validator_key.json"))
        else:
            # only the sole genesis validator may hold a validator key; otherwise block sync is disabled
            for f in (os.path.join(cfg_dir, "priv_validator_key.json"),
                      os.path.join(nodes_dir, node, "data", "priv_validator_state.json")):
                if os.path.exists(f):
                    os.remove(f)

    # ---- config.toml: same tuning as upstream, ports/peers for a single host ----
    for node in ("node0", "node1"):
        path = os.path.join(nodes_dir, node, "config", "config.toml")
        with open(path) as f:
            c = f.read()
        c = c.replace('create_empty_blocks_interval = "0s"', 'create_empty_blocks_interval = "5s"')
        c = c.replace('peer_gossip_sleep_duration = "100ms"', 'peer_gossip_sleep_duration = "10ms"')
        c = c.replace('flush_throttle_timeout = "100ms"', 'flush_throttle_timeout = "10ms"')
        c = c.replace('max_packet_msg_payload_size = 1024', 'max_packet_msg_payload_size = 10485760')
        c = c.replace('send_rate = 5120000', 'send_rate = 52428800')
        c = c.replace('recv_rate = 5120000', 'recv_rate = 102428800')
        c = c.replace('block_sync = false', 'block_sync = true')
        c = c.replace('prometheus = false', 'prometheus = true')
        # only node0 runs tendermint; nobody else to dial in single-sequencer mode
        c = re.sub(r'persistent_peers\s*=\s*".*?"', 'persistent_peers = ""', c)
        if node == "node0":
            c = c.replace('pex = true', 'pex = false')
            c = re.sub(r'(\[rpc\][\s\S]*?laddr\s*=\s*)"[^"]*"', rf'\1"tcp://127.0.0.1:{NODE0["rpc"]}"', c, count=1)
            c = re.sub(r'(\[p2p\][\s\S]*?laddr\s*=\s*)"[^"]*"', rf'\1"tcp://127.0.0.1:{NODE0["p2p"]}"', c, count=1)
            c = re.sub(r'prometheus_listen_addr\s*=\s*"[^"]*"', f'prometheus_listen_addr = "127.0.0.1:{NODE0["prometheus"]}"', c)
        with open(path, "w") as f:
            f.write(c)

    # ---- morph-geth datadirs ----
    jwt_src = os.path.join(docker_dir, "jwt-secret.txt")
    shutil.copyfile(jwt_src, os.path.join(run_dir, "jwt-secret.txt"))
    for i, el in enumerate(("el0", "el1")):
        datadir = os.path.join(run_dir, el)
        os.makedirs(os.path.join(datadir, "geth"), exist_ok=True)
        shutil.copyfile(os.path.join(docker_dir, f"nodekey{i}"), os.path.join(datadir, "geth", "nodekey"))
        run([a.geth, "--verbosity=3", "init", f"--datadir={datadir}", a.genesis])

    # static-nodes.json: the upstream file names containers; rewrite to loopback with our p2p ports
    with open(os.path.join(docker_dir, "static-nodes.json")) as f:
        static = json.load(f)
    fixed = []
    for enode in static:
        m = re.match(r"(enode://[0-9a-f]+)@morph-el-(\d):30303", enode)
        if not m:
            sys.exit(f"unexpected enode format: {enode}")
        port = EL[f"el{m.group(2)}"]["p2p"]
        fixed.append(f"{m.group(1)}@127.0.0.1:{port}")
    with open(os.path.join(run_dir, "el1", "geth", "static-nodes.json"), "w") as f:
        json.dump(fixed, f, indent=2)

    print()
    print("node0 id:", node_id(os.path.join(nodes_dir, "node0", "config", "node_key.json")))
    print("done:", run_dir)


if __name__ == "__main__":
    main()
