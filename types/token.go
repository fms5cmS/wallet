package types

import (
	"github.com/shopspring/decimal"
	"math/big"
)

type Token struct {
	Native      bool
	Contract    string
	Name        string   `json:"name"`
	Symbol      string   `json:"symbol"`
	Decimals    uint8    `json:"decimals"`
	TotalSupply *big.Int `json:"total_supply"`
}

type GetTokenBalanceReq struct {
	Owner    string
	Contract string
}

type TokenBalance struct {
	Contract string
	Decimals uint8
	MinUnit  *big.Int
	Readable decimal.Decimal
}
