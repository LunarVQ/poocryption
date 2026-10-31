// Package format defines the binary container in front of the ciphertext.
// The header is public and is fed to the AEAD as associated data, so it
// cannot be altered without detection.
//
//	"POOC" | version u8 | key fingerprint [32] | len(wrapped) u16 | wrapped key | nonce [12]
//
// "wrapped key" is the session secret locked under the recipient's public key
// (trapdoor.Wrap); its length equals the modulus size, hence the length prefix.
package format

import (
	"errors"

	"github.com/LunarVQ/poocryption/internal/aead"
)

const (
	Magic   = "POOC"
	Version = 1

	// FixedSize is the header length excluding the variable wrapped key.
	FixedSize = 4 + 1 + 32 + 2 + aead.NonceSize
)

var (
	ErrBadMagic  = errors.New("format: not a poocryption container")
	ErrTruncated = errors.New("format: container too short")
)

// Header is the public, authenticated metadata of a ciphertext.
type Header struct {
	KeyID   [32]byte // trapdoor.PublicKey.Fingerprint() of the intended recipient
	Wrapped []byte   // trapdoor.Wrap output
	Nonce   [aead.NonceSize]byte
}

// Marshal serialises the header (big-endian).
func (h *Header) Marshal() []byte {
	// TODO
	return nil
}

// Unmarshal parses a header from the front of b and returns the remaining
// bytes (the ciphertext).
func Unmarshal(b []byte) (*Header, []byte, error) {
	// TODO: check length and magic/version, read fields, bound len(wrapped).
	return nil, nil, ErrTruncated
}
