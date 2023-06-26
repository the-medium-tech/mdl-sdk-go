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

func TestBridgeScenario(t *testing.T) {
	fabric := "{\"type\":\"fabric\",\"publicKey\":\"0x663a3a56fd753f889a19c6ed4057cc4274523666fa4cc88e5f78193b36105a33\"}"
	deserializedFabricAddress, err := address.Deserialize([]byte(fabric))
	assert.NoError(t, err)
	deserializedFabric, err := GetContract(deserializedFabricAddress)
	assert.NoError(t, err)
	extractedAddress, err := deserializedFabric.Address(deserializedFabricAddress)
	assert.NoError(t, err)
	assert.Equal(t, "0x4057cc4274523666fa4cc88e5f78193b36105a33", extractedAddress)

	ethereum := "{\"type\":\"ethereum\",\"hash\":\"0xb82f66edc8ab69712adc0aa435956a9db69da12df3cd7e277277603cca8b3d42\",\"signature\":\"0x7cb7fe069294f9f6fbc3603312bcbdcb6372c5dcf0eacda47af51a3e6d0d950c637c957fa2d19ef0b8b5eebf2f77064ec8d99e9856eff9c8e95d2112d0fa309e01\"}"
	deserializedEthereumAddress, err := address.Deserialize([]byte(ethereum))
	assert.NoError(t, err)
	deserializedEthereum, err := GetContract(deserializedEthereumAddress)
	assert.NoError(t, err)
	extractedAddress, err = deserializedEthereum.Address(deserializedEthereumAddress)
	assert.NoError(t, err)
	assert.Equal(t, "0x93b2Cb3061e36Ed3099d003fF78cd685b424e95b", extractedAddress)
	bitcoin := "{\"type\":\"bitcoin\",\"hash\":\"0x289aed9cd03bbe0667af67013fcb5482c735eaa600ab94d232ff23425a9bd824\",\"signature\":\"0x30440220cfc2f04986dfa0c27026df63b7cb01e2b51be9e9a9ce8b31806af1b8adac4e8c02207b5ff5eb0405d3d8d41f3e7a117181ee9e2f243494722d0cb7851a175f969a9d\",\"publicKey\":\"0x036d8c7153055fcaa97796451185bbfd62904150319f5260e87336ac24cae04935\"}"
	deserializedBitcoinAddress, err := address.Deserialize([]byte(bitcoin))
	assert.NoError(t, err)
	deserializedBitcoin, err := GetContract(deserializedBitcoinAddress)
	assert.NoError(t, err)
	extractedAddress, err = deserializedBitcoin.Address(deserializedBitcoinAddress)
	assert.NoError(t, err)
	assert.Equal(t, "15VDTyzYK6SiH4kCdT89bEaskB15QS79F9", extractedAddress)
	rigo := "{\"type\":\"rigo\",\"hash\":\"0x34802bef80f8e1c2a580625d4e93466020d0034cdf9351064ff3bcbb5e17a224\",\"signature\":\"0x09b6cd62dc627e79eeab2a65806be88b60b3f1eaf65b22e577de5793a59706311246ac6a55276c24b05395ef656be0f699d07e9f4e1f4cb62edb404f42623fea\",\"publicKey\":\"0x036d8c7153055fcaa97796451185bbfd62904150319f5260e87336ac24cae04935\"}"
	deserializedRigoAddress, err := address.Deserialize([]byte(rigo))
	assert.NoError(t, err)
	deserializedRigo, err := GetContract(deserializedRigoAddress)
	assert.NoError(t, err)
	extractedAddress, err = deserializedRigo.Address(deserializedRigoAddress)
	assert.NoError(t, err)
	assert.Equal(t, "31368e3591140f51180fdde81993e22050dc19f8", extractedAddress)
}
