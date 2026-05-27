// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "cond",
		Eval: cond,
		Desc: `A conditional construct modeled after the LISP cond. All
arguments must be array of two elements. The first element must
evaluate to a boolean and the second can be any value. The value
of the first true first argument is returned. If none match nil
is returned.`,
	})
}

func cond(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func evalValue(root map[string]any, at any, value any) (result any) {
	_ = "STUB: not implemented"
	return *new(any)
}
