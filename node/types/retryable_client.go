package types

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync/atomic"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/eth/catalyst"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/ethclient/authclient"
	"github.com/morph-l2/go-ethereum/rpc"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

const (
	ConnectionRefused       = "connection refused"
	EOFError                = "EOF"
	JWTStaleToken           = "stale token"
	JWTExpiredToken         = "token is expired"
	MinerClosed             = "miner closed"
	ExecutionAborted        = "execution aborted"
	Timeout                 = "timed out"
	DiscontinuousBlockError = "discontinuous block number"
)

// configResponse represents the eth_config RPC response (EIP-7910)
type configResponse struct {
	Current *forkConfig `json:"current"`
	Next    *forkConfig `json:"next"`
	Last    *forkConfig `json:"last"`
}

// forkConfig represents a single fork configuration
type forkConfig struct {
	ActivationTime  uint64                 `json:"activationTime"`
	ChainId         string                 `json:"chainId"`
	ForkId          string                 `json:"forkId"`
	Precompiles     map[string]string      `json:"precompiles"`
	SystemContracts map[string]string      `json:"systemContracts"`
	Morph           *morphExtension        `json:"morph,omitempty"`
}

// morphExtension contains Morph-specific configuration fields
type morphExtension struct {
	UseZktrie   bool    `json:"useZktrie"`
	MPTForkTime *uint64 `json:"mptForkTime,omitempty"`
}

// fetchMPTForkTime fetches the MPT fork time from geth via eth_config API
func fetchMPTForkTime(rpcURL string, logger tmlog.Logger) (uint64, error) {
	client, err := rpc.Dial(rpcURL)
	if err != nil {
		return 0, fmt.Errorf("failed to connect to geth: %w", err)
	}
	defer client.Close()

	var result json.RawMessage
	if err := client.Call(&result, "eth_config"); err != nil {
		return 0, fmt.Errorf("eth_config call failed: %w", err)
	}

	var resp configResponse
	if err := json.Unmarshal(result, &resp); err != nil {
		return 0, fmt.Errorf("failed to parse eth_config response: %w", err)
	}

	// Try to get mptForkTime from current config
	if resp.Current != nil && resp.Current.Morph != nil && resp.Current.Morph.MPTForkTime != nil {
		mptTime := *resp.Current.Morph.MPTForkTime
		logger.Info("Fetched MPT fork time from geth", "mptForkTime", mptTime, "source", "current")
		return mptTime, nil
	}

	// Fallback to next config
	if resp.Next != nil && resp.Next.Morph != nil && resp.Next.Morph.MPTForkTime != nil {
		mptTime := *resp.Next.Morph.MPTForkTime
		logger.Info("Fetched MPT fork time from geth", "mptForkTime", mptTime, "source", "next")
		return mptTime, nil
	}

	// Fallback to last config
	if resp.Last != nil && resp.Last.Morph != nil && resp.Last.Morph.MPTForkTime != nil {
		mptTime := *resp.Last.Morph.MPTForkTime
		logger.Info("Fetched MPT fork time from geth", "mptForkTime", mptTime, "source", "last")
		return mptTime, nil
	}

	logger.Info("MPT fork time not configured in geth, MPT switch disabled")
	return 0, nil // No MPT fork configured, return 0 (never switch)
}

type RetryableClient struct {
	legacyAuthClient *authclient.Client
	legacyEthClient  *ethclient.Client
	authClient       *authclient.Client
	ethClient        *ethclient.Client
	mptTime          uint64
	mpt              atomic.Bool
	b                backoff.BackOff
	logger           tmlog.Logger
}

// NewRetryableClient creates a new retryable client that fetches MPT fork time from geth.
// The legacyEthAddr is used to fetch the MPT fork time via eth_config API.
// Will retry calling the api, if the connection is refused.
func NewRetryableClient(legacyAuthClient *authclient.Client, legacyEthClient *ethclient.Client, authClient *authclient.Client, ethClient *ethclient.Client, legacyEthAddr string, logger tmlog.Logger) *RetryableClient {
	logger = logger.With("module", "retryClient")

	// Fetch MPT fork time from legacy geth via eth_config API
	mptTime, err := fetchMPTForkTime(legacyEthAddr, logger)
	if err != nil {
		logger.Error("Failed to fetch MPT fork time from geth, using 0 (MPT switch disabled)", "error", err)
		mptTime = 0
	}

	return &RetryableClient{
		legacyAuthClient: legacyAuthClient,
		legacyEthClient:  legacyEthClient,
		authClient:       authClient,
		ethClient:        ethClient,
		mptTime:          mptTime,
		b:                backoff.NewExponentialBackOff(),
		logger:           logger,
	}
}

func (rc *RetryableClient) aClient() *authclient.Client {
	if !rc.mpt.Load() {
		return rc.legacyAuthClient
	}
	return rc.authClient
}

func (rc *RetryableClient) eClient() *ethclient.Client {
	if !rc.mpt.Load() {
		return rc.legacyEthClient
	}
	return rc.ethClient
}

// EnsureSwitched checks if MPT switch time has been reached and switches to MPT client if needed.
// This should be called when the block is already delivered (e.g., synced via P2P) to ensure
// the client switch happens even if NewL2Block is not called.
func (rc *RetryableClient) EnsureSwitched(ctx context.Context, timeStamp uint64, number uint64) {
	rc.switchClient(ctx, timeStamp, number)
}

func (rc *RetryableClient) switchClient(ctx context.Context, timeStamp uint64, number uint64) {
	if rc.mpt.Load() {
		return
	}
	if timeStamp <= rc.mptTime {
		return
	}

	rc.logger.Info("========================================")
	rc.logger.Info("MPT UPGRADE: Switch time reached!")
	rc.logger.Info("========================================")
	rc.logger.Info("MPT switch time reached, switching from legacy client to MPT client",
		"mpt_time", rc.mptTime,
		"current_time", timeStamp,
		"target_block", number)
	rc.logger.Info("Current status: connected to LEGACY geth, waiting for target geth to sync...")

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	startTime := time.Now()
	lastLogTime := startTime

	for {
		remote, err := rc.ethClient.BlockNumber(ctx)
		if err != nil {
			rc.logger.Error("Failed to get target geth block number",
				"error", err,
				"hint", "Please ensure target geth is running and accessible")
			<-ticker.C
			continue
		}

		if remote+1 >= number {
			// Get target geth's latest block hash for debugging
			targetHeader, headerErr := rc.ethClient.HeaderByNumber(ctx, big.NewInt(int64(remote)))
			targetBlockHash := "unknown"
			targetStateRoot := "unknown"
			if headerErr == nil && targetHeader != nil {
				targetBlockHash = targetHeader.Hash().Hex()
				targetStateRoot = targetHeader.Root.Hex()
			}

			rc.mpt.Store(true)
			rc.logger.Info("========================================")
			rc.logger.Info("MPT UPGRADE: Successfully switched!")
			rc.logger.Info("========================================")
			rc.logger.Info("Successfully switched to MPT client",
				"remote_block", remote,
				"target_block", number,
				"target_block_hash", targetBlockHash,
				"target_state_root", targetStateRoot,
				"wait_duration", time.Since(startTime))
			return
		}

		if time.Since(lastLogTime) >= 5*time.Second {
			rc.logger.Error("!!! WAITING: Node BLOCKED waiting for target geth !!!",
				"target_geth_block", remote,
				"target_block", number,
				"blocks_behind", number-remote-1,
				"wait_duration", time.Since(startTime))
			lastLogTime = time.Now()
		}

		<-ticker.C
	}
}

func (rc *RetryableClient) AssembleL2Block(ctx context.Context, number *big.Int, transactions eth.Transactions) (ret *catalyst.ExecutableL2Data, err error) {
	timestamp := uint64(time.Now().Unix())
	if retryErr := backoff.Retry(func() error {
		rc.switchClient(ctx, timestamp, number.Uint64())
		resp, respErr := rc.aClient().AssembleL2Block(ctx, &timestamp, number, transactions)
		if respErr != nil {
			rc.logger.Info("failed to AssembleL2Block", "error", respErr)
			if retryableError(respErr) {
				return respErr
			}
			err = respErr // stop retrying and put this error to response error field, if the error is not connection related
		}
		ret = resp
		return nil
	}, rc.b); retryErr != nil {
		return nil, retryErr
	}
	return
}

func (rc *RetryableClient) ValidateL2Block(ctx context.Context, executableL2Data *catalyst.ExecutableL2Data) (ret bool, err error) {
	rc.switchClient(ctx, executableL2Data.Timestamp, executableL2Data.Number)
	if retryErr := backoff.Retry(func() error {
		resp, respErr := rc.aClient().ValidateL2Block(ctx, executableL2Data)
		if respErr != nil {
			rc.logger.Info("failed to ValidateL2Block", "error", respErr)
			if retryableError(respErr) {
				return respErr
			}
			err = respErr
		}
		ret = resp
		return nil
	}, rc.b); retryErr != nil {
		return false, retryErr
	}
	return
}

func (rc *RetryableClient) NewL2Block(ctx context.Context, executableL2Data *catalyst.ExecutableL2Data, batchHash *common.Hash) (err error) {
	rc.logger.Info("NewL2Block called",
		"block_number", executableL2Data.Number,
		"parent_hash", executableL2Data.ParentHash.Hex(),
		"state_root", executableL2Data.StateRoot.Hex(),
		"mpt_switched", rc.mpt.Load())

	rc.switchClient(ctx, executableL2Data.Timestamp, executableL2Data.Number)

	rc.logger.Info("After switchClient",
		"block_number", executableL2Data.Number,
		"mpt_switched", rc.mpt.Load())

	if retryErr := backoff.Retry(func() error {
		rc.logger.Info("Sending block to geth",
			"block_number", executableL2Data.Number,
			"using_mpt_client", rc.mpt.Load())

		respErr := rc.aClient().NewL2Block(ctx, executableL2Data, batchHash)
		if respErr != nil {
			rc.logger.Error("NewL2Block failed",
				"block_number", executableL2Data.Number,
				"error", respErr)
			if retryableError(respErr) {
				return respErr
			}
			err = respErr
		} else {
			rc.logger.Info("NewL2Block succeeded", "block_number", executableL2Data.Number)
		}
		return nil
	}, rc.b); retryErr != nil {
		return retryErr
	}
	return
}

func (rc *RetryableClient) NewSafeL2Block(ctx context.Context, safeL2Data *catalyst.SafeL2Data) (ret *eth.Header, err error) {
	rc.switchClient(ctx, safeL2Data.Timestamp, safeL2Data.Number)
	if retryErr := backoff.Retry(func() error {
		resp, respErr := rc.aClient().NewSafeL2Block(ctx, safeL2Data)
		if respErr != nil {
			rc.logger.Info("failed to NewSafeL2Block", "error", respErr)
			if retryableError(respErr) {
				return respErr
			}
			err = respErr
		}
		ret = resp
		return nil
	}, rc.b); retryErr != nil {
		return nil, retryErr
	}
	return
}

func (rc *RetryableClient) CommitBatch(ctx context.Context, batch *eth.RollupBatch, signatures []eth.BatchSignature) (err error) {
	if retryErr := backoff.Retry(func() error {
		respErr := rc.aClient().CommitBatch(ctx, batch, signatures)
		if respErr != nil {
			rc.logger.Info("failed to CommitBatch", "error", respErr)
			if retryableError(respErr) {
				return respErr
			}
			err = respErr
		}
		return nil
	}, rc.b); retryErr != nil {
		return retryErr
	}
	return
}

func (rc *RetryableClient) AppendBlsSignature(ctx context.Context, batchHash common.Hash, signature eth.BatchSignature) (err error) {
	if retryErr := backoff.Retry(func() error {
		respErr := rc.aClient().AppendBlsSignature(ctx, batchHash, signature)
		if respErr != nil {
			rc.logger.Info("failed to call AppendBlsSignature", "error", respErr)
			if retryableError(respErr) {
				return respErr
			}
			err = respErr
		}
		return nil
	}, rc.b); retryErr != nil {
		return retryErr
	}
	return
}

func (rc *RetryableClient) BlockNumber(ctx context.Context) (ret uint64, err error) {
	if retryErr := backoff.Retry(func() error {
		resp, respErr := rc.eClient().BlockNumber(ctx)
		if respErr != nil {
			rc.logger.Info("failed to call BlockNumber", "error", respErr)
			if retryableError(respErr) {
				return respErr
			}
			err = respErr
		}
		ret = resp
		return nil
	}, rc.b); retryErr != nil {
		return 0, retryErr
	}
	return
}

func (rc *RetryableClient) HeaderByNumber(ctx context.Context, blockNumber *big.Int) (ret *eth.Header, err error) {
	if retryErr := backoff.Retry(func() error {
		resp, respErr := rc.eClient().HeaderByNumber(ctx, blockNumber)
		if respErr != nil {
			rc.logger.Info("failed to call BlockNumber", "error", respErr)
			if retryableError(respErr) {
				return respErr
			}
			err = respErr
		}
		ret = resp
		return nil
	}, rc.b); retryErr != nil {
		return nil, retryErr
	}
	return
}

func (rc *RetryableClient) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) (ret []byte, err error) {
	if retryErr := backoff.Retry(func() error {
		resp, respErr := rc.eClient().CallContract(ctx, call, blockNumber)
		if respErr != nil {
			rc.logger.Info("failed to call eth_call", "error", respErr)
			if retryableError(respErr) {
				return respErr
			}
			err = respErr
		}
		ret = resp
		return nil
	}, rc.b); retryErr != nil {
		return nil, retryErr
	}
	return
}

func (rc *RetryableClient) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) (ret []byte, err error) {
	if retryErr := backoff.Retry(func() error {
		resp, respErr := rc.eClient().CodeAt(ctx, contract, blockNumber)
		if respErr != nil {
			rc.logger.Info("failed to call eth_getCode", "error", respErr)
			if retryableError(respErr) {
				return respErr
			}
			err = respErr
		}
		ret = resp
		return nil
	}, rc.b); retryErr != nil {
		return nil, retryErr
	}
	return
}

// currently we want every error retryable, except the DiscontinuousBlockError
func retryableError(err error) bool {
	// return strings.Contains(err.Error(), ConnectionRefused) ||
	// 	strings.Contains(err.Error(), EOFError) ||
	// 	strings.Contains(err.Error(), JWTStaleToken) ||
	// 	strings.Contains(err.Error(), JWTExpiredToken) ||
	// 	strings.Contains(err.Error(), MinerClosed) ||
	// 	strings.Contains(err.Error(), ExecutionAborted) ||
	// 	strings.Contains(err.Error(), Timeout)
	return !strings.Contains(err.Error(), DiscontinuousBlockError)
}
