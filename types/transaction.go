package types

import "math/big"

type TransactionDataReq struct {
	// base parameters
	Sender   string
	Receiver string
	Value    *big.Int
	// parameters for smart contract transaction
	ContractTxn *ContractTransaction
	// fee
	Fee *EstimateFee
}

type TransactionDataResp struct {
	TransactionData []byte
	DataForSign     []byte
}

type ContractTransaction struct {
	Contract string
	Data     []byte
}

type EstimateFee struct {
	UnitLimit         uint64   // fee compute unit limit
	PriorityUnitPrice *big.Int // priority unit price to faster execution
	BaseOrStaticFee   *big.Int // base fee or static fee price
}

type TransactionDataOption func(*TransactionDataReq)

func WithContractData(info *ContractTransaction) TransactionDataOption {
	return func(req *TransactionDataReq) {
		req.ContractTxn = info
	}
}

func WithFeeData(fee *EstimateFee) TransactionDataOption {
	return func(req *TransactionDataReq) {
		req.Fee = fee
	}
}

type SendTransactionReq struct {
	TransactionData []byte
	Signature       []byte
}

type SendTransactionResp struct {
	TxId []byte
}

type SendTransactionOption func(*SendTransactionReq)

func WithSignature(signature []byte) SendTransactionOption {
	return func(req *SendTransactionReq) {
		req.Signature = signature
	}
}
