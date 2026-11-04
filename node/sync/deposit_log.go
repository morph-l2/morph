package sync

import (
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"

	"morph-l2/bindings/bindings"
	"morph-l2/node/types"
)

var (
	DepositEventABI              = "QueueTransaction(address,address,uint256,uint64,uint256,bytes)"
	DepositEventABIHash          = crypto.Keccak256Hash([]byte(DepositEventABI))
	DepositEventVersion0         = common.Hash{}
	L2CrossDomainMessengerABI, _ = bindings.L2CrossDomainMessengerMetaData.GetAbi()
)

// init fails fast at startup if the generated L2CrossDomainMessenger ABI cannot
// be parsed, instead of letting the dropped error above surface later as an
// obscure nil-dereference the first time L2CrossDomainMessengerABI is used.
func init() {
	if _, err := bindings.L2CrossDomainMessengerMetaData.GetAbi(); err != nil {
		panic(fmt.Sprintf("parse L2CrossDomainMessenger ABI: %v", err))
	}
}

func L1MessageTxFromEvent(event *bindings.L1MessageQueueWithGasPriceOracleQueueTransaction) eth.L1MessageTx {
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
