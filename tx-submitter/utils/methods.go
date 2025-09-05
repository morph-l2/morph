package utils

import (
	"fmt"
	"math/big"
	"reflect"
	"strconv"

	"github.com/morph-l2/go-ethereum"
	"github.com/morph-l2/go-ethereum/common"
	"github.com/morph-l2/go-ethereum/common/hexutil"
)

func ToCallArg(msg ethereum.CallMsg) interface{} {
	arg := map[string]interface{}{
		"from": msg.From,
		"to":   msg.To,
	}
	if len(msg.Data) > 0 {
		arg["data"] = hexutil.Bytes(msg.Data)
	}
	if msg.Value != nil {
		arg["value"] = (*hexutil.Big)(msg.Value)
	}
	if msg.Gas != 0 {
		arg["gas"] = hexutil.Uint64(msg.Gas)
	}
	if msg.GasPrice != nil {
		arg["gasPrice"] = (*hexutil.Big)(msg.GasPrice)
	}
	return arg
}

func MaxOfThreeBig(a, b, c *big.Int) *big.Int {
	max := a

	if b.Cmp(max) > 0 {
		max = b
	}

	if c.Cmp(max) > 0 {
		max = c
	}

	return max
}

// IntersectionOfAddresses
func IntersectionOfAddresses(a, b []common.Address) []common.Address {
	intersection := make([]common.Address, 0)
	aMap := make(map[common.Address]bool)

	// Create a map of addresses in set 'a' for faster lookup
	for _, addr := range a {
		aMap[addr] = true
	}

	// Check each address in 'b' against the map of 'a'
	for _, addr := range b {
		if aMap[addr] {
			intersection = append(intersection, addr)
		}
	}

	return intersection
}

func ParseStringToType[T any](s string) (T, error) {
	var result T
	var err error

	// 获取目标类型的名称
	switch any(result).(type) {
	case int:
		var v int64
		v, err = strconv.ParseInt(s, 10, 0)
		result = reflect.ValueOf(int(v)).Interface().(T)
	case int8:
		var v int64
		v, err = strconv.ParseInt(s, 10, 8)
		result = reflect.ValueOf(int8(v)).Interface().(T)
	case int16:
		var v int64
		v, err = strconv.ParseInt(s, 10, 16)
		result = reflect.ValueOf(int16(v)).Interface().(T)
	case int32:
		var v int64
		v, err = strconv.ParseInt(s, 10, 32)
		result = reflect.ValueOf(int32(v)).Interface().(T)
	case int64:
		var v int64
		v, err = strconv.ParseInt(s, 10, 64)
		result = reflect.ValueOf(v).Interface().(T)
	case uint:
		var v uint64
		v, err = strconv.ParseUint(s, 10, 0)
		result = reflect.ValueOf(uint(v)).Interface().(T)
	case uint8:
		var v uint64
		v, err = strconv.ParseUint(s, 10, 8)
		result = reflect.ValueOf(uint8(v)).Interface().(T)
	case uint16:
		var v uint64
		v, err = strconv.ParseUint(s, 10, 16)
		result = reflect.ValueOf(uint16(v)).Interface().(T)
	case uint32:
		var v uint64
		v, err = strconv.ParseUint(s, 10, 32)
		result = reflect.ValueOf(uint32(v)).Interface().(T)
	case uint64:
		var v uint64
		v, err = strconv.ParseUint(s, 10, 64)
		result = reflect.ValueOf(v).Interface().(T)
	case float32:
		var v float64
		v, err = strconv.ParseFloat(s, 32)
		result = reflect.ValueOf(float32(v)).Interface().(T)
	case float64:
		var v float64
		v, err = strconv.ParseFloat(s, 64)
		result = reflect.ValueOf(v).Interface().(T)
	case bool:
		var v bool
		v, err = strconv.ParseBool(s)
		result = reflect.ValueOf(v).Interface().(T)
	case string:
		result = reflect.ValueOf(s).Interface().(T)
	default:
		return *new(T), fmt.Errorf("unsupported type: %v", reflect.TypeOf(result))
	}

	if err != nil {
		return *new(T), err
	}

	return result, nil
}

// WeiToGwei converts wei value to gwei string representation
func WeiToGwei(wei *big.Int) string {
	if wei == nil {
		return "0"
	}
	gwei := new(big.Float).Quo(
		new(big.Float).SetInt(wei),
		new(big.Float).SetInt64(1e9),
	)
	return gwei.Text('f', 9)
}
