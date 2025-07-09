package constant

import "github.com/ethereum/go-ethereum/signer/core/apitypes"

const EIP712Domain = "EIP712Domain"

var EIP712DomainPrimaryType = []apitypes.Type{
	{Name: "name", Type: "string"},
	{Name: "version", Type: "string"},
	{Name: "chainId", Type: "uint256"},
	{Name: "verifyingContract", Type: "address"},
}
