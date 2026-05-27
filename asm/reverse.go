// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "reverse",
		Eval: reverse,
		Desc: `Reverse the items in an array and return a copy of it.`,
	})
}

func reverse(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// Make a copy so not to change the original.
