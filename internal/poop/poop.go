// Package poop encodes bytes as a string of 💩.
//
// A bare U+1F4A9 carries no information, so each 💩 is followed by invisible
// Unicode variation selectors. There are exactly 256 of them
// (U+FE00..U+FE0F and U+E0100..U+E01EF), so one selector = one byte. Fonts
// render the run as a plain 💩; the selectors survive copy/paste.
//
// This is transport only. It provides no secrecy.
package poop

import "errors"

// Emoji is the carrier character.
const Emoji = '\U0001F4A9'

// DefaultBytesPerEmoji is how many payload bytes trail each 💩.
const DefaultBytesPerEmoji = 8

var ErrInvalid = errors.New("poop: not valid poop-encoded data")

// Encode renders data as 💩 characters carrying variation selectors.
func Encode(data []byte, bytesPerEmoji int) string {
	// TODO: every bytesPerEmoji bytes emit Emoji; per byte emit
	//   b < 16  -> rune(0xFE00 + b)
	//   b >= 16 -> rune(0xE0100 + b - 16)
	return ""
}

// Decode reverses Encode. Whitespace is ignored; any other rune is ErrInvalid.
func Decode(s string) ([]byte, error) {
	// TODO
	return nil, ErrInvalid
}
