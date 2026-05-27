// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "string?",
		Eval: stringCheck,
		Desc: `Returns true if the single required argumement is a string
otherwise false is returned.`,
	})
	Define(&Fn{
		Name: "string",
		Eval: stringConv,
		Desc: `Converts a value into a string.`,
	})
}

func stringCheck(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func stringConv(root map[string]any, at any, args ...any) (s any) {
	_ = "STUB: not implemented"
	return *new(any)
}
