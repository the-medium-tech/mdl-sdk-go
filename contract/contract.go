package contract

import (
	"encoding/json"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"github.com/the-medium-tech/mdl-sdk-go/internal/crypto"
)

type FabricContract struct {
	*contract
}

func NewFabricContract() *FabricContract {
	return &FabricContract{
		contract: NewContract(nil),
	}
}

func (f *FabricContract) SubmitTransaction(contract *gateway.Contract, function string, args ...string) ([]byte, error) {
	f.contract.setFunction(function)
	f.contract.setArgs(args)
	return f.contract.SubmitTransaction(contract)
}

type EthereumContract struct {
	*contract
}

func NewEthereumContract(config *config) *EthereumContract {
	return &EthereumContract{
		NewContract(config),
	}
}

func (e *EthereumContract) SubmitTransaction(contract *gateway.Contract, function string, args ...string) ([]byte, error) {
	e.contract.setFunction(function)
	e.contract.setArgs(args)
	err := e.contract.sign()
	if err != nil {
		return nil, err
	}
	return e.contract.SubmitTransaction(contract)
}

type BitcoinContract struct {
	*contract
}

func NewBitcoinContract(config *config) *BitcoinContract {
	return &BitcoinContract{
		NewContract(config),
	}
}

func (b *BitcoinContract) SubmitTransaction(contract *gateway.Contract, function string, args ...string) ([]byte, error) {
	b.contract.setFunction(function)
	b.contract.setArgs(args)
	err := b.contract.compress()
	if err != nil {
		return nil, err
	}
	return b.contract.SubmitTransaction(contract)
}

type contract struct {
	Function string
	Msg      *message
	Config   *config
}

type message struct {
	Args      []string `json:"args"`
	PublicKey []byte   `json:"publicKey,omitempty"`
	Hash      []byte   `json:"hash,omitempty"`
	Signature []byte   `json:"signature,omitempty"`
}

func (m *message) serialize() ([]byte, error) {
	return json.Marshal(&m)
}

func NewContract(config *config) *contract {
	return &contract{
		Config: config,
	}
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

func (c *contract) SubmitTransaction(contract *gateway.Contract) ([]byte, error) {
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

func (c *contract) compress() error {
	key, err := crypto.LoadECDSA(c.Config.path)
	if err != nil {
		return err
	}
	c.Msg.PublicKey = crypto.CompressPubkey(&key.PublicKey)
	return nil
}
