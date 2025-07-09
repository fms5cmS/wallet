package util

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"wallet/constant"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"

	"github.com/ethereum/go-ethereum/accounts"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// LoadPrivateKey loads an ECDSA private key object from a private key string
// Supports hexadecimal private key strings with or without "0x" prefix
// Uses secp256k1 elliptic curve algorithm
//
// Parameters:
//   - privateKeyStr: Hexadecimal private key string, can include "0x" prefix
//
// Returns:
//   - *ecdsa.PrivateKey: ECDSA private key object
//   - error: Returns error if private key format is invalid or parsing fails
func LoadPrivateKey(privateKeyStr string) (*ecdsa.PrivateKey, error) {
	privateKeyStr = strings.TrimPrefix(privateKeyStr, constant.HexPrefix)
	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, fmt.Errorf("load PrivateKey: privateKey HexToECDSA failed: %w", err)
	}
	return privateKey, nil
}

// GetEVMAddressByPrivateKey generates the corresponding Ethereum address from a private key
// Derives the public key from the private key, then calculates the Keccak-256 hash of the public key to get the address
// This is the standard process for Ethereum address generation
//
// Parameters:
//   - privateKey: ECDSA private key object
//
// Returns:
//   - ethCommon.Address: Ethereum address
//   - error: Returns error if public key type conversion fails
func GetEVMAddressByPrivateKey(privateKey *ecdsa.PrivateKey) (ethCommon.Address, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return ethCommon.Address{}, fmt.Errorf("load PrivateKey: publicKey is not of type *ecdsa.PublicKey")
	}
	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}

// SignEVMMessageHash signs a message hash using an ECDSA private key
// The hash should be provided as a hexadecimal string (with or without "0x" prefix)
// Returns the signature in hexadecimal format with "0x" prefix
// The signature includes the recovery ID for public key recovery
//
// Parameters:
//   - privateKey: ECDSA private key object used for signing
//   - hash: Hexadecimal string of the message hash to sign
//
// Returns:
//   - string: Hexadecimal signature with "0x" prefix
//   - error: Returns error if hash decoding or signing fails
func SignEVMMessageHash(privateKey *ecdsa.PrivateKey, hash string) (string, error) {
	hash = strings.TrimPrefix(hash, constant.HexPrefix)
	hashBytes, err := hex.DecodeString(hash)
	if err != nil {
		return "", fmt.Errorf("decode hash failed: %w", err)
	}
	sig, err := crypto.Sign(hashBytes, privateKey)
	if err != nil {
		return "", fmt.Errorf("sign failed: %w", err)
	}
	if len(sig) != crypto.SignatureLength {
		return "", fmt.Errorf("wrong size for signature: got %d, want %d", len(sig), crypto.SignatureLength)
	}
	if sig[crypto.RecoveryIDOffset] < 27 {
		sig[crypto.RecoveryIDOffset] += 27
	}
	return constant.HexPrefix + hex.EncodeToString(sig), nil
}

// SignEVMMessage signs a plain text message using an ECDSA private key
// Prepends the Ethereum message prefix to the message before hashing
// The prefix format is: "\x19Ethereum Signed Message:\n" + message length + message
// This follows the Ethereum personal message signing standard
//
// Parameters:
//   - privateKey: ECDSA private key object used for signing
//   - message: Plain text message to sign
//
// Returns:
//   - string: Hexadecimal signature with "0x" prefix
//   - error: Returns error if signing process fails
func SignEVMMessage(privateKey *ecdsa.PrivateKey, message string) (string, error) {
	prefixedMessage := constant.EVMMessagePrefix + strconv.Itoa(len(message)) + message
	hash := crypto.Keccak256Hash([]byte(prefixedMessage))
	return SignEVMMessageHash(privateKey, hash.String())
}

// VerifyEVMSignature verifies an Ethereum message signature
// Reconstructs the public key from the signature and message
// Compares the recovered address with the provided address
// This follows the Ethereum personal message verification standard
//
// Parameters:
//   - signatureHex: Hexadecimal signature string (with or without "0x" prefix)
//   - message: Original plain text message that was signed
//   - addr: Ethereum address in hexadecimal format to verify against
//
// Returns:
//   - bool: True if signature is valid and matches the provided address
//   - error: Returns error if signature decoding or verification fails
func VerifyEVMSignature(signatureHex, message string, addr string) (bool, error) {
	signature, err := hex.DecodeString(strings.TrimPrefix(signatureHex, constant.HexPrefix))
	if err != nil {
		return false, fmt.Errorf("decode signatureHex(%s) failed: %s", signatureHex, err)
	}
	msg := accounts.TextHash([]byte(message))
	signature[crypto.RecoveryIDOffset] -= 27
	publicKey, err := crypto.SigToPub(msg, signature)
	if err != nil {
		return false, fmt.Errorf("sigToPub failed: %s", err)
	}
	recoveredAddr := crypto.PubkeyToAddress(*publicKey)
	return recoveredAddr.Cmp(ethCommon.HexToAddress(addr)) == 0, nil
}

// GenerateEIP712Hash computes the EIP-712 hash for the given typed data
//
// Parameters:
//   - types: apitypes.Types, a map of all involved type definitions (including EIP712Domain and any nested types)
//   - domain: apitypes.TypedDataDomain, the EIP-712 domain struct
//   - primaryType: string, the root type name for the message
//   - typeDataMsg: apitypes.TypedDataMessage, the message data to hash
//
// Returns:
//   - string: The EIP-712 hash as a hex string (with 0x prefix)
//   - string: The raw data string used for signing
//   - error: Any error encountered during hashing
func GenerateEIP712Hash(
	types apitypes.Types,
	domain apitypes.TypedDataDomain,
	primaryType string,
	typeDataMsg apitypes.TypedDataMessage,
) (string, string, error) {
	// Make a shallow copy to avoid mutating the caller's map
	typesCopy := make(apitypes.Types, len(types))
	for k, v := range types {
		typesCopy[k] = v
	}
	// Ensure EIP712Domain type definition is present
	if _, ok := typesCopy[constant.EIP712Domain]; !ok {
		typesCopy[constant.EIP712Domain] = constant.EIP712DomainPrimaryType
	}
	data := apitypes.TypedData{
		Types:       typesCopy,
		PrimaryType: primaryType,
		Domain:      domain,
		Message:     typeDataMsg,
	}
	hashBytes, rawData, err := apitypes.TypedDataAndHash(data)
	if err != nil {
		return "", "", err
	}
	return hexutil.Encode(hashBytes), rawData, nil
}
