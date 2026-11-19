# 版本记录(P1-9)

记录日期:2026-11-18

## 源码

| 项目 | 版本 |
|---|---|
| Morph 仓库 | `ae987b15`(2026-11-11,"add reproduceable code for CICD (#1088)"),工作分支 `native-hoodi` |
| go-ethereum 子模块 | `1c528b351`(2026-11-09),tag `morph-v2.2.8` |

## 构建环境

| 工具 | 版本 |
|---|---|
| Go | 1.25.7(Morph 的 `go.work` 要求 ≥ 1.24.0,满足) |
| Rust | 1.95.0(Docker 版用 1.88) |
| Node | v24.19.0 |
| yarn | 1.22.22 |
| Foundry(cast / forge) | 1.9.0 |
| gcc | 13 |
| solc(hardhat 自动下载) | 0.8.24,evm target cancun |

## 二进制(`ops/native-hoodi/bin/`)

| 路径 | 版本输出 | 大小 | 编译耗时 |
|---|---|---|---|
| `morph-geth/geth` | `Version: 2.2.8-mainnet` | 53 MB | 23 秒 |
| `morphnode/morphnode` | `--help` 正常(不支持 `--version`) | 52 MB | 35 秒(与 tendermint 一起) |
| `morphnode/tendermint` | `0.37.0-alpha.1` | 41 MB | |
| `tx-submitter/tx-submitter` | `v1.0.0-2.2.8-mainnet-ae987b15-1794400668` | 35 MB | 19 秒 |
| `gas-oracle/app` | 需要环境变量才能运行(见 I-004) | 25 MB | 2 分 20 秒 |

合计约 197 MB。

## 合约工具链

| 步骤 | 结果 |
|---|---|
| `yarn install --frozen-lockfile`(在 `contracts/`) | 安装 1039 个包,49 秒 |
| `yarn hardhat compile` | 223 个 Solidity 文件通过,31 秒,Node 24 可用 |
| `yarn build`(完整,即 `main.py` 部署前跑的那条) | 通过,37 秒;含 `autogen:artifacts`、`build:ts`、`typechain`、`go build` 差异测试工具 |
| 编译后 `git status` | 干净(产物都被 `.gitignore` 忽略) |

## 日志

所有构建日志在 `run/build-logs/`:`p1-2-morph-geth.log`、`p1-3-node.log`、`p1-4-tx-submitter.log`、`p1-5-gas-oracle.log`、`p1-6-yarn-install.log`、`p1-6-hardhat-compile.log`、`p1-6-yarn-build.log`。
