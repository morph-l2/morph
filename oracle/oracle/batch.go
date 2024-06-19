package oracle

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"morph-l2/bindings/bindings"
	"morph-l2/node/derivation"

	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/eth"
	"github.com/scroll-tech/go-ethereum/log"
)

var (
	RollupEventTopic     = "CommitBatch(uint256,bytes32)"
	RollupEventTopicHash = crypto.Keccak256Hash([]byte(RollupEventTopic))
)

type BatchInfoMap map[common.Hash][]BatchInfo
type RollupBatch struct {
	TxCount uint64
}

type BatchInfo struct {
	BatchIndex    uint64
	L1BlockNumber uint64
	L1TxHash      common.Hash
	L2BlockNumber uint64
	L2BlockCount  uint64
	L2TxCount     uint64
}

func (o *Oracle) GetStartBlock(nextBatchSubmissionIndex *big.Int) (uint64, error) {
	if nextBatchSubmissionIndex.Uint64() == 1 {
		return o.cfg.StartBlock, nil
	}
	bs, err := o.record.BatchSubmissions(nil, new(big.Int).Sub(nextBatchSubmissionIndex, big.NewInt(1)))
	if err != nil {
		return 0, err
	}
	return bs.RollupBlock.Uint64() + 1, nil
}

func (o *Oracle) GetBatchSubmission(ctx context.Context, startBlock uint64) ([]bindings.IRecordBatchSubmission, error) {
	var rLogs []types.Log
	for {
		endBlock := startBlock + o.cfg.MaxSize
		header, err := o.l1Client.HeaderByNumber(o.ctx, nil)
		if err != nil {
			return nil, fmt.Errorf("get latest header error:%v", err)
		}
		if startBlock >= header.Number.Uint64() {
			time.Sleep(defaultSleepTime)
			continue
		}
		if endBlock >= header.Number.Uint64() {
			endBlock = header.Number.Uint64()
		}
		rLogs, err = o.fetchRollupLog(ctx, startBlock, endBlock)
		if err != nil {
			return nil, fmt.Errorf("fetch rollupLog error:%v", err)
		}
		if len(rLogs) > 1 {
			break
		}
		startBlock = endBlock + 1
	}

	var recordBatchSubmissions []bindings.IRecordBatchSubmission
	for _, lg := range rLogs {
		tx, pending, err := o.l1Client.TransactionByHash(ctx, lg.TxHash)
		if err != nil {
			return nil, fmt.Errorf("get transaction by hash error:%v", err)
		}
		signer := types.NewLondonSignerWithEIP4844(tx.ChainId())
		msg, err := tx.AsMessage(signer, tx.GasFeeCap())
		if err != nil {
			return nil, err
		}
		header, err := o.l1Client.HeaderByNumber(context.Background(), big.NewInt(int64(lg.BlockNumber)))
		if err != nil {
			return nil, fmt.Errorf("get header by number error:%v", err)
		}
		if pending {
			return nil, errors.New("pending transaction")
		}
		abi, err := bindings.RollupMetaData.GetAbi()
		if err != nil {
			return nil, err
		}

		rollupCommitBatch, parseErr := o.rollup.ParseCommitBatch(lg)
		if parseErr != nil {
			log.Error("get l2 BlockNumber", "err", err)
			return nil, parseErr
		}
		args, err := abi.Methods["commitBatch"].Inputs.Unpack(tx.Data()[4:])
		if err != nil {
			if rollupCommitBatch.BatchIndex.Uint64() == 0 {
				continue
			}
			log.Error("fetch batch info failed", "txHash", lg.TxHash, "blockNumber", lg.BlockNumber, "error", err)
			return nil, err
		}
		rollupBatchData := args[0].(struct {
			Version                uint8     "json:\"version\""
			ParentBatchHeader      []uint8   "json:\"parentBatchHeader\""
			Chunks                 [][]uint8 "json:\"chunks\""
			SkippedL1MessageBitmap []uint8   "json:\"skippedL1MessageBitmap\""
			PrevStateRoot          [32]uint8 "json:\"prevStateRoot\""
			PostStateRoot          [32]uint8 "json:\"postStateRoot\""
			WithdrawalRoot         [32]uint8 "json:\"withdrawalRoot\""
		})

		var chunks []hexutil.Bytes
		for _, chunk := range rollupBatchData.Chunks {
			chunks = append(chunks, chunk)
		}
		batch := eth.RPCRollupBatch{
			Version:                uint(rollupBatchData.Version),
			ParentBatchHeader:      rollupBatchData.ParentBatchHeader,
			Chunks:                 chunks,
			SkippedL1MessageBitmap: rollupBatchData.SkippedL1MessageBitmap,
			PrevStateRoot:          common.BytesToHash(rollupBatchData.PrevStateRoot[:]),
			PostStateRoot:          common.BytesToHash(rollupBatchData.PostStateRoot[:]),
			WithdrawRoot:           common.BytesToHash(rollupBatchData.WithdrawalRoot[:]),
		}
		var batchData derivation.BatchInfo
		if err = batchData.ParseBatch(batch); err != nil {
			return nil, fmt.Errorf("parse batch error:%v", err)
		}
		log.Info("received new batch", "batch_index", rollupCommitBatch.BatchIndex.Uint64())
		recordBatchSubmission := bindings.IRecordBatchSubmission{
			Index:       rollupCommitBatch.BatchIndex,
			Submitter:   msg.From(),
			StartBlock:  big.NewInt(int64(batchData.FirstBlockNumber())),
			EndBlock:    big.NewInt(int64(batchData.LastBlockNumber())),
			RollupTime:  big.NewInt(int64(header.Time)),
			RollupBlock: big.NewInt(int64(lg.BlockNumber)),
		}
		recordBatchSubmissions = append(recordBatchSubmissions, recordBatchSubmission)
	}
	return recordBatchSubmissions, nil
}

func (o *Oracle) fetchRollupLog(ctx context.Context, start, end uint64) ([]types.Log, error) {
	opts := &bind.FilterOpts{
		Context: ctx,
		Start:   start,
		End:     &end,
	}
	iter, err := o.rollup.FilterCommitBatch(opts, nil, nil)
	if err != nil {
		return nil, err
	}
	defer iter.Close()
	var logs []types.Log
	for iter.Next() {
		logs = append(logs, iter.Event.Raw)
	}
	return logs, nil
}

func (o *Oracle) GetNextBatchSubmissionIndex() (*big.Int, error) {
	return o.record.NextBatchSubmissionIndex(nil)
}

func (o *Oracle) LastBatchIndex(opts *bind.CallOpts) (*big.Int, error) {
	if o.isFinalized {
		return o.rollup.LastFinalizedBatchIndex(opts)

	}
	return o.rollup.LastCommittedBatchIndex(opts)
}

func (o *Oracle) submitRecord() error {
	nextBatchSubmissionIndex, err := o.GetNextBatchSubmissionIndex()
	if err != nil {
		return fmt.Errorf("get next batch submission index failed:%v", err)
	}
	o.metrics.SetBatchEpoch(nextBatchSubmissionIndex.Uint64() - 1)
	lastBatchIndex, err := o.LastBatchIndex(nil)
	if err != nil {
		return fmt.Errorf("get last finalized batch index error:%v", err)
	}
	if nextBatchSubmissionIndex.Cmp(lastBatchIndex) > 0 {
		log.Info("already newest batch submission...", "lastBatchIndex", lastBatchIndex, "nextBatchSubmissionIndex", nextBatchSubmissionIndex)
		time.Sleep(defaultSleepTime)
		return nil
	}
	start, err := o.GetStartBlock(nextBatchSubmissionIndex)
	if err != nil {
		log.Error("get pre batch rollup block number failed", "error", err)
		return fmt.Errorf("get pre batch rollup block number error:%v", err)
	}
	batchSubmissions, err := o.GetBatchSubmission(context.Background(), start)
	if err != nil {
		return fmt.Errorf("get batch submission error:%v", err)
	}
	chainId, err := o.l2Client.ChainID(o.ctx)
	if err != nil {
		return fmt.Errorf("get chain id error:%v", err)
	}
	opts, err := bind.NewKeyedTransactorWithChainID(o.privKey, chainId)
	if err != nil {
		return fmt.Errorf("new keyed transaction error:%v", err)
	}
	tx, err := o.record.RecordFinalizedBatchSubmissions(opts, batchSubmissions)
	if err != nil {
		return fmt.Errorf("record finalized batch error:%v", err)
	}
	receipt, err := o.waitReceiptWithCtx(o.ctx, tx.Hash())
	if err != nil {
		log.Error("tx receipt failed", "error", err)
		return fmt.Errorf("wait tx receipt error:%v", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("record batch not success")
	}
	return nil
}
