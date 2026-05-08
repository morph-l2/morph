package derivation

import (
	"context"
	"errors"
	"fmt"

	"github.com/morph-l2/go-ethereum/common"
)

// SPEC-005 §3.6 / §5.1 admin RPC: operator-triggered rollback entry point.
//
// Exposes the ability to roll the local L2 chain back to a target (number,
// hash) pair. Per tech-design §3.3, the rollback **must** match by hash —
// rolling back to a number alone is unsafe because it can silently land
// on a different fork after a reorg.
//
// Authentication and the concrete wire-up (registering this with the
// node's existing admin RPC server) are blocked on SPEC-005 §8 #2:
//   - dev-mode only (current default below)
//   - operator-only via a node-local UNIX socket
//   - signed multisig request
// All three options keep the same public method signature.

// AdminAPI groups operator-only RPC entry points exposed by the
// derivation pipeline.
//
// TODO(spec-005-admin-rpc): wire this into morph/node/cmd/node/main.go
// once SPEC-005 §8 #2 (auth) is decided. Until then, AdminAPI is
// constructible but unregistered; tests can still exercise it directly.
type AdminAPI struct {
	d *Derivation
}

// NewAdminAPI returns the operator-only API surface bound to the given
// Derivation instance.
func NewAdminAPI(d *Derivation) *AdminAPI {
	return &AdminAPI{d: d}
}

// SetL2Head requests a rollback of the local L2 chain to the supplied
// (number, hash). The implementation must verify that hash matches the
// local block at the given number before delegating to the rollback
// executor (SPEC-005 §5.1 / §5.2).
//
// Returns an error if:
//   - the (number, hash) does not match the local canonical chain;
//   - the target is below finalized_head (SPEC-005 §3.6: halted);
//   - the rollback executor itself fails (the node enters halted).
func (a *AdminAPI) SetL2Head(ctx context.Context, number uint64, hash common.Hash) error {
	if a == nil || a.d == nil {
		return errors.New("admin API not bound to a derivation instance")
	}

	if err := a.d.checkRollbackBoundary(number); err != nil {
		return err
	}

	// TODO(spec-005-admin-rpc):
	//   1. Authenticate the request (SPEC-005 §8 #2).
	//   2. Verify hash matches local block at `number` via l2Client.
	//   3. Acquire sequencerMutex.AcquireRollback() / defer release.
	//   4. Call into rollbackLocalChain(number) — currently returns
	//      "not implemented" because the underlying go-ethereum
	//      hash-matched SetHead interface (SPEC-005 §8 #4) is not finalised.
	return fmt.Errorf("admin SetL2Head not yet implemented (number=%d, hash=%s)", number, hash.Hex())
}
