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
	// TransactionData get transaction data before sign and send
	TransactionData(ctx context.Context, req *types.TransactionDataReq, options ...types.TransactionDataOption) (*types.TransactionDataResp, error)
	// SendTransaction send transaction
	SendTransaction(ctx context.Context, req *types.SendTransactionReq, options ...types.TransactionDataOption) (*types.SendTransactionResp, error)
	// TransactionByHas get transaction by hash
	TransactionByHas(ctx context.Context, hash string)
}
