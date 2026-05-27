// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp

// Wildcard is used as a flag to indicate the path should be displayed in a
// wildcarded representation.
type Wildcard byte

// Append a fragment string representation of the fragment to the buffer
// then returning the expanded buffer.
func (f Wildcard) Append(buf []byte, bracket, first bool) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (f Wildcard) remove(value any) (out any, changed bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (f Wildcard) removeOne(value any) (out any, changed bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (f Wildcard) locate(pp Expr, data any, rest Expr, max int) (locs []Expr) {
	_ = "STUB: not implemented"
	return nil
}

// last one

// place holder

// last one

// place holder

// last one

// place holder

// last one

// place holder

// last one

// place holder

// last one

// place holder

// no match

// last one

// place holder

// Walk follows the all elements in a map or slice like element.
func (f Wildcard) Walk(rest, path Expr, nodes []any, cb func(path Expr, nodes []any)) {
	_ = "STUB: not implemented"
	return
}

func wildWalk(rest, path Expr, nodes []any, cb func(path Expr, nodes []any), f Frag) {
	_ = "STUB: not implemented"
	return
}

// A bypass to avoid using reflection in the default case.

// Iterate in reverse order as that puts values on the stack in reverse.
