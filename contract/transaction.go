package contract

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
	SetHeader(header *header)
	GetArgs(file, function string, args ...string) ([]string, error)
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

func GetTransaction(message [][]byte) (Transaction, error) {
	header := newHeader()
	if err := header.deserialize(message[0]); err != nil {
		return nil, err
	}
	transaction := NewTransactionFactory(header.Type)
	transaction.SetHeader(header)
	return transaction, nil
}
