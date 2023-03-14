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
	f.setFunction(function)
	f.setArgs(args)
	return f.contract.submitTransaction(contract)
}

func (f *FabricContract) Sign() error {
	return nil
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
	e.setFunction(function)
	e.setArgs(args)
	if err := e.Sign(); err != nil {
		return nil, err
	}
	return e.contract.submitTransaction(contract)
}

func (e *EthereumContract) Sign() error {
	msg, err := e.Msg.serialize()
	if err != nil {
		return err
	}
	e.Msg.Hash = crypto.Keccak256(msg)
	key, err := crypto.LoadECDSA(e.Config.path)
	if err != nil {
		return err
	}
	e.Msg.Signature, err = crypto.SignCompact(e.Msg.Hash, key)
	if err != nil {
		return err
	}
	return err
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
	b.setFunction(function)
	b.setArgs(args)
	if err := b.Compress(); err != nil {
		return nil, err
	}
	return b.contract.submitTransaction(contract)
}

func (b *BitcoinContract) Sign() error {
	msg, err := b.Msg.serialize()
	if err != nil {
		return err
	}
	b.Msg.Hash = crypto.DoubleHash(msg)
	key, err := crypto.LoadECDSA(b.Config.path)
	if err != nil {
		return err
	}
	b.Msg.Signature, err = crypto.Sign(b.Msg.Hash, key)
	if err != nil {
		return err
	}
	return err
}

func (b *BitcoinContract) Compress() error {
	key, err := crypto.LoadECDSA(b.Config.path)
	if err != nil {
		return err
	}
	b.Msg.PublicKey = crypto.CompressPubkey(&key.PublicKey)
	return nil
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

func (c *contract) submitTransaction(contract *gateway.Contract) ([]byte, error) {
	return contract.SubmitTransaction(c.Function, c.string())
}
