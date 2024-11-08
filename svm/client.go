package svm

import (
	_ "github.com/btcsuite/btcutil/base58"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/pkg/errors"
	"strings"
	"wallet-go/client"
	"wallet-go/constant"
	"wallet-go/types"
)

type Client struct {
	sdkClient   *rpc.Client
	nativeToken *types.Token
}

func NewClient(cfg *client.Config, options ...client.ConfigOption) (*Client, error) {
	for _, option := range options {
		option(cfg)
	}
	rpcEndpoint := strings.TrimSpace(cfg.Endpoints[client.EndpointTypeRPC])
	if rpcEndpoint == constant.EmptyStr {
		return nil, errors.Wrap(constant.InvalidConfigErr, "rpc endpoint is empty")
	}
	sdkClient := rpc.New(rpcEndpoint)
	return &Client{sdkClient: sdkClient, nativeToken: &types.Token{
		Native:      true,
		Contract:    constant.SVMZeroAddress,
		Name:        cfg.NativeTokenSymbol,
		Symbol:      cfg.NativeTokenSymbol,
		Decimals:    cfg.NativeTokenDecimals,
		TotalSupply: nil,
	}}, nil
}
