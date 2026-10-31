// Package poocryption encrypts data to a recipient's public key and renders
// the result as a string of 💩. Standard library only.
//
// Pipeline (hybrid encryption):
//
//	sender:    S ← random 32 bytes
//	           wrapped = trapdoor.Wrap(pub, S)            one step around the cycle, OAEP-padded
//	           K = HKDF(S, salt = fingerprint(pub), info)
//	           ct = AES-256-GCM(K, nonce, plaintext, aad = header)
//	           header ‖ ct ──poop.Encode──► 💩💩💩…
//	receiver:  S = trapdoor.Unwrap(priv, wrapped)         the shortcut around the cycle
//	           K = HKDF(S, …); plaintext = AES-GCM⁻¹(K, …)
//
// Asymmetry comes entirely from internal/trapdoor (RSA: x -> x^e mod N, with
// d = e^-1 mod lambda(N) as the shortcut). Everything after Unwrap is symmetric.
package poocryption

import (
	"errors"

	"github.com/LunarVQ/poocryption/internal/trapdoor"
)

const (
	// DefaultKeyBits is the modulus size used by GenerateKey.
	DefaultKeyBits = trapdoor.DefaultBits

	sessionSecretSize = 32
	dataKeyInfo       = "poocryption/v1/data-key"
)

var (
	ErrDecrypt      = errors.New("poocryption: decryption failed (wrong key or corrupted data)")
	ErrKeyMismatch  = errors.New("poocryption: ciphertext was not encrypted to this key")
	ErrNoPublicKey  = errors.New("poocryption: public key required")
	ErrNoPrivateKey = errors.New("poocryption: private key required")
)

// PublicKey and PrivateKey are re-exported so callers never import internal/.
type (
	PublicKey  = trapdoor.PublicKey
	PrivateKey = trapdoor.PrivateKey
)

// Options tune encryption. The zero value means "use defaults".
type Options struct {
	BytesPerEmoji int // zero -> poop.DefaultBytesPerEmoji
}

// GenerateKey creates a fresh keypair. bits <= 0 selects DefaultKeyBits.
func GenerateKey(bits int) (*PrivateKey, error) {
	if bits <= 0 {
		bits = DefaultKeyBits
	}
	return trapdoor.GenerateKey(bits)
}

// Encrypt returns plaintext encrypted to pub as a poop-emoji string.
func Encrypt(plaintext []byte, pub *PublicKey, opts *Options) (string, error) {
	if pub == nil {
		return "", ErrNoPublicKey
	}
	// TODO:
	//  1. S := random sessionSecretSize bytes; nonce := random aead.NonceSize bytes
	//  2. wrapped := pub.Wrap(S)
	//  3. fp := pub.Fingerprint(); K := kdf.Expand(S, fp[:], dataKeyInfo, aead.KeySize)
	//  4. hdr := format.Header{KeyID: fp, Wrapped: wrapped, Nonce: nonce}.Marshal()
	//  5. ct := aead.Seal(K, nonce, plaintext, hdr)
	//  6. return poop.Encode(hdr ‖ ct, opts.BytesPerEmoji)
	return "", errors.New("poocryption: Encrypt not implemented")
}

// Decrypt reverses Encrypt with the matching private key. Any failure after
// the key-ID check returns ErrDecrypt; callers must not distinguish causes.
func Decrypt(poopText string, priv *PrivateKey) ([]byte, error) {
	if priv == nil {
		return nil, ErrNoPrivateKey
	}
	// TODO:
	//  1. raw := poop.Decode(poopText); hdr, ct := format.Unmarshal(raw)
	//  2. if hdr.KeyID != priv.Fingerprint() -> ErrKeyMismatch
	//  3. S := priv.Unwrap(hdr.Wrapped)            (failure -> ErrDecrypt)
	//  4. K := kdf.Expand(S, hdr.KeyID[:], dataKeyInfo, aead.KeySize)
	//  5. aead.Open(K, hdr.Nonce, ct, hdr.Marshal()) (failure -> ErrDecrypt)
	return nil, errors.New("poocryption: Decrypt not implemented")
}
