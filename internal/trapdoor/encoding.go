package trapdoor

import "errors"

// Key serialisation. Keys are stored as a small binary record and rendered
// through internal/emojiset (ciphertext uses package poop; keys never do).
//
//	public:  "POOK" | version | len(N) u16 | N | len(E) u16 | E
//	private: "POOS" | version | len(P) u16 | P | len(Q) u16 | Q | len(E) u16 | E
//
// The private record carries only P, Q and E: N, Lambda and D are recomputed
// on load. Storing D and N as well would triple the emoji count for no gain.
// Big-endian, minimal big.Int.Bytes() encoding.

const (
	publicMagic  = "POOK"
	privateMagic = "POOS"
	keyVersion   = 1
)

var ErrBadKeyEncoding = errors.New("trapdoor: malformed key encoding")

// MarshalPublic serialises a public key.
func (pub *PublicKey) MarshalPublic() []byte {
	// TODO
	return nil
}

// UnmarshalPublic parses a public key.
func UnmarshalPublic(b []byte) (*PublicKey, error) {
	// TODO
	return nil, ErrBadKeyEncoding
}

// MarshalPrivate serialises a private key.
func (priv *PrivateKey) MarshalPrivate() []byte {
	// TODO
	return nil
}

// UnmarshalPrivate parses a private key and recomputes N, Lambda and D.
func UnmarshalPrivate(b []byte) (*PrivateKey, error) {
	// TODO
	return nil, ErrBadKeyEncoding
}
