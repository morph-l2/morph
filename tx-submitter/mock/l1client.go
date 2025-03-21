package mock

import (
	"context"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/ethclient"
)

// L1ClientWrapper is a mock implementation of the L1 client interface
type L1ClientWrapper struct {
	*ethclient.Client
	Block              *types.Block
	TipCap             *big.Int
	SendTxErr          error
	MockReorg          bool
	MockReorgDepth     uint64
	CallContractErr    error
	CallContractResult []byte
	CodeAtErr          error
	CodeAtResult       []byte
	BaseFee            *big.Int
	Header             *types.Header
	// Track transactions and their status
	transactions map[common.Hash]*types.Transaction
	receipts     map[common.Hash]*types.Receipt
	// Track block history for reorg simulation
	blockHistory map[uint64]*types.Block
	// Track original block hashes before reorg
	originalHashes map[uint64]common.Hash
}

// FilterLogs implements iface.Client.
func (l *L1ClientWrapper) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	panic("unimplemented")
}

// NonceAt implements iface.Client.
func (l *L1ClientWrapper) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	panic("unimplemented")
}

// PendingCodeAt implements iface.Client.
func (l *L1ClientWrapper) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	panic("unimplemented")
}

// SubscribeFilterLogs implements iface.Client.
func (l *L1ClientWrapper) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	panic("unimplemented")
}

// SuggestGasPrice implements iface.Client.
func (l *L1ClientWrapper) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	panic("unimplemented")
}

// NewL1ClientWrapper creates a new mock L1 client wrapper
func NewL1ClientWrapper() *L1ClientWrapper {
	header := &types.Header{
		Number:        big.NewInt(100),
		Time:          uint64(1234567890),
		BaseFee:       big.NewInt(1e9),
		ExcessBlobGas: new(uint64),
	}
	*header.ExcessBlobGas = 0
	block := types.NewBlockWithHeader(header)
	return &L1ClientWrapper{
		Client:         nil,
		Block:          block,
		TipCap:         big.NewInt(1e9),
		BaseFee:        big.NewInt(1e9),
		Header:         header,
		transactions:   make(map[common.Hash]*types.Transaction),
		receipts:       make(map[common.Hash]*types.Receipt),
		blockHistory:   make(map[uint64]*types.Block),
		originalHashes: make(map[uint64]common.Hash),
	}
}

// BlockNumber returns the current block number
func (l *L1ClientWrapper) BlockNumber(ctx context.Context) (uint64, error) {
	if l.Block != nil {
		return l.Block.NumberU64(), nil
	}
	return 0, nil
}

// HeaderByNumber returns the block header at the given number
func (l *L1ClientWrapper) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	if l.Header != nil {
		if l.Header.ExcessBlobGas == nil {
			l.Header.ExcessBlobGas = new(uint64)
			*l.Header.ExcessBlobGas = 0
		}
		return l.Header, nil
	}
	return nil, ethereum.NotFound
}

// BlockByNumber returns the block at the given number
func (l *L1ClientWrapper) BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error) {
	if number == nil {
		return l.Block, nil
	}

	blockNum := number.Uint64()
	if l.MockReorg && blockNum > l.MockReorgDepth {
		// If we haven't stored the original hash, store it
		if _, ok := l.originalHashes[blockNum]; !ok {
			if block, ok := l.blockHistory[blockNum]; ok {
				l.originalHashes[blockNum] = block.Hash()
			}
		}

		// Return a block with a different hash to simulate reorg
		header := &types.Header{
			Number:        number,
			Time:          uint64(1234567890),
			ParentHash:    common.HexToHash("0x" + fmt.Sprintf("%x", blockNum+1)),
			BaseFee:       big.NewInt(1e9),
			ExcessBlobGas: new(uint64),
		}
		*header.ExcessBlobGas = 0
		block := types.NewBlockWithHeader(header)
		l.blockHistory[blockNum] = block
		l.SimulateReorg()
		return block, nil
	}

	if block, ok := l.blockHistory[blockNum]; ok {
		return block, nil
	}

	// For blocks before reorg depth, return blocks with consistent hashes
	header := &types.Header{
		Number:        number,
		Time:          uint64(1234567890),
		ParentHash:    common.HexToHash("0x" + fmt.Sprintf("%x", blockNum)),
		BaseFee:       big.NewInt(1e9),
		ExcessBlobGas: new(uint64),
	}
	*header.ExcessBlobGas = 0
	block := types.NewBlockWithHeader(header)
	l.blockHistory[blockNum] = block
	return block, nil
}

func (l *L1ClientWrapper) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	if l.CallContractErr != nil {
		return nil, l.CallContractErr
	}
	return l.CallContractResult, nil
}

func (l *L1ClientWrapper) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	if l.MockReorg {
		// During reorg, all transactions should be considered missing
		return nil, ethereum.NotFound
	}
	if receipt, ok := l.receipts[txHash]; ok {
		return receipt, nil
	}
	return nil, ethereum.NotFound
}

func (l *L1ClientWrapper) TransactionByHash(ctx context.Context, hash common.Hash) (tx *types.Transaction, isPending bool, err error) {
	if l.MockReorg {
		// During reorg, all transactions should be considered missing
		return nil, false, ethereum.NotFound
	}
	if tx, ok := l.transactions[hash]; ok {
		// If there's no receipt, the transaction is pending
		// If there's a receipt, the transaction is not pending
		_, hasReceipt := l.receipts[hash]
		return tx, !hasReceipt, nil
	}
	return nil, false, ethereum.NotFound
}

func (l *L1ClientWrapper) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	return 0, nil
}

func (l *L1ClientWrapper) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	return l.TipCap, nil
}

func (l *L1ClientWrapper) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return 21000, nil
}

func (l *L1ClientWrapper) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	if l.SendTxErr != nil {
		return l.SendTxErr
	}
	l.transactions[tx.Hash()] = tx
	if !l.MockReorg {
		// Only create a receipt if we're not in reorg mode
		l.receipts[tx.Hash()] = &types.Receipt{
			TxHash:      tx.Hash(),
			BlockNumber: l.Block.Number(),
			Status:      1, // Success
		}
	}
	return nil
}

func (l *L1ClientWrapper) BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return big.NewInt(1e18), nil
}

// CodeAt mocks the CodeAt method
func (l *L1ClientWrapper) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	if l.CodeAtErr != nil {
		return nil, l.CodeAtErr
	}
	return l.CodeAtResult, nil
}

// SimulateReorg simulates a reorg by clearing transaction receipts
func (l *L1ClientWrapper) SimulateReorg() {
	l.receipts = make(map[common.Hash]*types.Receipt)
}

func (l *L1ClientWrapper) AddTx(tx *types.Transaction) {
	l.transactions[tx.Hash()] = tx
}
func (l *L1ClientWrapper) AddReceipt(receipt *types.Receipt) {
	l.receipts[receipt.TxHash] = receipt
}

// ClearReceipts removes all receipts to simulate reorg
func (l *L1ClientWrapper) ClearReceipts() {
	l.receipts = make(map[common.Hash]*types.Receipt)
}
