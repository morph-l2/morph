package node

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/morph-l2/bindings/bindings"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
	"github.com/scroll-tech/go-ethereum/crypto/bls12381"
	"github.com/tendermint/tendermint/blssignatures"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

const tmKeySize = ed25519.PubKeySize

type sequencerKey struct {
	index     uint64
	blsPubKey blssignatures.PublicKey
}

func (e *Executor) getBlsPubKeyByTmKey(tmPubKey []byte, height *uint64) *sequencerKey {
	var pk [tmKeySize]byte
	copy(pk[:], tmPubKey)

	if e.currentSequencerSet != nil && (height == nil || *height >= e.currentSequencerSet.startHeight) {
		seqKey, ok := e.currentSequencerSet.sequencerSet[pk]
		if ok {
			return &seqKey
		}
	}
	if e.currentSequencerSet.startHeight <= 1 { // means no previous sequencer set
		return nil
	}

	nextEndHeight := e.currentSequencerSet.startHeight - 1
	if len(e.previousSequencerSet) > 0 {
		for i := len(e.previousSequencerSet) - 1; i >= 0; i-- {
			if height == nil || (*height >= e.previousSequencerSet[i].startHeight && *height <= nextEndHeight) {
				seqKey, ok := e.previousSequencerSet[i].sequencerSet[pk]
				if ok {
					return &seqKey
				}
				nextEndHeight = e.previousSequencerSet[i].startHeight - 1
			}
		}
	}
	return nil
}

func (e *Executor) VerifySignature(tmPubKey []byte, messageHash []byte, blsSig []byte) (bool, error) {
	if e.devSequencer {
		e.logger.Info("we are in dev mode, do not verify the bls signature")
		return true, nil
	}
	if e.currentSequencerSet == nil {
		return false, errors.New("no available sequencers found in layer2")
	}

	seqKey := e.getBlsPubKeyByTmKey(tmPubKey, nil)
	if seqKey == nil {
		return false, errors.New("it is not a valid sequencer")
	}

	sig, err := blssignatures.SignatureFromBytes(blsSig)
	if err != nil {
		e.logger.Error("failed to recover bytes to signature", "error", err)
		return false, fmt.Errorf("failed to recover bytes to signature, error: %v", err)
	}
	return blssignatures.VerifySignature(sig, messageHash, seqKey.blsPubKey)
}

func (e *Executor) sequencerSetUpdates(curHeight *uint64) ([][]byte, error) {
	currentVersion, err := e.sequencerContract.CurrentVersion(nil)
	if err != nil {
		return nil, err
	}
	if e.currentSequencerSet != nil && e.currentSequencerSet.version == currentVersion.Uint64() {
		return e.nextValidators, nil
	}

	// found new version sequencerSet
	// move current sequencer set to previous sequencer set
	if e.currentSequencerSet != nil && curHeight != nil {
		e.previousSequencerSet = append(e.previousSequencerSet, *e.currentSequencerSet)
		if len(e.previousSequencerSet) > 2 {
			e.previousSequencerSet = e.previousSequencerSet[len(e.previousSequencerSet)-2:] // only reserves 2 elements
		}
	} else if currentVersion.Uint64() != 0 { // first time to fetch sequencer set, and it is not the first version
		preSequencerInfo, err := e.sequencerContract.GetSequencerInfos(nil, true)
		if err != nil {
			e.logger.Error("failed to call GetSequencerInfos", "previous", true, "err", err)
			return nil, err
		}
		_, preSequencerSet, err := e.convertSequencerSet(preSequencerInfo)
		if err != nil {
			return nil, err
		}
		preVersionHeight, err := e.sequencerContract.PreVersionHeight(nil)
		if err != nil {
			e.logger.Error("failed to call PreVersionHeight", "err", err)
			return nil, err
		}
		var preStartHeight uint64
		if preVersionHeight.Sign() > 0 {
			preStartHeight = preVersionHeight.Uint64() + 2
		}
		e.previousSequencerSet = []SequencerSetInfo{{
			version:      currentVersion.Uint64() - 1,
			startHeight:  preStartHeight,
			sequencerSet: preSequencerSet,
		}}
	}

	// fetch current sequencerSet info
	sequencersInfo, err := e.sequencerContract.GetSequencerInfos(nil, false)
	if err != nil {
		e.logger.Error("failed to call GetSequencerInfos", "previous", false, "err", err)
		return nil, err
	}
	newValidators, newSequencerSet, err := e.convertSequencerSet(sequencersInfo)
	if err != nil {
		return nil, err
	}
	var currentStartHeight uint64
	curVersionHeight, err := e.sequencerContract.CurrentVersionHeight(nil)
	if err != nil {
		e.logger.Error("failed to call CurrentVersionHeight", "err", err)
		return nil, err
	}
	if curVersionHeight.Uint64() != 0 {
		currentStartHeight = curVersionHeight.Uint64() + 2
	}
	e.currentSequencerSet = &SequencerSetInfo{
		version:      currentVersion.Uint64(),
		startHeight:  currentStartHeight,
		sequencerSet: newSequencerSet,
	}
	e.nextValidators = newValidators

	var before string
	if len(e.previousSequencerSet) > 0 {
		before = e.previousSequencerSet[len(e.previousSequencerSet)-1].String()
	}
	e.logger.Info("sequencers updates, version changed", "before", before, "current", e.currentSequencerSet.String())
	return newValidators, nil
}

func (e *Executor) convertSequencerSet(sequencersInfo []bindings.TypesSequencerInfo) ([][]byte, map[[32]byte]sequencerKey, error) {
	newValidators := make([][]byte, 0)
	newSequencerSet := make(map[[tmKeySize]byte]sequencerKey)
	for i := range sequencersInfo {
		blsPK, err := decodeBlsPubKey(sequencersInfo[i].BlsKey)
		if err != nil {
			e.logger.Error("failed to decode bls key", "key bytes", hexutil.Encode(sequencersInfo[i].BlsKey), "error", err)
			return nil, nil, err
		}
		newSequencerSet[sequencersInfo[i].TmKey] = sequencerKey{
			index:     uint64(i),
			blsPubKey: blsPK,
		}
		newValidators = append(newValidators, sequencersInfo[i].TmKey[:])
	}
	return newValidators, newSequencerSet, nil
}

func (e *Executor) batchParamsUpdates(height uint64) (*tmproto.BatchParams, error) {
	var (
		batchBlockInterval, batchMaxBytes, batchTimeout, batchMaxChunks *big.Int
		err                                                             error
	)

	if batchBlockInterval, err = e.govContract.BatchBlockInterval(nil); err != nil {
		return nil, err
	}
	if batchMaxBytes, err = e.govContract.BatchMaxBytes(nil); err != nil {
		return nil, err
	}
	if batchTimeout, err = e.govContract.BatchTimeout(nil); err != nil {
		return nil, err
	}
	if batchMaxChunks, err = e.govContract.MaxChunks(nil); err != nil {
		return nil, err
	}

	changed := e.batchParams.BlocksInterval != batchBlockInterval.Int64() ||
		e.batchParams.MaxBytes != batchMaxBytes.Int64() ||
		int64(e.batchParams.Timeout.Seconds()) != batchTimeout.Int64() ||
		e.batchParams.MaxChunks != batchMaxChunks.Int64()

	if changed {
		e.batchParams.BlocksInterval = batchBlockInterval.Int64()
		e.batchParams.MaxBytes = batchMaxBytes.Int64()
		e.batchParams.Timeout = time.Duration(batchTimeout.Int64() * int64(time.Second))
		e.batchParams.MaxChunks = batchMaxChunks.Int64()
		e.logger.Info("batch params changed", "height", height,
			"batchBlockInterval", batchBlockInterval.Int64(),
			"batchMaxBytes", batchMaxBytes.Int64(),
			"batchTimeout", batchTimeout.Int64(),
			"batchMaxChunks", batchMaxChunks.Int64())
		return &tmproto.BatchParams{
			BlocksInterval: batchBlockInterval.Int64(),
			MaxBytes:       batchMaxBytes.Int64(),
			Timeout:        time.Duration(batchTimeout.Int64() * int64(time.Second)),
			MaxChunks:      batchMaxChunks.Int64(),
		}, nil
	}
	return nil, nil
}

func (e *Executor) updateSequencerSet(curHeight *uint64) ([][]byte, error) {
	validatorUpdates, err := e.sequencerSetUpdates(curHeight)
	if err != nil {
		e.logger.Error("failed to get sequencer set from geth", "err", err)
		return nil, err
	}
	var tmPKBz [tmKeySize]byte
	copy(tmPKBz[:], e.tmPubKey)
	_, isSequencer := e.currentSequencerSet.sequencerSet[tmPKBz]
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
