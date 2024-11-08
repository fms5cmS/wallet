package constant

import "errors"

var (
	InvalidConfigErr          = errors.New("invalid config")
	InvalidAddressErr         = errors.New("invalid address")
	InvalidTransactionHashErr = errors.New("invalid transaction hash")
)
