// Copyright (c) 2020, Peter Ohler, All rights reserved.

package gen

// Big represents a number too large to be an int64 or a float64.
type Big string

// String representation of the number.
func (n Big) String() string {
	_ = "STUB: not implemented"

	// Alter returns the backing string.
	return ""
}

func (n Big) Alter() any {
	_ = "STUB: not implemented"

	// Simplify the Node into a string.
	return *new(any)
}

func (n Big) Simplify() any {
	_ = "STUB: not implemented"

	// Dup returns itself since it is immutable.
	return *new(any)
}

func (n Big) Dup() Node {
	_ = "STUB: not implemented"

	// Empty returns true if the backing string is empty.
	return *new(Node)
}

func (n Big) Empty() bool { _ = "STUB: not implemented"; return false }
