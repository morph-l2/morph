package types

import (
	"context"
	"errors"
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

// NewL2BlockV2 wraps engine_newL2BlockV2 — the reorg-capable variant of
// NewSafeL2Block introduced by go-ethereum PR #325. Unlike NewSafeL2Block
// (which requires parent == currentHead), V2 only requires the parent to
// exist on-chain; SetCanonical detects parentHash != currentHead and
// triggers EL forkchoice reorg automatically. With isSafe=true the EL
// skips verifyBlock + ValidateState (used for L1-confirmed blocks where
// the caller already trusts the block's content).
//
// Used by SPEC-005 §4.3 Path B self-heal: when local-rebuild produces a
// versioned hash that disagrees with L1, we pull the real blob, derive
// the batch using the true sequencer payload, and rewrite the locally
// divergent unsafe blocks via this API.
//
// Temporary note: the upstream PR https://github.com/morph-l2/go-ethereum/pull/325
// is still open. Once merged into main and the morph go-ethereum
// dependency is bumped to a release that contains the merged commit,
// no caller change is needed.
func (rc *RetryableClient) NewL2BlockV2(ctx context.Context, executableL2Data *catalyst.ExecutableL2Data, isSafe bool) (err error) {
	if retryErr := backoff.Retry(func() error {
		respErr := rc.authClient.NewL2BlockV2(ctx, executableL2Data, isSafe)
		if respErr != nil {
			rc.logger.Error("NewL2BlockV2 failed",
				"block_number", executableL2Data.Number,
				"parent_hash", executableL2Data.ParentHash,
				"is_safe", isSafe,
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
			if retryableError(respErr) {
				rc.logger.Info("failed to call HeaderByNumber, will retry", "error", respErr)
				return respErr
			}
			rc.logger.Error("failed to call HeaderByNumber, non-retryable", "error", respErr)
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
			if retryableError(respErr) {
				rc.logger.Info("failed to call BlockByNumber, will retry", "error", respErr)
				return respErr
			}
			rc.logger.Error("failed to call BlockByNumber, non-retryable", "error", respErr)
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
// retryableError reports whether an RPC error should trigger an exponential
// backoff retry inside RetryableClient. Errors not classified as retryable
// escape immediately so callers see the failure on the first poll cycle
// rather than after the 30-minute MaxElapsedTime budget runs out.
//
// Permanent classifications (do NOT retry):
//   - ethereum.NotFound: target block / header doesn't exist locally. With
//     SPEC-005 Path B reading L2 blocks the sequencer hasn't yet sealed
//     locally (snapshot too old, sync still catching up), this is a "wait
//     for sync" condition, not a transient RPC blip; retrying every
//     backoff tick for 30 minutes wastes the cycle and hides the gap from
//     the operator. The caller (e.g. verify_path_b) surfaces the missing
//     block, derivation logs an Error, and the next poll re-evaluates.
//   - DiscontinuousBlockError: structurally invalid input that no amount
//     of retry will fix.
func retryableError(err error) bool {
	if errors.Is(err, ethereum.NotFound) {
		return false
	}
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
