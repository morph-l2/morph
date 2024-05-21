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
	"morph-l2/bindings/predeploys"

	"github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/log"
)

var (
	RollupEpochTopic             = "RollupEpochUpdated(uint256, uint256)"
	RollupEpochTopicHash         = crypto.Keccak256Hash([]byte(RollupEpochTopic))
	SequencerSetUpdatedTopic     = "SequencerSetUpdated(address[], uint256)"
	SequencerSetUpdatedTopicHash = crypto.Keccak256Hash([]byte(SequencerSetUpdatedTopic))
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
	log.Info("generate rollup epoch", "startTime", startTime, "endBlockTime", endBlockTime)
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
		rollupEpochInfos = append(rollupEpochInfos, rollupEpochInfo)
		if endTime == endBlockTime {
			break
		}
		startTime = endTime
		index++
	}
	return rollupEpochInfos, nil
}

func (o *Oracle) recordRollupEpoch() error {
	epochIndex, err := o.record.NextRollupEpochIndex(nil)
	if err != nil {
		return err
	}
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
	if len(setsEpochs) != 0 {
		for _, setsEpoch := range setsEpochs {
			updateTime, err := o.GetUpdateTime(setsEpoch.EndBlock.Int64() - 1)
			if err != nil {
				return err
			}
			epochTime, err := o.gov.RollupEpoch(&bind.CallOpts{
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
		epochTime, err := o.gov.RollupEpoch(&bind.CallOpts{
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
		log.Info("rollup epoch infos length is zero", "startBlock", startBlock, "endBlock", endBlock)
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
	chainId, err := o.l2Client.ChainID(o.ctx)
	if err != nil {
		return err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(o.privKey, chainId)
	if err != nil {
		return err
	}
	tx, err := o.record.RecordRollupEpochs(opts, epochs)
	if err != nil {
		return err
	}
	log.Info("send record rollup epoch tx success", "txHash", tx.Hash().Hex(), "nonce", tx.Nonce())
	receipt, err := o.waitReceiptWithCtx(o.ctx, tx.Hash())
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
	rollupEpochLogs, err := o.fetchRollupEpochLog(o.ctx, start, end)
	if err != nil {
		return nil, err
	}
	var epochBlock []int
	for _, lg := range rollupEpochLogs {
		epochBlock = append(epochBlock, int(lg.BlockNumber))
	}
	sequencerSetUpdatedLogs, err := o.fetchSequencerSetUpdatedLog(o.ctx, start, end)
	if err != nil {
		return nil, err
	}
	for _, lg := range sequencerSetUpdatedLogs {
		epochBlock = append(epochBlock, int(lg.BlockNumber))
	}
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

func (o *Oracle) fetchRollupEpochLog(ctx context.Context, start, end uint64) ([]types.Log, error) {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0).SetUint64(start),
		ToBlock:   big.NewInt(0).SetUint64(end),
		Addresses: []common.Address{
			predeploys.GovAddr,
		},
		Topics: [][]common.Hash{
			{RollupEpochTopicHash},
		},
	}
	return o.l1Client.FilterLogs(ctx, query)
}

func (o *Oracle) fetchSequencerSetUpdatedLog(ctx context.Context, start, end uint64) ([]types.Log, error) {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0).SetUint64(start),
		ToBlock:   big.NewInt(0).SetUint64(end),
		Addresses: []common.Address{
			predeploys.SequencerAddr,
		},
		Topics: [][]common.Hash{
			{SequencerSetUpdatedTopicHash},
		},
	}
	return o.l1Client.FilterLogs(ctx, query)
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
