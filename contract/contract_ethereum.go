package contract

import (
	"errors"
	"github.com/the-medium-tech/mdl-sdk-go/address"
	"github.com/the-medium-tech/mdl-sdk-go/internal/crypto"
)

type EthereumContract struct {
}

func newEthereumContract() *EthereumContract {
	return &EthereumContract{}
}

func (e *EthereumContract) Sign(hash []byte, keyFile string) ([]byte, error) {
	key, err := crypto.LoadECDSA(keyFile)
	if err != nil {
		return nil, err
	}
	return crypto.SignCompact(hash, key)
}

func (e *EthereumContract) Hash(data []byte) []byte {
	return crypto.Keccak256(data)
}

func (e *EthereumContract) Verify(a *address.Address) bool {
	return true
}

func (e *EthereumContract) ExtractAddress(a *address.Address) string {
	recoveredPub, err := crypto.Ecrecover(a.Hash, a.Signature)
	if err != nil {
		return ""
	}
	pubKey, err := crypto.UnmarshalPubkey(recoveredPub)
	if err != nil {
		return ""
	}
	return crypto.PubkeyToAddress(*pubKey).String()
}

func (e *EthereumContract) ExtractPublickey(keyFile string) ([]byte, error) {
	return nil, errors.New("ethereum contract does not support extract public key function")
}

func (e *EthereumContract) StringsToBytes(args []string) []byte {
	var temp string
	for _, arg := range args {
		temp += arg
	}
	return []byte(temp)
}
