package node

import (
	"bytes"
	"slices"

	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/crypto/bls12381"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/upgrade"
)

const tmKeySize = ed25519.PubKeySize

// Keep the protected pre-cutover update path addressable for symbol-level
// audits even though all production entry points now use validatorSetAtHeight.
// The method bodies below intentionally remain unchanged apart from removing
// the old syncer-start side effect.
var _ = (*Executor).updateSequencerSet

func cloneValidatorKeys(keys [][]byte) [][]byte {
	cloned := make([][]byte, len(keys))
	for i := range keys {
		cloned[i] = bytes.Clone(keys[i])
	}
	return cloned
}

// validatorSetAtHeight is the single validator data-source boundary. The
// upgrade package intentionally switches strictly after the persisted height:
// H is historical, H+1 is static.
func (e *Executor) validatorSetAtHeight(height uint64) ([][]byte, error) {
	if upgrade.IsUpgraded(int64(height)) {
		return cloneValidatorKeys(e.staticValidatorTmKeys), nil
	}
	if e.legacyValidatorSetAtHeightFunc != nil {
		return e.legacyValidatorSetAtHeightFunc(int64(height))
	}
	return e.getValidatorsAtHeight(int64(height))
}

func (e *Executor) sequencerSetUpdates(height uint64) ([][]byte, error) {
	seqHash, err := e.sequencerCaller.SequencerSetVerifyHash(nil)
	if err != nil {
		return nil, err
	}
	if e.shouldReuseSequencerCache(height, seqHash) {
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
		if !e.shouldKeepSequencerAtHeight(height, stakesInfo[i].BlsKey) {
			e.logger.Error("sequencerSetUpdates: skip sequencer with invalid bls key", "height", height, "addr", stakesInfo[i].Addr)
			continue
		}
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

func (e *Executor) shouldReuseSequencerCache(height uint64, seqHash [32]byte) bool {
	if e.currentSeqHash == nil || !bytes.Equal(e.currentSeqHash[:], seqHash[:]) {
		return false
	}

	if e.blsKeyCheckForkHeight > 0 &&
		(height == e.blsKeyCheckForkHeight || height == e.blsKeyCheckForkHeight+1) {
		return false
	}
	return true
}

func (e *Executor) shouldKeepSequencerAtHeight(height uint64, blsKey []byte) bool {
	if isValidBlsKey(blsKey) {
		return true
	}

	return e.blsKeyCheckForkHeight > 0 && height <= e.blsKeyCheckForkHeight
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
		e.logger.Info("I am a sequencer")
	} else if e.isSequencer && !isSequencer {
		// as the derivation always need the syncer running, it should not be stopped

		//e.logger.Info("I am not a sequencer, stop syncing")
		//e.syncer.Stop()
	}
	e.isSequencer = isSequencer
	return validatorUpdates, nil
}

func isValidBlsKey(in []byte) bool {
	_, err := bls12381.NewG2().DecodePoint(in)
	return err == nil
}
