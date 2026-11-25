# 运行手册:Morph L2(原生进程)+ 公共 Hoodi 作为 L1

这份手册讲**怎么用**这套环境。想了解**为什么这样设计、中间踩过什么坑**,看 `PLAN.md` 和 `notes/issues.md`。

## 1. 它是什么

在本机用原生进程(不用 Docker)跑一条私有的 Morph L2(链 ID 53077),它的 L1 是**公共的以太坊 Hoodi 测试网**(链 ID 560048)。合约部署在 Hoodi 上,批次以 EIP-4844 blob 交易提交到 Hoodi。

```
 公共 Hoodi:执行层 RPC(主 ethpandaops,备 0xrpc.io / publicnode)+ 信标 RPC(ethpandaops / publicnode)
        ▲
   l1-proxy(本机 127.0.0.1:18546,rpc-consistency-proxy.js)——所有进程访问 L1 都经过它
        ▲
 ┌──────┴────────────────────────────────────────────────────────────────────────┐
 │ morph-el-0 ◀▶ node-0   出块节点(tendermint 单验证者,持有出块密钥)            │
 │ morph-el-1 ◀▶ node-1   layer1 校验模式:不跑 tendermint,独立从 L1 的 blob 还原  │
 │ morph-el-2 ◀▶ node-2   全节点:跑 tendermint,没有出块密钥。**只为让 node-0 能重启** │
 │ tx-submitter           把 L2 批次提交到 Hoodi 的 Rollup,并最终化               │
 │ gas-price-oracle       把 L1 费用写入 L2 的 GasPriceOracle                      │
 │ challenger(可选)     比对批次根,不一致时发起 Rollup.challengeState            │
 └────────────────────────────────────────────────────────────────────────────────┘
```

## 2. 前置条件

- 构建环境见 README 第 2 节。所有脚本通过 `scripts/lib/env.sh` 检查工具链;工具不在 PATH 上时,用环境变量 `MORPH_ENV_FILE` 指向一个设置 PATH 的文件。手工调用 `cast` 时需要自己保证它在 PATH 上。
- 二进制在 `bin/`:`morph-geth/geth`、`morphnode/morphnode`、`morphnode/tendermint`、`tx-submitter/tx-submitter`、`gas-oracle/app`、`challenger/challenger`(构建方法见 `notes/versions.md` 和 `TODO.md` 的 P1;`build-all.sh` 一并构建)。
- 密钥在 `env/secrets.env`(权限 600,**不要提交、不要贴到任何地方**),模板是 `env/hoodi.env.example`。5 个角色:deployer、sequencer、batch submitter、gas oracle、challenger(可选,`init-secrets.sh` 每次生成新的)。

## 3. 日常操作

全部通过 `scripts/morph-native.sh`:

| 命令 | 作用 |
|---|---|
| `status` | 9 个进程(加可选的 challenger)、三个 L2 的高度、Hoodi 高度、代理统计、Rollup 已提交批次 |
| `up [NAME...]` | 启动全部或指定进程(按依赖顺序)。`challenger` 不在默认列表里:用 `up challenger` 启动,或在 `env/secrets.env` 设 `CHALLENGER_ENABLE=1` |
| `down` | 停止全部进程(含 challenger),**保留链数据** |
| `kill NAME` / `start NAME` | 停 / 启动单个进程(`kill` 后不会被自动重启) |
| `logs NAME` | 跟随某个进程的日志(文件在 `run/l2/logs/`) |
| `setup` | 用 `deploy-l1.sh` 产出的创世文件生成运行目录(含 node-2) |
| `clean logs` / `clean l2` | **不删除**,只把数据挪到 `run/archive/`。`clean l2` 有重要后果,见第 7 节 |

进程名:`l1-proxy morph-el-0 node-0 morph-el-1 node-1 morph-el-2 node-2 tx-submitter gas-price-oracle`,另有可选的 `challenger`。每个进程都在一个自动重启的守护外壳里,崩溃后 2 秒内被拉起。

## 4. 端口(全部 127.0.0.1)

| 进程 | HTTP | WS | 引擎 | p2p | metrics |
|---|---|---|---|---|---|
| morph-el-0 | 8545 | 8546 | 8551 | 30303 | 6160 |
| morph-el-1 | 8645 | 8646 | 8651 | 30304 | 6161 |
| morph-el-2 | 8745 | 8746 | 8751 | 30305 | 6162 |
| node-0(tendermint) | | | | 26656 | 26660;tendermint rpc 26657 |
| node-1 | | | | 不跑 tendermint | 27660 |
| node-2(tendermint) | | | | 26666 | 26670;tendermint rpc 26667 |
| tx-submitter | | | | | 8060 |
| gas-price-oracle | | | | | 6070 |
| l1-proxy | 18546(`/stats` 看统计) | | | | |
| challenger(可选) | 8090 | | | | 6065 |

geth 的 metrics 端口用 6160 起,是因为有些环境里 6060 已被占用;challenger 同理避开它默认的 8080 / 6060(用 `CHALLENGER_HTTP_PORT` / `CHALLENGER_METRICS_PORT` 可改)。**这些端口(8545、8551、30303 等)与其他本地以太坊 devnet 常用端口相同**,不能同时跑;先 `down` 或换端口。

## 5. 从零开始的流程

1. `scripts/build-all.sh`:编译全部二进制、安装合约依赖(P1)。
2. 生成密钥(`cast wallet new`),写入 `env/secrets.env`;给 **deployer** 和 **batch submitter** 在 Hoodi 上领测试币。
3. 部署 L1 合约并生成 L2 创世:
   ```bash
   RPC_URL=https://rpc.hoodi.ethpandaops.io RUN_DIR=$PWD/run/l1 WAIT_FINALIZED=1 SUBMITTER_FUND_ETH=1 scripts/deploy-l1.sh
   RPC_URL=https://rpc.hoodi.ethpandaops.io scripts/verify-l1.sh      # 读链上状态核对,不信退出码
   ```
   约 20 分钟(其中约 13 分钟是等 Hoodi finalized)。**在公共链上花测试币、不可撤销。**
   `deploy-l1.sh` 最后还会跑 s34:把 challenger 加入 `Rollup.isChallenger` 并充值到 0.2 ETH(`CHALLENGER_FUND_ETH`)。`env/secrets.env` 里没有 `CHALLENGER_PRIVATE_KEY` 时跳过;已有的旧部署用 `scripts/register-challenger.sh` 补做。
4. `scripts/morph-native.sh setup`,然后 `up`。
5. `scripts/l2-bootstrap.sh`:给 deployer 和 oracle 在 L2 上转余额,并把 oracle 密钥加入 GasPriceOracle 白名单(否则 gas-price-oracle 没权限更新)。
6. 用第 6 节的脚本验证。
7. (可选)`scripts/morph-native.sh up challenger`,按第 6 节“挑战者”验证。

## 6. 验证脚本

| 脚本 | 验证什么 | 耗时 |
|---|---|---|
| `scripts/verify-l1.sh` | 读 Hoodi 上的合约状态:owner、互相引用、创世批次、submitter 质押、所有代理的 implementation 与 admin、第一个序列器(53 项) | 1 分钟 |
| `scripts/deposit-test.sh` | L1→L2 存款到一个全新地址,等两个节点都到账 | 约 7 分钟 |
| `scripts/withdraw-test.sh` | L2→L1 提款:本地生成默克尔证明,在 L1 上领取 | 约 10 分钟 |
| `scripts/restart-test.sh` | 崩溃与重启恢复(T1 到 T8,T7 需单独指定) | 约 35 分钟 |
| `scripts/rpc-consistency-proxy.selftest.js` | 代理的 6 项行为(不碰网络) | 10 秒 |
| `scripts/withdraw-proof.selftest.js` | 默克尔证明算法与合约逐行一致(210 组) | 5 秒 |
| `scripts/probe-hoodi.sh` | 公共 RPC 的限流探测 | 1 分钟 |
| `scripts/register-challenger.sh` | 把 challenger 加入 `Rollup.isChallenger` 并充值(幂等,读回链上结果确认;`deploy-l1.sh` 的 s34 就是它) | 1 分钟 |

比较 node-1 与 node-0 是否一致,用 `cast block N --field hash --rpc-url http://127.0.0.1:8545 / 8645 / 8745` 逐块比对(已在三个节点间比过 0 到 7562 共 7563 个块,0 个不一致)。

**挑战者(可选)。** challenger 是个 HTTP 服务,**不会自己轮询**,只在被请求时才动作:

- `GET /batch/{index}`:读该批次在 L1 的提交交易和 el-0 上同一个块,比对 state root 与 withdraw root。链正常时返回 `"verdict":"match"`,否则 `mismatch` / `inconclusive`。不发交易。
- `POST /challenge`,body `{"batchIndex":N,"batchHash":"0x..."}`:**只有**该 hash 正是 L1 上当前提交的那个、且根不一致时才发 `Rollup.challengeState`;其它情况都返回 4xx、不发交易。
- 监听 `127.0.0.1:8090`(无鉴权,所以只绑本机),指标 `127.0.0.1:6065`;L1 走 l1-proxy,L2 读 el-0(归档节点,`--gcmode=archive`,读 withdraw root 需要)。账户是 `env/secrets.env` 里的 `CHALLENGER_PRIVATE_KEY`(不是 deployer)。

```bash
scripts/morph-native.sh up challenger
curl -s http://127.0.0.1:8090/health                                   # {"status":"ok"}
FIN=$(cast call $ROLLUP 'lastFinalizedBatchIndex()(uint256)' --rpc-url "$L1_ETH_RPC")
curl -s http://127.0.0.1:8090/batch/$FIN            # 409 "batch already finalized"
curl -s http://127.0.0.1:8090/batch/$((FIN + 1))    # 已提交未 finalized 的批次:"verdict":"match"(还没有则 404)
```

成功的挑战是一笔真实交易:锁定保证金(`challengeDeposit`),并让 finalize 停住,直到 Rollup owner 调 `revertBatch`;超过 `proofWindow` 后,提交该批次的 submitter 会被罚没。不要随手发。全部配置项见 `challenger/env.example`。

## 7. 必须知道的限制和坑

1. **node-2 不能去掉。** 出块节点 node-0 重启后,必须有至少一个 tendermint 对端它才会恢复出块;没有时它静默停转(进程在、RPC 在、日志无错误、高度冻结)。node-1 是 layer1 校验模式,不算 tendermint 对端。见 `notes/issues.md` I-028。
2. **L2 和 L1 上的 Rollup 是一一绑定的。** L1 上已提交的批次绑定着当前这条 L2 链。把 L2 清空重新从创世出块,新区块与旧批次对不上,tx-submitter 无法继续。**重置 L2 必须同时重新部署 L1 合约**(新的 `RUN_DIR`),所以 `clean l2` 只是把数据挪走。
3. **存款的 gas 上限要给够。** 仓库示例里的 `110000` 在这套部署上会让 L2 上的执行失败(钱锁在 L1 网关里)。`deposit-test.sh` 默认 400000。
4. **公共 RPC 不可靠。** 背后是多个节点:刚上链的区块可能读不到,偶尔限流。`l1-proxy` 解决这两件事,所有进程都应经过它,不要直接连 Hoodi。
5. **node-1 不是实时的。** 它按批次推进,批次提交到 L1 后才前进,所以总会落后几百个块(约一个批次)。
6. **不要把 `hardhat` 任务的退出码当成功。** 它们出错时常常只打印一行然后退出码为 0。部署后用 `verify-l1.sh` 读链上状态。
7. **owner / 管理员 / 费用接收方都是 deployer,不是写死在仓库里的地址。** `deploy-l1.sh` 从 `env/secrets.env` 的密钥推导出地址,经环境变量传给 `hoodidev.ts`,并填进生成的创世配置(`run/l1/hoodidev-deploy-config.generated.json`)。谁来部署,谁就是 owner。序列器集合、Tendermint 与 BLS 公钥是固定的。
8. **改动在 git 里。** 对上游仓库的改动只有 4 个文件:`contracts/hardhat.config.ts`(加网络 `hoodidev`)、`contracts/src/deploy-config/hoodidev.ts`、`ops/l2-genesis/deploy-config/hoodidev-deploy-config.json`、`contracts/src/deploy-utils.ts`(可选的读重试)。
9. **challenger 不会自动挑战。** 它不轮询 L1,必须有人 `POST /challenge`。挑战成功后 finalize 会停住,需要 owner `revertBatch` 收尾;`proofWindow` 过后调 `proveState` 判挑战方胜,会罚没提交该批次的 submitter 并让合约进入 `pending revert request`。

## 8. 成本(实测)

| 项目 | 数值 |
|---|---|
| L1 合约部署 | 77 笔交易、5130 万 gas;按 1 gwei 约 0.06 ETH(含 0.01 ETH 的 submitter 质押) |
| tx-submitter 运行 | 约 0.0004 ETH / 批次;批次约每 7 分钟一个,约 0.09 ETH / 天 |
| challenger 充值 | 部署时转 0.2 ETH(`CHALLENGER_FUND_ETH`);一次挑战锁定 0.01 ETH 保证金 + 约 0.0002 ETH gas |
| 部署期间的损耗 | 前几次失败遗留的孤儿合约,合计不到 0.2 ETH |

## 9. 故障排查

| 现象 | 可能原因与做法 |
|---|---|
| L2 高度不动,进程都在 | 先看 node-0 日志是否停在 `derivation skipped: sequencer mode`:node-2 没跑,或没连上。`status` 看 node-2,日志里应有 `Added peer` |
| node-1 高度不涨 | 正常:等下一个批次提交。`status` 里 `lastCommittedBatchIndex` 在涨就没问题 |
| `failed to fetch L1 messages ... beyond current head` | 落后的 L1 后端;代理会重试。偶发无需处理 |
| gas-price-oracle 报 `ErrCallerNotAllowed` | 没跑 `l2-bootstrap.sh`,oracle 不在白名单 |
| 存款后 L2 没到账,但 L1 消息已入队 | 在 L2 上找类型 `0x7e` 的交易看回执:状态 0 通常是 `_gasLimit` 太小 |
| 起不来,端口被占 | 先 `scripts/morph-native.sh down`;本机上别的 devnet 或节点在用同样的端口也会冲突(`ss -ltn` 查) |
| `up challenger` 报缺 `CHALLENGER_PRIVATE_KEY` / `bin/challenger/challenger` | 旧的 `env/secrets.env` 没有该密钥(加一把 `cast wallet new`,再跑 `register-challenger.sh`);二进制要先跑 `build-all.sh` |
| `POST /challenge` 返回 `only challenger allowed` | challenger 账户没加入 `Rollup.isChallenger`:跑 `scripts/register-challenger.sh` |
| challenger 进程反复重启 | 看 `logs challenger`:L2 RPC 的链与 Rollup 对不上时它拒绝启动;确认 el-0 在跑、`L1_RUN` 指向当前部署 |

## 10. 重置与部署记录(2026-11-19 新增,已用这套流程完整重做过一次)

- `scripts/reset-env.sh plan|run|restore|list`:把运行数据(和可选的构建产物 `--with-build`、密钥 `--new-keys`)用 `mv` 归档到 `run/archive/<时间戳>/`,`restore` 可原样挪回。不删除任何东西。submitter 在 L1 上余额高于 0.02 ETH 时拒绝执行,先用 `scripts/sweep-submitter.sh`(要先 `down`)把钱转回 deployer。
- `scripts/export-deployment.sh NAME`:把一次部署的合约地址、创世文件、账户地址(无私钥)、版本、测试输出存到 `deployments/NAME/`。重置之前先导出。
- 已有记录:`deployments/devnet-1/`(旧)、`deployments/devnet-2/`(当前)。
- 重做顺序:sweep → export 旧部署 → reset-env run --with-build --new-keys → build-all → deploy-l1(约 43 分钟,其中 `s30_initialize` 约 20 分钟)→ verify-l1 → setup → up → l2-bootstrap → 存款/提款/重启测试。
