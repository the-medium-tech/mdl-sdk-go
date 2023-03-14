package contract

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransactionFactory(t *testing.T) {
	file := filepath.Join("testdata", "test.key")

	fabric := NewTransactionFactory(FABRIC, "")
	assert.Equal(t, "*contract.FabricContract", reflect.TypeOf(fabric).String())

	nilEth := NewTransactionFactory(ETHEREUM, "")
	assert.Nil(t, nilEth)

	eth := NewTransactionFactory(ETHEREUM, file)
	assert.Equal(t, "*contract.EthereumContract", reflect.TypeOf(eth).String())

	nilBit := NewTransactionFactory(BITCOIN, "")
	assert.Nil(t, nilBit)

	btc := NewTransactionFactory(BITCOIN, file)
	assert.Equal(t, "*contract.BitcoinContract", reflect.TypeOf(btc).String())
}
