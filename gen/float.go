// Copyright (c) 2020, Peter Ohler, All rights reserved.

package gen

// Float is a float64 Node.
type Float float64

// String returns a string representation of the Node.
func (n Float) String() string { _ = "STUB: not implemented"; return "" }

// Alter returns the backing float64 value of the Node.
func (n Float) Alter() any {
	_ = "STUB: not implemented"

	// Simplify returns the backing float64 value of the Node.
	return *new(any)
}

func (n Float) Simplify() any {
	_ = "STUB: not implemented"

	// Dup returns the backing float64 value of the Node.
	return *new(any)
}

func (n Float) Dup() Node {
	_ = "STUB: not implemented"

	// Empty returns false.
	return *new(Node)
}

func (n Float) Empty() bool { _ = "STUB: not implemented"; return false }
