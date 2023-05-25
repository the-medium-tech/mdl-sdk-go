package address

import (
	"encoding/json"
	"errors"

	"github.com/the-medium-tech/mdl-sdk-go/internal/common/hexutil"
)

type AddressType int

const (
	FABRIC AddressType = iota
	ETHEREUM
	BITCOIN
	NotSupoorted
)

var addressTypeStrings = map[AddressType]string{
	FABRIC:   "fabric",
	ETHEREUM: "ethereum",
	BITCOIN:  "bitcoin",
}

func AddressTypeToString(id AddressType) string {
	if res, found := addressTypeStrings[id]; found {
		return res
	}
	return ""
}

type Address struct {
	Type      string `json:"type"`
	PublicKey []byte `json:"publicKey,omitempty"`
	Hash      []byte `json:"hash,omitempty"`
	Signature []byte `json:"signature,omitempty"`
}

func (a *Address) Serialize() ([]byte, error) {
	return json.Marshal(&a)
}

func (a *Address) HashToHex() string {
	return hexutil.Encode(a.Hash)
}

func (a *Address) AppendArgs(args []string) ([]string, error) {
	var result []string
	serializedAddress, err := a.Serialize()
	if err != nil {
		return nil, err
	}
	result = append(result, string(serializedAddress))
	result = append(result, args...)
	return result, nil
}

func Deserialize(serializedAddress []byte) (*Address, error) {
	a := &Address{}
	err := json.Unmarshal(serializedAddress, &a)
	return a, err
}

func NewAddress(t AddressType, publicKey, hash, signature []byte) (*Address, error) {
	if t >= NotSupoorted {
		return nil, errors.New("not supported address type")
	}

	if t == FABRIC && publicKey == nil {
		return nil, errors.New("fabric address must have public key")
	} else if t == ETHEREUM && (hash == nil || signature == nil) {
		return nil, errors.New("ethereum address must have hash, signature")
	} else if t == BITCOIN && (publicKey == nil || hash == nil || signature == nil) {
		return nil, errors.New("bitcoin address must have public key, hash, signature")
	}

	return &Address{
		Type:      AddressTypeToString(t),
		PublicKey: publicKey,
		Hash:      hash,
		Signature: signature,
	}, nil
}
