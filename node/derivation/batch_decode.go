package derivation

import (
	"bytes"
	"encoding/binary"
	"math/big"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/rlp"
)

type BatchData struct {
	Txs           []*types.Transaction
	BlockContexts []*BlockInfo
	//Signature     *bindings.RollupBatchSignature
}

// number || timestamp || base_fee || gas_limit || num_txs || tx_hashs
type BlockInfo struct {
	Number    *big.Int
	Timestamp uint64
	BaseFee   *big.Int
	GasLimit  uint64
	NumTxs    uint16
}

// decode blockcontext
func (b *BatchData) DecodeBlockContext(endBlock uint64, bs []byte) error {
	b.BlockContexts = []*BlockInfo{}
	// [block1, block2, ..., blockN]
	reader := bytes.NewReader(bs)
	for {
		block := new(BlockInfo)
		// number || timestamp || base_fee || gas_limit || num_txs
		// Number(8) || Timestamp(8) || BaseFee(32) || GasLimit(8) || numTxs(2)
		bsBlockNumber := make([]byte, 8)
		if _, err := reader.Read(bsBlockNumber[:]); err != nil {
			return err
		}
		block.Number = new(big.Int).SetBytes(bsBlockNumber)

		if err := binary.Read(reader, binary.BigEndian, &block.Timestamp); err != nil {
			return err
		}
		// [32]byte uint256
		bsBaseFee := make([]byte, 32)
		if _, err := reader.Read(bsBaseFee[:]); err != nil {
			return err
		}
		block.BaseFee = new(big.Int).SetBytes(bsBaseFee)
		if err := binary.Read(reader, binary.BigEndian, &block.GasLimit); err != nil {
			return err
		}
		if err := binary.Read(reader, binary.BigEndian, &block.NumTxs); err != nil {
			return err
		}
		for i := 0; i < int(block.NumTxs); i++ {
			txHash := common.Hash{}
			if _, err := reader.Read(txHash[:]); err != nil {
				return err
			}
			// drop txHash
		}
		b.BlockContexts = append(b.BlockContexts, block)
		if block.Number.Uint64() == endBlock {
			break
		}
	}
	return nil
}

func (b *BatchData) DecodeTransactions(bs []byte) error {
	return rlp.DecodeBytes(bs, &b.Txs)
}
