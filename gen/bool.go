// Copyright (c) 2020, Peter Ohler, All rights reserved.

package gen

// Bool repreents a boolean value.
type Bool bool

// True is a true boolean value.
var True = Bool(true)

// False is a false boolean value.
var False = Bool(false)

// String returns a string representation of the Node.
func (n Bool) String() (s string) { _ = "STUB: not implemented"; return "" }

// Alter returns the backing boolean value of the Node.
func (n Bool) Alter() any {
	_ = "STUB: not implemented"

	// Simplify returns the backing boolean value.
	return *new(any)
}

func (n Bool) Simplify() any {
	_ = "STUB: not implemented"

	// Dup returns itself.
	return *new(any)
}

func (n Bool) Dup() Node {
	_ = "STUB: not implemented"

	// Empty returns false.
	return *new(Node)
}

func (n Bool) Empty() bool { _ = "STUB: not implemented"; return false }
