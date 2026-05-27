// Copyright (c) 2020, Peter Ohler, All rights reserved.

package oj

import (
	"io"

	"github.com/ohler55/ojg/alt"
	"github.com/ohler55/ojg/gen"
)

const (
	stackInitSize = 32 // for container stack { or [
	tmpInitSize   = 32 // for tokens and numbers
	mapInitSize   = 8
	readBufSize   = 4096
)

var emptySlice = []any{}

// Parser is a reusable JSON parser. It can be reused for multiple parsings
// which allows buffer reuse for a performance advantage.
type Parser struct {
	tracker
	tmp        []byte // used for numbers and strings
	runeBytes  []byte
	stack      []any
	starts     []int
	maps       []map[string]any
	cb         func(any)
	resultChan chan any
	ri         int // read index for null, false, and true
	mi         int
	num        gen.Number
	rn         rune
	result     any
	mode       string
	nextMode   string

	// Reuse maps. Previously returned maps will no longer be valid or rather
	// could be modified during parsing.
	Reuse bool
}

func recomposeToJSON(v any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

func init() {
	alt.DefaultRecomposer.RegisterUnmarshalerComposer(recomposeToJSON)
}

// Unmarshal parses the provided JSON and stores the result in the value
// pointed to by vp.
func (p *Parser) Unmarshal(data []byte, vp any, recomposer ...alt.Recomposer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Parse a JSON string in to simple types. An error is returned if not valid JSON.
func (p *Parser) Parse(buf []byte, args ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Skip BOM if present.

// ParseReader reads JSON from an io.Reader. An error is returned if not valid
// JSON.
func (p *Parser) ParseReader(r io.Reader, args ...any) (data any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Skip BOM if present.

func (p *Parser) parseBuffer(buf []byte, last bool) error { _ = "STUB: not implemented"; return nil }

// skip and continue

// Only modes with a close array are value, after, and numbers
// which are all over 256 long.

// valid finishing maps are one byte longer

func (p *Parser) add(n any) { _ = "STUB: not implemented"; return }
