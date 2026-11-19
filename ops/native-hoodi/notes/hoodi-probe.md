# Hoodi 公共端点探测结果(P2-1、P2-2)

探测时间:2026-11-18。脚本:`scripts/probe-hoodi.sh`(并发 20 × 3 轮,只读,对公共服务只做小规模并发)。原始输出:`notes/hoodi-probe-raw.txt`。

## 限流

| 端点 | 类型 | 60 个请求的结果 | 平均延迟 |
|---|---|---|---|
| `https://0xrpc.io/hoodi` | EL | 全部 200 | 0.92 秒 |
| `https://rpc.hoodi.ethpandaops.io` | EL | 全部 200 | 0.39 秒 |
| `https://beacon.hoodi.ethpandaops.io` | 信标 | 全部 200 | 0.46 秒 |
| `https://ethereum-hoodi-beacon-api.publicnode.com` | 信标 | 全部 200 | 0.06 秒 |

**这个并发量下没有任何限流。** 更高并发或长时间持续的上限没有测。

## eth_getLogs 区间上限

| 端点 | 1000 | 5000 | 10000 | 50000 |
|---|---|---|---|---|
| 0xrpc.io | ok | ok | 报 `invalid block params` | 报 `invalid block params` |
| ethpandaops | ok | ok | ok | 报 `exceeded max allowed range` |

**Morph 的 L1 同步每次只拉 100 个块**(`node/sync/config.go` 的 `DefaultFetchBlockRange = 100`),远低于两者的上限,不受影响。

## 信标 blob_sidecars

两个信标端点对近期 slot 都能返回 sidecar。第一次连取 5 个 slot 时各有 1 个失败,第二次复查 5 个全部成功,推测是某个 slot 没有区块或没有 blob,属偶发。对 Morph 的影响:取 blob 失败会按序回退到下一个端点,见 `node/derivation/beacon.go`。

## 结论

- EL 主用 **ethpandaops**(延迟低、getLogs 上限高);0xrpc.io 作备用。**这与最初计划里"主用 0xrpc.io"相反**,`env/hoodi.env.example` 与 `env/secrets.env` 同步调整。
- 信标 morphnode 配两个(ethpandaops 在前、publicnode 在后);gas-oracle 只能配一个。
