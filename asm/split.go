// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "split",
		Eval: split,
		Desc: `Split a string on using a specified separator.`,
	})
}

func split(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
