// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "trim",
		Eval: trim,
		Desc: `Trim white space from both ends of a string unless a second
argument provides an alternative cut set.`,
	})
}

func trim(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
