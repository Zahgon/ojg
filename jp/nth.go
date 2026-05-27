// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp

// Nth is a subscript operator that matches the n-th element in an array for a
// JSON path expression.
type Nth int

// Append a fragment string representation of the fragment to the buffer
// then returning the expanded buffer.
func (f Nth) Append(buf []byte, bracket, first bool) []byte { _ = "STUB: not implemented"; return nil }

func (f Nth) remove(value any) (out any, changed bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (f Nth) locate(pp Expr, data any, rest Expr, max int) (locs []Expr) {
	_ = "STUB: not implemented"
	return nil
}

// Walk follows the matching element in a slice or slice like element.
func (f Nth) Walk(rest, path Expr, nodes []any, cb func(path Expr, nodes []any)) {
	_ = "STUB: not implemented"
	return
}
