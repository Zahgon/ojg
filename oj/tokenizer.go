// Copyright (c) 2021, Peter Ohler, All rights reserved.

package oj

import (
	"io"

	"github.com/ohler55/ojg/gen"
)

const (
	objectStart = '{'
	arrayStart  = '['
)

// Tokenizer is used to tokenize a JSON document.
type Tokenizer struct {
	tracker
	tmp       []byte // used for numbers and strings
	runeBytes []byte
	starts    []byte
	handler   TokenHandler
	ri        int // read index for null, false, and true
	mi        int
	num       gen.Number
	rn        rune
	mode      string
	nextMode  string
}

// TokenizeString the provided JSON and call the handler functions for each
// token in the JSON.
func TokenizeString(data string, handler TokenHandler) error { _ = "STUB: not implemented"; return nil }

// Tokenize the provided JSON and call the TokenHandler functions for each
// token in the JSON.
func Tokenize(data []byte, handler TokenHandler) error { _ = "STUB: not implemented"; return nil }

// TokenizeLoad JSON from a io.Reader and call the TokenHandler functions for
// each token in the JSON.
func TokenizeLoad(r io.Reader, handler TokenHandler) error { _ = "STUB: not implemented"; return nil }

// Parse the JSON and call the handler functions for each token in the JSON.
func (t *Tokenizer) Parse(buf []byte, handler TokenHandler) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Skip BOM if present.

// Load aand parse the JSON and call the handler functions for each token in
// the JSON.
func (t *Tokenizer) Load(r io.Reader, handler TokenHandler) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Skip BOM if present.

func (t *Tokenizer) tokenizeBuffer(buf []byte, last bool) error {
	_ = "STUB: not implemented"
	return nil
}

// skip and continue

// Only modes with a close array are value, after, and numbers
// which are all over 256 long.

// valid finishing maps are one byte longer

func (t *Tokenizer) handleNum() { _ = "STUB: not implemented"; return }
