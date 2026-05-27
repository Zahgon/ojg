// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "nth",
		Eval: nth,
		Desc: `Returns a nth element of an array. The second argument must be
an integer that indicates the element of the array to return.
If the index is less than 0 then the index is from the end of
the array.`,
	})
}

func nth(root map[string]any, at any, args ...any) any { _ = "STUB: not implemented"; return *new(any) }
