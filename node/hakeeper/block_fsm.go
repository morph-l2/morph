package hakeeper

import (
	"encoding/binary"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/hashicorp/raft"
	tmlog "github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/types"
)

// FSMDecodeError is returned when a Raft log entry cannot be decoded into a BlockV2.
// This typically indicates a programming bug or proto incompatibility.
type FSMDecodeError struct{ Err error }

func (e *FSMDecodeError) Error() string  { return fmt.Sprintf("FSM decode: %v", e.Err) }
func (e *FSMDecodeError) Unwrap() error  { return e.Err }

// FSMApplyError is returned when the business callback (geth applyBlock / saveSignature) fails.
type FSMApplyError struct {
	Height uint64
	Err    error
}

func (e *FSMApplyError) Error() string  { return fmt.Sprintf("FSM apply height %d: %v", e.Height, e.Err) }
func (e *FSMApplyError) Unwrap() error  { return e.Err }

var _ raft.FSM = (*BlockFSM)(nil)

// BlockFSM implements raft.FSM for the Sequencer HA V2 module.
// It replaces the old RaftStateTracker: instead of storing full consensus payloads,
// it stores only the applied block height (for log compaction) and delivers decoded
// blocks to subscribers via a buffered channel.
type BlockFSM struct {
	logger tmlog.Logger
	mu     sync.RWMutex

	// appliedHeight is the block number of the most recently applied log entry.
	// Used exclusively by Snapshot for log compaction; NOT a full block reference.
	appliedHeight uint64

	// blockCh delivers applied blocks to Subscribe() consumers (broadcastRoutine).
	// Buffer of 200 gives ample room for transient subscriber slowness.
	blockCh chan *types.BlockV2

	// onApplied is the injected business callback. Protected by mu for safe concurrent set/read.
	onApplied func(*types.BlockV2) error
}

// NewBlockFSM creates a new BlockFSM.
func NewBlockFSM(logger tmlog.Logger) *BlockFSM {
	return &BlockFSM{
		logger:  logger,
		blockCh: make(chan *types.BlockV2, 200),
	}
}

// SetOnBlockApplied sets the business callback invoked on every FSM.Apply.
// Must be called before Start (i.e. before any Raft logs are applied).
func (f *BlockFSM) SetOnBlockApplied(fn func(*types.BlockV2) error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.onApplied = fn
}

// Apply implements raft.FSM.
// Called by the Raft library on the FSM goroutine after a log entry is committed.
// For the leader, raft.Apply blocks until this method returns (the Future completes).
// For followers, this runs asynchronously.
//
// Error handling:
//   - Decode failure → returns FSMDecodeError. For the leader this propagates via
//     Future.Response() and triggers a panic (invariant violation). For followers
//     it is logged by Raft.
//   - onApplied failure → returns FSMApplyError. For the leader this triggers a
//     panic via Commit(). For followers, the block is NOT delivered to blockCh
//     and appliedHeight is NOT advanced; the follower becomes degraded and
//     requires manual resync.
//   - Success → block is delivered to blockCh (for P2P broadcast) and
//     appliedHeight is advanced (for snapshot/log compaction).
func (f *BlockFSM) Apply(l *raft.Log) interface{} {
	// Skip non-command logs (configuration changes, barriers, etc.)
	if l.Type != raft.LogCommand {
		return nil
	}

	t0 := time.Now()

	block, err := decodeBlock(l.Data)
	if err != nil {
		return &FSMDecodeError{Err: err}
	}
	decodeDur := time.Since(t0)

	f.mu.RLock()
	fn := f.onApplied
	f.mu.RUnlock()

	var onAppliedDur time.Duration
	if fn != nil {
		t1 := time.Now()
		if err := fn(block); err != nil {
			return &FSMApplyError{Height: block.Number, Err: err}
		}
		onAppliedDur = time.Since(t1)
	}

	totalDur := time.Since(t0)

	f.logger.Debug("[PERF] BlockFSM.Apply",
		"height", block.Number,
		"decode_ms", float64(decodeDur.Microseconds())/1000.0,
		"onApplied_ms", float64(onAppliedDur.Microseconds())/1000.0,
		"total_ms", float64(totalDur.Microseconds())/1000.0,
		"txCount", len(block.Transactions),
		"dataBytes", len(l.Data),
	)

	select {
	case f.blockCh <- block:
	default:
		f.logger.Error("BlockFSM: blockCh full, subscriber too slow", "height", block.Number)
	}

	f.mu.Lock()
	f.appliedHeight = block.Number
	f.mu.Unlock()

	return nil
}

// Snapshot implements raft.FSM.
// Returns a snapshot containing only appliedHeight as an 8-byte big-endian uint64.
// This is for log compaction only -- it does NOT store full block data.
// If a follower falls behind beyond TrailingLogs and receives InstallSnapshot,
// it must be manually resynchronized (Fullnode sync + rejoin).
func (f *BlockFSM) Snapshot() (raft.FSMSnapshot, error) {
	f.mu.RLock()
	h := f.appliedHeight
	f.mu.RUnlock()
	return &blockSnapshot{height: h}, nil
}

// Restore implements raft.FSM.
// Reads the 8-byte appliedHeight from the snapshot. Does NOT call onApplied --
// geth state must be recovered independently (Fullnode P2P sync).
func (f *BlockFSM) Restore(rc io.ReadCloser) error {
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return fmt.Errorf("BlockFSM.Restore: read failed: %w", err)
	}
	if len(data) == 0 {
		return nil
	}
	if len(data) != 8 {
		return fmt.Errorf("BlockFSM.Restore: unexpected snapshot size %d, expected 8", len(data))
	}

	height := binary.BigEndian.Uint64(data)

	f.mu.Lock()
	f.appliedHeight = height
	f.mu.Unlock()

	f.logger.Info("BlockFSM.Restore: restored appliedHeight from snapshot", "height", height)
	return nil
}

// --- blockSnapshot ---

var _ raft.FSMSnapshot = (*blockSnapshot)(nil)

// blockSnapshot persists a single uint64 (appliedHeight) for log compaction.
type blockSnapshot struct {
	height uint64
}

// Persist implements raft.FSMSnapshot.
// Writes appliedHeight as 8-byte big-endian to the snapshot sink.
func (s *blockSnapshot) Persist(sink raft.SnapshotSink) error {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], s.height)
	if _, err := sink.Write(buf[:]); err != nil {
		sink.Cancel()
		return fmt.Errorf("blockSnapshot.Persist: write failed: %w", err)
	}
	return sink.Close()
}

// Release implements raft.FSMSnapshot. No-op.
func (s *blockSnapshot) Release() {}
