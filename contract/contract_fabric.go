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

func (f *FabricContract) Sign(hash []byte, keyFile string) ([]byte, error) {
	return nil, errors.New("fabric contract does not support sign function")
}

func (f *FabricContract) Hash(data []byte) []byte {
	return nil
}

func (f *FabricContract) Verify(a *address.Address) bool {
	return true
}

func (f *FabricContract) ExtractAddress(a *address.Address) string {
	return hexutil.Encode(common.BytesToAddress(a.PublicKey).Bytes())
}

func (f *FabricContract) ExtractPublickey(keyFile string) ([]byte, error) {
	cert, err := crypto.LoadCertificate(keyFile)
	if err != nil {
		return nil, err
	}
	return crypto.Keccak256(elliptic.Marshal(elliptic.P256(), cert.PublicKey.(*ecdsa.PublicKey).X, cert.PublicKey.(*ecdsa.PublicKey).Y)[12:]), nil
}

func (f *FabricContract) StringsToBytes(args []string) []byte {
	var temp string
	for _, arg := range args {
		temp += arg
	}
	return []byte(temp)
}
