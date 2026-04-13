package oracle

import (
	"container/list"
	"context"
	"fmt"
	"math"
	"math/big"
	"sort"
	"time"

	"morph-l2/bindings/bindings"
	"morph-l2/oracle/backoff"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/log"
)

var (
	MaxEpochCount = 50
)

type SequencerSetUpdateEpoch struct {
	Submitters []common.Address
	StartTime  *big.Int
	EndTime    *big.Int
	EndBlock   *big.Int
}

func (o *Oracle) generateRollupEpoch(index, startTime, rollupEpoch, updateTime, endBlock, endBlockTime, nextUpdateTime int64, sequencerSets []common.Address) ([]bindings.IRecordRollupEpochInfo, error) {
	var rollupEpochInfos []bindings.IRecordRollupEpochInfo
	if startTime == 0 {
		startTime = updateTime
	}
	epochsStart := startTime
	for {
		endTime := startTime + rollupEpoch
		if endTime > nextUpdateTime {
			endTime = nextUpdateTime
		}
		rollupEpochInfo := bindings.IRecordRollupEpochInfo{
			Index:     big.NewInt(index),
			Submitter: sequencerSets[(endTime-updateTime)/rollupEpoch%int64(len(sequencerSets))],
			StartTime: big.NewInt(startTime),
			EndTime:   big.NewInt(endTime),
			EndBlock:  big.NewInt(endBlock),
		}
		if endTime > endBlockTime {
			break
		}
		// TODO
		if o.rollupEpochMaxBlock == 1 && len(rollupEpochInfos) >= MaxEpochCount {
			rollupEpochInfo.EndBlock = big.NewInt(endBlock - 1)
			rollupEpochInfos = append(rollupEpochInfos, rollupEpochInfo)
			break
		}
		rollupEpochInfos = append(rollupEpochInfos, rollupEpochInfo)
		if endTime == endBlockTime {
			break
		}
		startTime = endTime
		index++
	}
	log.Info("generate rollup epoch", "startTime", epochsStart, "endBlockTime", endBlockTime, "epochLength", len(rollupEpochInfos))
	return rollupEpochInfos, nil
}

func (o *Oracle) recordRollupEpoch() error {
	epochIndex, err := o.record.NextRollupEpochIndex(nil)
	if err != nil {
		return err
	}
	o.metrics.SetRollupEpoch(epochIndex.Uint64() - 1)
	rollupEpoch, err := o.record.RollupEpochs(nil, new(big.Int).Sub(epochIndex, big.NewInt(1)))
	if err != nil {
		return err
	}
	startBlock := rollupEpoch.EndBlock.Uint64()
	blockNumber, err := o.l2Client.BlockNumber(o.ctx)
	if err != nil {
		return err
	}
	if startBlock+o.cfg.MinSize >= blockNumber {
		log.Info("too few blocks are newer than startBlock", "startBlock", startBlock, "latestBlock", blockNumber, "minSize", o.cfg.MinSize)
		time.Sleep(defaultSleepTime)
		return nil
	}
	endBlock := startBlock + o.rollupEpochMaxBlock
	if endBlock > blockNumber {
		endBlock = blockNumber
	}
	log.Info("record rollup epoch info start", "startBlock", startBlock, "endBlock", endBlock, "nextEpochIndex", epochIndex, "lastEpochInfo", rollupEpoch)
	setsEpochs, err := o.GetSequencerSetsEpoch(startBlock, endBlock)
	if err != nil {
		return err
	}
	var rollupEpochInfos []bindings.IRecordRollupEpochInfo
	var epochTime *big.Int
	if len(setsEpochs) != 0 {
		for _, setsEpoch := range setsEpochs {
			log.Info("received new sets change", "startTime", setsEpoch.StartTime, "endTime", setsEpoch.EndTime, "endBlock", setsEpoch.EndBlock)
			updateTime, err := o.GetUpdateTime(setsEpoch.EndBlock.Int64() - 1)
			if err != nil {
				return err
			}
			epochTime, err = o.gov.RollupEpoch(&bind.CallOpts{
				BlockNumber: big.NewInt(setsEpoch.EndBlock.Int64() - 1),
			})
			if err != nil {
				return err
			}
			epochs, err := o.generateRollupEpoch(epochIndex.Int64()+int64(len(rollupEpochInfos)), rollupEpoch.EndTime.Int64(), epochTime.Int64(), updateTime, setsEpoch.EndBlock.Int64(), setsEpoch.EndTime.Int64(), setsEpoch.EndTime.Int64(), setsEpoch.Submitters)
			if err != nil {
				return err
			}
			rollupEpochInfos = append(rollupEpochInfos, epochs...)
		}
	} else {
		updateTime, err := o.GetUpdateTime(int64(endBlock))
		if err != nil {
			return fmt.Errorf("get update time error:%v", err)
		}
		epochTime, err = o.gov.RollupEpoch(&bind.CallOpts{
			BlockNumber: big.NewInt(int64(endBlock)),
		})
		if err != nil {
			return fmt.Errorf("get rollup epoch time error:%v", err)
		}
		header, err := o.l2Client.HeaderByNumber(o.ctx, big.NewInt(int64(endBlock)))
		if err != nil {
			return fmt.Errorf("get header by number error:%v", err)
		}
		sets, err := o.sequencer.GetSequencerSet2(&bind.CallOpts{
			BlockNumber: big.NewInt(int64(endBlock)),
		})
		if err != nil {
			return fmt.Errorf("get sequencer set error:%v", err)
		}
		epochs, err := o.generateRollupEpoch(epochIndex.Int64(), rollupEpoch.EndTime.Int64(), epochTime.Int64(), updateTime, int64(endBlock), int64(header.Time), math.MaxInt64, sets)
		if err != nil {
			return fmt.Errorf("gernerate rollup epoch info error:%v", err)
		}
		rollupEpochInfos = append(rollupEpochInfos, epochs...)
	}
	if len(rollupEpochInfos) == 0 {
		log.Info("rollup epoch infos length is zero", "startBlock", startBlock, "endBlock", endBlock, "rollupEpochMaxBlock", o.rollupEpochMaxBlock, "epochTime", epochTime)
		time.Sleep(defaultSleepTime)
		return nil
	}
	log.Info("submit rollup epoch infos", "startBlock", startBlock, "endBlock", endBlock, "infoLength", len(rollupEpochInfos))
	err = o.submitRollupEpoch(rollupEpochInfos)
	if err != nil {
		if len(rollupEpochInfos) > 50 {
			if o.cfg.MinSize*2 <= o.rollupEpochMaxBlock {
				o.rollupEpochMaxBlock -= o.cfg.MinSize
			} else {
				o.rollupEpochMaxBlock = o.rollupEpochMaxBlock / 2
			}
		}
		return fmt.Errorf("submit rollup epoch info error:%v,rollupEpochMaxBlock:%v", err, o.rollupEpochMaxBlock)
	}
	if o.rollupEpochMaxBlock+o.cfg.MinSize <= o.cfg.MaxSize {
		o.rollupEpochMaxBlock += o.cfg.MinSize
	}

	log.Info("submit rollup epoch info success", "rollupEpochMaxBlock", o.rollupEpochMaxBlock)
	return nil
}

func (o *Oracle) submitRollupEpoch(epochs []bindings.IRecordRollupEpochInfo) error {
	callData, err := o.recordAbi.Pack("recordRollupEpochs", epochs)
	if err != nil {
		return err
	}
	tx, err := o.newRecordTxAndSign(callData)
	if err != nil {
		return err
	}
	log.Info("send record rollup epoch tx success", "txHash", tx.Hash().Hex(), "nonce", tx.Nonce())
	var receipt *types.Receipt
	err = backoff.Do(30, backoff.Exponential(), func() error {
		var err error
		receipt, err = o.waitReceiptWithCtx(o.ctx, tx.Hash())
		return err
	})
	if err != nil {
		return fmt.Errorf("receipt record rollup epochs error:%v", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("record rollup epochs not success")
	}
	log.Info("wait receipt success", "txHash", tx.Hash())
	return nil
}

func (o *Oracle) GetUpdateTime(blockNumber int64) (int64, error) {
	updateTime, err := o.sequencer.UpdateTime(&bind.CallOpts{
		BlockNumber: big.NewInt(blockNumber),
	})
	if err != nil {
		return 0, err
	}
	epochUpdateTime, err := o.gov.RollupEpochUpdateTime(&bind.CallOpts{
		BlockNumber: big.NewInt(blockNumber),
	})
	if err != nil {
		return 0, err
	}
	header, err := o.l2Client.HeaderByNumber(o.ctx, big.NewInt(1))
	if err != nil {
		return 0, err
	}
	if updateTime.Cmp(epochUpdateTime) <= 0 {
		updateTime = epochUpdateTime
	}
	if updateTime.Uint64() <= header.Time {
		updateTime = big.NewInt(int64(header.Time))
	}
	return updateTime.Int64(), nil
}

func (o *Oracle) GetSequencerSetsEpoch(start, end uint64) ([]SequencerSetUpdateEpoch, error) {
	var epochBlock []int
	rollupEpochUpdated, err := o.fetchRollupEpochUpdated(o.ctx, start, end)
	if err != nil {
		return nil, err
	}
	epochBlock = append(epochBlock, rollupEpochUpdated...)
	sequencerSetUpdated, err := o.fetchSequencerSetUpdated(o.ctx, start, end)
	if err != nil {
		return nil, err
	}
	epochBlock = append(epochBlock, sequencerSetUpdated...)
	sortedBlocks := removeDuplicatesAndSort(epochBlock)
	sort.Ints(sortedBlocks)
	var setsEpochInfos []SequencerSetUpdateEpoch
	for _, eb := range sortedBlocks {
		header, err := o.l2Client.HeaderByNumber(o.ctx, big.NewInt(int64(eb)))
		if err != nil {
			return nil, err
		}
		sequencerSets, err := o.sequencer.GetSequencerSet2(&bind.CallOpts{
			BlockNumber: big.NewInt(int64(eb - 1)),
		})
		if err != nil {
			return nil, err
		}
		lastTime, err := o.GetUpdateTime(header.Number.Int64() - 1)
		if err != nil {
			return nil, err
		}
		epochInfo := SequencerSetUpdateEpoch{
			Submitters: sequencerSets,
			StartTime:  big.NewInt(lastTime),
			EndTime:    big.NewInt(int64(header.Time)),
			EndBlock:   header.Number,
		}
		setsEpochInfos = append(setsEpochInfos, epochInfo)
	}
	return setsEpochInfos, nil
}

func (o *Oracle) fetchRollupEpochUpdated(ctx context.Context, start, end uint64) ([]int, error) {
	opts := &bind.FilterOpts{
		Context: ctx,
		Start:   start,
		End:     &end,
	}
	iter, err := o.gov.FilterRollupEpochUpdated(opts)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := iter.Close(); err != nil {
			log.Info("GovRollupEpochUpdatedIterator close failed", "error", err)
		}
	}()
	var blocks []int
	for iter.Next() {
		blocks = append(blocks, int(iter.Event.Raw.BlockNumber))
	}
	return blocks, nil
}

func (o *Oracle) fetchSequencerSetUpdated(ctx context.Context, start, end uint64) ([]int, error) {
	opts := &bind.FilterOpts{
		Context: ctx,
		Start:   start,
		End:     &end,
	}
	iter, err := o.sequencer.FilterSequencerSetUpdated(opts)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := iter.Close(); err != nil {
			log.Info("SequencerSequencerSetUpdatedIterator close failed", "error", err)
		}
	}()
	var blocks []int
	for iter.Next() {
		blocks = append(blocks, int(iter.Event.Raw.BlockNumber))
	}
	return blocks, nil
}

type set struct {
	list *list.List
}

func newSet() *set {
	return &set{list.New()}
}

func (s *set) add(value int) {
	for e := s.list.Front(); e != nil; e = e.Next() {
		if e.Value == value {
			return
		}
	}
	s.list.PushBack(value)
}

func (s *set) values() []int {
	values := make([]int, 0, s.list.Len())
	for e := s.list.Front(); e != nil; e = e.Next() {
		values = append(values, e.Value.(int))
	}
	return values
}

func removeDuplicatesAndSort(arr []int) []int {
	s := newSet()
	for _, v := range arr {
		s.add(v)
	}
	return s.values()
}
