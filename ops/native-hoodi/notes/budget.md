# 资金预算(P2-5 的结果)

依据:在 Hoodi 的 anvil 分叉上把 `scripts/deploy-l1.sh` 完整干跑了一遍(`run/dryrun-1/`),**所有步骤成功**。分叉里的交易不花公共 Hoodi 的任何币。

## 实测

| 项 | 数值 |
|---|---|
| deployer 发出的交易数 | **77** |
| 总 gas | **51,304,680**(最大的一笔 4,780,806,远低于区块 gas 上限 6000 万) |
| submitter 质押(由 deployer 代付) | 0.01 ETH(配置里调低过,默认是 1 ETH) |
| 干跑时 deployer 实际花费 | 0.0179 ETH(= 约 0.0079 ETH gas + 0.01 ETH 质押),对应分叉里约 0.15 gwei 的 gas 价格 |
| 耗时 | 部署约 2 分钟,初始化约 2 分钟,总共约 4 分钟(分叉里无需等待确认) |

## 按不同 gas 价格换算 deployer 的需求

(51,304,680 gas × 价格 + 0.01 ETH 质押)

| Hoodi gas 价格 | deployer 需要 |
|---|---|
| 0.15 gwei | 约 0.018 ETH |
| 0.5 gwei | 约 0.036 ETH |
| 1 gwei(我 09:19 测到的 1.08 gwei) | 约 0.061 ETH |
| 2 gwei | 约 0.113 ETH |
| 5 gwei | 约 0.267 ETH |

**建议 deployer 备 0.2 ETH**(覆盖 gas 涨到约 3 gwei,并留余量)。真正部署前再查一次当时的 gas 价格。

## submitter

- 这次**没有实测**:干跑里 submitter 没发交易(批次提交要等 L2 跑起来)。
- 它每个批次要发一笔 blob 交易。blob 费用本身很低(Hoodi 当时 blob base fee 约 0.05 gwei),主要成本是执行 gas。
- 批次间隔由配置决定(`govBatchBlockInterval=200` 个块 / `govBatchTimeout=600` 秒),大致每 10 分钟一个批次(L2 有交易时更快)。
- **估算(未实测):** 先备 **0.1 ETH**,P6 跑起来后按实际消耗再调;`TX_SUBMITTER_TX_FEE_LIMIT` 要调低。

## 结论

**所需的测试币比我最初担心的少很多:合计约 0.3 ETH 量级就够。** 水龙头单次通常能领到这个数量级,所以 R2(领不到足够测试币)的风险下降,但领币本身还没有验证过。
