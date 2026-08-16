package derivation

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/rpc"
	"github.com/stretchr/testify/require"
	tmlog "github.com/tendermint/tendermint/libs/log"

	"morph-l2/bindings/bindings"
	nodedb "morph-l2/node/db"
	nodetypes "morph-l2/node/types"
)

const selectorCursorTestL1Height = uint64(42)

// selectorCursorTestDB is the smallest in-memory implementation needed to
// observe whether derivationBlock committed its L1 scan cursor. The L1 message
// methods are unused by these tests but are required by Database.
type selectorCursorTestDB struct {
	latestDerivation *uint64
	derivationBlocks map[uint64]*nodedb.DerivationL1Block
}

func newSelectorCursorTestDB() *selectorCursorTestDB {
	return &selectorCursorTestDB{derivationBlocks: make(map[uint64]*nodedb.DerivationL1Block)}
}

func (db *selectorCursorTestDB) ReadLatestDerivationL1Height() *uint64 {
	if db.latestDerivation == nil {
		return nil
	}
	value := *db.latestDerivation
	return &value
}

func (db *selectorCursorTestDB) WriteLatestDerivationL1Height(latest uint64) {
	db.latestDerivation = new(uint64)
	*db.latestDerivation = latest
}

func (db *selectorCursorTestDB) ReadDerivationL1BlockRange(from, to uint64) []*nodedb.DerivationL1Block {
	if from > to {
		return nil
	}
	blocks := make([]*nodedb.DerivationL1Block, 0, to-from+1)
	for height := from; height <= to; height++ {
		if block := db.derivationBlocks[height]; block != nil {
			copy := *block
			blocks = append(blocks, &copy)
		}
	}
	return blocks
}

func (db *selectorCursorTestDB) WriteDerivationL1Block(block *nodedb.DerivationL1Block) {
	copy := *block
	db.derivationBlocks[block.Number] = &copy
}

func (db *selectorCursorTestDB) DeleteDerivationL1BlocksFrom(height uint64) {
	for number := range db.derivationBlocks {
		if number >= height {
			delete(db.derivationBlocks, number)
		}
	}
}

func (*selectorCursorTestDB) ReadLatestSyncedL1Height() *uint64 { return nil }

func (*selectorCursorTestDB) ReadL1MessagesInRange(uint64, uint64) []nodetypes.L1Message {
	return nil
}

func (*selectorCursorTestDB) ReadL1MessageByIndex(uint64) *nodetypes.L1Message { return nil }

func (*selectorCursorTestDB) WriteLatestSyncedL1Height(uint64) {}

func (*selectorCursorTestDB) WriteSyncedL1Messages([]nodetypes.L1Message, uint64) error {
	return nil
}

// selectorCursorEthAPI provides the exact JSON-RPC surface derivationBlock
// reaches before it either skips an unknown selector or aborts on malformed
// calldata.
type selectorCursorEthAPI struct {
	tx            *eth.Transaction
	from          common.Address
	rollupAddress common.Address
	header        *eth.Header
}

func (api *selectorCursorEthAPI) BlockNumber(context.Context) (hexutil.Uint64, error) {
	return hexutil.Uint64(selectorCursorTestL1Height), nil
}

func (api *selectorCursorEthAPI) GetLogs(context.Context, map[string]interface{}) ([]eth.Log, error) {
	return []eth.Log{{
		Address:     api.rollupAddress,
		Topics:      []common.Hash{RollupEventTopicHash},
		TxHash:      api.tx.Hash(),
		BlockHash:   api.header.Hash(),
		BlockNumber: selectorCursorTestL1Height,
	}}, nil
}

func (*selectorCursorEthAPI) Call(context.Context, map[string]interface{}, string) (hexutil.Bytes, error) {
	// Rollup.lastCommittedBatchIndex() == 1.
	return common.LeftPadBytes([]byte{1}, 32), nil
}

func (api *selectorCursorEthAPI) GetTransactionByHash(context.Context, common.Hash) (json.RawMessage, error) {
	raw, err := api.tx.MarshalJSON()
	if err != nil {
		return nil, err
	}
	var fields map[string]interface{}
	if err := json.Unmarshal(raw, &fields); err != nil {
		return nil, err
	}
	fields["blockNumber"] = hexutil.EncodeUint64(selectorCursorTestL1Height)
	fields["blockHash"] = api.header.Hash()
	fields["from"] = api.from
	return json.Marshal(fields)
}

func (api *selectorCursorEthAPI) GetBlockByNumber(context.Context, string, bool) (*eth.Header, error) {
	return api.header, nil
}

func TestDerivationBlockSelectorErrorControlsCursorAdvance(t *testing.T) {
	verifyModes := []string{VerifyModeLayer1, VerifyModeLocal}
	tests := []struct {
		name        string
		calldata    []byte
		wantAdvance bool
	}{
		{
			name:        "unknown selector is skipped and successful poll advances cursor",
			calldata:    []byte{0xde, 0xad, 0xbe, 0xef},
			wantAdvance: true,
		},
		{
			name:        "known selector with malformed calldata aborts poll without cursor advance",
			calldata:    []byte{0x41, 0xf7, 0x56, 0xda},
			wantAdvance: false,
		},
	}

	for _, verifyMode := range verifyModes {
		for _, test := range tests {
			t.Run(verifyMode+"/"+test.name, func(t *testing.T) {
				d, db := newSelectorCursorTestDerivation(t, verifyMode, test.calldata)

				d.derivationBlock(context.Background())

				cursor := db.ReadLatestDerivationL1Height()
				if test.wantAdvance {
					require.NotNil(t, cursor)
					require.Equal(t, selectorCursorTestL1Height, *cursor)
					require.Contains(t, db.derivationBlocks, selectorCursorTestL1Height,
						"a successful poll records the scanned L1 block before advancing")
				} else {
					require.Nil(t, cursor)
					require.Empty(t, db.derivationBlocks,
						"an aborted poll must not record the range it did not commit")
				}
			})
		}
	}
}

func newSelectorCursorTestDerivation(
	t *testing.T,
	verifyMode string,
	calldata []byte,
) (*Derivation, *selectorCursorTestDB) {
	t.Helper()

	key, err := crypto.GenerateKey()
	require.NoError(t, err)
	from := crypto.PubkeyToAddress(key.PublicKey)
	rollupAddress := common.HexToAddress("0x1000000000000000000000000000000000000001")
	unsigned := eth.NewTransaction(
		0,
		rollupAddress,
		big.NewInt(0),
		100_000,
		big.NewInt(1),
		calldata,
	)
	tx, err := eth.SignTx(unsigned, eth.NewEIP155Signer(big.NewInt(1)), key)
	require.NoError(t, err)

	header := &eth.Header{
		Number:     new(big.Int).SetUint64(selectorCursorTestL1Height),
		Difficulty: big.NewInt(0),
		GasLimit:   30_000_000,
		Time:       1,
		Extra:      []byte{},
	}
	server := rpc.NewServer()
	require.NoError(t, server.RegisterName("eth", &selectorCursorEthAPI{
		tx:            tx,
		from:          from,
		rollupAddress: rollupAddress,
		header:        header,
	}))
	rpcClient := rpc.DialInProc(server)
	t.Cleanup(rpcClient.Close)
	t.Cleanup(server.Stop)
	l1Client := ethclient.NewClient(rpcClient)

	rollup, err := bindings.NewRollup(rollupAddress, l1Client)
	require.NoError(t, err)
	db := newSelectorCursorTestDB()
	logger := tmlog.NewNopLogger()
	d := newFixtureDerivation(t)
	d.ctx = context.Background()
	d.db = db
	d.l1Client = l1Client
	d.RollupContractAddress = rollupAddress
	d.confirmations = rpc.LatestBlockNumber
	d.l2Client = nodetypes.NewRetryableClient(nil, l1Client, logger)
	d.logger = logger
	d.rollup = rollup
	d.metrics = newDiscardMetrics()
	d.startHeight = selectorCursorTestL1Height
	d.fetchBlockRange = 100
	d.verifyMode = verifyMode
	d.reorgCheckDepth = 10

	return d, db
}
