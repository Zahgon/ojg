// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "map?",
		Eval: mapEval,
		Desc: `Returns true if the single required argumement is a map
otherwise false is returned.`,
	})
}

func mapEval(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
