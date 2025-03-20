package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/params"
)

const (
	GENESIS_TIMESTAMP = uint64(1695902400) // Holesky 创世区块时间戳
	SECONDS_PER_SLOT  = uint64(12)         // 每个时隙的秒数
)

// HTTP 接口
type HTTP interface {
	Get(ctx context.Context, path string, headers http.Header) (*http.Response, error)
}

// 基本的HTTP客户端实现
type BasicHTTPClient struct {
	endpoint string
	client   *http.Client
}

func NewBasicHTTPClient(endpoint string) *BasicHTTPClient {
	// 确保endpoint以斜杠结尾
	trimmedEndpoint := strings.TrimSuffix(endpoint, "/") + "/"
	return &BasicHTTPClient{
		endpoint: trimmedEndpoint,
		client:   &http.Client{Timeout: 30 * time.Second},
	}
}

func (cl *BasicHTTPClient) Get(ctx context.Context, p string, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, cl.endpoint+p, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: 构建请求失败", err)
	}
	for k, values := range headers {
		for _, v := range values {
			req.Header.Add(k, v)
		}
	}
	return cl.client.Do(req)
}

// L1BeaconClient
type L1BeaconClient struct {
	cl         HTTP
	timeToSlot func(timestamp uint64) (uint64, error)
}

// 创建新的beacon客户端
func NewL1BeaconClient(cl HTTP) *L1BeaconClient {
	// 为Holesky网络设置固定的genesis时间戳
	timeToSlotFn := func(timestamp uint64) (uint64, error) {
		if timestamp < GENESIS_TIMESTAMP {
			return 0, fmt.Errorf("提供的时间戳(%v)早于创世时间戳(%v)", timestamp, GENESIS_TIMESTAMP)
		}
		return (timestamp - GENESIS_TIMESTAMP) / SECONDS_PER_SLOT, nil
	}

	return &L1BeaconClient{cl: cl, timeToSlot: timeToSlotFn}
}

func (cl *L1BeaconClient) apiReq(ctx context.Context, dest any, method string) error {
	headers := http.Header{}
	headers.Add("Accept", "application/json")
	resp, err := cl.cl.Get(ctx, method, headers)
	if err != nil {
		return fmt.Errorf("%w: http Get失败", err)
	}
	if resp.StatusCode != http.StatusOK {
		errMsg, _ := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		return fmt.Errorf("请求失败，状态码 %d: %s", resp.StatusCode, string(errMsg))
	}
	if err := json.NewDecoder(resp.Body).Decode(dest); err != nil {
		_ = resp.Body.Close()
		return err
	}
	if err := resp.Body.Close(); err != nil {
		return fmt.Errorf("%w: 关闭响应体失败", err)
	}
	return nil
}

// BlobSidecar 结构体
type BlobSidecar struct {
	BlockRoot     []byte      `json:"block_root"`
	Index         json.Number `json:"index"`
	Slot          json.Number `json:"slot"`
	Blob          string      `json:"blob"`
	KZGCommitment []byte      `json:"kzg_commitment"`
	KZGProof      []byte      `json:"kzg_proof"`
}

// API响应
type APIGetBlobSidecarsResponse struct {
	Data []*BlobSidecar `json:"data"`
}

// 索引blob哈希
type IndexedBlobHash struct {
	Index uint64      // 在区块中的绝对索引
	Hash  common.Hash // blob的哈希，用于一致性检查
}

// KZGToVersionedHash 将KZG承诺转换为带版本的哈希
func KZGToVersionedHash(commitment []byte) common.Hash {
	// EIP-4844规范:
	//  def kzg_to_versioned_hash(commitment: KZGCommitment) -> VersionedHash:
	//    return VERSIONED_HASH_VERSION_KZG + sha256(commitment)[1:]
	var out common.Hash
	h := sha256.New()
	h.Write(commitment)
	copy(out[:], h.Sum(nil))
	out[0] = params.BlobTxHashVersion
	return out
}

// GetBlobSidecarsEnhanced 获取blob sidecars的增强方法
func (cl *L1BeaconClient) GetBlobSidecarsEnhanced(ctx context.Context, blockTime uint64, hashes []IndexedBlobHash) ([]*BlobSidecar, error) {
	// 计算slot
	slot, err := cl.timeToSlot(blockTime)
	if err != nil {
		return nil, fmt.Errorf("计算slot失败: %w", err)
	}

	// 首先尝试使用带索引的URL
	if len(hashes) > 0 {
		builder := strings.Builder{}
		builder.WriteString("eth/v1/beacon/blob_sidecars/")
		builder.WriteString(strconv.FormatUint(slot, 10))
		builder.WriteRune('?')
		v := url.Values{}

		for i := range hashes {
			v.Add("indices", strconv.FormatUint(hashes[i].Index, 10))
		}
		builder.WriteString(v.Encode())

		var resp APIGetBlobSidecarsResponse
		if err := cl.apiReq(ctx, &resp, builder.String()); err == nil && len(resp.Data) > 0 {
			return resp.Data, nil
		}
	}

	// 如果第一种方法失败或者没有指定索引，尝试第二种方法
	method := fmt.Sprintf("eth/v1/beacon/blob_sidecars/%d", slot)
	var blobResp APIGetBlobSidecarsResponse
	if err := cl.apiReq(ctx, &blobResp, method); err != nil {
		return nil, fmt.Errorf("请求blob sidecars失败: %w", err)
	}

	return blobResp.Data, nil
}

func main() {
	// 设置日志
	logger := log.New(os.Stdout, "[BlobTest] ", log.LstdFlags|log.Lshortfile)
	logger.Println("开始增强版blob获取测试")

	// 创建上下文
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// 节点URL和交易哈希
	nodeUrl := "https://omniscient-serene-sanctuary.ethereum-holesky.quiknode.pro/37083bb58e379efb02453b2f07fe4853faf18d03/"
	txHashHex := "0x2ada86d083a7b1af5e59c916327ef2255ccb9068834bbf13efa5527ba411718f"

	// 创建ETH客户端
	ethClient, err := ethclient.Dial(nodeUrl)
	if err != nil {
		logger.Fatalf("连接ETH客户端失败: %v", err)
	}

	// 获取交易
	txHash := common.HexToHash(txHashHex)
	tx, isPending, err := ethClient.TransactionByHash(ctx, txHash)
	if err != nil {
		logger.Fatalf("获取交易失败: %v", err)
	}
	if isPending {
		logger.Fatalf("交易仍处于挂起状态")
	}

	// 获取交易收据，确定区块号
	receipt, err := ethClient.TransactionReceipt(ctx, txHash)
	if err != nil {
		logger.Fatalf("获取交易收据失败: %v", err)
	}

	blockNumber := receipt.BlockNumber.Uint64()
	logger.Printf("交易位于区块: %d", blockNumber)

	// 获取区块
	block, err := ethClient.BlockByNumber(ctx, receipt.BlockNumber)
	if err != nil {
		logger.Fatalf("获取区块失败: %v", err)
	}

	blockTime := block.Time()
	logger.Printf("区块时间戳: %d", blockTime)

	// 获取blob哈希
	blobHashes := tx.BlobHashes()
	if len(blobHashes) == 0 {
		logger.Fatalf("交易不包含blob")
	}
	logger.Printf("交易包含 %d 个blob", len(blobHashes))

	for i, hash := range blobHashes {
		logger.Printf("Blob #%d 哈希: %s", i, hash.Hex())
	}

	// 创建beacon客户端
	beaconEndpoint := "https://beacon.holesky.ethpandaops.io"
	beaconClient := NewL1BeaconClient(NewBasicHTTPClient(beaconEndpoint))

	// 构建索引哈希数组
	var indexedHashes []IndexedBlobHash
	for i, hash := range blobHashes {
		indexedHashes = append(indexedHashes, IndexedBlobHash{
			Index: uint64(i),
			Hash:  hash,
		})
	}

	// 调用GetBlobSidecarsEnhanced
	blobSidecars, err := beaconClient.GetBlobSidecarsEnhanced(ctx, blockTime, indexedHashes)
	if err != nil {
		logger.Fatalf("获取blob sidecars失败: %v", err)
	}

	if len(blobSidecars) == 0 {
		logger.Fatalf("未返回blob sidecars")
	}

	logger.Printf("成功获取 %d 个blob sidecar", len(blobSidecars))

	// 检查每个blob sidecar
	matchedCount := 0
	for i, sidecar := range blobSidecars {
		// 从索引获取slot
		slotNum, _ := sidecar.Slot.Int64()
		indexNum, _ := sidecar.Index.Int64()

		logger.Printf("Blob Sidecar #%d: Slot=%d, Index=%d", i, slotNum, indexNum)

		// 如果有承诺数据，验证与哈希的匹配
		if len(sidecar.KZGCommitment) > 0 {
			versionedHash := KZGToVersionedHash(sidecar.KZGCommitment)
			logger.Printf("  KZG承诺转换为哈希: %s", versionedHash.Hex())

			// 检查是否匹配任何期望的哈希
			for _, expectedHash := range blobHashes {
				if bytes.Equal(versionedHash[:], expectedHash[:]) {
					matchedCount++
					logger.Printf("  ✅ 匹配blob哈希: %s", expectedHash.Hex())

					// 解码并处理blob数据
					blobData, err := hexutil.Decode(sidecar.Blob)
					if err != nil {
						logger.Printf("  ❌ 解码blob数据失败: %v", err)
						continue
					}

					// 打印blob数据的前200字节作为预览
					previewLen := 200
					if len(blobData) < previewLen {
						previewLen = len(blobData)
					}
					logger.Printf("  Blob数据预览 (%d字节): 0x%x...", len(blobData), blobData[:previewLen])
					break
				}
			}
		}
	}

	logger.Printf("总结: 找到 %d 个匹配的blob (共期望 %d 个)", matchedCount, len(blobHashes))
	logger.Println("测试完成")
}
