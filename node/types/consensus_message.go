package types

import (
	"bytes"
	"github.com/scroll-tech/go-ethereum/core/types"
	"math/big"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/rlp"
)

// Configs need BLS signature
type BLSMessage struct {
	ParentHash common.Hash    `json:"parentHash"     gencodec:"required"`
	Miner      common.Address `json:"miner"          gencodec:"required"`
	Number     uint64         `json:"number"         gencodec:"required"`
	GasLimit   uint64         `json:"gasLimit"       gencodec:"required"`
	BaseFee    *big.Int       `json:"baseFeePerGas"  gencodec:"required"`
	Timestamp  uint64         `json:"timestamp"      gencodec:"required"`
}

func (bm *BLSMessage) MarshalBinary() ([]byte, error) {
	if bm == nil {
		return nil, nil
	}
	return rlp.EncodeToBytes(bm)
}

func (bm *BLSMessage) UnmarshalBinary(b []byte) error {
	return rlp.DecodeBytes(b, bm)
}

// Configs do NOT need BLS signature
type NonBLSMessage struct {
	// execution result
	StateRoot   common.Hash `json:"stateRoot"`
	GasUsed     uint64      `json:"gasUsed"`
	ReceiptRoot common.Hash `json:"receiptsRoot"`
	LogsBloom   []byte      `json:"logsBloom"`

	L1Messages []L1Message `json:"l1Messages"`
}

func (nbm *NonBLSMessage) MarshalBinary() ([]byte, error) {
	if nbm == nil {
		return nil, nil
	}
	return rlp.EncodeToBytes(nbm)
}

func (nbm *NonBLSMessage) UnmarshalBinary(b []byte) error {
	return rlp.DecodeBytes(b, nbm)
}

type RestMessage struct {
	NonBLSMessage
	Miner              common.Address `json:"miner"`
	BlockHash          common.Hash    `json:"blockHash"`
	ParentHash         common.Hash    `json:"parentHash"`
	NextL1MessageIndex uint64         `json:"nextL1MessageIndex"`
}

func (rm *RestMessage) MarshalBinary() ([]byte, error) {
	if rm == nil {
		return nil, nil
	}
	return rlp.EncodeToBytes(rm)
}

func (rm *RestMessage) UnmarshalBinary(b []byte) error {
	return rlp.DecodeBytes(b, rm)
}

type WrappedBlock struct {
	ParentHash         common.Hash          `json:"parentHash"     gencodec:"required"`
	Miner              common.Address       `json:"miner"          gencodec:"required"`
	Number             uint64               `json:"number"         gencodec:"required"`
	GasLimit           uint64               `json:"gasLimit"       gencodec:"required"`
	Timestamp          uint64               `json:"timestamp"      gencodec:"required"`
	StateRoot          common.Hash          `json:"stateRoot"`
	GasUsed            uint64               `json:"gasUsed"`
	ReceiptRoot        common.Hash          `json:"receiptsRoot"`
	LogsBloom          []byte               `json:"logsBloom"`
	WithdrawTrieRoot   common.Hash          `json:"withdrawTrieRoot"`
	RowConsumption     types.RowConsumption `json:"rowConsumption"`
	NextL1MessageIndex uint64               `json:"nextL1MessageIndex"`
	Hash               common.Hash          `json:"hash"`

	CollectedL1Messages []L1Message `json:"l1Messages" rlp:"optional"`
	BaseFee             *big.Int    `json:"baseFeePerGas"  rlp:"optional"`
}

func (wb *WrappedBlock) MarshalBinary() ([]byte, error) {
	if wb == nil {
		return nil, nil
	}
	return rlp.EncodeToBytes(wb)
}

func (wb *WrappedBlock) UnmarshalBinary(b []byte) error {
	return rlp.Decode(bytes.NewReader(b), wb)
}

func (wb *WrappedBlock) BlockContextBytes(txsNum, l1MsgNum int) []byte {
	// Number(8) || Timestamp(8) || BaseFee(32) || GasLimit(8) || numTxs(2) || numL1Messages(2)
	blsBytes := make([]byte, 60)
	copy(blsBytes[:8], Uint64ToBigEndianBytes(wb.Number))
	copy(blsBytes[8:16], Uint64ToBigEndianBytes(wb.Timestamp))
	if wb.BaseFee != nil {
		copy(blsBytes[16:48], wb.BaseFee.FillBytes(make([]byte, 32)))
	} else {
		copy(blsBytes[16:48], make([]byte, 32))
	}
	copy(blsBytes[48:56], Uint64ToBigEndianBytes(wb.GasLimit))
	copy(blsBytes[56:58], Uint16ToBigEndianBytes(uint16(txsNum)))
	copy(blsBytes[58:60], Uint16ToBigEndianBytes(uint16(l1MsgNum)))

	return blsBytes
}
