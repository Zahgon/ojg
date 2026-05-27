// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "include",
		Eval: include,
		Desc: `Returns true if a list first argument includes the second
argument. It will also return true if the first argument is a
string and the second string argument is included in the first.`,
	})
}

func include(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
