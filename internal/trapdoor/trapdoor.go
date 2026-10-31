// Package trapdoor is the asymmetric core: a cycle anyone can step around
// forward, with a shortcut only the private-key holder has.
//
// The map is x -> x^e mod N on Z_N*. Iterating it from a start value walks a
// cycle and eventually returns to the start (e.g. mod 33, e = 3:
// 2 -> 8 -> 17 -> 29 -> 2). Encryption is one step. Decryption is "the rest
// of the way around", and the private exponent d collapses that into a single
// exponentiation because d = e^-1 mod lambda(N). lambda(N) is only computable
// from the factorisation of N, which is the trapdoor.
//
// Raw x^e is deterministic and malleable, so the payload key is never locked
// raw: Wrap/Unwrap apply OAEP padding on top of the same map. Step/Walk are
// the bare primitives, kept for understanding and testing.
package trapdoor

import (
	"errors"
	"math/big"
)

// DefaultBits is the modulus size for GenerateKey.
const DefaultBits = 3072

// PublicKey is (N, e): the cycle anyone can step around.
type PublicKey struct {
	N *big.Int
	E *big.Int
}

// PrivateKey adds the shortcut. P, Q and Lambda are kept so the key can be
// exported/inspected and so Shortcut can reduce exponents mod lambda(N).
type PrivateKey struct {
	PublicKey
	D      *big.Int // e^-1 mod Lambda
	P, Q   *big.Int
	Lambda *big.Int // lcm(P-1, Q-1): the cycle period
}

var (
	ErrKeySize  = errors.New("trapdoor: modulus too small")
	ErrBadInput = errors.New("trapdoor: input out of range")
)

// GenerateKey produces a fresh keypair with an N of the given bit size.
func GenerateKey(bits int) (*PrivateKey, error) {
	// TODO: p, q := crypto/rand.Prime(bits/2); N = p*q; Lambda = lcm(p-1, q-1);
	// E = 65537 (check gcd(E, Lambda) == 1); D = E^-1 mod Lambda.
	// Alternatively wrap crypto/rsa.GenerateKey and copy the fields out.
	return nil, errors.New("trapdoor: GenerateKey not implemented")
}

// Step is one pass around the cycle: x^e mod N. Public operation.
func (pub *PublicKey) Step(x *big.Int) *big.Int {
	// TODO: new(big.Int).Exp(x, pub.E, pub.N)
	return nil
}

// Walk is t passes around the cycle. Costs t exponentiations for anyone
// without the private key. Useful for demonstrating the cycle and for a
// time-lock variant (x -> x^2 mod N, large t).
func (pub *PublicKey) Walk(x *big.Int, t uint64) *big.Int {
	// TODO: loop Step t times
	return nil
}

// Shortcut computes the same result as Walk(x, t) in one exponentiation,
// by reducing the exponent e^t mod Lambda first.
func (priv *PrivateKey) Shortcut(x *big.Int, t uint64) *big.Int {
	// TODO: k := e^t mod Lambda; return x^k mod N
	return nil
}

// Unlock inverts Step: c^d mod N. This is "finish going around the cycle"
// collapsed into one operation.
func (priv *PrivateKey) Unlock(c *big.Int) *big.Int {
	// TODO: new(big.Int).Exp(c, priv.D, priv.N)
	return nil
}

// Wrap locks a short secret (the payload key) under pub with OAEP padding.
// Output length equals the modulus length in bytes.
func (pub *PublicKey) Wrap(secret []byte) ([]byte, error) {
	// TODO: crypto/rsa.EncryptOAEP(sha256.New(), rand.Reader, pub.toRSA(), secret, label)
	return nil, errors.New("trapdoor: Wrap not implemented")
}

// Unwrap recovers a secret produced by Wrap.
func (priv *PrivateKey) Unwrap(wrapped []byte) ([]byte, error) {
	// TODO: crypto/rsa.DecryptOAEP
	return nil, errors.New("trapdoor: Unwrap not implemented")
}

// Size returns the modulus length in bytes (the Wrap output length).
func (pub *PublicKey) Size() int {
	// TODO: (pub.N.BitLen() + 7) / 8
	return 0
}

// Fingerprint identifies a public key: SHA-256 over the canonical encoding of (N, e).
func (pub *PublicKey) Fingerprint() [32]byte {
	// TODO
	return [32]byte{}
}
