package contract

import (
	"encoding/hex"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/the-medium-tech/mdl-sdk-go/contract/header"
	"github.com/the-medium-tech/mdl-sdk-go/internal/crypto"
)

func TestUserScenarioForFabricContract(t *testing.T) {
	file := filepath.Join("testdata", "cert.pem")
	fab := NewFabricContract()
	h1, err := fab.Header(file)
	assert.NoError(t, err)
	args := fab.GetArgs(h1, []string{"1,2,3,4"}...)

	bytes := StringArrayToTwoDimensionalArray(args)

	h2, err := header.UnmarshaledHeader(bytes[0])
	assert.NoError(t, err)
	c := GetContract(h2)
	assert.Equal(t, "*contract.FabricContract", reflect.TypeOf(c).String())
	assert.True(t, header.Verify(h2))
	assert.Equal(t, "0x4057cc4274523666fa4cc88e5f78193b36105a33", header.Address(h2))
}

func TestUserScenarioForEthereumContract(t *testing.T) {
	file := filepath.Join("testdata", "test.key")
	err := generateKey(file)
	assert.NoError(t, err)

	eth := NewEthereumContract()
	h1, err := eth.Header(file)
	assert.NoError(t, err)
	args := eth.GetArgs(h1, []string{"1,2,3,4"}...)

	bytes := StringArrayToTwoDimensionalArray(args)

	h2, err := header.UnmarshaledHeader(bytes[0])
	assert.NoError(t, err)
	c := GetContract(h2)
	assert.Equal(t, "*contract.EthereumContract", reflect.TypeOf(c).String())
	assert.True(t, header.Verify(h2))
	assert.Equal(t, "0x93b2Cb3061e36Ed3099d003fF78cd685b424e95b", header.Address(h2))
}

func TestUserScenarioForBitcoinContract(t *testing.T) {
	file := filepath.Join("testdata", "test.key")
	err := generateKey(file)
	assert.NoError(t, err)

	btc := NewBitcoinContract()
	h1, err := btc.Header(file)
	assert.NoError(t, err)
	args := btc.GetArgs(h1, []string{"1,2,3,4"}...)

	bytes := StringArrayToTwoDimensionalArray(args)

	h2, err := header.UnmarshaledHeader(bytes[0])
	assert.NoError(t, err)
	c := GetContract(h2)
	assert.NoError(t, err)
	assert.Equal(t, "*contract.BitcoinContract", reflect.TypeOf(c).String())
	assert.True(t, header.Verify(h2))
	assert.Equal(t, "15VDTyzYK6SiH4kCdT89bEaskB15QS79F9", header.Address(h2))
}

func TestUserScenarioForNilContract(t *testing.T) {
	bytes := StringArrayToTwoDimensionalArray([]string{"1,2,3,4"})

	h, err := header.UnmarshaledHeader(bytes[0])
	assert.Error(t, err)
	c := GetContract(h)
	assert.Nil(t, c)
}

func generateKey(file string) error {
	if _, err := os.Stat(file); err == os.ErrNotExist {
		key, err := crypto.GenerateKey()
		if err != nil {
			return err
		}
		return os.WriteFile(file, []byte(hex.EncodeToString(crypto.FromECDSA(key))), 0600)
	}
	return nil
}
