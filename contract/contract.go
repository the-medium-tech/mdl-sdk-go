package contract

import (
	"errors"
	"github.com/the-medium-tech/mdl-sdk-go/address"
)

type Contract interface {
	Sign(hash string, keyFile string) (string, error)
	Hash(data string) (string, error)
	Address(a *address.Address) (string, error)
	PublicKey(keyFile string) (string, error)
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
