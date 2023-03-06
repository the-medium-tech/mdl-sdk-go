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
)

var (
	expectedAddress = "0x4057cc4274523666fa4cc88e5f78193b36105a33"
)

func TestGetAddressWithSerializedIdentity(t *testing.T) {
	certBytes, err := ioutil.ReadFile(filepath.Join("testdata", "cert.pem"))
	assert.NoError(t, err)
	sId := &msp.SerializedIdentity{
		IdBytes: certBytes,
	}
	serializedIdentity, err := proto.Marshal(sId)
	assert.Equal(t, expectedAddress, GetAddressWithSerializedIdentity(serializedIdentity))
}

func TestGetAddressWithCert(t *testing.T) {
	certBytes, err := ioutil.ReadFile(filepath.Join("testdata", "cert.pem"))
	assert.NoError(t, err)
	block, _ := pem.Decode(certBytes)
	cert, err := x509.ParseCertificate(block.Bytes)
	assert.NoError(t, err)
	assert.Equal(t, expectedAddress, GetAddressWithCert(cert))
}
