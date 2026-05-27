// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp

type delFlagType struct{}

var delFlag = &delFlagType{}

// MustDel removes matching nodes and panics on error.
func (x Expr) MustDel(data any) { _ = "STUB: not implemented"; return }

// Del removes matching nodes.
func (x Expr) Del(data any) error { _ = "STUB: not implemented"; return nil }

// MustDelOne removes one matching node and panics on error.
func (x Expr) MustDelOne(data any) { _ = "STUB: not implemented"; return }

// DelOne removes at most one node.
func (x Expr) DelOne(data any) error { _ = "STUB: not implemented"; return nil }

// MustSet all matching child node values. If the path to the child does not
// exist array and map elements are added. Panics on error.
func (x Expr) MustSet(data, value any) { _ = "STUB: not implemented"; return }

// Set all matching child node values. An error is returned if it is not
// possible. If the path to the child does not exist array and map elements
// are added.
func (x Expr) Set(data, value any) error { _ = "STUB: not implemented"; return nil }

// SetOne child node value. An error is returned if it is not possible. If the
// path to the child does not exist array and map elements are added.
func (x Expr) SetOne(data, value any) error { _ = "STUB: not implemented"; return nil }

// MustSetOne child node value. If the path to the child does not exist array
// and map elements are added. Panics on error.
func (x Expr) MustSetOne(data, value any) { _ = "STUB: not implemented"; return }

func (x Expr) set(data, value any, fun string, one bool) error {
	_ = "STUB: not implemented"
	return nil
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

// first pass expands, second continues evaluation

// Put prev back and slide fi.

// Put prev back and slide fi.

// Put prev back and slide fi.

// Put prev back and slide fi.

// Put prev back and slide fi.

// Put prev back and slide fi.

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

func reflectSetChild(data any, key string, v any) bool { _ = "STUB: not implemented"; return false }

func reflectSetNth(data any, i int, v any) bool { _ = "STUB: not implemented"; return false }
