package contract

import (
	"encoding/hex"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/the-medium-tech/mdl-sdk-go/internal/crypto"
)

func TestUserScenarioForFabricContract(t *testing.T) {
	file := filepath.Join("testdata", "cert.pem")
	fab := NewFabricContract()
	args, err := fab.GetArgs(file, "function", []string{"1,2,3,4"}...)
	assert.NoError(t, err)

	bytes := StringArrayToTwoDimensionalArray(args)
	transaction, err := GetTransaction(bytes)
	assert.NoError(t, err)
	assert.Equal(t, "*contract.FabricContract", reflect.TypeOf(transaction).String())
	assert.True(t, transaction.Verify())
	assert.Equal(t, "0x4057cc4274523666fa4cc88e5f78193b36105a33", transaction.Address())
}

func TestUserScenarioForEthereumContract(t *testing.T) {
	file := filepath.Join("testdata", "test.key")
	err := generateKey(file)
	assert.NoError(t, err)

	eth := NewEthereumContract()
	args, err := eth.GetArgs(file, "function", []string{"1,2,3,4"}...)
	assert.NoError(t, err)

	bytes := StringArrayToTwoDimensionalArray(args)
	transaction, err := GetTransaction(bytes)
	assert.NoError(t, err)
	assert.Equal(t, "*contract.EthereumContract", reflect.TypeOf(transaction).String())
	assert.True(t, transaction.Verify())
	assert.Equal(t, "0x93b2Cb3061e36Ed3099d003fF78cd685b424e95b", transaction.Address())
}

func TestUserScenarioForBitcoinContract(t *testing.T) {
	file := filepath.Join("testdata", "test.key")
	err := generateKey(file)
	assert.NoError(t, err)

	btc := NewBitcoinContract()
	args, err := btc.GetArgs(file, "function", []string{"1,2,3,4"}...)
	assert.NoError(t, err)

	bytes := StringArrayToTwoDimensionalArray(args)
	transaction, err := GetTransaction(bytes)
	assert.NoError(t, err)
	assert.Equal(t, "*contract.BitcoinContract", reflect.TypeOf(transaction).String())
	assert.True(t, transaction.Verify())
	assert.Equal(t, "15VDTyzYK6SiH4kCdT89bEaskB15QS79F9", transaction.Address())
}

func TestUserScenarioForNilContract(t *testing.T) {
	bytes := StringArrayToTwoDimensionalArray([]string{"1,2,3,4"})
	transaction, err := GetTransaction(bytes)
	assert.Error(t, err)
	assert.Nil(t, transaction)
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
