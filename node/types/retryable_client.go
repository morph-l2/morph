package types

import (
	"context"
	"github.com/cenkalti/backoff/v4"
	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/eth/catalyst"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/ethclient/authclient"
	tmlog "github.com/tendermint/tendermint/libs/log"
	"math/big"
	"strings"
	"time"
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

type RetryableClient struct {
	legacyAuthClient *authclient.Client
	legacyEthClient  *ethclient.Client
	authClient       *authclient.Client
	ethClient        *ethclient.Client
	mptTime          uint64 // TODO rename
	b                backoff.BackOff
	logger           tmlog.Logger
}

// NewRetryableClient make the client retryable
// Will retry calling the api, if the connection is refused
func NewRetryableClient(legacyAuthClient *authclient.Client, legacyEthClient *ethclient.Client, authClient *authclient.Client, ethClient *ethclient.Client, logger tmlog.Logger) *RetryableClient {
	logger = logger.With("module", "retryClient")
	return &RetryableClient{
		legacyAuthClient: legacyAuthClient,
		legacyEthClient:  legacyEthClient,
		authClient:       authClient,
		ethClient:        ethClient,
		b:                backoff.NewExponentialBackOff(),
		logger:           logger,
	}
}

func (c *RetryableClient) aClient(timeStamp uint64) *authclient.Client {
	if c.mptTime >= timeStamp {
		return c.legacyAuthClient
	}
	return c.authClient
}

func (c *RetryableClient) eClient(timeStamp uint64) *ethclient.Client {
	if c.mptTime >= timeStamp {
		return c.legacyEthClient
	}
	return c.ethClient
}

func (rc *RetryableClient) AssembleL2Block(ctx context.Context, number *big.Int, transactions eth.Transactions) (ret *catalyst.ExecutableL2Data, err error) {
	timestamp := uint64(time.Now().Unix())
	if retryErr := backoff.Retry(func() error {
		resp, respErr := rc.aClient(timestamp).AssembleL2Block(ctx, &timestamp, number, transactions)
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
		resp, respErr := rc.aClient(executableL2Data.Timestamp).ValidateL2Block(ctx, executableL2Data)
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
		respErr := rc.aClient(executableL2Data.Timestamp).NewL2Block(ctx, executableL2Data, batchHash)
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
		resp, respErr := rc.aClient(safeL2Data.Timestamp).NewSafeL2Block(ctx, safeL2Data)
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
		// TODO timestamp
		respErr := rc.aClient(0).CommitBatch(ctx, batch, signatures)
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
		// TODO timestamp
		respErr := rc.aClient(0).AppendBlsSignature(ctx, batchHash, signature)
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
		// TODO timestamp
		resp, respErr := rc.eClient(0).BlockNumber(ctx)
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
		// TODO timestamp
		resp, respErr := rc.eClient(0).HeaderByNumber(ctx, blockNumber)
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
		// TODO timestamp
		resp, respErr := rc.eClient(0).CallContract(ctx, call, blockNumber)
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
		// TODO timestamp
		resp, respErr := rc.eClient(0).CodeAt(ctx, contract, blockNumber)
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
