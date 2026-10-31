package trapdoor

import "errors"

// Key serialisation. Keys are stored as a small binary record so they can be
// poop-encoded like everything else:
//
//	public:  "POOK" | version | len(N) u16 | N | len(E) u16 | E
//	private: "POOS" | version | public record | len(D) u16 | D | len(P) u16 | P | len(Q) u16 | Q
//
// Lambda is recomputed on load. Big-endian, minimal big.Int.Bytes() encoding.

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

// UnmarshalPrivate parses a private key and recomputes Lambda.
func UnmarshalPrivate(b []byte) (*PrivateKey, error) {
	// TODO
	return nil, ErrBadKeyEncoding
}
