// Copyright (c) 2020, Peter Ohler, All rights reserved.

package gen

// Int is a int64 Node.
type Int int64

// String returns a string representation of the Node.
func (n Int) String() string { _ = "STUB: not implemented"; return "" }

// Alter returns the backing int64 value of the Node.
func (n Int) Alter() any {
	_ = "STUB: not implemented"

	// Simplify returns the backing int64 value of the Node.
	return *new(any)
}

func (n Int) Simplify() any {
	_ = "STUB: not implemented"

	// Dup returns the backing int64 value of the Node.
	return *new(any)
}

func (n Int) Dup() Node {
	_ = "STUB: not implemented"

	// Empty returns false.
	return *new(Node)
}

func (n Int) Empty() bool { _ = "STUB: not implemented"; return false }
