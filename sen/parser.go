// Copyright (c) 2020, Peter Ohler, All rights reserved.

package sen

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
	emptyKey      = gen.Key("")
)

var (
	emptySlice = []any{}
)

// TokenFunc is a function that can be used to evaluate functions embedded in
// a SEN file.
type TokenFunc func(args ...any) any

// Parser is a reusable JSON parser. It can be reused for multiple parsings
// which allows buffer reuse for a performance advantage.
type Parser struct {
	tmp        []byte // used for numbers and strings
	runeBytes  []byte
	stack      []any
	starts     []int
	maps       []map[string]any
	cb         func(any)
	resultChan chan any
	line       int
	noff       int // Offset of last newline from start of buf. Can be negative when using a reader.
	ri         int // read index for null, false, and true
	mi         int
	num        gen.Number
	rn         rune
	result     any
	mode       string
	lastKey    gen.Key
	lastStrKey gen.Key
	tokenFuncs map[string]TokenFunc
	quoteDelim byte

	// Reuse maps. Previously returned maps will no longer be valid or rather
	// could be modified during parsing.
	Reuse bool

	// OnlyOne returns an error if more than one JSON is in the string or stream.
	OnlyOne bool

	plus bool
}

// AddTokenFunc add a token function that can appear in the data being
// parsed. As an example `[ISODate("2021-06-28T10:11:12Z")]` could be parsed
// to a time.Time.
func (p *Parser) AddTokenFunc(name string, tf TokenFunc) { _ = "STUB: not implemented"; return }

// Unmarshal parses the provided JSON and stores the result in the value
// pointed to by vp.
func (p *Parser) Unmarshal(data []byte, vp any, recomposer ...alt.Recomposer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// MustParse a JSON string in to simple types. Panics on error.
func (p *Parser) MustParse(buf []byte, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// Parse a SEN string in to simple types. An error is returned if not valid SEN.
func (p *Parser) Parse(buf []byte, args ...any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Skip BOM if present.

// MustParseReader a JSON io.Reader. Panics on error.
func (p *Parser) MustParseReader(r io.Reader, args ...any) (data any) {
	_ = "STUB: not implemented"
	return *new(any)
}

// ParseReader a SEN io.Reader. An error is returned if not valid SEN.
func (p *Parser) ParseReader(r io.Reader, args ...any) (data any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// Skip BOM if present.

func (p *Parser) parseBuffer(buf []byte, last bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// end of buf reached

// skip and continue

// skip and back to ccomment

// TBD maybe separarte add function or check here for time options

// Only modes with a close array are value, token, and numbers
// which are all over 256 long.

// can not fail appending to an array

// Store additional state (plus) to be used later in addString()
// instead of creating another set of modes for this semi-rare
// case (mongo or javascript only).

// Only modes with a close paren are value, token, and numbers
// which are all over 256 long.

// can not fail appending to a function argument set

// valid finishing maps are one byte longer

// number

// token

// only for non-string
func (p *Parser) add(n any, off int) error { _ = "STUB: not implemented"; return nil }

// object

// array

func (p *Parser) addToken(off int) { _ = "STUB: not implemented"; return }

// object

// Array or just a value

func (p *Parser) addTokenWith(s string, off int) { _ = "STUB: not implemented"; return }

// object

// Array or just a value

func (p *Parser) addString(s string, off int) { _ = "STUB: not implemented"; return }

// object

// TBD if time option for @ and length is over a certain size try as time

// Array or just a value

func (p *Parser) newError(off int, format string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Parser) byteError(off int, mode string, b byte, r rune) error {
	_ = "STUB: not implemented"
	return nil
}

func defaultTokenFunc(args ...any) (result any) { _ = "STUB: not implemented"; return *new(any) }
