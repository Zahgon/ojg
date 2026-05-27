// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp

// Union is a union operation for a JSON path expression which is a union of a
// Child and Nth fragment.
type Union []any

// Append a fragment string representation of the fragment to the buffer
// then returning the expanded buffer.
func (f Union) Append(buf []byte, _, _ bool) []byte { _ = "STUB: not implemented"; return nil }

// NewUnion creates a new Union with the provide keys.
func NewUnion(keys ...any) (u Union) { _ = "STUB: not implemented"; return *new(Union) }

func (f Union) hasN(n int64) bool { _ = "STUB: not implemented"; return false }

func (f Union) hasKey(key string) bool { _ = "STUB: not implemented"; return false }

func (f Union) removeOne(value any) (out any, changed bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// You would think that ns.SetLen() would work in a case like
// this but it panics as unaddressable so instead the length
// is calculated and then a second pass is made to assign the
// new slice values.

func (f Union) remove(value any) (out any, changed bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// You would think that ns.SetLen() would work in a case like
// this but it panics as unaddressable so instead the length
// is calculated and then a second pass is made to assign the
// new slice values.

func (f Union) locate(pp Expr, data any, rest Expr, max int) (locs []Expr) {
	_ = "STUB: not implemented"
	return nil
}

// last one

// Walk each element in a union.
func (f Union) Walk(rest, path Expr, nodes []any, cb func(path Expr, nodes []any)) {
	_ = "STUB: not implemented"
	return
}
