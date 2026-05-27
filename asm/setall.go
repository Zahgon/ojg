// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "setall",
		Eval: setall,
		Desc: `Sets multiple values in either the root ($) or local (@) data.
Two arguments are required, the first must be a path and the
second argument is evaluate to a value and inserted using the
jp.Set() function.`,
	})
}

func setall(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
