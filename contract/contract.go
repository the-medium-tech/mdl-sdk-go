package contract

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/json"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	"github.com/the-medium-tech/mdl-sdk-go/internal/common"
	"github.com/the-medium-tech/mdl-sdk-go/internal/common/hexutil"
	"github.com/the-medium-tech/mdl-sdk-go/internal/crypto"
)

type FabricContract struct {
	*contract
}

func NewFabricContract() *FabricContract {
	return &FabricContract{
		contract: newContract(),
	}
}

func (f *FabricContract) SetMessage(msg *message) {
	f.Msg = msg
}

func (f *FabricContract) SubmitTransaction(contract *gateway.Contract, file, function string, args ...string) ([]byte, error) {
	if err := f.makeTransaction(file, function, args...); err != nil {
		return nil, err
	}
	return f.submitTransaction(contract)
}

func (f *FabricContract) makeTransaction(file, function string, args ...string) error {
	var err error
	f.setConfig(LoadConfig(file))
	f.setFunction(function)
	msg := newMessage()
	msg.setArgs(args)
	msg.setType(f.Name())
	msg.PublicKey, err = f.setPublicKeyFromCertificate()
	if err != nil {
		return err
	}
	f.SetMessage(msg)
	return nil
}

func (f *FabricContract) Verify() bool {
	return true
}

func (f *FabricContract) Address() string {
	return hexutil.Encode(common.BytesToAddress(f.Msg.PublicKey).Bytes())
}

func (f *FabricContract) setPublicKeyFromCertificate() ([]byte, error) {
	cert, err := crypto.LoadCertificate(f.Config.path)
	if err != nil {
		return nil, err
	}
	return crypto.Keccak256(elliptic.Marshal(elliptic.P256(), cert.PublicKey.(*ecdsa.PublicKey).X, cert.PublicKey.(*ecdsa.PublicKey).Y)[12:]), nil
}

func (f *FabricContract) Name() string {
	return TransactionTypeToString(FABRIC)
}

type EthereumContract struct {
	*contract
}

func NewEthereumContract() *EthereumContract {
	return &EthereumContract{
		newContract(),
	}
}

func (e *EthereumContract) SetMessage(msg *message) {
	e.Msg = msg
}

func (e *EthereumContract) SubmitTransaction(contract *gateway.Contract, file, function string, args ...string) ([]byte, error) {
	if err := e.makeTransaction(file, function, args...); err != nil {
		return nil, err
	}
	return e.submitTransaction(contract)
}

func (e *EthereumContract) makeTransaction(file, function string, args ...string) error {
	e.setConfig(LoadConfig(file))
	e.setFunction(function)
	msg := newMessage()
	msg.setArgs(args)
	msg.setType(e.Name())
	if err := e.sign(msg); err != nil {
		return err
	}
	e.SetMessage(msg)
	return nil
}

func (e *EthereumContract) sign(msg *message) error {
	m, err := msg.serialize()
	if err != nil {
		return err
	}
	msg.Hash = crypto.Keccak256(m)
	key, err := crypto.LoadECDSA(e.Config.path)
	if err != nil {
		return err
	}
	msg.Signature, err = crypto.SignCompact(msg.Hash, key)
	if err != nil {
		return err
	}
	return err
}

func (e *EthereumContract) Verify() bool {
	return true
}

func (e *EthereumContract) Address() string {
	recoveredPub, err := crypto.Ecrecover(e.Msg.Hash, e.Msg.Signature)
	if err != nil {
		return ""
	}
	pubKey, err := crypto.UnmarshalPubkey(recoveredPub)
	if err != nil {
		return ""
	}

	return crypto.PubkeyToAddress(*pubKey).String()
}

func (e *EthereumContract) Name() string {
	return TransactionTypeToString(ETHEREUM)
}

type BitcoinContract struct {
	*contract
}

func NewBitcoinContract() *BitcoinContract {
	return &BitcoinContract{
		newContract(),
	}
}

func (b *BitcoinContract) SetMessage(msg *message) {
	b.Msg = msg
}

func (b *BitcoinContract) SubmitTransaction(contract *gateway.Contract, file, function string, args ...string) ([]byte, error) {
	if err := b.makeTransaction(file, function, args...); err != nil {
		return nil, err
	}
	return b.contract.submitTransaction(contract)
}

func (b *BitcoinContract) makeTransaction(file, function string, args ...string) error {
	b.setConfig(LoadConfig(file))
	b.setFunction(function)
	msg := newMessage()
	msg.setArgs(args)
	msg.setType(b.Name())
	if err := b.compress(msg); err != nil {
		return err
	}
	if err := b.sign(msg); err != nil {
		return err
	}
	b.SetMessage(msg)
	return nil
}

func (b *BitcoinContract) Verify() bool {
	sig, err := crypto.ParseSignature(b.Msg.Signature)
	if err != nil {
		return false
	}
	pubkey, err := crypto.DecompressPubkey(b.Msg.PublicKey)
	if err != nil {
		return false
	}
	return sig.Verify(b.Msg.Hash, pubkey)
}

func (b *BitcoinContract) Address() string {
	payload := crypto.Hash160(b.Msg.PublicKey)
	versionedPayload := append([]byte{0x00}, payload...)
	checksum := crypto.DoubleHash(payload)[:crypto.ChecksumLength]
	fullPayload := append(versionedPayload, checksum...)
	return crypto.Base58Encode(fullPayload)
}

func (b *BitcoinContract) sign(msg *message) error {
	m, err := msg.serialize()
	if err != nil {
		return err
	}
	msg.Hash = crypto.DoubleHash(m)
	key, err := crypto.LoadECDSA(b.Config.path)
	if err != nil {
		return err
	}
	msg.Signature, err = crypto.Sign(msg.Hash, key)
	return err
}

func (b *BitcoinContract) compress(msg *message) error {
	key, err := crypto.LoadECDSA(b.Config.path)
	if err != nil {
		return err
	}
	msg.PublicKey = crypto.CompressPubkey(&key.PublicKey)
	return err
}

func (b *BitcoinContract) Name() string {
	return TransactionTypeToString(BITCOIN)
}

type contract struct {
	Function string
	Msg      *message
	Config   *config
}

type message struct {
	Type      string   `json:"type"`
	Args      []string `json:"args"`
	PublicKey []byte   `json:"publicKey,omitempty"`
	Hash      []byte   `json:"hash,omitempty"`
	Signature []byte   `json:"signature,omitempty"`
}

func newMessage() *message {
	return &message{}
}

func (m *message) serialize() ([]byte, error) {
	return json.Marshal(&m)
}

func (m *message) Deserialize(payload []byte) error {
	return json.Unmarshal(payload, &m)
}

func (m *message) setArgs(args []string) {
	m.Args = args
}

func (m *message) setType(t string) {
	m.Type = t
}

func newContract() *contract {
	return &contract{}
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

func (c *contract) setConfig(config *config) {
	c.Config = config
}

func (c *contract) submitTransaction(contract *gateway.Contract) ([]byte, error) {
	return contract.SubmitTransaction(c.Function, c.string())
}
