package derivation

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
)

func TestGetBlob(t *testing.T) {
	url := os.Getenv("BLOB_URL")
	if url == "" {
		return
	}
	var (
		start uint64 = 1590159
		end   uint64 = 1590159
	)
	baseHttp := NewBasicHTTPClient(url, nil)
	// query blob
	l1BeaconClient := NewL1BeaconClient(baseHttp)
	l1Client, err := ethclient.Dial(url)
	require.NoError(t, err)
	logs, err := testTchRollupLog(l1Client, context.Background(), start, end)
	require.NoError(t, err)
	if len(logs) > 0 {
		for _, lg := range logs {
			txHash := lg.TxHash
			block, err := l1Client.BlockByNumber(context.Background(), big.NewInt(int64(lg.BlockNumber)))
			require.NoError(t, err)
			tx, _, err := l1Client.TransactionByHash(context.Background(), txHash)
			require.NoError(t, err)
			indexedBlobHashes := dataAndHashesFromTxs(block.Transactions(), tx)
			fmt.Println(indexedBlobHashes)
			header, err := l1Client.HeaderByNumber(context.Background(), big.NewInt(int64(lg.BlockNumber)))
			require.NoError(t, err)
			var bts eth.BlobTxSidecar
			if len(indexedBlobHashes) != 0 {
				bts, err = l1BeaconClient.GetBlobSidecar(context.Background(), L1BlockRef{
					Time: header.Time,
				}, indexedBlobHashes)
				require.NoError(t, err)
			}
			t.Log(len(bts.Blobs))
		}

	}

}

func testTchRollupLog(l1Client *ethclient.Client, ctx context.Context, from, to uint64) ([]eth.Log, error) {
	RollupContractAddress := common.HexToAddress("0x511d92b63ae7471fd5239bded29b76a446698a00")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0).SetUint64(from),
		ToBlock:   big.NewInt(0).SetUint64(to),
		Addresses: []common.Address{
			RollupContractAddress,
		},
		Topics: [][]common.Hash{
			{RollupEventTopicHash},
		},
	}
	return l1Client.FilterLogs(ctx, query)
}
