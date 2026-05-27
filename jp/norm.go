// Copyright (c) 2025, Peter Ohler, All rights reserved.

package jp

type norm byte

// Append a fragment string representation of the fragment to the buffer
// then returning the expanded buffer.
func (f norm) Append(buf []byte, bracket, first bool) []byte { _ = "STUB: not implemented"; return nil }

func (f norm) locate(pp Expr, data any, rest Expr, max int) (locs []Expr) {
	_ = "STUB: not implemented"
	return nil
}

// Walk continues with the next in rest.
func (f norm) Walk(rest, path Expr, nodes []any, cb func(path Expr, nodes []any)) {
	_ = "STUB: not implemented"
	return
}

func normalExpr(x Expr) Expr { _ = "STUB: not implemented"; return *new(Expr) }
