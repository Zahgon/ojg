// Copyright (c) 2020, Peter Ohler, All rights reserved.

package gen

import (
	"math"

	"github.com/ohler55/ojg"
)

const (
	// BigLimit is the limit before a number is converted into a Big
	// instance. (9223372036854775807 / 10 = 922337203685477580)
	BigLimit = math.MaxInt64 / 10

	// DivLimit is the divisor limit before a number is converted into a Big
	// instance. (nearest multiple of 10 below max int64)
	DivLimit = 1000000000000000000
)

// Number is used internally by parsers.
type Number struct {
	I          uint64
	Frac       uint64
	Div        uint64
	Exp        uint64
	Neg        bool
	NegExp     bool
	BigBuf     []byte
	Conv       ojg.NumConvMethod
	ForceFloat bool
}

// Reset the number.
func (n *Number) Reset() { _ = "STUB: not implemented"; return }

// AddDigit to a number.
func (n *Number) AddDigit(b byte) { _ = "STUB: not implemented"; return }

// AddFrac adds a fractional digit.
func (n *Number) AddFrac(b byte) { _ = "STUB: not implemented"; return }

// big

// AddExp adds an exponent digit.
func (n *Number) AddExp(b byte) { _ = "STUB: not implemented"; return }

// big

// FillBig fills the internal buffer with a big number.
func (n *Number) FillBig() { _ = "STUB: not implemented"; return }

// AsNum returns the number as best fit.
func (n *Number) AsNum() (num any) { _ = "STUB: not implemented"; return *new(any) }

// Remove trailing zeros as they can cause precision loss due to
// the way go or the hardware handles multiplication and division.

// A simple division loses precision yet dividing 1.0 by the
// divisor and then multiplying the fraction seems to solve the
// issue on arm64 anyway.

// AsNode returns the number as best fit.
func (n *Number) AsNode() (num Node) { _ = "STUB: not implemented"; return *new(Node) }
