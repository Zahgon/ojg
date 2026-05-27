// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "replace",
		Eval: replace,
		Desc: `Replace an occurrences the second argument with the third
argument. All three arguments must be strings.`,
	})
}

func replace(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
