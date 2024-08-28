package sync

import (
	"context"
	"flag"
	"math/big"
	"os"
	"testing"

	"github.com/morph-l2/go-ethereum/common"
	gethTypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/stretchr/testify/require"
	tmlog "github.com/tendermint/tendermint/libs/log"
	"github.com/urfave/cli"

	"morph-l2/node/db"
	"morph-l2/node/types"
)

func TestSyncer_GetL1Message(t *testing.T) {
	//prepare msg
	to := common.BigToAddress(big.NewInt(101))
	msg := types.L1Message{
		L1MessageTx: gethTypes.L1MessageTx{
			QueueIndex: 123,
			Gas:        500000,
			To:         &to,
			Value:      big.NewInt(3e9),
			Data:       []byte("0x1a2b3c"),
			Sender:     common.BigToAddress(big.NewInt(202)),
		},
		L1TxHash: common.BigToHash(big.NewInt(1111)),
	}

	//prepare context
	ctx := PrepareContext()

	//syncer
	store := prepareDB(msg)
	store.WriteLatestSyncedL1Height(100)
	syncConfig := DefaultConfig()
	_ = syncConfig.SetCliContext(ctx)
	syncer, err := NewSyncer(context.Background(), store, syncConfig, tmlog.NewTMLogger(tmlog.NewSyncWriter(os.Stdout)))
	require.NotNil(t, syncer)
	require.NoError(t, err)

	actualMsg, err := syncer.GetL1Message(123, msg.L1TxHash)
	require.NotNil(t, actualMsg)
	require.NoError(t, err)
	require.EqualValues(t, msg.QueueIndex, actualMsg.QueueIndex)
	require.EqualValues(t, msg.L1TxHash, actualMsg.L1TxHash)
	require.EqualValues(t, msg.Gas, actualMsg.Gas)
	require.EqualValues(t, msg.To.Bytes(), actualMsg.To.Bytes())
	require.EqualValues(t, msg.Value, actualMsg.Value)
	require.EqualValues(t, msg.Data, actualMsg.Data)
	require.EqualValues(t, msg.Sender, actualMsg.Sender)

}

func prepareDB(msg types.L1Message) *db.Store {
	db := db.NewMemoryStore()
	msgs := make([]types.L1Message, 0)
	msgs = append(msgs, msg)
	_ = db.WriteSyncedL1Messages(msgs, 0)
	return db
}

func PrepareContext() *cli.Context {
	env := map[string]string{
		"l1.rpc":                   "https://arb1.arbitrum.io/rpc",
		"sync.depositContractAddr": "0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9",
		"l2.engine":                "http://127.0.0.1:8551",
		"l2.eth":                   "http://127.0.0.1:8545",
		"l2.jwt-secret":            "../jwt-secret.txt",
	}
	flagSet := flag.NewFlagSet("testApp", flag.ContinueOnError)
	for k, v := range env {
		flagSet.String(k, v, "param")
		_ = flagSet.Set(k, v)
	}
	ctx := cli.NewContext(nil, flagSet, nil)
	return ctx
}
