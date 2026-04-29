package batch

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"time"

	"morph-l2/bindings/bindings"
	"morph-l2/tx-submitter/db"
	"morph-l2/tx-submitter/iface"
	"morph-l2/tx-submitter/types"
	"morph-l2/tx-submitter/utils"

	"github.com/holiman/uint256"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/consensus/misc/eip4844"
	ethtypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/eth"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/log"
	"github.com/stretchr/testify/require"
)

var pk = ""

func TestRollupWithProof(t *testing.T) {
	testDir := filepath.Join(t.TempDir(), "testleveldb")
	os.RemoveAll(testDir)
	t.Cleanup(func() {
		os.RemoveAll(testDir)
	})
	testDB, err := db.New(testDir)
	require.NoError(t, err)

	cache := NewBatchCache(nil, l1Client, []iface.L2Client{l2Client}, rollupContract, l2Caller, testDB)
	err = cache.InitFromRollupByRange()
	require.NoError(t, err)

	privateKey, err := crypto.HexToECDSA(pk[2:])
	require.NoError(t, err)
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	ctx := context.Background()
	l1ChainId, err := l1Client.ChainID(ctx)
	require.NoError(t, err)
	rollup, err := bindings.NewRollup(rollupAddr, l1Client)
	require.NoError(t, err)
	abi, err := bindings.RollupMetaData.GetAbi()
	require.NoError(t, err)
	latestCommitBatchIndex, err := rollup.LastCommittedBatchIndex(nil)
	require.NoError(t, err)

	batch, err := cache.Get(latestCommitBatchIndex.Uint64() + 1)
	require.NoError(t, err)
	h := crypto.Keccak256Hash(batch.CurrentSequencerSetBytes)
	t.Log("sequencer verify hash:", h.String())

	signature, err := buildSigInput(batch)
	require.NoError(t, err)
	rollupBatch := bindings.IRollupBatchDataInput{
		Version:           uint8(batch.Version),
		ParentBatchHeader: batch.ParentBatchHeader,
		LastBlockNumber:   batch.LastBlockNumber,
		NumL1Messages:     batch.NumL1Messages,
		PrevStateRoot:     batch.PrevStateRoot,
		PostStateRoot:     batch.PostStateRoot,
		WithdrawalRoot:    batch.WithdrawRoot,
	}
	tip, gasFeeCap, blobFee, head, err := getGasTipAndCap(l1Client)
	require.NoError(t, err)

	calldata, err := abi.Pack("commitBatch", rollupBatch, *signature)
	require.NoError(t, err)
	nonce, err := l1Client.NonceAt(context.Background(), address, nil)
	require.NoError(t, err)
	tx, err := createBlobTx(l1Client, batch, nonce, 3200000, tip, gasFeeCap, blobFee, calldata, head)
	require.NoError(t, err)
	transaction, err := sign(tx, ethtypes.LatestSignerForChainID(l1ChainId), privateKey)
	require.NoError(t, err)
	t.Log("txHash", transaction.Hash().String())
	err = sendTx(l1Client, 500000000000000000, transaction)
	require.NoError(t, err)
	time.Sleep(2 * time.Second)
	receipt, err := l1Client.TransactionReceipt(ctx, transaction.Hash())
	require.NoError(t, err)
	t.Log("receipt status", receipt.Status)
	t.Log("receipt", receipt)

}

func sign(tx *ethtypes.Transaction, signer ethtypes.Signer, prv *ecdsa.PrivateKey) (*ethtypes.Transaction, error) {
	signedTx, err := ethtypes.SignTx(tx, signer, prv)
	if err != nil {
		return nil, fmt.Errorf("sign tx error:%v", err)
	}
	return signedTx, nil
}

func createBlobTx(l1client *ethclient.Client, batch *eth.RPCRollupBatch, nonce, gas uint64, tip, gasFeeCap, blobFee *big.Int, calldata []byte, head *ethtypes.Header) (*ethtypes.Transaction, error) {
	versionedHashes := types.BlobHashes(batch.Sidecar.Blobs, batch.Sidecar.Commitments)
	sidecar := &ethtypes.BlobTxSidecar{
		Blobs:       batch.Sidecar.Blobs,
		Commitments: batch.Sidecar.Commitments,
	}
	chainID, err := l1client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	switch types.DetermineBlobVersion(head, chainID.Uint64()) {
	case ethtypes.BlobSidecarVersion0:
		sidecar.Version = ethtypes.BlobSidecarVersion0
		proof, err := types.MakeBlobProof(sidecar.Blobs, sidecar.Commitments)
		if err != nil {
			return nil, fmt.Errorf("gen blob proof failed %v", err)
		}
		sidecar.Proofs = proof
	case ethtypes.BlobSidecarVersion1:
		sidecar.Version = ethtypes.BlobSidecarVersion1
		proof, err := types.MakeCellProof(sidecar.Blobs)
		if err != nil {
			return nil, fmt.Errorf("gen cell proof failed %v", err)
		}
		sidecar.Proofs = proof
	default:
		return nil, fmt.Errorf("unsupported blob version")
	}

	return ethtypes.NewTx(&ethtypes.BlobTx{
		ChainID:    uint256.MustFromBig(chainID),
		Nonce:      nonce,
		GasTipCap:  uint256.MustFromBig(tip),
		GasFeeCap:  uint256.MustFromBig(gasFeeCap),
		Gas:        gas,
		To:         rollupAddr,
		Data:       calldata,
		BlobFeeCap: uint256.MustFromBig(blobFee),
		BlobHashes: versionedHashes,
		Sidecar:    sidecar,
	}), nil
}

func getGasTipAndCap(l1client *ethclient.Client) (*big.Int, *big.Int, *big.Int, *ethtypes.Header, error) {
	head, err := l1client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	if head.BaseFee != nil {
		log.Info("market fee info", "feecap", head.BaseFee)
	}

	tip, err := l1client.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, nil, nil, nil, err
	}
	log.Info("market fee info", "tip", tip)

	tip = new(big.Int).Mul(tip, big.NewInt(int64(200)))
	tip = new(big.Int).Div(tip, big.NewInt(100))

	var gasFeeCap *big.Int
	if head.BaseFee != nil {
		gasFeeCap = new(big.Int).Add(
			tip,
			new(big.Int).Mul(head.BaseFee, big.NewInt(2)),
		)
	} else {
		gasFeeCap = new(big.Int).Set(tip)
	}

	// calc blob fee cap
	var blobFee *big.Int
	if head.ExcessBlobGas != nil {
		id, err := l1client.ChainID(context.Background())
		if err != nil {
			return nil, nil, nil, nil, err
		}
		log.Info("market blob fee info", "excess blob gas", *head.ExcessBlobGas)
		blobConfig, exist := types.ChainConfigMap[id.Uint64()]
		if !exist {
			blobConfig = types.DefaultBlobConfig
		}
		blobFeeDenominator := types.GetBlobFeeDenominator(blobConfig, head.Time)
		blobFee = eip4844.CalcBlobFee(*head.ExcessBlobGas, blobFeeDenominator.Uint64())
		// Set to 3x to handle blob market congestion
		blobFee = new(big.Int).Mul(blobFee, big.NewInt(3))
	}

	return tip, gasFeeCap, blobFee, head, nil
}

func buildSigInput(batch *eth.RPCRollupBatch) (*bindings.IRollupBatchSignatureInput, error) {
	sigData := &bindings.IRollupBatchSignatureInput{
		SignedSequencersBitmap: common.Big0,
		SequencerSets:          batch.CurrentSequencerSetBytes,
		Signature:              []byte("0x"),
	}
	return sigData, nil
}

// send tx to l1 with business logic check
func sendTx(client iface.Client, txFeeLimit uint64, tx *ethtypes.Transaction) error {
	// fee limit
	if txFeeLimit > 0 {
		var fee uint64
		// calc tx gas fee
		if tx.Type() == ethtypes.BlobTxType {
			blobFee := new(big.Int).Mul(tx.BlobGasFeeCap(), new(big.Int).SetUint64(tx.BlobGas()))
			txFee := new(big.Int).Mul(tx.GasPrice(), new(big.Int).SetUint64(tx.Gas()))
			totalFee := new(big.Int).Add(blobFee, txFee)
			if !totalFee.IsUint64() || totalFee.Uint64() > txFeeLimit {
				return fmt.Errorf("%v:limit=%v,but got=%v", utils.ErrExceedFeeLimit, txFeeLimit, totalFee)
			}
			return client.SendTransaction(context.Background(), tx)
		} else {
			fee = tx.GasPrice().Uint64() * tx.Gas()
		}

		if fee > txFeeLimit {
			return fmt.Errorf("%v:limit=%v,but got=%v", utils.ErrExceedFeeLimit, txFeeLimit, fee)
		}
	}

	return client.SendTransaction(context.Background(), tx)
}
