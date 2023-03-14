package contract

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

type TransactionType int

const (
	FABRIC TransactionType = iota
	ETHEREUM
	BITCOIN
)

type Transaction interface {
	SubmitTransaction(contract *gateway.Contract, function string, args ...string) ([]byte, error)
}

func NewTransactionFactory(transactionType TransactionType, file string) Transaction {
	switch transactionType {
	case FABRIC:
		return NewFabricContract()
	case ETHEREUM:
		if file != "" {
			return NewEthereumContract(LoadConfig(file))
		}
	case BITCOIN:
		if file != "" {
			return NewBitcoinContract(LoadConfig(file))
		}
	default:
	}
	return nil
}
