package client

import (
	"context"
	"wallet-go/types"
)

type BaseClient interface {
	// NativeToken get native token info
	NativeToken() *types.Token
	// NativeTokenBalance get native token balance
	NativeTokenBalance(ctx context.Context, req *types.GetTokenBalanceReq) *types.TokenBalance
	// GetTransactionData get transaction data before sign and send
	GetTransactionData(ctx context.Context, req *types.GetTransactionReq, options ...types.GetTransactionOption) (*types.GetTransactionResp, error)
	// SendTransaction send transaction
	SendTransaction(ctx context.Context, req *types.SendTransactionReq, options ...types.GetTransactionOption) (*types.SendTransactionResp, error)
}
