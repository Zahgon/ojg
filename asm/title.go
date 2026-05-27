// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "title",
		Eval: title,
		Desc: `Convert a string to capitalized string. There must be exactly
one string argument.`,
	})
}

func title(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
