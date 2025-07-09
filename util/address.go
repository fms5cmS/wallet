package util

import (
	"errors"
	"log"
	"regexp"

	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	tronAddr "github.com/fbsobreira/gotron-sdk/pkg/address"
)

var (
	evmRegex  = regexp.MustCompile("^0x[A-Fa-f0-9]{40}")
	tronRegex = regexp.MustCompile("T[A-Za-z1-9]{33}")
	svmRegex  = regexp.MustCompile("[A-Za-z0-9]]{43,44}")
)

// ErrInvalidAddress is returned when an address format is invalid
var ErrInvalidAddress = errors.New("invalid address")

// ValidAddress checks if the given address string is valid for any supported blockchain.
// It validates addresses for EVM chains (0x...), Tron (T...), and SVM chains.
//
// Parameters:
//   - address: The address string to validate
//
// Returns:
//   - bool: true if the address is valid, false otherwise
func ValidAddress(address string) bool {
	return evmRegex.MatchString(address) || tronRegex.MatchString(address) || svmRegex.MatchString(address)
}

// TronToEVMAddress converts a Tron address to its equivalent EVM address format.
// If the input is already a valid EVM address, it returns the address as-is.
//
// Parameters:
//   - srcAddress: The Tron address to convert (should start with 'T')
//
// Returns:
//   - string: The converted EVM address in hex format (0x...)
//   - error: ErrInvalidAddress if the input address is invalid
func TronToEVMAddress(srcAddress string) (string, error) {
	if !tronRegex.MatchString(srcAddress) && !evmRegex.MatchString(srcAddress) {
		return "", ErrInvalidAddress
	}
	if ethCommon.IsHexAddress(srcAddress) {
		return ethCommon.HexToAddress(srcAddress).Hex(), nil
	}
	addr, err := tronAddr.Base58ToAddress(srcAddress)
	if err != nil {
		log.Printf("invalid tron address(%s): %v\n", srcAddress, err)
		return "", ErrInvalidAddress
	}
	return ethCommon.HexToAddress(addr.Hex()).Hex(), nil
}

// EVMToTronAddress converts an EVM address to its equivalent Tron address format.
// The conversion adds the Tron byte prefix if not already present.
//
// Parameters:
//   - srcAddress: The EVM address to convert (should be in 0x... format)
//
// Returns:
//   - string: The converted Tron address in base58 format (T...)
//   - error: ErrInvalidAddress if the input address is invalid
func EVMToTronAddress(srcAddress string) (string, error) {
	if !evmRegex.MatchString(srcAddress) {
		return "", ErrInvalidAddress
	}
	hexAddrBytes := ethCommon.FromHex(srcAddress)
	addrHex := hexAddrBytes
	if addrHex[0] != tronAddr.TronBytePrefix {
		addrHex = append([]byte{tronAddr.TronBytePrefix}, hexAddrBytes...)
	}
	return tronAddr.HexToAddress(hexutil.Encode(addrHex)).String(), nil
}
