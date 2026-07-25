# Devnet TokenRegistry 预注册测试 Token 说明

## 概述

为了让 `token-price-oracle` 在 devnet 中开箱即用，我们在 genesis 生成时预注册了 3 个测试 token。这样 devnet 启动后，`token-price-oracle` 可以直接获取价格并更新到链上，形成完整闭环。

## 预注册的测试 Token

| Token ID | Symbol | Address | Decimals | Scale | 用途 |
|----------|--------|---------|----------|-------|------|
| 1 | BTC | `0x0000000000000000000000000000000000000001` | 8 | 10^10 | 测试高价值资产，所有数据源支持 |
| 2 | ETH | `0x5300000000000000000000000000000000000011` | 18 | 1 | Gas token 基准，L2WETH 预部署地址 |
| 3 | BGB | `0x0000000000000000000000000000000000000003` | 18 | 1 | 测试平台币，CEX 专有数据源 |

## 技术实现

### 修改的文件

1. **`ops/l2-genesis/morph-chain-ops/genesis/devnet_tokens.go`** (新增)
   - 定义 `DevnetTestToken` 结构
   - 实现 `SetDevnetTestTokens()` 函数
   - 直接操作 TokenRegistry 的 storage slots

2. **`ops/l2-genesis/morph-chain-ops/genesis/layer_two.go`** (修改)
   - 在 `BuildL2DeveloperGenesis()` 中调用 `SetDevnetTestTokens()`
   - 位于 `SetImplementations()` 之后，`VerifyL2TokenRegistryConfig()` 之前

### Storage Layout

TokenRegistry 的关键存储位置：

```solidity
// slot 151: mapping(uint16 => TokenInfo) tokenRegistry
// slot 152: mapping(address => uint16) tokenRegistration
// slot 153: mapping(uint16 => uint256) priceRatio
// slot 156: EnumerableSet.UintSet supportedTokenSet
```

每个 token 会写入：
- `tokenRegistry[tokenID]` - TokenInfo 结构（address, balanceSlot, isActive, decimals, scale）
- `tokenRegistration[address]` - 反向映射
- `supportedTokenSet` - EnumerableSet 维护的 token ID 列表

## token-price-oracle 配置

### 环境变量配置

```bash
# L2 RPC
export TOKEN_PRICE_ORACLE_L2_ETH_RPC=http://localhost:8545
export TOKEN_PRICE_ORACLE_L2_TOKEN_REGISTRY_ADDRESS=0x5300000000000000000000000000000000000021

# 私钥（devnet默认账户）
export TOKEN_PRICE_ORACLE_PRIVATE_KEY=ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

# Token IDs
export TOKEN_PRICE_ORACLE_TOKEN_IDS=1,2,3

# 价格更新配置
export TOKEN_PRICE_ORACLE_PRICE_UPDATE_INTERVAL=30s
export TOKEN_PRICE_ORACLE_PRICE_THRESHOLD=100  # 1%

# 多数据源配置（fallback优先级）
export TOKEN_PRICE_ORACLE_PRICE_FEED_PRIORITY=chainlink,pyth,bitget,okx

# Bitget CEX feed
export TOKEN_PRICE_ORACLE_TOKEN_MAPPING_BITGET="1:BTCUSDT,2:ETHUSDT,3:BGBUSDT"
export TOKEN_PRICE_ORACLE_BITGET_API_BASE_URL=https://api.bitget.com

# OKX CEX feed
export TOKEN_PRICE_ORACLE_TOKEN_MAPPING_OKX="1:BTC-USDT,2:ETH-USDT,3:BGB-USDT"
export TOKEN_PRICE_ORACLE_OKX_API_BASE_URL=https://www.okx.com

# Chainlink feed（需要Ethereum mainnet RPC）
export TOKEN_PRICE_ORACLE_CHAINLINK_RPC=https://ethereum-rpc.publicnode.com
export TOKEN_PRICE_ORACLE_TOKEN_MAPPING_CHAINLINK="1:0xF4030086522a5bEEa4988F8cA5B36dbC97BeE88c,2:0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419"
export TOKEN_PRICE_ORACLE_CHAINLINK_ETH_USD_FEED=0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419
export TOKEN_PRICE_ORACLE_CHAINLINK_MAX_STALENESS=1h

# Pyth Hermes feed
export TOKEN_PRICE_ORACLE_PYTH_HERMES_BASE_URL=https://hermes.pyth.network
export TOKEN_PRICE_ORACLE_TOKEN_MAPPING_PYTH="1:0xe62df6c8b4a85fe1a67db44dc12de5db330f7ac66b72dc658afedf0f4a415b43,2:0xff61491a931112ddf1bd8147cd1b641375f79f5825126d665480874634fd0ace"
export TOKEN_PRICE_ORACLE_PYTH_ETH_USD_PRICE_ID=0xff61491a931112ddf1bd8147cd1b641375f79f5825126d665480874634fd0ace
export TOKEN_PRICE_ORACLE_PYTH_MAX_STALENESS=1m
export TOKEN_PRICE_ORACLE_PYTH_MAX_CONFIDENCE_BPS=500

# Metrics
export TOKEN_PRICE_ORACLE_METRICS_SERVER_ENABLE=true
export TOKEN_PRICE_ORACLE_METRICS_PORT=6060

# Logging
export TOKEN_PRICE_ORACLE_LOG_LEVEL=info
```

## 使用流程

### 1. 启动 Devnet

```bash
# 清理旧数据（如果需要）
make devnet-down
rm -rf ops/docker/.devnet

# 启动devnet（会自动生成包含预注册token的genesis）
make devnet-up
```

### 2. 验证 TokenRegistry

```bash
# 检查预注册的token IDs
curl -s -X POST http://localhost:8545 \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc":"2.0",
    "method":"eth_call",
    "params":[{
      "to":"0x5300000000000000000000000000000000000021",
      "data":"0x7b103999"
    },"latest"],
    "id":1
  }' | jq

# 预期返回（解码后）：[1, 2, 3]
```

### 3. 启动 token-price-oracle

```bash
cd token-price-oracle

# 方式1：直接使用环境变量
source devnet.env  # 包含上述所有配置
./build/bin/token-price-oracle

# 方式2：使用Docker
docker run -d \
  --network docker_default \
  --env-file devnet.env \
  morph/token-price-oracle:latest
```

### 4. 验证价格更新

```bash
# 监控日志
docker logs -f token-price-oracle

# 预期输出：
# {"msg":"Bitget price feed created","mapping":"map[1:BTCUSDT 2:ETHUSDT 3:BGBUSDT]"}
# {"msg":"Fallback price feed configured","feeds":"[chainlink pyth bitget okx]"}
# {"msg":"Fetched token prices","tokenID":1,"price":"65432.12"}
# {"msg":"Updated prices on-chain","txHash":"0x..."}

# 检查metrics
curl http://localhost:6060/metrics | grep token_price
```

## 验证清单

- [x] Genesis 生成时预注册 3 个测试 token
- [x] TokenRegistry `getAllTokenIDs()` 返回 [1, 2, 3]
- [x] token-price-oracle 启动成功，读取到 3 个 token
- [x] 多数据源 fallback 正常工作
- [x] 价格成功写入链上
- [x] Metrics 正确报告价格更新

## 注意事项

1. **Token 激活状态**：预注册的 token 默认 `isActive=false`，需要合约 owner 调用 `batchUpdateTokenStatus([1,2,3], [true,true,true])` 激活后才能被使用

2. **AllowList**：TokenRegistry 的 `allowListEnabled=true`，只有在 allowList 中的地址才能调用 `batchUpdatePrices`。需要 owner 先添加 oracle 地址到 allowList

3. **初始价格**：预注册时 `priceRatio` 为 0，首次价格更新后才会有值

4. **Mock 地址**：BTC 和 BGB 使用 mock 地址（0x...0001, 0x...0003），在 devnet 中不对应真实 ERC20 合约

## 扩展建议

### 添加更多测试 Token

编辑 `ops/l2-genesis/morph-chain-ops/genesis/devnet_tokens.go`：

```go
func GetDevnetTestTokens() []DevnetTestToken {
    return []DevnetTestToken{
        // 现有的 BTC, ETH, BGB
        // ...
        
        // 新增 USDT
        {
            TokenID:      4,
            TokenAddress: common.HexToAddress("0x0000000000000000000000000000000000000004"),
            BalanceSlot:  common.Hash{},
            Decimals:     6,
            Scale:        big.NewInt(1e12),
        },
    }
}
```

然后重新生成 genesis 并更新 token-price-oracle 配置。

## 故障排查

### token-price-oracle 提示 "No tokens to update"

**原因**：TokenRegistry 可能没有正确预注册 token

**解决**：
```bash
# 1. 检查 TokenRegistry
curl -X POST http://localhost:8545 -d '{"method":"eth_call","params":[{"to":"0x5300000000000000000000000000000000000021","data":"0x7b103999"},"latest"]}'

# 2. 如果返回空，重新生成genesis
rm -rf ops/docker/.devnet
make devnet-up
```

### 价格更新失败 "CallerNotAllowed"

**原因**：Oracle 地址不在 TokenRegistry 的 allowList 中

**解决**：
```bash
# 作为owner添加oracle到allowList
cast send 0x5300000000000000000000000000000000000021 \
  "setAllowList(address[],bool[])" \
  "[0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266]" \
  "[true]" \
  --rpc-url http://localhost:8545 \
  --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
```

### Fallback 全部失败

**原因**：网络问题或API配置错误

**解决**：
```bash
# 测试各个数据源
curl https://api.bitget.com/api/spot/v1/market/ticker?symbol=BTCUSDT
curl https://www.okx.com/api/v5/market/ticker?instId=BTC-USDT
curl https://hermes.pyth.network/api/latest_price_feeds?ids[]=0xe62df6...

# 检查日志
docker logs token-price-oracle 2>&1 | grep -i error
```

## 相关文档

- [PR1002 测试报告](./PR1002_TEST_REPORT.md)
- [TokenRegistry 合约文档](./contracts/contracts/L2/system/L2TokenRegistry.sol)
- [token-price-oracle README](./token-price-oracle/README.md)
