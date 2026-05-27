// Copyright (c) 2020, Peter Ohler, All rights reserved.

package gen

// Key use for parsing.
type Key string

// String returns the key as a string.
func (k Key) String() string {
	_ = "STUB: not implemented"

	// Alter converts the node into it's native type. Note this will modify
	// Objects and Arrays in place making them no longer usable as the
	// original type. Use with care!
	return ""
}

func (k Key) Alter() any {
	_ = "STUB: not implemented"

	// Simplify makes a copy of the node but as simple types.
	return *new(any)
}

func (k Key) Simplify() any {
	_ = "STUB: not implemented"

	// Dup returns a deep duplicate of the node.
	return *new(any)
}

func (k Key) Dup() Node {
	_ = "STUB: not implemented"

	// Empty returns true if the node is empty.
	return *new(Node)
}

func (k Key) Empty() bool { _ = "STUB: not implemented"; return false }
