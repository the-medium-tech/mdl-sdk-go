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
	err := fab.makeTransaction(file, "function", []string{"1,2,3,4"}...)
	assert.NoError(t, err)

	transaction, err := GetTransaction([]byte(fab.string()))
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
	err = eth.makeTransaction(file, "function", []string{"1,2,3,4"}...)
	assert.NoError(t, err)

	transaction, err := GetTransaction([]byte(eth.string()))
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
	err = btc.makeTransaction(file, "function", []string{"1,2,3,4"}...)
	assert.NoError(t, err)

	transaction, err := GetTransaction([]byte(btc.string()))
	assert.NoError(t, err)
	assert.Equal(t, "*contract.BitcoinContract", reflect.TypeOf(transaction).String())
	assert.True(t, transaction.Verify())
	assert.Equal(t, "15VDTyzYK6SiH4kCdT89bEaskB15QS79F9", transaction.Address())
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
