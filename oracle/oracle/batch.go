package oracle

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/morph-l2/bindings/bindings"
	"github.com/morph-l2/node/derivation"
	"github.com/scroll-tech/go-ethereum"
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

func (o *Oracle) GetPreBatchRollupBlock(nextBatchSubmissionIndex *big.Int) (uint64, error) {
	if nextBatchSubmissionIndex.Uint64() == 0 {
		return o.cfg.StartBlock, nil
	}
	bs, err := o.record.BatchSubmissions(nil, nextBatchSubmissionIndex)
	if err != nil {
		return 0, err
	}
	return bs.RollupBlock.Uint64(), nil
}

func (o *Oracle) GetBatchSubmission(ctx context.Context, startBlock, endBlock uint64) ([]bindings.IRecordBatchSubmission, error) {
	rLogs, err := o.fetchRollupLog(ctx, startBlock, endBlock)
	if err != nil {
		return nil, err
	}
	//lastBatchIndex,err:= o.rollup.LastCommittedBatchIndex(nil)
	//if err != nil {
	//	return nil,err
	//}
	//o.l2Client.GetRollupBatchByIndex(o.ctx,)
	var recordBatchSubmissions []bindings.IRecordBatchSubmission
	for _, lg := range rLogs {
		tx, pending, err := o.l1Client.TransactionByHash(ctx, lg.TxHash)
		if err != nil {
			return nil, err
		}
		signer := types.NewLondonSigner(tx.ChainId())
		msg, err := tx.AsMessage(signer, tx.GasFeeCap())
		if err != nil {
			return nil, err
		}
		header, err := o.l1Client.HeaderByNumber(context.Background(), big.NewInt(int64(lg.BlockNumber)))
		if err != nil {
			return nil, err
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
		if err != nil {
			return nil, fmt.Errorf("unpack commitBatch error:%v", err)
		}
		rollupBatchData := args[0].(struct {
			Version                uint8     "json:\"version\""
			ParentBatchHeader      []uint8   "json:\"parentBatchHeader\""
			Chunks                 [][]uint8 "json:\"chunks\""
			SkippedL1MessageBitmap []uint8   "json:\"skippedL1MessageBitmap\""
			PrevStateRoot          [32]uint8 "json:\"prevStateRoot\""
			PostStateRoot          [32]uint8 "json:\"postStateRoot\""
			WithdrawalRoot         [32]uint8 "json:\"withdrawalRoot\""
			Signature              struct {
				Version   *big.Int   "json:\"version\""
				Signers   []*big.Int "json:\"signers\""
				Signature []uint8    "json:\"signature\""
			} "json:\"signature\""
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

		//if rollupCommitBatch.BatchIndex.Uint64() != highestBatchIndex+1 {
		//	return nil, fmt.Errorf("invalid batch,batch index discontinuity or batch storage blockNumber too small,bs.BlockNumber:%v,HighestBatch.BlockNumber:%v,bs.BatchIndex:%v,HighestBatch.BatchIndex:%v,txHash:%v", lg.BlockNumber, HighestBatch.BlockNumber, rollupCommitBatch.BatchIndex.Uint64(), HighestBatch.BatchIndex, lg.TxHash)
		//}
		log.Info("received new batch", "batch_index", rollupCommitBatch.BatchIndex.Uint64())
		recordBatchSubmission := bindings.IRecordBatchSubmission{
			Index:      rollupCommitBatch.BatchIndex,
			Submitter:  msg.From(),
			StartBlock: big.NewInt(int64(batchData.FirstBlockNumber())),
			EndBlock:   big.NewInt(int64(batchData.LastBlockNumber())),
			RollupTime: big.NewInt(int64(header.Time)),
		}
		recordBatchSubmissions = append(recordBatchSubmissions, recordBatchSubmission)
	}
	return recordBatchSubmissions, nil
}

func (o *Oracle) fetchRollupLog(ctx context.Context, start, end uint64) ([]types.Log, error) {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0).SetUint64(start),
		ToBlock:   big.NewInt(0).SetUint64(end),
		Addresses: []common.Address{
			o.rollupAddr,
		},
		Topics: [][]common.Hash{
			{RollupEventTopicHash},
		},
	}
	return o.l1Client.FilterLogs(ctx, query)
}

func (o *Oracle) GetNextBatchSubmissionIndex() (*big.Int, error) {
	return o.record.NextBatchSubmissionIndex(nil)
}

func (o *Oracle) submitRecord() {
	nextBatchSubmissionIndex, err := o.GetNextBatchSubmissionIndex()
	lastFinalized, err := o.rollup.LastFinalizedBatchIndex(nil)
	if err != nil {
		log.Error("get last finalized batch index failed ", "error", err)
		return
	}
	if nextBatchSubmissionIndex.Cmp(lastFinalized) > 0 {
		log.Info("already newest batch...")
		return
	}
	start, err := o.GetPreBatchRollupBlock(nextBatchSubmissionIndex)
	if err != nil {
		log.Error("get pre batch rollup block number failed", "error", err)
		return
	}
	end, err := o.l1Client.BlockNumber(o.ctx)
	if err != nil {
		log.Error("get latest block number failed", "error", err)
		return
	}
	if start+o.cfg.MaxSize < end {
		end = start + o.cfg.MaxSize - 1
	}
	batchSubmissions, err := o.GetBatchSubmission(context.Background(), start, end)
	tx, err := o.record.RecordFinalizedBatchSubmissions(nil, batchSubmissions)
	receipt, err := o.l2Client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {

	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		// TODO
	}
}
