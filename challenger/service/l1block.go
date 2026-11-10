package service

import (
	"context"
	"fmt"
	"math/big"

	"github.com/morph-l2/go-ethereum/core/types"
)

type headerReader interface {
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
}

// blocksAtTime returns the first and last block up to head whose timestamp is ts.
// Block timestamps never decrease, so the matching blocks form one contiguous run.
// The search alternates interpolation and bisection: interpolation lands within a
// few blocks on a chain with a steady block time, and bisection bounds the worst case.
func blocksAtTime(ctx context.Context, l1 headerReader, ts uint64, head *types.Header) (uint64, uint64, error) {
	if head.Time < ts {
		return 0, 0, fmt.Errorf("timestamp %d is after head block %d", ts, head.Number.Uint64())
	}
	timeAt := func(n uint64) (uint64, error) {
		header, err := l1.HeaderByNumber(ctx, new(big.Int).SetUint64(n))
		if err != nil {
			return 0, fmt.Errorf("l1 header %d: %w", n, err)
		}
		return header.Time, nil
	}

	lo, hi := uint64(0), head.Number.Uint64()
	tlo, err := timeAt(lo)
	if err != nil {
		return 0, 0, err
	}
	thi := head.Time
	first, firstTime := lo, tlo
	if tlo < ts {
		// Invariant: time(lo) < ts <= time(hi).
		for step := 0; hi-lo > 1; step++ {
			mid := lo + (hi-lo)/2
			if step%2 == 0 {
				mid = lo + (ts-tlo)*(hi-lo)/(thi-tlo)
				if mid <= lo {
					mid = lo + 1
				}
				if mid >= hi {
					mid = hi - 1
				}
			}
			t, err := timeAt(mid)
			if err != nil {
				return 0, 0, err
			}
			if t < ts {
				lo, tlo = mid, t
			} else {
				hi, thi = mid, t
			}
		}
		first, firstTime = hi, thi
	}
	if firstTime != ts {
		return 0, 0, fmt.Errorf("no l1 block has timestamp %d", ts)
	}

	last := first
	for last < head.Number.Uint64() {
		t, err := timeAt(last + 1)
		if err != nil {
			return 0, 0, err
		}
		if t != ts {
			break
		}
		last++
	}
	return first, last, nil
}
