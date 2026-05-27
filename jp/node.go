// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp

import (
	"github.com/ohler55/ojg/gen"
)

type index int

// String returns the key as a string.
func (i index) String() string { _ = "STUB: not implemented"; return "" }

// Alter converts the node into it's native type. Note this will modify
// Objects and Arrays in place making them no longer usable as the
// original type. Use with care!
func (i index) Alter() any {
	_ = "STUB: not implemented"

	// Simplify makes a copy of the node but as simple types.
	return *new(any)
}

func (i index) Simplify() any {
	_ = "STUB: not implemented"

	// Dup returns a deep duplicate of the node.
	return *new(any)
}

func (i index) Dup() gen.Node {
	_ = "STUB: not implemented"

	// Empty returns true if the node is empty.
	return *new(gen.Node)
}

func (i index) Empty() bool {
	_ = "STUB: not implemented"

	// GetNodes the elements of the data identified by the path.
	return false
}

func (x Expr) GetNodes(n gen.Node) (results []gen.Node) {
	_ = "STUB: not implemented"

	// A bit of a cheat but to get 100% coverage the index interface
	// functions have to be called. The alternative would be to make th
	// eindex type public but then someone could use it as a value and
	// break the evaluation. Anyway, since evaluating an expty expression
	// is all but useless the index functions are called here and return
	// values ignored since they are never used in the real code.
	return nil
}

// frag index

// must have at least a data element and a fragment index

// last one

// last one

// last one

// last one

// first pass expands, second continues evaluation

// Put prev back and slide fi.

// last one

// Put prev back and slide fi.

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// Free up anything still on the stack.

// FirstNode returns the first matcning node.
func (x Expr) FirstNode(n gen.Node) (result gen.Node) {
	_ = "STUB: not implemented"
	return *new(gen.Node)
}

// frag index

// must have at least a data element and a fragment index

// last one

// last one

// last one

// last one

// first pass expands, second continues evaluation

// Put prev back and slide fi.

// last one

// Put prev back and slide fi.

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one
