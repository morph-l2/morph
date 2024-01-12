package sync

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/hashicorp/go-multierror"
	"github.com/holiman/uint256"
	"github.com/morph-l2/bindings/bindings"
	"github.com/morph-l2/node/types"
	"github.com/scroll-tech/go-ethereum/common"
	eth "github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/log"
)

var (
	DepositEventABI              = "QueueTransaction(address,address,uint256,uint64,uint256,bytes)"
	DepositEventABIHash          = crypto.Keccak256Hash([]byte(DepositEventABI))
	DepositEventVersion0         = common.Hash{}
	L2CrossDomainMessengerABI, _ = bindings.L2CrossDomainMessengerMetaData.GetAbi()
)

func L1MessageTxFromEvent(event *bindings.MorphPortalQueueTransaction) eth.L1MessageTx {
	return eth.L1MessageTx{
		QueueIndex: event.QueueIndex,
		Gas:        event.GasLimit.Uint64(),
		To:         &event.Target,
		Value:      event.Value,
		Data:       event.Data,
		Sender:     event.Sender,
	}
}

func (c *BridgeClient) deriveFromReceipt(receipts []*eth.Receipt) ([]types.L1Message, error) {
	var out []types.L1Message
	var result error
	for i, rec := range receipts {
		if rec.Status != eth.ReceiptStatusSuccessful {
			continue
		}
		for j, lg := range rec.Logs {
			if lg.Address == c.morphPortalAddress && len(lg.Topics) > 0 && lg.Topics[0] == DepositEventABIHash {
				event, err := c.filter.ParseQueueTransaction(*lg)
				if err != nil {
					result = multierror.Append(result, fmt.Errorf("malformatted L1 deposit log in receipt %d, log %d: %w", i, j, err))
				} else {
					if event == nil {
						continue
					}
					out = append(out, types.L1Message{
						L1MessageTx: L1MessageTxFromEvent(event),
						L1TxHash:    lg.TxHash,
					})
				}
			}
		}
	}
	return out, result
}

func deriveFromReceipt(receipts []*eth.Receipt, depositContractAddr common.Address) ([]types.L1Message, error) {
	var out []types.L1Message
	var result error
	for i, rec := range receipts {
		if rec.Status != eth.ReceiptStatusSuccessful {
			continue
		}
		for j, lg := range rec.Logs {
			if lg.Address == depositContractAddr && len(lg.Topics) > 0 && lg.Topics[0] == DepositEventABIHash {
				msg, err := UnmarshalDepositLogEvent(lg)
				if err != nil {
					result = multierror.Append(result, fmt.Errorf("malformatted L1 deposit log in receipt %d, log %d: %w", i, j, err))
				} else {
					if msg == nil {
						continue
					}
					out = append(out, types.L1Message{
						L1MessageTx: *msg,
						L1TxHash:    lg.TxHash,
					})
				}
			}
		}
	}
	return out, result
}

// UnmarshalDepositLogEvent decodes an EVM log entry emitted by the deposit contract into typed deposit data.
//
// parse log data for:
//
//	event TransactionDeposited(
//	    address indexed from,
//	    address indexed to,
//	    uint256 indexed version,
//	    bytes opaqueData
//	);
//
// Additionally, the event log-index and
func UnmarshalDepositLogEvent(ev *eth.Log) (*eth.L1MessageTx, error) {
	if len(ev.Topics) != 4 {
		return nil, fmt.Errorf("expected 4 event topics (event identity, indexed from, indexed to, indexed version), got %d", len(ev.Topics))
	}
	if ev.Topics[0] != DepositEventABIHash {
		return nil, fmt.Errorf("invalid deposit event selector: %s, expected %s", ev.Topics[0], DepositEventABIHash)
	}
	if len(ev.Data) < 64 {
		return nil, fmt.Errorf("incomplate opaqueData slice header (%d bytes): %x", len(ev.Data), ev.Data)
	}
	if len(ev.Data)%32 != 0 {
		return nil, fmt.Errorf("expected log data to be multiple of 32 bytes: got %d bytes", len(ev.Data))
	}

	// indexed 0
	from := common.BytesToAddress(ev.Topics[1][12:])
	log.Trace("Unmarshalling deposit log", "from", from.String())
	// indexed 1
	to := common.BytesToAddress(ev.Topics[2][12:])
	// indexed 2
	version := ev.Topics[3]
	// unindexed data
	// Solidity serializes the event's Data field as follows:
	// abi.encode(abi.encodPacked(uint256 mint, uint256 value, uint64 gasLimit, uint8 isCreation, bytes data))
	// Thus the first 32 bytes of the Data will give us the offset of the opaqueData,
	// which should always be 0x20.
	var opaqueContentOffset uint256.Int
	opaqueContentOffset.SetBytes(ev.Data[0:32])
	if !opaqueContentOffset.IsUint64() || opaqueContentOffset.Uint64() != 32 {
		return nil, fmt.Errorf("invalid opaqueData slice header offset: %d", opaqueContentOffset.Uint64())
	}
	// The next 32 bytes indicate the length of the opaqueData content.
	var opaqueContentLength uint256.Int
	opaqueContentLength.SetBytes(ev.Data[32:64])
	// Make sure the length is an uint64, it's not larger than the remaining data, and the log is using minimal padding (i.e. can't add 32 bytes without exceeding data)
	if !opaqueContentLength.IsUint64() || opaqueContentLength.Uint64() > uint64(len(ev.Data)-64) || opaqueContentLength.Uint64()+32 <= uint64(len(ev.Data)-64) {
		return nil, fmt.Errorf("invalid opaqueData slice header length: %d", opaqueContentLength.Uint64())
	}
	// The remaining data is the opaqueData which is tightly packed
	// and then padded to 32 bytes by the EVM.
	opaqueData := ev.Data[64 : 64+opaqueContentLength.Uint64()]

	var tx *eth.L1MessageTx
	var err error
	switch version {
	case DepositEventVersion0:
		tx, err = unmarshalDepositVersion0(to, opaqueData)
	default:
		return nil, fmt.Errorf("invalid deposit version, got %s", version)
	}
	if err != nil {
		if err == types.ErrNotFromCrossDomainMessenger {
			log.Warn("found the message not sent by L1CrossDomainMessenger, ignore it for now")
			return nil, nil
		}
		return nil, fmt.Errorf("failed to decode deposit (version %s): %w", version, err)
	}
	tx.Sender = from
	return tx, nil
}

func unmarshalDepositVersion0(to common.Address, opaqueData []byte) (*eth.L1MessageTx, error) {
	var message eth.L1MessageTx
	if len(opaqueData) < 32+32+8+1 {
		return nil, fmt.Errorf("unexpected opaqueData length: %d", len(opaqueData))
	}
	offset := uint64(0)
	// uint256 mint
	mint := new(big.Int).SetBytes(opaqueData[offset : offset+32])
	offset += 32
	log.Trace("Unmarshalling deposit log", "mint", mint)

	// uint256 value
	value := new(big.Int).SetBytes(opaqueData[offset : offset+32])
	offset += 32
	message.Value = value
	log.Trace("Unmarshalling deposit log", "value", value)

	// uint64 gas
	gas := new(big.Int).SetBytes(opaqueData[offset : offset+8])
	if !gas.IsUint64() {
		return nil, fmt.Errorf("bad gas value: %x", opaqueData[offset:offset+8])
	}
	message.Gas = gas.Uint64()
	offset += 8

	// uint8 isCreation
	// isCreation: If the boolean byte is 1 then dep.To will stay nil,
	// and it will create a contract using L2 account nonce to determine the created address.
	if opaqueData[offset] == 0 {
		message.To = &to
	}
	offset += 1

	// The remainder of the opaqueData is the transaction data (without length prefix).
	// The data may be padded to a multiple of 32 bytes
	txDataLen := uint64(len(opaqueData)) - offset

	// remaining bytes fill the data
	message.Data = opaqueData[offset : offset+txDataLen]

	// NOTE: currently we acquire the nonce from relayMessage input parsed from event data,
	// which means we only allow the cross message data which are formed as relayMessage for now.
	// in the future, the `nonce` is supposed to be exposed as one of the event fields
	relayMessage, err := unpackRelayMessage(message.Data)
	if err != nil {
		return nil, types.ErrNotFromCrossDomainMessenger
	}
	message.QueueIndex, err = types.DecodeNonce(relayMessage.nonce)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

type relayMessageData struct {
	nonce       *big.Int
	sender      common.Address
	target      common.Address
	value       *big.Int
	minGasLimit *big.Int
	message     []byte
}

func unpackRelayMessage(data []byte) (*relayMessageData, error) {
	abi := L2CrossDomainMessengerABI
	method, ok := abi.Methods["relayMessage"]
	if !ok {
		return nil, errors.New("can not find the method of relayMessage")
	}
	args := method.Inputs
	unpacked, err := args.Unpack(data[4:])
	if err != nil {
		return nil, err
	}
	if len(unpacked) != 6 {
		return nil, errors.New("wrong unpacked value length")
	}

	relayMessage := new(relayMessageData)
	relayMessage.nonce = unpacked[0].(*big.Int)
	relayMessage.sender = unpacked[1].(common.Address)
	relayMessage.target = unpacked[2].(common.Address)
	relayMessage.value = unpacked[3].(*big.Int)
	relayMessage.minGasLimit = unpacked[4].(*big.Int)
	relayMessage.message = unpacked[5].([]byte)
	return relayMessage, nil
}
