package service

import "fmt"

// RejectError is a challenge the contract would refuse. The HTTP API returns 409.
type RejectError struct {
	Reason string
}

func (e *RejectError) Error() string {
	return e.Reason
}

// NotCommittedError means committedBatches holds no hash for this index. The batch
// was never committed, or it was reverted and has not been committed again.
type NotCommittedError struct {
	BatchIndex uint64
}

func (e *NotCommittedError) Error() string {
	return "batch is not committed"
}

// CallError is an RPC or decoding failure while reading chain state.
type CallError struct {
	Err error
}

func (e *CallError) Error() string {
	return e.Err.Error()
}

func (e *CallError) Unwrap() error {
	return e.Err
}

// BusyError means this signer already has a transaction in flight.
type BusyError struct {
	Confirmed uint64
	Pending   uint64
}

func (e *BusyError) Error() string {
	return fmt.Sprintf("pending transactions exist (confirmed: %d, pending: %d)", e.Confirmed, e.Pending)
}

// TxError is returned after a challenge transaction has been broadcast.
type TxError struct {
	TxHash   string
	Pending  bool
	Reverted bool
	Err      error
}

func (e *TxError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return "challenge transaction failed"
}

func (e *TxError) Unwrap() error {
	return e.Err
}
