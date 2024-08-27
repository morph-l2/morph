package node

import (
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/eth/catalyst"

	"morph-l2/node/types"
)

type BlockConverter interface {
	Separate(l2Block *catalyst.ExecutableL2Data, l1Msg []types.L1Message) (blsMsg []byte, restMsg []byte, err error)
	Recover(blsMsg []byte, restMsg []byte, txs [][]byte) (l2Block *catalyst.ExecutableL2Data, l1Message []types.L1Message, err error)
}

type Version1Converter struct{}

func (bc *Version1Converter) Separate(l2Block *catalyst.ExecutableL2Data, l1Msg []types.L1Message) ([]byte, []byte, error) {
	// Number(8) || Timestamp(8) || BaseFee(32) || GasLimit(8) || numTxs(2)
	blsBytes := make([]byte, 58)
	copy(blsBytes[:8], types.Uint64ToBigEndianBytes(l2Block.Number))
	copy(blsBytes[8:16], types.Uint64ToBigEndianBytes(l2Block.Timestamp))
	if l2Block.BaseFee != nil {
		copy(blsBytes[16:48], l2Block.BaseFee.FillBytes(make([]byte, 32)))
	} else {
		copy(blsBytes[16:48], make([]byte, 32))
	}
	copy(blsBytes[48:56], types.Uint64ToBigEndianBytes(l2Block.GasLimit))
	copy(blsBytes[56:58], types.Uint16ToBigEndianBytes(uint16(len(l2Block.Transactions))))

	// bls signing context includes the tx hash
	for i, txBz := range l2Block.Transactions {
		var tx eth.Transaction
		if err := tx.UnmarshalBinary(txBz); err != nil {
			return nil, nil, fmt.Errorf("transaction %d is not valid: %v", i, err)
		}
		blsBytes = append(blsBytes, tx.Hash().Bytes()...)
	}

	rest := types.RestMessage{
		NonBLSMessage: types.NonBLSMessage{
			StateRoot:   l2Block.StateRoot,
			GasUsed:     l2Block.GasUsed,
			ReceiptRoot: l2Block.ReceiptRoot,
			LogsBloom:   l2Block.LogsBloom,
			L1Messages:  l1Msg,
		},
		BlockHash:          l2Block.Hash,
		ParentHash:         l2Block.ParentHash,
		Miner:              l2Block.Miner,
		NextL1MessageIndex: l2Block.NextL1MessageIndex,
	}
	restBytes, err := rest.MarshalBinary()
	if err != nil {
		return nil, nil, err
	}
	return blsBytes, restBytes, nil
}

func (bc *Version1Converter) Recover(blsMsg []byte, restMsg []byte, txs [][]byte) (*catalyst.ExecutableL2Data, []types.L1Message, error) {
	expectedBlsMsgLength := 58 + len(txs)*common.HashLength
	if len(blsMsg) != expectedBlsMsgLength {
		return nil, nil, fmt.Errorf("wrong blsMsg size, expected: %d, actual: %d", expectedBlsMsgLength, len(blsMsg))
	}
	rest := new(types.RestMessage)
	if err := rest.UnmarshalBinary(restMsg); err != nil {
		return nil, nil, err
	}
	if binary.BigEndian.Uint16(blsMsg[56:58]) != uint16(len(txs)) {
		return nil, nil, fmt.Errorf("wrong blsMsg, numTxs(%d) is not equal to the length of txs(%d)", binary.BigEndian.Uint16(blsMsg[56:58]), len(txs))
	}

	baseFee := new(big.Int).SetBytes(blsMsg[16:48])
	if baseFee.Cmp(big.NewInt(0)) == 0 {
		baseFee = nil
	}

	return &catalyst.ExecutableL2Data{
		Hash:               rest.BlockHash,
		ParentHash:         rest.ParentHash,
		Miner:              rest.Miner,
		Number:             binary.BigEndian.Uint64(blsMsg[:8]),
		Timestamp:          binary.BigEndian.Uint64(blsMsg[8:16]),
		BaseFee:            baseFee,
		GasLimit:           binary.BigEndian.Uint64(blsMsg[48:56]),
		StateRoot:          rest.StateRoot,
		GasUsed:            rest.GasUsed,
		ReceiptRoot:        rest.ReceiptRoot,
		LogsBloom:          rest.LogsBloom,
		NextL1MessageIndex: rest.NextL1MessageIndex,
		Transactions:       txs,
	}, rest.L1Messages, nil
}

type Version2Converter struct{}

func (bc *Version2Converter) Separate(l2Block *catalyst.ExecutableL2Data, l1Msg []types.L1Message) (blsMsg []byte, restMsg []byte, err error) {
	bm := &types.BLSMessage{
		ParentHash: l2Block.ParentHash,
		Miner:      l2Block.Miner,
		Number:     l2Block.Number,
		GasLimit:   l2Block.GasLimit,
		BaseFee:    l2Block.BaseFee,
		Timestamp:  l2Block.Timestamp,
	}
	if blsMsg, err = bm.MarshalBinary(); err != nil {
		return
	}
	nbm := &types.NonBLSMessage{
		StateRoot:   l2Block.StateRoot,
		GasUsed:     l2Block.GasUsed,
		ReceiptRoot: l2Block.ReceiptRoot,
		LogsBloom:   l2Block.LogsBloom,
		L1Messages:  l1Msg,
	}
	if restMsg, err = nbm.MarshalBinary(); err != nil {
		return
	}
	return
}

func (bc *Version2Converter) Recover(blsMsg []byte, restMsg []byte, txs [][]byte) (*catalyst.ExecutableL2Data, []types.L1Message, error) {
	bm := new(types.BLSMessage)
	if err := bm.UnmarshalBinary(blsMsg); err != nil {
		return nil, nil, err
	}
	nbm := new(types.NonBLSMessage)
	if err := nbm.UnmarshalBinary(restMsg); err != nil {
		return nil, nil, err
	}
	return &catalyst.ExecutableL2Data{
		ParentHash:   bm.ParentHash,
		Miner:        bm.Miner,
		Number:       bm.Number,
		GasLimit:     bm.GasLimit,
		BaseFee:      bm.BaseFee,
		Timestamp:    bm.Timestamp,
		StateRoot:    nbm.StateRoot,
		GasUsed:      nbm.GasUsed,
		ReceiptRoot:  nbm.ReceiptRoot,
		LogsBloom:    nbm.LogsBloom,
		Transactions: txs,
	}, nbm.L1Messages, nil
}
