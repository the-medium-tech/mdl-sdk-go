package contract

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"encoding/pem"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/the-medium-tech/mdl-sdk-go/internal/common"
	"github.com/the-medium-tech/mdl-sdk-go/internal/common/hexutil"
	"github.com/the-medium-tech/mdl-sdk-go/internal/crypto"
)

func GetAddress(transactionType string, args ...[]byte) string {
	switch transactionType {
	case TransactionTypeToString(FABRIC):
		return GetAddressWithSerializedIdentity(args[0])
	case TransactionTypeToString(ETHEREUM):
		return GetAddressWithSignature(args[0], args[1])
	case TransactionTypeToString(BITCOIN):
		return GetAddressWithCompressedPublicKey(args[0])
	default:
	}
	return ""
}

func GetAddressWithSerializedIdentity(serializedIdentity []byte) string {
	sId := &msp.SerializedIdentity{}
	err := proto.Unmarshal(serializedIdentity, sId)
	if err != nil {
		return ""
	}
	bl, _ := pem.Decode(sId.IdBytes)
	if bl == nil {
		return ""
	}
	cert, err := x509.ParseCertificate(bl.Bytes)
	if err != nil {
		return ""
	}

	return hexutil.Encode(common.BytesToAddress(crypto.Keccak256(elliptic.Marshal(elliptic.P256(), cert.PublicKey.(*ecdsa.PublicKey).X, cert.PublicKey.(*ecdsa.PublicKey).Y)[12:])).Bytes())
}

func GetAddressWithSignature(hash, sig []byte) string {
	recoveredPub, err := crypto.Ecrecover(hash, sig)
	if err != nil {
		return ""
	}
	pubKey, err := crypto.UnmarshalPubkey(recoveredPub)
	if err != nil {
		return ""
	}

	return crypto.PubkeyToAddress(*pubKey).String()
}

func GetAddressWithCompressedPublicKey(compressedPubKey []byte) string {
	payload := crypto.Hash160(compressedPubKey)
	versionedPayload := append([]byte{0x00}, payload...)
	checksum := crypto.DoubleHash(payload)
	checksum = checksum[:crypto.ChecksumLength]
	fullPayload := append(versionedPayload, checksum...)
	return crypto.Base58Encode(fullPayload)
}