package svm

import (
	"github.com/gagliardetto/solana-go"
	"strings"
	"wallet-go/constant"
)

func FormaterAddress(address string) (string, error) {
	if IsNativeAsset(address) {
		return address, nil
	}
	if _, err := solana.PublicKeyFromBase58(address); err != nil {
		return "", err
	}
	return address, nil
}

func NativeAssetAddress() string {
	return constant.SVMZeroAddress
}

func IsNativeAsset(contract string) bool {
	return strings.EqualFold(contract, constant.SVMZeroAddress)
}
