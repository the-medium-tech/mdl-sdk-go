package contract

import (
	"encoding/hex"
	"encoding/json"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"github.com/the-medium-tech/mdl-sdk-go/internal/crypto"
	"os"
)

type contract struct {
	Function string
	Msg      *message
	Config   *config
}

type message struct {
	Args      []string `json:"args"`
	Hash      []byte   `json:"hash,omitempty"`
	Signature []byte   `json:"signature,omitempty"`
}

func NewContract(config *config) *contract {
	return &contract{
		Config: config,
	}
}

func (m *message) serialize() ([]byte, error) {
	return json.Marshal(&m)
}

func (c *contract) string() string {
	msg, err := c.Msg.serialize()
	if err != nil {
		return ""
	}
	return string(msg)
}

func (c *contract) setFunction(function string) {
	c.Function = function
}

func (c *contract) setArgs(args []string) {
	c.Msg = &message{
		Args: args,
	}
}

func (c *contract) SubmitTransaction(contract *gateway.Contract, function string, args ...string) ([]byte, error) {
	c.setFunction(function)
	c.setArgs(args)
	err := c.sign()
	if err != nil {
		return nil, err
	}

	return contract.SubmitTransaction(c.Function, c.string())
}

func (c *contract) sign() error {
	msg, err := c.Msg.serialize()
	if err != nil {
		return err
	}
	c.Msg.Hash = crypto.Keccak256(msg)
	key, err := crypto.LoadECDSA(c.Config.path)
	if err != nil {
		return err
	}
	c.Msg.Signature, err = crypto.Sign(c.Msg.Hash, key)
	if err != nil {
		return err
	}
	return err
}

func GenerateKey(file string) error {
	key, err := crypto.GenerateKey()
	if err != nil {
		return err
	}
	return os.WriteFile(file, []byte(hex.EncodeToString(crypto.FromECDSA(key))), 0600)
}
