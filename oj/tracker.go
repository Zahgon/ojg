// Copyright (c) 2020, Peter Ohler, All rights reserved.

package oj

type tracker struct {
	line int
	noff int // Offset of last newline from start of buf. Can be negative when using a reader.

	// OnlyOne returns an error if more than one JSON is in the string or stream.
	OnlyOne bool
}

func (t *tracker) newError(off int, format string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *tracker) byteError(off int, mode string, b byte, r rune) error {
	_ = "STUB: not implemented"
	return nil
}
