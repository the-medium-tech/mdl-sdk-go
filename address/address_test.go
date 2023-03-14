package address

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/stretchr/testify/assert"
	"github.com/the-medium-tech/mdl-sdk-go/internal/crypto"
)

func TestGetAddressWithSerializedIdentity(t *testing.T) {
	certBytes, err := ioutil.ReadFile(filepath.Join("testdata", "cert.pem"))
	assert.NoError(t, err)
	sId := &msp.SerializedIdentity{
		IdBytes: certBytes,
	}
	serializedIdentity, err := proto.Marshal(sId)
	assert.Equal(t, "0x4057cc4274523666fa4cc88e5f78193b36105a33", GetAddressWithSerializedIdentity(serializedIdentity))
}

func TestGetAddressWithCert(t *testing.T) {
	certBytes, err := ioutil.ReadFile(filepath.Join("testdata", "cert.pem"))
	assert.NoError(t, err)
	block, _ := pem.Decode(certBytes)
	cert, err := x509.ParseCertificate(block.Bytes)
	assert.NoError(t, err)
	assert.Equal(t, "0x4057cc4274523666fa4cc88e5f78193b36105a33", GetAddressWithCert(cert))
}

func TestGetAddressWithPublicKey(t *testing.T) {
	file := filepath.Join("testdata", "test.key")
	key, err := crypto.LoadECDSA(file)
	assert.NoError(t, err)
	assert.Equal(t, "15VDTyzYK6SiH4kCdT89bEaskB15QS79F9", GetAddressWithCompressedPublicKey(crypto.CompressPubkey(&key.PublicKey)))
}

func TestGetAddressWithSignature(t *testing.T) {
	file := filepath.Join("testdata", "test.key")
	key, err := crypto.LoadECDSA(file)
	assert.NoError(t, err)
	msg := []byte("foo")
	hash := crypto.Keccak256(msg)
	sig, err := crypto.SignCompact(hash, key)
	assert.NoError(t, err)
	assert.Equal(t, "0x93b2Cb3061e36Ed3099d003fF78cd685b424e95b", GetAddressWithSignature(hash, sig))
}
