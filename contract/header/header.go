package header

import (
	"encoding/json"

	"github.com/the-medium-tech/mdl-sdk-go/internal/common"
	"github.com/the-medium-tech/mdl-sdk-go/internal/common/hexutil"
	"github.com/the-medium-tech/mdl-sdk-go/internal/crypto"
)

type HeaderType int

const (
	FABRIC HeaderType = iota
	ETHEREUM
	BITCOIN
)

var headerTypeStrings = map[HeaderType]string{
	FABRIC:   "fabric",
	ETHEREUM: "ethereum",
	BITCOIN:  "bitcoin",
}

func HeaderTypeToString(id HeaderType) string {
	if res, found := headerTypeStrings[id]; found {
		return res
	}
	return ""
}

type Header struct {
	Type      string `json:"type"`
	PublicKey []byte `json:"publicKey,omitempty"`
	Hash      []byte `json:"hash,omitempty"`
	Signature []byte `json:"signature,omitempty"`
}

func (h *Header) Serialize() ([]byte, error) {
	return json.Marshal(&h)
}

func NewHeader(t string) *Header {
	return &Header{
		Type: t,
	}
}

func UnmarshaledHeader(payload []byte) (*Header, error) {
	header := &Header{}
	err := json.Unmarshal(payload, &header)
	return header, err
}

func Verify(header *Header) bool {
	switch header.Type {
	case HeaderTypeToString(FABRIC):
		return true
	case HeaderTypeToString(ETHEREUM):
		return true
	case HeaderTypeToString(BITCOIN):
		sig, err := crypto.ParseSignature(header.Signature)
		if err != nil {
			return false
		}
		pubkey, err := crypto.DecompressPubkey(header.PublicKey)
		if err != nil {
			return false
		}
		return sig.Verify(header.Hash, pubkey)
	default:
		return false
	}
	return false
}

func Address(header *Header) string {
	switch header.Type {
	case HeaderTypeToString(FABRIC):
		return hexutil.Encode(common.BytesToAddress(header.PublicKey).Bytes())
	case HeaderTypeToString(ETHEREUM):
		recoveredPub, err := crypto.Ecrecover(header.Hash, header.Signature)
		if err != nil {
			return ""
		}
		pubKey, err := crypto.UnmarshalPubkey(recoveredPub)
		if err != nil {
			return ""
		}
		return crypto.PubkeyToAddress(*pubKey).String()
	case HeaderTypeToString(BITCOIN):
		payload := crypto.Hash160(header.PublicKey)
		versionedPayload := append([]byte{0x00}, payload...)
		checksum := crypto.DoubleHash(payload)[:crypto.ChecksumLength]
		fullPayload := append(versionedPayload, checksum...)
		return crypto.Base58Encode(fullPayload)
	default:
		return ""
	}
	return ""
}
