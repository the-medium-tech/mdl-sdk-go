package contract

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/the-medium-tech/mdl-sdk-go/contract/header"
)

func TestNewContractFactory(t *testing.T) {
	fabric := NewContractFactory(header.HeaderTypeToString(header.FABRIC))
	assert.Equal(t, "*contract.FabricContract", reflect.TypeOf(fabric).String())

	eth := NewContractFactory(header.HeaderTypeToString(header.ETHEREUM))
	assert.Equal(t, "*contract.EthereumContract", reflect.TypeOf(eth).String())

	btc := NewContractFactory(header.HeaderTypeToString(header.BITCOIN))
	assert.Equal(t, "*contract.BitcoinContract", reflect.TypeOf(btc).String())
}
