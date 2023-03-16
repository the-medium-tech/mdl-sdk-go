package contract

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransactionFactory(t *testing.T) {
	fabric := NewTransactionFactory(TransactionTypeToString(FABRIC))
	assert.Equal(t, "*contract.FabricContract", reflect.TypeOf(fabric).String())

	eth := NewTransactionFactory(TransactionTypeToString(ETHEREUM))
	assert.Equal(t, "*contract.EthereumContract", reflect.TypeOf(eth).String())

	btc := NewTransactionFactory(TransactionTypeToString(BITCOIN))
	assert.Equal(t, "*contract.BitcoinContract", reflect.TypeOf(btc).String())
}

func TestTransactionType(t *testing.T) {
	assert.Equal(t, "fabric", TransactionTypeToString(FABRIC))
	assert.Equal(t, "ethereum", TransactionTypeToString(ETHEREUM))
	assert.Equal(t, "bitcoin", TransactionTypeToString(BITCOIN))
}
