// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "list",
		Eval: list,
		Desc: `Creates a list from all the argument and return that list.`,
	})
}

func list(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
