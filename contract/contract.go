package contract

import (
	"errors"
	"github.com/the-medium-tech/mdl-sdk-go/address"
)

type Contract interface {
	Sign(hash []byte, keyFile string) ([]byte, error)
	Hash(data []byte) []byte
	Verify(a *address.Address) bool
	ExtractAddress(a *address.Address) string
	ExtractPublickey(keyFile string) ([]byte, error)
	StringsToBytes(args []string) []byte
}

func NewContract(addressType string) (Contract, error) {
	switch addressType {
	case address.AddressTypeToString(address.FABRIC):
		return newFabricContract(), nil
	case address.AddressTypeToString(address.ETHEREUM):
		return newEthereumContract(), nil
	case address.AddressTypeToString(address.BITCOIN):
		return NewBitcoinContract(), nil
	default:
	}
	return nil, errors.New("not supported address type")
}

func GetContract(a *address.Address) (Contract, error) {
	return NewContract(a.Type)
}
