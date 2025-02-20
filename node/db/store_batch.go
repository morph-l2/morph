package db

import (
	"errors"

	"github.com/morph-l2/go-ethereum/core/rawdb"
	"github.com/morph-l2/go-ethereum/core/types"
	eth "github.com/morph-l2/go-ethereum/eth"
)

func (s *Store) ImportBatch(batch *types.RollupBatch, signatures []*types.BatchSignature) error {

	dbBatch := s.db.NewBatch()
	rawdb.WriteRollupBatch(dbBatch, batch)
	for _, signature := range signatures {
		rawdb.WriteBatchSignature(dbBatch, batch.Hash, *signature)
	}
	return dbBatch.Write()
}

func (s *Store) GetBatchByIndex(index uint64) (*types.RollupBatch, []*types.BatchSignature, error) {
	rollupBatch, err := rawdb.ReadRollupBatch(s.db, index)
	if err != nil {
		return nil, nil, errors.New("failed to read batch")
	}
	if rollupBatch == nil {
		return nil, nil, nil
	}
	signatures, err := rawdb.ReadBatchSignatures(s.db, rollupBatch.Hash)
	if err != nil {
		return nil, nil, errors.New("failed to read signatures")
	}
	return rollupBatch, signatures, nil
}

func (s *Store) GetRollupBatchByIndex(index uint64) (*eth.RPCRollupBatch, error) {
	rollupBatch, err := rawdb.ReadRollupBatch(s.db, index)
	if err != nil {
		return nil, errors.New("failed to read batch")
	}
	if rollupBatch == nil {
		return nil, nil
	}
	signatures, err := rawdb.ReadBatchSignatures(s.db, rollupBatch.Hash)
	if err != nil {
		return nil, errors.New("failed to read signatures")
	}

	rpcSignatures := make([]eth.RPCBatchSignature, len(signatures))
	for i, sig := range signatures {
		rpcSignatures[i] = eth.RPCBatchSignature{
			Signer:       sig.Signer,
			SignerPubKey: sig.SignerPubKey,
			Signature:    sig.Signature,
		}
	}

	var sidecar types.BlobTxSidecar
	if rollupBatch.Sidecar != nil {
		sidecar = *rollupBatch.Sidecar
	}

	// var collectedL1Fee *hexutil.Big
	// l1DataFee := rawdb.ReadBatchL1DataFee(s.db, index)
	// if l1DataFee != nil {
	// 	collectedL1Fee = (*hexutil.Big)(l1DataFee)
	// }

	return &eth.RPCRollupBatch{
		Version:                  rollupBatch.Version,
		Hash:                     rollupBatch.Hash,
		ParentBatchHeader:        rollupBatch.ParentBatchHeader,
		BlockContexts:            rollupBatch.BlockContexts,
		CurrentSequencerSetBytes: rollupBatch.CurrentSequencerSetBytes,
		PrevStateRoot:            rollupBatch.PrevStateRoot,
		PostStateRoot:            rollupBatch.PostStateRoot,
		WithdrawRoot:             rollupBatch.WithdrawRoot,
		Sidecar:                  sidecar,
		Signatures:               rpcSignatures,
		// CollectedL1Fee:           collectedL1Fee,
	}, nil
}
