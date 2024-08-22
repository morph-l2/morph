package types

import (
	"context"
	"math/big"
	"strings"

	"github.com/cenkalti/backoff/v4"
	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/eth/catalyst"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/ethclient/authclient"
	tmlog "github.com/tendermint/tendermint/libs/log"
)

const (
	ConnectionRefused = "connection refused"
	EOFError          = "EOF"
	JWTStaleToken     = "stale token"
	JWTExpiredToken   = "token is expired"
	MinerClosed       = "miner closed"
)

type RetryableClient struct {
	authClient *authclient.Client
	ethClient  *ethclient.Client
	b          backoff.BackOff
	logger     tmlog.Logger
}

// NewRetryableClient make the client retryable
// Will retry calling the api, if the connection is refused
func NewRetryableClient(authClient *authclient.Client, ethClient *ethclient.Client, logger tmlog.Logger) *RetryableClient {
	logger = logger.With("module", "retryClient")
	return &RetryableClient{
		authClient: authClient,
		ethClient:  ethClient,
		b:          backoff.NewExponentialBackOff(),
		logger:     logger,
	}
}

func (rc *RetryableClient) AssembleL2Block(ctx context.Context, number *big.Int, transactions eth.Transactions) (ret *catalyst.ExecutableL2Data, err error) {
	if retryErr := backoff.Retry(func() error {
		resp, respErr := rc.authClient.AssembleL2Block(ctx, number, transactions)
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

func (rc *RetryableClient) NewL2Block(ctx context.Context, executableL2Data *catalyst.ExecutableL2Data, batchHash *common.Hash) (err error) {
	if retryErr := backoff.Retry(func() error {
		respErr := rc.authClient.NewL2Block(ctx, executableL2Data, batchHash)
		if respErr != nil {
			rc.logger.Info("failed to NewL2Block", "error", respErr)
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

func (rc *RetryableClient) CommitBatch(ctx context.Context, batch *eth.RollupBatch, signatures []eth.BatchSignature) (err error) {
	if retryErr := backoff.Retry(func() error {
		respErr := rc.authClient.CommitBatch(ctx, batch, signatures)
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
		respErr := rc.authClient.AppendBlsSignature(ctx, batchHash, signature)
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

func retryableError(err error) bool {
	return strings.Contains(err.Error(), ConnectionRefused) ||
		strings.Contains(err.Error(), EOFError) ||
		strings.Contains(err.Error(), JWTStaleToken) ||
		strings.Contains(err.Error(), JWTExpiredToken) ||
		strings.Contains(err.Error(), MinerClosed)

}
