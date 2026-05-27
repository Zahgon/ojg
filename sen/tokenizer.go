// Copyright (c) 2021, Peter Ohler, All rights reserved.

package sen

import (
	"io"

	"github.com/ohler55/ojg/gen"
	"github.com/ohler55/ojg/oj"
)

const (
	objectStart = '{'
	arrayStart  = '['
)

// Tokenizer is a reusable JSON tokenizer. It can be reused for multiple parsings
// which allows buffer reuse for a performance advantage.
type Tokenizer struct {
	tmp       []byte // used for numbers and strings
	runeBytes []byte
	starts    []byte
	handler   oj.TokenHandler
	line      int
	noff      int // Offset of last newline from start of buf. Can be negative when using a reader.
	ri        int // read index for null, false, and true
	mi        int
	num       gen.Number
	rn        rune
	mode      string
	exkey     bool

	// OnlyOne returns an error if more than one JSON is in the string or stream.
	OnlyOne bool
}

// TokenizeString the provided JSON and call the handler functions for each
// token in the JSON.
func TokenizeString(data string, handler oj.TokenHandler) error {
	_ = "STUB: not implemented"
	return nil
}

// Tokenize a JSON string in to simple types. An error is returned if not valid JSON.
func Tokenize(data []byte, handler oj.TokenHandler) error { _ = "STUB: not implemented"; return nil }

// TokenizeLoad a JSON io.Reader. An error is returned if not valid JSON.
func TokenizeLoad(r io.Reader, handler oj.TokenHandler) error {
	_ = "STUB: not implemented"
	return nil
}

// Parse a JSON string in to simple types. An error is returned if not valid JSON.
func (t *Tokenizer) Parse(buf []byte, handler oj.TokenHandler) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Skip BOM if present.

// Load a JSON io.Reader. An error is returned if not valid JSON.
func (t *Tokenizer) Load(r io.Reader, handler oj.TokenHandler) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Skip BOM if present.

func (t *Tokenizer) tokenizeBuffer(buf []byte, last bool) { _ = "STUB: not implemented"; return }

// end of buf reached

// skip and continue

// Only modes with a close array are value, token, and numbers
// which are all over 256 long.

// valid finishing maps are one byte longer

// number

// token

func (t *Tokenizer) addToken(s string) { _ = "STUB: not implemented"; return }

func (t *Tokenizer) addString(s string) { _ = "STUB: not implemented"; return }

func (t *Tokenizer) newError(off int, format string, args ...any) {
	_ = "STUB: not implemented"
	return
}

func (t *Tokenizer) byteError(off int, mode string, b byte) { _ = "STUB: not implemented"; return }

func (t *Tokenizer) handleNum(off int) { _ = "STUB: not implemented"; return }
