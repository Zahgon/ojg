// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "tolower",
		Eval: tolower,
		Desc: `Convert a string to lowercase. There must be exactly one
string argument.`,
	})
}

func tolower(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
