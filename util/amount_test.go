package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReadableAmountToMinUnit(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{"int", ReadableAmountToMinUnit(int(12), 2).String(), "1200"},
		{"int8", ReadableAmountToMinUnit(int8(12), 2).String(), "1200"},
		{"int16", ReadableAmountToMinUnit(int16(12), 2).String(), "1200"},
		{"int32", ReadableAmountToMinUnit(int32(12), 2).String(), "1200"},
		{"int64", ReadableAmountToMinUnit(int64(12), 2).String(), "1200"},
		{"uint", ReadableAmountToMinUnit(uint(12), 2).String(), "1200"},
		{"uint8", ReadableAmountToMinUnit(uint8(12), 2).String(), "1200"},
		{"uint16", ReadableAmountToMinUnit(uint16(12), 2).String(), "1200"},
		{"uint32", ReadableAmountToMinUnit(uint32(12), 2).String(), "1200"},
		{"uint64", ReadableAmountToMinUnit(uint64(12), 2).String(), "1200"},
		{"float32", ReadableAmountToMinUnit(float32(1.23), 2).String(), "123"},
		{"float64", ReadableAmountToMinUnit(1.23, 2).String(), "123"},
		{"string", ReadableAmountToMinUnit("1.23", 2).String(), "123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.got)
		})
	}
}

func TestMinUnitToReadableAmount(t *testing.T) {
	tests := []struct {
		name string
		got  string
		want string
	}{
		{"int", MinUnitToReadableAmount(int(1200), 2).String(), "12"},
		{"int8", MinUnitToReadableAmount(int8(120), 2).String(), "1.2"},
		{"int16", MinUnitToReadableAmount(int16(1200), 2).String(), "12"},
		{"int32", MinUnitToReadableAmount(int32(1200), 2).String(), "12"},
		{"int64", MinUnitToReadableAmount(int64(1200), 2).String(), "12"},
		{"uint", MinUnitToReadableAmount(uint(1200), 2).String(), "12"},
		{"uint8", MinUnitToReadableAmount(uint8(120), 2).String(), "1.2"},
		{"uint16", MinUnitToReadableAmount(uint16(1200), 2).String(), "12"},
		{"uint32", MinUnitToReadableAmount(uint32(1200), 2).String(), "12"},
		{"uint64", MinUnitToReadableAmount(uint64(1200), 2).String(), "12"},
		{"float32", MinUnitToReadableAmount(float32(123), 2).String(), "1.23"},
		{"float64", MinUnitToReadableAmount(123, 2).String(), "1.23"},
		{"string", MinUnitToReadableAmount("123", 2).String(), "1.23"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, tt.got)
		})
	}
}
