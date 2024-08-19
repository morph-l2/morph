package derivation

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/morph-l2/go-ethereum/common/hexutil"
)

type ErrorCode int

// InputError distinguishes an user-input error from regular rpc errors,
// to help the (Engine) API user divert from accidental input mistakes.
type InputError struct {
	Inner error
	Code  ErrorCode
}

func (ie InputError) Error() string {
	return fmt.Sprintf("input error %d: %s", ie.Code, ie.Inner.Error())
}

func (ie InputError) Unwrap() error {
	return ie.Inner
}

// Is checks if the error is the given target type.
// Any type of InputError counts, regardless of code.
func (ie InputError) Is(target error) bool {
	_, ok := target.(InputError)
	return ok || errors.Is(target, ie.Unwrap())
}

type Bytes32 [32]byte

func (b *Bytes32) UnmarshalJSON(text []byte) error {
	return hexutil.UnmarshalFixedJSON(reflect.TypeOf(b), text, b[:])
}

func (b *Bytes32) UnmarshalText(text []byte) error {
	return hexutil.UnmarshalFixedText("Bytes32", text, b[:])
}

func (b *Bytes32) MarshalText() ([]byte, error) {
	return hexutil.Bytes(b[:]).MarshalText()
}

func (b *Bytes32) String() string {
	return hexutil.Encode(b[:])
}

// TerminalString implements log.TerminalStringer, formatting a string for console
// output during logging.
func (b *Bytes32) TerminalString() string {
	return fmt.Sprintf("%x..%x", b[:3], b[29:])
}

type Bytes256 [256]byte

func (b *Bytes256) UnmarshalJSON(text []byte) error {
	return hexutil.UnmarshalFixedJSON(reflect.TypeOf(b), text, b[:])
}

func (b *Bytes256) UnmarshalText(text []byte) error {
	return hexutil.UnmarshalFixedText("Bytes32", text, b[:])
}

func (b *Bytes256) MarshalText() ([]byte, error) {
	return hexutil.Bytes(b[:]).MarshalText()
}

func (b *Bytes256) String() string {
	return hexutil.Encode(b[:])
}

// TerminalString implements log.TerminalStringer, formatting a string for console
// output during logging.
func (b *Bytes256) TerminalString() string {
	return fmt.Sprintf("%x..%x", b[:3], b[253:])
}

type BytesMax32 []byte

func (b *BytesMax32) UnmarshalJSON(text []byte) error {
	if len(text) > 64+2+2 { // account for delimiter "", and 0x prefix
		return fmt.Errorf("input too long, expected at most 32 hex-encoded, 0x-prefixed, bytes: %x", text)
	}
	return (*hexutil.Bytes)(b).UnmarshalJSON(text)
}

func (b *BytesMax32) UnmarshalText(text []byte) error {
	if len(text) > 64+2 { // account for 0x prefix
		return fmt.Errorf("input too long, expected at most 32 hex-encoded, 0x-prefixed, bytes: %x", text)
	}
	return (*hexutil.Bytes)(b).UnmarshalText(text)
}

func (b BytesMax32) MarshalText() ([]byte, error) {
	return (hexutil.Bytes)(b).MarshalText()
}

func (b BytesMax32) String() string {
	return hexutil.Encode(b)
}

type Bytes48 [48]byte

func (b *Bytes48) UnmarshalJSON(text []byte) error {
	return hexutil.UnmarshalFixedJSON(reflect.TypeOf(b), text, b[:])
}

func (b *Bytes48) UnmarshalText(text []byte) error {
	return hexutil.UnmarshalFixedText("Bytes32", text, b[:])
}

func (b *Bytes48) MarshalText() ([]byte, error) {
	return hexutil.Bytes(b[:]).MarshalText()
}

func (b *Bytes48) String() string {
	return hexutil.Encode(b[:])
}

// TerminalString implements log.TerminalStringer, formatting a string for console
// output during logging.
func (b *Bytes48) TerminalString() string {
	return fmt.Sprintf("%x..%x", b[:3], b[45:])
}

// Uint64String is a decimal string representation of an uint64, for usage in the Beacon API JSON encoding
type Uint64String uint64

func (v Uint64String) MarshalText() (out []byte, err error) {
	out = strconv.AppendUint(out, uint64(v), 10)
	return
}

func (v *Uint64String) UnmarshalText(b []byte) error {
	n, err := strconv.ParseUint(string(b), 0, 64)
	if err != nil {
		return err
	}
	*v = Uint64String(n)
	return nil
}
