package oracle

import (
	"container/list"
	"context"
	"fmt"
	"math/big"
	"sort"

	"github.com/morph-l2/bindings/bindings"
	"github.com/morph-l2/bindings/predeploys"
	"github.com/scroll-tech/go-ethereum"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/log"
)

var (
	RollupEpochTopic     = "RollupEpochUpdated(uint256, uint256)"
	RollupEpochTopicHash = crypto.Keccak256Hash([]byte(RollupEpochTopic))

	SequencerSetUpdatedTopic     = "SequencerSetUpdated(address[], uint256)"
	SequencerSetUpdatedTopicHash = crypto.Keccak256Hash([]byte(SequencerSetUpdatedTopic))
)

func (o *Oracle) recordRollupEpoch() error {
	epochIndex, err := o.record.NextRollupEpochIndex(nil)
	if err != nil {
		return err
	}
	rollupEpoch, err := o.record.RollupEpochs(nil, new(big.Int).Sub(epochIndex, big.NewInt(1)))
	if err != nil {
		return err
	}
	var startBlock uint64
	recordRollupEpochInfos, err := o.GetRollupEpoch(startBlock, startBlock, rollupEpoch.EndTime.Uint64(), rollupEpoch.Index.Uint64())
	chainId, err := o.l2Client.ChainID(o.ctx)
	if err != nil {
		return err
	}
	opts, err := bind.NewKeyedTransactorWithChainID(o.privKey, chainId)
	tx, err := o.record.RecordRollupEpochs(opts, recordRollupEpochInfos)
	log.Info("send record reward tx success", "txHash", tx.Hash().Hex(), "nonce", tx.Nonce())
	receipt, err := o.l2Client.TransactionReceipt(o.ctx, tx.Hash())
	if err != nil {
		return fmt.Errorf("receipt record reward epochs error:%v", err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		return fmt.Errorf("record reward epochs not success")
	}
	return nil
}

func (o *Oracle) GetRollupEpoch(start, end uint64, lastTime uint64, lastIndex uint64) ([]bindings.IRecordRollupEpochInfo, error) {
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
	//var latestIndex uint64 // TODO
	//var lastTime uint64
	var epochInfos []bindings.IRecordRollupEpochInfo
	for _, eb := range sortedBlocks {
		lastIndex += 1
		header, err := o.l2Client.HeaderByNumber(o.ctx, big.NewInt(int64(eb)))
		if err != nil {

		}
		epochInfo := bindings.IRecordRollupEpochInfo{
			Index: big.NewInt(int64(lastIndex)),

			//Submitter common.Address
			StartTime: big.NewInt(int64(lastTime)),
			EndTime:   big.NewInt(int64(header.Time)),
		}
		lastTime = header.Time
		epochInfos = append(epochInfos, epochInfo)
	}
	// TODO delete
	log.Info("epoch info", "epochInfo", epochInfos)
	return epochInfos, nil
}

func GetEpochBlock() (*big.Int, error) {
	return big.NewInt(0), nil
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
