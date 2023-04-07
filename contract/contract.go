package contract

import (
	"crypto/ecdsa"
	"crypto/elliptic"

	"github.com/the-medium-tech/mdl-sdk-go/contract/header"
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

func (f *FabricContract) GetArgs(file string, args ...string) ([]string, error) {
	h, err := f.makeHeader(file)
	return f.getArgs(h, args...), err
}

func (f *FabricContract) name() string {
	return header.HeaderTypeToString(header.FABRIC)
}

func (f *FabricContract) makeHeader(file string) (*header.Header, error) {
	var err error
	f.setConfig(LoadConfig(file))
	header := header.NewHeader(f.name())
	header.PublicKey, err = f.setPublicKeyFromCertificate()
	return header, err
}

func (f *FabricContract) setPublicKeyFromCertificate() ([]byte, error) {
	cert, err := crypto.LoadCertificate(f.Config.path)
	if err != nil {
		return nil, err
	}
	return crypto.Keccak256(elliptic.Marshal(elliptic.P256(), cert.PublicKey.(*ecdsa.PublicKey).X, cert.PublicKey.(*ecdsa.PublicKey).Y)[12:]), nil
}

type EthereumContract struct {
	*contract
}

func NewEthereumContract() *EthereumContract {
	return &EthereumContract{
		newContract(),
	}
}

func (e *EthereumContract) GetArgs(file string, args ...string) ([]string, error) {
	h, err := e.makeHeader(file)
	return e.getArgs(h, args...), err
}

func (e *EthereumContract) name() string {
	return header.HeaderTypeToString(header.ETHEREUM)
}

func (e *EthereumContract) makeHeader(file string) (*header.Header, error) {
	e.setConfig(LoadConfig(file))
	h := header.NewHeader(e.name())
	err := e.sign(h)
	return h, err
}

func (e *EthereumContract) sign(h *header.Header) error {
	m, err := h.Serialize()
	if err != nil {
		return err
	}
	h.Hash = crypto.Keccak256(m)
	key, err := crypto.LoadECDSA(e.Config.path)
	if err != nil {
		return err
	}
	h.Signature, err = crypto.SignCompact(h.Hash, key)
	if err != nil {
		return err
	}
	return err
}

type BitcoinContract struct {
	*contract
}

func NewBitcoinContract() *BitcoinContract {
	return &BitcoinContract{
		newContract(),
	}
}

func (b *BitcoinContract) GetArgs(file string, args ...string) ([]string, error) {
	h, err := b.makeHeader(file)
	return b.getArgs(h, args...), err
}

func (b *BitcoinContract) name() string {
	return header.HeaderTypeToString(header.BITCOIN)
}

func (b *BitcoinContract) makeHeader(file string) (*header.Header, error) {
	b.setConfig(LoadConfig(file))
	h := header.NewHeader(b.name())
	if err := b.compress(h); err != nil {
		return h, err
	}
	if err := b.sign(h); err != nil {
		return h, err
	}
	return h, nil
}

func (b *BitcoinContract) sign(h *header.Header) error {
	m, err := h.Serialize()
	if err != nil {
		return err
	}
	h.Hash = crypto.DoubleHash(m)
	key, err := crypto.LoadECDSA(b.Config.path)
	if err != nil {
		return err
	}
	h.Signature, err = crypto.Sign(h.Hash, key)
	return err
}

func (b *BitcoinContract) compress(h *header.Header) error {
	key, err := crypto.LoadECDSA(b.Config.path)
	if err != nil {
		return err
	}
	h.PublicKey = crypto.CompressPubkey(&key.PublicKey)
	return err
}

type contract struct {
	Args   []string
	Config *config
}

func newContract() *contract {
	return &contract{}
}

func (c *contract) getArgs(h *header.Header, a ...string) []string {
	var args []string
	header, err := h.Serialize()
	if err != nil {
		return nil
	}
	args = append(args, string(header))
	args = append(args, a...)
	return args
}

func (c *contract) setConfig(config *config) {
	c.Config = config
}
