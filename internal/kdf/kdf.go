// Package kdf expands the unwrapped session secret into the payload key.
package kdf

// Expand derives n bytes from secret bound to salt and info via HKDF-SHA256.
// Different info strings yield independent keys from the same secret.
func Expand(secret, salt []byte, info string, n int) ([]byte, error) {
	// TODO: crypto/hkdf.Key(sha256.New, secret, salt, info, n)   (stdlib, Go 1.24+)
	return nil, nil
}
