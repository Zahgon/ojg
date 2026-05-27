// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "append",
		Eval: appendEval,
		Desc: `Appends the second argument to the first argument which must be
an array.`,
	})
}

func appendEval(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
