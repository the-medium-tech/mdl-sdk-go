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

type Address interface {
	GetAddress(material []byte) string
}

type NormalImpl struct{}

func NewAddressNormal() *NormalImpl {
	return &NormalImpl{}
}

func (n *NormalImpl) GetAddress(material []byte) string {
	sId := &msp.SerializedIdentity{}
	err := proto.Unmarshal(material, sId)
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

type EthImpl struct{}

func NewAddressEth() *EthImpl {
	return &EthImpl{}
}

func (e *EthImpl) GetAddress(material []byte) string {

	// crypto.PubkeyToAddress()
	return ""
}
