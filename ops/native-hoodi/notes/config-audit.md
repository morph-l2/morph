# 配置审计(P3-4 ~ P3-7)

结论先行:**Hoodi 上不需要给任何 staking 地址或 BLOCK_SIGNER 打币;但 deployer 要多备一笔 submitter 的质押。**

## 1. 配置里的密钥与地址

| 配置项 | 性质 | 在哪里用 | Hoodi 上怎么办 |
|---|---|---|---|
| `BLOCK_SIGNER_PRIVATE_KEY` | **Anvil 公开私钥** | 只在 `ops/devnet-morph/devnet/__init__.py:190`,用来在 devnet 的 L1 上给 `l2StakingAddresses` 打 10 ETH;Go 创世工具**完全不读它** | 删除 |
| `BLOCK_SIGNER_ADDRESS` | 地址 | 代码里没有任何读取 | 保留字段,换成 sequencer 地址,无实际作用 |
| `l2StakingAddresses / Pks / TmKeys / BlsKeys` | 地址与**公钥** | 只进入 **L2 创世**(`immutables.go`,L2Staking 合约的序列器集合);从不在 L1 发交易 | 沿用;**不需要在 Hoodi 上给它们打币** |
| `finalSystemOwner / proxyAdminOwner / gasPriceOracleOwner / recordOracleAddress / morphTokenOwner` | L2 合约的 owner | L2 创世 | 全部换成 deployer 地址 |
| `l1FeeVaultRecipient / sequencerFeeVaultRecipient / l2BridgeFeeVaultRecipient` | 费用接收地址 | L2 创世 | 换成 deployer 地址 |

## 2. `l1StartingBlockTag`(P3-5)

- 格式:geth 的 `rpc.BlockNumberOrHash`,可以是 `"earliest"` / `"latest"` 等标签、十六进制块号字符串(如 `"0x3da3eb"`),或 `{"blockHash": ...}`。
- 创世工具(`cmd/genesis/cmd.go`)用 `BlockByNumber` / `BlockByHash` 取这个 L1 区块,取其**哈希、高度、时间戳**写进 `rollup.json`(`L2Time = l1StartBlock.Time()`)。
- 默认 `earliest` = Hoodi 第 0 块,会让 rollup.json 的 L1 起点是 2025 年 3 月的创世块。
- **取值方式:** 部署第一笔交易之前,取当时 Hoodi 的最新块号(十六进制),写进配置后再跑创世工具。
- 当前文件里放的是**故意无法解析**的占位符 `REPLACE_WITH_HOODI_START_BLOCK_HEX`,忘记替换会立刻报错,而不会静默用错值。
- 节点的 `MORPH_NODE_SYNC_START_HEIGHT` 也用这个块号。

## 3. main.py 运行时注入的字段(我们的 driver 要自己处理)

| main.py 写入的字段 | 实际是否有效 | 我们怎么做 |
|---|---|---|
| `l1GenesisBlockTimestamp` | **无效**。`DeployConfig` 结构里没有这个字段,会被忽略 | 不写 |
| `l1StartingBlockTag = earliest` | 有效,但不适合 Hoodi | 部署时替换为起始块号 |
| `l1StakingProxy = 0x...dEaD` | **必需**。为 0 时 `config.go` 先去找已部署的 L1Staking,找不到就报错;L1Staking 已退役 | 直接写进 `hoodidev-deploy-config.json`(已写) |

## 4. `register` 与质押(P3-6)

- `hardhat register`(`contracts/deploy/021-SubmitterRegister.ts`):owner 先 `addSubmitter(submitter)`,再调用 `stake(submitter, {value: minimumStake - currentStake})`。
- **质押的钱由 owner(deployer)付**,不是 submitter。
- 配置里 `submitterMinimumStake`、`submitterChallengeDeposit` 默认都是 **1 ETH**。合约只要求它们大于 0。
- 我们改成 **0.01 ETH**(`hoodidev.ts` 里),所以 deployer 额外只需要约 0.01 ETH 的质押。
- 这个小数值能否被 `parseEther` 之类的换算正确处理,会在 anvil 分叉干跑时实测。

## 5. `hardhat fund`(P3-7)

- 会给每个 `batchSubmitterPks` 里的账户转 **100 ETH**(余额低于约 10 ETH 时)。Hoodi 上不可行。
- **driver 不调用它。** 需要时手动向 submitter 转小额(额度由 P2-5 的干跑确定)。
