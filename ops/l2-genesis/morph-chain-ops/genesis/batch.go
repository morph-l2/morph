package genesis

import (
	"encoding/binary"

	"github.com/morph-l2/go-ethereum/common"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"

	nodetypes "morph-l2/node/types"
)

const (
	genesisBatchVersion      = uint8(0)
	genesisBatchHeaderLength = 249
)

// emptyBlobVersionedHash is the KZG versioned hash of an empty blob.
var emptyBlobVersionedHash = common.HexToHash("0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014")

// GenesisBatchHeader builds the raw batch header bytes committed to L1 alongside
// the L2 genesis block. It is called by ops/l2-genesis when bootstrapping a
// fresh chain.
func GenesisBatchHeader(genesisHeader *ethtypes.Header) ([]byte, error) {
	wb := nodetypes.WrappedBlock{
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
	// Data hash for a single-block batch with no L1 messages:
	// keccak256(blockContext[:58]). The last 2 bytes (numL1Messages) are excluded.
	dataHash := crypto.Keccak256Hash(blockContext[:58])

	// V0 batch header layout (249 bytes):
	// version(1) | batchIndex(8) | l1MsgPopped(8) | totalL1MsgPopped(8) |
	// dataHash(32) | blobVersionedHash(32) | prevStateRoot(32) |
	// postStateRoot(32) | withdrawalRoot(32) | sequencerSetVerifyHash(32) |
	// parentBatchHash(32)
	header := make([]byte, genesisBatchHeaderLength)
	header[0] = genesisBatchVersion
	binary.BigEndian.PutUint64(header[1:], 0)  // batchIndex
	binary.BigEndian.PutUint64(header[9:], 0)  // l1MessagePopped
	binary.BigEndian.PutUint64(header[17:], 0) // totalL1MessagePopped
	copy(header[25:], dataHash[:])
	copy(header[57:], emptyBlobVersionedHash[:])
	// prevStateRoot (89:121) — zero for genesis
	copy(header[121:], genesisHeader.Root[:]) // postStateRoot
	// withdrawalRoot (153:185) — zero for genesis
	// sequencerSetVerifyHash (185:217) — zero for genesis
	// parentBatchHash (217:249) — zero for genesis
	return header, nil
}
