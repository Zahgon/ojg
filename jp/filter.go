// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp

// Filter is a script used as a filter.
type Filter struct {
	Script
}

// NewFilter creates a new Filter.
func NewFilter(str string) (f *Filter, err error) { _ = "STUB: not implemented"; return nil, nil }

// MustNewFilter creates a new Filter and panics on error.
func MustNewFilter(str string) (f *Filter) { _ = "STUB: not implemented"; return nil }

// String representation of the filter.
func (f *Filter) String() string { _ = "STUB: not implemented"; return "" }

// Append a fragment string representation of the fragment to the buffer
// then returning the expanded buffer.
func (f *Filter) Append(buf []byte, _, _ bool) []byte { _ = "STUB: not implemented"; return nil }

func (f *Filter) remove(value any) (out any, changed bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// You would think that ns.SetLen() would work in a case like
// this but it panics as unaddressable so instead the length
// is calculated and then a second pass is made to assign the
// new slice values.

func (f *Filter) removeOne(value any) (out any, changed bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

// You would think that ns.SetLen() would work in a case like
// this but it panics as unaddressable so instead the length
// is calculated and then a second pass is made to assign the
// new slice values.

func (f *Filter) locate(pp Expr, data any, rest Expr, max int) (locs []Expr) {
	_ = "STUB: not implemented"
	return nil
}

// last one

// place holder

// Walk each element that matches the filter.
func (f *Filter) Walk(rest, path Expr, nodes []any, cb func(path Expr, nodes []any)) {
	_ = "STUB: not implemented"
	return
}
