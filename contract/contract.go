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

func (f *FabricContract) SetHeader(header *header) {
	f.Header = header
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
	f.setArgs(args)
	header := newHeader()
	header.setType(f.Name())
	header.PublicKey, err = f.setPublicKeyFromCertificate()
	if err != nil {
		return err
	}
	f.SetHeader(header)
	return nil
}

func (f *FabricContract) Verify() bool {
	return true
}

func (f *FabricContract) Address() string {
	return hexutil.Encode(common.BytesToAddress(f.Header.PublicKey).Bytes())
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

func (e *EthereumContract) SetHeader(header *header) {
	e.Header = header
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
	e.setArgs(args)
	header := newHeader()
	header.setType(e.Name())
	if err := e.sign(header); err != nil {
		return err
	}
	e.SetHeader(header)
	return nil
}

func (e *EthereumContract) sign(header *header) error {
	m, err := header.serialize()
	if err != nil {
		return err
	}
	header.Hash = crypto.Keccak256(m)
	key, err := crypto.LoadECDSA(e.Config.path)
	if err != nil {
		return err
	}
	header.Signature, err = crypto.SignCompact(header.Hash, key)
	if err != nil {
		return err
	}
	return err
}

func (e *EthereumContract) Verify() bool {
	return true
}

func (e *EthereumContract) Address() string {
	recoveredPub, err := crypto.Ecrecover(e.Header.Hash, e.Header.Signature)
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

func (b *BitcoinContract) SetHeader(header *header) {
	b.Header = header
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
	b.setArgs(args)
	header := newHeader()
	header.setType(b.Name())
	if err := b.compress(header); err != nil {
		return err
	}
	if err := b.sign(header); err != nil {
		return err
	}
	b.SetHeader(header)
	return nil
}

func (b *BitcoinContract) Verify() bool {
	sig, err := crypto.ParseSignature(b.Header.Signature)
	if err != nil {
		return false
	}
	pubkey, err := crypto.DecompressPubkey(b.Header.PublicKey)
	if err != nil {
		return false
	}
	return sig.Verify(b.Header.Hash, pubkey)
}

func (b *BitcoinContract) Address() string {
	payload := crypto.Hash160(b.Header.PublicKey)
	versionedPayload := append([]byte{0x00}, payload...)
	checksum := crypto.DoubleHash(payload)[:crypto.ChecksumLength]
	fullPayload := append(versionedPayload, checksum...)
	return crypto.Base58Encode(fullPayload)
}

func (b *BitcoinContract) sign(header *header) error {
	m, err := header.serialize()
	if err != nil {
		return err
	}
	header.Hash = crypto.DoubleHash(m)
	key, err := crypto.LoadECDSA(b.Config.path)
	if err != nil {
		return err
	}
	header.Signature, err = crypto.Sign(header.Hash, key)
	return err
}

func (b *BitcoinContract) compress(header *header) error {
	key, err := crypto.LoadECDSA(b.Config.path)
	if err != nil {
		return err
	}
	header.PublicKey = crypto.CompressPubkey(&key.PublicKey)
	return err
}

func (b *BitcoinContract) Name() string {
	return TransactionTypeToString(BITCOIN)
}

type contract struct {
	Function string
	Args     []string
	Header   *header
	Config   *config
}

type header struct {
	Type      string `json:"type"`
	PublicKey []byte `json:"publicKey,omitempty"`
	Hash      []byte `json:"hash,omitempty"`
	Signature []byte `json:"signature,omitempty"`
}

func newHeader() *header {
	return &header{}
}

func (h *header) serialize() ([]byte, error) {
	return json.Marshal(&h)
}

func (h *header) deserialize(payload []byte) error {
	return json.Unmarshal(payload, &h)
}

func (h *header) setType(t string) {
	h.Type = t
}

func newContract() *contract {
	return &contract{}
}

func (c *contract) getArgs() []string {
	var args []string
	header, err := c.Header.serialize()
	if err != nil {
		return nil
	}
	args = append(args, string(header))
	args = append(args, c.Args...)
	return args
}

func (c *contract) setFunction(function string) {
	c.Function = function
}

func (c *contract) setArgs(args []string) {
	c.Args = args
}

func (c *contract) setConfig(config *config) {
	c.Config = config
}

func (c *contract) submitTransaction(contract *gateway.Contract) ([]byte, error) {
	return contract.SubmitTransaction(c.Function, c.getArgs()...)
}
