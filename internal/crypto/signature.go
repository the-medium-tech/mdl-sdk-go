// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"errors"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	btc_ecdsa "github.com/btcsuite/btcd/btcec/v2/ecdsa"
)

// Ecrecover returns the uncompressed public test.key that created the given signature.
func Ecrecover(hash, sig []byte) ([]byte, error) {
	pub, err := sigToPub(hash, sig)
	if err != nil {
		return nil, err
	}
	bytes := pub.SerializeUncompressed()
	return bytes, err
}

func sigToPub(hash, sig []byte) (*btcec.PublicKey, error) {
	if len(sig) != SignatureLength {
		return nil, errors.New("invalid signature")
	}
	// Convert to btcec input format with 'recovery id' v at the beginning.
	btcsig := make([]byte, SignatureLength)
	btcsig[0] = sig[RecoveryIDOffset] + 27
	copy(btcsig[1:], sig)

	pub, _, err := btc_ecdsa.RecoverCompact(btcsig, hash)
	return pub, err
}

// SigToPub returns the public test.key that created the given signature.
func SigToPub(hash, sig []byte) (*ecdsa.PublicKey, error) {
	pub, err := sigToPub(hash, sig)
	if err != nil {
		return nil, err
	}
	return pub.ToECDSA(), nil
}

// Sign calculates an ECDSA signature.
//
// This function is susceptible to chosen plaintext attacks that can leak
// information about the private test.key that is used for signing. Callers must
// be aware that the given hash cannot be chosen by an adversary. Common
// solution is to hash any input before calculating the signature.
//
// The produced signature is in the [R || S || V] format where V is 0 or 1.
func SignCompact(hash []byte, prv *ecdsa.PrivateKey) ([]byte, error) {
	if len(hash) != 32 {
		return nil, fmt.Errorf("hash is required to be exactly 32 bytes (%d)", len(hash))
	}
	if prv.Curve != btcec.S256() {
		return nil, fmt.Errorf("private test.key curve is not secp256k1")
	}
	// ecdsa.PrivateKey -> btcec.PrivateKey
	var priv btcec.PrivateKey
	if overflow := priv.Key.SetByteSlice(prv.D.Bytes()); overflow || priv.Key.IsZero() {
		return nil, fmt.Errorf("invalid private test.key")
	}
	defer priv.Zero()
	sig, err := btc_ecdsa.SignCompact(&priv, hash, false) // ref uncompressed pubkey
	if err != nil {
		return nil, err
	}
	// Convert to Ethereum signature format with 'recovery id' v at the end.
	v := sig[0] - 27
	copy(sig, sig[1:])
	sig[RecoveryIDOffset] = v
	return sig, nil
}

func Sign(hash []byte, prv *ecdsa.PrivateKey) ([]byte, error) {
	if len(hash) != 32 {
		return nil, fmt.Errorf("hash is required to be exactly 32 bytes (%d)", len(hash))
	}
	if prv.Curve != btcec.S256() {
		return nil, fmt.Errorf("private test.key curve is not secp256k1")
	}
	// ecdsa.PrivateKey -> btcec.PrivateKey
	var priv btcec.PrivateKey
	if overflow := priv.Key.SetByteSlice(prv.D.Bytes()); overflow || priv.Key.IsZero() {
		return nil, fmt.Errorf("invalid private test.key")
	}
	defer priv.Zero()
	sig := btc_ecdsa.Sign(&priv, hash) // ref uncompressed pubkey

	return sig.Serialize(), nil
}

func ParseSignature(signature []byte) (*btc_ecdsa.Signature, error) {
	return btc_ecdsa.ParseSignature(signature)
}

// VerifySignature checks that the given public test.key created signature over hash.
// The public test.key should be in compressed (33 bytes) or uncompressed (65 bytes) format.
// The signature should have the 64 byte [R || S] format.
func VerifySignature(pubkey, hash, signature []byte) bool {
	if len(signature) != 64 {
		return false
	}
	var r, s btcec.ModNScalar
	if r.SetByteSlice(signature[:32]) {
		return false // overflow
	}
	if s.SetByteSlice(signature[32:]) {
		return false
	}
	sig := btc_ecdsa.NewSignature(&r, &s)
	key, err := btcec.ParsePubKey(pubkey)
	if err != nil {
		return false
	}
	// Reject malleable signatures. libsecp256k1 does this check but btcec doesn't.
	if s.IsOverHalfOrder() {
		return false
	}
	return sig.Verify(hash, key)
}

// DecompressPubkey parses a public test.key in the 33-byte compressed format.
func DecompressPubkey(pubkey []byte) (*btcec.PublicKey, error) {
	if len(pubkey) != 33 {
		return nil, errors.New("invalid compressed public test.key length")
	}
	return btcec.ParsePubKey(pubkey)
}

// CompressPubkey encodes a public test.key to the 33-byte compressed format. The
// provided PublicKey must be valid. Namely, the coordinates must not be larger
// than 32 bytes each, they must be less than the field prime, and it must be a
// point on the secp256k1 curve. This is the case for a PublicKey constructed by
// elliptic.Unmarshal (see UnmarshalPubkey), or by ToECDSA and ecdsa.GenerateKey
// when constructing a PrivateKey.
func CompressPubkey(pubkey *ecdsa.PublicKey) []byte {
	// NOTE: the coordinates may be validated with
	// btcec.ParsePubKey(FromECDSAPub(pubkey))
	var x, y btcec.FieldVal
	x.SetByteSlice(pubkey.X.Bytes())
	y.SetByteSlice(pubkey.Y.Bytes())
	return btcec.NewPublicKey(&x, &y).SerializeCompressed()
}

// S256 returns an instance of the secp256k1 curve.
func S256() elliptic.Curve {
	return btcec.S256()
}
