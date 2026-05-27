// Copyright (c) 2020, Peter Ohler, All rights reserved.

package gen

import (
	"io"
)

const (
	stackInitSize = 32 // for container stack { or [
	tmpInitSize   = 32 // for tokens and numbers
	mapInitSize   = 8
	readBufSize   = 4096
)

// Parser is a reusable JSON parser. It can be reused for multiple parsings
// which allows buffer reuse for a performance advantage.
type Parser struct {
	tmp        []byte // used for numbers and strings
	runeBytes  []byte
	stack      []Node
	starts     []int
	maps       []Object
	cb         func(Node)
	resultChan chan Node
	line       int
	noff       int // Offset of last newline from start of buf. Can be negative when using a reader.
	ri         int // read index for null, false, and true
	mi         int
	num        Number
	rn         rune
	result     Node
	mode       string
	nextMode   string

	// OnlyOne returns an error if more than one JSON is in the string or stream.
	OnlyOne bool

	// Reuse maps. Previously returned maps will no longer be valid or rather
	// could be modified during parsing.
	Reuse bool
}

// Parse a JSON string in to simple types. An error is returned if not valid JSON.
func (p *Parser) Parse(buf []byte, args ...any) (Node, error) {
	_ = "STUB: not implemented"
	return *new(Node), nil
}

// Skip BOM if present.

// ParseReader a JSON io.Reader. An error is returned if not valid JSON.
func (p *Parser) ParseReader(r io.Reader, args ...any) (data Node, err error) {
	_ = "STUB: not implemented"
	return *new(Node), nil
}

// Skip BOM if present.

func (p *Parser) parseBuffer(buf []byte, last bool) error { _ = "STUB: not implemented"; return nil }

// skip and continue

// Only modes with a close array are value, after, and numbers
// which are all over 256 long.

// valid finishing maps are one byte longer

func (p *Parser) add(n Node) { _ = "STUB: not implemented"; return }

func (p *Parser) newError(off int, format string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) byteError(off int, mode string, b byte, r rune) error {
	_ = "STUB: not implemented"
	return nil
}
