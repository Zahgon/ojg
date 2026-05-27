// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp

import (
	"regexp"
)

// Equation represents JSON Path script and filter equations. They are used to
// build a script. The purpose of the Equation is to allow scripts or filters
// to be created without using a parser which could return an error if an
// invalid string representation of the script is provided.
type Equation struct {
	o      *op
	result any
	left   *Equation
	right  *Equation
}

// MustParseEquation parses the string argument and returns an Equation or panics.
func MustParseEquation(str string) (eq *Equation) { _ = "STUB: not implemented"; return nil }

// Script creates and returns a Script that implements the equation.
func (e *Equation) Script() *Script { _ = "STUB: not implemented"; return nil }

// Inspect is a debugging function for inspecting an equation tree.
// func (e *Equation) Inspect(b []byte, depth int) []byte {
// 	indent := bytes.Repeat([]byte{' '}, depth)
// 	b = append(b, indent...)
// 	b = append(b, '{')
// 	if e.o == nil {
// 		b = e.appendValue(b, e.result)
// 		b = append(b, '}', '\n')
// 		return b
// 	}
// 	b = append(b, e.o.name...)
// 	b = append(b, '\n')
// 	if e.left == nil {
// 		b = append(b, indent...)
// 		b = append(b, "  nil\n"...)
// 	} else {
// 		b = e.left.Inspect(b, depth+2)
// 	}
// 	if e.right == nil {
// 		b = append(b, indent...)
// 		b = append(b, "  nil\n"...)
// 	} else {
// 		b = e.right.Inspect(b, depth+2)
// 	}
// 	b = append(b, indent...)

// 	return append(b, '}', '\n')
// }

// Filter creates and returns a Script that implements the equation.
func (e *Equation) Filter() (f *Filter) { _ = "STUB: not implemented"; return nil }

// ConstNil creates and returns an Equation for a constant of nil.
func ConstNil() *Equation { _ = "STUB: not implemented"; return nil }

// ConstNothing creates and returns an Equation for a constant of nothing.
func ConstNothing() *Equation { _ = "STUB: not implemented"; return nil }

// ConstBool creates and returns an Equation for a bool constant.
func ConstBool(b bool) *Equation { _ = "STUB: not implemented"; return nil }

// ConstInt creates and returns an Equation for an int64 constant.
func ConstInt(i int64) *Equation { _ = "STUB: not implemented"; return nil }

// ConstFloat creates and returns an Equation for a float64 constant.
func ConstFloat(f float64) *Equation { _ = "STUB: not implemented"; return nil }

// ConstString creates and returns an Equation for a string constant.
func ConstString(s string) *Equation { _ = "STUB: not implemented"; return nil }

// ConstList creates and returns an Equation for an []any constant.
func ConstList(list []any) *Equation { _ = "STUB: not implemented"; return nil }

// ConstRegex creates and returns an Equation for a regex constant.
func ConstRegex(rx *regexp.Regexp) *Equation { _ = "STUB: not implemented"; return nil }

// Get creates and returns an Equation for an expression get of the form
// @.child.
func Get(x Expr) *Equation { _ = "STUB: not implemented"; return nil }

// Eq creates and returns an Equation for an == operator.
func Eq(left, right *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// Neq creates and returns an Equation for a != operator.
func Neq(left, right *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// Lt creates and returns an Equation for a < operator.
func Lt(left, right *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// Gt creates and returns an Equation for a > operator.
func Gt(left, right *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// Lte creates and returns an Equation for a <= operator.
func Lte(left, right *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// Gte creates and returns an Equation for a >= operator.
func Gte(left, right *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// Or creates and returns an Equation for a || operator.
func Or(left, right *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// And creates and returns an Equation for a && operator.
func And(left, right *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// Not creates and returns an Equation for a ! operator.
func Not(arg *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// Add creates and returns an Equation for a + operator.
func Add(left, right *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// Sub creates and returns an Equation for a - operator.
func Sub(left, right *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// Multiply creates and returns an Equation for a * operator.
func Multiply(left, right *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// Divide creates and returns an Equation for a / operator.
func Divide(left, right *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// In creates and returns an Equation for an in operator.
func In(left, right *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// Empty creates and returns an Equation for an empty operator.
func Empty(left, right *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// Has creates and returns an Equation for a has operator.
func Has(left, right *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// Exists creates and returns an Equation for a exists operator.
func Exists(left, right *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// Regex creates and returns an Equation for a regex operator.
func Regex(left, right *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// Length creates and returns an Equation for a length function.
func Length(x Expr) *Equation { _ = "STUB: not implemented"; return nil }

// Count creates and returns an Equation for a count function.
func Count(x Expr) *Equation { _ = "STUB: not implemented"; return nil }

// Match creates and returns an Equation for a match function.
func Match(left, right *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// Search creates and returns an Equation for a search function.
func Search(left, right *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// Append a equation string representation to a buffer.
func (e *Equation) Append(buf []byte, parens bool) []byte { _ = "STUB: not implemented"; return nil }

func (e *Equation) appendValue(buf []byte, v any) []byte { _ = "STUB: not implemented"; return nil }

// String representation of the equation.
func (e *Equation) String() string { _ = "STUB: not implemented"; return "" }

func (e *Equation) buildScript(stack []any) []any { _ = "STUB: not implemented"; return nil }

// should always be an Expr

// Parsing of an equation is from left to right. Each equation is added to the
// equation right side with no regard for precedence. This function then
// reorganizes the equations to be in the correct evaluation order based on
// the precedent.
func precedentCorrect(e *Equation) *Equation { _ = "STUB: not implemented"; return nil }

// a result or empty/nothing

// The left precedence correction is called too many times. Could add a
// flag to Equation indicating it has already been corrected or just
// process more than once for a small performance hit on parsing the
// equation.

func reduceGroups(e *Equation, po *op) *Equation { _ = "STUB: not implemented"; return nil }

// a result or empty/nothing
