package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/rlp"
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
	ParentHash         common.Hash    `json:"parentHash"     gencodec:"required"`
	Miner              common.Address `json:"miner"          gencodec:"required"`
	Number             uint64         `json:"number"         gencodec:"required"`
	GasLimit           uint64         `json:"gasLimit"       gencodec:"required"`
	Timestamp          uint64         `json:"timestamp"      gencodec:"required"`
	StateRoot          common.Hash    `json:"stateRoot"`
	GasUsed            uint64         `json:"gasUsed"`
	ReceiptRoot        common.Hash    `json:"receiptsRoot"`
	LogsBloom          []byte         `json:"logsBloom"`
	WithdrawTrieRoot   common.Hash    `json:"withdrawTrieRoot"`
	RowConsumption     interface{}    `json:"rowConsumption"` // keep this field for compatibility
	NextL1MessageIndex uint64         `json:"nextL1MessageIndex"`
	Hash               common.Hash    `json:"hash"`

	CollectedL1TxHashes []common.Hash               `json:"l1TxHashes" rlp:"optional"`
	SkippedL1Txs        []*types.SkippedTransaction `json:"-" rlp:"optional"`
	BaseFee             *big.Int                    `json:"baseFeePerGas"  rlp:"optional"`
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
	// Number(8) || Timestamp(8) || BaseFee(32) || GasLimit(8) || numTxs(2) || numL1Messages(2) || Miner(20)
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

	if wb.Miner != types.EmptyAddress {
		blsBytes = append(blsBytes, make([]byte, common.AddressLength)...)
		copy(blsBytes[60:], wb.Miner.Bytes())
	}

	return blsBytes
}

func (wb *WrappedBlock) DecodeBlockContext(bc []byte) (uint16, uint16, error) {
	reader := bytes.NewReader(bc)
	bsBaseFee := make([]byte, 32)
	if err := binary.Read(reader, binary.BigEndian, &wb.Number); err != nil {
		return 0, 0, err
	}
	if err := binary.Read(reader, binary.BigEndian, &wb.Timestamp); err != nil {
		return 0, 0, err
	}
	if err := binary.Read(reader, binary.BigEndian, &bsBaseFee); err != nil {
		return 0, 0, err
	}
	wb.BaseFee = new(big.Int).SetBytes(bsBaseFee)
	if err := binary.Read(reader, binary.BigEndian, &wb.GasLimit); err != nil {
		return 0, 0, err
	}
	var txsNum, l1MsgNum uint16
	if err := binary.Read(reader, binary.BigEndian, &txsNum); err != nil {
		return 0, 0, err
	}
	if err := binary.Read(reader, binary.BigEndian, &l1MsgNum); err != nil {
		return 0, 0, err
	}
	return txsNum, l1MsgNum, nil
}

func WrappedBlockFromBytes(blockBytes []byte) (*WrappedBlock, error) {
	var curBlock = new(WrappedBlock)
	if err := curBlock.UnmarshalBinary(blockBytes); err != nil {
		return nil, err
	}
	return curBlock, nil
}

func HeightFromBlockBytes(blockBytes []byte) (uint64, error) {
	curBlock, err := WrappedBlockFromBytes(blockBytes)
	if err != nil {
		return 0, err
	}
	return curBlock.Number, nil
}

func HeightFromBlockContextBytes(blockContextBytes []byte) (uint64, error) {
	if len(blockContextBytes) != 60 {
		return 0, fmt.Errorf("wrong block context bytes length, input: %x", blockContextBytes)
	}
	return binary.BigEndian.Uint64(blockContextBytes[:8]), nil
}
