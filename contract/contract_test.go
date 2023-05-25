package contract

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/the-medium-tech/mdl-sdk-go/address"
)

func TestNewContract(t *testing.T) {
	fabric, err := NewContract(address.AddressTypeToString(address.FABRIC))
	assert.NoError(t, err)
	assert.Equal(t, "*contract.FabricContract", reflect.TypeOf(fabric).String())

	eth, err := NewContract(address.AddressTypeToString(address.ETHEREUM))
	assert.NoError(t, err)
	assert.Equal(t, "*contract.EthereumContract", reflect.TypeOf(eth).String())

	btc, err := NewContract(address.AddressTypeToString(address.BITCOIN))
	assert.NoError(t, err)
	assert.Equal(t, "*contract.BitcoinContract", reflect.TypeOf(btc).String())
}

func TestFabricContract(t *testing.T) {

	// create fabric contract
	fabric, err := NewContract(address.AddressTypeToString(address.FABRIC))
	assert.NoError(t, err)

	// extract public key from fabric certificate
	file := filepath.Join("testdata", "cert.pem")
	pub, err := fabric.ExtractPublickey(file)
	assert.NoError(t, err)

	// create address with public key
	addr, err := address.NewAddress(address.FABRIC, pub, nil, nil)
	assert.NoError(t, err)
	var args = []string{"1", "2"}
	mdlArgs, err := addr.AppendArgs(args)
	assert.NoError(t, err)
	addressBytes, err := addr.Serialize()
	assert.NoError(t, err)
	expectedMDLArgs := []string{string(addressBytes), "1", "2"}
	assert.Equal(t, expectedMDLArgs, mdlArgs)

	deserializedAddress, err := address.Deserialize(addressBytes)
	assert.NoError(t, err)

	deserializedFabric, err := GetContract(deserializedAddress)
	assert.NoError(t, err)

	ok := deserializedFabric.Verify(deserializedAddress)
	assert.True(t, ok)
	extractedAddress := deserializedFabric.ExtractAddress(deserializedAddress)
	assert.Equal(t, "0x4057cc4274523666fa4cc88e5f78193b36105a33", extractedAddress)
}

func TestEthereumContract(t *testing.T) {

	// create fabric contract
	ethereum, err := NewContract(address.AddressTypeToString(address.ETHEREUM))
	assert.NoError(t, err)

	// extract public key from fabric certificate
	file := filepath.Join("testdata", "test.key")
	var args = []string{"1", "2"}
	hash := ethereum.Hash(ethereum.StringsToBytes(args))
	signature, err := ethereum.Sign(hash, file)
	assert.NoError(t, err)

	_, err = ethereum.ExtractPublickey(file)
	assert.Error(t, err)

	// create address with hash, signature
	addr, err := address.NewAddress(address.ETHEREUM, nil, hash, signature)
	assert.NoError(t, err)
	mdlArgs, err := addr.AppendArgs(args)
	assert.NoError(t, err)
	addressBytes, err := addr.Serialize()
	assert.NoError(t, err)
	expectedMDLArgs := []string{string(addressBytes), "1", "2"}
	assert.Equal(t, expectedMDLArgs, mdlArgs)

	deserializedAddress, err := address.Deserialize(addressBytes)
	assert.NoError(t, err)

	deserializedEthereum, err := GetContract(deserializedAddress)
	assert.NoError(t, err)

	ok := deserializedEthereum.Verify(deserializedAddress)
	assert.True(t, ok)
	extractedAddress := deserializedEthereum.ExtractAddress(deserializedAddress)
	assert.Equal(t, "0x93b2Cb3061e36Ed3099d003fF78cd685b424e95b", extractedAddress)
}

func TestBitcoinContract(t *testing.T) {

	// create fabric contract
	bitcoin, err := NewContract(address.AddressTypeToString(address.BITCOIN))
	assert.NoError(t, err)

	// extract public key from fabric certificate
	file := filepath.Join("testdata", "test.key")
	var args = []string{"1", "2"}
	hash := bitcoin.Hash(bitcoin.StringsToBytes(args))
	signature, err := bitcoin.Sign(hash, file)
	assert.NoError(t, err)

	pub, err := bitcoin.ExtractPublickey(file)
	assert.NoError(t, err)

	// create address with public key, hash, signature
	addr, err := address.NewAddress(address.BITCOIN, pub, hash, signature)
	assert.NoError(t, err)
	mdlArgs, err := addr.AppendArgs(args)
	assert.NoError(t, err)
	addressBytes, err := addr.Serialize()
	assert.NoError(t, err)
	expectedMDLArgs := []string{string(addressBytes), "1", "2"}
	assert.Equal(t, expectedMDLArgs, mdlArgs)

	deserializedAddress, err := address.Deserialize(addressBytes)
	assert.NoError(t, err)

	deserializedBitcoin, err := GetContract(deserializedAddress)
	assert.NoError(t, err)
	
	ok := deserializedBitcoin.Verify(deserializedAddress)
	assert.True(t, ok)
	extractedAddress := deserializedBitcoin.ExtractAddress(deserializedAddress)
	assert.Equal(t, "15VDTyzYK6SiH4kCdT89bEaskB15QS79F9", extractedAddress)
}
