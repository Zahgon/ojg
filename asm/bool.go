// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "bool?",
		Eval: boolEval,
		Desc: `Returns true if the single required argumement is a boolean
otherwise false is returned.`,
	})
}

func boolEval(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
