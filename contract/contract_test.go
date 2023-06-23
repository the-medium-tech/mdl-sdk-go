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
	pub, err := fabric.PublicKey(file)
	assert.NoError(t, err)

	// create address with public key
	addr, err := address.NewAddress(address.FABRIC, pub, "", "")
	assert.NoError(t, err)
	var args = []string{"1", "2"}
	mdlArgs, err := addr.AppendArgs(args)
	assert.NoError(t, err)

	addressBytes, err := addr.Serialize()
	assert.NoError(t, err)
	t.Log(string(addressBytes))

	expectedMDLArgs := []string{string(addressBytes), "1", "2"}
	assert.Equal(t, expectedMDLArgs, mdlArgs)

	deserializedAddress, err := address.Deserialize(addressBytes)
	assert.NoError(t, err)

	deserializedFabric, err := GetContract(deserializedAddress)
	assert.NoError(t, err)

	extractedAddress, err := deserializedFabric.Address(deserializedAddress)
	assert.NoError(t, err)
	assert.Equal(t, "0x4057cc4274523666fa4cc88e5f78193b36105a33", extractedAddress)
}

func TestEthereumContract(t *testing.T) {

	// create fabric contract
	ethereum, err := NewContract(address.AddressTypeToString(address.ETHEREUM))
	assert.NoError(t, err)

	// extract public key from fabric certificate
	file := filepath.Join("testdata", "test.key")
	var args = []string{"1", "2"}
	hash, err := ethereum.Hash("247c0fc35ab8459f760210fc4c0fd2a2eebfc32a123538e0724582825ac95948")
	assert.NoError(t, err)
	signature, err := ethereum.Sign(hash, file)
	assert.NoError(t, err)

	_, err = ethereum.PublicKey(file)
	assert.Error(t, err)

	// create address with hash, signature
	addr, err := address.NewAddress(address.ETHEREUM, "", hash, signature)
	assert.NoError(t, err)

	mdlArgs, err := addr.AppendArgs(args)
	assert.NoError(t, err)

	addressBytes, err := addr.Serialize()
	assert.NoError(t, err)
	t.Log(string(addressBytes))

	expectedMDLArgs := []string{string(addressBytes), "1", "2"}
	assert.Equal(t, expectedMDLArgs, mdlArgs)

	deserializedAddress, err := address.Deserialize(addressBytes)
	assert.NoError(t, err)

	deserializedEthereum, err := GetContract(deserializedAddress)
	assert.NoError(t, err)

	extractedAddress, err := deserializedEthereum.Address(deserializedAddress)
	assert.NoError(t, err)
	assert.Equal(t, "0x93b2Cb3061e36Ed3099d003fF78cd685b424e95b", extractedAddress)
}

func TestBitcoinContract(t *testing.T) {

	// create fabric contract
	bitcoin, err := NewContract(address.AddressTypeToString(address.BITCOIN))
	assert.NoError(t, err)

	// extract public key from fabric certificate
	file := filepath.Join("testdata", "test.key")
	var args = []string{"1", "2"}
	hash, err := bitcoin.Hash("247c0fc35ab8459f760210fc4c0fd2a2eebfc32a123538e0724582825ac95948")
	assert.NoError(t, err)

	signature, err := bitcoin.Sign(hash, file)
	assert.NoError(t, err)

	pub, err := bitcoin.PublicKey(file)
	assert.NoError(t, err)

	// create address with public key, hash, signature
	addr, err := address.NewAddress(address.BITCOIN, pub, hash, signature)
	assert.NoError(t, err)

	mdlArgs, err := addr.AppendArgs(args)
	assert.NoError(t, err)

	addressBytes, err := addr.Serialize()
	assert.NoError(t, err)
	t.Log(string(addressBytes))

	expectedMDLArgs := []string{string(addressBytes), "1", "2"}
	assert.Equal(t, expectedMDLArgs, mdlArgs)

	deserializedAddress, err := address.Deserialize(addressBytes)
	assert.NoError(t, err)

	deserializedBitcoin, err := GetContract(deserializedAddress)
	assert.NoError(t, err)

	extractedAddress, err := deserializedBitcoin.Address(deserializedAddress)
	assert.NoError(t, err)
	assert.Equal(t, "15VDTyzYK6SiH4kCdT89bEaskB15QS79F9", extractedAddress)
}

func TestRigoContract(t *testing.T) {

	rigo, err := NewContract(address.AddressTypeToString(address.RIGO))
	assert.NoError(t, err)

	// extract public key from fabric certificate
	file := filepath.Join("testdata", "test.key")
	var args = []string{"1", "2"}
	hash, err := rigo.Hash("247c0fc35ab8459f760210fc4c0fd2a2eebfc32a123538e0724582825ac95948")
	assert.NoError(t, err)

	signature, err := rigo.Sign(hash, file)
	assert.NoError(t, err)

	pub, err := rigo.PublicKey(file)
	assert.NoError(t, err)

	// create address with public key, hash, signature
	addr, err := address.NewAddress(address.RIGO, pub, hash, signature)
	assert.NoError(t, err)

	mdlArgs, err := addr.AppendArgs(args)
	assert.NoError(t, err)

	addressBytes, err := addr.Serialize()
	assert.NoError(t, err)
	t.Log(string(addressBytes))

	expectedMDLArgs := []string{string(addressBytes), "1", "2"}
	assert.Equal(t, expectedMDLArgs, mdlArgs)

	deserializedAddress, err := address.Deserialize(addressBytes)
	assert.NoError(t, err)

	deserializedBitcoin, err := GetContract(deserializedAddress)
	assert.NoError(t, err)

	extractedAddress, err := deserializedBitcoin.Address(deserializedAddress)
	assert.NoError(t, err)
	assert.Equal(t, "31368e3591140f51180fdde81993e22050dc19f8", extractedAddress)
}
