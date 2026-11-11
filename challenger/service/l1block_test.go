package service

import (
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/morph-l2/go-ethereum/core/types"
)

// fakeHeaders serves block timestamps from timeAt and counts header reads.
type fakeHeaders struct {
	timeAt func(n uint64) uint64
	reads  int
}

func (f *fakeHeaders) HeaderByNumber(_ context.Context, number *big.Int) (*types.Header, error) {
	f.reads++
	return &types.Header{Number: new(big.Int).Set(number), Time: f.timeAt(number.Uint64())}, nil
}

func (f *fakeHeaders) head(n uint64) *types.Header {
	return &types.Header{Number: new(big.Int).SetUint64(n), Time: f.timeAt(n)}
}

const genesisTime = 1_600_000_000

func steady(n uint64) uint64 { return genesisTime + 12*n }

func TestBlocksAtTimeSteadyChain(t *testing.T) {
	l1 := &fakeHeaders{timeAt: steady}
	head := l1.head(24_000_000)
	for _, want := range []uint64{1, 7, 12_345_678, 23_999_999, 24_000_000} {
		l1.reads = 0
		from, to, err := blocksAtTime(context.Background(), l1, steady(want), head)
		if err != nil {
			t.Fatalf("block %d: %v", want, err)
		}
		if from != want || to != want {
			t.Fatalf("block %d: got %d..%d", want, from, to)
		}
		if l1.reads > 5 {
			t.Fatalf("block %d: %d header reads", want, l1.reads)
		}
	}
}

func TestBlocksAtTimeMissedSlots(t *testing.T) {
	// Every seventh slot is empty, so block times drift from a fixed 12s step.
	timeAt := func(n uint64) uint64 { return genesisTime + 12*(n+n/6) }
	l1 := &fakeHeaders{timeAt: timeAt}
	head := l1.head(1_000_000)
	for _, want := range []uint64{0, 5, 6, 499_999, 1_000_000} {
		l1.reads = 0
		from, to, err := blocksAtTime(context.Background(), l1, timeAt(want), head)
		if err != nil {
			t.Fatalf("block %d: %v", want, err)
		}
		if from != want || to != want {
			t.Fatalf("block %d: got %d..%d", want, from, to)
		}
		if l1.reads > 60 {
			t.Fatalf("block %d: %d header reads", want, l1.reads)
		}
	}
}

func TestBlocksAtTimeSharedTimestamp(t *testing.T) {
	// Blocks 100..103 share one timestamp, as on a dev chain that seals faster than 1s.
	timeAt := func(n uint64) uint64 {
		switch {
		case n < 100:
			return genesisTime + n
		case n <= 103:
			return genesisTime + 100
		default:
			return genesisTime + n - 3
		}
	}
	l1 := &fakeHeaders{timeAt: timeAt}
	from, to, err := blocksAtTime(context.Background(), l1, genesisTime+100, l1.head(500))
	if err != nil {
		t.Fatal(err)
	}
	if from != 100 || to != 103 {
		t.Fatalf("got %d..%d, want 100..103", from, to)
	}
}

func TestBlocksAtTimeMissingTimestamp(t *testing.T) {
	l1 := &fakeHeaders{timeAt: steady}
	_, _, err := blocksAtTime(context.Background(), l1, steady(50)+5, l1.head(1000))
	if err == nil || !strings.Contains(err.Error(), "no l1 block has timestamp") {
		t.Fatalf("err = %v", err)
	}
}

func TestBlocksAtTimeAfterHead(t *testing.T) {
	l1 := &fakeHeaders{timeAt: steady}
	_, _, err := blocksAtTime(context.Background(), l1, steady(1001), l1.head(1000))
	if err == nil || !strings.Contains(err.Error(), "after head block") {
		t.Fatalf("err = %v", err)
	}
}

func TestBlocksAtTimeGenesis(t *testing.T) {
	l1 := &fakeHeaders{timeAt: steady}
	from, to, err := blocksAtTime(context.Background(), l1, genesisTime, l1.head(1000))
	if err != nil {
		t.Fatal(err)
	}
	if from != 0 || to != 0 {
		t.Fatalf("got %d..%d", from, to)
	}
}
