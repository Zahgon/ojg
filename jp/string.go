// Copyright (c) 2023, Peter Ohler, All rights reserved.

package jp

const hex = "0123456789abcdef"

var (
	//   0123456789abcdef0123456789abcdef
	jMap = "" +
		`........btn.fr..................` + // 0x00
		`oo"oooo'oooooooooooooooooooooooo` + // 0x20
		`oooooooooooooooooooooooooooo\ooo` + // 0x40
		"ooooooooooooooooooooooooooooooo." + // 0x60
		`88888888888888888888888888888888` + // 0x80
		`88888888888888888888888888888888` + // 0xa0
		`88888888888888888888888888888888` + // 0xc0
		`88888888888888888888888888888888` //  0xe0
)

// AppendString to a buffer while escaping characters as necessary.
func AppendString(buf []byte, s string, delim byte) []byte { _ = "STUB: not implemented"; return nil }

// don't escape regexp as they are already escaped
