// Copyright (c) 2020, Peter Ohler, All rights reserved.

package gen

// String is a string Node.
type String string

// String returns a string representation of the Node.
func (n String) String() string { _ = "STUB: not implemented"; return "" }

// Alter returns the backing float64 value of the Node.
func (n String) Alter() any {
	_ = "STUB: not implemented"

	// Simplify returns the backing float64 value of the Node.
	return *new(any)
}

func (n String) Simplify() any {
	_ = "STUB: not implemented"

	// Dup returns the backing float64 value of the Node.
	return *new(any)
}

func (n String) Dup() Node {
	_ = "STUB: not implemented"

	// Empty returns false if the string has no characters and true otherwise.
	return *new(Node)
}

func (n String) Empty() bool { _ = "STUB: not implemented"; return false }
