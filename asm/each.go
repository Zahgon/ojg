// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "each",
		Eval: each,
		Desc: `Each .`,
	})
}

func each(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
