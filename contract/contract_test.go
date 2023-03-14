package contract

import (
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/the-medium-tech/mdl-sdk-go/internal/crypto"
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

func TestCompress(t *testing.T) {
	file := filepath.Join("testdata", "test.key")
	err := generateKey(file)
	assert.NoError(t, err)
	contract := NewContract(LoadConfig(file))
	contract.Function = "function"
	contract.Msg = &message{
		Args: []string{"1,2,3,4"},
	}
	err = contract.compress()
	assert.NoError(t, err)
	assert.NotEmpty(t, contract.Msg.PublicKey)
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
