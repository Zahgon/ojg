// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "size",
		Eval: size,
		Desc: `Returns the size or length of a string, array, or object (map).
For all other types zero is returned`,
	})
}

func size(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
