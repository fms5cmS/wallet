package util

import (
	"math/big"

	"github.com/shopspring/decimal"
	"golang.org/x/exp/constraints"
)

type AmountType interface {
	string | constraints.Float | constraints.Integer
}

// ReadableAmountToMinUnit converts a readable amount to its minimum unit representation.
// For example, converts "1.5" ETH to "1500000000000000000" wei (assuming 18 decimals).
//
// Parameters:
//   - iAmount: The readable amount to convert (can be string, float, or integer)
//   - decimals: The number of decimal places for the token
//
// Returns:
//   - *big.Int: The amount in minimum units (e.g., wei for ETH)
func ReadableAmountToMinUnit[T AmountType](iAmount T, decimals uint8) *big.Int {
	return convertToDecimal(iAmount).Shift(int32(decimals)).BigInt()
}

// MinUnitToReadableAmount converts a minimum unit amount to its readable representation.
// For example, converts "1500000000000000000" wei to "1.5" ETH (assuming 18 decimals).
//
// Parameters:
//   - iAmount: The amount in minimum units to convert (can be string, float, or integer)
//   - decimals: The number of decimal places for the token
//
// Returns:
//   - decimal.Decimal: The readable amount
func MinUnitToReadableAmount[T AmountType](iAmount T, decimals uint8) decimal.Decimal {
	return convertToDecimal(iAmount).Shift(-int32(decimals))
}

func convertToDecimal[T AmountType](iAmount T) decimal.Decimal {
	switch v := any(iAmount).(type) {
	case int:
		return decimal.NewFromInt(int64(v))
	case int8:
		return decimal.NewFromInt(int64(v))
	case int16:
		return decimal.NewFromInt(int64(v))
	case int32:
		return decimal.NewFromInt(int64(v))
	case int64:
		return decimal.NewFromInt(v)
	case uint:
		return decimal.NewFromUint64(uint64(v))
	case uint8:
		return decimal.NewFromUint64(uint64(v))
	case uint16:
		return decimal.NewFromUint64(uint64(v))
	case uint32:
		return decimal.NewFromUint64(uint64(v))
	case uint64:
		return decimal.NewFromUint64(v)
	case float32:
		return decimal.NewFromFloat32(v)
	case float64:
		return decimal.NewFromFloat(v)
	case string:
		amount, _ := decimal.NewFromString(v)
		return amount
	default:
		return decimal.Zero
	}
}
