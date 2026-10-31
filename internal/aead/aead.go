// Package aead is the authenticated cipher that protects the payload once the
// trapdoor has delivered the session key. AES-256-GCM, standard library only.
package aead

const (
	KeySize   = 32
	NonceSize = 12 // GCM standard nonce; fresh random per message
	Overhead  = 16
)

// Seal encrypts and authenticates plaintext, binding aad (the header) to it.
func Seal(key, nonce, plaintext, aad []byte) ([]byte, error) {
	// TODO: block := aes.NewCipher(key); g := cipher.NewGCM(block); g.Seal(nil, nonce, plaintext, aad)
	return nil, nil
}

// Open verifies and decrypts ciphertext produced by Seal.
func Open(key, nonce, ciphertext, aad []byte) ([]byte, error) {
	// TODO: cipher.NewGCM(aes.NewCipher(key)).Open(nil, nonce, ciphertext, aad)
	return nil, nil
}
