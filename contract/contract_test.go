package contract

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestNoSignMessage(t *testing.T) {
	file := filepath.Join("testdata", "test.key")
	err := generateKey(file)
	assert.NoError(t, err)
	contract := NewContract(LoadConfig(file))
	contract.Function = "function"
	contract.Msg = &message{
		Args: []string{"1,2,3,4"},
	}
	assert.Equal(t, `{"args":["1,2,3,4"]}`, contract.string())
}

func TestSignedMessage(t *testing.T) {
	file := filepath.Join("testdata", "test.key")
	err := generateKey(file)
	assert.NoError(t, err)
	contract := NewContract(LoadConfig(file))
	contract.Function = "function"
	contract.Msg = &message{
		Args: []string{"1,2,3,4"},
	}
	err = contract.sign()
	assert.NoError(t, err)
	assert.NotEmpty(t, contract.Msg.Hash)
	assert.NotEmpty(t, contract.Msg.Signature)
}

func generateKey(file string) error {
	if _, err := os.Stat(file); err == os.ErrNotExist {
		return GenerateKey(file)
	}
	return nil
}
