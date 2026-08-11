package derivation

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"slices"
	"strings"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/accounts/abi"
	"github.com/morph-l2/go-ethereum/common"
	eth "github.com/morph-l2/go-ethereum/core/types"
	"github.com/morph-l2/go-ethereum/crypto"
	geth "github.com/morph-l2/go-ethereum/eth"
)

var (
	ErrSnapshotChanged          = errors.New("fixed L1 snapshot changed")
	ErrCanonicalCommitNotFound  = errors.New("canonical CommitBatch log not found")
	ErrStoredBlobSourceNotFound = errors.New("stored blob source not found")
	ErrZeroBlobVersionedHash    = errors.New("pre-cutover zero blob versioned hash")
	zeroBlobVersionedHash       = common.HexToHash("0x010657f37554c781402a22917dee2f75def7ab966d7b770905398eba3c444014")
	canonicalResolverRollupABI  = mustCanonicalResolverRollupABI()
	canonicalCommitEventTopic   = canonicalResolverRollupABI.Events["CommitBatch"].ID
	canonicalRevertEventTopic   = canonicalResolverRollupABI.Events["RevertBatch"].ID
)

const canonicalResolverRollupABIJSON = `[
  {"type":"function","name":"committedBatches","stateMutability":"view","inputs":[{"name":"batchIndex","type":"uint256"}],"outputs":[{"name":"batchHash","type":"bytes32"}]},
  {"type":"function","name":"batchBlobVersionedHashes","stateMutability":"view","inputs":[{"name":"batchIndex","type":"uint256"}],"outputs":[{"name":"blobVersionedHash","type":"bytes32"}]},
  {"type":"event","name":"CommitBatch","anonymous":false,"inputs":[{"name":"batchIndex","type":"uint256","indexed":true},{"name":"batchHash","type":"bytes32","indexed":true}]},
  {"type":"event","name":"RevertBatch","anonymous":false,"inputs":[{"name":"batchIndex","type":"uint256","indexed":true},{"name":"batchHash","type":"bytes32","indexed":true}]}
]`

func mustCanonicalResolverRollupABI() abi.ABI {
	parsed, err := abi.JSON(strings.NewReader(canonicalResolverRollupABIJSON))
	if err != nil {
		panic(err)
	}
	return parsed
}

// FixedL1Snapshot pins both the number used by legacy JSON-RPC block-number
// calls and its hash. Every resolver verifies the hash before and after its
// reads, providing EIP-1898-equivalent fail-closed behavior on providers that
// do not support blockHash call options.
type FixedL1Snapshot struct {
	Number uint64
	Hash   common.Hash
}

type CanonicalCommitClient interface {
	HeaderByNumber(context.Context, *big.Int) (*eth.Header, error)
	CallContract(context.Context, ethereum.CallMsg, *big.Int) ([]byte, error)
	FilterLogs(context.Context, ethereum.FilterQuery) ([]eth.Log, error)
}

type StoredBlobSourceClient interface {
	CanonicalCommitClient
	TransactionByHash(context.Context, common.Hash) (*eth.Transaction, bool, error)
}

type CanonicalCommit struct {
	BatchIndex uint64
	BatchHash  common.Hash
	Log        eth.Log
}

type StoredBlobSource struct {
	BatchIndex uint64
	StoredHash common.Hash
	CommitLog  eth.Log
	Tx         *eth.Transaction
	BlobHashes []common.Hash
}

type RollupBatchDecoder func([]byte) (geth.RPCRollupBatch, error)

func PinL1Snapshot(ctx context.Context, client CanonicalCommitClient, number uint64) (FixedL1Snapshot, error) {
	header, err := client.HeaderByNumber(ctx, uint64Big(number))
	if err != nil {
		return FixedL1Snapshot{}, err
	}
	if header == nil || header.Number == nil || header.Number.Uint64() != number {
		return FixedL1Snapshot{}, fmt.Errorf("pin L1 snapshot %d: invalid header", number)
	}
	return FixedL1Snapshot{Number: number, Hash: header.Hash()}, nil
}

func ResolveCanonicalCommit(
	ctx context.Context,
	client CanonicalCommitClient,
	rollupAddress common.Address,
	batchIndex uint64,
	snapshotNumber uint64,
) (*CanonicalCommit, error) {
	snapshot, err := PinL1Snapshot(ctx, client, snapshotNumber)
	if err != nil {
		return nil, err
	}
	return ResolveCanonicalCommitAtSnapshot(ctx, client, rollupAddress, batchIndex, snapshot)
}

func ResolveCanonicalCommitAtSnapshot(
	ctx context.Context,
	client CanonicalCommitClient,
	rollupAddress common.Address,
	batchIndex uint64,
	snapshot FixedL1Snapshot,
) (*CanonicalCommit, error) {
	if err := verifyL1Snapshot(ctx, client, snapshot); err != nil {
		return nil, err
	}
	committedHash, err := readRollupHashAtSnapshot(ctx, client, rollupAddress, "committedBatches", batchIndex, snapshot.Number)
	if err != nil {
		return nil, err
	}
	if committedHash == (common.Hash{}) {
		if err := verifyL1Snapshot(ctx, client, snapshot); err != nil {
			return nil, err
		}
		return nil, nil
	}

	logs, err := client.FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: uint64Big(0),
		ToBlock:   uint64Big(snapshot.Number),
		Addresses: []common.Address{rollupAddress},
		Topics: [][]common.Hash{
			{canonicalCommitEventTopic, canonicalRevertEventTopic},
			{common.BigToHash(uint64Big(batchIndex))},
		},
	})
	if err != nil {
		return nil, err
	}
	ordered, err := validateAndSortLifecycleLogs(logs, rollupAddress, batchIndex, snapshot.Number)
	if err != nil {
		return nil, err
	}

	lastRevert := -1
	for i := range ordered {
		if ordered[i].Topics[0] == canonicalRevertEventTopic {
			lastRevert = i
		}
	}
	var selected *eth.Log
	for i := lastRevert + 1; i < len(ordered); i++ {
		log := ordered[i]
		if log.Topics[0] == canonicalCommitEventTopic && log.Topics[2] == committedHash {
			candidate := log
			selected = &candidate
		}
	}
	if err := verifyL1Snapshot(ctx, client, snapshot); err != nil {
		return nil, err
	}
	if selected == nil {
		return nil, fmt.Errorf("%w: batch %d, getter hash %s", ErrCanonicalCommitNotFound, batchIndex, committedHash.Hex())
	}
	return &CanonicalCommit{
		BatchIndex: batchIndex,
		BatchHash:  committedHash,
		Log:        *selected,
	}, nil
}

// ResolveStoredBlobSource deliberately does not apply the canonical resolver's
// "after the last Revert" filter. Rollup retains batchBlobVersionedHashes so a
// later blobless commitState/commitBatchWithProof can reuse a blob originally
// published by a pre-Revert Commit. The source remains identified by its full
// canonical L1 log and transaction identity.
func ResolveStoredBlobSource(
	ctx context.Context,
	client StoredBlobSourceClient,
	rollupAddress common.Address,
	batchIndex uint64,
	snapshotNumber uint64,
	decode RollupBatchDecoder,
) (*StoredBlobSource, error) {
	snapshot, err := PinL1Snapshot(ctx, client, snapshotNumber)
	if err != nil {
		return nil, err
	}
	return ResolveStoredBlobSourceAtSnapshot(ctx, client, rollupAddress, batchIndex, snapshot, decode)
}

func ResolveStoredBlobSourceAtSnapshot(
	ctx context.Context,
	client StoredBlobSourceClient,
	rollupAddress common.Address,
	batchIndex uint64,
	snapshot FixedL1Snapshot,
	decode RollupBatchDecoder,
) (*StoredBlobSource, error) {
	if decode == nil {
		return nil, errors.New("rollup batch decoder is nil")
	}
	if err := verifyL1Snapshot(ctx, client, snapshot); err != nil {
		return nil, err
	}
	storedHash, err := readRollupHashAtSnapshot(ctx, client, rollupAddress, "batchBlobVersionedHashes", batchIndex, snapshot.Number)
	if err != nil {
		return nil, err
	}
	if storedHash == (common.Hash{}) || storedHash == zeroBlobVersionedHash {
		if err := verifyL1Snapshot(ctx, client, snapshot); err != nil {
			return nil, err
		}
		if storedHash == zeroBlobVersionedHash {
			return nil, ErrZeroBlobVersionedHash
		}
		return nil, fmt.Errorf("%w: batch %d has no stored hash", ErrStoredBlobSourceNotFound, batchIndex)
	}

	logs, err := client.FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: uint64Big(0),
		ToBlock:   uint64Big(snapshot.Number),
		Addresses: []common.Address{rollupAddress},
		Topics: [][]common.Hash{
			{canonicalCommitEventTopic},
			{common.BigToHash(uint64Big(batchIndex))},
		},
	})
	if err != nil {
		return nil, err
	}
	ordered, err := validateAndSortLifecycleLogs(logs, rollupAddress, batchIndex, snapshot.Number)
	if err != nil {
		return nil, err
	}

	var source *StoredBlobSource
	for i := len(ordered) - 1; i >= 0; i-- {
		log := ordered[i]
		if log.Topics[0] != canonicalCommitEventTopic {
			continue
		}
		tx, pending, err := client.TransactionByHash(ctx, log.TxHash)
		if err != nil {
			return nil, err
		}
		if pending || tx == nil {
			return nil, fmt.Errorf("stored blob source transaction %s is pending or missing", log.TxHash.Hex())
		}
		if tx.Hash() != log.TxHash {
			return nil, fmt.Errorf("stored blob source transaction hash mismatch: requested %s, got %s", log.TxHash.Hex(), tx.Hash().Hex())
		}
		blobHashes := tx.BlobHashes()
		if len(blobHashes) == 0 {
			continue
		}
		batch, err := decode(tx.Data())
		if err != nil {
			return nil, fmt.Errorf("decode stored blob source transaction %s: %w", log.TxHash.Hex(), err)
		}
		if !storedBlobHashMatches(uint(batch.Version), blobHashes, storedHash) {
			continue
		}
		source = &StoredBlobSource{
			BatchIndex: batchIndex,
			StoredHash: storedHash,
			CommitLog:  log,
			Tx:         tx,
			BlobHashes: slices.Clone(blobHashes),
		}
		break
	}
	if err := verifyL1Snapshot(ctx, client, snapshot); err != nil {
		return nil, err
	}
	if source == nil {
		return nil, fmt.Errorf("%w: batch %d, stored hash %s", ErrStoredBlobSourceNotFound, batchIndex, storedHash.Hex())
	}
	return source, nil
}

func storedBlobHashMatches(version uint, blobHashes []common.Hash, storedHash common.Hash) bool {
	if len(blobHashes) == 0 {
		return false
	}
	if version != 2 {
		return len(blobHashes) == 1 && blobHashes[0] == storedHash
	}
	encoded := make([]byte, 0, len(blobHashes)*common.HashLength)
	for _, hash := range blobHashes {
		encoded = append(encoded, hash[:]...)
	}
	return crypto.Keccak256Hash(encoded) == storedHash
}

func readRollupHashAtSnapshot(
	ctx context.Context,
	client CanonicalCommitClient,
	rollupAddress common.Address,
	method string,
	batchIndex uint64,
	snapshotNumber uint64,
) (common.Hash, error) {
	input, err := canonicalResolverRollupABI.Pack(method, uint64Big(batchIndex))
	if err != nil {
		return common.Hash{}, err
	}
	output, err := client.CallContract(ctx, ethereum.CallMsg{To: &rollupAddress, Data: input}, uint64Big(snapshotNumber))
	if err != nil {
		return common.Hash{}, err
	}
	if len(output) != common.HashLength {
		return common.Hash{}, fmt.Errorf("%s returned %d bytes, want %d", method, len(output), common.HashLength)
	}
	values, err := canonicalResolverRollupABI.Methods[method].Outputs.Unpack(output)
	if err != nil {
		return common.Hash{}, err
	}
	if len(values) != 1 {
		return common.Hash{}, fmt.Errorf("%s returned %d values, want 1", method, len(values))
	}
	decoded := *abi.ConvertType(values[0], new([32]byte)).(*[32]byte)
	return common.Hash(decoded), nil
}

func validateAndSortLifecycleLogs(logs []eth.Log, rollupAddress common.Address, batchIndex, snapshotNumber uint64) ([]eth.Log, error) {
	indexTopic := common.BigToHash(uint64Big(batchIndex))
	ordered := slices.Clone(logs)
	for i := range ordered {
		log := ordered[i]
		if log.Removed {
			return nil, fmt.Errorf("removed lifecycle log at %d/%d/%d", log.BlockNumber, log.TxIndex, log.Index)
		}
		if log.Address != rollupAddress || log.BlockNumber > snapshotNumber || len(log.Topics) != 3 || log.Topics[1] != indexTopic {
			return nil, fmt.Errorf("malformed lifecycle log at %d/%d/%d", log.BlockNumber, log.TxIndex, log.Index)
		}
		if log.Topics[0] != canonicalCommitEventTopic && log.Topics[0] != canonicalRevertEventTopic {
			return nil, fmt.Errorf("unexpected lifecycle topic %s", log.Topics[0].Hex())
		}
	}
	slices.SortFunc(ordered, func(a, b eth.Log) int {
		if a.BlockNumber != b.BlockNumber {
			return compareUint64(a.BlockNumber, b.BlockNumber)
		}
		if a.TxIndex != b.TxIndex {
			return compareUint64(uint64(a.TxIndex), uint64(b.TxIndex))
		}
		return compareUint64(uint64(a.Index), uint64(b.Index))
	})
	return ordered, nil
}

func verifyL1Snapshot(ctx context.Context, client CanonicalCommitClient, snapshot FixedL1Snapshot) error {
	header, err := client.HeaderByNumber(ctx, uint64Big(snapshot.Number))
	if err != nil {
		return err
	}
	if header == nil || header.Number == nil || header.Number.Uint64() != snapshot.Number || header.Hash() != snapshot.Hash {
		return fmt.Errorf("%w: block %d expected %s", ErrSnapshotChanged, snapshot.Number, snapshot.Hash.Hex())
	}
	return nil
}

func uint64Big(value uint64) *big.Int {
	return new(big.Int).SetUint64(value)
}

func compareUint64(a, b uint64) int {
	switch {
	case a < b:
		return -1
	case a > b:
		return 1
	default:
		return 0
	}
}
