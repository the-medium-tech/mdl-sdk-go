package contract

import (
	"errors"
	"github.com/the-medium-tech/mdl-sdk-go/address"
	"github.com/the-medium-tech/mdl-sdk-go/internal/common/hexutil"
	"github.com/the-medium-tech/mdl-sdk-go/internal/crypto"
)

type EthereumContract struct {
}

func newEthereumContract() *EthereumContract {
	return &EthereumContract{}
}

func (e *EthereumContract) Sign(hash string, keyFile string) (string, error) {
	key, err := crypto.LoadECDSA(keyFile)
	if err != nil {
		return "", err
	}
	hashBytes, err := hexutil.Decode(hash)
	if err != nil {
		return "", err
	}
	signature, err := crypto.SignCompact(hashBytes, key)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(signature), nil
}

func (e *EthereumContract) Hash(data string) (string, error) {
	return hexutil.Encode(crypto.Keccak256([]byte(data))), nil
}

func (e *EthereumContract) Address(a *address.Address) (string, error) {
	hashBytes, err := a.HexToBytes(a.Hash)
	if err != nil {
		return "", err
	}
	signatureBytes, err := a.HexToBytes(a.Signature)
	if err != nil {
		return "", err
	}
	recoveredPub, err := crypto.Ecrecover(hashBytes, signatureBytes)
	if err != nil {
		return "", err
	}
	pubKey, err := crypto.UnmarshalPubkey(recoveredPub)
	if err != nil {
		return "", err
	}
	return crypto.PubkeyToAddress(*pubKey).String(), nil
}

func (e *EthereumContract) PublicKey(keyFile string) (string, error) {
	return "", errors.New("ethereum contract does not support extract public key function")
}

func (e *EthereumContract) StringsToBytes(args []string) []byte {
	var temp string
	for _, arg := range args {
		temp += arg
	}
	return []byte(temp)
}
