package types

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/eth/catalyst"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/ethclient/authclient"
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

	GethRetryMaxElapsedTime = 30 * time.Minute
)

type RetryableClient struct {
	authClient *authclient.Client
	ethClient  *ethclient.Client
	b          backoff.BackOff
	logger     tmlog.Logger
}

func NewRetryableClient(authClient *authclient.Client, ethClient *ethclient.Client, logger tmlog.Logger) *RetryableClient {
	logger = logger.With("module", "retryClient")
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = GethRetryMaxElapsedTime
	return &RetryableClient{
		authClient: authClient,
		ethClient:  ethClient,
		b:          bo,
		logger:     logger,
	}
}

func (rc *RetryableClient) AssembleL2Block(ctx context.Context, number *big.Int, transactions eth.Transactions) (ret *catalyst.ExecutableL2Data, err error) {
	timestamp := uint64(time.Now().Unix())
	if retryErr := backoff.Retry(func() error {
		resp, respErr := rc.authClient.AssembleL2Block(ctx, &timestamp, number, transactions)
		if respErr != nil {
			rc.logger.Info("failed to AssembleL2Block", "error", respErr)
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

func (rc *RetryableClient) ValidateL2Block(ctx context.Context, executableL2Data *catalyst.ExecutableL2Data) (ret bool, err error) {
	if retryErr := backoff.Retry(func() error {
		resp, respErr := rc.authClient.ValidateL2Block(ctx, executableL2Data)
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

func (rc *RetryableClient) NewL2Block(ctx context.Context, executableL2Data *catalyst.ExecutableL2Data) (err error) {
	if retryErr := backoff.Retry(func() error {
		respErr := rc.authClient.NewL2Block(ctx, executableL2Data)
		if respErr != nil {
			rc.logger.Error("NewL2Block failed",
				"block_number", executableL2Data.Number,
				"error", respErr)
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

func (rc *RetryableClient) NewSafeL2Block(ctx context.Context, safeL2Data *catalyst.SafeL2Data) (ret *eth.Header, err error) {
	if retryErr := backoff.Retry(func() error {
		resp, respErr := rc.authClient.NewSafeL2Block(ctx, safeL2Data)
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

func (rc *RetryableClient) BlockNumber(ctx context.Context) (ret uint64, err error) {
	if retryErr := backoff.Retry(func() error {
		resp, respErr := rc.ethClient.BlockNumber(ctx)
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
		resp, respErr := rc.ethClient.HeaderByNumber(ctx, blockNumber)
		if respErr != nil {
			rc.logger.Info("failed to call HeaderByNumber", "error", respErr)
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

func (rc *RetryableClient) BlockByNumber(ctx context.Context, blockNumber *big.Int) (ret *eth.Block, err error) {
	if retryErr := backoff.Retry(func() error {
		resp, respErr := rc.ethClient.BlockByNumber(ctx, blockNumber)
		if respErr != nil {
			rc.logger.Info("failed to call BlockByNumber", "error", respErr)
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
		resp, respErr := rc.ethClient.CallContract(ctx, call, blockNumber)
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
		resp, respErr := rc.ethClient.CodeAt(ctx, contract, blockNumber)
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

func (rc *RetryableClient) SetBlockTags(ctx context.Context, safeBlockHash common.Hash, finalizedBlockHash common.Hash) (err error) {
	if retryErr := backoff.Retry(func() error {
		respErr := rc.authClient.SetBlockTags(ctx, safeBlockHash, finalizedBlockHash)
		if respErr != nil {
			rc.logger.Info("failed to call SetBlockTags", "error", respErr)
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

// currently we want every error retryable, except the DiscontinuousBlockError
func retryableError(err error) bool {
	return !strings.Contains(err.Error(), DiscontinuousBlockError)
}

// ============================================================================
// L2NodeV2 methods for sequencer mode
// ============================================================================

// AssembleL2BlockV2 assembles a L2 block based on parent hash.
func (rc *RetryableClient) AssembleL2BlockV2(ctx context.Context, parentHash common.Hash, transactions eth.Transactions) (ret *catalyst.ExecutableL2Data, err error) {
	timestamp := uint64(time.Now().Unix())
	if retryErr := backoff.Retry(func() error {
		resp, respErr := rc.authClient.AssembleL2BlockV2(ctx, parentHash, &timestamp, transactions)
		if respErr != nil {
			rc.logger.Info("failed to AssembleL2BlockV2", "error", respErr)
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
