// Copyright (c) 2021, Peter Ohler, All rights reserved.

package ojg

const hex = "0123456789abcdef"

var (
	maxTokenLen = 64

	// Copied from sen/maps.go

	//   0123456789abcdef0123456789abcdef
	jMap = "" +
		`........btn.fr..................` + // 0x00
		`oo"ooohooooooooooooooooooooohoho` + // 0x20
		`oooooooooooooooooooooooooooo\ooo` + // 0x40
		"ooooooooooooooooooooooooooooooo." + // 0x60
		`88888888888888888888888888888888` + // 0x80
		`88888888888888888888888888888888` + // 0xa0
		`88888888888888888888888888888888` + // 0xc0
		`88888888888888888888888888888888` //  0xe0

	//   0123456789abcdef0123456789abcdef
	senMap = "" +
		`........bxx.fr..................` + // 0x00
		`xx"xoxhxxxooxoox0000000000xxhxho` + // 0x20
		`ooooooooooooooooooooooooooox\xoo` + // 0x40
		"oooooooooooooooooooooooooooxoxo." + // 0x60
		`88888888888888888888888888888888` + // 0x80
		`88888888888888888888888888888888` + // 0xa0
		`88888888888888888888888888888888` + // 0xc0
		`88888888888888888888888888888888` //  0xe0
)

// AppendJSONString appends a JSON encoding of a string to the provided byte
// slice.
func AppendJSONString(buf []byte, s string, htmlSafe bool) []byte {
	_ = "STUB: not implemented"
	return nil
}

// AppendSENString appends a SEN encoding of a string to the provided byte
// slice.
func AppendSENString(buf []byte, s string, htmlSafe bool) []byte {
	_ = "STUB: not implemented"
	return nil
}
