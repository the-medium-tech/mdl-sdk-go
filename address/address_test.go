package address

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/stretchr/testify/assert"
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
