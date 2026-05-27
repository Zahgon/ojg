// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "inspect",
		Eval: inspect,
		Desc: `Print the arguments as JSON unless the argument is an integer.
Integers are assumed to be the indentation for the arguments
that follow.`,
	})
}

func inspect(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
