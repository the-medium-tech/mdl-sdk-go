package contract

import (
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/stretchr/testify/assert"
	"github.com/the-medium-tech/mdl-sdk-go/address"
	"github.com/the-medium-tech/mdl-sdk-go/internal/crypto"
)

func TestUserScenarioForFabricContractWithCore(t *testing.T) {

	fab := NewFabricContract()
	fab.Function = "function"
	fab.Msg = &message{
		Args: []string{"1,2,3,4"},
	}
	
	certBytes, err := ioutil.ReadFile(filepath.Join("testdata", "cert.pem"))
	assert.NoError(t, err)
	sId := &msp.SerializedIdentity{
		IdBytes: certBytes,
	}
	serializedIdentity, err := proto.Marshal(sId)
	assert.Equal(t, "0x4057cc4274523666fa4cc88e5f78193b36105a33", address.GetAddressWithSerializedIdentity(serializedIdentity))
}

func TestUserScenarioForFabricContractWithChaincode(t *testing.T) {

	fab := NewFabricContract()
	fab.Function = "function"
	fab.Msg = &message{
		Args: []string{"1,2,3,4"},
	}

	certBytes, err := ioutil.ReadFile(filepath.Join("testdata", "cert.pem"))
	assert.NoError(t, err)
	block, _ := pem.Decode(certBytes)
	cert, err := x509.ParseCertificate(block.Bytes)
	assert.NoError(t, err)
	assert.Equal(t, "0x4057cc4274523666fa4cc88e5f78193b36105a33", address.GetAddressWithCert(cert))
}

func TestUserScenarioForEthereumContract(t *testing.T) {
	file := filepath.Join("testdata", "test.key")
	err := generateKey(file)
	assert.NoError(t, err)

	eth := NewEthereumContract(LoadConfig(file))
	eth.Function = "function"
	eth.Msg = &message{
		Args: []string{"1,2,3,4"},
	}
	err = eth.Sign()
	assert.NoError(t, err)
	assert.Equal(t, "0x93b2Cb3061e36Ed3099d003fF78cd685b424e95b", address.GetAddressWithSignature(eth.Msg.Hash, eth.Msg.Signature))
}

func TestUserScenarioForBitcoinContract(t *testing.T) {
	file := filepath.Join("testdata", "test.key")
	err := generateKey(file)
	assert.NoError(t, err)

	btc := NewBitcoinContract(LoadConfig(file))
	btc.Function = "function"
	btc.Msg = &message{
		Args: []string{"1,2,3,4"},
	}
	err = btc.Sign()
	assert.NoError(t, err)
	err = btc.Compress()
	assert.NoError(t, err)

	sig, err := crypto.ParseSignature(btc.Msg.Signature)
	assert.NoError(t, err)
	pubkey, err := crypto.DecompressPubkey(btc.Msg.PublicKey)
	assert.NoError(t, err)
	assert.True(t, sig.Verify(btc.Msg.Hash, pubkey))
}

func generateKey(file string) error {
	if _, err := os.Stat(file); err == os.ErrNotExist {
		key, err := crypto.GenerateKey()
		if err != nil {
			return err
		}
		return os.WriteFile(file, []byte(hex.EncodeToString(crypto.FromECDSA(key))), 0600)
	}
	return nil
}
