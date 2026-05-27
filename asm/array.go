// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "array?",
		Eval: arrayEval,
		Desc: `Returns true if the single required argumement is an array
otherwise false is returned.`,
	})
}

func arrayEval(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
