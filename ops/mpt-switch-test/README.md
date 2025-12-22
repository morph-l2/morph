# MPT Switch 测试

测试 sequencer node 和 sentry node 在 MPT 升级时间点的客户端切换逻辑。

## 架构

```
升级前:                              升级后 (交换):
┌─────────────┐                     ┌─────────────┐
│  Sequencer  │──► ZK Geth          │  Sequencer  │──► MPT Geth
│    Node     │    (:8545)          │    Node     │    (:9545)
└─────────────┘                     └─────────────┘
                        ──────►
┌─────────────┐                     ┌─────────────┐
│   Sentry    │──► MPT Geth         │   Sentry    │──► ZK Geth
│    Node     │    (:9545)          │    Node     │    (:8545)
└─────────────┘                     └─────────────┘
```

**关键点：**
- 两个 Node 共用同样的两个 Geth 实例
- 升级时间到达后，两个 Node 交换 Geth 连接
- Sequencer: ZK Geth → MPT Geth
- Sentry: MPT Geth → ZK Geth

## 前置条件

### 准备二进制文件

将所有二进制文件放到 `bin` 目录下：

```bash
ops/mpt-switch-test/bin/
├── geth
├── morphnode
└── tendermint
```

Genesis 文件 (`genesis-l2.json`) 已包含在目录中。

## 使用方法

```bash
cd /Users/corey.zhang/workspace/morph/ops/mpt-switch-test

# 1. 启动测试环境 (60秒后触发切换)
./test-mpt-switch-local.sh start 60

# 2. 监控 Sequencer 切换日志
./test-mpt-switch-local.sh monitor sequencer

# 3. 监控 Sentry 切换日志
./test-mpt-switch-local.sh monitor sentry

# 4. 查看状态
./test-mpt-switch-local.sh status

# 5. 停止服务
./test-mpt-switch-local.sh stop

# 6. 清理数据
./test-mpt-switch-local.sh clean
```

## 命令列表

| 命令 | 说明 |
|------|------|
| `start [delay]` | 启动测试环境，delay 为 MPT 切换延迟秒数 (默认 60) |
| `stop` | 停止所有服务 |
| `clean` | 清理所有测试数据 |
| `status` | 查看服务状态和区块高度 |
| `monitor [target]` | 监控日志 (sequencer/sentry/all) |
| `logs [service]` | 查看日志 (sequencer/sentry/zk-geth/mpt-geth/all) |

## 端口分配

| 服务 | HTTP | Engine | P2P |
|------|------|--------|-----|
| ZK Geth | 8545 | 8551 | 30303 |
| MPT Geth | 9545 | 9551 | 30304 |
| Sequencer Node | - | - | 26656 (RPC: 26657) |
| Sentry Node | - | - | 26756 (RPC: 26757) |

## 预期日志

两个 Node 都应该看到类似的切换日志：

```
MPT switch time reached, MUST wait for MPT node to sync
  mpt_time=<timestamp> current_time=<timestamp> target_block=<number>

Waiting for MPT node to sync...
  remote_block=<n> target_block=<m> blocks_behind=<diff>

Successfully switched to MPT client
  remote_block=<n> target_block=<m> wait_duration=<duration>
```

## 文件结构

```
mpt-switch-test/
├── test-mpt-switch-local.sh  # 测试脚本
├── README.md                 # 本文档
├── genesis-l2.json           # L2 创世文件 (已包含)
├── bin/                      # 二进制文件目录 (需手动放置)
│   ├── geth
│   ├── morphnode
│   └── tendermint
└── .testdata/                # 测试数据目录 (启动时生成)
    ├── zk-geth/              # ZK Geth 数据
    ├── mpt-geth/             # MPT Geth 数据
    ├── sequencer-node/       # Sequencer Node 数据
    ├── sentry-node/          # Sentry Node 数据
    ├── jwt-secret.txt        # JWT 密钥
    └── *.log                 # 日志文件
```
