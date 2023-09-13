package bridge

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/ripemd160"
)

func MultiSignToAddress2(encodeMsg string, bytesMultiSign string) ([]string, error) {
	msgHash := sha256.Sum256([]byte(encodeMsg))
	return MultiSignToAddress(msgHash, bytesMultiSign)
}

func MultiSignToAddress(msgHash [32]byte, bytesMultiSign string) ([]string, error) {
	hexSignList, err := DecodeMultiSign(bytesMultiSign)
	if err != nil {
		return nil, err
	}

	var address []string
	for _, hexSign := range hexSignList {
		addr, err := SignToAddress(msgHash, hexSign)
		if err != nil {
			return nil, err
		}
		address = append(address, addr)
	}

	return address, nil
}

func SignToAddress(mshHash [32]byte, hexSign string) (string, error) {
	pubKey, err := SignToPub(mshHash, hexSign)
	if err != nil {
		return "", err
	}
	rigoAddress := PubToAddress(pubKey)
	return rigoAddress, nil
}

func SignToPub(msgHash [32]byte, hexSign string) ([]byte, error) {
	signByte, err := decodeSign(hexSign)
	if err != nil {
		return nil, err
	}

	publicKeySrc, err := crypto.SigToPub(msgHash[:], signByte)
	if err != nil {
		return nil, err
	}

	pubKeyCompress := crypto.CompressPubkey(publicKeySrc)
	//hexPublicKey := hex.EncodeToString(pubKeyCompress)
	return pubKeyCompress, nil
}

func PubToAddress(publicKey []byte) string {
	pubKeyHash := sha256.Sum256(publicKey)

	rip := ripemd160.New()
	rip.Write(pubKeyHash[:])
	ripHash := rip.Sum(nil)
	rigoAddress := hex.EncodeToString(ripHash)

	return "0x" + rigoAddress
}

func SenderAddress(txid string, senderSign string) (string, error) {
	txidHash := sha256.Sum256([]byte(txid))
	senderAddr, err := SignToAddress(txidHash, senderSign)
	if err != nil {
		return "", err
	}

	return senderAddr, nil
}
