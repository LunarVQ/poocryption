// Package emojiset encodes bytes as a sequence of distinct emoji. It is the
// key format: keygen output is rendered through this, ciphertext through
// package poop, so the two are visually impossible to confuse.
//
// Alphabet rules (matter for unambiguous decoding):
//   - exactly 256 entries, one per byte value, all distinct
//   - each entry is a single Unicode code point: no ZWJ sequences, no skin
//     tone modifiers, no variation selectors, no flags (regional-indicator pairs)
//   - stable forever once keys exist; changing it invalidates every key
//
// Pure transport, no secrecy. A 3072-bit public key is ~390 emoji.
package emojiset

import "errors"

// Alphabet maps byte value -> emoji. Must satisfy the rules above.
// TODO: fill with 256 single-code-point emoji (e.g. ranges from U+1F300..U+1F5FF
// and U+1F600..U+1F64F, filtered to fully-qualified, non-modifier characters).
var Alphabet = [256]rune{}

// index is the reverse map, built from Alphabet at init.
var index map[rune]byte

func init() {
	// TODO: populate index; panic if Alphabet has duplicates or zero entries,
	// so a broken alphabet is caught at startup rather than at decode time.
}

var ErrInvalid = errors.New("emojiset: rune not in alphabet")

// Encode renders data as one emoji per byte.
func Encode(data []byte) string {
	// TODO
	return ""
}

// Decode reverses Encode. Whitespace is ignored; any other unknown rune is ErrInvalid.
func Decode(s string) ([]byte, error) {
	// TODO
	return nil, ErrInvalid
}
