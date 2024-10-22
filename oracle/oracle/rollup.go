package oracle

import (
	"container/list"
	"context"
	"errors"
	"fmt"
	"math/big"

	"morph-l2/bindings/bindings"
	"morph-l2/oracle/types"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/log"
)

func (o *Oracle) getLatestUsingL1ChangePoint() types.ChangePoint {
	return types.ChangePoint{}
}

func (o *Oracle) getLatestUsingL2ChangePoint() types.ChangePoint {
	return types.ChangePoint{}
}

func (o *Oracle) generateEpochs(lastEpoch *bindings.IRecordRollupEpochInfo, syncedEndTime uint64) ([]bindings.IRecordRollupEpochInfo, error) {
	// Check that lastEpoch.EndTime is valid
	// TODO
	if len(o.ChangeCtx.ChangePoints) < 2 {
		return nil, fmt.Errorf("invalid change points length,expect >= 2,have %v", len(o.ChangeCtx.ChangePoints))
	}
	if lastEpoch.EndTime.Uint64() < o.ChangeCtx.ChangePoints[0].TimeStamp || lastEpoch.EndTime.Uint64() >= o.ChangeCtx.ChangePoints[len(o.ChangeCtx.ChangePoints)-1].TimeStamp {
		return nil, errors.New("lastEpoch.EndTime must be greater than or equal to changePoints[0].TimeStamp")
	}
	var epochs []bindings.IRecordRollupEpochInfo
	startTime := lastEpoch.EndTime.Uint64()

	// Initialize epochIndex starting from lastEpoch.Index + 1
	epochIndex := lastEpoch.Index.Uint64() + 1

	for i, point := range o.ChangeCtx.ChangePoints {
		// Check if the next ChangePoint exists
		if i+1 == len(o.ChangeCtx.ChangePoints) {
			break
		}
		nextPoint := o.ChangeCtx.ChangePoints[i+1]
		epochDuration := point.EpochInterval

		// Create epochs until the nextPoint's timestamp is less than endTime
		for {
			if startTime >= nextPoint.TimeStamp {
				break
			}
			endTime := startTime + epochDuration
			if endTime > nextPoint.TimeStamp {
				if nextPoint.ChangeType == types.L1ChangePoint || nextPoint.ChangeType == types.L2ChangePoint {
					endTime = nextPoint.TimeStamp
				} else {
					break
				}
			}

			// Ensure StartTime and EndTime are valid
			epochInfo := bindings.IRecordRollupEpochInfo{
				Index:     new(big.Int).SetUint64(epochIndex), // Set the current epoch index
				Submitter: point.Submitters[int((startTime-point.TimeStamp)/point.EpochInterval)%len(point.Submitters)],
				StartTime: new(big.Int).SetUint64(startTime), // Use SetUint64 for a copy
				EndTime:   new(big.Int).SetUint64(endTime),   // Use SetUint64 for a copy
				EndBlock:  big.NewInt(0),                     // Can set the corresponding EndBlock
			}
			epochs = append(epochs, epochInfo)
			if len(epochs) > types.MaxEpochCount {
				return epochs, nil
			}
			// Increment epochIndex for the next epoch
			epochIndex++
			// Update startTime for the next epoch
			startTime = endTime
		}
	}
	return epochs, nil
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

func (o *Oracle) fetchL1StakerRemoved(ctx context.Context, start, end uint64) ([]int, error) {
	opts := &bind.FilterOpts{
		Context: ctx,
		Start:   start,
		End:     &end,
	}
	iter, err := o.l1Staking.FilterStakersRemoved(opts)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := iter.Close(); err != nil {
			log.Info("L1StakingStakersRemovedIterator close failed", "error", err)
		}
	}()
	var blocks []int
	for iter.Next() {
		blocks = append(blocks, int(iter.Event.Raw.BlockNumber))
	}
	return blocks, nil
}

func (o *Oracle) fetchL1StakerAdded(ctx context.Context, start, end uint64) ([]int, error) {
	opts := &bind.FilterOpts{
		Context: ctx,
		Start:   start,
		End:     &end,
	}
	iter, err := o.l1Staking.FilterRegistered(opts)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := iter.Close(); err != nil {
			log.Info("L1StakingRegisteredIterator close failed", "error", err)
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
