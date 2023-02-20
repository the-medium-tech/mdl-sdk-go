package address

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"encoding/pem"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/protobuf/proto"

	"github.com/hyperledger/fabric-protos-go/msp"
)

func GetAddress() string {

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
