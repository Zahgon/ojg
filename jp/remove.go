// Copyright (c) 2022, Peter Ohler, All rights reserved.

package jp

type remover interface {
	remove(element any) (altered any, changed bool)
}

type oneRemover interface {
	removeOne(element any) (altered any, changed bool)
}

// MustRemove removes matching nodes and panics on an expression error but
// silently makes no changes if there is no match for the expression. Removed
// slice elements are removed and the remaining elements are moveed to fill in
// the removed element. The slice is shortened.
func (x Expr) MustRemove(data any) any { _ = "STUB: not implemented"; return *new(any) }

// MustRemoveOne removes matching nodes and panics on an expression error
// but silently makes no changes if there is no match for the
// expression. Removed slice elements are removed and the remaining elements
// are moveed to fill in the removed element. The slice is shortened.
func (x Expr) MustRemoveOne(data any) any { _ = "STUB: not implemented"; return *new(any) }

// Remove removes matching nodes. An error is returned for an expression error
// but silently makes no changes if there is no match for the
// expression. Removed slice elements are removed and the remaining elements
// are moveed to fill in the removed element. The slice is shortened.
func (x Expr) Remove(data any) (result any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// RemoveOne removes at most one node. An error is returned for an expression
// error but silently makes no changes if there is no match for the
// expression. Removed slice elements are removed and the remaining elements
// are moveed to fill in the removed element. The slice is shortened.
func (x Expr) RemoveOne(data any) (result any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
