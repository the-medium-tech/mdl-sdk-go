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

var transactionTypeStrings = map[TransactionType]string{
	FABRIC:   "fabric",
	ETHEREUM: "ethereum",
	BITCOIN:  "bitcoin",
}

func TransactionTypeToString(id TransactionType) string {
	if res, found := transactionTypeStrings[id]; found {
		return res
	}
	return ""
}

type Transaction interface {
	SetMessage(msg *message)
	SubmitTransaction(contract *gateway.Contract, file, function string, args ...string) ([]byte, error)
	Verify() bool
	Address() string
}

func NewTransactionFactory(transactionType string) Transaction {
	switch transactionType {
	case TransactionTypeToString(FABRIC):
		return NewFabricContract()
	case TransactionTypeToString(ETHEREUM):
		return NewEthereumContract()
	case TransactionTypeToString(BITCOIN):
		return NewBitcoinContract()
	default:
	}
	return nil
}

func GetTransaction(message []byte) (Transaction, error) {
	msg := newMessage()
	if err := msg.Deserialize(message); err != nil {
		return nil, err
	}
	transaction := NewTransactionFactory(msg.Type)
	transaction.SetMessage(msg)
	return transaction, nil
}
