// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp

// Child is a child operation for a JSON path expression.
type Child string

// Append a fragment string representation of the fragment to the buffer
// then returning the expanded buffer.
func (f Child) Append(buf []byte, bracket, first bool) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (f Child) tokenOk() bool { _ = "STUB: not implemented"; return false }

func (f Child) remove(value any) (out any, changed bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// Can't remove a field from a struct so only a map can be modified.

func (f Child) locate(pp Expr, data any, rest Expr, max int) (locs []Expr) {
	_ = "STUB: not implemented"
	return nil
}

// Walk follows the matching element in a map or map like element.
func (f Child) Walk(rest, path Expr, nodes []any, cb func(path Expr, nodes []any)) {
	_ = "STUB: not implemented"
	return
}
