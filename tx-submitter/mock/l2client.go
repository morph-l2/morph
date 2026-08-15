package mock

import (
	"context"
	"math/big"
	"time"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/eth"
)

// L2ClientWrapper is a mock implementation of iface.L2Client
type L2ClientWrapper struct {
	CallContractErr    error
	CallContractResult []byte
	BlockByNumberErr   error
	Block              *types.Block
	Balance            *big.Int
	TipCap             *big.Int
	reorgDepth         int64
	reorgCount         int64
	sequencerSet       []common.Address
	BaseFee            *big.Int
}

func NewL2ClientWrapper() *L2ClientWrapper {
	// Create mock CallContractResult for epoch update time
	// Return a timestamp from 1 hour ago
	timestamp := time.Now().Add(-1 * time.Hour).Unix()
	result := make([]byte, 32)
	big.NewInt(timestamp).FillBytes(result)

	// Create a mock sequencer set with two addresses
	mockAddr1 := common.HexToAddress("0x1111111111111111111111111111111111111111")
	mockAddr2 := common.HexToAddress("0x2222222222222222222222222222222222222222")
	sequencerSet := []common.Address{mockAddr1, mockAddr2}

	return &L2ClientWrapper{
		TipCap:             big.NewInt(1e9),
		Balance:            big.NewInt(1e18),
		CallContractResult: result,
		sequencerSet:       sequencerSet,
		BaseFee:            big.NewInt(1e9),
	}
}

// CallContract implements the CallContract method
func (l *L2ClientWrapper) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	if l.CallContractErr != nil {
		return nil, l.CallContractErr
	}

	// If calling the sequencer set method, return the mock sequencer set
	if len(msg.Data) >= 4 {
		methodID := msg.Data[:4]
		// Method ID for GetSequencerSet2()
		if string(methodID) == "\x77\xd7\xdf\xfb" {
			// Encode the sequencer set into bytes
			result := make([]byte, 0)
			// First 32 bytes for offset
			offset := make([]byte, 32)
			big.NewInt(32).FillBytes(offset)
			result = append(result, offset...)
			// Next 32 bytes for length
			length := make([]byte, 32)
			big.NewInt(int64(len(l.sequencerSet))).FillBytes(length)
			result = append(result, length...)
			// Then the addresses
			for _, addr := range l.sequencerSet {
				addrBytes := make([]byte, 32)
				copy(addrBytes[12:], addr.Bytes())
				result = append(result, addrBytes...)
			}
			return result, nil
		}
	}

	return l.CallContractResult, nil
}

// BlockByNumber implements the BlockByNumber method
func (l *L2ClientWrapper) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	if l.BlockByNumberErr != nil {
		return nil, l.BlockByNumberErr
	}
	if l.Block != nil {
		return l.Block, nil
	}
	return nil, ethereum.NotFound
}

// HeaderByNumber implements the HeaderByNumber method
func (l *L2ClientWrapper) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	if l.Block != nil {
		return l.Block.Header(), nil
	}
	return nil, ethereum.NotFound
}

// BlockNumber implements the BlockNumber method
func (l *L2ClientWrapper) BlockNumber(ctx context.Context) (uint64, error) {
	if l.Block != nil {
		return l.Block.NumberU64(), nil
	}
	return 0, nil
}

// GetReorgDepth returns the mock reorg depth
func (l *L2ClientWrapper) GetReorgDepth() int64 {
	return l.reorgDepth
}

// GetReorgCount returns the mock reorg count
func (l *L2ClientWrapper) GetReorgCount() int64 {
	return l.reorgCount
}

// SetReorgDepth sets the mock reorg depth
func (l *L2ClientWrapper) SetReorgDepth(depth int64) {
	l.reorgDepth = depth
}

// SetReorgCount sets the mock reorg count
func (l *L2ClientWrapper) SetReorgCount(count int64) {
	l.reorgCount = count
}

// BalanceAt implements the BalanceAt method
func (l *L2ClientWrapper) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	if l.Balance != nil {
		return l.Balance, nil
	}
	return big.NewInt(0), nil
}

// TransactionByHash implements the TransactionByHash method
func (l *L2ClientWrapper) TransactionByHash(ctx context.Context, hash common.Hash) (*types.Transaction, bool, error) {
	return nil, false, ethereum.NotFound
}

// NonceAt implements the NonceAt method
func (l *L2ClientWrapper) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	return 0, nil
}

// TransactionReceipt implements the TransactionReceipt method
func (l *L2ClientWrapper) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	return nil, ethereum.NotFound
}

// CodeAt implements bind.ContractCaller
func (l *L2ClientWrapper) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	return nil, nil
}

// PendingCodeAt implements bind.ContractCaller
func (l *L2ClientWrapper) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	return nil, nil
}

// PendingNonceAt implements bind.ContractTransactor
func (l *L2ClientWrapper) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return 0, nil
}

// SuggestGasPrice implements bind.ContractTransactor
func (l *L2ClientWrapper) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return big.NewInt(1e9), nil
}

// EstimateGas implements bind.ContractTransactor
func (l *L2ClientWrapper) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	return 21000, nil
}

// SendTransaction implements bind.ContractTransactor
func (l *L2ClientWrapper) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return nil
}

// FilterLogs implements bind.ContractFilterer
func (l *L2ClientWrapper) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return nil, nil
}

// SubscribeFilterLogs implements bind.ContractFilterer
func (l *L2ClientWrapper) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return nil, nil
}

// GetBlockTraceByNumber implements L2Client
func (l *L2ClientWrapper) GetBlockTraceByNumber(ctx context.Context, number *big.Int) (*types.BlockTrace, error) {
	return nil, nil
}

// GetRollupBatchByIndex implements L2Client
func (l *L2ClientWrapper) GetRollupBatchByIndex(ctx context.Context, batchIndex uint64) (*eth.RPCRollupBatch, error) {
	return nil, nil
}

// SuggestGasTipCap implements the SuggestGasTipCap method
func (l *L2ClientWrapper) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	if l.TipCap != nil {
		return l.TipCap, nil
	}
	return big.NewInt(1e9), nil
}

// GetSequencerSet2 implements IL2Sequencer
func (m *L2ClientWrapper) GetSequencerSet2() ([]common.Address, error) {
	return m.sequencerSet, nil
}
