// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "toupper",
		Eval: toupper,
		Desc: `Convert a string to uppercase. There must be exactly one
string argument.`,
	})
}

func toupper(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
