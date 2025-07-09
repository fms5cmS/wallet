package util

import (
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"testing"
)

const (
	testPrivateKeyStr = "e84ae4ef6d329827ac01db8fc1b7595bf9ee2fe2c6cffd1ff783fc47d31311ff"
	testAddress       = "0xa4Bd37efCC3dB270aafa541b01845D712C5697B0"
)

func TestGetEVMAddressByPrivateKey(t *testing.T) {
	privateKey, err := LoadPrivateKey(testPrivateKeyStr)
	require.NoError(t, err)
	addr, err := GetEVMAddressByPrivateKey(privateKey)
	require.NoError(t, err)
	require.Equal(t, 0, addr.Cmp(ethCommon.HexToAddress(testAddress)))
}

func TestVerifyEVMSignature(t *testing.T) {
	privateKey, err := LoadPrivateKey(testPrivateKeyStr)
	require.NoError(t, err)
	message := "unsigned message"
	signature, err := SignEVMMessage(privateKey, message)
	require.NoError(t, err)
	verifyResult, err := VerifyEVMSignature(signature, message, testAddress)
	require.NoError(t, err)
	require.True(t, verifyResult)
}
