package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"io"
)

const (
	addressLen   = 20
	seedLen      = 32
	signatureLen = 64
)

type PrivateKey struct {
	key ed25519.PrivateKey // Ed25519 is generally considered to be more secure and efficient than ECDSA.
	// Ed25519 has a smaller signature size and faster signing and verification times.
}

// NewPrivateKeyFromString get an existed private key and Return *PrivateKey
func NewPrivateKeyFromString(s string) *PrivateKey {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return NewPrivateKeyFromSeed(b)
}

func NewPrivateKeyFromSeed(seed []byte) *PrivateKey {
	if len(seed) != seedLen {
		panic("invalid seed length, must be 32")
	}

	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}
}

func GeneratePrivateKey() *PrivateKey {
	seed := make([]byte, seedLen)
	_, err := io.ReadFull(rand.Reader, seed)
	if err != nil {
		panic(err)
	}
	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}
}

func (p *PrivateKey) Public() *PublicKey {
	pubKey := make([]byte, ed25519.PublicKeySize)
	copy(pubKey, p.key[32:])

	return &PublicKey{
		key: pubKey,
	}
}

func (p *PrivateKey) Bytes() []byte {
	return p.key
}
func (p *PrivateKey) Sign(msg []byte) *Signature {

	sign := ed25519.Sign(p.key, msg)

	return &Signature{
		value: sign,
	}
}

type PublicKey struct {
	key ed25519.PublicKey
}

func PublicKeyFromBytes(b []byte) *PublicKey {
	if len(b) != ed25519.PublicKeySize {
		panic("invalid public key length")
	}
	return &PublicKey{
		key: ed25519.PublicKey(b),
	}
}

func (p *PublicKey) Bytes() []byte {
	return p.key
}

type Signature struct {
	value []byte
}

func SignatureFromBytes(b []byte) *Signature {
	if len(b) != signatureLen {
		panic("length of the bytes not equal to 64")
	}
	return &Signature{
		value: b,
	}
}
func (s *Signature) Bytes() []byte {
	return s.value
}

type Address struct {
	value []byte
}

func (a Address) String() string {
	return hex.EncodeToString(a.value)
}
func (s *Signature) Verify(pubkey *PublicKey, msg []byte) bool {

	return ed25519.Verify(pubkey.key, msg, s.value)
}

// Bytes functions for serialization purposes
func (a Address) Bytes() []byte {
	return a.value
}
func (p *PublicKey) Address() Address {
	return Address{
		value: p.key[len(p.key)-addressLen:],
	}
}
