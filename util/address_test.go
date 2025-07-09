package util

import (
	ethCommon "github.com/ethereum/go-ethereum/common"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidAddress(t *testing.T) {
	tests := []struct {
		name    string
		address string
		want    bool
	}{
		{
			name:    "valid EVM address",
			address: "0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b6",
			want:    true,
		},
		{
			name:    "valid EVM address lowercase",
			address: "0x742d35cc6634c0532925a3b8d4c9db96c4b4d8b6",
			want:    true,
		},
		{
			name:    "valid Tron address",
			address: "TJRabPrwbZy45sbavfcjinPJC18kjpRTv8",
			want:    true,
		},
		{
			name:    "valid SVM address",
			address: "AxiW9VLvTjLw6LS7uZqioFWbq2wWLCVcyy48pnM3UiLs",
			want:    true,
		},
		{
			name:    "invalid address - too short",
			address: "0x123",
			want:    false,
		},
		{
			name:    "invalid address - wrong prefix",
			address: "1x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b6",
			want:    false,
		},
		{
			name:    "invalid address - empty string",
			address: "",
			want:    false,
		},
		{
			name:    "invalid address - random string",
			address: "not-an-address",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidAddress(tt.address)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestTronToEVMAddress(t *testing.T) {
	tests := []struct {
		name        string
		srcAddress  string
		wantAddress string
		wantErr     bool
	}{
		{
			name:        "valid Tron address",
			srcAddress:  "TJRabPrwbZy45sbavfcjinPJC18kjpRTv8",
			wantAddress: ethCommon.HexToAddress("0x5CBDD86A2FA8DC4BDDD8A8F69DBA48572EEC07FB").Hex(),
			wantErr:     false,
		},
		{
			name:        "already EVM address",
			srcAddress:  "0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b6",
			wantAddress: ethCommon.HexToAddress("0x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b6").Hex(),
			wantErr:     false,
		},
		{
			name:        "invalid Tron address",
			srcAddress:  "invalid-tron-address",
			wantAddress: "",
			wantErr:     true,
		},
		{
			name:        "empty string",
			srcAddress:  "",
			wantAddress: "",
			wantErr:     true,
		},
		{
			name:        "invalid EVM address",
			srcAddress:  "0x123",
			wantAddress: "",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAddress, gotErr := TronToEVMAddress(tt.srcAddress)

			if tt.wantErr {
				require.Error(t, gotErr)
				require.Equal(t, ErrInvalidAddress, gotErr)
				require.Empty(t, gotAddress)
			} else {
				require.NoError(t, gotErr)
				require.Equal(t, tt.wantAddress, gotAddress)
			}
		})
	}
}

func TestEVMToTronAddress(t *testing.T) {
	tests := []struct {
		name        string
		srcAddress  string
		wantAddress string
		wantErr     bool
	}{
		{
			name:        "valid EVM address",
			srcAddress:  "0x5CBDD86A2FA8DC4BDDD8A8F69DBA48572EEC07FB",
			wantAddress: "TJRabPrwbZy45sbavfcjinPJC18kjpRTv8",
			wantErr:     false,
		},
		{
			name:        "valid EVM address with different case",
			srcAddress:  "0x260e5b32aF70E1189Bd4779D4145f47afeCC2387",
			wantAddress: "TDSRsqYS87WJgRsrH6xS6jZcuoNU1RaTn1",
			wantErr:     false,
		},
		{
			name:        "invalid EVM address - too short",
			srcAddress:  "0x123",
			wantAddress: "",
			wantErr:     true,
		},
		{
			name:        "invalid EVM address - wrong prefix",
			srcAddress:  "1x742d35Cc6634C0532925a3b8D4C9db96C4b4d8b6",
			wantAddress: "",
			wantErr:     true,
		},
		{
			name:        "invalid EVM address - empty string",
			srcAddress:  "",
			wantAddress: "",
			wantErr:     true,
		},
		{
			name:        "invalid EVM address - random string",
			srcAddress:  "not-an-evm-address",
			wantAddress: "",
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAddress, gotErr := EVMToTronAddress(tt.srcAddress)

			if tt.wantErr {
				require.Error(t, gotErr)
				require.Equal(t, ErrInvalidAddress, gotErr)
				require.Empty(t, gotAddress)
			} else {
				require.NoError(t, gotErr)
				require.Equal(t, tt.wantAddress, gotAddress)
			}
		})
	}
}

func TestAddressConversionRoundTrip(t *testing.T) {
	// Test that converting from Tron to EVM and back gives the same result
	tronAddress := "TJRabPrwbZy45sbavfcjinPJC18kjpRTv8"

	// Convert Tron to EVM
	evmAddress, err := TronToEVMAddress(tronAddress)
	require.NoError(t, err)
	require.NotEmpty(t, evmAddress)

	// Convert EVM back to Tron
	convertedTronAddress, err := EVMToTronAddress(evmAddress)
	require.NoError(t, err)
	require.NotEmpty(t, convertedTronAddress)

	// Should be the same
	require.Equal(t, tronAddress, convertedTronAddress)
}
