package svm

import (
	"context"
	"github.com/gagliardetto/solana-go"
	"wallet-go/constant"
)

func (c *Client) TransactionByHash(ctx context.Context, hash string) error {
	signature, err := solana.SignatureFromBase58(hash)
	if err != nil {
		return constant.InvalidTransactionHashErr
	}
	// TODO
	// c.sdkClient.GetConfirmedTransactionWithOpts(ctx, signature, &rpc.GetTransactionOpts{
	// 	Encoding:   solana.EncodingJSONParsed,
	// 	Commitment: rpc.CommitmentFinalized,
	// })
}
