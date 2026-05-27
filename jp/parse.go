// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp

import (
	"regexp"
)

const (
	//   0123456789abcdef0123456789abcdef
	tokenMap = "" +
		"................................" + // 0x00
		"...o.o..........oooooooooooo...o" + // 0x20
		".oooooooooooooooooooooooooo...oo" + // 0x40
		".oooooooooooooooooooooooooooooo." + // 0x60
		"oooooooooooooooooooooooooooooooo" + // 0x80
		"oooooooooooooooooooooooooooooooo" + // 0xa0
		"oooooooooooooooooooooooooooooooo" + // 0xc0
		"oooooooooooooooooooooooooooooooo" //   0xe0

	// o for an operatio
	// v for a value start character
	//   0123456789abcdef0123456789abcdef
	eqMap = "" +
		"................................" + // 0x00
		".ov.v.ovv.oo.o.ovvvvvvvvvv..ooo." + // 0x20
		"v..............................." + // 0x40
		".....ov.oo....v.....v.......o.o." + // 0x60
		"................................" + // 0x80
		"................................" + // 0xa0
		"................................" + // 0xc0
		"................................" //   0xe0
)

// Performance is less a concern with Expr parsing as it is usually done just
// once if performance is important. Alternatively, an Expr can be built using
// function calls or bare structs. Parsing is more for convenience. Using this
// approach over modes only adds 10% so a reasonable penalty for
// maintainability.
type parser struct {
	buf []byte
	pos int
}

// ParseString parses a string into an Expr.
func ParseString(s string) (x Expr, err error) {
	_ = "STUB: not implemented"
	return *

	// MustParseString parses a string into an Expr and panics on error.
	new(Expr), nil
}

func MustParseString(s string) (x Expr) { _ = "STUB: not implemented"; return *new(Expr) }

// Parse parses a []byte into an Expr.
func Parse(buf []byte) (x Expr, err error) { _ = "STUB: not implemented"; return *new(Expr), nil }

// MustParse parses a []byte into an Expr and panics on error.
func MustParse(buf []byte) (x Expr) { _ = "STUB: not implemented"; return *new(Expr) }

func (p *parser) readExpr() (x Expr) { _ = "STUB: not implemented"; return *new(Expr) }

func (p *parser) nextFrag(first, lastDescent bool) (f Frag) {
	_ = "STUB: not implemented"
	return *new(Frag)
}

// done

// Any other character is the end of the Expr, figure out later if
// that is an error.

func (p *parser) afterDot() Frag { _ = "STUB: not implemented"; return *new(Frag) }

func (p *parser) afterDotDot() Frag { _ = "STUB: not implemented"; return *new(Frag) }

func (p *parser) afterBracket() Frag { _ = "STUB: not implemented"; return *new(Frag) }

// expect ]

// Kind of ugly but needed to attain full cod coverage as the cover tool
// and the compilier don't know about panics in functions so get the
// return and raise on the same line.

func (p *parser) readInt(b byte) (int, byte) {
	_ = "STUB: not implemented"
	// Allow numbers to begin with a zero.
	/*
		if b == '0' {
			if p.pos < len(p.buf) {
				b = p.buf[p.pos]
				p.pos++
			}
			return 0, b, nil
		}
	*/return 0, 0
}

func (p *parser) readNum(b byte) any { _ = "STUB: not implemented"; return *new(any) }

// Read digits first

func (p *parser) readSlice(i int) Frag { _ = "STUB: not implemented"; return *new(Frag) }

// read the end

func (p *parser) readUnion(v any, b byte) Frag { _ = "STUB: not implemented"; return *new(Frag) }

// next union member

func (p *parser) readHex(b byte) (i byte) { _ = "STUB: not implemented"; return 0 }

func (p *parser) readEscStr(start int, term byte) string { _ = "STUB: not implemented"; return "" }

func (p *parser) readStr(term byte) string { _ = "STUB: not implemented"; return "" }

func (p *parser) readRegex() *regexp.Regexp { _ = "STUB: not implemented"; return nil }

// skip and then continue

func (p *parser) readFilter() *Filter { _ = "STUB: not implemented"; return nil }

func (p *parser) readProc() *Proc { _ = "STUB: not implemented"; return nil }

// Reads an equation by reading the left value first and then seeing if there
// is an operation after that. If so it reads the next equation and decides
// based on precedent which is contained in the other.
func (p *parser) readEq() (eq *Equation) { _ = "STUB: not implemented"; return nil }

// probably reading array elements or function arguments

// reads just lower case alpha characters (a-z)
func (p *parser) readToken() []byte { _ = "STUB: not implemented"; return nil }

func (p *parser) readOpArgs(o *op) (eq *Equation) { _ = "STUB: not implemented"; return nil }

func (p *parser) readEqToken(token []byte) { _ = "STUB: not implemented"; return }

func (p *parser) readEqList() (list []any) { _ = "STUB: not implemented"; return nil }

func partialOp(token []byte, b byte) bool { _ = "STUB: not implemented"; return false }

func (p *parser) readEqOp() (o *op) { _ = "STUB: not implemented"; return nil }

func (p *parser) skipSpace() (b byte) { _ = "STUB: not implemented"; return 0 }

func (p *parser) nextNonSpace() (b byte) { _ = "STUB: not implemented"; return 0 }

func (p *parser) raise(format string, args ...any) { _ = "STUB: not implemented"; return }
