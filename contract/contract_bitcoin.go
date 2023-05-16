package contract

import (
	"github.com/the-medium-tech/mdl-sdk-go/address"
	"github.com/the-medium-tech/mdl-sdk-go/internal/crypto"
)

type BitcoinContract struct {
}

func NewBitcoinContract() *BitcoinContract {
	return &BitcoinContract{}
}

func (b *BitcoinContract) Sign(hash []byte, keyFile string) ([]byte, error) {
	key, err := crypto.LoadECDSA(keyFile)
	if err != nil {
		return nil, err
	}
	return crypto.Sign(hash, key)
}

func (b *BitcoinContract) Hash(data []byte) []byte {
	return crypto.DoubleHash(data)
}

func (b *BitcoinContract) Verify(a *address.Address) bool {
	sig, err := crypto.ParseSignature(a.Signature)
	if err != nil {
		return false
	}
	pub, err := crypto.DecompressPubkey(a.PublicKey)
	if err != nil {
		return false
	}
	return sig.Verify(a.Hash, pub)
}

func (b *BitcoinContract) ExtractAddress(a *address.Address) string {
	payload := crypto.Hash160(a.PublicKey)
	versionedPayload := append([]byte{0x00}, payload...)
	checksum := crypto.DoubleHash(payload)[:crypto.ChecksumLength]
	fullPayload := append(versionedPayload, checksum...)
	return crypto.Base58Encode(fullPayload)
}

func (b *BitcoinContract) ExtractPublickey(keyFile string) ([]byte, error) {
	key, err := crypto.LoadECDSA(keyFile)
	if err != nil {
		return nil, err
	}
	return crypto.CompressPubkey(&key.PublicKey), nil
}

func (b *BitcoinContract) StringsToBytes(args []string) []byte {
	var temp string
	for _, arg := range args {
		temp += arg
	}
	return []byte(temp)
}
