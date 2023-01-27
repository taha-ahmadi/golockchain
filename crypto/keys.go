package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"io"
)

const (
	addressLen = 20
	seedLen    = 32
)

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

type PrivateKey struct {
	key ed25519.PrivateKey // Ed25519 is generally considered to be more secure and efficient than ECDSA.
	// Ed25519 has a smaller signature size and faster signing and verification times.
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

func (p *PublicKey) Bytes() []byte {
	return p.key
}

type Signature struct {
	value []byte
}

func (s *Signature) Bytes() []byte {
	return s.value
}

type Address struct {
	value []byte
}

func NewPrivateKeyFromString(seed string) *PrivateKey {
	s, err := hex.DecodeString(seed)
	if err != nil {
		panic(err)
	}

	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(s),
	}

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
func main() {

}
