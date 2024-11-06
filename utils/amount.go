package utils

import (
	"errors"
	"github.com/shopspring/decimal"
	"golang.org/x/exp/constraints"
	"math/big"
)

const (
	ZeroExp = 0
)

var (
	NegativeAmount = errors.New("amount cannot be negative")
)

type ToNumber interface {
	string |
		constraints.Integer | constraints.Float |
		decimal.Decimal | *decimal.Decimal | *big.Int
}

// ToMinUnit convert the amount to the smallest unit
func ToMinUnit[T ToNumber](iAmount T, decimals uint8) *big.Int {
	amount := convert2Decimal(iAmount)
	return amount.Shift(int32(decimals)).BigInt()
}

// ToReadable convert the amount from the smallest unit to a more human-readable unit
func ToReadable[T ToNumber](iAmount T, decimals uint8) decimal.Decimal {
	amount := convert2Decimal(iAmount)
	return amount.Shift(-int32(decimals))
}

func convert2Decimal[T ToNumber](value T) decimal.Decimal {
	amount := decimal.NewFromFloat(0)
	switch v := any(value).(type) {
	case int:
		amount = decimal.NewFromInt(int64(v))
	case int8:
		amount = decimal.NewFromInt(int64(v))
	case int16:
		amount = decimal.NewFromInt(int64(v))
	case int32:
		amount = decimal.NewFromInt(int64(v))
	case int64:
		amount = decimal.NewFromInt(v)
	case uint:
		amount = decimal.NewFromUint64(uint64(v))
	case uint8:
		amount = decimal.NewFromUint64(uint64(v))
	case uint16:
		amount = decimal.NewFromUint64(uint64(v))
	case uint32:
		amount = decimal.NewFromUint64(uint64(v))
	case uint64:
		amount = decimal.NewFromUint64(v)
	case float32:
		amount = decimal.NewFromFloat32(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	case *big.Int:
		amount = decimal.NewFromBigInt(v, ZeroExp)
	case string:
		amount, _ = decimal.NewFromString(v)
	}
	return amount
}
