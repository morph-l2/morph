# 问题与处理记录

按时间顺序,每条记录:现象 → 原因 → 处理 → 对计划的影响。

---

## 2026-11-18

### I-001 morph-geth 按计划里的命令编译失败 — 已解决

- **现象:** `cd repos/morph/go-ethereum && go run build/ci.go install ./cmd/geth` 报
  `build/ci.go:63:2: use of internal package github.com/morph-l2/go-ethereum/internal/build not allowed`
- **原因:** `go-ethereum` 目录位于 Morph 仓库内,仓库根有 `go.work`,但 `go.work` 没把 `./go-ethereum` 列进去。Go 于是把 `github.com/morph-l2/go-ethereum` 当成**下载的外部依赖**,而不是本地目录,`internal` 包就不允许访问。Docker 版没这个问题,是因为它把 `go-ethereum` 单独拷到 `/go-ethereum`,上面没有 `go.work`。
- **处理:** 加 `GOWORK=off`。Morph 自己顶层 Makefile 的 `geth` 目标也是这么写的(`env GO111MODULE=on GOWORK=off go run build/ci.go install ./cmd/geth`)。我当初是从 Dockerfile 抄的命令,漏了这一点。
- **结果:** 编译成功,23 秒;`geth version` 显示 `Version: 2.2.8-mainnet`,与子模块 tag `morph-v2.2.8` 一致。
- **对计划的影响:** TODO P1-2 的命令已改。

### I-002 morphnode 不支持 `--version`

- **现象:** `morphnode --version` 报 `flag provided but not defined: -version`。
- **处理:** 版本用 `morphnode --help` 确认可运行,用 `tendermint version`(0.37.0-alpha.1)确认 tendermint。tx-submitter 支持 `--version`,输出 `v1.0.0-2.2.8-mainnet-ae987b15-...`。
- **对计划的影响:** 无。P1 的验收标准里"能打印版本"对 morphnode 改为"`--help` 正常输出"。

### I-003 tx-submitter 的版本串带 `mainnet`

- **现象:** `tx-submitter version v1.0.0-2.2.8-mainnet-ae987b15-...`;morph-geth 也是 `2.2.8-mainnet`。
- **说明:** 这只是构建标签,不是说它们要连主网。运行时的链由配置和 genesis 决定(`TX_SUBMITTER_BUILD_ENV=dev`、`--networkid=53077`)。记录在此,避免以后看到 `mainnet` 字样产生误会。

### I-004 gas-oracle 启动即 panic:日志目录无写权限 — 已找到处理方式

- **现象:** 直接运行 `bin/gas-oracle/app`,在 `src/main.rs:58` panic,`OutputIo(PermissionDenied)`。
- **原因:** 日志目录默认是 `/data/logs/morph-gas-price-oracle`(`main.rs` 的 `setup_logging`)。Docker 里以 root 运行,这个路径可写;本机没有 root,`/data` 不可写。
- **处理:** 设置环境变量 `GAS_ORACLE_LOG_DIR` 指向可写目录(计划用 `run/logs/gas-oracle/`)。设置后程序正常启动,打印 `Starting update gas oracle...`,随后因缺少 `L2_GAS_ORACLE_PRIVATE_KEY` 而 panic,这是预期行为。
- **对计划的影响:** P6-3 的 gas-oracle 环境变量清单要加上 `GAS_ORACLE_LOG_DIR`。

### I-005 gas-oracle 的信标地址只支持单个 — 已确认

- **现象/依据:** `gas_price_oracle.rs:36` 里 `l1_beacon_rpc: String`;`da_scalar/blob_client.rs` 里把它当 URL 前缀直接拼接 `/eth/v1/beacon/blob_sidecars/<slot>`,没有任何逗号拆分或回退。
- **结论:** `GAS_ORACLE_L1_BEACON_RPC` **不能**配逗号分隔的多个端点,只能配一个。这与 morphnode(支持多个、按序回退)不同。
- **处理:** gas-oracle 用 `https://beacon.hoodi.ethpandaops.io`;morphnode 配两个。
- **对计划的影响:** 这就是 TODO P6-3 里"待确认"的那一项,现已确认,`env/hoodi.env.example` 的注释同步修正。

### I-006 合约工具链在 Node 24 上可用 — 风险 R8 解除

- **结果:** `yarn hardhat compile` 与完整的 `yarn build` 都在 Node v24.19.0 上通过。
- **对计划的影响:** TODO P1-6 里"若失败则装 Node 18/20"的备选不再需要。PLAN 风险 R8 可关闭。

### I-007 node-1 不运行 tendermint,只有 node-0 运行 — 简化了 P5

- **依据:** `ops/devnet-morph/devnet/setup_nodes.py` 的注释:"node-1 runs with `MORPH_NODE_DERIVATION_VERIFY_MODE=layer1` and never starts tendermint; node-2 has no compose service at all"。`TENDERMINT_PEERS = ("node0", "ha-node0", "ha-node1", "ha-node2")`,不含 node1。
- **结论:** 单序列器模式下,只有 node-0 是 tendermint 节点;node-1 完全靠 L1 的批次和 blob 还原区块。因此 PBFT 对端的 `persistent_peers` 在这个模式下基本用不上。
- **对计划的影响:** PLAN 第 6 节端口表里 node-1 去掉 tendermint 端口;TODO P5-1 缩小范围。这也意味着 **node-1 能否与 node-0 一致,完全取决于信标链 blob 是否取得到**,P7-2 因此更关键。

### I-008 setup_nodes.py 只用标准库 — P1-8 通过

- **结果:** 依赖只有 `base64 hashlib json os re shutil subprocess sys`,Python 3.12.3 可直接运行,无需 `pip install`。
- **注意:** 脚本开头会检查 PATH 里有没有 `tendermint`,没有就调用 `make install-tendermint`。我们的 driver 复制其逻辑时,要直接用 `bin/morphnode/tendermint` 的全路径,不依赖 PATH。

### I-009 git 提交被 GPG 签名拦下 — 改为暂存区加补丁

- **现象:** `git commit` 报 `gpg failed to sign the data ... Operation cancelled`。
- **原因:** 该仓库的 git 配置强制提交签名,而当前环境没有可交互的 GPG 签名。
- **处理:** 我**没有**为了提交去关闭签名(那是你的签名策略)。三个改动文件保留在 **git 暂存区**,并导出补丁到 `patches/0001-hoodidev-config.patch`(166 行)。
- **核对:** 补丁里没有出现我们的任何私钥;公开 Anvil 私钥只出现一次,在 `hardhat.config.ts` 的"禁止使用"判断常量里。
- **需要你做:** 如果要在 `native-hoodi` 分支上留下提交,请你自己提交一次(能走你的签名),或告诉我允许对这个分支临时不签名。

### I-010 原计划里 P3 的几个"待确认"现已全部确认

见 `notes/config-audit.md`。最重要的两条:`BLOCK_SIGNER_PRIVATE_KEY` 只用于 devnet 给 staking 地址打币,Hoodi 上直接删除;`register` 的质押由 deployer 代付,已把额度从 1 ETH 降到 0.01 ETH。

### I-011 干跑脚本里 `printf '%x'` 对 1e20 溢出 — 已修

- **现象:** `printf: warning: 100000000000000000000: Numerical result out of range`,给分叉里的账户设置的余额被截成 2^64-1(约 18.4 ETH)。
- **原因:** bash 的 `printf` 只处理 64 位整数。
- **处理:** 改用 `cast to-hex`。对这次干跑没有影响(18.4 ETH 足够),只是余额不是原本想设的 100 ETH。

### I-012 分叉里的 gas 价格比我实测的低很多

- **现象:** 我 09:19 直接问 Hoodi 得到 `eth_gasPrice ≈ 1.08 gwei`,而干跑里 deployer 的实际 gas 开销折算出约 0.15 gwei。
- **原因(推测,未验证):** anvil 分叉按分叉点的区块 base fee 起算,而 Hoodi 的 gas 价格本身会波动。
- **处理:** 预算按总 gas(这个才是确定的)乘以多档价格,见 `notes/budget.md`。正式部署前再查一次实时价格。

### I-013 干跑后分叉被关闭,无法对链上状态做抽样核对

- **说明:** `dryrun-fork.sh` 结束时会停掉自己启动的 anvil,所以干跑后没法再 `cast call` 检查 Rollup / L1Sequencer 的状态(TODO P4-11)。
- **已有的保证:** 每一步都成功退出;`s11_verify_deploy` 检查了所有部署的合约都有代码;`s32` 检查了 L1Sequencer 的 owner 等于 deployer。
- **建议:** 正式部署后再做 P4-11 的抽样核对;或者另起一个**保持运行**的分叉,在上面继续做 P5~P7 的彩排(见下)。

### I-014 `tendermint testnet --n` 是**非验证者**数量,不是总数 — 已修

- **现象:** `tendermint testnet --v 1 --n 2 --hostname ... --hostname ...` 报 `testnet needs precisely 3 hostnames (number of validators plus non-validators)`。
- **原因:** 主机名个数 = `--v` + `--n`。上游用 `--v 1 --n 5` 配 6 个主机名,我误把 `--n` 当成总数。
- **处理:** 改成 `--v 1 --n 1`(1 个验证者 node0 + 1 个非验证者 node1)。
- **验证:** 用干跑留下的创世文件在 `run/nodes-test/` 里跑通 `scripts/setup-nodes.py`:两个 geth 初始化出的创世块哈希 `be3cbc..5540f9` 与干跑 `rollup.json` 里的 L2 创世哈希一致;node0 配置里 rpc/p2p/abci 为 26657/26656/26658,prometheus 为 26660,`block_sync=true`、`pex=false`、`persistent_peers=""`;node1 没有验证者密钥;创世里的验证者公钥与 node0 的 `priv_validator_key` 匹配。
- **注意:** `run/nodes-test/` 只是测试产物(用的是干跑的 L1 起点),**不能**用于正式运行;正式运行时用 `--run-dir run/l2`,并传入真实部署产出的 `genesis-l2.json`。

### I-015 三个进程的 metrics 端口默认值会互相冲突 — 已在 morph-native.sh 里分开

| 进程 | 默认 | 冲突对象 | 处理 |
|---|---|---|---|
| gas-oracle | `GAS_ORACLE_METRIC_ADDRESS` 默认 `0.0.0.0:6060` | morph-el-0 的 metrics 6060 | 设为 `127.0.0.1:6070` |
| morphnode(node-1,layer1 校验模式) | `MORPH_NODE_METRICS_PORT` 默认 26660 | node-0 的 tendermint prometheus 26660 | 设 `MORPH_NODE_METRICS_SERVER_ENABLE=true`、`MORPH_NODE_METRICS_PORT=27660` |
| tx-submitter | metrics 默认关闭(compose 里也没开) | — | 启用并设 `TX_SUBMITTER_METRICS_PORT=8060`,方便观测 |

### I-016 `scripts/morph-native.sh` 已写好并做了无副作用测试

- **已测:** 语法、无参数用法、`status`(无进程时)、部署前 `up` 会拒绝、未知进程名、从干跑的部署文件里读合约地址。
- **未测:** 真正启动进程。原因:本机上另一个 devnet 还占着 8545 和 8551(`status` 里 `el-0 :8545` 读到的是别的链的高度,恰好证实了这个冲突)。要等它停掉,且 L1 部署完成后才能端到端验证。

### I-017 deployer 换成用户自带的账户

- **背景:** 用户提供了自己的 deployer 私钥(不记录),在 Hoodi 上有 36 ETH。
- **提醒过:** 私钥最好不要经过对话。用户仍通过对话给出了私钥,所以该私钥已出现在对话记录里。它是测试网账户;**建议今后把它当作"已暴露"**,不要用它持有真实资产,也不要在主网复用。
- **处理:** 写入权限 600 的 `env/secrets.env`;核对推导出的地址与所给地址一致;把 `hoodidev.ts`(4 处)和 `hoodidev-deploy-config.json`(8 处)里的旧 deployer 地址全部换成新地址,重新暂存并导出补丁。补丁里 0 处私钥。
- **影响:** 所有合约的 owner 都是用户的这个账户。旧的生成账户 `0x0129…1A2c` 从未收到过币,作废。
- **从 deployer 向 submitter 转 1 ETH** 作为批次提交的运行资金(`SUBMITTER_FUND_ETH=1`)。

### I-018 真实部署第一次失败:公共 RPC 上"刚上链就读不到" — 已加固,重跑

- **现象:** 在真实 Hoodi 上部署的第 1 个合约(ProxyAdmin)已经上链(交易 `0xbbcab7c1…`),但紧接着的 `owner()` 调用返回空数据(`CALL_EXCEPTION, data="0x"`),部署脚本退出。anvil 分叉上同一套脚本完全正常。
- **验证:** 失败后立刻再查,这个合约有代码(2059 字节)、`owner()` 就是 deployer —— 合约本身没问题,是读请求早了一步。
- **原因(部分验证):** 部署代码对每个合约都是 `await contract.deployed()` 后立刻 `assertContractVariable(...)` 读状态。公共 RPC 背后是多个节点,确认上链的节点和接着处理读请求的节点可能不同。我探测到 `0xrpc.io` 的区块高度会在相邻两个值之间跳动(符合这个模型);ethpandaops 在我的 30 次快速探测里很稳,所以**这个机制是推断,不是对失败那一次的直接证明**。
- **处理:** 在 `contracts/src/deploy-utils.ts` 给两个断言函数(共 51 处调用)加**默认关闭**的重试:`MORPH_READ_RETRIES`(默认 0=原样)、`MORPH_READ_RETRY_DELAY_MS`(默认 2000)。`deploy-l1.sh` 默认设为 15 次、每次 2 秒。读不到或值不匹配都会重试,最后一次的结果原样抛出,断言报的仍是真实问题。
- **副作用:** 失败的那次留下了一个孤儿合约 `0x2637…EEF45`(ProxyAdmin),deployer nonce 为 1。重跑会重新部署一套,孤儿合约不用管,花费只有几美分级别的测试币。
- **局限:** 只覆盖断言函数里的读。`register` 步骤里还有几处别的读(如 `registered()`),如果在那里再遇到同类问题,需要同样处理。

### I-019 第二次真实部署失败 + 改用请求层的一致性代理 — 已重跑

- **现象:** 加了断言重试后再跑,失败点换到 `deploy/013-DeployProxys.ts:27/227` 的 `console.log(await IProxyContract.admin())`(一个调试打印)。这次已经部署了 9 笔交易(到 `Proxy__L1GatewayRouter`)。同一个根因的第二个读点;后面的 `register` 等步骤还有更多读点,逐个补不是办法。
- **处理:** 写了 `scripts/rpc-consistency-proxy.js`(约 100 行,无依赖),放在 hardhat / 创世工具 / cast 与公共 RPC 之间:
  1. 记住"已经见过的最高区块"(来自收据和 `eth_blockNumber`),叫 floor;
  2. 把所有对 `latest`/`pending` 的状态读改写成 floor 这个具体块号;
  3. 上游还没同步到时,等待并重试(500 毫秒一次,最多 30 秒);
  4. 写交易、显式指定块号的请求、其余方法原样透传。
- **一个必须验证才发现的细节:** 这个上游对"还没同步到的块"**不报错,而是返回 `result: null`**(`eth_call`、`getBalance`、`getCode`、`getTransactionCount`、`getStorageAt`、`getBlockByNumber` 都是;只有 `eth_estimateGas` 报 `header not found`)。我的第一版只认错误文字,等于没修。已改成:对被改写过的读,`null` 结果也视为"还没同步到"。
- **验证:** (1) 透传、改写、`eth_blockNumber` 单调递增、批量请求都正常;(2) 用测试钩子让代理以为自己见过 head+2:`owner()` 被拦住 17 秒(21 次重试)后返回正确的 deployer 地址,之后的读 276 毫秒;(3) 显式块号的请求立即返回,不受影响。
- **接入:** `deploy-l1.sh` 对 `https://` 的 RPC 默认启用(`USE_CONSISTENCY_PROXY=auto`),对本地 anvil 不启用;退出时自动停掉代理。
- **局限:** 这是针对我们实测到的行为(`null`/`header not found`)写的。换一个上游,其"未同步"的表现可能不同;代理只保证单调的读,不保证在重组(reorg)时的正确性,对 devnet 部署够用。
- **残留:** 前两次失败留下约 10 笔孤儿交易(几个孤儿合约),deployer 余额 35.99 ETH,影响可以忽略。

### I-020 第三次真实部署失败:上游返回 `rate-limit exceeded` — 代理加了退避与备用端点

- **现象:** 经一致性代理后部署推进了约 100 秒,在 `eth_call`(`admin()`)上收到 `ProviderError: rate-limit exceeded`,被 ethers 当作 `CALL_EXCEPTION` 抛出,部署退出。代理日志里没有一次读重试,说明这次不是读不一致,而是限流。
- **排查:** 单独用 `curl` 连续发请求:`eth_blockNumber`、`eth_call`、`eth_estimateGas`、`eth_getTransactionCount`、`eth_getTransactionReceipt`、`eth_getBlockByNumber` 各 150 个,加上另两个端点各 150 个 `eth_blockNumber`,**全部没有被限流**。所以限流不是按"持续请求数"或"某个方法"稳定触发的,更像是公共服务的**瞬时限流**(共享额度)。这是推断,没有对那一次的直接证明。
- **处理:** 代理新增:
  1. 限流(HTTP 429/503,或错误信息含 rate limit / too many requests / quota / exceeded)与网络错误时,**立刻切换到下一个上游**;所有上游都拒绝时,按 0.5→1→2→4→8 秒指数退避,最多 8 轮;
  2. 支持备用上游(逗号分隔);`deploy-l1.sh` 默认备用 `0xrpc.io` 和 `publicnode`;
  3. `/stats` 里按方法统计请求数、限流次数、故障切换次数,方便以后看流量构成。
- **验证(用本地假上游):** (A) 主上游一直限流、备用是真实 Hoodi:332 毫秒内完成故障切换,返回正确的链 ID 和余额;(B) 单上游前 3 次限流后恢复:退避 0.5+1+2 秒,总耗时 3568 毫秒后成功。
- **风险:** 多个上游的区块头可能略有先后。读请求已经被 floor 钉在具体块号上,落后的节点会返回 null 并被重试,所以这个风险被覆盖;发送交易可能去到任意一个上游,收据则由客户端轮询,最终一致。
- **残留:** 三次失败共留下约 20 笔孤儿交易;deployer 余额 35.98 ETH,约花了 0.017 ETH。

### I-021 第四次真实部署卡死在 `initialize`:一致性代理的 floor 冻结 — 代理的 bug,已修

- **现象:** `s10_deploy`(部署)一次通过,`s11`、`s20`(创世)、`s21` 也通过,随后 `s30_initialize` 在 "Upgrading the L1CrossDomainMessenger proxy..." 上卡了 18 分钟以上。进程几乎不占 CPU;代理统计里 `rateLimited=0`、`retries=0`、`eth_sendRawTransaction` 停在 38,而 `eth_call` 和 `eth_chainId` 在缓慢增加。
- **根因(已通过自测复现):** `deploy/015-MessengerInit.ts` 的写法是:发送 `upgradeToAndCall` 交易但**不等收据**,然后每 3 秒轮询 `implementation()` 等它变成新地址。我的代理把对 `latest` 的读钉在 floor 上,而 floor 只有在客户端调用 `eth_blockNumber` 或取到收据时才更新——这里两样都没有,floor 冻结在升级交易**之前**的区块,轮询永远读到旧状态。**这是我写的代理的 bug,不是 RPC 或链的问题。**
- **修复:**
  1. 每次处理被改写的读之前,最多每秒向上游问一次最新高度并更新 floor(`PROXY_HEAD_REFRESH_MS`,默认 1000);
  2. `pending` 的 nonce 单独处理:取"上游 pending"、"floor 区块的计数"、"代理自己记录的已发送交易下一个 nonce"三者的最大值,避免漏掉已发出但未打包的交易导致 nonce 重用;
  3. 新增可选节流 `PROXY_MIN_INTERVAL_MS`(默认 0=关),对应"限流时在请求之间加 sleep"的需求。
- **验证:** `scripts/rpc-consistency-proxy.selftest.js` 用一条每秒出块的假链复现了这个 bug:关闭刷新后轮询 12 秒仍读到旧值(FAIL);开启后在链越过变化块后 3 秒读到新值。四项全部通过(floor 跟随链、pending nonce、latest nonce 不被污染、显式块号透传)。
- **关于加 sleep:** 第四次运行里代理统计显示 `rateLimited=0`、`failovers=0`,**没有触发限流**,所以这次不需要在合约部署之间加 sleep。如果以后触发限流,打开 `PROXY_MIN_INTERVAL_MS`(例如 300)即可,不用改部署代码。
- **善后:** 卡住的那次运行目录保留在 `run/l1-attempt4-stuck/`。它已部署的合约和升级交易留在链上成为孤儿合约;因为 `initialize` 会重新部署实现合约、创世文件又嵌入了合约地址,不能原地续跑,所以整套重来。
- **另一个观察:** deployer 余额在此期间从 35.98 ETH 变成 68.93 ETH(不是我转的),看起来有人又转了币进来。

### I-022 `cast send` 不读 `ETH_PRIVATE_KEY`,且步骤运行器吞错误 — `setFirstSequencer` 一度没发出去,已修并核对

- **现象:** `s33_fund_submitter` 报 `Error accessing local wallet`;追查发现 `s32_first_sequencer` 其实**同样失败**,却在 0 秒内被标成 "ok"。链上 `getSequencer()` 仍是 `no sequencer configured`。
- **根因 1:** 这个 cast 版本(Foundry 1.9.0)的 `cast send` 不会读 `ETH_PRIVATE_KEY` 环境变量(之前我只在 `cast wallet address` 上测过,并错误地把结论推广到 `cast send`)。
- **根因 2:** `step()` 里 `if ! "${fn}" ...` 调用步骤函数。bash 在 `if` 条件上下文里**会关闭 `set -e`**,所以函数中间的命令失败不会中止,只有最后一条命令的状态决定结果;`s32` 的最后一条是 `echo`,于是"成功"。
- **修复:** (1) 新增 `scripts/send-tx.js`,用 ethers 发送交易,私钥只经环境变量 `KEY` 传入,不出现在命令行;(2) `step()` 改为 `set +e; ( set -e; fn ) ...; rc=$?; set -e`,让 errexit 在步骤内真正生效;(3) `s32`、`s33` 在步骤内**读回链上结果**再判成功。
- **教训:** 不能只看退出码。hardhat 任务出错时常常只打印一行然后 `return`,退出码仍是 0。所以新增 `scripts/verify-l1.sh`,完全不看步骤退出码,直接读链上状态与预期比对。
- **结果:** 修复后续跑,`s32`、`s33`、`s40` 全部成功;`verify-l1.sh` **53 项通过、0 失败**(Rollup/Messenger/MessageQueue/Submitter/L1Sequencer 的 owner 与互相引用、Rollup 的 L2 链 ID 与创世批次、submitter 登记与质押、所有 21 个代理的 implementation 与 admin、第一个序列器)。
- **善后:** 错误的 `s32` 标记保留为 `run/l1/state/s32_first_sequencer.done.WRONG-nothing-was-sent`。

### I-023 deployer 余额又增加了约 32 ETH

- 部署期间 deployer 余额从 68.93 ETH 变成 100.87 ETH(`deploy-l1.sh` 的 summary 里"spent"显示为负数,是因为用的是第一次尝试前记录的余额)。不是我转的,应该是你又转了币。仅记录。

### I-024 6060 端口被 Coder 工作区代理占用 — morph-geth 的 metrics 改到 6160/6161

- **现象:** 检查 Morph 要用的端口时,6060 被工作区自身的 agent 占用,不能动。所以 geth 的 metrics 端口改用 6160 起。
- **处理:** `morph-native.sh` 里 morph-geth 的 metrics 端口从 6060/6061 改为 6160/6161;PLAN 第 6 节端口表同步修改。
- **提醒:** 这也是 geth 的默认 metrics 端口,以后在这台机器上起 geth 类进程都要避开 6060。

### I-025 oracle 密钥没有权限更新 GasPriceOracle — 上游 devnet 里 owner 与 oracle 是同一个账户,我们拆开了

- **现象(启动 oracle 前发现):** L2 创世里 `gasPriceOracleOwner` 是用户的 deployer,而 oracle 进程用我另外生成的 `L2_GAS_ORACLE_PRIVATE_KEY`。链上读到 `GasPriceOracle.owner()` 是 deployer,oracle 密钥 `isAllowed=false`、L2 余额为 0。所有更新函数(`setL1BaseFee…`)是 `onlyAllowed`(owner 或白名单)。
- **原因:** 上游 devnet 里 `gasPriceOracleOwner` 和 `L2_GAS_ORACLE_PRIVATE_KEY` 是同一个 Anvil #0 账户,所以没人注意到这个耦合。
- **处理:** `scripts/l2-bootstrap.sh`:用 L2 创世里预置了余额的开发账户(公开的 Anvil #0 密钥,仅在这个只监听 127.0.0.1 的本机私有 L2 上使用)给 deployer 转 100 ETH、给 oracle 转 10 ETH,再由 deployer 调 `setAllowList([oracle],[true])`,读回 `isAllowed=true`。
- **结果:** oracle 随后第一次更新就成功(L2 交易 `0x7fa3…`)。

### I-026 node-1 对 L1 的扫描节奏只有每 15 秒 1~2 个块

- **现象:** node-1(layer1 校验模式)的日志里,每 15 秒拉取一次 L1 的 1~2 个区块(`derivation start pull rollupData form l1 startBlock=… end=…`)。Hoodi 每 12 秒一个块,即约 1.25 块/15 秒,而它每轮处理 1~2 块,追赶余量很小。
- **影响:** 批次在 L1 区块 4043361,node-1 在 4043352,约 2 分钟后才追上并还原出区块。不影响正确性,但**还原延迟随 L1 出块时间和这个轮询节奏而定**,不是实时的。
- **未做:** 没有去调轮询间隔或范围参数(未确认对应的 flag)。如果需要更低的还原延迟,这里是调优点。

### I-027 第一次 L1→L2 存款失败:L1 消息已被打包,但在 L2 上执行失败(状态 0)

- **现象:** `scripts/deposit-test.sh` 第一次运行:L1 存款交易在 Hoodi 区块 4043381 成功(状态 1,gas 241464),消息入队(`L1MessageQueue.nextCrossDomainMessageIndex = 1`)。但 L2 上全新接收地址 `0x50d3…831b` 的余额 10 分钟内始终为 0,脚本超时失败。
- **排查过程:**
  1. node-0 日志里只有 2 条 `failed to fetch L1 messages … block range extends beyond current head block`(见 I-026 之后加固代理的那一项),看起来像是读取问题;
  2. 但 node-0 的指标 `morphnode_syncer_message_count = 1`、`morphnode_executor_next_l1_message_queue_index = 1` 说明消息**已经被读到并被打进了区块**;
  3. 扫描 L2 全部区块,找到唯一一笔类型 `0x7e`(L1 消息)的交易,在 **L2 区块 342**:`to` 是 L2CrossDomainMessenger(`0x5300…0007`),`gas` 110000,**回执状态 0**,`gasUsed` 108798,0 条日志,调用的是 `relayMessage`。
- **结论(推断,待第二次存款验证):** 我沿用了仓库示例任务 `deposit-l1-eth` 里的 `_gasLimit = 110000`,在这套合约上 `relayMessage` 的整个调用链把 gas 用到 98.9% 就失败了,基本是 **gas 不足**。**不是 L1 消息没送到,也不是 L2 没打包。**
- **第一次存款的 0.01 ETH:** 已经锁在 L1 的网关里,L2 上没到账。这是测试币。Morph 的 L1 Messenger 有 `replayMessage`(可用更大的 gas 重放),本次没有去做重放。
- **另一个教训(我自己的):** 存款脚本我接了 `| cut`,输出被缓冲到脚本结束才显示,让我在等待的 10 分钟里看不到进展;已改成直接写文件。
- **对计划的影响:** `deposit-test.sh` 的默认 gas 上限改为 400000,可用 `GAS_LIMIT` 覆盖。第二次存款的结果见下一条。

#### I-027 结论(第二次存款,gas 上限 400000)— 已证实

- 在 Hoodi 区块 4043433 存入 0.01 ETH 到新地址 `0xA2ac…6357`(L1 gas 224386)。**node-0 上约 25 秒后到账**;**node-1 在 391 秒后到账**(它要等包含该 L2 区块的批次 4 提交到 L1,再从 blob 还原)。
- 所以第一次失败确实是 **`_gasLimit = 110000` 不足**,不是链路或配置问题。仓库示例任务里的 110000 在这套部署上不能用,存款至少给到几十万。
- 之后把 node-1 已有的全部区块(0 到 796)逐个与 node-0 比对哈希:**0 个不一致**。这个范围同时包含区块 342(含失败的 L1 消息交易)和成功存款所在的区块,说明两条路径下 node-1 还原的结果也和 node-0 一致。
- **未做:** 第一次失败存款的 0.01 ETH 的 `replayMessage` 重放(测试币,金额小)。

### I-028 出块节点 node-0 重启后永远不再出块 — 单个 tendermint 节点的拓扑不可重启(重启测试 T5 失败)

- **现象:** `scripts/restart-test.sh` 的 T5(停 15 秒再起 node-0):进程是 running、RPC 能答、日志没有任何错误,但 L2 高度冻结(7675),**链停了**。重启前的区块哈希不变(T5b 通过,没有分叉)。
- **日志特征(与 Morph `ops/README.md` 描述的"静默停转"一致):** `ABCI Handshake height=7675` → `Reactor waitSync=true` → `Already upgraded to sequencer mode, consensus reactor will not start height=3 upgradeHeight=2` → 最后一行是 `derivation skipped: sequencer mode`,**之后没有任何输出**,没有启动出块例程。
- **根因(已在代码里确认,不是猜测):**
  1. 出块例程只在 `blocksync/reactor.go` 里 `if bcR.pool.IsCaughtUp()` 之后才 `StartSequencerRoutines()`(没有任何配置能绕开);
  2. `blocksync/pool.go` 的 `IsCaughtUp()` 第一条是 `if len(pool.peers) == 0 { return false }`("Need at least 1 peer to be considered caught up");
  3. 我们的拓扑里 node-0 是**唯一**运行 tendermint 的节点(node-1 是 layer1 校验模式,不启动 tendermint),`persistent_peers` 为空,pool 没有对端。
  4. 首次启动能出块,是因为从高度 0 开始,靠单验证者先跑 PBFT 阶段、升级进入出块模式,不走这条路径;**重启后高度已超过升级高度,必须走 block pool,于是卡死。**
- **这是这个拓扑的固有限制,不是原生运行引入的:** Morph 的 `ops/README.md` 第 3 条前置条件写着"至少需要两个运行 tendermint 的节点,layer1 校验模式的节点不算",上游 docker 版的单序列器拓扑(node-0 + layer1 的 node-1)应有同样的限制。**我没有在 docker 版上验证过,这一句是按代码推断的。**
- **处理:** 新增 `scripts/add-node2.py`,在现有运行目录上追加一个标准全节点 **node-2(带 morph-el-2,没有出块密钥)**,它运行 tendermint,是 node-0 的 block pool 里永远存在的对端;node-0 的 `persistent_peers` 加入 node-2。**保留 node-1 不变**,因为它的"独立从 L1 blob 还原"验证我不想丢。端口:node-2 的 tendermint 26666/26667/26670,morph-el-2 的 8745/8746/8751/30305/6162。
- **此前的结论需要修正:** 我之前汇报"主线验证完成"时没有包含重启恢复,这正是当时列在"未验证"里的那一项,现在证明它确实有问题。

#### I-028 结论(加入 node-2 后重跑完整重启测试)

- `scripts/restart-test.sh`(T1~T6、T8)**19 项全部通过、0 失败**:
  - T1 `kill -9` node-1:5 秒内被守护外壳拉起,继续还原,tip 哈希一致;
  - T2 / T3 重启 morph-el-1 / node-1:继续还原,哈希一致;
  - T4 停 L1 代理 20 秒:L1 同步恢复、node-0 没停、所有进程都没退出;
  - **T5 重启出块节点 node-0:24 秒内恢复出块(原先永远停转),重启前的区块哈希不变,无分叉;**
  - T6 重启 morph-el-0:node-0 恢复出块;
  - T8 全部 down 再 up:高度从 8196 继续(不是从 0),历史区块哈希不变,所有进程在运行,node-1 tip 哈希一致。
- 恢复期间三个节点已有的 7563 个区块哈希逐个比对,0 个不一致:那次停转没有造成分叉。
- 测试脚本里一处文字("all 7 processes")在加入 node-2 后过时,检查本身用的是完整进程清单,已改文字。
- T7(重启 tx-submitter 并等下一批)单独跑,见下一条。

### 验证记录:L2→L1 提款(P7-5)通过

- `scripts/withdraw-test.sh`:L2 区块 8341 发起 0.005 ETH 提款(叶子下标 0,`from` 为 L2 的 ETH 网关,`to` 为 L1 的 ETH 网关);343 秒后 L1 Rollup 列出包含它的 withdrawal root;本地从 L2 的全部 `AppendMessage` 事件重建深度 32 的稀疏默克尔树,**重建的根与链上一致**;L1 `proveAndRelayMessage`(区块 4044651,状态 1,gas 167527)发出 `RelayedMessage`;新的 L1 接收地址余额 0.0 → 0.005 ETH。
- 证明算法的离线自测(`scripts/withdraw-proof.selftest.js`):逐行复刻合约的 `_appendMessageHash` / `getTreeRoot` / `verifyMerkleProof`,210 组(叶子数 1~20 × 每个下标)全部一致。
- 只验证了**一笔**提款(树里只有一个叶子)。多笔提款时证明的兄弟节点会包含真实叶子,算法已被离线自测覆盖,但链上还没有用多笔提款实测。
