// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "delall",
		Eval: delall,
		Desc: `Deletes the all matching values in either the root ($) or
local (@) data. Exactly one argument is required and it must be
a path. The jp.DelOne() function is used to delete the value.
The local (@) value is returned.`,
	})
}

func delall(root map[string]any, at any, args ...any) (list any) {
	_ = "STUB: not implemented"
	return *new(any)
}
