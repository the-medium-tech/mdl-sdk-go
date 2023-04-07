package contract

import (
	"github.com/the-medium-tech/mdl-sdk-go/contract/header"
)

type Contract interface {
	GetArgs(file string, args ...string) ([]string, error)
}

func NewContractFactory(headerType string) Contract {
	switch headerType {
	case header.HeaderTypeToString(header.FABRIC):
		return NewFabricContract()
	case header.HeaderTypeToString(header.ETHEREUM):
		return NewEthereumContract()
	case header.HeaderTypeToString(header.BITCOIN):
		return NewBitcoinContract()
	default:
	}
	return nil
}

func GetContract(h *header.Header) Contract {
	return NewContractFactory(h.Type)
}
