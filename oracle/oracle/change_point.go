package oracle

import (
	"fmt"
	"math/big"
	"sort"
	"time"

	"morph-l2/oracle/types"

	"github.com/morph-l2/go-ethereum/accounts/abi/bind"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/log"
)

//func (o *Oracle) queryActiveStakersByTime(timestamp uint64) ([]common.Address, error) {
//	for _, activeStakers := range o.ChangeCtx.ActiveStakersByTime {
//		if timestamp >= activeStakers.TimeStamp {
//			return activeStakers.Addresses, nil
//		}
//	}
//	return nil, fmt.Errorf("not found,invalid timestamp,expect bigger or equal than %v but have :%v ", o.ChangeCtx.ActiveStakersByTime[0].TimeStamp, timestamp)
//}

func (o *Oracle) syncL1ChangePoint(start, end, startTime, endTime uint64) error {
	var epochBlock []int
	stakerAddedPoint, err := o.fetchL1StakerAdded(o.ctx, start, end)
	if err != nil {
		return err
	}
	epochBlock = append(epochBlock, stakerAddedPoint...)
	stakerRemoved, err := o.fetchL1StakerRemoved(o.ctx, start, end)
	if err != nil {
		return err
	}
	epochBlock = append(epochBlock, stakerRemoved...)
	sortedBlocks := removeDuplicatesAndSort(epochBlock)
	sort.Ints(sortedBlocks)
	for _, sortBlock := range sortedBlocks {
		header, err := o.l1Client.HeaderByNumber(o.ctx, big.NewInt(int64(sortBlock)))
		if err != nil {
			return err
		}
		activeStakers, err := o.l1Staking.GetActiveStakers(&bind.CallOpts{
			BlockNumber: header.Number,
		})
		if err != nil {
			return err
		}
		l1Staker := types.ActiveStakers{
			TimeStamp:   header.Time,
			Addresses:   activeStakers,
			BlockNumber: uint64(sortBlock),
		}
		// TODO
		if header.Time > endTime {
			break
		}
		if header.Time > o.ChangeCtx.ActiveStakersByTime[len(o.ChangeCtx.ActiveStakersByTime)-1].TimeStamp {
			o.ChangeCtx.ActiveStakersByTime = append(o.ChangeCtx.ActiveStakersByTime, l1Staker)
			o.ChangeCtx.L1Synced = uint64(sortBlock)
		}

	}
	var changePoints []types.ChangePoint
	for _, eb := range stakerRemoved {
		header, err := o.l1Client.HeaderByNumber(o.ctx, big.NewInt(int64(eb)))
		if err != nil {
			return err
		}
		if header.Time < startTime {
			continue
		}
		if header.Time > endTime {
			break
		}
		changePoint := types.ChangePoint{
			//Submitters: activeSequencerSet,
			TimeStamp:   header.Time,
			BlockNumber: header.Number.Uint64(),
			ChangeType:  types.L1ChangePoint,
		}
		if header.Time > endTime {
			break
		}
		if len(changePoints) == 0 || header.Time > changePoints[len(changePoints)-1].TimeStamp {
			changePoints = append(changePoints, changePoint)
		}
		changePoints = append(changePoints, changePoint)
	}
	if len(stakerRemoved) == 0 {
		o.ChangeCtx.L1Synced = end
	}
	o.insertL1AndMerge(changePoints)
	return nil
}

func (o *Oracle) syncL2ChangePoint(start, end uint64) ([]types.ChangePoint, error) {
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
	var changePoints []types.ChangePoint
	for _, eb := range sortedBlocks {
		header, err := o.l2Client.HeaderByNumber(o.ctx, big.NewInt(int64(eb)))
		if err != nil {
			return nil, err
		}
		sequencerSets, err := o.sequencer.GetCurrentSequencerSet(&bind.CallOpts{
			BlockNumber: big.NewInt(int64(eb)),
		})
		if err != nil {
			return nil, err
		}
		epochInterval, err := o.gov.RollupEpoch(&bind.CallOpts{
			BlockNumber: header.Number,
		})
		if err != nil {
			return nil, err
		}
		changePoint := types.ChangePoint{
			Submitters:    sequencerSets,
			EpochInterval: epochInterval.Uint64(),
			TimeStamp:     header.Time,
			ChangeType:    types.L2ChangePoint,
		}
		changePoints = append(changePoints, changePoint)
		o.ChangeCtx.L2Sequencers = append(o.ChangeCtx.L2Sequencers, types.L2Sequencer{
			Addresses:     sequencerSets,
			EpochInterval: epochInterval.Uint64(),
			TimeStamp:     header.Time,
		})
	}
	return changePoints, nil
}

func (o *Oracle) recordRollupEpoch() error {
	l2Start := o.ChangeCtx.L2synced + 1
	l2Latest, err := o.l2Client.BlockNumber(o.ctx)
	if err != nil {
		return err
	}
	l2End := l2Latest
	if l2Start+o.rollupEpochMaxBlock < l2Latest {
		l2End = l2Start + o.rollupEpochMaxBlock - 1
	}
	epochIndex, err := o.record.NextRollupEpochIndex(nil)
	if err != nil {
		return err
	}
	lastEpoch, err := o.record.RollupEpochs(nil, epochIndex.Sub(epochIndex, big.NewInt(1)))
	if err != nil {
		return err
	}
	//o.PruneChangeCtx()
	changePointIndex := 0
	for i, cps := range o.ChangeCtx.ChangePoints {
		if cps.TimeStamp > lastEpoch.EndTime.Uint64() {
			changePointIndex = i
			break
		}
	}

	// Check if a valid index was found
	if changePointIndex > 0 {
		// Trim the slice to keep only the points after the found index
		o.ChangeCtx.ChangePoints = o.ChangeCtx.ChangePoints[changePointIndex-1:]
	}

	startHeader, err := o.l2Client.HeaderByNumber(o.ctx, big.NewInt(int64(l2Start)))
	if err != nil {
		return err
	}
	endHeader, err := o.l2Client.HeaderByNumber(o.ctx, big.NewInt(int64(l2End)))
	if err != nil {
		return err
	}
	startTime := startHeader.Time
	endTime := endHeader.Time
	// clean fake point
	if o.ChangeCtx.ChangePoints[len(o.ChangeCtx.ChangePoints)-1].ChangeType != types.L1ChangePoint && o.ChangeCtx.ChangePoints[len(o.ChangeCtx.ChangePoints)-1].ChangeType != types.L2ChangePoint {
		o.ChangeCtx.ChangePoints = o.ChangeCtx.ChangePoints[:len(o.ChangeCtx.ChangePoints)-1]
	}
	_, err = o.syncL2ChangePoint(l2Start, l2End)
	if err != nil {
		return err
	}
	l1Start := o.ChangeCtx.L1Synced + 1
	l1Latest, err := o.l1Client.BlockNumber(o.ctx)
	if err != nil {
		return err
	}
	l1End := l1Latest
	if l1Start+o.rollupEpochMaxBlock < l1Latest {
		l1End = l1Start + o.rollupEpochMaxBlock - 1
	}
	if err = o.syncL1ChangePoint(l1Start, l1End, startTime, endTime); err != nil {
		return err
	}
	// insert a fake change point
	if o.ChangeCtx.ChangePoints[len(o.ChangeCtx.ChangePoints)-1].TimeStamp < endTime {
		o.ChangeCtx.ChangePoints = append(o.ChangeCtx.ChangePoints, types.ChangePoint{TimeStamp: endTime})
	}
	// TODO

	if err = o.db.WriteLatestChangeContext(o.ChangeCtx); err != nil {
		return err
	}
	epochs, err := o.generateEpochs(lastEpoch, endTime)
	if err != nil {
		return err
	}
	if len(epochs) == 0 {
		time.Sleep(defaultSleepTime)
		log.Info("rollup epoch count too small", "startTime", lastEpoch.StartTime, "index", lastEpoch.Index)
		return nil
	}
	log.Info("submit rollup epoch infos", "l1Start", l1Start, "l1End", l1End, "l2Start", l2Start, "l2End", l2End, "infoLength", len(epochs))
	err = o.submitRollupEpoch(epochs)
	if err != nil {
		if len(epochs) > 50 {
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

func (o *Oracle) GetSequencerSetsByTime(t uint64) ([]common.Address, error) {
	for i, p := range o.ChangeCtx.L2Sequencers {
		if i+1 > len(o.ChangeCtx.L2Sequencers) {
			break
		}
		if t > p.TimeStamp && t < o.ChangeCtx.L2Sequencers[i+1].TimeStamp {
			return p.Addresses, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func (o *Oracle) GetActiveStakersByTime(t uint64) ([]common.Address, error) {
	for i, as := range o.ChangeCtx.ActiveStakersByTime {
		if i+1 > len(o.ChangeCtx.ActiveStakersByTime) {
			break
		}
		if t > as.TimeStamp && t < o.ChangeCtx.ActiveStakersByTime[i+1].TimeStamp {
			return as.Addresses, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func findIntersection(arr1, arr2 []common.Address) []common.Address {
	addressMap := make(map[common.Address]struct{})
	for _, addr := range arr1 {
		addressMap[addr] = struct{}{}
	}
	intersection := []common.Address{}
	for _, addr := range arr2 {
		if _, exists := addressMap[addr]; exists {
			intersection = append(intersection, addr)
		}
	}

	return intersection
}

// InsertL1AndMerge inserts new ChangePoints into the old array while maintaining order
func (o *Oracle) insertL1AndMerge(newPoints []types.ChangePoint) {
	for _, newPoint := range newPoints {
		// Flag to check if the new element has been inserted
		inserted := false
		for i := 0; i <= len(o.ChangeCtx.ChangePoints); i++ {
			// If we've reached the end of the array or the new element's timestamp is less than the current element
			if i == len(o.ChangeCtx.ChangePoints) || newPoint.TimeStamp < o.ChangeCtx.ChangePoints[i].TimeStamp {
				// If ChangeType is 1, assign the previous element's Submitters
				if newPoint.ChangeType == 1 && i > 0 {
					newPoint.Submitters = o.ChangeCtx.ChangePoints[i-1].Submitters
					newPoint.EpochInterval = o.ChangeCtx.ChangePoints[i-1].EpochInterval
				}
				// Insert the new element
				o.ChangeCtx.ChangePoints = append(o.ChangeCtx.ChangePoints[:i], append([]types.ChangePoint{newPoint}, o.ChangeCtx.ChangePoints[i:]...)...)
				inserted = true
				break
			}
		}
		// If the new element wasn't inserted, it means it's the largest element, so append it to the end
		if !inserted {
			newPoint.Submitters = o.ChangeCtx.ChangePoints[len(o.ChangeCtx.ChangePoints)-1].Submitters
			newPoint.EpochInterval = o.ChangeCtx.ChangePoints[len(o.ChangeCtx.ChangePoints)-1].EpochInterval
			o.ChangeCtx.ChangePoints = append(o.ChangeCtx.ChangePoints, newPoint)
		}
	}
}

func (o *Oracle) insertL2AndMerge(newPoints []types.ChangePoint) error {
	if len(newPoints) == 0 {
		return nil
	}
	if newPoints[0].TimeStamp <= o.ChangeCtx.ChangePoints[len(o.ChangeCtx.ChangePoints)-1].TimeStamp {
		return fmt.Errorf("invalid l2 newPoint,expect bigger than %v but have %v",
			o.ChangeCtx.ChangePoints[len(o.ChangeCtx.ChangePoints)-1].TimeStamp,
			newPoints[0].TimeStamp)
	}
	o.ChangeCtx.ChangePoints = append(o.ChangeCtx.ChangePoints, newPoints...)
	return nil
}
