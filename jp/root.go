// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp

// Root represents a root $ in a JSON path representation.
type Root byte

// Append a fragment string representation of the fragment to the buffer
// then returning the expanded buffer.
func (f Root) Append(buf []byte, bracket, first bool) []byte { _ = "STUB: not implemented"; return nil }

func (f Root) locate(pp Expr, data any, rest Expr, max int) (locs []Expr) {
	_ = "STUB: not implemented"
	return nil
}

// Walk continues with the next in rest.
func (f Root) Walk(rest, path Expr, nodes []any, cb func(path Expr, nodes []any)) {
	_ = "STUB: not implemented"
	return
}
