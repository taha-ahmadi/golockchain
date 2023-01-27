package crypto

import (
	"crypto/ed25519"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	require.Equal(t, len(privKey.Bytes()), ed25519.PrivateKeySize)
	pubKey := privKey.Public()
	require.Equal(t, len(pubKey.Bytes()), ed25519.PublicKeySize)
}
func TestPrivateKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	msg := []byte("Test message")
	sig := privKey.Sign(msg)

	require.True(t, sig.Verify(pubKey, msg))
	require.False(t, sig.Verify(pubKey, []byte("f")))

	invalidPrivKey := GeneratePrivateKey()
	invalidPubKey := invalidPrivKey.Public()
	require.False(t, sig.Verify(invalidPubKey, msg))
}

func TestPublicKeyToAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	publicKey := privKey.Public()
	address := publicKey.Address()
	require.Equal(t, addressLen, len(address.Bytes()))
}
