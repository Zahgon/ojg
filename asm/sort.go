// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "sort",
		Eval: sortEval,
		Desc: `Sort the items in an array and return a copy of the array. Valid
types for comparison are strings, numbers, and times. Any other
type returned or a type mismatch will raise an error.`,
	})
}

func sortEval(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// Make a copy so not to change the original.
