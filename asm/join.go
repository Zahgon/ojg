// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "join",
		Eval: join,
		Desc: `Join an array of strings with the provided separator. If a
separator is not provided as the second argument then an empty
string is used.`,
	})
}

func join(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
