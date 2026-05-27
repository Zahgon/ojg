// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp

// Expr is a JSON path expression composed of fragments. An Expr implements
// JSONPath as described by https://goessner.net/articles/JsonPath. Where the
// definition is unclear Oj has implemented the description based on the best
// judgement of the author.
type Expr []Frag

// String returns a string representation of the expression.
func (x Expr) String() string { _ = "STUB: not implemented"; return "" }

// BracketString returns a string representation of the expression using the
// bracket notation.
func (x Expr) BracketString() string { _ = "STUB: not implemented"; return "" }

// Append a string representation of the expression to a byte slice and return
// the expanded buffer.
func (x Expr) Append(buf []byte, brackets ...bool) []byte { _ = "STUB: not implemented"; return nil }

// Normal returns true if the only fragments in the expression are root, at,
// child, and nth.
func (x Expr) Normal() bool { _ = "STUB: not implemented"; return false }

// normal

func isNil(v any) bool { _ = "STUB: not implemented"; return false }
