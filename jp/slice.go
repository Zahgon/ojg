// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp

// Slice is a slice operation for a JSON path expression.
type Slice []int

// Append a fragment string representation of the fragment to the buffer
// then returning the expanded buffer.
func (f Slice) Append(buf []byte, _, _ bool) []byte { _ = "STUB: not implemented"; return nil }

func (f Slice) remove(value any) (out any, changed bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// Walk in reverse to handle the just-one condition.

func (f Slice) removeOne(value any) (out any, changed bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// Walk in reverse to handle the just-one condition.

// Walk in reverse to handle the just-one condition.

func inStep(i, start, end, step int) bool { _ = "STUB: not implemented"; return false }

func (f Slice) startEndStep(size int) (start, end, step int) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

func (f Slice) locate(pp Expr, data any, rest Expr, max int) (locs []Expr) {
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

// last one

// place holder

// Walk each element in a slice as defined by the Slice fragment.
func (f Slice) Walk(rest, path Expr, nodes []any, cb func(path Expr, nodes []any)) {
	_ = "STUB: not implemented"
	return
}
