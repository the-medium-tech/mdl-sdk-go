package contract

import (
	"errors"
	"github.com/the-medium-tech/mdl-sdk-go/address"
	"github.com/the-medium-tech/mdl-sdk-go/internal/common/hexutil"
	"github.com/the-medium-tech/mdl-sdk-go/internal/crypto"
)

type BitcoinContract struct {
}

func NewBitcoinContract() *BitcoinContract {
	return &BitcoinContract{}
}

func (b *BitcoinContract) Sign(hash string, keyFile string) (string, error) {
	key, err := crypto.LoadECDSA(keyFile)
	if err != nil {
		return "", err
	}
	hashBytes, err := hexutil.Decode(hash)
	if err != nil {
		return "", err
	}
	signature, err := crypto.Sign(hashBytes, key)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(signature), nil
}

func (b *BitcoinContract) Hash(data string) (string, error) {
	return hexutil.Encode(crypto.DoubleHash([]byte(data))), nil
}

func (b *BitcoinContract) verify(a *address.Address) bool {
	signatureBytes, err := a.HexToBytes(a.Signature)
	if err != nil {
		return false
	}
	sig, err := crypto.ParseSignature(signatureBytes)
	if err != nil {
		return false
	}
	publicKeyBytes, err := a.HexToBytes(a.PublicKey)
	if err != nil {
		return false
	}
	publicKey, err := crypto.DecompressPubkey(publicKeyBytes)
	if err != nil {
		return false
	}
	hashBytes, err := a.HexToBytes(a.Hash)
	if err != nil {
		return false
	}
	return sig.Verify(hashBytes, publicKey)
}

func (b *BitcoinContract) Address(a *address.Address) (string, error) {
	if !b.verify(a) {
		return "", errors.New("failed to verify signature")
	}
	publicKeyBytes, err := a.HexToBytes(a.PublicKey)
	if err != nil {
		return "", err
	}
	payload := crypto.Hash160(publicKeyBytes)
	versionedPayload := append([]byte{0x00}, payload...)
	checksum := crypto.DoubleHash(payload)[:crypto.ChecksumLength]
	fullPayload := append(versionedPayload, checksum...)
	return crypto.Base58Encode(fullPayload), nil
}

func (b *BitcoinContract) PublicKey(keyFile string) (string, error) {
	key, err := crypto.LoadECDSA(keyFile)
	if err != nil {
		return "", err
	}
	return hexutil.Encode(crypto.CompressPubkey(&key.PublicKey)), nil
}

func (b *BitcoinContract) StringsToBytes(args []string) []byte {
	var temp string
	for _, arg := range args {
		temp += arg
	}
	return []byte(temp)
}
