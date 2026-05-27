// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp

// Descent is used as a flag to indicate the path should be displayed in a
// recursive descent representation.
type Descent byte

// Append a fragment string representation of the fragment to the buffer
// then returning the expanded buffer.
func (f Descent) Append(buf []byte, bracket, first bool) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (f Descent) locate(pp Expr, data any, rest Expr, max int) (locs []Expr) {
	_ = "STUB: not implemented"
	// last one
	return nil
}

// place holder

// place holder

// Walk each element in the tree of elements.
func (f Descent) Walk(rest, path Expr, nodes []any, cb func(path Expr, nodes []any)) {
	_ = "STUB: not implemented"
	return
}
