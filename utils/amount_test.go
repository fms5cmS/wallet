package utils

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func TestToMinUnitAndToReadable(t *testing.T) {
	decimals := uint8(9)
	// string
	strVal := "2"
	strGet := ToMinUnit(strVal, decimals)
	strBig, _ := new(big.Int).SetString(strVal, 10)
	strWant := decimal.NewFromBigInt(strBig, int32(decimals))
	assert.Equal(t, strWant.String(), strGet.String())
	strReadable := ToReadable(strGet, decimals)
	assert.Equal(t, strVal, strReadable.String())
	// int
	intVal := 7
	intGet := ToMinUnit(intVal, decimals)
	intWant := decimal.NewFromBigInt(new(big.Int).SetInt64(int64(intVal)), int32(decimals))
	assert.Equal(t, intWant.String(), intGet.String())
	intReadable := ToReadable(intGet, decimals)
	assert.Equal(t, intVal, int(intReadable.IntPart()))
	// uint
	uintVal := uint(9)
	uintGet := ToMinUnit(uintVal, decimals)
	uintWant := decimal.NewFromBigInt(new(big.Int).SetUint64(uint64(uintVal)), int32(decimals))
	assert.Equal(t, uintWant.String(), uintGet.String())
	uintReadable := ToReadable(uintGet, decimals)
	assert.Equal(t, uintVal, uint(uintReadable.IntPart()))
	// decimal.Decimal
	decimalVal := decimal.NewFromFloat(2.3)
	decimalGet := ToMinUnit(decimalVal, decimals)
	decimalWant := decimalVal.Shift(int32(decimals))
	assert.Equal(t, decimalWant.String(), decimalGet.String())
	decimalReadable := ToReadable(decimalGet, decimals)
	assert.Equal(t, decimalVal.String(), decimalReadable.String())
	// *decimal.Decimal
	pointDecimalVal := &decimalVal
	pointDecimalGet := ToMinUnit(pointDecimalVal, decimals)
	assert.Equal(t, decimalWant.String(), pointDecimalGet.String())
	pointReadable := ToReadable(pointDecimalGet, decimals)
	assert.Equal(t, decimalVal.String(), pointReadable.String())
	// *big.Int
	bigIntVal := big.NewInt(5)
	bigIntGet := ToMinUnit(bigIntVal, decimals)
	bigIntWant := decimal.NewFromBigInt(bigIntVal, int32(decimals))
	assert.Equal(t, bigIntWant.String(), bigIntGet.String())
	bigIntReadable := ToReadable(bigIntGet, decimals)
	assert.Equal(t, bigIntVal.String(), bigIntReadable.String())
}
