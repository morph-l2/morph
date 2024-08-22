package sync

import (
	"context"
	"fmt"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/rpc"
	tmlog "github.com/tendermint/tendermint/libs/log"

	"morph-l2/bindings/bindings"
	nodecommon "morph-l2/node/common"
	"morph-l2/node/types"
)

type BridgeClient struct {
	l1Client           *ethclient.Client
	filter             *bindings.L1MessageQueueWithGasPriceOracleFilterer
	morphPortalAddress common.Address
	confirmations      rpc.BlockNumber
	logger             tmlog.Logger
}

func NewBridgeClient(l1Client *ethclient.Client, l1MessageQueueAddress common.Address, confirmations rpc.BlockNumber, logger tmlog.Logger) (*BridgeClient, error) {
	logger = logger.With("module", "bridge")
	filter, err := bindings.NewL1MessageQueueWithGasPriceOracleFilterer(l1MessageQueueAddress, l1Client)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize NewMorphPortalFilterer, err = %w", err)
	}
	return &BridgeClient{
		l1Client:           l1Client,
		filter:             filter,
		morphPortalAddress: l1MessageQueueAddress,
		confirmations:      confirmations,
		logger:             logger,
	}, nil
}

func (c *BridgeClient) L1Messages(ctx context.Context, from, to uint64) ([]types.L1Message, error) {
	opts := bind.FilterOpts{
		Start:   from,
		End:     &to,
		Context: ctx,
	}
	it, err := c.filter.FilterQueueTransaction(&opts, nil, nil)
	if err != nil {
		return nil, err
	}

	txs := make([]types.L1Message, 0)
	for it.Next() {
		event := it.Event
		c.logger.Debug("Received new L1 QueueTransaction event", "event", event)

		if !event.GasLimit.IsUint64() {
			return nil, fmt.Errorf("invalid QueueTransaction event: QueueIndex = %v, GasLimit = %v", event.QueueIndex, event.GasLimit)
		}
		txs = append(txs, types.L1Message{
			L1MessageTx: L1MessageTxFromEvent(event),
			L1TxHash:    event.Raw.TxHash,
		})
	}
	return txs, nil
}

func (c *BridgeClient) L1MessagesFromTxHash(ctx context.Context, txHash common.Hash) ([]types.L1Message, error) {
	receipt, err := c.l1Client.TransactionReceipt(ctx, txHash)
	if err != nil {
		return nil, err
	}
	latestConfirmed, err := c.getLatestConfirmedBlockNumber(ctx)
	if err != nil {
		return nil, err
	}
	if receipt.BlockNumber.Uint64() > latestConfirmed {
		c.logger.Error("the target block has not been considered to be confirmed", "latestConfirmedHeight", latestConfirmed, "receiptAtBlockHeight", receipt.BlockNumber.Uint64())
		return nil, types.ErrNotConfirmedBlock
	}

	return c.deriveFromReceipt([]*eth.Receipt{receipt})
}

func (c *BridgeClient) getLatestConfirmedBlockNumber(ctx context.Context) (uint64, error) {
	return nodecommon.GetLatestConfirmedBlockNumber(ctx, c.l1Client, c.confirmations)
}
