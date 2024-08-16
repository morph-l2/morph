package node

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
	"github.com/morph-l2/go-ethereum/crypto/bls12381"
	"github.com/tendermint/tendermint/blssignatures"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	"golang.org/x/exp/slices"
)

const tmKeySize = ed25519.PubKeySize

type validatorInfo struct {
	address   common.Address
	blsPubKey blssignatures.PublicKey
}

func (e *Executor) getBlsPubKeyByTmKey(tmPubKey []byte) (blssignatures.PublicKey, bool) {
	var tmKey [32]byte
	copy(tmKey[:], tmPubKey)
	val, found := e.valsByTmKey[tmKey]
	if !found {
		return blssignatures.PublicKey{}, false
	}
	return val.blsPubKey, true
}

func (e *Executor) VerifySignature(tmPubKey []byte, messageHash []byte, blsSig []byte) (bool, error) {
	if e.devSequencer {
		e.logger.Info("we are in dev mode, do not verify the bls signature")
		return true, nil
	}
	if len(e.valsByTmKey) == 0 {
		return false, errors.New("no available sequencers found in layer2")
	}

	blsKey, found := e.getBlsPubKeyByTmKey(tmPubKey)
	if !found {
		return false, errors.New("it is not a valid sequencer")
	}

	sig, err := blssignatures.SignatureFromBytes(blsSig)
	if err != nil {
		e.logger.Error("failed to recover bytes to signature", "error", err)
		return false, fmt.Errorf("failed to recover bytes to signature, error: %v", err)
	}
	return blssignatures.VerifySignature(sig, messageHash, blsKey)
}

func (e *Executor) sequencerSetUpdates() ([][]byte, error) {
	seqHash, err := e.sequencer.SequencerSetVerifyHash(nil)
	if err != nil {
		return nil, err
	}
	if e.currentSeqHash != nil && bytes.Equal(e.currentSeqHash[:], seqHash[:]) {
		return e.nextValidators, nil
	}

	sequencerSet0, err := e.sequencer.GetSequencerSet0(nil)
	if err != nil {
		return nil, err
	}
	sequencerSet1, err := e.sequencer.GetSequencerSet1(nil)
	if err != nil {
		return nil, err
	}
	sequencerSet2, err := e.sequencer.GetSequencerSet2(nil)
	if err != nil {
		return nil, err
	}

	cache := make(map[common.Address]struct{})
	requestAddrs := make([]common.Address, 0)
	for _, addr := range append(append(sequencerSet0, sequencerSet1...), sequencerSet2...) {
		_, ok := cache[addr]
		if !ok {
			cache[addr] = struct{}{}
			requestAddrs = append(requestAddrs, addr)
		}
	}
	stakesInfo, err := e.l2Staking.GetStakesInfo(nil, requestAddrs)
	if err != nil {
		e.logger.Error("failed to GetStakesInfo", "error", err)
		return nil, err
	}

	valsByTmKey := make(map[[tmKeySize]byte]validatorInfo)
	nextValidators := make([][]byte, 0)
	for i := range stakesInfo {
		blsPK, err := decodeBlsPubKey(stakesInfo[i].BlsKey)
		if err != nil {
			e.logger.Error("failed to decode bls key", "key bytes", hexutil.Encode(stakesInfo[i].BlsKey), "error", err)
			return nil, err
		}
		// sequencerSet2 is the latest updated sequencer set which is considered as the next validator set for tendermint
		if slices.Contains(sequencerSet2, stakesInfo[i].Addr) {
			nextValidators = append(nextValidators, stakesInfo[i].TmKey[:])
		}
		valsByTmKey[stakesInfo[i].TmKey] = validatorInfo{
			address:   stakesInfo[i].Addr,
			blsPubKey: blsPK,
		}
	}

	e.logger.Info("sequencers updates, sequencer verified hash changed")
	e.valsByTmKey = valsByTmKey
	e.nextValidators = nextValidators
	e.currentSeqHash = &seqHash
	return nextValidators, nil
}

func (e *Executor) batchParamsUpdates(height uint64) (*tmproto.BatchParams, error) {
	var (
		batchBlockInterval, batchTimeout, batchMaxChunks *big.Int
		err                                              error
	)

	if batchBlockInterval, err = e.govContract.BatchBlockInterval(nil); err != nil {
		return nil, err
	}
	if batchTimeout, err = e.govContract.BatchTimeout(nil); err != nil {
		return nil, err
	}
	if batchMaxChunks, err = e.govContract.MaxChunks(nil); err != nil {
		return nil, err
	}

	changed := e.batchParams.BlocksInterval != batchBlockInterval.Int64() ||
		int64(e.batchParams.Timeout.Seconds()) != batchTimeout.Int64() ||
		e.batchParams.MaxChunks != batchMaxChunks.Int64()

	if changed {
		e.batchParams.BlocksInterval = batchBlockInterval.Int64()
		e.batchParams.Timeout = time.Duration(batchTimeout.Int64() * int64(time.Second))
		e.batchParams.MaxChunks = batchMaxChunks.Int64()
		e.logger.Info("batch params changed", "height", height,
			"batchBlockInterval", batchBlockInterval.Int64(),
			"batchTimeout", batchTimeout.Int64(),
			"batchMaxChunks", batchMaxChunks.Int64())
		return &tmproto.BatchParams{
			BlocksInterval: batchBlockInterval.Int64(),
			Timeout:        time.Duration(batchTimeout.Int64() * int64(time.Second)),
			MaxChunks:      batchMaxChunks.Int64(),
		}, nil
	}
	return nil, nil
}

func (e *Executor) updateSequencerSet() ([][]byte, error) {
	validatorUpdates, err := e.sequencerSetUpdates()
	if err != nil {
		e.logger.Error("failed to get sequencer set from geth", "err", err)
		return nil, err
	}
	var tmPKBz [tmKeySize]byte
	copy(tmPKBz[:], e.tmPubKey)

	_, isSequencer := e.valsByTmKey[tmPKBz]
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

func decodeBlsPubKey(in []byte) (blssignatures.PublicKey, error) {
	g2P, err := bls12381.NewG2().DecodePoint(in)
	if err != nil {
		return blssignatures.PublicKey{}, err
	}
	return blssignatures.NewTrustedPublicKey(g2P), nil
}
