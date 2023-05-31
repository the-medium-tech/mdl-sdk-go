package contract

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"errors"

	"github.com/the-medium-tech/mdl-sdk-go/address"
	"github.com/the-medium-tech/mdl-sdk-go/internal/common"
	"github.com/the-medium-tech/mdl-sdk-go/internal/common/hexutil"
	"github.com/the-medium-tech/mdl-sdk-go/internal/crypto"
)

type FabricContract struct {
}

func newFabricContract() *FabricContract {
	return &FabricContract{}
}

func (f *FabricContract) Sign(hash string, keyFile string) (string, error) {
	return "", errors.New("fabric contract does not support sign function")
}

func (f *FabricContract) Hash(data string) (string, error) {
	return "", errors.New("fabric contract does not support hash function")
}

func (f *FabricContract) Address(a *address.Address) (string, error) {
	publicKeyBytes, err := hexutil.Decode(a.PublicKey)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(common.BytesToAddress(publicKeyBytes).Bytes()), nil
}

func (f *FabricContract) PublicKey(keyFile string) (string, error) {
	cert, err := crypto.LoadCertificate(keyFile)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(crypto.Keccak256(elliptic.Marshal(elliptic.P256(), cert.PublicKey.(*ecdsa.PublicKey).X, cert.PublicKey.(*ecdsa.PublicKey).Y)[12:])), nil
}

func (f *FabricContract) StringsToBytes(args []string) []byte {
	var temp string
	for _, arg := range args {
		temp += arg
	}
	return []byte(temp)
}
