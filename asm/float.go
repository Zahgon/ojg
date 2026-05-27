// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "float",
		Eval: floatEval,
		Desc: `Converts a value into a float if possible. I no conversion is
possible nil is returned.`,
	})
}

func floatEval(root map[string]any, at any, args ...any) (f any) {
	_ = "STUB: not implemented"
	return *new(any)
}
