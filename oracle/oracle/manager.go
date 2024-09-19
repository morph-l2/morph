package oracle

import (
	"context"
	"crypto/ecdsa"
	"crypto/rsa"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/morph-l2/externalsign"
	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	coretypes "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	"github.com/morph-l2/go-ethereum/ethclient"
	"github.com/morph-l2/go-ethereum/log"

	"morph-l2/bindings/bindings"
	"morph-l2/bindings/predeploys"
	"morph-l2/oracle/backoff"
	"morph-l2/oracle/config"
)

type RecordManager interface {
	UploadRollupEpoch(recordRollupEpochInfos []bindings.IRecordRollupEpochInfo) error
	LatestRollupEpoch() (*bindings.IRecordRollupEpochInfo, error)
	UploadRewardsEpoch(recordRewardsEpochInfos []bindings.IRecordRewardEpochInfo) error
	NextRewardEpochIndex() (*big.Int, error)
	UploadBatchEpoch(recordBatchSubmissions []bindings.IRecordBatchSubmission) error
	NextBatchEpochIndex() (*big.Int, error)
}

type RecordClient struct {
	l2Client   *ethclient.Client
	record     *bindings.Record
	recordAddr common.Address
	recordAbi  *abi.ABI
	ctx        context.Context

	privKey         *ecdsa.PrivateKey
	externalRsaPriv *rsa.PrivateKey
	signer          coretypes.Signer
	cfg             *config.Config
	chainId         *big.Int
}

func NewRecordClient(
	l2Client *ethclient.Client,
	record *bindings.Record,
	recordAddr common.Address,
	recordAbi *abi.ABI,
	ctx context.Context,
	privKey *ecdsa.PrivateKey,
	externalRsaPriv *rsa.PrivateKey,
	signer coretypes.Signer,
	cfg *config.Config,
	chainId *big.Int,
) RecordManager {
	return &RecordClient{
		l2Client:        l2Client,
		record:          record,
		recordAddr:      recordAddr,
		recordAbi:       recordAbi,
		ctx:             ctx,
		privKey:         privKey,
		externalRsaPriv: externalRsaPriv,
		signer:          signer,
		cfg:             cfg,
		chainId:         chainId,
	}
}

func (r *RecordClient) UploadRollupEpoch(epochs []bindings.IRecordRollupEpochInfo) error {
	callData, err := r.recordAbi.Pack("recordRollupEpochs", epochs)
	if err != nil {
		return err
	}
	tx, err := r.newRecordTxAndSign(callData)
	if err != nil {
		return err
	}
	log.Info("send record rollup epoch tx success", "txHash", tx.Hash().Hex(), "nonce", tx.Nonce())
	var receipt *coretypes.Receipt
	err = backoff.Do(30, backoff.Exponential(), func() error {
		var err error
		receipt, err = r.waitReceiptWithCtx(r.ctx, tx.Hash())
		return err
	})
	if err != nil {
		return fmt.Errorf("receipt record rollup epochs error:%v", err)
	}
	if receipt.Status != coretypes.ReceiptStatusSuccessful {
		return fmt.Errorf("record rollup epochs not success")
	}
	log.Info("wait receipt success", "txHash", tx.Hash())
	return nil
}

func (r *RecordClient) LatestRollupEpoch() (*bindings.IRecordRollupEpochInfo, error) {
	epochIndex, err := r.record.NextRollupEpochIndex(nil)
	if err != nil {
		return nil, err
	}
	lastEpoch, err := r.record.RollupEpochs(nil, epochIndex.Sub(epochIndex, big.NewInt(1)))
	if err != nil {
		return nil, err
	}
	return &bindings.IRecordRollupEpochInfo{
		Index:     lastEpoch.Index,
		Submitter: lastEpoch.Submitter,
		StartTime: lastEpoch.StartTime,
		EndTime:   lastEpoch.EndTime,
		EndBlock:  lastEpoch.EndBlock,
	}, nil
}

func (r *RecordClient) UploadRewardsEpoch(recordRewardsEpochInfos []bindings.IRecordRewardEpochInfo) error {
	callData, err := r.recordAbi.Pack("recordRewardEpochs", recordRewardsEpochInfos)
	if err != nil {
		return err
	}
	tx, err := r.newRecordTxAndSign(callData)
	if err != nil {
		return fmt.Errorf("record reward epochs error:%v", err)
	}
	err = r.l2Client.SendTransaction(r.ctx, tx)
	if err != nil {
		return fmt.Errorf("send transaction error:%v", err)
	}
	log.Info("send record reward tx success", "txHash", tx.Hash().Hex(), "nonce", tx.Nonce())
	var receipt *types.Receipt
	err = backoff.Do(30, backoff.Exponential(), func() error {
		var err error
		receipt, err = r.waitReceiptWithCtx(r.ctx, tx.Hash())
		return err
	})
	if err != nil {
		return fmt.Errorf("receipt record reward epochs error:%v", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("record reward epochs not success")
	}
	return nil
}

func (r *RecordClient) newRecordTxAndSign(callData []byte) (*types.Transaction, error) {
	from := common.HexToAddress(r.cfg.ExternalSignAddress)
	if !r.cfg.ExternalSign {
		from = crypto.PubkeyToAddress(r.privKey.PublicKey)
	}
	nonce, err := r.l2Client.NonceAt(r.ctx, from, nil)
	if err != nil {
		return nil, err
	}
	// tip and cap
	tip, err := r.l2Client.SuggestGasTipCap(r.ctx)
	if err != nil {
		return nil, err
	}
	head, err := r.l2Client.HeaderByNumber(r.ctx, nil)
	if err != nil {
		return nil, err
	}
	var gasFeeCap *big.Int
	if head.BaseFee != nil {
		gasFeeCap = new(big.Int).Add(
			tip,
			new(big.Int).Mul(head.BaseFee, big.NewInt(2)),
		)
	} else {
		gasFeeCap = new(big.Int).Set(tip)
	}
	gas, err := r.l2Client.EstimateGas(r.ctx, ethereum.CallMsg{
		From:      from,
		To:        &predeploys.RecordAddr,
		GasFeeCap: gasFeeCap,
		GasTipCap: tip,
		Data:      callData,
	})
	if err != nil {
		return nil, err
	}
	return r.sign(types.NewTx(&types.DynamicFeeTx{
		ChainID:   r.chainId,
		Nonce:     nonce,
		GasTipCap: tip,
		GasFeeCap: gasFeeCap,
		Gas:       gas,
		To:        &r.recordAddr,
		Data:      callData}))
}

func (r *RecordClient) sign(tx *types.Transaction) (*types.Transaction, error) {
	if r.cfg.ExternalSign {
		if externalSigner == nil {
			externalSigner = externalsign.NewExternalSign(r.cfg.ExternalSignAppid, r.externalRsaPriv, r.cfg.ExternalSignAddress, r.cfg.ExternalSignChain, r.signer)
		}
		signedTx, err := externalSigner.RequestSign(r.cfg.ExternalSignUrl, tx)
		if err != nil {
			return nil, fmt.Errorf("externalsign sign tx error:%v", err)
		}
		return signedTx, nil
	} else {
		signedTx, err := types.SignTx(tx, r.signer, r.privKey)
		if err != nil {
			return nil, fmt.Errorf("sign tx error:%v", err)
		}
		return signedTx, nil

	}
}

func (r *RecordClient) waitReceiptWithCtx(ctx context.Context, txHash common.Hash) (*coretypes.Receipt, error) {
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-ctx.Done():
			return nil, errors.New("timeout")
		case <-t.C:
			receipt, err := r.l2Client.TransactionReceipt(r.ctx, txHash)
			if errors.Is(err, ethereum.NotFound) {
				continue
			}
			if err != nil {
				return nil, err
			}
			if receipt != nil {
				t.Stop()
				return receipt, nil
			}
		}
	}
}

func (r *RecordClient) NextRewardEpochIndex() (*big.Int, error) {
	return r.record.NextRewardEpochIndex(nil)

}

func (r *RecordClient) UploadBatchEpoch(batchSubmissions []bindings.IRecordBatchSubmission) error {
	callData, err := r.recordAbi.Pack("recordFinalizedBatchSubmissions", batchSubmissions)
	if err != nil {
		return err
	}
	tx, err := r.newRecordTxAndSign(callData)
	if err != nil {
		return fmt.Errorf("record finalized batch error:%v,batchLength:%v", err, len(batchSubmissions))
	}
	log.Info("record finalized batch success", "txHash", tx.Hash(), "batchLength", len(batchSubmissions))
	var receipt *types.Receipt
	err = backoff.Do(30, backoff.Exponential(), func() error {
		var err error
		receipt, err = r.waitReceiptWithCtx(r.ctx, tx.Hash())
		return err
	})
	if err != nil {
		return fmt.Errorf("wait tx receipt error:%v,txHash:%v", err, tx.Hash())
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("record batch receipt failed,txHash:%v", tx.Hash())
	}
	return nil
}

func (r *RecordClient) NextBatchEpochIndex() (*big.Int, error) {
	nextBatchSubmissionIndex, err := r.record.NextBatchSubmissionIndex(nil)
	if err != nil {
		return nil, fmt.Errorf("get next batch submission index failed:%v", err)
	}
	return nextBatchSubmissionIndex, nil
}

type MockClient struct {
	recordRollupEpochInfo bindings.IRecordRollupEpochInfo
	nextRewardEpochIndex  *big.Int
	batchEpochIndex       *big.Int
}

func (m *MockClient) UploadRollupEpoch(recordRollupEpochInfos []bindings.IRecordRollupEpochInfo) error {
	m.recordRollupEpochInfo = recordRollupEpochInfos[len(recordRollupEpochInfos)-1]
	return nil
}

func (m *MockClient) LatestRollupEpoch() (*bindings.IRecordRollupEpochInfo, error) {
	return nil, nil
}

func (m *MockClient) UploadRewardsEpoch(recordRewardsEpochInfos []bindings.IRecordRewardEpochInfo) error {
	return nil
}

func (m *MockClient) NextRewardEpochIndex() (*big.Int, error) {
	if m.nextRewardEpochIndex == nil {
		return big.NewInt(1), nil
	}

	return m.nextRewardEpochIndex, nil
}

func (m *MockClient) UploadBatchEpoch(recordBatchSubmissions []bindings.IRecordBatchSubmission) error {
	m.batchEpochIndex = recordBatchSubmissions[len(recordBatchSubmissions)-1].Index
	return nil
}

func (m *MockClient) NextBatchEpochIndex() (*big.Int, error) {
	if m.batchEpochIndex == nil {
		return big.NewInt(1), nil
	}
	return m.batchEpochIndex, nil
}
