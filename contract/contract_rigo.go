package contract

import (
	"errors"
	"github.com/the-medium-tech/mdl-sdk-go/address"
	"github.com/the-medium-tech/mdl-sdk-go/internal/common/hexutil"
	"github.com/the-medium-tech/mdl-sdk-go/internal/crypto"
)

type RigoContract struct {
}

func NewRigoContract() *RigoContract {
	return &RigoContract{}
}

func (r *RigoContract) Sign(hash string, keyFile string) (string, error) {
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
	return hexutil.Encode(signature[:64]), nil
}

func (r *RigoContract) Hash(data string) (string, error) {
	return hexutil.Encode(crypto.Sha256([]byte(data))), nil
}

func (r *RigoContract) verify(a *address.Address) bool {
	signatureBytes, err := a.HexToBytes(a.Signature)
	if err != nil {
		return false
	}
	publicKeyBytes, err := a.HexToBytes(a.PublicKey)
	if err != nil {
		return false
	}
	hashBytes, err := a.HexToBytes(a.Hash)
	if err != nil {
		return false
	}
	return crypto.VerifySignature(publicKeyBytes, hashBytes, signatureBytes)
}

func (r *RigoContract) Address(a *address.Address) (string, error) {
	if !r.verify(a) {
		return "", errors.New("failed to verify signature")
	}
	publicKeyBytes, err := a.HexToBytes(a.PublicKey)
	if err != nil {
		return "", err
	}
	payload := crypto.Hash160(publicKeyBytes)
	return hexutil.ToHex(payload), nil
}

func (r *RigoContract) PublicKey(keyFile string) (string, error) {
	key, err := crypto.LoadECDSA(keyFile)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(crypto.CompressPubkey(&key.PublicKey)), nil
}

func (r *RigoContract) StringsToBytes(args []string) []byte {
	var temp string
	for _, arg := range args {
		temp += arg
	}
	return []byte(temp)
}
