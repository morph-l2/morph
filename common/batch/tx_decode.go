package batch

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/rlp"
)

func DecodeTxsFromBytes(txsBytes []byte) (eth.Transactions, error) {
	reader := bytes.NewReader(txsBytes)
	txs := make(eth.Transactions, 0)
	for {
		var (
			firstByte   byte
			fullTxBytes []byte
			innerTx     eth.TxData
			err         error
		)
		if err = binary.Read(reader, binary.BigEndian, &firstByte); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if firstByte == 0 {
			break
		}

		switch firstByte {
		case eth.AccessListTxType:
			if err := binary.Read(reader, binary.BigEndian, &firstByte); err != nil {
				return nil, err
			}
			innerTx = new(eth.AccessListTx)
		case eth.DynamicFeeTxType:
			if err := binary.Read(reader, binary.BigEndian, &firstByte); err != nil {
				return nil, err
			}
			innerTx = new(eth.DynamicFeeTx)
		case eth.SetCodeTxType:
			if err := binary.Read(reader, binary.BigEndian, &firstByte); err != nil {
				return nil, err
			}
			innerTx = new(eth.SetCodeTx)
		case eth.MorphTxType:
			if err := binary.Read(reader, binary.BigEndian, &firstByte); err != nil {
				return nil, err
			}
			innerTx = new(eth.MorphTx)
		default:
			if firstByte <= 0xf7 {
				return nil, fmt.Errorf("not supported tx type: %d", firstByte)
			}
			innerTx = new(eth.LegacyTx)
		}

		fullTxBytes, err = extractInnerTxFullBytes(firstByte, reader)
		if err != nil {
			return nil, err
		}
		if err = rlp.DecodeBytes(fullTxBytes, innerTx); err != nil {
			return nil, err
		}
		txs = append(txs, eth.NewTx(innerTx))
	}
	return txs, nil
}

func extractInnerTxFullBytes(firstByte byte, reader io.Reader) ([]byte, error) {
	sizeByteLen := firstByte - 0xf7

	sizeByte := make([]byte, sizeByteLen)
	if err := binary.Read(reader, binary.BigEndian, sizeByte); err != nil {
		return nil, err
	}
	size := binary.BigEndian.Uint32(append(make([]byte, 4-len(sizeByte)), sizeByte...))

	txRaw := make([]byte, size)
	if err := binary.Read(reader, binary.BigEndian, txRaw); err != nil {
		return nil, err
	}
	fullTxBytes := make([]byte, 1+uint32(sizeByteLen)+size)
	copy(fullTxBytes[:1], []byte{firstByte})
	copy(fullTxBytes[1:1+sizeByteLen], sizeByte)
	copy(fullTxBytes[1+sizeByteLen:], txRaw)

	return fullTxBytes, nil
}
