// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "int",
		Eval: intEval,
		Desc: `Converts a value into a integer if possible. I no conversion is
possible nil is returned.`,
	})
}

func intEval(root map[string]any, at any, args ...any) (i any) {
	_ = "STUB: not implemented"
	return *new(any)
}
