package node

import (
	"morph-l2/node/types"

	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
)

// GenesisBatchHeader builds the batch header committed to L1 alongside the L2 genesis block.
// It is invoked by ops/l2-genesis when bootstrapping a fresh chain and is the only producer
// of batch data on this side after the sequencer batch path was removed.
func GenesisBatchHeader(genesisHeader *eth.Header) (types.BatchHeaderV0, error) {
	wb := types.WrappedBlock{
		ParentHash:  genesisHeader.ParentHash,
		Miner:       genesisHeader.Coinbase,
		Number:      genesisHeader.Number.Uint64(),
		GasLimit:    genesisHeader.GasLimit,
		BaseFee:     genesisHeader.BaseFee,
		Timestamp:   genesisHeader.Time,
		StateRoot:   genesisHeader.Root,
		GasUsed:     genesisHeader.GasUsed,
		ReceiptRoot: genesisHeader.ReceiptHash,
	}
	blockContext := wb.BlockContextBytes(0, 0)
	batchData := types.NewBatchData()
	batchData.Append(blockContext, nil)

	return types.BatchHeaderV0{
		BatchIndex:           0,
		L1MessagePopped:      0,
		TotalL1MessagePopped: 0,
		DataHash:             batchData.DataHash(),
		BlobVersionedHash:    types.EmptyVersionedHash,
		PostStateRoot:        genesisHeader.Root,
		ParentBatchHash:      common.Hash{},
	}, nil
}
