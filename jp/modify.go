// Copyright (c) 2022, Peter Ohler, All rights reserved.

package jp

// MustModify modifies matching nodes and panics on an expression error. In
// go, maps can be modified in place as the map itself is modified. Slice
// elements can be replaced in place but elements can not be added or removed
// without potentially needing to replace the original slice with a new
// one. This function and the other jp.Modify functions allow a slice to be
// replaced by stopping at the parent of the target slice and applying a
// modifier function to the target which is then replaced in the
// parent. Without that functionality slice element can only be replaced.
//
// Modified elements replace the original element in the data. The modified
// data is returned. Unless the data is a slice and modified the returned data
// will be the same object as the original. The modifier function will be
// called with the elements that match the path and should return the original
// element or if altered the altered value along with setting the returned
// changed value to true.
func (x Expr) MustModify(data any, modifier func(element any) (altered any, changed bool)) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// MustModifyOne modifies matching nodes and panics on an expression error.
// Modified elements replace the original element in the data. The modified
// data is returned. Unless the data is a slice and modified the returned data
// will be the same object as the original. The modifier function will be
// called with the elements that match the path and should return the original
// element or if altered the altered value along with setting the returned
// changed value to true. The function returns after the first modification.
func (x Expr) MustModifyOne(data any, modifier func(element any) (altered any, changed bool)) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// Modify modifies matching nodes and panics on an expression error.  Modified
// elements replace the original element in the data. The modified data is
// returned. Unless the data is a slice and modified the returned data will be
// the same object as the original. The modifier function will be called with
// the elements that match the path and should return the original element or
// if altered the altered value along with setting the returned changed value
// to true.
func (x Expr) Modify(data any, modifier func(element any) (altered any, changed bool)) (result any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// ModifyOne modifies matching nodes and panics on an expression error.
// Modified elements replace the original element in the data. The modified
// data is returned. Unless the data is a slice and modified the returned data
// will be the same object as the original. The modifier function will be
// called with the elements that match the path and should return the original
// element or if altered the altered value along with setting the returned
// changed value to true. The function returns after the first modification.
func (x Expr) ModifyOne(data any, modifier func(element any) (altered any, changed bool)) (result any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func (x Expr) modify(data any, modifier func(element any) (altered any, changed bool), one bool) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// frag index

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// first pass expands, second continues evaluation

// Put prev back and slide fi.

// Put prev back and slide fi.

// Put prev back and slide fi.

// Put prev back and slide fi.

// Put prev back and slide fi.

// Put prev back and slide fi.

// last one

// last one

func stackAddValue(stack []any, v any) []any { _ = "STUB: not implemented"; return nil }

func descentAddValue(stack []any, v any, fi fragIndex) []any { _ = "STUB: not implemented"; return nil }
