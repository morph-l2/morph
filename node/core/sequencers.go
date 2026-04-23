package node

import (
	"bytes"
	"math/big"
	"slices"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

const tmKeySize = ed25519.PubKeySize

func (e *Executor) sequencerSetUpdates(height uint64) ([][]byte, error) {
	seqHash, err := e.sequencerCaller.SequencerSetVerifyHash(nil)
	if err != nil {
		return nil, err
	}
	if e.currentSeqHash != nil && bytes.Equal(e.currentSeqHash[:], seqHash[:]) {
		return e.nextValidators, nil
	}

	sequencerSet0, err := e.sequencerCaller.GetSequencerSet0(nil)
	if err != nil {
		return nil, err
	}
	sequencerSet1, err := e.sequencerCaller.GetSequencerSet1(nil)
	if err != nil {
		return nil, err
	}
	sequencerSet2, err := e.sequencerCaller.GetSequencerSet2(nil)
	if err != nil {
		return nil, err
	}

	cache := make(map[common.Address]struct{})
	requestAddrs := make([]common.Address, 0)
	for _, addr := range append(append(sequencerSet0, sequencerSet1...), sequencerSet2...) {
		if _, ok := cache[addr]; !ok {
			cache[addr] = struct{}{}
			requestAddrs = append(requestAddrs, addr)
		}
	}
	stakesInfo, err := e.l2StakingCaller.GetStakesInfo(nil, requestAddrs)
	if err != nil {
		e.logger.Error("failed to GetStakesInfo", "error", err)
		return nil, err
	}

	seqTmKeySet := make(map[[tmKeySize]byte]struct{}, len(stakesInfo))
	nextValidators := make([][]byte, 0, len(sequencerSet2))
	for i := range stakesInfo {
		// sequencerSet2 is the latest updated sequencer set which is considered as the next validator set for tendermint
		if slices.Contains(sequencerSet2, stakesInfo[i].Addr) {
			nextValidators = append(nextValidators, stakesInfo[i].TmKey[:])
		}
		seqTmKeySet[stakesInfo[i].TmKey] = struct{}{}
	}

	e.logger.Info("sequencers updates, sequencer verified hash changed", "height", height)
	e.seqTmKeySet = seqTmKeySet
	e.nextValidators = nextValidators
	e.currentSeqHash = &seqHash
	return nextValidators, nil
}

func (e *Executor) batchParamsUpdates(height uint64) (*tmproto.BatchParams, error) {
	var (
		batchBlockInterval, batchTimeout *big.Int
		err                              error
	)

	if batchBlockInterval, err = e.govCaller.BatchBlockInterval(nil); err != nil {
		return nil, err
	}
	if batchTimeout, err = e.govCaller.BatchTimeout(nil); err != nil {
		return nil, err
	}

	changed := e.batchParams.BlocksInterval != batchBlockInterval.Int64() ||
		int64(e.batchParams.Timeout.Seconds()) != batchTimeout.Int64()

	if changed {
		e.batchParams.BlocksInterval = batchBlockInterval.Int64()
		e.batchParams.Timeout = time.Duration(batchTimeout.Int64() * int64(time.Second))
		e.logger.Info("batch params changed", "height", height,
			"batchBlockInterval", batchBlockInterval.Int64(),
			"batchTimeout", batchTimeout.Int64())
		return &tmproto.BatchParams{
			BlocksInterval: batchBlockInterval.Int64(),
			Timeout:        time.Duration(batchTimeout.Int64() * int64(time.Second)),
		}, nil
	}
	return nil, nil
}

func (e *Executor) updateSequencerSet(height uint64) ([][]byte, error) {
	validatorUpdates, err := e.sequencerSetUpdates(height)
	if err != nil {
		e.logger.Error("failed to get sequencer set from geth", "err", err)
		return nil, err
	}
	var tmPKBz [tmKeySize]byte
	copy(tmPKBz[:], e.tmPubKey)

	_, isSequencer := e.seqTmKeySet[tmPKBz]
	if !e.isSequencer && isSequencer {
		e.logger.Info("I am a sequencer, start to launch syncer")
		if e.syncer == nil {
			syncer, err := e.newSyncerFunc()
			if err != nil {
				e.logger.Error("failed to create syncer", "error", err)
				return nil, err
			}
			e.syncer = syncer
			e.l1MsgReader = syncer // syncer works as l1MsgReader
			e.syncer.Start()
		} else {
			go e.syncer.Start()
		}
	} else if e.isSequencer && !isSequencer {
		e.logger.Info("I am not a sequencer, stop syncing")
		e.syncer.Stop()
	}
	e.isSequencer = isSequencer
	return validatorUpdates, nil
}
